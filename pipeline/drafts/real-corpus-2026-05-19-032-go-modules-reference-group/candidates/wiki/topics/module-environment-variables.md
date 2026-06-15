---
title: Module Environment Variables
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Module Environment Variables

## Overview

Environment variables play a central role in controlling the behavior of the `go` command, particularly regarding module discovery, proxy usage, and security verification. These variables allow users to customize how dependencies are fetched and verified without modifying source code or configuration files directly.

## Core Variables

### GOPROXY
The `GOPROXY` environment variable defines the list of proxy servers used for fetching modules. It supports comma-separated URLs and specific keywords like `direct` or `off`. This variable determines whether the build tool contacts a global proxy, local caches, or fetches directly from origin servers.

- See [[go-mod-directives]] for how proxies interact with other directives.
- The proxy protocol is defined by the [[proxy-protocol-specification]].

### GOSUMDB
This variable specifies the checksum database used to verify module integrity. By default, it points to `sum.golang.org`. Users can set this to `off` to disable security checks entirely or provide a custom URL for private repositories.

- Security and cache behavior is further discussed in [[module-security-and-cache]].

### GOVCS
Used when fetching modules from version control systems (e.g., Git, Mercurial). It specifies the VCS protocol to use for resolving pseudo-versioned dependencies.

## Behavior Defaults

If not explicitly set, the `go` command uses sensible defaults:
- **GOPROXY**: Typically points to the official Go proxy or is left open depending on OS and network configuration.
- **GOSUMDB**: Defaults to the global checksum database for security verification.
- **GOVCS**: Defaults to using Git by convention when supported.

These defaults ensure secure and reliable module resolution out of the box, while allowing flexibility for specialized environments.

## Related Topics

For further reading on how these variables interact with other system settings:
- [[go-mod-versioning]]
- [[minimal-version-selection]]
- [[indirect-dependencies-in-go-mod]]
