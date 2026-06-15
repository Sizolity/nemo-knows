---
title: Server Side Hooks
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/010-githooks-reference.md
confidence: medium
---

# Server Side Hooks

Server side hooks are a category of Git hooks that execute on the remote repository server rather than the client machine. They are designed to trigger actions at specific points during Git operations, particularly those involving push and receive events.

Common examples of server side hooks include `pre-receive` and `update`. These scripts are located by default in the `$GIT_DIR/hooks` directory or a path defined by `core.hooksPath`. To function correctly, these hook scripts must have the executable bit set; otherwise, they will be ignored by Git.

The working directory changes to `$GIT_DIR` in bare repositories or the repository root in non-bare ones before a server side hook runs. Environment variables like `GIT_DIR` and `GIT_WORK_TREE` are exported during execution. If a hook invokes Git for foreign repositories, it must unset these environment variables first.

Unlike notification-only hooks such as `post-commit`, exiting with a non-zero status in a server side hook generally aborts the triggering command.
