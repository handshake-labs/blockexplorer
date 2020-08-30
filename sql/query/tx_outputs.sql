-- name: InsertTxOutput :exec
INSERT INTO tx_outputs (tx_hash, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name, covenant_bid_hash, covenant_nonce, covenant_record_data, covenant_block_hash, covenant_version, covenant_address, covenant_claim_height, covenant_renewal_count)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);
