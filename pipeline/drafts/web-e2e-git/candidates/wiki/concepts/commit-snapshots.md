---
title: Commit Snapshots
kind: concept
sources:
  - source.md
  - raw/web/git-branching.md
confidence: medium
---

# Commit Snapshots

Git stores data as snapshots rather than just tracking file changes. A commit object includes metadata, a pointer to the root tree snapshot, and pointers to parent commits. Files are stored as blobs, directories as trees, and commits connect these snapshots into history.

## Branches as Pointers

A Git branch is essentially a small pointer file containing a commit checksum. Because branches act as movable pointers to commits rather than copying many files, they are considered lightweight. The default branch name is often `master`, though it holds no special status. When new commits are made, the current branch pointer moves forward. Git uses `HEAD` to identify the current local branch.

## Workflow Implications

Creating a branch with `git branch testing` creates a second pointer to the current commit without switching contexts. Switching branches with `git checkout testing` moves `HEAD` to point to the selected branch, updating the working directory to match that snapshot. After switching, new commits move the selected branch forward, which can create divergent history between branches. Git can show these relationships with commands such as `git log --oneline --decorate --graph --all`. This cheap branching model encourages frequent branch and merge workflows.
