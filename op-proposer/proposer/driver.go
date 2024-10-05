package proposer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-proposer/metrics"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mdehoog/op-enclave/bindings"
	"github.com/mdehoog/op-enclave/op-enclave/enclave"
)

var (
	ErrProposerNotRunning = errors.New("proposer is not running")
)

type OOContract interface {
	Version(*bind.CallOpts) (string, error)
	LatestL2Output(opts *bind.CallOpts) (bindings.TypesOutputProposal, error)
}

type DriverSetup struct {
	Log           log.Logger
	Metr          metrics.Metricer
	Cfg           ProposerConfig
	Txmgr         txmgr.TxManager
	L1Client      L1Client
	L2Client      L2Client
	RollupClient  RollupClient
	EnclaveClient enclave.RPC
}

// L2OutputSubmitter is responsible for proposing outputs
type L2OutputSubmitter struct {
	DriverSetup

	wg   sync.WaitGroup
	done chan struct{}

	ctx    context.Context
	cancel context.CancelFunc

	mutex   sync.Mutex
	running bool

	ooContract OOContract
	ooABI      *abi.ABI

	prover *Prover

	blocksBatched      map[uint64]struct{}
	blocksBatchedMutex sync.Mutex
}

// NewL2OutputSubmitter creates a new L2 Output Submitter
func NewL2OutputSubmitter(setup DriverSetup) (_ *L2OutputSubmitter, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	// The above context is long-lived, and passed to the `L2OutputSubmitter` instance. This context is closed by
	// `StopL2OutputSubmitting`, but if this function returns an error or panics, we want to ensure that the context
	// doesn't leak.
	defer func() {
		if err != nil || recover() != nil {
			cancel()
		}
	}()

	return newL2OOSubmitter(ctx, cancel, setup)
}

func newL2OOSubmitter(ctx context.Context, cancel context.CancelFunc, setup DriverSetup) (*L2OutputSubmitter, error) {
	ooContract, err := bindings.NewOutputOracleCaller(*setup.Cfg.L2OutputOracleAddr, setup.L1Client)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create L2OO at address %s: %w", setup.Cfg.L2OutputOracleAddr, err)
	}

	cCtx, cCancel := context.WithTimeout(ctx, setup.Cfg.NetworkTimeout)
	defer cCancel()
	version, err := ooContract.Version(&bind.CallOpts{Context: cCtx})
	if err != nil {
		cancel()
		return nil, err
	}
	log.Info("Connected to L2OutputOracle", "address", setup.Cfg.L2OutputOracleAddr, "version", version)

	parsed, err := bindings.OutputOracleMetaData.GetAbi()
	if err != nil {
		cancel()
		return nil, err
	}

	prover, err := NewProver(cCtx, setup.L1Client, setup.L2Client, setup.RollupClient, setup.EnclaveClient)
	if err != nil {
		cancel()
		return nil, err
	}

	return &L2OutputSubmitter{
		DriverSetup: setup,
		done:        make(chan struct{}),
		ctx:         ctx,
		cancel:      cancel,

		ooContract: ooContract,
		ooABI:      parsed,
		prover:     prover,

		blocksBatched: make(map[uint64]struct{}),
	}, nil
}

func (l *L2OutputSubmitter) StartL2OutputSubmitting() error {
	l.Log.Info("Starting Proposer")

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.running {
		return errors.New("proposer is already running")
	}
	l.running = true

	l.wg.Add(1)
	go l.loop()

	l.Log.Info("Proposer started")
	return nil
}

func (l *L2OutputSubmitter) StopL2OutputSubmittingIfRunning() error {
	err := l.StopL2OutputSubmitting()
	if errors.Is(err, ErrProposerNotRunning) {
		return nil
	}
	return err
}

func (l *L2OutputSubmitter) StopL2OutputSubmitting() error {
	l.Log.Info("Stopping Proposer")

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if !l.running {
		return ErrProposerNotRunning
	}
	l.running = false

	l.cancel()
	close(l.done)
	l.wg.Wait()

	l.Log.Info("Proposer stopped")
	return nil
}

