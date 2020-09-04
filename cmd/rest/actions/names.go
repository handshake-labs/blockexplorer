package actions

import (
	"database/sql"

	"github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/handshake-labs/blockexplorer/pkg/db"
	// "github.com/jinzhu/copier"
  "log"
)

type GetMostExpensiveNamesParams struct {
	Page      int16       `json:"page"`
}

type GetMostExpensiveNamesParamsResult struct {
	NameRows []db.NameRow `json:"namerows"`
	Count        int16         `json:"count"`
	Limit        int16         `json:"limit"`
}


func GetTopList(ctx *Context, params *GetMostExpensiveNamesParams) (*GetMostExpensiveNamesParamsResult, error) {
  log.Println("yoyo")
	result := GetMostExpensiveNamesParamsResult{}
	names, err := ctx.db.GetMostExpensiveNames(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
  result.NameRows = names
  result.Count = 50
  result.Limit = 50
  return &result, nil
}

type GetAuctionHistoryParams struct {
	Page      int16       `json:"page"`
	Name      string       `json:"name"`
}

type GetAuctionHistoryParamsResult struct {
	AuctionHistoryRows []db.AuctionHistoryRow `json:"history"`
	Count        int16         `json:"count"`
	Limit        int16         `json:"limit"`
}


// func GetAuctionHistory(ctx *Context, params *GetAuctionHistoryParams) (*GetAuctionHistoryParamsResult, error) {
//   log.Println("yoyo")
  
func GetAuctionHistoryByName(ctx *Context, params *GetAuctionHistoryParams) (*GetAuctionHistoryParamsResult, error) {
	result := GetAuctionHistoryParamsResult{}
	auctionRows, err := ctx.db.GetAuctionHistoryByName(ctx, types.Bytes(params.Name))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
  result.AuctionHistoryRows = auctionRows
  result.Count = 50
  result.Limit = 50
  return &result, nil
}


