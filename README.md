# Overview 

Backend for hnsnetwork.com. It consists of:

- hsd node, which has additional rpc method for full mempool, [link](https://github.com/handshake-labs/hsd/tree/hnsnetwork)
- postgresql 
- sync process for syncing data from hsd to postgresql
- rest process which is the backend itself

# Steps

## docker

For local testing you may start PostgreSQL and HSD node containers.

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
go run cmd/sync/*
```

## rest


```
go run cmd/rest/*
```

## go2ts

Converts go types into typescript types that are used at the frontend.

`go run -tags typescript github.com/handshake-labs/blockexplorer/cmd/rest > ../<frontend dir>/src/api.ts`


## Dependencies

`go mod download`

`go mod vendor`

## Docker builds

sync
```
docker build -t sync:blockexplorer -f Dockerfile.sync .
```

rest
```
docker build -t rest:blockexplorer -f Dockerfile.rest .
```

Be aware of .dockerignore which should differ for `docker-compose` and for `docker build`.


## Additonal

Showing addresses with a lot of inputs/outputs was in production at cloud, postgresql `enable_hashjoin = off` helped.

Feel free to reach us out at https://t.me/hnsnetwork.

