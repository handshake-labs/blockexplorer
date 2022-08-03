# Overview 

Backend for hnsnetwork.com. It consists of:

- hsd node
- postgresql 
- sync process for syncing data from hsd to postgresql
- rest process which is the backend itself

## Steps

### Dependencies

`go mod download`

`go mod vendor`

### Environment

Load environment variables into current shell session

```
. ./env
```

### Goose
Run SQL migrations. You need to install [goose](https://github.com/pressly/goose) first:

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

```
goose -dir sql/schema postgres $POSTGRES_URI up
```
### SQLC
Generate types and methods from SQL code

```
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
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


## Additonal settings

Showing addresses with a lot of inputs/outputs was slow in production at cloud, postgresql `enable_hashjoin = off` helped.

Feel free to reach us [here](https://t.me/hnsnetwork).

### Docker builds

It's possible to use the explorer as docker container.

sync
```
docker build -t sync:blockexplorer -f Dockerfile.sync .
```

rest
```
docker build -t rest:blockexplorer -f Dockerfile.rest .
```

Be aware of .dockerignore which should differ for `docker-compose` and for `docker build`.
