---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

### Chunk Context
This chunk covers **Section 2.2.2 Headers** within the Fetch Standard, detailing HTTP header management in the browser environment. It defines headers as ordered lists of key-value pairs (multimaps), explaining serialization/parsing logic, splitting algorithms for comma-separated values, and specific handling for `Set-Cookie` headers. The section concludes with a transition to **Section 2.2.3 Statuses**, defining HTTP status code ranges (e.g., OK, Redirect) and null body statuses.

### Local Summary
The document establishes that an HTTP header list is an ordered collection of name-value pairs where duplicate keys are permitted but non-`Set-Cookie` headers are combined into single values when exposed to JavaScript. It details the algorithms for getting, setting, deleting, and combining headers, including a specific sorting mechanism that preserves `Set-Cookie` entries individually while merging others. The text also defines "structured field values" (currently byte sequences) and provides rigorous rules for splitting header values by comma while respecting quoted strings. Finally, it outlines CORS safety lists, forbidden request/response headers, and the classification of HTTP status codes.

### Key Claims
- **Header Representation**: A header list is a specialized multimap; implementations may optimize storage for non-`Set-Cookie` headers since they are combined on the client side.
- **Structured Fields**: Currently, Fetch supports header values only as byte sequences; future versions might preserve object structures end-to-end (RFC9651).
- **Splitting Logic**: Header values are split by comma (`U+002C`) unless enclosed in double quotes (`U+0022`), which allows for complex multi-value headers like `Accept`.
- **CORS Safety**: A header is "CORS-safelisted" only if it passes length checks (128 bytes) and contains only safe characters or specific MIME types. Headers starting with `Sec-` are reserved for future-safe APIs.
- **Forbidden Headers**: Specific headers like `Accept`, `Cookie`, `Host`, and those prefixed with `proxy-` or `sec-` are forbidden on requests to prevent user agent bypassing controls.

### Entities And Concepts
- **Header List**: An ordered list of zero or more headers (key-value pairs).
- **Structured Field Value**: Objects that HTTP can serialize efficiently; currently represented as byte sequences in Fetch.
- **CORS-Safelisted Request Header**: Headers allowed to be sent with cross-origin requests, subject to length and character restrictions.
- **Forbidden Request Header**: Headers restricted to maintain user agent control (e.g., `Cookie`, `Host`, `Content-Length`).
- **Range Status**: HTTP status codes 206 or 416 indicating partial content or range not satisfiable.
- **Null Body Status**: Status codes 101, 103, 204, 205, or 304 where the response body is empty or irrelevant.

### Procedures And API Details
- **Getting a Header**: Returns `null` if absent; otherwise returns all values separated by `0x2C 0x20`.
- **Parsing Structured Fields**: Converts byte sequences into structured objects (currently limited to byte sequence representation).
- **Sorting Headers**: Converts header names to a sorted-lowercase set, preserving individual `Set-Cookie` entries while combining others.
- **Building Content Range**: Constructs the `bytes <start>-<end>/<fullLength>` string for partial content requests.
- **CORS Safelist Check**: Validates header length and character sets; rejects headers with unsafe bytes (e.g., `<`, `>`, `:`) or disallowed MIME types.

### Nuance Or Contradictions
- **Set-Cookie Handling**: Unlike other headers, `Set-Cookie` cannot be combined and requires complex handling in the Headers object to avoid leaking this complexity into requests; it is semantically a response header but appears in request contexts as forbidden.
- **MIME Type Parsing**: The standard explicitly avoids using the "extract a MIME type" algorithm because it is too forgiving, preventing servers from misinterpreting request bodies (e.g., treating a `text/plain` body as JSON).
- **Range Syntax**: While `bytes=-500` is syntactically valid per the range parsing rules, browsers historically do not emit such ranges, so they are excluded from the safelist.

### Candidate Wiki Hints
- **Topic: HTTP Headers in Fetch API** – Explaining how headers are stored as ordered lists and combined for client-side access.
- **Topic: CORS Header Safety** – Detailing the logic behind `CORS-safelisted` vs. forbidden request headers and byte restrictions.
- **Topic: Parsing Multi-Value Headers** – Algorithms for splitting comma-separated values while handling quoted strings (e.g., `Accept`, `Link`).
- **Topic: Range Requests** – Implementation details of parsing `Range` headers and constructing content range responses.
