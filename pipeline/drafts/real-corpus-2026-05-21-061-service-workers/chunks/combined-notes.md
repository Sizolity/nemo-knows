## chunk-01

---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

## Chunk Context
**Heading path:** Document → Service Workers → 1. Motivations, 2. Model (Service Worker state, timing)
**Line range:** 1–497
**Coverage:** Abstract, Status, Table of Contents, Motivations, and initial Model definitions including Service Worker attributes and timing structs.

## Local Summary
This chunk introduces the Service Workers specification as a W3C Candidate Recommendation Draft (April 2026). It explains that service workers are event-driven worker contexts designed to intercept network requests, enable offline functionality, and manage background tasks without relying on document lifetimes. The text details motivations for flexibility over older cache APIs, emphasizes recoverable error handling, and outlines the asynchronous, single-threaded execution model. The Model section begins by defining core attributes (state, script URL, type) and timing metadata exposed to APIs.

## Key Claims
- Service workers are generic, event-driven, time-limited script contexts running at an origin.
- They provide an event destination for network interception when other destinations do not exist or are inappropriate.
- The specification aims for maximum flexibility (procedural model) at the cost of added complexity.
- Errors must always be recoverable; update processes are designed to avoid unrecoverable states seen in Application Cache.
- Service workers may start and terminate without attached documents, resembling Chrome Event Pages.
- They can handle push notifications, background sync, cross-origin requests, and centralized data updates (e.g., geolocation).
- APIs are almost entirely asynchronous to avoid blocking document/resource loading.

## Entities And Concepts
- **Service Worker**: A web worker executing in the registering client’s origin; manages network interception and offline behavior.
- **State**: One of "parsed", "installing", "installed", "activating", "activated", "redundant".
- **Script URL**: The URL of the service worker script.
- **Type**: Either "classic" or "module"; defaults to "classic".
- **Service Worker Registration**: Contains the service worker and manages lifecycle (installing, waiting, active).
- **Timing Info**: Struct with `startTime`, `fetchEventDispatchTime`, `workerRouterEvaluationStart`, `workerCacheLookupStart`, `workerMatchedRouterSource`, `workerFinalRouterSource`.
- **Event Types**: Lifecycle (`install`, `activate`), functional (`fetch`, etc.), and `message`/`messageerror`.

## Procedures And API Details
- **Skip Waiting**: A flag to skip the waiting phase, allowing immediate activation.
- **Classic Scripts Imported Flag**: Tracks whether classic scripts have been imported.
- **Used Scripts Set**: Prunes unused resources after installation based on old worker map.
- **Event Loop Running**: Defines when a service worker is considered "running".
- **Service Worker Queue**: A parallel queue associated with the worker.
- **Termination Conditions**: No event to handle or abnormal operation (e.g., infinite loops, time limit exceeded).

## Nuance Or Contradictions
- Service workers are described as both similar to Shared Workers and distinct in that they never handle messages from documents; they process events only.
- The specification is a "living document" with unimplemented features and potential changes; citation is discouraged outside of work-in-progress context.
- Timing values are initially 0 but updated during execution; exact semantics depend on WPT results and user agent implementation.

## Candidate Wiki Hints
- **Service Worker Lifecycle**: Document states, events (install/activate), and termination policies.
- **Offline First Architecture**: How service workers intercept fetches and override default network behavior.
- **Timing API Integration**: Use of `DOMHighResTimeStamp` for performance tracing in service worker contexts.
- **Error Recovery Design**: Contrast with Application Cache; emphasis on avoiding unrecoverable states.
- **Event-Driven Model**: Comparison to Chrome Event Pages and Shared Workers; implications for resource conservation.

## chunk-02

---
title: Chunk 02 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

## Chunk Context
This chunk defines the internal structure of a **Service Worker Registration**, detailing its associated workers (installing, waiting, active), task queues, and update settings. It further explains the lifecycle of these registrations during user agent shutdown. The text then transitions to defining **Service Worker Clients** (window, worker) and how they acquire an active service worker based on origin matching and fetch routing. Finally, it covers specific **Task Sources** for events and the rules for persisting registration state across restarts. The latter half introduces the `ServiceWorker` and `ServiceWorkerRegistration` IDL interfaces, their attributes (like `scriptURL`, `scope`, `state`), and the logic for creating their corresponding objects in an environment settings map.

## Local Summary
A service worker registration is a persistent tuple containing a scope URL, a storage key, and up to three associated workers representing different lifecycle states: installing, waiting, and active. It also tracks task queues, update cache modes, and navigation preload headers. Clients (window or worker) are environments that may be controlled by an active service worker; this control is determined by origin matching rules during creation or navigation. The `ServiceWorker` and `ServiceWorkerRegistration` interfaces expose read-only attributes for these states and URLs, along with methods like `postMessage` and `update`. Object instantiation logic ensures a one-to-one mapping between the internal registration state and the exposed objects within an environment's map.

## Key Claims
- A service worker registration persists across user agent restarts but discards installing workers; waiting workers promote to active upon restart.
- Task queues for a registration back up tasks from the active worker's event loop when terminated, re-queuing them when the worker spins off.
- A window client's active service worker is null if the origin is opaque or differs from the creator document's origin during navigation (unless routed via HTTP fetch in specific ways).
- The `ServiceWorkerRegistration` interface exposes three optional attributes: `installing`, `waiting`, and `active`, each holding a `ServiceWorker?` object.
- The `postMessage` method on `ServiceWorker` serializes the message, handles transferable objects, and dispatches an event at the destination `ServiceWorkerGlobalScope`.

## Entities And Concepts
- **Service Worker Registration**: A tuple of scope URL, storage key, and worker states (installing/waiting/active).
- **Service Worker Client**: An environment that is either a window client or a worker client.
- **Window Client**: Controlled by a service worker if the origin matches and fetch is routed appropriately.
- **Worker Client**: Inherits active service worker from owner unless origins differ or data/blob URLs are used.
- **Task Sources**: `handle fetch` and `handle functional event` (e.g., push).
- **ServiceWorker State Enum**: `parsed`, `installing`, `installed`, `activating`, `activated`, `redundant`.
- **ServiceWorkerUpdateViaCache Enum**: `imports`, `all`, `none`.

## Procedures And API Details
- **Getting ServiceWorker objects**: Retrieve from the environment settings object map; create a new one if missing, associating it with the internal service worker and copying its state.
- **Getting ServiceWorkerRegistration objects**: Retrieve from the environment settings registration object map; initialize attributes to null or the corresponding `ServiceWorker` objects if the internal workers are non-null.
- **postMessage(message, options)**:
  1. Serialize message with transfer support.
  2. Check if event dispatching should be skipped.
  3. Run parallel substeps to determine source (window/client) and destination (`ServiceWorkerGlobalScope`).
  4. Deserialize transferred values.
  5. Create an `ExtendableMessageEvent` with data, ports, and origin.
  6. Dispatch the event at the destination.
- **scriptURL**: Returns the serialized script URL of the associated service worker (e.g., `https://example.com/service_worker.js`).
- **scope**: Returns the serialized scope URL (e.g., `https://example.com/`).

## Nuance Or Contradictions
- The text notes that behavior in section 2.5 is non-normative and will be specified in the HTML Standard, implying current implementation details might evolve.
- Sandboxed iframes without specific directives (`allow-same-origin`, `allow-scripts`) result in an opaque origin, causing their active service worker to be null.
- Data URLs and Blob URLs inherently result in an opaque origin for window/worker clients, leading to a null active service worker value regardless of creator context.

