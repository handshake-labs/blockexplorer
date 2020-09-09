package actions

import (
	"database/sql"

	"golang.org/x/crypto/sha3"
	// "github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/handshake-labs/blockexplorer/pkg/db"
	// "github.com/jinzhu/copier"
	// "log"
)

type GetListExpensiveParams struct {
	Page int16 `json:"page"`
}

type GetListExpensiveResult struct {
	NameRows []db.NameRow `json:"names"`
	Count    int32        `json:"count"`
	Limit    int16        `json:"limit"`
}

func GetListExpensive(ctx *Context, params *GetListExpensiveParams) (*GetListExpensiveResult, error) {
	result := GetListExpensiveResult{}
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

type GetListLockupVolumeParams struct {
	Page int16 `json:"page"`
}

type GetListLockupVolumeResult struct {
	NameRows []db.NameVolumeRow `json:"names"`
	Count    int32              `json:"count"`
	Limit    int16              `json:"limit"`
}

func GetListLockupVolume(ctx *Context, params *GetListLockupVolumeParams) (*GetListLockupVolumeResult, error) {
	result := GetListLockupVolumeResult{}
	result.Limit = 50

	var page int16 = params.Page
	if page < 0 {
		page = 0
	}

	names, err := ctx.db.GetMostLockupVolumeNames(ctx, db.GetMostLockupVolumeNamesParams{
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

type GetListRevealVolumeParams struct {
	Page int16 `json:"page"`
}

type GetListRevealVolumeResult struct {
	NameRows []db.NameVolumeRow `json:"names"`
	Count    int32              `json:"count"`
	Limit    int16              `json:"limit"`
}

func GetListRevealVolume(ctx *Context, params *GetListRevealVolumeParams) (*GetListRevealVolumeResult, error) {
	result := GetListRevealVolumeResult{}
	result.Limit = 50

	var page int16 = params.Page
	if page < 0 {
		page = 0
	}

	names, err := ctx.db.GetMostRevealVolumeNames(ctx, db.GetMostRevealVolumeNamesParams{
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

type GetListBidsParams struct {
	Page int16 `json:"page"`
}

type GetListBidsResult struct {
	NameRows []db.NameVolumeRow `json:"names"`
	Count    int32              `json:"count"`
	Limit    int16              `json:"limit"`
}

func GetListBids(ctx *Context, params *GetListBidsParams) (*GetListBidsResult, error) {
	result := GetListBidsResult{}
	result.Limit = 50

	var page int16 = params.Page
	if page < 0 {
		page = 0
	}

	names, err := ctx.db.GetMostBidsNames(ctx, db.GetMostBidsNamesParams{
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

// type GetListTransfersParams struct {
// 	Page int16 `json:"page"`
// }
//
// type GetListTransfersResult struct {
// 	NameRows []db.NameRow `json:"names"`
// 	Count    int32        `json:"count"`
// 	Limit    int16        `json:"limit"`
// }
//
// func GetListTransfers(ctx *Context, params *GetListTransfersParams) (*GetListTransfersResult, error) {
// 	result := GetListTransfersResult{}
// 	result.Limit = 50
//
// 	var page int16 = params.Page
// 	if page < 0 {
// 		page = 0
// 	}
//
// 	names, err := ctx.db.GetMostTransfersNames(ctx, db.GetMostTransfersNamesParams{
// 		Limit:  result.Limit,
// 		Offset: page * result.Limit,
// 	})
//
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	result.NameRows = names
// 	result.Count = names[0].Count
// 	result.Limit = 50
// 	return &result, nil
// }
//
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
