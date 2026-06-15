---
title: Shadow Dom
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Shadow Dom

Shadow DOM provides encapsulation for a DOM subtree by attaching a **shadow tree** to a host element. The shadow tree has its own [[dom-nodes|shadow root]] (`ShadowRoot`) and is hidden from the document’s main tree.
- **Slot-based projection** (manual or named assignment) maps the host’s light DOM children into the shadow tree.
- Events that cross shadow boundaries are **retargeted** to preserve encapsulation; see [[dom-events]] for event retargeting details.
