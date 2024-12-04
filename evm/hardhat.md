# Hardhat

## Hardhat Runner

 - Main thing you interact with, it's a task runner, it's all based around **tasks** and **plugins**.
 - E.g. `npx hardhat compile` runs the built-in `compile` task.
 - Task can call other tasks, and "Users" and plugins can override tasks.

## Installation

 - Nomic has a Solidity plugin for VSCode.
 - It's an NPM app.
 - Make a new folder and `npm init` then `npm install --save-dev hardhat`.
 - Use `npx` to run `npx hardhat init`.
 - They recommend using TypeScript and WSL for Windows.
 - `npx hardhat` provides usage docs.

## Project Setup

 - Looks like contracts go in a `contracts/` subfolder.
 - `npx hardhat compile` from the project root, will also make TS bindings via TypeChain.
 - Project initialized with Mocha, Chai, Ethers.js and Hardhat Ignition.
 - Tests are in `test/` subfolder and are TypeScript code files.
 - Appears to use a `describe(name, fn())` function to bundle tests for a contract by `name`.
 - Within that function, an optional initial async function stands as a test fixture to deploy the contract, presumably once per test?
 - Then there are further calls to `describe(name, fn())` where the function body is the actual test code.
 - These appear to use `it()` and `expect()` or `await expect()` to build the test up.
 - `npx hardhat test` from project root with run it.

## Deploying Contracts

 - Uses **Hardhat Ignition** module which has `.ts` files in `ignition/modules` sub.
 - Each code file in that folder seems to be responsible for deploying contracts.
 - `buildModule(name, (m) => {})` seems to prepare values for the contract constructor which is then default exported.
 - `npx hardhat ignition deploy ./ignition/modules/Lock.ts` then deploys.
 - Of course there's nowhere to deploy it to! Run `npx hardhat node` to run a local blockchain enumlator and EVM.
 - This local node will expose a JSON-RPC service at `http://localhost:8545` to which you can connect your wallet.
 - Run `npx hardhat ignition deploy [path] --network localhost` to deploy to it (not sure if that network name must be in `hardhat.config.ts`).
 - Can run `npx hardhat node --hostname 0.0.0.0 --port 8545` to expose to public IPs.

## Config

 - The config is just TypeScript which is executed before any task, so can do other setup.
 - `hardhat.config.ts` has various sections all detailed in the official docs.
 - Networks are JSON-RPC services and can be configured with `url`, `chainId`, `from` (`msg.sender` address, else first account of node), `gas`, `gasPrice`, `gasMultiplier`, `accounts` (it can use "the node's accounts"?!), `httpHeaders`, `timeout`.
 - The default EVM version is determined by the Solidity version but can override.

## Testing

 - Uses ethers.js to connect to Hardhart Network, and on Mocha and Chai for the tests.
 - Not going to repeat everything that's in the HH docs.
 - May need to pull in helpers for reading from the HH network, blockchain:

 ```js
 import { time } from "@nomicfoundation/hardhat-toolbox/network-helpers";
 ```

 - `time.latest()` returns the timestamp of the last mined block.
 - To deploy, use normal `ethers.deployContract` passing an object with transaction params.
 - The contract instance is returned which has methods which can be called to check stuff.

 ```js
 expect(await myContract.someGetter()).to.equal(someValue);
 ```

 - The following code shows how to test a function reverts:

 ```js
 it("Should revert with the right error if called too soon", async function () {
  // ...deploy the contract as before...
  await expect(lock.withdraw()).to.be.revertedWith("You can't withdraw yet");
});
```

**Note** - In the above, the whole expect is awaited because "it has to wait until the transaction is mined." The second mutating function needs a transaction to be executed while the first is just a simple "call".

 - You can alter the global state of the network, e.g. the block time, `await time.increaseTo(unlockTime);`.
 - Since we're just using ethers.js then you can get other signers for testing stuff with.

```js
const [owner, otherAccount] = await hre.ethers.getSigners();
```

 - Import `loadFixture` and use that in each test to reset and rerun setup needed.
 - You can test which events are emitted, too: https://hardhat.org/hardhat-chai-matchers/docs/reference#.emit
 - Coverage is done via command line, `npx hardhat coverage`
 - Reporting gas `REPORT_GAS=true npx hardhat test`
 - Can run in parallel but can cause issues, see docs.

 