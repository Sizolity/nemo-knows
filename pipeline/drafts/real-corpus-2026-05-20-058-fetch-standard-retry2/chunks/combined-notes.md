## chunk-01

---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
- Source: Fetch Standard (Living Standard)
- Lines: 1–447
- Coverage: Document metadata, Preface, and Infrastructure (URL schemes, HTTP basics, Method normalization).

Local Summary
The Fetch Standard unifies fetching across the web platform, handling URL schemes, redirects, cross-origin semantics, CSP, Service Workers, and Mixed Content. It supersedes previous `Origin` header semantics. The document defines a unified architecture for APIs like `<img>`, `<script>`, `navigator.sendBeacon()`, and the `fetch()` API, ensuring consistent behavior for redirects and CORS.

Key Claims
- Fetching is conceptually simple (request in, response out) but involves complex details previously inconsistent across APIs.
- The standard covers all URL schemes including "about", "blob", "data", "file", and HTTP(S).
- `fetch()` exposes low-level networking functionality via a unified API.
- Specific HTTP methods are categorized: CORS-safelisted (`GET`, `HEAD`, `POST`), forbidden (`CONNECT`, `TRACE`, `TRACK`), and normalized (uppercase for `DELETE`, `GET`, etc.).
- Methods like `PATCH` are preferred over `patch` to avoid 405 errors due to normalization rules.

Entities And Concepts
- **Fetch Standard**: The specification defining requests, responses, and the fetching process.
- **Fetch Controller**: A struct enabling callers to perform operations after a fetch starts (state, timing info, abort reason).
- **Fetch Params**: A bookkeeping struct containing request details, body processing algorithms, task destinations, and controller references.
- **URL Schemes**: Local schemes ("about", "blob", "data"), HTTP(S) schemes ("http", "https"), and fetch schemes (including "file").
- **HTTP Methods**: Normalized methods (`GET`, `POST`, etc.) vs. non-normalized custom methods.

Procedures And API Details
- **Method Normalization**: Byte-uppercase methods matching `DELETE`, `GET`, `HEAD`, `OPTIONS`, `POST`, or `PUT` for backwards compatibility and consistency, though methods are technically case-sensitive.
- **Fetch Controller State**: Can be "ongoing", "terminated", or "aborted". Aborting sets state to "aborted" and serializes an error; terminating sets state to "terminated".
- **Timing Info Struct**: Tracks start time, redirect times, network request/response times, and service worker timing.
- **Body Processing**: Includes algorithms for processing request body chunks, end-of-body events, and response consumption.

Nuance Or Contradictions
- **Case Sensitivity vs. Normalization**: While methods are technically case-sensitive (e.g., `Egg` is valid), normalization forces uppercase for standard verbs. Using lowercase standard verbs like `patch` is likely to fail with 405, whereas `PATCH` succeeds.
- **Whitespace Handling**: HTTP tab/space (U+0009/U+0020) is preferred for header values over general ASCII whitespace, which excludes U+000C (FF).

Candidate Wiki Hints
- Page: **Fetch Standard Overview** (Introduction to goals and unified architecture)
- Page: **HTTP Method Normalization** (Rules for uppercase conversion and forbidden methods)
- Page: **Fetch Controller Lifecycle** (States, aborting, and timing info)

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

### Chunk Context
This chunk covers **Section 2.2.2 Headers** within the Fetch Standard, detailing HTTP header management in the browser environment. It defines headers as ordered lists of key-value pairs (multimaps), explaining serialization/parsing logic, splitting algorithms for comma-separated values, and specific handling for `Set-Cookie` headers. The section concludes with a transition to **Section 2.2.3 Statuses**, defining HTTP status code ranges (e.g., OK, Redirect) and null body statuses.

### Local Summary
The document establishes that an HTTP header list is an ordered collection of name-value pairs where duplicate keys are permitted but non-`Set-Cookie` headers are combined into single values when exposed to JavaScript. It details the algorithms for getting, setting, deleting, and combining headers, including a specific sorting mechanism that preserves `Set-Cookie` entries individually while merging others. The text also defines "structured field values" (currently byte sequences) and provides rigorous rules for splitting header values by comma while respecting quoted strings. Finally, it outlines CORS safety lists, forbidden request/response headers, and the classification of HTTP status codes.

### Key Claims
- **Header Representation**: A header list is a specialized multimap; implementations may optimize storage for non-`Set-Cookie` headers since they are combined on the client side.
- **Structured Fields**: Currently, Fetch supports header values only as byte sequences; future versions might preserve object structures end-to-end (RFC9651).
- **Splitting Logic**: Header values are split by comma (`U+002C`) unless enclosed in double quotes (`U+0022`), which allows for complex multi-value headers like `Accept`.
- **CORS Safety**: A header is "CORS-safelisted" only if it passes length checks (128 bytes) and contains only safe characters or specific MIME types. Headers starting with `Sec-` are reserved for future-safe APIs.
- **Forbidden Headers**: Specific headers like `Accept`, `Cookie`, `Host`, and those prefixed with `proxy-` or `sec-` are forbidden on requests to prevent user agent bypassing controls.

### Entities And Concepts
- **Header List**: An ordered list of zero or more headers (key-value pairs).
- **Structured Field Value**: Objects that HTTP can serialize efficiently; currently represented as byte sequences in Fetch.
- **CORS-Safelisted Request Header**: Headers allowed to be sent with cross-origin requests, subject to length and character restrictions.
- **Forbidden Request Header**: Headers restricted to maintain user agent control (e.g., `Cookie`, `Host`, `Content-Length`).
- **Range Status**: HTTP status codes 206 or 416 indicating partial content or range not satisfiable.
- **Null Body Status**: Status codes 101, 103, 204, 205, or 304 where the response body is empty or irrelevant.

### Procedures And API Details
- **Getting a Header**: Returns `null` if absent; otherwise returns all values separated by `0x2C 0x20`.
- **Parsing Structured Fields**: Converts byte sequences into structured objects (currently limited to byte sequence representation).
- **Sorting Headers**: Converts header names to a sorted-lowercase set, preserving individual `Set-Cookie` entries while combining others.
- **Building Content Range**: Constructs the `bytes <start>-<end>/<fullLength>` string for partial content requests.
- **CORS Safelist Check**: Validates header length and character sets; rejects headers with unsafe bytes (e.g., `<`, `>`, `:`) or disallowed MIME types.

### Nuance Or Contradictions
- **Set-Cookie Handling**: Unlike other headers, `Set-Cookie` cannot be combined and requires complex handling in the Headers object to avoid leaking this complexity into requests; it is semantically a response header but appears in request contexts as forbidden.
- **MIME Type Parsing**: The standard explicitly avoids using the "extract a MIME type" algorithm because it is too forgiving, preventing servers from misinterpreting request bodies (e.g., treating a `text/plain` body as JSON).
- **Range Syntax**: While `bytes=-500` is syntactically valid per the range parsing rules, browsers historically do not emit such ranges, so they are excluded from the safelist.

### Candidate Wiki Hints
- **Topic: HTTP Headers in Fetch API** – Explaining how headers are stored as ordered lists and combined for client-side access.
- **Topic: CORS Header Safety** – Detailing the logic behind `CORS-safelisted` vs. forbidden request headers and byte restrictions.
- **Topic: Parsing Multi-Value Headers** – Algorithms for splitting comma-separated values while handling quoted strings (e.g., `Accept`, `Link`).
- **Topic: Range Requests** – Implementation details of parsing `Range` headers and constructing content range responses.

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context
This chunk details the structural definition of a `body` object within the Fetch Standard, specifically covering sections 2.2.4 (Bodies) and associated algorithms for cloning, incrementally reading, fully reading, and handling content codings. It outlines the lifecycle of a body's stream, source, and length, and defines the microsteps required to process data chunks via parallel queues.

# Local Summary
A `body` is defined by three members: a `stream` (a ReadableStream), a `source` (initially null, representing raw data or a Blob/FormData), and a `length` (initially null). The chunk describes the procedure to clone a body by teeing its stream. It further defines two primary reading algorithms: "incrementally read," which processes chunks as they arrive using provided callback algorithms (`processBodyChunk`, `processEndOfBody`, `processBodyError`), and "fully read," which consumes the entire stream before invoking a final processing algorithm. Both reading mechanisms utilize parallel queues or global objects to manage asynchronous tasks. Finally, it introduces the concept of handling content codings by decoding bytes according to HTTP standards if supported.

# Key Claims
- A body is composed of a `stream`, a `source`, and a `length`.
- Cloning a body involves teeing the stream into two outputs (`out1` and `out2`) and assigning them to the cloned body's members.
- Incremental reading requires specific algorithms for handling chunks, end-of-body signals, and errors.
- The incrementally-read loop queues fetch tasks based on the state of the reader (chunk received, close requested, or error encountered).
- Fully reading a body involves reading all bytes from the stream and invoking a success callback with the full byte sequence or an error callback if reading fails.
- Content coding handling attempts to decode bytes using provided codings; if unsupported or decoding fails, the raw bytes are returned.

# Entities And Concepts
- **Body**: An object representing the payload of a request or response.
- **ReadableStream**: The underlying stream mechanism used by a body.
- **Source**: The origin of the data (null, byte sequence, Blob, or FormData).
- **Cloning**: Creating a new body instance sharing the same underlying stream via `tee()`.
- **Parallel Queue**: An execution context used to queue fetch tasks without blocking the main thread.
- **Content Codings**: Headers indicating compression or encoding applied to the payload (e.g., gzip, deflate).
- **Fetch Task**: A microtask queued for asynchronous execution.

# Procedures And API Details
**Cloning a Body**
1.  Let `« out1, out2 »` be the result of `tee()` on the body's stream.
2.  Set the body's stream to `out1`.
3.  Return a new body object where the stream is `out2` and other members are copied from the original.

**Incrementally Reading a Body**
Inputs: `body`, `processBodyChunk`, `processEndOfBody`, `processBodyError`, optional `taskDestination`.
Steps:
1.  Initialize `taskDestination` to a new parallel queue if null.
2.  Get a reader for the body's stream.
3.  Execute the "incrementally-read loop" with the reader and callbacks.

**The Incrementally-Read Loop Logic**
Inputs: `reader`, `taskDestination`, `processBodyChunk`, `processEndOfBody`, `processBodyError`.
Loop Actions:
-   **Chunk Steps**:
    -   If chunk is not a `Uint8Array`, run `processBodyError` with a `TypeError`.
    -   Otherwise, copy the chunk (implementation should minimize this copy).
    -   Run `processBodyChunk` with the bytes.
    -   Recursively perform the incrementally-read loop.
-   **Close Steps**: Queue a fetch task to run `processEndOfBody`.
-   **Error Steps**: Queue a fetch task to run `processBodyError` with the exception.

**Fully Reading a Body**
Inputs: `body`, `processBody`, `processBodyError`, optional `taskDestination`.
Steps:
1.  Initialize `taskDestination` if null.
2.  Define success steps: Queue a fetch task to run `processBody` with bytes.
3.  Define error steps: Queue a fetch task to run `processBodyError` with the exception.
4.  Get a reader for the body's stream; if getting the reader throws, run error steps and return.
5.  Read all bytes from the reader using success and error steps.

