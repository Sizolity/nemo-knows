---
title: Fetch Standard
kind: source
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## What It Is

The **Fetch Standard** (Living Standard) defines a unified architecture for fetching resources in the web platform. It supersedes previous inconsistent semantics, providing a single definition that applies to all web APIs—including `fetch()`, `<img>`, `<script>`, and Service Workers—ensuring consistent behavior for redirects, caching, and security constraints across different URL schemes like HTTP(S), data, and blob.

## Summary

The specification details the complete lifecycle of an HTTP request and response, from infrastructure setup to final body consumption. The process begins with **Infrastructure**, defining URL parsing, method normalization (converting verbs to uppercase), and header handling rules. It moves into **Connection Management**, where connections are partitioned by origin and security context (e.g., network partition keys) to prevent data leakage between sites.

The core logic resides in the **Fetching Algorithms**. These handle scheme resolution, caching strategies (default, no-store, reload), redirect logic (up to 20 redirects), and security protocols like CORS and Mixed Content blocking. The standard emphasizes **Security First**, implementing strict MIME type extraction, blocking requests on "bad ports" (like FTP or Telnet), and enforcing Cross-Origin-Embedder-Policy (COEP) checks.

Finally, the **Fetch API** section defines the JavaScript interfaces (`Headers`, `Request`, `Response`) that abstract these low-level operations. These classes expose attributes like `status`, `headers`, and `body` while managing internal states such as timing info and abort signals. The document concludes with specific handling for `data:` URLs, garbage collection semantics, and extensive browser compatibility matrices.

## Key Claims

*   **Unified Fetching Architecture**: Replaces fragmented implementations with a single algorithm governing all web platform fetching APIs.
*   **Method Normalization**: HTTP methods are normalized to uppercase (e.g., `GET`, `POST`) for compatibility, while custom methods require careful handling.
*   **Response Tainting**: Responses are categorized into three states—`basic`, `cors`, and `opaque`—which dictate header exposure, body readability, and redirect handling.
*   **CORS Safety & Preflight**: Headers are classified as "CORS-safelisted" or forbidden. Cross-Origin Resource Sharing requires explicit opt-in via headers (`Access-Control-Allow-Origin`) and often involves a two-step preflight `OPTIONS` request.
*   **Body Streams**: Request and response bodies are streams (`ReadableStream`), supporting incremental reading, cloning via `tee()`, and content decoding without loading the entire payload into memory immediately.
*   **Connection Partitioning**: Connections are strictly partitioned by a "Network Partition Key" (site origin + credentials) to ensure cookies and sensitive data do not leak between different origins or security contexts.
*   **Buffering Limits**: Request bodies are buffered to 64 KiB when the source is a stream to handle potential resends due to timeouts, with deferred fetching quotas managing memory usage hierarchically (640 KiB top-level, 64 KiB concurrent per origin).
*   **Security by Default**: The standard prioritizes security over legacy compatibility, blocking mixed content, enforcing strict MIME type extraction, and ignoring dangerous header parameters.

## Suggested Links

none
