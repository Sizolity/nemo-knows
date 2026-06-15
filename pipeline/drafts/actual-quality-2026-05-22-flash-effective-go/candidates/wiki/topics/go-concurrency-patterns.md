---
title: Go Concurrency Patterns
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Concurrency Patterns

## Share Memory by Communicating

Go's concurrency philosophy is captured by the slogan "share memory by communicating". Instead of guarding shared variables with locks, programs pass ownership of values through channels. Only one goroutine holds a value at any time, so data races are eliminated by design. This approach treats communication itself as the synchronisation mechanism, much like Unix pipelines, making explicit mutexes unnecessary in the common case.

## Goroutines

Goroutines are lightweight, independently executing functions that share a single address space. They start with small stacks that grow and shrink as needed by allocating and freeing heap storage, making them cheap to create and multiplex onto OS threads. The model avoids the overhead and naming confusion of traditional threads, coroutines, or processes.

## Concurrency vs Parallelism

Go distinguishes *concurrency* – structuring software as independently executing components – from *parallelism* – performing computations simultaneously on multiple CPUs. The language’s design supports concurrent composition, not all forms of automatic parallelisation. Understanding this difference helps writers apply goroutines and channels effectively without expecting them to solve every performance problem.
