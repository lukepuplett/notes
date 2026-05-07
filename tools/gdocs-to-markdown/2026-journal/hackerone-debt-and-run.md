# HackerOne - Debt and Run

---



## Summary:

Debt Bypass via `refundBySeller`



The `refundBySeller` function does not settle outstanding seller debts before processing refunds. This allows sellers to extract funds that should first go toward repaying the arbiter, potentially causing permanent financial loss to the arbiter.



`refundBySeller` lacks a call to `_settleDebt(msg.sender)` at the beginning of the function, unlike `withdraw()` which correctly settles debts first.



Debts are created when the arbiter covers chargebacks/refunds on behalf of sellers who lack sufficient balance.



## Steps To Reproduce:



### Attack Scenario (1M USDC Example)



#### Step 1: Seller accumulates debt

1. Seller operates high-volume business, receives 5M USDC in payments

1. Seller withdraws most funds: `balances[seller]` = 100K USDC

1. Customer disputes/chargebacks occur totalling 1M USDC

1. Arbiter calls `refundByArbiter` to cover these chargebacks

1. Arbiter pays from their balance: `balances[arbiter]` -= 1M

1. Result: `debts[seller]` = 1M USDC, `balances[seller]` = 100K USDC



#### Step 2: Seller receives new payment

1. New customer pays 900K USDC for goods/services

1. Result: `balances[seller]` = 1M USDC, `debts[seller]` = 1M USDC



#### Step 3: Seller exploits the vulnerability

1. Seller (legitimately or collusively) decides to issue refund on the 900K payment

1. Seller calls `refundBySeller(newPaymentID)`

1. Code checks: 900K <= 1M ✓ passes

1. Code executes: `balances[seller]` = 1M - 900K = 100K

1. Customer receives 900K USDC refund

1. Result: `balances[seller]` = 100K, debts[seller] = 1M USDC (unchanged!)



#### Step 4: Seller tries to exit

1. Seller attempts `withdraw()` to extract remaining 100K

1. `_settleDebt(msg.sender)` executes automatically

1. 100K → arbiter, reducing debt to 900K

1. Withdrawal fails (insufficient funds)

1. Seller gets nothing, leaves platform

1. Final permanent loss: 900K USDC (the amount that should have been settled but was used for refund instead)



