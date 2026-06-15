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
