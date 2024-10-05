package flags

import (
	"github.com/ethereum-optimism/optimism/op-batcher/flags"
	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/urfave/cli/v2"
)

func prefixEnvVar(name string) []string {
	return opservice.PrefixEnvVar(flags.EnvVarPrefix, name)
}

var (
	ProposerRpcFlag = &cli.StringFlag{
		Name:     "proposer-rpc",
		Usage:    "HTTP provider URL for the Proposer",
		EnvVars:  prefixEnvVar("PROPOSER_RPC"),
		Required: true,
	}
)

var requiredFlags = []cli.Flag{
	ProposerRpcFlag,
}

func init() {
	Flags = append(requiredFlags, flags.Flags...)
}

var Flags []cli.Flag
