---
title: Pull Push Workflow
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/005-remote-branches.md
confidence: medium
---

# Pull Push Workflow

The pull push workflow describes the process of synchronizing changes between local repositories and remote servers using Git. This synchronization is managed through **remote-tracking branches**, which act as bookmarks pointing to the last known state of a branch on a remote server. These local references are automatically updated by Git during network communication and cannot be manually moved by the user.

## Fetching (Pull)

To synchronize with a remote repository, users typically run `git fetch <remote>`. This command updates local remote-tracking branches without modifying the working directory. To integrate these changes into the current branch, `git pull` is executed; this performs a fetch followed by a merge.

When creating a tracking branch for automatic merge behavior during pulls, checking out a local branch from a remote-tracking branch establishes an associated upstream. This enables automatic fetch targets for future `git pull` operations. The names of these references follow the format `<remote>/<branch>`, such as `origin/master`. The remote name (like "origin") is not special and can be customized during cloning.

## Pushing

Local branches are not automatically synchronized to remotes; users must explicitly push desired branches using commands like `git push origin branchname`. To remove a branch from a remote server, the `--delete` option on `git push` is used (`git push <remote> --delete <branch>`), though the server often retains data temporarily for recovery.
