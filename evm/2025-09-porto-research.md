# Porto Research Notes - September 2025

## Overview

Porto is a next-generation account abstraction stack for Ethereum that provides secure, scalable, and user-friendly account management. It solves key Web3 problems around onboarding complexity, cross-chain fragmentation, unpredictable fees, and poor user experience.

## EIP-7702: The Foundation

**EIP-7702 is the key unlock** that makes Porto's architecture possible. It allows Externally Owned Accounts (EOAs) to delegate their execution to smart contract code while maintaining the EOA's address as `msg.sender`.

### How EIP-7702 Works

**Traditional EOAs**: Can only send transactions with `msg.sender = EOA address`, but cannot execute arbitrary code.

**EIP-7702 EOAs**: Can delegate execution to a smart contract implementation, enabling:
- **Contract-like behavior**: EOAs can now execute complex logic
- **Preserved identity**: `msg.sender` remains the EOA address in all calls
- **Onward calls**: The delegated code can make multiple calls, all with `msg.sender = EOA`
- **Contract-triggered execution**: Authorized contracts can call the EOA, which then executes delegated code

### The Key Unlock

**Contract-to-EOA Calls**: This is the revolutionary aspect:
1. **Contract calls EOA**: An authorized contract (like Porto's Orchestrator) calls the EOA
2. **EOA executes code**: The EOA's delegated implementation runs
3. **Onward calls**: The implementation makes multiple calls to other contracts
4. **Identity preserved**: All calls have `msg.sender = EOA address`

**Example Flow**:
```solidity
// Orchestrator contract calls the EOA
EOA.transfer(recipient, amount);  // msg.sender = EOA in this call

// Inside EOA's delegated implementation:
function transfer(address recipient, uint256 amount) {
    // This code runs with msg.sender = EOA
    token.transfer(recipient, amount);  // msg.sender = EOA
    emit Transfer(EOA, recipient, amount);  // msg.sender = EOA
    // Can make many more calls, all with EOA as sender
}
```

This enables Porto to:
- **Batch operations**: Execute multiple calls atomically with EOA identity
- **Gas abstraction**: Pay fees in ERC20 tokens while maintaining EOA identity
- **Permission systems**: Implement complex authorization logic
- **Cross-chain operations**: Execute complex multi-contract flows

## Architecture

### Core Components

1. **Porto Account**: Smart contract that holds user funds and manages permissions
2. **Orchestrator**: Privileged contract that facilitates trustless interactions between relay and account
3. **Relay Service**: Infrastructure that builds, simulates, and sends transactions
4. **Simulator**: Utility for accurate gas estimation

## Passkey System

### How It Works

Porto uses WebAuthn (Web Authentication API) for hardware-backed authentication while solving account discoverability and cross-chain compatibility challenges.

#### Key Types Supported
- **WebAuthnP256**: Enables passkey support using WebAuthn standard
- **P256**: Standard ECDSA key on secp256r1 curve (browser session keys)
- **Secp256k1**: Standard ECDSA key on secp256k1 curve (Ethereum EOA private keys)
- **External**: Custom signature verification contracts

#### Account Creation Process

**New User (Ephemeral Key Approach):**
1. Generate EOA with ephemeral private key
2. Sign EIP7702 authorization to upgrade EOA to smart account
3. **Discard ephemeral private key** (this is the key insight!)
4. Account now controlled by authorized keys (passkeys, session keys)

**The Ephemeral Key Mystery Solved:**
- The "ephemeral" key **IS** the EOA's private key during creation
- This private key is used **once** to sign the EIP7702 delegation
- After delegation is signed and submitted, the private key is **immediately discarded**
- The EOA becomes a smart account controlled by authorized keys (passkeys, session keys)
- **No future transactions use the ephemeral key** - all subsequent operations use authorized keys

**Existing User:**
- Can upgrade existing EOA to smart account
- Original private key is NOT discarded (user retains control)

#### Authentication Flow

1. **Passkey Signs Digest**: User authenticates with passkey, signs transaction digest
2. **Signature Serialization**: WebAuthn signature serialized with metadata (authenticatorData, clientDataJSON, challenge)
3. **Smart Contract Verification**: Account contract verifies WebAuthn signature, extracts public key
4. **Execution**: Account executes calls based on authorized key permissions

#### WebAuthn Passkey Integration Deep Dive

**WebAuthn Standard Implementation:**
- **Hardware Security**: Uses platform authenticators (Touch ID, Face ID, Windows Hello, hardware security keys)
- **Biometric Authentication**: Leverages device biometrics for user verification
- **Cross-Device Sync**: Same passkey can be used across multiple devices via cloud sync
- **Phishing Protection**: Origin binding prevents credential theft

**Passkey Creation Process:**
1. **Credential Generation**: Device generates P256 key pair (public/private)
2. **User Handle**: Porto uses account address as `userHandle` for account discovery
3. **Registration**: Public key stored in account contract, private key stays on device
4. **Cross-Device**: Same passkey can be registered on multiple devices

**Authentication Flow Details:**
```typescript
// WebAuthn signature structure
{
  signature: { r, s },           // ECDSA signature components
  raw: authenticatorData + clientDataJSON + signature,
  metadata: {
    authenticatorData,           // Contains public key info and flags
    clientDataJSON,             // Contains challenge, origin, type
    challenge                   // Random challenge for replay protection
  }
}
```

**Smart Contract Verification:**
- **Signature Extraction**: Contract extracts ECDSA signature from WebAuthn format
- **Public Key Recovery**: Recovers public key from signature and message digest
- **Key Hash Matching**: Compares recovered key hash with stored authorized keys
- **Permission Check**: Verifies key has required permissions for operation

**Account Discovery Mechanism:**
- **User Handle**: Account address used as `userHandle` in WebAuthn
- **Credential Discovery**: Browser/device can find existing credentials by userHandle
- **Seamless UX**: Users don't need to remember account addresses or seed phrases

## ERC20 Gas Payment System

### Core Mechanism

Porto allows users to pay gas fees using ERC20 tokens (USDC, USDT) instead of native tokens (ETH) through an intent-based payment system.

#### Payment Flow

1. **Intent Creation**: User specifies `paymentToken` and `paymentMaxAmount`
2. **Relay Analysis**: Relay estimates gas costs and sets `paymentAmount`
3. **Payment Execution**: 
   - Relay calls `account.pay()` to transfer ERC20 tokens
   - Account increments nonce (payment happens even if execution fails)
4. **Work Execution**: Relay calls `account.execute(calls)` with actual transaction calls

#### Intent Structure
```typescript
struct Intent {
    address eoa;                    // User's smart account
    bytes executionData;            // Encoded calls to execute
    uint256 nonce;                 // Anti-replay protection
    address paymentToken;          // ERC20 token for gas payment
    uint256 paymentMaxAmount;      // Max user willing to pay
    uint256 paymentAmount;         // Actual amount relay requests
    address paymentRecipient;      // Who receives payment
    bytes signature;               // User's authorization
    // ... other fields
}
```

#### Supported Payment Tokens
- **Base**: ETH, USDC, USDT
- **Optimism**: ETH, USDC, USDT
- **Arbitrum**: ETH, USDC, USDT
- **Ethereum**: ETH, USDC, USDT
- **Polygon**: POL, USDC, USDT
- **BNB Chain**: BNB, USDT
- **Celo**: CELO, USDC, USDT

## The Orchestrator

### Role and Function

The Orchestrator is a separate smart contract that acts as a trusted intermediary between the relay service and user accounts.

#### Key Responsibilities

1. **Atomic Execution**: Ensures payment and execution happen atomically
2. **Signature Verification**: Verifies intent signatures once, avoiding re-verification
3. **Nonce Management**: Handles nonce incrementing to prevent replay attacks
4. **Trustless Guarantees**: Ensures relay gets paid even if execution fails

#### Execution Flow

```solidity
// Relay calls orchestrator
orchestrator.execute(encodedIntent)

// Orchestrator processes:
1. account.pay(paymentAmount, keyHash, intentDigest, encodedIntent)
2. account.checkAndIncrementNonce(nonce)
3. account.execute(executionMode, executionData)
```

**Relay Implementation**: The relay service (`src/rpc/relay.rs`) processes intents through:
1. **Preparation** (`wallet_prepareCalls`) - Validate, simulate, and generate signed quote
2. **Execution** (`wallet_sendPreparedCalls`) - Verify signature and broadcast transaction
3. **Monitoring** (`wallet_getCallsStatus`) - Track status and provide updates

#### Special Privileges

The orchestrator has privileged access to account functions:
- Can call `pay()` function to collect ERC20 payments
- Can call `checkAndIncrementNonce()` for nonce management
- Can call `execute()` without signature verification (orchestrator already verified)

## Multi-Chain Support

### Cross-Chain Architecture

- **Unified Account**: Same account works across all supported chains
- **Cross-Chain Signatures**: Multichain nonce prefix allows same signature across chains
- **Relay Infrastructure**: Centralized relay service handles routing and execution

### Supported Networks

**Production Networks:**
- Base (Chain ID: 8453)
- Optimism (Chain ID: 10)
- Arbitrum (Chain ID: 42161)
- BNB Chain (Chain ID: 56)
- Polygon (Chain ID: 137)
- Ethereum (Chain ID: 1)
- Celo (Chain ID: 42220)

**Test Networks:**
- Base Sepolia (Chain ID: 84532)
- Optimism Sepolia (Chain ID: 11155420)
- Arbitrum Sepolia (Chain ID: 421614)
- Katana (Chain ID: 747474)

## Key Insights and Clarifications

### Intent Terminology

**Important Distinction**: Porto's "intents" are not true intents in the MEV/trading sense:

- **Porto's "Intents"**: Signed transaction bundles with payment details (deterministic execution)
- **MEV Intents**: High-level goals without execution details (solvers find optimal paths)

Porto's "intents" are really "signed transaction bundles with payment abstraction."

**Confirmed from Relay Source**: The relay implements `IntentV05` struct which contains:
- `executionData`: Encoded array of calls (ERC7579 batch execution encoding)
- `paymentToken`: ERC20 token for gas payment
- `paymentMaxAmount`: Maximum user willing to pay
- `paymentAmount`: Actual amount relay requests
- `signature`: User's EIP-712 signature
- Cross-chain fields for multichain operations

### Account Upgrade Process

- **New Users**: Ephemeral key is discarded after EIP7702 upgrade
- **Existing Users**: Original private key is NOT discarded (user retains control)
- **Smart Account**: After upgrade, account controlled by authorized keys (passkeys, session keys)

### The Ephemeral Key Architecture Explained

**Why Use Ephemeral Keys?**
- **Security**: Private key exists only briefly, reducing attack surface
- **Simplicity**: No need for users to manage seed phrases or private keys
- **UX**: Seamless onboarding without complex key management

**The Complete Flow:**
1. **User Initiates**: Wants to create Porto account
2. **Key Generation**: System generates random secp256k1 private key
3. **EOA Creation**: Creates Ethereum EOA with this private key
4. **EIP7702 Authorization**: Uses private key to sign delegation to smart account implementation
5. **Key Discard**: Private key is immediately deleted from memory
6. **Smart Account Active**: EOA now controlled by authorized keys (passkeys, session keys)

**Security Implications:**
- **No Key Recovery**: If ephemeral key is lost, account is lost (but this is intentional)
- **No Seed Phrase**: Users don't need to manage 12/24 word seed phrases
- **Hardware Security**: Passkeys provide hardware-backed security for ongoing operations
- **Permission-Based**: All future operations require authorized key signatures

**Comparison with Traditional Wallets:**
- **Traditional**: User manages private key/seed phrase permanently
- **Porto**: User manages passkeys (hardware-backed, biometric-protected)
- **Trade-off**: Slightly less user control, significantly better UX and security

### Gas Efficiency

- **>50% more gas efficient** than alternatives on payments
- **State-of-the-art transaction latency**
- **Predictable, transparent pricing** in familiar tokens

## Developer Experience

### Integration

- **Wagmi Compatibility**: Seamless integration with Wagmi and Viem libraries
- **EIP-6963 Support**: Auto-discovery in wallet connection libraries
- **Third-Party Library Support**: Works with Privy, ConnectKit, AppKit, Dynamic, RainbowKit, Thirdweb
- **TypeScript First**: Full TypeScript support with excellent type safety

### Modes

1. **Dialog Mode**: Hosted UI for applications (default)
2. **Relay Mode**: Direct integration for wallets and account managers
3. **Hybrid Mode**: Fallback between modes based on environment

## Security Model

### Key Management

- **Hierarchical Keys**: Admin keys (full control) and session keys (limited permissions)
- **Permission System**: Granular control over operations, spending limits, time restrictions
- **Nonce Management**: 2D nonce system prevents replay attacks

### Trust Model

- **Orchestrator Trust**: Immutable privileged entity for payment and execution coordination
- **Relay Trust**: Infrastructure provider that routes transactions and manages state
- **User Control**: Users maintain control through authorized keys and permissions

## Use Cases and Applications

### Built-in Examples

- **Authentication**: Passkey-based login systems
- **Payments**: NFT purchases, subscription payments
- **Permissions**: Creator tipping, limited spending
- **Sponsoring**: Gasless transactions for users
- **Theming**: Custom UI theming for applications

### Real-World Applications

- **DeFi**: Cross-chain trading, yield farming
- **NFT Marketplaces**: Seamless purchases with passkeys
- **Gaming**: In-game transactions and asset management
- **Social Apps**: Creator monetization, tipping systems
- **Enterprise**: Corporate treasury management, multi-sig operations

## Technical Implementation

### Contract Architecture

- **Account Contract**: Manages keys, permissions, and execution
- **Orchestrator Contract**: Handles intent processing and payment
- **Relay Service**: Off-chain infrastructure for transaction routing
- **Simulator Contract**: Gas estimation and validation

### Relay Service Architecture

**Confirmed from Source Code**: The relay is a Rust-based service with the following components:

- **RPC Server** (`src/rpc/relay.rs`): JSON-RPC API with `wallet_*` namespace
- **Transaction Service** (`src/transactions/service.rs`): Handles transaction lifecycle and queuing
- **Storage Layer** (`src/storage/`): PostgreSQL backend with SQLx for type-safe queries
- **Price Oracle** (`src/price/oracle.rs`): CoinGecko integration for token pricing
- **Cross-Chain Operations** (`src/interop/`): LayerZero integration for multichain intents
- **Signers** (`src/signers/`): P256, WebAuthn, and dynamic signer support

### Standards Compliance

- **EIP-7702**: Proxy delegation for smart account upgrades
- **EIP-5792**: Wallet capabilities for enhanced functionality
- **EIP-6963**: Provider discovery for wallet integration
- **ERC-7579**: Batch execution encoding
- **WebAuthn**: Standard for passkey authentication

## Conclusion

Porto represents a sophisticated approach to account abstraction that addresses fundamental Web3 UX challenges:

1. **Simplified Onboarding**: Passkey-based authentication eliminates seed phrases and browser extensions
2. **Gas Abstraction**: ERC20 token payments make transactions predictable and accessible
3. **Cross-Chain Compatibility**: Unified account experience across multiple networks
4. **Developer-Friendly**: Seamless integration with existing Web3 tooling
5. **Security**: Hardware-backed authentication with granular permission controls

The system's use of ephemeral keys for account creation, intent-based payment processing, and orchestrator-mediated execution creates a robust foundation for next-generation Web3 applications.

## Additional Findings from Relay Source Code

### Intent Structure Confirmed

The relay implements `IntentV05` with the following confirmed fields:
- `executionData`: ERC7579 batch execution encoding of calls
- `paymentToken`: ERC20 token address for gas payment
- `paymentMaxAmount`: Maximum amount user is willing to pay
- `paymentAmount`: Actual amount relay requests (≤ paymentMaxAmount)
- `signature`: EIP-712 signature from user's authorized key
- `encodedPreCalls`: Optional calls executed before main intent
- `encodedFundTransfers`: Cross-chain fund transfer data
- `settler`: Settlement contract for cross-chain operations
- `funder`: Fund sourcing contract for cross-chain liquidity

### Cross-Chain Implementation

**Bundle State Machine**: The relay manages complex cross-chain state transitions:
1. `Init` → `LiquidityLocked` → `SourceQueued` → `SourceConfirmed` → `DestinationQueued` → `DestinationConfirmed` → `SettlementsQueued` → `Done`

**LayerZero Integration**: Uses LayerZero for secure cross-chain messaging and settlement verification.

### Multichain Nonce System

**Confirmed**: The relay uses `MULTICHAIN_NONCE_PREFIX` (0xc1d0) to signal cross-chain intents:
- When nonce prefix is 0xc1d0, EIP-712 signature excludes chain ID
- Allows same signature to be valid across multiple chains
- Enables atomic cross-chain execution

---

*Research conducted through codebase analysis of the Porto project repository and Ithaca Relay source code. All technical details confirmed from available documentation and implementation code.*
