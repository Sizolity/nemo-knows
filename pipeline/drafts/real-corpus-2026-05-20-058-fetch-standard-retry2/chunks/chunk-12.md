---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- Heading path: 1. Set response to a network error. > 4.6. HTTP-network-or-cache fetch
- Line range: 3798-4083 (Lines 3798–4083)
- Scope: Defines the `HTTP-network-or-cache` fetch algorithm, handling request cloning, header construction, cache interactions (validation, storage, stale-while-revalidate), credential inclusion, proxy authentication prompts, and negative caching.

Local Summary
This section details how a user agent performs an HTTP fetch that may utilize caching. It specifies steps to prepare the request (cloning bodies if needed), construct headers (including `Content-Length`, `Referer`, `Origin`, `User-Agent`, etc.), manage keep-alive connections within size limits, interact with the HTTP cache (checking for stored responses, sending validation requests if stale, storing new responses or network errors), handle authentication (401/407 status codes and user prompts), and recursively re-run the fetch logic for redirects or credential updates.

Key Claims
- Browser caches do not widely support partial content caching per HTTP Caching standards.
- Request bodies must be cloned before potential redirects or authentication attempts to avoid failure due to consumed streams.
- If the sum of a request's body length and in-flight keep-alive bytes exceeds 64 KiB, the fetch returns a network error to prevent indefinite resource usage.
- Cache modes (`default`, `no-cache`, `no-store`, `reload`, `only-if-cached`, `force-cache`) dictate whether stored responses are used, validated, or bypassed.
- A status code of 304 (Not Modified) with a revalidating flag set updates the stored response in cache to a "validated" state.
- Network errors returned by servers can be cached ("negative caching") alongside successful responses.
- 401 and 407 responses trigger user prompts for credentials or proxy authentication, potentially leading to a recursive fetch call with updated credentials.

Entities And Concepts
- `HTTP-network-or-cache fetch`: The core algorithm combining network fetching and cache logic.
- `httpFetchParams` / `httpRequest`: Internal structures holding the fetch parameters and cloned request object.
- `includeCredentials`: Boolean flag determining if `Cookie` and `Authorization` headers are added.
- `Content-Length` header: Manually appended for POST/PUT requests with null bodies (value 0) or actual body lengths.
- `Keep-alive` management: Tracks in-flight bytes to enforce a 64 KiB limit on long-lived connections.
- HTTP Cache Partitioning: Determines the cache partition based on the request URL and credentials.
- Stale-while-revalidate: A caching strategy where a stale response is served immediately while a background fetch validates it.
- Negative Caching: Storing network errors in the cache to avoid repeated failed requests.
- Authentication Entry: Stores username/password or token for realms, triggered by 401/407 responses.
- WebDriver BiDi: Protocol steps referenced for sending and receiving data (likely internal implementation detail).

Procedures And API Details
- **Request Preparation**:
  - Clone `request` if user prompts or redirects are possible to spare a copy of the body stream.
  - Append headers: `Content-Length` (if applicable), `Referer`, `Origin`, Fetch metadata, `Sec-Purpose` (for prefetch), `User-Agent` (if missing).
- **Header Modification**:
  - Adjust cache-related headers based on `cache mode` (e.g., add `Pragma: no-cache` or `Cache-Control: max-age=0`).
  - Append `Accept-Encoding: identity` if a `Range` header is present to handle partial content correctly.
  - Normalize headers per HTTP spec, avoiding duplicates.
- **Credential Handling**:
  - Add `Cookie` header if credentials mode is "include" or "same-origin" (with basic tainting).
  - Add `Authorization` header if an auth entry exists or URL contains credentials and `isAuthenticationFetch` is true.
  - Handle proxy authentication separately, independent of request credentials mode.
- **Cache Logic**:
  - Determine cache partition; set mode to "no-store" if no cache exists.
  - Select stored response if applicable (stale-while-revalidate).
  - If stale and revalidating: send validation request with `If-None-Match` or `If-Modified-Since`.
  - On 304 response, update stored response headers and mark as "validated".
  - On new response (2xx/other), store in cache (including network errors).
- **Error Handling**:
  - Return network error if fetch is canceled.
  - Return network error if `only-if-cached` mode finds no cached response.
  - Invalidate stored responses for unsafe methods with 2xx status codes.

Nuance Or Contradictions
- The spec notes that while HTTP caching allows partial content, browser support is limited.
- There is a tension between cloning request bodies (to support retries/redirects) and efficiency; the spec encourages avoiding teeing streams when not needed.
- The 64 KiB limit on keep-alive bytes is an implementation constraint to prevent indefinite resource holding, distinct from standard HTTP limits.
- Some header manipulations (like appending `User-Agent`) are phrased as "user agents should" rather than strict normative requirements in some contexts.
- The handling of multiple `WWW-Authenticate` or `Proxy-Authenticate` headers is marked as needing testing.

Candidate Wiki Hints
- **HTTP Fetch Algorithm**: A page explaining the high-level flow of `fetch()` including network vs cache interactions.
- **Request Cloning and Streams**: Notes on when to clone request bodies (redirects, auth) and implications for stream consumption.
- **Cache Modes Explained**: A guide covering `default`, `no-store`, `reload`, `only-if-cached`, etc., with examples of behavior changes.
- **Negative Caching**: Concept of storing network errors in the cache to improve performance on repeated failures.
- **Stale-While-Revalidate**: Pattern for serving cached data while updating it asynchronously.
- **Authentication Flows**: How 401/407 responses trigger credential prompts and recursive fetch calls.
