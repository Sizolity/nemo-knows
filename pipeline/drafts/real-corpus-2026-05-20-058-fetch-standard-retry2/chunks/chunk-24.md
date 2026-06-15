---
title: Chunk 24 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
# Chunk Context

This section covers the `data:` URL API within the Fetch Standard, specifically detailing browser support matrices for various response properties and headers. It includes compatibility tables for desktop browsers (Firefox, Safari, Chrome, Edge), mobile browsers, Android WebViews, and Node.js environments. The data indicates version requirements (e.g., Firefox 39+, Chrome 42+) and notes specific limitations or missing support in legacy browsers like IE or certain mobile variants.

# Local Summary

The chunk provides a comprehensive compatibility reference for the `data:` URL scheme across modern web browsers and runtime environments. It lists specific minimum versions required for accessing response properties (such as `response.json_static`, `response.redirected`) and CORS-related headers (like `Access-Control-Allow-Origin`). The source distinguishes between "In all current engines" and those supporting only one engine, highlighting gaps in support for older or specialized browsers like IE and legacy Edge.

# Key Claims

- The `data:` URL API is supported in Firefox 39+, Safari 10.1+, and Chrome 42+ (for basic response properties).
- `response.json_static` requires higher versions: Firefox 115+, Chrome 105+, and Edge 105+.
- `response.redirected` support begins at Firefox 49+, Safari 10.1+, and Chrome 57+.
- CORS headers like `Access-Control-Allow-Origin` are supported in Firefox 3.5+, Safari 4+, and Chrome 4+.
- The `Headers/Sec-Purpose` header is currently supported only in Firefox 115+ and not in other major browsers or Node.js environments listed.
- Legacy browsers like Internet Explorer have varying levels of support, with some headers unsupported ("IENone") while others are present ("IEYes").

# Entities And Concepts

- **data: URLs**: A URL scheme used to reference data directly embedded in the document.
- **Fetch API**: The web API used for network requests, which includes handling `data:` URLs.
- **Response Properties**: Attributes of the fetch response object such as `.json_static`, `.redirected`, `.status`.
- **CORS Headers**: HTTP headers related to Cross-Origin Resource Sharing (e.g., `Access-Control-Allow-*`).
- **Browser Compatibility**: The matrix showing which browser versions support specific features.
- **Android WebView**: The web view component used in Android applications, noted for varying support levels.

# Procedures And API Details

To check if a feature is available:
1.  Consult the compatibility table under the relevant section (e.g., "Response" or "Headers").
2.  Identify the browser column (e.g., Firefox, Chrome).
3.  Look for the version number (e.g., "Firefox39+") indicating the minimum supported version.
4.  Note entries marked with "?" as unsupported or unknown, and "✔MDN" as fully supported according to MDN standards.

# Nuance Or Contradictions

- **Legacy Browser Gaps**: While modern browsers show consistent support ("In all current engines"), legacy browsers like Internet Explorer often have incomplete support (marked as "IENone" for many features).
- **Mobile Variability**: Mobile browsers sometimes lack support even when desktop counterparts do, or vice versa. For instance, `response.json_static` shows "None" for iOS Safari and Android WebView in some contexts.
- **Header Specifics**: Some headers like `Headers/Sec-Purpose` are explicitly marked as supported in only one current engine (Firefox 115+), indicating a significant disparity in implementation across the web platform.

# Candidate Wiki Hints

- Create a page documenting browser compatibility for the Fetch API's `data:` URL handling.
- Document specific CORS header support and version requirements for secure cross-origin requests.
- Add a section on legacy browser limitations when using `data:` URLs with modern APIs.
