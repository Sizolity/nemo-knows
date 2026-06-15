---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk details the specification for handling network errors and establishing connections within the Fetch standard. It covers three primary mechanisms:
1.  **HTTP-network fetch**: The core algorithm for initiating requests, managing connections (WebSocket/WebTransport), handling response headers, buffering request bodies, decoding content, and managing CORS/TAO checks.
2.  **CORS-preflight fetch**: The specific logic for issuing `OPTIONS` requests to validate cross-origin resource sharing policies before main requests.
3.  **CORS-preflight cache**: The structure and management of a user agent's cache used to store preflight results to minimize redundant network calls.

# Local Summary

The document defines the `HTTP-network fetch` algorithm, which orchestrates the lifecycle of a network request. It handles connection acquisition (including WebSocket and WebTransport modes), manages response parsing including interim responses (1xx), and implements logic for buffering request bodies when sources are null (e.g., ReadableStreams). The specification details how to handle content encoding, decompressing data, and tracking decoded vs. encoded sizes. It also outlines the `CORS-preflight fetch` process, which constructs an `OPTIONS` request with specific headers (`Accept`, `Access-Control-Request-Method`, `Access-Control-Request-Headers`) and validates the response against a set of rules regarding allowed methods, headers, and origins. Finally, it describes the `CORS-preflight cache` data structure, which stores entries keyed by network partition key, origin, URL, method, and header name, along with a `max-age` for expiration.

# Key Claims

*   **Request Body Buffering**: When a request body's source is null (indicating it is created from a `ReadableStream`), the user agent must buffer up to 64 kibibytes of the body. If the read exceeds this limit and resending is required (e.g., due to timeout), a network error is returned because the body cannot be recreated.
*   **Interim Response Handling**: For status codes in the range 100–199, the algorithm updates timing information. Specifically for WebSocket upgrades (status 101), the loop breaks immediately after receiving the upgrade response.
*   **CORS Preflight Construction**: The preflight request is a clone of the original request but with the method set to `OPTIONS`, mode set to `"cors"`, and specific headers appended: `Accept: */*`, `Access-Control-Request-Method`, and optionally `Access-Control-Request-Headers`.
*   **Cache Entry Structure**: A CORS-preflight cache entry contains a network partition key, byte-serialized origin, URL, max-age, credentials mode (boolean), method (`null`, `*`, or specific method), and header name (`null`, `*`, or specific header).
*   **TAO Check Logic**: The Timing-Allow-Origin (TAO) check returns success if the response contains `Timing-Allow-Origin` with a wildcard `*`, the specific request origin, or if the credentials mode is "basic".

# Entities And Concepts

*   **HTTP-network fetch**: The main algorithm for fetching resources over HTTP.
*   **CORS-preflight cache**: A user agent data structure storing preflight validation results to avoid redundant network requests.
*   **Network Partition Key**: A value used to group connections and determine routing, derived from the request context.
*   **ReadableStream**: Used as the source for request bodies; requires buffering when not backed by a non-null source.
*   **CORS-unsafe request-header names**: Headers that require explicit permission via `Access-Control-Allow-Headers` in preflight responses.
*   **WebDriver BiDi response started steps**: A referenced algorithm step for handling WebDriver-specific response events.

# Procedures And API Details

**HTTP-network fetch Steps:**
1.  Initialize variables: `request`, `response` (null), `timingInfo`, `networkPartitionKey`, `newConnection`.
2.  **Connection Acquisition**:
    *   If mode is `"websocket"`: Obtain a WebSocket connection.
    *   If mode is `"webtransport"`: Obtain a WebTransport connection using the network partition key.
    *   Otherwise: Obtain an HTTP connection using the network partition key, URL, credentials flag, and new connection flag.
