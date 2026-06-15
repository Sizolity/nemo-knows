---
title: Type Narrowing Strategies
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

# Type Narrowing Strategies

Type narrowing in Python refers to the process of restricting a variable's inferred or declared type to a more specific subset based on runtime conditions. This mechanism allows static type checkers and IDEs to provide better autocomplete suggestions, reduce false positives, and catch potential errors earlier in development.

## Implementation via `typing` Module

The `typing` module provides specialized forms for defining narrowing predicates:

- **`TypeIs`**: Narrows types by intersection. It is used with predicate functions that return a boolean value, allowing the type checker to infer the narrowed type within the conditional block where the predicate is `True`.
- **`TypeGuard`**: Narrows strictly to the guard type. Similar to `TypeIs`, it works with predicate functions but is designed for cases where the narrowing logic specifically identifies a particular subtype rather than an intersection of types.

Both forms are intended to be used alongside predicate functions that evaluate conditions at runtime, enabling the static analysis tools to understand that the type has changed within the scope of that condition.

## Related Concepts

- [[lint]]
- [[static-checking-utilities]]
- [[python-type-system-overview]]
