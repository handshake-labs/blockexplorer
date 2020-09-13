package main

import (
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/sha3"
	"golang.org/x/net/idna"
	"log"
	"os"
)

type A struct {
	Pizda json.RawMessage `json:"pizda"`
}

func main() {
	// log.Println("www")
	// x := db.GetMostExpensiveNames()

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
	log.Printf("%+v", x)
	log.Println(err2)

	txid := types.Bytes("19350b99af07048a497018d76a5ff08b36e08454a71c98aa1b11c967ecce9a26")
	log.Println(txid)

	txid2, _ := hex.DecodeString("19350b99af07048a497018d76a5ff08b36e08454a71c98aa1b11c967ecce9a26")
	log.Println(types.Bytes(txid2))

	xx, err9 := q.GetTransactionByTxid(context.Background(), txid2)
	log.Println("wwwwwwwwwwwwwwwwwww")
	log.Println(xx)
	log.Println(err9)

	a := types.Bytes("4c7ac3f47d4ba73bb289f06aaf7e672967a165dcfc5fdcb5cf7bec599ba6fff5")
	b := types.Bytes("zzzzыы")
	c := "zcxvы"
	log.Println(a)
	log.Println(len(a))
	log.Println(b)
	log.Println(len(b))
	log.Println(c)
	log.Println(len(c))
	name, err := idna.ToASCII(c)
	if err != nil {
		log.Println(err)
	}
	log.Println(name)

}
