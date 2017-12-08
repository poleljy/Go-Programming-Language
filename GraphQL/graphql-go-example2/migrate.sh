#!/bin/bash
set -e
cd "$(dirname "$0")"/
./migrate -database postgres://postgres:123456@192.168.1.189:5432/graphql?sslmode=disable -path ./migrations $@
