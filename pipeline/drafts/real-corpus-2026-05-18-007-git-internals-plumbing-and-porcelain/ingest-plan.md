---
kind: topic
sources: [raw/web/corpus-2026-05-18/007-git-internals-plumbing-and-porcelain.md]
status: draft
---

# Ingest Plan

## Source Summary
- The source defines Git as a content-addressable filesystem with a user interface layer, distinguishing between low-level "plumbing" commands and high-level "porcelain" commands.
- It details the structure of the `.git` directory, specifically focusing on `objects`, `refs`, `HEAD`, and the `index` as the core components storing content, pointers, current branch state, and staging area data respectively.
- The document serves as a deep implementation reference intended for scripts and tooling rather than manual command-line usage by beginners.

## Candidate Wiki Pages
- wiki/sources/git-internals-plumbing-porcelain.md — Primary source documentation for the plumbing/porcelain architecture.
- wiki/concepts/content-addressable-filesystem.md — Explains the foundational storage model described in the text.
- wiki/topics/.git-directory-structure.md — Documents the specific subdirectories (objects, refs, hooks, etc.) mentioned in the `.git` initialization section.

## Suggested Links
- https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain

## Review Checklist
- [ ] Verify that candidate pages do not reference forbidden schema or index files.
- [ ] Ensure all candidate page paths are immediate children of `wiki/sources/`, `wiki/concepts/`, or `wiki/topics/`.
