---
title: Go Concurrency Patterns
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Concurrency Patterns

Go’s concurrency model is built around the idea that communication, rather than shared memory, should synchronise independent activities. The guiding principle is “share memory by communicating”—values are passed through channels and remain owned by a single goroutine at any moment, eliminating data races by design.

**Goroutines** are the units of concurrent execution. Each goroutine is a function executing in the same address space as other goroutines, but with its own lightweight, growable stack. They are cheap to create and are multiplexed onto operating system threads so that blocking a goroutine (e.g. on a channel or I/O) does not block the underlying thread.

**Channels** are typed conduits that connect goroutines. An unbuffered channel synchronises sender and receiver: a send blocks until a concurrent receive, and vice versa. A buffered channel provides decoupling, blocking sends only when the buffer is full and receives only when it is empty. The `select` statement allows a goroutine to wait on multiple channel operations simultaneously, enabling patterns like timeouts, non-blocking communication, and fan‑in.

Go draws a clear line between **concurrency** (structuring a program as independently executing components) and **parallelism** (executing computations simultaneously across multiple CPUs). Its primitives make concurrent structure natural, but the language is fundamentally concurrent, not parallel; not every parallelisation problem maps cleanly to goroutines and channels. The runtime uses `runtime.GOMAXPROCS` to set the number of OS threads available for goroutine scheduling, typically defaulting to the number of logical CPUs.

Concurrency tools also simplify problems that are not inherently parallel. The “leaky buffer” pattern—drawn from an RPC package—manages a pool of reusable buffers with a channel of free buffers, avoiding explicit locks. This illustrates how channel-based communication can clarify resource management even in single-threaded or loosely coupled scenarios.
