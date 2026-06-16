# LLM Wiki Core Concept

This document explains the product and engineering idea behind `nemo-knows`.
It is project documentation, not a maintained knowledge-base page.

## One-Sentence Summary

`nemo-knows` turns curated source material into a persistent, LLM-maintained
Markdown wiki so knowledge compounds over time instead of being rediscovered on
every question.

## Why This Is Not Just RAG

In a typical RAG workflow, each query retrieves chunks from raw documents and
asks the model to synthesize an answer on the fly. That is useful, but it means
the system repeatedly reconstructs the same context.

The LLM wiki pattern makes a different tradeoff:

```text
source material -> reviewed draft -> maintained wiki page -> future query
```

The LLM reads a source once, writes or updates durable Markdown pages, adds
links, records contradictions, and keeps summaries current. Future queries read
from that maintained layer instead of starting from raw sources every time.

## Project Layers

`nemo-knows` keeps three layers separate:

```text
pipeline/raw/          immutable test source material
pipeline/drafts/       generated candidate pages for review
pipeline/evals/runs/   deterministic evaluation outputs
wiki/sources/          accepted production source pages
wiki/                  accepted LLM-maintained knowledge base
```

`pipeline/raw/` is immutable development test source material and should not be
modified by automation. Production wiki source pages live under
`wiki/sources/`.

`pipeline/drafts/` is the engineering safety buffer. Local model output lands
here first so humans or agents can inspect it before anything becomes
maintained wiki content.

`wiki/` is the accepted knowledge layer. It contains source summaries, entity
pages, concept pages, topic pages, indexes, and logs that should stay internally
linked and current.

`AGENTS.md` defines the maintenance contract for agents that edit the wiki.
Project-level Cursor rules under `.cursor/rules/` define coding behavior for
agents working on this repository.

## Core Workflows

### Draft

The development pipeline turns a `pipeline/raw/` source and a prompt template
into draft files:

```text
pipeline/raw/<source>.md + prompts/<template>.md
  -> llama.cpp
  -> pipeline/drafts/<name>.raw.txt
  -> pipeline/drafts/<name>.md
```

The raw draft preserves model/runtime output for debugging. The cleaned draft is
a candidate Markdown page, not automatically accepted wiki content.

### Local Ingest Bundle

The local ingest MVP extends the draft workflow by generating a small bundle of
review artifacts from one `pipeline/raw/` source:

```text
pipeline/raw/<source>.md
  -> pipeline/drafts/<source>/source.raw.txt
  -> pipeline/drafts/<source>/source.md
  -> pipeline/drafts/<source>/ingest-plan.raw.txt
  -> pipeline/drafts/<source>/ingest-plan.md
```

The bundle is still outside the maintained wiki. Its purpose is to test whether
the local 9B model can assist with source summaries and maintenance planning
before any accepted wiki edits are made.

### Reviewed Ingest Helper

The reviewed ingest helper reads a local ingest bundle and produces an
`apply-plan.md` review artifact:

```text
pipeline/drafts/<source>/source.md
pipeline/drafts/<source>/ingest-plan.md
  -> deterministic validation
  -> pipeline/drafts/<source>/apply-plan.md
```

This step does not call the model and does not write to `wiki/`. It validates
the draft structure, extracts candidate wiki page paths, and records the manual
steps required before accepted wiki edits.

### Reviewed Candidate Draft Generation

Candidate draft generation turns reviewed concept and topic candidate paths into
separate draft pages:

```text
pipeline/drafts/<source>/apply-plan.md
pipeline/drafts/<source>/source.md
  -> pipeline/drafts/<source>/candidates/wiki/entities|concepts|topics/*.md
```

This step calls the model, but still does not write to `wiki/`. It gives
approved apply full page drafts to validate and apply, instead of promoting
one-line candidate descriptions from the ingest plan.

### Candidate Draft Evaluation Harness

Candidate draft evaluation scores generated concept and topic drafts before
approved apply can consume them:

```text
pipeline/drafts/<source>/candidates/wiki/entities|concepts|topics/*.md
  -> pipeline/evals/runs/<run-id>/candidate-scores.json
  -> pipeline/evals/runs/<run-id>/candidate-trace.md
```

The harness is deterministic. It checks frontmatter, source references, title
and heading consistency, semantic wikilink safety when links are present, draft
length, and whether the draft is mostly copied from `source.md`.

### Candidate Review And Link Repair

Candidate review turns evaluation findings into a deterministic repair report:

```text
candidate eval trace
  -> pipeline/evals/runs/<run-id>/candidate-review.md
```

This step is advisory. It does not call the model, rewrite candidate pages, or
write to `wiki/`. It explains what a reviewer should fix, such as converting
source-unsupported wikilinks to plain text or regenerating candidates that lack
durable source references.

### Multi-Source Regression Evals

Regression evals run the deterministic harness over many fixture cases:

```text
pipeline/evals/cases/*/bundle
  -> pipeline/evals/runs/<run-id>/regression-summary.json
  -> pipeline/evals/runs/<run-id>/regression-summary.md
```

This catches regressions in review and scoring logic across source shapes such
as short articles, technical documents, and loose meeting notes before the tool
is trusted on more real ingests.

### Candidate Link Quality Gate

