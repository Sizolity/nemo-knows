---
title: Git Merge Conflict Resolution
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/003-basic-branching-and-merging.md
confidence: medium
---

# Git Merge Conflict Resolution

A merge conflict occurs when two branches modify the same section of a file differently. When Git encounters such a situation, it pauses the merge process and marks the affected files as "unmerged." These files contain specific conflict markers to indicate the areas of disagreement: `<<<<<<<`, `=======`, and `>>>>>>>`.

## Resolution Process

To resolve a conflict, you must manually edit the conflicted files. The standard approach involves choosing one version of the code or merging both versions logically within the conflict markers. Once the file is edited to reflect the intended state, stage the resolved changes using `git add <filename>`. Finally, complete the merge by running `git commit`.

## Tools and Strategies

While manual editing is common, you can use graphical tools to assist in resolving conflicts more effectively. The command `git mergetool` can launch external applications such as `opendiff`, `kdiff3`, or `meld` to visualize differences side-by-side.

## Merge Contexts

Understanding the type of merge helps in managing expectations:
- **Fast-forward merges**: Occur when merging a branch that is directly ahead of the current HEAD, moving the pointer forward without creating a new commit or conflicts.
- **Three-way merges**: Triggered when histories have diverged; Git uses the two branch tips and their common ancestor to create a new snapshot. If the changes overlap during this process, a conflict arises.

After resolving conflicts and committing the merge, it is good practice to clean up by deleting merged branches using `git branch -d <branch-name>`.
