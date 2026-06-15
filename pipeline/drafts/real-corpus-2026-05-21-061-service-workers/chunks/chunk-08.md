---
title: Chunk 08 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---
Chunk Context
- Source Path: `raw/web/corpus-2026-05-18/061-service-workers.md`
- Heading: 7. Extensibility > Appendix A: Algorithms
- Range: Lines 3014–3350

Local Summary
This chunk defines the core internal data structures and algorithms for managing Service Worker registrations, updates, and unregistrations. It details the lifecycle of a "job" (the abstract unit of work), including its creation, scheduling, execution, and promise resolution/rejection. The text also outlines validation rules for script URLs and scope URLs, distinguishing between "classic" and "module" worker types during the fetch phase.

Key Claims
- A **job** is an abstraction representing a request to register, update, or unregister a service worker registration.
- Two jobs are considered **equivalent** if they share specific attributes (type, scope URL, script URL, worker type, cache mode) depending on whether they are register/update or unregister operations.
- The **scope URL** and **script URL** fragments are explicitly discarded by the user agent; only the scheme and path matter for identification.
- Security checks enforce that the script origin, referrer origin, and scope origin must all be "same origin" to prevent cross-origin attacks.
- If a registration already exists and matches the new job's parameters exactly (for register/update), the existing registration is returned immediately without re-fetching.

Entities And Concepts
- **Job**: An internal data structure containing type, storage key, scope URL, script URL, worker type, cache mode, client, referrer, promise, and status flags.
- **Job Queue**: A thread-safe queue used to synchronize concurrent jobs per scope URL.
- **Registration Map**: An ordered map linking (storage key, serialized scope URL) pairs to service worker registrations.
- **Worker Type**: Either `"classic"` or `"module"`, affecting how the script graph is fetched.
- **Bad Import Script Response**: Defined as an error type response, non-ok status, or a MIME type that is not JavaScript.

Procedures And API Details
- **Create Job**: Constructs a new job object by setting its attributes (type, keys, URLs, promise, client) and initializing referrer if a client exists.
- **Schedule Job**: Determines the appropriate job queue for the scope URL. If the queue is empty, it triggers execution (`Run Job`). If not empty but equivalent to the last pending job, it appends to that job's list of equivalents; otherwise, it enqueues.
- **Run Job**: Queues a task to execute `Register`, `Update`, or `Unregister` in parallel with the DOM manipulation task source. Execution is delayed until after a `DOMContentLoaded` event.
- **Finish Job**: Dequeues the completed job from its queue and triggers `Run Job` if the queue is not empty.
- **Resolve Job Promise**: Resolves the job's promise (and equivalent jobs' promises) with either a service worker registration object or the provided value, queued on the client's event loop.
- **Start Register**: Validates script URL and scope URL (rejecting non-HTTP/HTTPS schemes, invalid characters like `%2f` or `%5c`, and null values). It creates a job and schedules it.
- **Register Algorithm**: Performs security checks (origin trust, same-origin policy). If valid, it checks for an existing matching registration; if found, resolves the promise with the existing registration. Otherwise, it sets up the new registration via `Set Registration` and calls `Update`.
- **Update Algorithm**: Retrieves the existing registration. If null or if a newer worker exists with a different script URL, it rejects the promise. It then proceeds to fetch resources based on worker type.

Nuance Or Contradictions
- **Fragment Handling**: The specification explicitly notes that fragments (hashes) in both script URLs and scope URLs are set to `null` and have no effect on identification. This prevents confusion where different hashes point to the same logical resource.
- **Equivalence Logic**: Equivalence is defined differently for register/update jobs (requiring all attributes to match) versus unregister jobs (requiring only scope URL match).
- **Parallel Execution**: The `Run Job` step uses "in parallel" when invoking sub-algorithms, implying these operations can occur concurrently with other tasks on the event loop.

Candidate Wiki Hints
- Service Worker Job Lifecycle and State Machine
- Understanding Equivalent Jobs in Service Workers
- Script URL and Scope URL Validation Rules
