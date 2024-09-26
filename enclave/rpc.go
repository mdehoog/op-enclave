package enclave

import (
	"context"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

const Namespace = "enclave"

type RPC interface {
	SignerPublicKey(ctx context.Context) (hexutil.Bytes, error)
	SignerAttestation(ctx context.Context) (hexutil.Bytes, error)
	DecryptionPublicKey(ctx context.Context) (hexutil.Bytes, error)
	DecryptionAttestation(ctx context.Context) (hexutil.Bytes, error)
	EncryptedSignerKey(ctx context.Context, attestation hexutil.Bytes) (hexutil.Bytes, error)
	SetSignerKey(ctx context.Context, encrypted hexutil.Bytes) error
	ExecuteStateless(
		ctx context.Context,
		config *RollupConfig,
		l1Origin *types.Header,
		l1Receipts types.Receipts,
		previousBlockTxs types.Transactions,
		blockHeader *types.Header,
		blockTxs types.Transactions,
		witness hexutil.Bytes,
		messageAccount *eth.AccountResult,
		prevMessageAccountHash common.Hash,
	) (*Proposal, error)
	Aggregate(ctx context.Context, configHash common.Hash, prevOutputRoot common.Hash, proposals []*Proposal) (*Proposal, error)
}
