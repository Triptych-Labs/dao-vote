#!/bin/zsh

GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o qrcode.wasm
