---
kind: topic
sources: [raw/web/git-branching.md]
status: draft
---

# Ingest Plan

## Source Summary
- Git stores data as snapshots; commits link trees and blobs into history.
- Branches are lightweight, movable pointers to commits; `HEAD` tracks the current branch.
- Creating a branch just copies a commit checksum to a new pointer file.
- Switching branches updates the working tree to match the destination commit, enabling divergent histories.

## Candidate Wiki Pages
- wiki/sources/git-branching.md — raw notes summarising the “Branches in a Nutshell” chapter
- wiki/concepts/git-branch.md — canonical definition of a Git branch as a movable pointer
- wiki/concepts/git-head.md — explanation of the HEAD reference and branch switching
- wiki/topics/git-lightweight-branching.md — how Git’s cheap branching enables frequent branching workflows

## Suggested Links
- External: https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell

## Review Checklist
- [ ] All facts extracted match the source snapshot
- [ ] Slugs are consistent with existing naming conventions
- [ ] Concept pages define terms without step-by-step instructions
- [ ] Topic page focuses on workflows and practical implications
- [ ] Cross-references between related pages are sensible
