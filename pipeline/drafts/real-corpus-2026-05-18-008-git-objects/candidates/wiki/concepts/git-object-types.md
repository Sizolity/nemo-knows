---
title: Git Object Types
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/008-git-objects.md
confidence: medium
---

# Git Object Types

## Overview
Git's low-level storage model treats repository data as addressable objects. Content is written into the object database, assigned an identifier derived from that content, and later retrieved by that identifier. This object database is the foundation underneath Git's higher-level version-control commands.

The core object vocabulary in this source is blobs, trees, and commits. Together they model file bytes, directory shape, and project history.

## Object Types
- **Blobs** store file content without carrying the filename as part of the object itself.
- **Trees** describe directory entries by connecting names, modes, and object identifiers.
- **Commits** point at a tree and add history metadata such as parent commits, author information, and the message.

## Storage Mechanism
Loose objects are stored below `.git/objects` using a path derived from the object identifier: the leading characters select the directory and the remaining characters name the file. Before storage, Git records the object type and size with the content and compresses the result.

## Key Properties
- **Content addressability**: object identity follows from the bytes Git stores, not from the working-tree path that led to those bytes.
- **Layered structure**: blobs, trees, and commits build on each other rather than all storing complete repository state in the same format.
- **Inspectability**: plumbing commands expose the object database directly, which makes the storage model testable from the command line.

## Manipulation
The chapter uses commands such as `git hash-object`, `git cat-file`, `git write-tree`, and `git commit-tree` to show how the object database can be created and inspected without porcelain commands.
