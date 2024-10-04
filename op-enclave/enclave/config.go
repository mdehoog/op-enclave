package enclave

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-chain-ops/deployer/state"
	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

const (
	version0 uint64 = 0
)

var chainConfigTemplate params.ChainConfig
var rollupConfigTemplate rollup.Config

func init() {
	deployConfig := state.DefaultDeployConfig()
	deployConfig.L2ChainID = 1

	var err error
	chainConfigTemplate, rollupConfigTemplate, err = newChainConfigTemplate(&deployConfig)
	if err != nil {
		panic(err)
	}
}

type ChainConfig struct {
	*params.ChainConfig
	*PerChainConfig
}

func newChainConfigTemplate(cfg *genesis.DeployConfig) (params.ChainConfig, rollup.Config, error) {
	l1StartHeader := &types.Header{
		Time:   1,
		Number: big.NewInt(0),
	}
	g, err := genesis.NewL2Genesis(cfg, l1StartHeader)
	if err != nil {
		return params.ChainConfig{}, rollup.Config{}, err
	}

	cfg.OptimismPortalProxy = common.Address{1}
	cfg.SystemConfigProxy = common.Address{1}
	rollupConfig, err := cfg.RollupConfig(l1StartHeader, common.Hash{}, 0)
	if err != nil {
		return params.ChainConfig{}, rollup.Config{}, err
	}

	return *g.Config, *rollupConfig, nil
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
	cfg := rollupConfigTemplate
	cfg.L2ChainID = p.ChainID
	cfg.Genesis = p.Genesis
	cfg.BlockTime = p.BlockTime
	cfg.DepositContractAddress = p.DepositContractAddress
	cfg.L1SystemConfigAddress = p.L1SystemConfigAddress
	return &cfg
}

func (p *PerChainConfig) ForceDefaults() {
	p.BlockTime = 1
	p.Genesis.L2.Number = 0
	p.Genesis.SystemConfig.Overhead = eth.Bytes32{}
}

func (p *PerChainConfig) Hash() common.Hash {
	return crypto.Keccak256Hash(p.MarshalBinary())
}

func (p *PerChainConfig) MarshalBinary() (data []byte) {
	data = binary.BigEndian.AppendUint64(data, version0)
	chainIDBytes := p.ChainID.Bytes()
	data = append(data, make([]byte, 32-len(chainIDBytes))...)
	data = append(data, chainIDBytes...)
	data = append(data, p.Genesis.L1.Hash[:]...)
	data = append(data, p.Genesis.L2.Hash[:]...)
	data = binary.BigEndian.AppendUint64(data, p.Genesis.L2Time)
	data = append(data, p.Genesis.SystemConfig.BatcherAddr.Bytes()...)
	data = append(data, p.Genesis.SystemConfig.Scalar[:]...)
	data = binary.BigEndian.AppendUint64(data, p.Genesis.SystemConfig.GasLimit)
	data = append(data, p.DepositContractAddress.Bytes()...)
	data = append(data, p.L1SystemConfigAddress.Bytes()...)
	return data
}
