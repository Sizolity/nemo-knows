---
title: Proxy Protocol Specification
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Proxy Protocol Specification

The **Proxy Protocol Specification** outlines the operational standards for proxies within Go's dependency management ecosystem. It establishes the workflow for fetching metadata and resolving dependencies via proxies, which are configured using the `GOPROXY` environment variable. This mechanism allows the `go` command to consult proxy servers when searching the build list for matching module prefixes, preferring the longest matching path found.

## Integrity and Security

A critical component of the specification is the verification of module integrity. The system ensures that untrusted proxies or origin servers cannot tamper with dependencies by comparing SHA-256 hashes stored in the local `go.sum` file against a global checksum database (`sum.golang.org`). This process is central to the security model of Go modules.

## Environment and Configuration

The behavior of the `go` command, including module mode activation, proxy usage, and security checks, is largely controlled by environment variables with defined defaults. Key variables include:
- `GOPROXY`: Defines the list of proxies for dependency resolution.
- `GOSUMDB`: Specifies the global checksum database for verification.
- `GOVCS`: Controls version control system access.

## Compatibility and Behavior

The specification addresses legacy compatibility, detailing how Go handles pre-module repositories and GOPATH mode alongside modern modules. It also clarifies that while build commands like `go build` and `go test` can utilize vendored packages when enabled, other commands such as `go mod tidy` and `go get` continue to interact with the network or cache regardless of the vendoring status. Only vendor directories located at the main module's root are respected during these operations.
