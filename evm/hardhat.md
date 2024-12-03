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

