---
kind: source
sources:
  - web:https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell
confidence: medium
---

# Git Branches in a Nutshell

## What It Is

A Pro Git chapter explaining Git branches as lightweight movable pointers to
commits, how `HEAD` tracks the current branch, and why branching and switching
are cheap operations in Git.

## Summary

The chapter contrasts Git's snapshot-based object model with older version
control branching approaches. A Git commit points to a tree snapshot and parent
commits. A branch is a small pointer to a commit, and `HEAD` identifies the
current branch. Creating a branch creates another pointer; switching branches
moves `HEAD` and updates the working tree. This lightweight design encourages
frequent branching and merging.

## Key Claims

- Git stores data as snapshots rather than a series of file differences.
- A branch is a lightweight movable pointer to a commit.
- `HEAD` points to the current local branch.
- Switching branches changes the working tree to match the selected branch.
- Branches are cheap because they are represented by small pointer files.

## Suggested Links

- Git
- branch
- commit object
- HEAD
- working tree
