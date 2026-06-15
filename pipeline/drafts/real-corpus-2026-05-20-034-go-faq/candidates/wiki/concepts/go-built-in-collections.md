---
title: Go Built In Collections
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/034-go-faq.md
confidence: medium
---

# Go Built In Collections

Go provides three primary built-in collection types: **arrays**, **slices**, **maps**, and **channels**. These types serve as the fundamental building blocks for data storage and concurrency within the language.

## Types and Semantics

*   **Arrays**: Arrays are value types, meaning they store their data directly. A slice of arrays is not supported in Go.
*   **Slices**: Slices act as references (descriptors) pointing to shared underlying data rather than storing values themselves. They are the preferred dynamic collection for most use cases.
*   **Maps**: Maps are built-in associative collections implemented using hash tables. Like slices and maps, they act as references pointing to shared data structures.
    *   **Key Constraints**: Go maps cannot accept slices as keys because equality is not well-defined for them. While maps can hold pointers or interface values, primitive types like integers and strings serve as standard keys.
*   **Channels**: Channels are typed conduits used for passing messages between goroutines. They act as references to the underlying buffer and synchronization logic managed by the runtime.

## Concurrency Behavior

Go collections designed for concurrent use—specifically slices, maps, and channels—exhibit specific behaviors regarding atomicity and safety:

*   **Atomicity**: Map operations are not atomic by default. While concurrent read-only access to a map is generally safe (due to internal locking mechanisms in the runtime), writes require explicit synchronization using `sync.Mutex` or similar constructs.
*   **Goroutines**: Channels facilitate communication between goroutines, which are lightweight coroutines managed by the Go runtime. The runtime allows hundreds of thousands of goroutines to exist simultaneously, making channels essential for structuring concurrent programs via Communicating Sequential Processes (CSP) concepts.

## Design Principles

The inclusion of these collections reflects core Go design decisions:
*   **Simplicity**: Go avoids complex type hierarchies and exceptions in favor of simple, explicit mechanisms like maps for key-value storage and channels for synchronization.
*   **Safety**: By treating slices and maps as references to shared data, the language encourages developers to manage lifetimes carefully (often utilizing garbage collection) rather than relying on manual allocation/deallocation or unsafe pointer arithmetic.
*   **Structural Typing**: Interfaces in Go are satisfied implicitly based on method sets, allowing these collections to be passed between functions without explicit type declarations when their behavior matches the interface contract.
