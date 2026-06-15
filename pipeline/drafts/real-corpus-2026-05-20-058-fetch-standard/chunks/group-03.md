---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This group of notes synthesizes the detailed specification for the Fetch Standard, specifically focusing on the lifecycle of network requests and responses. The content spans from the core HTTP infrastructure (URLs, methods, headers) to the complex algorithms governing connection management, caching, and security policies like CORS (Cross-Origin Resource Sharing). It further details the Fetch API interfaces (`Headers`, `Request`, `Response`) and advanced mechanisms such as deferred fetching with quota management, data URL processing, and garbage collection rules for fetch operations.

# Cross-Chunk Summary

The document defines a comprehensive model for resource fetching over HTTP. The process begins with **Infrastructure**, establishing how URLs are parsed and HTTP requests are formed via methods and headers. The core logic resides in the **Fetch** algorithms:
1.  **Connection & Request**: Requests are routed through network partitions (connections). If the body source is null (e.g., a `ReadableStream`), it requires buffering to prevent data loss if retransmission fails.
2.  **Response Handling**: Responses are parsed, interim status codes (1xx) are handled, and bodies are decoded and buffered.
3.  **Security & Privacy**: CORS protocols dictate how cross-origin requests are validated via preflight (`OPTIONS`) checks and caching. Headers like `Origin`, `Content-Type`, and security policies (`X-Content-Type-Options`, `Cross-Origin-Resource-Policy`) control access.
4.  **API Interfaces**: The Fetch API exposes this logic through `Headers` (with mutation guards), `Request` (with modes/credentials), and `Response` (including static methods for errors and JSON).
5.  **Advanced Features**: Deferred fetching allows lazy loading with strict quota limits, while data URLs are processed internally without network calls.

# Repeated Or Central Claims

*   **Network Partitioning**: Connections are managed via "network partition keys" to group requests logically (e.g., by origin or protocol).
*   **CORS Preflight Caching**: To minimize redundant network overhead, the user agent caches preflight (`OPTIONS`) validation results. These entries include specific keys for origin, URL, method, headers, and max-age.
*   **Body Buffering Limits**: When a request body is derived from a `ReadableStream` (null source), it is unrecreatable. Consequently, the spec mandates buffering up to 64 KiB; exceeding this results in a network error if retransmission is needed.
*   **Quota Management for Deferred Fetching**: A top-level quota of 640 kibibytes exists, with sub-allocations (8 kibibytes) for cross-origin nested documents and a strict per-origin limit of 64 kibibytes to prevent resource exhaustion.
*   **Header Guards**: The `Headers` interface uses guard states ("immutable", "request", "response") to enforce mutation rules and CORS safety, automatically removing privileged headers in "no-cors" contexts.
*   **Observable Termination**: User agents may garbage collect ongoing fetches if their state (headers/status) is accessible but the body stream is not being actively read/observed.

# Important Local Details

*   **Request Body Handling**: The `processBodyChunk` algorithm manages reading from sources. For null sources, it tracks encoded vs. decoded sizes and suspends fetching if the 64 KiB buffer limit is exceeded.
*   **CORS Validation Logic**: A request is valid only if the response status is OK, the TAO (Timing-Allow-Origin) check passes, and the requested method/headers are listed in `Access-Control-Allow-Methods`/`Allow-Headers`.
*   **Deferred Fetch Quota Calculation**: The "total request length" includes the serialized URL, referrer, header lengths, and body length. This sum is checked against available quota before queuing a deferred fetch.
*   **Data URL Decoding**: The `data:` URL processor handles MIME types ending in `;base64` by performing an isomorphic decode followed by forgiving-base64 decoding. If no base64 encoding is detected, the body remains percent-decoded.
*   **Response Static Methods**: `Response.json()` serializes data to JSON bytes and sets the content type automatically. `Response.redirect()` constructs a response with a default status of 302 and appends the target URL to the `Location` header.

# Candidate Wiki Hints

*   **Page: Fetch Algorithm Overview**
    *   *Topic*: High-level flow from request initiation to response consumption.
    *   *Key Content*: Connection acquisition, body buffering, interim response handling, and finalization.
*   **Page: CORS Preflight Cache**
    *   *Topic*: Structure and lifecycle of preflight validation entries.
    *   *Key Content*: Cache entry fields (origin, URL, method, headers), matching logic, and max-age expiration.
*   **Page: Request Body Constraints**
    *   *Topic*: Handling `ReadableStream` bodies and buffering limits.
    *   *Key Content*: The 64 KiB buffer requirement for null sources and the resulting network error conditions.
*   **Page: Deferred Fetching API**
    *   *Topic*: Usage of `fetchLater()` and quota management.
    *   *Key Content*: Quota limits (640k total, 64k per origin), `Permissions-Policy` header, and activation states.
*   **Page: Fetch API Interfaces**
    *   *Topic*: Detailed documentation of `Headers`, `Request`, and `Response` classes.
    *   *Key Content*: Guard states, mutation methods, body init types, and static factory methods.
*   **Page: Data URL Processing**
    *   *Topic*: Internal handling of `data:` URLs within the Fetch API.
    *   *Key Content*: Parsing MIME types, base64 decoding steps, and struct composition.

# Gaps Or Cautions

*   **Missing Raw Text**: The provided chunk notes summarize algorithms and concepts but do not contain the full raw text of the specification (e.g., specific algorithm step numbers or exact code blocks for error handling).
*   **Future Work Notes**: The notes mention that detailed parsing specifications for `multipart/form-data` are to be written later, indicating potential implementation variance or future updates.
*   **Service Worker Context Nuance**: The text explicitly omits "serviceworker" from `RequestDestination` in the IDL because it cannot be observed from JavaScript, though implementations must support it internally; this distinction is crucial for developers debugging service worker interactions.
*   **Quota Policy Defaults**: There is a specific default behavior where `deferred-fetch-minimal` is enabled for all origins ("*"), while `deferred-fetch` is restricted to the top-level origin by default, which can be confusing without explicit policy configuration examples.
