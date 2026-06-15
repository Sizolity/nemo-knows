---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/032-go-modules-reference.md
confidence: medium
---
Chunk Context
This chunk covers the internal mechanics of Go modules, specifically focusing on how `go` creates and validates `.zip` files for module distribution. It details constraints on file paths, sizes, and content (e.g., ignoring vendor directories or symlinks), special handling for LICENSE files, and security controls via the `GOVCS` environment variable to restrict which version control systems are used when downloading from public vs. private servers.

Local Summary
The `go` command packages module contents into `.zip` files after locating the module root. These archives are authenticated before extraction. The chunk outlines strict rules for what goes into a module zip (no vendor dirs, no nested go.mod dirs, specific path prefixes) and how to handle license files across subdirectories. It also explains the `GOVCS` variable, which balances the need for decentralized package hosting with security by restricting default VCS usage to Git and Mercurial for public modules while allowing others for private ones via proxies.

Key Claims
- Module `.zip` files are authenticated before extraction into the cache, similar to proxy downloads.
- Vendor directories and nested modules (subdirectories with `go.mod`) are excluded from module zip files.
- If a subdirectory module lacks a `LICENSE` file, `go` copies one from the repository root if present.
- By default, `go` uses Git and Mercurial for public servers; other VCS tools are reserved for private modules or proxied downloads.
- The `GOVCS` environment variable allows explicit control over allowed version control systems per module path pattern.
- Module zip files are limited to 500 MiB total size (compressed and uncompressed).

Entities And Concepts
- `go mod download`: Command to fetch and extract modules.
- `GOPROXY`: Environment variable listing proxy URLs for module downloads.
- `GOVCS`: Environment variable controlling allowed version control systems per path pattern.
- `GONOPROXY`, `GONOSUMDB`, `GOPRIVATE`: Variables configuring private module access and security checks.
- `.zip` file format constraints (path prefixes, size limits, ignored metadata).
- Version Control Systems: Git, Mercurial, Bazaar, Fossil, Subversion.

Procedures And API Details
- **Setting GOVCS**: Use a comma-separated list of `pattern:vcslist` rules.
  - Example: `GOVCS=github.com:git,evil.com:off,*:git|hg`
  - Default logic: `public:git|hg,private:all`.
- **Disabling VCS usage**: Set `GOVCS=*:off`.
- **Enabling all VCS**: Set `GOVCS=*:all`.
- **Configuring private proxy**:
  - Set `GOPROXY=https://proxy.corp.example.com` and `GONOSUMDB=corp.example.com`.
- **Direct access to private modules**:
  - Set `GOPRIVATE=corp.example.com`.
  - Ensure repository URLs use HTTPS or have a VCS suffix (e.g., `.git`).

Nuance Or Contradictions
- While `go` ignores symbolic links in zip files, authors can manually copy license files into subdirectory modules if the root lacks one.
- The default restriction to Git and Mercurial for public servers is a security measure; other VCS tools are allowed for private modules or when using the proxy mirror (proxy.golang.org).
- `GOVCS` patterns match leading elements of module/import paths; the earliest matching pattern applies, even if later ones also match.

Candidate Wiki Hints
- Page: Go Modules Security and Version Control Configuration
  - Focus on `GOVCS`, `GOPRIVATE`, and managing private module access.
- Page: Module Zip File Constraints
  - Document size limits, path prefixes, and ignored file types (vendor, symlinks).
