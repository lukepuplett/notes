# Zipwire architecture rules

From `Tz/src/AGENTS.md`, `guts-pattern.mdc`, `vertical-communication.mdc`, `web-app-patterns.mdc` (API/auth parts only), `SYNC_VS_EVENT_ORCHESTRATION.md`.

---

## Application shape

- **Monolithic** server application with a **vertical-sliced** backend (and a legacy server-rendered UI in `Tz` — stack not documented here).
- **Vertical slices** by feature area; each vertical = **two libraries**:

### Public library (e.g. `Evoq.Timesheets.Web.Payment`)

- Public types and API.
- Minimal dependencies (Data, other **public** libs).
- May be referenced by any library.

### Guts library (e.g. `Evoq.Timesheets.Web.Payment.Guts`)

- Internal implementation only.
- **Must not** reference other **Guts** libraries.
- **May** reference: cross-cutting (Data, Framework), own public lib, other **public** libs.
- **Must not** be referenced by other feature areas.

---

## Vertical communication (do not mix patterns)

Choose **one** per cross-vertical need:

### 1. Synchronous — CQRS facades

- **Commands** — state changes, writes (`*Commands`).
- **Queries** — reads (`*Queries`).
- Use when **immediate** knowledge of outcome is required (payment verification, UI that must show new state in the same response).
- Internal signatures: any types. **Cross-vertical** signatures: simple types only.

### 2. Asynchronous — events

- Domain events via in-process `EventHub` (no external bus).
- Each vertical has a **service** that subscribes in constructor, implements `OnEvent`, uses `EventContext`.
- Singleton lifetime; idempotent handlers where possible; consider ordering, races, USNs/timestamps, retries.
- Use for side effects and background work (e.g. payment completion notifications).

**Rule:** Never mix sync facade and event for the same concern — pick based on whether the caller must know the outcome **now**.

---

## Sync vs event orchestration (when to deviate from “events default”)

**Default:** event choreography for side effects.

**Do synchronously when:**

1. **Race risk** — delete then create in the same request; event delete might not finish before create finds stale row.  
   **Rule:** delete synchronously when creation immediately follows.

2. **UI response correctness** — response or next client repaint must reflect new/absent state (new account key, workflow, removed sender visible in assignment).  
   **Rule:** do synchronously whatever the UI needs present or gone on next fetch/repaint.

**Can stay eventual:** cleanup the caller does not need in the same response (e.g. rate plan cleanup on `UserDeleted` when create path does not query that store; team/OAuth cleanup if UI does not need confirmation).

**Reference:** `SenderInvitationOrchestrator` — class remarks in Workflow Guts.

---

## Workplace-scoped API controllers (`WebOpApiController`)

1. Authenticate with **restricted** `op` — `ValidateApiKeyAsync`.
2. Resolve admin workplace — `ResolveAdminWorkplaceAsync(op, teamQueries, workplaceKey)` (helper may derestrict for team/account discovery).
3. After success: `var dOp = op.Derestricted()` for workplace-scoped Commands/Queries.
4. **Do not** derestrict inside command implementations.

Pattern matches `AssignmentsApiController`, `ClientsApiController`, `WorkplaceTeamsApiController`. Details: `Tz/docs/WORKPLACE_ADMIN_AUTH_PATTERNS.md` §6.

---

## Data model completeness

If tempted to create `OpContext` mid-method, break encapsulation, or add a query for one missing metadata field downstream:

- Stop and ask: **is that data already on the domain model upstream?**
- Prefer adding metadata to the source DTO/model, not querying around layers.

---

## HTTP / UI correctness (stack-neutral)

- If the client must show new or removed state on the **next** load or repaint, perform the underlying writes **synchronously** in the request (see sync vs event section above) — applies regardless of frontend framework.
- Web composition root for the existing app: `Evoq.Timesheets.AspNetCoreMvc` (implementation detail; future UI may differ).
