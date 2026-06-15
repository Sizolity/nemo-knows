# Multi-Source Regression Evals MVP

This document describes the tenth MVP: running deterministic ingest regression
checks across multiple eval cases.

## Goal

Move from one-off validation to repeatable regression checks:

```text
pipeline/evals/cases/*/bundle -> regression-summary.json + regression-summary.md
```

Earlier MVPs proved each stage on the `llm-wiki` case. MVP-10 checks whether
the review and evaluation rules stay stable across different source shapes.

## Case Layout

Each case lives under `pipeline/evals/cases/<case-name>/`:

```text
pipeline/evals/cases/<case-name>/
  expected.json
  bundle/
    source.md
    ingest-plan.md
    apply-plan.md
```

`expected.json` records the source label and minimum acceptable scores:

```json
{
  "case": "technical-doc",
  "source": "fixture:technical-doc",
  "bundle": "pipeline/evals/cases/technical-doc/bundle",
  "minimum_scores": {
    "schema": "pass",
    "wiki_safety": "pass",
    "candidate_paths": "pass",
    "duplicate_detection": "pass",
    "apply_readiness": "pass",
    "overall": "pass"
  }
}
```

MVP-10 uses fixture bundles rather than adding new `pipeline/raw/` sources. This
keeps `pipeline/raw/` immutable while still exercising the deterministic
harness over multiple artifact shapes.

## Command

```sh
go run ./cmd/nemo \
  -eval-regression pipeline/evals/cases \
  -out-dir pipeline/evals/runs/regression
```

## Outputs

```text
pipeline/evals/runs/regression/regression-summary.json
pipeline/evals/runs/regression/regression-summary.md
```

The summary records every case, actual scores, expected minimum scores, pass/fail
status, and trace entries from the underlying bundle evaluator.

## Initial Case Types

MVP-10 starts with three case shapes:

- `llm-wiki`: short blog/source-summary style.
- `technical-doc`: structured technical documentation with mechanisms and paths.
- `meeting-notes`: loose notes with decisions, entities, and follow-ups.

## Acceptance Criteria

This MVP is successful when:

- `go test ./...` passes.
- the regression runner evaluates every case under `pipeline/evals/cases`.
- case failures are represented in JSON, not hidden in prose.
- the real command produces a passing summary for the initial cases.
- no command writes to `wiki/`.
