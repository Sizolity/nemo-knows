---
title: Dom Nodes
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Nodes

A DOM node is the fundamental building block of every Document Object Model tree. The DOM Standard defines a unified node hierarchy rooted at `Node`, with specialised subtypes such as `Element`, `Text`, `Document`, `DocumentFragment`, and `ShadowRoot`. Nodes are arranged in parent‑child relationships that respect tree order, and they may carry attributes, text content, or child node collections.

Node tree mutations—insertion, removal, moving, and replacement—are performed through a precisely sequenced algorithm that includes insertion steps, removing steps, post‑connection steps, and children‑changed steps. This design keeps mutations atomic and guarantees correct lifecycle notifications, including custom element callbacks.

Shadow DOM introduces encapsulation by allowing a host element to attach a separate shadow tree. Slot‑based projection (either named or manual) controls how light‑DOM content is distributed, and event retargeting preserves the illusion of encapsulation across shadow boundaries (see [[shadow-dom]]).

Asynchronous observation of node changes is handled by mutation observers ([[mutation-observers]]). They collect batches of `MutationRecord` objects and deliver them via a microtask‑based callback, superseding the older synchronous `MutationEvent` mechanism. This avoids the performance and correctness pitfalls of synchronous observation.