## Candidate Wiki Hints
- **Service Worker Registration Lifecycle**: Explain the three worker states (installing, waiting, active) and how they map to the registration object.
- **Client Control Logic**: Detail how `navigator.serviceWorker.controller` is determined for window clients versus worker clients.
- **Message Passing Protocol**: Breakdown of `postMessage` serialization, transfer handling, and event dispatching mechanics.
- **Interface Definitions**: Summarize `ServiceWorker` and `ServiceWorkerRegistration` attributes and their underlying internal counterparts.

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---
Chunk Context
This chunk documents the Service Worker API's lifecycle management methods (`update()`, `unregister()`), the `navigator.serviceWorker` and `ServiceWorkerContainer` interfaces (including registration, retrieval, and messaging capabilities), event dispatching mechanisms, navigation preload configuration, and execution context details for the `ServiceWorkerGlobalScope`. It also includes a practical code example for serving cached resources.

Local Summary
The text defines algorithms for updating and unregistering service worker registrations, specifying error conditions like "InvalidStateError" when the worker is installing or already active. It details the `ServiceWorkerContainer` API for managing registrations (`register`, `getRegistration`, `getRegistrations`) and enabling client messaging via `startMessages()`. The section on events lists specific event types dispatched to different objects based on state changes (e.g., `updatefound`, `controllerchange`). Finally, it describes the `NavigationPreloadManager` for controlling navigation preload headers and provides a JavaScript snippet demonstrating caching strategies using `caches.open()` and `fetch` event handlers.

Key Claims
- The `update()` method rejects with an "InvalidStateError" if the relevant global object's associated service worker is in the "installing" state.
- The `unregister()` method only affects subsequent navigations; it does not immediately unload currently controlled clients.
- A `ServiceWorkerContainer` is created and associated with every `Navigator` or `WorkerNavigator` object upon creation.
- The `ready` attribute on `ServiceWorkerContainer` returns a promise that resolves to the registration once an active worker is present and never rejects.
- Client messaging is disabled by default and must be enabled via the `startMessages()` method or setting the `onmessage` event handler.
- The `skipWaiting()` method allows a waiting service worker to become active immediately, bypassing the normal waiting period for clients.

Entities And Concepts
- ServiceWorkerRegistration: Interface representing a registration of a service worker script.
- ServiceWorkerContainer: Interface providing capabilities to register, unregister, and update service workers; exposes `controller` and `ready`.
- ServiceWorkerGlobalScope: The global execution context of a service worker, extending `WorkerGlobalScope`.
- NavigationPreloadManager: Interface for managing navigation preload headers (`enable`, `disable`, `setHeaderValue`).
- Clients: A collection of all clients controlled by the service worker.
- Caching API: Used in the example code to store and retrieve resources via `caches.open()` and `cache.addAll()`.

Procedures And API Details
- **update()**: Runs "Get Newest Worker" algorithm; rejects if newest worker is null or if the current worker is installing. Schedules a job to perform the update.
- **unregister()**: Creates a job with the "unregister" type to remove the registration, effective only for future navigations until all clients unload.
- **register(scriptURL, options)**: Parses the script URL (or trusted string), optionally sets scope and type ("classic"), and invokes "Start Register".
- **getRegistration(clientURL)**: Validates the client URL origin against the service worker client's origin; returns undefined if no matching registration exists.
- **skipWaiting()**: Sets the skip waiting flag and invokes "Try Activate" to progress from waiting to active state.
- **Event Handlers**: Supported handlers include `onupdatefound` (Registration), `oncontrollerchange`, `onmessage`, `onmessageerror` (Container), and `oninstall`, `onactivate`, `onfetch`, `onmessage`, `onmessageerror` (GlobalScope).

Nuance Or Contradictions
- The `ready` promise resolves eventually even if not immediately, as long as a matching registration with an active worker is found.
- `navigator.serviceWorker.controller` returns null specifically during a force refresh (Shift+Refresh) or when no active worker exists.
- Synchronous requests must not be initiated inside a service worker; the environment is strictly asynchronous.

Candidate Wiki Hints
- Service Worker Lifecycle Management: Explaining the flow from registration to activation, including `skipWaiting()` and `update()`.
- Client Messaging API: Detailing how to enable messaging via `startMessages()` or event listeners.
- Navigation Preload Control: How to use `NavigationPreloadManager` to optimize navigation performance.
- Caching Strategies in Service Workers: Best practices for `caches.open()` and handling fetch events for offline support.

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
This chunk defines the `Client` and `WindowClient` interfaces, their attributes (such as `url`, `id`, `type`, `visibilityState`, `focused`), and methods (`postMessage`, `focus`, `navigate`). It also details the `Clients` global interface with methods like `get`, `matchAll`, `openWindow`, and `claim`. The section concludes by introducing the `ExtendableEvent` interface used for service worker lifecycle events (`install` and `activate`) and extensions.

Local Summary
The chunk outlines how clients are represented within a Service Worker context, distinguishing between generic clients (workers, shared workers) and window clients (top-level browsing contexts). It specifies how to retrieve client details, send messages via `postMessage`, manage focus states, navigate to new URLs, and open windows from the service worker scope. The `Clients` interface provides mechanisms to query active clients, filtering by type or control status, and handles the association between clients and their respective service worker registrations.

Key Claims
- A `Client` object represents a service worker client with an associated frame type ("auxiliary", "top-level", "nested", "none") and a unique `id`.
- A `WindowClient` extends `Client` to include browsing context-specific properties like `visibilityState`, `focused`, and `ancestorOrigins`.
- The `postMessage` method allows communication between the service worker and its clients, supporting structured serialization with transferable objects.
- The `Clients` interface is created when a `ServiceWorkerGlobalScope` object is created and manages all associated clients.
- `matchAll` returns an array of clients sorted by focus state (focused first) and creation order, respecting the `type` filter in options.
- `claim()` allows a service worker to take control of its clients, handling unloading if the client's active service worker differs from the claiming one.

Entities And Concepts
- Client: Interface representing a service worker client.
- WindowClient: Sub-interface for clients associated with window browsing contexts.
- FrameType: Enum defining frame hierarchy ("auxiliary", "top-level", "nested", "none").
- Clients: Global interface to manage and query active clients.
- ExtendableEvent: Interface extending Event, used for service worker lifecycle events (`install`, `activate`).
- ServiceWorkerGlobalScope: Scope where the `Clients` object is instantiated.

Procedures And API Details
- **Client Attributes**:
  - `url`: Returns the serialized creation URL of the associated service worker client.
  - `type`: Returns "window", "worker", or "sharedworker" based on the client type.
- **Client Methods**:
  - `postMessage(message, options)`: Sends a message to the client, handling serialization and deserialization with transferable objects.
  - `focus()`: Requests focus for the browsing context; returns a promise resolving to the `WindowClient` if successful.
  - `navigate(url)`: Navigates the associated browsing context to a new URL.
- **Clients Interface Methods**:
  - `get(id)`: Returns a promise resolving to a specific client by ID or `undefined`.
  - `matchAll(options)`: Returns a promise resolving to an array of clients matching criteria (e.g., type, includeUncontrolled).
  - `openWindow(url)`: Opens a new window and returns a `WindowClient` if the origin matches.
  - `claim()`: Takes control of all clients associated with the service worker.
- **ExtendableEvent**:
  - `waitUntil(promise)`: Keeps the service worker alive until the promise resolves or is rejected.

Nuance Or Contradictions
- The `type` getter logic includes a fallback to "window" if the client is an environment settings object, which might seem redundant given the specific checks for window/worker/sharedworker clients.
- In `matchAll`, sorting prioritizes focused window clients over unfocused ones and separates them from worker clients regardless of creation order within their respective groups.
- The `claim()` method explicitly rejects if the service worker is not active, ensuring control can only be taken by an active worker.

