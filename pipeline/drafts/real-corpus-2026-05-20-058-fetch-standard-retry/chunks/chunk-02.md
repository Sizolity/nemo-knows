---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context
This chunk covers HTTP header handling within the Fetch Standard, specifically focusing on header lists, structured field values, parsing algorithms, and CORS/security categorizations. It also introduces status code definitions. The text spans sections 2.2.2 (Headers) and 2.2.3 (Statuses).

## Local Summary
The specification defines HTTP headers as ordered key-value pairs (multimaps) where duplicate keys are allowed but combined for non-`Set-Cookie` headers exposed to JavaScript. It details algorithms for getting, setting, combining, deleting, and sorting headers. Special attention is given to parsing header values which can contain commas and quotes, requiring specific splitting logic. The chunk further categorizes headers into security buckets (CORS-safelisted, forbidden, privileged) and defines status codes as integers with specific ranges for OK, Redirect, Range, etc.

## Key Claims
- HTTP headers are referred to as "fields" in the spec but "headers" colloquially on the web platform.
- Header lists act as specialized multimaps; keys are case-insensitive matches.
- `Set-Cookie` headers are treated separately due to inability to combine values and client-side complexity.
- Header values are byte sequences, not native objects, in the current Fetch implementation (RFC9651 noted for future object support).
- Parsing header values involves isomorphic decoding and splitting on commas/quotes while handling whitespace.
- CORS categorization relies on value length limits (<128 bytes) and specific allowed character sets or MIME types.
- Forbidden headers include standard proxy/control headers (`Host`, `Connection`) and those starting with `proxy-` or `sec-`.
- Status codes are integers 0–999, with specific ranges for OK (2xx), Redirect (3xx), etc.

## Entities And Concepts
- **Header List**: An ordered list of key-value pairs; essentially a multimap.
- **Structured Field Value**: Objects that can be serialized to byte sequences for headers.
- **CORS-safelisted Request Header**: Headers with length <128 bytes and restricted character sets (e.g., `Accept`, `Content-Type`).
- **Forbidden Request Header**: Headers like `Cookie`, `Host`, `Origin`, or those prefixed with `proxy-`/`sec-`.
- **Privileged No-CORS Request Header**: Headers like `Range` that can be set by privileged APIs.
- **No-CORS-safelisted Request Header**: A subset of request headers allowed in CORS contexts (e.g., `Accept`, `Content-Type`).
- **Status Code**: An integer 0–999 representing HTTP response status (OK, Redirect, etc.).

## Procedures And API Details
### Getting a Structured Field Value
1. Assert type is "dictionary", "list", or "item".
2. Get name from list; if null, return null.
3. Parse structured fields with input_string and header_type.
4. Return result or null if parsing fails.

### Setting a Structured Field Value
1. Serialize the structured value.
2. Set (name, serializedValue) in list.

### Getting Header Name (Combining Values)
1. If name not in list, return null.
2. Return values of all headers with matching name, separated by `0x2C 0x20` (comma-space).

### Parsing Header Value (`get, decode, and split`)
1. Isomorphic decode the value.
2. Collect code points that are not `"` or `,`.
3. If `"`, collect HTTP quoted string.
4. Trim whitespace.
5. Split by `,` if present.

### Sorting and Combining Header List
1. Convert all names to a sorted-lowercase set.
2. Iterate through names:
   - If `set-cookie`, append all values in order.
   - Otherwise, combine into single value or append if missing.

### Determining CORS-Safelisted Request Header
1. Check value length > 128 bytes -> false.
2. Byte-lowercase name and check specific headers (`accept`, `content-type`, etc.):
   - Verify character set (alphanumeric, specific symbols).
   - For `content-type`, ensure MIME essence is `application/x-www-form-urlencoded`, `multipart/form-data`, or `text/plain`.
3. Return false for others.

### Forbidden Request Header Check
Returns true if name matches:
- `Accept-Charset`, `Accept-Encoding`, `Connection`, `Cookie`, `Host`, `Origin`, etc.
- Names starting with `proxy-` or `sec-`.
- X-HTTP-Method headers with forbidden methods.

## Nuance Or Contradictions
- **Header Value Definition**: Defined as byte sequences without leading/trailing whitespace or NUL/newlines, differing from standard ABNF `field-value` production for compatibility with deployed content.
- **CORS Safelisting Exceptions**: The algorithm does not use `extract a MIME type` due to its forgiving nature; naive servers might misinterpret request bodies if strict parsing isn't used.
- **Range Header Parsing**: Supports omitted ends/starts (e.g., `bytes=0-`, `bytes=-500`) but notes browsers historically omit ranges like `bytes=-500`.
- **Status Code Edge Cases**: Mapping HTTP/1 status codes to the Fetch concept involves edge cases tracked in issue #1156.

## Candidate Wiki Hints
- **HTTP Header Lists and Multimaps**: Explain how headers are stored and combined, especially regarding `Set-Cookie` exceptions.
- **Structured Fields in Fetch**: Discuss the transition from byte sequences to objects (RFC9651) and serialization requirements.
- **CORS Header Categorization**: Detail the logic behind safelisted, forbidden, and privileged headers, including character set restrictions.
- **Parsing HTTP Values**: Break down the algorithm for splitting comma-separated values with quoted strings and whitespace handling.
- **Status Code Definitions**: Define ranges for OK, Redirect, Range, and Null Body statuses.
