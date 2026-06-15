---
title: Chunk 18 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

## Chunk Context
**Heading:** 8. Acknowledgements
**Coverage:** Section 8 containing a comprehensive table of Service Worker API support across browser engines (Firefox, Safari, Chrome, Edge, Opera, and Android WebViews). The text lists specific properties, methods, and events with their corresponding version numbers for each engine.

## Local Summary
This section documents the feature parity and version history for the Service Worker API within various browser implementations. It details support for lifecycle management (install/activate/fetch events), client interaction (`postMessage`, `clients`), registration handling (`ServiceWorkerContainer`), and navigation strategies (`NavigationPreloadManager`). Specific attention is given to edge cases like `FetchEvent/respondWith` availability and the transition from legacy Edge versions to modern equivalents.

## Key Claims
- **Universal Support:** Core Service Worker objects (e.g., `ServiceWorker`, `ServiceWorkerContainer`) are available in "all current engines" starting at Firefox 44, Safari 11.1, and Chrome 40.
- **Event Handling:** Lifecycle events such as `install_event` and `fetch_event` follow the same version baseline (Firefox 44+, Safari 11.1+, Chrome 40+) as the core object.
- **Client Interaction:** The `postMessage` method for cross-context communication is supported in Firefox 44+, Safari 11.1+, and Chrome 40+.
- **Navigation Preload:** The `NavigationPreloadManager` interface (including `enable`, `disable`, `getState`) requires newer versions: Firefox 99+, Safari 15.4+, and Chrome 59+.
- **Legacy Edge:** Older versions of Microsoft Edge (labeled "Edge Legacy") show support starting at version 17, while modern Edge aligns with the Chrome versioning scheme (e.g., 79+).

## Entities And Concepts
- **Service Worker API:** The core interface for background scripts.
- **Lifecycle Events:** `install_event`, `activate_event`, `fetch_event`.
- **Client Management:** `clients` property on `ServiceWorkerGlobalScope`.
- **Registration Objects:** `ServiceWorkerContainer`, `ServiceWorkerRegistration`.
- **Navigation Strategies:** `NavigationPreloadManager`.
- **Cross-Process Messaging:** `postMessage`.

## Procedures And API Details
- **Registering a Worker:** Use `navigator.serviceWorker.register()`.
  - *Support:* Firefox 44+, Safari 11.1+, Chrome 40+.
- **Handling Requests:** Implement `fetch_event` in the worker script.
  - *Support:* Firefox 44+, Safari 11.1+, Chrome 40+.
- **Updating Workers:** Call `registration.update()`.
  - *Support:* Firefox 44+, Safari 11.1+, Chrome 45+.
- **Disabling Preload:** Use `NavigationPreloadManager.disable()`.
  - *Support:* Firefox 99+, Safari 15.4+, Chrome 59+.

## Nuance Or Contradictions
- **Version Discrepancies:** The data shows distinct versioning for "Edge Legacy" (e.g., 17+) versus modern "Edge" (e.g., 79+), suggesting a divergence in feature adoption or API naming between the two browser iterations.
- **Android WebView Variance:** Support for `ServiceWorkerContainer/getRegistrations` on Android WebViews is listed as starting at version 40+, whereas `startMessages` requires Firefox 64+ and Chrome 74+, indicating asynchronous updates to specific capabilities within the same engine family.
- **Missing Data Points:** Several entries (marked with `?`) indicate uncertain or unverified support for Opera Mobile and Android WebViews in specific high-level APIs like `FetchEvent/request`.

## Candidate Wiki Hints
- **Page:** Service Worker API Reference
  - *Reason:* The chunk provides a definitive version matrix for all major properties and methods, suitable for a dedicated reference page.
- **Page:** Browser Compatibility for Service Workers
  - *Reason:* The extensive table of browser versions supports a compatibility checker or comparison guide.
