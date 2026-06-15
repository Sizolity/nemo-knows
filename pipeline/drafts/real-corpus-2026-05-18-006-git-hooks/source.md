---
title: Git Hooks
kind: source
sources:
  - raw/web/corpus-2026-05-18/006-git-hooks.md
confidence: medium
---

# Git Hooks Summary

## What It Is
Git hooks are custom scripts that fire off automatically when specific important actions occur within a repository. They function as event-driven automation tools available in two groups: **client-side** and **server-side**. Client-side hooks trigger on local operations like committing and merging, while server-side hooks run on network operations such as receiving pushed commits.

## Summary
The Git documentation describes hooks as a mechanism to enforce policies or automate tasks without modifying the core workflow. Hooks are stored in the `.git/hooks` subdirectory. When initializing a new repository with `git init`, Git populates this directory with example scripts ending in `.sample`. To activate a hook, the file must be named appropriately (without an extension) and marked as executable. While examples are shell scripts, any properly named executable script works.

**Important Note:** Client-side hooks are **not** copied when cloning a repository. If the intent is to enforce a policy that applies to all clones, it should be implemented on the server side.

## Key Claims
- **Client-Side Hooks:**
  - **Committing-Workflow:**
    - `pre-commit`: Runs before the commit message is typed; used for inspection (linting, tests) or aborting the commit if non-zero exit code. Can be bypassed via `git commit --no-verify`.
    - `prepare-commit-msg`: Runs before the commit message editor opens but after the default message is created; useful for templated messages.
    - `commit-msg`: Takes the path to the temporary file containing the commit message; aborts the commit if the script exits non-zero.
    - `post-commit`: Runs after the entire commit process; typically used for notifications.
  - **Email Workflow (via `git am`):**
    - `applypatch-msg`: Validates or normalizes the proposed commit message from a patch.
    - `pre-applypatch`: Runs after the patch is applied but before the commit; allows inspection of the snapshot or running tests.
    - `post-applypatch`: Runs after the commit is made for notifications.
  - **Other Client Hooks:**
    - `pre-rebase`: Halts rebasing if it exits non-zero (e.g., to disallow rebasing pushed commits).
    - `post-rewrite`: Runs when commits are replaced (e.g., `git commit --amend`, `git rebase`).
    - `post-checkout`: Runs after a successful checkout; useful for setting up the working directory.
    - `post-merge`: Runs after a successful merge; useful for restoring data Git cannot track (permissions) or copying external files.
    - `pre-push`: Runs during push before object transfer; can validate ref updates and abort the push if necessary.
    - `pre-auto-gc`: Invoked just before automatic garbage collection.

- **Server-Side Hooks:**
  - Run before and after pushes to the server to enforce policies.
  - `pre-receive`: Runs once per push; receives list of references via stdin; can reject all refs if it exits non-zero. Used for access control or checking fast-forward status.
  - `update`: Similar to `pre-receive` but runs once per branch being updated. Takes reference name, old SHA-1, and new SHA-1 as arguments. Only rejects the specific reference if non-zero.
  - `post-receive`: Runs after the push process completes; cannot stop the push but can update other services (CI servers, ticket tracking) or notify users.

## Suggested Links
- https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks
