## group-01

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This group of notes synthesizes the **Fetch Standard** (Living Standard) document, specifically covering the infrastructure and core algorithms governing HTTP fetching in the web platform. The content spans from the document preface through detailed specifications for URL schemes, HTTP methods, headers, request/response objects, bodies, authentication, CORS protocols, cache management, and various fetch modes (scheme, network, redirect). It also covers specialized topics like data URLs, cookies, cross-origin resource policies (CORS, COEP), and response filtering.

The document unifies fetching logic across APIs such as `fetch()`, `<img>`, `<script>`, `navigator.sendBeacon()`, and Service Workers, ensuring consistent behavior for redirects, caching, and security constraints. It supersedes previous inconsistent semantics, particularly regarding the `Origin` header.

# Cross-Chunk Summary

The Fetch Standard defines a unified architecture where fetching is conceptually simple (request in, response out) but involves complex underlying details. The flow generally proceeds as follows:

1.  **Infrastructure & Request Setup**: The process begins with URL parsing and method normalization (uppercase for standard verbs). Requests are constructed with attributes like `method`, `url`, `headers`, `body`, `mode` (same-origin, cors, no-cors, navigate), and `credentials`.
2.  **Fetching Algorithm**: A request is processed through a series of fetch steps:
    *   **Scheme Fetch**: Determines the protocol (`http`, `https`, `data`, etc.).
    *   **HTTP Fetch**: Handles the network interaction, potentially checking the cache first (if mode allows).
    *   **Redirects**: If redirects are followed, the process loops back to scheme fetch.
    *   **Network Error**: Returns a network error if the request fails or is blocked (e.g., mixed content, forbidden headers).
3.  **Response Handling**: The response is constructed with a status code, headers, and body. It has an associated type (`basic`, `cors`, `default`, `error`, `opaque`).
4.  **Security & Policies**: Throughout the process, security checks occur, including CORS preflight validation, Content Security Policy (CSP) hooking, Cross-Origin-Embedder-Policy (COEP), and Mixed Content blocking.
5.  **Body Processing**: Request and response bodies are handled as streams (`ReadableStream`), supporting incremental reading, cloning via `tee()`, and content decoding.

# Repeated Or Central Claims

*   **Unified Fetching Architecture**: The standard provides a single definition for fetching that applies to all web APIs, replacing fragmented implementations.
*   **Method Normalization**: HTTP methods are technically case-sensitive, but the standard normalizes them to uppercase (e.g., `GET`, `POST`) for compatibility. Custom methods must be carefully handled.
*   **Header Handling**: Headers are stored as ordered lists (multimaps). Non-`Set-Cookie` headers are combined into single values when exposed to JavaScript, while `Set-Cookie` is preserved individually.
*   **Request/Response Objects**: Both `Request` and `Response` objects encapsulate the state of a fetch operation, including attributes like `body`, `headers`, `status`, and various flags (`keepalive`, `aborted`).
*   **CORS Safety**: Headers are categorized as "CORS-safelisted" (safe to send cross-origin) or forbidden. This classification depends on character sets and length limits.
*   **Body Streams**: Bodies are represented by streams, allowing for incremental reading without loading the entire payload into memory immediately. Cloning a body shares the underlying stream via `tee()`.
*   **Security First**: The standard emphasizes security through CSP directives, COEP checks, origin serialization (preventing redirect taint leakage), and filtering responses to prevent information leakage.

# Important Local Details

*   **URL Schemes**: The standard supports "about", "blob", "data", "file" (fetch scheme), and HTTP(S).
*   **Method Categories**:
    *   *CORS-safelisted*: `GET`, `HEAD`, `POST` (and others not in forbidden lists).
    *   *Forbidden*: `CONNECT`, `TRACE`, `TRACK`.
    *   *Normalized*: Uppercase versions of standard verbs.
*   **Cache Modes**:
    *   `default`: Uses HTTP cache, revalidates if needed.
    *   `no-store`: Does not use the cache.
    *   `only-if-cached`: Returns cached response or network error (requires same-origin).
