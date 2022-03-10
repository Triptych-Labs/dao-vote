---------
Its important to note
Any holder identity needs to be an _Instruction Argument_ and *not* an _Account Key_. This prevents the address from being associated to a Transaction.
---------

In order to interact with the dao, the user is forced to sign a message that is deferred to as certification to elevate dao interactions.

Dao interactions are executed in browser-space via WebAssembly which allows us to do many things around the newfound obfuscation principle:
    * Securely deliver private keys to the client
    * Securely uphold sensitive business logic
    * Impale front running attacks

I suppose such a dao principle requires masking of participating identities so by therein embedding an authoritative identity that is able to co-operate amongst the user experience, we can sidecar (employ) transaction executions via embedded authority identity. Kind of.

As I learned that string literals can be interpreted after disassembly, we implement a unique security pattern. A seperate onchain program to that of the dao will maintain a decipher key that is only given to the caller upon verification criteria. Principally, the data is always encrypted at rest when presiding outside of an internal L2, script, and Web2 asset that is locally owned.

This secondary program must _not_ present any delineating vector of association to its caller (Signer/Holder/Member) so to uphold someones anonymity. This is technically achieved with a centralised L2 via `Anchor Events`.

Event data structures require obfuscation as such data would be publically accessible on a block explorer. 

The process follows -

1) `User` commits some auth/registration transaction to auth contract.
2) L2 commits 2FA data of `User` to auth contract.
3) WebAssembly fetches 2FA of `User`.
4) `User` accesses dao.

