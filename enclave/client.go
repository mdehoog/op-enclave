package enclave

import (
	"context"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	rpc.Client
}

var _ RPC = (*Client)(nil)

func (c *Client) callContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return c.CallContext(ctx, result, Namespace+"_"+method, args...)
}

func (c *Client) SignerPublicKey(ctx context.Context) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	return result, c.callContext(ctx, &result, "signerPublicKey")
}

func (c *Client) SignerAttestation(ctx context.Context) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	return result, c.callContext(ctx, &result, "signerAttestation")
}

func (c *Client) DecryptionPublicKey(ctx context.Context) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	return result, c.callContext(ctx, &result, "decryptionPublicKey")
}

func (c *Client) DecryptionAttestation(ctx context.Context) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	return result, c.callContext(ctx, &result, "decryptionAttestation")
}

func (c *Client) EncryptedSignerKey(ctx context.Context, attestation hexutil.Bytes) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	return result, c.callContext(ctx, &result, "encryptedSignerKey", attestation)
}

func (c *Client) SetSignerKey(ctx context.Context, encrypted hexutil.Bytes) error {
	return c.callContext(ctx, nil, "setSignerKey", encrypted)
}

func (c *Client) ExecuteStateless(ctx context.Context, config *RollupConfig, l1Origin *types.Header, l1Receipts types.Receipts, previousBlockTxs []*types.Transaction, block *Block, witness hexutil.Bytes, messageAccount *eth.AccountResult, prevMessageAccountHash common.Hash) (*Proposal, error) {
	var result Proposal
	return &result, c.callContext(ctx, &result, "executeStateless", config, l1Origin, l1Receipts, previousBlockTxs, block, witness, messageAccount, prevMessageAccountHash)
}

func (c *Client) Aggregate(ctx context.Context, configHash common.Hash, prevOutputRoot common.Hash, proposals []*Proposal) (*Proposal, error) {
	var result Proposal
	return &result, c.callContext(ctx, &result, "aggregate", configHash, prevOutputRoot, proposals)
}
