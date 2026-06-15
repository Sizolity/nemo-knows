---
title: Git Hooks
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/010-githooks-reference.md
confidence: medium
---

# Git Hooks

Git hooks are executable programs placed in a repository's hooks directory that trigger actions at specific points during Git operations. By default, these scripts are located in `$GIT_DIR/hooks`, though the path can be defined by `core.hooksPath`.

## Functionality

Hooks allow for the automation of various tasks within the version control workflow, including:
- Validating commits
- Enforcing branch policies
- Handling push and receive events

The available hooks cover both client-side operations (e.g., `pre-commit`) and server-side operations (e.g., `pre-receive`, `update`). Documentation details their invocation contexts, parameter formats, and exit status behaviors.

## Behavior and Rules

- **Execution Permission**: Hooks are ignored if they do not have the executable bit set.
- **Working Directory**: The working directory changes to `$GIT_DIR` in bare repositories or the repository root in non-bare ones before a hook runs.
- **Environment Variables**: Environment variables like `GIT_DIR` and `GIT_WORK_TREE` are exported. Hooks must unset these if invoking Git on foreign repositories.
- **Exit Status**: Exiting with a non-zero status generally aborts the triggering command, except for notification-only hooks like `post-commit`.

## Specific Integrations

The system supports specific hooks for:
- P4 integration (`p4-changelist`, `p4-pre-submit`)
- Filesystem monitoring (`fsmonitor-watchman`)
