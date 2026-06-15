---
title: Rebasing
kind: source
sources:
  - raw/web/corpus-2026-05-18/004-rebasing.md
confidence: medium
---

# Rebasing

## What It Is
Rebasing is a Git command that integrates changes from one branch into another by replaying the commits of one branch onto another. Instead of creating a merge commit, it takes the patch of the change introduced in a commit and reapplies it on top of the latest snapshot of the target branch. This operation works by identifying the common ancestor, saving diffs to temporary files, resetting the current branch, and applying each change in turn.

## Summary
The Git book distinguishes between two main ways to integrate changes: merging and rebasing. While merging creates a new commit representing the three-way merge of snapshots, rebasing replays commits sequentially to create a linear history. The primary use case for rebasing is to ensure local commits apply cleanly onto a remote branch (like `origin/master`) before submitting patches, ensuring the maintainer only needs a fast-forward or clean apply. However, rebasing must be avoided for public commits that others have based work on, as rewriting shared history can cause significant confusion and require collaborators to re-merge their work.

## Key Claims
- **Linear History:** Rebasing makes the commit log appear linear, showing all work happened in series even if it occurred in parallel.
- **Snapshot Equivalence:** The final snapshot resulting from a rebase is identical to that of a merge; only the history graph differs.
- **Public Commit Rule:** Do not rebase commits that exist outside your repository and that people may have based work on. Rewriting public history forces collaborators to resolve conflicts and can lead to messy state.
- **Patch-ID Resolution:** Git uses patch checksums (patch-id) in addition to SHA-1 to determine if rewritten commits are duplicates, allowing `git pull --rebase` to automatically apply unique local changes on top of force-pushed rebase work.
- **Philosophical Divide:** Some view commit history as a historical record that should not be tampered with (favoring merge), while others view it as a story of the project's creation and prefer a clean, coherent narrative using tools like `rebase` and `filter-branch`.

## Suggested Links
- https://git-scm.com/book/en/v2/Git-Branching-Rebasing
