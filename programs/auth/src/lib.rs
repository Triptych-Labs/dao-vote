#![deny(rustdoc::all)]
#![allow(rustdoc::missing_doc_code_examples)]

use anchor_lang::prelude::*;
use anchor_lang::Key;
use std::convert::Into;
use std::vec::Vec;

mod state;

pub use state::*;

declare_id!("9LS5eDSs36coxwkfhvUyWn7ejvVYnKAebsHTuLmky4aT");

#[program]
pub mod auth {
    use super::*;

    pub fn holder_register(ctx: Context<Register>, access_code: String) -> ProgramResult {
        emit!(RegisterHolderEvent {
            holder: ctx.accounts.holder.key(),
            access_code: access_code,
        });

        Ok(())
    }
    pub fn holder_enroll(
        ctx: Context<Enroll>,
        _bump: u8,
        holder_hash: String,
        access_code: String,
    ) -> ProgramResult {
        let enrollment = &mut ctx.accounts.enrollment;
        enrollment.holder = holder_hash;
        enrollment.access_code = access_code;
        enrollment.cipher_key = format!("somecipherkeyhash");

        Ok(())
    }
}

#[derive(Accounts)]
#[instruction()]
pub struct Register<'info> {
    #[account(mut)]
    pub holder: Signer<'info>,
    pub system_program: Program<'info, System>,
}
#[derive(Accounts)]
#[instruction(bump: u8, holder_hash: String)]
pub struct Enroll<'info> {
    #[account(
        init,
        seeds = [
            b"enroll".as_ref(),
            holder_hash.as_ref(),
        ],
        bump,
        payer = oracle,
        space = Enrollment::space(),
    )]
    pub enrollment: Account<'info, Enrollment>,
    #[account(mut)]
    pub oracle: Signer<'info>,
    pub system_program: Program<'info, System>,
}
