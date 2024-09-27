package flags

import (
	opservice "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/urfave/cli/v2"
)

const EnvVarPrefix = "OP_DA"

func prefixEnvVar(name string) []string {
	return opservice.PrefixEnvVar(EnvVarPrefix, name)
}

var (
	PortFlag = &cli.IntFlag{
		Name:     "port",
		Usage:    "Port to listen on",
		EnvVars:  prefixEnvVar("PORT"),
		Required: true,
	}
	DAURLFlag = &cli.StringFlag{
		Name:     "da-url",
		Usage:    "URL for DA (file, http, S3)",
		EnvVars:  prefixEnvVar("DA_URL"),
		Required: true,
	}
)

var requiredFlags = []cli.Flag{
	PortFlag,
	DAURLFlag,
}

func init() {
	Flags = append(requiredFlags, oplog.CLIFlags(EnvVarPrefix)...)
}

var Flags []cli.Flag
