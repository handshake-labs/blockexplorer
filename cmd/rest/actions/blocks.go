package actions

import (
	"database/sql"

	"github.com/jinzhu/copier"
)

type GetBlockByHeightParams struct {
	Height int32 `json:"height"`
}

type GetBlockByHeightResult struct {
	Block           Block `json:"block"`
	BlocksMaxHeight int32 `json:"maxHeight"`
}

func GetBlockByHeight(ctx *Context, params *GetBlockByHeightParams) (*GetBlockByHeightResult, error) {
	result := GetBlockByHeightResult{}
	block, err := ctx.db.GetBlockByHeight(ctx, params.Height)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	copier.Copy(&result.Block, &block)
	if height, err := ctx.db.GetBlocksMaxHeight(ctx); err == nil {
		result.BlocksMaxHeight = height
	} else {
		return nil, err
	}
	return &result, nil
}
