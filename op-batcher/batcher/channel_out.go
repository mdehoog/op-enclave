package batcher

import (
	"errors"

	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/core/types"
)

var ErrWithdrawalDetected = errors.New("withdrawal detected")

func NewChannelOut(cfg batcher.ChannelConfig, rollupCfg *rollup.Config) (derive.ChannelOut, error) {
	co, err := batcher.NewChannelOut(cfg, rollupCfg)
	if err != nil {
		return nil, err
	}
	return &channelOut{
		ChannelOut: co,
	}, nil
}

type channelOut struct {
	derive.ChannelOut
	fullErr error
}

func (c *channelOut) AddBlock(config *rollup.Config, block *types.Block) (*derive.L1BlockInfo, error) {
	if block.Bloom().Test(predeploys.L2ToL1MessagePasserAddr.Bytes()) {
		c.fullErr = ErrWithdrawalDetected
	}
	return c.ChannelOut.AddBlock(config, block)
}

func (c *channelOut) FullErr() error {
	if c.fullErr != nil {
		return c.fullErr
	}
	return c.ChannelOut.FullErr()
}
