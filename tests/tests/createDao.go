package tests

import (
	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/triptychdao"
)

func CreateDao() {
	triptychdao.CreateDao(solana.PrivateKey{}, 0)
}

