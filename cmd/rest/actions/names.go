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

type GetNameParams struct {
	Name string `json:"name"`
}

type GetNameResult struct {
	ReservedName *ReservedName `json:"reserved,omitempty"`
	BidsCount    int32         `json:"bids_count"`
	RecordsCount int32         `json:"records_count"`
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
	result := GetNameResult{nil, counts.BidsCount, counts.RecordsCount}
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
