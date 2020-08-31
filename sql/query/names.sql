-- name: GetNameOpeningBlock :one
SELECT (height) FROM tx_outputs, transactions, blocks WHERE covenant_name = $1 AND covenant_action = 'OPEN' AND
tx_outputs.txid = transactions.txid AND blocks.hash = block_hash;


-- get auction history of a name with reveals, input parameter - hash of the name 
-- name: GetAuctionHistory :many
SELECT height, transactions.txid, A.value AS lockup, B.value AS reveal, A.covenant_action FROM transactions, blocks,
tx_outputs A LEFT OUTER JOIN tx_outputs B ON (A.covenant_name_hash = B.covenant_name_hash AND B.covenant_action =
  'REVEAL' AND A.address = B.address AND A.covenant_action = 'BID') WHERE A.txid=transactions.txid AND
transactions.block_hash = blocks.hash AND A.covenant_name_hash=$1 ORDER BY height DESC;


-- domain top list by lockup with count and revealed values 
-- name: GetTopList :many
SELECT bids.covenant_name AS name, max(lockups.value) AS max_lockup, max(reveals.value) AS max_revealed,
count(distinct(reveals.txid, reveals.index_out)) AS bidcount FROM tx_outputs lockups, tx_outputs reveals LEFT OUTER JOIN
tx_outputs bids ON (reveals.covenant_name_hash = bids.covenant_name_hash) WHERE lockups.covenant_name_hash =
bids.covenant_name_hash AND reveals.covenant_action = 'REVEAL' AND (bids.covenant_action='BID') GROUP BY
bids.covenant_name ORDER BY lockup DESC;
