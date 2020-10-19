package actions

import (
	"database/sql"
	"encoding/hex"
	"strconv"

	"github.com/handshake-labs/blockexplorer/pkg/types"
	"golang.org/x/net/idna"
)

type SearchParams struct {
	Query string `json:"query"`
}

type SearchResult struct {
	Transaction string `json:"transaction"`
	BlockHeight int32  `json:"block"`
	Name        string `json:"name"`
}

func Search(ctx *Context, params *SearchParams) (*SearchResult, error) {
	var tx, name string
	var blockHeight int32
	var result SearchResult
	query := params.Query
	if len(query) == 64 {
		if hash, err := hex.DecodeString(query); err == nil {
			hexString := types.Bytes(hash)
			//check if it's a transaction hash, if there is such a tx, then redirect there, otherwise give a name result
			if _, err := ctx.db.GetTransactionByTxid(ctx, hexString); err != sql.ErrNoRows {
				tx = query
			}
			//check if it's a block hash, if there is a block of such hash, redirect there, otherwsie give a name resulkt
			if block, err := ctx.db.GetBlockByHash(ctx, hexString); err != sql.ErrNoRows {
				blockHeight = block.Height
			}
		}
	}
	if height, err := strconv.Atoi(query); err == nil {
		//otherwise check if it's a string of ints, therefore it's a block
		blockHeight = int32(height)
	}
	punycoded_name, err := idna.ToASCII(query)
	if err == nil {
		name = string(punycoded_name)
		if name == "localhost" || name == "local" || name == "invalid" || name == "test" || name == "example" {
			return &result, nil
		}
	}
	result.BlockHeight = blockHeight
	result.Transaction = tx
	result.Name = name
	return &result, nil
}
