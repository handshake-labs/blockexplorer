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



TODOS:
- check what happens if the node stops and has to resync from block 0. will the sync and the db be ok?
- parse the record data
- add icann restricted names


Docker build

## Pass dependencies

`go mod download`
`go mod vendor`


## Build

sync
```
docker build -t sync:blockexplorer -f Dockerfile.sync .
```

rest
```
docker build -t rest:blockexplorer -f Dockerfile.rest .
```

Be aware of .dockerignore which should differ for docker-compose and for docker build

