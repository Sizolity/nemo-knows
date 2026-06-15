---
title: Remote Tracking Branches
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/005-remote-branches.md
confidence: medium
---

# Remote Tracking Branches

Remote tracking branches are local references that Git maintains to track the state of branches, tags, and other items in remote repositories. They act as bookmarks pointing to the last known state of a branch on a remote server. These local references are automatically updated by Git during network communication and cannot be manually moved by the user.

## Naming Convention
Remote tracking branch names follow the format `<remote>/<branch>`. Common examples include `origin/master` or `upstream/develop`. The remote name (like "origin") is not special and can be customized during cloning.

## Synchronization
Running `git fetch <remote>` updates local remote-tracking branches without modifying the working directory, while `git pull` performs a fetch followed by a merge. Local branches are not automatically synchronized to remotes; users must explicitly push desired branches (e.g., `git push origin branchname`).

## Creating Tracking Branches
Checking out a local branch from a remote tracking branch creates a "tracking branch" with an associated upstream, enabling automatic fetch targets for `git pull`. This setup facilitates the standard [[pull-push-workflow]].

## Deletion
Remote branches can be deleted using `git push <remote> --delete <branch>`, though the server often retains data temporarily for recovery.
