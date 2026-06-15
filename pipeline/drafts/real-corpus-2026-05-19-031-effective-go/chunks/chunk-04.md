---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Chunk Context

This chunk covers lines 1044–1546 of the source document, focusing on built-in data structures (maps), formatted printing capabilities in the `fmt` package, slice manipulation via `append`, variable initialization strategies, the role of `init` functions, and the mechanics of methods including receiver types (`value` vs. `pointer`) and interface implementation.

# Local Summary

The section details how to construct and access maps using composite literals and the "comma ok" idiom for handling missing keys. It explores the `fmt` package's printing functions, highlighting differences from C (type-based formatting), custom format verbs (`%v`, `%+v`, `%T`, `%q`), and defining `String()` methods to control output. The text explains the design of the built-in `append` function for slices, including variadic usage and array resizing semantics. Finally, it covers initialization techniques, `init` functions for package-level setup, and method receivers, illustrating how pointer receivers enable state mutation and satisfy interfaces like `io.Writer`.

# Key Claims

- Maps associate keys (any equality-defined type) with values; slices cannot be keys.
- Fetching a missing map key returns the zero value of the entry type.
- The "comma ok" idiom (`val, ok := map[key]`) distinguishes between zero values and missing entries.
- `fmt` functions use argument types to decide numeric formatting properties rather than flags.
- Custom types can define `String() string` methods to control their printed representation.
- Defining a `String()` method that calls `Sprintf` with `%s` on itself causes infinite recursion; converting to the base type prevents this.
- The built-in `append` function returns a new slice because the underlying array may be copied.
- Value methods can be called on pointers; pointer methods require pointers but are automatically rewritten when invoked on addressable values.
- An `init()` function runs after all declarations in a package are evaluated and before real execution begins.

# Entities And Concepts

- **Map**: A built-in data structure associating keys with values.
- **"Comma ok" idiom**: Multiple assignment pattern to check for map key presence.
- **`fmt` package**: Provides formatted printing functions (`Printf`, `Println`, etc.).
- **Format verbs**: `%v` (default value), `%+v` (struct fields annotated), `%T` (type name), `%q` (quoted string).
- **Variadic functions**: Functions accepting arbitrary arguments via `...interface{}`.
- **`append` function**: Built-in function to add elements or slices to a slice; returns a new slice.
- **`init()` function**: Niladic function defined per source file for package initialization.
- **Method receivers**: Value vs. pointer types determining mutability and interface satisfaction.
- **Interface implementation**: Types satisfying interfaces like `io.Writer` or `sort.Interface`.

# Procedures And API Details

- **Creating a map**:
  ```go
  var timeZone = map[string]int{
      "UTC": 0*60*60,
      "EST": -5*60*60,
  }
  ```
- **Safe map access**:
  ```go
  var seconds int
  var ok bool
  seconds, ok = timeZone[tz]
  ```
- **Defining a custom String method**:
  ```go
  func (b ByteSize) String() string {
      switch {
      case b >= YB:
          return fmt.Sprintf("%.2fYB", b/YB)
      // ... other cases
      }
      return fmt.Sprintf("%.2fB", b)
  }
  ```
- **Avoiding recursion in String method**:
  ```go
  func (m MyString) String() string {
      return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
  }
  ```
- **Appending to a slice with multiple elements**:
  ```go
  x := []int{1,2,3}
  x = append(x, 4, 5, 6)
  ```
- **Appending a slice to another slice**:
  ```go
  y := []int{4,5,6}
  x = append(x, y...)
  ```
- **Pointer receiver enabling interface satisfaction**:
  ```go
  func (p *ByteSlice) Write(data []byte) (n int, err error) {
      // implementation details
      return len(data), nil
  }
  ```

# Nuance Or Contradictions

- **Formatting flags**: Unlike C's `printf`, Go's `%d` does not take flags for signedness or size; the argument type dictates this.
- **Receiver types**: Value methods work on both values and pointers, but pointer methods strictly require pointers unless the value is addressable (in which case the compiler inserts `&`).
- **Map key types**: While structs and arrays can be keys, slices cannot because equality is undefined for them.
- **String method recursion**: Calling `Sprintf` with `%s` on a custom type within its own `String()` method causes infinite recursion; converting to the base type (e.g., `string(m)`) resolves this.

# Candidate Wiki Hints

- Map Access and Existence Checking
- Go Format Verbs Explained
- Implementing io.Writer Interface
- Avoiding Infinite Recursion in String Methods
- Variadic Functions and Append Usage
