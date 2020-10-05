-- name: GetTxOutputsByTxid :many
SELECT DISTINCT ON(t1.index)
  t1.*,
  COALESCE(CONVERT_FROM(t2.covenant_name, 'SQL_ASCII'), '')::text AS name
FROM
  tx_outputs t1
  LEFT JOIN tx_outputs t2 ON (t1.covenant_name_hash = t2.covenant_name_hash AND t2.covenant_name IS NOT NULL)
WHERE t1.txid = $1
ORDER BY t1.index;

-- name: InsertTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name, covenant_bid_hash, covenant_nonce, covenant_record_data, covenant_block_hash, covenant_version, covenant_address, covenant_claim_height, covenant_renewal_count)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);
