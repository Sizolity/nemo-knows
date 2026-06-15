---
title: Service Worker Lifecycle
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

# Service Worker Lifecycle

Service Workers are generic, event-driven script contexts that run at an origin to intercept network requests and manage offline functionality. They operate asynchronously within a background context (`ServiceWorkerGlobalScope`) distinct from the main document thread. Unlike shared workers, they never handle messages directly but process specific events such as `fetch`, `install`, and `activate`.

## Lifecycle States

A service worker registration persists across user agent restarts but manages three distinct worker instances based on their state:

- **Installing**: The current script is being installed with route definitions.
- **Waiting**: A newer script is present but pending activation.
- **Active**: The worker is currently controlling clients within its scope.

The lifecycle workflow typically involves registering a script, installing it, activating it to take control of clients, and handling events via extendable handlers. Communication between the global scope and client environments occurs asynchronously using `postMessage()`.

## State Transitions

The transition from the `waiting` state to the `active` state is handled by the `update()` method. To bypass the waiting phase entirely, the `skipWaiting()` method can be used.

## Security and Scope

Service workers must execute in secure contexts (typically HTTPS). Exceptions exist for localhost (`127.0.0.0/8`, `::1/128`) during development. Both the worker and the registering client must be secure. The scope defaults to the directory containing the script file (e.g., a script at `/js/sw.js` controls `/js/`). This can be expanded using the `Service-Worker-Allowed` HTTP response header, with scope matching being prefix-based.

## Caching

Caching is handled via the Cache API (`self.caches`), which is isolated from the browser's HTTP cache. Resources are managed as immutable resources keyed by request info. Operations include `match`, `matchAll`, `put`, `add`, `delete`, and `keys`. The architecture supports multiple named caches with responses stored as Basic Filtered, CORS Filtered, or Opaque Filtered.

## Event Handling

Standard events are extended with methods like `event.waitUntil()` (to defer task completion) and `event.respondWith()` (to intercept requests). Specific event types include `InstallEvent`, `FetchEvent`, and `ExtendableMessageEvent`. The specification distinguishes service workers from shared workers by noting that they process these specific events rather than handling messages directly.

## Error Recovery

Unlike Application Cache, service workers avoid unrecoverable states. Errors during installation or activation trigger specific behaviors (e.g., moving to a redundant state) rather than leaving the system broken.
