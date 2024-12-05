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
 - To manually setup TypeScript stuff in your project, like so:

 ```bash
 npm install --save-dev ts-node typescript
 npm install --save-dev chai@4 @types/node @types/mocha @types/chai@4
```

 - `npx hardhat` provides usage docs.

**Note** - Hardhat won't auto type check when any tasks are run, so you'll need to use `--typecheck` for that.

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
 - Because it's just JavaScript, you can read environment vars and use packages.

 ```js
 const { vars } = require("hardhat/config");
 const ETHERSCAN_API_KEY = vars.get("ETHERSCAN_API_KEY", "fallback");
 ```

  - I think this is using Hardhat's own config vars feature, which is stored in plaintext on the local disk.
  - Run `npx hardhat vars path` to find the path.
  - There's also `set`, `get`, `list`, `delete` and `setup` which shows all those in use and needing values.
  - Does not recommend using popular **dotenv** due to risk of accidental `.env` file upload.
  Local environment vars override configuration vars, so this works:

```bash
HARDHAT_VAR_MY_KEY=123 npx hardhat some-task
```

## Testing

 - Run them all with `npx hardhat test` and maybe add  `--typecheck`.
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

## Deploying Contracts

 - Declarative JS **Hardhat Ignition** as we know.

```js
  const lock = m.contract("Lock", [unlockTime], {
    value: lockedAmount,
  });
```

 - I think the object passed in `{ value: x }` is setting fields on the raw transaction.
 - Remember to start the network, first: `npx hardhat node`

```bash
npx hardhat ignition deploy ./ignition/modules/Lock.ts --network localhost
```

 - That's it.

## Verifying Contracts

 - To **Verify** a contract means to publish its source code and compiler settings.
 - This is sent to Etherscan so people can create the same exact bytecode.
 - Need an API key from Etherscan to store as a **configuration variable**.
 - Store this using `npx hardhat vars set ETHERSCAN_API_KEY` which prompts for the value.
 - Sepolia is the main dev-test network, and will need a JSON-RPC proxy like Alchemy, or own node.
 - Need to send some Sepolia ether to address making the deployment; get this from a faucet.

**Note** - I'm not sure what that means, the "address making the deployment".

 - The sample code in the docs is already verified so it'll error, so you need to add something unique, like a comment, to the code.
 - Run this:

```bash
npx hardhat ignition deploy ignition/modules/Lock.ts --network sepolia --deployment-id sepolia-deployment
```

 - The `--deployment-id` flag is optional way to name the deployment to refer to later.
 - To verify the contract, run:

```bash
npx hardhat ignition verify sepolia-deployment
```

 - This does both, combined:

```bash
npx hardhat ignition deploy ignition/modules/Lock.ts --network sepolia --verify
```

 - If you see an error that the address has no bytecode, it usually means Etherscan has not indexed it yet.

## Writing Tasks

 - Hardhat is a task runner with built-in `compile` and `test`.
 - You can author your own custom tasks.
 - The following is a custom task sample which should go in your hardhat config file:

```js
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});
```

 - And you can run it with:

```bash
npx hardhat accounts
```

 - The `hre` object is the Hardhat Runtime Environment, and its properties are all injected into the global namespace, too, maybe...

**Note** - It also says "When using TypeScript nothing will be available in the global scope and you will need to import everything explicitly using" in another part of the docs.

 - TypeScript users need `import { ethers } from "hardhat"` to import stuff into the global scope.
 - You can do anything you want in the task function.

**Note** - Hardhat won't auto type check when any tasks are run, so you'll need to use `--typecheck` for that.

## Using TypeScript

 - For Hardhat to generate types for your contracts, install and use `@typechain/hardhat` which generates typing files (*.d.ts) based on ABIs.
 - But if you installed `@nomicfoundation/hardhat-toolbox` you don't need typechain.
 - See this for TS path mapping: https://hardhat.org/hardhat-runner/docs/guides/typescript#support-for-path-mappings
