# C# / .NET rules

Merged from `Tz/src/.cursor/rules/dotnet-standards.mdc`, `always.mdc`, `Tz/AGENTS.md`, `Tz/src/AGENTS.md`, and `~/.cursor/rules/dependency-injection-seams-only.mdc`.

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
- Use `this.` when referencing class members.
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
- Do **not** use `#region` — refactor into smaller classes instead.

### Member order (class)

1. Fields and constants  
2. Constructor  
3. Properties  
4. Methods / functions  
5. Public members before private  
6. Group related members; separate groups with `//` on its own line  
7. Place `return` on its own line with a blank line above  

---

## Naming

| Kind | Convention |
|------|------------|
| Public types, methods, properties | PascalCase |
| Private fields, parameters | camelCase |
| Private fields | **No** leading underscore `_` |
| Interfaces | `I` prefix |
| Implementations | `Impl` suffix |
| DTOs | `Dto` suffix |
| Names | Meaningful, reflect purpose/intent |

When **logging or printing member names**, include the type for clarity (e.g. `ChatOrchestratorResult.Ok`, not `Ok`).

### Test method names (examples)

- `IAccount__HasSimilarName__when__luke_ashley_puplett_vs_puplett_l_a__then__true`
- `AgeClassifer__when__now_plus_2__then__returns_B`

---

## Namespaces and test layout

- Unlike typical .NET layouts, **each namespace segment gets its own subfolder** (including tests).
- **Bad:** extra single-name folder under a feature (`.../Proofs/AttestedMerkleProofBuilderTests.cs` nested wrong).
- **Good:** namespace path matches folder path (`.../Attestations/Proofs/AttestedMerkleProofBuilderTests.cs`).
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
