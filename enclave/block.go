package enclave

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Block struct {
	*types.Block
}

type txs struct {
	Transactions []*types.Transaction
}

func (b *Block) UnmarshalJSON(input []byte) error {
	h := &types.Header{}
	err := json.Unmarshal(input, h)
	if err != nil {
		return err
	}
	h.Root = common.Hash{}
	h.ReceiptHash = common.Hash{}

	txs := &txs{}
	err = json.Unmarshal(input, txs)
	if err != nil {
		return err
	}

	b.Block = types.NewBlockWithHeader(h).WithBody(types.Body{Transactions: txs.Transactions})
	return nil
}
