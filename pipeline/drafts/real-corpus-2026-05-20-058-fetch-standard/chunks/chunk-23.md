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
