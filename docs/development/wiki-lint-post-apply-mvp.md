# Wiki Lint / Post-Apply Validation MVP

This document describes the ninth MVP: deterministic validation of the
maintained `wiki/` after approved apply.

## Goal

Turn the `AGENTS.md` lint workflow into a repeatable report:

```text
wiki/ -> wiki-lint.json + wiki-lint.md
```

The command is read-only. It reports likely maintenance issues and waits for a
human decision before any wiki edits.

## Command

```sh
go run ./cmd/nemo \
  -lint-wiki \
  -out-dir evals/runs/wiki-lint-real
```

## Initial Checks

MVP-9 checks:

- missing YAML frontmatter on wiki pages,
- invalid or missing `kind`,
- missing `sources` on source/entity/concept/topic pages,
- missing or invalid `confidence` on source/entity/concept/topic pages,
- duplicated index entries,
- wikilinks that point to missing pages,
- orphan pages that are not linked from any other page or `wiki/index.md`,
- invalid `wiki/log.md` entry actions.

## Outputs

```text
evals/runs/<run-id>/wiki-lint.json
evals/runs/<run-id>/wiki-lint.md
```

The JSON output is intended for regression checks. The Markdown output is for
human review.

## Non-Goals

This MVP does not:

- edit `wiki/`,
- decide which orphan or stub should be kept,
- detect semantic contradictions,
- run an LLM judge.

## Acceptance Criteria

This MVP is successful when:

- `go test ./...` passes.
- lint reports invalid frontmatter and invalid log actions.
- lint reports duplicate index entries.
- lint reports missing wikilink targets.
- lint reports orphan pages.
- real `wiki/` can be linted without edits.
