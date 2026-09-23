# Agent behaviour and tooling

Rules for **AI assistants** (Cursor, Claude Code) and cross-cutting workflow — not language syntax.

---

## Global Claude Code (`~/.claude/CLAUDE.md`)

- **Never use plan mode** (`EnterPlanMode`) — proceed with implementation; if stuck, ask the user.
- **Never `git commit`** unless the user explicitly asked to commit (non-negotiable).
- Approach: ask when unclear; prior art first; watch for transcription errors.
- Code: SRP, small classes, meaningful names, smaller functions.
- Errors: custom types; `Unable to {action}. {detail}.`; validate early.
- Tests: focused; assert messages; edge cases; targeted runs + `/tmp` + `rg`.
- Security: no secrets in code; validate inputs; least privilege.
- Misc: do not create temp files just to pretty-print; if Edit fails, re-read and retry Edit (no `cat` workaround).

---

## Zipwire repo agents (`Tz/AGENTS.md` / `CLAUDE.md`)

Same critical blocks as global where duplicated:

- Filtered `dotnet test` only; no `[Ignore]` without permission.
- **No subagents for trivial lookups** — grep/read yourself unless truly parallel work or user asked.
- User **questions are not instructions** (e.g. “Did you write that to tmp?” → answer, don’t do unless asked).
- Don’t let side notes derail the task.

### Cursor / long tasks (`src/AGENTS.md`)

- User says **Wait/Stop** → stop and clarify.
- Long multi-step work: use task tracking; keep status updated.

### Git commits (`src/AGENTS.md`)

- Do not commit automatically; ask “Ready to commit?” when appropriate.

---

## Repo-specific vs global conflicts

| Topic | Global Claude | Some repos (e.g. ProofPack, evoq-ethereum-eas) |
|-------|---------------|-----------------------------------------------|
| Plan mode | Forbidden | Plan written to file; user review before implement |

**Rule:** When working in a repo, follow **that repo’s** `CLAUDE.md` / `AGENTS.md` over global defaults when they conflict.

---

## User-global Cursor rule

**File:** `~/.cursor/rules/dependency-injection-seams-only.mdc` (always apply)

Same DI seams content as [csharp-dotnet.md](./csharp-dotnet.md#dependency-injection-seams-only) — register only external/config seams; `new` internal helpers; gate question before `Add*`.

---

## Security and ops (agents)

- No hardcoded secrets (`always.mdc`, global Claude).
- gcloud: always `--project`; no `danger-live` without explicit approval.
- Staging deploys on tag push — do not manually trigger staging builds (see `Tz/src/AGENTS.md` deployment section).

---

## Zipwire product context (for correct changes)

From `always.mdc`:

- **Zipwire** products: **Approve** (timesheets/journal), **Collect** (documents/ID), **Attest** (standalone ID + blockchain proofs).

---

## Documentation principles (`src/AGENTS.md`)

- DRY; docs near code; update with behaviour changes; prefer concrete examples.
