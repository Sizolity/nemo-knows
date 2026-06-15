---
title: Dom Ranges
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Ranges

Dom Ranges represent a contiguous span of content inside a document, defined by two boundary points that each consist of a [[dom-nodes]] reference and a numeric offset. The DOM Standard layers three range interfaces on this boundary‑point model: the abstract `AbstractRange`, the lightweight and non‑live `StaticRange`, and the live `Range`. A live `Range` tracks mutations of the node tree and adjusts its boundaries automatically, staying aligned with the original selection. In contrast, a `StaticRange` does not react to tree changes and remains cheaper to create for one‑off snapshot use cases.
