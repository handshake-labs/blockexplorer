package actions

import (
	// "database/sql"
	// "context"
	// "golang.org/x/crypto/sha3"
	// "github.com/handshake-labs/blockexplorer/pkg/types"
	// "github.com/handshake-labs/blockexplorer/pkg/db"
	// "github.com/jinzhu/copier"
	"encoding/json"
	"golang.org/x/net/idna"
	"log"
	"strconv"
)

type SearchParams struct {
	Query string `json:"query"`
}

type SearchResult struct {
	// NameRows []db.NameRow `json:"namerows"`
	// Count int32 `json:"count"`
	// Limit int16 `json:"limit"`
	Response json.RawMessage `json:"result"`
}

func Search(ctx *Context, params *SearchParams) (*SearchResult, error) {

	//QUERY IS A HASH
	query := params.Query
	if len(query) == 64 {
		//check if it's a transaction hash, if there is such a tx, then redirect there, otherwise give a name result
		//check if it's a block hash, if there is a block of such hash, redirect there, otherwsie give a name resulkt
	}

	if height, err := strconv.Atoi(query); err == nil {
		log.Println(height)
		//otherwise check if it's a string of ints, therefore it's a block
		//propose either the block of given height, or a namestring of the same integer
		return nil, nil
	}

	_, err := idna.ToASCII(query)
	if err != nil {
		return nil, nil
		// log.Println(err)
	}

	return nil, nil
}
