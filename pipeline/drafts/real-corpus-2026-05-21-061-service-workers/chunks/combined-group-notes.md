## group-01

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

## group-02

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

# Group Context

This group synthesizes the core implementation details and security constraints of the Service Worker specification. It bridges the high-level model (Motivations, Model) with the low-level execution context (Algorithms). The content covers the entire lifecycle from registration and script fetching to installation, activation, event handling, caching strategies, and secure termination. Key themes include the strict requirements for secure contexts, the mechanism of "jobs" for managing state transitions, the internal algorithms for fetch routing, and the security implications of persistent storage and cross-origin resource handling.

# Cross-Chunk Summary

The specification defines a robust lifecycle management system centered on **Jobs**. A job represents a unit of work (register, update, unregister) containing attributes like scope URL, script URL, and worker type. Jobs are queued per scope URL; if the queue is empty or the new job differs from the pending one, it is enqueued. Execution happens in parallel with DOM manipulation tasks.

The **Service Worker Lifecycle** flows through distinct states:
1.  **Installation:** Triggered by fetching a new script. The `Install` algorithm runs, updating state to "installing," queuing an `updatefound` event, and waiting for asynchronous extensions (`extend lifetime promises`). If successful, it dispatches the `install` event.
2.  **Activation:** The worker transitions from "waiting" to "activated." It terminates the previous active worker (if redundant or idle) and matches clients to the new registration, resolving their `readyPromise`.
3.  **Termination:** Workers are cleared when no longer needed. The `Terminate Service Worker` algorithm sets a closing flag, backs up functional events (fetch/push), and discards message tasks.

**Fetch Handling** is managed by the `Handle Fetch` algorithm. It determines if a worker can handle a request based on security context. It matches registrations using storage keys and evaluates **Router Conditions**. These conditions support `urlPattern`, `requestMethod`, and logical operators (`_or`, `not`). The system supports "Soft Updates" to refresh stale workers without blocking, ensuring the worker is current before responding.

**Caching** is handled via the Cache API (`self.caches` and `CacheStorage`). Responses are stored as either **Basic Filtered**, **CORS Filtered**, or **Opaque Filtered**. The spec enforces strict validation for cache operations (e.g., HTTP/HTTPS schemes) and provides atomic `Batch Cache Operations` with rollback capabilities on failure.

# Repeated Or Central Claims

-   **Secure Context Requirement:** Service workers must execute in secure contexts (typically HTTPS). Exceptions exist only for localhost, 127.0.0.0/8, and ::1/128 during development. Both the worker and the registering client must be secure.
-   **Origin Relativity:** A service worker executes within the registering client's origin. It cannot host resources on a CDN directly unless specific headers are present. This prevents cross-origin attacks and enforces strict same-origin policies for scripts.
-   **Job Equivalence:** Two jobs are equivalent if they share specific attributes (type, scope URL, script URL, worker type, cache mode). If a new job is equivalent to a pending one, it is appended to that job's list rather than creating a duplicate entry.
-   **Fetch Routing Logic:** The `Handle Fetch` algorithm determines the response source (cache, network, or fetch-event listener) based on router conditions. This logic includes race responses for GET requests and soft updates for stale workers.
-   **Router Condition Limits:** To prevent performance degradation, router condition complexity is constrained: total nested conditions cannot exceed 1024, and nesting depth is limited to 10 levels.
-   **Security on Regex:** User-defined regular expressions in URL patterns are prohibited due to security concerns. If a pattern has regexp groups, the router condition verification fails.

# Important Local Details

-   **Fragment Handling:** The specification explicitly discards fragments (hashes) in both script URLs and scope URLs during identification. Only the scheme and path matter.
-   **Bad Import Responses:** Responses with bad MIME types or non-ok status codes for `importScripts()` are ignored for the byte-to-byte update check but may still populate the cache. This prevents minor fetch errors from triggering unnecessary re-installs.
-   **Event Skipping:** To avoid performance delays, the user agent may skip event dispatch if no listeners exist deterministically added during the first script execution.
-   **Batch Cache Operations:** These operations perform atomic writes. If an exception occurs (e.g., quota exceeded), the implementation rolls back all changes made during the batch job.
-   **Clear Registration:** A registration is cleared only if its installing, waiting, and active workers are null or have no pending events. This ensures data integrity before removing a worker from storage.
-   **Scope Matching:** URL string matching for service worker scopes is prefix-based rather than path-structural. It maintains same-origin security but relies on trailing slashes in serialized URLs for safety.

