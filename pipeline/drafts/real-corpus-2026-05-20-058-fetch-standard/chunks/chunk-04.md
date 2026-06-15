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
