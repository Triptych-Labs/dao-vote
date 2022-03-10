package events

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/authl2/typestructs"
	"triptychlabs.io/dao/v2/src/generated/auth"
	"triptychlabs.io/dao/v2/src/keys"
)

func registerHolder(event *typestructs.RegisterHolderEvent) {
	oracle := keys.GetProvider(0)
	hash := md5.New()

	io.WriteString(hash, event.Holder.String())
	hexhash := fmt.Sprintf("%x", hash.Sum(nil))
	enrollment, bump := GetEnrollment(hexhash)

	ix := auth.NewHolderEnrollInstructionBuilder().
		SetAccessCode(event.AccessCode).
		SetBump(bump).
		SetEnrollmentAccount(enrollment).
		SetHolderHash(hexhash).
		SetOracleAccount(oracle.PublicKey()).
		SetSystemProgramAccount(solana.SystemProgramID)

	ixs := append(
		make([]solana.Instruction, 0),
		ix.Build(),
	)
	signers := append(
		make([]solana.PrivateKey, 0),
		oracle,
	)

	SendTx(
		"Enrollment",
		ixs,
		signers,
		oracle.PublicKey(),
	)

}
