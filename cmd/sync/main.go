package main

import (
	"database/sql"
	"os"
	"time"

	"log"

	"github.com/handshake-labs/blockexplorer/pkg/node"

	_ "github.com/lib/pq"
)

func main() {
	pg, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatalln(err)
	}
	nc := node.NewClient(os.Getenv("NODE_API_ORIGIN"), os.Getenv("NODE_API_KEY"))
	for {
		if err := syncMempool(pg, nc); err != nil {
			log.Println(err)
			time.Sleep(time.Second)
		}
		if err := syncBlocks(pg, nc); err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		time.Sleep(os.Getenv("UPDATE_DB_INTERVAL") * time.Second)
	}
}
