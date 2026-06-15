---
title: Go Concurrency Patterns
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Concurrency Patterns

Go's concurrency model is built around the principle of "sharing memory by communicating." This approach prioritizes safety and simplicity over performance optimizations like complex locking mechanisms.

## Core Concepts

### Goroutines
Goroutines are lightweight threads managed by the Go runtime. Unlike OS threads, they have a small stack size that grows and shrinks as needed, allowing for massive concurrency at low cost.

### Channels
Channels are typed send/receive buffers that facilitate communication between goroutines. They serve as synchronization primitives, replacing locks and mutexes to avoid data races.
- **Unbuffered channels**: Blocks until both a sender and receiver are ready.
- **Buffered channels**: Can hold a specified number of values before blocking.

## Patterns

### Communication Over Shared Memory
The preferred strategy is to keep shared memory away from goroutines. Instead, use channels to pass data between them. This eliminates the need for complex locking logic and reduces the risk of deadlocks.

```go
// Example: Sharing memory by communicating
var jobs = make(chan int)
for i := 0; i < 10; i++ {
    jobs <- i
}
close(jobs)
```

### Semaphore Pattern
Buffered channels can be used to limit concurrency, acting as a semaphore. This is useful for controlling the number of concurrent goroutines accessing a shared resource (e.g., network connections or CPU cores).

### Error Handling in Concurrency
When working with goroutines, errors must be propagated explicitly. Do not discard errors silently; instead, collect them and return them to the caller. Use `recover` only within deferred functions during stack unwinding to catch panics that occur during shutdown or cleanup.
