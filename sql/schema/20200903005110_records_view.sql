-- +goose Up
CREATE MATERIALIZED VIEW records AS
SELECT
blocks.height,
covenant_record_data,
covenant_name_hash,
tx_outputs.txid,
tx_outputs.index
FROM tx_outputs, blocks WHERE
blocks.hash = tx_outputs.block_hash AND
covenant_record_data IS NOT NULL
GROUP BY covenant_action, blocks.height, covenant_record_data, covenant_name_hash, txid, index;

CREATE UNIQUE INDEX records_index ON records (txid, index);

-- +goose Down
DROP MATERIALIZED VIEW records;
