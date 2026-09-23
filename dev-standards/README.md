# Dev standards index

Personal map of where coding rules, architecture guidance, and agent workflows live across repos and tooling. **Source of truth stays in each linked file** — this page is an index, not a copy.

**Last reviewed:** 2026-09-23

## How to use

- When starting work in a repo, open that repo’s **AGENTS.md** / **CLAUDE.md** and **`.cursor/rules/`** first.
- Cursor loads rules from the workspace; global rules live under `~/.cursor/rules/` and `~/.claude/CLAUDE.md`.
- Zipwire app details that drift: prefer `Tz/src/AGENTS.md` over older snapshots (e.g. `me/zipwire-technicals.md`).

Discover sibling repos from Tz: [RELATED_REPOS.md](https://github.com/lukepuplett/Tz/blob/master/RELATED_REPOS.md) (beacon files under `../`).

---

## Master hubs

| What | Where |
|------|--------|
| Zipwire repo root (tests, git, DI, logging) | [Tz/AGENTS.md](https://github.com/lukepuplett/Tz/blob/master/AGENTS.md), [Tz/CLAUDE.md](https://github.com/lukepuplett/Tz/blob/master/CLAUDE.md) |
| Zipwire main dev guide | [Tz/src/AGENTS.md](https://github.com/lukepuplett/Tz/blob/master/src/AGENTS.md) |
| Always-on Cursor rules (Zipwire) | `Tz/src/.cursor/rules/always.mdc` |
| Global Cursor DI rule | `~/.cursor/rules/dependency-injection-seams-only.mdc` |
| Global Claude Code | `~/.claude/CLAUDE.md` |

Local clone paths (this machine): `~/Git/Tz`, `~/Git/Hub/notes`, `~/Git/zwcli`, etc.

---

## By language / stack

### C# / .NET (Zipwire)

| Topic | Location |
|-------|----------|
| Style, testing, DI, OpContext, params records | `Tz/src/.cursor/rules/dotnet-standards.mdc` |
| Firestore/GCS, validators, OpContext | `Tz/src/.cursor/rules/data-access.mdc` |
| Verticals, CQRS, events, MCP pitfalls | `Tz/src/AGENTS.md` |
| Testing (macOS SSL, sandbox, seed data) | [Tz/tests/README.md](https://github.com/lukepuplett/Tz/blob/master/tests/README.md) |
| ProofPack .NET | [ProofPack/CLAUDE.md](https://github.com/lukepuplett/ProofPack/blob/main/CLAUDE.md) |
| EAS library | `Hub/evoq-ethereum-eas/CLAUDE.md` (local) |

### TypeScript / JavaScript (Zipwire MPA)

| Topic | Location |
|-------|----------|
| TS + Knockout, minimal inline JS | `Tz/src/.cursor/rules/typescript-usage.mdc` |
| MPA, Gulp, workplace API auth | `Tz/src/.cursor/rules/web-app-patterns.mdc` |
| Web app commands | `Tz/src/Evoq.Timesheets.AspNetCoreMvc/AGENTS.md` |

### Razor / Bootstrap

| Topic | Location |
|-------|----------|
| `_Bootstrap4*` partials | `Tz/src/.cursor/rules/bootstrap-components.mdc` |

### Go

| Project | Location |
|---------|----------|
| Zipwire CLI | [zwcli/AGENTS.md](https://github.com/lukepuplett/zwcli/blob/master/AGENTS.md) |
| Sentimentals (STM) | `~/Git/cards/AGENTS.md` |

### TypeScript / SvelteKit / Bun (zwinv)

| Topic | Location |
|-------|----------|
| TypeScript classes / typing | `~/Git/zwinv/.cursor/rules/typescript-conventions.mdc` |
| Svelte 5 / SvelteKit | `~/Git/zwinv/.cursor/rules/sveltekit-conventions.mdc` |
| Bun | `~/Git/zwinv/.cursor/rules/bun-runtime.mdc` |

### Deluge (Zoho)

| Topic | Location |
|-------|----------|
| Deluge conventions | `~/Git/Hub/2025-swaine-zoho-qb/.cursor/rules/general-deluge.mdc` |

### Public docs (GitBook)

| Topic | Location |
|-------|----------|
| Doc workflow, SUMMARY.md | [zw-gitbook-docs/AGENTS.md](https://github.com/lukepuplett/zw-gitbook-docs) |

---

## Architectural patterns (Zipwire)

| Pattern | Location |
|---------|----------|
| Public + Guts libraries | `Tz/src/.cursor/rules/guts-pattern.mdc`, `web-app-patterns.mdc` |
| CQRS facades vs events | `Tz/src/.cursor/rules/vertical-communication.mdc`, `Tz/src/AGENTS.md` |
| When to sync vs choreograph | [Tz/docs/architecture/SYNC_VS_EVENT_ORCHESTRATION.md](https://github.com/lukepuplett/Tz/blob/master/docs/architecture/SYNC_VS_EVENT_ORCHESTRATION.md) |
| DI — register only at seams | `dotnet-standards.mdc`, global `dependency-injection-seams-only.mdc` |
| OpContext / data protection | [OP_CONTEXT_DATA_PROTECTION.md](https://github.com/lukepuplett/Tz/blob/master/docs/architecture/OP_CONTEXT_DATA_PROTECTION.md), `data-access.mdc` |
| Workplace admin API auth | [WORKPLACE_ADMIN_AUTH_PATTERNS.md](https://github.com/lukepuplett/Tz/blob/master/docs/WORKPLACE_ADMIN_AUTH_PATTERNS.md), `web-app-patterns.mdc` |
| MCP surfaces vs dispatch/DI | `Tz/docs/architecture/MCP_SURFACES_AND_PERSONAS.md`, `MCP_TOOL_DISPATCH_AND_DI.md` |
| Attest / ProofPack lifecycle | `MERKLE_TREE_ATTESTATION_PROOFPACK_LIFECYCLE.md` (under `Tz/docs/architecture/`) |
| API response / domain specs | `Tz/docs/specs/SPEC_*.md` |

High-level architecture snapshot (may age): [me/zipwire-technicals.md](../me/zipwire-technicals.md).

---

## Testing and quality

| Rule | Where |
|------|--------|
| Never `dotnet test` without `--filter` | `Tz/AGENTS.md` |
| Fakes not mocks; MockHttp for external HTTP | `dotnet-standards.mdc`, `Tz/AGENTS.md` |
| Build single projects, not whole solution | `Tz/AGENTS.md` |
| Assert messages; less noisy test output | `Tz/assets/CLAUDE.md` |
| zwcli doc-driven output tests | `zwcli/AGENTS.md` |

---

## Security and ops

| Topic | Where |
|-------|--------|
| Secrets, validation, least privilege | `Tz/src/.cursor/rules/always.mdc`, `~/.claude/CLAUDE.md` |
| gcloud project name; danger-live caution | `Tz/src/AGENTS.md` |
| Deploy / tags / staging | `Tz/src/AGENTS.md`, `Tz/docs/README.md` |

---

## Agent workflow (meta)

| Topic | Where |
|-------|--------|
| Commit only when asked | `Tz/AGENTS.md`, `~/.claude/CLAUDE.md` |
| No `[Ignore]` on tests without permission | `Tz/AGENTS.md` |
| Avoid subagents for trivial lookups | `Tz/AGENTS.md` |
| Questions are not implicit instructions | `Tz/AGENTS.md` |
| ProofPack: plan + review (repo-local) | `ProofPack/CLAUDE.md` vs global “no plan mode” in `~/.claude/CLAUDE.md` — **follow the repo you’re in** |
| Time tracking / zw | `Tz/docs/prompt_helpers/HOWTO_ZW_TIME_TRACKING.md` |

---

## Other research repos (process, not syntax)

| Repo | Role |
|------|------|
| [invoice-factoring-notes](https://github.com/lukepuplett/invoice-factoring-notes) | Research workflow, open questions (private) |
| `Hub/investment_analysis/.cursor/rules/` | Analysis doc layout |

---

## Maintenance

When adding a new codebase:

1. Add `AGENTS.md` or `CLAUDE.md` at repo root if agents will work there.
2. Add `.cursor/rules/*.mdc` for language- or area-specific rules.
3. Link both from this README under the right section.

Optional later: script under `notes/tools/` to scan `~/Git` for `AGENTS.md` and `.cursor/rules/*.mdc` and regenerate sections.
