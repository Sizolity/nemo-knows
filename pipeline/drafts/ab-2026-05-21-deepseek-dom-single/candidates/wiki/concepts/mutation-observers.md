---
title: Mutation Observers
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Mutation Observers

The `MutationObserver` API provides asynchronous, batched reporting of changes to a [[dom-nodes]] tree. It replaces the deprecated synchronous `MutationEvent`.
An observer is created with a callback and configured to watch specific mutation types (child list, attributes, character data). Mutations are collected and delivered as a sequence of `MutationRecord` objects after the current task’s microtasks, keeping observation non-blocking. Calling `disconnect()` stops observation, and `takeRecords()` drains any pending records synchronously.
