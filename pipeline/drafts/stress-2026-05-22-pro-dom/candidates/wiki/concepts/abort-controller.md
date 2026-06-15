---
title: Abort Controller
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Abort Controller

The `AbortController` interface provides a standard way to signal cancellation of ongoing activities. It creates an associated `AbortSignal` object that can be passed to APIs that support aborting.

When the controller’s `abort()` method is called, the signal transitions to an aborted state. This causes registered abort handlers to run and notifies consumers of the signal. For example, passing an `AbortSignal` to `addEventListener()` via the `signal` option (see [[dom-events]]) will remove the event listener once the signal is aborted. The same mechanism can abort `fetch` requests and other promise‑based operations.
