---
title: Git References Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/009-git-references.md
confidence: medium
---

# Git References

## What It Is
Git references (refs) are simple names or pointers that store SHA-1 values of Git objects, allowing users to refer to commits without remembering raw hashes. These references are stored in the `.git/refs` directory structure and include branches, tags, and remote tracking references. The HEAD file acts as a symbolic reference pointing to the current branch or, in detached states, directly to a commit SHA-1.

## Summary
This document explains the internal mechanics of Git references, covering how they store commit history pointers, the distinction between lightweight and annotated tags, and the role of remote references. It details commands like `git update-ref` for managing refs safely and describes the HEAD file's behavior in normal versus detached states. The text also outlines the structure of tag objects and how remote references differ from local branches by being read-only bookmarks to server states.

## Key Claims
- **Reference Storage**: Simple names (refs) store SHA-1 values, stored in `.git/refs/heads`, `.git/refs/tags`, and `.git/refs/remotes`.
- **HEAD Behavior**: The HEAD file normally points to a branch (`ref: refs/heads/master`) but can contain a raw SHA-1 value when in a "detached HEAD" state.
- **Tag Types**: Lightweight tags are direct references that never move, while annotated tags create an intermediate tag object pointing to the commit.
- **Remote References**: Stored in `refs/remotes`, these are read-only bookmarks representing the last known state of remote branches and cannot be updated via commit commands.
- **Safe Updates**: Direct file editing of refs is discouraged; `git update-ref` and `git symbolic-ref` are the recommended commands for modifying references.

## Suggested Links
https://git-scm.com/book/en/v2/Git-Internals-Git-References
