---
title: Branch Fast Forward Vs Three Way
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/003-basic-branching-and-merging.md
confidence: medium
---

# Branch Fast Forward Vs Three Way

In Git workflows, merging a branch involves two distinct behaviors depending on the history of the branches being combined. The primary distinction lies in whether the merge creates a new commit or simply updates the current branch pointer.

## Fast-Forward Merge

A **fast-forward** merge occurs when the branch to be merged is directly ahead of the current HEAD. In this scenario, Git does not create a new merge commit. Instead, it simply moves the pointer of the current branch forward to point to the tip of the other branch. This preserves a linear history where the changes appear as if they were committed sequentially.

## Three-Way Merge

A **three-way** merge is performed when the histories of the two branches have diverged. Git analyzes three points in time:
1. The current state of the branch being merged into.
2. The tip of the branch to be merged.
3. The common ancestor (last shared commit) before the divergence.

Using these three snapshots, Git attempts to automatically integrate changes. If conflicts exist where both branches modified the same lines differently, Git pauses the merge and marks files as "unmerged" with conflict markers (`<<<<<<<`, `=======`, `>>>>>>>`). Users must then manually resolve these conflicts by editing the files, staging them with `git add`, and finalizing the process with `git commit`. This results in a new **merge commit** that records the integration point.

## Merge Tools

To assist with the resolution of conflicts during a three-way merge, users can employ graphical tools such as `opendiff`, `kdiff3`, or `meld` via the command `git mergetool`. These tools provide a visual interface to compare versions and select the appropriate changes before completing the merge.
