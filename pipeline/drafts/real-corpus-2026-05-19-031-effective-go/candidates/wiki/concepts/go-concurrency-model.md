---
title: Go Concurrency Model
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Concurrency Model

The Go concurrency model is built around the principle of **"share by communicating"**. This approach encourages avoiding shared mutable state in favor of passing data through channels. Goroutines are lightweight functions managed by the runtime that allow for concurrent execution.

Communication between goroutines often utilizes channels, which act as rendezvous points. Unbuffered channels block until both sender and receiver are ready, while buffered channels limit concurrency based on their capacity. This design promotes safe parallelism without requiring complex locking mechanisms.

For more details on related concepts, see:
- [[go-allocation-new-vs-make]]
- [[go-control-flow-idioms]]
- [[go-defer-mechanics]]
- [[go-error-handling-panic-recovery]]
