---
title: Service Worker Scope Registration
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

# Service Worker Scope Registration

A **Service Worker** is a generic, event-driven script context that runs at an origin to intercept network requests and manage offline functionality. Unlike shared workers, service workers never handle messages directly but process specific events like `fetch`, `install`, and `activate`. They operate asynchronously within a background execution context (`ServiceWorkerGlobalScope`) distinct from the main document thread.

## Registration Lifecycle

The registration object tracks three distinct lifecycle states for worker instances:
- **`installing`**: The current script is being installed.
- **`waiting`**: A newer script is pending activation.
- **`active`**: The worker is currently controlling clients.

The workflow involves registering a script, installing it with route definitions, and activating it to take control of clients within its scope. Communication occurs asynchronously using `postMessage()`. Resources are managed through the Cache API (`self.caches`), which is isolated from the browser's HTTP cache. Security is enforced through secure contexts (HTTPS/localhost) and strict path restrictions.

## Scope Definitions

A service worker's scope defaults to the directory containing its script file. For example, a script located at `/js/sw.js` controls resources under `/js/`. This relationship is determined by prefix-based matching. The scope can be expanded using the `Service-Worker-Allowed` HTTP response header.

## Security and Constraints

Service workers must execute in secure contexts, typically HTTPS. Exceptions exist for localhost, 127.0.0.0/8, and ::1/128 during development. Both the worker and the registering client must be secure. Origin relativity, CORS restrictions on cross-origin resources, and strict script request handling are enforced. The `importScripts` method is subject to origin restrictions.

## Related Concepts

- [[service-worker-lifecycle]]
- [[service-worker-client-management]]
- [[service-worker-security-constraints]]
