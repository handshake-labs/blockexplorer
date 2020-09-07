-- +goose Up
CREATE VIEW records AS
SELECT
blocks.height,
covenant_record_data,
covenant_name_hash
FROM tx_outputs, blocks WHERE
blocks.hash = tx_outputs.block_hash AND
covenant_record_data IS NOT NULL
GROUP BY covenant_action, blocks.height, covenant_record_data, covenant_name_hash;

-- +goose Down
DROP VIEW records;
