# Go rules

---

## zwcli (Zipwire CLI)

**Source:** `~/Git/zwcli/AGENTS.md` (symlinked `CLAUDE.md`)

### Architecture

- Cobra commands; clean architecture with interfaces; `internal/api` (real + fake), `commands`, `models`, `output`, `utils`.
- LLMs: prefer `--format structured` for machine-readable output.

### Development discipline

- **Search-first** before new formatters or helpers (`grep -r` across `internal/commands/`).
- **“Shapes in holes”** — map data types to existing hierout helpers (duration, timestamps, bullets, errors).
- One helper at a time; verify tests after each change.
- Avoid duplicate functions; remove dead code carefully (`deadcode -test`).
- Ask mode: if user asks for code changes while Ask Mode is on, refuse until approval mode.

(Full command/API debugging notes remain in repo `AGENTS.md`.)

---

## Sentimentals / cards (`stm`)

**Source:** `~/Git/cards/AGENTS.md`

### Stack

- Go + **Gin**; **Zerolog** JSON logs (GCP Cloud Logging); Tailwind in Docker build.
- Module `stm`; imports `stm/internal/...`; env prefix **`STM_`**.

### API design (agent-facing)

- **Hypermedia for LLMs:** every response includes `state`, `session`, `data`, `controls`.
- `data.items` is **always an array**.
- Controls: **link** (GET) or **form** (POST/PUT + JSON Schema body).
- Agents should not need external docs — API self-describing.
- **Never hardcode URLs** in controls — use named routes: `routing.RouteNameMap[routes.RouteName…]`.
- Route constants in `internal/features/<feature>/api/routes/routes.go`.
- JSON body handlers: generic **`BoundHandler[T]`** for type-safe schema generation.

(Deploy/gcloud commands in repo — specify project `stm-danger-live`.)
