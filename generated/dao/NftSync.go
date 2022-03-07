// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dao

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// NftSync is the `nftSync` instruction.
type NftSync struct {

	// [0] = [SIGNER] oracle
	//
	// [1] = [] ballot
	//
	// [2] = [] dao
	//
	// [3] = [] previousMember
	//
	// [4] = [] currentMember
	//
	// [5] = [WRITE] nft
	//
	// [6] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewNftSyncInstructionBuilder creates a new `NftSync` instruction builder.
func NewNftSyncInstructionBuilder() *NftSync {
	nd := &NftSync{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetOracleAccount sets the "oracle" account.
func (inst *NftSync) SetOracleAccount(oracle ag_solanago.PublicKey) *NftSync {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(oracle).SIGNER()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *NftSync) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBallotAccount sets the "ballot" account.
func (inst *NftSync) SetBallotAccount(ballot ag_solanago.PublicKey) *NftSync {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ballot)
	return inst
}

// GetBallotAccount gets the "ballot" account.
func (inst *NftSync) GetBallotAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetDaoAccount sets the "dao" account.
func (inst *NftSync) SetDaoAccount(dao ag_solanago.PublicKey) *NftSync {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(dao)
	return inst
}

// GetDaoAccount gets the "dao" account.
func (inst *NftSync) GetDaoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPreviousMemberAccount sets the "previousMember" account.
func (inst *NftSync) SetPreviousMemberAccount(previousMember ag_solanago.PublicKey) *NftSync {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(previousMember)
	return inst
}

// GetPreviousMemberAccount gets the "previousMember" account.
func (inst *NftSync) GetPreviousMemberAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCurrentMemberAccount sets the "currentMember" account.
func (inst *NftSync) SetCurrentMemberAccount(currentMember ag_solanago.PublicKey) *NftSync {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(currentMember)
	return inst
}

// GetCurrentMemberAccount gets the "currentMember" account.
func (inst *NftSync) GetCurrentMemberAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetNftAccount sets the "nft" account.
func (inst *NftSync) SetNftAccount(nft ag_solanago.PublicKey) *NftSync {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(nft).WRITE()
	return inst
}

// GetNftAccount gets the "nft" account.
func (inst *NftSync) GetNftAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *NftSync) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *NftSync {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *NftSync) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst NftSync) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_NftSync,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst NftSync) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *NftSync) Validate() error {
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
			return errors.New("accounts.PreviousMember is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CurrentMember is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Nft is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *NftSync) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("NftSync")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        oracle", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        ballot", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("           dao", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("previousMember", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta(" currentMember", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           nft", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj NftSync) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *NftSync) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewNftSyncInstruction declares a new NftSync instruction with the provided parameters and accounts.
func NewNftSyncInstruction(
	// Accounts:
	oracle ag_solanago.PublicKey,
	ballot ag_solanago.PublicKey,
	dao ag_solanago.PublicKey,
	previousMember ag_solanago.PublicKey,
	currentMember ag_solanago.PublicKey,
	nft ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *NftSync {
	return NewNftSyncInstructionBuilder().
		SetOracleAccount(oracle).
		SetBallotAccount(ballot).
		SetDaoAccount(dao).
		SetPreviousMemberAccount(previousMember).
		SetCurrentMemberAccount(currentMember).
		SetNftAccount(nft).
		SetSystemProgramAccount(systemProgram)
}