---
title: Git Branches Pointers
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/002-branches-in-a-nutshell.md
confidence: medium
---

# Git Branches Pointers

Git's branching model is described as a "killer feature" due to its lightness and speed. Unlike older Version Control Systems (VCS) that require copying entire source code directories, Git branches are simply **pointers** to specific commits. Creating a branch involves writing a 40-character SHA-1 checksum to a file, making the operation nearly instantaneous.

## Definition

A branch is defined as a lightweight movable pointer to one of the commit objects. The default branch created by `git init` is named `master`, which is not special in itself.

The concept relies on two primary pointers:
*   **Branches**: Lightweight references that track a specific commit history.
*   **HEAD**: A special pointer in Git that references the current local branch, distinct from other VCS concepts.

## Mechanics

Git stores data as a series of snapshots (blobs, trees, commits) rather than change sets. When switching branches:
1.  The `HEAD` pointer moves to reference the target branch.
2.  The working directory reverts to the snapshot of that target branch.
3.  If local changes conflict with the target branch state, the switch is blocked.

This mechanism enables isolated development paths that can be merged later without performance penalties. Modern commands like `git switch` (available from Git version 2.23 onwards) allow for switching branches or creating and switching to a new branch simultaneously.
