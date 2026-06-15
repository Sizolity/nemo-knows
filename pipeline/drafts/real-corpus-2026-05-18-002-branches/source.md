---
title: Branches in a Nutshell Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/002-branches-in-a-nutshell.md
confidence: medium
---

# Branches in a Nutshell Summary

## What It Is
This document is a summary of Chapter 3.1 ("Branches in a Nutshell") from the Git Book (2nd Edition), hosted on git-scm.com. It explains Git's lightweight branching model, contrasting it with traditional Version Control Systems (VCS) that rely on expensive file copying. The text details how Git stores data as snapshots rather than change sets, defines branches as movable pointers to commits, and describes the mechanics of creating, switching, and managing divergent histories using commands like `git branch`, `git checkout`, and `git switch`.

## Summary
Git's branching model is described as a "killer feature" due to its lightness and speed. Unlike older VCS tools that require copying entire source code directories to create branches (a slow process for large projects), Git branches are simply pointers to specific commits. Creating a branch involves writing a 40-character SHA-1 checksum to a file, making the operation nearly instantaneous. This allows developers to frequently branch and merge without performance penalties. The document explains that `git init` creates a default branch named `master` (which is not special), and introduces the `HEAD` pointer which tracks the current branch. It covers how switching branches moves the `HEAD` pointer and reverts the working directory to the snapshot of the target branch, enabling isolated development paths that can be merged later.

## Key Claims
- **Lightweight Branching**: Git branches are pointers to commits, not copies of files, making creation and destruction instantaneous compared to other VCS tools.
- **Snapshot Storage**: Git stores data as a series of snapshots (blobs, trees, commits) rather than change sets.
- **Branch Definition**: A branch is a lightweight movable pointer to one of the commit objects; `master` is simply the default name given by `git init`.
- **HEAD Pointer**: `HEAD` is a special pointer in Git that references the current local branch (distinct from other VCS concepts).
- **Working Directory Changes**: Switching branches automatically updates the working directory files to match the snapshot of the target branch; if changes exist locally that conflict with the target, the switch is blocked.
- **Divergent History**: By creating a new branch and switching back to `master`, developers can work on isolated features or fixes that diverge from the main line without affecting it immediately.
- **Modern Commands**: From Git version 2.23 onwards, `git switch` can be used instead of `git checkout` to switch branches (`git switch <branch>`) or create and switch to a new branch simultaneously (`git switch -c <branch>`).

## Suggested Links
- https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell
