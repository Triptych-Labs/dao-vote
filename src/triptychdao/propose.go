package triptychdao

import (
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/generated/dao"
	"triptychlabs.io/dao/v2/src/utils"
)

func Propose(daoIndex, ballotIndex uint64) {
	oracle := solana.MustPrivateKeyFromBase58("4sXP2Zebxsv1dt382f9ym4B4qyaPQV66YksF4ENG2zVtRBDCy2kvC8CQGxoFJWT15JHvMHhazjXnmM2ooycudpkL")

	daoPDA, _, err := utils.FindDaoAddress(oracle.PublicKey(), daoIndex)
	if err != nil {
		panic(err)
	}

	ballotPDA, ballotBump, err := utils.FindBallotAddress(
		daoPDA,
		oracle.PublicKey(),
		ballotIndex,
	)
	if err != nil {
		panic(err)
	}

	proposalIx := dao.NewCreateProposalInstructionBuilder().
		SetBallotAccount(ballotPDA).
		SetBallotIndex(ballotIndex).
		SetBump(ballotBump).
		SetDaoAccount(daoPDA).
		SetDescription("FUCK ANDREW GOWER").
		SetEnd(uint64(time.Now().UTC().Unix()) + (5 * 60)).
		SetName("I love RSC").
		SetOracleAccount(oracle.PublicKey()).
		SetSystemProgramAccount(solana.SystemProgramID)

	ixs := make([]solana.Instruction, 0)
	ixs = append(
		ixs,
		proposalIx.Build(),
	)
	signers := make([]solana.PrivateKey, 0)
	signers = append(
		signers,
		oracle,
	)

	SendTx(
		"Create Proposal",
		ixs,
		signers,
		oracle.PublicKey(),
	)

	fmt.Println("Success")
}

