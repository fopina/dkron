#!/bin/bash

cd $(dirname $0)

rm -fr dist
mkdir -p dist
cp CHANGELOG.md README.md LICENSE dist/

docker build -t tmp_dkron_build .

docker run --rm -v $(pwd)/dist:/go/bin/ tmp_dkron_build go install -ldflags '-s -w' ./...

cd dist

tar zvcpf ../dkron_0.11.0_darwin_amd64.tar.gz *
