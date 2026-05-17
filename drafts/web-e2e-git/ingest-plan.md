---
kind: topic
sources: [raw/web/git-branching.md]
status: draft
---

# Ingest Plan

## Source Summary
- The Pro Git chapter explains that Git branching is lightweight because Git stores data as snapshots and branches are movable pointers to commits.
- A Git branch is a pointer to one commit, where the default branch name is often `master`, but it is not special; creating a branch does not switch context until explicitly checked out.
- Switching branches changes the working tree to match the selected snapshot, allowing divergent history that can be visualized using specific log commands.

## Candidate Wiki Pages
- wiki/sources/git-branching.md — Documents the raw source material regarding Git branching mechanics and storage model.
- wiki/concepts/commit-snapshots.md — Captures the core concept that Git stores data as snapshots connected by commit objects.
- wiki/topics/cheap-branches-workflow.md — Explores the topic of why branches are cheap pointers encouraging frequent merge workflows.

## Suggested Links
- none

## Review Checklist
- [ ] Verify candidate slugs match wiki directory conventions.
- [ ] Ensure no schema or index files were inadvertently proposed as candidates.
