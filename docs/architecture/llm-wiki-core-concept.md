# LLM Wiki Core Concept

This document explains the product and engineering idea behind `nemo-knows`.
It is project documentation, not a maintained knowledge-base page.

## One-Sentence Summary

`nemo-knows` turns raw source files into a persistent, LLM-maintained Markdown
wiki so knowledge compounds over time instead of being rediscovered on every
question.

## Why This Is Not Just RAG

In a typical RAG workflow, each query retrieves chunks from raw documents and
asks the model to synthesize an answer on the fly. That is useful, but it means
the system repeatedly reconstructs the same context.

The LLM wiki pattern makes a different tradeoff:

```text
raw source -> reviewed draft -> maintained wiki page -> future query
```

The LLM reads a source once, writes or updates durable Markdown pages, adds
links, records contradictions, and keeps summaries current. Future queries read
from that maintained layer instead of starting from raw sources every time.

## Project Layers

`nemo-knows` keeps three layers separate:

```text
raw/      immutable source material
drafts/   generated candidate pages for review
wiki/     accepted LLM-maintained knowledge base
```

`raw/` is the source of truth and should not be modified by automation.

`drafts/` is the engineering safety buffer. Local model output lands here first
so humans or agents can inspect it before anything becomes maintained wiki
content.

`wiki/` is the accepted knowledge layer. It contains source summaries, concept
pages, topic pages, indexes, and logs that should stay internally linked and
current.

`AGENTS.md` defines the maintenance contract for agents that edit the wiki.
Project-level Cursor rules under `.cursor/rules/` define coding behavior for
agents working on this repository.

## Core Workflows

### Draft

The Go command turns a raw source and a prompt template into draft files:

```text
raw/<source>.md + prompts/<template>.md
  -> llama.cpp
  -> drafts/<name>.raw.txt
  -> drafts/<name>.md
```

The raw draft preserves model/runtime output for debugging. The cleaned draft is
a candidate Markdown page, not automatically accepted wiki content.

### Local Ingest Bundle

The local ingest MVP extends the draft workflow by generating a small bundle of
review artifacts from one raw source:

```text
raw/<source>.md
  -> drafts/<source>/source.raw.txt
  -> drafts/<source>/source.md
  -> drafts/<source>/ingest-plan.raw.txt
  -> drafts/<source>/ingest-plan.md
```

The bundle is still outside the maintained wiki. Its purpose is to test whether
the local 9B model can assist with source summaries and maintenance planning
before any accepted wiki edits are made.

### Reviewed Ingest Helper

The reviewed ingest helper reads a local ingest bundle and produces an
`apply-plan.md` review artifact:

```text
drafts/<source>/source.md
drafts/<source>/ingest-plan.md
  -> deterministic validation
  -> drafts/<source>/apply-plan.md
```

This step does not call the model and does not write to `wiki/`. It validates
the draft structure, extracts candidate wiki page paths, and records the manual
steps required before accepted wiki edits.

### Reviewed Candidate Draft Generation

Candidate draft generation turns reviewed concept and topic candidate paths into
separate draft pages:

```text
drafts/<source>/apply-plan.md
drafts/<source>/source.md
  -> drafts/<source>/candidates/wiki/concepts|topics/*.md
```

This step calls the model, but still does not write to `wiki/`. It gives
approved apply full page drafts to validate and apply, instead of promoting
one-line candidate descriptions from the ingest plan.

### Candidate Draft Evaluation Harness

Candidate draft evaluation scores generated concept and topic drafts before
approved apply can consume them:

```text
drafts/<source>/candidates/wiki/concepts|topics/*.md
  -> evals/runs/<run-id>/candidate-scores.json
  -> evals/runs/<run-id>/candidate-trace.md
```

The harness is deterministic. It checks frontmatter, source references, title
and heading consistency, wikilink safety when links are present, draft length,
and whether the draft is mostly copied from `source.md`.

### Candidate Review And Link Repair

Candidate review turns evaluation findings into a deterministic repair report:

```text
candidate eval trace
  -> evals/runs/<run-id>/candidate-review.md
```

This step is advisory. It does not call the model, rewrite candidate pages, or
write to `wiki/`. It explains what a reviewer should fix, such as converting
source-unsupported wikilinks to plain text or regenerating candidates that lack
durable source references.

### Multi-Source Regression Evals

Regression evals run the deterministic harness over many fixture cases:

```text
evals/cases/*/bundle
  -> evals/runs/<run-id>/regression-summary.json
  -> evals/runs/<run-id>/regression-summary.md
```

This catches regressions in review and scoring logic across source shapes such
as short articles, technical documents, and loose meeting notes before the tool
is trusted on more real ingests.

### Candidate Link Quality Gate

Candidate generation is constrained by an Allowed Links list assembled from
source-supported existing wiki pages and the concept/topic candidates in the
reviewed apply plan:

```text
source.md + wiki/*.md + apply-plan candidate paths
  -> prompt Allowed Links
  -> candidate draft wikilink normalization
  -> candidate eval missing-target check
```

This keeps model-invented links such as `[[RAG]]` or `[[Memex]]` from becoming
broken wiki references unless those pages already exist or are explicitly part
of the reviewed candidate set.

The gate does not require every candidate to contain a wikilink. A missing link
is acceptable when the source does not support a strong cross-reference; weak
navigation is worse than plain text.

The gate also distinguishes a valid target from a useful target. A wikilink to
an existing page can still be marked `borderline` when the link target is not
supported by the current source material or by the reviewed candidate set. This
keeps wikilinks aligned with the wiki's purpose: navigation should make the
knowledge base more accurate and more connected, not merely satisfy a structural
check.

### Ingest Evaluation Harness

The ingest evaluation harness scores generated and reviewed artifacts:

```text
drafts/<source>/source.md
drafts/<source>/ingest-plan.md
drafts/<source>/apply-plan.md
  -> evals/runs/<run-id>/scores.json
  -> evals/runs/<run-id>/trace.md
```

This is the OpenAI-style eval layer around the Claude-style
plan/generate/review workflow. It makes prompt, profile, and review-logic
changes comparable across repeated runs.

### Approved Wiki Apply

The approved apply step is the first guarded write into `wiki/`:

```text
drafts/<source>/apply-plan.md
evals/runs/<run-id>/scores.json
  -> explicit approval
  -> wiki/sources|concepts|topics
  -> wiki/index.md
  -> wiki/log.md
  -> drafts/<source>/apply-report.md
```

It keeps the project aligned with the original goal: the wiki must actually
grow, but accepted wiki edits remain explicit and auditable. Source pages can be
applied from `source.md`; concept and topic pages require matching reviewed
drafts under `drafts/<source>/candidates/` before they can be written.

For public web material, the CLI may explicitly persist a testing source under
`raw/web/<slug>.md` before bundle generation. This is opt-in and create-only:
automation must not overwrite existing raw files. Once persisted, generated
source summaries and candidate pages cite the durable `raw/web/...` path instead
of a temporary `drafts/...` input, preserving the same source-of-truth boundary as
normal raw-file ingest.

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
frontmatter issues, duplicate index entries, missing wikilink targets, orphan
pages, and invalid log actions after approved apply.

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
