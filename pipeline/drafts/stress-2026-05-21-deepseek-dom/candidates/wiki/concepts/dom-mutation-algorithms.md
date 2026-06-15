---
title: Dom Mutation Algorithms
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Mutation Algorithms

The DOM’s low‑level mutation primitives—insert, move, replace, and remove—form the foundation for all tree modifications. They enforce the strict parent–child contracts defined by the [[dom-node-tree]] and are reused internally by every higher‑level DOM method. These primitives are interwoven with live [[dom-ranges]] (in particular through range pre‑remove steps) and with `MutationObserver` record queuing, so tree consistency is preserved across all structural changes. An older event‑driven mutation scheme (`MutationEvent`) has been removed from the platform; see [[dom-legacy-and-removed-apis]].
