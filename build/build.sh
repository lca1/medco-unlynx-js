#!/usr/bin/env bash
set -Eeuo pipefail

pushd "$GOPATH"/src/github.com/lca1/medco-unlynx-js
go get -v -d ./javascriptLibrary/...
gopherjs build -m ./javascriptLibrary/ -o medco-unlynx-js.js
chmod a+rw medco-unlynx-js.js medco-unlynx-js.js.map
popd
