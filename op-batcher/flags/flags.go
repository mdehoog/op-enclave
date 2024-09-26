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
	DAURLFlag = &cli.StringFlag{
		Name:     "da-url",
		Usage:    "URL for the Data Availability uploads",
		EnvVars:  prefixEnvVar("DA_URL"),
		Required: true,
	}
)

var requiredFlags = []cli.Flag{
	DAURLFlag,
}

func init() {
	Flags = append(requiredFlags, flags.Flags...)
}

var Flags []cli.Flag
