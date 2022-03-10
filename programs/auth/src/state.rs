use crate::*;

#[event]
pub struct RegisterHolderEvent {
    #[index]
    pub holder: Pubkey,
    pub access_code: String, // md5 hash from frontend
}

#[account]
#[derive(Default, Debug, PartialEq)]
pub struct Enrollment {
    pub holder: String,      // md5 hash
    pub access_code: String, // md5 hash
    pub cipher_key: String,  // md5 hash
}

impl Enrollment {
    pub fn space() -> usize {
        //8 + 32 + 4 + 32 + 4 + 32
        8 + 32 + 32 + 32
    }
}
