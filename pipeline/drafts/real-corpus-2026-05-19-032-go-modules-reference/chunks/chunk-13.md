---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Chunk Context

The document discusses the behavior of the `go` command regarding module-related environment variables. It lists specific variables like `GO111MODULE`, `GOMODCACHE`, `GOINSECURE`, and others, detailing their functions and default behaviors. Additionally, it includes a glossary defining key terms related to Go modules, such as "build constraint," "canonical version," "direct dependency," and "minimal version selection (MVS)." The chunk concludes with navigation links and footer information from the go.dev website.

# Local Summary

This section outlines various environment variables that control the `go` command's behavior in module-aware or GOPATH modes. Key variables include `GO111MODULE` for toggling module mode, `GOPROXY` for managing proxy URLs, and `GOSUMDB` for checksum verification. The glossary provides definitions for essential concepts like "main module," "pseudo-version," and "workspace."

# Key Claims

- The `go` command supports three modes for `GO111MODULE`: off, on, and auto.
- `GOPROXY` can be set to URLs or keywords like "off" and "direct".
- `GOSUMDB` defaults to `sum.golang.org`, the Go checksum database run by Google.
- Minimal version selection (MVS) determines the versions of all modules used in a build.
- Major version suffixes are required at v2.0.0 and later.

# Entities And Concepts

- **GO111MODULE**: Controls module-aware mode vs. GOPATH mode.
- **GOMODCACHE**: Directory for storing downloaded modules.
- **GOPROXY**: List of module proxy URLs or keywords ("off", "direct").
- **GOSUMDB**: Checksum database identifier and URL.
- **GOVCS**: Controls version control tools for downloading modules.
- **GOWORK**: Enables workspace mode using a `go.work` file.
- **Build constraint**: Condition determining if a Go source file is used.
- **Minimal version selection (MVS)**: Algorithm for selecting module versions.

# Procedures And API Details

- To disable module checksum verification, set `GOPRIVATE` or `GONOSUMDB`.
- Use `GOPROXY=direct` to download modules directly from version control systems.
- Set `GO111MODULE=off` to ignore `go.mod` files and run in GOPATH mode.

# Nuance Or Contradictions

- In Go 1.15 and lower, `auto` was the default for `GO111MODULE`.
- `GOPROXY` defaults to `https://proxy.golang.org,direct`, contacting Google's mirror first.
- Pseudo-versions encode revision identifiers and timestamps for compatibility with non-module repositories.

# Candidate Wiki Hints

- **Environment Variables in Go Modules**: A comprehensive guide to configuring `go` command behavior using environment variables like `GO111MODULE`, `GOPROXY`, and `GOSUMDB`.
- **Minimal Version Selection (MVS)**: Explanation of how MVS determines module versions for builds.
- **Major Version Suffixes**: Guidelines on when and how to use major version suffixes in module paths.