**Handling Content Codings**
Inputs: `codings`, `bytes`.
Steps:
1.  If `codings` are not supported, return `bytes`.
2.  Return the result of decoding `bytes` with `codings` per HTTP specifications, or failure if decoding errors occur.

# Nuance Or Contradictions
- **Stream Sharing**: Cloning a body does not duplicate the stream data but splits the single underlying stream into two distinct readable streams (`out1` and `out2`). Consuming one will affect the other.
- **Copy Minimization**: The spec explicitly encourages implementations to avoid copying the chunk bytes when possible, suggesting that the internal buffer might be used directly if safe.
- **Task Destination**: Both incremental and full reading default to creating a new parallel queue if none is provided, ensuring non-blocking execution of callbacks.

# Candidate Wiki Hints
- **Body Object Structure**: A page defining the members of a `body` (stream, source, length) and their initial states.
- **Body Cloning**: A guide on how to duplicate a request/response body while maintaining stream integrity using `tee()`.
- **Reading Algorithms**: A comparative overview of "incrementally reading" vs. "fully reading," detailing when to use callbacks for chunks versus the full payload.
- **Content Decoding**: An explanation of the content coding handling process, including fallback behavior when codings are unsupported.

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- **Source**: `raw/web/corpus-2026-05-18/058-fetch-standard.md`
- **Heading**: 2. Read a chunk from reader given readRequest. > 2.2.5. Requests
- **Range**: Lines 1045–1499
- **Scope**: Detailed specification of the `request` input to the fetch algorithm, covering attributes, flags, modes, and their interactions with CORS, CSP, and navigation logic.

Local Summary
This section defines the structure and behavior of a `request` object in the Fetch standard. It details every attribute associated with a request (e.g., method, URL, headers, body, client) and explains their default states. The chunk further elaborates on specialized flags like `keepalive`, `unsafe-request`, and `use-URL-credentials`. A significant portion is dedicated to `mode` (same-origin, cors, no-cors, navigate), `cache mode` behaviors, and the relationship between `initiator`, `destination`, and Content Security Policy (CSP) directives.

Key Claims
- The default method for a request is `GET`.
- The default request mode is `no-cors`, which restricts methods to safe ones and returns an opaque response.
- The default credentials mode is `same-origin`, though this changes to `include` when the mode is `navigate`.
- A request's `origin` starts as `"client"` and resolves to a specific origin during fetching.
- `cache-mode: "default"` allows fetch to inspect the HTTP cache for fresh or stale responses before or alongside network requests.
- The `keepalive` flag allows requests (like those from `navigator.sendBeacon()` or `<img>`) to outlive their environment settings object.

Entities And Concepts
- **Request Attributes**: method, URL, header list, body, client, origin, referrer, credentials mode, cache mode, redirect mode, destination.
- **Modes**: `same-origin`, `cors`, `no-cors`, `navigate`, `websocket`, `webtransport`.
- **Cache Modes**: `default`, `no-store`, `reload`, `no-cache`, `force-cache`, `only-if-cached`.
- **Initiators**: Sources like `fetch`, `script`, `link`, `ping`, etc., mapped to specific CSP directives (e.g., `connect-src`).
- **Destination Types**: Categories like `document`, `script`, `style`, `image` used for CSP and routing.
- **Flags**: `keepalive`, `unsafe-request`, `use-CORS-preflight`, `timing allow failed`.

Procedures And API Details
- **Setting Up a Request**: Start by seeing "Setting up a request" (referenced at the start of the section).
- **CSP Hooks**: The `form-action` CSP directive must hook directly into HTML's navigate or form submission algorithms.
- **Cache Behavior**:
  - `default`: Inspects HTTP cache; returns fresh/stale responses or performs conditional/non-conditional network fetches.
  - `reload`: Ignores cache, creates a normal request, updates cache.
  - `only-if-cached`: Returns cached response or network error (requires same-origin mode).
- **CORS Preflight**: Triggered if event listeners are on an `XMLHttpRequestUpload` or a `ReadableStream` is used in a request.

Nuance Or Contradictions
- **Default Mode Safety**: Although `no-cors` is the default, standards are discouraged from using it for new features due to safety concerns (opaque responses).
- **Credentials vs URL**: Modern specs avoid setting the `use-URL-credentials` flag because putting credentials in URLs is discouraged.
- **Navigational Requests**: When mode is `navigate`, the credentials mode is assumed to be `include`, overriding other values.

Candidate Wiki Hints
- **Page: Fetch Request Modes** – Explain the differences between `cors`, `no-cors`, and `same-origin`, including their impact on response tainting and allowed methods.
- **Page: HTTP Cache Control in Fetch** – Detail how `cache-mode` affects interaction with the HTTP cache (freshness, staleness, revalidation).
- **Page: Content Security Policy and Fetch** – Map request initiators and destinations to CSP directives (e.g., `script-src`, `connect-src`).
- **Page: Request Attributes Reference** – A comprehensive list of all `request` attributes with their default values and usage contexts.

## chunk-05

---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- Heading path: `2. Read a chunk from reader given readRequest. > 2.2.5. Requests`
- Line range: 1501–1595

Local Summary
This section defines request classifications (subresource vs non-subresource vs navigation), the algorithm to compute redirect-taint, origin serialization rules for reporting, request cloning mechanics (including WebDriver ID handling), Range header construction, and a security warning regarding partial response aggregation. It concludes with an algorithm to check Cross-Origin-Embedder-Policy credential allowance.

Key Claims
- Request destinations determine their classification: "audio", "audioworklet", "font", etc., are subresources; "document", "embed", etc., are non-subresources; a subset of the latter forms navigation requests.
- Redirect-taint is computed by iterating through a request’s URL list, comparing origins between consecutive URLs and the current origin to yield "same-origin", "same-site", or "cross-site".
- Origin serialization returns `null` if redirect-taint is not "same-origin" to avoid leaking redirect target information.
- Cloning a request copies all fields except body and WebDriver id; the new request receives a freshly generated random UUID as its WebDriver id.
- Range headers denote inclusive byte ranges where `first=0` and `last=500` implies 501 bytes.
- Features combining multiple responses into one logical resource are historically prone to security bugs and require security review.

Entities And Concepts
- **Subresource request**: Destination in `["audio", "audioworklet", "font", "image", "json", "manifest", "paintworklet", "script", "style", "text", "track", "video", "xslt", ""]`.
- **Non-subresource request**: Destination in `["document", "embed", "frame", "iframe", "object", "report", "serviceworker", "sharedworker", "worker"]`.
- **Navigation request**: Subset of non-subresources with destination in `["document", "embed", "frame", "iframe", "object"]`.
- **Redirect-taint**: A classification ("same-origin", "same-site", "cross-site") derived from origin comparisons across a request’s URL list.
- **Cross-Origin-Embedder-Policy**: A policy container value (`"credentialless"`) checked to determine if credentials are allowed for a given request.

Procedures And API Details
- `compute redirect-taint`: Iterates `request`’s URL list, tracking `lastURL`. Returns "cross-site" immediately if consecutive origins differ from the current origin in site terms; otherwise updates taint to "same-site" if origins differ but sites match, finally returning current taint.
- `serialize request origin`: Asserts origin is not `"client"`. Returns `"null"` if redirect-taint is not `"same-origin"`, otherwise returns serialized origin.
- `byte-serialize request origin`: Wraps the result of `serialize request origin` with isomorphic encoding.
- `clone request`: Creates a copy excluding body and WebDriver id; generates a new random UUID for the WebDriver id; clones the body if non-null.
- `add range header`: Asserts valid bounds, constructs `bytes=first-last` (or `bytes=first-` if last omitted), serializes components isomorphically, and appends to request headers.
- `check COEP allows credentials`: Asserts origin not `"client"`. Returns `true` unless mode is `"no-cors"` AND client exists AND policy container embedder policy value is `"credentialless"` AND redirect-taint is `"same-origin"`.

Nuance Or Contradictions
- The definition of "Range header denotes an inclusive byte range" includes a clarifying example: `first=0`, `last=500` represents 501 bytes, which aligns with standard HTTP range semantics (inclusive endpoints).
- The warning about features combining multiple responses into one logical resource explicitly flags them as historical sources of security bugs.

Candidate Wiki Hints
- Page: **Fetch Requests Classification** – Define subresource, non-subresource, and navigation request destinations.
- Page: **Request Redirect-Taint Computation** – Detail the algorithm for determining origin taint across URL lists.
- Page: **Origin Serialization Security** – Explain why `null` is returned when redirect-taint is not "same-origin" to prevent redirect leakage.
- Page: **Request Cloning Mechanics** – Document the exclusion of body and WebDriver id, and UUID generation.
- Page: **Range Header Construction** – Steps for adding Range headers and byte-range semantics.
- Page: **Cross-Origin-Embedder-Policy Credential Check** – Logic flow for determining credential allowance under COEP.

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
The chunk covers sections 2.2.6 (Responses), 2.2.7 (Miscellaneous), 2.3 (Authentication entries), 2.4 (Fetch groups), and 2.5 (Resolving domains). It details the structure of a `Response`, its lifecycle states, filtering mechanisms, cloning procedures, caching classifications, authentication entry structures, fetch group state management, and origin resolution logic.

Local Summary
A `Response` evolves over time with fields becoming available sequentially. It possesses an associated type ("basic", "cors", "default", "error", "opaque", "opaqueredirect"), a status code (default 200), headers, body, cache state, and various flags like `aborted`, `timing allow passed`, and `range-requested`. Filtered responses provide limited views to prevent information leakage. The chunk also defines `fetch` destinations, authentication entry tuples, fetch group records for managing concurrent requests, and the `resolve an origin` algorithm which maps origins to IP addresses via DNS or implementation-defined operations.

Key Claims
- A response's type defaults to "default" unless specified otherwise.
- Responses over HTTP/2 always have an empty status message.
- Filtered responses expose a limited view; their internal response is accessible only for legacy reasons like feeding image data to decoders.
- New specifications should not build on opaque filtered responses or opaque-redirect filtered responses due to architectural limitations.
- The `location` URL algorithm returns null if the status is not a redirect status.
- A fetch group holds lists of fetch records and deferred fetch records to manage state during document unloading or inactivity.

Entities And Concepts
- Response (type, status, headers, body, cache state)
- Filtered response (basic, CORS, opaque, opaque-redirect)
- Network error
- Aborted network error
- Fetch destination ("fetch" or string)
- Authentication entry (username, password, realm)
- Proxy-authentication entry
- Fetch group (fetch records, deferred fetch records)
- Fetch record (request, controller)
- Deferred fetch record (request, notify invoked, invoke state)
- Resolve an origin (network partition key, origin, IP addresses)

Procedures And API Details
- Cloning a response:
  1. If filtered, return a new identical filtered response with a cloned internal response.
  2. Copy the response except for its body.
  3. If the original body is non-null, clone it and set it on the new response.
  4. Return the new response.
- Translating a potential destination:
  1. If "fetch", return the empty string.
  2. Assert it is a destination.
  3. Return the destination.
- Resolving an origin:
  1. If host is an IP, return the host as a set.
  2. If host is "localhost" or "localhost.", return « ::1, 127.0.0.1 ».
  3. Perform implementation-defined operation (e.g., DNS query) to get a set of IP addresses.
  4. Return failure if unsuccessful.

