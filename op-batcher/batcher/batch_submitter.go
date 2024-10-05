package batcher

import (
	"context"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-batcher/flags"
	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum/go-ethereum/rpc"
	thisflags "github.com/mdehoog/op-enclave/op-batcher/flags"
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

		proposerClient, err := rpc.DialContext(cliCtx.Context, cliCtx.String(thisflags.ProposerRpcFlag.Name))
		if err != nil {
			return nil, fmt.Errorf("failed to connect to Proposer: %w", err)
		}

		l.Info("Initializing Batch Submitter")
		channelFactoryOpt := func(setup *batcher.DriverSetup) {
			metricer := NewMetricer(setup.Metr, setup.Log, proposerClient)
			setup.Metr = metricer
			setup.ChannelOutFactory = ChannelOutFactory(metricer)
		}
		return batcher.BatcherServiceFromCLIConfig(cliCtx.Context, version, cfg, l, channelFactoryOpt)
	}
}
