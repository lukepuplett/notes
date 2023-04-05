# Mastering Ethereum - Building Smart Contracts and DApps

Notes taken from the book by Andreas Antonopoulos and Dr Gavin Wood.

Repo for the book: https://github.com/ethereumbook/ethereumbook

### Intro

- TCP port 30303, and UDP 30303.
- Various client like GoEth, Geth and Parity.
- Greek letter Xi Ξ or ♦1 is 1 Ether.
- 1 _wei_ is 1 quintillionth of 1 Ether, or 1x10^18.
- Ether is always represented using `uint` denominated in _wei_.
- Metamask wallet address and private key is same on all networks.
- Use the testnets Fauncet to fund your wallet.
- Use etherscan.io to look up the transaction.
- Use Metamask plus test Fauncet to donate and return test ETH.
- Hoarding test ETH is frowned upon; though donating costs gas.
- EOA is Externally Owned Account and has private key.
- Contract accounts have no private key.
- Contracts cannot initiate transactions.
- `function() public payable { }` is the default function called if a transfer is made to the contract address; it makes the contract payable.
- Remix is an IDE with a compiler.
- Deploying a contract is done by tooling and Metamask by sending to destination 0x0 the "zero address".
- JavaScript limit is 10^17 so gwei is treated as a string and BigNumber.
- Tooling like Remix allows a contract to be inspected by its address and it'll make a little GUI with the contract's methods available to execute.
- Transactions sent by the contract are called "internal transactions" or messages and show differently in block explorers.
- `mg.sender.transfer(addr);`

### Ethereum Clients

- "Remote Client" is a wallet with an API like web3.js
- Ganache is a local blockchain emulator.
- "Light Client" validates transactions and evalutates their effects like a full node.
- Running a full node for a testnet is feasible.
- Since THE MERGE two nodes are now needed: an execution layer node (EL) and a consensus layer node (CL).
- Geth 500GB "snap sync" and >12GB for full node with all data.
- DappNode and Avado will sell a prebuilt, pre-setup box as a node.
- DappNode is also a "launcher" and "control center" app you can use on your own kit.
- Others include: eth-docker, Stereum, NiceNode and Sedge.
- Execution Clients: Besu, Erigon, Geth, Nethermind.
- Consensus Clients: Lighthouse, Lodestar, Nimbus, Prysm, Teku.
- Use GnuPG `> gpg` to verify signature or open source binaries or build/compile from source yourself.
- OR `> sha256sum teku-22.6.1.tar.gz` will print the hash so you can check it matches what it says online.
- Before starting, choose the network and synch mode.
- You can also configure pruning options to remove old data.
- Choose to enable HTTP or WebSockets and local path for JWTSECRET used to authenticate EL<->CL node communications.
- Lodestar has the lowest CPU% and Teku the most. As of Feb 23 need C250GB for Goerli testnet node.
- Nimbus uses the least RAM <2GB, Teku most at 8GB. Lighthouse leaked until it crashed.
- Teku uses the least disk and Lighthouse the most >100GB.
- All use <200 ops/s and >80h to synch, >3 weeks on archive mode!
- All <6Mb/s network.
- Ganache is a local blockchain simulator.
- It has no other users, so no competition and sequencing and MEV is not representative or live env.
- No other contracts; no stablecoins or lib contracts.
- JSON RPC service runs on HTTP port 8545.

### Cryptography

- User EOAs are generated from public key part, but contract addresses are not gen from pub or priv key.
- Private key needed to sign transactions and prove ownership/control.
- Random entropy sent to Keccak256 hasher to get a private key.
- Ethereum uses uncompressed publick key presentation standard.
- 04 concat-with x-coord 64 hex chars concat-with y-coord 64 hex chars.
- secp256k1 elliptic curve in OpenSSL and libsecp256k1
- Ethereum uses originally designed Keccak256 and _not_ the NIST SHA-3 standard, due to Ed Snowden.
- You can use a "test vector" known input output to see if your Keccak hasher lib is correct.
- 04 is not included when you compute your own addresses since 04 just denotes the type of formatting in hex.
- Ethereum addresses are last 20 bytes of hash of public key; 0x denotes hex format.
- ICAP Inter Exchange Client Address Protocol encoding for Ethereum addresses partly compatible with IBAN encoding.
- IBAN = 34 alphanumeric case insensitive chars of country code, checksum, and bank account ID.
- ICAP uses XE as country code then 2-char checksum then any one of these account IDs:
  - Direct: big endian base36 up to 30 chars representing the 155 least significant bits of Ethereum address. 155 is fewer that 160 bits of a full Ethereum address so address must be (is assumed to be) 0 at beginning. Might want to keep trying to create addresses until you find one with zeroes at the beginning to have it work with IBAN.
  - Basic: same as above but 31 chars, so incompatible with IBAN field validation.
  - Indirect: encodes an ENS name. Uses 16 chars: asset ID "ETH", name service "XREG", and 9 char human readable name "KITTYKATS" so e.g. XE##ETHXREGKITTYKATS where ## is the checksum.
