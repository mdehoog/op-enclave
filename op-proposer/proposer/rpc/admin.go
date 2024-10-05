package rpc

import (
	"context"

	"github.com/ethereum/go-ethereum/log"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/rpc"
)

type ProposerDriver interface {
	BlocksBatched(numbers []uint64) error
}

type adminAPI struct {
	*rpc.CommonAdminAPI
	b ProposerDriver
}

func NewAdminAPI(dr ProposerDriver, m metrics.RPCMetricer, log log.Logger) gethrpc.API {
	return gethrpc.API{
		Namespace: "admin",
		Service: &adminAPI{
			CommonAdminAPI: rpc.NewCommonAdminAPI(m, log),
			b:              dr,
		},
	}
}

func (a *adminAPI) BlocksBatched(_ context.Context, numbers []uint64) error {
	return a.b.BlocksBatched(numbers)
}
