package main

import (
	. "github.com/handshake-labs/blockexplorer/cmd/rest/actions"
)

var routes = map[string]interface{}{
	"/block":           GetBlockByHeight,
	"/block/txs":       GetTransactionsByBlockHash,
	"/tx":              GetTransactionByTxid,
	"/lists/expensive": GetTopList,
	"/lists/bids":      GetTopList,
	"/names/records":   GetRecordsByName,
	"/names/auction":   GetAuctionHistoryByName,
	"/names/reserved":  GetAuctionHistoryByName,
	"/names/info":      GetAuctionHistoryByName,
}
