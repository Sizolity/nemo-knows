---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

# Group Context

This group synthesizes the core implementation details and security constraints of the Service Worker specification. It bridges the high-level model (Motivations, Model) with the low-level execution context (Algorithms). The content covers the entire lifecycle from registration and script fetching to installation, activation, event handling, caching strategies, and secure termination. Key themes include the strict requirements for secure contexts, the mechanism of "jobs" for managing state transitions, the internal algorithms for fetch routing, and the security implications of persistent storage and cross-origin resource handling.

# Cross-Chunk Summary

The specification defines a robust lifecycle management system centered on **Jobs**. A job represents a unit of work (register, update, unregister) containing attributes like scope URL, script URL, and worker type. Jobs are queued per scope URL; if the queue is empty or the new job differs from the pending one, it is enqueued. Execution happens in parallel with DOM manipulation tasks.

The **Service Worker Lifecycle** flows through distinct states:
1.  **Installation:** Triggered by fetching a new script. The `Install` algorithm runs, updating state to "installing," queuing an `updatefound` event, and waiting for asynchronous extensions (`extend lifetime promises`). If successful, it dispatches the `install` event.
2.  **Activation:** The worker transitions from "waiting" to "activated." It terminates the previous active worker (if redundant or idle) and matches clients to the new registration, resolving their `readyPromise`.
3.  **Termination:** Workers are cleared when no longer needed. The `Terminate Service Worker` algorithm sets a closing flag, backs up functional events (fetch/push), and discards message tasks.

**Fetch Handling** is managed by the `Handle Fetch` algorithm. It determines if a worker can handle a request based on security context. It matches registrations using storage keys and evaluates **Router Conditions**. These conditions support `urlPattern`, `requestMethod`, and logical operators (`_or`, `not`). The system supports "Soft Updates" to refresh stale workers without blocking, ensuring the worker is current before responding.

**Caching** is handled via the Cache API (`self.caches` and `CacheStorage`). Responses are stored as either **Basic Filtered**, **CORS Filtered**, or **Opaque Filtered**. The spec enforces strict validation for cache operations (e.g., HTTP/HTTPS schemes) and provides atomic `Batch Cache Operations` with rollback capabilities on failure.

# Repeated Or Central Claims

-   **Secure Context Requirement:** Service workers must execute in secure contexts (typically HTTPS). Exceptions exist only for localhost, 127.0.0.0/8, and ::1/128 during development. Both the worker and the registering client must be secure.
-   **Origin Relativity:** A service worker executes within the registering client's origin. It cannot host resources on a CDN directly unless specific headers are present. This prevents cross-origin attacks and enforces strict same-origin policies for scripts.
-   **Job Equivalence:** Two jobs are equivalent if they share specific attributes (type, scope URL, script URL, worker type, cache mode). If a new job is equivalent to a pending one, it is appended to that job's list rather than creating a duplicate entry.
-   **Fetch Routing Logic:** The `Handle Fetch` algorithm determines the response source (cache, network, or fetch-event listener) based on router conditions. This logic includes race responses for GET requests and soft updates for stale workers.
-   **Router Condition Limits:** To prevent performance degradation, router condition complexity is constrained: total nested conditions cannot exceed 1024, and nesting depth is limited to 10 levels.
-   **Security on Regex:** User-defined regular expressions in URL patterns are prohibited due to security concerns. If a pattern has regexp groups, the router condition verification fails.

# Important Local Details

-   **Fragment Handling:** The specification explicitly discards fragments (hashes) in both script URLs and scope URLs during identification. Only the scheme and path matter.
-   **Bad Import Responses:** Responses with bad MIME types or non-ok status codes for `importScripts()` are ignored for the byte-to-byte update check but may still populate the cache. This prevents minor fetch errors from triggering unnecessary re-installs.
-   **Event Skipping:** To avoid performance delays, the user agent may skip event dispatch if no listeners exist deterministically added during the first script execution.
-   **Batch Cache Operations:** These operations perform atomic writes. If an exception occurs (e.g., quota exceeded), the implementation rolls back all changes made during the batch job.
-   **Clear Registration:** A registration is cleared only if its installing, waiting, and active workers are null or have no pending events. This ensures data integrity before removing a worker from storage.
-   **Scope Matching:** URL string matching for service worker scopes is prefix-based rather than path-structural. It maintains same-origin security but relies on trailing slashes in serialized URLs for safety.

# Candidate Wiki Hints

-   **Service Worker Job Lifecycle:** A deep dive into how `Jobs` manage state transitions, queue synchronization, and promise resolution.
-   **Secure Context Enforcement:** Explaining the requirements for HTTPS, localhost exceptions, and CSP headers.
-   **Fetch Event Routing:** Detailing the `Handle Fetch` algorithm, router conditions, soft updates, and race responses.
-   **Cache API Security:** Covering response types (Basic/CORS/Opaque filtered), validation rules, and batch operations with rollback.
-   **Router Condition Architecture:** Documenting the structure of router conditions (`urlPattern`, `requestMethod`, `_or`, `not`) and their complexity limits.
-   **Installation vs. Activation:** A visual or textual guide to the state machine (Installing -> Waiting -> Active) and event dispatching.

# Gaps Or Cautions

-   **Algorithm Implementation Flexibility:** While the spec defines minimal setup for `Setup ServiceWorkerGlobalScope`, implementations can optimize this step as long as observable behavior (results of security checks like CSP) remains equivalent.
-   **Activation Robustness:** Activation handlers are designed to do non-essential work because they may not complete if the browser terminates during activation. Handlers must function properly even if they fail.
-   **Plugin Discouragement:** Plugins are explicitly discouraged from loading via service workers due to origin handling limitations in the `Handle Fetch` algorithm, specifically regarding security origins.
-   **Privacy and Data Purge:** Persistent storages (registration map, cache name to map, script resource map) must be cleared when users purge data, but specific details on *how* this happens across different storage mechanisms may vary by implementation.
-   **Empty Fetch Listeners:** The spec notes that empty function bodies in `fetch` listeners (`() => {}`) are used by some sites to signal PWA status but may have performance implications.
