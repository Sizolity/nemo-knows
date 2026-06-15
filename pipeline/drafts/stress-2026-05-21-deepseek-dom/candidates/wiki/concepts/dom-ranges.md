---
title: Dom Ranges
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Ranges

A **range** represents a contiguous fragment of a [[dom-node-tree]], defined by a start and end boundary (node plus offset). The DOM standard provides two kinds: live `Range` objects and their lighter static counterpart, `StaticRange`. Live ranges automatically adapt to tree changes through integration with [[dom-mutation-algorithms]]; the range’s boundaries are updated by pre-remove steps and other mutation primitives to keep the range consistent after insertions, deletions, or node moves. `StaticRange` is an immutable snapshot that does not follow mutations. Together they provide a programmatic way to select, extract, and manipulate document content without requiring full serialisation.
