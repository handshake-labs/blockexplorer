package actions

import (
	"database/sql"

	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	"golang.org/x/crypto/sha3"
	// "github.com/jinzhu/copier"
	// "log"
)

type GetMostExpensiveNamesParams struct {
	Page int16 `json:"page"`
}

type GetMostExpensiveNamesParamsResult struct {
	NameRows []db.NameRow `json:"namerows"`
	Count    int32        `json:"count"`
	Limit    int16        `json:"limit"`
}

func GetTopList(ctx *Context, params *GetMostExpensiveNamesParams) (*GetMostExpensiveNamesParamsResult, error) {
	result := GetMostExpensiveNamesParamsResult{}
	result.Limit = 50

	var page int16 = params.Page
	if page < 0 {
		page = 0
	}

	names, err := ctx.db.GetMostExpensiveNames(ctx, db.GetMostExpensiveNamesParams{
		Limit:  result.Limit,
		Offset: page * result.Limit,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	result.NameRows = names
	result.Count = names[0].Count
	result.Limit = 50
	return &result, nil
}

type GetAuctionHistoryParams struct {
	Page int16  `json:"page"`
	Name string `json:"name"`
}

type GetAuctionHistoryParamsResult struct {
	AuctionHistoryRows []db.AuctionHistoryRow `json:"history"`
	Count              int16                  `json:"count"`
	Limit              int16                  `json:"limit"`
}

func GetAuctionHistoryByName(ctx *Context, params *GetAuctionHistoryParams) (*GetAuctionHistoryParamsResult, error) {
	result := GetAuctionHistoryParamsResult{}
	result.Limit = 50
	var page int16 = params.Page
	if page < 0 {
		page = 0
	}

	auctionRows, err := ctx.db.GetAuctionHistoryByName(ctx, db.GetAuctionHistoryByNameParams{
		Name:   params.Name,
		Limit:  result.Limit,
		Offset: page * result.Limit,
	})

	if len(auctionRows) == 0 {
		return &result, nil
	}
	// result.Count = int16(len(auctionRows))
	result.Count = auctionRows[0].Count

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	result.AuctionHistoryRows = auctionRows
	return &result, nil
}

type GetRecordsParams struct {
	Name string `json:"name"`
	Page int16  `json:"page"`
}

type GetRecordsParamsResult struct {
	Records []db.RecordRow `json:"records"`
	Count   int16          `json:"count"`
	Limit   int16          `json:"limit"`
}

func GetRecordsByName(ctx *Context, params *GetRecordsParams) (*GetRecordsParamsResult, error) {
	sha := sha3.New256()
	sha.Write([]byte(params.Name))
	nameHash := sha.Sum(nil)

	result := GetRecordsParamsResult{}

	result.Limit = 50
	var page int16 = params.Page
	if page < 0 {
		page = 0
	}

	recordRows, err := ctx.db.GetNameRecordHistoryByNameHash(ctx, db.GetNameRecordHistoryByNameHashParams{
		NameHash: nameHash,
		Limit:    result.Limit,
		Offset:   page * result.Limit,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	result.Records = recordRows
	// result.Count = int16(len(recordRows))
	result.Count = recordRows[0].Count
	result.Limit = 50
	return &result, nil
}

type GetNameInfoParams struct {
	Name string `json:"name"`
	// Page int16  `json:"page"`
}

type GetNameInfoParamsResult struct {
	Records []db.RecordRow `json:"records"`
	Count   int16          `json:"count"`
	Limit   int16          `json:"limit"`
}

func GetNameInfoByName(ctx *Context, params *GetNameInfoParams) (*GetNameInfoParamsResult, error) {
	// sha := sha3.New256()
	// sha.Write([]byte(params.Name))
	// nameHash := sha.Sum(nil)
	//
	// result := GetNameInfoParamsResult{}
	//
	// result.Limit = 50
	// var page int16 = params.Page
	// if page < 0 {
	// 	page = 0
	// }
	//
	res, err := ctx.db.CheckReservedName(ctx, types.Bytes(params.Name))
	if err == sql.ErrNoRows {

	}

	recordRows, err := ctx.db.GetNameRecordHistoryByNameHash(ctx, db.GetNameRecordHistoryByNameHashParams{
		NameHash: nameHash,
		Limit:    result.Limit,
		Offset:   page * result.Limit,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	result.Records = recordRows
	// result.Count = int16(len(recordRows))
	result.Count = recordRows[0].Count
	result.Limit = 50
	return &result, nil
}
