package actions

import (
	"database/sql"

	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/sha3"
	"strings"
)

const treeInterval = 36
const blocksPerDay = 144

func nameHash(name string) (types.Bytes, error) {
	sha := sha3.New256()
	_, err := sha.Write([]byte(name))
	if err != nil {
		return nil, err
	}
	return types.Bytes(sha.Sum(nil)), nil
}

//returns modulo of uint represented by array of bytes
func modulo(x []byte, j int32) int32 {
	var m int32
	for i := 0; i < len(x); i++ {
		m <<= 8
		m += (int32(x[i]) & 0xff)
		m %= j
	}
	return m
}

type GetNameParams struct {
	Name string `json:"name"`
}

type GetNameResult struct {
	ReservedName       *ReservedName `json:"reserved,omitempty"`
	ReleaseBlockHeight int32         `json:"release_block,omitempty"`
	BidsCount          int32         `json:"bids_count,omitempty"`
	RecordsCount       int32         `json:"records_count,omitempty"`
	State              *State        `json:"state,omitempty"`
}

func GetName(ctx *Context, params *GetNameParams) (*GetNameResult, error) {
	result := GetNameResult{}
	nameString := strings.ToLower(params.Name)
	if len(nameString) > 63 || nameString == "localhost" || nameString == "local" || nameString == "invalid" || nameString == "test" || nameString == "example" {
		return &result, nil
	}
	hash, err := nameHash(nameString)
	if err != nil {
		return nil, err
	}
	counts, err := ctx.db.GetNameCountsByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	height, err := ctx.db.GetBlocksMaxHeight(ctx)
	if err != nil {
		return nil, err
	}
	state, err := getStateByName(ctx, height, nameString)
	if err != nil {
		return nil, err
	}
	result = GetNameResult{nil, ReleaseBlock(nameString), counts.BidsCount, counts.RecordsCount, state}
	name, err := ctx.db.GetReservedName(ctx, nameString)
	if err == nil {
		result.ReservedName = &ReservedName{}
		copier.Copy(&result.ReservedName, &name)
	} else if err != sql.ErrNoRows {
		return nil, err
	}
	return &result, nil
}

type State struct {
	OpenHeight      int32        `json:"open_height,omitempty"`
	CurrentState    AuctionState `json:"current_state"`
	AuctionComplete bool         `json:"auction_completed"`
}

//get state of the name relative to the block
//if the auction has not concluded then name can be opened again after TreeInterval is elapsed
func getStateByName(ctx *Context, height int32, name string) (*State, error) {
	state := State{}
	nameHash, _ := nameHash(name)
	openHeightParams := db.GetLastNameBlockHeightByActionAndHashParams{db.CovenantAction("OPEN"), &nameHash}
	openHeight, err := ctx.db.GetLastNameBlockHeightByActionAndHash(ctx, openHeightParams)
	if err == sql.ErrNoRows || openHeight == -1 {
		state.CurrentState = AuctionStateClosed
		state.AuctionComplete = false
	} else if err != nil {
		return nil, err
	}
	state.OpenHeight = openHeight
	if openHeight+treeInterval >= height {
		state.CurrentState = AuctionStateClosed
		state.AuctionComplete = false
		return &state, nil
	}
	if openHeight+treeInterval+1+blocksPerDay*5 >= height {
		state.CurrentState = AuctionStateBid
		state.AuctionComplete = false
		return &state, nil
	}
	if openHeight+treeInterval+1+blocksPerDay*15 >= height {
		state.CurrentState = AuctionStateReveal
		state.AuctionComplete = false
		return &state, nil
	}
	revealHeightParams := db.GetLastNameBlockHeightByActionAndHashParams{db.CovenantAction("REVEAL"), &nameHash}
	revealHeight, err := ctx.db.GetLastNameBlockHeightByActionAndHash(ctx, revealHeightParams)
	if revealHeight >= openHeight {
		state.CurrentState = AuctionStateClosed
		state.AuctionComplete = true
		return &state, nil
	}
	if err == sql.ErrNoRows {
		state.CurrentState = AuctionStateClosed
		state.AuctionComplete = false
		return &state, nil
	} else if err != nil {
		return nil, err
	}
	claimHeightParams := db.GetLastNameBlockHeightByActionAndHashParams{db.CovenantAction("CLAIM"), &nameHash}
	_, err = ctx.db.GetLastNameBlockHeightByActionAndHash(ctx, claimHeightParams)
	return &state, nil
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
	nameString := strings.ToLower(params.Name)
	hash, err := nameHash(nameString)
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
	nameString := strings.ToLower(params.Name)
	hash, err := nameHash(nameString)
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

func ReleaseBlock(name string) int32 {
	hash, _ := nameHash(name)
	return modulo([]byte(hash), 52)*blocksPerDay*7 + 2016
}
