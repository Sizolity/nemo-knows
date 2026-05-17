# Web E2E Production Observation

Date: 2026-05-17

Eval scores, traces, and candidate-review Markdown under `evals/runs/` are
**local outputs** from `nemo -eval-*` / `-review-*` commands; that directory is
gitignored. Re-run those commands after changing bundles or prompts if you need
fresh JSON/Markdown artifacts.

## Inputs

- `drafts/web-e2e-sources/qwen-llama-cpp.md`
- `drafts/web-e2e-sources/sqlite-wal.md`
- `drafts/web-e2e-sources/git-branching.md`

The initial run used public-web excerpts stored under `drafts/` for testing.
After the source-durability change, the same inputs were explicitly persisted to:

- `raw/web/qwen-llama-cpp.md`
- `raw/web/sqlite-wal.md`
- `raw/web/git-branching.md`

This keeps web-derived tests aligned with the normal ingest rule that durable
claims trace to `raw/...` or `wiki/sources/...`.

## Runtime

| Stage | Qwen llama.cpp | SQLite WAL | Git Branching |
| --- | ---: | ---: | ---: |
| Bundle generation, first run | 21.070s | 15.917s | 16.946s |
| Candidate generation, first run | 21.418s | 17.344s | 17.265s |
| Bundle generation, durable raw run | 22.325s | 17.212s | 16.385s |
| Candidate generation, durable raw run | 18.018s | 21.477s | 18.527s |
| Bundle generation, source-normalized raw run | 22.590s | 16.621s | 18.703s |
| Candidate generation after source normalization | 26.398s | 16.401s | 15.490s |

No `.fallback.raw.txt` files were produced. No visible thinking markers were
found in raw outputs.

## Bundle Evaluation

All three reviewed bundles passed deterministic bundle evaluation (run
`-eval-bundle` per bundle with an `-out-dir` of your choice; compare to
`expected.json` shapes under `evals/cases/*-web/` where applicable).

Each bundle produced one source page candidate plus two concept/topic page
candidates.

## Candidate Evaluation

The initial candidate evaluation failed for all three bundles because generated
candidate frontmatter used:

```yaml
sources:
  - source.md
```

The candidate evaluator requires a durable `raw/...` or `wiki/sources/...`
reference. This was correct for production: a temporary `drafts/` input should
not be silently accepted as a durable wiki source.

After rerunning with durable `raw/web/...` inputs and source-supported link
constraints, candidate source attribution passed for all three bundles. Current
candidate results:

- Qwen llama.cpp: `pass`.
- SQLite WAL: `pass`.
- Git Branching: `pass`.

Observed content issues:

- Some Qwen and SQLite candidate links were target-resolvable but semantically
  weak. The new semantic link gate correctly marks these as `borderline`.
- Git candidates improved on rerun: source attribution, links, length, and
  originality all passed.

## Assessment

The local inference path is operationally stable for these short public-web
inputs. Source durability is now handled by an explicit create-only
`raw/web/<slug>.md` persistence step before bundle generation.

The link quality gate now prevents broken links and also reports weak semantic
links. This does not deviate from the core wiki idea: it preserves immutable raw
sources, keeps model output in `drafts/`, and uses evaluation to prevent
structurally valid but poorly justified links from entering the maintained wiki
without review.

## Candidate Review

MVP-11 added a deterministic review artifact on top of candidate eval (written
to `-out-dir` when you run `-review-candidates`).

The review reports convert candidate eval findings into manual repair guidance.
For an intermediate run:

- Qwen llama.cpp produced one repair item for weak links in
  `wiki/concepts/llama-cli-arguments.md`.
- SQLite WAL produced one repair item for weak links in
  `wiki/concepts/database-journal-modes.md`.
- Git Branching produced no repair items.

The review step remains advisory and deterministic: it does not call the model,
rewrite candidate drafts, or write to `wiki/`.

## Link Root-Cause Fix

After reviewing the weak-link cases, the pipeline was changed to treat wikilinks
as optional rather than mandatory. Candidate generation now passes only
source-supported existing wiki pages plus reviewed candidate pages in the
Allowed Links list, and deterministic cleanup no longer inserts a wikilink just
to satisfy structure.

After regenerating candidates with this policy:

- Qwen llama.cpp: `pass`
- SQLite WAL: `pass`
- Git Branching: `pass`

The candidate review reports then contained zero repair items for all three web E2E
bundles. This supports the project principle that plain text is preferable to
weak navigation: the wiki should become more connected only when the source
material supports the connection.

## Approved Apply Dry-Run

MVP-12 validated the final promotion step in an isolated temporary root rather
than writing to the real `wiki/`:

```text
raw/web/... -> bundle -> candidates -> eval pass -> review pass -> approved apply -> wiki lint
```

The first dry-run surfaced two promotion gaps:

- `source.md` drafts did not always include valid `confidence` frontmatter, so
  applied `wiki/sources/*` pages could fail wiki lint.
- New `wiki/sources/*` pages were written to disk but not added to
  `wiki/index.md`, leaving them as orphan pages.

Both gaps were fixed deterministically:

- bundle generation now normalizes `source.md` frontmatter with `title`,
  `kind: source`, durable `sources`, and `confidence: medium`;
- approved apply now indexes newly created source pages in `## Sources`, just as
  it already indexed concept and topic pages.

The final isolated dry-run applied all three web E2E bundles successfully:

- Qwen llama.cpp wrote source, concept, topic, index, log, and apply report;
- SQLite WAL wrote source, concept, topic, index, log, and apply report;
- Git Branching wrote source, concept, topic, index, log, and apply report.

Post-apply wiki lint reported:

```text
total issues: 0
pages checked: 21
```

This confirms the production promotion loop is structurally ready for reviewed
candidate pages: applied pages keep durable `raw/web/...` source references,
`wiki/index.md` is updated for sources/concepts/topics, `wiki/log.md` receives
append-only ingest entries, and `apply-report.md` records no skipped files.
