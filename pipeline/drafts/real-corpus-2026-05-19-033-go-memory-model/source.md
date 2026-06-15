---
title: Go Memory Model
kind: source
sources:
  - raw/web/corpus-2026-05-18/033-go-memory-model.md
confidence: medium
---

# Go Memory Model

## What It Is
A specification detailing the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes in a different goroutine. The model defines requirements on program executions composed of memory operations (reads, writes, and synchronizing operations) to ensure data-race-free programs behave as if executed sequentially.

## Summary
The Go memory model aims for simplicity and understandability, closely following the approach presented by Boehm and Adve regarding C++ concurrency. It guarantees that data-race-free (DRF) programs execute in a sequentially consistent manner (DRF-SC). While programmers are strongly encouraged to use synchronization to avoid data races, the implementation may react to a race by reporting it and terminating the program, rather than exhibiting undefined behavior as seen in C or C++.

## Key Claims
- **Sequential Consistency for Race-Free Programs:** Any Go program that is data-race-free can only have outcomes explained by some sequentially consistent interleaving of goroutine executions.
- **Implementation Constraints:** In the absence of data races, programs behave as if all goroutines are multiplexed onto a single processor. If a race occurs, an implementation may halt execution and report the race.
- **Synchronization Primitives:** Access to shared data must be serialized using channel operations or primitives from the `sync` and `sync/atomic` packages.
- **Atomic Semantics:** Atomic operations in `sync/atomic` behave as if executed in some sequentially consistent order, similar to C++ atomics or Java volatile variables.
- **Compiler Restrictions:** Compilers must not introduce writes that do not exist, assume loops terminate without synchronization, or reload local variables from shared memory in a way that violates the happens-before relationship.

## Suggested Links
- The Go Programming Language: https://go.dev/ref/mem
