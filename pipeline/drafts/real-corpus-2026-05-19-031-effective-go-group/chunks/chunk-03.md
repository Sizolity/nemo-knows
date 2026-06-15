---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

Chunk Context
This chunk covers control flow (switch, defer), function design (multiple returns, named results), memory allocation primitives (`new` vs `make`), and collection types (arrays, slices, maps).

Local Summary
The text details idiomatic Go patterns for error handling via multiple return values, resource management using `defer`, and the distinction between `new` (pointer to zero value) and `make` (initialized slice/map/channel). It also explains the mechanics of slices, including appending and two-dimensional structures.

Key Claims
- Functions in Go can return multiple values; this replaces C-style error codes and reference parameters.
- Named result parameters act as documentation and simplify logic when using bare `return`.
- `defer` schedules a call to run before the function returns, evaluated at execution time for arguments (LIFO order).
- `new(T)` allocates zeroed storage and returns a pointer (`*T`). Since Go 1.26, it accepts an initial value expression.
- `make(T, args)` initializes slices, maps, and channels; it does not return a pointer.
- Arrays are values (copied on assignment); slices wrap arrays and pass references to underlying data.
- Slices can be sliced (`buf[0:32]`) without copying data; appending grows capacity or reallocates if necessary.

Entities And Concepts
- `continue` with labels (loop control)
- Type switch syntax (`switch t := t.(type)`)
- Multiple return values `(n int, err error)`
- Named result parameters `(value, nextPos int)`
- `defer` execution timing and LIFO order
- `new(T)` vs `make(T)` allocation semantics
- Zero value initialization strategy
- Composite literals for struct/map/slice/arrays
- Arrays (values) vs Slices (references)
- Slice capacity (`cap`) and length (`len`)
- Two-dimensional slices (array-of-slices or slice-of-slices)

Procedures And API Details
- **Compare byte slices**: Uses two switch statements to compare elements then lengths.
  ```go
  func Compare(a, b []byte) int {
      for i := 0; i < len(a) && i < len(b); i++ {
          switch {
          case a[i] > b[i]:
              return 1
          case a[i] < b[i]:
              return -1
          }
      }
      switch {
      case len(a) > len(b):
          return 1
      case len(a) < len(b):
          return -1
      }
      return 0
  }
  ```
- **Reading numbers**: `nextInt` returns value and next position; used in a loop to scan input.
  ```go
  func nextInt(b []byte, i int) (int, int) { ... }
  for i := 0; i < len(b); {
      x, i = nextInt(b, i)
      fmt.Println(x)
  }
  ```
- **File reading with defer**: `Contents` opens a file and defers its close.
  ```go
  func Contents(filename string) (string, error) {
      f, err := os.Open(filename)
      if err != nil { return "", err }
      defer f.Close()
      // read loop
      return string(result), nil
  }
  ```
- **Appending to a slice**: Manual implementation showing capacity check and reallocation.
  ```go
  func Append(slice, data []byte) []byte {
      l := len(slice)
      if l + len(data) > cap(slice) {
          newSlice := make([]byte, (l+len(data))*2)
          copy(newSlice, slice)
          slice = newSlice
      }
      slice = slice[0:l+len(data)]
      copy(slice[l:], data)
      return slice
  }
  ```

Nuance Or Contradictions
- `new` returns a pointer to zeroed memory; since Go 1.26, it can initialize with an expression (e.g., `new(int64(300))`).
- Returning the address of a composite literal is safe because the storage survives function return.
- `make([]int)` initializes the slice structure and underlying array; `new([]int)` returns a pointer to a nil slice structure.

Candidate Wiki Hints
- Multiple Return Values in Go
- Named Result Parameters
- Defer Mechanics and Evaluation Order
- new vs make: Allocation Primitives
- Slices: Underlying Arrays, Capacity, and Appending
- Two-Dimensional Slices Strategies
