---
title: Service Worker Client Management
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

# Service Worker Client Management

Service Workers operate asynchronously within a background context, distinct from the main document thread. They are defined as generic, event-driven script contexts that run at an origin to intercept network requests and manage offline functionality. Unlike shared workers, service workers never handle messages directly; instead, they process specific events such as `fetch`, `install`, and `activate`.

## Lifecycle and Registration

Service Workers are managed via a registration object that tracks three distinct lifecycle states:
- **`installing`**: The current script is being installed.
- **`waiting`**: A newer script is pending activation but has not yet taken control of clients.
- **`active`**: The worker is currently controlling clients within its scope.

The workflow involves registering a script, installing it with route definitions, and activating it to take control of clients. The `update()` method handles upgrading from the `waiting` state to `active`, while `skipWaiting()` bypasses the waiting phase entirely. Errors during installation or activation trigger specific recovery behaviors rather than leaving the system in a broken state.

## Client Control and Communication

Clients (browsing contexts like windows or workers) are controlled by a service worker if they share an origin and the scope matches. The current controlling worker can be accessed via `navigator.serviceWorker.controller`. An active worker can take control of existing clients immediately using the `claim()` method.

Communication between the service worker global scope and client environments occurs asynchronously using `postMessage()`.

## Scope and Security

A service worker's scope defaults to the directory containing its script file (e.g., a script at `/js/sw.js` controls `/js/`). This scope can be expanded using the `Service-Worker-Allowed` HTTP response header. Scope matching is prefix-based.

Security is enforced through secure contexts, typically requiring HTTPS. Exceptions exist for localhost, 127.0.0.0/8, and ::1/128 during development. Both the worker and the registering client must be in a secure context.

## Resource Management

Resources are managed through the Cache API (`self.caches`), which is isolated from the browser's HTTP cache. Responses are stored as immutable resources keyed by request info. Operations include `match`, `matchAll`, `put`, `add`, `delete`, and `keys`. Caching is handled via `CacheStorage`, allowing multiple named caches with batch operations that support atomic writes.

## Event Handling

Standard events are extended with methods like `event.waitUntil()` to defer task completion and `event.respondWith()` to intercept requests. Specific event types include `InstallEvent`, `FetchEvent`, and `ExtendableMessageEvent`. The specification also covers navigation preload headers and extensibility mechanisms for the fetch handler.
