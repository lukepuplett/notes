# Developing a simple blockchain pplication with a web UI

Notes taken from:

https://dev.to/dabit3/the-complete-guide-to-full-stack-ethereum-development-3j13?signin=true

See also:

https://dev.to/dabit3/building-graphql-apis-on-ethereum-4poa

### Hardhat

Development environment and framework. Other's are Ganache, Truffle and Foundry.

### Ethers.js

Javascript library for interacting with the Ethereum chain. I'm not sure if it goes direct or via some insecure middleman HTTP API. Another option is **web3.js**

### Metamask

Browser plug-in wallet. Helps with account management and connecting the current user to the chain. I think it adds `window.ethereum` to the browser JS namespace. Can pop-up requests to sign transactions for the user.

### The Graph

Exposes GraphQL in a decentralized way that makes querying the chain much simpler than doing it directly.

### Configuring Hardhat

Run Hardhat directly from its NPM package using `npx hardhat` and choose to create a sample project. It'll make some files and folders with sample stuff in then, including a deploy script.

The default `hardhat.config.js` needs changing to set `module.exports.networks.hardhat.chainId` to `1337`.

### Smart contract

The code looks like JavaScript and imports other `.sol` files. It seems to use a `contract` keyword instead of `class` and it can have private fields. The constructor appears to accept arguments to initialize field values with, I guess at deploy time. I think the constructor runs at deploy.

At the very least this is state on a blockchain. I don't know where it's stored or what the length limits are. The template code doesn't truncate it, which is odd for a public app; someone could send 1G string.

Writes seem to be considered as transactions and the gas fee must be paid. Presumably reading is free because it requires no validation work from the network.

A contract is compiled to produce an **Application Binary Interface** which seems to be like IDL from Windows COM. Hardhat will make the ABI, else you can find them on Etherscan.

Run `npx hardhat compile` and it'll create `src/artifacts/contracts/MyContract.json` which has the ABI as a property. Being JSON, it is then imported into your JavaScript client like so:

    import MyContract from './artifacts/contracts/MyContract.sol/MyContract.json

And reference it like so:

    console.log("My ABI:", MyContract.abi)
    
Learn about Human-Readable ABIs in Ethers.js here: https://blog.ricmoo.com/human-readable-contract-abis-in-ethers-js-141902f4d917

Hardhat has a built-in EVM node emulator. Run `npx hardhat node` and it'll spew 20 test account addresses and private keys onto the console, each with 10,000 fake ETH.

Rename `sample-script.js` to `deploy.js` and run it:

    > npx hardhat run scripts/deploy.js --network localhost
    MyContract deployed to: 0xdeadbeef

**Note** - this runs as the first (of 20) test account.

