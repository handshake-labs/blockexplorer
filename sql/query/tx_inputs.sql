-- name: InsertTxInput :exec
INSERT INTO tx_inputs (tx_hash, index, hash_prevout, index_prevout, sequence)
VALUES ($1, $2, $3, $4, $5);
