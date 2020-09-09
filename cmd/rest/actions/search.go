package actions

import (
	"database/sql"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	"golang.org/x/net/idna"
	"log"
	"strconv"
)

type SearchParams struct {
	Query string `json:"query"`
}

type SearchResult struct {
	Transactions []string `json:"transactions"`
	Blocks       []int32  `json:"blocks"`
	Names        []string `json:"names"`
}

func Search(ctx *Context, params *SearchParams) (*SearchResult, error) {
	var txs, names []string
	var blocks []int32
	var result SearchResult
	query := params.Query
	if len(query) == 64 {
		//check if it's a transaction hash, if there is such a tx, then redirect there, otherwise give a name result
		// log.Println("zzzzzz")
		log.Println(types.Bytes(query))
		// if t, err := ctx.db.GetTransactionByTxid(ctx, types.Bytes(query)); err != sql.ErrNoRows {
		t, err := ctx.db.GetTransactionByTxid(ctx, types.Bytes(query))
		// 	log.Println(t)
		// 	txs = append(txs, query)
		// }
		log.Println(t)
		log.Println(err)
		//check if it's a block hash, if there is a block of such hash, redirect there, otherwsie give a name resulkt
		if block, err3 := ctx.db.GetBlockByHash(ctx, types.Bytes(query)); err3 != sql.ErrNoRows {
			log.Println(err3)
			blocks = append(blocks, (block.Height))
		}
		// log.Println(block)
	}

	if height, err := strconv.Atoi(query); err == nil {
		//otherwise check if it's a string of ints, therefore it's a block
		blocks = append(blocks, int32(height))
	}

	punycoded_name, err := idna.ToASCII(query)
	if err == nil {
		names = append(names, string(punycoded_name))
	}
	result.Blocks = blocks
	result.Transactions = txs
	result.Names = names
	return &result, nil
}
