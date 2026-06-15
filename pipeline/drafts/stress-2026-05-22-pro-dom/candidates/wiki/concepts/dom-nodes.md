---
title: Dom Nodes
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Nodes

A **DOM node** is one of the fundamental building blocks of a document tree, as defined in the DOM Standard. Nodes come in several varieties—`Element`, `Text`, `Comment`, `Document`, `DocumentFragment`, `Attr`, and processing instructions—and follow a strict tree structure with parent, child, and sibling relationships. A tree order is maintained across the document.

The standard specifies precise **mutation algorithms** for inserting, removing, replacing, and moving nodes, which automatically trigger mutation observers. Common behaviours are shared through mixins such as `ParentNode`, `ChildNode`, and `Slottable`.

All nodes act as **event targets**; they participate in event propagation (capture, target, bubble) as described in [[dom-events]]. Additionally, [[dom-ranges]] represent continuous portions of the node tree by using a pair of (node, offset) boundary points, which live ranges adjust when mutations occur.
