// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dao

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// HolderRegister is the `holderRegister` instruction.
type HolderRegister struct {
	Bump   *uint8
	Holder *ag_solanago.PublicKey

	// [0] = [SIGNER] oracle
	//
	// [1] = [WRITE] dao
	//
	// [2] = [WRITE] member
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewHolderRegisterInstructionBuilder creates a new `HolderRegister` instruction builder.
func NewHolderRegisterInstructionBuilder() *HolderRegister {
	nd := &HolderRegister{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetBump sets the "bump" parameter.
func (inst *HolderRegister) SetBump(bump uint8) *HolderRegister {
	inst.Bump = &bump
	return inst
}

// SetHolder sets the "holder" parameter.
func (inst *HolderRegister) SetHolder(holder ag_solanago.PublicKey) *HolderRegister {
	inst.Holder = &holder
	return inst
}

// SetOracleAccount sets the "oracle" account.
func (inst *HolderRegister) SetOracleAccount(oracle ag_solanago.PublicKey) *HolderRegister {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(oracle).SIGNER()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *HolderRegister) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetDaoAccount sets the "dao" account.
func (inst *HolderRegister) SetDaoAccount(dao ag_solanago.PublicKey) *HolderRegister {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(dao).WRITE()
	return inst
}

// GetDaoAccount gets the "dao" account.
func (inst *HolderRegister) GetDaoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMemberAccount sets the "member" account.
func (inst *HolderRegister) SetMemberAccount(member ag_solanago.PublicKey) *HolderRegister {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(member).WRITE()
	return inst
}

// GetMemberAccount gets the "member" account.
func (inst *HolderRegister) GetMemberAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *HolderRegister) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *HolderRegister {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *HolderRegister) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst HolderRegister) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_HolderRegister,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst HolderRegister) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *HolderRegister) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Bump == nil {
			return errors.New("Bump parameter is not set")
		}
		if inst.Holder == nil {
			return errors.New("Holder parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Dao is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Member is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *HolderRegister) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("HolderRegister")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  Bump", *inst.Bump))
						paramsBranch.Child(ag_format.Param("Holder", *inst.Holder))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       oracle", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          dao", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       member", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj HolderRegister) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Holder` param:
	err = encoder.Encode(obj.Holder)
	if err != nil {
		return err
	}
	return nil
}
func (obj *HolderRegister) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Holder`:
	err = decoder.Decode(&obj.Holder)
	if err != nil {
		return err
	}
	return nil
}

// NewHolderRegisterInstruction declares a new HolderRegister instruction with the provided parameters and accounts.
func NewHolderRegisterInstruction(
	// Parameters:
	bump uint8,
	holder ag_solanago.PublicKey,
	// Accounts:
	oracle ag_solanago.PublicKey,
	dao ag_solanago.PublicKey,
	member ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *HolderRegister {
	return NewHolderRegisterInstructionBuilder().
		SetBump(bump).
		SetHolder(holder).
		SetOracleAccount(oracle).
		SetDaoAccount(dao).
		SetMemberAccount(member).
		SetSystemProgramAccount(systemProgram)
}
