---
title: Go Allocation New Vs Make
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Allocation New Vs Make

In Go, choosing between `new` and `make` depends on the type of data structure you are initializing. Both functions return pointers, but they operate differently regarding memory initialization and supported types.

## Core Distinction

The primary difference lies in what they allocate and initialize:
- `new(T)` allocates zeroed storage for a value of type `T` and returns a pointer to it (`*T`). It works with any type.
- `make(T, args)` creates initialized slices, maps, or channels. These data structures require initialization (e.g., setting length) before use, which `make` handles automatically.

## When to Use `new`

Use `new` when you need a pointer to a zeroed value of any type. This is common for:
- Allocating pointers to structs.
- Initializing basic types like integers or booleans where the default zero value is acceptable.
- Creating pointers to functions.

**Example:**
```go
// Allocates a pointer to an int initialized to 0
p := new(int)
```

## When to Use `make`

Use `make` exclusively for dynamic collections: slices, maps, and channels. Attempting to use `new` on these types results in runtime panics because the underlying storage is not initialized with the correct capacity or length.

**Example:**
```go
// Allocates an empty slice with a default capacity (implementation-defined)
s := make([]int, 0)

// Allocates a map with zero initial size
m := make(map[string]int)

// Allocates an unbuffered channel
c := make(chan int)
```

## Memory Semantics

- **`new`** returns a pointer to a memory location filled with the type's default zero value (e.g., `0` for integers, `nil` for slices/maps). If you use `new([]int)`, you get a pointer to a nil slice, not an empty slice.
- **`make`** returns a pointer to a fully initialized data structure ready for immediate use.

## Summary Table

| Function | Returns         | Supported Types          | Initialization Behavior           |
|----------|-----------------|--------------------------|-----------------------------------|
| `new`    | `*T`            | Any type                 | Allocates zeroed storage          |
| `make`   | `interface{}`   | Slices, maps, channels   | Initializes capacity/length/buffer |

## Related Concepts

- [[go-data-structures-slices-maps]]
- [[go-defer-mechanics]]
