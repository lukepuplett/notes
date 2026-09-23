# TypeScript, JavaScript, and frontend

## Zipwire (`Tz`) — ASP.NET MPA

**Sources:** `typescript-usage.mdc`, `web-app-patterns.mdc`, `bootstrap-components.mdc`, `src/AGENTS.md`.

### When to use TypeScript

- Complex UI only — viewmodels in **MVVM** with **KnockoutJS**.
- Type-safe structures; not for every script.

### When to use minimal JS

- Simple interactions: data attributes, Bootstrap built-ins, basic event handlers.
- Keep **inline JavaScript minimal**.

### Bootstrap UI

- Use **`_Bootstrap4*`** partial views with matching **model classes**.
- Components handle validation state, responsive layout, accessibility.
- Examples: `_Bootstrap4InputSet`, `_Bootstrap4ButtonAnchor`, `_Bootstrap4FormSubmitButton`, `_Bootstrap4TabListItemAnchor`, `_Bootstrap4Card`, `_Bootstrap4Alert`, `_Bootstrap4Run`, `_Bootstrap4HtmlTable`, `_Bootstrap4BusinessCard`.

### Build

- **Gulp 4:** TS (source maps), Sass (Bootstrap overrides), static copy, BrowserSync for dev.
- Source: `ts/`, `sass/` → `wwwroot/`.

---

## zwinv — SvelteKit 5 + Bun

**Sources:** `zwinv/.cursor/rules/*.mdc`

### Audience (ask-mode.mdc)

- Luke: strong backend background; new to SPAs, Svelte 5, and TypeScript at scale.
- Prefer **Svelte 5 / runes**, not v4 patterns.
- Comments: **why** and “magic”, not restating **what**.

### Svelte 5 / SvelteKit

- Consult https://svelte.dev/docs/svelte/llms.txt
- Runes in `.svelte.js|ts`: `$state`, `$derived`, `$effect`, `$props`, `host()`.
- Snippets instead of slots; `onclick` HTML-style events (no `|once`).
- `$state.snapshot`, `$inspect`.
- Routes: kebab-case; components in `src/lib/components/`; `+page.svelte`, `+layout.svelte`; API under `src/routes/api/`.
- Components: PascalCase files; TS modules camelCase.
- Tailwind via `cn()` from `$lib/utils`.
- Vitest; `@testing-library/svelte`; tests as `.spec.ts` beside source.

### TypeScript (zwinv)

- **Strict** tsconfig; interfaces for props and API types.
- Class-based layout: static utils, stateful classes, service classes.
- Arrow functions in classes for `this`.
- Avoid `any`; use `unknown` + guards.
- Prefer async/await.
- Naming: long descriptive names; `is`/`has`/`can` for booleans; verbNoun (`getUser`, `validateToken`).
- Imports: `$lib/` absolute; group external → SvelteKit → local; named exports for classes.
- Utils in `src/lib/utils/`; services in `src/lib/server/`; business logic not in HTTP handlers.
- JSDoc for **why**, security assumptions, non-obvious behaviour.

### Bun

- `bun install`; commit `bun.lock`.
- Dev: `bun run dev`.
- Keep standard `package.json` scripts for portability.

### Project layout

- Main app: `svk/`; reference demo: `demo_front/`; docs in `docs/`.
- Keep `AGENTS.md` concise; absolute paths in tool calls when possible.
