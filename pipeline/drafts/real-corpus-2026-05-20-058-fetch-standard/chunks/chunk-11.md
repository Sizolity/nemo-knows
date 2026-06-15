---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
# Chunk Context
This chunk details the algorithms for **Scheme Fetch** (handling `about`, `blob`, `data`, `file` schemes) and **HTTP fetch**, including service worker interactions, CORS preflight logic, cache handling, redirect management, and environment determination. It also outlines the steps for **HTTP-redirect fetch**.

# Local Summary
The section defines how a User Agent (UA) processes requests based on their URL scheme. For `blob` URLs, it handles range requests by slicing the blob object, adjusting byte ranges to account for inclusive vs. exclusive boundaries in the slice algorithm. The HTTP fetch algorithm manages service worker interception, timing info updates, and various error conditions (CORS failures, TAO checks, cross-origin resource policy). It explicitly defines conditions under which redirects result in network errors or require method changes (e.g., POST to GET on 301/302).

# Key Claims
- **Blob Range Requests**: A range header denotes an inclusive byte range, whereas the `slice` algorithm expects a different boundary convention; specifically, `rangeEnd` must be incremented before passing it to the slice operation.
- **Service Worker Interception**: If a request's service-workers mode is "all", the UA clones the request, pipes its body through a TransformStream (handling cancellation), and invokes `handle fetch for`.
- **CORS Preflight Caching**: The UA checks a cache for method/header matches before performing a CORS-preflight fetch. This minimizes redundant fetches to ensure resources are familiar with the CORS protocol.
- **Redirect Method Normalization**: If a redirect status is 301 or 302 and the method is `POST`, the request's method is changed to `GET` and the body cleared. Similarly, 303 redirects require the method to be `GET` or `HEAD`.
- **Cross-Origin Header Removal**: When redirecting to a different origin, headers with CORS non-wildcard names (like `Authorization`) are deleted from the request's header list.

# Entities And Concepts
- **Blob URL Entry**: Used to obtain blob objects for `blob:` URLs.
- **TransformStream**: Used in service worker contexts to stream and potentially transform request bodies.
- **CORS-safelisted method**: Methods allowed without preflight checks (e.g., GET, HEAD, POST).
- **TAO check**: Timing Allow Origin check, which can set the "timing allow failed" flag.
- **Cross-origin resource policy**: A security policy that can block responses based on origin and destination.
- **Filtered response**: A response object containing an internal response, used for opaque redirects.

# Procedures And API Details
**Blob Fetch Procedure (Range Request)**
1. Parse the `Range` header to get `(rangeStart, rangeEnd)`.
2. Adjust ranges: If `rangeStart` is null, set it to `fullLength - rangeEnd` and `rangeEnd` to `rangeStart + rangeEnd - 1`.
3. If `rangeEnd` is null or exceeds content, cap it at `fullLength - 1`.
4. **Crucial Step**: Increment `rangeEnd` by 1 before invoking the `slice` algorithm because the slice input range is exclusive on the end.
5. Build a `Content-Range` header and set status to 206 (`Partial Content`).

**HTTP Fetch Service Worker Handling**
1. Clone request if service-workers mode is "all".
2. Pipe body through a TransformStream configured with an algorithm that handles cancellation and chunk enqueueing.
3. Invoke `handle fetch for` on the cloned request.
4. Update timing info with service worker start time.

**HTTP-redirect Fetch Logic**
1. Extract `locationURL`.
2. Check redirect count (max 20).
3. If `cors` mode and destination includes credentials, return network error if origins differ.
4. Normalize method: If status is 301/302 and method is `POST`, switch to `GET`. Delete body headers.
5. Remove CORS non-wildcard headers (e.g., `Authorization`) when crossing origins.
6. Call `main fetch` recursively with updated params.

# Nuance Or Contradictions
- **Range Header Semantics**: There is a specific distinction between the inclusive byte range denoted by the HTTP Range header and the exclusive end boundary expected by the underlying `slice` operation. The text explicitly notes that `rangeEnd` must be incremented to bridge this gap.
- **Redirect Method Changes**: While standard HTTP behavior often normalizes POST to GET on 301/302 redirects, the algorithm explicitly enforces this normalization and clears headers associated with the body. This prevents potential data loss or security issues if a user agent blindly followed a redirect preserving the POST body.
- **Service Worker vs. Network Redirects**: The algorithm distinguishes between redirects coming from the network versus those handled by service workers. Redirects from the network are not exposed to service workers; thus, the mode is switched to "none" for network-originated redirects after the initial fetch response handling.

# Candidate Wiki Hints
- **Blob URL Fetching**: How `blob:` URLs handle range requests and the specific adjustment required for byte ranges in the slice algorithm.
- **CORS Preflight Optimization**: The role of caching in CORS preflight requests to minimize network overhead.
- **HTTP Redirect Normalization**: Rules governing method changes (POST -> GET) and header stripping during redirects.
