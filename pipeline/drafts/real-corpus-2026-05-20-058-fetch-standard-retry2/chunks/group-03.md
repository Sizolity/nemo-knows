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
