-- get auction history of a name with reveals, input parameter - hash of the name 
-- name: GetAuctionHistoryByNameHash :many
SELECT * FROM auctions WHERE covenant_name_hash=$1 ORDER BY height DESC;

-- get auction history of a name with reveals, input parameter - the name 
-- name: GetAuctionHistoryByName :many
SELECT * FROM auctions WHERE covenant_name=$1 ORDER BY height DESC;

-- name: GetMostExpensiveNames :many
SELECT * FROM names ORDER BY max_lockup desc;

-- name: GetNameRecordHistoryByNameHash :many 
SELECT height, covenant_record_data FROM records WHERE covenant_name_hash = $1 ORDER BY height DESC;
