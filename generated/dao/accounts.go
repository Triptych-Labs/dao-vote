// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dao

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type DAO struct {
	Oracle      ag_solanago.PublicKey
	Name        string
	Description string
}

var DAODiscriminator = [8]byte{242, 60, 23, 196, 237, 48, 173, 129}

func (obj DAO) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DAODiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Oracle` param:
	err = encoder.Encode(obj.Oracle)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Description` param:
	err = encoder.Encode(obj.Description)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DAO) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DAODiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[242 60 23 196 237 48 173 129]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Oracle`:
	err = decoder.Decode(&obj.Oracle)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Description`:
	err = decoder.Decode(&obj.Description)
	if err != nil {
		return err
	}
	return nil
}

type Ballot struct {
	Votes       int64
	End         int64
	Dao         ag_solanago.PublicKey
	Name        string
	Description string
}

var BallotDiscriminator = [8]byte{3, 232, 121, 204, 232, 137, 138, 164}

func (obj Ballot) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BallotDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Votes` param:
	err = encoder.Encode(obj.Votes)
	if err != nil {
		return err
	}
	// Serialize `End` param:
	err = encoder.Encode(obj.End)
	if err != nil {
		return err
	}
	// Serialize `Dao` param:
	err = encoder.Encode(obj.Dao)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Description` param:
	err = encoder.Encode(obj.Description)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Ballot) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BallotDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[3 232 121 204 232 137 138 164]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Votes`:
	err = decoder.Decode(&obj.Votes)
	if err != nil {
		return err
	}
	// Deserialize `End`:
	err = decoder.Decode(&obj.End)
	if err != nil {
		return err
	}
	// Deserialize `Dao`:
	err = decoder.Decode(&obj.Dao)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Description`:
	err = decoder.Decode(&obj.Description)
	if err != nil {
		return err
	}
	return nil
}

type Member struct {
	Votes         int64
	LastVoteEpoch int64
	Holder        ag_solanago.PublicKey
	Dao           ag_solanago.PublicKey
	Description   string
}

var MemberDiscriminator = [8]byte{54, 19, 162, 21, 29, 166, 17, 198}

func (obj Member) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MemberDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Votes` param:
	err = encoder.Encode(obj.Votes)
	if err != nil {
		return err
	}
	// Serialize `LastVoteEpoch` param:
	err = encoder.Encode(obj.LastVoteEpoch)
	if err != nil {
		return err
	}
	// Serialize `Holder` param:
	err = encoder.Encode(obj.Holder)
	if err != nil {
		return err
	}
	// Serialize `Dao` param:
	err = encoder.Encode(obj.Dao)
	if err != nil {
		return err
	}
	// Serialize `Description` param:
	err = encoder.Encode(obj.Description)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Member) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MemberDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[54 19 162 21 29 166 17 198]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Votes`:
	err = decoder.Decode(&obj.Votes)
	if err != nil {
		return err
	}
	// Deserialize `LastVoteEpoch`:
	err = decoder.Decode(&obj.LastVoteEpoch)
	if err != nil {
		return err
	}
	// Deserialize `Holder`:
	err = decoder.Decode(&obj.Holder)
	if err != nil {
		return err
	}
	// Deserialize `Dao`:
	err = decoder.Decode(&obj.Dao)
	if err != nil {
		return err
	}
	// Deserialize `Description`:
	err = decoder.Decode(&obj.Description)
	if err != nil {
		return err
	}
	return nil
}

type Casting struct{}

var CastingDiscriminator = [8]byte{131, 254, 180, 136, 238, 215, 239, 59}

func (obj Casting) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(CastingDiscriminator[:], false)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Casting) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(CastingDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[131 254 180 136 238 215 239 59]",
				fmt.Sprint(discriminator[:]))
		}
	}
	return nil
}

type NFT struct {
	Dao    ag_solanago.PublicKey
	Ballot ag_solanago.PublicKey
	Member ag_solanago.PublicKey
}

var NFTDiscriminator = [8]byte{97, 230, 6, 21, 131, 208, 111, 115}

func (obj NFT) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(NFTDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Dao` param:
	err = encoder.Encode(obj.Dao)
	if err != nil {
		return err
	}
	// Serialize `Ballot` param:
	err = encoder.Encode(obj.Ballot)
	if err != nil {
		return err
	}
	// Serialize `Member` param:
	err = encoder.Encode(obj.Member)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NFT) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(NFTDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[97 230 6 21 131 208 111 115]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Dao`:
	err = decoder.Decode(&obj.Dao)
	if err != nil {
		return err
	}
	// Deserialize `Ballot`:
	err = decoder.Decode(&obj.Ballot)
	if err != nil {
		return err
	}
	// Deserialize `Member`:
	err = decoder.Decode(&obj.Member)
	if err != nil {
		return err
	}
	return nil
}