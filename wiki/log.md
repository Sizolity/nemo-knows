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

## [2026-06-15] query-filed | SQLite WAL concurrency and checkpoint trade-offs
Question: SQLite WAL concurrency and checkpoint trade-offs
Touched:
- topics/sqlite-wal-concurrency-checkpoint-trade-offs.md (created)
- index.md (updated)
References:
- wiki/sources/sqlite-wal.md
Open: review whether this filed answer should be expanded after future ingests.

## [2026-06-15] schema-change | index.md is navigation-only

Removed the per-section usage descriptions and the top overview prose
from `index.md`, making it navigation-only (category headings plus
single-line Markdown relative-link entries, or `(none yet)`). Category
meanings now live solely in `AGENTS.md` §0, and §2 gained an "Index
format" convention plus a §9 rule forbidding explanatory prose in the
index. The index append/normalize tooling no longer depends on this
prose and now replaces the `(none yet)` placeholder when a section
gains its first entry.
Touched:
- AGENTS.md (updated)
- index.md (updated)
Open: none.

## [2026-06-15] ingest | SQLite Database Engine Overview
Source: pipeline/raw/web/sqlite.md
Applied bundle: pipeline/drafts/entity-sqlite-20260615
Touched:
- wiki/concepts/serverless-database.md (created)
- wiki/index.md (updated)
- wiki/entities/sqlite.md (created)
- wiki/sources/sqlite.md (created)
- wiki/topics/data-preservation-formats.md (created)
Open: review skipped candidates before creating entity, concept, or topic pages.

## [2026-06-15] lint | sqlite slug uniqueness and title fix
Resolved a cross-category slug collision surfaced during E2E testing: the source
page that ingest created at `wiki/sources/sqlite.md` shared the slug `sqlite`
with `wiki/entities/sqlite.md`, leaving `[[sqlite]]` ambiguous and (via a
slug-only index check) keeping the source page out of the index. Adopted
wiki-wide slug uniqueness; source pages are named after the external document
rather than the entity, so the source page was renamed to `sqlite-overview`.
The entity title casing was also corrected from the slug-derived `Sqlite` back
to `SQLite`.
Touched:
- wiki/sources/sqlite-overview.md (renamed from wiki/sources/sqlite.md)
- wiki/entities/sqlite.md (updated: title cased SQLite; sources cite wiki/sources/sqlite-overview.md)
- wiki/concepts/serverless-database.md (updated: sources cite wiki/sources/sqlite-overview.md)
- wiki/topics/data-preservation-formats.md (updated: sources cite wiki/sources/sqlite-overview.md)
- wiki/index.md (updated: added sqlite-overview under Sources; corrected the entity entry to SQLite)
- wiki/log.md (updated)
Open: none.

## [2026-06-15] query-filed | What makes SQLite suitable for long-term data preservation
Question: What makes SQLite suitable for long-term data preservation
Touched:
- topics/makes-sqlite-suitable-long-term-data.md (created)
- index.md (updated)
References:
- wiki/entities/sqlite.md
- wiki/concepts/serverless-database.md
- wiki/sources/sqlite-overview.md
- wiki/topics/data-preservation-formats.md
- wiki/sources/sqlite-wal.md
Open: review whether this filed answer should be expanded after future ingests.

## [2026-06-16] schema-change | wiki/assets image reference convention

Defined how wiki pages reference binary assets. Images, diagrams, and slides
live under `wiki/assets/<group>/<file>` and are embedded with standard Markdown
image syntax using a page-relative path, e.g.
`![alt](../assets/<group>/<file>.png)`. Local image targets must resolve to an
existing file under `wiki/assets/`; `http(s)://` images are allowed and not
existence-checked. The tooling now backs this convention end to end: the web
console renders `![alt](src)` as inline `<img>` and serves `wiki/assets/`
read-only under `/assets/`, and the lint pass reports missing local images as
`missing-image`. A minimal source page plus a slide figure under
`wiki/assets/c10s1-slides/` exercise the mechanism end to end; additional
MIT 6.004 slide PNGs can be dropped into the same directory and referenced the
same way.

