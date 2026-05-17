# Candidate Review Link Repair MVP

This document describes the eleventh MVP: turning candidate evaluation findings
into a deterministic review artifact before approved apply.

## Goal

Bridge the gap between candidate scoring and human repair:

```text
candidates/wiki/concepts|topics/*.md
  -> candidate-scores.json
  -> candidate-review.md
  -> manual repair or approved apply decision
```

The review step does not call the model and does not edit candidate files. It
summarizes problems that are already visible in candidate eval traces and
suggests conservative fixes.

## Command

```sh
go run ./cmd/nemo \
  -review-candidates drafts/web-e2e-sqlite \
  -out-dir evals/runs/web-e2e/sqlite-review
```

Expected output:

```text
evals/runs/<run-id>/candidate-review.md
```

## Repair Policy

The command is intentionally advisory:

- It never writes to `wiki/`.
- It never rewrites candidate drafts automatically.
- It includes a no-auto-apply warning in the report.
- It recommends plain text over weak wikilinks unless source support is added.
- It treats missing links, weak semantic links, missing durable sources,
  title/frontmatter problems, short drafts, and copied prose as review items.

This keeps the wiki aligned with the core project model: `raw/` is the durable
source layer, `drafts/` is the model-output buffer, and `wiki/` changes remain
explicit and auditable.

## Acceptance Criteria

This MVP is successful when:

- `go test ./...` passes.
- `-review-candidates` writes `candidate-review.md`.
- weak semantic links produce a recommendation to convert unsupported wikilinks
  to plain text unless support is added.
- clean candidate eval results produce a review report with no suggested repairs.
- the command remains deterministic and does not call llama.cpp.
