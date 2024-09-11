package enclave

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
)

type Block struct {
	*types.Block
}

type txs struct {
	Transactions []*types.Transaction
}

func (b *Block) UnmarshalJSON(input []byte) error {
	t := &txs{}
	err := json.Unmarshal(input, t)
	if err != nil {
		return err
	}

	h := &types.Header{}
	err = json.Unmarshal(input, h)
	if err != nil {
		return err
	}
	h.TxHash = types.DeriveSha(types.Transactions(t.Transactions), trie.NewStackTrie(nil))

	h.Root = common.Hash{}
	h.ReceiptHash = common.Hash{}

	b.Block = types.NewBlockWithHeader(h).WithBody(types.Body{Transactions: t.Transactions})
	return nil
}
