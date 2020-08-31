# Tools

## docker

Start PostgreSQL and HSD node containers.

```
docker-compose up
```

## env

Load environment variables into current shell session

```
. ./env
```

## goose

Run SQL migrations

```
goose -dir sql/schema postgres $POSTGRES_URI up
```

## sqlc

Generate types and methods from SQL code

```
sqlc generate
```

## sync

Synchronize the database
```
go run services/blocks_sync/*.go
```
