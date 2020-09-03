// +build !typescript

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
	"github.com/handshake-labs/blockexplorer/pkg/db"

	_ "github.com/lib/pq"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	pg, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatalln(err)
	}
	q := db.New(pg)
  log.Println("asf")

	handlers := make(map[string]http.HandlerFunc, 0)
	for path, function := range routes {
		handlers[path] = actions.NewAction(function).BuildHandlerFunc(q)
	}

	srv := &http.Server{
		Addr: os.Getenv("REST_ADDR"),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			path := r.URL.Path
			if path == "/" {
				w.WriteHeader(http.StatusOK)
				return
			}
      log.Printf("+v%", path)
      log.Printf("+v%", w)
      log.Printf("+v%", r)
			if handler, ok := handlers[path]; ok {
        log.Println("bbb")
				handler(w, r)
				return
			}
      log.Println("aaa")
			w.WriteHeader(http.StatusNotFound)
		}),
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
