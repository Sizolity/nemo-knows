## group-01

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This document synthesizes the **Fetch Standard** specification, covering the entire lifecycle of resource fetching in web browsers and environments. The notes span from high-level infrastructure (URLs, HTTP methods) through detailed request/response handling, security headers, CORS protocols, caching strategies, and the Fetch API itself. Key concepts include the distinction between filtered and internal responses, the mechanics of body streaming, and the complex logic behind cross-origin resource sharing.

# Cross-Chunk Summary

The specification defines a unified fetching model where diverse web APIs (e.g., `<img>`, `fetch()`, `XMLHttpRequest`) adhere to a common protocol. The process begins with **Infrastructure** (Section 2), defining URL schemes, HTTP method normalization, and header handling. It progresses to the core **Fetch Algorithm** (Sections 3-5), which manages the request lifecycle: initiating, resolving domains, reading bodies incrementally, handling redirects, and applying security policies like CORS and CSP. The final sections cover the **Fetch API** classes (`Headers`, `Request`, `Response`) and specific URL types like `data:` URLs.

Throughout the document, a recurring theme is **security through obscurity and filtering**. Responses are categorized into "filtered" (e.g., opaque, cors) and internal types to prevent scripts from accessing sensitive headers or redirect targets unless explicitly allowed by security policies. Similarly, request headers are strictly classified as CORS-safelisted or forbidden to mitigate cross-origin attacks.

# Repeated Or Central Claims

- **HTTP Method Normalization**: HTTP methods are technically case-sensitive per the protocol but are normalized to uppercase (`GET`, `POST`) within the Fetch Standard for consistency. Forbidden verbs (e.g., `CONNECT`, `TRACE`) remain disallowed regardless of casing.
- **Header Handling and Safety**: Headers are treated as a specialized ordered multimap where keys are byte-case-insensitive. The standard strictly separates "CORS-safelisted" headers (safe for cross-origin requests) from "forbidden" headers (scripts cannot set them). `Set-Cookie` is uniquely handled to prevent header collisions.
- **Response Types and Filtering**: Responses evolve over time and are categorized by type (`basic`, `cors`, `opaque`, `error`). Filtered responses provide a limited view (e.g., hiding status or headers) to scripts, while internal responses hold the full data. This distinction is critical for security, ensuring that sensitive information isn't leaked via cross-origin requests.
- **Body Streaming**: The body of a response is treated as a `ReadableStream`. Reading is incremental, utilizing parallel queues (`taskDestination`) to handle chunks, errors, and completion without blocking. Cloning a body involves "teeing" the stream rather than copying data immediately.
- **CORS Protocol**: Cross-Origin Resource Sharing relies on a complex interplay of headers (`Origin`, `Access-Control-*`), request modes (`same-origin`, `cors`, `no-cors`), and response tainting. Preflight checks are triggered for specific methods/headers, and the protocol includes exceptions for content types and credentials.
- **Redirect Handling**: The fetch algorithm handles redirects by iterating through a list of URLs. It computes `redirect-taint` to determine if the chain crosses site boundaries. Redirect targets are obscured in reporting (using the first URL) to prevent leaking internal redirect logic.

# Important Local Details

- **Forbidden Methods**: Case-insensitive matches for `CONNECT`, `TRACE`, and `TRACK` are forbidden. Using lowercase versions may result in `405 Method Not Allowed`.
- **CORS Safelisted Headers**: Only specific headers (e.g., `Accept`, `Content-Type` with restricted MIME types) are considered safe for cross-origin requests. Values exceeding 128 bytes are rejected for safety checks.
- **Status Code Ranges**: Status codes are integers from 0 to 999. Specific ranges define success (`2xx`), redirection (`3xx`), client errors (`4xx`), and server errors (`5xx`).
- **Request Classification**: Requests are classified as subresource (e.g., images, scripts), non-subresource, or navigation requests based on their destination type. This classification influences caching behavior and service worker interactions.
- **Cache Modes**: The `cache mode` attribute dictates interaction with the HTTP cache (`default`, `only-if-cached`, `no-cache`, `reload`). `only-if-cached` is restricted to `same-origin` mode.
- **Data URLs**: The specification includes a dedicated section for `data:` URLs, which are handled as a distinct scheme within the fetch infrastructure.

# Candidate Wiki Hints

- **Page: Fetch Standard Overview** – High-level goals and unification across APIs.
- **Page: HTTP Method Handling** – Normalization rules, CORS-safelisted vs forbidden methods.
- **Page: Fetch Controller Internals** – State machine, timing info, abort handling.
- **Page: Header List Structure** – Multimap logic, `Set-Cookie` uniqueness, parsing algorithms.
- **Page: Response Lifecycle** – Filtered vs internal responses, type evolution, cloning logic.
- **Page: CORS Protocol Deep Dive** – Safelisted headers, preflight checks, credential handling.
- **Page: Request Object Attributes** – Method, URL, cache mode, destination mapping, traversability.
- **Page: Body Streaming Mechanics** – ReadableStream integration, incremental reading, cloning via teeing.
- **Page: Redirect and Taint Logic** – `redirect-taint` computation, URL serialization for reporting.
- **Page: Fetch API Classes** – `Headers`, `Request`, `Response` initialization and behavior.

# Gaps Or Cautions

- **Content-Type Parsing**: The specification notes that servers are not expected to implement strict MIME type extraction; thus, the standard algorithm avoids using `extract a MIME type` in certain contexts to prevent failures.
- **Opaque Response Ambiguity**: `opaque` and `opaqueredirect` filtered responses lack accessible headers and status, making them nearly indistinguishable from network errors without inspecting internal state. New APIs should avoid relying on these distinctions.
- **DNS Resolution Caching**: While the spec allows for DNS caching to improve performance, it leaves implementation details (partition keys) somewhat open, potentially leading to inconsistent behavior across different environments.
- **Security Review Required**: Combining multiple responses into a single logical resource is flagged as a historical source of security bugs and requires a specific security review before implementation.
- **URL Credential Handling**: Modern specifications discourage setting the `use-URL-credentials` flag due to security concerns, even though it exists in the algorithm for overriding authentication entries.

