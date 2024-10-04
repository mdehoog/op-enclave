package bindings

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type deployChainDeployMarshal struct {
	ChainID    *big.Int
	ConfigHash common.Hash
	OutputRoot common.Hash
	BatchInbox common.Address
	Addresses  DeployChainDeployAddresses
	Raw        types.Log
}

func (d *DeployChainDeploy) MarshalJSON() ([]byte, error) {
	m := deployChainDeployMarshal{
		ChainID:    d.ChainID,
		ConfigHash: d.ConfigHash,
		OutputRoot: d.OutputRoot,
		BatchInbox: d.BatchInbox,
		Addresses:  d.Addresses,
		Raw:        d.Raw,
	}
	return json.Marshal(m)
}

func (d *DeployChainDeploy) UnmarshalJSON(data []byte) error {
	var m deployChainDeployMarshal
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	d.ChainID = m.ChainID
	d.ConfigHash = m.ConfigHash
	d.OutputRoot = m.OutputRoot
	d.BatchInbox = m.BatchInbox
	d.Addresses = m.Addresses
	d.Raw = m.Raw
	return nil
}
