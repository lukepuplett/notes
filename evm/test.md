## RPC Methods for ERC-4337: A Comprehensive Guide

This section provides a detailed specification for the RPC methods introduced by ERC-4337, enabling interaction with the Account Abstraction infrastructure. These methods cater to both standard user operations and debugging functionalities.

### 1. eth Namespace Methods:

These methods primarily handle user-facing operations related to UserOperations.

**1.1. eth_sendUserOperation:**

* **Purpose:** Submits a UserOperation to the client's UserOperation pool for processing and potential inclusion in a bundle.
* **Parameters:**
* `UserOperation` (Object): A complete UserOperation structure with all fields as hex values. Empty fields like `initCode` should be "0x". Ensure either both `factory` and `factoryData` are present or both are absent. Similarly, all paymaster fields should either be present or absent together.
* `EntryPoint` (Address): The address of the desired EntryPoint contract to process this UserOperation. This MUST be one of the addresses returned by the `eth_supportedEntryPoints` RPC call.
* **Return Value:**
* **Success:** `userOpHash` (String): The calculated hash of the UserOperation if it passes validation and is accepted into the pool.
* **Failure:** Error object with `code` and `message` fields:
* `code: -32602`: Invalid UserOperation structure or field values.
* `code: -32500`: Transaction rejected during simulation by the EntryPoint, either during account creation or validation. The `message` field should contain the EntryPoint's "AAxx" error message.
* `code: -32501`: Transaction rejected by the paymaster's `validatePaymasterUserOp` function. The `message` field should ideally contain the revert message from the paymaster. The `data` field MUST include the `paymaster` address.
* `code: -32502`: Transaction rejected due to opcode validation failure.
* `code: -32503`: UserOperation time-range invalid. Either the wallet or paymaster returned a time-range that is already expired or expires too soon. The `data` field should include the `validUntil` and `validAfter` values, and optionally the `paymaster` address if the error originated there.
* `code: -32504`: Transaction rejected because the paymaster or signature aggregator is throttled or banned. The `data` field should include either the `paymaster` or `aggregator` address.
* `code: -32505`: Transaction rejected because the paymaster or signature aggregator has insufficient stake or an inadequate unstake delay. The `data` field should include the `paymaster` or `aggregator` address, along with the `minimumStake` and `minimumUnstakeDelay` values.
* `code: -32506`: Transaction rejected because the wallet specified an unsupported signature aggregator. The `data` field should include the `aggregator` address.
* `code: -32507`: Transaction rejected due to wallet signature verification failure. This could also be due to paymaster signature failure if the paymaster uses its data as a signature.

**1.2. eth_estimateUserOperationGas:**

* **Purpose:** Provides gas estimations for a given UserOperation, including pre-verification, verification, and execution costs.
* **Parameters:**
* Same as `eth_sendUserOperation`.
* Gas limit and price parameters are optional. If provided, they are used for estimation. If not, `maxFeePerGas` and `maxPriorityFeePerGas` default to zero, implying no payment requirement for simulation.
* Optionally accepts the "State Override Set" (refer to `eth_call` documentation) to modify state during gas estimation.
* **Return Values:**
* `preVerificationGas` (Quantity): Gas overhead for the UserOperation itself.
* `verificationGasLimit` (Quantity): Estimated gas limit required for validating the UserOperation.
* `paymasterVerificationGasLimit` (Quantity): Estimated gas limit for paymaster verification (only if a paymaster is specified).
* `callGasLimit` (Quantity): Estimated gas limit for the internal account execution.
* **Error Codes:**
* Same as `eth_sendUserOperation`.
* Additionally, errors may occur if the internal call to the account contract or the paymaster's `postOp` call reverts.

**1.3. eth_getUserOperationByHash:**

* **Purpose:** Retrieves a UserOperation based on its hash (`userOpHash`) obtained from `eth_sendUserOperation`.
* **Parameters:**
* `hash` (String): The `userOpHash` value.
* **Return Value:**
* **Included in block:** Full UserOperation object, including `entryPoint`, `blockNumber`, `blockHash`, and `transactionHash`.
* **Pending in mempool:** May return `null` or a full UserOperation object with the `entryPoint` field and `null` values for `blockNumber`, `blockHash`, and `transactionHash`.
* **Not found:** `null`.

