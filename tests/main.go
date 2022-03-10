package main

import (
	"triptychlabs.io/dao/v2/src/keys"
	"triptychlabs.io/dao/v2/tests/tests"
)

func init() {
	// dao.SetProgramID(solana.MustPublicKeyFromBase58("E18jUpqrxp8w4u556G4CcE1jHW7iAvt7i6JH7SkGjaD8"))
	// auth.SetProgramID(solana.MustPublicKeyFromBase58("9LS5eDSs36coxwkfhvUyWn7ejvVYnKAebsHTuLmky4aT"))
	keys.SetupProviders()
}

func main() {
	tests.Proposel()
}

