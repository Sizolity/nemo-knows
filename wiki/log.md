---
title: Log
kind: log
---

# Log

Append-only record of every ingest, filed query answer, lint pass, and
schema change. Each entry begins with a heading in the canonical
format so it is greppable:

```
## [YYYY-MM-DD] <action> | <subject>
```

`<action>` is one of: `ingest`, `query-filed`, `lint`, `schema-change`,
`note`. The full format and conventions are defined in
[`AGENTS.md`](../AGENTS.md) §6.

To skim recent activity:

```sh
grep "^## \[" wiki/log.md | tail -10
```

---

## [2026-04-22] note | repository bootstrap

Initial scaffold created: `README.md`, `AGENTS.md`, empty `raw/`, and
`wiki/` with `index.md` plus this `log.md`. No sources ingested yet —
the next entry should be the first real `ingest`.
