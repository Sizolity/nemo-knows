# Candidate Draft Evaluation Harness MVP

This document describes the eighth MVP: deterministic scoring for reviewed
candidate page drafts before approved apply writes them to `wiki/`.

## Goal

Catch low-quality concept and topic drafts before they become maintained wiki
pages:

```text
candidates/wiki/concepts|topics/*.md -> candidate-scores.json + candidate-trace.md
```

MVP-7 generates candidate drafts. MVP-8 checks whether those drafts satisfy the
minimum contract expected by approved apply and the wiki schema.

## Command

```sh
go run ./cmd/nemo \
  -eval-candidates drafts/mvp7-real-test-llm-wiki \
  -out-dir evals/runs/mvp7-real-test
```

The command does not call the model and does not write to `wiki/`.

## Scoring Dimensions

Each candidate draft is scored on:

- `frontmatter`: YAML frontmatter exists, includes `title`, `kind`, `sources`,
  and `confidence`, and `kind` matches the target directory.
- `sources`: `sources` includes `source.md` and at least one durable source
  reference such as `raw/...` or `wiki/sources/...`.
- `title`: frontmatter title exists and the body has a matching top-level
  heading.
- `wikilinks`: wikilinks are optional; when present, every target resolves to
  either an existing `wiki/` page or a concept/topic candidate named in the
  reviewed apply plan; and existing-page links are marked `borderline` when the
  target is not supported by the current source material or reviewed candidates.
- `length`: body is long enough to be useful but not a full article.
- `originality`: draft is not mostly copied line-for-line from `source.md`.

Scores are coarse and stable: `pass`, `borderline`, or `fail`.

## Outputs

```text
evals/runs/<run-id>/candidate-scores.json
evals/runs/<run-id>/candidate-trace.md
```

The JSON output records aggregate scores and one result per candidate. The trace
explains which checks passed, failed, or need review.

## Acceptance Criteria

This MVP is successful when:

- `go test ./...` passes.
- missing candidate drafts fail the aggregate score.
- missing `sources` frontmatter fails the candidate.
- missing wikilinks is allowed; no link is better than a weak link.
- wikilinks that target missing pages are `borderline`, so link-quality issues
  are visible before approved apply and post-apply lint.
- wikilinks that target existing but source-unsupported pages are also
  `borderline`, so structurally valid but semantically weak links are visible
  before approved apply.
- candidate generation now normalizes linked or mismatched H1 headings and
  downgrades unsupported wikilinks to plain text, so eval failures in these
  dimensions usually indicate a regression in deterministic cleanup.
- short drafts fail length scoring.
- near-copy drafts are marked `borderline`.
- real MVP-7 generated candidates can be evaluated without writing to `wiki/`.
