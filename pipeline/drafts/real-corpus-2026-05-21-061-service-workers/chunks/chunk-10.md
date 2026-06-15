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