Nuance Or Contradictions
- The `Location` header parsing requires an absolute URL with fragment if the response was constructed via the Response constructor (null URL).
- Opaque and opaque-redirect filtered responses are nearly indistinguishable from network errors, making properties like `response.ok` useless for them.
- Origin resolution order of IP addresses can differ between invocations.

Candidate Wiki Hints
- Fetch API: Response Object Structure and Types
- Security: Filtered Responses and Information Leakage Prevention
- Networking: Origin Resolution and DNS Caching in Fetch

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
This chunk covers the Fetch Standard's sections on connection management (2.6), network partition keys (2.7), HTTP cache partitions (2.8), port blocking rules (2.9), MIME type restrictions for scripts (2.10), and cookie infrastructure including `Cookie`/`Set-Cookie` headers (3.1).

Local Summary
The specification defines how a user agent manages a pool of connections identified by network partition keys, origins, and credentials. It details the process of obtaining or creating connections, handling proxies, negotiating ALPN protocols, and recording timing information to avoid exposing reused connection details. The chunk also outlines logic for determining cache partitions based on these keys, blocking requests from bad ports (e.g., telnet, ftp), restricting certain MIME types for script destinations, and managing cookies via secure headers and infrastructure algorithms.

Key Claims
- A user agent maintains an ordered set of connections identified by a network partition key, origin, and credentials boolean.
- Connection timing info includes timestamps for domain lookup, connection start/end, and secure connection start, along with the ALPN negotiated protocol.
- The "clamp and coarsen" algorithm ensures reused connection details are not exposed by resetting times to a default or coarse value.
- Connections can be created with specific settings regarding unreliable transport (e.g., HTTP/3) and custom certificate verification for WebTransport.
- Network partition keys consist of a site and an optional implementation-defined second key, used to determine HTTP cache partitions.
- Fetching is blocked if the URL scheme is HTTP(S) and the port is listed as a "bad port" (e.g., 23 telnet).
- Responses with MIME types like `audio/`, `image/`, `video/`, or `text/csv` are blocked for script-like destinations.
- Cookie infrastructure includes logic to determine same-site modes (`lax-or-less`, `strict-or-less`) and serialize cookies for the `Cookie` header.

Entities And Concepts
- **Connection Pool**: An ordered set of connections maintained by a user agent.
- **Network Partition Key**: A tuple (site, secondKey) used to identify connections and cache partitions.
- **Connection Timing Info**: A struct tracking timing metrics like domain lookup duration and ALPN protocol.
- **Bad Port**: Ports listed in a table (e.g., 21 ftp, 23 telnet) that trigger request blocking.
- **Same-Site Mode**: Values such as `lax-or-less` or `strict-or-less` determining cookie behavior across origins.
- **WebTransport**: A feature requiring specific ALPN settings (`SETTINGS_ENABLE_WEBTRANSPORT`) on HTTP/3 connections.

Procedures And API Details
- **Obtain a Connection**: Involves checking the pool, finding proxies (defaulting to "DIRECT"), resolving hosts, and racing connection attempts.
- **Create a Connection**: Sets timing info, establishes transport (handling TLS certificates and ALPN), and returns the connection or failure.
- **Record Connection Timing Info**: Defines when `connection end time` and `secure connection start time` are recorded relative to handshake completion and early data usage.
- **Append Cookie Header**: Retrieves cookies based on security context, host/path, and same-site mode, then serializes them into the header value.
- **Parse Set-Cookie Headers**: Iterates through response headers, parsing each `Set-Cookie`, storing it, and performing garbage collection for the host.

Nuance Or Contradictions
- The specification notes that connection management details are intentionally vague to allow implementer discretion, particularly regarding proxy auto-config (PAC) and IP address racing.
- Timing info clamping logic differs based on whether a cross-origin isolated capability is present, affecting how times are coarsened.
- ALPN Protocol ID identification must account for tunnelled protocols when a proxy is configured versus the first hop to the proxy.

Candidate Wiki Hints
- Connection Pooling and Partition Keys
- HTTP/3 and WebTransport Setup
- Port Blocking List Reference
- Same-Site Cookie Logic

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context

This chunk covers sections **3.2** and **3.4**, focusing on the `Origin` request header, its serialization grammar, and its usage rules within the Fetch standard. It also details **3.3**, which defines the CORS protocol, including preflight requests, response headers, credential handling, and specific exceptions for certain `Content-Type` values. Section **3.4** briefly outlines the processing model for the `Content-Length` header.

## Local Summary

The `Origin` header indicates the source of a fetch but omits the path, differing from the legacy `Referer`. It is sent based on response tainting ("cors"), request method (non-GET/HEAD), and referrer policy. The chunk provides strict ABNF for serializing origins, enforcing lowercase ASCII schemes/domains and specific IPv6 formatting rules.

The CORS protocol allows cross-origin resource sharing via opt-in headers (`Access-Control-Allow-Origin`, `Allow-Credentials`, etc.). It distinguishes between preflight requests (using `OPTIONS`) and actual requests. The section details how credentials impact header values (e.g., `*` is forbidden for `Allow-Origin` when credentials are included) and lists specific `Content-Type` exceptions for non-preflighted cross-origin requests.

## Key Claims

- The `Origin` header is a path-less version of `Referer`, used for CORS-tainted responses and non-GET/HEAD methods.
- Origin serialization is stricter than RFC 3986: schemes/domains are lowercase ASCII, IPv6 addresses cannot be elided to single zero blocks, and leading zeros are forbidden.
- A CORS request includes an `Origin` header; a CORS-preflight request specifically uses the `OPTIONS` method with `Access-Control-Request-Method` and `Access-Control-Request-Headers`.
- Successful responses to CORS requests can use any status code if they include appropriate headers, while preflight responses are restricted to "ok" statuses (e.g., 200).
- When credentials mode is "include", `Access-Control-Allow-Origin` cannot be `*`, and `Access-Control-Allow-Credentials` must be present with the value `true`.
- The protocol includes exceptions for specific `Content-Type` headers (`application/csp-report`, etc.) to allow non-preflighted cross-origin requests.

## Entities And Concepts

- **Origin Header**: Indicates fetch origin; does not reveal path.
- **CORS Protocol**: Mechanism for cross-origin resource sharing.
- **Preflight Request**: `OPTIONS` request checking CORS support.
- **Credentials Mode**: Controls whether cookies/auth headers are sent ("omit", "same-origin", "include").
- **Access-Control-Allow-Origin**: Header declaring allowed origins (`*` or specific origin).
- **Access-Control-Allow-Credentials**: Header permitting shared responses with credentials (`true`).
- **Access-Control-Expose-Headers**: Header listing response headers exposed to client JavaScript.
- **Referrer Policy**: Determines when `Origin` is sent based on privacy settings (e.g., "no-referrer").
- **Content-Type Exceptions**: Specific MIME types allowed without preflight for security reasons.

## Procedures And API Details

**Appending the `Origin` Header:**
1. Assert request origin is not "client".
2. Byte-serialize the request origin to get `serializedOrigin`.
3. If response tainting is "cors" or mode is "websocket"/"webtransport", append (`Origin`, `serializedOrigin`).
4. Otherwise, if method is neither `GET` nor `HEAD`:
   - Check referrer policy:
     - `"no-referrer"`: Set origin to `null`.
     - `"no-referrer-when-downgrade"`, `"strict-origin"`, `"strict-origin-when-cross-origin"`: If origin scheme is "https" and current URL scheme is not, set origin to `null`.
     - `"same-origin"`: If origins differ, set origin to `null`.
   - Append (`Origin`, `serializedOrigin`) if applicable.

**CORS Header Values (ABNF):**
- `Access-Control-Allow-Origin`: `origin-or-null` or `*`.
- `Access-Control-Allow-Credentials`: `%s"true"` (case-sensitive).
- `Access-Control-Expose-Headers`: List of field names or `*` (only without credentials).
- `Access-Control-Max-Age`: Delta seconds.

## Nuance Or Contradictions

- The `Origin` header is technically a variant of `Referer` but omits the path to prevent leaking specific resource locations.
- A response can technically be "successful" with status 403 if it includes the necessary CORS headers, distinguishing between HTTP success and CORS permission logic.
- `Access-Control-Allow-Origin` values are case-sensitive; `"True"` is invalid even though `"true"` is valid.
- The `Allow` header is explicitly irrelevant for CORS protocol compliance.

## Candidate Wiki Hints

- **Origin Header**: Explain the difference from `Referer`, serialization rules, and when it is sent.
- **CORS Protocol**: Overview of headers (`Allow-Origin`, `Allow-Credentials`, `Expose-Headers`) and preflight logic.
- **Credentials in CORS**: Detail the restrictions on using `*` for origins and the requirement for `Allow-Credentials: true`.
- **Content-Type Exceptions**: List the specific MIME types exempt from preflight requirements.

## chunk-09

---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context

This chunk covers sections 3.5 through 4 of the Fetch Standard, detailing HTTP header processing logic and the core fetching algorithm. It includes specifications for `Content-Type`, `X-Content-Type-Options`, `Cross-Origin-Resource-Policy`, and `Sec-Purpose` headers, followed by the main steps to initiate a fetch operation.

## Local Summary

The section defines algorithms for extracting MIME types from header lists, handling charset parameters, and legacy encoding extraction. It specifies logic for determining if a response should be blocked based on `nosniff` directives relative to script or style destinations. The `Cross-Origin-Resource-Policy` (CORP) header is detailed with its ABNF, internal check mechanisms, and violation reporting workflows. Finally, the main "To fetch" algorithm outlines request population, early hint processing, preloaded resource handling, header insertion (Accept, Accept-Language), priority setting, and the invocation of the main fetch controller.

## Key Claims

- The `Content-Type` header extraction model differs from standard HTTP models when applied to web content; it prioritizes safety over strict compatibility with legacy features that have caused security vulnerabilities.
- MIME type parameters can typically be safely ignored during extraction unless they define the essence or charset, which are critical for format validation.
- The `X-Content-Type-Options: nosniff` directive triggers a blocking mechanism only for "script-like" and "style" destinations; image destinations are explicitly noted as incompatible with this exploit pattern in current deployments.
- The `Cross-Origin-Resource-Policy` header allows origins to enforce restrictions on resources loaded via "no-cors" requests, supporting values like `same-origin`, `same-site`, and `cross-origin`.
- Fetching algorithms support suspension and resumption of ongoing operations, with specific constraints regarding HTTP cache updates for "no-store" responses.

## Entities And Concepts

- **MIME Type Extraction**: Algorithm to parse `Content-Type` headers, handling multiple values, essence matching, and charset parameters.
- **Legacy Encoding Extraction**: A fallback method to retrieve encoding from a MIME type, returning a fallback encoding if the type is failure or lacks a charset parameter.
- **Nosniff Check**: Logic to determine if a response's `Content-Type` matches the request destination (script/style) when `nosniff` is present.
- **Cross-Origin Resource Policy (CORP)**: Mechanism to check request origin against response URL origin for "no-cors" requests, supporting violation reporting.
- **Fetch Params**: An object encapsulating request details, timing info, callbacks, and controller references during the fetch lifecycle.
- **Preloaded Response Candidate**: A state indicating whether a preloaded resource is available or pending for a specific request.
- **Accept Header Insertion**: Logic to populate `Accept` headers based on destination type (document, image, json, style, text) if missing.

