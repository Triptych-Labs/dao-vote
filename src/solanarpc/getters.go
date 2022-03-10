package solanarpc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	ag_binary "github.com/gagliardetto/binary"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptychlabs.io/dao/v2/src/cluster"
	"triptychlabs.io/dao/v2/src/generated/auth"
	"triptychlabs.io/dao/v2/src/generated/dao"
)

type AccountMeta struct {
	Account struct {
		Data []string `json:"data"`
	} `json:"account"`
}
type genericAccount struct {
	Result []AccountMeta `json:"result"`
}

type tokenAccountMetaValue struct {
	Account struct {
		Data struct {
			Parsed struct {
				Info struct {
					Mint        string `json:"mint"`
					Owner       string `json:"owner"`
					TokenAmount struct {
						UIAmount float64 `json:"uiAmount"`
					} `json:"tokenAmount"`
				} `json:"info"`
			} `json:"parsed"`
		} `json:"data"`
	} `json:"account"`
	Pubkey string `json:"pubkey"`
}
type tokenAccountMetaResult struct {
	Value []tokenAccountMetaValue `json:"value"`
}

type tokenAccountMeta struct {
	Result tokenAccountMetaResult `json:"result"`
}

type Signatures struct {
	Result []struct {
		BlockTime int64  `json:"blockTime"`
		Signature string `json:"signature"`
	} `json:"result"`
}

type Transaction struct {
	Result struct {
		Transaction struct {
			Message struct {
				RecentBlockhash string `json:"recentBlockhash"`
				AccountKeys     []struct {
					Pubkey string `json:"pubkey"`
				} `json:"accountKeys"`
			} `json:"message"`
		} `json:"transaction"`
	} `json:"result"`
}

func FetchEnrollmentData(enrollment solana.PublicKey) (*auth.Enrollment, error) {
	url := cluster.RPCEndpoint
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
     "jsonrpc": "2.0",
      "id": 1,
      "method": "getAccountInfo",
      "params": [
        "%s",
        {
          "encoding": "base64"
        }
      ]
    }`, enrollment))

	client := http.DefaultClient
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Result interface{} `json:"result"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	var accountInfo rpc.GetAccountInfoResult
	err = json.Unmarshal(func(inp interface{}) []byte {
		b, e := json.Marshal(inp)
		if e != nil {
			return []byte{}
		}
		return b
	}(resp.Result), &accountInfo)
	if err != nil {
		return nil, err
	}
	if accountInfo.Value == nil {
		return nil, err
	}
	decoder := ag_binary.NewBorshDecoder(accountInfo.Value.Data.GetBinary())
	var enrollmentData auth.Enrollment
	err = enrollmentData.UnmarshalWithDecoder(decoder)
	if err != nil {
		return nil, err
	}

	return &enrollmentData, nil

}

func FetchProposalsData(daoProgramId solana.PublicKey, disc string) (*[]dao.Ballot, error) {
	url := cluster.RPCEndpoint
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
       "jsonrpc": "2.0",
       "id": 1,
       "method": "getProgramAccounts",
       "params": [
           "%s",
           {
               "encoding": "base64",
               "filters": [
                   {
                       "dataSize": 128
                   },
                   {
                       "memcmp": {
                           "offset": 0,
                           "bytes": "%s"
                       }
                   }
               ]
           }
       ]
    }`, daoProgramId, disc))

	client := http.DefaultClient
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Result interface{} `json:"result"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	var accountsInfo genericAccount
	err = json.Unmarshal(func(inp interface{}) []byte {
		b, e := json.Marshal(inp)
		if e != nil {
			return []byte{}
		}
		return b
	}(resp), &accountsInfo)
	if err != nil {
		return nil, err
	}
	proposals := make([]dao.Ballot, 0)
	for _, accountInfo := range accountsInfo.Result {
		b64, err := base64.StdEncoding.DecodeString(accountInfo.Account.Data[0])
		if err != nil {
			return nil, err
		}
		decoder := ag_binary.NewBorshDecoder(b64)
		var proposal dao.Ballot
		err = proposal.UnmarshalWithDecoder(decoder)
		if err != nil {
			return nil, err
		}
		proposals = append(
			proposals,
			proposal,
		)
	}

	return &proposals, nil

}
