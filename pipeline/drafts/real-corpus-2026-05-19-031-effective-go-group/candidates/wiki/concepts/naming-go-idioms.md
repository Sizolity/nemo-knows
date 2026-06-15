---
title: Naming Go Idioms
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Naming Go Idioms

Go naming conventions prioritize clarity, brevity, and consistency to support code written at scale. The language explicitly advises against translating idioms from C++ or Java.

## Package Names
Package names should be short, concise, evocative, lower-case, and single-word identifiers. Examples include `bufio` rather than `BufReader`.

## Type and Method Naming
- **Agent Nouns**: Interface names follow the `<method>-er` pattern (e.g., `Reader`, `Writer`).
- **Getters**: Use PascalCase methods to expose fields (e.g., `Owner()`). Avoid using `Get` prefixes or prefixed fields.
- **Multiword Identifiers**: Use MixedCaps (or mixedCaps) for multi-word identifiers; underscores are not used.

## Formatting and Comments
While distinct from naming, formatting conventions impact code readability. Issues are handled automatically by `gofmt` (or `go fmt`), which aligns columns and manages indentation using tabs. Comments should be aligned vertically if they form a block.

## Error Handling
Library functions should prioritize masking problems or working around them rather than causing the entire program to halt. Errors follow the convention of implementing the `error` interface (`type error interface { Error() string }`). Internal panics should be caught, converted into error values, and returned to callers.
