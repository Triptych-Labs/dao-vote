#!/bin/zsh

anchor build
solana program deploy target/deploy/dao.so
solana program deploy target/deploy/auth.so
anchor-go --src target/idl/dao.json
anchor-go --src target/idl/auth.json
cp -r generated src