*   **Response Types**:
    *   `basic`/`cors`: Standard responses with headers accessible.
    *   `default`: Responses from `fetch()` in non-CORS modes.
    *   `opaque`: Responses where status and headers are hidden (often used for `no-cors`).
    *   `error`: Indicates a network error occurred.
*   **CORS Preflight**: A two-step process involving an `OPTIONS` request to check permissions before sending the actual request. This involves caching preflight results.
*   **Range Requests**: Support for `Range` headers and `Content-Range` responses, allowing partial content retrieval.
*   **Filtered Responses**: Used to expose limited views of a response (e.g., only image data) to prevent leaking sensitive header info or full body content in certain contexts.

# Candidate Wiki Hints

*   **Page: Fetch Standard Overview** – Introduction to goals, unified architecture, and API coverage.
*   **Page: HTTP Method Normalization** – Rules for uppercase conversion, forbidden methods, and CORS-safelisted lists.
*   **Page: Request & Response Objects** – Detailed attributes, default values, cloning mechanics, and lifecycle states.
*   **Page: Fetch Modes & Cache Interaction** – Explaining `same-origin`, `cors`, `no-cors`, `navigate` modes and cache behavior (`default`, `reload`, etc.).
*   **Page: Body Streams & Cloning** – How bodies work as streams, `tee()` usage, incremental reading algorithms.
*   **Page: CORS Protocol Deep Dive** – Preflight logic, header filtering, credentials handling, and exceptions.
*   **Page: Security Policies (CSP/COEP)** – Content Security Policy hooks, Cross-Origin-Embedder-Policy checks, and Mixed Content blocking.
*   **Page: Origin Resolution & Taint** – How origins are resolved, serialized, and how redirect-taint is computed to prevent leakage.

# Gaps Or Cautions

*   **Missing Raw Text**: While the chunk outline provides headings and summaries, specific algorithmic pseudocode (e.g., exact microstep sequences for `fetch()` implementation) is not fully detailed in the provided notes beyond high-level descriptions.
*   **Data URL Specifics**: Chunks 18–24 cover "data: URLs" but only list them as a heading without detailed content in the notes, potentially missing nuances on parsing or security constraints specific to data URIs.
*   **Service Worker Interaction**: The preface mentions Service Workers, but the detailed interaction between fetch algorithms and Service Worker event handlers (e.g., `fetch` events) is not elaborated upon in these chunks.
*   **Implementation Variations**: The notes mention "implementation-defined operations" (e.g., DNS queries) which may vary across browsers or environments; caution is needed when relying on specific IP resolution orders.
*   **Legacy Compatibility**: Some sections note that certain features (like opaque filtered responses for legacy image decoding) are discouraged for new specifications due to architectural limitations, implying potential deprecation or migration paths not fully detailed here.

## group-02

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This document synthesizes the Fetch Standard specification covering the complete lifecycle of a network request and response. It spans from initial infrastructure (URLs, HTTP methods, headers) through connection management, security policies (CORS, CORP), the core fetching algorithms (scheme fetches, redirects, caching), and finally the Fetch API types (`Headers`, `Request`, `Response`). The notes bridge low-level network mechanics with high-level API usage patterns.

# Cross-Chunk Summary

The specification describes a request as an interaction between a user agent and a remote server. This process begins with **Infrastructure** (defining URLs, HTTP methods, and headers), moves to **Connection Management** (partitioning connections by origin/credentials, handling bad ports), and applies security protocols like **CORS** and **Content-Type** safety checks. The core logic resides in the **Fetching** algorithms, which handle scheme resolution (HTTP, HTTPS, Data, Blob), caching strategies (validation, negative caching), redirect logic, and service worker integration. Finally, the **Fetch API** section defines the JavaScript classes (`Headers`, `Request`, `Response`) that expose these low-level operations to developers.

# Repeated Or Central Claims

