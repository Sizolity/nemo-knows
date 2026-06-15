---
title: Http Fetch Lifecycle
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Http Fetch Lifecycle

The **Fetch Standard** defines a unified architecture for fetching resources across the web platform, superseding previous inconsistent semantics. This single definition applies to all web APIs, including `fetch()`, `<img>`, `<script>`, and Service Workers, ensuring consistent behavior for redirects, caching, and security constraints.

## Lifecycle Phases

The specification details the complete lifecycle of an HTTP request and response, progressing through distinct phases:

### Infrastructure Setup
The process begins with **Infrastructure**, which defines URL parsing, method normalization (converting verbs to uppercase), and header handling rules. This foundational layer ensures that all subsequent operations work with standardized inputs.

### Connection Management
Next, the lifecycle moves into **Connection Management**. Here, connections are partitioned by origin and security context (e.g., network partition keys) to prevent data leakage between sites. This phase strictly manages how resources are routed based on the site origin and credentials.

### Fetching Algorithms
The core logic resides in the **Fetching Algorithms**. These algorithms handle:
- Scheme resolution for various URL types.
- Caching strategies (default, no-store, reload).
- Redirect logic, supporting up to 20 redirects.
- Security protocols like CORS and Mixed Content blocking.

### Body Consumption
Finally, the lifecycle concludes with body consumption via **Fetch Body Streams**. Request and response bodies are implemented as `ReadableStream` objects, supporting incremental reading, cloning via `tee()`, and content decoding without loading the entire payload into memory immediately.

## Security and Response States

The standard emphasizes **Security First**, implementing strict MIME type extraction and blocking requests on "bad ports" (like FTP or Telnet). It enforces Cross-Origin-Embedder-Policy (COEP) checks and ignores dangerous header parameters.

**Response Tainting** categorizes responses into three states:
1.  **Basic**: Full access to headers and body.
2.  **CORS**: Restricted header exposure based on CORS-safelisted headers.
3.  **Opaque**: Limited access, typically used for cross-origin requests where the body cannot be read directly.

## API Abstractions

The **Fetch API** section defines the JavaScript interfaces (`Headers`, `Request`, `Response`) that abstract these low-level operations. These classes expose attributes like `status`, `headers`, and `body` while managing internal states such as timing info and abort signals.

### Memory Management
To handle potential resends due to timeouts, request bodies are buffered to 64 KiB when the source is a stream. This is managed through **deferred-fetching-quota**, which handles memory usage hierarchically (640 KiB top-level, 64 KiB concurrent per origin).

## Key Concepts

- [[http-fetch-lifecycle]]
- [[request-response-objects]]
- [[fetch-body-streams]]
- [[cors-preflight-cache]]
- [[deferred-fetching-quota]]
- [[fetch-security-policies]]
