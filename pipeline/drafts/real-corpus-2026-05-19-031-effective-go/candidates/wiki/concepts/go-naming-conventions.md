---
title: Go Naming Conventions
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Naming Conventions

Naming conventions in Go are designed to ensure code clarity and consistency. The language relies heavily on implicit rules that developers follow rather than strict enforcement.

## General Principles

- **Package names** should be lowercase, single words, and concise (e.g., `bytes`). They serve as the primary namespace for imported modules.
- **Function and variable names** are typically lowercase to indicate they are exported only if their first letter is uppercase.
- **Type names** follow PascalCase conventions.

## Specific Patterns

### Getters and Accessors

Rather than using a `Get` prefix (e.g., `GetName()`), Go prefers returning capitalized field names directly (e.g., `Name`). This pattern reduces boilerplate code and makes the API more intuitive when accessing struct fields or method receivers.

### Interface Names

Interface names often follow the `<Method>-er` suffix pattern to indicate the action they support. Common examples include:
- `Reader` for reading data.
- `Writer` for writing data.
- `Closer` for closing resources.

## Formatting and Style

While not strictly naming conventions, style guides like **Effective Go** dictate that programs should be formatted by `gofmt`. This tool enforces standard indentation using tabs and avoids arbitrary line length limits, ensuring that the structure of names and code remains uniform across projects.
