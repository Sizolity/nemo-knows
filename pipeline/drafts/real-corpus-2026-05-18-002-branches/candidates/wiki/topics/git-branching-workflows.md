---
title: Git Branching Workflows
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/002-branches-in-a-nutshell.md
confidence: medium
---

# Git Branching Workflows

Git's branching model is often described as a "killer feature" due to its lightness and speed. Unlike traditional Version Control Systems (VCS) that rely on expensive file copying, Git stores data as snapshots rather than change sets. This architectural difference allows developers to frequently branch and merge without performance penalties.

## Lightweight Branching

In Git, branches are not copies of files but lightweight movable pointers to commits. Creating a branch involves writing a 40-character SHA-1 checksum to a file, making the operation nearly instantaneous. The default branch created by `git init` is named `master`, which is not special in terms of functionality.

## Branch Definition and HEAD Pointer

A branch is defined as a pointer to one of the commit objects. The `HEAD` pointer is a special reference that tracks the current local branch, distinct from concepts found in older VCS tools. When switching branches, the `HEAD` pointer moves, and the working directory reverts to the snapshot of the target branch.

## Managing Divergent Histories

Developers can create isolated development paths by creating a new branch and switching back to `master`. This allows work on features or fixes that diverge from the main line without immediately affecting it. However, if local changes exist that conflict with the target branch, the switch is blocked to prevent data loss.

## Modern Commands

From Git version 2.23 onwards, the command `git switch` can be used instead of `git checkout`. This command allows users to switch branches (`git switch <branch>`) or create and switch to a new branch simultaneously (`git switch -c <branch>`).
