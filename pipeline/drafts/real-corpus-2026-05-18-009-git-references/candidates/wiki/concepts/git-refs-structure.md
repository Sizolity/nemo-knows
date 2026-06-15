---
title: Git Refs Structure
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/009-git-references.md
confidence: medium
---

# Git Refs Structure

Git references, or **refs**, are simple names that act as pointers to store SHA-1 values of Git objects. They allow users to refer to commits without needing to remember raw hashes. These references are stored within the `.git/refs` directory structure and primarily include branches, tags, and remote tracking references. The `HEAD` file functions as a symbolic reference pointing to the current branch or, in specific cases, directly to a commit SHA-1.

## Storage Locations

References are organized into specific directories within the repository:

- **Branches**: Stored in `.git/refs/heads`.
- **Tags**: Stored in `.git/refs/tags`.
- **Remote Tracking**: Stored in `.git/refs/remotes`.

Direct file editing of these references is generally discouraged. Instead, it is recommended to use commands like `git update-ref` and `git symbolic-ref` to modify references safely.

## The HEAD File

The `HEAD` file acts as a pointer to the current state of the repository:

- **Normal State**: Typically contains a relative reference pointing to a branch (e.g., `ref: refs/heads/master`).
- **Detached Head State**: Contains a raw SHA-1 value, indicating that `HEAD` points directly to a specific commit rather than a branch name. This state is known as the [[detached-head-state]].

## Tags

Tags in Git can be of two types:

- **Lightweight Tags**: These are direct references that never move and point straight to a commit.
- **Annotated Tags**: These create an intermediate tag object that points to the commit, allowing for metadata like authorship and message storage.

## Remote References

Remote tracking references differ from local branches in their behavior and purpose:

- They serve as read-only bookmarks representing the last known state of remote branches.
- They cannot be updated via standard commit commands.
- Their structure is maintained to reflect the state fetched from a remote server.
