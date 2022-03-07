#![deny(rustdoc::all)]
#![allow(rustdoc::missing_doc_code_examples)]

use anchor_lang::prelude::*;
use anchor_lang::Key;
use std::convert::Into;
use std::vec::Vec;
use vipers::unwrap_int;

mod state;

pub use state::*;

declare_id!("E18jUpqrxp8w4u556G4CcE1jHW7iAvt7i6JH7SkGjaD8");

#[program]
pub mod dao {
    use super::*;

    pub fn create_dao(
        ctx: Context<Create>,
        _bump: u8,
        _dao_index: u64,
        name: String,
        description: String,
    ) -> ProgramResult {
        let dao_authority = &mut ctx.accounts.dao;
        dao_authority.oracle = ctx.accounts.oracle.key();
        dao_authority.name = name;
        dao_authority.description = description;

        Ok(())
    }

    pub fn create_ballot(
        ctx: Context<Propose>,
        _bump: u8,
        _ballot_index: u64,
        name: String,
        description: String,
    ) -> ProgramResult {
        let dao_authority = &mut ctx.accounts.dao;
        dao_authority.oracle = ctx.accounts.oracle.key();
        dao_authority.name = name;
        dao_authority.description = description;

        Ok(())
    }

    pub fn holder_register(
        ctx: Context<Register>,
        _bump: u8,
        holder: Pubkey,
    ) -> ProgramResult {
        require!(ctx.accounts.oracle.key() == ctx.accounts.dao.oracle.key(), BadOracle);

        let member = &mut ctx.accounts.member;
        member.holder = holder;
        member.dao = ctx.accounts.dao.key();

        Ok(())
    }

    pub fn nft_enter(
        ctx: Context<Enter>,
        _bump: u8,
        _mint: Pubkey,
    ) -> ProgramResult {
        require!(ctx.accounts.oracle.key() == ctx.accounts.dao.oracle.key(), BadOracle);
        let nft = &mut ctx.accounts.nft;
        nft.member = ctx.accounts.member.key();
        nft.dao = ctx.accounts.dao.key();

        Ok(())
    }

    pub fn nft_sync(
        ctx: Context<Sync>,
    ) -> ProgramResult {
        require!(ctx.accounts.oracle.key() == ctx.accounts.dao.oracle.key(), BadOracle);
        require!(ctx.accounts.ballot.key() == ctx.accounts.nft.ballot.key(), BadNFTBallot);
        require!(ctx.accounts.dao.key() == ctx.accounts.nft.dao.key(), BadNFTDAO);
        let nft = &mut ctx.accounts.nft;
        nft.member = ctx.accounts.current_member.key();
        let previous_member = &mut ctx.accounts.previous_member;
        previous_member.votes = unwrap_int!(previous_member.votes.checked_sub(1));

        Ok(())
    }

    pub fn holder_vote(
        ctx: Context<Cast>,
        _bump: u8,
        votes: u64,
    ) -> ProgramResult {
        require!(ctx.accounts.oracle.key() == ctx.accounts.dao.oracle.key(), BadOracle);
        require!(ctx.accounts.dao.key() == ctx.accounts.dao.key(), BadDAO);
        let member = &mut ctx.accounts.member;

        member.votes = i64::try_from(votes).unwrap();

        Ok(())
    }
}

#[derive(Accounts)]
#[instruction(bump: u8, dao_index: u64)]
pub struct Create<'info> {
    pub oracle: Signer<'info>,
    #[account(
        init,
        seeds = [
            b"dao".as_ref(),
            oracle.key().to_bytes().as_ref(),
            &dao_index.to_le_bytes(),
        ],
        bump,
        payer = oracle,
        space = DAO::space(),
    )]
    pub dao: Account<'info, DAO>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(bump: u8, holder: Pubkey)]
pub struct Register<'info> {
    pub oracle: Signer<'info>,
    #[account(mut)]
    pub dao: Account<'info, DAO>,
    #[account(
        init,
        seeds = [
            b"member".as_ref(),
            dao.key().to_bytes().as_ref(),
            holder.as_ref(),
        ],
        bump,
        payer = oracle,
        space = Member::space(),
    )]
    pub member: Account<'info, Member>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(bump: u8, ballot_index: u64)]
pub struct Propose<'info> {
    pub oracle: Signer<'info>,
    #[account(mut)]
    pub dao: Account<'info, DAO>,
    #[account(
        init,
        seeds = [
            b"ballot".as_ref(),
            dao.key().to_bytes().as_ref(),
            ballot_index.to_le_bytes().as_ref(),
        ],
        bump,
        payer = oracle,
        space = Ballot::space(),
    )]
    pub member: Account<'info, Member>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(bump: u8, mint: Pubkey)]
pub struct Enter<'info> {
    pub oracle: Signer<'info>,
    pub ballot: Account<'info, Ballot>,
    #[account(mut)]
    pub dao: Account<'info, DAO>,
    #[account(
        init,
        seeds = [
            b"NFT".as_ref(),
            dao.key().to_bytes().as_ref(),
            ballot.key().to_bytes().as_ref(),
            mint.key().to_bytes().as_ref(),
        ],
        bump,
        payer = oracle,
        space = NFT::space(),
    )]
    pub nft: Account<'info, NFT>,
    pub member: Account<'info, Member>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction()]
pub struct Sync<'info> {
    pub oracle: Signer<'info>,
    pub ballot: Account<'info, Ballot>,
    pub dao: Account<'info, DAO>,
    pub previous_member: Account<'info, Member>,
    pub current_member: Account<'info, Member>,
    #[account(mut)]
    pub nft: Account<'info, NFT>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(bump: u8)]
pub struct Cast<'info> {
    pub oracle: Signer<'info>,
    #[account(mut)]
    pub ballot: Account<'info, Ballot>,
    #[account(mut)]
    pub dao: Account<'info, DAO>,
    #[account(
        init,
        seeds = [
            b"electorate".as_ref(),
            dao.key().to_bytes().as_ref(),
            ballot.key().to_bytes().as_ref(),
            member.key().to_bytes().as_ref(),
        ],
        bump,
        payer = oracle,
        space = Casting::space(),
    )]
    pub casting: Account<'info, Casting>,
    pub member: Account<'info, Member>,
    pub system_program: Program<'info, System>,
}


#[error]
pub enum ErrorCode {
    #[msg("Bad Oracle.")]
    BadOracle,
    #[msg("Bad Ballot.")]
    BadBallot,
    #[msg("Bad DAO.")]
    BadDAO,
    #[msg("Bad NFT DAO.")]
    BadNFTDAO,
    #[msg("Bad Ballot.")]
    BadNFTBallot,
}
