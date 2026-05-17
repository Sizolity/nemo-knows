---
title: Git Branches in a Nutshell
kind: source
sources:
  - raw/web/git-branching.md
confidence: medium
---

## What It Is

A summary of the Pro Git chapter on Git branching, which explains that Git branching is lightweight because Git stores data as snapshots and branches are movable pointers to commits.

## Summary

Git branching is considered lightweight because Git stores data as snapshots and branches act as movable pointers to commits. A commit object includes metadata, a pointer to the root tree snapshot, and pointers to parent commits. Files are stored as blobs, directories as trees, and commits connect snapshots into history. A Git branch is a pointer to one commit; the default branch name is often `master`, but it is not special. When new commits are made, the current branch pointer moves forward. Git uses `HEAD` to identify the current local branch. Creating a branch with `git branch testing` creates a second pointer to the current commit without switching to that branch. Switching branches with `git checkout testing` moves `HEAD` to point to the selected branch. The chapter emphasizes that switching branches changes the working tree; if the selected branch points to an older commit, the working directory is updated to match that snapshot. After switching, new commits move the selected branch forward, which can create divergent history between branches. Git can show these relationships with commands such as `git log --oneline --decorate --graph --all`. The key claim is that Git branches are cheap because a branch is essentially a small pointer file containing a commit checksum. This contrasts with older version control systems where branching often meant copying many files into a new directory. Git's cheap branches encourage frequent branch and merge workflows.

## Key Claims

- Git branching is lightweight because Git stores data as snapshots and branches are movable pointers to commits.
- A branch is essentially a small pointer file containing a commit checksum, making it cheap compared to older version control systems that required copying many files.
- Git's cheap branches encourage frequent branch and merge workflows.

## Suggested Links

https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell
