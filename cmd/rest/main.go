package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/handshake-labs/blockexplorer/cmd/rest/actions"
	"github.com/handshake-labs/blockexplorer/cmd/rest/handler"
	"github.com/handshake-labs/blockexplorer/pkg/db"

	_ "github.com/lib/pq"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	pg, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatalln(err)
	}

	acs := map[string]interface{}{
		"/block":     actions.GetBlockByHeight,
		"/block/txs": actions.GetTransactionsByBlockHash,
	}

	srv := &http.Server{
		Addr:    os.Getenv("REST_ADDR"),
		Handler: handler.NewHandler(db.New(pg), acs),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
