---
title: Service Workers Specification Summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

## What It Is

This document is a specification for **Service Workers**, defined as generic, event-driven script contexts that run at an origin to intercept network requests and manage offline functionality. The specification (W3C Candidate Recommendation Draft, April 2026) details the lifecycle, API interfaces (`ServiceWorker`, `ServiceWorkerRegistration`, `Client`, `Cache`), execution contexts, security constraints, and browser compatibility requirements. It distinguishes service workers from shared workers by noting that they never handle messages directly but process events like `fetch`, `install`, and `activate`.

## Summary

Service Workers operate asynchronously within a background context (`ServiceWorkerGlobalScope`) distinct from the main document thread. They are managed via a registration object that tracks three lifecycle states: `installing` (current script installation), `waiting` (newer script pending activation), and `active` (currently controlling clients). The workflow involves registering a script, installing it with route definitions, activating it to take control of clients within its scope, and handling events via extendable handlers. Communication occurs asynchronously using `postMessage()`. Resources are managed through the Cache API (`self.caches`), which is isolated from the browser's HTTP cache. Security is enforced through secure contexts (HTTPS/localhost), origin relativity, and strict path restrictions. The specification also covers extensibility mechanisms, navigation preload headers, and detailed browser support matrices.

## Key Claims

- **Event-Driven Architecture:** Service workers function on an event loop driven by specific events rather than document lifetimes. They are asynchronous to avoid blocking resource loading.
- **Registration Lifecycle States:** A registration persists across user agent restarts but manages three distinct worker instances: `installing`, `waiting`, and `active`. The `update()` method handles upgrading from waiting to active, while `skipWaiting()` bypasses the waiting phase.
- **Client Control & Communication:** Clients (windows or workers) are controlled by a service worker if they share an origin and the scope matches. `navigator.serviceWorker.controller` returns the currently controlling worker. `claim()` allows an active worker to take control of existing clients. Communication is handled via `postMessage()` between the global scope and client environments.
- **Caching Separation:** Caches are isolated from the browser's HTTP cache, requiring manual population and management. Responses are stored as immutable resources keyed by request info. Operations include `match`, `matchAll`, `put`, `add`, `delete`, and `keys`.
- **Scope and Registration:** A service worker's scope defaults to the directory containing its script file (e.g., `/js/sw.js` controls `/js/`). This can be expanded using the `Service-Worker-Allowed` HTTP response header. Scope matching is prefix-based.
- **Secure Context Requirement:** Service workers must execute in secure contexts (typically HTTPS). Exceptions exist for localhost, 127.0.0.0/8, and ::1/128 during development. Both the worker and registering client must be secure.
- **Fetch Handling & Routing:** The `Handle Fetch` algorithm determines response sources based on router conditions (`urlPattern`, `requestMethod`). It supports race responses for GET requests and soft updates for stale workers to ensure current behavior without blocking.
- **Extendable Events:** Standard events are extended with methods like `event.waitUntil()` (to defer task completion) and `event.respondWith()` (to intercept requests). Specific event types include `InstallEvent`, `FetchEvent`, and `ExtendableMessageEvent`.
- **Error Recovery:** Unlike Application Cache, service workers avoid unrecoverable states. Errors during installation or activation trigger specific behaviors (e.g., moving to redundant state) rather than leaving the system broken.
- **Job-Based State Management:** The lifecycle flows through distinct states managed by "Jobs" (units of work). A job contains attributes like scope URL and script URL. Jobs are queued per scope URL; execution happens in parallel with DOM manipulation tasks.
- **Cache Storage Architecture:** Caching is handled via `CacheStorage` allowing multiple named caches. Responses are stored as Basic Filtered, CORS Filtered, or Opaque Filtered. Batch cache operations support atomic writes with rollback on failure.
- **Security Constraints:** Includes origin relativity, CORS restrictions on cross-origin resources, path restrictions, and strict script request handling. `importScripts` is subject to origin restrictions.

## Suggested Links

- none
