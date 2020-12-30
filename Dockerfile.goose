FROM golang:1.14.3-alpine
RUN apk add gcc g++ git
RUN go get github.com/pressly/goose/cmd/goose

WORKDIR /src
COPY . .

ENTRYPOINT exec /go/bin/goose -dir sql/schema postgres "$POSTGRES_URI" up
