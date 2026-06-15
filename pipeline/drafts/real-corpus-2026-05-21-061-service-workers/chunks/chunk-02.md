---
title: Chunk 02 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

## Chunk Context
This chunk defines the internal structure of a **Service Worker Registration**, detailing its associated workers (installing, waiting, active), task queues, and update settings. It further explains the lifecycle of these registrations during user agent shutdown. The text then transitions to defining **Service Worker Clients** (window, worker) and how they acquire an active service worker based on origin matching and fetch routing. Finally, it covers specific **Task Sources** for events and the rules for persisting registration state across restarts. The latter half introduces the `ServiceWorker` and `ServiceWorkerRegistration` IDL interfaces, their attributes (like `scriptURL`, `scope`, `state`), and the logic for creating their corresponding objects in an environment settings map.

## Local Summary
A service worker registration is a persistent tuple containing a scope URL, a storage key, and up to three associated workers representing different lifecycle states: installing, waiting, and active. It also tracks task queues, update cache modes, and navigation preload headers. Clients (window or worker) are environments that may be controlled by an active service worker; this control is determined by origin matching rules during creation or navigation. The `ServiceWorker` and `ServiceWorkerRegistration` interfaces expose read-only attributes for these states and URLs, along with methods like `postMessage` and `update`. Object instantiation logic ensures a one-to-one mapping between the internal registration state and the exposed objects within an environment's map.

## Key Claims
- A service worker registration persists across user agent restarts but discards installing workers; waiting workers promote to active upon restart.
- Task queues for a registration back up tasks from the active worker's event loop when terminated, re-queuing them when the worker spins off.
- A window client's active service worker is null if the origin is opaque or differs from the creator document's origin during navigation (unless routed via HTTP fetch in specific ways).
- The `ServiceWorkerRegistration` interface exposes three optional attributes: `installing`, `waiting`, and `active`, each holding a `ServiceWorker?` object.
- The `postMessage` method on `ServiceWorker` serializes the message, handles transferable objects, and dispatches an event at the destination `ServiceWorkerGlobalScope`.

## Entities And Concepts
- **Service Worker Registration**: A tuple of scope URL, storage key, and worker states (installing/waiting/active).
- **Service Worker Client**: An environment that is either a window client or a worker client.
- **Window Client**: Controlled by a service worker if the origin matches and fetch is routed appropriately.
- **Worker Client**: Inherits active service worker from owner unless origins differ or data/blob URLs are used.
- **Task Sources**: `handle fetch` and `handle functional event` (e.g., push).
- **ServiceWorker State Enum**: `parsed`, `installing`, `installed`, `activating`, `activated`, `redundant`.
- **ServiceWorkerUpdateViaCache Enum**: `imports`, `all`, `none`.

## Procedures And API Details
- **Getting ServiceWorker objects**: Retrieve from the environment settings object map; create a new one if missing, associating it with the internal service worker and copying its state.
- **Getting ServiceWorkerRegistration objects**: Retrieve from the environment settings registration object map; initialize attributes to null or the corresponding `ServiceWorker` objects if the internal workers are non-null.
- **postMessage(message, options)**:
  1. Serialize message with transfer support.
  2. Check if event dispatching should be skipped.
  3. Run parallel substeps to determine source (window/client) and destination (`ServiceWorkerGlobalScope`).
  4. Deserialize transferred values.
  5. Create an `ExtendableMessageEvent` with data, ports, and origin.
  6. Dispatch the event at the destination.
- **scriptURL**: Returns the serialized script URL of the associated service worker (e.g., `https://example.com/service_worker.js`).
- **scope**: Returns the serialized scope URL (e.g., `https://example.com/`).

## Nuance Or Contradictions
- The text notes that behavior in section 2.5 is non-normative and will be specified in the HTML Standard, implying current implementation details might evolve.
- Sandboxed iframes without specific directives (`allow-same-origin`, `allow-scripts`) result in an opaque origin, causing their active service worker to be null.
- Data URLs and Blob URLs inherently result in an opaque origin for window/worker clients, leading to a null active service worker value regardless of creator context.

## Candidate Wiki Hints
- **Service Worker Registration Lifecycle**: Explain the three worker states (installing, waiting, active) and how they map to the registration object.
- **Client Control Logic**: Detail how `navigator.serviceWorker.controller` is determined for window clients versus worker clients.
- **Message Passing Protocol**: Breakdown of `postMessage` serialization, transfer handling, and event dispatching mechanics.
- **Interface Definitions**: Summarize `ServiceWorker` and `ServiceWorkerRegistration` attributes and their underlying internal counterparts.
