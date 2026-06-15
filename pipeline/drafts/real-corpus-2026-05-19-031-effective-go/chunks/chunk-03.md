---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Chunk Context
This chunk covers control flow (`continue` with labels), comparison of byte slices, type switches for interface variables, multiple return values (including named results and `io.ReadFull` examples), `defer` mechanics and tracing examples, memory allocation primitives (`new` vs `make`), composite literals, array semantics, and slice fundamentals including appending and two-dimensional structures.

# Local Summary
The text explains Go's control flow quirks, demonstrates idiomatic comparison functions using switch statements, details how to inspect interface types via type switches, discusses multiple return values with named result parameters to clarify intent, explores `defer` for resource management and its evaluation timing, contrasts `new` (pointer to zero value) with `make` (initialized slice/map/channel), describes composite literals for struct/array/map creation, outlines array pass-by-value semantics versus slices as references, and introduces two-dimensional slices via arrays-of-arrays or slices-of-slices.

# Key Claims
- The `continue` statement accepts an optional label but applies only to loops.
- Comparison of byte slices can be implemented with multiple switch statements handling element-wise comparison followed by length comparison.
- Type switches use the syntax `switch t := t.(type) { ... }` to discover dynamic types; declared variables in clauses hold the corresponding type for each case.
- Go functions can return multiple values; named result parameters act as documentation and simplify code when using unadorned returns.
- `defer` schedules function calls to run immediately before the enclosing function returns, evaluated at execution time (not declaration time), allowing LIFO ordering and tracing patterns.
- `new(T)` allocates zeroed storage returning a pointer (`*T`); starting in Go 1.26 it can accept an initial value expression. `make(T, args)` creates initialized slices, maps, or channels and returns the value (not a pointer).
- Composite literals create new instances; field labels allow out-of-order initialization, and omitting fields leaves them as zero values. Returning the address of a composite literal is idiomatic because storage persists after return.
- Arrays are values; assignment copies elements, and passing an array to a function passes a copy unless a pointer is explicitly used. Slices wrap arrays, hold references, and allow modification visible to callers.
- Appending to a slice may trigger reallocation if capacity is exceeded; the built-in `append` handles this logic. Two-dimensional slices are implemented as slices of slices, allowing independent inner lengths or shared underlying arrays for efficiency.

# Entities And Concepts
- `continue` (with optional label)
- Byte slice comparison routine
- Type switch (`switch t := t.(type)`)
- Multiple return values
- Named result parameters
- `io.ReadFull` pattern with named results
- `defer` statement and evaluation timing
- Tracing routines using `defer`
- `new` vs `make` allocation primitives
- Composite literals (structs, arrays, slices, maps)
- Arrays (value semantics, type includes size)
- Slices (reference semantics, length/capacity, slicing syntax)
- Two-dimensional slices (`[][]T`)

# Procedures And API Details
### Compare byte slices
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

### Type switch example
```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)
case bool:
    fmt.Printf("boolean %t\n", t)
case int:
    fmt.Printf("integer %d\n", t)
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t)
case *int:
    fmt.Printf("pointer to integer %d\n", *t)
}
```

### Multiple return values with named results
```go
func nextInt(b []byte, pos int) (value, nextPos int) {
    for ; pos < len(b) && !isDigit(b[pos]); pos++ {
    }
    x := 0
    for ; pos < len(b) && isDigit(b[pos]); pos++ {
        x = x*10 + int(b[pos]) - '0'
    }
    return x, pos // named results used implicitly
}

func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return // uses current values of named results
}
```

### Defer usage patterns
```go
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close() // guaranteed execution on any return path

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...)
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err
        }
    }
    return string(result), nil
}

func a() {
    defer un(trace("a")) // argument evaluated when defer executes
    fmt.Println("in a")
}
```

### Allocation with new vs make
```go
// Idiomatic slice allocation:
v := make([]int, 100)

// Equivalent but less idiomatic pointer form:
p := new([]int) // *p == nil; rarely useful
```

### Composite literal for struct
```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    return &File{fd: fd, name: name} // missing fields zeroed
}
```

### Slicing a buffer
```go
n, err := f.Read(buf[0:32])
```

### Appending to a slice (manual implementation)
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

### Two-dimensional slice allocation strategies
```go
// Independent allocation per row:
picture := make([][]uint8, YSize)
for i := range picture {
    picture[i] = make([]uint8, XSize)
}

// Single allocation sliced into rows:
pixels := make([]uint8, XSize*YSize)
picture := make([][]uint8, YSize)
for i := range picture {
    picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```

# Nuance Or Contradictions
- `new` returns a pointer to zeroed storage; in Go 1.26 it can accept an initial value expression, but traditionally only zeros memory. `make` initializes slices, maps, and channels specifically because they require internal state setup before use.
- Arrays are values (copy on assignment), whereas slices are references to underlying arrays; this distinction influences function parameter design and mutation semantics.
- Returning the address of a local variable (including composite literals) is safe in Go because the storage persists after the function returns, unlike C.
- `defer` arguments are evaluated when the deferred call executes, not when the `defer` statement is encountered, enabling patterns like tracing where arguments capture current state at execution time.

# Candidate Wiki Hints
- Type Switches for Interface Variables
- Multiple Return Values and Named Result Parameters
- Defer: Resource Management and Evaluation Timing
- Allocation: new vs make
- Composite Literals in Structs, Slices, Maps, and Arrays
- Array Value Semantics vs Slice Reference Semantics
- Slicing Buffers and Appending to Slices
- Two-Dimensional Slices: Independent vs Shared Underlying Arrays
