# Reviewed Ingest Apply Plan

Bundle: `drafts/stress-2026-05-21-deepseek-moby-dick-forced-chunk`

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

- `wiki/concepts/fast-fish-loose-fish-principle.md` — create new page.
- `wiki/sources/moby-dick-novel.md` — create new page.
- `wiki/topics/ahab-monomania-and-crew-fate.md` — create new page.
- `wiki/topics/cetology-and-whale-anatomy-in-moby-dick.md` — create new page.
- `wiki/topics/ishmael-and-queequeg-friendship.md` — create new page.
- `wiki/topics/moby-dick-the-white-whale.md` — create new page.
- `wiki/topics/whaling-life-practices-and-customs.md` — create new page.

## Required Manual Steps

1. Compare each candidate page against the raw source and cleaned drafts.
2. Create or update approved `wiki/sources/`, `wiki/concepts/`, and `wiki/topics/` pages.
3. Update `wiki/index.md` after accepted page changes.
4. Append an `ingest` entry to `wiki/log.md` after accepted page changes.
5. Re-run wiki lint checks before committing.
