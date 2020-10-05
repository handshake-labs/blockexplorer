package db

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

func (row *GetTransactionByTxidRow) BlockHeight() *int32 {
	return nullableInt32(row.BlockHeightNotNull)
}

func (row *GetTransactionsByBlockHeightRow) BlockHeight() *int32 {
	return nullableInt32(row.BlockHeightNotNull)
}

func (row *GetNameBidsByHashRow) BlockHeight() *int32 {
	return nullableInt32(row.BlockHeightNotNull)
}

func (row *GetNameBidsByHashRow) RevealValue() *int64 {
	return nullableInt64(row.RevealValueNotNull)
}

func (row *GetNameRecordsByHashRow) BlockHeight() *int32 {
	return nullableInt32(row.BlockHeightNotNull)
}
