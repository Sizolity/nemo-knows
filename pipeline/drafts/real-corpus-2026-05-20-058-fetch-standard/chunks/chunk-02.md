---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context
This chunk details the **Headers** infrastructure within the Fetch Standard (Section 2.2.2), covering header list structures, parsing/splitting algorithms, CORS safety logic, forbidden headers, and status code definitions. It transitions into Section 3 regarding HTTP Statuses at the end.

## Local Summary
The specification defines a `header list` as an ordered multimap of key-value pairs. It provides algorithms for getting, setting, deleting, and combining headers, with special handling for `Set-Cookie`. The text details the complex logic for determining CORS-safelisted request/response headers and lists specific forbidden header names. Finally, it defines status codes (null body, ok, redirect) as integer ranges.

## Key Claims
- A header list is essentially a specialized multimap where keys are byte-case-insensitive matches.
- `Set-Cookie` headers are treated uniquely; they cannot be combined and require complex handling in the Headers object, so they are forbidden on requests to avoid leaking complexity.
- Header values must not contain NUL bytes or HTTP newline bytes and must have no leading/trailing HTTP whitespace.
- A header is "CORS-safelisted" only if it passes strict value length checks (e.g., >128 bytes returns false) and contains only safe characters.
- Headers starting with `Sec-` are reserved for future use to ensure they remain safe from developer-controlled APIs like XMLHttpRequest.
- A status code is an integer in the range 0 to 999 inclusive.

## Entities And Concepts
- **Header List**: An ordered list of key-value pairs representing HTTP headers; acts as a specialized multimap.
- **CORS-safelisted request-header**: Headers safe for cross-origin requests (e.g., `Accept`, `Content-Type` with restricted MIME types).
- **Forbidden Request Header**: Headers that must not be set by scripts (e.g., `Host`, `Origin`, `Cookie`).
- **Range Header**: Used for partial content retrieval; parsed into start/end values.
- **Status Code**: Integer representation of HTTP response status (0–999).

## Procedures And API Details
### Getting a Structured Field Value
1. Assert type is "dictionary", "list", or "item".
2. Get value from list using name.
3. If null, return null.
4. Parse structured fields with input_string set to value.
5. Return result or null if parsing failed.

### Getting, Decoding, and Splitting a Header Name
1. Call `getting name` on the list.
2. If null, return null.
3. Apply `getting, decoding, and splitting` to the value.

### Checking CORS Safety for Request Headers
1. Check if value length > 128; if so, return false.
2. Byte-lowercase name and check against safelist:
   - `accept`: Must not contain unsafe bytes.
   - `content-type`: Must be one of `application/x-www-form-urlencoded`, `multipart/form-data`, or `text/plain`.
3. Return true only if all checks pass; otherwise false.

### Parsing Range Header Values
1. Ensure value starts with "bytes".
2. Parse start and end values (allowing omission for suffix ranges like `bytes=-500`).
3. Validate that start is not greater than end if both are present.

## Nuance Or Contradictions
- **Parsing vs. Presence**: The algorithm for getting a structured field value does not distinguish between a header being absent and its value failing to parse; both return null. This ensures uniform processing but can hide errors in parsing logic.
- **CORS Safelisting Exceptions**: While most headers are safe, `Content-Type` has limited exceptions documented in CORS protocol exceptions. The specification notes that the standard algorithm does not use `extract a MIME type` because servers are not expected to implement it strictly.

## Candidate Wiki Hints
- **HTTP Headers Structure**: A page explaining the difference between header fields and headers, and how `Set-Cookie` is handled differently in the Fetch API.
- **CORS Header Safelisting**: A guide detailing which headers are allowed in cross-origin requests and why `Content-Type` restrictions exist.
- **Forbidden HTTP Headers**: A list of headers that scripts cannot set (e.g., `Host`, `Connection`) to maintain user agent control.
