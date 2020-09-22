package db

import (
	"context"
)

const refreshViews = `--name: RefreshViews :exec
REFRESH MATERIALIZED VIEW CONCURRENTLY namehash;
REFRESH MATERIALIZED VIEW CONCURRENTLY names;
REFRESH MATERIALIZED VIEW CONCURRENTLY records;
REFRESH MATERIALIZED VIEW CONCURRENTLY auctions;
`

func (q *Queries) RefreshViews(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, refreshViews)
	return err
}
