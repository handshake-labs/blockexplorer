-- +goose Up
CREATE MATERIALIZED VIEW records AS
SELECT
blocks.height,
covenant_record_data,
covenant_name_hash,
tx_outputs.txid,
tx_outputs.index
FROM tx_outputs, transactions, blocks WHERE
transactions.txid = tx_outputs.txid AND
transactions.block_hash = blocks.hash AND
covenant_record_data IS NOT NULL
GROUP BY covenant_action, blocks.height, covenant_record_data, covenant_name_hash, tx_outputs.txid, tx_outputs.index ;

CREATE UNIQUE INDEX records_index ON records (txid, index);

-- +goose Down
DROP MATERIALIZED VIEW records;
