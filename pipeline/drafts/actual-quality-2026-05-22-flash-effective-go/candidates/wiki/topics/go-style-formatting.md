---
title: Go Style Formatting
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Style Formatting

Go eliminates formatting debates by providing a machine‑enforced canonical style. The `gofmt` tool (invoked as `go fmt` at the package level) reads Go source and rewrites it with standard indentation, vertical alignment, and comment formatting.

Indentation uses tabs; spaces appear only when necessary for alignment. There is no fixed line‑length limit—if a line feels too long, wrap it and indent the continuation with an extra tab.

Control structures (`if`, `for`, `switch`) omit parentheses around their conditions, and the opening brace must stay on the same line because the lexer inserts semicolons automatically. This rule makes brace style uniform across all Go code (see [[go-control-structures]]).

Go offers C‑style block comments (`/* */`) and line comments (`//`). Line comments are the norm; block comments serve mainly as package‑level documentation or to disable large sections of code. A comment placed immediately before a top‑level declaration, with no intervening blank lines, is treated as a doc comment for that declaration.
