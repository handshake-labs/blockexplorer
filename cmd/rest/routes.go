package main

import (
	. "github.com/handshake-labs/blockexplorer/cmd/rest/actions"
)

var routes = map[string]interface{}{
	"/blocks":        GetBlocks,
	"/block":         GetBlockByHeight,
	"/block/txs":     GetTransactionsByBlockHeight,
	"/name":          GetName,
	"/name/bids":     GetNameBidsByHash,
	"/name/records":  GetNameRecordsByHash,
	"/search":        Search,
	"/tx":            GetTransactionByTxid,
	"/mempool":       GetMempoolTxs,
	"/block/height/": GetBlockByHeight,
	// "/lists/transfer":      GetListTransfers,    //names with the most transfers made
}
