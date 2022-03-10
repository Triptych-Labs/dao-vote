package integrations

import (
	"crypto/md5"
	"fmt"
	"io"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/authl2/events"
	"triptychlabs.io/dao/v2/src/cryptog"
	"triptychlabs.io/dao/v2/src/generated/auth"
	"triptychlabs.io/dao/v2/src/solanarpc"
	"triptychlabs.io/dao/v2/wasm/utils"
)

func Certify(this js.Value, args []js.Value) interface{} {
	user := solana.MustPublicKeyFromBase58(args[0].String())
	programKey := args[1].String()
	accessCode := args[2].String()

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
			auth.SetProgramID(solana.MustPublicKeyFromBase58(*programID))

			hash := md5.New()
			io.WriteString(hash, user.String())
			hexhash := fmt.Sprintf("%x", hash.Sum(nil))
			enrollment, _ := events.GetEnrollment(hexhash)

			data, err := solanarpc.FetchEnrollmentData(enrollment)
			if err != nil {
				resolve.Invoke(false)
			}

			hash = md5.New()
			io.WriteString(hash, accessCode)
			accessHash := fmt.Sprintf("%x", hash.Sum(nil))
			if data.AccessCode != accessHash {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}

			resolve.Invoke(true)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}
