---
title: Control Flow Go
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Control Flow Go

In Go, control flow structures are designed to be explicit and readable, avoiding implicit behaviors common in other languages like C or Java. The lexer implicitly inserts semicolons after specific tokens before newlines (except before closing braces), allowing for flexible line breaks without cluttering code with extra symbols.

## Statements

### `if`
The `if` statement supports optional initialization clauses. Its body must always be brace-delimited. Unlike C, the condition expression can include variable declarations and initializations directly within the statement.

```go
for i := 0; ; i++ {
    if x := getVal(i); x > 0 {
        println(x)
    }
}
```

### `switch`
Go's `switch` is more flexible than C's, supporting non-constant expressions and omitting the traditional fall-through logic (cases run top-to-bottom). Like `if`, the body requires braces. A common pattern involves using `fallthrough` to explicitly continue execution into the next case, or relying on default behavior where cases are distinct blocks.

```go
switch os {
case "windows":
    println("Win")
case "linux", "darwin":
    println("Unix-like")
default:
    println("Unknown")
}
```

### `for`
The `for` loop unifies C-style `for`, `while`, and infinite loops. It supports a single initialization clause, a condition (which can be omitted for an infinite loop), and an update expression. Range clauses manage iteration over arrays, slices, strings, and maps, returning the index and value (or just the key/value for maps). The underscore `_` is frequently used to discard unwanted range values.

```go
// Infinite loop
for { }

// Standard loop
for i := 0; i < len(s); i++ { }

// Range over map
for k, v := range m { }
```

## Error Handling in Control Flow

While not strictly a control structure, error handling often dictates control flow. Go does not use exceptions for normal errors. Instead, functions return an `error` interface value alongside results. The "comma ok" idiom is standard for checking existence or validity:

```go
val, ok := map[string]int{"key": 1}["missing"]
if !ok {
    // Handle missing key
}
```

Panics are reserved for unrecoverable states and halt execution. `recover` can only catch panics within deferred functions during stack unwinding. Internal panics should be caught, converted into error values, and returned to callers rather than crashing the program.

## References

- [[go-error-handling-strategies]]
- [[go-fmt-conventions]]
- [[naming-go-idioms]]
