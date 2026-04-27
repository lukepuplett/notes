# All About Derivatives - Notes from Michael Durbin

## Introduction & Fundamentals

### Quotes & Pricing Basics
- A **quote** is a price at which one party is willing to buy or sell
- **Bid**: price I will buy at
- **Ask** (or **Offer**): price I will sell at
- Quotes are usually accompanied by a **size** (how many contracts are available at that price)
- Three basic components: **bid**, **value**, and **ask**
- Bid and ask sizes may differ

### Pricing Concepts
- Pricing a contract is mathematically fairly simple; the main task is adjusting values for time
- **Swap pricing** involves:
  - **Spot rate**
  - **Forward rate**
  - **Yield curve**
- The hard part in pricing is not the math, but rather:
  - Which **interest rate** to use
  - Which price to plug in for the **underlier**

### Option Pricing Intuition
- To price an option, construct an imaginary portfolio of non-option-like things (whose prices are easy to obtain) such that the portfolio payoff mimics the option's payoff

### Law of One Price
- **Law of One Price**: Two things with identical payoffs must cost the same (to prevent arbitrage)
- *Note: This rule was encountered at Centauri where a pricing model broke because it violated this arbitrage principle*

### Mathematical Complexity
- The heavy lifting in derivatives math involves **Partial Differential Equations (PDEs)**
- If you're going to struggle with derivatives mathematics, it will be on PDEs

---

## Chapter 2: Forward Contract

### Forward Rate Agreement (FRA)
- **FRA** (forward rate agreement): a contract on future interest rates for a specified amount
- Pronounced like the beginning of "France"
- **Swaps** can be thought of as portfolios of FRAs

### Counterparty Risk Management
- You can **post collateral** with the counterparty to mitigate counterparty risk
- **Futures** contracts can be used to avoid counterparty credit default risk altogether

### Payoff & Payoff Diagrams

**Payoff** is the value of a contract at delivery.

**Payoff Diagram**: Illustrates the payoff across a whole range of spot prices at once. Essentially a function graph showing payoff as a continuous function of spot price.

Two payoff functions exist:
- One for the **long** (buyer)
- One for the **short** (seller)

#### Payoff Formulas
- **P<sub>long</sub> = S - K**
- **P<sub>short</sub> = K - S**

Where:
- **K** = delivery price (possibly related to "strike")
- **S** = spot price (the current price of the contract/underlier)

#### Payoff Diagram Characteristics
- Vertical axis: **P** (payoff)
- Horizontal axis: **S** (spot price)
- The diagram resembles a **T on its side**
- The line bisects the S-axis at point K
- As S increases, P increases linearly
- As S decreases, P goes negative
- The payoff line is diagonal, starting low on the P-axis, crossing through K on the S-axis, and heading out to the upper right

**Interpretation**: Whether a contract represents a "gain" depends on your intended action. If you immediately sell, you've gained. If you hold, you're buying at a different time—this is more about perspective and strategy than absolute profit.

---

## Chapter 3: Futures Contract

### Key Characteristics
- **Highly standardized** contracts
- Executed on an **exchange** (not OTC)
- Standard contract **sizes**
- **Anonymous counterparties** (though counterparties are known to the exchange, and the exchange acts as gatekeeper)
- All trades are anonymous to each other

### Settlement - The Critical Difference
- **OTC contracts**: Payoff realized at delivery → exposes you to **counterparty risk**
- **Futures contracts**: Payoff realized at the **end of every trading day** via **daily settlement**
- This daily settlement mechanism eliminates counterparty credit risk

### Short Seller's Payoff Diagram
- The short seller's payoff diagram mirrors the long buyer's
- Instead of a line pointing up-right, it points down-right (from upper left to lower right)

---

## Chapter 4: Swap Contract

### Basic Structure
- **Cash flows** stem from interest payments
- One side pays a **floating rate**
- The other side pays a **fixed rate**
- Parties are essentially **swapping income streams**

### Long vs. Short in Swaps
- Determining long/short in swaps is less intuitive than in other derivatives
- The **fixed-rate payer** may be considered the long side
- The **variable-rate payer** is betting that rates will go down

### Fixed vs. Variable Rate Perspective
- If you're **long on rates** (expecting them to rise), you don't want a higher floating rate
- Example: Mortgage holders prefer fixed rates when they expect rates to rise
- If you choose a **variable rate**, you're betting that it will go down
- *Note: Some confusion here about the relationship between rate expectations and long/short positioning—this could use further clarification*

### Key Swap Mechanics

**Notional Amount**
- The **notional** is the amount of the obligation—the agreed number that interest payments are based on
- Example: A home loan of $250,000
- Important: The notional has no direct relationship to the contract value; it's purely the reference number for calculating interest payments

**Multi-Currency Swaps**
- When the two legs are in different currencies, parties may wire each other the notional amounts
- This helps mitigate **exchange rate** risk from currency fluctuations

**Net Present Value & Pricing**
- No fee is exchanged in a swap
- The **fixed rate** is adjusted so its **net present value (NPV)** matches that of the variable rate
- The two sides **net off**—they offset each other

**Tenor & Coupon Frequency**
- **Tenor** (or **coupon frequency**) is the payment schedule of cash flows
- Common tenors:
  - 1 month
  - 3 months
  - 6 months
- Important: The two legs of a swap do **not** need to have the same tenor

### Master Agreements & Standardization
- **Swap counterparties** typically execute a **master agreement** ahead of time
- A master agreement simplifies future swap contracts
- Usually contains or refers to **ISDA definitions** (ISDA is a standards body that defines swap terms)
- Each **coupon** is really just a **forward contract** in disguise

### Non-Vanilla Interest Rate Derivatives

