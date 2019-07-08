FROM golang:1.12-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /go/src/github.com/nandaryanizar/golang-webservice-example

CMD go get -v ./... && CGO_ENABLED=0 go build -o ./bin/golang-webservice-example ./cmd/... && ./bin/golang-webservice-example