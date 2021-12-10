# Overview 

Backend for hnsnetwork.com. It consists of:

- hsd node, which has additional rpc method for full mempool, [link](https://github.com/handshake-labs/hsd/tree/hnsnetwork)
- postgresql 
- sync process for syncing data from hsd to postgresql
- rest process which is the backend itself

## Steps

### Dependencies

`go mod download`

`go mod vendor`

### Docker

For local testing you may start PostgreSQL and HSD node containers.

```
docker-compose up
```

For production you should have your own working postgresql.

### Environment

Load environment variables into current shell session

```
. ./env
```

### Goose
Run SQL migrations

```
goose -dir sql/schema postgres $POSTGRES_URI up
```
### SQLC
Generate types and methods from SQL code

```
sqlc generate
```
### Sync

Now we can run the sync process which will synchronize the database.

```
go run cmd/sync/*
```
### Rest

And now you can start rest API.

```
go run github.com/handshake-labs/blockexplorer/cmd/rest
```

### go2ts

Converts go types into typescript types that are used at the frontend.

`go run -tags typescript github.com/handshake-labs/blockexplorer/cmd/rest > ../<frontend dir>/src/api.ts`

### Docker builds

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

Feel free to reach us [here](https://t.me/hnsnetwork).

