---
title: Synchronization Primitives
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/033-go-memory-model.md
confidence: medium
---

# Synchronization Primitives

In concurrent programming, **synchronization primitives** are mechanisms used to coordinate access to shared resources among multiple goroutines or threads. They ensure that operations on shared data occur in a defined order and prevent data races.

## Purpose and Guarantees

The primary goal of synchronization primitives is to serialize access to shared state, ensuring that a read operation observes values produced by prior write operations in other threads. In the context of the Go memory model, these primitives are essential for achieving **data-race-free semantics**. When used correctly, they allow programs to behave as if executed sequentially (sequentially consistent), even though execution is actually interleaved across multiple goroutines.

## Usage in Go

In Go, shared data must be accessed using synchronization to avoid undefined behavior. The standard library provides several categories of primitives:

- **Channel Operations**: Channels act as both communication and synchronization mechanisms. Sending or receiving on a channel creates an implicit happens-before relationship between the sender and receiver.
- **`sync` Package**: This package offers high-level abstractions such as `Mutex`, `RWMutex`, `Cond`, and `WaitGroup` to manage complex locking patterns and coordination.
- **`sync/atomic` Package**: Provides low-level atomic operations (e.g., `AddInt32`, `SwapInt64`) that behave sequentially consistently without requiring locks, suitable for simple counters or flags.

## Compiler and Implementation Behavior

Compilers must respect the happens-before relationships established by these primitives. They cannot reorder memory operations in a way that violates these guarantees, nor can they assume loops terminate without explicit synchronization. If a data race occurs despite these measures, the Go implementation may halt execution and report the race rather than exhibiting undefined behavior.
