-- name: GetTxOutputsByTxid :many
SELECT tx_outputs.*, namehash.covenant_name FROM tx_outputs, namehash WHERE tx_outputs.covenant_name_hash = namehash.covenant_name_hash AND tx_outputs.txid = $1
ORDER BY index;

-- name: InsertTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name, covenant_bid_hash, covenant_nonce, covenant_record_data, covenant_block_hash, covenant_version, covenant_address, covenant_claim_height, covenant_renewal_count) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);

-- name: InsertNONETxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action) VALUES ($1, $2, $3, $4, 'NONE');

-- name: InsertCLAIMTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name) VALUES ($1, $2, $3, $4, $5, 'CLAIM', $6, $7);

-- name: InsertOPENTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name) VALUES ($1, $2, $3, $4, $5, 'OPEN', $6, $7);

-- name: InsertBIDTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name, covenant_bid_hash) VALUES ($1, $2, $3, $4, $5, 'BID', $6, $7, $8);

-- name: InsertREVEALTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_nonce) VALUES ($1, $2, $3, $4, $5, 'REVEAL', $6, $7);

-- name: InsertREDEEMTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height) VALUES ($1, $2, $3, $4, $5, 'REDEEM', $6);

-- name: InsertREGISTERTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_record_data, covenant_block_hash) VALUES ($1, $2, $3, $4, $5, 'REGISTER', $6, $7, $8);

-- name: InsertUPDATETxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_record_data) VALUES ($1, $2, $3, $4, $5, 'UPDATE', $6, $7);

-- name: InsertRENEWTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_block_hash) VALUES ($1, $2, $3, $4, $5, 'RENEW', $6, $7);

-- name: InsertTRANSFERTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_version, covenant_address) VALUES ($1, $2, $3, $4, $5, 'TRANSFER', $6, $7, $8);

-- name: InsertFINALIZETxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name, covenant_claim_height, covenant_renewal_count, covenant_block_hash) VALUES ($1, $2, $3, $4, $5, 'FINALIZE', $6, $7, $8, $9, $10);

-- name: InsertREVOKETxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height) VALUES ($1, $2, $3, $4, $5, 'REVOKE', $6);

