---
title: Go Formatting Conventions
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Formatting Conventions

Go source formatting is handled automatically by the `gofmt` tool (exposed as `go fmt` for package‑level operations). Instead of prescribing a detailed style guide, the language relies on the machine to produce a single, standard layout for every `.go` file. `gofmt` standardises indentation, vertical alignment, and comment placement—for example, aligning per‑field comments inside struct declarations so that developers never need to manually line up columns. Because all code in the standard library is formatted with `gofmt`, the community shares a consistent visual style without spending time on formatting debates.
