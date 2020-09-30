package main

import (
	"context"
	"database/sql"
	"encoding/hex"
	// "encoding/json"
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/node"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	// "github.com/handshake-labs/blockexplorer/rest/actions"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/sha3"
	// "golang.org/x/net/idna"
	"log"
	"os"
)

func main() {
	pg, err1 := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	log.Println(err1)
	q := db.New(pg)
	// x, err2 := q.GetMostExpensiveNames(context.Background())
	// x, err2 := q.GetAuctionHistoryByName(context.Background(), "js")

	nc := node.NewClient(os.Getenv("NODE_API_ORIGIN"), os.Getenv("NODE_API_KEY"))
	txs, _ := nc.GetMempool(context.Background())
	// log.Println(txs)
	azxc := *txs
	qwer := azxc[0]
	log.Printf("%+v", qwer)

	sha := sha3.New256()
	sha.Write([]byte("js"))
	// sum := sha.Sum(nil)
	// log.Print(sum)
	// rr, err99 := q.GetAuctionHistoryByName(context.Background(), db.GetAuctionHistoryByNameParams{Name: "ximik", Offset: 0, Limit: 50})
	// log.Println(rr)
	// log.Println(err99)
	//
	// check, err0 := q.CheckReservedName(context.Background(), types.Bytes("ximik"))
	// if err0 == sql.ErrNoRows {
	// 	log.Printf("yoyoy")
	//
	// }
	// log.Printf("%+v", check)
	//

	// aa, err101 := q.GetTxOutputsByTxid(context.Background(), types.Bytes("82f324aff4d93df901b9c0579d63fc9fa33b63c8a2495905723de974c47581ba"))
	z, _ := hex.DecodeString("82f324aff4d93df901b9c0579d63fc9fa33b63c8a2495905723de974c47581ba")
	// aa, err101 := q.GetTxOutputsByTxid(context.Background(), types.Bytes("42c3b89479ac26d01d199ec551ecf6924ea63ba1e2dd0c5f3d99bc19eeaef2f6"))
	aa, err101 := q.GetTxOutputsByTxid(context.Background(), types.Bytes(z))
	log.Println(aa)
	log.Println(err101)

	namehash, _ := hex.DecodeString("0002e258f7e171410ad572139df53d6c45e9e855c6f148592d78d7b9050225e9")
	log.Println(namehash)
	pizda, adz := q.GetNameByNameHash(context.Background(), types.Bytes(namehash))

	log.Println(adz)
	log.Println(pizda)

	// log.Println(len(a))
	// log.Println(b)
	// log.Println(len(b))
	// log.Println(c)
	// log.Println(len(c))
	// name, err := idna.ToASCII(c)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(name)
	//
}
