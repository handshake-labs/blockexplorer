-- +goose Up
-- +goose StatementBegin
CREATE VIEW names AS
SELECT
max(blocks.height) AS open_height,
bids.covenant_name_hash,
bids.covenant_name AS name,
max(lockups.value) AS max_lockup,
max(reveals.value) AS max_revealed,
count(distinct(reveals.txid, reveals.index)) AS bidcount
FROM transactions, blocks, tx_outputs lockups, tx_outputs opens, tx_outputs reveals LEFT OUTER JOIN
tx_outputs bids ON (bids.covenant_name_hash = reveals.covenant_name_hash) WHERE
lockups.covenant_name_hash = bids.covenant_name_hash AND
reveals.covenant_action = 'REVEAL' AND
bids.covenant_action='BID' AND
opens.covenant_name_hash = bids.covenant_name_hash AND --covenant_name = $1 AND
opens.covenant_action = 'OPEN' AND
opens.txid = transactions.txid AND
blocks.hash = transactions.block_hash
GROUP BY
bids.covenant_name, bids.covenant_name_hash;


CREATE VIEW records AS
SELECT
blocks.height,
covenant_record_data,
covenant_name_hash
FROM tx_outputs, blocks WHERE
blocks.hash = tx_outputs.block_hash AND
covenant_record_data IS NOT NULL
GROUP BY covenant_action, blocks.height, covenant_record_data, covenant_name_hash;


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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW names;
DROP VIEW records;
DROP VIEW auctions;
-- +goose StatementEnd
