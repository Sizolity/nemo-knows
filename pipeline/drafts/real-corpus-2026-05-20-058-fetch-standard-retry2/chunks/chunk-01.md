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
