package da

import (
	"github.com/mdehoog/op-nitro/op-da/flags"
	"github.com/urfave/cli/v2"

	oplog "github.com/ethereum-optimism/optimism/op-service/log"
)

type CLIConfig struct {
	Port      int
	DAURL     string
	LogConfig oplog.CLIConfig
}

func NewConfig(ctx *cli.Context) *CLIConfig {
	return &CLIConfig{
		Port:      ctx.Int(flags.PortFlag.Name),
		DAURL:     ctx.String(flags.DAURLFlag.Name),
		LogConfig: oplog.ReadCLIConfig(ctx),
	}
}
