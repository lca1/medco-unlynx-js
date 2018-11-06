#!/usr/bin/env bash
set -Eeuo pipefail
# script that generates the javascript file from the Go code
# requires docker and docker-compose

SCRIPT_FOLDER="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

pushd "$SCRIPT_FOLDER/build"
docker-compose up
popd
