# Foundry

A Solidity smart contract development toolset written in Rust by Paradigm and Ithaca.

## Basics

```
curl -L https://foundry.paradigm.xyz | bash
# Install forge, cast, anvil, chisel
foundryup
```

- Use `forge` for build, test, debug, deploy and verify smart contracts.
- Use `anvil` as a local development blockchain node which I think can clone a live chain, serves JSON-RPC.
- Use `cast` for interacting with onchain apps, e.g. `cast call 0x1... "balanceOf(address)" 0x2... --rpc-url https...`
- Use `chisel` as a Solidity REPL.

### Forge

- `forge init Counter` created a project folder, cloned down a Foundry Git repo with standard bits and pieces in called forge-std, then creates a bunch of subfolders like lib, script, src, test and a `foundry.toml` and a lock file.

- It also created a scaffold for a Counter smart contract with `setNumber` and `increment` functions.

- The forge-std repo in `Counter/lib/forge-std/` is structured like a Foundry project itself, with its own scripts, src etc. folders. Inside `src` are a bunch of Solidiity files and an interfaces folder with e.g. `IERC721.sol`

- The `forge build` command, compiled 23 files and downloaded and installed the Solc compiler to do this. It also warned on `/Users/lukepuplett/.foundry/cache/signatures` being missing.

- I later realised that I wasn't supposed to see the verbose debug of what was happening but a Rust logging env var was still set from my Rust coding time weeks ago.

- The `out` directory contains the contract artifact, such as the ABI.

- Interestingly, the `forge test` command actually ran two tests defined in `Counter/tests/Counter.t.sol` one for each of the functions in the contract, including a fuzz of SetNumber.

- The `forge` tool enables table tests, fuzzing, invariant testing, coverage reports and gas tracking.

- In the docs, early on is an invitation which I didn't take up:

```bash
# Run tests against live chain state by forking
forge test --fork-url https://reth-ethereum.ithaca.xyz/rpc
```

- **Important** - Foundry calls a copy of chain state a 'fork' which is confused with a chain's feature or 'protocol fork', e.g. Shanghai.

- I suspect I could stand up a Sepolia local node and then fork from that to get a private space to play with stuff deployed to it, like EAS, and without going via an RPC service.

- There is no built-in base fork or version for Foundry's node, in the state sense - i.e. it is clean.

- Deploying contracts appears to be done via scripts written in normal Solidity and the init command has created a demo `Counter/scripts/Counter.s.sol` which is a single contract deriving from `Script` and defines `setUp` and `run` functions. Deployment is thus a case of running the script with some args.

```bash
# Use forge scripts to deploy contracts
# Set your private key
export PRIVATE_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
 
# Deploy to local anvil instance
forge script script/Counter.s.sol --rpc-url http://127.0.0.1:8545 --broadcast --private-key $PRIVATE_KEY
```

- There's a linter, `forge lint` though when run I just see DEBUG spam.

- Can use the following to clone a verified contract into a new subfolder WETH9 and configure it as a Foundry project:

```bash
forge clone 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2 WETH9 --etherscan-api-key <YOUR_API_KEY>
```

- Forge manages dependencies using git submodules by default, though Soldeer is the new thing. I devote a section to Soldeer, below.

- To install a dependency, use `forge install` which will pull the GitHub repo `vendor/repo` with optional `@v0.1.2` suffix and stage the `.gitmodules` file and make a commit. The submodule is put in the `lib` folder.

- There's a `remappings.txt` of `key/=folder/path/` pairs to make import statements shorter or pointed at the right version, and these can also be set in foundry.toml.

- Conflicts in remappings and git submodules can be resolved with syntax like this:

```
lib/lib_1/:@openzeppelin/=lib/lib_1/node_modules/@openzeppelin/
lib/lib_2/:@openzeppelin/=lib/lib_2/node_modules/@openzeppelin/
```

- Use `forge update lib/subfolder` to update a particular dep, and `remove` to get rid.

- Forge also supports Hardhat-style projects where dependencies are npm packages (stored in node_modules) and contracts are stored in contracts as opposed to src. It says "To enable Hardhat compatibility mode pass the `--hh` flag." but doesn't say on which command.

- Can configure the paths for libs and contracts via CLI params or the `foundry.toml` file.

- Using `--lib-paths node_modules --contracts contracts` gives HardHat support.

- Can use `extends` in `foundry.toml` to (single level) inherit from a shared base config using a relative path, see docs.

- The docs hint at profiles via `[profile.default]` and `[profile.production]` though I have not otherwise seen this mentioned.

### Anvil

- Start a local node with `anvil` no switches.

- Forking mainnet looks very simple and appears to need just an RPC URL. God knows how long it takes.

```bash
anvil --fork-url https://reth-ethereum.ithaca.xyz/rpc
```

- The docs mention custom `anvil_` methods, account impersonation, state manipulation and mining control, and forking.

### Cast

- They describe it as a Swiss Army Knife for interaction, calls, transactions, getting any type of chain data.

- We saw the balanceOf example above, and there's a way to get balances via `balance`, as well as `send`:

```bash
cast balance vitalik.eth --ether --rpc-url https://reth-ethereum.ithaca.xyz/rpc
cast send 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 --value 10000000 --private-key $PRIVATE_KEY
```

- Running all these command line tools so far results in loads of DEBUG message spam.

- Cast also has `rpc` action which takes a JSON-RPC method and args, and `--rpc-url` arg and a `block-number` action and `2h` action which appears to convert numbers to hex perhaps.

- There's a short section on Chisel, the REPL, but I cannot imagine using it.

