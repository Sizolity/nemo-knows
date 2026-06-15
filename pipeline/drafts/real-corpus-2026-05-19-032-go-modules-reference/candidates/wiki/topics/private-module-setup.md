---
title: Private Module Setup
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Private Module Setup

Setting up private modules in Go requires configuring the module system to recognize internal paths as private, preventing them from being queried via public proxies or validated against external checksum databases. This ensures that dependencies hosted on internal servers or within a private network are resolved correctly without triggering proxy timeouts or security warnings.

## Configuration

To mark specific module paths as private, set the `GOPRIVATE` environment variable. When Go encounters a dependency matching a path listed in `GOPRIVATE`, it bypasses configured proxies (controlled by `GOPROXY`) and skips checksum verification via the public database. This allows builds to proceed using local caches or internal package servers.

## Environment Variables

The following variables control privacy and proxy behavior:

- **`GOPRIVATE`**: Specifies module paths to treat as private. Paths can be absolute or relative; matching prefixes are treated as private.
- **`GOPROXY`**: Controls the list of URLs for querying public proxies. Setting this to `direct` or `off` is often required when combined with private modules to ensure local resolution takes precedence.
- **`GONOSUMDB`**: Specifies patterns for modules that should not be checked against the public checksum database (`sum.golang.org`). This complements `GOPRIVATE` by allowing selective trust decisions.

## Security Considerations

While marking modules as private improves build reliability in internal environments, it reduces security guarantees. Private modules are not validated against the official checksum database, which increases the risk of supply chain attacks if the internal server is compromised. Therefore, internal verification mechanisms should be employed alongside `GOPRIVATE` to ensure code integrity.
