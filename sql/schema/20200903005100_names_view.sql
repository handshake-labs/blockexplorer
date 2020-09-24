-- +goose Up
CREATE MATERIALIZED VIEW names AS
SELECT
max(blocks.height) AS open_height,
bids.covenant_name_hash,
bids.covenant_name AS name,
max(lockups.value) AS max_lockup,
sum(lockups.value) AS sum_lockup,
max(reveals.value) AS max_revealed,
sum(reveals.value) AS sum_revealed,
count(distinct(reveals.txid, reveals.index)) AS bidcount
FROM transactions, blocks, tx_outputs lockups, tx_outputs opens, tx_outputs reveals LEFT OUTER JOIN
tx_outputs bids ON (bids.covenant_name_hash = reveals.covenant_name_hash) WHERE
transactions.block_hash IS NOT NULL AND
lockups.covenant_name_hash = bids.covenant_name_hash AND
reveals.covenant_action = 'REVEAL' AND
bids.covenant_action='BID' AND
opens.covenant_name_hash = bids.covenant_name_hash AND --covenant_name = $1 AND
opens.covenant_action = 'OPEN' AND
opens.txid = transactions.txid AND
blocks.hash = transactions.block_hash
GROUP BY
bids.covenant_name, bids.covenant_name_hash;

CREATE UNIQUE INDEX names_hash_index ON names (name);

-- +goose Down
DROP MATERIALIZED VIEW names;
