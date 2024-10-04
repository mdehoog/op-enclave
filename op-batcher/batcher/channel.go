package batcher

import (
	"errors"

	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-batcher/metrics"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var ErrWithdrawalDetected = errors.New("withdrawal detected")

func NewChannel(log log.Logger, metr metrics.Metricer, cfg batcher.ChannelConfig, rollupCfg *rollup.Config, latestL1OriginBlockNum uint64) (batcher.Channel, error) {
	co, err := batcher.NewChannel(log, metr, cfg, rollupCfg, latestL1OriginBlockNum)
	if err != nil {
		return nil, err
	}
	return &channel{
		Channel: co,
	}, nil
}

type channel struct {
	batcher.Channel
	fullErr error
}

func (c *channel) AddBlock(block *types.Block) (*derive.L1BlockInfo, error) {
	if block.Bloom().Test(predeploys.L2ToL1MessagePasserAddr.Bytes()) {
		c.fullErr = ErrWithdrawalDetected
	}
	return c.Channel.AddBlock(block)
}

func (c *channel) FullErr() error {
	if c.fullErr != nil {
		return c.fullErr
	}
	return c.Channel.FullErr()
}
