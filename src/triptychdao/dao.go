package triptychdao

import (
	"fmt"
	"log"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/generated/dao"
	"triptychlabs.io/dao/v2/src/utils"
)

func CreateDao(daoIndex uint64) {
	dao.SetProgramID(solana.MustPublicKeyFromBase58("E18jUpqrxp8w4u556G4CcE1jHW7iAvt7i6JH7SkGjaD8"))

	oracle := solana.MustPrivateKeyFromBase58("4sXP2Zebxsv1dt382f9ym4B4qyaPQV66YksF4ENG2zVtRBDCy2kvC8CQGxoFJWT15JHvMHhazjXnmM2ooycudpkL")

	daoPDA, daoBump, err := utils.FindDaoAddress(oracle.PublicKey(), daoIndex)
	if err != nil {
		panic(err)
	}

	createDaoIx := dao.NewCreateDaoInstructionBuilder().
		SetBump(daoBump).
		SetDaoAccount(daoPDA).
		SetDaoIndex(daoIndex).
		SetDescription("asdfasd").
		SetName("asd").
		SetOracleAccount(oracle.PublicKey()).
		SetSystemProgramAccount(solana.SystemProgramID)

	log.Println("....")
	ixs := make([]solana.Instruction, 0)
	ixs = append(
		ixs,
		createDaoIx.Build(),
	)
	log.Println("....")
	signers := make([]solana.PrivateKey, 0)
	signers = append(
		signers,
		oracle,
	)

	SendTx(
		"Create DAO",
		ixs,
		signers,
		oracle.PublicKey(),
	)

	fmt.Println("Success")
}

