// Package cluster - points rpcs
package cluster

import "fmt"

const MainnetEndpoint = "falling-empty-tree.solana-mainnet.quiknode.pro/934c447edf1a84cdee3f3e392d5e5f33f2b9bb48/"
const DevnetEndpoint = "sparkling-dark-shadow.solana-devnet.quiknode.pro/0e9964e4d70fe7f856e7d03bc7e41dc6a2b84452/"

const Endpoint = DevnetEndpoint

var RPCEndpoint = fmt.Sprint("https://", Endpoint)
var WSEndpoint = fmt.Sprint("wss://", Endpoint)

