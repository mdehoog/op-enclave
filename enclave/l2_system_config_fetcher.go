package enclave

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type l2SystemConfigFetcher struct {
	config *rollup.Config
	hash   common.Hash
	header *types.Header
	txs    []*types.Transaction
}

func NewL2SystemConfigFetcher(config *rollup.Config, hash common.Hash, header *types.Header, txs []*types.Transaction) derive.SystemConfigL2Fetcher {
	return &l2SystemConfigFetcher{
		config: config,
		hash:   hash,
		header: header,
		txs:    txs,
	}
}

func (l *l2SystemConfigFetcher) SystemConfigByL2Hash(ctx context.Context, hash common.Hash) (eth.SystemConfig, error) {
	if hash != l.hash {
		return eth.SystemConfig{}, errors.New("not found")
	}
	return blockToSystemConfig(l.config, l.header, l.txs)
}

// Copy of https://github.com/ethereum-optimism/optimism/blob/8b61225d51105b142580d40bde43adde791423b8/op-node/rollup/derive/payload_util.go#L54
// but takes a header/txs rather than an execution payload.
func blockToSystemConfig(rollupCfg *rollup.Config, header *types.Header, txs []*types.Transaction) (eth.SystemConfig, error) {
	hash := header.Hash()
	if header.Number.Uint64() == rollupCfg.Genesis.L2.Number {
		if hash != rollupCfg.Genesis.L2.Hash {
			return eth.SystemConfig{}, fmt.Errorf(
				"expected L2 genesis hash to match L2 block at genesis block number %d: %s <> %s",
				rollupCfg.Genesis.L2.Number, hash, rollupCfg.Genesis.L2.Hash)
		}
		return rollupCfg.Genesis.SystemConfig, nil
	} else {
		if len(txs) == 0 {
			return eth.SystemConfig{}, fmt.Errorf("l2 block is missing L1 info deposit tx, block hash: %s", hash)
		}
		tx := txs[0]
		if tx.Type() != types.DepositTxType {
			return eth.SystemConfig{}, fmt.Errorf("first payload tx has unexpected tx type: %d", tx.Type())
		}
		info, err := derive.L1BlockInfoFromBytes(rollupCfg, header.Time, tx.Data())
		if err != nil {
			return eth.SystemConfig{}, fmt.Errorf("failed to parse L1 info deposit tx from L2 block: %w", err)
		}
		if isEcotoneButNotFirstBlock(rollupCfg, header.Time) {
			// Translate Ecotone values back into encoded scalar if needed.
			// We do not know if it was derived from a v0 or v1 scalar,
			// but v1 is fine, a 0 blob base fee has the same effect.
			info.L1FeeScalar[0] = 1
			binary.BigEndian.PutUint32(info.L1FeeScalar[24:28], info.BlobBaseFeeScalar)
			binary.BigEndian.PutUint32(info.L1FeeScalar[28:32], info.BaseFeeScalar)
		}
		return eth.SystemConfig{
			BatcherAddr: info.BatcherAddr,
			Overhead:    info.L1FeeOverhead,
			Scalar:      info.L1FeeScalar,
			GasLimit:    header.GasLimit,
		}, err
	}
}

func isEcotoneButNotFirstBlock(rollupCfg *rollup.Config, l2BlockTime uint64) bool {
	return rollupCfg.IsEcotone(l2BlockTime) && !rollupCfg.IsEcotoneActivationBlock(l2BlockTime)
}
