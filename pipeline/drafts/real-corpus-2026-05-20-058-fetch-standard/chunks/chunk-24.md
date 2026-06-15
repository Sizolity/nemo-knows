---
title: Chunk 24 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This section details browser and engine support for the Fetch API and associated Headers. It covers `Request` methods, various `Response` properties (including static methods like `json_static`), specific response status handling (`redirected`, `statusText`), and a comprehensive list of CORS-related headers (`Access-Control-*`, `Cross-Origin-Resource-Policy`, `Sec-Purpose`) along with legacy header support.

# Local Summary

The chunk outlines compatibility matrices for Fetch API features across major browsers (Firefox, Safari, Chrome, Edge, Opera) and their variants (Android, iOS WebView, Samsung Internet). It specifies version requirements for standard properties like `Response.url` and `fetch`, noting that some features are supported in "all current engines" while others have specific version cutoffs (e.g., `Response/json_static` requires Firefox 115+ or Chrome 105+). The text also highlights headers available in all engines versus those restricted to single engines or marked with warnings regarding MDN documentation status.

# Key Claims

- **Universal Support**: Properties like `Response/url`, `Response/status`, and the `fetch` function itself are supported in "all current engines" starting from Firefox 39, Safari 10.1, and Chrome 42.
- **Version Specifics**: `Response/json_static` is not universally available; it requires Firefox 115+, Chrome 105+, or Edge 105+.
- **Legacy Support**: `Edge (Legacy)` supports many features starting from version 14, while `IE` support varies or is marked as non-existent for modern standards.
- **CORS Headers**: Standard CORS headers (`Access-Control-Allow-Origin`, etc.) are supported in Firefox 3.5+, Safari 4+, and Chrome 4+.
- **Warning Flags**: Certain headers like `Headers/Origin` and features in Android WebViews carry warning indicators (⚠) or question marks, suggesting potential inconsistencies or incomplete documentation.

# Entities And Concepts

- **Fetch API**: The core interface for making network requests (`fetch`).
- **Response Object**: Properties including `type`, `url`, `status`, `redirected`, and static methods like `json_static`.
- **Request Object**: Represents the request initiated by `fetch`.
- **CORS Headers**: A set of headers used to control cross-origin resource sharing (e.g., `Access-Control-Allow-Credentials`, `Cross-Origin-Resource-Policy`).
- **Browser Variants**: Specific builds for mobile operating systems such as Firefox for Android, iOS Safari, and Samsung Internet.

# Procedures And API Details

- **Checking Support**: The text implies a procedure of checking version numbers (e.g., "Firefox 39+") to determine if a specific feature or property is available in a given engine.
- **Static Methods**: `Response.json_static` is identified as a method that parses response data, with strict version requirements for its availability.
- **Header Enumeration**: The chunk lists specific header names (e.g., `Access-Control-Expose-Headers`) and indicates their support status across different browser versions.

# Nuance Or Contradictions

- **Documentation Status**: Some entries are marked with "⚠MDN" or question marks, indicating that while the feature might exist in the browser, its documentation or standardization status is uncertain compared to those marked simply with "✔MDN".
- **Engine Variations**: While desktop browsers show clear version numbers (e.g., Chrome 42+), mobile variants often have missing data points (represented by "?"), suggesting gaps in testing or reporting for those specific environments.
- **Legacy vs. Modern**: There is a distinction made between modern Edge and "Edge (Legacy)", with the latter supporting older versions (14+) but having different compatibility profiles than the modern Chromium-based Edge.

# Candidate Wiki Hints

- **Fetch API Compatibility Table**: Create a dedicated page or section summarizing browser support for Fetch API properties, distinguishing between universal support and version-specific requirements.
- **CORS Headers Reference**: Compile a list of CORS-related headers with their minimum browser versions and any known limitations or partial implementations in mobile webviews.
- **Response Object Properties**: Document the `Response` object properties, specifically highlighting the availability of static methods like `json_static` across different browsers.
