## chunk-01

---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

**Heading Path**: Document → Fetch Standard → Infrastructure → HTTP → Methods
**Line Range**: Lines 1–447 (covering Preface, URL, HTTP methods, and related definitions)
**Source**: `raw/web/corpus-2026-05-18/058-fetch-standard.md`

# Local Summary

This chunk introduces the Fetch Standard's goal of unifying resource fetching across web APIs. It defines core infrastructure concepts including URL schemes, HTTP method normalization, CORS-safelisted methods, and forbidden methods. It also outlines fetch parameters, controllers, timing info structures, and basic HTTP string parsing logic.

# Key Claims

- The Fetch Standard aims to unify fetching behavior across various web APIs (e.g., `<img>`, `navigator.sendBeacon()`, `fetch()`).
- Fetching involves handling URL schemes, redirects, CORS, CSP, service workers, and mixed content uniformly.
- HTTP methods are technically case-sensitive but normalized for consistency (`GET`, `POST`, etc.).
- Methods like `PATCH` are preferred over `patch` to avoid 405 errors.
- Arbitrary method names (e.g., `CHICKEN`) are allowed unless they match forbidden verbs.

# Entities And Concepts

- **Fetch Standard**: Living standard defining requests, responses, and fetching logic.
- **Fetch Params / Controller**: Internal structures managing fetch state, timing, and abort reasons.
- **URL Schemes**: Local (`about`, `blob`, `data`), HTTP(S), and others like `file`.
- **HTTP Methods**: Normalized uppercase forms; CORS-safelisted (`GET`, `HEAD`, `POST`) vs forbidden (`CONNECT`, `TRACE`, `TRACK`).
- **Fetch Timing Info**: Struct for tracking timing data across fetch stages.
- **Response Body Info**: Tracks encoded/decoded sizes and content types.

# Procedures And API Details

**Method Normalization**:
1. If method matches `DELETE`, `GET`, `HEAD`, `OPTIONS`, `POST`, or `PUT` (case-insensitive), convert to uppercase.
2. Other methods retain original casing unless explicitly forbidden.

**CORS-Safelisted Methods**:
- Must be one of: `GET`, `HEAD`, `POST`.

**Forbidden Methods**:
- Case-insensitive match for: `CONNECT`, `TRACE`, `TRACK`.

**Fetch Controller State**:
- `"ongoing"` (default), `"terminated"`, `"aborted"`.
- Aborted if controller state is `"aborted"`.
- Canceled if state is `"aborted"` or `"terminated"`.

# Nuance Or Contradictions

- HTTP methods are technically case-sensitive per protocol, but normalization is applied for API consistency.
- Using lowercase `patch` may yield a `405 Method Not Allowed`; uppercase `PATCH` is safer.
- Arbitrary method names are allowed (e.g., `Egg`, `eGg`), though uppercase casing is encouraged.

# Candidate Wiki Hints

- **Page: Fetch Standard Overview** – High-level goals and unification across APIs.
- **Page: HTTP Method Handling** – Normalization rules, CORS-safelisted vs forbidden methods.
- **Page: Fetch Controller Internals** – State machine, timing info, abort handling.
- **Page: URL Scheme Classification** – Local, HTTP(S), fetch schemes usage.

## chunk-02

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

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
This chunk details the structure of a `Body` object, defining its components (stream, source, length) and cloning logic. It further elaborates on algorithms for incrementally reading a body, including handling chunks, closing, and errors via parallel queues. Finally, it covers fully reading a body and handling content codings.

Local Summary
A Body is composed of a ReadableStream, an optional source (null, byte sequence, Blob, or FormData), and an optional length. Cloning involves teeing the stream and duplicating properties. Incremental reading utilizes specific algorithms (`processBodyChunk`, `processEndOfBody`, `processBodyError`) queued as fetch tasks to handle incoming bytes, completion, or errors. Full reading aggregates these steps to consume the entire stream. Content codings are decoded if supported; otherwise, raw bytes are returned.

Key Claims
- A body consists of a stream, source (initially null), and length (initially null).
- Cloning a body tees its stream into two outputs (`out1`, `out2`); the original uses `out1`, while the clone uses `out2`.
- Incremental reading requires algorithms for processing chunks, end-of-body, and errors.
- If a chunk is not a Uint8Array, a TypeError is queued via `processBodyError`.
- Implementations are strongly encouraged to avoid copying byte sequences where possible during incremental reads.
- Fully reading a body queues tasks for success (byte sequence) and error (exception) handling.
- Content codings are decoded if supported; failure results in returning the raw bytes.

Entities And Concepts
- `Body`: An object representing the response body.
- `ReadableStream`: The underlying stream of data within a Body.
- `source`: The initial source of the body (null, byte sequence, Blob, FormData).
- `length`: The size of the body content.
- `processBodyChunk`: Algorithm to handle incoming byte sequences.
- `processEndOfBody`: Algorithm to handle the end of the stream.
- `processBodyError`: Algorithm to handle exceptions during reading.
- `taskDestination`: A parallel queue or global object for scheduling tasks.
- `Uint8Array`: The expected type for body chunks; deviation triggers a TypeError.

Procedures And API Details
- **Cloning a Body**:
  1. Tee the body's stream into `out1` and `out2`.
  2. Set the original body's stream to `out1`.
  3. Return a new body with stream `out2` and copied members.
- **Incrementally Reading a Body**:
  - Requires algorithms: `processBodyChunk(bytes)`, `processEndOfBody()`, `processBodyError(exception)`.
  - Initializes `taskDestination` to a new parallel queue if null.
  - Gets a reader for the body's stream.
  - Loops through chunks, queuing tasks for processing, closing, or errors.
- **Fully Reading a Body**:
  - Requires algorithms: `processBody(bytes)`, `processBodyError(exception)`.
  - Initializes `taskDestination` to a new parallel queue if null.
  - Reads all bytes from the reader, invoking success or error steps accordingly.

Nuance Or Contradictions
- The spec notes that getting a reader for a body's stream will not throw an exception, implying robustness in stream access.
- While copying byte sequences is part of the standard algorithm, implementations are strongly encouraged to avoid this copy to optimize performance.
- Content coding decoding fails if an error occurs during decoding, resulting in failure; otherwise, it returns the decoded bytes. If codings are unsupported, raw bytes are returned immediately.

