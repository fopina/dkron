#!/bin/bash

cd $(dirname $0)

DKRON_VERSION=0.11.1

rm -fr dist

docker build -t tmp_dkron_build .

docker run --rm -v $(pwd)/dist:/go/bin/ tmp_dkron_build bash -c "GOOS=darwin GOARCH=amd64 go install -ldflags '-s -w' ./... && GOOS=linux GOARCH=amd64 go install -ldflags '-s -w' ./..."

cd dist

tar zvcpf ../dkron_${DKRON_VERSION}_linux_amd64.tar.gz dkron*
cd darwin_amd64
tar zvcpf ../../dkron_${DKRON_VERSION}_darwin_amd64.tar.gz dkron*
