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
