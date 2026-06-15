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
