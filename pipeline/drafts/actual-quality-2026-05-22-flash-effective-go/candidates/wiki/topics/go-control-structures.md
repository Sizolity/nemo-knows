---
title: Go Control Structures
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Control Structures

Go’s control structures blend simplicity with powerful features adapted from C, with distinct idioms that encourage readable code. Semicolons are inserted automatically by the lexer according to a rule that makes the opening brace of `if`, `for`, `switch`, and `select` mandatory on the same line as the control keyword—placing it on the next line would trigger an unwanted semicolon. For more on this and other formatting conventions, see [[go-style-formatting]].

`if` and `switch` both accept an optional short initialization statement before the condition, similar to the `for` clause. This pattern is especially common for error handling: the statement creates variables that remain scoped to the enclosing block, and the body often ends with a `return`, `break`, `continue`, or `goto`, making an `else` branch unnecessary. The result is a linear sequence of guard clauses where the happy path runs down the page and failures exit early.

```go
f, err := os.Open(name)
if err != nil {
    return err
}
// use f
```

The `for` loop is the only looping construct; it eliminates both `while` and `do`‑`while`. A `for` with a single condition behaves like `while`, and an empty `for` clause signifies an infinite loop. The classic three‑part `for` (init; condition; post) remains available.

Go’s `switch` breaks automatically from each case, so `break` is rarely needed. A `switch` with no expression is equivalent to a chain of `if else` statements. Case bodies can be comma-separated lists, and `break` and `continue` can carry labels to control flow in nested constructs.

Type switches let you branch on the dynamic concrete type of an interface value using the `type` keyword. The initialisation form is again useful to bind the typed value within each `case`.

```go
switch v := x.(type) {
case nil:   // …
case int:   // v is int
case string: // v is string
default:    // v is the same interface type
}
```

Together, these mechanisms support clear, terse control flow without the ceremony of explicit booleans or manual break statements, encouraging a style where logic reads naturally from top to bottom.
