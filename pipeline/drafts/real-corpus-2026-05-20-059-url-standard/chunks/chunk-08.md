---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---

## Chunk Context
This section (Heading path: 6. API > 6.3. URL APIs elsewhere) details browser and runtime support for the `URL` interface and `URLSearchParams` object. It lists specific properties (e.g., `hostname`, `searchParams`) and methods (e.g., `append`, `delete`, `sort`) alongside their minimum supported versions for browsers like Firefox, Chrome, Safari, Edge, and Node.js.

## Local Summary
The chunk provides a comprehensive compatibility matrix for URL parsing and query string manipulation APIs. It distinguishes between the core `URL` interface (introduced in Node 10) and the more granular properties available in Node 20+. The `URLSearchParams` API is shown to be widely supported across modern browsers since around 2015-2017, with specific methods like `delete` and `sort` having slightly later adoption dates.

## Key Claims
- **Node.js Support**: The core `URL` class requires Node.js 10.0.0+. The `URLSearchParams` interface is supported from Node.js 10.0.0+, while specific properties like `URL.pathname` and methods on `URLSearchParams` generally require Node.js 7.5.0+ or higher (specifically 7.7.0+ for `sort`).
- **Browser Support**: All current engines support the `URL` interface. Firefox requires version 19+, Chrome 32+, and Safari 7+. The `URLSearchParams` object is supported from Firefox 29, Chrome 49, and Safari 10.1.
- **Method Granularity**: While basic access (e.g., `get`) is available in Firefox 29, methods like `delete` require Firefox 14 (Safari) or specific versions, and `sort` is not universally supported until newer versions (Firefox 54, Chrome 61).
- **Legacy Browsers**: Edge (Legacy) support is noted for specific versions (e.g., IE10+ for the base `URL`), while Opera support data is often marked with a question mark or specific version numbers like 79+.

## Entities And Concepts
- **URL Interface**: Represents a parsed URL string. Properties include `href`, `hostname`, `pathname`, `search`, etc.
- **URLSearchParams**: A convenient object for working with query strings. Methods include `append`, `delete`, `get`, `getAll`, `set`, `has`, `keys`, `values`, `forEach`, `toString`.
- **Node.js Versions**: Ranges from 7.0.0 to 20.0.0+ are referenced for API availability.
- **Browser Engines**: Firefox, Chrome, Safari, Edge (Legacy and Chromium), Opera, Samsung Internet.

## Procedures And API Details
- **URL Object Creation**: Supported in all current engines; Node.js 10.0.0+.
- **Property Access**:
    - `href`: Available since Firefox 19, Chrome 32, Safari 7, Node 10.
    - `hostname`: Available since Firefox 22, Chrome 32, Node 7.
    - `pathname`: Available since Firefox 22, Chrome 32, Node 7.
    - `searchParams`: Available since Firefox 29, Chrome 49, Safari 10.1, Node 7.5.
- **URLSearchParams Methods**:
    - `append`, `get`, `getAll`, `has`, `set`, `keys`, `values`, `toString`: Supported from Firefox 29/Chrome 49/Safari 10.1/Node 7.5.
    - `delete`: Supported from Firefox 29 (note: Safari requires 14 in some contexts or is implied by browser age), Chrome 49, Node 7.5.
    - `forEach`, `size`: Supported from Firefox 44/Chrome 49/Safari 10.1/Node 7.5.
    - `sort`: Supported from Firefox 54/Chrome 61/Safari 11/Node 7.7.

## Nuance Or Contradictions
- **Safari Versioning**: Some entries list Safari 10.1 for general `URLSearchParams`, while others specify Safari 17 for the `size` property, indicating that not all methods were available at the object's introduction.
- **Edge Discrepancies**: "Edge (Legacy)" often cites IE versions (e.g., IE10) alongside Edge numbers, suggesting a conflation of support timelines between Internet Explorer and early Chromium-based Edge in the source data.
- **Opera Ambiguity**: Many Opera entries are marked with "?", indicating uncertain or unverified support data for that browser compared to others.

## Candidate Wiki Hints
- **Page: URL_API_Compatibility** - Documenting the specific versions of browsers and runtimes where `URL` and `URLSearchParams` became available.
- **Page: URLSearchParams_Methods** - Detailing which methods (`append`, `delete`, `sort`) require polyfills for older environments.
- **Topic: Polyfill_Necessity** - Identifying that `URLSearchParams.sort()` might require a polyfill for browsers prior to Firefox 54 or Chrome 61.
