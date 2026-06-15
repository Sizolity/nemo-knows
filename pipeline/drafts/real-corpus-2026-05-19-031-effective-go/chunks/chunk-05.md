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
