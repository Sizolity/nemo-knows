---
title: Service Worker Event Handling
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

# Service Worker Event Handling

Service Workers function on an event loop driven by specific events rather than document lifetimes. This asynchronous architecture avoids blocking resource loading, allowing the worker to operate within a background context (`ServiceWorkerGlobalScope`) distinct from the main document thread.

## Core Events and Lifecycle

The specification defines standard events extended with methods such as `event.waitUntil()` to defer task completion and `event.respondWith()` to intercept requests. The primary event types include:

- **`InstallEvent`**: Triggered during script installation.
- **`FetchEvent`**: Handles network request interception.
- **`ExtendableMessageEvent`**: Facilitates asynchronous communication.

The workflow involves registering a script, installing it with route definitions, and activating it to take control of clients within its scope. The registration object tracks three lifecycle states: `installing`, `waiting`, and `active`. The `update()` method handles upgrading from the waiting to active state, while `skipWaiting()` bypasses the waiting phase entirely.

## Communication and Control

Clients (windows or workers) are controlled by a service worker if they share an origin and the scope matches. The currently controlling worker is accessible via `navigator.serviceWorker.controller`. An active worker can take control of existing clients using the `claim()` method.

Communication occurs asynchronously using `postMessage()`. This allows data exchange between the global scope and client environments without direct message handling, distinguishing service workers from shared workers.

## Caching Strategy

Resources are managed through the Cache API (`self.caches`), which is isolated from the browser's HTTP cache. Operations include `match`, `matchAll`, `put`, `add`, `delete`, and `keys`. Responses are stored as immutable resources keyed by request info, supporting strategies such as basic filtering, CORS filtering, or opaque filtering.

## Security Constraints

Service Workers must execute in secure contexts, typically HTTPS. Exceptions exist for localhost, 127.0.0.0/8, and ::1/128 during development. Both the worker and registering client must be secure. Additional constraints include origin relativity, CORS restrictions on cross-origin resources, path restrictions, and strict script request handling via `importScripts`.

## Error Recovery

Unlike Application Cache, service workers avoid unrecoverable states. Errors during installation or activation trigger specific behaviors (e.g., moving to a redundant state) rather than leaving the system broken. Job-based state management ensures the lifecycle flows through distinct states managed by units of work queued per scope URL.
