// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dao

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// NftEnter is the `nftEnter` instruction.
type NftEnter struct {
	Bump *uint8
	Mint *ag_solanago.PublicKey

	// [0] = [SIGNER] oracle
	//
	// [1] = [] ballot
	//
	// [2] = [WRITE] dao
	//
	// [3] = [WRITE] nft
	//
	// [4] = [] member
	//
	// [5] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewNftEnterInstructionBuilder creates a new `NftEnter` instruction builder.
func NewNftEnterInstructionBuilder() *NftEnter {
	nd := &NftEnter{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetBump sets the "bump" parameter.
func (inst *NftEnter) SetBump(bump uint8) *NftEnter {
	inst.Bump = &bump
	return inst
}

// SetMint sets the "mint" parameter.
func (inst *NftEnter) SetMint(mint ag_solanago.PublicKey) *NftEnter {
	inst.Mint = &mint
	return inst
}

// SetOracleAccount sets the "oracle" account.
func (inst *NftEnter) SetOracleAccount(oracle ag_solanago.PublicKey) *NftEnter {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(oracle).SIGNER()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *NftEnter) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBallotAccount sets the "ballot" account.
func (inst *NftEnter) SetBallotAccount(ballot ag_solanago.PublicKey) *NftEnter {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ballot)
	return inst
}

// GetBallotAccount gets the "ballot" account.
func (inst *NftEnter) GetBallotAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetDaoAccount sets the "dao" account.
func (inst *NftEnter) SetDaoAccount(dao ag_solanago.PublicKey) *NftEnter {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(dao).WRITE()
	return inst
}

// GetDaoAccount gets the "dao" account.
func (inst *NftEnter) GetDaoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetNftAccount sets the "nft" account.
func (inst *NftEnter) SetNftAccount(nft ag_solanago.PublicKey) *NftEnter {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(nft).WRITE()
	return inst
}

// GetNftAccount gets the "nft" account.
func (inst *NftEnter) GetNftAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMemberAccount sets the "member" account.
func (inst *NftEnter) SetMemberAccount(member ag_solanago.PublicKey) *NftEnter {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(member)
	return inst
}

// GetMemberAccount gets the "member" account.
func (inst *NftEnter) GetMemberAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *NftEnter) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *NftEnter {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *NftEnter) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst NftEnter) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_NftEnter,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst NftEnter) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *NftEnter) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Bump == nil {
			return errors.New("Bump parameter is not set")
		}
		if inst.Mint == nil {
			return errors.New("Mint parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Ballot is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Dao is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Nft is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Member is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *NftEnter) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("NftEnter")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Bump", *inst.Bump))
						paramsBranch.Child(ag_format.Param("Mint", *inst.Mint))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       oracle", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       ballot", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          dao", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          nft", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       member", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj NftEnter) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	return nil
}
func (obj *NftEnter) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	return nil
}

// NewNftEnterInstruction declares a new NftEnter instruction with the provided parameters and accounts.
func NewNftEnterInstruction(
	// Parameters:
	bump uint8,
	mint ag_solanago.PublicKey,
	// Accounts:
	oracle ag_solanago.PublicKey,
	ballot ag_solanago.PublicKey,
	dao ag_solanago.PublicKey,
	nft ag_solanago.PublicKey,
	member ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *NftEnter {
	return NewNftEnterInstructionBuilder().
		SetBump(bump).
		SetMint(mint).
		SetOracleAccount(oracle).
		SetBallotAccount(ballot).
		SetDaoAccount(dao).
		SetNftAccount(nft).
		SetMemberAccount(member).
		SetSystemProgramAccount(systemProgram)
}
