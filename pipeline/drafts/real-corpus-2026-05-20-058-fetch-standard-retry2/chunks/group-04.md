---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This group of notes aggregates the concluding sections of the Fetch Standard specification. It covers the transition from core fetch algorithms to specific scheme handling (`data:` URLs), followed by comprehensive legal metadata (copyrights, licenses), a glossary of referenced terms, and extensive browser compatibility matrices for the Fetch API interfaces (`Request`, `Response`, `Headers`). The notes span from administrative headers regarding authorship and licensing to technical definitions of attributes, methods, and version-specific support across major engines.

# Cross-Chunk Summary

The progression of these chunks moves from high-level legal and structural metadata (Chunks 19) into the specific implementation details of `data:` URL handling (Chunks 20–24). While Chunks 20–23 focus on definitions, attributes, and methods for the Fetch API generally (including `Request`, `Response`, and initialization dictionaries), Chunk 24 specifically narrows the scope to compatibility tables for `data:` URLs and specific response properties. The content synthesizes the distinction between the "Living Standard" text (CC BY 4.0) and source code portions (BSD 3-Clause), defines the interface structures including enums like `RequestMode` and `RequestCache`, and provides a reference for browser support ranging from modern Chromium/Firefox/Safari to legacy environments like Internet Explorer and Android WebViews.

# Repeated Or Central Claims

- **Interface Definitions:** The `Request` and `Response` interfaces are central, featuring read-only attributes (e.g., `url`, `type`, `status`) and methods (e.g., `clone`, `json`, `text`). Initialization dictionaries (`RequestInit`, `ResponseInit`) allow configuration of behavior such as caching, redirection, and credentials.
- **Data URL Handling:** `data:` URLs are treated as a distinct resource type within the Fetch API, processed via a specific "data: URL processor" rather than standard network traversal. They support MIME types like `text`, `video`, and `xslt`.
- **Security and Origin:** The specification references security contexts, CORS protocols (including preflight checks), and origin headers. Specific flags like `use-CORS-preflight` and `unsafe-request` dictate handling logic for cross-origin or local resources.
- **Browser Compatibility:** A recurring theme is the variation in support across engines. Modern browsers generally share features ("In all current engines"), but legacy browsers (IE, Edge Legacy) and specific mobile environments show significant gaps or require higher version numbers for features like `signal` or `cache`.

# Important Local Details

- **Licensing Distinction:** The specification text is licensed under Creative Commons Attribution 4.0 International (CC BY 4.0), while source code portions incorporating the standard are licensed under the BSD 3-Clause License.
- **Authorship and Copyright:** Anne van Kesteren (Apple) is identified as the author, with copyright held by WHATWG (representing Apple, Google, Mozilla, Microsoft).
- **Versioning:** The document distinguishes between the "Living Standard" (current version) and a separate "Patent-Review Version."
- **Attribute Details:**
  - `Request` attributes include `destination`, `referrer`, `mode`, `credentials`, `cache`, `redirect`, `integrity`, `keepalive`, `signal`, and `duplex`.
  - `Response` attributes include `type`, `url`, `status`, `ok`, `headers`, and methods like `clone()`.
  - Enums defined include `RequestDestination` (e.g., `"script"`, `"image"`), `RequestMode` (`"navigate"`, `"cors"`), and `ResponseType` (`"basic"`, `"opaque"`).
- **Compatibility Thresholds:**
  - Basic `data:` URL support: Firefox 39+, Safari 10.1+, Chrome 42+.
  - `response.json_static`: Requires higher versions (e.g., Firefox 115+).
  - `Headers/Sec-Purpose`: Currently supported only in Firefox 115+.

# Candidate Wiki Hints

- **Page: Fetch API Interfaces**: A reference page detailing the `Request`, `Response`, and `Headers` objects, their attributes, methods, and initialization dictionaries.
- **Page: Data URL Scheme**: Documentation covering the syntax, processing logic, and MIME type handling for embedded resources.
- **Page: Browser Compatibility Matrix**: A comparative table summarizing support for Fetch API features across Firefox, Safari, Chrome, Edge, Node.js, and legacy browsers.
- **Page: Licensing Overview**: Explaining the split between CC BY 4.0 (spec text) and BSD 3-Clause (source code).

# Gaps Or Cautions

- **Legacy Environment Limitations:** Many features are unsupported or marked as "None" in Internet Explorer and certain mobile webviews, which may require polyfills or fallback logic.
- **Streaming Body Caution:** The specification notes that cloning a request (`request.clone()`) shares the body stream; accessing the body after cloning may result in errors if the stream is consumed, making `bodyUsed` checks critical.
- **Header Specificity:** Some headers (e.g., `Sec-Purpose`) are not universally supported, appearing only in specific modern versions of Firefox, which could lead to inconsistencies in cross-browser implementations.
- **Incomplete Reference Data:** While a glossary and IDL index are provided, some terms rely on external specifications (HTML, DOM) that evolve independently, requiring careful linkage to avoid drift.