- Has an extensive LLM demo prompt for good use of AI to write contracts. This converts naturally to a .cursorrules file by Cursor itself.

- The spam was found to be where I had RUST_LOG env var from my Rust coding time.

### Soldeer

- Soldeer is new and still early in December 2025 as I write. It has its home at soldeer.xyz, is written in Rust and aims to be a package manager for Solidity.

- However, according to these docs at least, it is initialised via Forge, `forge soldeer init` which I assume is like `npm init` and can be done at any project stage.

- There's a package search site and installing is via `forge soldeer install @openzepellin-contracts~5.0.2`.

- It doesn't use lib but instead a new `dependencies` subfolder.

- Define a `[dependencies]` section in `foundry.toml` or it can use `soldeer.toml`.

- Use `--url` to specify a dependency outside of the official Soldeer registry, and it'll appear like this in config:

```toml
[dependencies]
"@custom-dependency" = { version = "1.0.0", url = "https://my-website.com/custom-dependency-1-0-0.zip" }
```

- Similarly, you can use `--git` with `--rev`, `--tag`, `--branch` though Git is discouraged.

- Use `forge soldeer update` on a pulled repo and it'll install all the deps in the TOML file, like `npm install` I guess.

- `uninstall` will clear up the dependencies section, `soldeer.lock` and remappings.

- Further configuration can be made to the remappings functionality via a `[soldeer]` TOML section.

- The strangest design choice is that sub-dependencies don't seem to be installed automatically(!) so you must remember the `--recursive-deps` arg or set `recursive_deps = true` in the `[soldeer]` section.

- Like NPM, you can publish reusable stuff but I won't explain all that here.

- If you use other dependency managers, such as git submodules or npm, ensure you don't duplicate dependencies between soldeer and the other manager.

- Remappings targeting dependencies installed without Soldeer are not modified or removed when using Soldeer commands, unless the --regenerate-remappings flag is specified or the remappings_regenerate = true option is set.

## Scripting in Solidity

- https://getfoundry.sh/guides/scripting-with-solidity/

- Docs say that `forge create` is limited and not user friendly vs. scripting.

- Scripts are run on EVM so they can do EVM type stuff.

- They have four phases;
  - Local Simulation - any external call (non-static, non-internal) from a vm.broadcast and/or vm.startBroadcast will be appended to a list (of transactions, I think). This is done in rpc/fork context, if a URL is supplied.
  - Onchain Simulation - Optional - If rpc/fork URL supplied then collected transactions are executed sequentially here.
  - Broadcasting - Optional - If `--broadcast` specified, then transactions are broadcast (to where?)
  - Verification - Optional - If `--verify` specified and there's an API key, it'll try to verify the contract via e.g. etherscan.

- Use `--resume` to continue on after failures or time-outs.

- The following shows how to wrap three function calls so that they're broadcast for real, later.

- By default, scripts are executed by calling the function named `run`, our entrypoint.

- Follows is an example to deploy the demo Counter contract:

```
forge script --chain sepolia script/Counter.s.sol:CounterScript --rpc-url $SEPOLIA_RPC_URL --broadcast --verify -vvvv --interactives 1
```

```solidity
vm.startBroadcast();
myContract.setValue(42);
myContract.doSomethingElse();
myContract.anotherCall();
vm.stopBroadcast();
```

- Thus, to deploy a contract, the instantiation 'newing up' of the contract from another (the script) is how it works:

```solidity
vm.startBroadcast();
counter = new Counter(); // instatiation is simulated and then replayed
vm.stopBroadcast();
```

- The broadcast transaction logs will be stored in the `broadcast` directory by default. You can change the logs location by setting broadcast in your foundry.toml file.

- The broadcasting sender is determined by checking the following in order:
  - If `--sender` argument was provided, that address is used.
  - If exactly one signer (e.g. private key, hardware wallet, keystore) is set, that signer is used.
  - Otherwise, the default Foundry sender (`0x1804c8AB1F12E6bbf3894d4083f33e07309d1f38`) is attempted to be used.

- I don't get this. It's not clear. We can specify a private key in an .env file, or by using `--interactives 1` and giving it on a terminal prompt for input, or via a hardware wallet and other options via command line, but then there is the `--sender` argument. Perhaps it means that when there are many private keys around, then the one to use can be set using `--sender`?

- It alludes to using a one-time private key for deployment. This makes sense if you consider that the deployer account doesn't matter long term, if your deployer happens to set an owner or admin during construction which would be a public key/address. So you'd deploy but then the owner/admin would be some other account and the deployer can be forgotten.

- Previously failed transactions can be retried using `--resume`.

- **Note** - due to frontrunning etc. the results may differ in reality from that simulated.

- Foundry automatically loads 'an .env file' if present in your project folder.

- The environment variables can be referenced by values in `foundry.toml` using `${ENV_KEY}` syntax, so this is how to configure things like private keys.

- I don't think it is safe to put private keys in `.env` files and they acknowledge this and how Foundry supports various wallets. They link to a documentation page which really just details command line arguments which include various options for wallets but there's no explainer.

https://getfoundry.sh/forge/reference/script/#wallet-options---raw

- To deploy a script, you'll need the following TOML settings:

```toml
[rpc_endpoints]
sepolia = "${SEPOLIA_RPC_URL}"
 
[etherscan]
sepolia = { key = "${ETHERSCAN_API_KEY}" }
```

- This creates a **RPC alias** for Sepolia and loads the Etherscan API key. However this does not affect the `getChain` method.

- Apparently for complex multi-chain deployments they can use a standard premade script called **Config** in the forge-std library which centralizes config in TOML files.

https://getfoundry.sh/guides/scripting-with-config/

- 