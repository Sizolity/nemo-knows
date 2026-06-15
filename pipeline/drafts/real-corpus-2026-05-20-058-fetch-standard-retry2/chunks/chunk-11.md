---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- Source location: `raw/web/corpus-2026-05-18/058-fetch-standard.md`, lines 3479–3797.
- Heading path: `5. Return response. > 4.3. Scheme fetch`.
- Coverage includes subsections: `4.3. Scheme fetch`, `4. Return a network error.`, `4.4. HTTP fetch`, and `4.5. HTTP-redirect fetch`.

Local Summary
This chunk details the Fetch specification's algorithmic steps for handling different URL schemes (`about`, `blob`, `data`, `file`), defining how to determine the request environment, and outlining the core logic for HTTP fetching including service-worker interactions, CORS preflight checks, cache usage, cross-origin resource policy (CORP) checks, timing updates, and redirect handling. It covers conditions under which a network error is returned versus a valid response being constructed or forwarded.

Key Claims
- **Scheme Handling:** The `about:blank` scheme returns an empty body with status `OK` and content type `text/html;charset=utf-8`. Other `about:` URLs (e.g., `about:config`) result in network errors during fetching.
- **Blob Fetching:** Only the `GET` method is allowed for blob URLs; other methods return a network error. The algorithm handles range requests by calculating byte ranges, slicing the blob, and returning a `206 Partial Content` response with appropriate headers (`Content-Range`, `Content-Length`).
- **Data URL Fetching:** Data URLs are processed via a specific processor, returning the body directly with an `OK` status and the serialized MIME type.
- **Environment Determination:** The environment for a request is derived from its reserved client, fallback to its standard client, or null if neither exists.
- **HTTP Fetch Logic:** If service-workers mode is "all", the request body is piped through a TransformStream before being handled. Timing info is updated upon response arrival.
- **Error Conditions:** A network error is returned if the response type is "error", specific CORS/redirect mismatches occur, or cross-origin resource policy blocks the access.
- **CORS Preflight:** If `makeCORSPreflight` is true and caching conditions aren't met for a safe method, a preflight fetch is initiated to populate the cache.
- **Redirect Handling:** Redirects update timing info, strip non-wildcard CORS headers on cross-origin jumps, normalize methods (e.g., POST becomes GET on 301/302), and enforce redirect count limits (max 20).

Entities And Concepts
- **Fetch Params:** The structured input containing the request and controller for fetch operations.
- **Blob URL Entry:** Metadata associated with a blob URL used to extract or slice blob objects.
- **Request Environment:** The context (Window, navigable, or null) determining how resources are loaded.
- **Service Worker Timing Info:** Data tracking when service workers handle parts of the fetch lifecycle.
- **CORS Unsafe Request Header Names:** Headers that require preflight checks before cross-origin access.
- **Cross-Origin Resource Policy (CORP):** A security check distinct from CORS, ensuring embedder policies are respected across origins.
- **TransformStream:** Used to stream and potentially transform request bodies in service-worker contexts.

Procedures And API Details
- **Safely Extracting Blob:** Converts a blob object into a body with associated type information.
- **Parsing Range Header:** Validates and calculates `rangeStart` and `rangeEnd`, adjusting for inclusive byte range semantics versus slice algorithm requirements.
- **Building Content Range:** Constructs the `Content-Range` header value based on start, end, and total length.
- **CORS-Preflight Fetch:** An algorithm to check cache or perform a preflight request to ensure resource familiarity with CORS protocols.
- **HTTP-Network-or-Cache Fetch:** The underlying mechanism to retrieve responses from network or cache before final validation steps.
- **Main Fetch Invocation:** Called recursively during manual redirects to ensure response tainting is correctly updated.

Nuance Or Contradictions
- **File Scheme:** The document explicitly leaves `file:` URLs as an exercise for the reader, noting it is "unfortunate."
- **Range Header Semantics:** There is a specific distinction between inclusive byte ranges in HTTP headers versus the exclusive end indices required by the slice algorithm, necessitating an increment of `rangeEnd`.
- **303 Redirects:** The specification excludes status 303 from certain RST_STREAM frame recommendations, as some communities attribute special significance to it.
- **Manual Redirects:** When redirect mode is "manual", the recursive flag is set to false, altering the flow compared to standard "follow" redirects.

Candidate Wiki Hints
- **Page: Fetch API Algorithms** – A comprehensive guide covering `scheme fetch`, `HTTP fetch`, and `redirect` logic.
- **Page: Understanding Blob Fetching** – Deep dive into blob URL handling, range requests, and slicing mechanics.
- **Page: CORS and Preflight Logic** – Explaining the interaction between `makeCORSPreflight`, caching, and service workers.
- **Page: Redirect Handling in Fetch** – Focusing on header stripping, method normalization, and timing updates during redirects.