Candidate Wiki Hints
- **Service Worker Client Interface**: Overview of `Client` and `WindowClient` attributes and methods.
- **Clients API**: Guide to querying clients via `get()`, `matchAll()`, `openWindow()`, and `claim()`.
- **Interacting with Clients**: Best practices for using `postMessage` and handling client types.
- **Service Worker Lifecycle**: Explanation of `ExtendableEvent` usage in `install` and `activate` events.

## chunk-05

---
title: Chunk 05 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

## Chunk Context

This chunk covers the Service Worker specification details for `ExtendableEvent` and its specific event types: `InstallEvent`, `FetchEvent`, and `ExtendableMessageEvent`. It defines their attributes, lifecycle behaviors (specifically how `waitUntil()` manages promise lifetimes), routing rules via `addRoutes()`, and response handling logic. The text concludes with the definition of the `CacheStorage` interface (`self.caches`) and guidelines on cache lifetimes.

## Local Summary

The chunk details how Service Workers extend event lifetimes using promises attached to `ExtendableEvent`. It specifies the attributes for `InstallEvent` (including router rules), `FetchEvent` (request/response/client IDs), and `ExtendableMessageEvent` (data/origin/source). It explains that adding routes via `addRoutes()` implicitly extends event lifetime. Finally, it introduces the `CacheStorage` API (`self.caches`) as a mechanism for authors to manually manage offline content caches, distinct from the browser's HTTP cache.

## Key Claims

- The `waitUntil(f)` method adds a promise to an event's "extend lifetime promises" list; if all resolve successfully, the worker is treated as installed/activated.
- `event.addRoutes(rules)` registers rules for offloading tasks and implicitly extends the event's lifetime similar to `waitUntil()`.
- `FetchEvent` attributes include `request`, `preloadResponse`, various client IDs (`clientId`, `resultingClientId`, etc.), and `handled`.
- `respondWith(r)` is used by developers to provide a response; it automatically extends the event lifetime.
- Cache instances are not part of the browser's HTTP cache and must be manually managed (populated/deleted) by authors.

## Entities And Concepts

- **ExtendableEvent**: Base interface for service worker events supporting lifetime extension via promises.
- **InstallEvent**: Lifecycle event fired when a worker registration is installing; supports `addRoutes()`.
- **FetchEvent**: Functional event for HTTP requests; supports `respondWith()` and `preloadResponse`.
- **ExtendableMessageEvent**: Legacy functional event for message passing (`data`, `origin`, `source`).
- **RouterRule / RouterCondition**: Types used in `InstallEvent` to define routing logic (URL patterns, request methods).
- **CacheStorage**: The storage object accessed via `self.caches` for managing named caches.
- **Response Object**: Returned by service workers or fetched from the network.

## Procedures And API Details

### Managing Event Lifetime (`waitUntil`)
To add a lifetime promise to an event:
1. Check if the event's `isTrusted` attribute is false; throw `InvalidStateError` if so.
2. Check if the event is active; throw `InvalidStateError` if not.
3. Add the promise to the event's list of extend lifetime promises.
4. Increment the pending promises count (even if the promise is already settled).
5. Upon promise settlement, decrement the count. If count reaches 0:
    - Try clearing the registration if unregistered.
    - Invoke `Try Activate` with the registration if not null.

### Registering Routes (`addRoutes`)
Steps to register router rules:
1. Normalize `rules` input.
2. Verify each rule's condition against the service worker.
3. Ensure `fetch` is in the event types to handle if source requires it.
4. Create a new lifetime promise and add it to the event (extending lifetime).
5. Enqueue steps to the service worker queue to update the list of router rules.

### Handling Fetch Responses (`respondWith`)
Steps when calling `respondWith(r)`:
1. Verify the event's dispatch flag is set and respond-with entered flag is unset; throw `InvalidStateError` otherwise.
2. Add `r` (response or promise) to the event lifetime promises.
3. Set flags: stop propagation, respond-with entered, wait to respond.
4. Upon rejection of `r`: Set error flag, unset wait-to-respond flag.
5. Upon fulfillment with a `Response` object:
    - If response is not a `Response` object, set error flag.
    - Otherwise, create a new stream from the response body (handling chunks, reading, closing, or erroring).
    - Set the potential response and unset wait-to-respond flag.

## Nuance Or Contradictions

- **Lifetime Extension Default**: The text notes that `addRoutes()` extends lifetime "as if" `waitUntil(promise)` was called, and similarly for `respondWith()`. This implies automatic extension without explicit `waitUntil` calls in those contexts.
- **Cache Isolation**: Caches are explicitly stated to be isolated from the browser's HTTP cache and do not update automatically; authors must version caches by name and manage them manually.
- **Promise Settlement**: The pending promises count is incremented even if a promise has already settled, with the decrement handled in a microtask upon reaction to the promise.

## Candidate Wiki Hints

- Service Worker Event Lifetime Management (`waitUntil`)
- InstallEvent Router Rules API
- FetchEvent Response Handling Guide
- ExtendableMessageEvent Attributes
- CacheStorage and Manual Cache Management

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---
Chunk Context
- Heading path: 5. Caches > 5.4. Cache
- Line range: 2371–2819
- Scope: Specification of the `Cache` and `CacheStorage` interfaces, including algorithms for match/matchAll/add/put/delete/keys, batch operation structures, security checks (cross-origin, Vary headers), and async storage management.

Local Summary
This chunk defines the API surface and internal algorithmic steps for caching requests/responses in service workers. It distinguishes between the `Cache` interface (per-cache operations) and `CacheStorage` (global cache map management). Operations are largely asynchronous and rely on a "batch operation" struct to serialize state changes. Strict validation occurs on schemes (`http`/`https`), methods (`GET`), status codes, and `Vary` header contents.

Key Claims
- A `Cache` object represents a request/response list shared across documents/workers.
- `match()` returns the first matching response or `undefined`; `matchAll()` returns all matches as a frozen array of `Response` objects.
- Batch operations (`put`, `delete`) are serialized via an internal job queue; failures reject the operation promise with specific errors (e.g., `TypeError`, `QuotaExceededError`).
- Only `GET` requests over `http`/`https` can be cached; other methods or schemes cause immediate rejection.
- Responses with a `Vary` header containing `*` are rejected to prevent invalid caching policies.
- `CacheStorage` behaves like an async map, explicitly excluding synchronous iteration methods (`forEach`, etc.) pending TC39 async iteration standards.
- After deleting a cache entry via `CacheStorage.delete()`, existing DOM objects referencing the old entries remain functional until garbage collected or manually cleaned.

Entities And Concepts
- **Cache**: Interface for operations on a single named cache (match, add, put, delete).
- **CacheStorage**: Interface managing the global map of named caches (`open`, `delete`, `keys`).
- **Batch Cache Operations**: Internal struct used to serialize state changes (type, request, response, options).
- **RequestInfo / Request**: Inputs for cache operations; strings are converted via the `Request` constructor.
- **CacheQueryOptions / MultiCacheQueryOptions**: Options dictionaries controlling query behavior (`ignoreSearch`, `ignoreMethod`, `ignoreVary`, `cacheName`).
- **FrozenArray**: Return type for `matchAll()` and `keys()` to ensure immutability of results.
- **QuotaExceededError**: Error thrown when opening a cache exceeds storage limits.

Procedures And API Details
- **match(request, options)**: Returns first match; resolves with `undefined` if none found. Handles non-GET requests based on `ignoreMethod`.
- **matchAll(request, options)**: Returns all matches in a frozen array; rejects cross-origin blocked resources.
- **add(request)**: Equivalent to fetching the request and storing its response (if available); returns `undefined`.
- **put(request, response)**: Stores a specific response. Requires full body read (locking), validates status (not 206), and checks `Vary` headers. Returns upon completion of the async job.
- **delete(request, options)**: Removes matching entries. Returns `true` if at least one entry was deleted, `false` otherwise.
- **keys(request, options)**: Returns a frozen array of `Request` objects found in the cache (or empty array if none).
- **open(cacheName)**: Creates or retrieves a `Cache` object. Throws `QuotaExceededError` on storage limit breach.
- **Batch Cache Operations**: Internal step invoked by mutating methods to serialize writes/deletes and handle errors uniformly.

