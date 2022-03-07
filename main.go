package main

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	sendAndConfirmTransaction "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"triptychlabs.io/dao/v2/generated/dao"
)

func init() {
	dao.SetProgramID(solana.MustPublicKeyFromBase58("E18jUpqrxp8w4u556G4CcE1jHW7iAvt7i6JH7SkGjaD8"))
}

func main() {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	daoPDA, daoBump, err := findDaoAddress(oracle.PublicKey(), 0)
	if err != nil {
		panic(err)
	}

	createDaoIx := dao.NewCreateDaoInstructionBuilder().
		SetBump(daoBump).
		SetDaoAccount(daoPDA).
		SetDaoIndex(0).
		SetDescription("asdfasd").
		SetName("asd").
		SetOracleAccount(oracle.PublicKey()).
		SetSystemProgramAccount(solana.SystemProgramID)

	ixs := make([]solana.Instruction, 0)
	ixs = append(
		ixs,
		createDaoIx.Build(),
	)
	signers := make([]solana.PrivateKey, 0)
	signers = append(
		signers,
		oracle,
	)

	sig, err := sendTransaction(
		oracle,
		ixs,
		signers,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(sig)

	fmt.Println("Success")
}

func sendTransaction(
	feePayer solana.PrivateKey,
	instructions []solana.Instruction,
	signers []solana.PrivateKey,
) (solana.Signature, error) {
	client := rpc.New(rpc.DevNet_RPC)

	recent, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, err
	}

	tx, err := solana.NewTransaction(
		instructions,
		recent.Value.Blockhash,
		solana.TransactionPayer(feePayer.PublicKey()),
	)
	if err != nil {
		return solana.Signature{}, err
	}

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		for _, candidate := range signers {
			if candidate.PublicKey().Equals(key) {
				return &candidate
			}
		}
		return nil
	})
	if err != nil {
		return solana.Signature{}, err
	}
	spew.Dump(tx)

	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		return solana.Signature{}, err
	}

	return sendAndConfirmTransaction.SendAndConfirmTransaction(
		context.TODO(),
		client,
		wsClient,
		tx,
	)
}

func findDaoAddress(oracle solana.PublicKey, daoIndex uint64) (solana.PublicKey, uint8, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, daoIndex)
	addr, bump, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("dao"),
			oracle.Bytes(),
			buf,
		},
		dao.ProgramID,
	)
	return addr, bump, err
}
