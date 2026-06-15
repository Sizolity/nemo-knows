---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
# Chunk Context
The chunk details the module-related environment variables controlling `go` command behavior, including cache paths, proxy configurations, and security settings. It concludes with a glossary defining core module terminology such as "main module," "direct dependency," and "minimal version selection."

# Local Summary
This section lists environment variables (GO111MODULE, GOMODCACHE, GOPROXY, etc.) that configure how the `go` command handles modules, proxies, and checksums. It also provides a glossary of terms used throughout the Go module system documentation.

# Key Claims
- The `go` command's module behavior is configurable via environment variables.
- `GO111MODULE` controls whether to run in module-aware mode or GOPATH mode (`off`, `on`, `auto`).
- `GOPROXY` defaults to `https://proxy.golang.org,direct`.
- `GOSUMDB` defaults to `sum.golang.org`.
- `GOVCS` defaults to using `git` and `hg` for public modules.
- Minimal version selection (MVS) determines the build list of module versions.

# Entities And Concepts
- **Environment Variables**: GO111MODULE, GOMODCACHE, GOINSECURE, GONOPROXY, GONOSUMDB, GOPATH, GOPRIVATE, GOPROXY, GOSUMDB, GOVCS, GOWORK.
- **Module Terms**: main module, direct dependency, indirect dependency, build list, canonical version, selected version, pseudo-version, release version, pre-release version.
- **Modes**: module-aware mode, GOPATH mode, workspace mode, single-module mode.
- **Proxies/Security**: module proxy, insecure download, checksum database validation.

# Procedures And API Details
- **Setting GOPROXY**: Use commas (`,`) for fallback on 404/410 errors; use pipes (`|`) to fall back on any error (including timeouts).
- **Configuring GOSUMDB**: Format is `database-name` or `database-name+<publickey> <url>`. Set to `off` to disable checksum verification.
- **Disabling Module Proxy**: Set `GOPROXY=direct` or use `GONOPROXY`/`GOPRIVATE` patterns.
- **Enabling Single-Module Mode**: Use `GOWORK=off go build .` when no `go.work` file is desired.

# Nuance Or Contradictions
- `GO111MODULE=auto` was the default in Go 1.15 and lower, whereas it defaults to `on` (or unset) in newer versions unless explicitly set to `off`.
- `GOPRIVATE` serves as a default value for both `GONOPROXY` and `GONOSUMDB`, meaning private modules are excluded from proxy checks by default if not overridden.
- Setting `GOSUMDB=off` or using `-insecure` bypasses checksum validation entirely, accepting all unrecognized modules without security guarantees.

# Candidate Wiki Hints
- **Topic: Module Environment Variables**: A reference page listing all module-related env vars with their defaults and usage examples.
- **Topic: Minimal Version Selection (MVS)**: Explanation of how MVS determines the build list and its impact on transitive dependencies.
- **Topic: Proxy Configuration**: Guide on configuring `GOPROXY`, handling errors, and using local file proxies.
