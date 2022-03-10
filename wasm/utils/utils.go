package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/cluster"
)

func TxToBase64(
	instructions []solana.Instruction,
) (string, error) {
	blockhash, err := GetBlockhash()
	if err != nil {
		return "", err
	}

	tx, err := solana.NewTransaction(
		instructions,
		blockhash,
	)
	if err != nil {
		return "", err
	}

	return tx.MustToBase64(), nil
}

func GetBlockhash() (solana.Hash, error) {
	payload := strings.NewReader(`
  {"jsonrpc":"2.0","id":1, "method":"getRecentBlockhash"}
`)

	res, err := http.DefaultClient.Post(cluster.RPCEndpoint, "application/json", payload)
	if err != nil {
		return solana.Hash{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return solana.Hash{}, err
	}
	type blockhashRespT struct {
		Result struct {
			Value struct {
				Blockhash string `json:"blockhash"`
			} `json:"value"`
		} `json:"result"`
	}

	var blockhashResp blockhashRespT

	err = json.Unmarshal(body, &blockhashResp)
	if err != nil {
		fmt.Println(err)
		return solana.Hash{}, err
	}

	return solana.MustHashFromBase58(blockhashResp.Result.Value.Blockhash), nil
}