func (l *L2OutputSubmitter) BlocksBatched(numbers []uint64) error {
	l.blocksBatchedMutex.Lock()
	defer l.blocksBatchedMutex.Unlock()
	for _, number := range numbers {
		l.blocksBatched[number] = struct{}{}
	}
	return nil
}

func (l *L2OutputSubmitter) LatestBlockBatched(ctx context.Context) (uint64, error) {
	syncStatus, err := l.RollupClient.SyncStatus(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get sync status from Rollup: %w", err)
	}
	batched := syncStatus.FinalizedL2.Number
	if l.Cfg.AllowNonFinalized {
		batched = syncStatus.PendingSafeL2.Number

		l.blocksBatchedMutex.Lock()
		defer l.blocksBatchedMutex.Unlock()

		for number := range l.blocksBatched {
			if number <= batched {
				delete(l.blocksBatched, number)
			}
		}

		// iterate through the batched blocks to find the last contiguous batched block number
		for i := batched + 1; ; i++ {
			if _, ok := l.blocksBatched[i]; !ok {
				return i - 1, nil
			}
		}
	}
	return batched, nil
}

// loop is responsible for creating & submitting the next outputs
// The loop regularly polls the L2 chain to infer whether to make the next proposal.
func (l *L2OutputSubmitter) loop() {
	defer l.wg.Done()
	defer l.Log.Info("loop returning")
	ctx := l.ctx
	ticker := time.NewTicker(l.Cfg.PollInterval)
	defer ticker.Stop()
	var lastProposal *Proposal
	for {
		select {
		case <-ticker.C:
			// prioritize quit signal
			select {
			case <-l.done:
				return
			default:
			}

			// A note on retrying: the outer ticker already runs on a short
			// poll interval, which has a default value of 6 seconds. So no
			// retry logic is needed around output fetching here.
			proposal, shouldPropose, err := l.generateNextProposal(ctx, lastProposal)
			lastProposal = proposal
			if err != nil {
				l.Log.Warn("Error getting output", "err", err)
				continue
			} else if !shouldPropose {
				// debug logging already in Fetch(DGF|L2OO)Output
				continue
			}

			l.proposeOutput(ctx, proposal)
		case <-l.done:
			return
		}
	}
}

