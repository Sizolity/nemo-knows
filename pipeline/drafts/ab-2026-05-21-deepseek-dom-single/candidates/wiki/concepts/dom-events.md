---
title: Dom Events
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Events

DOM events are notifications that signal occurrences, not actions. The `Event` interface and its subtypes (such as `CustomEvent`) model these occurrences. Events propagate through capture, target, and bubble phases on `EventTarget` objects, which are typically [[dom-nodes]].

Listeners are registered with `addEventListener()` and removed with `removeEventListener()`. When an [[abort-signal]] is passed to `addEventListener()`, the listener is automatically removed when the signal’s controller calls `abort()`.

In [[shadow-dom]] trees, events are retargeted as they cross shadow boundaries to maintain encapsulation. Handlers can call `event.preventDefault()` to indicate that any associated default browser action should not be performed.
