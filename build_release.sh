#!/bin/bash

cd $(dirname $0)

DKRON_VERSION=$(cat dkron/version.go | grep Version | cut -d '"' -f2)

echo "Building ${DKRON_VERSION}"

rm -fr dist

docker build -t tmp_dkron_build .

docker run --rm -v $(pwd)/dist:/go/bin/ tmp_dkron_build bash -c \
	   "CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go install -ldflags '-s -w' ./... && \
	    CGO_ENABLED=0 GOOS=linux go install -ldflags '-s -w' ./..."

cd dist

tar zvcpf ../dkron_${DKRON_VERSION}_linux_amd64.tar.gz dkron*
cd darwin_amd64
tar zvcpf ../../dkron_${DKRON_VERSION}_darwin_amd64.tar.gz dkron*
