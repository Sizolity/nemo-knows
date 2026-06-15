---
title: Data Race Free Semantics
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/033-go-memory-model.md
confidence: medium
---

# Data Race Free Semantics

Data race free semantics is a specification detailing the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes in a different goroutine. It defines requirements on program executions composed of memory operations (reads, writes, and synchronizing operations) to ensure data-race-free programs behave as if executed sequentially.

## Sequential Consistency for Race-Free Programs

Any Go program that is data-race-free can only have outcomes explained by some sequentially consistent interleaving of goroutine executions. In the absence of data races, programs behave as if all goroutines are multiplexed onto a single processor.

## Implementation Constraints

While programmers are strongly encouraged to use synchronization to avoid data races, the implementation may react to a race by reporting it and terminating the program, rather than exhibiting undefined behavior. Access to shared data must be serialized using channel operations or primitives from the `sync` and `sync/atomic` packages. Atomic operations behave as if executed in some sequentially consistent order.

## Compiler Restrictions

Compilers must not introduce writes that do not exist, assume loops terminate without synchronization, or reload local variables from shared memory in a way that violates the happens-before relationship.
