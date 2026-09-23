# Testing rules

---

## .NET / Zipwire — critical

**Sources:** `Tz/AGENTS.md`, `always.mdc`, `dotnet-standards.mdc`, `src/AGENTS.md`, `assets/CLAUDE.md`, `tests/README.md`.

### Never unfiltered test runs

- **NEVER** run `dotnet test` without `--filter`.
- Non–`AllInMemory` tests may hit real infrastructure and **erase staging data**.
- Ask the user before any unfiltered run.
- **Never run ALL tests** unless explicitly asked.

Examples:

```bash
dotnet test --filter "FullyQualifiedName~ClassName"
dotnet test --filter "TestCategory=AllInMemory"
```

### Never `[Ignore]` without permission

- Do not add `[Ignore]` to skip broken tests.
- Simplify test + TODO and let it fail until user approves ignore.

### Frameworks

- **Fakes, not mocks** — no Moq, NSubstitute, FakeItEasy, etc.
- **Exception:** `RichardSzalay.MockHttp` for external HTTP (OpenAI, FreeAgent, blockchain).
- Arrange–Act–Assert; one focused behaviour per test; assert messages on every assert.
- Document in the test **why** it exists and what it expects; do not change expectations to match wrong behaviour.
- Comment the goal of the test; test edge cases.
- TDD when possible: tests first, stub impl, then implement.
- Prefer basic tests first, then implementation (`dotnet-standards.mdc`).

### Test runs (mechanics)

- Before runs: `dotnet msbuild -t:Clean`; `dotnet test /nodeReuse:false`.
- Targeted runs only; pipe to `/tmp/*.txt` and analyze with `rg` (avoid flooding context).
- Logger: `--logger "console;verbosity=detailed" --verbosity quiet` for failures without build noise.
- Less verbosity in output when possible but still show pass/fail clearly (`assets/CLAUDE.md`).

### Test account keys (MCP / API tests)

Format **`PREFIX_SUFFIX`** — must split on `_` into exactly two parts:

- Valid: `00_accn_test_actor_inv_01`, `00_accn_phone_norm_01`
- Invalid: `test-actor-phone-norm` (no underscore → 401 on MCP endpoints)

Convention: `00_accn_<description>_<number>`.

### Test data

- Use unique keys (timestamp/GUID) for parallel-safe tests.
- Rebuild test project if failures look stale.

### Cursor terminal

- Prefer **Terminal.app** or IDE test explorer for .NET tests — Cursor sandbox breaks IPC/TLS (see `tests/README.md`).
- macOS ARM64 + .NET 9: SSL/CSSM issues with external API tests; Docker or filter `Category!=ExternalAPI`.

### In-memory Firestore docs

See [data-access-and-storage.md](./data-access-and-storage.md) — **`InMemoryDocRef` required**.

---

## zwcli (Go)

**Source:** `zwcli/AGENTS.md`

- **Documentation-driven tests:** expected output in `PLAN_HIEROUT.md`; `sample_extractor.go` + `output_comparator.go`.
- Whitespace/indent in samples must match exactly.
- **Search before new formatters** — reuse `formatDurationHuman`, `formatCreatedAtHuman`, hierout helpers.
- File split: `journal.go` logic, `*_hierout_helpers.go`, `hierout_shared_helpers.go`, `*_test.go`.
- Incremental output changes: one formatting fix at a time.
- Duration/date formats are strict (e.g. `Mon 02 Jan 2006`, ‧ symbols for human duration).

---

## zwinv

- Vitest; `.spec.ts` beside source; test errors and edge cases; mock crypto where needed.