# Candidate Wiki Hints

-   **Service Worker Job Lifecycle:** A deep dive into how `Jobs` manage state transitions, queue synchronization, and promise resolution.
-   **Secure Context Enforcement:** Explaining the requirements for HTTPS, localhost exceptions, and CSP headers.
-   **Fetch Event Routing:** Detailing the `Handle Fetch` algorithm, router conditions, soft updates, and race responses.
-   **Cache API Security:** Covering response types (Basic/CORS/Opaque filtered), validation rules, and batch operations with rollback.
-   **Router Condition Architecture:** Documenting the structure of router conditions (`urlPattern`, `requestMethod`, `_or`, `not`) and their complexity limits.
-   **Installation vs. Activation:** A visual or textual guide to the state machine (Installing -> Waiting -> Active) and event dispatching.

# Gaps Or Cautions

-   **Algorithm Implementation Flexibility:** While the spec defines minimal setup for `Setup ServiceWorkerGlobalScope`, implementations can optimize this step as long as observable behavior (results of security checks like CSP) remains equivalent.
-   **Activation Robustness:** Activation handlers are designed to do non-essential work because they may not complete if the browser terminates during activation. Handlers must function properly even if they fail.
-   **Plugin Discouragement:** Plugins are explicitly discouraged from loading via service workers due to origin handling limitations in the `Handle Fetch` algorithm, specifically regarding security origins.
-   **Privacy and Data Purge:** Persistent storages (registration map, cache name to map, script resource map) must be cleared when users purge data, but specific details on *how* this happens across different storage mechanisms may vary by implementation.
-   **Empty Fetch Listeners:** The spec notes that empty function bodies in `fetch` listeners (`() => {}`) are used by some sites to signal PWA status but may have performance implications.

## group-03

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

# Group Context

This group of notes synthesizes the technical specification for **Service Workers**, covering their motivations, model, registration lifecycle, client contexts, execution environments, cache storage management, security considerations, extensibility, and browser compatibility. The source document details the API definitions, algorithms, and conformance requirements necessary to implement a service worker environment that enables offline capabilities, resource caching, and network interception.

# Cross-Chunk Summary

The specification defines Service Workers as background scripts that run in a separate execution context (`ServiceWorkerGlobalScope`) from the main thread. The core workflow involves:
1.  **Registration**: A client script registers a service worker script URL via `navigator.serviceWorker.register()`.
2.  **Lifecycle**: The script undergoes states (`parsed`, `installing`, `installed`, `activating`, `activated`, `redundant`) managed by the `ServiceWorkerRegistration` object.
3.  **Activation**: Once active, the worker controls a specific URL scope. It can intercept network requests via the `fetch` event and cache responses.
4.  **Communication**: The worker communicates with clients using `postMessage()`. Clients query workers via the `clients` property or the `Clients` API (`matchAll`).
5.  **Caching**: Workers use the Cache Storage API to store and retrieve resources, managing versions and lifetimes independently of network availability.
6.  **Security & Scope**: By default, a service worker controls URLs relative to its script location. This can be overridden via the `Service-Worker-Allowed` HTTP response header. The worker must run in a secure context (HTTPS or localhost).

# Repeated Or Central Claims

*   **Lifecycle States are Critical**: Service Workers transition through distinct states (`waiting`, `active`, `redundant`, `unregistered`). An installed worker enters the `waiting` state until it activates, at which point it controls the scope. The `skipWaiting()` method allows skipping this queue.
*   **Scope is Path-Based**: A service worker's scope defaults to the directory containing its script file (e.g., a script at `/js/sw.js` controls `/js/`). This restriction can be explicitly expanded using the `Service-Worker-Allowed` header in the HTTP response.
*   **Extendable Events**: Standard events (`fetch`, `install`, `activate`) are extended to include methods like `event.waitUntil()` (to defer task completion) and `event.respondWith()` (to intercept requests).
*   **Cache Storage is Persistent**: The Cache API allows storing responses as `Response` objects. It supports operations like `put`, `match`, `delete`, and `keys`. Caches are managed via a global storage accessible by name.
*   **Client Management is Robust**: The `Clients` interface allows querying for clients (`get`, `matchAll`) based on type (window, worker), focus state, visibility, and frame hierarchy. Workers can claim existing clients to gain control of them.

