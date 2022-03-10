package integrations

import (
	"fmt"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/cryptog"
	"triptychlabs.io/dao/v2/wasm/utils"
)

func GetProgramID(this js.Value, args []js.Value) interface{} {
	programKey := args[0].String()

	fmt.Println(programKey, utils.AuthProgramID)
	programID := cryptog.Decrypt(programKey, utils.AuthProgramID)

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

			programId := solana.MustPublicKeyFromBase58(*programID).Bytes()

			dst := js.Global().Get("Uint8Array").New(len(programId))
			js.CopyBytesToJS(dst, programId)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}
