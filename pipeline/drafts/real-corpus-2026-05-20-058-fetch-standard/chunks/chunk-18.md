---
title: Chunk 18 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk covers section **6. data: URLs**, detailing the processing algorithm for `data:` URLs, followed by background information on HTTP header layers, CORS protocols, WebSockets, and general advice on setting up and invoking fetch operations.

# Local Summary

The document defines a `data:` URL as a struct containing a MIME type and a byte sequence body. It outlines a step-by-step processor that validates the scheme, parses the MIME type (handling base64 encoding), and decodes the body. The text further explains HTTP header layer divisions for fetching, safety considerations regarding redirects and CORS, the role of `Vary` headers in caching, and how WebSockets interact with fetch. Finally, it provides guidance on constructing requests, handling response bodies (completion vs. streaming), and manipulating ongoing fetch operations.

# Key Claims

- A `data:` URL struct consists of a MIME type and a body (byte sequence).
- The `data:` URL processor asserts the scheme is "data" and strips the leading prefix before parsing.
- If a MIME type ends with `;base64`, the body is isomorphically decoded and then forgiving-base64 decoded.
- Redirects are not exposed to APIs to prevent leaking secrets via cross-site scripting attacks.
- Using `Access-Control-Allow-Origin: *` is safe for resources accessible via curl/wget but unsafe for those protected by IP authentication or firewalls.
- The `Vary: Origin` header must be used if CORS requirements depend on the origin to prevent caching non-CORS responses for subsequent CORS requests.
- WebSockets initiate a special fetch with mode "websocket" and share policy decisions like HSTS.

# Entities And Concepts

- **data: URL**: A URL scheme carrying data directly in the URL string, defined by RFC 2397.
- **MIME Type**: Part of the `data:` URL struct indicating the media type of the body.
- **Fetch Controller**: An object returned by `fetch` used to abort, report timing, or handle manual redirects.
- **CORS (Cross-Origin Resource Sharing)**: A protocol mechanism allowing web pages on one domain to access resources from another domain via specific headers.
- **Vary Header**: Used in HTTP caching to indicate which request headers affect the response selection (e.g., `Vary: Origin`).
- **WebSockets**: A communication protocol that allows full-duplex communication over a single TCP connection, integrated with fetch via mode "websocket".

# Procedures And API Details

### Data: URL Processing Steps
1. Assert scheme is "data".
2. Serialize the URL excluding fragments.
3. Remove leading "data:" prefix.
4. Collect code points for MIME type (excluding commas).
5. Strip whitespace from MIME type.
6. If position reaches end, return failure.
7. Advance position; remainder becomes encoded body.
8. Percent-decode the encoded body.
9. Check if MIME type ends with `;base64` (case-insensitive):
    - Isomorphic decode the body string.
    - Forgiving-base64 decode the result.
    - Remove " base64" and the semicolon from MIME type.
10. If MIME type starts with ";", prepend "text/plain".
11. Parse MIME type; if failure, default to `text/plain;charset=US-ASCII`.
12. Return struct with parsed MIME type and decoded body.

### Request Setup Guidance
- Set URL and method (`POST`/`PUT` require a body).
- Choose destination (affects CSP and `Sec-Fetch-Dest`).
- Set client to environment settings object.
- Set mode: "same-origin" or "cors" (default for web-exposed features).
- Set credentials mode ("include") if cross-origin credentials are needed.
- Pass initiator type for Resource Timing reporting.
- Set header list for custom headers (may trigger CORS preflight).
- Set cache mode (e.g., "default", "no-store").
- Set redirect mode (e.g., "error" to disable redirects).

### Invoking Fetch Callbacks
- **processResponseConsumeBody**: Called upon completion; receives response and body contents (`null`, `failure`, or byte sequence).
- **processResponse**: Called when headers are received; caller streams the body. Useful for streaming video/images or handling non-OK status codes.
- **processResponseEndOfBody**: Called after fully reading the response body.
- **processEarlyHintsResponse**: For 103 Early Hints (currently handled by navigations).
- **useParallelQueue**: Set to `true` to run internal operations in a parallel queue, allowing interaction with the main thread via callbacks.

# Nuance Or Contradictions

- **CORS Safety**: The text clarifies that `Access-Control-Allow-Origin: *` is safe only if the resource is already publicly accessible (e.g., via curl/wget). If access requires IP auth or firewalls, CORS is unsafe regardless of headers.
- **Redirect Exposure**: Exposing redirects leaks secrets contained in redirect URLs. Therefore, redirects are not exposed to APIs by design.
- **Caching Behavior**: Without `Vary: Origin`, a user agent might cache a response lacking `Access-Control-Allow-Origin` from a non-CORS request and incorrectly reuse it for a subsequent CORS request, breaking the CORS check.

# Candidate Wiki Hints

- **data: URLs**: A dedicated page explaining the structure of `data:` URLs, their MIME type/body composition, and decoding rules (percent-decoding, base64 handling).
- **Fetch API Security**: Notes on redirect security, CORS header usage, and `Vary` header importance for caching.
- **Request Construction**: A reference guide for setting up fetch requests (URL, method, mode, credentials, cache mode, headers).
- **Response Handling**: Strategies for consuming responses fully versus streaming them chunk-by-chunk using callback arguments.
