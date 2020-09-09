-- name: CheckReservedName :one
SELECT *
FROM reserved_names
WHERE name = $1;

-- name: CheckReservedNameHash :one
SELECT *
FROM reserved_names
WHERE name_hash = $1;
