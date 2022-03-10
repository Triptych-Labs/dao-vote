package tests

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/generated/auth"
)

func tryRecovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func Try() {
	defer tryRecovery()

	fmt.Println("....")
	auth.SetProgramID(solana.MustPublicKeyFromBase58("9LS5eDSs36coxwkfhvUyWn7ejvVYnKAebsHTuLmky4aT"))
	fmt.Println("....")
	auth.SetProgramID(solana.MustPublicKeyFromBase58("9LS5eDSs36coxwkfhvUyWn7ejvVYnKAebsHTuLmky4aT"))
	fmt.Println("....")
	auth.SetProgramID(solana.MustPublicKeyFromBase58("9LS5eDSs36coxwkfhvUyWn7ejvVYnKAebsHTuLmky4aT"))
	fmt.Println("....")
	auth.SetProgramID(solana.MustPublicKeyFromBase58("9LS5eDSs36coxwkfhvUyWn7ejvVYnKAebsHTuLmky4aT"))
}