# Important Local Details

*   **HTTP Headers**:
    *   `Service-Worker`: Request header identifying the resource as a service worker script.
    *   `Service-Worker-Allowed`: Response header allowing the server to override the default scope path restriction (e.g., allowing `/` from a script at `/js/sw.js`).
*   **API Definitions**:
    *   `ServiceWorkerRegistration`: Manages lifecycle (`installing`, `waiting`, `active`), update mechanisms (`update()`, `unregister()`), and navigation preload settings.
    *   `ServiceWorkerContainer`: Provides global registration management (`getRegistrations`, `startMessages`).
    *   `NavigationPreloadManager`: Enables/disables the `Link: preload=` header behavior.
*   **Error Handling**:
    *   Fetch tasks can fail with a `TypeError` (e.g., if stream construction fails).
    *   The `messageerror` event fires when a client sends an unserializable message to a worker.
*   **Conformance**:
    *   Normative text uses RFC 2119 keywords (`MUST`, `SHOULD`).
    *   Algorithms are defined for clarity; implementations may optimize steps (e.g., parallelizing cache writes).

# Candidate Wiki Hints

*   **Service Worker Scope and Registration**: A guide explaining how to register a worker, the default scope rules, and how to use `Service-Worker-Allowed` headers to override them.
*   **The Service Worker Lifecycle**: An interactive diagram or page detailing the transitions between `parsed`, `installing`, `waiting`, and `active` states, including the role of `skipWaiting()`.
*   **Caching Strategies**: A section covering the Cache Storage API, including `Cache.match()`, `Cache.put()`, and strategies for managing cache lifetimes (e.g., versioning).
*   **Client Communication**: Documentation on using `postMessage()` for cross-context communication, handling messages in the worker (`onmessage`), and querying clients from the main thread.
*   **Fetch Event Handling**: A tutorial on implementing network interception, using `event.respondWith()`, and managing race conditions between cached and network responses.
*   **Browser Compatibility Matrix**: A reference table summarizing support for Service Worker features across major browsers (Chrome, Firefox, Safari, Edge) and versions.

# Gaps Or Cautions

*   **Algorithm Omission**: Chunks 8 through 12 cover "Appendix A: Algorithms" but do not contain the actual algorithmic steps in the provided notes. The synthesis must rely on the high-level descriptions of behavior (e.g., state transitions) rather than specific step-by-step logic found in the full source.
*   **Security Nuance**: The specification emphasizes that Service Workers require a secure context. The `Service-Worker` header is used for threat detection, implying potential security risks if misused or if headers are spoofed.
*   **Legacy Edge Distinction**: Compatibility data distinguishes between "Edge Legacy" and modern "Edge," which can be confusing when checking version support. Implementers must ensure compatibility with the specific browser engine generation in use.
*   **Relative URL Resolution**: The `Service-Worker-Allowed` header relies on relative URLs being resolved against the script's own URL, not the base document URL. This is a specific resolution rule that differs from standard absolute path assumptions.

## group-04

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

# Group Context

This document group provides a comprehensive specification and compatibility reference for the Service Worker API. The content spans from initial motivations and architectural models to detailed execution contexts, client management, caching strategies, security constraints, extensibility mechanisms (algorithms), and finally, browser support matrices. The primary focus is on defining the lifecycle, events, registration states, and interaction patterns between service workers, their clients, and the user agent.

# Cross-Chunk Summary

