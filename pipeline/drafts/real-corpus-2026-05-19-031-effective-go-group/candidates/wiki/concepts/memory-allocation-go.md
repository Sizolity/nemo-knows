---
title: Memory Allocation Go
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Memory Allocation Go

Memory allocation in Go is a fundamental aspect of the language's design, emphasizing simplicity and efficiency. The approach differs significantly from languages like C++ or Java, avoiding complex manual management patterns in favor of automatic garbage collection combined with specific tools for initialization.

## Initialization Tools: `make` vs `new`

Go provides two primary built-in functions for dynamic allocation, each serving a distinct purpose based on the data structure being created.

### Using `make`
The `make` function is specifically designed to initialize complex types that require internal state or memory management beyond a simple pointer dereference. It is used for:
- **Slices**: Allocates an underlying array and sets the length and capacity.
  ```go
  s := make([]int, 5) // Creates a slice of 5 zero-initialized ints
  ```
- **Maps**: Initializes the hash table required for key-value storage.
  ```go
  m := make(map[string]int)
  ```
- **Channels**: Sets up the communication pipe and buffer if specified.
  ```go
  c := make(chan int, 10)
  ```

### Using `new`
The `new` function allocates zeroed memory for a single value of a type and returns a pointer to it. It is primarily used for creating pointers to structs or other types where the internal fields do not require special initialization logic provided by `make`. Since Go 1.26, `new` also accepts an initial value expression.
```go
p := new(string) // Returns a *string pointing to ""
// With Go 1.26+
p := new(string, "hello") // Returns a *string pointing to "hello"
```

## Slices and Arrays

Understanding slices is crucial for memory management in Go. A slice is not an array itself but a descriptor (pointer, length, capacity) that wraps an underlying array. When a slice is passed to a function, only the descriptor is copied; the underlying data remains shared unless re-sliced or reallocated.

- **Range Iteration**: The `range` clause iterates over slices, strings, arrays, and maps. It allows discarding unwanted values using the blank identifier `_`.
  ```go
  for _, v := range mySlice {
      // process v
  }
  ```

## Maps

Maps in Go are hash tables backed by a pointer to an underlying structure.

- **Lookup Pattern**: A map lookup returns a two-element tuple: `(value, ok)`. The `ok` boolean indicates whether the key exists. This is known as the "comma ok" idiom.
  ```go
  value, ok := myMap["key"]
  if !ok {
      // Key does not exist
  }
  ```

- **Safe Deletion**: The `delete` function removes an entry from a map. It is safe to call even if the key does not exist; it simply performs no operation in that case.
  ```go
  delete(myMap, "key")
  ```

## Guidelines and Anti-Patterns

- **Avoid Manual Allocation**: Unlike C or Java, there is generally no need to manually `malloc` or `new` individual values unless specifically creating a pointer for a struct via `new`.
- **No Direct Translations**: Code should not attempt direct translations of idioms from C++ (e.g., using pointers everywhere) or Java. Go's memory model relies on the combination of `make`, slices, and garbage collection.
- **Simplicity**: The goal is to write clear code that avoids complex idioms. Relying on the standard library functions for allocation ensures correctness and leverages the runtime's optimizations.
