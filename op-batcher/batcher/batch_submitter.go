package batcher

import (
	"context"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-batcher/flags"
	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	thisflags "github.com/mdehoog/op-nitro/op-batcher/flags"
	"github.com/urfave/cli/v2"
)

// Main is the entrypoint into the Batch Submitter.
// This method returns a cliapp.LifecycleAction, to create an op-service CLI-lifecycle-managed batch-submitter with.
func Main(version string) cliapp.LifecycleAction {
	return func(cliCtx *cli.Context, closeApp context.CancelCauseFunc) (cliapp.Lifecycle, error) {
		if err := flags.CheckRequired(cliCtx); err != nil {
			return nil, err
		}
		cfg := batcher.NewConfig(cliCtx)
		if err := cfg.Check(); err != nil {
			return nil, fmt.Errorf("invalid CLI flags: %w", err)
		}

		l := oplog.NewLogger(oplog.AppOut(cliCtx), cfg.LogConfig)
		oplog.SetGlobalLogHandler(l.Handler())
		opservice.ValidateEnvVars(flags.EnvVarPrefix, thisflags.Flags, l)

		l.Info("Initializing Batch Submitter")
		return BatcherServiceFromCLIConfig(cliCtx.Context, version, cfg, l)
	}
}
