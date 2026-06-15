---
title: Interfaces And Embedding
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Interfaces And Embedding

In Go, an **interface** is a set of methods. A type implements an interface if it has the required methods, adhering to duck typing principles rather than explicit declaration. This mechanism allows for flexible composition where a type can satisfy multiple interfaces simply by possessing the necessary behavior.

The language also supports **embedding**, which promotes methods from a parent type to a child type while retaining the inner receiver type unless explicitly changed. Embedding is often used to build complex types that inherit specific behaviors without creating explicit interface implementations.

To verify if a variable satisfies an interface at compile time, developers can use the blank identifier `_` in a declaration like `var _ Interface = value`. This ensures the code compiles only if the type implements the required methods. While the blank identifier is also used to discard values or errors during range operations, discarding errors in practice is generally discouraged in favor of explicit handling.

When defining types that represent agents or actions, naming conventions follow the `<method>-er` pattern (e.g., `Reader`, `Writer`). For multiword identifiers, MixedCaps are preferred over underscores. Package names should remain short, concise, evocative, lower-case, and single-word where possible.
