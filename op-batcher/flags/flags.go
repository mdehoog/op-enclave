package flags

import (
	"github.com/ethereum-optimism/optimism/op-batcher/flags"
	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/urfave/cli/v2"
)

const EnvVarPrefix = "OP_BATCHER"

func prefixEnvVars(name string) []string {
	return opservice.PrefixEnvVar(EnvVarPrefix, name)
}

var (
	// Required flags
	//EnclaveRpcFlag = &cli.StringFlag{
	//	Name:     "enclave-rpc",
	//	Usage:    "HTTP provider URL for the enclave service",
	//	EnvVars:  prefixEnvVars("ENCLAVE_RPC"),
	//	Required: true,
	//}
	DAURLFlag = &cli.StringFlag{
		Name:     "da-url",
		Usage:    "URL for the Data Availability uploads",
		EnvVars:  prefixEnvVars("DA_URL"),
		Required: true,
	}
	//OutputOracleAddressFlag = &cli.StringFlag{
	//	Name:     "output-oracle-address",
	//	Usage:    "Address of the output oracle contract",
	//	EnvVars:  prefixEnvVars("OUTPUT_ORACLE_ADDRESS"),
	//	Required: true,
	//}
)

var requiredFlags = []cli.Flag{
	DAURLFlag,
}

func init() {
	Flags = append(requiredFlags, flags.Flags...)
}

var Flags []cli.Flag
