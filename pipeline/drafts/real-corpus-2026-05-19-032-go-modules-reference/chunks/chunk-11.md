---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
# Chunk Context

This chunk details the mechanics of module distribution via `.zip` files, security constraints on version control usage (`GOVCS`), and configuration for accessing private modules (proxies, direct VCS access, authentication). It covers file size limits, path constraints within zip archives, and environment variables like `GOPRIVATE`, `GOPROXY`, and `GONOSUMDB`.

# Local Summary

The Go toolchain distributes modules as authenticated `.zip` files containing only the module root contents (excluding vendor directories or nested modules). The `go` command enforces strict file constraints to ensure cross-platform compatibility and security, such as ignoring symbolic links and limiting zip sizes to 500 MiB. To manage security risks associated with untrusted version control servers, the `GOVCS` environment variable restricts allowed VCS tools (defaulting to `git` and `hg` for public paths) unless overridden. For private modules, administrators can configure `GOPROXY` and `GONOSUMDB` to use internal proxies or bypass checksum verification, while `GOPRIVATE` marks module paths as private to avoid external proxy lookups. Authentication for private resources is handled via `.netrc` files or embedded credentials in URLs, with support for non-interactive Git configurations.

# Key Claims

- Module zip files are authenticated before extraction; they do not include vendor directories or nested modules (subdirectories with `go.mod`).
- If a module lacks a `LICENSE` file in its root, the `go` command copies it from the repository root if present in the same revision.
- Symbolic links are excluded from module zip files to ensure portability across operating systems.
- The default behavior for public modules is to download via proxy (`proxy.golang.org`) rather than directly cloning version control repositories.
- `GOVCS` allows fine-grained control over which VCS tools (e.g., `git`, `hg`) are permitted for specific module paths.
- Private modules can be accessed without a private proxy by setting `GOPRIVATE` and configuring direct VCS access, potentially requiring non-interactive authentication.

# Entities And Concepts

- **Module Zip Files**: Distributed artifacts containing module contents; created, downloaded, and extracted automatically by the `go` command.
- **GOVCS**: Environment variable controlling allowed version control systems for downloading code.
- **GOPRIVATE**: Glob patterns marking module paths as private to bypass public proxies and checksum databases.
- **GOPROXY**: List of module proxy URLs; used to download modules sequentially, falling back to direct VCS access if unavailable.
- **GONOSUMDB**: Patterns for modules not to be checked against the public checksum database (`sum.golang.org`).
- **.netrc**: File used to store credentials for HTTP basic authentication with proxy servers.
- **Direct Access**: Configuration mode where `go` downloads private modules directly from VCS repositories without a proxy intermediary.

# Procedures And API Details

- **Setting GOVCS**: Define a comma-separated list of rules (e.g., `github.com:git,evil.com:off,*:git|hg`) to restrict VCS usage per module path pattern.
- **Configuring Private Proxy (All Modules)**: Set `GOPROXY` to the internal proxy URL and `GONOSUMDB` to the module prefix to disable public checksum checks.
- **Configuring Direct Access**: Set `GOPRIVATE=corp.example.com` to bypass proxies for private modules; ensure VCS tools have non-interactive authentication configured (e.g., `.gitconfig`).
- **Authenticating with .netrc**: Place credentials in `$HOME/.netrc` (or `%USERPROFILE%\_netrc`) using the `machine`, `login`, and `password` fields for proxy servers.
- **Embedding Credentials in URLs**: Append username:password to the `GOPROXY` URL (e.g., `https://user:pass@proxy.example.com`), though this may expose secrets in shell history or logs.

# Nuance Or Contradictions

- While `GOVCS` defaults to restricting public paths to `git` and `hg`, public modules served via the module mirror (`proxy.golang.org`) are still accessible even if they originate from unsupported VCS systems like Bazaar, because the proxy handles the download securely.
- The special case for copying `LICENSE` files applies only to files named exactly `LICENSE` (no extension); extending this behavior would break cryptographic sums used for authentication.
- Empty directories can be included in zip files but are not extracted by the `go` command.
- File names up to the first dot cannot be reserved Windows filenames (e.g., `CON`, `COM1`) regardless of case, even if the full path is valid.

# Candidate Wiki Hints

- **Module Distribution Format**: Explain the structure and security constraints of Go module zip files.
- **GOVCS Configuration Guide**: Document how to configure version control restrictions for mixed public/private ecosystems.
- **Private Module Setup**: Step-by-step guide for configuring `GOPRIVATE`, `GOPROXY`, and authentication for internal repositories.
- **License Handling in Modules**: Clarify the automatic `LICENSE` file copying behavior and its limitations.
