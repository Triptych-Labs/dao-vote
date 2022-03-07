In order to interact with the dao, the user is forced to sign a message that is deferred to as certification to elevate dao interactions.

Dao interactions are executed in browser-space via WebAssembly which allows us to do many things around the newfound obfuscation principle:
    * Securely deliver private keys to the client
    * Securely uphold sensitive business logic
    * Impale front running attacks

I suppose such a dao principle requires masking of participating identities so by therein embedding an authoritative identity that is able to co-operate amongst the user experience, we can sidecar(employ) transaction executions via embedded authority identity.




---------

****
****
Its important to note
****
****
Any holder identity needs to be an _Instruction Argument_ and **not** an _Account Key_. This prevents the address from being associated to a Transaction.
****
****





So `casting` accounts will retain `number_of_votes` per `member` account.
