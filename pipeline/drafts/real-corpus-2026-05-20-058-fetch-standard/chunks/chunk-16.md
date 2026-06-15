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
