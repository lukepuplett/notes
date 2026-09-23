# C# / .NET rules

Merged from `Tz/src/.cursor/rules/dotnet-standards.mdc`, `always.mdc`, `Tz/AGENTS.md`, `Tz/src/AGENTS.md`, and `~/.cursor/rules/dependency-injection-seams-only.mdc`.

---

## Design goal: lower cognitive load

The point of these rules is to **reduce cognitive load** and **remove ambiguity** — fewer mental speed bumps when reading, navigating, and reasoning about code.

Everything below should be read through that lens: expressive names, explicit `this.`, predictable layout, folder structure that “screams” the domain, tests that state circumstance and expectation in the name. When a rule feels picky, ask whether it buys clarity at a glance.

---

## Preparation and approach

- Read relevant `*.md` documentation before coding.
- Ask if the way forward is unclear; plan and look for prior art in the codebase.
- Use CLI commands when helpful.
- **Very expressive names** for variables and functions is the most important coding rule (`Tz/AGENTS.md`).
- User may use transcription — if a command or request seems wrong, stop and ask rather than guess.

---

## General coding

- Do **not** use `= null!;` — mark nullable, or use a good default/empty pattern.
- Use **`this.`** when referencing **class members** (fields, properties, methods on `this` instance).
- Wrap **all** `if` conditions in `{ }`, including single-line bodies.
- Break long methods into smaller private methods; look for reuse among them.
- Favor many small, single-purpose classes (SRP); sealed when no inheritance needed.
- Enable `#nullable enable` in new classes.
- Use **record structs** (or records) for complex inputs/outputs.
- Avoid magic strings; use constants.
- Use blank lines to group related code.
- Keep methods small and focused.
- Use `async/await` consistently.
- Prefer immutable properties when possible.
- Null-check and validate inputs early.
- Do **not** use **`#region`** (IDE collapsible regions) — see [Regions are a smell](#regions-are-a-smell).

### `this.` on class members

**Rule:** Write `this.field`, `this.Method()`, not bare `field` / `Method()` when referring to the instance.

**Why:** In a method you always have locals, parameters, and class members in scope. Without `this.`, the eye must stop and classify each identifier: is that a parameter? a local? the field? Explicit `this.` marks “this belongs to the object” immediately — less ambiguity, lower cognitive load. Same idea when discussing or logging members: prefer names that disambiguate (e.g. `ChatOrchestratorResult.Ok` not bare `Ok`).

### No `_` prefix on private fields

**Rule:** Private fields use camelCase; **no** leading underscore.

**Why:** `_field` plus optional bare `field` is a second visual dialect and often redundant once you standardize on `this.field`. Underscore prefixes look crufty here and do not add safety if `this.` is the disambiguator. Pick one clear pattern and stick to it.

### Member order (class)

**Rule:**

1. Fields and constants  
2. Constructor  
3. Properties  
4. Methods / functions  
5. Public members before private  
6. Group related members; separate groups with `//` on its own line  
7. Place `return` on its own line with a blank line above  

**Why:** Tidiness with a purpose — you know **where** to look, you see related state and behaviour **together**, and you can reason about the class in one vertical scan without hunting. Fixed order beats “whatever the last editor did.”

### Regions are a smell

**Rule:** Do not use `#region` / collapsible regions to organize a class.

**Why:** Regions hide structure instead of fixing it. If a file is so large that folding blocks is the only way to cope, the class (or method) is doing too much — **refactor into smaller types** instead. `#region` is a sign to split, not to fold harder. (Same spirit as “no `#region` — refactor into small classes” in Tz rules.)

---

## Naming

| Kind | Convention |
|------|------------|
| Public types, methods, properties | PascalCase |
| Private fields, parameters | camelCase |
| Private fields | **No** leading underscore `_` (use `this.` instead) |
| Interfaces | `I` prefix |
| Implementations | `Impl` suffix |
| DTOs | `Dto` suffix |
| Names | Meaningful, reflect purpose/intent |

When **logging or printing member names**, include the type for clarity (e.g. `ChatOrchestratorResult.Ok`, not `Ok`).

### Test method names

**Rule:** Name tests so they state **subject**, **circumstances** (`when__…`), and **expected outcome** (`then__…`) — no need to open the test body to know what is being proved.

**Why:** Failed test output should read like a spec sentence. Cognitive load drops when the name alone explains the scenario and expectation; debugging starts from the name, not from spelunking Arrange–Act–Assert.

**Examples:**

- `IAccount__HasSimilarName__when__luke_ashley_puplett_vs_puplett_l_a__then__true`
- `AgeClassifer__when__now_plus_2__then__returns_B`

Add assert messages too so the runner states which expectation broke.

---

## Namespaces, folders, and “screaming architecture”

**Rule:** Unlike many .NET repos, **each namespace segment is its own folder** (including tests). Folder path mirrors namespace path — no extra gratuitous nesting.

**Why:** Layout **screams** what the system is: feature and boundary visible from the tree, not buried in a flat list of files. You navigate by domain, not by alphabet soup. Same mental model in tests as in production code.

- **Bad:** single-name subfolder that hides the namespace (`.../Attestations/` → `Proofs/` → file one level too deep).
- **Good:** `.../Attestations/Proofs/AttestedMerkleProofBuilderTests.cs` matches the namespace segments.
- Put tests in files aligned with the class/system under test.

---

## Documentation and comments

- XML documentation on **public** APIs: parameters, returns, exceptions.
- Document error conditions and non-obvious behaviour.
- **Comments explain why, not what** — no narrating obvious steps (`// increment counter`).
- Do comment: gnarly logic, domain quirks, provider rules, security gates, ordering constraints (“list before clear so idents survive API failure”).
- Use clear narrative English in `///` remarks or block comments where helpful.
- Keep comments accurate when behaviour changes.
- Straightforward logic: code is the documentation.

---

## Error handling

- Use **specific** exception types, not bare `Exception`.
- Custom exceptions for domain errors; support serialization where appropriate; use inner exceptions.
- Message shape: `Unable to {action}. {what went wrong}.` (include context).
- Consider `Option`, `StatusOption`, `StatusMessage` for operations that may “fail” without throwing.

---

## Solution conventions

- **`OpContext op` is always the first parameter** on commands/queries and similar entry points.
- Default for logged-in callers: **restricted** `op`. Call `op.Derestricted()` only at **entry points** (controllers, event handlers) **after** auth gates — **not** inside Commands/Queries implementations. See [architecture-zipwire.md](./architecture-zipwire.md) and [data-access-and-storage.md](./data-access-and-storage.md).

### Params DTO records

- When a method would take many parameters or the same bundle is passed repeatedly, use `public record …Params(...)`.
- `OpContext op` stays first; params record is typically second.
- Extend the record for new optional inputs instead of adding parameters.
- Do **not** introduce a params DTO for trivial two-argument private helpers.

---

## Dependency injection (seams only)

Register in DI **only** when:

1. **External seam** — database, blob storage, third-party HTTP, message bus, clock (if faked at container in tests).
2. **Config/environment swap** — in-memory vs Firestore repos, active vs inactive email gateway, etc.

**Register examples:** Firestore/GCS repos, OAuth/API gateways, `*Commands`/`*Queries` implementations, storage factories, env-specific providers.

**Do not register:**

- Webhook validators/publishers, handlers used from one controller.
- Orchestrators that only compose already-registered services.
- Single-use helpers — **`new` them** in the owning controller/handler (dependencies can still come from DI / `ControllerInitializationContext` / `op.App`).
- Interfaces with one implementation and a “future stub” — use concrete class until a second real impl exists.
- Types “for testability” when tests can use in-memory seams or `new` with constructor args.

**Pattern:** `new InternalHelper(seamFromDi, seamFromDi)` — not `services.AddTransient<InternalHelper>()`.

**Gate before every `services.Add*`:** Will config/environment pick a different implementation? Will integration tests replace this registration? If both are **no**, do not register.

**Examples (not registered):** `AccountInitializer` (controller uses `new` with queries/commands from DI); Xero webhook validator/publisher (`new` in controller with config + init context).

---

## Logging

- Log **success** paths for important operations, not only failures.
- Log **counts** for batch/enrichment (e.g. number of activities added).
- Do **not** prefix messages with class names when using `ILogger<T>` (type is already in context).
- Before LLM/critical paths, log what context is being enriched.
- Do not catch/swallow without logging what failed and why.

---

## Build and tooling

- **Build individual projects**, not the whole solution (~3 minutes wasted otherwise): `dotnet build [ProjectPath]`.
- **`sd`** for find/replace (Rust): `sd 'old' 'new' file`, `sd -p` preview, `sd -r src/` recursive.

---

## Git (commits)

- **Do not commit** unless the user explicitly asks.
- Short commit subject; brief bullets only when needed for later understanding.
- Tz has **two remotes**: `origin` (GCP mirror) and `github` (PRs). Fetch/pull **both** when checking for upstream changes.

---

## gcloud

- Every `gcloud` command must specify **project name**.
- Do **not** run against projects with `danger-live` in the name without **explicit user approval**.
