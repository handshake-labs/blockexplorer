package actions

import (
	"database/sql"

	"github.com/jinzhu/copier"

	"github.com/handshake-labs/blockexplorer/pkg/db"
)

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
	height, err := ctx.db.GetBlocksMaxHeight(ctx)
	if err != nil {
		return nil, err
	}
	result := GetBlocksResult{}
	result.Count = height + 1
	copier.Copy(&result.Blocks, &blocks)
	return &result, nil
}

type GetBlockByHeightParams struct {
	Height int32 `json:"height"`
}

type GetBlockByHeightResult struct {
	Block          Block `json:"block"`
	MaxBlockHeight int32 `json:"height"`
}

func GetBlockByHeight(ctx *Context, params *GetBlockByHeightParams) (*GetBlockByHeightResult, error) {
	block, err := ctx.db.GetBlockByHeight(ctx, params.Height)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	height, err := ctx.db.GetBlocksMaxHeight(ctx)
	if err != nil {
		return nil, err
	}
	result := GetBlockByHeightResult{}
	result.MaxBlockHeight = height
	copier.Copy(&result.Block, &block)
	return &result, nil
}
