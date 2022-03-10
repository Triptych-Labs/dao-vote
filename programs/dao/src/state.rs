use anchor_lang::prelude::*;

#[account]
#[derive(Default, Debug, PartialEq)]
pub struct DAO {
    pub oracle: Pubkey,
    pub name: String,
    pub description: String,
}

impl DAO {
    pub fn space() -> usize {
        8 + 32 + 4 + 32 + 4 + 32
    }
}

#[account]
#[derive(Default, Debug, PartialEq)]
pub struct Ballot {
    pub votes: i64,
    pub end: u64,
    pub dao: Pubkey,
    pub name: String,
    pub description: String,
}

impl Ballot {
    pub fn space() -> usize {
        8 + 8 + 8 + 32 + 4 + 32 + 4 + 32
    }
}

#[account]
#[derive(Default, Debug, PartialEq)]
pub struct Member {
    pub votes: i64,
    pub last_vote_epoch: i64,
    pub holder: Pubkey,
    pub dao: Pubkey,
    pub description: String,
}

impl Member {
    pub fn space() -> usize {
        8 + 8 + 8 + 32 + 32 + 4 + 128
    }
}

#[account]
#[derive(Default, Debug, PartialEq)]
pub struct Casting {}

impl Casting {
    pub fn space() -> usize {
        8 + 8
    }
}

#[account]
#[derive(Default, Debug, PartialEq)]
pub struct NFT {
    pub dao: Pubkey,
    pub ballot: Pubkey,
    pub member: Pubkey,
}

impl NFT {
    pub fn space() -> usize {
        8 + 32 + 32 + 32
    }
}
