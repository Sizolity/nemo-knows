---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
This chunk covers **Appendix A: Algorithms** within the Service Workers specification (Section 7. Extensibility). It details the low-level algorithms governing request handling, event dispatching, router condition matching, registration lifecycle management, and cleanup procedures during shutdown or client unloading.

Local Summary
The text defines a comprehensive set of algorithms for service worker execution flow. Key areas include the main fetch handler logic (handling race conditions, soft updates, and error propagation), URL pattern parsing for routing, complex nested router condition verification (supporting `or`, `not`, and status checks), and lifecycle management functions like clearing registrations and handling user agent shutdowns.

Key Claims
- **Soft Updates:** The spec introduces "soft update" logic to refresh stale registrations or workers under specific conditions (e.g., non-subresource requests on stale registrations).
- **Event Skipping:** To avoid performance delays, the UA may skip event dispatch if no listeners exist deterministically added during the first script execution.
- **Router Limits:** Router condition complexity is constrained: total nested conditions cannot exceed 1024, and nesting depth is limited to 10 levels to prevent exponential computation.
- **Security on Regex:** User-defined regular expressions in URL patterns are prohibited due to security concerns; if a pattern has regexp groups, the router condition verification fails.
- **Body Cancellation:** If a fetch response is null but the request body source is null and the body is unusable, the body is cancelled with `undefined`.

Entities And Concepts
- **Service Worker State:** Workers transition between states like "activating", "activated", "redundant".
- **FetchEvent Attributes:** Includes `respondWith`, `waitUntil` (implied via flags), `clientId`, `replacesClientId`, and `preloadResponse`.
- **Router Conditions:** Structured objects supporting `urlPattern`, `requestMethod`, `requestMode`, `runningStatus`, `_or`, and `not` operators.
- **Registration Map:** A storage mechanism mapping scope/storage keys to service worker registration objects.
- **Extended Events Set:** A collection tracking active extended events for a specific worker, cleaned up upon dispatch or inactivity.

Procedures And API Details
- **Parse URL Pattern:** Converts a raw pattern string into a `URLPattern` object relative to the service worker's script URL. Throws exceptions if invalid.
- **Verify Router Condition:** Recursively validates router conditions. Returns `false` if regex groups exist, forbidden methods are used, or nested depth/counts exceed limits.
- **Match Router Condition:** Evaluates if a request matches a specific condition. Supports short-circuit logic for `or` (match any) and `not` (invert match).
- **Count Router Inner Conditions:** Helper algorithm to track recursion depth and total condition count against the 1024/10 limits.
- **Fire Functional Event:** Creates, initializes (optional), and dispatches an event on the active worker's global object, handling stale registration updates in parallel.
- **Clear Registration:** Terminates installing, waiting, or active workers associated with a registration and updates their state to "redundant" or clears the registration entry.

Nuance Or Contradictions
- **Mutual Exclusivity of Router Logic:** The spec explicitly states that `_or` and `not` conditions are mutually exclusive with other conditions (like `urlPattern`) for ease of understanding, enforced by returning `false` if combined improperly.
- **Abort Handling:** If a fetch controller is terminated/aborted, the algorithm deserializes the abort reason and signals it via an AbortController; if the task handling this is discarded, it defaults to treating it as a fetch failure (`handleFetchFailed`).
- **Stale Worker Handling:** When skipping events or handling failures on stale workers, the spec mandates running "Soft Update" in parallel rather than blocking.

Candidate Wiki Hints
- Service Worker Router Architecture
- Fetch Event Lifecycle and Race Responses
- Soft Update Mechanism
- URL Pattern Security Constraints
