---
title: Git Rebase Mechanics
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/004-rebasing.md
confidence: medium
---

# Git Rebase Mechanics

Git rebasing is a history-rewriting integration operation. It takes commits that were originally based on one commit and recreates equivalent commits on top of a different base. The resulting branch can contain the same file state that a merge would produce, but its graph records the work as if it had been made after the new base.

Mechanically, a rebase identifies the commits that are unique to the current branch, temporarily represents their changes as patches, moves the branch to the target base, and applies those changes again in order. This differs from merge-based integration, which preserves both lines of development and adds a merge commit when needed.

## Key Characteristics

### Linear History
Rebasing can make a project history easier to read by replacing a diverged branch shape with a straight sequence of commits.

### Snapshot Equivalence
A successful rebase may end at the same project snapshot as an equivalent merge. The important difference is the ancestry recorded in the commit graph.

### Public Commit Rule
The source warns against rebasing commits that others may already have based work on. Once a commit is shared, replacing it with a rewritten commit changes the identity that collaborators are using.

### Patch-ID Resolution
Git can compare the content of changes as well as commit object IDs. That lets rebase-related workflows recognize that two different commits may represent the same patch and avoid replaying duplicate work.

### Philosophical Divide
The chapter frames merge and rebase as different attitudes toward history: one preserves the visible record of collaboration, while the other edits local history into a cleaner story before publication.

## References

- [[git-history-management-best-practices]]
- [[log]]
