---
title: Rust API Guidelines
kind: source
created: 2026-05-18
updated: 2026-05-18
sources:
  - raw/web/curated-web-corpus-2026-05-18.md
  - https://rust-lang.github.io/api-guidelines/
tags: [rust, web-corpus]
confidence: medium
---

# Rust API Guidelines

## Fetch Metadata

- Corpus item: 42
- Category: Rust
- Source URL: https://rust-lang.github.io/api-guidelines/
- Final URL: https://rust-lang.github.io/api-guidelines/
- Retrieved: 2026-05-18
- Content-Type: text/html; charset=utf-8
- Fetch status: ok
- Test value: Tests design guidance extraction.
- Fetched page title: About - Rust API Guidelines

## Retrieved Text

1. About
2. Checklist
3. 1. Naming
4. 2. Interoperability
5. 3. Macros
6. 4. Documentation
7. 5. Predictability
8. 6. Flexibility
9. 7. Type safety
10. 8. Dependability
11. 9. Debuggability
12. 10. Future proofing
13. 11. Necessities
14. External links

(BUTTON) (BUTTON)
* (BUTTON) Light (default)
* (BUTTON) Rust
* (BUTTON) Coal
* (BUTTON) Navy
* (BUTTON) Ayu

(BUTTON)

Rust API Guidelines

____________________

Rust API Guidelines

This is a set of recommendations on how to design and present APIs for
the Rust programming language. They are authored largely by the Rust
library team, based on experiences building the Rust standard library
and other crates in the Rust ecosystem.

These are only guidelines, some more firm than others. In some cases
they are vague and still in development. Rust crate authors should
consider them as a set of important considerations in the development
of idiomatic and interoperable Rust libraries, to use as they see fit.
These guidelines should not in any way be considered a mandate that
crate authors must follow, though they may find that crates that
conform well to these guidelines integrate better with the existing
crate ecosystem than those that do not.

This book is organized in two parts: the concise checklist of all
individual guidelines, suitable for quick scanning during crate
reviews; and topical chapters containing explanations of the guidelines
in detail.

If you are interested in contributing to the API guidelines, check out
contributing.md and join our Gitter channel.
