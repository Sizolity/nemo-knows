# Reviewed Ingest Apply Plan

Bundle: `drafts/autonomy-r2-2026-05-23-deepseek-moby-dick`

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

- `wiki/sources/moby-dick-raw.md` — create new page.
- `wiki/topics/ahab-and-moby-dick.md` — create new page.
- `wiki/topics/cetology-and-whale-lore-in-moby-dick.md` — create new page.
- `wiki/topics/ishmael-and-queequeg.md` — create new page.
- `wiki/topics/moby-dick-themes-of-fate-and-the-sea.md` — create new page.
- `wiki/topics/prophecy-and-symbolism-in-moby-dick.md` — create new page.
- `wiki/topics/whaling-practice-and-law-in-moby-dick.md` — create new page.

## Required Manual Steps

1. Compare each candidate page against the raw source and cleaned drafts.
2. Create or update approved `wiki/sources/`, `wiki/concepts/`, and `wiki/topics/` pages.
3. Update `wiki/index.md` after accepted page changes.
4. Append an `ingest` entry to `wiki/log.md` after accepted page changes.
5. Re-run wiki lint checks before committing.
