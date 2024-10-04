package proposer

import (
	"github.com/ethereum-optimism/optimism/op-proposer/proposer"
	"github.com/mdehoog/op-enclave/op-proposer/flags"
	"github.com/urfave/cli/v2"
)

type CLIConfig struct {
	*proposer.CLIConfig
	L2EthRpc            string
	EnclaveRpc          string
	MinProposalInterval uint64
}

func NewConfig(ctx *cli.Context) *CLIConfig {
	return &CLIConfig{
		CLIConfig:           proposer.NewConfig(ctx),
		L2EthRpc:            ctx.String(flags.L2EthRpcFlag.Name),
		EnclaveRpc:          ctx.String(flags.EnclaveRpcFlag.Name),
		MinProposalInterval: ctx.Uint64(flags.MinProposalIntervalFlag.Name),
	}
}
