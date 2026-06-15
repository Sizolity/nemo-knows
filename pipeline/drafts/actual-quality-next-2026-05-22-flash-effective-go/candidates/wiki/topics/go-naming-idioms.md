---
title: Go Naming Idioms
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Naming Idioms

Names in Go carry semantic weight: an identifier whose first character is uppercase is exported and visible outside its package, while a lowercase start keeps it unexported. This single rule underpins most naming conventions in the language.

Package names are short, lowercase, single-word, and avoid underscores or mixed caps. They should be concise and evocative, enabling the import path to serve as a natural prefix—callers can refer to `bytes.Buffer` without repeating the package name in the symbol. The package name usually matches the directory that holds its source files.

Getter methods for unexported fields omit the `Get` prefix. A field named `owner` yields an exported getter `Owner()`, not `GetOwner()`, preserving directness and consistency with the capital‑case export rule.

One‑method interfaces conventionally end with `‑er`, such as `Reader` and `Writer`. This “doer” pattern immediately signals the single responsibility of the interface.

When a name consists of multiple words, `MixedCaps` or `mixedCaps` (camelCase) is used rather than underscores, making identifiers both readable and uniform across variables, types, and functions.
