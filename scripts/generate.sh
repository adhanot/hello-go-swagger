#!/bin/bash

echo "Generate ENT"

set -e

find ../pkg/ent/ ! -name 'generate.go' ! -name "ent"  -exec rm -rf {} +

go generate ../pkg/ent/

echo "Generate API"

cd ..

/usr/bin/swagger generate server -f /home/salty/go/src/hello/gen/api/swagger.yaml --default-scheme http

/usr/bin/swagger generate client -f /home/salty/go/src/hello/gen/api/swagger.yaml --default-scheme http

docker-compose stop
docker-compose up

