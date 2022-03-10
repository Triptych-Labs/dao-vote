#!/bin/zsh

GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o dao.wasm
scp dao.wasm ddigiacomo@10.145:/noshit/triptych_labs/homepage/public
scp dao.wasm ddigiacomo@10.145:/noshit/triptych_labs/homepage/public/dao