Nuance Or Contradictions
- **Synchronous vs Async Map**: The spec explicitly notes that `CacheStorage` conforms to an async map pattern, intentionally omitting standard synchronous iteration methods found in ECMAScript 6 Maps.
- **Body Handling**: While optimizations like streaming directly to disk are noted as possible, the reference implementation requires reading all bytes into memory (`clonedResponse`) before committing to ensure body locking and consistency.
- **Cross-Origin Policy**: `match()` checks cross-origin resource policies; if blocked, it rejects with a `TypeError`. This check happens on opaque responses during resolution.
- **Post-Delete Functionality**: The spec clarifies that deleting a cache name does not immediately invalidate existing DOM objects referencing the old data; they remain functional until GC or manual intervention.

Candidate Wiki Hints
- Page: `ServiceWorker Cache API Overview` (Summary of `Cache` vs `CacheStorage`).
- Page: `ServiceWorker Caching Strategies` (Using `match`, `matchAll` for retrieval).
- Page: `ServiceWorker Storage Limits` (Handling `QuotaExceededError` and `open()` failures).
- Page: `ServiceWorker Cache Security` (`Vary` header restrictions, cross-origin checks).

## chunk-07

---
title: Chunk 07 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---
Chunk Context
This chunk covers Section 6 (Security Considerations) and Section 7 (Extensibility) of the Service Workers specification. It details requirements for secure contexts, Content Security Policy enforcement, origin relativity restrictions, path limitations, CORS handling for cached resources, header requirements for script requests, implementer concerns regarding plugins and legacy code, privacy implications of persistent storage, and mechanisms for extending the API via functional events.

Local Summary
Service workers are restricted to secure contexts (typically HTTPS) and must execute within the registering client's origin. They enforce Content Security Policies on their scripts and restrict access to CDN-hosted resources unless specific CORS headers are present. Path restrictions prevent a single script from serving multiple unrelated scopes, though this can be overridden via `Service-Worker-Allowed` headers. The spec mandates headers like `Service-Worker` and specific MIME types for scripts. Privacy is addressed by requiring the clearing of persistent storage maps when users clear their data. Extensibility is supported through partial interface definitions for registrations, functional events, and event handlers.

Key Claims
- Service workers must execute in secure contexts; clients must also be secure contexts to register or interact with them.
- `localhost`, `127.0.0.0/8`, and `::1/128` are exceptions for development purposes.
- Content Security Policy (CSP) headers on the script resource are enforced or monitored by the user agent.
- Service workers execute in the registering client's origin, preventing hosting on CDNs directly.
- Resources from other origins can be fetched and cached only if appropriate CORS headers are set; stored responses are either CORS filtered or opaque filtered.
- Path restrictions limit a service worker script to its specific scope path unless overridden by `Service-Worker-Allowed`.
- Service worker scripts must include the `Service-Worker` header and be served with a JavaScript MIME type.
- Plugins should not load via service workers because the embedding worker cannot handle their security origins.
- Legacy networking stack code may require auditing for interactions with service workers.
- Persistent storages (registration map, cache name to map, script resource map) must be cleared when users purge data.
- Specifications can extend the API using partial interface definitions on `ServiceWorkerRegistration`, `ExtendableEvent`, and `ServiceWorkerGlobalScope`.

Entities And Concepts
- Secure Context
- Content Security Policy (CSP) / Content-Security-Policy-Report-Only
- Origin Relativity
- Service Worker Registration Scope
- Import Scripts
- Caches API
- CORS Filtered Response / Opaque Filtered Response
- Path Restriction
- Service-Worker-Allowed Header
- Service-Worker Header
- Persistent Storage (Registration Map, Cache Map, Script Resource Map)
- Extensibility
- Partial Interface Definition
- Functional Event
- ExtendableEvent

Procedures And API Details
- **Run Service Worker Algorithm**: Enforces CSP if `Content-Security-Policy` header matches policy; monitors if `Content-Security-Policy-Report-Only` matches.
- **importScripts(urls)**:
  - Checks worker state ("parsed" or "installing").
  - Validates script resource map against URL.
  - Sets service-workers mode to "none".
  - Determines cache mode based on registration update mode, force bypass flag, or staleness.
  - Fetches request and updates the response map if safe.
- **Event.respondWith(r)**: Can accept Response objects whose corresponding responses are basic filtered, CORS filtered, or opaque filtered, but cannot create them programmatically.
- **Fire Functional Event**: Invoked by specifications to dispatch a functional event to the active worker of a service worker registration.

Nuance Or Contradictions
- While path restriction offers some protection for multi-user content on the same origin, origins are the only hard security boundary; sites should use different origins for secure isolation.
- Unlike same-origin resources managed in Cache as basic filtered responses, off-origin resources stored in Cache are CORS or opaque filtered and cannot be meaningfully created programmatically.
- Plugins are explicitly discouraged from loading via service workers due to origin handling limitations in the Handle Fetch algorithm.

Candidate Wiki Hints
- Service Worker Security Requirements
- Secure Contexts and HTTPS
- Content Security Policy for Service Workers
- Origin Relativity and CDN Restrictions
- CORS and Cached Off-Origin Resources
- Path Restriction and Service-Worker-Allowed
- Privacy and Persistent Storage Clearing
- Extending the Service Worker API

## chunk-08

---
title: Chunk 08 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---
Chunk Context
- Source Path: `raw/web/corpus-2026-05-18/061-service-workers.md`
- Heading: 7. Extensibility > Appendix A: Algorithms
- Range: Lines 3014–3350

Local Summary
This chunk defines the core internal data structures and algorithms for managing Service Worker registrations, updates, and unregistrations. It details the lifecycle of a "job" (the abstract unit of work), including its creation, scheduling, execution, and promise resolution/rejection. The text also outlines validation rules for script URLs and scope URLs, distinguishing between "classic" and "module" worker types during the fetch phase.

Key Claims
- A **job** is an abstraction representing a request to register, update, or unregister a service worker registration.
- Two jobs are considered **equivalent** if they share specific attributes (type, scope URL, script URL, worker type, cache mode) depending on whether they are register/update or unregister operations.
- The **scope URL** and **script URL** fragments are explicitly discarded by the user agent; only the scheme and path matter for identification.
- Security checks enforce that the script origin, referrer origin, and scope origin must all be "same origin" to prevent cross-origin attacks.
- If a registration already exists and matches the new job's parameters exactly (for register/update), the existing registration is returned immediately without re-fetching.

Entities And Concepts
- **Job**: An internal data structure containing type, storage key, scope URL, script URL, worker type, cache mode, client, referrer, promise, and status flags.
- **Job Queue**: A thread-safe queue used to synchronize concurrent jobs per scope URL.
- **Registration Map**: An ordered map linking (storage key, serialized scope URL) pairs to service worker registrations.
- **Worker Type**: Either `"classic"` or `"module"`, affecting how the script graph is fetched.
- **Bad Import Script Response**: Defined as an error type response, non-ok status, or a MIME type that is not JavaScript.

