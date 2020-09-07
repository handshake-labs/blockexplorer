package db

import (
	"context"
)

const refreshViews = `--name: RefreshViews :exec
REFRESH MATERIALIZED VIEW names;
`

func (q *Queries) RefreshViews(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, refreshViews)
	return err
}
