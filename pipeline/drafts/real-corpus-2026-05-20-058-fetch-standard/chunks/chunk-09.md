---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context
This chunk covers sections 3.5 through 3.8 of the Fetch Standard, detailing header processing logic for `Content-Type`, `X-Content-Type-Options`, and `Cross-Origin-Resource-Policy`. It concludes with section 4, "Fetching," which outlines the high-level algorithm for initiating network requests, including body handling, early hints, parallel queue usage, and default header population (Accept, Accept-Language). The text notes that sections 3.6 and 3.7 are repeated or expanded in this specific view compared to a standard summary.

## Local Summary
The chunk defines algorithms for extracting MIME types from headers, treating incorrect essence as fatal errors. It details the `X-Content-Type-Options` header logic, specifically how "nosniff" blocks requests destined for scripts or styles if the MIME type does not match expectations. The `Cross-Origin-Resource-Policy` (CORP) section explains how to check origins against response headers (`same-origin`, `same-site`, `cross-origin`) and queue violation reports. Finally, the fetching algorithm initializes request properties, handles preloaded resources, sets default headers like `Accept` based on destination type, and manages priority and client context.

## Key Claims
- **MIME Type Extraction:** Extracting a MIME type returns failure or a type; if the essence is incorrect for the format, it is treated as a fatal error. Parameters can typically be safely ignored during this process.
- **Nosniff Logic:** The `X-Content-Type-Options: nosniff` header requires checking the response's `Content-Type` against the request destination. If the destination is script-like or "style" and the MIME type is invalid or incorrect, the response is blocked.
- **CORP Check:** The `Cross-Origin-Resource-Policy` header allows enforcing policies like "same-origin" or "same-site". It interacts with an embedder policy (e.g., "credentialless", "require-corp") to determine if a request is allowed or blocked, potentially queuing violation reports.
- **Default Headers:** If `Accept` is missing, the user agent should populate it with `*/*` or specific values based on the request destination (e.g., images, JSON). Similarly, `Accept-Language` is populated if the client has an emulated language or appropriate header value exists.

## Entities And Concepts
- **Content-Type Header:** Used to define MIME types; extraction logic handles charset and essence validation.
- **X-Content-Type-Options:** A security header where "nosniff" triggers MIME type checking against script/style destinations.
- **Cross-Origin-Resource-Policy (CORP):** A header used to restrict resource loading based on origin relationships (same-origin, same-site, cross-origin).
- **Fetch Algorithm:** The core process for retrieving resources, handling bodies, parallel queues, and preloaded responses.
- **Accept Header:** Automatically populated if missing, with values derived from the document's Accept header or specific defaults for images/text/JSON.
- **Embedder Policy:** Internal setting (e.g., "unsafe-none", "credentialless") influencing CORP enforcement.

## Procedures And API Details
### Extract a MIME Type
1. Get, decode, and split `Content-Type` from headers.
2. If null, return failure.
3. Parse each value; skip if essence is "*/*".
4. Determine the final MIME type and charset.
5. Return failure or the MIME type (with essence).

### Legacy Extract Encoding
1. If MIME type is failure, return fallback encoding.
2. If no charset parameter exists, return fallback encoding.
3. Get encoding from charset parameter; if failure, return fallback.
4. Return tentative encoding.

### Determine Nosniff
1. Get `X-Content-Type-Options` values.
2. If null or not "nosniff", return false.
3. Return true if first value matches "nosniff" (case-insensitive).

### Cross-Origin Resource Policy Internal Check
1. Handle "unsafe-none" embedder policy by returning allowed for navigation.
2. Get `Cross-Origin-Resource-Policy` header value; normalize to null if invalid or multiple headers exist.
3. Switch on policy and embedder policy value:
    - **null:** Return allowed (or blocked based on specific conditions).
    - **same-origin:** Allow only if origins match exactly.
    - **same-site:** Allow only if schemelessly same site and secure transport rules apply (HTTPS matching).
4. Queue violation reports if checks fail or policy differs from embedder expectations.

### Main Fetch Steps
1. Assert mode is "navigate" or early hints are null.
2. Populate request from client context (origin, global object).
3. Initialize timing info and fetch params.
4. Convert byte sequence body to a body object if applicable.
5. Run WebDriver BiDi clone network request body steps.
6. Check for preloaded resources; update candidate status.
7. Append `Accept` header:
    - Default to `*/*`.
    - If initiator is "prefetch", use document's Accept value.
    - Otherwise, set based on destination (image, json, style, text).
8. Append `Accept-Language` if client has emulated language or header is missing.
9. Set internal priority using request attributes.
10. Run main fetch given fetch params and return controller.

## Nuance Or Contradictions
- **MIME Type Fatal Errors:** The standard states that incorrect MIME essence should be a fatal error, noting that existing web features have historically ignored this pattern, leading to security vulnerabilities. However, it also states parameters can be safely ignored, creating a distinction between the core type (essence) and metadata (parameters).
- **CORP Header Matching:** Multiple `Cross-Origin-Resource-Policy` headers are noted to have the same effect as a single one. Furthermore, a value like "same-site,same-origin" will never match anything if the embedder policy is "unsafe-none", effectively rendering that header combination useless in that context.
- **Secure Transport Matching:** The `Cross-Origin-Resource-Policy: same-site` does not consider a response delivered via secure transport to match a non-secure requesting origin, even if hosts are the same site. This implies strict scheme matching is enforced for "same-site" policies.

## Candidate Wiki Hints
- **Fetch Standard Headers:** A page summarizing security headers like `X-Content-Type-Options` and `Cross-Origin-Resource-Policy`.
- **MIME Type Extraction Algorithm:** Documentation on how the browser parses `Content-Type` headers, handling charsets and essence validation.
- **Preloaded Resources:** Notes on the interaction between `Accept` headers and preloaded resource candidates during fetch initialization.