Procedures And API Details
- **Create Job**: Constructs a new job object by setting its attributes (type, keys, URLs, promise, client) and initializing referrer if a client exists.
- **Schedule Job**: Determines the appropriate job queue for the scope URL. If the queue is empty, it triggers execution (`Run Job`). If not empty but equivalent to the last pending job, it appends to that job's list of equivalents; otherwise, it enqueues.
- **Run Job**: Queues a task to execute `Register`, `Update`, or `Unregister` in parallel with the DOM manipulation task source. Execution is delayed until after a `DOMContentLoaded` event.
- **Finish Job**: Dequeues the completed job from its queue and triggers `Run Job` if the queue is not empty.
- **Resolve Job Promise**: Resolves the job's promise (and equivalent jobs' promises) with either a service worker registration object or the provided value, queued on the client's event loop.
- **Start Register**: Validates script URL and scope URL (rejecting non-HTTP/HTTPS schemes, invalid characters like `%2f` or `%5c`, and null values). It creates a job and schedules it.
- **Register Algorithm**: Performs security checks (origin trust, same-origin policy). If valid, it checks for an existing matching registration; if found, resolves the promise with the existing registration. Otherwise, it sets up the new registration via `Set Registration` and calls `Update`.
- **Update Algorithm**: Retrieves the existing registration. If null or if a newer worker exists with a different script URL, it rejects the promise. It then proceeds to fetch resources based on worker type.

Nuance Or Contradictions
- **Fragment Handling**: The specification explicitly notes that fragments (hashes) in both script URLs and scope URLs are set to `null` and have no effect on identification. This prevents confusion where different hashes point to the same logical resource.
- **Equivalence Logic**: Equivalence is defined differently for register/update jobs (requiring all attributes to match) versus unregister jobs (requiring only scope URL match).
- **Parallel Execution**: The `Run Job` step uses "in parallel" when invoking sub-algorithms, implying these operations can occur concurrently with other tasks on the event loop.

Candidate Wiki Hints
- Service Worker Job Lifecycle and State Machine
- Understanding Equivalent Jobs in Service Workers
- Script URL and Scope URL Validation Rules

## chunk-09

---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
This chunk details the "Appendix A: Algorithms" section of the Service Workers specification. It focuses on the `fetch` hook implementation for updating scripts, the logic for determining if resources have been updated (`hasUpdatedResources`), and the sequential algorithms governing the lifecycle of a service worker: **Install**, **Activate**, and **Try Activate**. It also introduces the `Setup ServiceWorkerGlobalScope` algorithm, which handles the creation of the execution context (realm, agent, settings object) required for security checks like CSP.

Local Summary
The text describes how the user agent fetches scripts via a specific hook to handle the unique processing model of service workers. The process involves appending headers (`Service-Worker/script`), managing cache modes, and validating MIME types. If the main script or imported scripts change (byte-for-byte comparison), the worker is marked as updated. The chunk then outlines the state machine transitions:
1.  **Install**: Runs when a new script is fetched and validated. It updates internal states, queues an `updatefound` event, and eventually dispatches an `install` event after waiting for asynchronous extensions (like `extend lifetime promises`). If installation fails or is skipped, the worker becomes redundant.
2.  **Activate**: Transitions the worker from a "waiting" state to "activating" and finally "activated". It terminates the previous active worker if necessary, matches clients to the new registration, and dispatches an `activate` event.
3.  **Try Activate**: A helper algorithm that checks conditions (e.g., existing active worker is null or has no pending events) before invoking `Activate`.
4.  **Setup ServiceWorkerGlobalScope**: Ensures the necessary global scope object exists for security checks, creating a new realm and agent if needed.

Key Claims
-   Service workers use a separate script fetching mechanism (`Update algorithm`) compared to other web workers, requiring a specific environment settings object approach.
-   The `fetch` hook for service workers must append `Service-Worker/script` headers and set the redirect mode to "error".
-   A worker is considered updated if the main script URL changes, the MIME type differs, or the body is not byte-for-byte identical. Imported scripts are checked separately; bad import responses are ignored for the update check but populate the cache.
-   The `Install` algorithm waits for all `extend lifetime promises` associated with the `install` event to settle before proceeding to activation logic.
-   Activation occurs only if the waiting worker is valid and the current active worker is either null, redundant, or has no pending events/clients using it.
-   The `Setup ServiceWorkerGlobalScope` algorithm ensures a `ServiceWorkerGlobalScope` object exists for CSP checks before any security validation occurs.

Entities And Concepts
-   **Service Worker Registration**: Manages the state (installing, waiting, active) and maps workers to storage keys/scopes.
-   **Job Promise**: Used to coordinate asynchronous operations during the fetch/update cycle; can be rejected with "SecurityError" or "TypeError".
-   **Environment Settings Object**: The execution context for service worker algorithms, distinct from classic workers.
-   **Realm**: The isolation boundary containing the global object and agent.
-   **ServiceWorkerGlobalScope**: The global scope object returned by `Setup ServiceWorkerGlobalScope`, used for security checks.
-   **Extend Lifetime Promises**: Asynchronous extensions that can delay the installation process; their settlement is awaited before activation logic proceeds.
-   **Update Via Cache Mode**: Determines how updates are fetched (e.g., "all", "none").

Procedures And API Details
-   **Fetch Hook Steps**:
    -   Append `Service-Worker/script` header.
    -   Set cache mode to "no-cache" if registration is stale or force bypass is set.
    -   Set redirect mode to "error".
    -   Fetch request; on response, extract MIME type and headers (`Service-Worker-Allowed`).
    -   Validate scope permissions against `maxScopeString`.
    -   Check byte-for-byte identity of main script and imported scripts.
-   **Install Algorithm**:
    -   Update registration state to "installing".
    -   Resolve job promise with the worker registration.
    -   Fire `updatefound` event on registration objects.
    -   Run `Run Service Worker` (awaiting async extensions).
    -   If successful, dispatch `InstallEvent`.
    -   Terminate waiting worker and transition to "waiting" state for the new worker.
-   **Activate Algorithm**:
    -   Terminate active worker if present.
    -   Transition states: active -> redundant, waiting -> null, installing -> active.
    -   Match clients by creation URL and scope.
    -   Resolve `readyPromise` for clients.
    -   Dispatch `ActivateEvent`.
-   **Try Activate**: Checks if activation can proceed (no active worker or active worker is idle/no clients).

Nuance Or Contradictions
-   **Bad Import Responses**: The spec explicitly states that bad responses for `importScripts()` are ignored for the purpose of the byte-to-byte update check. Only good responses from the incumbent and potential update workers count. This prevents minor fetch errors from triggering unnecessary re-installs.
-   **Activation Robustness**: Activation handlers are designed to do non-essential work (like cleanup) because they may not complete if the browser terminates during activation. The worker must function properly even if handlers fail.
-   **Implementation Flexibility**: While the spec defines a minimal setup for `Setup ServiceWorkerGlobalScope`, implementations can optimize this step as long as the observable behavior (results of security checks like CSP) remains equivalent.

Candidate Wiki Hints
-   **Service Worker Lifecycle**: Create a page explaining the Install -> Activate cycle, detailing state transitions and event dispatching.
-   **Update Strategies**: Document how `fetch` caching modes ("all", "none") interact with the `Service-Worker` header and stale checks.
-   **Security Scope**: Explain the role of `ServiceWorker-Allowed` headers and scope validation in preventing unauthorized updates.
-   **Global Scope Setup**: A technical note on why `Setup ServiceWorkerGlobalScope` is necessary for Content Security Policy (CSP) enforcement in service workers.

## chunk-10

---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---
Chunk Context
- **Heading**: 7. Extensibility > Appendix A: Algorithms
- **Coverage**: The chunk details algorithms for Service Worker lifecycle management, specifically the `Run Service Worker`, `Terminate Service Worker`, and `Handle Fetch` procedures, along with supporting concepts like module maps, policy containers, and fetch routing logic.

Local Summary
This section outlines the core execution model for Service Workers. It defines how a worker is initialized (including handling classic vs. module scripts), how it handles incoming network requests (`Handle Fetch`), how it manages cache interactions and router rules, and the specific steps taken when terminating the worker or clearing its event loop tasks.

