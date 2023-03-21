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

