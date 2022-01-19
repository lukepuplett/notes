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

