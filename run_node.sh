#!/bin/sh

## This is mainly used while development.
fuser -k 8090/tcp || true
HTTP_ENDPOINT=8090 nohup go run cmd/node/main.go &
exit