- `helpeth` command line tool makes ICAP addresses.
- EIP-55 is another format that uses the hex capitalization to encode the checksum!

### Wallets

- They're a GUI to Ethereum but technically they're a store of keys; a keyring that can create and sign transactions.
- Nondeterministic wallet: keys not related and are random each time: JBOK.
- Deterministic wallet: all keys derive from a single master key, often use BIP-32/44 tree structure.
- Heirarchically Deterministic (HD) wallets are created from a 128, 256 or 512-bit seed.
- Mnemonic code words are a backup of the seed key using BIP-39 standard.
- Best practice to use a new address (new private key) each time you receive funds.
- HD BIP-32 Bitcoin standard produces a tree where a key can produce children, which can produce grandchildren.
- Branches can be used to organise addresses perhaps in a corporation or organisation.
- Can create public keys without the private key, so a wallet can make addresses to receive funds but cannot spend because it does not have the private key needed to sign the transaction.
- Page 86 gives example for how to generate mnemonic words.
- Most heirarchically deterministic wallets follow BIP-32 and "Paths" BIP-43/44.
- Most/many BIP-32 implementations are designed for Bitcoin so use an Ethereum one!
- bip32.org is useful tool for testing.
- The term for using/deriving child/subchild keys is "extending" by taking a key and appending a "chain code".
- If the key is private, it becomes an "extended private key" prefixed xprv.
- An extended public key is prefixed xpub.
- Can derive xpub without the private keys; xpubs can come from the child private key or from parent public key.
- Can deploy xpub on a public web server insecure and safe to create a diffferent address to receive funds.
- **Risk!** If a child private key leaks then the chain code can be gotten from the xpub and used to compute **all** the child private keys!!
- **Worse!!** You can the compute the parent private key!!
- To counter this risk, HD wallets use "hardened derivation" which breaks the link between the parent public key and child chain code.
- Hardened derivation uses the parent private key to compute the child chain code.
- **Best practice** is to have level 1 children of the master keys always derived by hardened derivation in order to prevent compromise of the master keys; this installs a kind of firewall/break.
- An "index number" is an int32 index from the parent key which identifies and helps generate the child key.
- The int32 key space is ranged so as to identify hardened and non-hardened index keys.
- See page 95.
- Keys are identified via a "path" convention like m/0/1 or M/0/1, see page 95.
- Various BIP proposals for managing, organising and naming the paths, see page 96.

### Transactions

- These are signed messages originated by an EOA, transmitted by the network, recorded on the blockchain.
- They're the only things that can mutate state.
- The network serialization is the only standard form.
  - Nonce - sequence no. issued by EOA for anti-replay.
  - Gas Price - the amount, in wei, originator is willing to pay for gas.
  - Gas Limit - the max gas the originator is willing to buy for this transaction.
  - Recipient - destination address.
  - Value - optional amount to send to the address.
  - Data - optional variable-length binary payload.
  - v,r,s - three parts of ECDSA digital signature of the originating EOA.