**1.4. eth_getUserOperationReceipt:**

* **Purpose:** Fetches the receipt for a UserOperation given its hash (`userOpHash`).
* **Parameters:**
* `hash` (String): The `userOpHash` value.
* **Return Value:**
* **Not yet included:** `null`.
* **Included in block:** An object containing:
* `userOpHash` (String): The UserOperation hash.
* `entryPoint` (Address): The EntryPoint address.
* `sender` (Address): The sender's account address.
* `nonce` (Quantity): The UserOperation nonce.
* `paymaster` (Address): The paymaster address (or empty if none).
* `actualGasCost` (Quantity): The actual amount paid for the UserOperation.
* `actualGasUsed` (Quantity): Total gas used, including pre-verification, creation, validation, and execution.
* `success` (Boolean): Indicates whether execution completed without reverting.
* `reason` (String): Revert reason if the execution failed.
* `logs` (Array): Logs generated by this UserOperation (excluding logs from other UserOperations in the same bundle).
* `receipt` (Object): The TransactionReceipt object for the entire bundle.

**1.5. eth_supportedEntryPoints:**

* **Purpose:** Returns an array of EntryPoint contract addresses supported by the client.
* **Parameters:** None.
* **Return Value:** An array of EntryPoint addresses. The first element should be the client's preferred EntryPoint.

**1.6. eth_chainId:**

* **Purpose:** Returns the EIP-155 Chain ID.
* **Parameters:** None.
* **Return Value:** The Chain ID as a hexadecimal string (e.g., "0x1" for Ethereum mainnet).

### 2. debug Namespace Methods:

These methods are intended for testing and debugging purposes and should be disabled in production environments.

**2.1. debug_bundler_clearState:**

* **Purpose:** Clears the bundler's mempool and resets reputation data for paymasters, accounts, factories, and aggregators.
* **Parameters:** None.
* **Return Value:** "ok".

**2.2. debug_bundler_dumpMempool:**

* **Purpose:** Retrieves a dump of the current UserOperation mempool for the specified EntryPoint.
* **Parameters:**
* `EntryPoint` (Address): The EntryPoint address.
* **Return Value:** An array of UserOperation objects currently in the mempool.

**2.3. debug_bundler_sendBundleNow:**

* **Purpose:** Forces the bundler to immediately build and execute a bundle from the mempool as a `handleOps` transaction.
* **Parameters:** None.
* **Return Value:** The transaction hash of the submitted bundle.

**2.4. debug_bundler_setBundlingMode:**

* **Purpose:** Switches the bundling mode between "manual" and "auto". In "manual" mode, an explicit call to `debug_bundler_sendBundleNow` is required to send a bundle.
* **Parameters:**
* `mode` (String): Either "manual" or "auto".
* **Return Value:** "ok".

**2.5. debug_bundler_setReputation:**

* **Purpose:** Sets or modifies the reputation data for specific addresses.
* **Parameters:**
* `reputationEntries` (Array): An array of reputation entry objects, each containing:
* `address` (Address): The target address.
* `opsSeen` (Quantity): Number of UserOperations seen and added to the mempool involving this address.
* `opsIncluded` (Quantity): Number of UserOperations included on-chain involving this address.
* `EntryPoint` (Address): The EntryPoint address.
* **Return Value:** "ok".

**2.6. debug_bundler_dumpReputation:**

* **Purpose:** Retrieves the reputation data for all observed addresses associated with a specific EntryPoint.
* **Parameters:**
* `EntryPoint` (Address): The EntryPoint address.
* **Return Value:** An array of reputation objects, each containing:
* `address` (Address): The address.
* `opsSeen` (Quantity): Number of UserOperations seen.
* `opsIncluded` (Quantity): Number of UserOperations included on-chain.
* `status` (String): The address's status: "ok", "throttled", or "banned".

**2.7. debug_bundler_addUserOps:**

* **Purpose:** Directly injects UserOperations into the mempool without performing validation, assuming they are valid.
* **Parameters:**
* `userOps` (Array): An array of UserOperation objects.
* **Return Value:** "ok".

These RPC methods provide a comprehensive interface for interacting with the ERC-4337 Account Abstraction infrastructure, enabling developers and users to manage UserOperations, monitor the system's state, and debug their implementations effectively.
