# Approved Wiki Apply MVP

This document describes the fifth MVP: applying a reviewed and evaluated ingest
bundle to the maintained wiki only after explicit approval.

## Goal

Return the implementation to the core purpose of `nemo-knows`:

```text
raw source -> reviewed bundle -> evaluated apply plan -> approved wiki edits
```

The previous MVPs created safety rails around model output. This MVP validates
that those rails can support a real wiki-maintenance step.

## Command

```sh
go run ./cmd/nemo \
  -apply-approved drafts/actual-use-llm-wiki \
  -approve
```

Without `-approve`, the command must fail without writing files.

If the same bundle already appears in `wiki/log.md` as an `ingest` entry, the
command must also fail by default. Re-running an apply requires explicit force:

```sh
go run ./cmd/nemo \
  -apply-approved drafts/actual-use-llm-wiki \
  -approve \
  -force-apply
```

## Preconditions

The bundle must contain:

```text
source.md
ingest-plan.md
apply-plan.md
```

The bundle may also contain reviewed candidate page drafts:

```text
candidates/wiki/concepts/<slug>.md
candidates/wiki/topics/<slug>.md
```

The path below `candidates/` mirrors the final wiki path. For example:

```text
drafts/actual-use-llm-wiki/candidates/wiki/concepts/llm-maintenance-pattern.md
  -> wiki/concepts/llm-maintenance-pattern.md
```

The bundle must pass the deterministic eval harness:

```text
overall = pass
```

## Apply Policy

The command is intentionally conservative:

- It never writes to `raw/`.
- It only writes to `wiki/sources/`, `wiki/concepts/`, `wiki/topics/`,
  `wiki/index.md`, and `wiki/log.md`.
- It applies source-summary content only when a safe target can be determined.
- It indexes newly created source pages under `## Sources`.
- It applies concept/topic candidates only when a matching reviewed draft exists
  under `candidates/`.
- It does not create concept or topic pages from `ingest-plan.md` alone.
- It skips duplicate candidate creation and records the reason in
  `apply-report.md`.
- It refuses to apply the same bundle twice unless `-force-apply` is present.

Candidate drafts must:

- Have YAML frontmatter.
- Use `kind: concept` for `wiki/concepts/` targets.
- Use `kind: topic` for `wiki/topics/` targets.
- Include a `sources` frontmatter field.

For this MVP, a source candidate marked as a possible duplicate of an existing
source page is redirected to the existing source page. This avoids creating
parallel pages such as both `llm-wiki.md` and `llm-wiki-pattern.md`.

## Outputs

The command writes:

```text
wiki/sources/<slug>.md        # only when source apply is safe
wiki/concepts/<slug>.md       # only with reviewed candidate draft
wiki/topics/<slug>.md         # only with reviewed candidate draft
wiki/index.md                 # only when a new accepted page needs indexing
wiki/log.md                   # append-only apply record
drafts/<bundle>/apply-report.md
```

`apply-report.md` records:

- which files were written,
- which candidates were skipped,
- why they were skipped,
- whether manual follow-up is required.

## Acceptance Criteria

This MVP is successful when:

- `go test ./...` passes.
- `nemo -apply-approved` refuses to run without `-approve`.
- the command refuses bundles whose eval score is not `overall: pass`.
- the command does not create possible duplicate pages.
- the command can update a safe source-summary target.
- the command adds newly created source pages to `wiki/index.md`.
- the command can create concept/topic pages from reviewed candidate drafts.
- the command skips concept/topic candidates that lack matching reviewed drafts.
- the command rejects candidate drafts whose `kind` does not match the target
  directory.
- the command appends to `wiki/log.md`.
- the command writes `apply-report.md`.
- the command rejects repeated applies by default.

## Non-Goals

This MVP does not:

- generate concept/topic page content,
- resolve semantic duplicates with an LLM judge,
- apply concept/topic candidates without reviewed draft files,
- commit changes.

## Next Step

After this MVP, add a reviewed concept/topic drafting stage so approved applies
can safely create or update more than source summary pages.
