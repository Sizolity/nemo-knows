---
title: Chunk 19 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---
Chunk Context
The provided chunk is titled "8. Acknowledgements" and contains a list of Service Worker API properties and headers with their browser support status across various engines (Firefox, Safari, Chrome, Edge, Opera, etc.) and versions.

Local Summary
This section lists specific Service Worker API attributes such as `ServiceWorkerRegistration` methods (`update`, `viaCache`, `waiting`) and `WindowClient` properties (`focus`, `focused`, `navigate`, `visibilityState`, `caches`), along with the `Headers` interface property `Service-Worker-Navigation-Preload`. It details the minimum version requirements for each browser to support these features.

Key Claims
- All listed Service Worker APIs are supported in all current engines as of the data source date.
- Specific version thresholds exist for individual properties (e.g., `WindowClient/focus` requires Chrome 42+, Firefox 44+).
- The `Headers/Service-Worker-Navigation-Preload` header is available in a Firefox preview, Safari 15.4+, and Chrome 59+.

Entities And Concepts
- Service Worker Registration APIs: `update`, `viaCache`, `waiting`.
- Window Client APIs: `focus`, `focused`, `navigate`, `visibilityState`.
- Cache Management: `caches`.
- HTTP Headers: `Service-Worker-Navigation-Preload`.
- Browser Engines: Firefox, Safari, Chrome, Edge (Legacy and Current), Opera Mobile.

Procedures And API Details
No specific procedural steps are outlined in this chunk; it serves as a compatibility reference table for Service Worker features.

Nuance Or Contradictions
The chunk indicates support status with question marks (`?`) for mobile versions of Firefox, iOS Safari, Chrome for Android, WebView, Samsung Internet, and Opera Mobile, suggesting data might be pending or unverified for those specific environments compared to the desktop counterparts which have explicit version numbers.

Candidate Wiki Hints
- **Topic**: Service Worker Browser Compatibility
- **Page Idea**: A compatibility matrix page for Service Worker APIs and headers.
