package actions

import (
	"database/sql"

	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/sha3"
)

func nameHash(name string) (types.Bytes, error) {
	sha := sha3.New256()
	_, err := sha.Write([]byte(name))
	if err != nil {
		return nil, err
	}
	return types.Bytes(sha.Sum(nil)), nil
}

//returns modulo of uint represented by array of bytes
func modulo(x []byte, j int) int {
	var m int
	for i := 0; i < len(x); i++ {
		m <<= 8
		m += (int(x[i]) & 0xff)
		m %= j
	}
	return m
}

type GetNameParams struct {
	Name string `json:"name"`
}

type GetNameResult struct {
	ReservedName *ReservedName `json:"reserved,omitempty"`
	ReleaseBlock int           `json:"release_block"`
	BidsCount    int32         `json:"bids_count"`
	RecordsCount int32         `json:"records_count"`
	State        string        `json:"state"`
}

//get state of the name relative to the block
func getStateByName(ctx *Context, height int, name string) string {
	nameHash, _ := nameHash(name)
	openHeightParams := db.GetLastHeightByActionByHashParams{db.CovenantAction("OPEN"), &nameHash}
	openHeight, err := ctx.db.GetLastHeightByActionByHash(ctx, openHeightParams)
	if err != sql.ErrNoRows {
		return "Auction has not opened."
	}
	if openHeight+17 <= height {
		return "Auction is opening."
	}
	if openHeight+17+144*5 >= height {
		return "Auction is in bid state."
	}
	if openHeight+17+144*15 >= height {
		return "Auction is in reveal state."
	}
	// openHeight += 2
	revealHeightParams := db.GetLastHeightByActionByHashParams{db.CovenantAction("REVEAL"), &nameHash}
	_, err = ctx.db.GetLastHeightByActionByHash(ctx, revealHeightParams)
	if err != sql.ErrNoRows {
		return "Auction has been finished."
	}
	if err == sql.ErrNoRows {
		return "Auction has not been concluded. New auction can be opened at block " + string(openHeight+17+144*15+144*10) + "."
	}
	return "Error"
}

func GetName(ctx *Context, params *GetNameParams) (*GetNameResult, error) {
	hash, err := nameHash(params.Name)
	if err != nil {
		return nil, err
	}
	counts, err := ctx.db.GetNameCountsByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	height, _ := ctx.db.GetBlocksMaxHeight(ctx)
	result := GetNameResult{nil, ReleaseBlock(params.Name), counts.BidsCount, counts.RecordsCount, getStateByName(ctx, int(height), params.Name)}
	name, err := ctx.db.GetReservedName(ctx, params.Name)
	if err == nil {
		result.ReservedName = &ReservedName{}
		copier.Copy(&result.ReservedName, &name)
	} else if err != sql.ErrNoRows {
		return nil, err
	}
	return &result, nil
}

type GetNameBidsByHashParams struct {
	Name   string `json:"name"`
	Limit  int8   `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetNameBidsByHashResult struct {
	NameBids []NameBid `json:"bids"`
}

func GetNameBidsByHash(ctx *Context, params *GetNameBidsByHashParams) (*GetNameBidsByHashResult, error) {
	hash, err := nameHash(params.Name)
	if err != nil {
		return nil, err
	}
	bids, err := ctx.db.GetNameBidsByHash(ctx, db.GetNameBidsByHashParams{
		NameHash: hash,
		Limit:    int32(params.Limit),
		Offset:   params.Offset,
	})
	if err != nil {
		return nil, err
	}
	result := GetNameBidsByHashResult{[]NameBid{}}
	copier.Copy(&result.NameBids, &bids)
	return &result, nil
}

type GetNameRecordsByHashParams struct {
	Name   string `json:"name"`
	Limit  int8   `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetNameRecordsByHashResult struct {
	NameRecords []NameRecord `json:"records"`
}

func GetNameRecordsByHash(ctx *Context, params *GetNameRecordsByHashParams) (*GetNameRecordsByHashResult, error) {
	hash, err := nameHash(params.Name)
	if err != nil {
		return nil, err
	}
	records, err := ctx.db.GetNameRecordsByHash(ctx, db.GetNameRecordsByHashParams{
		NameHash: hash,
		Limit:    int32(params.Limit),
		Offset:   params.Offset,
	})
	if err != nil {
		return nil, err
	}
	result := GetNameRecordsByHashResult{[]NameRecord{}}
	copier.Copy(&result.NameRecords, &records)
	return &result, nil
}

func ReleaseBlock(name string) int {
	hash, _ := nameHash(name)
	return modulo([]byte(hash), 52)*144*7 + 2016
}
