---
title: Dom Node Tree
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Node Tree

The DOM represents a document as a tree of nodes whose strict parent–child contracts determine the allowed hierarchy and traversal order. Each node has at most one parent; only node types permitted by the tree rules can appear as children.

Event propagation follows the tree structure through a path-based capture-then-bubble flow (see [[dom-event-model]]). Low‑level mutation primitives—insert, move, replace, remove—operate directly on the node tree and are reused by higher‑level methods (see [[dom-mutation-algorithms]]). Live ranges and `MutationObserver` rely on pre‑remove steps and mutation‑record queuing to stay consistent with tree changes (see [[dom-ranges]]).

The standard defines the entire node hierarchy, from `Document` and `DocumentFragment` down to leaf types such as `Text` and `Comment`, and includes shadow trees for encapsulation and slot assignment.
