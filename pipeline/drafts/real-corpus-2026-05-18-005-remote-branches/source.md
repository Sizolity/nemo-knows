---
title: Remote Branches
kind: source
sources:
  - raw/web/corpus-2026-05-18/005-remote-branches.md
confidence: medium
---

# Remote Branches

## What It Is
Remote branches are references (pointers) to the state of branches, tags, and other items in remote repositories. Git manages these via **remote-tracking branches**, which act as bookmarks pointing to the last known state of a branch on a remote server. These local references are automatically updated by Git during network communication and cannot be manually moved by the user.

## Summary
The document explains how Git handles synchronization between local and remote repositories using remote-tracking branches (e.g., `origin/master`). It details how to view lists of remote references, fetch updates from specific remotes, and push local branches to share them with others. The text covers creating tracking branches for automatic merge behavior during pulls, managing upstream relationships, and deleting remote branches via the `--delete` option on `git push`.

## Key Claims
- **Remote-tracking branches** are local references that Git moves automatically to reflect the state of remote branches; they serve as bookmarks for remote repository states.
- **Remote branch names** follow the format `<remote>/<branch>` (e.g., `origin/master`). The remote name (like "origin") is not special and can be customized during cloning.
- **Synchronization**: Running `git fetch <remote>` updates local remote-tracking branches without modifying the working directory, while `git pull` performs a fetch followed by a merge.
- **Pushing**: Local branches are not automatically synchronized to remotes; users must explicitly push desired branches (e.g., `git push origin branchname`).
- **Tracking Branches**: Checking out a local branch from a remote-tracking branch creates a "tracking branch" with an associated upstream, enabling automatic fetch targets for `git pull`.
- **Deletion**: Remote branches can be deleted using `git push <remote> --delete <branch>`, though the server often retains data temporarily for recovery.

## Suggested Links
- https://git-scm.com/book/en/v2/Git-Branching-Remote-Branches
