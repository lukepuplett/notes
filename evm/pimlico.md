# Pimlico

Ethereum EVM service for hosted bundler and paymasters.

## Resources

- Website: https://pimlico.io
- Documentation: https://docs.pimlico.io

## Overview

Pimlico provides infrastructure services for Ethereum, specifically:
- **Bundler**: Hosted bundler service for EIP-4337 (Account Abstraction)
- **Paymaster**: Paymaster service for sponsoring user transactions

## API

### Bundler
- Uses **Alto** bundler (OSS implementation, hosted by Pimlico)
- Dev/test endpoint (no API key required): `https://public.pimlico.io/v2/{chain_id}/rpc`
- JSON-RPC methods:
  - `eth_sendUserOperation`
  - `eth_estimateUserOperationGas`
  - `eth_getUserOperationReceipt`
  - `eth_getUserOperationByHash`
  - `eth_supportedEntryPoints`
  - `pimlico_getUserOperationGasPrice`
  - `pimlico_getUserOperationStatus`
- **Recommended**: Use `pimlico_getUserOperationStatus` endpoint to check the status of your user operation

#### Bundler Payment Mechanism

How onchain bundler payment works:
1. When constructing your user operation, use `pimlico_getUserOperationGasPrice` to determine the totals and limits for your user operation
2. When submitting to the bundler, the bundler bundles the user operation into a transaction with a gas price where the surcharge is stripped
3. The bundler makes profit from the small difference between: user operation totals and limits vs. bundler transaction gas price (surcharge stripped)

#### Bundler Surcharges (Tips)

Bundlers adding a surcharge/tip to gas fees is normal and necessary:
- **Purpose**: Compensates bundlers for infrastructure costs, L1 gas fees paid upfront, price volatility risks, and priority fees paid to block builders
- **Impact on Fees**: The bundler's RPC estimates (via `eth_estimateUserOperationGas`) generally include this tip automatically
- **Recommendation**: It's often recommended to add a small manual buffer (10-25%) to the `maxPriorityFeePerGas` to ensure timely inclusion, especially during network congestion

### Paymaster
- URL: `https://api.pimlico.io/v2/{chain}/rpc?apikey=[YOUR_API_KEY_HERE]`
- Supported chains:
  - `base` (Base mainnet)
  - `base-sepolia` (Base Sepolia testnet)
  - `unichain` (Unichain mainnet)
  - `unichain-testnet` (Unichain testnet)

#### Paymaster URL Usage

The paymaster URL acts as a JSON-RPC endpoint used by clients (Viem) and bundlers to interact with the paymaster service:
- **Viem/Client**: Uses the URL with a `paymasterClient` to call methods like `pm_getPaymasterData`. The client receives the necessary `paymasterAndData` field to complete the UserOperation
- **Flow**: dApp -> Paymaster URL -> Bundler URL -> EntryPoint Contract

### Request Routing

**Assumption**: The service handling the URL route determines how to handle paymaster traffic vs bundler requests (likely based on JSON-RPC method called).

## Errors and Codes

Various errors and codes documented in their docs. (See documentation for details)

## Local Development / Testing

For dev/test/CI/CD environments:
- **Bundler**: Can use **Prool** bundler
- **Alto**: Can be used via **Anvil** (part of Foundry)

## Integration

Pimlico recommends using **permissionless.js** to interface with their Paymaster and Bundler services.

### Using Bundler RPC Methods in Viem

You must use your specific bundler's RPC methods to ensure fees are calculated correctly and adhere to their requirements.

**Viem Implementation**:
- Use `createBundlerClient` with your bundler's specific URL

**Key Methods**:
- `estimateUserOperationGas`: Fetches required `callGasLimit`, `verificationGasLimit`, etc.
- `getUserOperationGasPrice`: (or custom methods like `pimlico_getUserOperationGasPrice`) Retrieves optimized `maxFeePerGas` and `maxPriorityFeePerGas`
- **Helper**: Viem's `prepareUserOperation` helper automates these calls internally

#### How Viem Maps Actions to JSON-RPC Methods

Viem knows which RPC method to call because its "actions" are explicitly mapped to specific JSON-RPC method names internally:
- For standard ERC-4337 methods, Viem has built-in support
- For a bundler's special prefixed methods, you may need to use a dedicated client from a specialized SDK (like permissionless.js or a specific bundler's SDK) or manually extend the Viem client

**Standard ERC-4337 Methods**:
- When you use the Viem BundlerClient's standard actions (e.g., `estimateUserOperationGas`), Viem automatically translates this into the corresponding standard JSON-RPC method name defined in the EIP: `eth_estimateUserOperationGas`
- Example:
  - Viem Action: `await bundlerClient.estimateUserOperationGas(...)`
  - JSON-RPC Call (under the hood): `{ "method": "eth_estimateUserOperationGas", ... }`

**Special Prefixed (Custom) Methods**:
For bundlers that use non-standard, custom-prefixed methods (e.g., `pm_...` for paymaster interactions or `pimlico_...` for specific gas price estimates), you have a few options:

1. **Use Specialized SDKs**: Libraries like permissionless.js, which is built on top of Viem, provide type-safe client extensions for these specific methods. You use actions like `await pimlicoClient.getUserOperationGasPrice()` which the library internally maps to the correct custom RPC method (`pimlico_getUserOperationGasPrice`).

2. **`paymaster: true` Configuration**: If your bundler endpoint also supports the paymaster `pm_` methods, you can configure your bundler client with `paymaster: true`. Viem will then assume those methods are available at that URL and use them when needed.

3. **Manual `client.request`**: Viem clients expose a generic `request` method that allows you to manually specify any custom RPC method string and its parameters. This bypasses Viem's built-in actions.
   ```typescript
   const gasPrice = await bundlerClient.request({
     method: 'pimlico_getUserOperationGasPrice', // The exact custom method name
     params: []
   });
   ```

**Summary**: Viem uses predefined mappings for standard actions. For custom ones, you either rely on an extension library that defines the mapping or use the generic `request` method to specify the exact method name string provided by your bundler's documentation.

## Why Pimlico

Chosen over competitors because:
- Easiest to setup
- Simplest documentation

## Competitors

- **Alchemy**: https://www.alchemy.com
- **Infura**: https://www.infura.io

## Related

- See [[eip-4337.md]] for Account Abstraction details
- See [[eip4337-blog-notes.md]] for additional EIP-4337 notes
