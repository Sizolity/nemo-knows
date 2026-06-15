---
title: Request Response Objects
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Request Response Objects

In the **Fetch Standard**, request and response objects are the core interfaces that abstract the low-level details of HTTP interactions. These objects represent the infrastructure setup, connection management, and final body consumption stages of a unified fetching architecture.

## Core Concepts

### Unified Architecture
The standard replaces fragmented implementations with a single algorithm governing all web platform fetching APIs, including `fetch()`, `<img>`, `<script>`, and Service Workers. This ensures consistent behavior for redirects, caching, and security constraints across different URL schemes like HTTP(S), data, and blob.

### Object Lifecycle
The lifecycle of an HTTP request and response begins with **Infrastructure**, defining URL parsing, method normalization (converting verbs to uppercase), and header handling rules. It moves into **Connection Management**, where connections are partitioned by origin and security context to prevent data leakage between sites. The core logic resides in the **Fetching Algorithms**, which handle scheme resolution, caching strategies, redirect logic, and security protocols like CORS.

### Security and State
The standard emphasizes **Security First**, implementing strict MIME type extraction, blocking requests on "bad ports," and enforcing Cross-Origin-Embedder-Policy (COEP) checks. Response objects are categorized into three states—`basic`, `cors`, and `opaque`—which dictate header exposure, body readability, and redirect handling. Headers are classified as "CORS-safelisted" or forbidden, requiring explicit opt-in via headers like `Access-Control-Allow-Origin`.

### Body Handling
Request and response bodies are represented as streams (`ReadableStream`). This design supports incremental reading, cloning via `tee()`, and content decoding without loading the entire payload into memory immediately. To handle potential resends due to timeouts, request bodies are buffered to 64 KiB when the source is a stream.

## Related Concepts

- [[fetch-standard-overview]]
- [[http-fetch-lifecycle]]
- [[cors-preflight-cache]]
- [[fetch-body-streams]]
- deferred-fetch-quota
