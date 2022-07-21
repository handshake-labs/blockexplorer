package actions

import (
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/jinzhu/copier"
)

type GetAddressHistoryParams struct {
	Address string `json:"address"`
	Limit   int8   `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetAddressHistoryResult struct {
	History []HistoryEntry `json:"history"`
}

func GetAddressHistory(ctx *Context, params *GetAddressHistoryParams) (*GetAddressHistoryResult, error) {
	param := db.GetTxOutputsByAddressParams{}
	copier.Copy(&param, &params)
	history, err := ctx.db.GetTxOutputsByAddress(ctx, param)
	if err != nil {
		return nil, err
	}
	result := GetAddressHistoryResult{}
	for _, dbEntry := range history {
		var outEntry HistoryEntry
		copier.Copy(&outEntry.TxOutput, &dbEntry)
		copier.Copy(&outEntry.TxInput, &dbEntry)
		result.History = append(result.History, outEntry)
	}
	return &result, nil
}

type GetAddressInfoParams struct {
	Address string `json:"address"`
}

func GetAddressInfo(ctx *Context, params *GetAddressInfoParams) (*db.GetAddressInfoRow, error) {
	info, err := ctx.db.GetAddressInfo(ctx, params.Address)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
