package main

import (
  "log"
	"github.com/handshake-labs/blockexplorer/pkg/db"
	// "github.com/handshake-labs/blockexplorer/pkg/types"
	"database/sql"
  "os"
	"context"
	_ "github.com/lib/pq"
)

func main() {
  // log.Println("www")
  // x := db.GetMostExpensiveNames()

	pg, err1 := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
  log.Println(err1)
	q := db.New(pg)
	// x, err2 := q.GetMostExpensiveNames(context.Background())
  x, err2 := q.GetAuctionHistoryByName(context.Background(), "js")
  log.Println(err2)
	
  log.Printf("%+v",x)

}