## Procedures And API Details

### Extract MIME Type
1. Initialize `charset`, `essence`, and `mimeType` to null.
2. Decode and split the `Content-Type` header from the list.
3. Iterate through values, parsing each; skip failures or "*/*" essences.
4. Set `mimeType` on the first valid parse.
5. If subsequent parses have a different essence, reset `charset` to null and update `essence`.
6. If `mimeType` parameters lack "charset" but `charset` is non-null, assign it to the parameters.
7. Return failure if no valid `mimeType` was found; otherwise, return the final `mimeType`.

### Determine Nosniff
1. Decode and split `X-Content-Type-Options`.
2. If null, return false.
3. If the first value matches "nosniff" (case-insensitive), return true.
4. Return false.

### Cross-Origin Resource Policy Internal Check
1. If `forNavigation` is true and embedder policy is "unsafe-none", return allowed.
2. Retrieve `Cross-Origin-Resource-Policy` from the response header list.
3. Normalize policy: if not one of `same-origin`, `same-site`, or `cross-origin`, treat as null.
4. Switch on embedder policy value:
   - "unsafe-none": Do nothing (policy remains null).
   - "credentialless": Set policy to `same-origin` if request includes credentials or is navigation.
   - "require-corp": Set policy to `same-origin`.
5. Evaluate based on normalized policy:
   - `null`: Return allowed.
   - `cross-origin`: Return allowed.
   - `same-origin`: Return allowed only if origins match; otherwise blocked.
   - `same-site`: Return allowed only if schemes match (HTTPS/HTTP) and sites are schemelessly same; otherwise blocked.

### To Fetch
1. Assert request mode is "navigate" or early hints handler is null.
2. Initialize task destination and cross-origin isolated capability.
3. Populate request from client details.
4. If client exists, set task destination to client's global object and capture cross-origin isolated capability.
5. If parallel queue is used, start a new one for the task destination.
6. Create fetch timing info with coarsened current time.
7. Initialize fetch params with request, callbacks, and controller references.
8. Convert byte sequence body to a body object if present.
9. Run WebDriver BiDi clone network request body steps.
10. Check for preloaded resources if URL is HTTP(S), mode is safe, client is Window, method is GET, and unsafe flag is unset:
    - Define `onPreloadedResponseAvailable` callback.
    - Invoke consume a preloaded resource algorithm.
    - Update fetch params preloaded response candidate state if found.
11. If `Accept` header is missing:
    - Default to `*/*`.
    - Override for "prefetch" initiator with document value.
    - Otherwise, set based on destination (document, image, json, style, text).
    - Append (`Accept`, value) to header list.
12. If `Accept-Language` is missing and client exists:
    - Use emulated language from WebDriver BiDi if available.
    - Encode and append (`Accept-Language`, encodedEmulatedLanguage).
13. If `Accept-Language` still missing, suggest appending an appropriate value.
14. Set request's internal priority using implementation-defined logic based on priority, initiator, destination, and render-blocking.
15. For subresource requests, create a fetch record and append to client's fetch group.
16. Run main fetch given fetch params.
17. Return fetchParams’s controller.

## Nuance Or Contradictions

- **MIME Type Safety**: The standard explicitly notes that existing web platform features have not always followed the strict MIME type extraction pattern, leading to security vulnerabilities. This suggests a divergence from historical behavior toward a more secure, albeit stricter, model.
- **Image Destination and Nosniff**: The text notes that considering "image" as a target for `nosniff` blocking was not compatible with deployed content, implying that current implementations likely ignore this check for images despite the general algorithmic possibility.
- **Policy Header Matching**: Multiple `Cross-Origin-Resource-Policy` headers or values like `same-site, same-origin` are noted to have specific matching behaviors (e.g., `same-site` ignoring secure transport mismatches) which can result in unexpected "allowed" states if not carefully parsed.

## Candidate Wiki Hints

- **MIME Type Extraction Algorithm**: A reusable procedure for parsing and validating MIME types with charset handling, useful for browser engine documentation.
- **X-Content-Type-Options Security**: Notes on the `nosniff` header's scope (script/style only) and its role in preventing content-type sniffing attacks.
- **Cross-Origin Resource Policy (CORP)**: Detailed breakdown of CORP header values, internal check logic, and interaction with Embedder Policies.
- **Fetch Algorithm Initialization**: The "To fetch" algorithm steps serve as a foundational reference for implementing network request lifecycles in web engines.

## chunk-10

---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context
This chunk details the **Main fetch** algorithm (steps 4.1) and introduces the **Override fetch** concept (step 4.2). It covers URL security upgrades, referrer logic, scheme switching to HTTPS via HSTS/SVCB, response tainting (basic, cors, opaque), filtering responses for CORS headers, integrity checks (SRI), timing information collection, and stream transformation for body handling. The chunk concludes by defining the `override fetch` mechanism which allows user agents to intercept requests and return synthetic responses or errors before the standard fetch proceeds.

## Local Summary
The fetch algorithm validates security constraints (local-only URLs, mixed content, CSP) and handles referrer logic. It upgrades HTTP requests to HTTPS for known HSTS hosts. The algorithm branches based on request mode (navigate, no-cors, same-origin) and scheme (data, http). For CORS requests, it manages preflight responses and exposes headers. Finally, it constructs a filtered response, checks integrity metadata, handles timing info, and pipes the body stream to notify completion. `override fetch` allows user agents to return synthetic responses for specific hosts or paths.

## Key Claims
- **Security Upgrades**: Requests are blocked if local-only flags conflict with non-local URLs, mixed content is unsafe, or CSP/Integrity policies block them. HTTP requests are automatically upgraded to HTTPS if the host matches HSTS rules (superdomain or congruent) and DNS resolves an HTTPS RR.
- **Response Tainting**: Responses are categorized into "basic", "cors", or "opaque". This determines how headers are exposed and whether the response body is accessible.
- **CORS Handling**: If `Access-Control-Expose-Headers` contains `*` and credentials mode is not "include", all headers are exposed. Preflight responses clear cache entries if they fail.
- **Integrity Checks**: If integrity metadata is set, the body must match; otherwise, a network error is returned.
- **Timing Info**: Server-timing headers are decoded from `Server-Timing` and reported via resource timing APIs for secure contexts.
- **Stream Transformation**: The response body stream is piped through an identity transform stream to trigger callbacks when the body ends or fails.

## Entities And Concepts
- **Fetch Params**: Object containing request, recursive flag, timing info, controller, etc.
- **Request Object**: Contains current URL, referrer, mode, credentials mode, headers, flags (local-URLs-only, mixed content).
- **Response Object**: Internal response, status, body, header list, tainting, cache state.
- **Filtered Response**: A wrapper exposing only allowed headers based on tainting ("basic", "cors", "opaque").
- **Override Fetch**: An algorithm entry point allowing interception to return a direct response or null.
- **Scheme Fetch / HTTP Fetch**: Sub-algorithms invoked by override fetch for specific schemes.
- **HSTS / SVCB**: Protocols used to determine if an HTTPS upgrade is valid based on DNS records and domain matching.
- **SRI (Subresource Integrity)**: Mechanism to verify response body integrity against metadata.

## Procedures And API Details
- **Main Fetch Steps**:
  1. Validate local URL constraints.
  2. Check CSP violations.
  3. Upgrade URLs if HSTS/SVCB conditions met.
  4. Apply referrer policy and determine referrer value.
  5. Switch scheme to "https" for eligible HTTP hosts.
  6. Branch based on `recursive` flag or initial response candidate (preload).
  7. Handle specific modes: navigate, websocket, webtransport, same-origin, no-cors.
  8. For CORS: manage preflight, set tainting to "cors", run HTTP fetch.
  9. Construct filtered response based on tainting.
 10. Set redirect taint and timing allow passed flag.
 11. Block if mixed content, CSP, MIME type, or nosniff policies fail.
 12. Handle range requests and null body statuses.
 13. Validate integrity metadata; abort if mismatch.
 14. Run fetch response handover: decode `Server-Timing`, update controller timing info.
 15. Queue tasks for end-of-body processing and resource timing reporting.
 16. Pipe body stream through a TransformStream to trigger flush callbacks.
- **Override Fetch**:
  - Input: type ("scheme-fetch" or "http-fetch"), fetchParams, makeCORSPreflight.
  - Logic: Execute potentially override response. If non-null, return immediately.
  - Implementation: Default returns null; user agents may inject synthetic responses (e.g., shims for unsafe.example/widget.js).

## Nuance Or Contradictions
- **DNS Timing**: DNS operations are implementation-defined and might occur earlier or later than traditionally expected to support HSTS upgrades, requiring potential logic unwinding.
- **Range Requests**: APIs traditionally accept ranged responses even if the range wasn't requested, preventing partial data leaks in service workers.
- **Header Exposure**: One of the header names in `Access-Control-Expose-Headers` can be `*`, which matches only headers named `*`, a specific edge case in CORS logic.
- **User Agent Intervention**: The `potentially override response` step is implementation-defined, meaning user agents can silently block or synthesize responses without following the standard fetch flow for certain domains.

## Candidate Wiki Hints
- **Fetch Algorithm Overview**: Explain the high-level steps of fetching resources, including security checks and mode handling.
- **CORS Filtering**: Detail how response headers are filtered based on tainting and `Access-Control-Expose-Headers`.
- **HSTS Upgrades**: Describe the conditions under which HTTP URLs are upgraded to HTTPS via HSTS or SVCB records.
- **Override Fetch Pattern**: Discuss the architecture allowing user agents to intercept fetch requests and return synthetic responses (shimming).
- **Response Tainting**: Define "basic", "cors", and "opaque" states and their impact on browser behavior.

## chunk-11

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

## chunk-12

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

## chunk-13

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

## chunk-14

---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk covers the **Deferred Fetching** algorithm within the Fetch specification, detailing how requests are queued to execute at a later time (e.g., when a fetch group terminates or after a timeout). It also includes the **Fetch API** overview.

The content spans from section 4.12 ("Deferred fetching") through 4.12.1 ("Deferred fetching quota"), providing detailed algorithms for queuing, processing, and managing quotas for deferred requests. Finally, it introduces section 5, offering a high-level overview of the `fetch()` method and practical usage examples.

# Local Summary

The specification defines a mechanism to defer network requests until necessary, prioritizing them in a specific task source to ensure scripts reflecting the latest state run before dependent scripts. A strict quota system limits memory usage:
- **Top-level quota**: 640 kibibytes.
- **Same-origin nested documents**: Share the parent's quota.
- **Cross-origin nested documents**: Receive a default allocation of 8 kibibytes, configurable via `Permissions-Policy`.
- **Per-origin concurrent limit**: Only 64 kibibytes can be used concurrently for the same reporting origin to prevent third-party libraries from hoarding memory.

The chunk concludes with an overview of the `fetch()` API, highlighting its capabilities for low-level resource fetching, handling responses as Blobs or JSON, working with URL parameters, and progressive consumption of request bodies.

# Key Claims

