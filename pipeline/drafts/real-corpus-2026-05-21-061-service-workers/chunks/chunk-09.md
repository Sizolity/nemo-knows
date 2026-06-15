---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
This chunk details the "Appendix A: Algorithms" section of the Service Workers specification. It focuses on the `fetch` hook implementation for updating scripts, the logic for determining if resources have been updated (`hasUpdatedResources`), and the sequential algorithms governing the lifecycle of a service worker: **Install**, **Activate**, and **Try Activate**. It also introduces the `Setup ServiceWorkerGlobalScope` algorithm, which handles the creation of the execution context (realm, agent, settings object) required for security checks like CSP.

Local Summary
The text describes how the user agent fetches scripts via a specific hook to handle the unique processing model of service workers. The process involves appending headers (`Service-Worker/script`), managing cache modes, and validating MIME types. If the main script or imported scripts change (byte-for-byte comparison), the worker is marked as updated. The chunk then outlines the state machine transitions:
1.  **Install**: Runs when a new script is fetched and validated. It updates internal states, queues an `updatefound` event, and eventually dispatches an `install` event after waiting for asynchronous extensions (like `extend lifetime promises`). If installation fails or is skipped, the worker becomes redundant.
2.  **Activate**: Transitions the worker from a "waiting" state to "activating" and finally "activated". It terminates the previous active worker if necessary, matches clients to the new registration, and dispatches an `activate` event.
3.  **Try Activate**: A helper algorithm that checks conditions (e.g., existing active worker is null or has no pending events) before invoking `Activate`.
4.  **Setup ServiceWorkerGlobalScope**: Ensures the necessary global scope object exists for security checks, creating a new realm and agent if needed.

Key Claims
-   Service workers use a separate script fetching mechanism (`Update algorithm`) compared to other web workers, requiring a specific environment settings object approach.
-   The `fetch` hook for service workers must append `Service-Worker/script` headers and set the redirect mode to "error".
-   A worker is considered updated if the main script URL changes, the MIME type differs, or the body is not byte-for-byte identical. Imported scripts are checked separately; bad import responses are ignored for the update check but populate the cache.
-   The `Install` algorithm waits for all `extend lifetime promises` associated with the `install` event to settle before proceeding to activation logic.
-   Activation occurs only if the waiting worker is valid and the current active worker is either null, redundant, or has no pending events/clients using it.
-   The `Setup ServiceWorkerGlobalScope` algorithm ensures a `ServiceWorkerGlobalScope` object exists for CSP checks before any security validation occurs.

Entities And Concepts
-   **Service Worker Registration**: Manages the state (installing, waiting, active) and maps workers to storage keys/scopes.
-   **Job Promise**: Used to coordinate asynchronous operations during the fetch/update cycle; can be rejected with "SecurityError" or "TypeError".
-   **Environment Settings Object**: The execution context for service worker algorithms, distinct from classic workers.
-   **Realm**: The isolation boundary containing the global object and agent.
-   **ServiceWorkerGlobalScope**: The global scope object returned by `Setup ServiceWorkerGlobalScope`, used for security checks.
-   **Extend Lifetime Promises**: Asynchronous extensions that can delay the installation process; their settlement is awaited before activation logic proceeds.
-   **Update Via Cache Mode**: Determines how updates are fetched (e.g., "all", "none").

Procedures And API Details
-   **Fetch Hook Steps**:
    -   Append `Service-Worker/script` header.
    -   Set cache mode to "no-cache" if registration is stale or force bypass is set.
    -   Set redirect mode to "error".
    -   Fetch request; on response, extract MIME type and headers (`Service-Worker-Allowed`).
    -   Validate scope permissions against `maxScopeString`.
    -   Check byte-for-byte identity of main script and imported scripts.
-   **Install Algorithm**:
    -   Update registration state to "installing".
    -   Resolve job promise with the worker registration.
    -   Fire `updatefound` event on registration objects.
    -   Run `Run Service Worker` (awaiting async extensions).
    -   If successful, dispatch `InstallEvent`.
    -   Terminate waiting worker and transition to "waiting" state for the new worker.
-   **Activate Algorithm**:
    -   Terminate active worker if present.
    -   Transition states: active -> redundant, waiting -> null, installing -> active.
    -   Match clients by creation URL and scope.
    -   Resolve `readyPromise` for clients.
    -   Dispatch `ActivateEvent`.
-   **Try Activate**: Checks if activation can proceed (no active worker or active worker is idle/no clients).

Nuance Or Contradictions
-   **Bad Import Responses**: The spec explicitly states that bad responses for `importScripts()` are ignored for the purpose of the byte-to-byte update check. Only good responses from the incumbent and potential update workers count. This prevents minor fetch errors from triggering unnecessary re-installs.
-   **Activation Robustness**: Activation handlers are designed to do non-essential work (like cleanup) because they may not complete if the browser terminates during activation. The worker must function properly even if handlers fail.
-   **Implementation Flexibility**: While the spec defines a minimal setup for `Setup ServiceWorkerGlobalScope`, implementations can optimize this step as long as the observable behavior (results of security checks like CSP) remains equivalent.

Candidate Wiki Hints
-   **Service Worker Lifecycle**: Create a page explaining the Install -> Activate cycle, detailing state transitions and event dispatching.
-   **Update Strategies**: Document how `fetch` caching modes ("all", "none") interact with the `Service-Worker` header and stale checks.
-   **Security Scope**: Explain the role of `ServiceWorker-Allowed` headers and scope validation in preventing unauthorized updates.
-   **Global Scope Setup**: A technical note on why `Setup ServiceWorkerGlobalScope` is necessary for Content Security Policy (CSP) enforcement in service workers.
