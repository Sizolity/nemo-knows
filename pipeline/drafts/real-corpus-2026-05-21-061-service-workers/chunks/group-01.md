---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

# Group Context

This document synthesizes the Service Workers specification (W3C Candidate Recommendation Draft, April 2026). It covers the full lifecycle from motivations and model definitions to execution contexts, client management, event handling, caching strategies, and security considerations. The content details how service workers function as event-driven worker contexts that intercept network requests, manage offline functionality, and execute background tasks without relying on document lifetimes. The scope includes the core API interfaces (`ServiceWorker`, `ServiceWorkerRegistration`, `Client`, `Cache`), their internal states, and the algorithms governing their interaction with the browser environment.

# Cross-Chunk Summary

The specification defines service workers as generic, event-driven script contexts running at an origin, distinct from shared workers in that they never handle messages directly but process events. They are managed via a registration object containing three worker states: `installing`, `waiting`, and `active`. The lifecycle begins with installation (adding routes), activation (taking control of clients), and termination (redundancy). Communication occurs asynchronously via `postMessage` between the global scope (`ServiceWorkerGlobalScope`) and client environments (window or worker). Caching is handled through a separate storage system (`CacheStorage`) distinct from the browser's HTTP cache, allowing manual management of offline resources. Security constraints include origin relativity, CORS restrictions on cross-origin resources, and strict path/script requests.

# Repeated Or Central Claims

- **Event-Driven Execution:** Service workers operate on an event loop driven by events (fetch, install, activate, message) rather than document lifecycle. They are asynchronous to avoid blocking resource loading.
- **Registration Lifecycle:** A registration persists across user agent restarts but manages three distinct worker instances: `installing` (current script), `waiting` (newer script pending activation), and `active` (currently controlling clients). The `update()` method handles upgrading from waiting to active, while `skipWaiting()` bypasses the waiting phase.
- **Client Control:** Clients (windows or workers) are controlled by a service worker if they share an origin and the service worker is registered for that scope. `navigator.serviceWorker.controller` returns the currently controlling worker (or null). `claim()` allows an active worker to take control of clients it does not currently control.
- **Caching Separation:** Caches (`self.caches`) are isolated from the browser's HTTP cache. They must be manually populated, versioned, and deleted. Responses are stored as immutable resources keyed by request info.
- **Error Recovery:** Unlike Application Cache, service workers are designed to avoid unrecoverable states. Errors during installation or activation trigger specific behaviors (e.g., moving to redundant state) rather than leaving the system in a broken state.

# Important Local Details

- **Service Worker Global Scope Attributes:**
  - `clients`: Returns a `Clients` object representing all controlled clients.
  - `registration`: The associated `ServiceWorkerRegistration`.
  - `serviceWorker`: The current `ServiceWorker` instance (if any).
  - `skipWaiting()`: A boolean flag or method to skip the waiting period.
- **Client Interface Attributes:**
  - `url`: The creation URL of the client's browsing context.
  - `type`: "window", "worker", or "sharedworker".
  - `frameType`: "auxiliary", "top-level", "nested", or "none".
  - `visibilityState` and `focused`: Specific to `WindowClient`.
- **Event Attributes:**
  - `InstallEvent`: Includes `event.addRoutes(rules)` for routing logic.
  - `FetchEvent`: Includes `event.request`, `event.respondWith(r)`, `event.clientId`, `event.handled`.
  - `ExtendableMessageEvent`: Includes `event.data`, `event.origin`, `event.source`.
- **Cache API Operations:**
  - `match()`: Returns the first matching response.
  - `matchAll()`: Returns a frozen array of all matches.
  - `put()`: Stores a request/response pair; requires full body read.
  - `delete()`: Removes entries; returns boolean indicating success.
- **Timing Struct:** Includes fields like `startTime`, `fetchEventDispatchTime`, and various worker router evaluation timestamps for performance tracing.

# Candidate Wiki Hints

- **Service Worker Lifecycle Guide:** Explain the transition from installation to activation, including the roles of `installing`, `waiting`, and `active` states, and how `skipWaiting()` interacts with this flow.
- **Client Management API:** Detail the usage of `navigator.serviceWorker.controller`, `clients.matchAll()`, and `clients.get()` for querying active clients.
- **Caching Strategies:** Provide examples of using `caches.open()`, `cache.put()`, and `cache.match()` to implement offline-first architectures, including cache versioning strategies.
- **Event Handling Patterns:** Describe how to use `event.waitUntil()` and `event.respondWith()` to extend event lifetimes and handle fetch requests without blocking the main thread.
- **Security Considerations:** Summarize constraints on script loading (`importScripts`), cross-origin resource handling, and path restrictions for service worker scripts.

# Gaps Or Cautions

- **Appendix Content:** Chunks 8 through 12 cover "Appendix A: Algorithms" in detail but do not contain the actual algorithmic pseudocode text within these specific chunk notes; they serve as placeholders or structural markers for detailed algorithm definitions.
- **Living Document Status:** The specification is explicitly noted as a "living document" with unimplemented features and potential changes. Citation outside of a work-in-progress context is discouraged.
- **Timing Implementation:** Timing values (e.g., `startTime`) are initially 0 and updated during execution; exact semantics depend on WPT results and user agent implementation details not fully elaborated in the provided notes.
- **Opaque Origins:** Sandboxed iframes without specific directives (`allow-same-origin`, `allow-scripts`) result in an opaque origin, causing their active service worker to be null. Similarly, data URLs and Blob URLs inherently result in opaque origins for clients.
- **Synchronous Iteration:** The `CacheStorage` interface explicitly excludes synchronous iteration methods (like `forEach`) pending TC39 async iteration standards, requiring users to handle cache storage asynchronously.
