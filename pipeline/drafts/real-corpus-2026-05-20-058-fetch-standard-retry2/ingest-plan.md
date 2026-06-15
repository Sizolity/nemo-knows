---
kind: topic
sources: [raw/web/corpus-2026-05-18/058-fetch-standard.md]
status: draft
---

# Ingest Plan

## Source Summary
- The document defines the **Fetch Standard** (Living Standard), unifying fetching logic across all web APIs (`fetch()`, `<img>`, `fetch()` in Service Workers) to replace fragmented implementations.
- It covers the complete lifecycle from **Infrastructure** (URLs, HTTP methods, headers) through **Connection Management**, **Fetching Algorithms** (scheme fetches, redirects, caching), and finally **Fetch API** interfaces (`Headers`, `Request`, `Response`).
- Security is prioritized via strict MIME type extraction, blocking bad ports, CORS preflight validation, and origin serialization to prevent redirect taint leakage.

## Candidate Wiki Pages
- wiki/concepts/fetch-standard-overview.md — Introduction to the unified architecture, API coverage (`fetch`, `<img>`, `navigator.sendBeacon`), and goals of replacing inconsistent semantics.
- wiki/topics/http-fetch-lifecycle.md — Detailed breakdown of the request lifecycle: infrastructure setup -> scheme fetch -> redirect handling -> response tainting (basic/cors/opaque).
- wiki/concepts/fetch-security-policies.md — Coverage of CORS protocol, Content Security Policy hooks, Cross-Origin-Embedder-Policy checks, and Mixed Content blocking.
- wiki/topics/request-response-objects.md — Documentation of `Request` and `Response` object attributes, lifecycle states, cloning mechanics, and the Fetch API initialization dictionaries.
- wiki/concepts/fetch-body-streams.md — Explanation of bodies as streams (`ReadableStream`), incremental reading algorithms, `tee()` usage for cloning, and buffering limits (64 KiB).
- wiki/topics/cors-preflight-cache.md — Logic behind caching preflight responses to avoid redundant network calls, including TAO checks and cache entry structure.
- wiki/concepts/deferred-fetching-quota.md — Breakdown of the hierarchical quota system (top-level, same-origin children) and `Permissions-Policy` usage for resource management.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that "data: URLs" section is consolidated into a single canonical page under wiki/concepts/ or wiki/topics/.
- [ ] Ensure browser compatibility matrices (Chrome, Firefox, Safari, IE) are referenced or linked if external sources exist.
- [ ] Confirm that the distinction between CC BY 4.0 (spec text) and BSD 3-Clause (source code) is noted in the overview page.
- [ ] Check that all "Gaps Or Cautions" regarding legacy compatibility and implementation discretion are summarized in the checklist or a separate note.
