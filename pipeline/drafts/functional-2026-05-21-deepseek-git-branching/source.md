---
title: The Anatomy of Git Branching
kind: source
sources:
  - raw/web/git-branching.md
confidence: medium
---

# The Anatomy of Git Branching

## What It Is
A condensed reference from the Pro Git book explaining how Git stores data and why branches are lightweight.

## Summary
Git stores content as snapshots. A commit object holds metadata, a pointer to the root tree snapshot, and pointers to parent commits. Files are stored as blobs, directories as trees, and commits form the history graph. A Git branch is just a movable pointer to a commit; the default branch name is often `master`, but it is not special. `HEAD` identifies the current local branch. Creating a branch with `git branch testing` adds a new pointer to the current commit without switching to it. Switching branches with `git checkout testing` moves `HEAD` to that branch and updates the working directory to match the target commit’s snapshot. Subsequent commits move only the current branch forward, leading to divergent histories. Graph relationships can be inspected with `git log --oneline --decorate --graph --all`.

## Key Claims
- Git branches are cheap: a branch is merely a tiny pointer file containing a commit checksum, unlike older version control systems that copied many files into a new directory.
- Because branching is cheap, Git encourages frequent branch and merge workflows.

## Suggested Links
- [Pro Git: Branches in a Nutshell](https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell)
