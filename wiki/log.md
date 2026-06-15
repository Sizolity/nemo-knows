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

## [2026-05-16] query-filed | LLM wiki core concept

Filed a quick-reference explanation of the LLM wiki pattern and filled
missing concept pages referenced by the index.
Touched:
- wiki/topics/llm-wiki-core-concept.md (created)
- wiki/concepts/ingest.md (created)
- wiki/concepts/query.md (created)
- wiki/concepts/lint.md (created)
- wiki/concepts/wiki-as-compounding-artifact.md (created)
- wiki/concepts/persistent-wiki.md (updated)
- wiki/sources/llm-wiki.md (updated)
- wiki/topics/nemo-knows-mvp.md (updated)
- wiki/index.md (updated)
- wiki/log.md (updated)
Open: none.

## [2026-05-16] ingest | drafts/actual-use-llm-wiki
Touched:
- wiki/sources/llm-wiki.md
- skipped: wiki/concepts/llm-maintenance-pattern.md — manual reviewed content required
- skipped: wiki/topics/persistent-wiki-architecture.md — manual reviewed content required
Open: review skipped candidates before creating concept or topic pages.

## [2026-05-16] ingest | drafts/actual-use-llm-wiki
Touched:
- wiki/concepts/llm-maintenance-pattern.md
- wiki/index.md
- wiki/sources/llm-wiki.md
- wiki/topics/persistent-wiki-architecture.md
Open: review skipped candidates before creating concept or topic pages.

## [2026-05-16] lint | wiki confidence frontmatter
Touched:
- wiki/sources/llm-wiki.md (updated)
- wiki/log.md (updated)
Open: none.

## [2026-05-26] schema-change | runtime-regenerated wiki pages

The tracked wiki state was reduced to the required skeleton files
(`wiki/index.md`, `wiki/log.md`, and directory placeholders). Runtime pages
under `wiki/sources/`, `wiki/entities/`, `wiki/concepts/`, `wiki/topics/`, and
`wiki/assets/` are regenerated from deployment inputs and are no longer stored
in Git.
Touched:
- AGENTS.md (updated)
- .gitignore (updated)
- wiki/index.md (updated)
- wiki/log.md (updated)
- wiki/sources/.gitkeep (created)
- wiki/entities/.gitkeep (created)
- wiki/concepts/.gitkeep (created)
- wiki/topics/.gitkeep (created)
- wiki/assets/.gitkeep (created)
Open: generated wiki pages should be reviewed on the server before promotion to
tracked source material.


## [2026-06-15] schema-change | core concepts from llm-wiki.md

Reviewed Karpathy's `raw/llm-wiki.md` against the current schema. Three
core concepts were missing or under-emphasized and have been added:

1. **Human/LLM boundary** (§0): The human curates sources and directs;
   the LLM does all bookkeeping, cross-referencing, and maintenance.
   The LLM owns the wiki layer; the human owns source selection and
   interpretation.

2. **Schema as co-evolved document** (preamble): AGENTS.md is a living
   document that the human and LLM co-evolve as patterns emerge, not a
   fixed specification.

3. **Why the pattern works** (§0): Added the justification — LLMs don't
   get bored, don't forget cross-references, and can touch many files
   in one pass. The maintenance cost is near zero, which is why the
   wiki stays current when human-maintained wikis are abandoned.

Also strengthened the query-filing concept (§4): good answers are
valuable and should compound in the wiki, not disappear into chat
history.

Touched:
- AGENTS.md (updated: preamble, §0 rewritten, §4 strengthened)
- wiki/log.md (updated)

## [2026-06-15] note | current-stage boundary clarification

The human/LLM boundary described in the schema is aspirational for the
full product. At the current stage:

- **Human** provides source documents, placed under `wiki/sources/`.
- **LLM** processes those sources into entity / concept / topic pages
  within `wiki/`, following the ingest workflow.
- **Auto-maintenance** (`nemo -maintain-wiki`) and **chat-query filing**
  are future priorities, not yet implemented in the production path.

The immediate work is to organize and optimize the repository file
structure, separating development artifacts from the production wiki.

## [2026-06-15] note | repository structure cleanup

Organized the repository file structure to align with the updated mental
model (wiki/ is the system, everything else is development infrastructure):

- Removed empty directories: `scripts/`, `docs/deployment/`
- Cleaned `.gitignore`: removed dead entries for `backups/`, `scratch/`,
  `tmp/`, `experiments/`; added `.agents/` and `.codex/`
- Rewrote `README.md` to reflect the new model: wiki/ is the product,
  raw/drafts/evals are development infrastructure, sources go into
  wiki/sources/ not raw/

Touched:
- scripts/ (removed)
- docs/deployment/ (removed)
- .gitignore (updated)
- README.md (rewritten)
- wiki/log.md (updated)

## [2026-06-15] schema-change | consolidate dev infrastructure under pipeline/

Moved `raw/`, `drafts/`, and `evals/` under a single `pipeline/` directory
to consolidate ingest pipeline test infrastructure. Created `tmp/` as
scratch space for future ad-hoc development tests.

Promoted four standalone source documents from `pipeline/raw/` to
`wiki/sources/` for use as initial wiki content:
- `wiki/sources/llm-wiki.md` (Karpathy's original concept)
- `wiki/sources/git-branching.md`
- `wiki/sources/qwen-llama-cpp.md`
- `wiki/sources/sqlite-wal.md`

These are now tracked in Git via explicit `.gitignore` exceptions. The
remaining test corpus under `pipeline/raw/web/corpus-2026-05-18/` stays
in the development area.

Touched:
- pipeline/ (created: raw/, drafts/, evals/ moved here)
- wiki/sources/ (4 source docs added)
- tmp/ (created with .gitkeep)
- .gitignore (updated paths)
- AGENTS.md (updated §1)
- README.md (updated)
- wiki/log.md (updated)

## [2026-06-15] schema-change | split project and wiki schemas

Created `wiki/AGENTS.md` as the dedicated wiki maintenance contract,
separating it from the project-level `AGENTS.md`. This follows
Karpathy's pattern: the schema lives with the wiki so the wiki is a
self-contained system.

- `wiki/AGENTS.md` — the contract for LLM agents maintaining the wiki:
  mental model, directory conventions, file conventions, ingest/query/
  lint workflows, log format, contradictions, writing rules.
- `AGENTS.md` (project root) — now focuses on development
  infrastructure: Go CLI, pipeline testing, deployment, repository
  conventions. Points to `wiki/AGENTS.md` for wiki maintenance.

This completes the project/product separation: copy `wiki/` anywhere
and an LLM agent can maintain it by reading `wiki/AGENTS.md` alone.

Touched:
- wiki/AGENTS.md (created)
- AGENTS.md (rewritten — project scope only)
- wiki/log.md (updated)
