package actions

import (
	"database/sql"

	"github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/jinzhu/copier"

	"github.com/handshake-labs/blockexplorer/pkg/db"
)

type GetBlockByHashParams struct {
	Hash types.Bytes `json:"hash"`
}

type GetBlockByHashResult struct {
	Block Block `json:"block"`
}

func GetBlockByHash(ctx *Context, params *GetBlockByHashParams) (*GetBlockByHashResult, error) {
	block, err := ctx.db.GetBlockByHash(ctx, params.Hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	result := &GetBlockByHashResult{}
	copier.Copy(&result.Block, &block)
	return result, nil
}

type GetBlocksParams struct {
	Limit  int8  `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetBlocksResult struct {
	Blocks []Block `json:"blocks"`
	Count  int32   `json:"count"`
}

func GetBlocks(ctx *Context, params *GetBlocksParams) (*GetBlocksResult, error) {
	blocks, err := ctx.db.GetBlocks(ctx, db.GetBlocksParams{
		Limit:  int32(params.Limit),
		Offset: params.Offset,
	})
	if err != nil {
		return nil, err
	}
	maxHeight, err := ctx.db.GetBlocksMaxHeight(ctx)
	if err != nil {
		return nil, err
	}
	result := &GetBlocksResult{make([]Block, 0), maxHeight + 1}
	copier.Copy(&result.Blocks, &blocks)
	return result, nil
}