**Basis Swap**
- Has two **floating legs** (not one fixed and one floating)
- Each leg might be a different variable rate
- Example: LIBOR vs. an investment bank's standard variable rate

**Currency Swap**
- Two parties transfer each other the **principal** (in different currencies)
- At the end, they transfer the principal back
- Result: Both parties end up with what they originally had
- Note: If one currency has appreciated, it has more buying power, but the nominal amount remains the same

### Foreign Exchange & Interest Rate Derivatives
- **Foreign exchange derivatives** are totally separate from interest rate swaps
- They operate under different pricing and risk dynamics

---

## Chapter 5: Options

### Option Fundamentals

**The Two Parties**
- **Writer** (seller): Creates the option contract
- **Holder** (buyer): Owns the right to exercise

**Premium**
- The **writer** receives a **premium** based on the option's value upon **execution**
- **Execution** = the creation and instantiation of the option contract (when it's written)
- **Exercise** = when the holder actually uses their right (different from execution)
- Important: These terms can be confused because they're similar

**Rights & Obligations**
- The **buyer (holder)** has the **right** to exercise, but is **not obligated** to
- The buyer can choose to:
  - **Call**: Right to buy the underlying
  - **Put**: Right to sell the underlying
- The buyer decides whether or not to exercise based on whether it's profitable

### Premium vs. Forwards and Swaps
- The **premium** makes options radically different from **forwards** and **swaps**
- Forwards and swaps have **no value upon execution** (no premium is paid)
- With options, the premium is later factored into the **payoff diagrams**

### Strike Price & Exercise

**Strike Price (Exercise Price)**
- The **strike price** (or **exercise price**) is the agreed price at which the holder can buy or sell the underlying
- Set when the option is created

**Assignment**
- With **exchange-traded options**, when a holder exercises, the **exchange must find a writer** to fulfill the obligation
- This process of matching a holder with a writer is called **assignment**
- The assignment is made to the **writer**

### Expiry
- All options specify an **expiry date** (the last day the option can be exercised)

### American vs. European Options

**American Option**
- Allows exercise **on or before expiry**
- More flexible than European options

**European Option**
- Allows exercise **only on expiry**
- More restrictive but simpler to price

### Warrants

**Warrant**
- An option where the **writer and the asset issuer are the same party**
- The issuer writes the option on its own asset
- *Note: This may explain some past trading strategies—if a company owns the underlying assets it's hedging, it might be issuing warrants*

**Dilution Effect**
- When a company issues new stock upon warrant exercise, it can **dilute** existing stock value
- Creates a somewhat circular feedback loop: exercising the warrant increases share count, potentially reducing per-share value

### Price Paths & Random Walks

**Price Path**
- The **price path** is the value of the underlying asset over time
- Typically modeled as a **random walk** in derivatives mathematics
- Used to understand how option values change as the underlying moves

### Option Notation & Ticker Symbols

Options are often denoted like ticker symbols with components:
- **Prefix**: lowercase = European, uppercase = American
  - Example: `cZED60` = European call on ZED with strike 60
  - Example: `CZED62` = American call on ZED with strike 62
- **Ticker**: The underlying asset (e.g., `ZED`)
- **Strike Price**: The exercise price (e.g., `60`, `62`, `64`)

### In the Money & Out of the Money

**In the Money (ITM)**
- A call option is **in the money** when the spot price is above the strike price
- Example: If ZED trades at $62, then cZED60 is $2 in the money
- The intrinsic value is the difference between spot and strike

**Out of the Money (OTM)**
- A call option is **out of the money** when the spot price is below the strike price
- Has no intrinsic value (yet)

**Deep In the Money**
- Options that are deeply in the money are so likely to expire in the money that they start to resemble **forwards**
- The holder essentially owns the underlying with certainty

**Exercise Decisions**
- When an option goes in the money, the holder **can** exercise it
- However, holders may wait if they expect the option to go even deeper in the money
- This betting on further price movement is why options have time value

### American vs. European Value

- **American options are always more valuable than identical European options**
- The flexibility to exercise anytime (rather than only at expiry) has real value

### Option Payoff Diagrams

**Key Difference from Forwards/Swaps**
- Option payoff diagrams have **no negative payoff**
- There is **no obligation to lose money** with options (unlike forwards/swaps)
- The holder can simply choose not to exercise

**Holder's Payoff**
- The holder's payoff function never goes negative
- The line tracks along the S-axis (spot price), then heads up and to the right
- Reflects the holder's ability to walk away if unprofitable

**Writer's Payoff**
- The **writer's payoff diagram is the mirror** of the holder's
- The writer's side goes **only negative** (or breaks even at best)
- The line tracks along the S-axis, then heads down and to the right
- The writer takes the opposite side of the bet

**Visual Pattern**
- Both diagrams follow the same S-axis (spot price) and angle in opposite directions
- Holder's: up-right; Writer's: down-right (or vice versa depending on call vs. put)

### Put Option Payoff Diagrams

**Long Put (Buyer)**
- The payoff line descends from the **top left** at a 45-degree angle
- It meets the S-axis at the **strike price K** (where payoff = 0)
- Then it **tracks rightward along the S-axis** (stays at zero for higher spot prices)
- Shape: descending line until K, then flat at zero to the right
- The buyer profits when the price falls (left side), capped at the strike

**Short Put (Writer)**
- The **mirror image** of the long put
- The payoff line ascends from the **bottom left** at a 45-degree angle
- It meets the S-axis at the strike price K
- Then it **tracks rightward along the S-axis** (stays at zero for higher spot prices)
- Shape: ascending line from negative until K, then flat at zero to the right
- The writer profits when the price stays high (right side), loses when it falls

**Practical Example of a Short Put**
- Example: Writing a short put on a Porsche is like making this promise: "I will buy your Porsche for $10,000 in 2030"
- **If the car is worth $20,000 in 2030**: The holder won't exercise (they can sell for more). The writer keeps the premium.
- **If the car is worth $5,000 in 2030**: The holder WILL exercise (they get $10,000 from the writer vs. $5,000 market value). The writer must buy the car for $10,000 when it's only worth $5,000, losing $5,000.
- The writer's loss is capped at (strike price × quantity), but the holder's profit is unlimited upside and capped downside.

### Premium & Profit vs. Payoff

**Premium Defined**
- The **premium** is the upfront cost paid by the buyer to the writer for the right to enter the contract
- This is paid at the initiation of the contract, regardless of whether the option is exercised

**How Premium Affects Payoff Diagrams**
- **Payoff diagrams** show the value at exercise without considering the premium paid
- **Profit diagrams** account for the premium, which shifts the entire payoff diagram down
- For a long put buyer:
  - The profit line starts at **negative** (by the amount of the premium paid)
  - Remains negative while out of the money
  - **Crosses zero** when the intrinsic value equals the premium (break-even point)
  - Then becomes positive as the option goes deeper into the money
- The premium essentially creates a hurdle: the option must go into the money by at least the premium amount before the buyer becomes profitable

**Example with Porsche Put**
- If you buy a put with strike price K = $10,000 and pay a premium of $1,000
- Your profit/loss starts at -$1,000 (the premium you paid)
- You only start making a profit when the car's value drops below $9,000 (K - Premium)
- At $5,000 car value: You profit (10,000 - 5,000) - 1,000 = $4,000

### Payoff Diagram Visual Reference

For visual representations of payoff and profit diagrams:
- ![Option Payoff Diagrams](https://analystprep.com/blog/wp-content/uploads/2020/08/cfa-frm-options-payoffs.png) - Basic payoff diagrams
- ![Options Payoffs and Profits](https://analystprep.com/blog/wp-content/uploads/2020/08/cfa-frm-options-payoffs-2.png) - Payoff and profit diagrams (accounting for premium)

These diagrams illustrate how premiums shift payoff diagrams downward for buyers and show the break-even points where options become profitable.

### Interest Rate Options

**Cap**
- Acts as a **maximum** on interest rates
- Protects against rates rising above a certain level

**Floor**
- Acts as a **minimum** on interest rates
- Protects against rates falling below a certain level

**Collar**
- Combines a **cap** and a **floor**
- Provides both an upper and lower bound on rates

### Swaptions

**Swaption**
- The **right** (but not obligation) to enter into a swap in the future
- Gives flexibility to lock in swap terms at a predetermined rate

---

## Options Strategies

### Straddle

**Definition**
- A **straddle** is a strategy where you buy both a **call** and a **put** on the same underlying asset, with the same strike price and expiration date

**Use Case**
- Used when you expect a **strong move** in one direction or the other, but you're unsure which direction
- You profit if the asset moves significantly in either direction (up or down)
- You're betting on **volatility** and price movement, not direction

**Risk & Loss**
- If the asset **doesn't move**, you lose because you paid **two premiums** (one for the call, one for the put)
- Your maximum loss is **2 × Premium** if the asset stays exactly at the strike price
- Break-even points occur when the move is large enough to exceed both premiums paid

### Strangle

**Definition**
- A **strangle** is similar to a straddle, but the **call and put have different strike prices**
- Creates a "gap" between the two strike prices where the strategy loses money

**Use Case**
- Similar to a straddle: betting on a big move in either direction
- The different strikes make it cheaper to enter than a straddle (lower total premium)

**Drawback**
- You lose money for longer because the asset must move beyond both strike prices to become profitable
- The wider the gap between strikes, the larger the move required to profit

### Covered Call

**Definition**
- A strategy where you **hold the underlying asset** and **write (sell) a call option** on it
- Also called a **"buy-write"** because you must buy the underlying as you write/sell the contract

**Purpose**
- The underlying asset **protects the writer** should the asset skyrocket
- If the buyer exercises the call, you deliver the asset you already own
- Reduces risk for the call writer

**Mechanics**
- You own the asset and collect the premium from selling the call
- If the price stays below the strike, you keep both the asset and the premium
- If the price rises above the strike, your asset gets called away at the strike price

### Naked Options

**Definition**
- **Naked options** are options written without the writer holding the underlying asset
- Contrasts with covered calls, where the writer owns the underlying

**Reality**
- The book implies that **no one really writes naked options** in practice
- Too risky: the writer has unlimited liability if the market moves against them

---

## Chapter 6: Credit Derivatives

### Credit Default Swap (CDS)

**Overview**
- The **CDS is the vanilla product** of the credit derivatives world
- The simplest and most commonly traded credit derivative

**Basic Structure**
- An investor buys **$10 million in bonds** over 5 years (with coupon/interest payments)
- The **creditworthiness** of the bond issuer changes over time
- This impacts the **bond price**, which may fall well below $10 million

**CDS Protection & Premium**
- Investor enters a **credit default swap** with an insurance company or large investment bank
- Pays **1.5% of the notional** (e.g., $150,000 on $10 million) over the 5-year term
- If the **business goes bust**, the insurer/swap provider buys the bonds at **full face value** ($10 million)

**CDS Premium as Spread**
- The CDS premium is often called the **CDS spread**
- There is **no explicit spread per se**—the spread is **implicit** in the premium
- The terminology is **borrowed from bond markets**, where spread refers to the portion of yield attributable to credit risk

**Relationship to Bond Credit Spread**
- **Bond credit spread** = Corporate bond yield - Risk-free rate
- This spread represents the **risk premium** required to compensate for credit risk
- The bond credit spread **informs the CDS spread/premium**
- Both are used to assess the **probability of default**
- The CDS spread essentially prices the risk captured in the bond's credit spread

**Parties to the CDS**
- **Protection Buyer** - the investor who buys protection (pays the premium)
- **Protection Seller** - the insurance company/bank that sells protection (receives the premium)
- These are the two official parties to the CDS agreement

**Reference Entity**
- The **reference entity** is the borrower/issuer of the bond being protected
- **NOT a party** to the CDS contract
- Completely separate from the CDS transaction
- Typically **unaware** that a CDS exists on their bonds
- The CDS is a pure agreement between buyer and seller about what happens if this third party defaults

### CDS Types

**Single Name CDS**
- Covers the credit risk of **one reference entity**
- Payoff triggered if that entity experiences a credit event

**Portfolio CDS**
- Covers the credit risk of **multiple reference entities**
- Payoff may be triggered by:
  - **Any one** entity defaulting (first-to-default basket)
  - **More than one** entity defaulting (nth-to-default baskets)
  - Creates more complex risk profiles

**Other CDS Variants**
- **Basket Default Swap**: Triggered when multiple entities in a basket default
- **Binary CDS**: Has binary payoff structure (either full payout or nothing)
- **Forward CDS**: CDS that begins at a future date
- **CDS Option**: Option to enter into a CDS at a predetermined spread at a future date

### CDS Sellers

- Tend to be **large insurers** and **financial institutions**
- These entities have the capital and risk management to absorb credit risk
- Similar to how we saw swaps are traded by banks and market makers

### CDS as Put Option

**CDS Analogy to Put Options**
- A CDS premium is **like a put option premium**
- **CDS notional** ($10M) = **Put strike price**
- **Current bond market value** ($9M) = **Spot price**
- If a **credit event occurs**, the contract goes **in the money**
- **Payoff calculation** works like a **put option**
  - Max(Strike - Spot, 0) = Max($10M - Bond Value, 0)
- Protection buyer receives payout if bond value drops below the notional

### CDS as Cash Flow Model

**Reframing CDS**
- A CDS can be reframed as a **series of cash flows**:
  - **Protection payments** (the premiums paid by buyer to seller, regular intervals)
  - **Compensation payment** (the payout if a credit event occurs)
- These cash flows can be combined into a **single floating rate cash flow**

**Modeling Approach**
- CDS can be **remodeled and analyzed similarly to interest rate swaps**
- Just as swaps involve analyzing fixed vs. floating cash flows
- CDS involves analyzing premium payments vs. potential compensation payouts

**Understanding CDS as Insurance**
- Think of a CDS as a **tradable insurance policy**
- Unlike traditional insurance, CDS is tradable on the secondary market
- This allows investors to buy/sell credit protection dynamically

**Settlement Methods**
- CDS settlement can occur in two ways:
  - **Cash settlement**: Payout equals the price drop (e.g., bond drops from $10M to $8M → $2M payout)
  - **Physical settlement**: Insurance company (protection seller) buys the bond at face value ($10M)
  - Both methods have the **same economic effect** for the protection buyer

**Credit Events** (ISDA-Defined)
- Only **credit events** trigger a payout from the CDS
- **ISDA-defined credit events include**:
  - **Bankruptcy**: The reference entity declares bankruptcy
  - **Failure to Pay**: The entity fails to meet an obligation payment
  - **Obligation Acceleration**: Outstanding obligations are accelerated due to default
  - **Restructuring**: The terms of outstanding obligations are restructured unfavorably
  
- **Obligation Default**:
  - Essentially a **breach of the lending contract**
  - The borrower fails to meet some contractual obligation
  - Could include missing interest payments, violating covenants, or other contractual violations

- **Price declines from other causes do NOT trigger a payout**
  - Example: Bond price falls due to rising U.S. Treasury rates (interest rate risk)
  - This is an economic event, not a credit event
  - The CDS buyer is not protected against this scenario

### Other Credit Derivatives

**Total Return Swap (TRS)**
- Also called **Total Rate of Return Swap**
- A swap where two parties exchange cash flows based on the **total economic performance** of a bond or asset

**TRS Mechanics**
- **Protection seller** (e.g., insurance company) makes periodic payments to the buyer
  - Payments are typically **LIBOR + spread** (e.g., LIBOR + 30 basis points)
  - These are regular, scheduled payments
- **Protection buyer** transfers all income from the bond to the seller
  - Includes **coupon payments** (interest income)
  - Includes **capital gains and losses** (realized and unrealized)
- **Buyer continues to own the asset** for accounting purposes
  - Asset remains **off the balance sheet** (accounting treatment)
  - Even though the seller gets all economic returns

**TRS Compensation**
- Seller compensates buyer for value loss due to **any reason**, not just credit events
- Includes:
  - Credit losses
  - Market price changes
  - Any depreciation in asset value
- **Asset can be anything**, not just debt securities
  - Could be stocks, commodities, other derivatives, etc.

**Key Feature**
- **Bi-directional payments** are normally **lined up to be simultaneous**
- This creates a true **swap** (two-way exchange of cash flows)
- The buyer receives protection against downside while appearing to own the asset
- The seller receives all economic benefit in exchange for floating rate payments

### Credit Linked Note (CLN)

**Overview**
- A **vehicle for raising funds** based on credit risk
- Transfers credit risk from one party to investors
- Structure is similar to a **CDS** (premium paid for compensation)

**CLN Mechanics**
- Invested funds are **held in reserve** in case needed to compensate protection buyer
- In event of credit loss, reserves are used to pay compensation
- Structure: Investor buys the note, capital is held in reserve

**Investor Returns**
- Investor receives **premiums** (minus the seller's cut/fee)
- Investor also receives **interest on the capital** invested
- Return depends on whether credit event occurs

**CLN Issuance**
- Often issued by an **SPV (Special Purpose Vehicle)**
- **SPV is spun up clean**: a new entity created specifically for this purpose
- The SPV isolates the credit risk from the originating institution
- Allows the seller to raise capital while transferring credit risk to investors

### CDS Documentation

**Confirmation Letter**
- The agreement between parties can be negotiated following **ISDA guidelines**
- A **confirmation letter** documents the specific terms agreed upon
- Allows flexibility while following standardized frameworks

### CDS Pricing Challenges

**Pricing Complexity**
- **Forwards, swaps, and options** have well-known, standardized pricing models
- **CDS do NOT have universally accepted standard models** for pricing
- **Limited venues for price discovery**
  - OTC market is less transparent than exchange-traded derivatives
  - Pricing is often negotiated between parties

**Pricing Approach**
- CDS pricing uses an approach of **discounting the expected risk-neutral path**
- **Discounting**: Calculating the present value of future cash flows using a risk-free interest rate
- The expected payoff is discounted back to today's value

**Expected Value Calculation**
- For a vanilla CDS of $10 million on IBM:
  - Compute the **expected value** of the $10 million payment
  - **Expected Value = Payout × Probability of Default**
  - Example: 1% probability of default × $10M = $100,000 expected value
  - This expected value is then discounted to present value

**Recovery Rate**

- Upon bankruptcy, debt holders often recover **some of their money** through liquidation
- The **recovery rate** is the percentage of face value recovered in default
- CDS premium must **factor in the recovery rate**
  - If recovery rate is 40%, expected loss on $10M = $10M × (1 - 0.40) = $6M
  - The protection seller's expected payout is reduced by the recovery amount
- The **higher the expected recovery rate**, the **lower the CDS premium** required

**Computing Probability of Default from Bond Spreads**

- The **probability of default** can be estimated by looking at the **current risk premium** on the reference entity's bonds
- **Risk Premium = Corporate Bond Yield - Risk-Free Rate (Treasury Yield)**
- Example:
  - Corporate bond yields 5%
  - US Treasury yields 3%
  - **Risk premium = 5% - 3% = 2%**
- This 2% spread reflects the market's assessment of default probability and recovery rate
- Higher spreads indicate higher perceived default probability

### Credit Derivatives Market

**Trading Structure**
- Traded **OTC** (over-the-counter), like interest rate swaps and forex forwards
- Market makers publish **indicative prices** that are **non-binding**
- Allows flexible terms tailored to specific needs

---

## Chapter 7: Using Derivatives to Manage Risk

*This chapter covered practical applications and real-world scenarios for using derivatives to hedge and manage risk. The content was mostly illustrative examples and situations that were fairly self-evident, so detailed notes were not taken.*

---

## Chapter 8: Pricing Forwards and Futures

### Core Pricing Concept

**The Essence of Forward/Futures Pricing**
- The core principle is **adjusting the spot price for time**
- Unlike options (where the transaction is uncertain), forwards and futures involve a transaction that is **certain to occur**
- This certainty makes pricing simpler than options pricing

**Key Insight: Price Irrelevance**
- With forwards and futures, you can **disregard changes in the price of the underlier**
- Why? Because the transaction is certain - you will buy/sell at the agreed delivery price
- The future market price of the asset doesn't affect the forward/futures pricing
- Only the **current spot price** and **adjustments for time and costs** matter

### Adjustments to Spot Price

**Time and Cost Adjustments**
- When pricing a forward/futures contract, adjust the spot price for:
  - **Interest rates** (cost of capital/carrying cost)
  - **Storage costs** (warehousing, insurance, etc.)
  - Other **out-of-pocket expenses**
- These adjustments **increase** the contract price above the current spot price

**Why Contracts Cost More**
- The **contract price will generally be more expensive** than the current spot price
- You're paying for the time value of money and physical costs of carrying the asset

### Delivery Price & Contract Pricing

**Delivery Price (Contract Price)**
- The **delivery price** is the price the parties agree to buy/sell at
- Also called the **contract price**
- This is the central focus of Chapter 8: calculating this price

**Zero-Value Principle**
- For a **new forward/futures contract**, the delivery price is set such that:
  - The **value of the contract is zero** at initiation
  - Neither party gains or loses at the moment of contract creation
  - This makes the contract fair and balanced
  - This is because the forward contract is a **zero-sum game**

**Terminology Clarification**
- **Execution** = the instantiation/creation of the contract
- Do not confuse with **exercise** (which applies to options)

### Forward Price vs. Value

**Once the Deal is Done**
- Once the contract is created and the underlier's price begins changing
- The **value of the contract changes** (no longer zero)
- The delivery price remains fixed, but the contract's value fluctuates

**Forward Price**
- The **forward price** is the delivery price of a **theoretical new zero-value contract**
- It's the price at which a new contract would be created with zero value today
- As market conditions change, the forward price changes
- But the **original delivery price** of an existing contract stays fixed

**Cost of Carry**
- All costs associated with **waiting for the delivery date** are bundled into the term **"cost of carry"**

**Components of Cost of Carry**

- **Storage Costs** (for physical commodities)
  - Common for commodities like **oil**, **wheat**, **wine**, etc.
  - Includes physical warehousing, insurance, handling, etc.

- **Interest/Financing Costs**
  - The **interest component** compensates the **short party** (seller of the forward) for the **opportunity cost** of not having cash up front
  - When you sell a forward, you're foregoing the ability to sell the asset immediately and invest the proceeds
  - The interest component captures this lost opportunity

- **Dividends** (for stock forwards)
  - If the underlying asset is a stock that pays dividends, the short party loses those dividend payments
  - This is a cost of carry for equity forwards

- **Convenience Yield**
  - The **benefit of physical possession** due to a shortage
  - In quotes in the book because it's a specific type of benefit
  - Example: During a shortage, owning the physical commodity provides value beyond the normal market price
  - The ability to sell the physical asset during a shortage is valuable
  - Reduces the cost of carry (or even creates a benefit to owning)

**Overall Formula**
- **Forward Price = Spot Price + Cost of Carry**
- Cost of carry aggregates all these components into a single adjustment to the spot price

**Key Terminology**
- **Agreed Price (Delivery Price)** = the fixed price at which the buy/sell transaction occurs
- **Value** = the changing measure of how much better or worse off the parties are as market conditions change
- **Forward Price** = the delivery price of a theoretical new zero-value contract (changes as market changes)
- **Value of an Existing Forward** = Present Value of (Current Forward Price - Original Delivery Price)

### Time Value of Money

**Adjusting Money for Time**
- Money invested, lent, or held grows over time due to interest
- **Adjusting Forward (Future Value)**:
  - Calculate what an investment will be worth in the future
  - Involves **compounding** using an interest rate
  - Future value grows as time passes
- **Adjusting Backward (Present Value)**:
  - Calculate what a future cash flow is worth in today's dollars
  - This process is called **discounting**
  - **Present Value** = the value today of a future cash flow

**Timeline Visualization**
- Think of money value on a timeline that ascends over time
- Money grows as you move forward in time
- To find present value from a future value, you "slide backwards" on the timeline
- Going backwards on the growth chart means the value goes down
- This is how we "discount" future cash flows back to today

**Interest Rates**
- The book describes value growth due to the **"price of borrowing"** (i.e., interest)
- *Personal note: Interest is often viewed as compensation for inflation, not just the price of borrowing—different perspectives exist on what interest fundamentally represents*

### Simple Interest & Growth Calculations

**Simple Interest Formula**

FV = PV × (1 + rt)

Where:
- **FV** = Future Value (the amount at the end of the period)
- **PV** = Present Value (the initial amount)
- **r** = interest rate (per period)
- **t** = time period (in years, or the number of periods)

This formula calculates how much money will grow using **simple interest**, where interest is calculated only on the principal, not on accumulated interest.

### Compounding Frequency & Continuous Compounding

**The Effect of Compounding Frequency**
- As you compound interest more frequently, you earn slightly more interest each time
- **Common compounding frequencies**:
  - Annually (once per year)
  - Semi-annually (twice per year)
  - Monthly
  - Weekly
  - Daily
- The more frequently you compound, the more total interest you accumulate over the same period

**The Limit: Continuous Compounding**
- The ultimate refinement is to compound **infinitely frequently**—at every instant
- This is called **continuous compounding**
- Continuous compounding is the most accurate representation of how interest actually grows
- The formula for continuous compounding uses **Euler's number**

**Euler's Number (e)**
- **Euler's number** (denoted *e*) is a mathematical constant, similar to π (pi)
- Approximate value: **e ≈ 2.71828...**
- It's an **irrational constant** that appears frequently in mathematics, particularly in:
  - Growth and decay processes
  - Continuous compounding calculations
  - Exponential functions
- In derivatives pricing, Euler's number is fundamental to calculating present and future values under continuous compounding

**Continuous Compounding Formula**

FV = PV × *e*<sup>rt</sup>

Where:
- **FV** = Future Value (the amount at the end of the period)
- **PV** = Present Value (the initial amount)
- ***e*** = Euler's number (≈ 2.71828)
- **r** = interest rate (per period)
- **t** = time period (in years, or the number of periods)

**Example Calculation**

FV = 100 × *e*<sup>0.06×1</sup> = 100 × *e*<sup>0.06</sup> ≈ **106.18**

This means an initial investment of $100 at a 6% interest rate, compounded continuously for 1 year, grows to approximately $106.18.

**Key Difference from Simple Interest**
- **Simple Interest**: FV = PV × (1 + rt) → 100 × (1 + 0.06) = 106.00
- **Continuous Compounding**: FV = PV × *e*<sup>rt</sup> → 100 × *e*<sup>0.06</sup> ≈ 106.18
- Continuous compounding yields slightly more interest (0.18 more in this example) because interest is being compounded at every instant

**Discounting Under Continuous Compounding (Reversing the Formula)**

To reverse the continuous compounding formula and calculate **present value** from a **future value**:

PV = FV × *e*<sup>-rt</sup>

Where:
- **PV** = Present Value (what the future amount is worth today)
- **FV** = Future Value (the amount in the future)
- ***e*** = Euler's number (≈ 2.71828)
- **r** = interest rate (per period)
- **t** = time period (in years, or the number of periods)
- **The negative exponent (-rt)** reverses the compounding effect

**Example Calculation (Reversing the Previous Example)**

PV = 106.18 × *e*<sup>-0.06×1</sup> = 106.18 × *e*<sup>-0.06</sup> = 106.18 × 0.9418 ≈ **100**

This undoes the previous continuous compounding calculation: $106.18 in the future is equivalent to $100 today (at a 6% continuous discount rate for 1 year).

**Alternative Form of the Discounting Formula**

An equivalent way to express the discounting formula uses division instead of a negative exponent:

PV = FV / *e*<sup>rt</sup>

This is mathematically identical to PV = FV × *e*<sup>-rt</sup>

**Example Using Division Format**

PV = 106.18 / *e*<sup>0.06</sup> = 106.18 / 1.0618 ≈ **100**

Both forms (multiplication by *e*<sup>-rt</sup> and division by *e*<sup>rt</sup>) produce the same result and are used interchangeably.

### Forward Payoff and the Delivery Price

**Long Position Payoff**

FV<sub>long</sub> = (S − d) + C

Where:
- **FV<sub>long</sub>** = Future Value for the long party (buyer)
- **S** = Spot price (current market price of the underlying)
- **d** = Delivery price (the agreed price in the forward contract)
- **C** = Cost of carry (storage, interest, dividends, convenience yield, etc.)

**Short Position Payoff**

FV<sub>short</sub> = (d − S) + C

Where:
- **FV<sub>short</sub>** = Future Value for the short party (seller)
- **d** = Delivery price
- **S** = Spot price
- **C** = Cost of carry

*Note: The book has not yet discussed adjusting these for time.*

**Finding the Delivery Price**

When setting up a new forward contract, the **delivery price (d)** is the unknown variable we need to compute. Rearranging the long position formula:

FV<sub>long</sub> = (S − d) + C

**Solving for d (the mysterious variable)**

Here are the algebraic steps to isolate d:

**Step 1:** Start with the formula
FV<sub>long</sub> = (S − d) + C

**Step 2:** Subtract C from both sides
FV<sub>long</sub> − C = S − d

**Step 3:** Subtract S from both sides
FV<sub>long</sub> − C − S = −d

**Step 4:** Multiply both sides by −1 (to eliminate the negative on d)
−(FV<sub>long</sub> − C − S) = d

Which simplifies to:
−FV<sub>long</sub> + C + S = d

**Step 5:** Rearrange to standard form
d = S + C − FV<sub>long</sub>

**Final Formula for Delivery Price**

d = S + C − FV<sub>long</sub>

Or more simply, when FV<sub>long</sub> = 0 (as it is for a new forward contract with zero initial value):

**d = S + C**

This shows that the **delivery price equals the spot price plus the cost of carry**.

**The Zero-Value Principle**

For a **new forward contract**, the value must be zero at initiation. This is expressed as:

0 = S + C − d

Since:
- The **spot price (S)** is not in your control (it's the current market price)
- The **cost of carry (C)** is not in your control (it's determined by interest rates, storage costs, etc.)
- You can **only vary the delivery price (d)** in the contract negotiation

The delivery price must be set such that whatever S + C equals, d must equal the same amount to make the equation equal zero.

Therefore:
**d = S + C**

**Numerical Example**

Let's say:
- Spot price (S) = 123.45
- Cost of carry (C) = 123.45

Then the equation becomes:
0 = 123.45 + 123.45 − d

Solving for d:
0 = 246.9 − d

Therefore:
**d = 246.9**

This delivery price of 246.9 ensures the new forward contract has zero value at initiation, making it fair for both parties.

**Key Principle: The delivery price equals the spot price plus the cost of carry (d = S + C).** This fundamental relationship ensures that a new forward contract has zero value at initiation, making it fair for both parties.

### Breaking Down Cost of Carry: Costs vs. Benefits

The cost of carry (C) can be decomposed into two components from the **short party's perspective**:

- **c** (lowercase) = **costs** to the short party (storage, interest, insurance, etc.)
- **b** (lowercase) = **benefits** to the short party (convenience yield, dividends, etc.)

**Cost of Carry Formula with Components**

C = c − b

Therefore, the delivery price formula becomes:

**d = S + c − b**

Where:
- **S** = Spot price
- **c** = Costs to the short party
- **b** = Benefits to the short party

This shows that costs **increase** the delivery price (the short party must be compensated), while benefits **decrease** the delivery price (the short party receives compensation from the benefit).

---

## Example: One-Year Forward on Oil

Consider a one-year forward contract on oil with a spot price of $100 at **execution** (the creation and instantiation of the contract):

**The Long Party's Perspective**
- The long party (buyer) **gets a fixed price** on the oil
- They can take the $100 they would have spent and **invest it at 6% annual interest** (using US Treasury rates, for example)
- Over one year, they can earn **$6 in interest** on that money
- This $6 is a **benefit** that compensates them for locking in a fixed price

**The Short Party's Perspective**
- The short party (seller) **must wait to receive payment**
- They lose the opportunity to invest the $100 immediately and earn that $6 in interest
- **Interest cost to short**: $6 (opportunity cost of delayed payment)
- The short party must also **physically store the barrel of oil**
- **Storage cost**: $7 per year

**Calculating the Delivery Price**

Using the formula: d = S + c − b

Where:
- **S** = Spot price = $100
- **c** = Costs to short = $6 (interest) + $7 (storage) = $13
- **b** = Benefits to short = $0 (oil provides no benefit like dividends)

Therefore:
**d = $100 + $13 = $113**

The delivery price of the forward contract must be **$113** to fairly compensate the short party for:
- The $6 opportunity cost of not having immediate cash to invest
- The $7 cost of storing the barrel of oil

**Additional Notes**

**Risk Neutrality in Derivatives Pricing**
- The book states that we can safely price derivatives **assuming all investors are risk neutral**
- **Risk neutral** means investors don't care whether prices go up or down—they're indifferent to market direction
- This is a key assumption in derivatives pricing and is explained in more detail later in the book

**LIBOR as the Risk-Free Rate**
- **LIBOR** (London Interbank Offered Rate) is commonly used as the **risk-free rate** for derivatives pricing
- It's the standard instead of US Treasury rates in many derivative contracts

**Forward on a Stock (Special Case)**
- Forward contracts on stocks have **no cost of carry**
- Stocks don't require storage or have significant carrying costs
- Therefore, the forward price is simply: **d = S**
- (Note: If the stock pays dividends, those would be subtracted as a benefit to the short party)

**Stock Forward Pricing with Interest Rate**

Even though stocks have no cost of carry, the forward price still accounts for the **time value of money** through the interest rate. The formula is:

F = S × *e*<sup>rt</sup>

Which is equivalent to:

F = S × (2.718...)^(rt)

**Example: Forward on a Stock**

Given:
- Spot price per share: **$15.50**
- Interest rate (annual): **2%** (0.02)
- Time period: **6 months** (0.5 years)
- Total position: 100 shares

**Step 1:** Calculate the total spot value
S = $15.50 × 100 = **$1,550**

**Step 2:** Apply the forward pricing formula

F = $1,550 × *e*<sup>0.02 × 0.5</sup>

F = $1,550 × *e*<sup>0.01</sup>

F = $1,550 × 1.01005

F = **$1,565.58**

**Step 3:** Calculate forward price per share
Forward price per share = $1,565.58 ÷ 100 = **$15.66**

**Interpretation**
The forward price of **$15.66 per share** is higher than the spot price of **$15.50** because of the **2% annual interest rate**. Even with no cost of carry (storage, convenience yield, etc.), the time value of money still increases the forward price. The short party must be compensated for the time value of delaying receipt of the $1,550 until the contract settles.

### Forward Pricing with Storage Costs

When the underlying asset has **storage costs** (such as physical commodities like oil), the forward pricing formula becomes:

F = (S + U) × *e*<sup>rt</sup>

Where:
- **F** = Forward price
- **S** = Spot price
- **U** = **Present value of all storage costs** (discounted back to today)
- ***e*** = Euler's number
- **r** = Interest rate
- **t** = Time period

**Key Insight**
- Storage costs are paid at **different times** (not all upfront)
- Each storage payment must be **discounted back to present value** using the continuous compounding formula
- **U** is the sum of all discounted storage payments
- Then the total (S + U) is compounded forward to the delivery date

**Example: Oil Forward with Quarterly Storage Costs**

Consider a forward contract on oil with:
- Quarterly storage cost: **$7,500** (paid each quarter)
- Interest rate: **3% annual** (0.03)
- Time period: **1 year**

The storage costs are paid **4 times over the year**:
- **Quarter 1**: Storage cost paid immediately (no discounting needed)
- **Quarters 2, 3, 4**: Storage costs paid in future quarters, requiring discounting

**Calculating Present Value of Storage Costs**

*Note: U (used in the forward pricing formula) is equivalent to PV<sub>storage</sub>. This is how it's written in the notes; notation conventions may vary across different sources.*

PV<sub>storage</sub> = 7,500 + 7,500 × *e*<sup>-0.03 × 0.25</sup> + 7,500 × *e*<sup>-0.03 × 0.5</sup> + 7,500 × *e*<sup>-0.03 × 0.75</sup>

Breaking this down by quarter:

**Q1 (paid today):** 7,500

**Q2 (paid in 3 months = 0.25 years):** 7,500 × *e*<sup>-0.03 × 0.25</sup>

**Q3 (paid in 6 months = 0.5 years):** 7,500 × *e*<sup>-0.03 × 0.5</sup>

**Q4 (paid in 9 months = 0.75 years):** 7,500 × *e*<sup>-0.03 × 0.75</sup>

Once all quarterly payments are discounted and summed to get **U** (the present value of storage), the forward price is calculated as:

F = (S + U) × *e*<sup>rt</sup>

### Proportional Storage Costs

**What is Proportional Storage?**

**Proportional storage** is a type of storage cost where the **storage price changes proportionally with the value of the underlying asset**.

**Key Insight**
- Regular storage costs are typically **fixed dollar amounts** (like the $7,500 per quarter in the oil example above)
- Proportional storage costs are **percentage-based** (a percentage of the asset value)
- Because proportional storage depends on the asset value, it behaves similarly to **interest rates**, which also depend on the asset value
- Since both interest (r) and proportional storage (u) are proportional to the asset value, they can be **combined into a single rate** in the formula

**Simplified Formula with Proportional Storage**

F = S × *e*<sup>(r + u)t</sup>

Where:
- **F** = Forward price
- **S** = Spot price
- **r** = Interest rate (lowercase)
- **u** = Proportional storage rate (lowercase)
- **t** = Time period
- ***e*** = Euler's number

**Why the Simplification Works**
- Both r and u are rates that apply proportionally to the asset value
- They can be combined into a single combined rate (r + u)
- This eliminates the need to calculate and discount storage payments separately
- The formula is much simpler than the quarterly storage cost approach

**Personal Note: Spoilage as Storage Cost**

If a commodity is expected to spoil or deteriorate during storage, this spoilage can be modeled as a **storage cost** using the proportional storage approach. 

For example, **LNG (Liquefied Natural Gas) experiences boil-off**—natural evaporation of the cargo during storage and transport. This boil-off rate represents a real economic loss and can be incorporated as a proportional storage cost (u) in the forward pricing formula. This approach was used in options pricing models for LNG trading at Centauri (on their options pricing supercomputer), where the boil-off rate was a critical factor in pricing LNG forwards and derivatives.

---

*That's the end of my notes, thus far.*

## Notes & Observations

- The book makes heavy use of mathematical formulas and concepts
- Jargon is critical to understanding derivatives; bold formatting helps distinguish key terms
- The author doesn't provide extensive explanations of some concepts (e.g., detailed counterparty risk mitigation)
- Payoff diagrams are a powerful visual tool for understanding derivative behavior
- The line between conceptual understanding and mathematical rigor is important in derivatives
- Credit derivatives protect against credit risk specifically, not other market risks (e.g., interest rate risk)
