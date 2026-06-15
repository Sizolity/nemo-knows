---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- **Source**: raw/web/corpus-2026-05-18/058-fetch-standard.md
- **Chunk ID**: 12 of 24
- **Line Range**: 3798–4083
- **Heading**: 1. Set response to a network error. > 4.6. HTTP-network-or-cache fetch

Local Summary
This section defines the `HTTP-network-or-cache fetch` algorithm, which orchestrates fetching resources over HTTP while handling caching strategies (validation, stale-while-revalidate), authentication prompts, and specific header manipulations required by the Fetch standard. It details how to construct request headers (including cookies, authorization, and cache directives), manage connection limits for keepalive requests, handle 401/407 responses via user prompts, and update or invalidate cached responses based on HTTP status codes like 304 and 200–399.

Key Claims
- Browser caches do not widely support partial content caching, though some implementations might attempt it per HTTP Caching standards.
- Request bodies must be cloned before potential redirects or authentication steps to prevent failure if the original stream is consumed.
- If the sum of a request's body length and active keepalive bytes exceeds 64 kibibytes, the fetch returns a network error to prevent indefinite memory usage.
- Specific cache modes (e.g., `no-store`, `reload`) automatically append `Pragma: no-cache` or `Cache-Control: no-cache` headers if missing.
- Range requests require `Accept-Encoding: identity` to avoid server failures when handling partial encoded content.
- 401 and 407 responses trigger user prompts for credentials; subsequent fetches reuse stored authentication entries unless explicitly overridden.
- Successful 2xx responses on unsafe methods invalidate stored cached responses in the HTTP cache partition.

Entities And Concepts
- **HTTP-network-or-cache fetch**: The main algorithm combining network fetching with cache logic.
- **httpFetchParams / httpRequest**: Internal parameters and cloned request objects used during the fetch process.
- **includeCredentials**: A boolean flag determining if cookies and auth headers are included based on credentials mode and CORS tainting.
- **httpCache**: The HTTP cache partition determined for the current request.
- **revalidatingFlag**: Tracks whether a stored response needs validation via conditional requests (ETag, Last-Modified).
- **storedResponse**: Responses retrieved from or updated within the HTTP cache.
- **isNewConnectionFetch**: A flag indicating whether to force a new TCP connection.
- **WebDriver BiDi**: Protocol steps invoked before sending and after receiving responses.

Procedures And API Details
- **Header Construction**:
  - Append `Content-Length` if the body is non-null; set to `0` for empty bodies on POST/PUT.
  - Add `Referer`, `Origin`, and Fetch metadata headers.
  - Inject `Sec-Purpose: prefetch` if initiator is "prefetch".
  - Ensure `User-Agent` header exists using the environment default.
  - Modify cache-control headers based on `cache mode`:
    - `default`: Convert to `no-store` if conditional headers (`If-Modified-Since`, etc.) are present.
    - `no-cache`: Append `Cache-Control: max-age=0` unless already set.
    - `no-store`/`reload`: Append `Pragma: no-cache` and `Cache-Control: no-cache`.
  - Add `Accept-Encoding: identity` for Range requests.
- **Authentication Handling**:
  - Include cookies if credentials mode is "include" or "same-origin" (with basic tainting).
  - Override includeCredentials to false if Cross-Origin-Embedder-Policy disallows it.
  - Append `Authorization` header from stored auth entries or URL credentials if `isAuthenticationFetch` is true.
  - Handle proxy authentication via separate entries, independent of request credentials mode.
- **Caching Logic**:
  - Select stored response if cache mode allows (e.g., `default`, `only-if-cached`).
  - For stale responses in `default` mode, run a parallel revalidation fetch with `no-cache`.
  - On 304 Not Modified: Update stored response headers and mark as "validated".
  - On 2xx unsafe method: Invalidate cached copies.
  - Store new responses or network errors (negative caching) in the cache if mode permits.

Nuance Or Contradictions
- **Partial Content Caching**: While HTTP Caching supports partial content caching, browser caches generally do not implement this widely.
- **Header Duplication**: The algorithm explicitly forbids appending headers already present in the list to avoid duplication issues.
- **Body Consumption Risk**: Redirects and authentication fail if the request body stream is teed unnecessarily; cloning is mandatory when the source is null or potentially consumed.
- **Proxy Auth Testing**: Multiple `Proxy-Authenticate` headers and parsing edge cases are noted as needing testing, indicating potential non-normative behavior in current implementations.
- **Range Header Semantics**: Servers often ignore `Range` headers if a non-identity encoding is accepted, which contradicts strict HTTP expectations but is noted as a common server mistake.

Candidate Wiki Hints
- **HTTP Fetch Algorithm Details**: A page explaining the specific steps of `HTTP-network-or-cache fetch`, including cloning logic and header construction.
- **Caching Strategies in Fetch**: Documentation on how different cache modes (`default`, `no-store`, `reload`) interact with HTTP caching headers and validation flags.
- **Authentication Prompts**: Notes on handling 401/407 responses, credential storage, and the role of `isAuthenticationFetch`.
- **Keepalive Limits**: A technical note on the 64 kibibyte limit for combined content length and inflight keepalive bytes.
