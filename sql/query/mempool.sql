-- name: DeleteMempool :exec
DELETE FROM transactions WHERE block_hash IS NULL;