Touched:
- AGENTS.md (updated: §2 gained an "Images and assets" convention)
- wiki/assets/c10s1-slides/ (slide figures for the demo source)
- wiki/sources/mit-6004-c10s1-assembly-models.md (created: demo source page)
- wiki/index.md (updated: added the source under Sources)
- wiki/log.md (updated)
Open: ingest-side handling of binary assets (copying source attachments into
wiki/assets/ during import) is not yet implemented; see the development report.

## [2026-06-16] schema-change | adopt Markdown relative links, retire wikilinks

Replaced the Obsidian-style double-bracket cross-reference convention with
standard Markdown relative links so internal references are direct file jumps
that need no slug-to-path resolution layer in the renderer. A reference target
slug now resolves to the target page's path, and the link is written relative to
the linking page's folder: a sibling page is `name.md` and a page in another
wiki folder is `../folder/name.md`. Labels follow the rule "explicit label wins,
otherwise the target page's frontmatter title, otherwise the slug text"; a
reference whose target page does not exist degrades to plain text instead of a
dangling link.

Existing wiki content was migrated with the shared `internal/wiki` converter: 2
references were rewritten, both pointing at the Data Preservation Formats topic
via a relative `../topics/data-preservation-formats.md` path —
`wiki/entities/sqlite.md` and `wiki/concepts/serverless-database.md`. The lint
and candidate-eval gates now report any residual double-bracket reference as
`forbidden-wikilink` (error) and any relative link whose target file is missing
as `missing-link-target` (error). Slugs stay globally unique and are still used
for page identity and de-duplication; only the cross-reference syntax changed.

Touched:
- wiki/entities/sqlite.md (updated: cross-reference converted to a relative link)
- wiki/concepts/serverless-database.md (updated: cross-reference converted)
- wiki/AGENTS.md (updated: §"Links" convention and lint rules)
- wiki/log.md (updated)
Open: a rename/move workflow that rewrites inbound relative links automatically,
so `missing-link-target` stays empty without manual edits after a page moves.

## [2026-06-16] ingest | llama.cpp
Source: pipeline/raw/web/llama-cpp.md
Applied bundle: pipeline/drafts/e2e-llama-cpp-2026-06-16
Touched:
- wiki/concepts/gguf.md (created)
- wiki/index.md (updated)
- wiki/concepts/llama-cpp-backends.md (created)
- wiki/concepts/quantization.md (created)
- wiki/entities/georgi-gerganov.md (created)
- wiki/entities/llama-cpp.md (created)
- wiki/sources/llama-cpp-overview.md (created)
Open: review skipped candidates before creating entity, concept, or topic pages.

## [2026-06-16] query-filed | How does llama.cpp combine the GGUF format and
Question: How does llama.cpp combine the GGUF format and quantization to enable CPU-first local inference?
Touched:
- topics/llama-cpp-gguf-quantization-cpu-inference.md (created)
- index.md (updated)
References:
- wiki/concepts/gguf.md
- wiki/entities/llama-cpp.md
- wiki/sources/llama-cpp-overview.md
- wiki/concepts/quantization.md
- wiki/concepts/llama-cpp-backends.md
Open: review whether this filed answer should be expanded after future ingests.

## [2026-06-16] ingest | The Transformer Architecture
Source: pipeline/raw/web/corpus-2026-05-18/081-attention-is-all-you-need.md
Applied bundle: tmp/gen-stress-2026-06-16/attention-r1
Touched:
- wiki/concepts/transformer-architecture.md (created)
- wiki/index.md (updated)
- wiki/sources/081-attention-is-all-you-need.md (created)
- wiki/topics/sequence-transduction-models.md (created)
Open: review skipped candidates before creating entity, concept, or topic pages.
