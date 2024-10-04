package enclave

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type l1ReceiptsFetcher struct {
	hash     common.Hash
	header   *types.Header
	receipts types.Receipts
}

func NewL1ReceiptsFetcher(hash common.Hash, header *types.Header, receipts types.Receipts) derive.L1ReceiptsFetcher {
	return &l1ReceiptsFetcher{
		hash:     hash,
		header:   header,
		receipts: receipts,
	}
}

func (l *l1ReceiptsFetcher) InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error) {
	if l.hash != hash {
		return nil, errors.New("not found")
	}
	return headerInfo{
		hash:   l.hash,
		Header: l.header,
	}, nil
}

func (l *l1ReceiptsFetcher) FetchReceipts(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Receipts, error) {
	info, err := l.InfoByHash(ctx, blockHash)
	if err != nil {
		return nil, nil, err
	}
	return info, l.receipts, nil
}

type headerInfo struct {
	hash common.Hash
	*types.Header
}

var _ eth.BlockInfo = (*headerInfo)(nil)

func (h headerInfo) Hash() common.Hash {
	return h.hash
}

func (h headerInfo) ParentHash() common.Hash {
	return h.Header.ParentHash
}

func (h headerInfo) Coinbase() common.Address {
	return h.Header.Coinbase
}

func (h headerInfo) Root() common.Hash {
	return h.Header.Root
}

func (h headerInfo) NumberU64() uint64 {
	return h.Header.Number.Uint64()
}

func (h headerInfo) Time() uint64 {
	return h.Header.Time
}

func (h headerInfo) MixDigest() common.Hash {
	return h.Header.MixDigest
}

func (h headerInfo) BaseFee() *big.Int {
	return h.Header.BaseFee
}

func (h headerInfo) BlobBaseFee() *big.Int {
	if h.Header.ExcessBlobGas == nil {
		return nil
	}
	return eip4844.CalcBlobFee(*h.Header.ExcessBlobGas)
}

func (h headerInfo) ReceiptHash() common.Hash {
	return h.Header.ReceiptHash
}

func (h headerInfo) GasUsed() uint64 {
	return h.Header.GasUsed
}

func (h headerInfo) GasLimit() uint64 {
	return h.Header.GasLimit
}

func (h headerInfo) ParentBeaconRoot() *common.Hash {
	return h.Header.ParentBeaconRoot
}

func (h headerInfo) HeaderRLP() ([]byte, error) {
	return rlp.EncodeToBytes(h.Header)
}
