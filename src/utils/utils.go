package utils

import (
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/generated/dao"
)

func FindDaoAddress(oracle solana.PublicKey, index uint64) (solana.PublicKey, uint8, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
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

func FindBallotAddress(
	daoPDA solana.PublicKey,
	oracle solana.PublicKey,
	index uint64,
) (solana.PublicKey, uint8, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	addr, bump, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("ballot"),
			daoPDA.Bytes(),
			oracle.Bytes(),
			buf,
		},
		dao.ProgramID,
	)
	return addr, bump, err
}
