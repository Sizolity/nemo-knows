---
title: Go Structural Typing
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/034-go-faq.md
confidence: medium
---

# Go Structural Typing

In the Go programming language, type compatibility is determined by **structure** rather than declaration. This approach, often referred to as "duck typing," allows a type to satisfy an interface implicitly if it possesses the required methods and fields, regardless of whether it explicitly declares that it implements the interface.

## Design Philosophy

This design decision aligns with Go's broader **[[go-design-principles]]**, which prioritize simplicity and maintainability over features found in languages like C++ or Java. By avoiding explicit inheritance hierarchies and complex type declarations, Go reduces boilerplate code and makes the relationship between types more discoverable through inspection of the code itself.

## Mechanism

A type `T` satisfies an interface `I` if `T` has all the methods defined in `I`. The language does not require or support explicit implementation declarations (e.g., `implements I`).

### Example

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}

type BufferedWriter struct{}

func (b *BufferedWriter) Write(p []byte) (n int, err error) {
    // implementation
}

// Even though BufferedWriter does not declare "implements Writer",
// it satisfies the Writer interface because it has the required method.
var w Writer = &BufferedWriter{}
```

## Comparison with Inheritance

Unlike object-oriented languages that rely on explicit inheritance chains, Go's structural typing simplifies the type system:

- **No Explicit Declaration:** You do not need to write `type MyType struct { } // implements Interface`.
- **Implicit Satisfaction:** The compiler checks for method presence at compile time.
- **Flexibility:** New types can satisfy existing interfaces without modifying their definitions, provided they implement the necessary methods.

## Limitations and Nuances

While structural typing offers flexibility, it is subject to Go's specific type rules:

- **Value vs. Pointer Semantics:** Methods on interface-typed values must be accessible. If a method is defined on a pointer receiver, the underlying value must be addressable (usually requiring a pointer to satisfy the interface).
- **No Private Methods:** A type can only satisfy an interface if its methods are exported (start with an uppercase letter). Private methods do not count toward interface satisfaction.
- **Type Safety:** Despite being implicit, Go's compiler enforces strict checking to ensure that all required methods exist before a type is considered compatible with an interface.

## Related Concepts

For further reading on related topics:

- [[go-built-in-collections]]
- [[go-concurrency-model]]
- [[go-error-handling]]
- [[go-generics-and-syntax]]
- [[go-memory-and-allocation]]