- **Security by Default**: The standard prioritizes security over legacy compatibility. This is evident in strict MIME type extraction (ignoring dangerous parameters), blocking requests on "bad ports" (e.g., telnet, ftp), enforcing CORS preflight checks for unsafe methods, and using `nosniff` directives to prevent content-type sniffing attacks.
- **Connection Pooling**: Connections are not treated as generic resources but are strictly partitioned by a "Network Partition Key" derived from the site origin and credentials (secure/insecure). This ensures that cookies and other sensitive data do not leak between different origins or security contexts.
- **Response Tainting**: Responses are categorized into three distinct states: `basic`, `cors`, and `opaque`. This tainting mechanism dictates whether headers are exposed to client JavaScript, whether the body is readable, and how redirects are handled.
- **CORS as an Exception**: Cross-Origin Resource Sharing is not a default behavior; it requires explicit opt-in via headers (`Access-Control-Allow-Origin`). The protocol distinguishes between preflight requests (checking support) and actual requests, with specific rules for credentials and exposed headers.
- **Caching Complexity**: Caching is sophisticated but constrained by browser limitations. It supports standard modes (`default`, `no-store`, `reload`) and advanced strategies like "stale-while-revalidate" and negative caching (caching network errors). Request bodies must be cloned to support redirects or authentication retries without stream exhaustion.
- **Fetch API Abstraction**: The JavaScript API provides a high-level interface that abstracts the underlying fetch algorithms. However, the classes (`Request`, `Response`) still maintain internal state regarding timing info, controller references, and headers, reflecting the complexity of the underlying network operations.

# Important Local Details

## Connection and Network Logic
- **Network Partition Key**: Composed of a "site" (scheme + host) and an optional implementation-defined second key. This tuple identifies connections and determines HTTP cache partitions.
- **Bad Ports**: Fetching is explicitly blocked if the URL scheme is HTTP(S) and the port matches a list of "bad ports" (e.g., 21, 23).
- **Timing Info**: Connections track timestamps for domain lookup, connection start/end, and secure connection establishment. A "clamp and coarsen" algorithm resets these times to default values to prevent exposing reused connection details to the client.
- **WebTransport**: For WebTransport connections, specific ALPN settings (`SETTINGS_ENABLE_WEBTRANSPORT`) are required on HTTP/3 transports.

## Header Processing
- **Origin Header**: Unlike `Referer`, the `Origin` header omits the path. It is appended based on response tainting ("cors") or request method (non-GET/HEAD) combined with referrer policy. Serialization rules enforce lowercase ASCII schemes and forbid leading zeros in IPv6 addresses.
- **Content-Type Extraction**: The algorithm parses headers to extract MIME types, prioritizing safety. It ignores parameters unless they define the essence or charset. Legacy encodings are extracted as a fallback.
- **CORS Headers**: `Access-Control-Allow-Origin` accepts an origin string or `*`. When credentials are included (`Access-Control-Allow-Credentials: true`), `*` is forbidden for the allow-origin value. `Access-Control-Expose-Headers` can list specific headers or `*` (though `*` only matches a header literally named `*`).

## Fetch Algorithms
- **Main Fetch**: Orchestrates security upgrades (HSTS/SVCB to HTTPS), referrer logic, and response filtering. It branches based on request mode (`navigate`, `same-origin`, `no-cors`) and scheme.
- **Scheme Fetch**: Handles specific URL schemes:
  - `about:blank`: Returns empty body with `OK`.
  - `data:` URLs: Processed via a specific processor returning the serialized MIME type.
  - `blob:` URLs: Only allow `GET`; range requests calculate byte ranges for slicing.
- **Redirect Handling**: Supports manual, follow, and error modes. Redirects update timing info, strip non-wildcard CORS headers on cross-origin jumps, and normalize methods (e.g., POST to GET on 301/302). Max redirect count is 20.
- **CORS Preflight Cache**: Maintains a cache for preflight responses (`OPTIONS` requests) to avoid repeated network calls. Entries are appended if the method/headers match cached entries, subject to TAO (Token Authentication Object) checks.

