// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dao

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// HolderVote is the `holderVote` instruction.
type HolderVote struct {
	Bump  *uint8
	Votes *uint64

	// [0] = [SIGNER] oracle
	//
	// [1] = [WRITE] ballot
	//
	// [2] = [WRITE] dao
	//
	// [3] = [WRITE] casting
	//
	// [4] = [] member
	//
	// [5] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewHolderVoteInstructionBuilder creates a new `HolderVote` instruction builder.
func NewHolderVoteInstructionBuilder() *HolderVote {
	nd := &HolderVote{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetBump sets the "bump" parameter.
func (inst *HolderVote) SetBump(bump uint8) *HolderVote {
	inst.Bump = &bump
	return inst
}

// SetVotes sets the "votes" parameter.
func (inst *HolderVote) SetVotes(votes uint64) *HolderVote {
	inst.Votes = &votes
	return inst
}

// SetOracleAccount sets the "oracle" account.
func (inst *HolderVote) SetOracleAccount(oracle ag_solanago.PublicKey) *HolderVote {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(oracle).SIGNER()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *HolderVote) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBallotAccount sets the "ballot" account.
func (inst *HolderVote) SetBallotAccount(ballot ag_solanago.PublicKey) *HolderVote {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ballot).WRITE()
	return inst
}

// GetBallotAccount gets the "ballot" account.
func (inst *HolderVote) GetBallotAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetDaoAccount sets the "dao" account.
func (inst *HolderVote) SetDaoAccount(dao ag_solanago.PublicKey) *HolderVote {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(dao).WRITE()
	return inst
}

// GetDaoAccount gets the "dao" account.
func (inst *HolderVote) GetDaoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCastingAccount sets the "casting" account.
func (inst *HolderVote) SetCastingAccount(casting ag_solanago.PublicKey) *HolderVote {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(casting).WRITE()
	return inst
}

// GetCastingAccount gets the "casting" account.
func (inst *HolderVote) GetCastingAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMemberAccount sets the "member" account.
func (inst *HolderVote) SetMemberAccount(member ag_solanago.PublicKey) *HolderVote {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(member)
	return inst
}

// GetMemberAccount gets the "member" account.
func (inst *HolderVote) GetMemberAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *HolderVote) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *HolderVote {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *HolderVote) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst HolderVote) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_HolderVote,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst HolderVote) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *HolderVote) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Bump == nil {
			return errors.New("Bump parameter is not set")
		}
		if inst.Votes == nil {
			return errors.New("Votes parameter is not set")
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
			return errors.New("accounts.Casting is not set")
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

func (inst *HolderVote) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("HolderVote")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" Bump", *inst.Bump))
						paramsBranch.Child(ag_format.Param("Votes", *inst.Votes))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       oracle", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       ballot", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          dao", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      casting", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       member", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj HolderVote) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Votes` param:
	err = encoder.Encode(obj.Votes)
	if err != nil {
		return err
	}
	return nil
}
func (obj *HolderVote) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Votes`:
	err = decoder.Decode(&obj.Votes)
	if err != nil {
		return err
	}
	return nil
}

// NewHolderVoteInstruction declares a new HolderVote instruction with the provided parameters and accounts.
func NewHolderVoteInstruction(
	// Parameters:
	bump uint8,
	votes uint64,
	// Accounts:
	oracle ag_solanago.PublicKey,
	ballot ag_solanago.PublicKey,
	dao ag_solanago.PublicKey,
	casting ag_solanago.PublicKey,
	member ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *HolderVote {
	return NewHolderVoteInstructionBuilder().
		SetBump(bump).
		SetVotes(votes).
		SetOracleAccount(oracle).
		SetBallotAccount(ballot).
		SetDaoAccount(dao).
		SetCastingAccount(casting).
		SetMemberAccount(member).
		SetSystemProgramAccount(systemProgram)
}
