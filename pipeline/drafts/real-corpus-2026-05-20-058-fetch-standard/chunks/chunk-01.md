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
