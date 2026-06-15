# Reviewed Ingest Apply Plan

Bundle: `drafts/real-corpus-2026-05-19-031-effective-go-group`

This is a review artifact. Do not apply this plan automatically.

## Validation

- [x] `source.md` has YAML frontmatter
- [x] `ingest-plan.md` has YAML frontmatter
- [x] `source.md` frontmatter `kind` is `source`
- [x] `ingest-plan.md` frontmatter `kind` is `topic`
- [x] `source.md` includes required section `What It Is`
- [x] `source.md` includes required section `Summary`
- [x] `source.md` includes required section `Key Claims`
- [x] `source.md` includes required section `Suggested Links`
- [x] `ingest-plan.md` includes required section `Source Summary`
- [x] `ingest-plan.md` includes required section `Candidate Wiki Pages`
- [x] `ingest-plan.md` includes required section `Suggested Links`
- [x] `ingest-plan.md` includes required section `Review Checklist`

## Candidate Changes

- `wiki/concepts/control-flow-go.md` — create new page.
- `wiki/concepts/go-fmt-conventions.md` — create new page.
- `wiki/concepts/interfaces-and-embedding.md` — create new page.
- `wiki/concepts/memory-allocation-go.md` — create new page.
- `wiki/concepts/naming-go-idioms.md` — create new page.
- `wiki/sources/effective-go.md` — create new page.
- `wiki/topics/go-concurrency-patterns.md` — create new page.
- `wiki/topics/go-error-handling-strategies.md` — create new page.

## Required Manual Steps

1. Compare each candidate page against the raw source and cleaned drafts.
2. Create or update approved `wiki/sources/`, `wiki/concepts/`, and `wiki/topics/` pages.
3. Update `wiki/index.md` after accepted page changes.
4. Append an `ingest` entry to `wiki/log.md` after accepted page changes.
5. Re-run wiki lint checks before committing.
