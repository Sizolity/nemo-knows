---
title: Module Security And Cache
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Module Security And Cache

The Go Modules system manages dependency security and caching through a combination of checksum verification, proxy protocols, and local storage mechanisms. Integrity is ensured by comparing SHA-256 hashes stored in the local `go.sum` file against a global checksum database defined via the `GOSUMDB` environment variable. This process prevents tampering by untrusted proxies or origin servers during the fetching of metadata.

## Security Verification

Module security relies on cryptographic verification of downloaded packages. The build tool verifies integrity by comparing hashes in the local `go.sum` file against a global checksum database. This mechanism ensures that code retrieved from the network has not been altered by intermediate nodes or malicious actors.

- **Checksum Database:** Controlled via the `GOSUMDB` environment variable, defaulting to `sum.golang.org`.
- **Verification Logic:** SHA-256 hashes are computed for downloaded modules and compared against the stored values before installation.

## Proxy Configuration

The system utilizes proxies to manage network requests and enforce security policies at the transport level. Proxy behavior is driven by the `GOPROXY` environment variable, which defines a list of URLs (comma-separated) that the `go` command will consult for fetching metadata and packages. This allows organizations to route traffic through internal caches or secure gateways.

- **Protocol Specification:** The interaction between clients and proxies follows the `proxy-protocol-specification`, ensuring consistent behavior across different implementations.
- **Environment Control:** Behavior is largely controlled by environment variables with defined defaults, such as `GOPROXY` and `GOVCS`.

## Cache Management

The Go toolchain maintains a local cache to accelerate builds and reduce network load. This cache stores downloaded modules and their metadata (e.g., `go.sum`, `go.mod`). Commands like `go build` and `go test` utilize this cache, while others like `go mod tidy` interact with the network as needed unless vendoring is explicitly enabled.

- **Local Storage:** The cache resides in the module-aware file system, allowing for deterministic selection of versions via algorithms like Minimal Version Selection (MVS).
- **Lazy Loading:** Since Go 1.17, indirect dependencies are recorded separately to enable lazy loading and graph pruning, optimizing cache usage.
- **Vendoring Interaction:** When vendoring is enabled, `go build` uses local copies from the vendor directory instead of the network/cache for the main module's root, though other commands continue to use the cache regardless of vendoring status.

## Environment Variables

The core behavior regarding security and caching is driven by environment variables:

- **GOPROXY:** Defines the proxy servers to consult.
- **GOSUMDB:** Defines the checksum database for security verification.
- **GOVCS:** Configures the version control system used for fetching source code when needed.
