---
title: Go Control Flow Idioms
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Control Flow Idioms

Go control flow idioms represent the standard patterns for managing program execution, including variable declaration, conditionals, loops, and error handling. These patterns prioritize clarity, explicitness, and safety over brevity or cleverness.

## Variable Declaration and Initialization

### Short Declarations
The `:=` operator is used to declare and initialize variables in a single statement. This is particularly common for loop counters and temporary values within function scopes.

```go
sum := 0
```

### Reassignment with Error Handling
When handling errors returned by functions, the blank identifier `_` allows developers to discard unused return values or ignore specific parts of a result tuple.

```go
err := func() error { ... }()
if err != nil {
    // handle error
}
```

## Conditionals and Loops

### Semicolon Omission
Go automatically inserts semicolons in most contexts. They are rarely visible in source code unless they appear in specific constructs like `for` loop clauses.

```go
for i := 0; i < len(slice); i++ { ... }
```

### Explicit Control Structures
Control structures such as `if`, `switch`, and `else` do not require parentheses around their conditions or expressions. This reduces visual clutter compared to languages like C or Java.

```go
if x > 0 {
    // positive
} else if x < 0 {
    // negative
} else {
    // zero
}
```

### The `else` Clause Omission
In functions that return an error, the `else` clause is often omitted after an `if` statement checking for errors. This allows successful flow to continue naturally without an explicit `else` block.

```go
if err != nil {
    return err
}
// success path continues here
```

## Error Handling Patterns

### Returning Errors
Errors are typically returned as the second value in a function's return list, implementing the `error` interface with a single `Error() string` method.

```go
func openFile(path string) (*os.File, error) { ... }
```

### Panic and Recovery
The `panic` function stops execution immediately unless recovered from inside a deferred function. It is intended for unrecoverable errors or impossible conditions. Library functions should generally avoid panicking unless during initialization, preferring to mask problems to keep programs running.

```go
func recover() {
    if r := recover(); r != nil {
        // handle panic
    }
}
defer func() { recover() }()
```

## Data Structure Manipulation

### Map Access Idioms
Maps associate keys (any equality-defined type) with values. Missing map keys return the zero value of the entry type. The "comma ok" idiom (`val, ok := map[key]`) distinguishes between zero values and missing entries.

```go
val, ok := myMap[key]
if !ok {
    // key not found
}
```

### Slice Operations
Slices are references to underlying arrays. The `append` function is used to add elements, while composite literals allow creating slices with specific initial values.

```go
slice := []int{1, 2, 3}
slice = append(slice, 4)
```

## Concurrency Patterns

### Share by Communicating
Go encourages avoiding shared mutable state in favor of passing data through channels. Goroutines are lightweight functions managed by the runtime; unbuffered channels act as rendezvous points, while buffered ones limit concurrency.

```go
done := make(chan bool)
go func() {
    // do work
    done <- true
}()
<-done
```
