package typestructs

import (
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type RegisterHolderEvent struct {
	Holder     ag_solanago.PublicKey
	AccessCode string
}

func (obj RegisterHolderEvent) MarshalWithEncoder(encoder *ag_binary.Encoder, eventDiscriminator []byte) (err error) {
	err = encoder.WriteBytes(eventDiscriminator[:], false)
	if err != nil {
		return err
	}
	err = encoder.Encode(obj.Holder)
	if err != nil {
		return err
	}
	err = encoder.Encode(obj.AccessCode)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RegisterHolderEvent) UnmarshalWithDecoder(decoder *ag_binary.Decoder, eventDiscriminator []byte) (err error) {
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(eventDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[66 0 62 83 227 66 175 18]",
				string(discriminator[:]))
		}
	}
	err = decoder.Decode(&obj.Holder)
	if err != nil {
		return err
	}
	err = decoder.Decode(&obj.AccessCode)
	if err != nil {
		return err
	}
	return nil
}
