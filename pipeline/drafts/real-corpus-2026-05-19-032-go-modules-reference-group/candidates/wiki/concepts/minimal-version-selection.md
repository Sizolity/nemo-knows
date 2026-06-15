---
title: Minimal Version Selection
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---

# Minimal Version Selection

## Overview

**Minimal Version Selection** (MVS) is an algorithm used by the Go build tool to determine the exact set of dependency versions required for a specific build. It traverses the module graph to ensure deterministic selection, choosing the minimal set of versions necessary to satisfy all direct and indirect dependencies defined in the `go.mod` file.

## Context

The algorithm operates within the broader workflow of fetching metadata and resolving dependencies, which is heavily influenced by environment variables such as `GOPROXY`. MVS ensures that when building a project, the tool selects specific versions rather than arbitrary ones, maintaining consistency across different build environments.

## Interaction with Indirect Dependencies

Since Go 1.17, indirect dependencies are recorded in a separate block within the module file to enable module graph pruning and lazy loading. MVS plays a critical role here by identifying which of these indirect dependencies are actually needed for the current build target versus those that can be excluded to optimize the build process.

## Security Implications

The deterministic nature of MVS contributes to build reproducibility, which is a key component of module security. By selecting specific versions and verifying them against checksums in the local `go.sum` file, the system prevents tampering by untrusted proxies or origin servers. The behavior of this verification process is largely controlled by environment variables like `GOSUMDB`.

## Related Concepts

- [[go-mod-directives]]
- [[indirect-dependencies-in-go-mod]]
- [[module-environment-variables]]
