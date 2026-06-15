---
title: Abort Signal
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Abort Signal

An **AbortSignal** communicates a request to abort an ongoing, cancelable operation. It is typically obtained from an `AbortController` or through the static factory methods `AbortSignal.abort()`, `AbortSignal.timeout()`, and `AbortSignal.any()`.

A signaled abort can be checked programmatically via `throwIfAborted()`. In the DOM, `AbortSignal` integrates with event listener registration: `addEventListener` and `removeEventListener` accept a signal, allowing registered listeners to be automatically removed when the signal aborts (see [[dom-events]]).
