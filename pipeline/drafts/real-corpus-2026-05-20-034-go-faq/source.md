---
title: Go FAQ
kind: source
sources:
  - raw/web/corpus-2026-05-18/034-go-faq.md
confidence: medium
---

## What It Is

The **Go FAQ** is a comprehensive question-and-answer resource for The Go Programming Language, hosted at `https://go.dev/doc/faq`. Originally created in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson to address multicore complexity, it serves as the primary reference for understanding Go's design philosophy, language semantics, concurrency models, and tooling. The document covers the history of Go (open-sourced in 2009), its orthogonality principles, explicit error handling strategies, structural typing rules, and testing best practices. It also details compiler technology (the self-hosting `gc` compiler), binary optimization strategies, and guidelines for contributing to the ecosystem.

## Summary

The Go FAQ aggregates technical answers to common questions about using, compiling, and extending Go. The content begins with metadata regarding its retrieval from `go.dev`, followed by a historical overview of its creation and open-source release. The core body explains key design decisions: rejecting exceptions and assertions in favor of multi-value returns; adopting Communicating Sequential Processes (CSP) for concurrency via channels and goroutines; implementing structural typing instead of inheritance; and managing memory through garbage collection rather than manual allocation.

Technical discussions span from basic syntax rules (like the absence of implicit numeric conversions and the ternary operator) to advanced runtime behaviors (escape analysis, parallelism limits, and closure capture evolution in Go 1.22). The document also outlines the module system introduced in Go 1.11, binary linking strategies (static linking by default), and testing philosophies that prioritize table-driven tests over assertion libraries. Finally, it concludes with community engagement pathways, including bug filing, pull requests, and social media connectivity.

## Key Claims

- **Design Philosophy:** Go prioritizes simplicity, safety, and maintainability over features like exceptions, assertions, or complex type hierarchies found in C++/Java. It uses structural typing where a type satisfies an interface implicitly if it possesses the required methods.
- **Concurrency Model:** The language utilizes CSP concepts with channels as first-class objects for communication and goroutines (lightweight coroutines with resizable stacks) for execution. Goroutines are managed by the runtime, allowing hundreds of thousands to exist simultaneously.
- **Error Handling:** Go avoids exceptions in favor of returning multiple values where the last is an error. Assertions are discouraged because they encourage ignoring errors; explicit handling ensures servers remain running and provides precise error messages.
- **Built-in Collections:** Maps are built-in due to their power but cannot accept slices as keys because equality is not well-defined for them. Slices, maps, and channels act as references (descriptors pointing to shared data), whereas arrays are values. Map operations are not atomic by default; concurrent read-only access is safe, but writes require synchronization.
- **Memory Management:** The runtime uses a parallel mark-and-sweep garbage collector with sub-millisecond pauses, eliminating manual memory management. Escape analysis determines heap vs. stack allocation.
- **Testing Strategy:** Go avoids assertion libraries to ensure all tests run after a failure and prefers native testing frameworks. Table-driven tests are encouraged to amortize the cost of writing good error messages across many cases.
- **Generics & Syntax:** Generics were added in Go 1.18 to balance complexity with utility. Type parameters use square brackets (`[]`) to avoid ambiguity with the less-than operator. Generic methods are not supported to prevent infinite implementation strategies.
- **Compiler & Binary:** The default compiler `gc` is now self-hosting (written in Go) with a recursive descent parser. Binaries are statically linked by default, including the runtime and type info; size can be reduced with `-ldflags=-w`.

## Suggested Links

- none
