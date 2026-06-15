---
title: Git Hooks Reference Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/010-githooks-reference.md
confidence: medium
---

# Git Hooks Reference Summary

## What It Is
A documentation resource for Git hooks, describing executable programs placed in a repository's hooks directory to trigger actions at specific points during Git operations. These scripts are located by default in `$GIT_DIR/hooks` or a path defined by `core.hooksPath`.

## Summary
Git hooks allow automation of tasks such as validating commits, enforcing branch policies, and handling push/receive events. The documentation details the available hooks, their invocation contexts, parameter formats, and exit status behaviors. It covers both client-side hooks (e.g., `pre-commit`) and server-side hooks (e.g., `pre-receive`, `update`).

## Key Claims
- Hooks are ignored if they do not have the executable bit set.
- The working directory changes to `$GIT_DIR` in bare repositories or the repository root in non-bare ones before a hook runs.
- Environment variables like `GIT_DIR` and `GIT_WORK_TREE` are exported; hooks must unset them if invoking Git on foreign repositories.
- Exiting with a non-zero status generally aborts the triggering command, except for notification-only hooks like `post-commit`.
- Specific hooks exist for P4 integration (`p4-changelist`, `p4-pre-submit`) and filesystem monitoring (`fsmonitor-watchman`).

## Suggested Links
https://git-scm.com/docs/githooks
