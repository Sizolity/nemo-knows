---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This group covers the **HTTP Fetch Standard**, detailing the complete lifecycle of fetching resources over a network. It encompasses infrastructure setup (URLs, HTTP methods/statuses), connection management (pools, partitioning, timing), security extensions (CORS, cookies, Content-Type options), and the core fetch algorithms (scheme handling, redirection, caching, and response processing). The content bridges the gap between raw network I/O and high-level Fetch API interactions.

# Cross-Chunk Summary

The document progresses from foundational infrastructure to complex security protocols and finally to the execution engine of fetching.

1.  **Infrastructure & HTTP Basics**: Early chunks establish the environment: URL parsing, HTTP methods, status codes, and response bodies.
2.  **Connection Management**: The system manages connections via a pool keyed by network partition keys (origin + credentials). Timing info is clamped to prevent privacy leakage. Specific handling exists for bad ports (e.g., echo/daytime) and MIME type blocking.
3.  **Security Extensions**: A significant portion covers headers that govern security contexts:
    *   **Cookies**: Infrastructure for `Cookie` and `Set-Cookie`, including SameSite modes and HttpOnly flags.
    *   **CORS**: The Cross-Origin Resource Sharing protocol, involving the `Origin` header serialization, preflight requests (`OPTIONS`), and response headers like `Access-Control-Allow-Origin`.
    *   **Content-Type Safety**: Headers like `X-Content-Type-Options` (nosniff) and `Cross-Origin-Resource-Policy` (CORP) restrict content interpretation and loading.
4.  **Fetching Algorithms**: The core logic is split into:
    *   **Main Fetch**: Initialization, security checks (CSP, HSTS), response tainting (`basic`, `cors`, `opaque`), and header population.
    *   **Scheme Fetch / HTTP Fetch**: Handling specific schemes (blob, data) vs. network protocols, including redirect logic and method normalization.
    *   **Caching & Validation**: Strategies for storing responses, revalidation via 304s, and handling stale-while-revalidate modes.
5.  **Response Processing**: Final steps include transforming bodies, integrity validation, and returning the final response object or a network error.

# Repeated Or Central Claims

*   **Connection Pooling & Partitioning**: Connections are strictly isolated by a "network partition key" (derived from top-level site, origin, and credentials). TLS session identifiers are not reused across different credential sets to prevent cross-origin leakage.
*   **Origin Header Serialization**: The `Origin` header is byte-serialized using strict ABNF rules (lowercase ASCII, specific IPv6 limits) rather than raw RFC 3986 parsing. It is appended based on response tainting ("cors") or request method (non-GET/HEAD).
*   **CORS Protocol Mechanics**: CORS is opt-in. Successful responses can have any status code if headers permit. `Access-Control-Allow-Origin` cannot be `*` when credentials are included. The protocol relies heavily on preflight caching to avoid redundant requests.
*   **Redirect Normalization**: Redirects (301/302) often force a method change from `POST` to `GET` and clear the body to prevent data loss or security issues. Cross-origin redirects strip CORS-sensitive headers like `Authorization`.
*   **Response Tainting**: Responses are categorized as "basic" (same origin, no CORS), "cors" (cross-origin sharing allowed), or "opaque" (no access to status/body). This dictates how JavaScript contexts can read the response.
*   **Security Header Enforcement**: Headers like `X-Content-Type-Options: nosniff` and `Cross-Origin-Resource-Policy` act as gates; if a MIME type is invalid for the destination or origin policy fails, the fetch is blocked entirely.

# Important Local Details

*   **Bad Port Blocking**: Ports listed in a specific table (e.g., 0, 1, 7, 9) trigger immediate request blocking to prevent interaction with legacy protocols like echo or daytime.
*   **MIME Type Fatal Errors**: Incorrect MIME essence is treated as a fatal error during extraction, though parameters (charset) are often ignored. This contrasts with historical behavior where browsers were lenient.
*   **Timing Info Clamping**: To hide connection reuse patterns, timing details (domain lookup, connection start/end) are clamped and coarsened before being exposed to the user agent.
*   **CORS Preflight Cache**: The user agent maintains a cache for CORS preflight requests (`OPTIONS`) keyed by method/headers. Entries are added only if they don't already exist, optimizing performance.
*   **Blob Range Semantics**: When fetching `blob:` URLs with range headers, the byte range boundaries must be adjusted (incrementing the end index) because the underlying `slice` algorithm uses an exclusive upper bound, unlike the HTTP inclusive convention.
*   **Keepalive Limits**: The sum of a request's body length and active keepalive bytes is capped at 64 kibibytes to prevent indefinite memory usage.

# Candidate Wiki Hints

*   **Fetch API Architecture**: A high-level overview of how browsers parse URLs, manage connections, and execute fetch requests.
*   **CORS Protocol Deep Dive**: Explaining the interaction between `Origin`, `Access-Control-Allow-Origin`, and credential handling modes.
*   **Connection Pooling & Partitioning**: How browsers isolate network resources by origin, credentials, and site to enhance privacy.
*   **HTTP Cache Strategies**: Details on `no-store`, `reload`, and stale-while-revalidate behaviors within the Fetch Standard.
*   **Security Headers Guide**: A reference for `X-Content-Type-Options`, `Cross-Origin-Resource-Policy`, and their impact on content loading.
*   **Redirect Behavior**: Rules governing method changes (POST -> GET) and header stripping during 3xx redirects.

# Gaps Or Cautions

*   **Implementation Discretion**: The standard explicitly leaves nuances like IP address selection (IPv6 vs IPv4), retry logic, and proxy resolution strategies to implementers.
*   **Partial Content Caching**: While the HTTP spec supports partial content caching, browser implementations generally do not widely support it, potentially leading to unexpected behavior with `Range` headers.
*   **MIME Type Leniency**: Although the standard states incorrect MIME essence is a fatal error, existing web features often ignore this, creating a gap between strict specification and practical reality.
*   **Referrer Policy Ambiguity**: The handling of the `Origin` header in conjunction with various referrer policies (e.g., "no-referrer", "strict-origin") can be complex, particularly regarding scheme downgrades from HTTPS to HTTP.
*   **Data URL Repetition**: Chunks 18-24 cover "data: URLs" extensively but appear to repeat the same heading path in the index, suggesting a potential segmentation artifact or extensive detail on data handling that needs careful consolidation.