func (l *L2OutputSubmitter) generateNextProposal(ctx context.Context, lastProposal *Proposal) (*Proposal, bool, error) {
	proposed, err := l.ooContract.LatestL2Output(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, false, fmt.Errorf("failed to get latest proposed block number from Oracle: %w", err)
	}
	proposedBlockNumber := proposed.L2BlockNumber.Uint64()
	lastProposalBlockNumber := proposedBlockNumber

	// purge reorged blocks
	if lastProposal != nil {
		lastProposalBlockRef := lastProposal.BlockRef
		proposedHeader, err := l.L2Client.HeaderByNumber(ctx, new(big.Int).SetUint64(lastProposalBlockRef.Number))
		if err != nil {
			return nil, false, fmt.Errorf("failed to get header for block %d: %w", lastProposalBlockRef.Number, err)
		}
		if lastProposalBlockRef.Hash != proposedHeader.Hash() {
			l.Log.Warn("Last proposal block hash does not match the L2 block hash, possible reorg",
				"last_proposal", lastProposalBlockRef.Hash, "l2_block", proposedHeader.Hash())
			// TODO rather than clearing all aggregated proposals, store snapshots and binary search back to the common ancestor
			lastProposal = nil
		} else {
			lastProposalBlockNumber = lastProposalBlockRef.Number
		}
	}

	// generate new proposals up to the latest block
	batchedBlockNumber, err := l.LatestBlockBatched(ctx)
	if err != nil {
		return nil, false, err
	}

	// TODO implement proposal array limit (aggregate in chunks)
	// TODO implement a pool of go-routines for parallel proof generation
	// TODO generate proofs for unsafe blocks ahead of time
	var proposals []*Proposal
	if lastProposal != nil {
		proposals = append(proposals, lastProposal)
	}
	shouldPropose := lastProposalBlockNumber < batchedBlockNumber &&
		l.Cfg.MinProposalInterval > 0 && batchedBlockNumber-proposedBlockNumber > l.Cfg.MinProposalInterval
	for i := lastProposalBlockNumber + 1; i <= batchedBlockNumber; i++ {
		proposal, anyWithdrawals, err := l.prover.Generate(ctx, i)
		if err != nil {
			return nil, false, fmt.Errorf("failed to generate proof for block %d: %w", i, err)
		}
		proposals = append(proposals, proposal)
		shouldPropose = shouldPropose || anyWithdrawals
		l.Log.Info("Generated proof for block", "block", i, "batched", batchedBlockNumber, "shouldPropose", shouldPropose, "output", proposal.Output.OutputRoot.String())
	}

	if len(proposals) == 0 {
		return nil, false, nil
	}

	if len(proposals) > 1 {
		log.Info("Aggregating proofs", "proposals", len(proposals))
		lastProposal, err = l.prover.Aggregate(ctx, proposed.OutputRoot, proposals)
		if err != nil {
			return nil, false, fmt.Errorf("failed to aggregate proofs: %w", err)
		}
	} else {
		lastProposal = proposals[0]
	}

	if shouldPropose {
		latestL1BlockHeader, err := l.L1Client.HeaderByNumber(l.ctx, nil)
		if err != nil {
			log.Warn("Failed to get latest block header", "err", err)
			shouldPropose = false
		} else if lastProposal.BlockRef.L1Origin.Number <= latestL1BlockHeader.Number.Uint64()-256 {
			// only submit onchain if within the blockhash window
			log.Warn("Not submitting proposal, block is too old", "l1Origin", lastProposal.BlockRef.L1Origin.Number, "l1Latest", latestL1BlockHeader.Number.Uint64())
			shouldPropose = false
		}
	}

	return lastProposal, shouldPropose, nil
}

func (l *L2OutputSubmitter) proposeOutput(ctx context.Context, proposal *Proposal) {
	cCtx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	if err := l.sendTransaction(cCtx, proposal); err != nil {
		l.Log.Error("Failed to send proposal transaction",
			"err", err,
			"block", proposal.BlockRef)
		return
	}
	l.Metr.RecordL2BlocksProposed(proposal.BlockRef)
}

// sendTransaction creates & sends transactions through the underlying transaction manager.
func (l *L2OutputSubmitter) sendTransaction(ctx context.Context, proposal *Proposal) error {
	l.Log.Info("Proposing output root", "output", proposal.Output.OutputRoot, "block", proposal.BlockRef)
	data, err := l.ProposeL2OutputTxData(proposal)
	if err != nil {
		return err
	}
	receipt, err := l.Txmgr.Send(ctx, txmgr.TxCandidate{
		TxData:   data,
		To:       l.Cfg.L2OutputOracleAddr,
		GasLimit: 0,
	})
	if err != nil {
		return err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		l.Log.Error("Proposer tx successfully published but reverted", "tx_hash", receipt.TxHash)
	} else {
		l.Log.Info("Proposer tx successfully published", "tx_hash", receipt.TxHash)
	}
	return nil
}

// ProposeL2OutputTxData creates the transaction data for the ProposeL2Output function
func (l *L2OutputSubmitter) ProposeL2OutputTxData(proposal *Proposal) ([]byte, error) {
	return proposeL2OutputTxData(l.ooABI, proposal)
}

// proposeL2OutputTxData creates the transaction data for the ProposeL2Output function
func proposeL2OutputTxData(abi *abi.ABI, proposal *Proposal) ([]byte, error) {
	sig := make([]byte, len(proposal.Output.Signature))
	copy(sig, proposal.Output.Signature)
	sig[64] += 27
	return abi.Pack(
		"proposeL2Output",
		proposal.Output.OutputRoot,
		new(big.Int).SetUint64(proposal.BlockRef.Number),
		new(big.Int).SetUint64(proposal.BlockRef.L1Origin.Number),
		sig,
	)
}
