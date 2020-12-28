package main

import (
	"context"
	"database/sql"
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	"os"
	// "github.com/handshake-labs/blockexplorer/rest/actions"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/sha3"
	"log"
)

func nameHash(name string) (types.Bytes, error) {
	sha := sha3.New256()
	_, err := sha.Write([]byte(name))
	if err != nil {
		return nil, err
	}
	return types.Bytes(sha.Sum(nil)), nil
}

func modulo(x []byte, j int) int {
	var m int
	for i := 0; i < len(x); i++ {
		m <<= 8
		m += (int(x[i]) & 0xff)
		m %= j
	}
	return m
}

func ReleaseBlock(name string) int {
	hash, _ := nameHash(name)
	w := modulo(hash, 52)*144*7 + 2016
	log.Println(w)
	return w
}

func main() {
	pg, _ := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	q := db.New(pg)
	z, _ := nameHash("js")
	// q.GetLastNameBlockHeightByActionAndHash
	params := db.GetLastNameBlockHeightByActionAndHashParams{db.CovenantAction("OPEN"), &z}
	// params := q.GetLastNameBlockHeightByActionByHash{db.CovenantAction("OPEN"), &z}
	c, t := q.GetLastNameBlockHeightByActionAndHash(context.Background(), params)
	// GetLastNameBlockHeightByActionAndHash
	log.Println(c)
	log.Println(t)
	z, _ = nameHash("stalin")
	params = db.GetLastNameBlockHeightByActionAndHashParams{db.CovenantAction("OPEN"), &z}
	c, t = q.GetLastNameBlockHeightByActionAndHash(context.Background(), params)
	log.Println("%+v", c)
	log.Println(t)
	// nc := node.NewClient(os.Getenv("NODE_API_ORIGIN"), os.Getenv("NODE_API_KEY"))

	sha := sha3.New256()
	sha.Write([]byte("js"))
	ReleaseBlock("js")
	ReleaseBlock("stalin")

}