The document is structured into distinct logical sections:
1.  **Motivations & Model (Chunks 01-02):** Defines the reasons for Service Workers and establishes the core model including lifetime management, events, timing, registration states (`installing`, `waiting`, `active`), and client contexts. It details how workers interact with window clients versus worker clients.
2.  **Client Context & Execution (Chunks 03-04):** Expands on the `ServiceWorkerRegistration` interface (including `update()`, `unregister()`), the global navigator scope, container APIs (`getRegistrations`, `skipWaiting`), and the execution context within the `ServiceWorkerGlobalScope`. It introduces the `Client` interface for managing active clients, including visibility state, focus management, and navigation capabilities.
3.  **Events & Caching (Chunks 05-06):** Covers extendable events (`waitUntil`, `respondWith`), specific event types (`InstallEvent`, `FetchEvent`, `MessageEvent`), and the caching architecture via `Cache` and `CacheStorage` interfaces for storing requests, responses, and managing cache lifetimes.
4.  **Security & Extensibility (Chunks 07-13):** Addresses security considerations such as secure contexts, Content Security Policy, origin relativity, CORS, and path restrictions. It outlines extensibility via algorithms and extended HTTP headers.
5.  **Appendices & Compatibility (Chunks 14-19):** Concludes with acknowledgements and detailed browser support matrices for specific API properties, version requirements across desktop and mobile engines, and notes on pending or unverified data points.

# Repeated Or Central Claims

*   **Lifecycle Management:** The Service Worker lifecycle is strictly state-based (`installing` -> `waiting` -> `active`), managed via the `ServiceWorkerRegistration` interface.
*   **Client Interoperability:** Communication between the service worker and its controlling clients is asynchronous, primarily utilizing `postMessage()` with optional transferable objects or options.
*   **Cache Architecture:** Caching is handled through a structured storage system (`CacheStorage`) allowing multiple named caches, with specific methods for adding, matching, deleting, and retrieving entries based on request metadata.
*   **Event Handling:** Service Workers operate within an execution context where events (Fetch, Install, Message) are extendable, allowing workers to pause or extend event lifecycles using `waitUntil()`.
*   **Security Model:** Service Workers require a secure context (HTTPS or localhost) and adhere to strict origin and path restrictions to prevent unauthorized access to resources or interference with navigation.

# Important Local Details

*   **API Properties:** The specification details granular properties such as `visibilityState` on clients, `focused` status, `ancestorOrigins`, and specific event attributes like `event.replacesClientId` for fetch events.
*   **Navigation Preload:** A dedicated mechanism (`navigationPreload`) exists to allow workers to signal resource preload strategies, exposing headers via `NavigationPreloadManager`.
*   **Scope & Registration:** The `scope` property defines the URL path under which the worker operates, and `updateViaCache` determines how updates are handled (cache-only vs. network-first).
*   **WindowClient Specifics:** Distinction is made between standard `Client` objects and `WindowClient` objects, particularly regarding browsing context types and frame types (`frameType`).
*   **Browser Support Nuances:** Mobile versions of browsers (iOS Safari, Android Chrome, Firefox Mobile) often have question marks in support tables, indicating potential lack of verification compared to desktop counterparts.

# Candidate Wiki Hints

*   **Service Worker Lifecycle:** Documentation on the transition between `install`, `activate`, and `fetch` phases.
*   **Client Management Guide:** A guide on using `navigator.serviceWorker.getRegistrations()` and managing active clients via `clients.matchAll()`.
*   **Caching Strategies:** Patterns for implementing offline-first architectures using `self.caches` and `CacheStorage`.
*   **Security Best Practices:** Checklist for ensuring secure context, proper origin handling, and preventing CORS issues in SW scripts.
*   **Browser Compatibility Matrix:** A dynamic table comparing feature support across Chrome, Firefox, Safari, Edge, and Opera (desktop vs. mobile).

# Gaps Or Cautions

*   **Mobile Data Uncertainty:** The compatibility section flags mobile browsers with question marks (`?`), suggesting that version data for iOS Safari, Android WebView, and mobile-specific builds may be incomplete or unverified at the time of writing.
*   **Algorithm Implementation:** Chunks 08-12 cover "Appendix A: Algorithms" but do not explicitly detail the step-by-step logic within the notes; these are references to external algorithm definitions rather than summarized procedural steps.
*   **Transferable Objects:** While `postMessage` supports transfer, the specific performance implications or serialization costs of transferring large blobs vs. strings are not elaborated beyond the API signature.
*   **Cache Storage Limits:** While cache lifetimes are mentioned, specific quota limits or eviction policies for `CacheStorage` are not detailed in this group's notes.

