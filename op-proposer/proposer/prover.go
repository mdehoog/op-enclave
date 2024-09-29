package proposer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hashicorp/go-multierror"
	"github.com/mdehoog/op-nitro/enclave"
)

type prover struct {
	config     *enclave.PerChainConfig
	configHash common.Hash
	l1         L1Client
	l2         L2Client
	rollup     RollupClient
	enclave    enclave.RPC
}

type proposal struct {
	output   *enclave.Proposal
	blockRef eth.L2BlockRef
}

func newProver(
	ctx context.Context,
	l1 L1Client,
	l2 L2Client,
	rollup RollupClient,
	enclav enclave.RPC,
) (*prover, error) {
	rollupConfig, err := rollup.RollupConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rollup config: %w", err)
	}
	cfg := enclave.FromRollupConfig(rollupConfig)
	configHash, err := cfg.Hash()
	if err != nil {
		return nil, fmt.Errorf("failed to hash rollup config: %w", err)
	}

	return &prover{
		config:     cfg,
		configHash: configHash,
		l1:         l1,
		l2:         l2,
		rollup:     rollup,
		enclave:    enclav,
	}, nil
}

func (o *prover) Generate(ctx context.Context, blockNumber uint64) (*proposal, bool, error) {
	blockCh := await(func() (*types.Block, error) {
		return o.l2.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	}, func(err error) error {
		return fmt.Errorf("failed to fetch L2 block: %w", err)
	})

	blockResult := <-blockCh
	if blockResult.err != nil {
		return nil, false, blockResult.err
	}
	block := blockResult.value

	witnessCh := await(func() ([]byte, error) {
		return o.rollup.WitnessAtBlock(ctx, block.Hash())
	}, func(err error) error {
		return fmt.Errorf("failed to fetch witness: %w", err)
	})

	messageAccountCh := await(func() (*eth.AccountResult, error) {
		return o.l2.GetProof(ctx, predeploys.L2ToL1MessagePasserAddr, block.Hash())
	}, func(err error) error {
		return fmt.Errorf("failed to fetch message account proof: %w", err)
	})

	previousBlockCh := await(func() (*types.Block, error) {
		return o.l2.BlockByHash(ctx, block.ParentHash())
	}, func(err error) error {
		return fmt.Errorf("failed to fetch previous L2 block: %w", err)
	})

	prevMessageAccountCh := await(func() (*eth.AccountResult, error) {
		return o.l2.GetProof(ctx, predeploys.L2ToL1MessagePasserAddr, block.ParentHash())
	}, func(err error) error {
		return fmt.Errorf("failed to fetch previous message account proof: %w", err)
	})

	blockRef, err := derive.L2BlockToBlockRef(o.config.ToRollupConfig(), block)
	if err != nil {
		return nil, false, fmt.Errorf("failed to derive block ref from L2 block: %w", err)
	}

	l1OriginCh := await(func() (*types.Header, error) {
		return o.l1.HeaderByHash(ctx, blockRef.L1Origin.Hash)
	}, func(err error) error {
		return fmt.Errorf("failed to fetch L1 origin header: %w", err)
	})

	l1ReceiptsCh := await(func() (types.Receipts, error) {
		return o.l1.BlockReceipts(ctx, blockRef.L1Origin.Hash)
	}, func(err error) error {
		return fmt.Errorf("failed to fetch L1 receipts: %w", err)
	})

	var errors []error

	witness := <-witnessCh
	errors = appendNonNil(errors, witness.err)

	messageAccount := <-messageAccountCh
	errors = appendNonNil(errors, messageAccount.err)

	previousBlock := <-previousBlockCh
	errors = appendNonNil(errors, previousBlock.err)

	l1Origin := <-l1OriginCh
	errors = appendNonNil(errors, l1Origin.err)

	l1Receipts := <-l1ReceiptsCh
	errors = appendNonNil(errors, l1Receipts.err)

	prevMessageAccount := <-prevMessageAccountCh
	errors = appendNonNil(errors, prevMessageAccount.err)

	if len(errors) > 0 {
		return nil, false, &multierror.Error{Errors: errors}
	}

	output, err := o.enclave.ExecuteStateless(
		ctx,
		o.config,
		l1Origin.value,
		l1Receipts.value,
		previousBlock.value.Transactions(),
		block.Header(),
		block.Transactions(),
		witness.value,
		messageAccount.value,
		prevMessageAccount.value.StorageHash,
	)
	if err != nil {
		return nil, false, fmt.Errorf("failed to execute enclave state transition: %w", err)
	}
	anyWithdrawals := block.Bloom().Test(predeploys.L2ToL1MessagePasserAddr.Bytes())
	return &proposal{
		output:   output,
		blockRef: blockRef,
	}, anyWithdrawals, nil
}

func (o *prover) Aggregate(ctx context.Context, prevOutputRoot common.Hash, proposals []*proposal) (*proposal, error) {
	prop := make([]*enclave.Proposal, len(proposals))
	for i, p := range proposals {
		prop[i] = p.output
	}
	output, err := o.enclave.Aggregate(ctx, o.configHash, prevOutputRoot, prop)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate proposals: %w", err)
	}
	return &proposal{
		output:   output,
		blockRef: proposals[len(proposals)-1].blockRef,
	}, nil
}

type result[E any] struct {
	value E
	err   error
}

func await[E any](f func() (E, error), w func(err error) error) chan result[E] {
	ch := make(chan result[E], 1)
	go func() {
		value, err := f()
		if err != nil {
			err = w(err)
		}
		ch <- result[E]{value, err}
	}()
	return ch
}

func appendNonNil(r []error, e error) []error {
	if e != nil {
		r = append(r, e)
	}
	return r
}
