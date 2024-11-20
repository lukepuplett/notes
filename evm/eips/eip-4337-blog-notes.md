# ERC-4337 Account Abstraction - Alchemy Blog Notes

## Part 1

https://www.alchemy.com/blog/account-abstraction

The following notes seem incredibly hard to follow. I think all the blog posts on the topic are bad and that's reflected here. Read my notes on EIP 4337 instead.

### Smart Contract Wallet

- Each person wanting to, e.g. use multiple signatories to transfer an asset, must have their own Smart Contract Wallet on the blockchain.

- The SCW serves as a kind of proxy to which a EOA issues commands.

- Blog I'm reading explains this without using the term proxy but it does look much like a proxy where the transaction calls a method on the SCW passing details which look a lot like a further transaction/operation to execute, complete with sig and nonce.

- To call the SCW a dedicated EOA is created for this purpose and holds enough ETH to pay gas to make the call(s).

- The SCW is known as an account in the ERC-4337.

- The EOA that calls the method on the SCW can be thought of as the executor.

- This account might be called bundler or relayer.

- The SCW sends ETH to the executor to reimburse them for the actual gas used.

- To ensure the SCW is honorable, another audited contract called the entrypoint is created to a) check SCW has funds for a worst case execution b) make the call to the SCW method and track actual gas used c) send some of the SCW's ETH to the executor to pay gas.

- For (c) to work, before making the call, someone needs to prefund the entrypoint via a special method and another method to take any surplus funds out. These are the deposit and withdrawTo methods.

- The problem is that if the SCW is reimbursing gas, if just anyone calls it and validation correctly fails, then that costs gas which is reimbursed by the SCW to the executor. A malicious caller can burn up all the SCW's funds to stop it functioning.

- To fix this, the SCW has a validation method and an execution method.

- The entrypoint now calls validate, then reserves ETH to pay max gas (or reject if not enough) then calls the execute method and refund the executor EOA from reserved funds and return rest to wallet's pool.

- The SCW enforces that only the entrypoint can call its execution method.

- The validation method is restricted so its unsuitable for anything other than validation.

- Problem now is the executor still pays for the gas of validation compute cycles.

- So the validation method must not use certain op-codes and access certain storage so that it can be simulated in a way that is predictably similar to how it'll run for real.

- With split validate and execute ops, the entrypoint now requires the validate op to deposit funds and fail it if it doesn't cough up.

- Unused gas must be returned to the SCW.

- But when writing a contract, it's risky to send ETH to an unknown contract because doing so calls code which could fail, use unpredictable gas, or attempt a reentrancy attack on us.

- So the entrypoint can't directly send funds to the SCW so it holds it and allows the SCW to pull it out via withdrawTo.

- The means the entrypoint could fund future calls from the ETH spare change it holds.

- The executor needs a tipping system because it needs incentivizing to do its thankless work. The maxPriorityFeePerGas field communicates a priority bribe.

- And the executor, when it sends its transaction to call the entrypoint, can choose a lower bribe and pocket the difference.

- The entrypoint is a shared singleton contract since it contains nothing custom, and so it takes an address of the SCW to work on.

Not sure if these are the exact official interfaces.

