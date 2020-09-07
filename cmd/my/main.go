package main

import (
	"context"
	"database/sql"
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/node"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/sha3"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	// log.Println("www")
	// x := db.GetMostExpensiveNames()

	nc := node.NewClient(os.Getenv("NODE_API_ORIGIN"), os.Getenv("NODE_API_KEY"))
	txs, err := nc.GetMempool(context.Background())
	log.Println(txs)
	log.Println(err)
	pg, err1 := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	log.Println(err1)
	q := db.New(pg)
	// x, err2 := q.GetMostExpensiveNames(context.Background())
	// x, err2 := q.GetAuctionHistoryByName(context.Background(), "js")

	sha := sha3.New256()
	sha.Write([]byte("js"))
	sum := sha.Sum(nil)
	// log.Print(sum)

	x, err2 := q.GetNameRecordHistoryByNameHash(context.Background(), db.GetNameRecordHistoryByNameHashParams{NameHash: types.Bytes(sum)})
	// x, err2 := q.GetNameRecordHistoryByNameHash(context.Background(), types.Bytes("7d9ed1a61b37b5a103316e2be8b6db155200dc15b7ef65be0031beba890c93e6"))
	log.Println(err2)

	log.Printf("%+v", x)

}
