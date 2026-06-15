---
title: Detached Head State
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/009-git-references.md
confidence: medium
---

# Detached Head State

In Git, the **HEAD** file normally acts as a symbolic reference pointing to a branch, such as `ref: refs/heads/master`. However, it can contain a raw SHA-1 value instead. When this occurs, the repository is in a "detached HEAD" state.

In this mode, HEAD points directly to a specific commit object rather than a named branch. This allows users to inspect or modify commits without affecting any existing branch history, but changes made to files will not be automatically associated with a branch unless one is created manually.