## group-02

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

## group-03

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

## group-04

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This document synthesizes notes from Chunks 19 through 24 of the Fetch Standard specification (`raw/web/corpus-2026-05-18/058-fetch-standard.md`). The content exclusively covers **Section 6: data: URLs**, transitioning from introductory metadata and licensing into the core algorithmic definitions, structural parsing rules, API interfaces, and browser compatibility matrices for handling `data:` URIs.

# Cross-Chunk Summary

The selected chunks provide a comprehensive view of the `data:` URL scheme within the Fetch Standard:
*   **Chunks 19–20** establish the legal framework (WHATWG licensing) and define the structural representation (`data: URL struct`) for parsing these URLs into components like MIME type, charset, and base64 content.
*   **Chunk 21** details the fetch algorithm specific to `data:` URLs, clarifying that they do not trigger network I/O but instead parse the fragment identifier directly. It highlights restrictions on cross-origin access (CORS) and the inability to perform redirects.
*   **Chunks 22–24** focus on the API layer and implementation details. They define the `Request` and `Response` interfaces relevant to fetching, list normative references (HTML, Fetch Metadata), and provide extensive compatibility matrices for browser engines (Chrome, Firefox, Safari, Edge) regarding properties like `body`, `json()`, and CORS headers.

# Repeated Or Central Claims

*   **No Network I/O**: A consistent theme across Chunks 20 and 21 is that fetching a `data:` URL does not involve network connections or cache partitions; it is an immediate processing step within the fetch algorithm.
*   **Parsing Logic**: The standard requires parsing the URI to extract the media type and parameters from the fragment identifier (Chunk 20, 21).
*   **CORS Restrictions**: While `data:` URLs are local resources, Chunks 21 and 24 note that they are subject to security policies like `Cross-Origin-Opener-Policy` and general CORS checks, often resulting in errors if treated as cross-origin resources.
*   **Interface Definitions**: The chunks repeatedly reference the `Request`, `Response`, and `Headers` interfaces, defining their attributes (e.g., `method`, `url`, `headers`) and methods (e.g., `append`, `get`, `json`).
*   **Compatibility Nuances**: Chunks 23 and 24 emphasize that support for specific features (like `Request.body` or static `Response.json_static`) varies significantly by browser version, with legacy browsers (Edge Legacy) and mobile webviews showing gaps in coverage.

# Important Local Details

*   **Licensing Distinction**: The text explicitly distinguishes between the Creative Commons Attribution 4.0 International license for the specification text and the BSD 3-Clause License for incorporated source code portions (Chunk 19).
*   **Data URL Structure**: A `data: URL struct` holds parsed components including scheme, base64 content, media type, and charset (Chunk 20).
*   **Empty Body Default**: By default, requests for `data:` URLs are created with an empty body, which is populated by reading the data from the URL's fragment during parsing (Chunk 21).
*   **Redirection Failure**: Attempts to redirect a `data:` URL fetch fail because they do not support the redirection mechanism used in HTTP network fetches (Chunk 21).
*   **Specific API Methods**: The chunks detail methods such as `Response.error()`, `Response.redirect()`, and `Response.json()` for creating synthetic responses, alongside the `fetchLater()` method for deferred fetching.
*   **Normative References**: Section 6 inherits definitions from external standards including W3C Living Standards, RFCs (e.g., [HTTP], [HTML]), and security advisories like [HTTPVERBSEC1] (Chunk 22).

# Candidate Wiki Hints

*   **Topic: Data URLs in Fetch API**
    *   A dedicated page explaining the `data:` URL scheme, its parsing logic, and its place within the Fetch Standard's lifecycle.
*   **Concept: Cross-Origin Policies for Data URIs**
    *   Documentation on how CORS policies (`Cross-Origin-Opener-Policy`, etc.) apply to local resources like `data:` URLs and why they might trigger security errors.
*   **API Reference: Request & Response Interfaces**
    *   A reference guide for the `Request` and `Response` objects, detailing attributes (`mode`, `credentials`, `integrity`) and utility methods (`clone`, `json`, `text`).
*   **Guide: Browser Compatibility for Fetch API**
    *   A matrix summarizing support for key features (e.g., `body` access, `fetchLater()`, specific CORS headers) across Chrome, Firefox, Safari, Edge, and mobile webviews.
*   **Topic: Deferred Fetching (`fetchLater`)**
    *   Notes on the experimental or newer `fetchLater()` method and its `DeferredRequestInit` dictionary for lazy loading strategies.

# Gaps Or Cautions

*   **Incomplete Mobile Data**: Chunks 23 and 24 contain numerous question marks (?) in their compatibility matrices, particularly for Android WebViews and iOS Safari, indicating a lack of confirmed support data or testing gaps for these environments.
*   **Legacy Edge Discontinuity**: The text notes a sharp discontinuity between "Edge (Legacy)" (EdgeHTML) and modern Chromium-based Edge, with the former lacking support for features like `Response.body` despite supporting others.
*   **MDN Documentation Status**: Some headers and features are flagged with warning indicators (⚠MDN), suggesting that while they may exist in the browser implementation, their standardization or documentation status is uncertain compared to universally supported items.
*   **Body Readability Variance**: There is a specific contradiction regarding `Request.body` support; while listed as generally available, Firefox only supports it from version 65+, contradicting a blanket "all current engines" claim for older versions.

