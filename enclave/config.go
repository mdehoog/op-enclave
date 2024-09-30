package enclave

import (
	"math/big"

	"github.com/ethereum-optimism/optimism/op-chain-ops/deployer/state"
	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

const (
	version0 uint64 = 0
)

var uint256Type abi.Type
var uint64Type abi.Type
var addressType abi.Type
var bytes32Type abi.Type

var chainConfigTemplate params.ChainConfig

func init() {
	uint256Type, _ = abi.NewType("uint256", "", nil)
	uint64Type, _ = abi.NewType("uint64", "", nil)
	addressType, _ = abi.NewType("address", "", nil)
	bytes32Type, _ = abi.NewType("bytes32", "", nil)

	deployConfig := state.DefaultDeployConfig()
	deployConfig.L2ChainID = 1
	var err error
	chainConfigTemplate, err = newChainConfigTemplate(&deployConfig)
	if err != nil {
		panic(err)
	}
}

type ChainConfig struct {
	*params.ChainConfig
	*PerChainConfig
}

func newChainConfigTemplate(cfg *genesis.DeployConfig) (params.ChainConfig, error) {
	l1StartHeader := &types.Header{
		Time: 1,
	}
	g, err := genesis.NewL2Genesis(cfg, l1StartHeader)
	if err != nil {
		return params.ChainConfig{}, err
	}
	return *g.Config, nil
}

func NewChainConfig(cfg *PerChainConfig) *ChainConfig {
	cfg.ForceDefaults()
	chainConfig := chainConfigTemplate
	chainConfig.ChainID = cfg.ChainID
	return &ChainConfig{
		ChainConfig:    &chainConfig,
		PerChainConfig: cfg,
	}
}

type PerChainConfig struct {
	ChainID *big.Int `json:"chain_id"`

	Genesis   rollup.Genesis `json:"genesis"`
	BlockTime uint64         `json:"block_time"`

	DepositContractAddress common.Address `json:"deposit_contract_address"`
	L1SystemConfigAddress  common.Address `json:"l1_system_config_address"`
}

func FromRollupConfig(cfg *rollup.Config) *PerChainConfig {
	p := &PerChainConfig{
		ChainID:                cfg.L2ChainID,
		Genesis:                cfg.Genesis,
		BlockTime:              cfg.BlockTime,
		DepositContractAddress: cfg.DepositContractAddress,
		L1SystemConfigAddress:  cfg.L1SystemConfigAddress,
	}
	p.ForceDefaults()
	return p
}

func (p *PerChainConfig) ToRollupConfig() *rollup.Config {
	return &rollup.Config{
		L2ChainID:              p.ChainID,
		Genesis:                p.Genesis,
		BlockTime:              p.BlockTime,
		DepositContractAddress: p.DepositContractAddress,
		L1SystemConfigAddress:  p.L1SystemConfigAddress,
	}
}

func (p *PerChainConfig) ForceDefaults() {
	p.BlockTime = 1
	p.Genesis.L2.Number = 0
	p.Genesis.SystemConfig.Overhead = eth.Bytes32{}
}

func (p *PerChainConfig) Hash() (common.Hash, error) {
	data, err := p.MarshalBinary()
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(data), nil
}

func (p *PerChainConfig) MarshalBinary() (data []byte, err error) {
	args := abi.Arguments{
		{Name: "version", Type: uint64Type},
		{Name: "chainID", Type: uint256Type},
		{Name: "genesisL1Hash", Type: bytes32Type},
		{Name: "genesisL1Number", Type: uint64Type},
		{Name: "genesisL2Hash", Type: bytes32Type},
		{Name: "genesisL2Time", Type: uint64Type},
		{Name: "genesisBatcherAddress", Type: addressType},
		{Name: "genesisScalar", Type: bytes32Type},
		{Name: "genesisGasLimit", Type: uint64Type},
		{Name: "depositContractAddress", Type: addressType},
		{Name: "l1SystemConfigAddress", Type: addressType},
	}
	return args.Pack(
		version0,
		p.ChainID,
		p.Genesis.L1.Hash,
		p.Genesis.L1.Number,
		p.Genesis.L2.Hash,
		p.Genesis.L2Time,
		p.Genesis.SystemConfig.BatcherAddr,
		p.Genesis.SystemConfig.Scalar,
		p.Genesis.SystemConfig.GasLimit,
		p.DepositContractAddress,
		p.L1SystemConfigAddress,
	)
}
