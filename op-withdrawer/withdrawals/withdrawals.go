package withdrawals

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/mdehoog/op-enclave/bindings"
)

type ProofClient interface {
	GetProof(context.Context, common.Address, []string, *big.Int) (*gethclient.AccountResult, error)
}

type EthClient interface {
	TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error)
	BlockByNumber(context.Context, *big.Int) (*types.Block, error)
}

type OutputOracle interface {
	LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error)
	GetL2OutputIndexAfter(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error)
}

type Portal interface {
	ProveAndFinalizeWithdrawalTransaction(opts *bind.TransactOpts, _tx bindings.TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof bindings.TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error)
}

func WaitForOutputBlock(ctx context.Context, outputOracle *bindings.OutputOracle, blockNumber *big.Int, pollInterval time.Duration) (*big.Int, error) {
	for {
		l2OutputBlock, err := outputOracle.LatestBlockNumber(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}
		if l2OutputBlock.Cmp(blockNumber) >= 0 {
			return l2OutputBlock, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(pollInterval):
		}
	}
}

func ProveAndFinalizeWithdrawal(ctx context.Context, l2ProofCl ProofClient, l2Client EthClient, opts *bind.TransactOpts, outputOracle OutputOracle, portal Portal, withdrawalTxHash common.Hash, l2OutputBlock *big.Int) (*types.Transaction, error) {
	l2OutputIndex, err := outputOracle.GetL2OutputIndexAfter(&bind.CallOpts{}, l2OutputBlock)
	if err != nil {
		log.Fatalf("Error getting L2 output index: %v", err)
	}

	withdrawal, err := withdrawals.ProveWithdrawalParametersForBlock(ctx, l2ProofCl, l2Client, l2Client, withdrawalTxHash, l2OutputBlock, l2OutputIndex)
	if err != nil {
		log.Fatalf("Error proving withdrawal parameters: %v", err)
	}

	outputRootProof := bindings.TypesOutputRootProof{
		Version:                  withdrawal.OutputRootProof.Version,
		StateRoot:                withdrawal.OutputRootProof.StateRoot,
		MessagePasserStorageRoot: withdrawal.OutputRootProof.MessagePasserStorageRoot,
		LatestBlockhash:          withdrawal.OutputRootProof.LatestBlockhash,
	}

	return portal.ProveAndFinalizeWithdrawalTransaction(
		opts,
		bindings.TypesWithdrawalTransaction{
			Nonce:    withdrawal.Nonce,
			Sender:   withdrawal.Sender,
			Target:   withdrawal.Target,
			Value:    withdrawal.Value,
			GasLimit: withdrawal.GasLimit,
			Data:     withdrawal.Data,
		},
		withdrawal.L2OutputIndex,
		outputRootProof,
		withdrawal.WithdrawalProof,
	)
}

func WaitForReceipt(ctx context.Context, client EthClient, txHash common.Hash, pollInterval time.Duration) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if errors.Is(err, ethereum.NotFound) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(pollInterval):
			}
			continue
		}
		if err != nil {
			return nil, err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return nil, errors.New("unsuccessful receipt status")
		}
		return receipt, nil
	}
}
