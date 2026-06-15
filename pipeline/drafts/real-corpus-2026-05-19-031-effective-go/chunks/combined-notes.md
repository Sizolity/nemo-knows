## chunk-01

---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## Chunk Context
This chunk introduces the "Effective Go" document, a style and language idiom source for the Go programming language. It contains metadata confirming the source URL (https://go.dev/doc/effective_go), retrieval date (2026-05-18), and categorization within the web corpus under the "Go" tag.

## Local Summary
The text serves as a header and metadata block for the "Effective Go" guide, establishing it as an authoritative resource on Go coding idioms. It confirms the document's origin from the official Go programming language documentation site.

## Key Claims
- The document titled "Effective Go" is a style and language idiom source.
- The content is hosted at `https://go.dev/doc/effective_go`.
- The specific chunk was retrieved on 2026-05-18 with a successful fetch status.

## Entities And Concepts
- **Effective Go**: A guide detailing idiomatic usage of the Go programming language.
- **Style and language idiom source**: The classification of the document's content type.
- **Go Programming Language**: The software development environment to which the style guide applies.

## Procedures And API Details
No specific procedures or API details are described in this chunk; it is purely introductory metadata.

## Nuance Or Contradictions
None observed in this introductory segment.

## Candidate Wiki Hints
- **Page: Effective Go** (Draft)
  - *Purpose*: Create a summary page linking to the official documentation and listing key style guidelines introduced in subsequent chunks.

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## Chunk Context
Lines 27-550 of `raw/web/corpus-2026-05-18/031-effective-go.md`. The text covers the "Effective Go" documentation, specifically focusing on formatting conventions (indentation, line length), naming conventions (packages, getters, interfaces), syntax details (semicolons, control structures like `if`, `for`, `switch`), and idiomatic code patterns (error handling, blank identifiers).

## Local Summary
This section of the document serves as a guide to writing clear, idiomatic Go code. It emphasizes that while formatting is often automated by `gofmt`, understanding its rules (tabs for indentation, no line length limits) is crucial. The text details naming conventions such as using MixedCaps for multi-word names, avoiding unnecessary getters (preferring capitalized field access), and adhering to standard interface suffixes like `-er`. It also explains the automatic semicolon insertion rule and how it influences code layout, particularly regarding control structures. Furthermore, it outlines the syntax of `if`, `for`, and `switch` statements, highlighting idioms like omitting `else` after error checks and using short declarations with `:=`.

## Key Claims
- Go programs should be formatted by `gofmt` to ensure consistency; manual formatting is discouraged unless necessary.
- Package names should be lowercase, single words, and concise (e.g., `bytes`, not `byteutil`).
- Getters should return capitalized field names (e.g., `Owner`) rather than using the prefix `Get`.
- Interface names often follow the pattern of `<Method>-er` (e.g., `Reader`, `Writer`).
- Semicolons are automatically inserted by the lexer unless they appear in specific contexts like `for` loop clauses.
- Control structures (`if`, `switch`) do not require parentheses around their expressions or conditions.
- The `else` clause is often omitted after an `if` statement that returns an error, allowing successful flow to continue naturally.
- Short declarations (`:=`) allow reassignment of existing variables in the same scope (e.g., `err`).
- The blank identifier `_` is used to ignore values in range loops when not needed.

## Entities And Concepts
- **gofmt**: A tool that automatically formats Go source code.
- **Package Name**: Should be short, concise, evocative, and lowercase (e.g., `bytes`).
- **Getter/Setter**: Methods for accessing/modifying fields; getters should not use the `Get` prefix.
- **Interface Naming**: Conventionally uses `-er` suffix for agent nouns (e.g., `Reader`).
- **Semicolon Insertion**: A lexer rule that inserts semicolons automatically at newlines under certain conditions.
- **Blank Identifier**: The underscore `_` used to discard values in loops or assignments.
- **Control Structures**: `if`, `for`, `switch`, `select`.
- **Short Declaration**: The `:=` operator for declaring and initializing variables.

## Procedures And API Details
- **Formatting Rules**:
  - Use tabs for indentation.
  - No line length limit; wrap lines if they feel too long.
  - Do not align comments manually; let `gofmt` handle it.
- **Naming Conventions**:
  - Exported names use uppercase first letter (e.g., `Owner`).
  - Avoid `Get` prefix for getters (use `Owner` instead of `GetOwner`).
  - Use MixedCaps for multi-word names (e.g., `MixedCaps`).
- **Code Patterns**:
  - Omit `else` when the `if` body returns an error.
  - Use `:=` to declare and assign variables in a single statement.
  - Use `_` in range loops to ignore unused values (e.g., `for _, value := range array`).
- **Control Structure Syntax**:
  - `if` statements do not require parentheses: `if condition { }`.
  - `switch` can be used as an if-else chain without an expression.
  - Labels are used with `break` to exit nested loops or switches (e.g., `break Loop`).

## Nuance Or Contradictions
- **Semicolon Visibility**: Although the formal grammar uses semicolons, they rarely appear in source code due to automatic insertion by the lexer. However, they are still required in specific places like `for` loop clauses and multiple statements on a line.
- **Getter Naming**: While getters are not idiomatic by default, providing them is allowed if appropriate; however, the naming convention should avoid the `Get` prefix.
- **Interface Names**: Standard interfaces like `Reader`, `Writer` have canonical signatures; creating methods with similar names but different signatures can cause confusion.

## Candidate Wiki Hints
- **Effective Go Formatting Guide**
- **Go Naming Conventions (Packages, Getters, Interfaces)**
- **Semicolon Insertion Rules in Go**
- **Idiomatic Control Structures in Go (`if`, `for`, `switch`)**
- **Blank Identifier Usage in Range Loops**

## chunk-03

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

## chunk-04

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

## chunk-05

---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## Chunk Context
Lines 1548–2017 of `raw/web/corpus-2026-05-18/031-effective-go.md`, covering the transition from `Sequence` printing logic to advanced interface patterns, type switches vs. assertions, blank identifier usage, function receivers, and embedding.

## Local Summary
This section demonstrates how Go leverages interfaces and type conversions for polymorphism without explicit subclassing. It covers:
- Using multiple types (`Sequence`, `sort.IntSlice`, `[]int`) to share behavior.
- Type switches vs. type assertions, including the "comma, ok" idiom.
- Returning interface values from constructors (e.g., `hash.Hash32`).
- Defining methods on any type, including functions (`HandlerFunc`).
- The blank identifier `_` for discarding values or side-effect imports.
- Compile-time interface checks via `var _ Interface = Value`.
- Embedding interfaces and structs to combine behavior without forwarding methods.

## Key Claims
- **Polymorphism via types**: A value can be converted among multiple types (concrete and interface) to delegate parts of a job, which is unusual but effective in Go.
- **Type switches vs. assertions**: Type switches handle concrete values or interface-to-interface conversions; type assertions extract a specific type from an interface.
- **Blank identifier safety**: The blank identifier `_` safely discards values (e.g., unused function return values) and imports, preventing compile errors during development.
- **Interface implementation is implicit**: A type satisfies an interface if it implements the required methods; no explicit declaration is needed.
- **Compile-time contract enforcement**: `var _ Interface = Value` forces the compiler to verify that a type satisfies an interface, even when static conversions are absent.
- **Embedding avoids forwarding**: Embedding structs or interfaces within another struct/interface automatically promotes their methods and satisfies combined interfaces without manual forwarding functions.

## Entities And Concepts
- `Sequence`, `sort.IntSlice`, `[]int`: Types used to demonstrate multiple type delegation for printing and sorting.
- `Stringer` interface: Defines `String() string`.
- Type switch vs. type assertion syntax (`value.(typeName)`).
- "Comma, ok" idiom: `str, ok := value.(string)`.
- `hash.Hash32` interface: Returned by constructors like `crc32.NewIEEE`.
- `crypto/cipher.Block` and `Stream` interfaces: Block and streaming cipher abstractions.
- `http.Handler`, `http.ResponseWriter`, `http.Request`: HTTP handler ecosystem types.
- `Counter` struct vs. `Counter int`: Demonstrating method receivers on values.
- `Chan` type with method: Using a channel as a receiver to send notifications.
- `HandlerFunc`: Adapter type allowing functions to act as `http.Handler`.
- Blank identifier `_`: Used for discarding values, unused imports, and interface contract checks.
- `json.Marshaler` interface: Runtime-checked marshaling capability.
- `io.Reader`, `io.Writer`, `io.ReadWriter`: I/O operation interfaces.
- Embedding in structs/interfaces: Combining behavior without explicit field names or forwarding methods.

## Procedures And API Details
### Type Switch vs. Assertion
```go
// Type switch
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}

// Type assertion with "comma, ok"
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
}
```

### Function as HTTP Handler
```go
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}

// Register using HandlerFunc adapter
http.Handle("/args", http.HandlerFunc(ArgServer))
```

### Interface Contract Check
```go
var _ json.Marshaler = (*RawMessage)(nil)
```

### Embedding Interfaces
```go
type ReadWriter interface {
    Reader
    Writer
}
```

### Embedding Structs
```go
type ReadWriter struct {
    *Reader // embeds bufio.Reader
    *Writer // embeds bufio.Writer
}
```

## Nuance Or Contradictions
- **Implicit vs. explicit interface satisfaction**: Go does not require types to declare they implement an interface; implementation is determined by method presence. However, runtime checks (e.g., in `encoding/json`) may still assert this property.
- **Blank identifier misuse warning**: Discarding errors without checking (`_, _ := os.Stat(path)`) is flagged as bad practice unless intentional for debugging or temporary work-in-progress code.
- **Embedding limitations**: Only interfaces can be embedded within other interfaces; struct embedding requires pointers to avoid nil dereference and implicitly promotes methods.

## Candidate Wiki Hints
- **Page: Blank Identifier in Go**
  Summarizes the `_` idiom for discarding values, unused imports, and interface contract checks.
- **Page: Embedding Interfaces and Structs**
  Explains how embedding combines behavior without forwarding methods, using `io.Reader`, `io.Writer`, and `bufio.ReadWriter` as examples.
- **Page: Function Receivers and Adapter Types**
  Details the creation of `HandlerFunc` and other adapter types that enable functions to satisfy interfaces like `http.Handler`.
- **Page: Compile-Time Interface Contracts**
  Describes the `var _ Interface = Value` pattern for enforcing implicit interface implementation at compile time.

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

Chunk Context
This chunk covers advanced Go topics including embedding vs. subclassing, concurrency models ("Share by communicating"), goroutine mechanics, channel usage (buffered/unbuffered), parallelization strategies, garbage-collected buffer management via channels, and error handling conventions using the `error` interface and `panic`.

Local Summary
The section explains how embedding types allows method forwarding while keeping the receiver on the embedded type. It contrasts this with subclassing behavior. Concurrency principles are introduced, emphasizing that shared memory should be avoided in favor of communication via channels. Goroutines are described as lightweight functions managed by the Go runtime. Channels provide synchronization and data exchange; unbuffered channels act as rendezvous points, while buffered ones can limit concurrency (semaphores). Parallelization across CPU cores is achieved by launching goroutines for independent tasks and collecting completion signals. Finally, error handling patterns using the `error` interface and panic recovery are discussed.

Key Claims
- Embedding a type makes its methods available on the outer type, but the receiver remains the embedded type.
- Go encourages avoiding shared mutable state in favor of passing data through channels to prevent data races.
- Goroutines are lightweight functions that run concurrently within the same address space; they are multiplexed onto OS threads.
- Unbuffered channels synchronize senders and receivers, while buffered channels can limit throughput or act as semaphores.
- Errors in Go are typically returned as second return values of type `error`, which implements an interface with a single `Error() string` method.
- The `panic` function causes the program to stop immediately unless recovered from; it is intended for unrecoverable errors or impossible conditions.

Entities And Concepts
- Embedding (Go composite literals, field embedding)
- Goroutine (`go` keyword, lightweight threads)
- Channel (`chan`, buffered/unbuffered, send/receive operations)
- Semaphore pattern using channels
- Parallelization via CPU core detection (`runtime.NumCPU`)
- Error handling (`error` interface, `os.PathError`)
- Panic and recovery (though only panic definition is covered here)

Procedures And API Details
```go
type Job struct {
    Command string
    *log.Logger
}

func (job *Job) Printf(format string, args ...interface{}) {
    job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}

go func() {
    time.Sleep(delay)
    fmt.Println(message)
}()

c := make(chan int)
go func() {
    list.Sort()
    c <- 1
}()
<-c

func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}

func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit
}

type Request struct {
    args   []int
    f      func([]int) int
    resultChan chan int
}

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
        v[i] += u.Op(v[i])
    }
    c <- 1
}

var numCPU = runtime.GOMAXPROCS(0)

func client() {
    for {
        var b *Buffer
        select {
        case b = <-freeList:
            // Got one; nothing more to do.
        default:
            b = new(Buffer)
        }
        load(b)
        serverChan <- b
    }
}

func server() {
    for {
        b := <-serverChan
        process(b)
        select {
        case freeList <- b:
            // Buffer on free list; nothing more to do.
        default:
            // Free list full, just carry on.
        }
    }
}

func CubeRoot(x float64) float64 {
    z := x / 3
    for i := 0; i < 1e6; i++ {
        prevz := z
        z -= (z*z*z - x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}

type error interface {
    Error() string
}

type PathError struct {
    Op   string // "open", "unlink", etc.
    Path string // The associated file.
    Err  error  // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```

Nuance Or Contradictions
- Before Go 1.22, using a loop variable inside a goroutine launched in a loop was buggy because the loop variable is shared across all goroutines unless captured carefully.
- While channels are first-class values and can be passed around, care must be taken not to create unbounded resource consumption when spawning goroutines per request without limiting concurrency.
- The distinction between concurrency (structuring as independent components) and parallelism (executing computations in parallel for efficiency) is emphasized; Go supports concurrency but does not guarantee parallel execution unless explicitly orchestrated.

Candidate Wiki Hints
- Embedding vs Subclassing
- Concurrency Model: Share by Communicating
- Goroutines and Channels Basics
- Channel Types: Buffered vs Unbuffered
- Semaphore Pattern with Channels
- Parallelizing Workloads Using CPU Cores
- Error Handling in Go
- Panic Usage and Recovery

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## Chunk Context
This chunk concludes the "Effective Go" chapter, focusing on advanced error handling patterns involving `panic` and `recover`, followed by a complete example of a QR code web server using the `html/template` package. The text ends with footer navigation links from go.dev.

## Local Summary
The section first advises that library functions should avoid panicking unless during initialization, preferring to mask or work around problems to keep programs running. It then details how to use `recover` inside deferred functions to catch panics and allow a goroutine to exit cleanly without crashing the whole server. A specific pattern is shown where `panic` reports parse errors, which are caught by a recovery block that converts them into return values (`nil` and an error). Finally, a complete Go program is presented that acts as a web re-server for generating QR codes from text inputs using a template engine.

## Key Claims
- Library functions should avoid panicking; if a problem can be masked or worked around, it is better to let the program continue running rather than taking it down.
- `recover` is only useful inside deferred functions because it stops stack unwinding and returns the argument passed to `panic`.
- A recovery pattern allows a failing goroutine to log its failure and exit cleanly without disturbing other executing goroutines.
- Library routines that use `panic` and `recover` can be called from deferred code without failing, provided `recover` is not invoked directly outside of a defer block.
- Panics should generally be confined within a package; internal panics can be converted into error values for external consumers (e.g., the `Compile` function).
- The `html/template` package allows dynamic rewriting of HTML text by substituting data items, with automatic escaping to ensure safety.

## Entities And Concepts
- **panic**: Immediately stops execution of the current function and begins unwinding the stack; used for run-time errors like index out of bounds or failed type assertions.
- **recover**: Built-in function that regains control of a goroutine when called from a deferred function, returning the value passed to `panic` or `nil`.
- **deferred functions**: Only code running during stack unwinding; the only place where `recover` is effective.
- **Error interface**: Satisfied by custom types like `type Error string` to report parse errors via panic recovery.
- **html/template**: Package for executing templates with data items, providing automatic escaping and conditional rendering (e.g., `{{if .}}`).
- **QR code generation**: A web service example using `chart.apis.google.com` to generate QR codes from text input via a form handler.

## Procedures And API Details
### Panic Recovery Pattern
1. Define a custom error type satisfying the `error` interface (e.g., `type Error string`).
2. Create an internal method that panics with this error type (e.g., `func (regexp *Regexp) error(err string)`).
3. Use a deferred function in functions like `Compile` to catch the panic:
   ```go
   defer func() {
       if e := recover(); e != nil {
           regexp = nil // Clear return value.
           err = e.(Error) // Re-panic if not a parse error.
       }
   }()
   ```
4. Check the type of recovered error to distinguish between expected parse errors and unexpected run-time failures.

### Web Server Implementation
1. Define a `main` function that parses flags, binds a handler (`QR`) to the root path, and starts `http.ListenAndServe`.
2. Use `flag.String` to set default ports (e.g., `:1718`).
3. Load an HTML template using `template.Must(template.New("qr").Parse(templateStr))`.
4. Implement the handler `QR(w http.ResponseWriter, req *http.Request)` which executes the template with the form value `req.FormValue("s")`.
5. Define `templateStr` containing HTML and Go expressions (`{{if .}}`, `{{.}}`) to conditionally render QR images based on input data.

## Nuance Or Contradictions
- While `panic` is used for exceptional conditions, the text notes that if an unexpected failure occurs (e.g., index out of bounds) within a recovery block, the type assertion `err = e.(Error)` will fail, causing a new panic and continuing stack unwinding. This ensures that only intended parse errors are handled gracefully.
- The re-panic idiom changes the panic value if an actual error occurs; both original and new failures appear in crash reports, which is usually sufficient but can be filtered if only the original value is desired.

## Candidate Wiki Hints
- **Topic: Panic Recovery Pattern** – A reusable source for implementing robust error handling in Go libraries where panics are internal signals for specific conditions (like parse errors) and external consumers expect `error` values.
- **Topic: Deferred Recover Usage** – Notes on the constraint that `recover` must be called from a deferred function to effectively stop stack unwinding.
- **Topic: Internal vs External Errors** – Guidance on keeping panic handling within a package and converting internal panics to returnable errors for public APIs.

