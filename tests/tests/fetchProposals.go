package tests

import (
	"crypto/sha256"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/mr-tron/base58"
	"triptychlabs.io/dao/v2/src/solanarpc"
)

// hWys35ZYZAY - account:DAO
// euzBDdhcHR - account:Ballot

func Proposel() {
	daoProgramId := solana.MustPublicKeyFromBase58("E18jUpqrxp8w4u556G4CcE1jHW7iAvt7i6JH7SkGjaD8")

	// disc := dao.BallotDiscriminator
	discriminator := []byte("account:Ballot")
	h := sha256.New()
	h.Write(discriminator[:])
	disc := base58.Encode(h.Sum(nil)[:8])

	ballots, err := solanarpc.FetchProposalsData(daoProgramId, disc)
	if err != nil {
		panic(err)
	}

	for _, ballot := range *ballots {
		fmt.Println(ballot)
	}
}