## API Classes
- **Headers**: Represents a collection of header entries. Supports iteration and appending.
- **Request**: Encapsulates request details including timing info, callbacks, controller references, and body. Handles `BodyInit` unions for body data.
- **Response**: Wraps the internal response object, exposing status, headers, and body. Supports integrity checks (SRI) where the body must match metadata.

# Gaps Or Cautions

- **Implementation Discretion**: The specification intentionally leaves certain details vague to allow implementer discretion. Notable areas include Proxy Auto-Config (PAC) handling, IP address racing logic for connection establishment, and specific DNS timing operations.
- **File Scheme Ambiguity**: The document explicitly notes that `file:` URLs are left as an exercise for the reader and their behavior is "unfortunate" or undefined in the current standard context.
- **MIME Type Safety vs. Legacy**: While the spec defines a strict MIME type extraction model for safety, it acknowledges that existing web platform features have not always followed this pattern, implying potential divergence from historical browser behavior.
- **Partial Content Limitations**: Although HTTP standards allow partial content caching, the text notes that browser implementations do not widely support this, which may lead to unexpected behavior regarding range requests.
- **Referrer Policy Nuance**: The handling of `Origin` headers is heavily dependent on Referrer Policy settings (`no-referrer`, `strict-origin`, etc.). Misconfiguring these policies can inadvertently leak or hide origin information contrary to security expectations.

## group-03

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This group of notes synthesizes the core mechanics of the Fetch specification, covering the lifecycle of an HTTP request from infrastructure setup to response handling. It details algorithms for network errors, CORS preflight logic, deferred fetching quotas, and the specific APIs for `Headers`, `Body`, `Request`, and `Response`. The content further addresses data URL processing, garbage collection semantics, and security headers related to cross-origin resource sharing and content type enforcement.

# Cross-Chunk Summary

The Fetch lifecycle is structured into distinct phases: infrastructure (URLs, HTTP methods/headers), request/response construction (`Headers`, `Request`, `Response`), fetching algorithms (network, redirect, CORS-preflight), and body handling. A critical theme spans chunks 13–14 regarding resource management: the **CORS-preflight cache** is used to minimize redundant network calls, while **deferred fetching** manages memory via strict quotas (640 KiB top-level, 64 KiB concurrent per origin). Security and privacy are woven throughout via headers like `Access-Control-Allow-Origin`, `Timing-Allow-Origin`, and the handling of redirects to prevent secret leakage. Finally, chunk boundaries 18–24 focus on specific URL schemes (`data:`) and the low-level consumption of response bodies into various JavaScript types (Blobs, JSON, Text).

# Repeated Or Central Claims

- **Buffering Limits**: The specification explicitly restricts buffering request bodies to 64 KiB when the source is a stream (`null`) to handle potential resends due to timeouts. This limit applies specifically to network layers where streams are involved.
- **CORS Preflight Caching**: CORS preflight requests (OPTIONS) are cached by the user agent to avoid repeated checks for the same origin, method, and headers. Cache entries include key, origin, URL, max-age, credentials mode, method, and header name.
- **Quota Management**: Deferred fetching utilizes a hierarchical quota system: 640 KiB top-level, shared quotas for same-origin nested documents, and default 8 KiB for cross-origin nested documents (configurable via `Permissions-Policy`). Concurrent usage is strictly limited to 64 KiB per reporting origin.
- **Body Consumption**: Response bodies are consumed via algorithms that return promises rejecting with `TypeError` if the body is disturbed or locked. Methods like `formData()` handle multipart parsing, while `json()` and `text()` rely on MIME type detection.
- **Redirect Handling**: Redirects must be handled before origin checks in certain contexts; documents created via redirects only know their final origin after resolution. Quota calculations assume redirects are resolved prior to origin-based checks.

# Important Local Details

