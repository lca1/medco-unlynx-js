#!/usr/bin/env bash
set -Eeuo pipefail

pushd "$GOPATH"/src/github.com/lca1/medco-unlynx-js
go get -d ./javascriptLibrary/...
gopherjs build -m ./javascriptLibrary/ -o medco-unlynx-js.js
popd
