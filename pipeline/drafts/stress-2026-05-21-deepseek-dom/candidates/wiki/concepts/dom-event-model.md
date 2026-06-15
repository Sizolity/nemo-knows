---
title: Dom Event Model
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Event Model

The DOM event model defines a path-based dispatching mechanism: events traverse the [[dom-node-tree]] in a capture‑then‑bubble flow, following the ancestor chain from the document root down to the target and back. Events are notification objects rather than direct representations of user actions.