- **CORS Check & TAO Check**: The `CORS check` validates `Access-Control-Allow-Origin`, credentials mode, and `Access-Control-Allow-Credentials`. The `TAO check` (Timing-Allow-Origin) validates timing headers against the request origin, including special cases for "nested navigables" to prevent timing info leakage.
- **Headers Guard States**: A `Headers` object maintains a guard state ("immutable", "request", "response", "none", or "request-no-cors"). Validation throws `TypeError` if the name/value is invalid or if the guard is "immutable". In "request-no-cors" mode, privileged headers are removed upon modification by unprivileged code.
- **Request Constructor**: The `new Request(input, init)` constructor defaults to `"cors"` mode for string inputs. Specific referrer values like `"about:client"` and `"no-referrer"` are handled internally. Cloning a request requires the signal to be non-null.
- **Response Static Methods**: Includes `error()` (network error response), `redirect(url, status)` (default 302), and `json(data, init)`. The constructor `new Response(body, init)` initializes headers and body, throwing errors for invalid status codes or null bodies at certain status ranges.
- **Data URL Structure**: A `data:` URL struct consists of a MIME type and a body byte sequence. The processor parses the scheme, validates the MIME type (handling base64 markers), decodes the body (percent-decoding or base64), and returns failure if parsing fails.

# Candidate Wiki Hints

- **Topic: CORS Preflight Mechanism** – Explain how `OPTIONS` requests are constructed, cached to reduce overhead, and validated against server responses including `Access-Control-Allow-*` headers.
- **Topic: Fetch Body Buffering** – Detail the 64 KiB buffer constraint for stream-based request bodies in network layers and the implications of exceeding this limit during resends.
- **Topic: Deferred Fetching Quota** – Break down the quota hierarchy (Top-level -> Same-origin children -> Cross-origin children), explain `Permissions-Policy` usage, and provide examples of request sizes triggering exhaustion.
- **Topic: Response Interface & Consumption** – Document the `Response` attributes, `ResponseType` enum, static methods (`error`, `redirect`, `json`), and algorithms for consuming bodies into `ArrayBuffer`, `Blob`, `Text`, or `FormData`.
- **Topic: Request Construction Options** – Detail the `RequestInit` dictionary, default modes, referrer handling, abort signal usage, and duplex mode restrictions.

# Gaps Or Cautions

- **Multipart Parsing Approximation**: The `formData()` implementation is noted as a "rough approximation," with detailed parsing specification pending, particularly regarding `_charset_` parameters and default content types.
- **WebSocket Standard Shift**: While the text notes that fetch used to define WebSocket connections directly, these are now defined in the WebSockets standard; fetch mode "websocket" still exists but refers to the separate spec.
- **Null Body Status Codes**: The specification includes status codes 101 and 103 in "null body status" for validation purposes, even though they are typically used elsewhere (e.g., switching protocols). This may confuse developers expecting these codes to allow a body.
- **Realm Guarding Differences**: Different static methods and constructors use different realm guards ("immutable" vs "response"), affecting whether the response can be modified by scripts. The `error()` method returns an immutable response, whereas standard construction might return a "response" guard.
- **Service Worker Origins**: A request can have an origin different from the current client when handled by a service worker, specifically for navigation requests, which may impact how origins are tracked for quota or security policies.

## group-04

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This group of notes aggregates the concluding sections of the Fetch Standard specification. It covers the transition from core fetch algorithms to specific scheme handling (`data:` URLs), followed by comprehensive legal metadata (copyrights, licenses), a glossary of referenced terms, and extensive browser compatibility matrices for the Fetch API interfaces (`Request`, `Response`, `Headers`). The notes span from administrative headers regarding authorship and licensing to technical definitions of attributes, methods, and version-specific support across major engines.

# Cross-Chunk Summary

The progression of these chunks moves from high-level legal and structural metadata (Chunks 19) into the specific implementation details of `data:` URL handling (Chunks 20–24). While Chunks 20–23 focus on definitions, attributes, and methods for the Fetch API generally (including `Request`, `Response`, and initialization dictionaries), Chunk 24 specifically narrows the scope to compatibility tables for `data:` URLs and specific response properties. The content synthesizes the distinction between the "Living Standard" text (CC BY 4.0) and source code portions (BSD 3-Clause), defines the interface structures including enums like `RequestMode` and `RequestCache`, and provides a reference for browser support ranging from modern Chromium/Firefox/Safari to legacy environments like Internet Explorer and Android WebViews.

