---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context
This chunk details the algorithms for handling network errors and CORS-related fetches within the Fetch specification. It covers HTTP-network fetch mechanics, including connection management, body buffering, and response parsing. It further defines CORS-preflight logic, caching strategies, and specific checks for CORS permissions (CORS check) and Timing-Allow-Origin (TAO check).

## Local Summary
The section outlines the `HTTP-network fetch` algorithm, detailing steps for establishing connections, handling HTTP/1.x and HTTP/2 specifics, managing request/response bodies via streams, and processing headers. It then transitions to `CORS-preflight fetch`, explaining how preflight requests are constructed and validated against server responses. Finally, it describes the structure and matching logic of the `CORS-preflight cache` and the specific algorithms for `CORS check` and `TAO check`.

## Key Claims
- The user agent may buffer up to 64 kibibytes of a request body if the source is null (stream-based) to handle potential resends due to timeouts.
- CORS-preflight requests are effectively `OPTIONS` requests with specific headers appended to check protocol understanding and populate a cache.
- A CORS-preflight cache entry includes key, origin, URL, max-age, credentials mode, method, and header name.
- The `CORS check` algorithm verifies the `Access-Control-Allow-Origin`, credentials mode, and `Access-Control-Allow-Credentials` headers.
- The `TAO check` validates the `Timing-Allow-Origin` header to determine if timing information can be exposed.

## Entities And Concepts
- **HTTP-network fetch**: The core algorithm for making network requests.
- **CORS-preflight fetch**: A specific type of request (OPTIONS) used to validate cross-origin resource sharing permissions before a main request.
- **CORS-preflight cache**: A data structure storing preflight results to minimize redundant requests.
- **CORS check**: Logic to determine if a response is valid for the requesting origin and credentials mode.
- **TAO check**: Logic validating the `Timing-Allow-Origin` header against the request origin.
- **Network Partition Key**: Used to group connections and manage network state.
- **Content-Encoding**: Headers indicating compression or encoding of the response body.

## Procedures And API Details
### HTTP-network Fetch Steps
1.  **Connection Acquisition**: Based on mode ("websocket", "webtransport", or standard), obtain a connection using the network partition key, URL, and credentials settings.
2.  **HTTP Request Execution**:
    -   Follow HTTP requirements and caching rules.
    -   Buffer request body (up to 64 KiB) if source is null.
    -   Track timing info for response start.
3.  **Response Handling**:
    -   Parse status codes; handle interim responses (1xx) appropriately.
    -   If TLS client certificate dialog occurs, make it available in the navigable or return error.
4.  **Body Transmission**:
    -   Queue tasks for processing request body end-of-body and chunks.
    -   Handle content decoding (`Content-Encoding`) and update size counters (encoded vs. decoded).
5.  **Stream Setup**: Create a `ReadableStream` for the response body with pull and cancel algorithms.
6.  **Cleanup**: Close connections appropriately, transmitting RST_STREAM frames if using HTTP/2 upon abort or error.

### CORS-preflight Fetch Steps
1.  **Request Construction**: Create a new request with method `OPTIONS`, cloning URL list, setting mode to "cors", and adding `Accept: */*` and `Access-Control-Request-*` headers.
2.  **Execution**: Run `HTTP-network-or-cache fetch`.
3.  **Validation**: If status is OK and CORS check succeeds:
    -   Extract `Access-Control-Allow-Methods`, `Access-Control-Allow-Headers`, and `Access-Control-Max-Age`.
    -   Update or create cache entries for methods and headers based on the response.
4.  **Failure**: Return a network error if checks fail.

### Cache Entry Creation
-   Initialize entry with network partition key, byte-serialized origin, URL, max-age, credentials boolean, method, and header name.
-   Append to the user agent's CORS-preflight cache list.

## Nuance Or Contradictions
- **Buffering Limit**: The specification explicitly allows buffering only 64 KiB of a request body when the source is null (stream), necessitating error return if reading beyond this limit during a resend scenario.
- **Cache Matching Logic**: Cache entry matching depends on specific conditions regarding credentials mode (true vs. false/non-"include") and header name wildcards (`*`) versus non-wildcard headers.
- **TAO Check Context**: The TAO check includes a special case for "nested navigables" where the request origin differs from the container document's origin, potentially causing failure to prevent timing info leakage to the container.

## Candidate Wiki Hints
-   **Topic: CORS Preflight Mechanism** – Explain how `OPTIONS` requests are constructed and cached to reduce overhead.
-   **Topic: Fetch Body Buffering** – Detail the 64 KiB buffer constraint for stream-based request bodies in network layers.
-   **Topic: CORS Cache Entry Structure** – Define the fields of a cache entry (key, origin, URL, max-age, etc.) and matching rules.
