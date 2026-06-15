---
kind: topic
sources: [raw/web/corpus-2026-05-18/003-basic-branching-and-merging.md]
status: draft
---

# Ingest Plan

## Source Summary
- Provides a practical, step-by-step walkthrough of Git branching workflows using real-world scenarios (user stories vs. hotfixes).
- Covers fundamental concepts including branch creation, switching, fast-forward merges, three-way merges, and conflict resolution.
- Includes detailed explanations of merge commit mechanics, conflict markers, and manual resolution strategies via `git add` and tools.

## Candidate Wiki Pages
- wiki/sources/git-basic-branching-and-merging.md — Stores the raw ingestion metadata and source content for future reference.
- wiki/concepts/branch-fast-forward-vs-three-way.md — Explains the difference between fast-forward and recursive merge strategies with Git examples.
- wiki/topics/git-merge-conflict-resolution.md — Focuses on resolving conflicts using markers, manual editing, and graphical tools like `git mergetool`.

## Suggested Links
- https://git-scm.com/book/en/v2/Git-Branching-Basic-Branching-and-Merging

## Review Checklist
- [ ] Verify all command examples are syntactically correct and up-to-date with modern Git versions.
- [ ] Ensure conflict resolution steps clearly distinguish between staged and unstaged changes.
- [ ] Confirm that the distinction between fast-forward and non-fast-forward merges is visually illustrated if possible.