- This structure is serialized using Recursive Length Prefix (RLP) encoding.
- All numbers in Ethereum encode as big-endian integers of lengths that are multiples of 8 bits.
- RLP does not endode the field names.
- EOA's address can be computed from v,r,s so there's no need for a from field.
- The transaction ID is computed from field data.
- Nonce is count of the number of tx sent from the address, or in the case of "accounts with associated code", the number of contract creations made by the account.
- Nonce provides an order of execution–it seems that "later" nonces are "ignored" but presumably they must be retried periodically because they eventually get processed if transaction with "before" nonces turn up to fill "the gap".
- I think the nonce is computed/looked-up from counting tx on the real blockchain by the originator when it sends, and by the validator when it checks.
- `>web3.eth.getTransactionCount("0x9e61...")`
- It appears that the need to use a valid incrementing nonce will mean designing a system around a single machine that issues transactions for an address, or a CAS-op must be run against an atomic database like MSSQL.
- When the system above starts, it'll have to wait a bit for all pending tx to complete on the blockchain, then query the blockchain to get the next nonce value.
- If you use a nonce that causes a "gap" then the tx will sit in the queue forever waiting for transactions with the missing nonce(s).
- **Fuck!** If a tx fails then it creates a gap and subsequent tx will sit forever until you reissue a good tx with the same bad nonce, then the blocked ones will all go through.
- Transactions cannot be recalled or cancelled.

### Transaction Gas

- Gas is NOT Ether! It's its own currency with its own exchange rate vs. Ether.
- gasPrice in the tx is the exchange rate in wei per gas unit that the originator is willing to pay for gas.
- It's not the max price but the actual bid and a higher bid will up the chances/incentive for the node to process your tx.
- gasLimit is the max units of gas the originator is willing to buy to process the tx.
- Simple payments from EOA to EOA in Ether are fixed at 21,000 gas, though you still need to bid on the gas.
- If the tx involves a smart contract then the gas needed to "run" can only be estimated.
- You're only charged for gas you actually burn.
- **Note!** Addresses in the tx are not validated as actually existing; they may not and payments will succeed regardless.
- A tx with only a value is a payment; a tx with only data is an invocation. See example commands on page 108, 109 using `web3.eth.sendTransaction()`
- Use `address().balance` in Solidity to get your contract's balance.
- data in the tx is an ABI-compatible hex payload laid out:
  - Function selector: first 4 bytes of Keccak hash of the function's prototype (see below).
  - Arguments as per ABI specification encoding.
- Prototype of a function is `functionName(type, type, ...)` hashed, e.g. `withdraw(uint)` hashed and first 4 bytes is `0x2e1a7d4d`
- See example on page 111 for encoding arguments.
- Contract creation is sent to 0x0 a.k.a. "the zero address" with the payload as compiled byte code.
- Set a value only if you want to seed a balance against the contract address.
- The tx receipt contains the new contract address.

### Digital Signatures

- A transaction's Keccak of the RLP-encoded tx is signed, see page 116.
- The final signature has two parts, r and s.
- To verify a tx you need r and s, and the serialized tx and the public key.
- JavaScript in Node sample code using a library can be seen on page 119.
- EIP-155 adds three new fields to the transaction: the chain ID, 0 and 0.
- "Public Key Recovery" is a method which uses v,r,s in the tx to compute the ephemeral public key of the sender of the tx, where v "signature prefix value" is used to determine which of two possible public keys (computed from r and s) is the right one.
- You can sign a tx on one device but then send it from another device via `web3.eth.sendSignedTransaction(...)` so decrypted keys are not in RAM on the network node that sends the tx.

### Transaction Propagation

- Uses "flood routing" protocol, p2p like mesh of equal peers; a tx is validated and forwarded to node's usually ~13 direct neighbours.
- They validate and propagate it to their neighbours, and so on.
- Some nodes are miners (PoW era book).
- Multi-sig is done by smart contract custom code.

### Smart Contracts