Key Claims
- A Service Worker's execution context is established by creating a `WorkerGlobalScope` object, which holds the script URL, policy container, and type.
- The `Run Service Worker` algorithm ensures the worker is running before returning control; if initialization fails (e.g., blocked CSP), it returns failure.
- Fetch handling (`Handle Fetch`) involves matching registrations based on storage keys, evaluating router rules (cache vs. network vs. fetch-event), and potentially triggering soft updates for stale workers.
- The `Terminate Service Worker` algorithm sets a closing flag, backs up functional events (fetch/push) to the registration's task queues, and discards other tasks like message events.

Entities And Concepts
- **WorkerGlobalScope**: The global object representing the service worker environment; stores the module map, policy container, and URL.
- **Service Worker Registration**: The object linking a script to its clients; manages active workers, router rules, and storage keys.
- **Router Source**: A classification for fetch handling outcomes (e.g., "cache", "network", "fetch-event").
- **Soft Update**: An algorithm run in parallel when cache or network routes are selected to ensure the worker is current before responding.
- **Race Response/Result**: Mechanisms used to coordinate between network requests and fetch event listeners, ensuring one takes precedence or both are handled appropriately.

Procedures And API Details
- **Run Service Worker**:
  - Sets `workerGlobalScope` properties (url, policy container, type).
  - Evaluates the script (classic or module); aborts if evaluation fails or is rejected.
  - Queues tasks from the registration's task queues to the worker's event loop.
  - Runs the responsible event loop until destruction.
- **Terminate Service Worker**:
  - Sets `closing` flag on the global object.
  - Clears extended events and aborts the current script.
  - Moves fetch/functional event tasks to the registration queue; discards message event tasks.
- **Handle Fetch**:
  - Determines if a worker should handle the request based on client security context and destination.
  - Matches a registration using `obtain a storage key` and `Match Service Worker Registration`.
  - Evaluates router rules: checks cache, then potentially races network with fetch handler for GET requests.
  - Creates a `Service-Worker-Navigation-Preload` header if navigation preload is enabled.

Nuance Or Contradictions
- **Empty Fetch Listeners**: The spec notes that empty function bodies in `fetch` listeners (e.g., `() => {}`) are used by some sites to signal PWA status but may have performance implications.
- **Task Queue Prioritization**: During termination, only fetch and functional events are backed up; other tasks (like messages) are discarded without processing.
- **Global Object Creation Timing**: The spec notes that `ServiceWorkerGlobalScope` might be created during cache lookups for CORS checks, though implementations may not always create it there as expected.

Candidate Wiki Hints
- Service Worker Initialization and Lifecycle
- Fetch Event Routing and Cache Strategies
- Managing Service Worker Task Queues

## chunk-11

---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
This chunk covers **Appendix A: Algorithms** within the Service Workers specification (Section 7. Extensibility). It details the low-level algorithms governing request handling, event dispatching, router condition matching, registration lifecycle management, and cleanup procedures during shutdown or client unloading.

Local Summary
The text defines a comprehensive set of algorithms for service worker execution flow. Key areas include the main fetch handler logic (handling race conditions, soft updates, and error propagation), URL pattern parsing for routing, complex nested router condition verification (supporting `or`, `not`, and status checks), and lifecycle management functions like clearing registrations and handling user agent shutdowns.

Key Claims
- **Soft Updates:** The spec introduces "soft update" logic to refresh stale registrations or workers under specific conditions (e.g., non-subresource requests on stale registrations).
- **Event Skipping:** To avoid performance delays, the UA may skip event dispatch if no listeners exist deterministically added during the first script execution.
- **Router Limits:** Router condition complexity is constrained: total nested conditions cannot exceed 1024, and nesting depth is limited to 10 levels to prevent exponential computation.
- **Security on Regex:** User-defined regular expressions in URL patterns are prohibited due to security concerns; if a pattern has regexp groups, the router condition verification fails.
- **Body Cancellation:** If a fetch response is null but the request body source is null and the body is unusable, the body is cancelled with `undefined`.

Entities And Concepts
- **Service Worker State:** Workers transition between states like "activating", "activated", "redundant".
- **FetchEvent Attributes:** Includes `respondWith`, `waitUntil` (implied via flags), `clientId`, `replacesClientId`, and `preloadResponse`.
- **Router Conditions:** Structured objects supporting `urlPattern`, `requestMethod`, `requestMode`, `runningStatus`, `_or`, and `not` operators.
- **Registration Map:** A storage mechanism mapping scope/storage keys to service worker registration objects.
- **Extended Events Set:** A collection tracking active extended events for a specific worker, cleaned up upon dispatch or inactivity.

Procedures And API Details
- **Parse URL Pattern:** Converts a raw pattern string into a `URLPattern` object relative to the service worker's script URL. Throws exceptions if invalid.
- **Verify Router Condition:** Recursively validates router conditions. Returns `false` if regex groups exist, forbidden methods are used, or nested depth/counts exceed limits.
- **Match Router Condition:** Evaluates if a request matches a specific condition. Supports short-circuit logic for `or` (match any) and `not` (invert match).
- **Count Router Inner Conditions:** Helper algorithm to track recursion depth and total condition count against the 1024/10 limits.
- **Fire Functional Event:** Creates, initializes (optional), and dispatches an event on the active worker's global object, handling stale registration updates in parallel.
- **Clear Registration:** Terminates installing, waiting, or active workers associated with a registration and updates their state to "redundant" or clears the registration entry.

Nuance Or Contradictions
- **Mutual Exclusivity of Router Logic:** The spec explicitly states that `_or` and `not` conditions are mutually exclusive with other conditions (like `urlPattern`) for ease of understanding, enforced by returning `false` if combined improperly.
- **Abort Handling:** If a fetch controller is terminated/aborted, the algorithm deserializes the abort reason and signals it via an AbortController; if the task handling this is discarded, it defaults to treating it as a fetch failure (`handleFetchFailed`).
- **Stale Worker Handling:** When skipping events or handling failures on stale workers, the spec mandates running "Soft Update" in parallel rather than blocking.

Candidate Wiki Hints
- Service Worker Router Architecture
- Fetch Event Lifecycle and Race Responses
- Soft Update Mechanism
- URL Pattern Security Constraints

## chunk-12

---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---
Chunk Context
This chunk details the Service Worker specification's "Appendix A: Algorithms," defining state transition logic for registrations, worker lifecycle management, client resolution, cache querying, and race condition handling. It covers atomic operations for registration matching and batched cache modifications with rollback mechanisms.

Local Summary
The document outlines specific algorithms for managing the `ServiceWorkerRegistration` lifecycle, including clearing registrations when no workers are active, updating worker states (installing/waiting/active), and notifying controllers of changes. It defines logic for matching clients to scopes via prefix-based URL comparison, retrieving the newest worker, and resolving client promises with security checks. The chunk also specifies cache operations like `Query Cache` and `Batch Cache Operations`, which include strict validation for HTTP schemes and methods, as well as race response lookup mechanisms.

Key Claims
- A registration is cleared only if its installing, waiting, and active workers are null or have no pending events.
- URL string matching for service worker scopes is prefix-based rather than path-structural, though it maintains same-origin security.
- `Batch Cache Operations` perform atomic writes; if an exception occurs (e.g., quota exceeded), the implementation rolls back all changes made during the batch job.
- `Resolve Get Client Promise` rejects with a "SecurityError" DOMException if the client is not in a secure context or has an untrustworthy creation URL.

Entities And Concepts
- ServiceWorkerRegistration
- ServiceWorker (states: installing, waiting, active, parsed)
- WindowClient / Client object
- Cache API (CacheQueryOptions, CacheBatchOperation)
- DOMException ("SecurityError", "InvalidStateError", "QuotaExceededError")
- Race Response Map