- **Deferred Fetching Task Source**: User agents must prioritize tasks in the deferred fetch task source before other sources (like DOM manipulation) to ensure the most recent state of a `fetchLater()` call is reflected before running dependent scripts.
- **Quota Allocation**: The default top-level quota is 640 kibibytes. By default, 128 kibibytes are reserved for delegating to cross-origin nested documents (`deferred-fetch-minimal` policy).
- **Concurrent Usage Limit**: Out of the allocated quota, only 64 kibibytes can be used concurrently for a specific reporting origin (the request's URL's origin).
- **Quota Calculation**: Total request length includes the URL (without fragment), referrer, header list lengths, and body length.
- **Policy Control**: The feature is identified by `"deferred-fetch"` (default allowlist: `"self"`) and `"deferred-fetch-minimal"` (default allowlist: `"*"`).

# Entities And Concepts

- **Deferred Fetch Record**: A record created when `fetchLater()` is called, containing the request and a notification function.
- **Fetch Group**: The container managing deferred fetch records for a client.
- **Task Source**: Specifically the "deferred fetch task source," prioritized for executing deferred requests.
- **Permissions Policy**: Used to control how much quota is delegated to cross-origin nested documents (e.g., `Permissions-Policy: deferred-fetch=(self "https://fratop.example.com")`).
- **Reporting Origin**: The origin derived from the request's URL, which has its own concurrent usage limit of 64 kibibytes.
- **Top-level Traversable**: A tab or window acting as the root for quota allocation (default 640 kibibytes).

# Procedures And API Details

### Queuing a Deferred Fetch
To queue a deferred fetch given a `request`, `activateAfter` (optional timestamp), and `onActivatedWithoutTermination`:
1. Populate request from client.
2. Set service-workers mode to "none".
3. Set keepalive to true.
4. Create a new deferred fetch record.
5. Append the record to the fetch group's deferred fetch records.
6. If `activateAfter` is provided, wait for the timeout or a reason to believe scripts are about to be lost (e.g., backgrounding) before processing.

### Computing Total Request Length
Used to check against quota limits:
1. Length of request’s URL (exclude fragment).
2. Plus length of request’s referrer.
3. Plus sum of name and value lengths for all headers.
4. Plus length of request’s body.

### Processing Deferred Fetches
When a fetch group is terminated or timeout occurs:
1. Iterate through deferred fetch records in the group.
2. For each record with "pending" invoke state:
   - Set state to "sent".
   - Fetch the request.
   - Queue a global task on the deferred fetch task source to run the notification function.

### Managing Quota (Algorithm Overview)
To get available quota for a document and origin:
1. Determine if it is a top-level traversable and if policies allow usage.
2. Calculate base quota (e.g., 640, 512, or 0 based on policy).
3. Subtract reserved quotas for cross-origin children and pending requests.
4. Return the remaining quota or 0 if exhausted.

# Nuance Or Contradictions

- **Quota Reset on Navigation**: If a navigable container navigates to a same-origin document, any reserved quota for that container is freed. Conversely, navigating away from a cross-origin frame releases its specific reservation back into the pool.
- **Redirect Handling**: Quota calculations assume redirects are handled before origin checks. A document created via redirect knows its final origin only after redirects are resolved, affecting when quota can be reserved or checked.
- **Cross-Origin vs Same-Origin Frames**: Same-origin frames share the parent's quota pool. Cross-origin frames default to 8 kibibytes unless explicitly delegated more via `Permissions-Policy`. If a cross-origin frame navigates to a different origin, it loses any previously granted specific quota for that path.

# Candidate Wiki Hints

- **Topic: Deferred Fetching**
  - Explain the purpose of deferring requests (memory management, user interaction timing).
  - Detail the `fetchLater()` API and its parameters.

- **Topic: Fetch Quota Management**
  - Break down the quota hierarchy (Top-level -> Same-origin children -> Cross-origin children).
  - Explain the `Permissions-Policy` header usage for controlling cross-origin delegation.
  - Provide examples of request sizes that trigger quota exhaustion.

- **Topic: Fetch API Overview**
  - Summarize the capabilities of `fetch()` compared to `XMLHttpRequest`.
  - Include code snippets for Blob extraction, JSON handling, and progressive reading.

## chunk-15

---
title: Chunk 15 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context
This chunk covers sections 5.1 through 5.3 of the Fetch specification, detailing the `Headers` class, `BodyInit` unions, and the `Body` mixin interface. It also includes algorithmic steps for validating headers, handling CORS guards (specifically "request-no-cors"), extracting request/response bodies from various input types (Blob, BufferSource, FormData, etc.), and consuming body streams into different JavaScript value types.

## Local Summary
The text defines the `Headers` API with methods like `append`, `delete`, `get`, `set`, and `has`, emphasizing validation rules based on header guards ("immutable", "request", "response", "none", "request-no-cors"). It outlines specific logic for appending headers in CORS contexts, where privileged no-CORS request-headers are removed upon modification by unprivileged code. The chunk then transitions to body handling, defining `BodyInit` types and the extraction/consumption algorithms that convert byte sequences or streams into `ArrayBuffer`, `Blob`, `FormData`, `json()`, or `text`. Finally, it details the `Body` mixin methods (`arrayBuffer()`, `blob()`, `bytes()`, `formData()`, `json()`, `text()`) and their underlying consumption logic, including MIME type detection and multipart parsing rules.

## Key Claims
- A `Headers` object maintains an associated header list and a guard state ("immutable", "request", "response", "none", or "request-no-cors").
- Validation of headers throws a `TypeError` if the name/value is invalid or if the guard is "immutable".
- For "request-no-cors" guards, appending headers requires checking for no-CORS-safelisted status; privileged headers are removed if modified by unprivileged code.
- The `BodyInit` type union includes `Blob`, `BufferSource`, `FormData`, `URLSearchParams`, `USVString`, and `ReadableStream`.
- The `consume body` algorithm ensures a promise is returned, rejecting with `TypeError` if the object is unusable (body disturbed/locked).
- The `formData()` method handles "multipart/form-data" and "application/x-www-form-urlencoded" MIME types, parsing parts into entries based on `Content-Disposition` headers.

## Entities And Concepts
- **Headers Class**: Represents HTTP headers with associated guards for security and immutability.
- **HeadersGuard**: States including "immutable", "request", "response", "none", and "request-no-cors".
- **BodyInit**: A union type for request/response body sources (e.g., `Blob`, `FormData`, `ReadableStream`).
- **BodyMixin**: An interface mixin providing methods to read the body as different types (`arrayBuffer()`, `blob()`, `text()`, etc.).
- **CORS Guards**: Specific logic for "request-no-cors" headers involving safelisted checks and removal of privileged headers.
- **MIME Type Extraction**: Logic to determine content type from headers (e.g., `multipart/form-data`, `application/json`).

## Procedures And API Details
- **`Headers.append(name, value)`**: Normalizes the value, validates against guards, handles "request-no-cors" merging logic, and removes privileged headers if needed.
- **`Headers.delete(name)`**: Validates (with a dummy value), checks guard constraints, deletes from list, and cleans up privileged headers in "request-no-cors" mode.
- **`BodyInit` Extraction**: Converts inputs like `Blob` (via stream), `FormData` (multipart encoding), or `URLSearchParams` into a readable stream with an associated type.
- **`consume body` Algorithm**: Reads the body stream, resolves promises with converted data (e.g., `ArrayBuffer`, `Blob`), or rejects if the stream is disturbed/locked.
- **`formData()` Parsing**: Parses multipart parts using the boundary parameter; parts with `filename` become `File` objects, others become strings.

## Nuance Or Contradictions
- The specification notes that steps for "request-no-cors" are not shared in a simple way because a fake value cannot always succeed for CORS-safelisted headers.
- The `formData()` implementation is noted as a "rough approximation," with a more detailed parsing specification pending, particularly regarding the handling of `_charset_` parameters and default content types.
- The `bytes()` method can reject with a `RangeError`, while `arrayBuffer()` and `blob()` can also fail under specific conditions (e.g., if the stream is disturbed).

## Candidate Wiki Hints
- **Headers API**: Documenting the `Headers` interface, including initialization from objects or arrays, guard states, and CORS-specific restrictions.
- **Body Consumption**: Explaining how to convert request/response bodies into various JavaScript types using the `Body` mixin methods.
- **FormData Parsing**: Detailing the logic behind `formData()` for handling multipart and form-encoded data, including `File` object creation.

## chunk-16

---
title: Chunk 16 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
The chunk details the **Request class** within the Fetch API specification. It defines the `Request` interface, its associated attributes (such as `url`, `method`, `headers`, and various modes like `mode`, `credentials`, and `cache`), the `RequestInit` dictionary for initialization options, enumeration types for attribute values, and the internal steps for constructing a new `Request` object via the `new Request(input, init)` constructor. It also covers getters for attributes, specific handling of referrers, and the logic for cloning a request.

Local Summary
This section explains how to create a `Request` object, detailing the properties available in the `init` argument (like `method`, `headers`, `body`, `referrer`, etc.). It outlines the internal constructor steps, including URL parsing, mode validation, referrer processing, header sanitization, and body handling. Getters for standard attributes are described, along with the specific logic for the `clone()` method which creates a new request object linked to the original signal.

Key Claims
- The `Request` interface represents a resource that is being fetched.
- The `new Request(input, init)` constructor accepts a URL string or an existing `Request` object as input.
- If `input` is a string, it must be parsed; if parsing fails or includes credentials, a `TypeError` is thrown.
- The default mode for a new request from a string is `"cors"`.
- Specific referrer values like `"about:client"` and `"no-referrer"` are handled internally during construction.
- The `clone()` method returns a clone of the request but requires the signal to be non-null.
- Headers added in the network layer (e.g., `Host`) are not included in the returned `Headers` object.

Entities And Concepts
- **Request**: The main interface for an HTTP request.
- **RequestInfo**: A union type representing either a `Request` object or a URL string (`USVString`).
- **RequestInit**: A dictionary containing optional initialization parameters for the `Request` constructor.
- **Headers**: An object used to manage HTTP headers associated with the request.
- **AbortSignal**: Used to cancel an ongoing fetch operation.
- **Referrer Policy**: Controls how referrer information is sent with requests.
- **CORS**: Cross-Origin Resource Sharing mode, distinct from `"same-origin"` or `"navigate"`.
- **SRI (Subresource Integrity)**: Cryptographic hash metadata for verifying resource integrity.

Procedures And API Details
- **Creating a Request**:
  - Pass a URL string to `new Request(url)`; default method is `"GET"`, default mode is `"cors"`.
  - Pass an existing `Request` object to copy its properties, optionally overriding them via the second argument.
  - Set `init.method` to change the HTTP verb (e.g., `"POST"`).
  - Set `init.headers` to provide a `Headers` object, literal, or array of arrays.
  - Set `init.body` to provide the request body content.
- **Referrer Handling**:
  - If `init.referrer` is an empty string, it sets the internal referrer to `"no-referrer"`.
  - If `init.referrer` is `"about:client"`, it sets the internal referrer to `"about:client"`.
  - If the parsed referrer's origin is not same-origin or scheme is `"about"` with path `"client"`, it defaults to `"client"`.
- **Cloning**:
  - Call `request.clone()` to get a copy.
  - The clone shares the abort signal dependency.
  - Throws a `TypeError` if the request is unusable.

