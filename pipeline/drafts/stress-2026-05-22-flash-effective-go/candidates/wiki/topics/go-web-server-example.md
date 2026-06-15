---
title: Go Web Server Example
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Web Server Example

The *Effective Go* guide closes with a small web server that demonstrates idiomatic ways to structure HTTP handlers and combine them with [[go-concurrency]].
A handler only needs to implement `ServeHTTP`. The simplest stateful handler is a `Counter` struct with an integer field, incremented on each request and written back to the `ResponseWriter`. To make it even lighter, the handler can be defined directly on an `int` type; the pointer receiver lets the count persist across calls (real‑world use would need synchronisation, e.g. from `sync/atomic`).

Channels can serve as handlers too. A `chan *http.Request` handler sends the incoming request into the channel and replies immediately, notifying an internal component without sharing memory.
When a plain function fits the desired signature, `http.HandlerFunc` turns it into a handler: registering `http.Handle("/args", http.HandlerFunc(ArgServer))` makes the function serve arguments on each visit.

The example then grows into a miniature RPC framework. A `Request` struct bundles the arguments, the operation function, and a result channel. A goroutine‑based server reads from a shared request channel, computes the result, and sends it back on the client’s result channel. The pattern achieves parallel, rate‑limited request handling without explicit mutexes, aligning with the “share memory by communicating” principle.

Together, these snippets illustrate [[go-idioms]] for building HTTP services and integrating concurrency primitives cleanly.
