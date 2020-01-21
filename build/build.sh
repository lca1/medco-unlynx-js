#!/usr/bin/env bash
set -Eeuo pipefail

pushd /src

# instruct go to compile to js/wasm
export GOOS=js GOARCH=wasm

# get dependencies
go get -d -v ./javascriptLibrary/...

# hack onet to get WASM compilation to work
pushd /go/pkg/mod/go.dedis.ch/onet
chmod u+w -R .
cat > $(find . -maxdepth 4 -name rtime.go) <<EOL
// +build freebsd linux darwin js

package monitor

// Converts microseconds to seconds.
func iiToF(sec int64, usec int64) float64 {
        return float64(sec) + float64(usec)/1000000.0
}

// Returns the system and the user CPU time used by the current process so far.
func getRTime() (tSys, tUsr float64) {
        return iiToF(int64(1), int64(1)), iiToF(int64(1), int64(1))
}
EOL
popd

# generate the wasm and js
go build -o npm/medco-unlynx-js.wasm ./javascriptLibrary/...
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" npm/
if [[ ! -z ${USER_ID} ]]; then
    chown "${USER_ID}" npm/wasm_exec.js npm/medco-unlynx-js.wasm
fi
if [[ ! -z ${GROUP_ID} ]]; then
    chgrp "${GROUP_ID}" npm/wasm_exec.js npm/medco-unlynx-js.wasm
fi
popd
