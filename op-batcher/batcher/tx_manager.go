package batcher

import (
	"context"
	"errors"
	"math/big"
	"sync/atomic"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

func NewTxManager(log log.Logger) (txmgr.TxManager, error) {
	return &txManager{
		log: log,
	}, nil
}

type txManager struct {
	log    log.Logger
	closed atomic.Bool
}

func (t *txManager) Send(ctx context.Context, candidate txmgr.TxCandidate) (*types.Receipt, error) {
	ch := make(chan txmgr.SendResponse)
	t.SendAsync(ctx, candidate, ch)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case resp := <-ch:
		return resp.Receipt, resp.Err
	}
}

func (t *txManager) SendAsync(ctx context.Context, candidate txmgr.TxCandidate, ch chan txmgr.SendResponse) {
	//TODO implement S3 write
	t.log.Info("SendAsync tx", "len", len(candidate.TxData))
	ch <- txmgr.SendResponse{
		Receipt: &types.Receipt{},
	}
}

func (t *txManager) From() common.Address {
	return common.Address{}
}

func (t *txManager) BlockNumber(context.Context) (uint64, error) {
	return 0, errors.New("not implemented")
}

func (t *txManager) API() rpc.API {
	return rpc.API{}
}

func (t *txManager) Close() {
	// TODO
	t.closed.Store(true)
}

func (t *txManager) IsClosed() bool {
	return t.closed.Load()
}

func (t *txManager) SuggestGasPriceCaps(context.Context) (tipCap *big.Int, baseFee *big.Int, blobBaseFee *big.Int, err error) {
	return nil, nil, nil, errors.New("not implemented")
}
