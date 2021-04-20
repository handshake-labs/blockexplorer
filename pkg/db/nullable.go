package db

import (
	"github.com/handshake-labs/blockexplorer/pkg/types"
)

func nullableInt32(v int32) *int32 {
	if v < 0 {
		return nil
	}
	return &v
}

func nullableInt64(v int64) *int64 {
	if v < 0 {
		return nil
	}
	return &v
}

func nullableBytes(v types.Bytes) *types.Bytes {
	if len(v) == 0 {
		return nil
	}
	return &v
}

func (row *GetTransactionByTxidRow) BlockHeight() *int32 {
	return nullableInt32(row.BlockHeightNotNull)
}

func (row *GetTransactionsByBlockHeightRow) BlockHeight() *int32 {
	return nullableInt32(row.BlockHeightNotNull)
}

func (row *GetNameBidsByHashRow) BlockHeight() *int32 {
	return nullableInt32(row.BlockHeightNotNull)
}

func (row *GetNameBidsByHashRow) RevealHeight() *int32 {
	return nullableInt32(row.RevealHeightNotNull)
}

func (row *GetNameBidsByHashRow) RevealIndex() *int32 {
	return nullableInt32(row.RevealIndexNotNull)
}

func (row *GetNameBidsByHashRow) Index() *int32 {
	return nullableInt32(row.RevealIndexNotNull)
}

func (row *GetNameBidsByHashRow) RevealValue() *int64 {
	return nullableInt64(row.RevealValueNotNull)
}

func (row *GetNameRecordsByHashRow) BlockHeight() *int32 {
	return nullableInt32(row.BlockHeightNotNull)
}

func (row *GetTxOutputsByAddressRow) SpendHeight() *int32 {
	return nullableInt32(row.SpendHeightNotNull)
}

func (row *GetTxOutputsByAddressRow) Height() *int32 {
	if row.HeightNotNull == 2147483647 {
		return nil
	}
	return nullableInt32(row.HeightNotNull)
}

func (row *GetTxOutputsByAddressRow) IndexPrevout() *int64 {
	return nullableInt64(row.IndexPrevoutNotNull)
}

func (row *GetTxOutputsByAddressRow) HashPrevout() *types.Bytes {
	return nullableBytes(row.HashPrevoutNotNull)
}
