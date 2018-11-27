#!/usr/bin/env bash
set -Eeuo pipefail

# pre-fetch unlynx in the desired version
UNLYNX_REPO="github.com/lca1/unlynx"
UNLYNX_VERSION="v1.2-alpha"

mkdir -p /go/src/${UNLYNX_REPO}
pushd /go/src/${UNLYNX_REPO}
git clone --depth 1 --branch ${UNLYNX_VERSION} https://${UNLYNX_REPO}.git .
popd

pushd "$GOPATH"/src/github.com/lca1/medco-unlynx-js
go get -v -d ./javascriptLibrary/...
gopherjs build -m ./javascriptLibrary/ -o medco-unlynx-js.js
chmod a+rw medco-unlynx-js.js medco-unlynx-js.js.map
popd
