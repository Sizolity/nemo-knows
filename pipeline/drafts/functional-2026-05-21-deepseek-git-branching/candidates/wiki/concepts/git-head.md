---
title: Git Head
kind: concept
sources:
  - source.md
  - raw/web/git-branching.md
confidence: medium
---

# Git Head

In Git, `HEAD` is a symbolic reference that identifies the current local branch. It tells Git which snapshot is active in the working directory. When you check out a different branch, `HEAD` moves to point to that branch pointer. Because a branch is merely a lightweight pointer ([[git-branch]]), `HEAD` ultimately determines which commit’s content is visible.

Inspecting the commit graph with `git log --decorate` reveals `HEAD`’s position alongside branch pointers ([[log]]).
