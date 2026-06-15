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
