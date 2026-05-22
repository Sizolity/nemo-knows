# Ingest Evaluation Harness MVP

This document describes the fourth MVP: a repeatable harness for evaluating
local ingest runs.

## Goal

Move from prompt-only iteration to harness engineering:

```text
case -> generate/review artifacts -> deterministic scores -> trace -> compare runs
```

The harness should answer whether a change to prompts, profiles, review logic,
or model settings improves the ingest workflow over time.

## Architecture

`nemo-knows` should use a hybrid workflow:

```text
OpenAI-style eval harness outside
Claude-style plan/generate/review loop inside
```

The inner loop handles one real ingest:

```text
raw source -> source.md -> ingest-plan.md -> apply-plan.md
```

The outer harness evaluates the artifacts:

```text
apply-plan.md + bundle artifacts -> scores.json + trace.md
```

This keeps the LLM workflow auditable while making regressions measurable.

## Why Harness, Not Prompt Only

Prompt changes alone are not enough because the project needs to preserve a
long-lived wiki. A better prompt can still create duplicate pages, miss existing
pages, omit required frontmatter, or produce an apply plan that looks plausible
but is not safe to use.

The harness makes quality explicit and repeatable:

- schema checks,
- safety checks,
- candidate path checks,
- duplicate-detection checks,
- source-completeness checks,
- apply-readiness checks,
- run traces for debugging.

## Initial Directory Layout

```text
evals/
  cases/
    llm-wiki/
      expected.json
  runs/
    <run-id>/
      scores.json
      trace.md
```

The first MVP can evaluate an existing bundle rather than regenerating model
output. This keeps the harness deterministic and fast.

## Command

```sh
go run ./cmd/nemo \
  -eval-bundle drafts/actual-use-llm-wiki \
  -out-dir evals/runs/llm-wiki-manual
```

Expected outputs:

```text
evals/runs/llm-wiki-manual/scores.json
evals/runs/llm-wiki-manual/trace.md
```

## Scoring

Use coarse scores, not false-precision numeric grades:

```text
pass
borderline
fail
```

Initial score dimensions:

- `schema`: required files, frontmatter, kinds, and sections.
- `wiki_safety`: no automatic wiki writes; apply plan remains a review artifact.
- `candidate_paths`: candidate paths are legal wiki knowledge-page paths.
- `duplicate_detection`: exact existing pages or likely duplicate pages are
  called out.
- `source_completeness`: generated source drafts do not make completeness claims
  when the referenced raw source has an explicit truncation marker.
- `apply_plan_coverage`: every page listed in the apply plan is represented by
  a generated artifact before the bundle is treated as ready.
- `apply_readiness`: the apply plan is good enough for human review.
- `overall`: aggregate decision.

## Initial Rules

`schema` is `pass` when the reviewed bundle produces an apply plan with all
required validation checks.

`wiki_safety` is `pass` when the apply plan includes "Do not apply this plan
automatically." and no command writes into `wiki/`.

`candidate_paths` is `pass` when all candidate paths are under
`wiki/sources/`, `wiki/concepts/`, or `wiki/topics/`.

`duplicate_detection` is:

- `pass` when no candidate appears duplicated or when every suspected duplicate
  is explicitly labeled.
- `borderline` when new candidates look likely to overlap existing pages but no
  duplicate label appears.
- `fail` when an exact existing page is proposed as a create.

`source_completeness` is:

- `pass` when referenced raw sources have no truncation marker, or when the
  generated source draft/ingest plan acknowledges the truncation boundary.
- `borderline` when a referenced raw source contains `[truncated at ...]` but
  the generated drafts do not mention truncation or incompleteness.
- `fail` when a referenced raw source contains `[truncated at ...]` and the
  generated drafts claim the source is complete, entire, full text, all
  chapters, final chapters, or without abridgment.

`apply_plan_coverage` is:

- `pass` when each planned `wiki/concepts/` or `wiki/topics/` page has a
  generated candidate draft, and planned `wiki/sources/` pages are represented
  by the generated `source.md`.
- `borderline` when multiple planned `wiki/sources/` pages share the single
  generated `source.md` artifact.
- `fail` when planned concept/topic candidate drafts are missing, or when a
  planned source page has no valid `source.md` artifact.

`apply_readiness` is:

- `pass` when all dimensions pass.
- `borderline` when only duplicate, source-completeness, apply-plan coverage, or
  candidate quality needs human judgment.
- `fail` when schema, safety, candidate paths, source completeness, or apply
  plan coverage fail.

## Acceptance Criteria

This MVP is successful when:

- `go test ./...` passes.
- `nemo -eval-bundle` writes `scores.json` and `trace.md`.
- the first real `llm-wiki` bundle receives deterministic scores.
- failures are represented in `scores.json`, not hidden in prose.
- `wiki/` remains unchanged by the harness.

## Next Step

After this MVP, add multiple eval cases and compare runs across prompt/profile
changes. Later, add an optional LLM judge for source fidelity, but keep
deterministic checks as the safety baseline.