Nuance Or Contradictions
- **Service Worker Origins**: A request can have an origin different from the current client when handled by a service worker, specifically for navigation requests.
- **Window Disassociation**: The `window` property in `RequestInit` can only be set to `null` to disassociate the request from any Window; setting it to a non-null value throws a `TypeError`.
- **Duplex Mode**: Currently, only `"half"` is valid for `duplex`. `"full"` is reserved for future use where response processing occurs before the full request body is sent.
- **No-CORS Restrictions**: If mode is `"no-cors"`, the method must be a CORS-safelisted method; otherwise, a `TypeError` is thrown during construction.

Candidate Wiki Hints
- [Request API](https://wiki.local/Request_API) - Core interface and attributes.
- [Request Constructor Options](https://wiki.local/Request_Constructor_Options) - Detailed breakdown of `RequestInit`.
- [Cloning Requests](https://wiki.local/Cloning_Requests) - Usage of the `clone()` method.
- [Referrer Policies in Fetch](https://wiki.local/Referrer_Policies_in_Fetch) - Handling of referrer strings and policies.

## chunk-17

---
title: Chunk 17 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

**Chunk Heading:** parsedReferrer’s origin is not same origin with > 5.5. Response class
**Lines:** 5769–6167
**Coverage:**
- Section 5.5: Response class interface, attributes, and constructor steps.
- Section 5.6: Fetch methods including `fetch()`, `fetchLater()`, and related APIs.
- Section 5.7: Garbage collection rules regarding observable fetch termination.

# Local Summary

This chunk defines the Web API for constructing and manipulating HTTP responses via the `Response` interface, detailing its attributes like `status`, `headers`, and `type`. It specifies static factory methods (`error()`, `redirect()`, `json()`) and the internal steps to initialize a response object. The text then moves to the `fetch()` method implementation within `WindowOrWorkerGlobalScope`, explaining promise resolution, abort signal handling, and client-side request lifecycle management. A new API, `fetchLater()`, is introduced for deferring network requests until a future time or document activation state, complete with quota checks and abort mechanisms. Finally, the chunk addresses garbage collection semantics, distinguishing between fetches that can be terminated by the user agent (unobservable bodies) versus those protected by script-observable promises or streams.

# Key Claims

- The `Response` interface is exposed on both `Window` and `Worker` contexts.
- A `Response` object has an associated response body, headers (initially null), and a type enum (`basic`, `cors`, `default`, `error`, `opaque`, `opaqueredirect`).
- The constructor `new Response(body, init)` creates a new response with an optional body and initialization dictionary.
- Static methods include:
  - `Response.error()`: Creates a network error response.
  - `Response.redirect(url, status)`: Creates a redirect response (default status 302).
  - `Response.json(data, init)`: Encodes data as JSON and creates a response with `Content-Type` set to `application/json`.
- The `fetch()` method returns a promise that resolves to a `Response` object or rejects if aborted/network error occurs.
- `fetchLater()` allows deferring a fetch request until the document is fully active or a specified time (`activateAfter`) has passed.
- Garbage collection can terminate ongoing fetches if the result (response/promise) is not observable through script arguments or return values.

# Entities And Concepts

- **Response**: The core interface representing an HTTP response, including properties like `url`, `status`, `ok`, `redirected`, and `headers`.
- **ResponseType**: An enum defining how a response body is exposed (e.g., `cors` allows reading, `opaque` does not).
- **FetchLaterResult**: A result object returned by `fetchLater()` indicating whether the deferred request has been activated.
- **DeferredRequestInit**: A dictionary extending `RequestInit` with an optional `activateAfter` timestamp.
- **Garbage Collection Termination**: The mechanism where the user agent may cancel a fetch if it cannot be observed via script, preventing resource leaks in unmanaged contexts.

# Procedures And API Details

### Creating a Response Object
1.  **Constructor (`new Response(body, init)`)**:
    - Creates a new `Response` with an associated response.
    - Initializes headers from the provided list.
    - Calls `initialize a response` to set status, status text, and body.
2.  **Static Methods**:
    - `error()`: Returns a network error response (realm: "immutable").
    - `redirect(url, status)`: Parses URL, validates redirect status, sets Location header, returns response (realm: "immutable").
    - `json(data, init)`: Serializes data to JSON bytes, extracts body, creates response with "response" guard, and initializes with `Content-Type: application/json`.

### Initializing a Response
1.  Validate `init["status"]` is between 200 and 599; otherwise throw `RangeError`.
2.  Validate `init["statusText"]` matches reason-phrase token production or is empty; otherwise throw `TypeError`.
3.  Set response status and status message.
4.  If `init["headers"]` exists, fill the headers list.
5.  If body is non-null:
    - Throw `TypeError` if status indicates a null body (e.g., 101, 103).
    - Set response body to the provided body's body.
    - If body type is non-null and `Content-Type` header is missing, append it.

### Fetch Lifecycle
1.  **Abort Handling**: Checks `requestObject.signal`. If aborted, aborts the call immediately with the signal's reason.
2.  **Promise Resolution**: Upon receiving a response, creates a `Response` object (immutable guard) and resolves the promise.
3.  **Network Error**: Rejects the promise with a `TypeError` if the response is a network error.

### FetchLater Implementation
1.  Parse input into a `Request` object.
2.  Validate abort signal state; throw if aborted.
3.  Check for HTTP(S) scheme and potentially trustworthy URL.
4.  Ensure body length is known (streams cannot be deferred).
5.  Check deferred-fetch quota against requested size.
6.  Queue the request to activate after `activateAfter` ms or when document becomes fully active.
7.  Return a `FetchLaterResult` object with an `activated` getter.

### Garbage Collection Rules
- **Terminable**: Fetches where the body is not readable or the response is ignored (e.g., `.then(res => res.headers)`).
- **Protected**: Fetches stored in variables (`window.promise = fetch(...)`) or where the stream reader's `closed` event is observed via a handler.

# Nuance Or Contradictions

- **Null Body Status**: The specification explicitly notes that status codes 101 and 103 are included in "null body status" for validation purposes, even though they are typically used elsewhere (e.g., switching protocols).
- **Observable Termination**: There is a clear distinction between fetches that can be terminated by the user agent (unobservable) and those that cannot. The chunk provides examples showing that accessing `res.body.getReader().closed` prevents termination because it makes the stream state observable to script. Conversely, ignoring the response body or headers alone allows termination.
- **Realm Guarding**: Different static methods and constructors use different realm guards ("immutable" vs "response"), affecting whether the response can be modified by scripts.

# Candidate Wiki Hints

- **Response Interface Guide**: A comprehensive page detailing all `Response` attributes, the `ResponseType` enum, and step-by-step examples of using `error()`, `redirect()`, and `json()`.
- **Fetch API Lifecycle**: An article explaining the internal steps of `fetch()`, including abort signal propagation and promise resolution behavior.
- **Deferred Fetching (`fetchLater`)**: A dedicated page or note on the `fetchLater()` API, covering its usage for background tasks, quota management, and the `FetchLaterResult` object.
- **Garbage Collection in Fetch**: A technical note explaining how the browser decides when to terminate fetches based on observability, with code examples of safe vs. unsafe patterns.

## chunk-18

---
title: Chunk 18 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
This chunk covers section "6. data: URLs", detailing the processing algorithm for `data:` URL schemes (based on RFC 2397), followed by background reading on HTTP header layers, atomic redirect handling, CORS safety protocols, WebSocket connections, and practical guidance for setting up requests and invoking fetch operations including response processing strategies.

Local Summary
The section defines the internal structure of a `data:` URL as containing a MIME type and a body byte sequence. It outlines an algorithmic processor that parses the scheme, validates the MIME type (handling base64 encoding markers), decodes the body, and returns a structured object. The text further explains HTTP header layer divisions, warns against exposing redirects to prevent cross-site scripting leaks via secrets in redirect URLs, clarifies safe CORS setups using `Access-Control-Allow-Origin: *`, discusses the role of `Vary` headers for caching CORS responses, describes WebSocket connection establishment as a special fetch mode, and provides guidance on constructing requests (URL, method, body, client settings) and invoking fetch with various response processing callbacks.

Key Claims
- A `data:` URL struct consists of a MIME type and a body byte sequence.
- The `data:` URL processor asserts the scheme is "data", strips the prefix, parses the MIME type, decodes the body (percent-decoding or base64 if indicated), and returns failure if parsing fails.
- If the MIME type starts with ";", it defaults to "text/plain;charset=US-ASCII".
- Exposing redirects can leak secrets; a redirect from an authenticated URL to a new URL containing a secret should not be exposed via APIs.
- Using `Access-Control-Allow-Origin: *` is safe for resources protected by IP or firewall, as it shares the resource with tools like curl and wget but does not reveal auth/cookie info if the resource itself isn't accessible from random devices.
- If CORS requirements are complex (not just static origin or "*"), the `Vary: Origin` header must be used to prevent caching non-CORS responses for subsequent CORS requests.
- WebSocket connections are established via a special fetch with mode "websocket".
- Fetch callbacks can handle response processing upon completion, chunk-by-chunk streaming, or ignoring the response entirely (e.g., `navigator.sendBeacon`).

Entities And Concepts
- data: URL struct
- MIME type
- body (byte sequence)
- isomorphic decode
- forgiving-base64 decode
- Access-Control-Allow-Origin
- Vary header
- WebSocket object
- fetch controller
- processResponseConsumeBody
- processResponse
- processResponseEndOfBody
- Sec-Fetch-Dest
- Content Security Policy (CSP)
- Referrer Policy

Procedures And API Details
**Data: URL Processor Steps:**
1. Assert scheme is "data".
2. Serialize URL excluding fragment.
3. Remove leading "data:".
4. Collect MIME type characters until comma (U+002C).
5. Strip ASCII whitespace from MIME type.
6. If position reaches end, return failure.
7. Advance position by 1.
8. Remainder is encodedBody.
9. Percent-decode encodedBody to get body.
10. Check if mimeType ends with ";", optional spaces, and "base64" (case-insensitive).
    - If yes: isomorphic decode body, then forgiving-base64 decode. Return failure if base64 decode fails. Remove trailing ";", spaces, and "base64" from mimeType.
11. If mimeType starts with ";", prepend "text/plain".
12. Parse mimeType into mimeTypeRecord. Default to "text/plain;charset=US-ASCII" on failure.
13. Return new data: URL struct with parsed MIME type and decoded body.

**Request Setup Guidance:**
- Set request's URL and method (e.g., POST/PUT with byte sequence or ReadableStream body).
- Choose destination per table to affect Content Security Policy (`Sec-Fetch-Dest`).
- Set client to environment settings object (or null for background fetching).
- Set mode to "same-origin" if same-origin required, otherwise "cors" for web-exposed features.
- Set credentials mode to "include" for cross-origin requests requiring credentials.
- Pass initiator type for Resource Timing reporting.
- Set header list for custom headers (be aware of CORS-preflight implications).
- Set cache mode if overriding default caching.
- Set redirect mode to "error" if redirects should not be followed.

**Fetch Invocation:**
- Pass `processResponseConsumeBody` to read entire body upon completion (returns null, failure, or byte sequence).
- Pass `processResponse` to handle headers and stream body chunk-by-chunk.
- Pass `processResponseEndOfBody` to handle after full download.
- Use fetch controller to abort, report timing, or handle manual redirects.

Nuance Or Contradictions
- The text notes that fetch used to define obtaining/establishing WebSocket connections directly, but these are now defined in the WebSockets standard.
- Handling "no-cors" responses requires caution; while the body is passed as a byte sequence, it should not be directly exposed to scripts in the embedding document (security restriction).
- Caching behavior differs significantly based on whether `Vary: Origin` is used; without it, cached non-CORS responses lack `Access-Control-Allow-Origin`, breaking subsequent CORS requests.

Candidate Wiki Hints
- data-url-processing-algorithm
- http-header-layers-fetch
- cors-safety-and-vary-header
- websocket-fetch-mode
- fetch-response-callbacks

## chunk-19

---
title: Chunk 19 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
The chunk covers the "data: URLs" section header and includes extensive acknowledgments, copyright/license information (CC BY 4.0 for the standard, BSD 3-Clause for source code portions), a note on the Living Standard versus the Patent-Review version, and an index reference to defined terms.

Local Summary
This segment of the specification document lists contributors thanked for their work on the standard, identifies Anne van Kesteren as the author, states the licensing terms, distinguishes between the Living Standard and the patent-review draft, and introduces the terms section.

Key Claims
- The standard is authored by Anne van Kesteren (Apple).
- Copyright is held by WHATWG (Apple, Google, Mozilla, Microsoft).
- The work is licensed under a Creative Commons Attribution 4.0 International License.
- Portions incorporated into source code are licensed under the BSD 3-Clause License.
- This version is the "Living Standard"; a patent-review version exists as a separate draft.

Entities And Concepts
- WHATWG: Working Group holding copyright.
- Living Standard: The current version of the specification.
- Patent-Review Version: A separate draft for patent review purposes.
- Creative Commons Attribution 4.0 International License: Primary license for the standard text.
- BSD 3-Clause License: License for source code portions incorporating parts of the standard.

Procedures And API Details
None applicable in this chunk; it contains administrative and legal metadata rather than technical procedures or APIs.

Nuance Or Contradictions
The document clarifies a licensing split: the specification text uses CC BY 4.0, while any implementation source code incorporating portions of the spec uses the BSD 3-Clause License instead.

Candidate Wiki Hints
- **data-urls**: A page summarizing the data: URL standard, its authors, and legal context.
- **WHATWG Licensing**: A comparison page for CC BY 4.0 (spec) vs. BSD 3-Clause (source code).
- **Living Standard vs. Patent Review**: Explaining the difference between the live standard and the patent-review draft.

## chunk-20

---
title: Chunk 20 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context
This chunk covers section **6. data: URLs**, focusing on the Fetch API's handling of `data:` URL schemes. It defines specific structures and processing logic for these URLs within the fetch environment. The text lists numerous definitions, algorithms, and terminology related to the broader HTTP fetch specification (e.g., `fetch()`, `Request`, `Response`, CORS), but the primary focus here is the interaction between the fetch mechanism and local data resources.

## Local Summary
The section establishes a formal definition for `data: URL` structs within the Fetch environment. It distinguishes between standard network fetches and those involving `data:` URLs, which are processed differently (often without network traversal). The chunk references algorithms for processing these URLs, likely involving parsing the media type and data payload directly. It also touches upon related concepts like MIME types (`media-type`), encoding, and how these resources fit into the broader cache and fetch timing models.

## Key Claims
- **Data URL Structure**: A specific "data: URL struct" is defined in § 6 to handle `data:` URLs within the fetch API.
- **Processing Logic**: The document outlines a "data: URL processor" responsible for handling these resources, distinct from standard HTTP network requests.
- **MIME Type Handling**: There is a relationship between `data:` URLs and MIME types (`media-type`), suggesting that the data payload's type is extracted or validated during processing.
- **Fetch Integration**: While `data:` URLs are local resources, they integrate into the Fetch API's terminology (e.g., `fetch()` calls might resolve to them under specific conditions, though often treated as opaque or distinct from network responses).

## Entities And Concepts
- **data: URL struct**: A defined structure for representing data URLs in the Fetch environment.
- **data: URL processor**: The algorithmic component responsible for resolving and handling `data:` URLs.
- **media-type**: The MIME type associated with the data payload within a `data:` URL.
- **fetch()**: The primary function invoked to retrieve resources, which may resolve to a `data:` URL depending on implementation details or user agent behavior (though standard fetch typically targets network resources).
- **Response**: The output object, potentially wrapping the parsed content of a `data:` URL.

## Procedures And API Details
- **Define data: URL struct**: An algorithm step is referenced for creating this specific structure (§ 6).
- **Process data: URL**: A procedure exists to handle the resolution and parsing of `data:` URLs, distinct from network fetch algorithms.
- **Extract media-type**: Logic is implied or defined for extracting the MIME type from the `data:` URL scheme.

## Nuance Or Contradictions
The chunk lists standard Fetch API terms (like `fetch()`, `Request`, `Response`) alongside specific `data: URL` terminology. This suggests that while `data:` URLs are part of the same specification ecosystem, they may bypass standard network fetch algorithms or have unique handling paths (e.g., regarding caching or redirection) not detailed in this specific text segment. The distinction between "network" and "data" processing is implied but not fully elaborated in this snippet.

## Candidate Wiki Hints
- **Page: Data URLs in Fetch** – A dedicated page explaining how the Fetch API handles `data:` URLs, including parsing rules and MIME type extraction.
- **Concept: Local Resource Fetching** – Notes on how non-network resources (like `data:` URLs) fit into the broader fetch model without traversing the network.

## chunk-21

---
title: Chunk 21 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- Heading path: 6. data: URLs
- Line range: 6581–7175
- Scope: Covers the "data:" URL scheme definition and usage within the Fetch Standard, including associated attributes, headers, and security flags referenced across multiple sections (§ 2.x, § 3.x, § 4.x, § 5.x).

Local Summary
This chunk details the `data:` URL scheme's integration into the Fetch specification. It defines how such URLs are handled as request bodies or response content, referencing specific attributes like `url` and `type` on `Request` and `Response` objects. The text also links to related concepts such as CORS preflight checks (`use-CORS-preflight flag`), unsafe request handling (`unsafe-request flag`), and credential usage (`use-URL-credentials flag`). It notes associations with other schemes (e.g., `video`, `text`, `xslt`) and specific headers like `X-Content-Type-Options`.

Key Claims
- The `data:` URL scheme is a defined resource type within the Fetch standard.
- Attributes such as `url` and `type` exist on both `Request` (§ 5.4) and `Response` (§ 5.5) objects to describe data sources.
- Security flags like `use-CORS-preflight`, `unsafe-request`, and `use-URL-credentials` are relevant contexts for handling these requests.
- The standard references content types including `"text"`, `"video"`, and `"xslt"` in the context of data URLs (§ 5.4).

Entities And Concepts
- **data: URL**: A URI scheme that carries data within the URL itself, often used for embedding resources like images or text directly in HTML/JS.
- **Request / Response Attributes**: Specific properties (`url`, `type`) attached to Fetch API objects to identify source content.
- **Security Flags**: Mechanisms like `use-CORS-preflight` and `unsafe-request` that dictate how data URLs interact with cross-origin policies or caching.
- **Content Types**: MIME types associated with embedded data, explicitly mentioning `"text"`, `"video"`, and `"xslt"`.

Procedures And API Details
- **Attribute Access**: The `url` attribute is defined for both `Request` (§ 5.4) and `Response` (§ 5.5).
- **Type Definition**: The `type` attribute exists on the `Response` object (§ 5.5).
- **Flag Usage**: Flags such as `use-CORS-preflight` (§ 2.2.5) and `unsafe-request` (§ 2.2.5) are referenced in the context of processing these URLs.

Nuance Or Contradictions
- The chunk lists various MIME types (`text`, `video`, `xslt`) alongside the generic `data:` scheme, implying that the standard treats data URLs flexibly regarding content type identification, though specific handling rules for each type may be defined elsewhere in the full document.

Candidate Wiki Hints
- **Page: Data URL Scheme** – A dedicated page explaining the syntax and security implications of `data:` URLs in web development.
- **Concept: Fetch API Attributes** – Documentation covering the `url` and `type` attributes on Request/Response objects.

## chunk-22

---
title: Chunk 22 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
Heading path: 6. data: URLs
Line range: 7177-7884
This chunk defines a comprehensive glossary of terms used throughout the Fetch Standard, referencing external specifications such as HTML, DOM, ECMAScript, and various security policies. It also includes normative and non-normative references to RFCs and W3C standards, along with an IDL index for key interfaces like `Headers`, `Body`, and `Request`.

Local Summary
This section enumerates terms defined by multiple specifications (e.g., HTML, DOM, Fetch Metadata) that are relevant to the Fetch Standard. It also lists numerous normative references (RFCs, W3C specs) and non-normative resources. The chunk concludes with an IDL index defining interfaces like `Headers`, `Body`, and `Request`.

Key Claims
- The chunk does not make explicit claims but rather serves as a reference list for terms and specifications used in the Fetch Standard.
- Terms are defined by referencing external specifications (e.g., `[HTML]` defines "active document", `[DOM]` defines "Document").
- Normative references include RFCs (e.g., [RFC3986], [RFC9110]) and W3C standards (e.g., [HTML], [DOM]).
- Non-normative references cover security vulnerabilities (e.g., HTTP TRACE method) and older timing APIs (e.g., Navigation Timing).

Entities And Concepts
- **Terms Defined**: Includes concepts like "active document", "Document", "Blob", "File", "URL", "Request", "Response", "Headers", "Body", "AbortSignal", "Promise", "ReadableStream", "WritableStream", etc.
- **Specifications Referenced**: HTML, DOM, ECMAScript, Fetch Metadata, File API, High Resolution Time, HTTP, HTTP Caching, HTTP/1.1, HTTP/3, Infra, MIME Sniffing, Mixed Content, Permissions Policy, Referrer Policy, Reporting API, Resource Timing, Secure Contexts, Subresource Integrity, Streams, Service Workers, TLS, Upgrade Insecure Requests, URL, Web Crypto, WebDriver BiDi, Web IDL, WebSockets, WebTransport, XMLHttpRequest.
- **Interfaces**: `Headers`, `Body`, `Request`, `Response`, `ReadableStream`, `WritableStream`, `TransformStream`, `ServiceWorkerGlobalScope`.

Procedures And API Details
- The chunk does not describe procedures but lists terms and references.
- IDL definitions for `Headers` include methods like `append()`, `delete()`, `get()`, `has()`, `set()`.
- `Body` mixin includes attributes like `body`, `bodyUsed` and methods like `arrayBuffer()`, `blob()`, `bytes()`, `formData()`, `json()`, `text()`.
- `Request` interface constructor takes `RequestInfo` and optional `RequestInit`.

Nuance Or Contradictions
- The chunk is primarily a reference list and does not present contradictions or nuances in itself.
- Some terms are defined by reference to external specifications, which may evolve independently.

Candidate Wiki Hints
- **Fetch Standard Glossary**: A wiki page summarizing key terms defined across various specifications referenced in the Fetch Standard.
- **Normative References Overview**: A page listing and categorizing normative references (RFCs, W3C specs) used in web standards.
- **Non-Normative References**: A resource documenting non-normative references, including security advisories and legacy APIs.
- **IDL Index for Fetch API**: A technical reference detailing IDL definitions for `Headers`, `Body`, `Request`, and related interfaces.

## chunk-23

---
title: Chunk 23 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

**Heading:** `6. data: URLs`
**Line Range:** 7886–8596
**Subject:** Definitions and browser support matrices for the Fetch API, specifically focusing on `Request`, `Response`, `Headers`, and related attributes/methods.

# Local Summary

This chunk defines the interface structure for the Fetch API within the Web APIs standard. It details the `Request` object's read-only attributes (such as `destination`, `referrer`, `mode`, `credentials`, `cache`, `redirect`, `integrity`, `keepalive`, `signal`, and `duplex`) and its methods (`clone`). It also defines the `Response` interface, including static constructors like `error()`, `redirect()`, and `json()`, along with attributes like `type`, `url`, `status`, `ok`, and `headers`. The chunk further introduces `RequestInit` and `ResponseInit` dictionaries for initialization options. A significant portion of the text is dedicated to browser support tables (compatibility matrices) for various methods and attributes across engines like Firefox, Safari, Chrome, Edge, Opera, and Node.js. Specific attention is given to features like `fetchLater`, `DeferredRequestInit`, and specific body handling (`arrayBuffer`, `blob`, `body`, `text`, `json`).

# Key Claims

- The `Request` interface includes numerous read-only attributes defining the nature of the request (destination, referrer policy, mode, credentials, cache behavior, redirect status, integrity check, keepalive flag, signal for aborting, and duplex type).
- The `Response` interface supports static factory methods (`error`, `redirect`, `json`) and standard attributes (`type`, `url`, `status`, `ok`, `headers`).
- Initialization dictionaries (`RequestInit`, `ResponseInit`) allow configuration of method, headers, body, referrer policy, credentials, cache mode, redirect behavior, integrity, keepalive, signal, duplex, priority, and window context.
- Specific request/response types exist for handling data: `arrayBuffer`, `blob`, `text`, `json`.
- The `Request` interface supports form data (`formData`) via both request and response contexts.
- Browser support varies significantly for specific attributes (e.g., `cache` requires Firefox 48+, Safari 10.1+, Chrome 64+; `signal` requires Firefox 57+, Safari 12.1+, Chrome 66+).

# Entities And Concepts

- **Request Interface:** Represents a request object with attributes like `destination`, `referrer`, `mode`, `credentials`, `cache`, `redirect`, `integrity`, `keepalive`, `isReloadNavigation`, `isHistoryNavigation`, `signal`, and `duplex`.
- **Response Interface:** Represents a response object with attributes like `type`, `url`, `redirected`, `status`, `statusText`, `ok`, and `headers`. Includes methods `clone()`.
- **Headers Interface:** A mixin interface for accessing HTTP headers (though detailed method support is listed in tables).
- **RequestInit Dictionary:** Contains optional parameters for creating a request (`method`, `headers`, `body`, `referrer`, `referrerPolicy`, `mode`, `credentials`, `cache`, `redirect`, `integrity`, `keepalive`, `signal`, `duplex`, `priority`, `window`).
- **ResponseInit Dictionary:** Contains optional parameters for creating a response (`status`, `statusText`, `headers`).
- **RequestDestination Enum:** Values include `"audio"`, `"audioworklet"`, `"document"`, `"embed"`, `"font"`, `"frame"`, `"iframe"`, `"image"`, `"json"`, `"manifest"`, `"object"`, `"paintworklet"`, `"report"`, `"script"`, `"sharedworker"`, `"style"`, `"text"`, `"track"`, `"video"`, `"worker"`, `"xslt"`.
- **RequestMode Enum:** Values include `"navigate"`, `"same-origin"`, `"no-cors"`, `"cors"`.
- **RequestCredentials Enum:** Values include `"omit"`, `"same-origin"`, `"include"`.
- **RequestCache Enum:** Values include `"default"`, `"no-store"`, `"reload"`, `"no-cache"`, `"force-cache"`, `"only-if-cached"`.
- **RequestRedirect Enum:** Values include `"follow"`, `"error"`, `"manual"`.
- **RequestDuplex Enum:** Value is `"half"`.
- **RequestPriority Enum:** Values are `"high"`, `"low"`, `"auto"`.
- **ResponseType Enum:** Values include `"basic"`, `"cors"`, `"default"`, `"error"`, `"opaque"`, `"opaqueredirect"`.
- **FetchLaterResult Interface:** Used with `fetchLater()` method, containing an `activated` attribute.
- **DeferredRequestInit Dictionary:** Extends `RequestInit` with `activateAfter` timestamp.

# Procedures And API Details

**Constructing a Request:**
Use the `new Request()` constructor (implied by `RequestInit`). Key properties to set include `method`, `headers`, `body`, `referrer`, `referrerPolicy`, `mode`, `credentials`, `cache`, `redirect`, `integrity`, `keepalive`, `signal`, and `duplex`.

**Cloning a Request:**
Call `request.clone()` to create a new request object that shares the same body stream. Note: The original request's body becomes unusable after cloning if accessed.

**Creating Responses:**
- Use `Response.error()` for network errors.
- Use `Response.redirect(url, status)` for redirects (default 302).
- Use `Response.json(data, init)` to create a response with JSON data and optional headers/status.
- Use `new Response(body, init)` for custom responses.

**Accessing Headers:**
- Use `headers.get(name)`, `headers.has(name)`, `headers.set(name, value)`, `headers.append(name, value)`, or `headers.delete(name)`.
- For cookie headers specifically, use `getSetCookie()`.

**Reading Response Body:**
- Call `response.arrayBuffer()` to read as an ArrayBuffer.
- Call `response.blob()` to read as a Blob.
- Call `response.text()` to read as text.
- Call `response.json()` to parse JSON directly.
- Check `response.bodyUsed` to see if the body stream has been consumed.

**Deferred Fetching:**
Use `window.fetchLater(input, init)` which returns a `FetchLaterResult`. If configured with `activateAfter` in `DeferredRequestInit`, the request activates at the specified time.

# Nuance Or Contradictions

- **Browser Support Discrepancies:** While many attributes are marked "In all current engines," specific versions vary widely. For example, `Request.cache` requires Firefox 48+, Safari 10.1+, Chrome 64+, while `Request.signal` requires Firefox 57+, Safari 12.1+, Chrome 66+. Older browsers (e.g., Edge Legacy) or mobile webviews often lack support for newer attributes like `cache`, `signal`, or specific body types depending on the engine version.
- **Body Usage Warning:** Cloning a request (`request.clone()`) creates a new object sharing the same body stream. Accessing the body of one request after cloning typically results in errors if the stream is consumed, implying that `bodyUsed` behavior is critical when cloning.
- **Legacy Edge Support:** "Edge (Legacy)" (EdgeHTML) supports very few modern features compared to Chromium-based Edge, often listed as version 14+ with `IENone` indicating IE compatibility notes.

# Candidate Wiki Hints

- **Page: Fetch API Reference**
  - **Content:** Comprehensive documentation of the `Request`, `Response`, and `Headers` interfaces, including attribute descriptions, method signatures, and examples.
  - **Source Support:** Directly supported by the definitions in this chunk.

- **Page: Request Initialization Options**
  - **Content:** A guide to the properties within `RequestInit` (`method`, `headers`, `body`, `referrerPolicy`, `cache`, `credentials`, etc.) and their effects on request behavior (e.g., CORS modes, caching strategies).
  - **Source Support:** Supported by the `RequestInit` dictionary definition.

- **Page: Response Handling Strategies**
  - **Content:** Best practices for handling `Response` objects, including using static methods (`error`, `redirect`, `json`) versus custom responses, and reading body content via `arrayBuffer()`, `blob()`, `text()`, or `json()`.
  - **Source Support:** Supported by the `Response` interface definition and body accessors.

- **Page: Browser Compatibility for Fetch API**
  - **Content:** A compatibility matrix detailing which browsers support specific Fetch API features (e.g., `signal`, `cache`, `priority`) and their minimum versions, helping developers implement polyfills or fallbacks.
  - **Source Support:** Supported by the extensive browser support tables in this chunk.

## chunk-24

---
title: Chunk 24 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
# Chunk Context

This section covers the `data:` URL API within the Fetch Standard, specifically detailing browser support matrices for various response properties and headers. It includes compatibility tables for desktop browsers (Firefox, Safari, Chrome, Edge), mobile browsers, Android WebViews, and Node.js environments. The data indicates version requirements (e.g., Firefox 39+, Chrome 42+) and notes specific limitations or missing support in legacy browsers like IE or certain mobile variants.

# Local Summary

The chunk provides a comprehensive compatibility reference for the `data:` URL scheme across modern web browsers and runtime environments. It lists specific minimum versions required for accessing response properties (such as `response.json_static`, `response.redirected`) and CORS-related headers (like `Access-Control-Allow-Origin`). The source distinguishes between "In all current engines" and those supporting only one engine, highlighting gaps in support for older or specialized browsers like IE and legacy Edge.

# Key Claims

- The `data:` URL API is supported in Firefox 39+, Safari 10.1+, and Chrome 42+ (for basic response properties).
- `response.json_static` requires higher versions: Firefox 115+, Chrome 105+, and Edge 105+.
- `response.redirected` support begins at Firefox 49+, Safari 10.1+, and Chrome 57+.
- CORS headers like `Access-Control-Allow-Origin` are supported in Firefox 3.5+, Safari 4+, and Chrome 4+.
- The `Headers/Sec-Purpose` header is currently supported only in Firefox 115+ and not in other major browsers or Node.js environments listed.
- Legacy browsers like Internet Explorer have varying levels of support, with some headers unsupported ("IENone") while others are present ("IEYes").

# Entities And Concepts

- **data: URLs**: A URL scheme used to reference data directly embedded in the document.
- **Fetch API**: The web API used for network requests, which includes handling `data:` URLs.
- **Response Properties**: Attributes of the fetch response object such as `.json_static`, `.redirected`, `.status`.
- **CORS Headers**: HTTP headers related to Cross-Origin Resource Sharing (e.g., `Access-Control-Allow-*`).
- **Browser Compatibility**: The matrix showing which browser versions support specific features.
- **Android WebView**: The web view component used in Android applications, noted for varying support levels.

# Procedures And API Details

To check if a feature is available:
1.  Consult the compatibility table under the relevant section (e.g., "Response" or "Headers").
2.  Identify the browser column (e.g., Firefox, Chrome).
3.  Look for the version number (e.g., "Firefox39+") indicating the minimum supported version.
4.  Note entries marked with "?" as unsupported or unknown, and "✔MDN" as fully supported according to MDN standards.

# Nuance Or Contradictions

- **Legacy Browser Gaps**: While modern browsers show consistent support ("In all current engines"), legacy browsers like Internet Explorer often have incomplete support (marked as "IENone" for many features).
- **Mobile Variability**: Mobile browsers sometimes lack support even when desktop counterparts do, or vice versa. For instance, `response.json_static` shows "None" for iOS Safari and Android WebView in some contexts.
- **Header Specifics**: Some headers like `Headers/Sec-Purpose` are explicitly marked as supported in only one current engine (Firefox 115+), indicating a significant disparity in implementation across the web platform.

# Candidate Wiki Hints

- Create a page documenting browser compatibility for the Fetch API's `data:` URL handling.
- Document specific CORS header support and version requirements for secure cross-origin requests.
- Add a section on legacy browser limitations when using `data:` URLs with modern APIs.

