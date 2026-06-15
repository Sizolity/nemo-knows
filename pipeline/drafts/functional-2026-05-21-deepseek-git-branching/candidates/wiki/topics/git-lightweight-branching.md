---
title: Git Lightweight Branching
kind: topic
sources:
  - source.md
  - raw/web/git-branching.md
confidence: medium
---

# Git Lightweight Branching

Git treats branches as movable pointers to commits, not as heavy copies of a directory tree. Each branch is stored as a small file that contains the checksum of the commit it points to. Creating a new branch therefore adds almost no overhead—unlike older version control systems that duplicated entire file structures.

The [[git-head]] pointer tracks which branch is currently active. Running `git branch <name>` creates a new pointer to the present commit without switching context; `git checkout` (or `git switch`) moves `HEAD` to that branch and updates the working tree to reflect the target snapshot. Because the underlying pointer is so cheap, Git encourages workflows where isolated work happens on short-lived branches that are merged back frequently.

You can inspect branch topology and commit relationships with [[log]] commands such as `git log --oneline --decorate --graph --all`. This lightweight model means switching branches and experimenting with new ideas carries almost no cost, making frequent branching a natural part of Git’s design.
