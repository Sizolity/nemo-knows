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