Candidate Wiki Hints
- [Fetch API: Body Structure](https://wiki.example.com/fetch-api/body-structure)
- [Fetch API: Incremental Reading Algorithms](https://wiki.example.com/fetch-api/incremental-reading)
- [Fetch API: Handling Content Codings](https://wiki.example.com/fetch-api/content-codings)

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
This chunk documents the detailed properties and attributes of an HTTP `Request` object within the Fetch standard. It covers initialization flags, security contexts (CORS, CSP), caching behavior, service worker interactions, and internal bookkeeping details used by the fetch algorithm. The content spans from basic request attributes like method and URL to complex scenarios involving redirects, user prompts, and WebDriver integration.

Local Summary
The section defines a `Request` as an object containing a method (defaulting to `GET`), a URL, and numerous optional flags that dictate behavior during fetching. Key properties include headers, body, client context, origin, referrer settings, and cache modes. The chunk details how the request interacts with service workers based on its destination type and initiator. It also outlines specific constraints for CORS (Cross-Origin Resource Sharing), Content Security Policy (CSP) directives mapped to initiators/destinations, and caching strategies like `no-store` or `reload`. Finally, it lists internal flags used by the browser's HTML navigation algorithm and WebDriver protocols.

Key Claims
- A request defaults to `GET` method and has a null body unless explicitly set.
- The `unsafe-request` flag is managed by APIs like `fetch()` and `XMLHttpRequest` to trigger CORS-preflight checks but does not override API restrictions on forbidden methods/headers.
- Requests have a `traversable for user prompts` attribute that determines where UI (e.g., auth dialogs) appears: `"no-traversable"`, `"client"`, or a specific navigable.
- The `destination` property maps to CSP directives (e.g., `script-src`, `img-src`) and dictates which origins can load the resource.
- Cache modes range from `default` (uses HTTP cache with conditional requests) to `only-if-cached` (returns network error if cache miss).
- CORS mode defaults to `"no-cors"`, restricting methods/headers but returning an opaque response on success; standards are discouraged from using this for new features.
- The `redirect count` and `response tainting` (`basic`, `cors`, `opaque`) serve as bookkeeping for the fetch algorithm.

Entities And Concepts
- **Request Object**: The core entity holding all attributes for an HTTP request in the Fetch API.
- **CORS (Cross-Origin Resource Sharing)**: A security mechanism involving `mode` (`same-origin`, `cors`, `no-cors`) and response tainting.
- **Content Security Policy (CSP)**: Uses `initiator` and `destination` to determine which CSP directives apply (e.g., `script-src` for scripts).
- **HTTP Cache**: Managed via `cache mode` attributes (`default`, `no-store`, `reload`, etc.).
- **Service Workers**: Controlled by `service-workers mode` and `destination` types like `"serviceworker"`.
- **User Prompts**: UI behavior is tied to the `traversable for user prompts` attribute.

Procedures And API Details
- **Request Initialization**: A request starts with a URL, optional method (default `GET`), headers list, and various flags set to defaults (`null`, `unset`, or specific strings).
- **CORS Preflight**: Triggered if the `use-CORS-preflight` flag is set, which happens when event listeners are on an `XMLHttpRequestUpload` or a `ReadableStream` is used in the request.
- **Cache Behavior**:
  - `default`: Checks cache for fresh/stale responses; makes conditional network fetch if needed.
  - `no-cache`: Makes conditional request if cache exists, otherwise normal request.
  - `only-if-cached`: Returns cached response or network error; only allowed in `same-origin` mode.
- **Destination Mapping**:
  - `"script"` maps to `script-src`.
  - `"image"` maps to `img-src`.
  - `"audio"` maps to `media-src`.
  - `"serviceworker"` skips service worker events.

Nuance Or Contradictions
- **Origin Resolution**: The request's `origin` starts as `"client"` and is resolved to a specific origin during fetching, simplifying standards that need the origin without explicit setting.
- **Credentials Handling**: When `mode` is `"navigate"`, the `credentials mode` is implicitly treated as `"include"`, overriding other values unless HTML changes occur.
- **URL Credentials**: The `use-URL-credentials` flag allows URL username/password to override authentication entries, though modern specs avoid setting this flag due to security discouragement.
- **CSP Granularity**: The request's initiator is not currently granular enough for all features; it primarily assists in defining CSP and Mixed Content rules.

Candidate Wiki Hints
- Page: **Request Attributes** (Summarizes all properties of the Request object).
- Page: **CORS Modes Explained** (Detailing `same-origin`, `cors`, `no-cors` behaviors).
- Page: **CSP Initiators and Destinations** (Mapping request types to CSP directives).
- Page: **HTTP Cache Strategies** (Explaining `cache mode` options).

## chunk-05

---
title: Chunk 05 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- **Heading**: 2. Read a chunk from reader given readRequest. > 2.2.5. Requests
- **Line Range**: 1501–1595 of the source document `raw/web/corpus-2026-05-18/058-fetch-standard.md`.

Local Summary
This chunk defines request classification (subresource, non-subresource, navigation), specifies the algorithm to compute `redirect-taint` and serialize request origins, details how to clone a request while handling bodies and WebDriver IDs, describes adding Range headers for partial responses, warns about security implications of combining multiple responses, and outlines serialization for response URLs to prevent leaking redirect targets. It also includes logic for checking Cross-Origin-Embedder-Policy credential allowance.

Key Claims
- A subresource request has destinations like "audio", "font", "image", "script", etc., while a non-subresource request includes "document", "embed", "worker", etc.
- A navigation request is restricted to "document", "embed", "frame", "iframe", or "object".
- `redirect-taint` computation iterates through a request's URL list, comparing origins of subsequent URLs against the previous one and the request's origin to determine if the taint is "same-origin", "same-site", or "cross-site".
- Serializing a request origin returns "null" if the redirect-taint is not "same-origin"; otherwise, it returns the serialized origin.
- Cloning a request creates a new copy excluding the original body and WebDriver id, then generates a random UUID for the new WebDriver id and clones the body if it exists.
- Adding a Range header constructs a value like `bytes=<first>-<last>` (or just `<first>-` if last is omitted) and appends it to the request's header list.
- Combining multiple responses into one logical resource is historically a source of security bugs and should undergo security review.
- Serializing a response URL for reporting uses the first URL in the list to avoid leaking information about redirect targets.
- Cross-Origin-Embedder-Policy allows credentials if the request mode is not "no-cors", the client is null, the embedder policy is not "credentialless", or specific origin/redirect-taint conditions are met.

Entities And Concepts
- **Request Classification**: subresource request, non-subresource request, navigation request.
- **Redirect Taint**: same-origin, same-site, cross-site.
- **Serialization**: request origin serialization, byte-serialization, response URL serialization for reporting.
- **Request Cloning**: handling body cloning and WebDriver id generation.
- **Range Headers**: inclusive byte range notation (`bytes=<first>-<last>`).
- **Security Concepts**: redirect target leakage prevention, Cross-Origin-Embedder-Policy credential allowance, partial response security bugs.

Procedures And API Details
- **Compute Redirect Taint**:
  1. Assert request's origin is not "client".
  2. Initialize `lastURL` to null and `taint` to "same-origin".
  3. Iterate through each URL in the request's URL list:
     - If `lastURL` is null, set it to the current URL and continue.
     - If current URL's origin is not same site with `lastURL`'s origin AND request's origin is not same site with `lastURL`'s origin, return "cross-site".
     - If current URL's origin is not same origin with `lastURL`'s origin AND request's origin is not same origin with `lastURL`'s origin, set `taint` to "same-site".
     - Set `lastURL` to the current URL.
  4. Return the final `taint`.

- **Serialize Request Origin**:
  1. Assert request's origin is not "client".
  2. If redirect-taint is not "same-origin", return "null".
  3. Return the serialized request's origin.

- **Clone a Request**:
  1. Create a copy of the request excluding body and WebDriver id.
  2. Generate a random UUID for the new request's WebDriver id.
  3. If the original body is non-null, clone the body for the new request.
  4. Return the new request.

- **Add Range Header**:
  1. Assert `last` is not given or `first` <= `last`.
  2. Initialize `rangeValue` with `bytes=`.
  3. Serialize and isomorphic encode `first`, append to `rangeValue`.
  4. Append `-` (0x2D) to `rangeValue`.
  5. If `last` is given, serialize and isomorphic encode it, append to `rangeValue`.
  6. Append (`Range`, `rangeValue`) to the request's header list.

- **Serialize Response URL for Reporting**:
  1. Assert response's URL list is not empty.
  2. Copy the first URL from the response's URL list (not the current response URL) to avoid leaking redirect targets.
  3. Set username and password to empty strings for that URL.
  4. Return the serialization of the URL with the fragment excluded.

- **Check Cross-Origin-Embedder-Policy Credential Allowance**:
  1. Assert request's origin is not "client".
  2. If request mode is not "no-cors", return true.
  3. If request's client is null, return true.
  4. If request's client's policy container's embedder policy value is not "credentialless", return true.
  5. If request's origin is same origin with current URL's origin AND redirect-taint is not "same-origin", return true.
  6. Return false.

Nuance Or Contradictions
- The definition of a navigation request overlaps significantly with non-subresource requests (both include "document", "embed"), suggesting navigation is a subset of non-subresource requests in terms of destination types, but distinguished by specific intent or handling not fully detailed here.
- The `redirect-taint` computation relies on comparing origins across the URL list; if any step triggers a "cross-site" condition, it immediately returns that status, prioritizing strict cross-origin detection over cumulative site-level analysis.
- Serializing a response URL specifically avoids using the current response URL to prevent leaking redirect targets, implying that the current URL might differ from the final resolved URL after redirects.

Candidate Wiki Hints
- **Request Classification**: Define and differentiate subresource, non-subresource, and navigation requests with their respective destination types.
- **Redirect Taint Algorithm**: Document the step-by-step logic for computing `redirect-taint` values ("same-origin", "same-site", "cross-site").
- **Request Origin Serialization**: Explain conditions under which a request origin serializes to "null" versus returning the actual origin, focusing on redirect-taint implications.
- **Request Cloning**: Detail the process of cloning a request, specifically handling body cloning and WebDriver id regeneration.
- **Range Headers**: Describe the format and construction of Range headers for partial content retrieval.
- **Response URL Serialization**: Explain the security rationale for using the first URL in the list rather than the current response URL when reporting.
- **Cross-Origin-Embedder-Policy**: Outline the conditions under which credentials are allowed or disallowed based on request mode, client state, and policy settings.

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
This chunk covers sections 2.2.6 (Responses), 2.2.7 (Miscellaneous), 2.3 (Authentication entries), 2.4 (Fetch groups), and the beginning of 2.5 (Resolving domains) within the Fetch Standard. It details the structure, lifecycle flags, and filtering logic for `Response` objects, network errors, and fetch group state management.

Local Summary
A `fetch` call returns a `response` that evolves over time. Responses have specific types ("basic", "cors", "default", "error", "opaque", "opaqueredirect") and associated attributes like status, headers, body, and CORS-exposed header lists. The standard distinguishes between filtered responses (limited views for security) and internal responses. It also defines fetch groups to manage active requests and deferred invocations, and outlines the domain resolution process involving IP address sets.

Key Claims
- A response's type defaults to "default" unless specified otherwise.
- Responses over HTTP/2 connections always have an empty status message.
- `response.ok` may return useless results for filtered responses like opaque types.
- Cloning a filtered response involves cloning its internal response if the body is non-null.
- Fetch groups manage active `fetch records` and `deferred fetch records` until termination.
- Resolving an origin typically involves DNS queries and returns a set of IP addresses.

Entities And Concepts
- **Response**: The result of a fetch, evolving over time with fields like status, headers, body, and type.
- **Filtered Response**: A response offering a limited view (e.g., excluding headers) to prevent information leakage; includes "basic", "cors", "opaque", and "opaqueredirect" types.
- **Network Error**: A response with type "error", status 0, and null body.
- **Fetch Group**: Holds fetch records and deferred fetch records within an environment settings object.
- **Fetch Record**: Contains a `request` and a `controller`.
- **Deferred Fetch Record**: Maintains state to invoke a fetch later (e.g., on document unload).

Procedures And API Details
- **Cloning a Response**:
  1. If the response is filtered, return a new identical filtered response cloning the internal response.
  2. Copy the response except for its body.
  3. Clone the body if it is non-null.
  4. Return the new response.
- **Location URL Algorithm**: Used for redirect handling; returns null if status is not a redirect, extracts the `Location` header, parses the URL (handling fragments), and returns the location.
- **Resolving an Origin**:
  1. Returns host IPs if the host is an IP address.
  2. Returns `::1, 127.0.0.1` for "localhost".
  3. Performs implementation-defined operations (e.g., DNS queries) to return a set of IP addresses.
  4. Returns failure otherwise.

Nuance Or Contradictions
- **Filtered vs. Internal**: While filtered responses expose limited data, specification algorithms may access the internal response for legacy reasons (e.g., feeding image decoders), but new APIs should avoid this to prevent leaks.
- **Opaque Responses**: `opaque` and `opaqueredirect` filtered responses are nearly indistinguishable from network errors because they lack headers, status, and body info accessible to scripts.
- **Caching DNS Results**: The results of resolving an origin may be cached, but implementations vary on whether they can account for the partition key in local caching or rely solely on DNS server caching.

Candidate Wiki Hints
- [Fetch Standard / Response Types](https://wiki.local/fetch/response-types)
- [Filtered Responses Security Model](https://wiki.local/fetch/filtered-responses)
- [Fetch Group Lifecycle](https://wiki.local/fetch/groups)

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
This chunk details the HTTP Fetch Standard's mechanisms for connection management, specifically focusing on how user agents obtain connections from a pool, handle network partition keys, manage HTTP cache partitions, block bad ports, and process cookies. It also covers timing information for connections to prevent leaking reuse details and defines infrastructure for SameSite cookie modes.

Local Summary
A user agent maintains an ordered set of connections identified by a network partition key, origin, and credentials. Obtaining a connection involves checking the pool, finding proxies (including non-standard ones like WPAD/PAC), and creating new connections if necessary. Connection timing is recorded and coarsened to prevent exposing reuse details. Network partition keys are derived from the top-level site. Bad ports (like echo or daytime) trigger blocking. Cookies are managed via specific headers with infrastructure for SameSite modes, HttpOnly flags, and path serialization.

Key Claims
- Connections are keyed by a network partition key, origin, and credentials; TLS session identifiers are not reused across connections with different credentials.
- Obtaining a connection involves a sequence: check pool, find proxies, resolve hosts (racing IPv4/IPv6), create connection, and potentially add to the pool.
- Connection timing info includes domain lookup, connection start/end, secure connection start, and ALPN protocol, which are clamped and coarsened to hide reuse details.
- A port is considered "bad" if listed in a specific table (e.g., 0, 1, 7, 9), leading to request blocking for HTTP(S) requests targeting those ports.
- MIME types like `audio/*`, `image/*`, `video/*`, and `text/csv` are blocked when the destination is script-like.
- The `Cookie` header is appended based on retrieved cookies filtered by security context, host, path, and SameSite mode.
- The `Set-Cookie` header is parsed and stored independently for each occurrence, with garbage collection of cookies.

Entities And Concepts
- **Connection Pool**: An ordered set of connections associated with a user agent.
- **Network Partition Key**: A tuple consisting of a site and an optional implementation-defined value (second key).
- **Connection Timing Info**: A struct tracking domain lookup times, connection start/end times, secure connection start time, and ALPN negotiated protocol.
- **Bad Port**: Ports listed in a standard table (e.g., 0, 1, 7) that should block fetching requests.
- **HTTP Cache Partition**: The unique HTTP cache associated with a network partition key.
- **SameSite Mode**: Determined for requests to handle cookie security contexts (strict-or-less, lax-or-less, unset-or-less).

Procedures And API Details
- **Obtain Connection**:
  1. Check pool for existing connection matching key, origin, and credentials.
  2. If none or `requireUnreliable` is true, find proxies (defaulting to "DIRECT").
  3. Resolve hosts (handling failures).
  4. Create connection in parallel if multiple proxies exist; select one return value.
  5. Add to pool unless `new` setting is "yes-and-dedicated".
- **Create Connection**:
  - Set connection start time.
  - Establish transport (TCP/UDP/TLS), handling proxy tunnels and ALPN.
  - Handle WebTransport requirements (SETTINGS_ENABLE_WEBTRANSPORT, H3_DATAGRAM).
  - Manage TLS certificates (client cert if credentials=true, custom verification via webTransportHashes).
- **Record Timing**:
  - Set connection end time immediately after establishing transport/TLS handshake sufficient to request resource.
  - Secure connection start time set before handshake.
  - For HTTP/3, connection start and secure connection start times must be equal.
- **Determine Network Partition Key**:
  - Get top-level origin from environment or creation URL.
  - Obtain site for the origin.
  - Return tuple (topLevelSite, secondKey).
- **Append Cookie Header**:
  - Check cookie disable configuration.
  - Determine SameSite mode and isSecure.
  - Retrieve cookies based on host, path, httpOnlyAllowed, and SameSite.
  - Serialize and append to header list.

Nuance Or Contradictions
- Connection management details are intentionally vague regarding nuances like IP address selection (favoring IPv6) or retry logic, leaving discretion to implementers.
- The "second key" in network partition keys is noted as evolving (referencing issue #1035).
- Timing info clamping ensures reused connections do not expose timing details; specific rules apply for TLS False Start and early data scenarios.
- Proxy handling allows non-standard technologies like WPAD/PAC to influence the "DIRECT" or proxy value.

Candidate Wiki Hints
- **Connection Pooling Strategy**: How browsers manage connection reuse, pooling, and lifecycle based on credentials and origins.
- **Network Partition Key Definition**: The tuple structure (site, secondKey) used to isolate network resources for privacy and caching.
- **Connection Timing Coarsening**: Techniques to hide connection reuse patterns by clamping timestamps.
- **Bad Port Blocking List**: A curated list of historically dangerous or obsolete ports (e.g., echo, daytime) that fetch should block.
- **Cookie Infrastructure in Fetch**: How the Fetch API handles `Cookie` and `Set-Cookie` headers, including SameSite logic and path serialization.

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
- Location: Section 3. HTTP extensions, subsections 3.2 (`Origin` header) and 3.3 (CORS protocol), followed by a brief note on `Content-Length`.
- Scope: Defines the `Origin` header serialization rules, conditions for including it in requests, CORS preflight mechanics, response headers for cross-origin sharing, credential handling, and specific exceptions for certain content types.

Local Summary
This chunk details how the `Origin` header is serialized and when it must be included in fetch requests based on response tainting, request method, and referrer policy. It then introduces the CORS protocol as an opt-in mechanism to allow cross-origin resource sharing while preventing data leakage. The text covers preflight requests (`OPTIONS`), relevant headers like `Access-Control-Allow-Origin`, handling of credentials, and specific rules for header exposure. Finally, it notes limited exceptions for certain non-safelisted content types in the CORS context.

Key Claims
- The `Origin` header indicates where a fetch originates and is used for all HTTP fetches with response tainting "cors" or methods other than `GET`/`HEAD`.
- The `Origin` header serialization is more constrained than RFC 3986: it uses lower-case ASCII without percent encoding, forbids leading zeros in IPv4, requires lowercase hex for IPv6, and limits the use of `::` to at most 6 blocks.
- CORS is an opt-in protocol necessary to prevent leaking data from responses behind firewalls and sensitive data with credentials.
- A successful CORS response can have any status code if it includes the appropriate headers; a preflight response must be "ok" (e.g., 200 or 204).
- When `Access-Control-Allow-Credentials` is present, `Access-Control-Allow-Origin` cannot be `*` and must match the request origin exactly.
- The `Content-Length` header extraction logic validates that values are consistent ASCII digits; inconsistent values result in failure.

Entities And Concepts
- **Origin Header**: Indicates fetch source; serialized as scheme + host + optional port.
- **Serialized Origin**: Strict ABNF grammar for IPv4, IPv6, and domain names (lowercase only).
- **Response Tainting**: Determines if CORS headers are required ("cors").
- **CORS Protocol**: Mechanism allowing cross-origin resource sharing via specific headers.
- **Preflight Request**: An `OPTIONS` request checking server CORS support.
- **Access-Control-Allow-Origin**: Response header specifying which origins can access the resource.
- **Credentials Mode**: "include" or "same-origin"; impacts whether cookies/authorization headers are sent and how CORS rules apply.
- **Referrer Policy**: Affects whether an `Origin` header is included (e.g., "no-referrer").

Procedures And API Details
**Appending `Origin` Header Steps:**
1. Assert request origin is not "client".
2. Byte-serialize the request origin.
3. If response tainting is "cors" or mode is "websocket"/"webtransport", append (`Origin`, serializedOrigin).
4. Otherwise, if method is not `GET`/`HEAD`:
   - If mode is not "cors", apply referrer policy rules:
     - "no-referrer": Set origin to `null`.
     - "no-referrer-when-downgrade", "strict-origin", "strict-origin-when-cross-origin": Set origin to `null` if scheme downgrades from HTTPS to non-HTTPS.
     - "same-origin": Set origin to `null` if not same origin as current URL.
   - Append (`Origin`, serializedOrigin).

**CORS Response Headers:**
- **Preflight**: Must return 2xx/3xx status with `Access-Control-Allow-Origin`, `Allow-Methods`, `Allow-Headers`, and optionally `Max-Age`.
- **Actual Request**: Can return any status with `Access-Control-Allow-Origin` (matching origin or `*`), `Allow-Credentials` (if needed), and `Expose-Headers` (list of names or `*` if no credentials).

**Credential Logic:**
- If `credentials` mode is "include", `Access-Control-Allow-Origin` must not be `*`.
- `Access-Control-Allow-Credentials: true` is required when sharing with credentials.
- `Set-Cookie` headers are functional only if the request includes credentials and CORS allows it.

Nuance Or Contradictions
- **Origin Header Ambiguity**: The `Origin` header is present in all non-simple requests (non-GET/HEAD) regardless of whether they participate in the CORS protocol, making it unreliable as a sole indicator of CORS participation without checking other headers or context.
- **Status Code Flexibility**: A "successful" response to a non-preflight CORS request can technically be 403 if the server intends to share it (by including headers), which contradicts typical HTTP semantics where 4xx implies failure, though the spec clarifies that side channels might still leak data.
- **Case Sensitivity**: `Access-Control-Allow-Credentials: true` is byte-case-sensitive; "True" or "TRUE" are invalid.

Candidate Wiki Hints
- **CORS Protocol Overview**: Explain the opt-in nature, preflight requests, and key headers (`Allow-Origin`, `Allow-Credentials`, `Expose-Headers`).
- **Origin Header Serialization**: Detail the strict ABNF constraints compared to RFC 3986 (lowercase, IPv6 limits).
- **Referrer Policy and Origin**: How policies like "no-referrer" affect the `Origin` header inclusion.
- **Credentials in CORS**: Rules governing `Access-Control-Allow-Credentials` and why `*` is forbidden with credentials.
- **CORS Exceptions**: List of allowed non-safelisted content types (`application/csp-report`, etc.).

## chunk-09

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

## chunk-10

---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
This chunk details the `main fetch` algorithm (section 4.1) and the `override fetch` algorithm (section 4.2) within the Fetch Standard. It covers security checks, URL upgrades, response tainting logic, filtering, timing information handling, and specific behaviors for preloaded responses or data URLs.

Local Summary
The `main fetch` algorithm initializes a request, performs various security and policy checks (local-only flags, CSP, mixed content, referrer policy, HSTS), and determines the appropriate response type based on the URL scheme and CORS mode. It handles parallel execution for non-recursive requests, applies filtering to expose headers appropriately, manages timing information, validates integrity metadata, and processes the response body via a TransformStream. The `override fetch` algorithm allows user agents to intervene before standard fetching occurs, enabling features like content shimming or blocking unsafe domains.

Key Claims
- **Security Checks**: Before fetching, the algorithm checks for local-only flags, Content Security Policy violations, mixed content risks, bad ports, and Integrity Policy blocks.
- **URL Upgrades**: Requests are upgraded to HTTPS if the host matches an HSTS domain or has a matching SVCB HTTPS RR.
- **Response Tainting**: The response is tainted as "basic", "cors", or "opaque" depending on the request's mode, URL scheme, and redirect settings.
- **Header Exposure**: For CORS requests, `Access-Control-Expose-Headers` determines which headers are exposed; if credentials are not included but headers contain `*`, all headers are exposed.
- **Integrity Validation**: If integrity metadata is present, the response body must match it; otherwise, a network error is returned.
- **Override Fetch**: Allows user agents to intercept requests and return custom responses (e.g., shimming resources) or errors before standard fetching proceeds.

Entities And Concepts
- **Fetch Params**: The structured data containing the request, client info, and callbacks for handling response end-of-body.
- **Response Tainting**: A state (`basic`, `cors`, `opaque`) indicating how a response is exposed to JavaScript contexts.
- **HSTS (HTTP Strict Transport Security)**: Mechanism to enforce HTTPS connections via known host domain matching.
- **SVCB (Service Binding)**: Protocol for specifying alternative server bindings, including HTTPS requirements.
- **CORS (Cross-Origin Resource Sharing)**: Security mechanism controlling cross-origin resource access via headers like `Access-Control-Allow-Origin`.
- **TransformStream**: Used to wrap the response body stream to notify when the end is reached (`processResponseEndOfBody`).

Procedures And API Details
- **Main Fetch Steps**:
  1. Initialize `response` as null.
  2. Check local-only flag; if set and URL is not local, return network error.
  3. Run CSP violation reports.
  4. Upgrade request to trustworthy URL if applicable (HSTS/SVCB).
  5. Handle referrer policy and determination.
  6. Set scheme to "https" for HSTS/SVCB matches.
  7. If recursive is false, run steps in parallel; otherwise, return current response.
  8. Determine response source (preloaded, data URL, same-origin, no-cors, HTTP(S), etc.).
  9. Apply CORS filtering based on `Access-Control-Expose-Headers`.
  10. Set internal response redirect taint and timing flags.
  11. Block responses if blocked by mixed content, CSP, MIME type, or nosniff.
  12. Handle ranged responses (return network error if range requested but header missing).
  13. Nullify body for HEAD/CONNECT methods or null body status.
  14. Validate integrity metadata; abort on mismatch.
- **Fetch Response Handover**:
  - Extract `Server-Timing` headers if in a secure context.
  - Update timing info and mark resource timing.
  - Queue tasks for reporting timing and processing response end-of-body.
  - Pipe response body through a TransformStream to handle completion notifications.
- **Override Fetch**:
  - Executes `potentially override response` which returns null by default or a custom response/error.
  - Supports "scheme-fetch" (runs `scheme fetch`) and "http-fetch" (runs `HTTP fetch`).

Nuance Or Contradictions
- **DNS Operations Timing**: DNS resolution for HTTPS RRs may happen before connection attempts, requiring implementation-defined logic to unwind earlier steps if the scheme needs upgrading.
- **Ranged Responses**: Traditionally, APIs accept ranged responses even without a range request header, but the standard prevents partial responses from being provided to APIs that didn't explicitly request a range.
- **User Agent Intervention**: The default behavior of `potentially override response` is trivial (return null), but user agents often implement complex logic like blocking unsafe domains while shimming specific resources.

Candidate Wiki Hints
- **Fetch Algorithm Overview**: Summarize the high-level flow of fetching resources, including security checks and response handling.
- **Response Tainting Explained**: Detail the differences between "basic", "cors", and "opaque" responses and their impact on JavaScript access.
- **HSTS and URL Upgrades**: Explain how HSTS and SVCB protocols force HTTPS upgrades based on host matching.
- **CORS Header Exposure**: Describe how `Access-Control-Expose-Headers` interacts with credentials mode to determine visible headers.
- **Fetch Response Body Handling**: Outline the use of TransformStreams for monitoring response body completion and integrity validation.

## chunk-11

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

## chunk-12

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

## chunk-13

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

## chunk-14

---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

**Heading Path:** `2. Append entry to the user agent’s CORS-preflight cache. > 4.12. Deferred fetching`
**Line Range:** 4503–4863
**Coverage:** Section 4.12 (Deferred fetching algorithm and task sources), Section 4.12.1 (Deferred fetching quota mechanics, policies, and examples), and the beginning of Section 5 (Fetch API usage patterns).

# Local Summary

This chunk details the "deferred fetch" feature in the Fetch Standard. It defines how to queue a fetch operation to execute later—specifically when a fetch group terminates or after a timeout. The text specifies the task source priority for deferred fetches, algorithms for queuing and processing these requests, and the calculation of total request length (including URL, referrer, headers, and body).

Section 4.12.1 focuses on quota management. It establishes a default top-level quota of 640 kibibytes, with sub-allocations for cross-origin nested documents (8 kibibytes each) via the `deferred-fetch-minimal` policy. The chunk explains how to calculate available quota per reporting origin (capped at 64 kibibytes), how permissions policies (`Permissions-Policy: deferred-fetch`) allow delegation of this quota, and provides examples of navigation scenarios affecting quota reservation and release. Finally, it introduces the high-level `fetch()` API, contrasting it with XMLHttpRequest and providing code snippets for handling blobs, headers, JSON parsing, URL parameters, and progressive streaming.

# Key Claims

- Deferred fetching allows requests to be invoked at the latest possible moment (e.g., when a fetch group terminates).
- The deferred fetch task source prioritizes tasks before other sources (like DOM manipulation) to reflect the most recent state of `fetchLater()` calls before running dependent scripts.
- To compute total request length, the algorithm sums the serialized URL (excluding fragment), referrer, header list lengths, and body length.
- The top-level traversable is allocated a default deferred-fetch quota of 640 kibibytes.
- By default, 128 kibibytes are reserved for delegating to cross-origin nested documents.
- Only 64 kibibytes of the allocated quota can be used concurrently for the same reporting origin to prevent opportunistic reservation by third-party libraries.
- The `Permissions-Policy` header controls delegation; `deferred-fetch-minimal` is enabled by default for all origins, while `deferred-fetch` is enabled only for the top-level document's origin by default.
- Cross-origin or cross-agent iframes receive a default of 8 kibibytes of quota.
- Same-origin nested documents share their parent's quota and do not reserve additional space unless navigation occurs under specific conditions.

# Entities And Concepts

- **Deferred Fetch**: A mechanism to queue fetch requests for later execution.
- **`fetchLater()`**: The conceptual API (represented in algorithms) used to queue a deferred fetch.
- **Deferred Fetch Record**: An object containing the request, notify function, and invoke state.
- **Deferred Fetch Task Source**: A specific task source prioritized before script-dependent sources.
- **Total Request Length**: A metric including URL, referrer, headers, and body used to check quota limits.
- **Quota**: The memory limit (default 640 kibibytes) for deferred fetches per top-level traversable.
- **Reporting Origin Quota**: The sub-limit of 64 kibibytes per origin.
- **`Permissions-Policy`**: Header used to control quota delegation (`deferred-fetch`, `deferred-fetch-minimal`).
- **Top-Level Traversable**: A "tab" or top-level document context holding the main quota pool.
- **Navigable Container**: An element (like an iframe) capable of having its own reserved quota.
- **`fetch()`**: The standard Fetch API method for fetching resources, supporting blobs, headers, JSON, and streaming.

# Procedures And API Details

**Algorithm: Queue a Deferred Fetch**
1. Populate request from client given request.
2. Set request’s service-workers mode to "none".
3. Set request’s keepalive to true.
4. Create a new deferred fetch record (containing request and `notify invoked` function).
5. Append record to request’s client’s fetch group’s deferred fetch records.
6. If `activateAfter` is non-null: Wait until `activateAfter` ms passes OR user agent believes scripts will be lost (background, hidden state), then process deferredRecord.
7. Return deferredRecord.

**Algorithm: Compute Total Request Length**
1. Get length of request’s URL (serialized with exclude fragment true).
2. Increment by length of request’s referrer (serialized).
3. For each header (name, value), increment by `name.length + value.length`.
4. Increment by request’s body’s length.
5. Return total.

**Algorithm: Process Deferred Fetches**
1. Iterate over fetch group’s deferred fetch records.
2. Call process a deferred fetch for each record.

**Algorithm: Process a Deferred Fetch**
1. If invoke state is not "pending", return.
2. Set invoke state to "sent".
3. Fetch the request.
4. Queue a global task on the deferred fetch task source to run `notify invoked`.

**Algorithm: Get Available Quota (Simplified Logic)**
- Inputs: Document, Origin.
- Logic checks if top-level, policy allowed, and current reserved quota.
- Returns available quota (e.g., 640, 512, or 0 depending on state).

**Code Example: Fetch Blob**
```javascript
fetch("/music/pk/altes-kamuffel.flac")
  .then(res => res.blob())
  .then(playBlob)
```

**Code Example: Check Header and Parse JSON**
```javascript
fetch("https://pk.example/berlin-calling.json", {mode:"cors"})
.then(res => {
  if (res.headers.get("content-type") &&
      res.headers.get("content-type").toLowerCase().indexOf("application/json") >= 0) {
    return res.json()
  } else {
    throw new TypeError()
  }
})
.then(processJSON)
```

**Code Example: Progressive Streaming**
```javascript
function consume(reader) {
  var total = 0
  return pump()
  function pump() {
    return reader.read().then(({done, value}) => {
      if (done) return
      total += value.byteLength
      log(`received ${value.byteLength} bytes (${total} bytes in total)`)
      return pump()
    })
  }
}

fetch("/music/pk/altes-kamuffel.flac")
.then(res => consume(res.body.getReader()))
.then(() => log("consumed the entire body without keeping the whole thing in memory!"))
.catch(e => log("something went wrong: " + e))
```

# Nuance Or Contradictions

- **Quota Calculation Complexity**: The algorithm for getting available quota involves checking multiple nested conditions regarding top-level status, policy permissions (`deferred-fetch` vs `deferred-fetch-minimal`), and current reserved quotas. The text notes that 640kb should be enough but clarifies the calculation results in specific values (512, etc.) based on these flags.
- **Policy Defaults**: There is a distinction between the default allowlist for `deferred-fetch` ("self") and `deferred-fetch-minimal` ("*"). The text explicitly states that disabling `deferred-fetch-minimal` for the top-level document collects the 128 kibibytes back into the main pool.
- **Navigation Effects**: When a navigable container navigates to a cross-origin URL, it may reserve quota (64kb or 8kb). If it later navigates away or changes state, that reserved quota is freed or lost, which affects subsequent calculations.
- **Request Length Definition**: The definition of "total request length" explicitly includes the referrer and header lengths, which might be counter-intuitive if one assumes only payload size matters for quota.

# Candidate Wiki Hints

1. **Page: Deferred Fetching**
   - *Content*: Explain the concept of deferred fetching, its use case (lazy loading heavy resources), and the `fetchLater()` mechanism.
   - *Sections*: Task Sources, Quota Management, Permissions Policy (`Permissions-Policy: deferred-fetch`).

2. **Page: Fetch API Overview**
   - *Content*: High-level introduction to `fetch()`, comparison with XMLHttpRequest, and common patterns (blobs, JSON, streaming).
   - *Sections*: Basic Usage, Response Handling, Streaming Bodies.

3. **Page: Request Length Calculation**
   - *Content*: Technical detail on how request size is computed for quota purposes (URL + referrer + headers + body).
   - *Note*: This is a niche but important implementation detail for developers managing large payloads.

## chunk-15

---
title: Chunk 15 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk details the specification for the **Headers** and **Body** classes within the Fetch API. It covers the `Headers` interface (construction, mutation methods like `append`, `delete`, `get`, `set`, and guard states), the `BodyInit` union types for request bodies, and the `Body` mixin methods (`arrayBuffer`, `blob`, `json`, `text`, etc.) along with MIME type extraction logic.

# Local Summary

The document defines the `Headers` interface as a container for header name-value pairs with associated guards ("immutable", "request", "response", "request-no-cors", "none") to control mutation and CORS safety. It specifies validation steps, normalization, and the removal of privileged headers in "request-no-cors" contexts. The chunk also defines `BodyInit` unions (Blob, BufferSource, FormData, URLSearchParams, USVString, ReadableStream) and the `Body` mixin, which provides methods to read response/request bodies as streams or parsed types (JSON, text, etc.).

# Key Claims

- **Headers Object**: Has an associated header list and a guard state.
- **Guard States**: "immutable", "request", "response", "request-no-cors", "none".
- **Mutation Logic**: `set` replaces the first header with the given name; `append` adds to the list.
- **CORS Safety**: In "request-no-cors" guards, privileged no-CORS request headers are removed upon modification.
- **BodyInit Unions**: Supports Blob, BufferSource, FormData, URLSearchParams, USVString, and ReadableStream.
- **MIME Types**: `FormData` maps to `multipart/form-data` or `application/x-www-form-urlencoded`.
- **Consumption**: The `consume body` algorithm handles reading streams into promises (ArrayBuffer, Blob, etc.).

# Entities And Concepts

- **Headers**: Interface for managing HTTP headers.
- **HeadersInit**: Union of record or sequence for initializing headers.
- **Guard**: A state controlling header mutation and CORS behavior.
- **BodyInit**: Union type for request body sources (Blob, FormData, string, etc.).
- **ReadableStream**: Used to stream body data.
- **FormData**: Represents form data with multipart or URL-encoded boundaries.
- **BodyMixin**: Provides access methods (`arrayBuffer`, `blob`, `json`, `text`) for response/request bodies.
- **MIME Type Extraction**: Logic to derive content types from header lists.

# Procedures And API Details

### Headers Initialization
```javascript
const meta = { "Content-Type": "text/xml", "Breaking-Bad": "<3" };
new Headers(meta);
// Equivalent to:
const meta2 = [ ["Content-Type", "text/xml"], ["Breaking-Bad", "<3"] ];
new Headers(meta2);
```

### Mutation Methods
- `headers.append(name, value)`: Appends a header.
- `headers.delete(name)`: Removes a header (validates name first).
- `headers.set(name, value)`: Replaces existing headers with the same name.
- `headers.get(name)`: Returns comma-space separated values.
- `headers.getSetCookie()`: Returns list of Set-Cookie values.
- `headers.has(name)`: Checks existence.

### Validation Steps for Headers
1. Validate name/value types (throw TypeError if invalid).
2. Check guard state (throw TypeError if "immutable").
3. Check forbidden headers based on guard ("request" or "response").
4. Return true/false based on safety.

### BodyInit Processing
- **Blob**: Sets source to object, length to size, type to attribute value.
- **FormData**: Encodes as multipart/form-data with boundary.
- **URLSearchParams**: Serializes as `application/x-www-form-urlencoded`.
- **USVString**: Encodes as UTF-8 text/plain.

### Body Mixin Methods
- `arrayBuffer()`: Returns Promise<ArrayBuffer>.
- `blob()`: Returns Promise<Blob> with MIME type.
- `bytes()`: Returns Promise<Uint8Array>.
- `formData()`: Parses multipart or urlencoded data.
- `json()`: Parses JSON (rejects SyntaxError).
- `text()`: UTF-8 decodes body to string.

# Nuance Or Contradictions

- **Forbidden Headers**: The spec distinguishes between "forbidden request-header" and "forbidden response-header" based on the guard state, returning false rather than throwing an error for forbidden headers in mutable contexts.
- **CORS Privileged Headers**: In "request-no-cors" mode, privileged no-CORS request headers are automatically removed from the header list when modified by unprivileged code.
- **FormData Parsing**: The spec notes that a detailed parsing specification for `multipart/form-data` is to be written later, suggesting the current implementation is a rough approximation.
- **Stream Safety**: Reading a ReadableStream body throws a TypeError if the stream is disturbed, locked, or if `keepalive` is true.

# Candidate Wiki Hints

1. **Fetch Headers Guard States** (Concept)
2. **Headers Mutation Rules** (Concept)
3. **BodyInit Types and MIME Mapping** (Concept)
4. **Consuming Fetch Body Streams** (Procedure)
5. **FormData Parsing Logic** (Concept)

## chunk-16

---
title: Chunk 16 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
- **Source**: raw/web/corpus-2026-05-18/058-fetch-standard.md
- **Heading**: 3. Throw a TypeError. > 5.4. Request class
- **Lines**: 5268–5768
- **Coverage**: Defines the `Request` interface, its attributes, associated slots, constructor arguments (`init`), enumeration values for modes/credentials/cache, and detailed algorithmic steps for creating a `Request` object (including referrer parsing logic and body handling).

Local Summary
This section details the Web IDL definition of the `Request` class. It specifies that a `Request` has an associated request, headers, signal, and body. The text outlines the `new Request(input, init)` constructor, defining how to handle both string URLs and existing `Request` objects. It lists valid values for enumerations like `RequestMode` ("navigate", "same-origin", "no-cors", "cors") and explains validation rules, such as throwing a `TypeError` if the referrer scheme is "about" with path "client" while not being same-origin, or if the mode is "navigate". The section concludes with getter implementations for properties like `.url`, `.headers`, and `.referrer`.

Key Claims
- A `Request` object includes a `Body` interface.
- The `RequestInit` dictionary allows setting method, headers, body, referrer, mode, credentials, cache, redirect, integrity, keepalive, signal, window (null only), duplex, and priority.
- If the input to `new Request()` is a string, the default mode is "cors".
- The `referrer` property can be set to a same-origin URL, "about:client", or an empty string; if invalid parsing occurs, a `TypeError` is thrown.
- A `Request` object's `.mode` cannot be set to "navigate" via the constructor; doing so throws a `TypeError`.
- If `init["cache"]` is "only-if-cached" and mode is not "same-origin", a `TypeError` is thrown.
- The `.referrerPolicy` getter returns the policy associated with the request.

Entities And Concepts
- **Request**: The primary interface representing an HTTP request.
- **RequestInit**: Dictionary for constructor arguments.
- **BodyInit**: Type for request body data.
- **AbortSignal**: Used to abort requests.
- **Headers**: Object representing HTTP headers.
- **ReferrerPolicy**: Policy controlling referrer information (e.g., "strict-origin").
- **CORS-safelisted method**: Methods allowed when mode is "no-cors".

Procedures And API Details
- **Constructor Logic**:
  - Parses string input or extracts from an existing `Request` object.
  - Validates `init["window"]`; if non-null, throws `TypeError`.
  - Handles referrer parsing: if `parsedReferrer` scheme is "about" and path is "client" (and not same-origin), sets referrer to "client".
  - Sets default mode to "cors" for string inputs.
- **Body Handling**:
  - Throws `TypeError` if body exists and method is `GET` or `HEAD`.
  - Adds `Content-Type` header automatically if type is known and missing.
  - Handles `duplex` flag: required if body source is null and mode is neither "same-origin" nor "cors".
- **Clone Method**:
  - Requires the object not to be unusable.
  - Creates a dependent abort signal from the original signal.
  - Returns a new `Request` object with cloned request data and headers guard preserved.

Nuance Or Contradictions
- **Service Worker Context**: The text notes that "serviceworker" is omitted from `RequestDestination` because it cannot be observed from JavaScript, though implementations must support it internally. Similarly, "websocket" and "webtransport" are omitted from `RequestMode`.
- **Origin Propagation**: For navigation requests handled by a service worker, the request's origin might differ from the current client. The constructor explicitly sets the origin to "client" if `init` is not empty and mode is "navigate", altering the apparent source of the request (e.g., from an image in a cross-origin style sheet) to the service worker itself.
- **Referrer Logic**: There is specific logic where if `parsedReferrer` fails parsing or has certain characteristics (scheme "about" with path "client" while not same-origin), the referrer defaults to "client".

Candidate Wiki Hints
- **Page: Request Interface**: Documenting the `Request` IDL definition, attributes, and standard properties.
- **Page: Request Constructor Arguments**: Explaining the `init` dictionary parameters and their default behaviors (e.g., mode defaults to "cors").
- **Page: Referrer Handling in Requests**: Detailing how the `.referrer` property is computed, including handling of "about:client" and same-origin checks.
- **Page: Request Body Constraints**: Covering restrictions on bodies for `GET`/`HEAD` methods and the `duplex` flag usage.

## chunk-17

---
title: Chunk 17 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk covers the specification for the `Response` class (section 5.5), `Fetch` methods (section 5.6), and garbage collection rules regarding fetch termination (section 5.7). It details how to construct response objects, handle redirects, serialize data, abort requests, defer fetching until a document is active, and determine if a fetch can be terminated by the user agent based on observability of its state.

# Local Summary

The `Response` interface represents an HTTP response received from a server or generated by the browser. It includes properties like `status`, `url`, `headers`, and methods like `clone()`. Static factory methods include `error()`, `redirect()`, and `json()`. The chunk also defines the `fetch()` method for asynchronous resource retrieval and introduces `fetchLater()` for deferring requests until a document is fully active, subject to quota limits. Finally, it outlines garbage collection rules where user agents may terminate fetches if their termination cannot be observed via script (e.g., ignoring response bodies or headers without reading streams).

# Key Claims

- A `Response` object can be created directly from a body and initialization options, or via static methods like `Response.error()` and `Response.json()`.
- The `redirect()` method creates a redirect response with a specific status code (default 302) and updates the `Location` header.
- The `fetchLater()` method allows deferring a fetch request until a specified time (`activateAfter`) or when the document becomes fully active, but only for potentially trustworthy URLs with known body lengths.
- User agents may terminate an ongoing fetch if the termination is not observable through script arguments and return values.
- Accessing response headers without reading the body makes the fetch terminable, whereas registering a handler on `res.body.getReader().closed` prevents termination as the closure becomes observable.

# Entities And Concepts

- **Response**: An interface representing an HTTP response, including properties like `type`, `url`, `redirected`, `status`, `ok`, `statusText`, and `headers`.
- **Fetch**: A method to asynchronously fetch a resource, returning a Promise of a Response.
- **FetchLater**: A method to defer fetching until the document is active or a specified time passes.
- **Observable through script**: A condition determining whether a user agent can terminate a fetch; if termination isn't observable (e.g., ignoring response body), the fetch may be garbage collected.
- **DeferredRequestInit**: An extension of `RequestInit` that includes an `activateAfter` timestamp for deferred requests.
- **FetchLaterResult**: A result object from `fetchLater()` indicating whether the request has been activated.

# Procedures And API Details

### Creating a Response Object
To create a `Response` object given a response, headers guard, and realm:
1. Create a new `Response` object with the specified realm.
2. Set its associated response.
3. Initialize its headers with the appropriate guard.
4. Return the object.

### Static Methods
- **`Response.error()`**: Creates a network error response.
- **`Response.redirect(url, status)`**: Parses `url`, validates it's a redirect status, creates an immutable response, sets the status, and appends the serialized URL to the `Location` header.
- **`Response.json(data, init)`**: Serializes data to JSON bytes, extracts them as a body, creates a response, initializes it with the JSON type, and returns it.

### Fetch Method Logic
1. Create a new promise.
2. Invoke the `Request` constructor with input and init; reject if an exception occurs.
3. Check for abort signals on the request object.
4. Determine the relevant realm and client global object.
5. Add abort steps to handle signal abortion by rejecting the promise and cancelling bodies.
6. Call `fetch` with the request, process response by creating a Response object, resolving or rejecting the promise accordingly.

### FetchLater Method Logic
1. Invoke the `Request` constructor.
2. Validate the request signal is not aborted.
3. Extract `activateAfter` from init if provided; ensure it's non-negative.
4. Verify the document is fully active and the URL is HTTP(S) and potentially trustworthy.
5. Ensure body length is known (streams are not supported).
6. Check deferred-fetch quota against requested length; throw `QuotaExceededError` if exceeded.
7. Queue a deferred fetch record with the activate state set to false initially.
8. Return a `FetchLaterResult` object where the `activated` getter reflects the current state.

# Nuance Or Contradictions

- **Terminability**: The distinction between observable and non-observable fetch termination is nuanced. Simply accessing headers or status makes a fetch terminable, but reading from the body stream (via `getReader()`) and observing its closure prevents termination because the event becomes observable.
- **Deferred Fetch Constraints**: `fetchLater()` strictly requires known body lengths; dynamic streams are rejected immediately. Additionally, it only works on active windows, not detached ones like those in removed iframes.

# Candidate Wiki Hints

- **Response API Reference**: A page detailing all properties and static methods of the `Response` interface, including examples of creating responses for JSON, errors, and redirects.
- **Deferred Fetching Guide**: Documentation explaining `fetchLater()`, its parameters (`activateAfter`, body constraints), usage patterns with `AbortSignal`, and examples of activation checking.
- **Fetch Garbage Collection Rules**: An explainer on how user agents decide to terminate fetches based on observability, with code examples distinguishing between terminable and non-terminable scenarios.

## chunk-18

---
title: Chunk 18 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk covers section **6. data: URLs**, detailing the processing algorithm for `data:` URLs, followed by background information on HTTP header layers, CORS protocols, WebSockets, and general advice on setting up and invoking fetch operations.

# Local Summary

The document defines a `data:` URL as a struct containing a MIME type and a byte sequence body. It outlines a step-by-step processor that validates the scheme, parses the MIME type (handling base64 encoding), and decodes the body. The text further explains HTTP header layer divisions for fetching, safety considerations regarding redirects and CORS, the role of `Vary` headers in caching, and how WebSockets interact with fetch. Finally, it provides guidance on constructing requests, handling response bodies (completion vs. streaming), and manipulating ongoing fetch operations.

# Key Claims

- A `data:` URL struct consists of a MIME type and a body (byte sequence).
- The `data:` URL processor asserts the scheme is "data" and strips the leading prefix before parsing.
- If a MIME type ends with `;base64`, the body is isomorphically decoded and then forgiving-base64 decoded.
- Redirects are not exposed to APIs to prevent leaking secrets via cross-site scripting attacks.
- Using `Access-Control-Allow-Origin: *` is safe for resources accessible via curl/wget but unsafe for those protected by IP authentication or firewalls.
- The `Vary: Origin` header must be used if CORS requirements depend on the origin to prevent caching non-CORS responses for subsequent CORS requests.
- WebSockets initiate a special fetch with mode "websocket" and share policy decisions like HSTS.

# Entities And Concepts

- **data: URL**: A URL scheme carrying data directly in the URL string, defined by RFC 2397.
- **MIME Type**: Part of the `data:` URL struct indicating the media type of the body.
- **Fetch Controller**: An object returned by `fetch` used to abort, report timing, or handle manual redirects.
- **CORS (Cross-Origin Resource Sharing)**: A protocol mechanism allowing web pages on one domain to access resources from another domain via specific headers.
- **Vary Header**: Used in HTTP caching to indicate which request headers affect the response selection (e.g., `Vary: Origin`).
- **WebSockets**: A communication protocol that allows full-duplex communication over a single TCP connection, integrated with fetch via mode "websocket".

# Procedures And API Details

### Data: URL Processing Steps
1. Assert scheme is "data".
2. Serialize the URL excluding fragments.
3. Remove leading "data:" prefix.
4. Collect code points for MIME type (excluding commas).
5. Strip whitespace from MIME type.
6. If position reaches end, return failure.
7. Advance position; remainder becomes encoded body.
8. Percent-decode the encoded body.
9. Check if MIME type ends with `;base64` (case-insensitive):
    - Isomorphic decode the body string.
    - Forgiving-base64 decode the result.
    - Remove " base64" and the semicolon from MIME type.
10. If MIME type starts with ";", prepend "text/plain".
11. Parse MIME type; if failure, default to `text/plain;charset=US-ASCII`.
12. Return struct with parsed MIME type and decoded body.

### Request Setup Guidance
- Set URL and method (`POST`/`PUT` require a body).
- Choose destination (affects CSP and `Sec-Fetch-Dest`).
- Set client to environment settings object.
- Set mode: "same-origin" or "cors" (default for web-exposed features).
- Set credentials mode ("include") if cross-origin credentials are needed.
- Pass initiator type for Resource Timing reporting.
- Set header list for custom headers (may trigger CORS preflight).
- Set cache mode (e.g., "default", "no-store").
- Set redirect mode (e.g., "error" to disable redirects).

### Invoking Fetch Callbacks
- **processResponseConsumeBody**: Called upon completion; receives response and body contents (`null`, `failure`, or byte sequence).
- **processResponse**: Called when headers are received; caller streams the body. Useful for streaming video/images or handling non-OK status codes.
- **processResponseEndOfBody**: Called after fully reading the response body.
- **processEarlyHintsResponse**: For 103 Early Hints (currently handled by navigations).
- **useParallelQueue**: Set to `true` to run internal operations in a parallel queue, allowing interaction with the main thread via callbacks.

# Nuance Or Contradictions

- **CORS Safety**: The text clarifies that `Access-Control-Allow-Origin: *` is safe only if the resource is already publicly accessible (e.g., via curl/wget). If access requires IP auth or firewalls, CORS is unsafe regardless of headers.
- **Redirect Exposure**: Exposing redirects leaks secrets contained in redirect URLs. Therefore, redirects are not exposed to APIs by design.
- **Caching Behavior**: Without `Vary: Origin`, a user agent might cache a response lacking `Access-Control-Allow-Origin` from a non-CORS request and incorrectly reuse it for a subsequent CORS request, breaking the CORS check.

# Candidate Wiki Hints

- **data: URLs**: A dedicated page explaining the structure of `data:` URLs, their MIME type/body composition, and decoding rules (percent-decoding, base64 handling).
- **Fetch API Security**: Notes on redirect security, CORS header usage, and `Vary` header importance for caching.
- **Request Construction**: A reference guide for setting up fetch requests (URL, method, mode, credentials, cache mode, headers).
- **Response Handling**: Strategies for consuming responses fully versus streaming them chunk-by-chunk using callback arguments.

## chunk-19

---
title: Chunk 19 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
- **Heading**: 6. data: URLs
- **Location**: Chapter 6 of the Fetch Standard specification.
- **Content Focus**: Acknowledgments, copyright licensing, standard versioning (Living Standard vs. Review Draft), and index reference.

Local Summary
This section opens the "data: URLs" chapter with a comprehensive list of contributors thanked for their work on the standard. It establishes the legal framework by defining the copyright holder as WHATWG (a consortium including Apple, Google, Mozilla, and Microsoft) and specifying dual licensing: Creative Commons Attribution 4.0 International for the text and BSD 3-Clause License for incorporated source code portions. It directs readers seeking the patent-review version to the Living Standard Review Draft.

Key Claims
- The standard is authored by Anne van Kesteren of Apple.
- The work is licensed under CC BY 4.0, with a specific exception for source code portions which fall under the BSD 3-Clause License.
- The "Living Standard" is the current version; patent-review details are found in the Review Draft.

Entities And Concepts
- **WHATWG**: The working group and copyright holder (Apple, Google, Mozilla, Microsoft).
- **Anne van Kesteren**: Primary author of this standard.
- **Creative Commons Attribution 4.0 International License**: License for the specification text.
- **BSD 3-Clause License**: License for source code portions incorporated into the spec.
- **Living Standard**: The current version of the specification.
- **Living Standard Review Draft**: The version containing patent review details.

Procedures And API Details
- No specific procedures or API definitions are present in this chunk; it serves as introductory metadata and attribution.

Nuance Or Contradictions
- None observed in this specific chunk. The distinction between the standard text license (CC BY 4.0) and source code license (BSD 3-Clause) is a deliberate nuance rather than a contradiction.

Candidate Wiki Hints
- **data URLs**: A page introducing the concept of data: URLs within the Fetch Standard, including legal and attribution context.

## chunk-20

---
title: Chunk 20 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk covers Section **6. data: URLs** of the Fetch Standard specification. It focuses on the definitions, structures, and processing logic related to `data:` URIs within the HTTP fetch algorithm. The content defines how a `data:` URL is parsed into a specific structure and outlines the constraints and behaviors associated with fetching resources from such URLs (e.g., they are typically treated as opaque or require specific MIME type handling).

# Local Summary

The chunk details the `data: URL struct` defined in § 6. It explains that when a fetch request targets a `data:` URL, the algorithm must parse the URI into its constituent parts (scheme, base64 content, media type, charset). The text implies that `data:` URLs are handled as a specific case within the general fetch mechanism, often resulting in an immediate response generation without network I/O, subject to integrity checks and MIME type validation.

# Key Claims

*   **Structure Definition**: A `data: URL struct` is explicitly defined in § 6 to hold parsed components of a data URI.
*   **Processing Logic**: The standard includes specific steps for processing `data:` URLs, distinct from HTTP-network fetches.
*   **MIME Type Handling**: The system extracts and validates the MIME type associated with the data content within the URL structure.

# Entities And Concepts

*   **data: URL struct**: A structured representation of a parsed data URI.
*   **data: URL processor**: The algorithmic component responsible for handling `data:` URIs.
*   **MIME Type**: Extracted from the data URI to determine content handling (e.g., text, image).
*   **Fetch**: The high-level operation that may target a `data:` URL.

# Procedures And API Details

*   **Parsing**: The `data: URL processor` parses the input string into a structured format.
*   **Extraction**: The system extracts specific fields such as media type and charset from the data URI structure.
*   **Algorithm Steps**: Section 6 outlines the steps required to create and utilize a `data: URL struct`.

# Nuance Or Contradictions

The text distinguishes between standard HTTP fetches (network or cache) and the specialized handling of `data:` URLs, which do not involve network connections but require parsing and integrity verification. The term "opaque" is often associated with responses from CORS-preflight failures or specific security contexts, though this chunk focuses on the structural definition of data URIs themselves.

# Candidate Wiki Hints

*   **Topic**: `data: URL` handling in the Fetch Standard.
*   **Concept**: Structure and parsing of `data:` URIs within browser APIs.
*   **Reference**: Section 6 definitions for `data: URL struct`.

## chunk-21

---
title: Chunk 21 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk covers **Section 6: data: URLs** within the Fetch Standard. It defines how fetch algorithms handle requests where the URL is a `data:` URI, detailing the parsing of the MIME type and optional parameters from the fragment identifier, the creation of an empty body, and the interaction with CORS (Cross-Origin Resource Sharing) settings.

# Local Summary

The section establishes that when a request targets a `data:` URL, the algorithm must parse the URL to extract the media type and any associated metadata. It specifies that such requests have an empty body by default. The handling of `cross-origin` resources via `data:` URLs is restricted; specifically, a fetch from one origin for a `data:` URL from another origin (conceptually) is not permitted in the same way as network resources, often resulting in a CORS error if the resource is treated as cross-origin or if specific headers like `Cross-Origin-Opener-Policy` are violated. The text also notes that `data:` URLs cannot be used to load resources that require specific network-level behaviors (like redirects) and that they are generally processed synchronously within the context of the initiator unless explicitly handled otherwise.

# Key Claims

*   **Parsing**: The algorithm parses the `data:` URL to separate the media type from any optional parameters in the fragment.
*   **Body Content**: Requests for `data:` URLs result in an empty body unless specific logic dictates otherwise (though typically, `data:` URIs contain the data in the fragment). *Correction based on standard behavior implied*: The text implies the data is read from the URL's fragment as the body.
*   **CORS Restrictions**: Fetching a `data:` URL from a different origin than the initiator is generally not allowed or treated as an error depending on the specific CORS context (specifically regarding `Cross-Origin-Opener-Policy`).
*   **Redirection**: `data:` URLs do not support redirection; any attempt to redirect would fail.

# Entities And Concepts

*   **data: URL**: A Uniform Resource Identifier that contains data directly in the URI rather than pointing to a remote resource.
*   **MIME Type**: The media type identifier extracted from the `data:` URL fragment (e.g., `text/plain`, `image/png`).
*   **Cross-Origin-Opener-Policy**: A security policy mentioned in relation to `data:` URLs, restricting how they can be accessed across different browsing contexts.
*   **Fetch Algorithm**: The set of steps defined by the Fetch Standard for handling network requests and `data:` URLs.

# Procedures And API Details

*   **Parse Data URL**: The algorithm involves parsing the URL string to identify the media type and parameters.
*   **Create Request**: A new request object is created with an empty body initially, which is then populated if the data is read from the URL fragment.
*   **CORS Check**: Before processing the `data:` URL, the system checks for CORS violations, particularly regarding opener policies.

# Nuance Or Contradictions

The text implies a distinction between loading resources that require network behavior versus those that are self-contained. There is a potential nuance in how "cross-origin" applies to `data:` URLs compared to HTTP resources; since `data:` URLs are local to the document, the concept of "origin" for the resource itself differs from the initiator's origin in network contexts.

# Candidate Wiki Hints

*   **data: URL Structure**: How to parse and use `data:` URIs in web applications.
*   **CORS with data: URLs**: Security implications and restrictions when using `data:` URIs across origins.
*   **Fetch API and Local Resources**: Differences between fetching network resources and local `data:` resources.

## chunk-22

---
title: Chunk 22 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk defines the "6. data: URLs" section of a web specification, listing all external standards and specifications referenced to define terms used within that context. It includes normative references (W3C Living Standards, RFCs) and non-normative references (security advisories). The chunk concludes with an IDL Index defining the `Headers` and `Request` interfaces relevant to fetching data.

# Local Summary

The document defines terminology for handling "data:" URLs by referencing a wide array of Web Platform APIs and standards, including cookies, content security policies, fetch metadata, file APIs, HTTP caching, mixed content, permissions policy, reporting, resource timing, secure contexts, subresource integrity, streams, service workers, upgrade-insecure-requests, URL parsing, web crypto, web driver bidi, web IDL, websockets, and web transport. It also lists specific normative and non-normative references for these standards. Finally, it provides an IDL Index defining the `Headers` interface (with methods like `append`, `delete`, `get`, `set`) and the `Request` interface (with attributes like `method`, `url`, `headers`), along with typedefs for `BodyInit` and `XMLHttpRequestRequestBodyInit`.

# Key Claims

*   The terms defined in section 6 ("data: URLs") are inherited from or defined by specific external specifications.
*   A comprehensive list of normative references (e.g., [HTML], [FETCH-METADATA], [HTTP-CACHING]) is provided to define the terminology used.
*   A list of non-normative references (e.g., security advisories like [HTTPVERBSEC1]) is provided for context.
*   The `Headers` interface allows manipulation of request/response headers via methods like `append`, `delete`, `get`, `has`, and `set`.
*   The `Request` interface represents a request, possessing attributes such as `method`, `url`, and `headers`.
*   Typedefs exist for `BodyInit` (representing the body of a fetch request) and `XMLHttpRequestRequestBodyInit`.

# Entities And Concepts

*   **data: URLs**: The subject of section 6.
*   **Normative References**: Standards that define behavior (e.g., [HTML], [RFC9110]).
*   **Non-Normative References**: Informational documents or advisories (e.g., [HTTPVERBSEC1]).
*   **Headers Interface**: An interface for managing HTTP headers.
    *   `append(name, value)`: Adds a header.
    *   `delete(name)`: Removes a header.
    *   `get(name)`: Retrieves a header value.
    *   `has(name)`: Checks if a header exists.
    *   `set(name, value)`: Sets a header value.
*   **Request Interface**: Represents an HTTP request.
    *   `method`: The HTTP method (e.g., GET, POST).
    *   `url`: The request URL.
    *   `headers`: A `Headers` object containing the request headers.
*   **BodyInit**: A type representing possible values for a fetch request body (`ReadableStream`, `XMLHttpRequestBodyInit`).
*   **XMLHttpRequestRequestBodyInit**: A type for XHR request bodies (Blob, BufferSource, FormData, etc.).

# Procedures And API Details

*   **Headers Construction**: The `Headers` interface can be constructed using an optional `HeadersInit` argument.
*   **Header Manipulation**:
    *   `append(ByteString name, ByteString value)`: Appends a header field with the given name and value.
    *   `delete(ByteString name)`: Deletes all headers whose name matches the provided string (case-insensitive).
    *   `get(ByteString name)`: Returns the value of the first header whose name matches the provided string, or null if no match is found.
    *   `has(ByteString name)`: Returns true if there are any headers whose name matches the provided string.
    *   `set(ByteString name, ByteString value)`: Sets the header's value to the given value, replacing any existing headers with that name.
*   **Request Construction**: The `Request` interface is constructed using a `RequestInfo` input (either a `Request` object or a USVString URL) and an optional `RequestInit` dictionary.
*   **Body Access**: The `Body` mixin provides methods to read the request/response body in various formats:
    *   `arrayBuffer()`: Returns a Promise resolving to an `ArrayBuffer`.
    *   `blob()`: Returns a Promise resolving to a `Blob`.
    *   `bytes()`: Returns a Promise resolving to a `Uint8Array`.
    *   `formData()`: Returns a Promise resolving to a `FormData` object.
    *   `json()`: Returns a Promise resolving to the parsed JSON value.
    *   `text()`: Returns a Promise resolving to a USVString.

# Nuance Or Contradictions

*   The document distinguishes between "Normative References" (which define the standard) and "Non-Normative References" (which provide context or security advisories).
*   The `Headers` interface supports both appending new headers (`append`) and setting/replacing existing ones (`set`).
*   The `get` method returns the *first* matching header value, implying multiple headers with the same name are possible.

# Candidate Wiki Hints

*   **data: URLs**: A brief overview of data URL schemes and their usage in web contexts.
*   **Web Platform References**: A page linking to key W3C standards (HTML, Fetch, Cookies, etc.) for developers needing precise definitions.
*   **Request/Response Headers API**: Documentation for the `Headers` interface methods (`append`, `delete`, `get`, `has`, `set`).
*   **Fetch Request Body Types**: Explaining `BodyInit` and how to provide different types of request bodies (streams, blobs, forms).
*   **Security References**: A curated list of security-related RFCs and advisories for web developers.

## chunk-23

---
title: Chunk 23 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

**Heading:** 6. data: URLs
**Source Range:** Lines 7886–8596 of `raw/web/corpus-2026-05-18/058-fetch-standard.md`
**Coverage:** Entire section "6. data: URLs"

# Local Summary

This chunk details the Fetch API specification, focusing on the `Request` and `Response` interfaces, their associated dictionaries (`RequestInit`, `ResponseInit`), and a comprehensive list of browser compatibility for various attributes and methods (e.g., `fetch`, `clone`, `headers`). It defines request modes, credential handling, cache directives, and response types. A significant portion is dedicated to a compatibility matrix listing support versions for Chrome, Firefox, Safari, Edge, Opera, and Node.js across numerous properties like `body`, `json`, `text`, `arrayBuffer`, `blob`, and `formData`.

# Key Claims

- The Fetch API includes the `Request` interface with attributes such as `destination`, `referrer`, `mode`, `credentials`, `cache`, `redirect`, `integrity`, `keepalive`, `isReloadNavigation`, `isHistoryNavigation`, `signal`, and `duplex`.
- The `RequestInit` dictionary defines optional initialization parameters for requests, including `method`, `headers`, `body`, `referrer`, `mode`, `credentials`, `cache`, `redirect`, `integrity`, `keepalive`, `signal`, `duplex`, `priority`, and `window`.
- Enums define valid values for request attributes: `RequestDestination` (e.g., "audio", "document"), `RequestMode` ("navigate", "cors"), `RequestCredentials` ("omit", "include"), `RequestCache` ("no-store", "reload"), `RequestRedirect` ("follow", "error"), `RequestDuplex` ("half"), and `RequestPriority` ("high", "low").
- The `Response` interface provides methods like `static Response error()`, `static Response redirect()`, and `static Response json()` to create responses. It includes attributes `type`, `url`, `redirected`, `status`, `ok`, `statusText`, and `headers`.
- The `ResponseType` enum lists values: "basic", "cors", "default", "error", "opaque", "opaqueredirect".
- The global scope (Window/Worker) exposes the `fetch()` method via `partial interface mixin WindowOrWorkerGlobalScope`.
- A new `fetchLater()` method is exposed on the Window interface, returning a `FetchLaterResult`, with support for a `DeferredRequestInit` dictionary containing an `activateAfter` timestamp.
- Extensive compatibility data confirms that most core Fetch API features are supported in "all current engines" (Firefox 39+, Safari 10.1+, Chrome 40/42+, Edge 79+, Opera 27+), with specific exceptions noted for older versions or mobile browsers.

# Entities And Concepts

- **Interfaces:** `Request`, `Response`, `WindowOrWorkerGlobalScope`, `Window`
- **Dictionaries:** `RequestInit`, `ResponseInit`, `DeferredRequestInit`, `FetchLaterResult`
- **Enums:** `RequestDestination`, `RequestMode`, `RequestCredentials`, `RequestCache`, `RequestRedirect`, `RequestDuplex`, `RequestPriority`, `ResponseType`
- **Methods:** `fetch()`, `clone()` (on Request/Response), `error()`, `redirect()`, `json()` (on Response), `fetchLater()`
- **Attributes:** `destination`, `referrer`, `referrerPolicy`, `mode`, `credentials`, `cache`, `redirect`, `integrity`, `keepalive`, `isReloadNavigation`, `isHistoryNavigation`, `signal`, `duplex`, `priority`, `url`, `redirected`, `status`, `ok`, `statusText`, `type`, `headers`, `body`, `bodyUsed`, `arrayBuffer`, `blob`, `formData`, `json`, `text`.
- **Browser Engines:** Chrome, Firefox, Safari, Edge (Legacy/Current), Opera, Node.js, Android WebView, Samsung Internet.

# Procedures And API Details

**Creating a Request:**
```javascript
const req = new Request(input, init);
// Or using global fetch with init:
fetch(input, init)
```

**Request Attributes & Init Options:**
- `method`: ByteString (e.g., "GET", "POST")
- `headers`: HeadersInit
- `body`: BodyInit?
- `referrer`: USVString
- `referrerPolicy`: ReferrerPolicy
- `mode`: RequestMode ("navigate", "same-origin", "no-cors", "cors")
- `credentials`: RequestCredentials ("omit", "same-origin", "include")
- `cache`: RequestCache ("default", "no-store", "reload", "no-cache", "force-cache", "only-if-cached")
- `redirect`: RequestRedirect ("follow", "error", "manual")
- `integrity`: DOMString (SubtleCrypto hash)
- `keepalive`: boolean
- `signal`: AbortSignal?
- `duplex`: RequestDuplex ("half")
- `priority`: RequestPriority ("high", "low", "auto")
- `window`: any (can be set to null)

**Response Attributes & Init Options:**
- `status`: unsigned short (default 200)
- `statusText`: ByteString (default "")
- `headers`: HeadersInit
- `type`: ResponseType ("basic", "cors", "default", "error", "opaque", "opaqueredirect")

**Response Creation Methods:**
- `static Response error()`: Creates an error response.
- `static Response redirect(url, status = 302)`: Creates a redirect response.
- `static Response json(data, init = {})`: Serializes data to JSON and creates a response with appropriate headers (`Content-Type: application/json`).

**Request Cloning:**
```javascript
const clone = request.clone(); // Returns a new Request object with the same properties (body stream is duplicated)
```

**Response Cloning:**
```javascript
const clone = response.clone(); // Returns a new Response object (body stream is duplicated)
```

**Deferred Fetching:**
```javascript
// Window interface
const result = window.fetchLater(input, { activateAfter: timestamp });
// result.activated indicates if the request was activated
```

# Nuance Or Contradictions

- **Body Readability:** `Request.body` is not supported in Firefox until version 65+, whereas Safari 11.1+ and Chrome 105+ support it earlier. This contradicts the "In all current engines" claim for older versions or specific mobile browsers listed with question marks.
- **GetSetCookie Header:** `Headers.getSetCookie()` is only supported in newer versions (Firefox 112+, Safari 17+, Chrome 113+) compared to standard `get()`, indicating a nuance in handling multiple Set-Cookie headers.
- **Mobile Browser Support:** Many mobile browsers (iOS Safari, Android WebView, Samsung Internet, Opera Mobile) have question marks or specific version constraints that differ from desktop counterparts, suggesting implementation gaps or timing differences.
- **Legacy Edge:** "Edge (Legacy)" (based on EdgeHTML) shows support starting at version 14 for many features, but `Response.body` is marked as unsupported ("IENone"), highlighting a discontinuity between legacy and current Chromium-based Edge.

# Candidate Wiki Hints

1.  **Page: Fetch API Overview**
    *   **Topic:** Introduction to the Fetch API, Request/Response lifecycle.
    *   **Source Support:** Definitions of `Request`, `Response`, `fetch()`.

2.  **Page: Request and Response Attributes**
    *   **Topic:** Detailed breakdown of `destination`, `referrerPolicy`, `cache`, `redirect`, `integrity`, etc.
    *   **Source Support:** Enum definitions (`RequestDestination`, `RequestMode`) and attribute lists.

3.  **Page: Creating Responses**
    *   **Topic:** Using `Response.error()`, `Response.redirect()`, `Response.json()`.
    *   **Source Support:** Static methods of the `Response` interface.

4.  **Page: Deferred Fetching with fetchLater()**
    *   **Topic:** The experimental or newer `fetchLater()` method for lazy loading.
    *   **Source Support:** `FetchLaterResult`, `DeferredRequestInit`, `window.fetchLater()`.

5.  **Page: Browser Compatibility Guide (Fetch API)**
    *   **Topic:** A reference table for attribute/method support across Chrome, Firefox, Safari, Edge, Opera, Node.js.
    *   **Source Support:** The extensive compatibility matrix in the chunk text.

## chunk-24

---
title: Chunk 24 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This section details browser and engine support for the Fetch API and associated Headers. It covers `Request` methods, various `Response` properties (including static methods like `json_static`), specific response status handling (`redirected`, `statusText`), and a comprehensive list of CORS-related headers (`Access-Control-*`, `Cross-Origin-Resource-Policy`, `Sec-Purpose`) along with legacy header support.

# Local Summary

The chunk outlines compatibility matrices for Fetch API features across major browsers (Firefox, Safari, Chrome, Edge, Opera) and their variants (Android, iOS WebView, Samsung Internet). It specifies version requirements for standard properties like `Response.url` and `fetch`, noting that some features are supported in "all current engines" while others have specific version cutoffs (e.g., `Response/json_static` requires Firefox 115+ or Chrome 105+). The text also highlights headers available in all engines versus those restricted to single engines or marked with warnings regarding MDN documentation status.

# Key Claims

- **Universal Support**: Properties like `Response/url`, `Response/status`, and the `fetch` function itself are supported in "all current engines" starting from Firefox 39, Safari 10.1, and Chrome 42.
- **Version Specifics**: `Response/json_static` is not universally available; it requires Firefox 115+, Chrome 105+, or Edge 105+.
- **Legacy Support**: `Edge (Legacy)` supports many features starting from version 14, while `IE` support varies or is marked as non-existent for modern standards.
- **CORS Headers**: Standard CORS headers (`Access-Control-Allow-Origin`, etc.) are supported in Firefox 3.5+, Safari 4+, and Chrome 4+.
- **Warning Flags**: Certain headers like `Headers/Origin` and features in Android WebViews carry warning indicators (⚠) or question marks, suggesting potential inconsistencies or incomplete documentation.

# Entities And Concepts

- **Fetch API**: The core interface for making network requests (`fetch`).
- **Response Object**: Properties including `type`, `url`, `status`, `redirected`, and static methods like `json_static`.
- **Request Object**: Represents the request initiated by `fetch`.
- **CORS Headers**: A set of headers used to control cross-origin resource sharing (e.g., `Access-Control-Allow-Credentials`, `Cross-Origin-Resource-Policy`).
- **Browser Variants**: Specific builds for mobile operating systems such as Firefox for Android, iOS Safari, and Samsung Internet.

# Procedures And API Details

- **Checking Support**: The text implies a procedure of checking version numbers (e.g., "Firefox 39+") to determine if a specific feature or property is available in a given engine.
- **Static Methods**: `Response.json_static` is identified as a method that parses response data, with strict version requirements for its availability.
- **Header Enumeration**: The chunk lists specific header names (e.g., `Access-Control-Expose-Headers`) and indicates their support status across different browser versions.

# Nuance Or Contradictions

- **Documentation Status**: Some entries are marked with "⚠MDN" or question marks, indicating that while the feature might exist in the browser, its documentation or standardization status is uncertain compared to those marked simply with "✔MDN".
- **Engine Variations**: While desktop browsers show clear version numbers (e.g., Chrome 42+), mobile variants often have missing data points (represented by "?"), suggesting gaps in testing or reporting for those specific environments.
- **Legacy vs. Modern**: There is a distinction made between modern Edge and "Edge (Legacy)", with the latter supporting older versions (14+) but having different compatibility profiles than the modern Chromium-based Edge.

# Candidate Wiki Hints

- **Fetch API Compatibility Table**: Create a dedicated page or section summarizing browser support for Fetch API properties, distinguishing between universal support and version-specific requirements.
- **CORS Headers Reference**: Compile a list of CORS-related headers with their minimum browser versions and any known limitations or partial implementations in mobile webviews.
- **Response Object Properties**: Document the `Response` object properties, specifically highlighting the availability of static methods like `json_static` across different browsers.

