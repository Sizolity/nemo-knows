---
title: Go Memory And Allocation
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/034-go-faq.md
confidence: medium
---

# Go Memory And Allocation

In Go, memory management is automated by the runtime through a parallel mark-and-sweep garbage collector. This approach eliminates the need for manual allocation and deallocation, providing sub-millisecond pauses during collection cycles [[go-memory-and-allocation]].

## Heap vs Stack Allocation

The Go runtime performs **escape analysis** to determine whether a variable should be allocated on the stack or the heap. Variables that are guaranteed not to escape the current function scope can reside on the stack, while those that interact with external pointers or outlive their scope are moved to the heap [[go-memory-and-allocation]].

## Built-in Collections and References

Go includes built-in collections such as slices, maps, and channels, which act as references (descriptors pointing to shared data) rather than values [[go-built-in-collections]]. In contrast, arrays are treated as values. While maps provide significant utility, they cannot accept slices as keys because equality is not well-defined for them [[go-built-in-collections]].

## Concurrency and Memory

Go's concurrency model relies on goroutines and channels to manage execution without explicit locks in many cases [[go-concurrency-model]]. Goroutines are lightweight coroutines with resizable stacks managed by the runtime, allowing hundreds of thousands to exist simultaneously. This design supports high concurrency while keeping memory usage efficient through the garbage collector.

## Safety and Atomicity

Map operations are not atomic by default. While concurrent read-only access is safe, writes to maps require explicit synchronization to prevent data races. This explicit requirement aligns with Go's focus on safety and maintainability over implicit complexity [[go-memory-and-allocation]].
