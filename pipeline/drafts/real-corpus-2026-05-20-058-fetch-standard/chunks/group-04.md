---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This document synthesizes notes from Chunks 19 through 24 of the Fetch Standard specification (`raw/web/corpus-2026-05-18/058-fetch-standard.md`). The content exclusively covers **Section 6: data: URLs**, transitioning from introductory metadata and licensing into the core algorithmic definitions, structural parsing rules, API interfaces, and browser compatibility matrices for handling `data:` URIs.

# Cross-Chunk Summary

The selected chunks provide a comprehensive view of the `data:` URL scheme within the Fetch Standard:
*   **Chunks 19–20** establish the legal framework (WHATWG licensing) and define the structural representation (`data: URL struct`) for parsing these URLs into components like MIME type, charset, and base64 content.
*   **Chunk 21** details the fetch algorithm specific to `data:` URLs, clarifying that they do not trigger network I/O but instead parse the fragment identifier directly. It highlights restrictions on cross-origin access (CORS) and the inability to perform redirects.
*   **Chunks 22–24** focus on the API layer and implementation details. They define the `Request` and `Response` interfaces relevant to fetching, list normative references (HTML, Fetch Metadata), and provide extensive compatibility matrices for browser engines (Chrome, Firefox, Safari, Edge) regarding properties like `body`, `json()`, and CORS headers.

# Repeated Or Central Claims

*   **No Network I/O**: A consistent theme across Chunks 20 and 21 is that fetching a `data:` URL does not involve network connections or cache partitions; it is an immediate processing step within the fetch algorithm.
*   **Parsing Logic**: The standard requires parsing the URI to extract the media type and parameters from the fragment identifier (Chunk 20, 21).
*   **CORS Restrictions**: While `data:` URLs are local resources, Chunks 21 and 24 note that they are subject to security policies like `Cross-Origin-Opener-Policy` and general CORS checks, often resulting in errors if treated as cross-origin resources.
*   **Interface Definitions**: The chunks repeatedly reference the `Request`, `Response`, and `Headers` interfaces, defining their attributes (e.g., `method`, `url`, `headers`) and methods (e.g., `append`, `get`, `json`).
*   **Compatibility Nuances**: Chunks 23 and 24 emphasize that support for specific features (like `Request.body` or static `Response.json_static`) varies significantly by browser version, with legacy browsers (Edge Legacy) and mobile webviews showing gaps in coverage.

# Important Local Details

*   **Licensing Distinction**: The text explicitly distinguishes between the Creative Commons Attribution 4.0 International license for the specification text and the BSD 3-Clause License for incorporated source code portions (Chunk 19).
*   **Data URL Structure**: A `data: URL struct` holds parsed components including scheme, base64 content, media type, and charset (Chunk 20).
*   **Empty Body Default**: By default, requests for `data:` URLs are created with an empty body, which is populated by reading the data from the URL's fragment during parsing (Chunk 21).
*   **Redirection Failure**: Attempts to redirect a `data:` URL fetch fail because they do not support the redirection mechanism used in HTTP network fetches (Chunk 21).
*   **Specific API Methods**: The chunks detail methods such as `Response.error()`, `Response.redirect()`, and `Response.json()` for creating synthetic responses, alongside the `fetchLater()` method for deferred fetching.
*   **Normative References**: Section 6 inherits definitions from external standards including W3C Living Standards, RFCs (e.g., [HTTP], [HTML]), and security advisories like [HTTPVERBSEC1] (Chunk 22).

# Candidate Wiki Hints

*   **Topic: Data URLs in Fetch API**
    *   A dedicated page explaining the `data:` URL scheme, its parsing logic, and its place within the Fetch Standard's lifecycle.
*   **Concept: Cross-Origin Policies for Data URIs**
    *   Documentation on how CORS policies (`Cross-Origin-Opener-Policy`, etc.) apply to local resources like `data:` URLs and why they might trigger security errors.
*   **API Reference: Request & Response Interfaces**
    *   A reference guide for the `Request` and `Response` objects, detailing attributes (`mode`, `credentials`, `integrity`) and utility methods (`clone`, `json`, `text`).
*   **Guide: Browser Compatibility for Fetch API**
    *   A matrix summarizing support for key features (e.g., `body` access, `fetchLater()`, specific CORS headers) across Chrome, Firefox, Safari, Edge, and mobile webviews.
*   **Topic: Deferred Fetching (`fetchLater`)**
    *   Notes on the experimental or newer `fetchLater()` method and its `DeferredRequestInit` dictionary for lazy loading strategies.

# Gaps Or Cautions

*   **Incomplete Mobile Data**: Chunks 23 and 24 contain numerous question marks (?) in their compatibility matrices, particularly for Android WebViews and iOS Safari, indicating a lack of confirmed support data or testing gaps for these environments.
*   **Legacy Edge Discontinuity**: The text notes a sharp discontinuity between "Edge (Legacy)" (EdgeHTML) and modern Chromium-based Edge, with the former lacking support for features like `Response.body` despite supporting others.
*   **MDN Documentation Status**: Some headers and features are flagged with warning indicators (⚠MDN), suggesting that while they may exist in the browser implementation, their standardization or documentation status is uncertain compared to universally supported items.
*   **Body Readability Variance**: There is a specific contradiction regarding `Request.body` support; while listed as generally available, Firefox only supports it from version 65+, contradicting a blanket "all current engines" claim for older versions.
