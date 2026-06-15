---
kind: topic
sources: [raw/web/corpus-2026-05-18/004-rebasing.md]
status: draft
---

# Ingest Plan

## Source Summary
- **Content**: A comprehensive section from the Pro Git book covering rebasing mechanics, advanced workflows (e.g., `--onto`), and the critical warning against rebasing public history.
- **Key Concepts**: Linear history vs. merge commits, replaying patches, force-pushing consequences, and patch-id conflict resolution.
- **Utility**: Provides definitive rules for team collaboration regarding history rewriting and clean integration strategies.

## Candidate Wiki Pages
- wiki/sources/git-rebasing-guide.md — To store the full text of this specific chapter as a curated source artifact.
- wiki/concepts/git-rebase-mechanics.md — To document the technical process of replaying commits onto a new base branch.
- wiki/topics/git-history-management-best-practices.md — To synthesize guidelines on when to rebase locally versus never rebasing public history.

## Suggested Links
- https://git-scm.com/book/en/v2/Git-Branching-Rebasing

## Review Checklist
- [ ] Verify that the "Do not rebase public commits" warning is highlighted in the summary.
- [ ] Ensure `wiki/concepts/git-rebase-mechanics.md` explains the `--onto` option clearly.
- [ ] Confirm that `wiki/topics/git-history-management-best-practices.md` includes the `pull.rebase` configuration tip.
