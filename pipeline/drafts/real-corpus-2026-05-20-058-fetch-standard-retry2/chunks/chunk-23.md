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