3.  **Request Execution Loop**:
    *   Check for connection failure.
    *   Coarsen and clamp timing info.
    *   Set final network-request start time.
    *   Make HTTP request. Handle TLS certificate dialogs (make available in traversable if allowed, else error).
    *   **Body Transmission**: If body is non-null, read incrementally. Define `processBodyChunk` and `processEndOfBody`.
    *   **Response Handling**: While true loop:
        *   Set final network-response start time upon receiving first byte.
        *   Wait for all headers.
        *   Run WebDriver BiDi response started steps.
        *   Handle status codes 100–199 (update interim timing, break on 101 WebSocket upgrade, queue task for 103 early hints).
    *   **Body Decoding**: Loop through transmitted bytes:
        *   Extract `Content-Encoding`.
        *   Determine `filteredCoding` ("@unknown", empty string, "multiple", or specific coding).
        *   Update body info (encoded size, decoded size).
        *   Handle content codings.
        *   Append bytes to internal buffer.
        *   Suspend fetch if buffer exceeds upper limit.
    *   **Completion/Abortion**: Close stream on normal completion or error. If aborted, set aborted flag, error stream, and transmit RST_STREAM for HTTP/2.

**CORS-preflight fetch Steps:**
1.  Create `preflight` request: Clone original URL list, set method to `OPTIONS`, mode to `"cors"`, response tainting to `"cors"`.
2.  Append headers: `Accept: */*`, `Access-Control-Request-Method`, and `Access-Control-Request-Headers` (comma-separated list of unsafe header names).
3.  Execute `HTTP-network-or-cache fetch`.
4.  **Validation**: If status is OK and CORS check returns success:
    *   Extract `Access-Control-Allow-Methods` and `Access-Control-Allow-Headers`.
    *   Handle null methods if `use-CORS-preflight` flag is set.
    *   Validate request method against allowed methods (reject if not allowed, not safe, and credentials are included).
    *   Validate unsafe headers against allowed headers.
    *   Extract `Access-Control-Max-Age`.
    *   Update or create cache entries for matching/non-matching methods and header names with the new max-age.
5.  Return response or network error.

**CORS-preflight Cache Entry Creation:**
*   Initialize entry with: Network partition key, byte-serialized origin, URL, max-age, credentials boolean, method, header name.
*   Append to user agent's CORS-preflight cache list.

**TAO Check Steps:**
1.  Assert request origin is not "client".
2.  Return failure if `timing allow failed` flag is set.
3.  Decode and split `Timing-Allow-Origin`.
4.  Return success if values contain `*` or the request origin.
5.  Return failure if mode is `"navigate"` and origins differ (nested navigable scenario).
6.  Return success if response tainting is `"basic"`.

# Nuance Or Contradictions

*   **Body Source Null**: The specification explicitly distinguishes between request bodies with non-null sources (recreatable) and those with null sources (ReadableStream, unrecreatable). This necessitates the 64 KiB buffer limit for the latter to prevent data loss on connection reset.
*   **Interim Responses**: The algorithm treats responses in the 100–199 range as interim. It specifically notes that these are eventually followed by a "final" response, but the loop logic handles the 101 upgrade case immediately without waiting for the final response if it is a WebSocket upgrade.
*   **Cache Matching Logic**: The definition of a "cache entry match" involves complex boolean logic regarding credentials modes (true vs. false/non-include) and header name wildcards (`*`), which must be carefully evaluated to determine if an existing cache entry can be reused or if a new one must be created/updated.
*   **Content-Encoding Filtering**: The `filteredCoding` variable has multiple states: `"@unknown"`, empty string, `"multiple"`, or a specific lowercase coding from the registry. This reflects the complexity of handling multiple encodings or unknown/unsupported ones during decompression.

# Candidate Wiki Hints

*   **Page: CORS Preflight Cache Structure**
    *   *Topic*: The internal data structure used by browsers to store preflight validation results.
    *   *Key Content*: Definition of cache entry fields (key, origin, URL, max-age, credentials, method, header name) and the logic for creating or updating entries based on HTTP response headers.
*   **Page: Fetch Body Buffering Limits**
    *   *Topic*: Handling request bodies that are ReadableStreams without a backing source.
    *   *Key Content*: Explanation of the 64 KiB buffer requirement, why it is needed (unrecreatability), and the conditions under which exceeding this limit results in a network error.
*   **Page: Timing-Allow-Origin (TAO) Validation**
    *   *Topic*: The logic determining whether timing information can be shared across origins.
    *   *Key Content*: Steps for parsing `Timing-Allow-Origin` headers, handling wildcards, checking request origins, and special cases for "basic" response tainting and nested navigables.
