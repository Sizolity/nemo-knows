---
title: DOM Standard
kind: source
sources:
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

## What It Is

The DOM Standard defines the platform‑neutral Document Object Model: a tree‑based interface for accessing and manipulating the content, structure, and style of documents.

## Summary

This snapshot captures the complete normative text of the DOM Standard. It opens with foundational infrastructure—trees, ordered sets, selectors, and name‑validation rules—then specifies the event system and the cooperative cancellation pattern (`AbortController`/`AbortSignal`). The core of the work defines the node hierarchy: document trees, shadow trees with slot assignment, all low‑level mutation algorithms (insert, move, replace, remove), and the full set of interfaces from `Node` down to leaf types like `Text` and `Comment`. The standard also covers live and static ranges (`AbstractRange`, `StaticRange`, `Range`), traversal (`NodeIterator`, `TreeWalker`), token lists (`DOMTokenList`), and legacy XPath/XSLT APIs. A long historical chapter catalogues every removed feature and provides extensive browser‑compatibility tables derived from MDN. The entire specification emphasizes live collections, strict mutation rules that integrate with custom elements and `MutationObserver`, and shadow‑DOM encapsulation.

## Key Claims

- **Event dispatching is a path‑based capture‑then‑bubble flow**; events do not represent actions.
- **The DOM is a tree of nodes with strict parent‑child contracts**, governing hierarchies and traversal order.
- **Mutation algorithms (insert, move, replace, remove) are low‑level primitives** reused by higher‑level methods and back‑integrated with live ranges, custom element callbacks, and `MutationObserver`.
- **`AbortSignal` objects are designed for cooperative cancellation**, with dependent signals, static factory methods, and a binding pattern for promise‑returning web APIs.
- **Shadow trees provide encapsulation through slots and assigned nodes**; the signal‑and‑assign cycle and flattened tree model underpin custom element behaviour and event retargeting.
- **Live ranges and `MutationObserver` maintain consistency** after tree changes, with range pre‑remove steps and mutation‑record queuing interwoven with the mutation primitives.
- **Name‑validation rules are intentionally loose**, matching the HTML parser’s output rather than strict XML constraints.
- The standard records **an exhaustive list of historically removed features** (e.g., `MutationEvent`, `Entity`, `DOMConfiguration`) and frozen compatibility snapshots for all core interfaces.

## Suggested Links

none
