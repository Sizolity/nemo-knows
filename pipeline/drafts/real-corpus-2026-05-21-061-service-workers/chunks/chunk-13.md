---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
- **Location**: Appendix B: Extended HTTP headers (Section 7. Extensibility).
- **Focus**: Defines the `Service-Worker` request header and the `Service-Worker-Allowed` response header for controlling script resource requests and scope restrictions.
- **Lines**: 5242–5308.

Local Summary
This section details the HTTP headers exchanged when fetching a service worker script. The `Service-Worker` header identifies the request type for logging and threat detection. The `Service-Worker-Allowed` header allows the server to override the default scope path restriction, enabling a script to control URLs outside its physical location. Examples illustrate how these headers interact with the `navigator.serviceWorker.register()` API to either succeed or fail based on scope validity.

Key Claims
- A request to fetch a service worker script includes the `Service-Worker` header.
- The `Service-Worker-Allowed` response header allows the user agent to override the path restriction limiting the maximum allowed scope URL.
- If `Service-Worker-Allowed` is omitted, the maximum allowed scope defaults to the path where the script sits (e.g., `/js/`).
- A relative URL in `Service-Worker-Allowed` is parsed against the script's URL.
- Validation for `Service-Worker-Allowed` uses the URL parsing algorithm rather than ABNF.

Entities And Concepts
- **Service-Worker**: An HTTP request header indicating a service worker script resource request.
- **Service-Worker-Allowed**: An HTTP response header defining the overridden maximum allowed scope URL.
- **Scope Restriction**: The default security constraint limiting a service worker's scope to its installation path unless overridden.
- **ABNF**: Augmented Backus-Naur Form, used for syntax definitions (though not strictly applied to `Service-Worker-Allowed` values).

Procedures And API Details
- **Default Scope Behavior**:
  - Script location: `/js/sw.js`
  - Default max allowed scope: `/js/`
  - Registration code: `navigator.serviceWorker.register("/js/sw.js")`
- **Overriding with Header (Success Case)**:
  - Response includes: `Service-Worker-Allowed: /`
  - Script location: `/js/sw.js`
  - Requested scope: `/`
  - Result: Installation succeeds as the overridden max allowed scope is `/`.
- **Overriding Without Header (Failure Case)**:
  - Response has no `Service-Worker-Allowed` header.
  - Script location: `/js/sw.js`
  - Requested scope: `/`
  - Result: Installation fails due to path restriction violation.
- **Partial Override Failure**:
  - Response includes: `Service-Worker-Allowed: /foo`
  - Script location: `/foo/bar/sw.js`
  - Requested scope: `/`
  - Result: Installation fails because the requested scope is still outside the overridden maximum allowed scope (`/foo`).

Nuance Or Contradictions
- **Syntax vs. Logic**: While `Service-Worker` uses ABNF (`%x73.63.72.69.70.74 ; "script"`), the validation of `Service-Worker-Allowed` values relies on the URL parsing algorithm, not ABNF.
- **Relative URLs**: The documentation notes that relative URLs in `Service-Worker-Allowed` must be resolved against the script's own URL, which is a specific resolution rule distinct from absolute paths.

Candidate Wiki Hints
- Service Worker Scope Restrictions and Overrides
- HTTP Headers for Service Workers (`Service-Worker`, `Service-Worker-Allowed`)
- Understanding the Default Service Worker Scope
