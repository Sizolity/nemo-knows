---
title: Go Data Structures Slices Maps
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Data Structures Slices Maps

In Go, **slices** and **maps** are fundamental built-in data structures that differ significantly in their semantics regarding values, references, and key usage. Understanding these distinctions is crucial for writing idiomatic and efficient code.

## Slices

A slice in Go is a dynamic-sized view into an underlying array. Unlike arrays, which are values copied on assignment, slices act as **references** to their underlying data. This means modifying a slice can affect the original array it references.

### Allocation
When creating a new slice, use `make` rather than `new`. The `make` function initializes the slice's underlying array and sets its length and capacity, whereas `new` simply allocates zeroed storage for a pointer type.

```go
s := make([]int, 5) // Creates a slice of 5 ints initialized to 0
```

### Manipulation
Slices support efficient manipulation techniques such as `append`, which grows the slice by adding elements to the end. Composite literals can also be used to create slices from existing ones or specific values.

## Maps

Maps associate keys with values in Go. They are implemented as hash tables and offer O(1) average-time complexity for lookups, insertions, and deletions.

### Key Constraints
A critical semantic difference between slices and maps is that **maps cannot be used as map keys**. Only types that define equality can serve as keys; slices do not define equality in the standard sense, so they are prohibited as keys. Missing map keys return the zero value of the entry type.

```go
m := make(map[string]int)
val, ok := m["key"] // The "comma ok" idiom distinguishes missing entries from zero values
if !ok {
    val = 0
}
```

### Memory and Semantics
While slices reference underlying arrays, maps are distinct entities. When iterating over a map, the order of keys is not guaranteed to be insertion order (prior to Go 1.22).

## Concurrency Considerations

Both slices and maps can be used in concurrent contexts, but care must be taken with shared mutable state.
- **Slices**: Since slices reference underlying arrays, concurrent modification requires synchronization (e.g., using `sync.Mutex`) unless the slice is read-only or accessed via channels.
- **Maps**: Similar to slices, map writes require protection. It is often recommended to use a `sync.Map` for concurrent access patterns to avoid locking contention in simple cases, though standard maps are more performant when protected by a single lock.

## Style and Formatting

Adhering to idiomatic Go style improves code readability:
- Use `gofmt` to format code; manual formatting is discouraged.
- When declaring variables inside loops or scopes where only one variable is needed, short declarations (`:=`) are preferred.
- The blank identifier `_` should be used to ignore values in range loops or discard unused return values from functions like `make` or map lookups when the value is irrelevant.

## References

For further reading on idiomatic usage and advanced patterns:
- [[go-allocation-new-vs-make]]
- [[go-concurrency-model]]
- [[go-defer-mechanics]]
- [[go-error-handling-panic-recovery]]
- [[go-interface-patterns]]
