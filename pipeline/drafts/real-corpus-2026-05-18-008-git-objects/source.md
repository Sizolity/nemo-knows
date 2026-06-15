---
title: Git Objects
kind: source
sources:
  - raw/web/corpus-2026-05-18/008-git-objects.md
confidence: medium
---

# Git Objects

## What It Is
Git is a content-addressable filesystem that functions as a simple key-value data store. At its core, it allows users to insert any kind of content into a repository, for which Git returns a unique key used later to retrieve that specific content. These objects are stored in the `.git/objects` directory and form the basis of version control history.

## Summary
Git stores all content as three primary types of objects: blobs, trees, and commits.
- **Blobs** store raw file contents (inodes).
- **Trees** correspond to UNIX directory entries, grouping files and subdirectories together with metadata like mode and filename.
- **Commits** represent snapshots of the project history, linking to a specific tree, author information, timestamps, and commit messages.

Objects are stored as separate files in the `.git/objects` directory, named using a SHA-1 checksum hash (40 characters total). The first two characters of the hash form the subdirectory name, while the remaining 38 characters form the filename. Every object includes a header identifying its type and size, followed by the content, which is compressed using zlib before being written to disk.

## Key Claims
- **Content Addressability**: Git does not rely on filenames for storage identity; instead, it uses SHA-1 hashes of the content plus a header to uniquely identify objects.
- **Object Structure**:
  - A blob header starts with `blob` followed by the size in bytes and a null byte.
  - Commits and trees have specific internal formats distinct from blobs.
- **Storage Mechanism**: Content is concatenated with its header, compressed with zlib, and then written to the filesystem under a path derived from the SHA-1 hash.
- **Plumbing Commands**: Low-level commands like `git hash-object`, `git cat-file`, `git write-tree`, and `git commit-tree` are used to manipulate these objects directly without using high-level frontend commands.

## Suggested Links
https://git-scm.com/book/en/v2/Git-Internals-Git-Objects
