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
