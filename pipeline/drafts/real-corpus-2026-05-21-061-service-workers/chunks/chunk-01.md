---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

## Chunk Context
**Heading path:** Document → Service Workers → 1. Motivations, 2. Model (Service Worker state, timing)
**Line range:** 1–497
**Coverage:** Abstract, Status, Table of Contents, Motivations, and initial Model definitions including Service Worker attributes and timing structs.

## Local Summary
This chunk introduces the Service Workers specification as a W3C Candidate Recommendation Draft (April 2026). It explains that service workers are event-driven worker contexts designed to intercept network requests, enable offline functionality, and manage background tasks without relying on document lifetimes. The text details motivations for flexibility over older cache APIs, emphasizes recoverable error handling, and outlines the asynchronous, single-threaded execution model. The Model section begins by defining core attributes (state, script URL, type) and timing metadata exposed to APIs.

## Key Claims
- Service workers are generic, event-driven, time-limited script contexts running at an origin.
- They provide an event destination for network interception when other destinations do not exist or are inappropriate.
- The specification aims for maximum flexibility (procedural model) at the cost of added complexity.
- Errors must always be recoverable; update processes are designed to avoid unrecoverable states seen in Application Cache.
- Service workers may start and terminate without attached documents, resembling Chrome Event Pages.
- They can handle push notifications, background sync, cross-origin requests, and centralized data updates (e.g., geolocation).
- APIs are almost entirely asynchronous to avoid blocking document/resource loading.

## Entities And Concepts
- **Service Worker**: A web worker executing in the registering client’s origin; manages network interception and offline behavior.
- **State**: One of "parsed", "installing", "installed", "activating", "activated", "redundant".
- **Script URL**: The URL of the service worker script.
- **Type**: Either "classic" or "module"; defaults to "classic".
- **Service Worker Registration**: Contains the service worker and manages lifecycle (installing, waiting, active).
- **Timing Info**: Struct with `startTime`, `fetchEventDispatchTime`, `workerRouterEvaluationStart`, `workerCacheLookupStart`, `workerMatchedRouterSource`, `workerFinalRouterSource`.
- **Event Types**: Lifecycle (`install`, `activate`), functional (`fetch`, etc.), and `message`/`messageerror`.

## Procedures And API Details
- **Skip Waiting**: A flag to skip the waiting phase, allowing immediate activation.
- **Classic Scripts Imported Flag**: Tracks whether classic scripts have been imported.
- **Used Scripts Set**: Prunes unused resources after installation based on old worker map.
- **Event Loop Running**: Defines when a service worker is considered "running".
- **Service Worker Queue**: A parallel queue associated with the worker.
- **Termination Conditions**: No event to handle or abnormal operation (e.g., infinite loops, time limit exceeded).

## Nuance Or Contradictions
- Service workers are described as both similar to Shared Workers and distinct in that they never handle messages from documents; they process events only.
- The specification is a "living document" with unimplemented features and potential changes; citation is discouraged outside of work-in-progress context.
- Timing values are initially 0 but updated during execution; exact semantics depend on WPT results and user agent implementation.

## Candidate Wiki Hints
- **Service Worker Lifecycle**: Document states, events (install/activate), and termination policies.
- **Offline First Architecture**: How service workers intercept fetches and override default network behavior.
- **Timing API Integration**: Use of `DOMHighResTimeStamp` for performance tracing in service worker contexts.
- **Error Recovery Design**: Contrast with Application Cache; emphasis on avoiding unrecoverable states.
- **Event-Driven Model**: Comparison to Chrome Event Pages and Shared Workers; implications for resource conservation.