- Compiled as bytecode, deployed to 0x0, addressable, no special privileges for the owner, unless added in code.
- Always single-threaded execution, and atomic.
- A failed tx is still recorded has having been attempted on the blockchain.
- A contract cannot be changed but it can be deleted!
- Contracts must call `selfdestruct` which has negative gas to incentivize freeing resources.
- Contract languages include: LLL, Serpent, Solidity, Vyper, Bamboo.
- Solidity is by far the most popular; `solc` compiler https://github.com/ethereum/solidity
- Notes here are all for Solidity 0.4.24
- `solc --optimize --bin faucet.sol`
- Contract ABI is a JSON array of function descriptions and events.
- Function=type, name, inputs, outputs, constant, payable.
- Events=type, name, inputs, anonymous.
- `solc --abi faucet.sol`
- Web3 just needs the ABI and contract address to be able to construct tx to interact with the contract.
- Use the `pragma solidity ^0,4,19;` to ensure the code is not compiled by an out-dated compiler that's not aware of new features.
- Solidity data types: `bool`, `int`, `uint`, `fixed`, `ufixed`, `address`, `byte array`, `Enum`, `Arrays`, `Struct`, `Mapping`, `seconds`, `minutes`, `hours`, `days`, `wei`, `finney`, `szabo`, `ether`.
- `int` and `uint` should be declared with a number of bits size in increments of 8-256 and defaults to 256 if not specified, e.g. uint16.
- `fixed` and `ufixed` need a size (8-256) and the number of decimals, e.g. `ufixed32x2`.
- Byte arrays can be fixed like `bytes1` to `bytes32` or variable size declared as just `bytes` or `string`.
- Arrays e.g. `uint32[][5]` is a fixed array of 5 dynamic arrays of 32-bit unsigned integers.
- Contracts have access to some global variables at runtime: `block`, `msg` and `tx`.
- `msg.sender`, `msg.value`, `msg.gas`, `msg.data`, `msg.sig`
- `tx.gasprice`, `tx.origin` which is **unsafe**.
- `block.blockhash`, `block.coinbase`, `block.difficulty`, `block.gaslimit`, `block.number`, `block.timestamp`.
- `address.balance`, `address.transfer(addr)`, `address.send(addr)`, `address.call(payload)`, `address.callcode(payload)`, `address.delegatecall()`
- Some built-in functions worth noting: `addmod`, `mulmod`, `keccak256`, `sha256`, `sha3`, `ripemd160`, `ecrecover`, `selfdestruct(addr)`, `this`.
- Object types `contract`, `interface`, `library`.
- The library is a contract deployed once, intended to be used by other contracts via `delegatecall(...)`.

### Functions

- Functions can be `public`, `external`, `internal`, `private`.
- **Important** - `private` and `internal` only affect language compilation since **all code and state is public** on the blockchain.
- A function's behaviour can be further marked with `constant`/`view`, `pure`, `payable`.
- Contracts can also have a `constructor() { }` function that runs in the same tx as the contract creation and initializes state etc.
- Contracts can also have a function that calls `selfdestruct(addr)` to tear down resources. It accepts an address to which the EVM pays what is effectively a reward for being good and clearing up.
- It's convention to set `this.owner` field to `msg.sender` in the constructor and then to check this field in the destructor to ensure only the owner can delete the contract.
- `function destroy() public { require(msg.sender = owner); selfdestruct(owner); }`
- You can write function modifiers which are like attributes in C#, e.g. `function destroy() public onlyOwner` where `onlyOwner` is the modifier which is simply an outer function that's called before the inner `destroy`.
- Inheritence is done via `is` keyword e.g. `contract Child is Parent { ... }` and multiple inheritence is via comma separated list.
- It's common to use an `Owned` contract which has a constructor which sets `this.owner = msg.sender` and has a destructor, too.
- Error handling is done via `assert`, `require` and `revert` where `assert` and `require` work the same way but `assert` is used when the condition should always be `true` and `require` is used when it's sometimes expected to be `false`, like checking argument values so `require` can include a friendly message.
- The `revert(msg)` is used like `throw` in other languages.
- Some methods like `transfer` do their own checks but it's worth spending gas to do the check and provide a clear error message.

### Events

- Whether a transaction completes successfully or not, it returns a receipt which contains log entries which are constructed from events.
- Events take args that are serialized and recorded in the transaction logs on the blockchain.
- Supplying the keyword `indexed` _before_ an argument makes the value part of a hashtable that can be "searched by an app" (!?)
- `event Deposit(address indexed from, uint amount);`
- Use `emit Deposit(msg.sender, msg.value);`

### Calling Other Contracts

- **Potentially dangerous** because you could be calling into a dodgy contract _and_ a dodgy contract could be calling your code!
- Safest to instantiate the other contract yourself using `new` and the `import` statement which references the source code.
- This actually creates the other contract on the blockchain! So you must remember to destroy it again from your destructor method.
- **Note** - the owner of the other contract will be the contract that created it, not the EOA that create the transaction.
- Another way is to cast the address of an existing contract to a known interface.
- **Important** - it is vital you are sure of the type of the existing contract, e.g. when an address is passed in to a method, that address might be for a dodgy contract and the caller who passed it in could be The Devil Incarnate.
- **Most dangerous of all** is to use low-level `call` and `delegatecall` etc. to invoke a method at an address and pass raw arguments in! It'll return `false` is there's a problem but this method opens your code up to _reentrancy attacks_, see page 173.
- With `delegatecall` the context is preserved so the executing code thinks it's in the context that called it, the `msg.sender` is maintained.
- `delegatecall`, page 155, is most commonly used to call library code, but if that code is not designed to be a library then you could open a gateway to Hell.
- Good example of calling other code and the varying values of `tx.origin`, `msg.sender` and `this` for four different ways, on page 157.

