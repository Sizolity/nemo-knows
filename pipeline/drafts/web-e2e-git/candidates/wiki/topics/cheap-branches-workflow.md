---
title: Cheap Branches Workflow
kind: topic
sources:
  - source.md
  - raw/web/git-branching.md
confidence: medium
---

# Cheap Branches Workflow

Git branching is considered lightweight because the system stores data as snapshots rather than file deltas. In this model, a branch acts as a movable pointer to a specific commit. When new commits are made, the current branch pointer moves forward, creating a history of snapshots.

This architecture contrasts with older version control systems where branching often required copying many files into a new directory. Because a Git branch is essentially a small pointer file containing a commit checksum, it is inexpensive to create and switch between branches. This efficiency encourages frequent branch creation and merging workflows.

When switching branches, the working tree is updated to match the snapshot pointed to by the selected branch. After switching, new commits move the selected branch forward, which can create divergent history between branches. These relationships can be visualized using commands that display the commit graph.
