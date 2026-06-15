---
title: Basic Branching and Merging Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/003-basic-branching-and-merging.md
confidence: medium
---

## What It Is

A procedural guide describing Git workflows for managing project changes through branching and merging. The document outlines a real-world scenario involving feature development (user stories), hotfixes, and the subsequent integration of work back into a main branch.

## Summary

The text explains how to create branches using `git checkout -b`, perform commits, switch between branches, and merge them. It details the concept of "fast-forward" merges where the branch pointer moves forward without creating a new commit, versus three-way merges that create a merge commit when histories diverge. The guide also covers resolving merge conflicts manually or using tools like `git mergetool`, staging resolved files with `git add`, and finalizing the merge with `git commit`. Finally, it instructs on cleaning up by deleting merged branches with `git branch -d`.

## Key Claims

- **Branching Workflow**: You can create a new branch and switch to it simultaneously using `git checkout -b <branch-name>`.
- **Fast-Forward Merges**: When merging a branch that is directly ahead of the current HEAD, Git simply moves the pointer forward. This is indicated by the phrase "fast-forward."
- **Three-Way Merges**: If histories have diverged, Git performs a three-way merge using the two branch tips and their common ancestor to create a new snapshot (merge commit).
- **Merge Conflicts**: Occur when both branches modify the same part of a file differently. Git pauses the merge and marks files as "unmerged" with conflict markers (`<<<<<<<`, `=======`, `>>>>>>>`).
- **Conflict Resolution**: Users must manually edit conflicted files to choose one version or merge them, stage the changes with `git add`, and complete the merge with `git commit`.
- **Merge Tools**: The command `git mergetool` can launch graphical tools (e.g., opendiff, kdiff3, meld) to assist in resolving conflicts.
- **Cleanup**: Once a branch is merged, it should be deleted using `git branch -d <branch-name>`.

## Suggested Links

- https://git-scm.com/book/en/v2/Git-Branching-Basic-Branching-and-Merging
