# Reviewed Ingest Apply Plan

Bundle: `drafts/real-corpus-2026-05-19-032-go-modules-reference-group`

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

- `wiki/concepts/go-mod-directives.md` — create new page.
- `wiki/concepts/go-mod-edit-manual.md` — create new page.
- `wiki/concepts/go-mod-versioning.md` — create new page.
- `wiki/concepts/go-modules-vendoring-behavior.md` — create new page.
- `wiki/concepts/go-workspace-files.md` — create new page.
- `wiki/concepts/legacy-module-compatibility.md` — create new page.
- `wiki/concepts/minimal-version-selection.md` — create new page.
- `wiki/concepts/module-security-and-cache.md` — create new page.
- `wiki/sources/go-modules-reference.md` — create new page.
- `wiki/topics/indirect-dependencies-in-go-mod.md` — create new page.
- `wiki/topics/module-environment-variables.md` — create new page.
- `wiki/topics/module-upgrade-strategies.md` — create new page.
- `wiki/topics/proxy-protocol-specification.md` — create new page.

## Required Manual Steps

1. Compare each candidate page against the raw source and cleaned drafts.
2. Create or update approved `wiki/sources/`, `wiki/concepts/`, and `wiki/topics/` pages.
3. Update `wiki/index.md` after accepted page changes.
4. Append an `ingest` entry to `wiki/log.md` after accepted page changes.
5. Re-run wiki lint checks before committing.
