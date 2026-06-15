---
title: Chunk 18 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
This chunk covers section "6. data: URLs", detailing the processing algorithm for `data:` URL schemes (based on RFC 2397), followed by background reading on HTTP header layers, atomic redirect handling, CORS safety protocols, WebSocket connections, and practical guidance for setting up requests and invoking fetch operations including response processing strategies.

Local Summary
The section defines the internal structure of a `data:` URL as containing a MIME type and a body byte sequence. It outlines an algorithmic processor that parses the scheme, validates the MIME type (handling base64 encoding markers), decodes the body, and returns a structured object. The text further explains HTTP header layer divisions, warns against exposing redirects to prevent cross-site scripting leaks via secrets in redirect URLs, clarifies safe CORS setups using `Access-Control-Allow-Origin: *`, discusses the role of `Vary` headers for caching CORS responses, describes WebSocket connection establishment as a special fetch mode, and provides guidance on constructing requests (URL, method, body, client settings) and invoking fetch with various response processing callbacks.

Key Claims
- A `data:` URL struct consists of a MIME type and a body byte sequence.
- The `data:` URL processor asserts the scheme is "data", strips the prefix, parses the MIME type, decodes the body (percent-decoding or base64 if indicated), and returns failure if parsing fails.
- If the MIME type starts with ";", it defaults to "text/plain;charset=US-ASCII".
- Exposing redirects can leak secrets; a redirect from an authenticated URL to a new URL containing a secret should not be exposed via APIs.
- Using `Access-Control-Allow-Origin: *` is safe for resources protected by IP or firewall, as it shares the resource with tools like curl and wget but does not reveal auth/cookie info if the resource itself isn't accessible from random devices.
- If CORS requirements are complex (not just static origin or "*"), the `Vary: Origin` header must be used to prevent caching non-CORS responses for subsequent CORS requests.
- WebSocket connections are established via a special fetch with mode "websocket".
- Fetch callbacks can handle response processing upon completion, chunk-by-chunk streaming, or ignoring the response entirely (e.g., `navigator.sendBeacon`).

Entities And Concepts
- data: URL struct
- MIME type
- body (byte sequence)
- isomorphic decode
- forgiving-base64 decode
- Access-Control-Allow-Origin
- Vary header
- WebSocket object
- fetch controller
- processResponseConsumeBody
- processResponse
- processResponseEndOfBody
- Sec-Fetch-Dest
- Content Security Policy (CSP)
- Referrer Policy

Procedures And API Details
**Data: URL Processor Steps:**
1. Assert scheme is "data".
2. Serialize URL excluding fragment.
3. Remove leading "data:".
4. Collect MIME type characters until comma (U+002C).
5. Strip ASCII whitespace from MIME type.
6. If position reaches end, return failure.
7. Advance position by 1.
8. Remainder is encodedBody.
9. Percent-decode encodedBody to get body.
10. Check if mimeType ends with ";", optional spaces, and "base64" (case-insensitive).
    - If yes: isomorphic decode body, then forgiving-base64 decode. Return failure if base64 decode fails. Remove trailing ";", spaces, and "base64" from mimeType.
11. If mimeType starts with ";", prepend "text/plain".
12. Parse mimeType into mimeTypeRecord. Default to "text/plain;charset=US-ASCII" on failure.
13. Return new data: URL struct with parsed MIME type and decoded body.

**Request Setup Guidance:**
- Set request's URL and method (e.g., POST/PUT with byte sequence or ReadableStream body).
- Choose destination per table to affect Content Security Policy (`Sec-Fetch-Dest`).
- Set client to environment settings object (or null for background fetching).
- Set mode to "same-origin" if same-origin required, otherwise "cors" for web-exposed features.
- Set credentials mode to "include" for cross-origin requests requiring credentials.
- Pass initiator type for Resource Timing reporting.
- Set header list for custom headers (be aware of CORS-preflight implications).
- Set cache mode if overriding default caching.
- Set redirect mode to "error" if redirects should not be followed.

**Fetch Invocation:**
- Pass `processResponseConsumeBody` to read entire body upon completion (returns null, failure, or byte sequence).
- Pass `processResponse` to handle headers and stream body chunk-by-chunk.
- Pass `processResponseEndOfBody` to handle after full download.
- Use fetch controller to abort, report timing, or handle manual redirects.

Nuance Or Contradictions
- The text notes that fetch used to define obtaining/establishing WebSocket connections directly, but these are now defined in the WebSockets standard.
- Handling "no-cors" responses requires caution; while the body is passed as a byte sequence, it should not be directly exposed to scripts in the embedding document (security restriction).
- Caching behavior differs significantly based on whether `Vary: Origin` is used; without it, cached non-CORS responses lack `Access-Control-Allow-Origin`, breaking subsequent CORS requests.

Candidate Wiki Hints
- data-url-processing-algorithm
- http-header-layers-fetch
- cors-safety-and-vary-header
- websocket-fetch-mode
- fetch-response-callbacks
