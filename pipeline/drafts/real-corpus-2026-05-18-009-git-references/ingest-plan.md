---
kind: topic
sources: [raw/web/corpus-2026-05-18/009-git-references.md]
status: draft
---

# Ingest Plan

## Source Summary
- Extracts technical definitions and file structures for Git references (heads, tags, remotes) from the official Pro Git book.
- Explains the `HEAD` file behavior, including symbolic links vs. detached states and SHA-1 storage.
- Differentiates between lightweight tags, annotated tags, and remote reference tracking mechanisms.

## Candidate Wiki Pages
- wiki/sources/git-references.md — Documents the ingestion metadata and source reliability for Git internals content.
- wiki/concepts/git-refs-structure.md — Captures core concepts of `.git/refs` hierarchy (heads, tags, remotes).
- wiki/topics/detached-head-state.md — Explains the specific state where HEAD points directly to a commit SHA.

## Suggested Links
- https://git-scm.com/book/en/v2/Git-Internals-Git-References

## Review Checklist
- [ ] Verify `.git/refs` directory structure examples match current Git versions.
- [ ] Confirm distinction between lightweight and annotated tag creation commands is accurate.
