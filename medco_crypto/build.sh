#!/usr/bin/env bash

tsc && tsc -m es6 --outDir lib-esm
npx webpack --config webpack.config.js