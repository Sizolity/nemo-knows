---
title: Go Data Allocation
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Data Allocation

Go provides two built-in allocation primitives with distinct roles: `new` and `make`.

`new(T)` allocates memory for a new variable of type `T`, sets it to the zero value, and returns a pointer (`*T`). Starting with Go 1.26, `new` may also accept an initial value expression. Because many types are designed so that their zero value is immediately usable (for example, a `sync.Mutex` or `bytes.Buffer`), `new` often gives back a ready-to-use object without further initialisation.

`make` is reserved for slices, maps, and channels. It returns an **initialised value** of type `T` (not a pointer), because these types internally reference data structures that must be set up before use. For instance, `make([]int, 10, 100)` allocates a backing array of 100 ints and returns a slice descriptor with length 10 and capacity 100, pointing to the first 10 elements. In contrast, `new([]int)` yields a pointer to a nil slice—useless until an actual slice structure is created.

Slices, maps, and channels are reference types: their variables hold descriptors that point to underlying data. Assigning one slice to another copies only the descriptor, so both refer to the same backing array. When a function receives a slice argument, modifications to its elements are visible to the caller, much like passing a pointer to the underlying array.

Composite literals offer a concise syntax for constructing these allocation‑backed values (arrays, slices, maps, structs) without calling `new` or `make` directly. Combined with the fact that `make` initialises internal state while `new` merely zeroes memory, this two‑pronged design ensures that freshly allocated data is always in a safe, predictable state.
