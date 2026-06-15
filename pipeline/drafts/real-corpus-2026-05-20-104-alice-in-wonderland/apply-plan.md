# Reviewed Ingest Apply Plan

Bundle: `drafts/real-corpus-2026-05-20-104-alice-in-wonderland`

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

- `wiki/concepts/dream-within-a-dream-literary-technique.md` — create new page.
- `wiki/concepts/mushroom-wonderland-side-effects.md` — create new page.
- `wiki/concepts/sea-school-curryculum.md` — create new page.
- `wiki/concepts/wonderland-arbitrary-authority.md` — create new page.
- `wiki/concepts/wonderland-size-manipulation.md` — create new page.
- `wiki/sources/alice-wonderland-project-gutenberg-fetch.md` — create new page.
- `wiki/topics/caucus-race-nonsense.md` — create new page.
- `wiki/topics/cheshire-cat-metaphysics.md` — create new page.
- `wiki/topics/lobster-quadrille-dance.md` — create new page.
- `wiki/topics/mad-tea-party-time-personification.md` — create new page.

## Required Manual Steps

1. Compare each candidate page against the raw source and cleaned drafts.
2. Create or update approved `wiki/sources/`, `wiki/concepts/`, and `wiki/topics/` pages.
3. Update `wiki/index.md` after accepted page changes.
4. Append an `ingest` entry to `wiki/log.md` after accepted page changes.
5. Re-run wiki lint checks before committing.
