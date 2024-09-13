package enclave

import (
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

const (
	version0 uint64 = 0
)

var uint256Type abi.Type
var uint64Type abi.Type
var boolType abi.Type
var addressType abi.Type
var bytes32Type abi.Type

func init() {
	uint256Type, _ = abi.NewType("uint256", "", nil)
	uint64Type, _ = abi.NewType("uint64", "", nil)
	boolType, _ = abi.NewType("bool", "", nil)
	addressType, _ = abi.NewType("address", "", nil)
	bytes32Type, _ = abi.NewType("bytes32", "", nil)
}

type RollupConfig struct {
	params.ChainConfig

	Genesis   rollup.Genesis `json:"genesis"`
	BlockTime uint64         `json:"block_time"`

	DepositContractAddress common.Address `json:"deposit_contract_address"`
	L1SystemConfigAddress  common.Address `json:"l1_system_config_address"`
}

func (c *RollupConfig) ToRollupConfig() *rollup.Config {
	return &rollup.Config{
		Genesis:                c.Genesis,
		BlockTime:              c.BlockTime,
		DepositContractAddress: c.DepositContractAddress,
		L1SystemConfigAddress:  c.L1SystemConfigAddress,
	}
}

func (c *RollupConfig) Hash() (common.Hash, error) {
	data, err := c.MarshalBinary()
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(data), nil
}

func (c *RollupConfig) MarshalBinary() (data []byte, err error) {
	args := abi.Arguments{
		{Name: "version", Type: uint64Type},
		{Name: "chainID", Type: uint256Type},
		{Name: "homesteadBlock", Type: uint256Type},
		{Name: "daoForkBlock", Type: uint256Type},
		{Name: "daoForkSupport", Type: boolType},
		{Name: "eip150Block", Type: uint256Type},
		{Name: "eip155Block", Type: uint256Type},
		{Name: "eip158Block", Type: uint256Type},
		{Name: "byzantiumBlock", Type: uint256Type},
		{Name: "constantinopleBlock", Type: uint256Type},
		{Name: "petersburgBlock", Type: uint256Type},
		{Name: "istanbulBlock", Type: uint256Type},
		{Name: "muirGlacierBlock", Type: uint256Type},
		{Name: "berlinBlock", Type: uint256Type},
		{Name: "londonBlock", Type: uint256Type},
		{Name: "arrowGlacierBlock", Type: uint256Type},
		{Name: "grayGlacierBlock", Type: uint256Type},
		{Name: "mergeNetsplitBlock", Type: uint256Type},
		{Name: "shanghaiTime", Type: uint64Type},
		{Name: "cancunTime", Type: uint64Type},
		{Name: "pragueTime", Type: uint64Type},
		{Name: "verkleTime", Type: uint64Type},
		{Name: "bedrockBlock", Type: uint256Type},
		{Name: "regolithTime", Type: uint64Type},
		{Name: "canyonTime", Type: uint64Type},
		{Name: "ecotoneTime", Type: uint64Type},
		{Name: "fjordTime", Type: uint64Type},
		{Name: "graniteTime", Type: uint64Type},
		{Name: "interopTime", Type: uint64Type},
		{Name: "terminalTotalDifficulty", Type: uint256Type},
		{Name: "terminalTotalDifficultyPassed", Type: boolType},
		{Name: "cliquePeriod", Type: uint64Type},
		{Name: "cliqueEpoch", Type: uint64Type},
		{Name: "optimismEIP1559Elasticity", Type: uint64Type},
		{Name: "optimismEIP1559Denominator", Type: uint64Type},
		{Name: "optimismEIP1559DenominatorCanyon", Type: uint64Type},
		{Name: "genesisL1Hash", Type: bytes32Type},
		{Name: "genesisL1Number", Type: uint64Type},
		{Name: "genesisL2Hash", Type: bytes32Type},
		{Name: "genesisL2Number", Type: uint64Type},
		{Name: "genesisL2Time", Type: uint64Type},
		{Name: "genesisSystemConfigBatcherAddr", Type: addressType},
		{Name: "genesisSystemConfigOverhead", Type: bytes32Type},
		{Name: "genesisSystemConfigScalar", Type: bytes32Type},
		{Name: "genesisSystemConfigGasLimit", Type: uint64Type},
		{Name: "blockTime", Type: uint64Type},
		{Name: "depositContractAddress", Type: addressType},
		{Name: "l1SystemConfigAddress", Type: addressType},
	}
	return args.Pack(
		version0,
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		c.IstanbulBlock,
		c.MuirGlacierBlock,
		c.BerlinBlock,
		c.LondonBlock,
		c.ArrowGlacierBlock,
		c.GrayGlacierBlock,
		c.MergeNetsplitBlock,
		c.ShanghaiTime,
		c.CancunTime,
		c.PragueTime,
		c.VerkleTime,
		c.BedrockBlock,
		c.RegolithTime,
		c.CanyonTime,
		c.EcotoneTime,
		c.FjordTime,
		c.GraniteTime,
		c.InteropTime,
		c.TerminalTotalDifficulty,
		c.TerminalTotalDifficultyPassed,
		c.Clique.Period,
		c.Clique.Epoch,
		c.Optimism.EIP1559Elasticity,
		c.Optimism.EIP1559Denominator,
		c.Optimism.EIP1559DenominatorCanyon,
		c.Genesis.L1.Hash,
		c.Genesis.L1.Number,
		c.Genesis.L2.Hash,
		c.Genesis.L2.Number,
		c.Genesis.L2Time,
		c.Genesis.SystemConfig.BatcherAddr,
		c.Genesis.SystemConfig.Overhead,
		c.Genesis.SystemConfig.Scalar,
		c.Genesis.SystemConfig.GasLimit,
		c.BlockTime,
		c.DepositContractAddress,
		c.L1SystemConfigAddress,
	)
}
