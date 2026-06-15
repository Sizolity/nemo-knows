---
title: Go Fmt Conventions
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Fmt Conventions

Go code style is primarily defined by the `gofmt` tool (invoked as `go fmt`), which automatically handles formatting to ensure consistency across projects. This section covers the specific conventions enforced by this tool and the guidelines for writing clear, idiomatic code.

## Formatting Rules

The `gofmt` utility aligns columns and manages indentation using **tabs** rather than spaces. It automatically inserts semicolons where necessary to handle statements ending with newlines, except before closing braces.

When writing block comments, they should be aligned vertically if they are part of a logical group. While `gofmt` handles the mechanical aspects of formatting, developers should still prioritize writing code that is simple and avoids complex idioms borrowed from other languages like C++ or Java.

## Practical Examples

```go
package main

import "fmt"

// This comment block will be aligned vertically by gofmt if part of a group.
func main() {
    // gofmt handles semicolons automatically before newlines (except before '}')
    fmt.Printf("Hello, %s!\n", "World")
}
```
