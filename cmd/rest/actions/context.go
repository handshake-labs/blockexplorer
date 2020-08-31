package actions

import (
	"context"
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"net/http"
)

type Context struct {
	context.Context
	db *db.Queries
}

func NewContext(db *db.Queries, r *http.Request) *Context {
	return &Context{r.Context(), db}
}