# Repeated Or Central Claims

- **Interface Definitions:** The `Request` and `Response` interfaces are central, featuring read-only attributes (e.g., `url`, `type`, `status`) and methods (e.g., `clone`, `json`, `text`). Initialization dictionaries (`RequestInit`, `ResponseInit`) allow configuration of behavior such as caching, redirection, and credentials.
- **Data URL Handling:** `data:` URLs are treated as a distinct resource type within the Fetch API, processed via a specific "data: URL processor" rather than standard network traversal. They support MIME types like `text`, `video`, and `xslt`.
- **Security and Origin:** The specification references security contexts, CORS protocols (including preflight checks), and origin headers. Specific flags like `use-CORS-preflight` and `unsafe-request` dictate handling logic for cross-origin or local resources.
- **Browser Compatibility:** A recurring theme is the variation in support across engines. Modern browsers generally share features ("In all current engines"), but legacy browsers (IE, Edge Legacy) and specific mobile environments show significant gaps or require higher version numbers for features like `signal` or `cache`.

# Important Local Details

- **Licensing Distinction:** The specification text is licensed under Creative Commons Attribution 4.0 International (CC BY 4.0), while source code portions incorporating the standard are licensed under the BSD 3-Clause License.
- **Authorship and Copyright:** Anne van Kesteren (Apple) is identified as the author, with copyright held by WHATWG (representing Apple, Google, Mozilla, Microsoft).
- **Versioning:** The document distinguishes between the "Living Standard" (current version) and a separate "Patent-Review Version."
- **Attribute Details:**
  - `Request` attributes include `destination`, `referrer`, `mode`, `credentials`, `cache`, `redirect`, `integrity`, `keepalive`, `signal`, and `duplex`.
  - `Response` attributes include `type`, `url`, `status`, `ok`, `headers`, and methods like `clone()`.
  - Enums defined include `RequestDestination` (e.g., `"script"`, `"image"`), `RequestMode` (`"navigate"`, `"cors"`), and `ResponseType` (`"basic"`, `"opaque"`).
- **Compatibility Thresholds:**
  - Basic `data:` URL support: Firefox 39+, Safari 10.1+, Chrome 42+.
  - `response.json_static`: Requires higher versions (e.g., Firefox 115+).
  - `Headers/Sec-Purpose`: Currently supported only in Firefox 115+.

# Candidate Wiki Hints

- **Page: Fetch API Interfaces**: A reference page detailing the `Request`, `Response`, and `Headers` objects, their attributes, methods, and initialization dictionaries.
- **Page: Data URL Scheme**: Documentation covering the syntax, processing logic, and MIME type handling for embedded resources.
- **Page: Browser Compatibility Matrix**: A comparative table summarizing support for Fetch API features across Firefox, Safari, Chrome, Edge, Node.js, and legacy browsers.
- **Page: Licensing Overview**: Explaining the split between CC BY 4.0 (spec text) and BSD 3-Clause (source code).

# Gaps Or Cautions

- **Legacy Environment Limitations:** Many features are unsupported or marked as "None" in Internet Explorer and certain mobile webviews, which may require polyfills or fallback logic.
- **Streaming Body Caution:** The specification notes that cloning a request (`request.clone()`) shares the body stream; accessing the body after cloning may result in errors if the stream is consumed, making `bodyUsed` checks critical.
- **Header Specificity:** Some headers (e.g., `Sec-Purpose`) are not universally supported, appearing only in specific modern versions of Firefox, which could lead to inconsistencies in cross-browser implementations.
- **Incomplete Reference Data:** While a glossary and IDL index are provided, some terms rely on external specifications (HTML, DOM) that evolve independently, requiring careful linkage to avoid drift.

