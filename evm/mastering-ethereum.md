. . , .

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

- 
