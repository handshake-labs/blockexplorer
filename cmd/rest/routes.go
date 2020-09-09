package main

import (
	. "github.com/handshake-labs/blockexplorer/cmd/rest/actions"
)

var routes = map[string]interface{}{
	"/block":               GetBlockByHeight,
	"/block/txs":           GetTransactionsByBlockHash,
	"/lists/expensive":     GetListExpensive,    //the most expensive names
	"/lists/lockup_volume": GetListLockupVolume, //names with the most auction volume
	"/lists/reveal_volume": GetListRevealVolume, //names with the most auction volume
	"/lists/bids":          GetListBids,         //names with the most bids in the auction
	// "/lists/transfer":      GetListTransfers,    //names with the most transfers made
	"/names/records": GetRecordsByName,
	"/names/auction": GetAuctionHistoryByName,
	"/search":        Search,
}
