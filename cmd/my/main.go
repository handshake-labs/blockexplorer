package main

import (
	"context"
	"database/sql"
	// "encoding/hex"
	// "encoding/json"
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	// "github.com/handshake-labs/blockexplorer/rest/actions"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/sha3"
	// "golang.org/x/net/idna"
	"log"
	"os"
)

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
	// sum := sha.Sum(nil)
	// log.Print(sum)
	rr, err99 := q.GetAuctionHistoryByName(context.Background(), db.GetAuctionHistoryByNameParams{Name: "ximik", Offset: 0, Limit: 50})
	log.Println(rr)
	log.Println(err99)

	check, err0 := q.CheckReservedName(context.Background(), types.Bytes("ximik"))
	if err0 == sql.ErrNoRows {
		log.Printf("yoyoy")

	}
	log.Printf("%+v", check)

	// a := types.Bytes("4c7ac3f47d4ba73bb289f06aaf7e672967a165dcfc5fdcb5cf7bec599ba6fff5")
	// b := types.Bytes("zzzzыы")
	// c := "zcxvы"
	// log.Println(a)
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
