package actions

import (
	"database/sql"

	"github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/jinzhu/copier"
)

type GetBlockByHeightParams struct {
	Height int32 `json:"height"`
}

type GetBlockByHeightResult struct {
	Block         Block       `json:"block"`
	PrevBlockHash types.Bytes `json:"prev_block_hash,omitempty"`
	NextBlockHash types.Bytes `json:"next_block_hash,omitempty"`
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
	if hash, err := ctx.db.GetBlockHashByHeight(ctx, block.Height-1); err == nil {
		result.PrevBlockHash = hash
	} else if err != sql.ErrNoRows {
		return nil, err
	}
	if hash, err := ctx.db.GetBlockHashByHeight(ctx, block.Height+1); err == nil {
		result.NextBlockHash = hash
	} else if err != sql.ErrNoRows {
		return nil, err
	}
	return &result, nil
}