### Gas Considerations

- More details on page 314.
- When all consumed, "out of gas" exception is thrown and everything reverts and gas is taken.
- Avoid dynamically sized arrays.
- Avoid calls to other contracts.
- Example JS (I think) so estimate gas cost of a contract method on page 159.
- Seems that Truffle can run JS files.

### Vyper

- Vyper, as of 2019 book, has stuff from Solidity omitted to make it safer.
- Not in the book but on quick research, Vyper may produce less gassy bytecode.
- Modifiers: (attributes) omitted since they essentially hide logic. I thought this, too!
- Inheritence: I also thought composition is better than inheritence and esp. didn't like multiple inheritence!
- Inline Assembly: low-level access would be a hole in Vyper's ethos. Obvs!
- Function Overloading: confusion over which function is called at runtime. This makes sense!
- Recursive Calling: when recursion is allowed, there's always scope for gas limit attack.
- Infinite-length Loops: as above, really.
- Binary Fixed Point: decimal fixed point is better since it doesn't approximate and lead to numeric errors.
- Vyper enforces member ordering like Python.
- Vyper has built-in overflow protection and an equivalent SafeMath library.

### Smart Contract Security

- **Complexity is the enemy of security**.
- Aim for as few lines as possible.
- Reuse code: do not reinvent the wheel.
- Aerospace engineering quality.
- Readability == auditability.
- Test coverage.

#### Reentrancy

- You are **urged** to study such attacks online and learn from others' mistakes.
- Can occur when a contract sends Ether to an unknown address.
- Attacker crafts a dodgy contract containing malicious code in its fallback address and runs the attacker's code.
- Typically the malicious code executes a function on the vulnerable contract.
- It's called "reentrant" because the malicious contract calls back into the very contract that invoked it.
- Page 173 provides an example.
- **Do** use the `transfer` method to send ether because it limits gas to 2,300 which limits the cycles.
- **Do** ensure the transfer is done after any other state changes such as deducting from their pot or anything else that the transfer is contingent on, such as daily withdrawal limits. This is known as the checks-effects-interactions pattern.
- **Consider** introducing a mutex so that the contract can only be used one at a time.

#### Arithmetic Over and Underflows

- Solidity does **not** check for overflows or underflows like most programming languages and so setting a `uint8` to 256 will set it to 0.
- Math code can easily overflow a variable and make it "go around the clock" like a car tachometer, esp. when subtracting.
- **Always** validate user inputs will not over/underflow a variable, esp. ensure they are positive numbers.
- **Do** read and re-read best practice on the web for preventing these attacks.
- **Do** use trusted libraries for doing math and **don't** reinvent the wheel, so avoid `+=` esp. on small data types, see "SafeMath" lib.
- **Remember** that all state is public so attackers can see the current value of a field and pass in the "right" number to over or underflow a variable.
- **Consider** using a language that has built-in checks.
- `pots[msg.sender] = pots[msg.sender].add(msg.value);`

#### Unexpected Ether

- Contract code that assumes and checks that fields tracking the total ether sent to the contract address are equal are open to attack by hackers sending ether to the contract via two other possible means; via the `selfdestruct(addr)` reward and to a less likely extent by an attacker preempting the contract's address before it is created which is possible because the contract address is deterministic.
- Such checks can thus be made to fail and cause the contract to malfunction and potentially brick the contract so funds are stuck in it.
- **Avoid** using `this.balance` and other ways to read the blockchain balance data, but instead account for ether using fields in your contract.

#### DELEGATECALL

- Smart Contracts store data in indexed slots from [0] onwards.
- When a contract is used like a library, any state in fields will also be allocated to slot beginning from [0].
- This means **state will conflict** causing all kinds of bugs and security holes.
- **DO NOT** use a normal Smart Contract i.e. `contract` keyword for a library.
- DO use the `library` keyword which causes the compiler to prevent the unit from having state; forces statelessness.
- Using a normal contract as a library means all its functions are exposed and callable like any other contract.

