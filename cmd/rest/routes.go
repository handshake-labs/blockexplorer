package main

import (
	. "github.com/handshake-labs/blockexplorer/cmd/rest/actions"
)

var routes = map[string]interface{}{
	"/blocks":        GetBlocks,
	"/block":         GetBlockByHeight,
	"/block/txs":     GetTransactionsByBlockHeight,
	"/name":          GetName,
	"/name/bids":     GetNameBids,
	"/name/records":  GetNameRecords,
	"/name/actions":  GetNameActions,
	"/search":        Search,
	"/tx":            GetTransactionByTxid,
	"/mempool":       GetMempoolTxs,
	"/block/height/": GetBlockByHeight,
	"/address":       GetAddressHistory,
	"/address/info":  GetAddressInfo,
	// "/lists/transfer":      GetListTransfers,    //names with the most transfers made
}
