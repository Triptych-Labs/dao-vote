package integrations

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/cryptog"
	"triptychlabs.io/dao/v2/src/solanarpc"
	"triptychlabs.io/dao/v2/wasm/utils"
)

func FetchProposals(this js.Value, args []js.Value) interface{} {
	arg0 := args[0].String()
	programID := cryptog.Decrypt(arg0, utils.DAOProgramID)
	fmt.Println("program id", *programID)

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			if programID == nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}

			disc := cryptog.Decrypt(arg0, utils.ProposalDisc)
			fmt.Println("disc", *disc)
			if disc == nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}
			data, err := solanarpc.FetchProposalsData(solana.MustPublicKeyFromBase58(*programID), *disc)
			fmt.Println(data)
			if err != nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}

			type obfuscate struct {
				ProposalNum         int    `json:"proposalNum"`
				ProposalType        string `json:"proposalType"`
				ProposalDescription string `json:"proposalDescription"`
			}

			var proposals []obfuscate
			for i, proposal := range *data {
				proposals = append(
					proposals,
					obfuscate{
						i,
						proposal.Name,
						proposal.Description,
					},
				)

			}

			payload, err := json.Marshal(proposals)
			if err != nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}

			resolve.Invoke(string(payload))
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}
