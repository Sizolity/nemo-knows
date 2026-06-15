---
title: Go Concurrency
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Concurrency

Go approaches concurrency through goroutines and channels, guided by the principle “share memory by communicating.” Instead of multiple threads locking shared variables, values travel over channels so that only one goroutine owns a value at any moment. This eliminates data races by design.

A goroutine is a lightweight function executing concurrently with others in the same address space. Its stack starts small and grows as needed, making goroutines cheap enough to spawn many thousands. The runtime multiplexes goroutines onto OS threads; by default it uses all available CPU cores, though this can be adjusted with `runtime.GOMAXPROCS`.

Channels form the communication backbone. An unbuffered channel synchronises send and receive, while a buffered channel can act as a semaphore. More elaborate patterns, such as channels of channels, support request–response interactions. Concurrency is about structuring independent components—parallel execution is a separate concern, not guaranteed by the concurrent design.

This concurrency model is woven into Go’s idiomatic style ([[go-idioms]]). The focus on communication over shared state helps produce clearer, correct programs, and even simplifies some non‑concurrent designs like a leaky‑buffer free list.