Procedures And API Details
- **Clear Registration**: Validates that no workers are using the registration before clearing state.
- **Update Registration State**: Sets the `installing`, `waiting`, or `active` worker on a registration and updates associated objects via queued tasks.
- **Match Service Worker Registration**: Uses prefix matching to find a registration based on storage key and client URL origin.
- **Batch Cache Operations**: Validates operation types ("delete", "put"), checks for existing matches before deletion, ensures HTTP/HTTPS schemes for PUTs, and handles rollback on failure.
- **Resolve Get Client Promise**: Checks security context and creation URL; constructs `Client` or `WindowClient` objects based on frame type and visibility state.

Nuance Or Contradictions
- The specification notes that "parsed" is the initial state but asserts a service worker is never updated to this state, implying it is an internal initialization marker rather than a runtime lifecycle stage.
- URL matching for scopes treats `https://example.com/prefix-of/resource.html` as matching a scope of `https://example.com/prefix`, relying on trailing slashes in serialized URLs for safety.

Candidate Wiki Hints
- **Service Worker Lifecycle**: Explaining the states and transitions between installing, waiting, and active workers.
- **Cache API Security**: Detailing the security checks performed when resolving client promises and interacting with caches.
- **Batched Cache Operations**: How to perform atomic cache updates safely with rollback capabilities.
- **Scope Matching Logic**: The specific prefix-based algorithm used to associate clients with service worker scopes.

## chunk-13

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

## chunk-14

---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
Section 8 (Acknowledgements) and the subsequent "Conformance" preamble, covering lines 5309–5393. The chunk lists individuals who contributed to the specification's development via workshops, design discussions, feedback, and tooling. It then defines how conformance requirements are expressed (using RFC 2119 terminology), distinguishes normative text from examples/notes, and clarifies that imperative algorithm steps inherit their modality ("must", "should") from the surrounding context.

Local Summary
The chunk acknowledges contributors to the service worker specification and outlines the document's conventions for conformance: normative vs. informative content, interpretation of RFC 2119 key words, and how algorithmic imperatives map to modal verbs.

Key Claims
- Andrew Betts organized a workshop that advanced the work; EdgeConf sessions on "Offline" created connections enabling progress.
- Anne van Kesteren's prior work on URLs, HTTP Fetch, Promises, and DOM is foundational; Ian Hickson's Web Worker spec is also essential.
- A long list of individuals provided design guidance and discussion (e.g., Domenic Denicola, Jake Archibald, etc.).
- Jason Weber, Chris Wilson, Paul Kinlan, Ehsan Akhgari, and Daniel Austin gave well-timed feedback on requirements and standardization.
- Dimitri Glazkov's scripts and formatting tools were essential for producing the specification.
- Professional support was provided by Vivian Cromwell, Greg Simon, Alex Komoroske, Wonsuk Lee, and Seojin Kim.
- Conformance requirements use RFC 2119 key words ("MUST", "SHOULD", etc.), though they appear in mixed case for readability.
- All text is normative except sections explicitly marked non-normative, examples, and notes.
- Examples are introduced with "for example" or set apart via class="example".
- Informative notes begin with "Note" and use class="note".
- Imperative steps in algorithms inherit their modality from the introducing key word; implementations may optimize beyond the prescribed algorithms.

Entities And Concepts
- RFC 2119 (key words: MUST, MUST NOT, REQUIRED, SHALL, SHALL NOT, SHOULD, SHOULD NOT, RECOMMENDED, MAY, OPTIONAL)
- Normative text vs. informative examples/notes
- Class attributes: class="example", class="note"
- Conformance algorithms (intended for clarity, not performance; implementers encouraged to optimize)

Procedures And API Details
None in this chunk.

Nuance Or Contradictions
Readability convention: RFC 2119 key words are written in mixed case rather than all caps throughout the specification. Algorithms are defined for ease of understanding and equivalence, not for performance; implementers should optimize as needed.

Candidate Wiki Hints
- Page: Service Worker Conformance Model (normative text, examples, notes, algorithmic modality)
- Topic: Contributors to the Service Worker Specification (acknowledgements summary)

## chunk-15

---
title: Chunk 15 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---
## Chunk Context
The chunk titled "8. Acknowledgements" serves as a comprehensive index or glossary, listing terms defined by reference throughout the Service Worker specification. It covers lifecycle states (e.g., `active`, `waiting`, `redundant`), event types (`activate`, `fetch`, `install`), API methods (`waitUntil`, `postMessage`, `skipWaiting`), and internal concepts (e.g., `race response`, `job queue`).

## Local Summary
This section acts as a lexical reference for the Service Worker API, enumerating attributes, events, interfaces, and algorithmic steps. It details the relationship between clients (`WindowClient`, `DedicatedWorkerClient`), caches (`CacheStorage`, `matchAll`), and service worker registrations (`ServiceWorkerRegistration`). Key concepts include the handling of fetch requests, navigation preload headers, and the management of extended events.

## Key Claims
- Service Workers operate within specific lifecycle states (`running`, `not-running`) managed by a registration object.
- Fetch events are dispatched to handlers which can respond using `respondWith()` or rely on network responses via race conditions.
- Communication between clients and workers occurs via `postMessage()`, supporting structured data and transferable objects.
- Cache operations include adding (`addAll`), deleting, and querying items, with support for versioning strategies like `ServiceWorkerUpdateViaCache`.

## Entities And Concepts
- **Interfaces**: `ServiceWorkerRegistration`, `ServiceWorkerGlobalScope`, `Client`, `CacheStorage`, `FetchEvent`, `ExtendableMessageEvent`.
- **States**: `waiting`, `active`, `redundant`, `unregistered`.
- **Events**: `activate`, `fetch`, `install`, `messageerror`, `statechange`.
- **Clients**: `WindowClient`, `DedicatedWorkerClient`, `SharedWorkerClient`.
- **Router Concepts**: `RouterRule`, `RouterCondition`, `RouterSourceDict`.

## Procedures And API Details
- **Registration**: `register(scriptURL, options)`, `getRegistration()`, `unregister()`.
- **Caching**: `open(cacheName)`, `match(request)`, `put(request, response)`, `delete(request)`.
- **Communication**: `postMessage(message, transfer)`, `onmessage` handler.
- **Lifecycle Control**: `skipWaiting()`, `waitUntil(f)`, `activate`, `install`, `fetch`.
- **Error Handling**: `messageerror` event, `handle fetch task source`.

## Nuance Or Contradictions
The text distinguishes between "active" (currently controlling the scope) and "waiting" (installed but awaiting activation). The specification also notes distinctions between "classic scripts imported flag" and standard import mechanisms, suggesting an evolution in how worker scripts are handled.

## Candidate Wiki Hints
- **Service Worker Lifecycle**: Explaining `waiting` vs `active` states and the `skipWaiting()` method.
- **Client Management**: Overview of `WindowClient`, `DedicatedWorkerClient`, and the `Clients` API (`matchAll`).
- **Cache Storage API**: Detailed guide on `CacheStorage`, `put`, `delete`, and `matchAll`.
- **Event Handling**: Deep dive into `ExtendableEvent`, `FetchEvent`, and `InstallEvent`.

## chunk-16

---
title: Chunk 16 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

## Chunk Context

This chunk, titled "8. Acknowledgements," serves as a comprehensive index of terms defined by various web standards (e.g., [CSP-NEXT], [FETCH], [HTML], [WEBIDL]) relevant to the Service Workers specification. It concludes with an IDL Index defining interfaces like `ServiceWorker`, `ServiceWorkerRegistration`, `ServiceWorkerContainer`, and event handlers such as `oninstall` and `onfetch`.

## Local Summary

The section lists specific terminology adopted from external specifications, categorizing them by their source standard (e.g., DOM, Fetch, ECMAScript). Following the term list, it provides normative references to these standards. Finally, it details the Interface Definitions Language (IDL) for key Service Worker APIs, including lifecycle states (`ServiceWorkerState`), registration management, client querying, and event handling within the `ServiceWorkerGlobalScope`.

