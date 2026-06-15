---
title: Dom Abort Signal
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/060-dom-standard.md
confidence: medium
---

# Dom Abort Signal

An `AbortSignal` object represents a cooperative cancellation token.
It can be associated with one or more dependent signals, and the platform provides a static `AbortSignal.abort()` factory for an already-aborted signal.
Cancellation is opt‑in: a web API that accepts a signal will observe its aborted state and stop any in‑progress work; the signal does not forcibly terminate operations.
The signal itself is an `EventTarget` and fires an `abort` event when cancellation is triggered, but the DOM Standard models the pattern primarily through the paired `AbortController`/`AbortSignal` interface rather than by extending the general event‑dispatching flow.
```
