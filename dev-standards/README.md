# Dev standards compendium

Extracted **rules and guidelines** (not just links) from Cursor rules, AGENTS/CLAUDE files, and architecture docs on this machine. Sources are cited in each section; when sources disagree, both are noted.

**Last full pass:** 2026-09-23

## Documents

| Doc | Contents |
|-----|----------|
| [csharp-dotnet.md](./csharp-dotnet.md) | C# / .NET style, DI, OpContext, params records, logging |
| [architecture-zipwire.md](./architecture-zipwire.md) | Guts/Public, CQRS, events, sync vs async, services |
| [data-access-and-storage.md](./data-access-and-storage.md) | Firestore/GCS, repos, validators, OpContext enforcement |
| [typescript-and-frontend.md](./typescript-and-frontend.md) | Zipwire MPA (TS/Knockout/Bootstrap) + zwinv (Svelte/Bun) |
| [testing.md](./testing.md) | .NET, zwcli, zwinv test rules and pitfalls |
| [go.md](./go.md) | zwcli, Sentimentals (cards) |
| [agent-and-tooling.md](./agent-and-tooling.md) | Claude Code global, Cursor, git, gcloud, agent behaviour |

## Source inventory (files read for this compendium)

### Zipwire (`~/Git/Tz`)

| File | Role |
|------|------|
| `AGENTS.md` / `CLAUDE.md` | Critical test/git/subagent rules; high-level .NET/logging/git |
| `src/AGENTS.md` | Architecture, CQRS, DI, code style, OpContext, test account keys |
| `src/.cursor/rules/always.mdc` | Always apply: org, errors, tests, security, DI summary |
| `src/.cursor/rules/dotnet-standards.mdc` | C# globs `*.cs` |
| `src/.cursor/rules/data-access.mdc` | Storage/repos/OpContext |
| `src/.cursor/rules/guts-pattern.mdc` | Public + Guts libraries |
| `src/.cursor/rules/vertical-communication.mdc` | CQRS vs events (always apply) |
| `src/.cursor/rules/web-app-patterns.mdc` | MVC, workplace API auth |
| `src/.cursor/rules/bootstrap-components.mdc` | `_Bootstrap4*` partials |
| `src/.cursor/rules/typescript-usage.mdc` | TS/Knockout in Zipwire |
| `assets/CLAUDE.md` | Test verbosity / assert messages |
| `docs/architecture/SYNC_VS_EVENT_ORCHESTRATION.md` | When to sync vs event |

### Global tooling

| File | Role |
|------|------|
| `~/.claude/CLAUDE.md` | Global Claude Code rules |
| `~/.cursor/rules/dependency-injection-seams-only.mdc` | User-global Cursor (always apply) |

### Other repos on this machine

| File | Role |
|------|------|
| `~/Git/zwcli/AGENTS.md` | Go CLI: output tests, reuse, formatting |
| `~/Git/zwinv/.cursor/rules/*.mdc` | SvelteKit 5, TS, Bun, ask-mode |
| `~/Git/cards/AGENTS.md` | Go/Gin API design for agents |

Not yet inlined (link-only in repo-specific docs): `Hub/ProofPack/CLAUDE.md`, `Hub/2025-swaine-zoho-qb/.cursor/rules/general-deluge.mdc`, `zw-gitbook-docs/AGENTS.md`, full `tests/README.md` (see [testing.md](./testing.md) for extracted test rules).

## Maintenance

When you change rules in Tz or global Cursor/Claude files, update the matching section here or re-run a “refresh dev-standards compendium” pass with an agent.

**Duplication:** Zipwire repeats DI and C# rules in `always.mdc`, `dotnet-standards.mdc`, and `AGENTS.md`. This compendium merges them once in [csharp-dotnet.md](./csharp-dotnet.md).

**Conflicts:** Global `~/.claude/CLAUDE.md` forbids plan mode; some repos (e.g. ProofPack) require plan-and-review. **Use the repo you are in.**
