---
title: Content Addressable Filesystem
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/007-git-internals-plumbing-and-porcelain.md
confidence: medium
---

# Content Addressable Filesystem

Git is fundamentally a content-addressable filesystem with a version control system user interface built on top of it. In this architecture, data is stored based on its content hash rather than its filename or location in a directory tree.

## Core Components

The internal storage of a Git repository resides primarily within the `.git` directory. This directory contains almost everything Git stores and manipulates; copying this single directory allows for backing up or cloning a repository. The core components include:

- **objects/**: Stores all content for the database, organized by type and hash.
- **refs/**: Stores pointers into commit objects, representing branches, tags, and remotes.
- **HEAD**: Points to the currently checked-out branch.
- **index**: Stores staging area information.
- **config**, **hooks/**, and **info/** handle configuration, client/server scripts, and ignored patterns respectively.

## Plumbing and Porcelain

Git distinguishes between two layers of command interfaces:

- **Porcelain**: User-friendly commands designed for daily use.
- **Plumbing**: Low-level commands designed for scripting and chaining.

While early versions of Git emphasized the filesystem complexity, modern Git has refined its UI to be user-friendly while retaining powerful low-level capabilities. Understanding the content-addressable filesystem layer is fundamentally important to appreciating Git's power, though it can be confusing for beginners. Plumbing commands are generally not meant to be used manually on the command line but serve as building blocks for new tools and custom scripts.