```
contract SCW {
    function validateOp(UserOperation op, uint256 requiredPayment);
    function executeOp(UserOperation op);
}

contract EntryPoint {
    function handleOp(UserOperation[] ops);
    function deposit(address scw) payable;
    function withdrawTo(address destination);
}

- When the owner of the SCW wants to perform an action, they craft a user op and, off-chain, ask an executor to handle it for them.

- The executor simulates the SCW validate method against the op to check it.

- If cool, the executor sends a transaction to the entrypoint to call handleOp

- The entrypoint then handles validation and execution on-chain and refunds ETH to the executor out of the wallet's deposited funds.

- Bundling is a badly named gas-saving batching optimization which lets the executor run a batch of different people's work in a single transaction. The fixed-priced transaction fee is saved and there's savings on accessing storage slots a second time.

- All validations are done before executions with failed validations having been discarded.

- The bundle cannot do >1 ops for the same SCW because its validate method could do funky shit that makes any subsequent validation behave differently to that under simulation.

- The executor can earn Maximal Extractable Value (MEV) by arranging ops in its bundle in a way that's profitable.

- The executor is thus called a bundler.

- Bundlers can work cooperatively with others to store and share (broadcast) validated ops.

- A bundler can benefit by also being a block builder; if they can choose the block that their bundle is included in, they can reduce or eliminate ops that fail execution after passing simulation.

- Bundlers and block builders might merge into the same role.


## Part 2

https://www.alchemy.com/blog/account-abstraction-paymasters

- What if we want someone other than the SCW owner to fund the SCW?

- Why? For blockchain noobs acquiring ETH is an adoption hurdle, learning curve. OR a dapp might be willing to foot the bill. OR a sponsor might allow the SCW to pay in USDC etc. OR for privacy, mixers.

- A paymaster is a type of contract that sees if it wants to pay for a user's op and looks something like this:

```
contract Paymaster {
    function validatePaymasterOp(UserOperation op);
}
```

- When a SCW submits an op is indicates which paymaster (if any) will pay gas: a new field on UserOperation struct can do this, and a new bytes field to pass arbitrary helpful data to the paymaster like "pleeeeease pay".

- The entrypoint now calls SCW.validateOp() and if the paymaster address is set, call pm.validatePaymasterOp(), validate as normal and execute op and track gas, then transfer ETH to the executor/bundler to pay for that gas. If the op has a paymaster, ETH comes from them, else from the SCW as before.

- Paymasters deposit ETH via entrypoint.deposit() before it can pay for ops.

- The same problem as before: the bundler wants to avoid submitting ops that fail paymaster validation because the paymaster won't pay and the bundler would be on the hook.

- But the same restrictions and simulation technique can't be applied here. That's because a paymaster's storage is shared across all the ops in the bundle that use that paymaster. So the actions of one validatePaymasterOp could potentially cause validation to fail for many other ops in the bundle that use the same paymaster.

- A malicious paymaster could use this to DoS the system. To prevent this, a reputation system is introduced!

- The bundler must now keep track of how often each paymaster has failedd validation recently and penalize those that fail a lot by throttling or banning ops using that one.

- To prevent paymasters just creating many instances of itself (a Sybil attack) the paymaster must stake ETH.

```
contract EntryPoint {
    // ..
    function addStake() payable;
    function unclockStake();
    function withdrawStake(address payable destination);
}
```

- Stake cannot be removed until some time after calling entrypoint.unlockStake()

- But if the paymaster only ever accesses the wallet's associated storage and not the paymaster's, then it doesn't need to stake since the storage accessed by >1 ops in the bundle will not overlap.

- Reputation rules are here: https://github.com/ethereum/EIPs/blob/master/EIPS/eip-4337.md#specification-2

- Bundlers can implement their own reputation logic.

- Stakes are never slashed; the exist to ensure Cybil attacks are costly.

- The paymaster needs a new method paymaster.postOp() which is a callback which is used to tell it how much gas was used so it can do stuff like pull down USDC.

```
contract Paymaster {
    function validatePaymasterOp(UserOperation) returns (bytes context);
    function postOp(bool hasAlreadyReverted, bytes context, uint256 actualGasCost);
}
```

- Presumably the paymaster checked the user had enough USDC before it validated the op, but if the op gives away all its USDC during execution, the paymaster can't then extract payment at the end.

- So there's a way for the paymaster to cause the op to fail after it's done and extract the payment anyway because it already agreed to pay gas during validatePaymasterOp().

- The entrypoint first calls paymaster.postOp() as part of the SCW.executeOp() so any revert will undo everything.

- If so, the entrypoint calls paymaster.postOp() once more and because we've undone whatever occurred in the execution the paymaster can extract its dues, hence the hasAlreadyReverted parameter.

```
struct UserOperation {
    // ..
    address paymaster;
    bytes paymasterData;
}
```

- The paymasters deposit ETH into the entrypoint just as a SCW paying its own way would.

- The entrypoint.handleOps() method, in addition to calling scw.validateOp() will also call the op's paymaster.validatePaymasterOp() method, then execute the op and finally calls paymaster.postOp()

- Due to problems with simulating paymaster validation, paymasters must also stake ETH.

```
contract Paymaster {
    // ..
    function addStake() payable;
    function unlockStake();
    function withdrawStake(address payable destination);
}
```

- Most of the above makes up ERC-4337 Account Abstraction, but not quite.


## Part 3

https://www.alchemy.com/blog/account-abstraction-wallet-creation

- Everyone's SCW has to get on the blockchain in the first place. The EOA can't do that since it wouldn't be very abstractiony.

- The goal is for a noob wanting a SCW should be able to get one published either by paying their own gas (even though they don't have a wallet yet) or by finding a paymaster to pick up the tab, without ever creating an EOA.

- Another less obvious important goal: when creating a new EOA,  generate the private key locally and claim the account without sending a transaction! I can then tell people my address and receive assets.

- Our SCW should have this capability, too; tell people its address before its even deployed.

- To do this we will need to know what address our EOA or SCW will have when created so anything sent to the address(es) are just waiting there.

- This address is called a counterfactual address and requires the CREATE2 opcode which deploys a contract at an address that can be deterministically computed using the address of the contract calling CREATE2, a 32-byte salt and the init code of the contract being deployed.

- The init code (code as in bytecode) is a blob of EVM bytecode which is an encoded function which returns a blob that is the contract to deploy.

- UserOperation would need a field so the user can pass in the init code.

- Now the entrypoint.handleOps() must now use CREATE2 to deploy the contract if initCode is specified, then proceed as normal which means calling the just-created SCW.validateOp() and call paymaster.validatePaymasterOp() if specified.

- A user can now deploy arbitrary contracts at determinable addresses, sponsored by a paymaster or paid by the user if they deposit ETH at the address where the contract will end up.

- However, the paymaster can't look at the bytecode to know if it wants to pay for running it, and the user can't tell if the code is good, either. The bundler is also stuck because it can't trust the bytecode either.

- So instead of taking in bytecode the user can pick a factory contract which calls CREATE2.

- Different factories can make different types of SCW, like multi-sig or whatever.

```
contract Factory {
    function deployContract(bytes data) returns (address);
}
```

- Users can simulate the call and get the actual address so they can fund it before calling it for real.

```
struct UserOperation {
    // ..
    address factory;
    bytes factoryData;
}
```

- Factories are obviously shared singletons, audited, and the paymaster can choose to pay for deployments for certain factories.

- Bundlers restrict factories in similar ways to before so they can accurately simulate them.

- And factories must also stake ETH using entrypoint.addStake() and bundlers can throttle/ban based on recent bad simulations. They don't need to stake if they don't do access certain storage.

- This is the entirety of ERC-4337.

## Part 4

https://www.alchemy.com/blog/account-abstraction-aggregate-signatures

- Optional gas saving optimization.