Candidate generation is constrained by an Allowed Links list assembled from
source-supported existing wiki pages and the entity/concept/topic candidates in
the reviewed apply plan:

```text
source.md + wiki/*.md + apply-plan candidate paths
  -> prompt Allowed Links
  -> candidate draft wikilink normalization
  -> candidate eval missing-target check
```

This keeps model-invented semantic links such as `[[RAG]]` or `[[Memex]]` from
becoming broken wiki references unless those pages already exist or are
explicitly part of the reviewed candidate set.

The gate does not require every candidate to contain a wikilink. A missing link
is acceptable when the source does not support a strong cross-reference; weak
semantic linking is worse than plain text.

The gate also distinguishes a valid target from a useful target. A wikilink to
an existing page can still be marked `borderline` when the link target is not
supported by the current source material or by the reviewed candidate set. This
keeps wikilinks aligned with the wiki's purpose: navigation should make the
knowledge base more accurate and more connected, not merely satisfy a structural
check.

Both `wiki/index.md` entries and body-page semantic cross-references use the
same form: standard Markdown links whose target is a path relative to the page
that contains them, for example `[sqlite-wal](sources/sqlite-wal.md)` from the
index or `[SQLite](../entities/sqlite.md)` from a sibling folder. Obsidian-style
`[[wikilinks]]` are no longer part of the schema; the lint and candidate-eval
gates report any residual `[[...]]` so links stay real, resolvable file jumps.

### Ingest Evaluation Harness

The ingest evaluation harness scores generated and reviewed artifacts:

```text
pipeline/drafts/<source>/source.md
pipeline/drafts/<source>/ingest-plan.md
pipeline/drafts/<source>/apply-plan.md
  -> pipeline/evals/runs/<run-id>/scores.json
  -> pipeline/evals/runs/<run-id>/trace.md
```

This is the OpenAI-style eval layer around the Claude-style
plan/generate/review workflow. It makes prompt, profile, and review-logic
changes comparable across repeated runs.

### Approved Wiki Apply

The approved apply step is the first guarded write into `wiki/`:

```text
pipeline/drafts/<source>/apply-plan.md
pipeline/evals/runs/<run-id>/scores.json
  -> explicit approval
  -> wiki/sources|entities|concepts|topics
  -> wiki/index.md
  -> wiki/log.md
  -> pipeline/drafts/<source>/apply-report.md
```

It keeps the project aligned with the original goal: the wiki must actually
grow, but accepted wiki edits remain explicit and auditable. Source pages can be
applied from `source.md`; entity, concept, and topic pages require matching
reviewed drafts under `pipeline/drafts/<source>/candidates/` before they can be
written.

For public web material, the CLI may explicitly persist a testing source under
`pipeline/raw/web/<slug>.md` before bundle generation. This is opt-in and
create-only: automation must not overwrite existing raw files. Once persisted,
generated source summaries and candidate pages cite the durable
`pipeline/raw/web/...` path instead of a temporary `pipeline/drafts/...` input,
preserving the same source-of-truth boundary as normal test-source ingest.

### Ingest

Ingest is the reviewed step that moves knowledge from a source or draft into the
wiki. It may create a source summary, update concept or topic pages, refresh the
index, and append to the log.

### Query

Query answers should start from `wiki/index.md`, read the relevant wiki pages,
and synthesize an answer from maintained knowledge. If the answer captures a
useful comparison, explanation, or decision, it can be filed back as a topic page.

### Lint

Lint is a maintenance pass over `wiki/`. It looks for broken links, missing
pages, contradictions, stale claims, and concepts that deserve their own page.
The first automated lint harness is deterministic and read-only: it reports
frontmatter issues, duplicate Markdown or legacy wikilink index entries,
missing semantic wikilink targets, orphan pages, and invalid log actions after
approved apply.

## Engineering Implications

The implementation should preserve review boundaries. Model output may be
generated automatically, but accepted wiki edits should remain explicit and
auditable.

The first useful CLI surface is intentionally small: render a prompt, call the
local model, clean the output, and write drafts. Higher-level commands such as
`nemo ingest`, `nemo query`, and `nemo lint` can build on that once the file
contracts are stable.

The system should prefer plain files over infrastructure. Markdown, Git, simple
validation, and explicit logs are enough for the early version; embeddings,
servers, databases, and workflow frameworks are optional future extensions.

## Related Documentation

- [`README.md`](../../README.md) for the user-facing project overview.
- [`docs/development/minimal-go-implementation.md`](../development/minimal-go-implementation.md) for the first Go implementation plan.
- [`docs/development/local-ingest-mvp.md`](../development/local-ingest-mvp.md) for the local-only ingest bundle design.
- [`docs/development/reviewed-ingest-helper-mvp.md`](../development/reviewed-ingest-helper-mvp.md) for reviewing generated bundles before wiki edits.
- [`docs/development/ingest-eval-harness-mvp.md`](../development/ingest-eval-harness-mvp.md) for evaluating ingest runs over time.
- [`docs/development/approved-wiki-apply-mvp.md`](../development/approved-wiki-apply-mvp.md) for explicitly approved wiki writes.
- [`docs/development/qwen3-5-huggingface-parameters.md`](../development/qwen3-5-huggingface-parameters.md) for Qwen3.5 Hugging Face generation settings.
- [`AGENTS.md`](../../AGENTS.md) for the wiki maintenance contract.
