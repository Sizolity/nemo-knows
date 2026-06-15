---
title: Dom Ranges
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Ranges

A DOM range captures a span of content inside a [[dom-nodes|node tree]] using a start and an end boundary point. Each boundary point is a (node, offset) pair. The shared `AbstractRange` interface exposes read‑only `startContainer`, `startOffset`, `endContainer`, `endOffset`, and a `collapsed` flag.

Two range implementations exist:

- **Live `Range`** – automatically re‑anchors itself when the tree mutates, continuing to represent the same logical content. This makes it suitable for editing and selections, but updates on every tree change can be expensive.
- **`StaticRange`** – a lightweight, immutable snapshot whose boundaries never shift.

Ranges can only represent node‑based content; attribute values are excluded.
