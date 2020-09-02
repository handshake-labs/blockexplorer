-- +goose Up
CREATE VIEW auctions AS
SELECT
height,
transactions.txid,
lockups.covenant_name,
lockups.value AS lockup,
reveals.value AS reveal,
lockups.covenant_action,
lockups.covenant_record_data,
lockups.covenant_name_hash
FROM transactions, blocks, tx_outputs lockups LEFT OUTER JOIN tx_outputs reveals ON
(lockups.covenant_name_hash = reveals.covenant_name_hash AND
reveals.covenant_action = 'REVEAL' AND
lockups.address = reveals.address AND
lockups.covenant_action = 'BID') WHERE
lockups.txid=transactions.txid AND
transactions.block_hash = blocks.hash;

-- +goose Down
DROP VIEW auctions;
