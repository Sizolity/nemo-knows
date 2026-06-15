---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
- **Source**: `raw/web/corpus-2026-05-18/058-fetch-standard.md`
- **Lines**: 1–447
- **Heading Path**: Document > Fetch Standard
- **Coverage**: Preface, Infrastructure (URL, HTTP), HTTP Methods

Local Summary
The Fetch Standard unifies resource fetching across the web platform, ensuring consistent handling of URL schemes, redirects, cross-origin semantics, and related headers. It supersedes previous `Origin` header semantics and provides a unified architecture for APIs like `navigator.sendBeacon()`, `<img>`, and `script` elements. The standard defines core infrastructure including URL classifications (local vs HTTP(S)), HTTP method normalization, and structural data types (`fetch params`, `fetch controller`, `fetch timing info`).

Key Claims
- Fetching is conceptually simple (request in, response out) but implementation details vary; the standard unifies these to handle redirects and CORS consistently.
- The standard defines a `fetch()` JavaScript API that exposes low-level networking functionality.
- Methods like `PATCH` are encouraged over `patch` to avoid `405 Method Not Allowed` errors due to normalization rules.
- Specific methods (`GET`, `HEAD`, `POST`) are CORS-safelisted, while `CONNECT`, `TRACE`, and `TRACK` are forbidden.

Entities And Concepts
- **Fetch Standard**: The living specification for web fetching.
- **Fetch Params**: A struct used for bookkeeping in the fetch algorithm (contains request body processing callbacks, task destination, controller reference).
- **Fetch Controller**: A struct enabling post-start operations on a fetch operation (tracks state: "ongoing", "terminated", "aborted").
- **Fetch Timing Info**: A struct maintaining timing data for Resource and Navigation Timing APIs.
- **URL Schemes**: Defined as "local" (`about`, `blob`, `data`) vs HTTP(S) schemes.
- **HTTP Methods**: Normalized to uppercase for consistency (e.g., `DELETE`, `GET`), though technically case-sensitive in practice.

Procedures And API Details
- **Normalize Method**: Convert methods like `delete`, `get`, `post` to uppercase (`DELETE`, `GET`, `POST`). This is done for backwards compatibility and cross-API consistency.
- **Abort Fetch Controller**: Set state to "aborted", serialize the error (defaulting to `AbortError`), and store the serialized reason.
- **Terminate Fetch Controller**: Set state to "terminated".
- **Collect HTTP Quoted String**: An algorithm to parse quoted strings from input, handling escape sequences (`\`) and end quotes (`"`).

Nuance Or Contradictions
- **Method Case Sensitivity**: While methods are technically case-sensitive (e.g., `Egg` is valid), normalization to uppercase is encouraged for consistency. Using lowercase variants like `patch` often results in errors compared to `PATCH`.
- **Whitespace Handling**: HTTP headers prefer tab/space over other ASCII whitespace, though the standard distinguishes between them based on context (headers vs MIME types).

Candidate Wiki Hints
- **Page: Fetch Standard Overview** – Summarize goals, scope, and relationship to existing standards.
- **Page: Fetch Infrastructure** – Detail `fetch params`, `fetch controller`, and timing info structures.
- **Page: HTTP Method Handling** – Explain normalization rules, CORS-safelisted methods, and forbidden methods.
