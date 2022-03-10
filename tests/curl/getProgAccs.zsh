#!/bin/zsh

curl http://api.devnet.solana.com -X POST -H "Content-Type: application/json" -d '
  {
    "jsonrpc": "2.0",
    "id": 1,
    "method": "getProgramAccounts",
    "params": [
      "E18jUpqrxp8w4u556G4CcE1jHW7iAvt7i6JH7SkGjaD8",
      {
        "encoding": "base64",
        "filters": [
          {
            "memcmp": {
              "offset": 0,
              "bytes": "hWys35ZYZAY"
            }
          }
        ]
      }
    ]
  }'
