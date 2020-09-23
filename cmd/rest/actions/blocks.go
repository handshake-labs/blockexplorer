package actions

import (
	"database/sql"

	"github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/jinzhu/copier"
)

type GetBlockByHashParams struct {
	Hash types.Bytes `json:"hash"`
}

type GetBlockByHashResult struct {
	Block             Block `json:"block"`
	TransactionsCount int32 `json:"txs_count"`
}

func GetBlockByHash(ctx *Context, params *GetBlockByHashParams) (*GetBlockByHashResult, error) {
	block, err := ctx.db.GetBlockByHash(ctx, params.Hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	transactionsCount, err := ctx.db.CountTransactionsByBlockHash(ctx, params.Hash)
	if err != nil {
		return nil, err
	}
	result := GetBlockByHashResult{Block{}, transactionsCount}
	copier.Copy(&result.Block, &block)
	return &result, nil
}
