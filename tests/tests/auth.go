package tests

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/authl2/events"
	"triptychlabs.io/dao/v2/src/generated/auth"
	"triptychlabs.io/dao/v2/src/keys"
	"triptychlabs.io/dao/v2/src/solanarpc"
)

func Login() {
	user := keys.GetProvider(1)

	ix := auth.NewHolderRegisterInstructionBuilder().
		SetAccessCode("IxfteXJFUe2sxK5d0hhZRFoLIbNbCxWa").
		SetHolderAccount(user.PublicKey()).
		SetSystemProgramAccount(solana.SystemProgramID)

	err := ix.Validate()
	if err != nil {
		panic(err)
	}

	SendTx(
		"auth",
		append(
			make([]solana.Instruction, 0),
			ix.Build(),
		),
		append(
			make([]solana.PrivateKey, 0),
			user,
		),
		user.PublicKey(),
	)

}

/*
   Here we need to simulate the operation of authenticating with the `auth` program
*/
func Integration() {
	auth.SetProgramID(solana.MustPublicKeyFromBase58("9LS5eDSs36coxwkfhvUyWn7ejvVYnKAebsHTuLmky4aT"))

	/*
		    Login()

			fmt.Println("Sle")
			time.Sleep(5 * time.Second)
	*/

	user := solana.MustPublicKeyFromBase58("4sCkw32Lpo4rCu96ZcGYjnB6xkFcxWjVCpGQ82jTx1RH")
	hash := md5.New()

	io.WriteString(hash, user.String())
	hexhash := fmt.Sprintf("%x", hash.Sum(nil))
	enrollment, _ := events.GetEnrollment(hexhash)

	data, err := solanarpc.FetchEnrollmentData(enrollment)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
