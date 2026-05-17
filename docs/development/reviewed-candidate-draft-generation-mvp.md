# Reviewed Candidate Draft Generation MVP

This document describes the seventh MVP: generating reviewed candidate page
drafts from a reviewed `apply-plan.md` without writing to `wiki/`.

## Goal

Fill the gap between reviewed candidate paths and approved wiki apply:

```text
apply-plan.md -> candidates/wiki/concepts|topics/*.md -> approved apply
```

MVP-6 proved that `-apply-approved` can safely consume separate candidate draft
files. MVP-7 creates those files as draft artifacts so the model still cannot
turn one-line candidate descriptions directly into permanent wiki pages.

## Command

```sh
go run ./cmd/nemo \
  -generate-candidates drafts/actual-use-llm-wiki \
  -profile stable
```

The command reads `drafts/<bundle>/apply-plan.md` and generates drafts for
candidate paths under:

```text
wiki/concepts/
wiki/topics/
```

It ignores source candidates because source pages already come from `source.md`.

## Inputs

The bundle must contain:

```text
source.md
apply-plan.md
```

The command uses `source.md` as the source material for each candidate draft.
The reviewed apply plan determines which candidate paths are eligible.

## Outputs

Generated candidate drafts mirror their eventual wiki paths below
`candidates/`:

```text
drafts/<bundle>/candidates/wiki/concepts/<slug>.md
drafts/<bundle>/candidates/wiki/topics/<slug>.md
```

Each generated draft also keeps the raw model output next to it:

```text
drafts/<bundle>/candidates/wiki/concepts/<slug>.raw.txt
```

## Policy

- The command never writes to `wiki/`.
- The command never writes to `raw/`.
- Only `wiki/concepts/` and `wiki/topics/` candidate paths are generated.
- Existing candidate draft files may be overwritten because they are draft
  artifacts, not accepted wiki pages.
- Candidate prompts receive an Allowed Links list built from existing `wiki/`
  pages plus concept/topic candidates in the reviewed apply plan.
- Generated candidate drafts are normalized after cleaning: wikilinks outside
  the Allowed Links list are downgraded to plain text before the draft is saved.
- Candidate normalization also enforces a deterministic `# <title>` heading.
- Candidate generation does not force wikilinks. If the source does not support
  a strong cross-reference, the candidate should use plain text.
- Approved apply remains a separate explicit step.

## Acceptance Criteria

This MVP is successful when:

- `go test ./...` passes.
- source candidates in `apply-plan.md` are ignored.
- concept/topic candidates generate mirrored draft files under `candidates/`.
- generated drafts are cleaned Markdown, with raw model output preserved.
- generated drafts do not preserve model-invented wikilinks to pages outside
  the current wiki or reviewed candidate set.
- generated drafts satisfy the candidate eval title and link-safety contract even
  when the model uses a linked H1 or omits wikilinks.
- `-apply-approved` remains the only command that writes candidate pages to
  `wiki/`.