## Key Claims

- The specification incorporates terms from diverse standards including [CSP-NEXT], [DOM], [ECMASCRIPT], [FETCH], [HTML], [INFRA], [STREAMS], [URL], and [WEBIDL].
- The `ServiceWorker` interface extends `EventTarget` and includes attributes like `scriptURL`, `state`, and event handlers such as `onstatechange`.
- The `ServiceWorkerRegistration` interface manages the lifecycle of a service worker, exposing properties like `installing`, `waiting`, `active`, and methods like `update()` and `unregister()`.
- Service Workers are exposed via the `serviceWorker` attribute on both `Navigator` and `WorkerNavigator`.
- The `ServiceWorkerContainer` interface allows for registration (`register`), retrieval of registrations, and listening to controller changes.
- The `NavigationPreloadManager` enables control over navigation preload headers via methods like `enable()`, `disable()`, and `setHeaderValue()`.
- Client management is handled by the `Clients` interface, which supports querying clients (`get`, `matchAll`) and opening windows (`openWindow`).
- Events within the service worker environment are extended (e.g., `ExtendableEvent`, `InstallEvent`, `FetchEvent`) to support lifecycle hooks like `waitUntil`.

## Entities And Concepts

- **Service Worker Lifecycle States**: `parsed`, `installing`, `installed`, `activating`, `activated`, `redundant` (defined in `ServiceWorkerState`).
- **Registration Update Cache**: Options for caching updates: `"imports"`, `"all"`, `"none"` (defined in `ServiceWorkerUpdateViaCache`).
- **Client Types**: Distinguish between `"window"`, `"worker"`, `"sharedworker"`, and `"all"`.
- **Frame Types**: Categorizes clients as `"auxiliary"`, `"top-level"`, `"nested"`, or `"none"`.
- **Running Status**: Indicates whether a service worker is `"running"` or `"not-running"`.
- **Router Sources**: Defines routing sources such as `"cache"`, `"fetch-event"`, `"network"`, and `"race-network-and-fetch-handler"`.
- **Navigation Preload**: A mechanism to preload navigation resources, managed by `NavigationPreloadManager`.

## Procedures And API Details

### ServiceWorker Registration
To register a service worker:
```webidl
Promise<ServiceWorkerRegistration> register((TrustedScriptURL or USVString) scriptURL, optional RegistrationOptions options = {});
```
- **Parameters**: `scriptURL` (the location of the worker script), `options` (containing `scope`, `type`, and `updateViaCache`).

### Client Querying
To query available clients:
```webidl
Promise<(Client or undefined)> get(DOMString id);
Promise<FrozenArray<Client>> matchAll(optional ClientQueryOptions options = {});
```
- **Options**: `includeUncontrolled` (boolean), `type` (e.g., `"window"`).

### Event Handling in Global Scope
The `ServiceWorkerGlobalScope` exposes event handlers:
- `oninstall`: Called when the service worker is being installed.
- `onactivate`: Called when a new service worker becomes active.
- `onfetch`: Called to handle network requests.
- `onmessage`: Called to handle messages from clients.

### Fetch Event Handling
```webidl
interface FetchEvent : ExtendableEvent {
  [SameObject] readonly attribute Request request;
  readonly attribute Promise<any> preloadResponse;
  DOMString clientId;
  DOMString resultingClientId;
  DOMString replacesClientId;
  Promise<undefined> handled;

  undefined respondWith(Promise<Response> r);
};
```
- **Methods**: `respondWith()` allows the handler to respond to the request.
- **Attributes**: Access to `request`, `clientId`, and `handled` promise.

### Navigation Preload Management
```webidl
interface NavigationPreloadManager {
  Promise<undefined> enable();
  Promise<undefined> disable();
  Promise<undefined> setHeaderValue(ByteString value);
  Promise<NavigationPreloadState> getState();
};
```
- **State**: `enabled` (boolean), `headerValue` (ByteString).

## Nuance Or Contradictions

- The specification distinguishes between different client types (`window`, `worker`) and frame types, which affects how clients are queried and interacted with.
- The `NavigationPreloadManager` is optional in some contexts but required for full navigation preload control, allowing developers to toggle this behavior dynamically.
- Service workers can be exposed via both `Navigator` and `WorkerNavigator`, indicating their availability across different browsing contexts.

## Candidate Wiki Hints

- **Service Worker Lifecycle**: Document the states (`parsed`, `installing`, etc.) and transitions between them.
- **Client Management**: Create a guide on querying and managing clients using the `Clients` interface.
- **Navigation Preload**: Explain how to enable, disable, and configure navigation preload headers.
- **Event Handlers**: Detail the lifecycle events (`oninstall`, `onactivate`, `onfetch`) and their typical use cases.

## chunk-17

---
title: Chunk 17 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
This chunk details the "Acknowledgements" section of a technical specification regarding Service Workers. It provides the formal interface definitions for the `Cache` and `CacheStorage` APIs, along with a comprehensive compatibility matrix listing supported methods across various browsers (Firefox, Chrome, Safari, Edge, Opera) and mobile environments.

Local Summary
The section outlines the Web Cache API, defining how service workers interact with cached resources via the `Cache` interface and manage storage via `CacheStorage`. It lists specific methods for adding, deleting, matching, and retrieving cached responses. A significant portion of the text is dedicated to tracking browser compatibility for these APIs and related objects (like `Client`, `FetchEvent`, and `ExtendableEvent`), noting version requirements and support status for desktop and mobile browsers.

Key Claims
- The `Cache` interface is exposed in Secure Contexts within Window and Worker environments.
- Service workers utilize a unique processing model that differs from other web workers, specifically regarding the use of an environment settings object during script fetching.
- The standard `fetch` algorithms for classic and module worker scripts take `job’s client` as an argument, which is null when passed from the Soft Update algorithm.
- Browser support varies significantly by method; for example, `CacheStorage/match` requires Chrome 54+, while basic `CacheStorage` operations are supported in Chrome 43+.

Entities And Concepts
- **Cache**: An interface representing a cache object where requests can be stored and retrieved.
- **CacheStorage**: An interface providing access to the global Cache API, allowing creation of new caches or retrieving existing ones by name.
- **Service Worker**: A background script that enables caching and network interception.
- **FetchEvent**: An event type triggered when a request is intercepted by a service worker.
- **ExtendableEvent**: A base interface for events like FetchEvent and MessageEvent, allowing extension with custom properties.

Procedures And API Details
- **Cache.match**: Returns a promise resolving to a `Response` or `undefined` based on the provided request.
- **Cache.put**: Stores a response in the cache, returning a promise that resolves when complete.
- **CacheStorage.open**: Opens an existing cache named by the provided string and returns it as a `Promise<Cache>`.
- **CacheStorage.delete**: Permanently deletes a cache with the specified name.
- **CacheQueryOptions**: A dictionary containing flags like `ignoreSearch`, `ignoreMethod`, and `ignoreVary` to customize query behavior.

Nuance Or Contradictions
The document notes that certain behaviors are not fully specified yet and will be addressed in the HTML Standard via future issues and pull requests. Specifically, the use of a "to-be-created environment settings object" is highlighted as a temporary measure necessitated by the unique processing model of service workers compared to other web workers.

Candidate Wiki Hints
- **Web Cache API**: A page detailing how to cache resources using the `Cache` and `CacheStorage` interfaces.
- **Service Worker Lifecycle**: Explaining the interaction between `job’s client` and the script fetching algorithms during updates.
- **Browser Compatibility Matrix for Service Workers**: A reference table summarizing version support for specific API methods across major browsers.

## chunk-18

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

## chunk-19

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

