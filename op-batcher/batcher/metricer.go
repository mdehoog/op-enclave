package batcher

import (
	"github.com/ethereum-optimism/optimism/op-batcher/metrics"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

type Metricer interface {
	metrics.Metricer
	RegisterChannel(out ChannelOut)
}

type metricer struct {
	metrics.Metricer
	log            log.Logger
	proposerClient *rpc.Client
	channels       map[derive.ChannelID]ChannelOut
}

func NewMetricer(m metrics.Metricer, log log.Logger, proposerClient *rpc.Client) Metricer {
	return &metricer{
		Metricer:       m,
		log:            log,
		proposerClient: proposerClient,
		channels:       make(map[derive.ChannelID]ChannelOut),
	}
}

func (m *metricer) RegisterChannel(out ChannelOut) {
	m.channels[out.ID()] = out
}

func (m *metricer) RecordChannelFullySubmitted(id derive.ChannelID) {
	m.Metricer.RecordChannelFullySubmitted(id)

	channel, ok := m.channels[id]
	if !ok {
		return
	}
	delete(m.channels, id)
	var numbers []uint64
	for _, b := range channel.Blocks() {
		numbers = append(numbers, b.NumberU64())
	}
	if err := m.proposerClient.Call(nil, "admin_blocksBatched", numbers); err != nil {
		m.log.Error("failed to notify Proposer of batched blocks", "err", err)
	}
}

func (m *metricer) RecordChannelTimedOut(id derive.ChannelID) {
	m.Metricer.RecordChannelTimedOut(id)
	delete(m.channels, id)
}
