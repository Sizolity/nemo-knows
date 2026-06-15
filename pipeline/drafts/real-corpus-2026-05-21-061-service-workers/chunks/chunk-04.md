---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
This chunk defines the `Client` and `WindowClient` interfaces, their attributes (such as `url`, `id`, `type`, `visibilityState`, `focused`), and methods (`postMessage`, `focus`, `navigate`). It also details the `Clients` global interface with methods like `get`, `matchAll`, `openWindow`, and `claim`. The section concludes by introducing the `ExtendableEvent` interface used for service worker lifecycle events (`install` and `activate`) and extensions.

Local Summary
The chunk outlines how clients are represented within a Service Worker context, distinguishing between generic clients (workers, shared workers) and window clients (top-level browsing contexts). It specifies how to retrieve client details, send messages via `postMessage`, manage focus states, navigate to new URLs, and open windows from the service worker scope. The `Clients` interface provides mechanisms to query active clients, filtering by type or control status, and handles the association between clients and their respective service worker registrations.

Key Claims
- A `Client` object represents a service worker client with an associated frame type ("auxiliary", "top-level", "nested", "none") and a unique `id`.
- A `WindowClient` extends `Client` to include browsing context-specific properties like `visibilityState`, `focused`, and `ancestorOrigins`.
- The `postMessage` method allows communication between the service worker and its clients, supporting structured serialization with transferable objects.
- The `Clients` interface is created when a `ServiceWorkerGlobalScope` object is created and manages all associated clients.
- `matchAll` returns an array of clients sorted by focus state (focused first) and creation order, respecting the `type` filter in options.
- `claim()` allows a service worker to take control of its clients, handling unloading if the client's active service worker differs from the claiming one.

Entities And Concepts
- Client: Interface representing a service worker client.
- WindowClient: Sub-interface for clients associated with window browsing contexts.
- FrameType: Enum defining frame hierarchy ("auxiliary", "top-level", "nested", "none").
- Clients: Global interface to manage and query active clients.
- ExtendableEvent: Interface extending Event, used for service worker lifecycle events (`install`, `activate`).
- ServiceWorkerGlobalScope: Scope where the `Clients` object is instantiated.

Procedures And API Details
- **Client Attributes**:
  - `url`: Returns the serialized creation URL of the associated service worker client.
  - `type`: Returns "window", "worker", or "sharedworker" based on the client type.
- **Client Methods**:
  - `postMessage(message, options)`: Sends a message to the client, handling serialization and deserialization with transferable objects.
  - `focus()`: Requests focus for the browsing context; returns a promise resolving to the `WindowClient` if successful.
  - `navigate(url)`: Navigates the associated browsing context to a new URL.
- **Clients Interface Methods**:
  - `get(id)`: Returns a promise resolving to a specific client by ID or `undefined`.
  - `matchAll(options)`: Returns a promise resolving to an array of clients matching criteria (e.g., type, includeUncontrolled).
  - `openWindow(url)`: Opens a new window and returns a `WindowClient` if the origin matches.
  - `claim()`: Takes control of all clients associated with the service worker.
- **ExtendableEvent**:
  - `waitUntil(promise)`: Keeps the service worker alive until the promise resolves or is rejected.

Nuance Or Contradictions
- The `type` getter logic includes a fallback to "window" if the client is an environment settings object, which might seem redundant given the specific checks for window/worker/sharedworker clients.
- In `matchAll`, sorting prioritizes focused window clients over unfocused ones and separates them from worker clients regardless of creation order within their respective groups.
- The `claim()` method explicitly rejects if the service worker is not active, ensuring control can only be taken by an active worker.

Candidate Wiki Hints
- **Service Worker Client Interface**: Overview of `Client` and `WindowClient` attributes and methods.
- **Clients API**: Guide to querying clients via `get()`, `matchAll()`, `openWindow()`, and `claim()`.
- **Interacting with Clients**: Best practices for using `postMessage` and handling client types.
- **Service Worker Lifecycle**: Explanation of `ExtendableEvent` usage in `install` and `activate` events.
