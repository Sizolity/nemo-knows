---
title: Dom Events
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Events

In the DOM Standard, events are objects dispatched to signal an occurrence such as network activity or user interaction. Objects that receive events implement the `EventTarget` interface.

- Event listeners are registered with `addEventListener()` and removed with `removeEventListener()`. An `AbortSignal` passed to `addEventListener()` ties the listener’s lifetime to an [[abort-controller]]: calling `abort()` on the controller removes the listener.
- Events implement the `Event` interface (or a derived interface like `CustomEvent`). The `type` attribute identifies the kind of occurrence; `target` points to the object the event was dispatched to.
- Applications can create synthetic events using constructors such as `new CustomEvent("cat", { detail: … })` and dispatch them via the object’s `dispatchEvent()` method.
- `dispatchEvent()` returns `false` if `preventDefault()` was called on the event, allowing callers to conditionally abort subsequent logic.
- When an event is dispatched to an object that participates in a tree (e.g., an element), propagation follows a capture phase (ancestor listeners with `capture: true`, in tree order) and, if the event’s `bubbles` is `true`, a bubble phase (ancestor listeners without capture, in reverse tree order). Calling `preventDefault()` during propagation signals that an associated default action should not be performed.
