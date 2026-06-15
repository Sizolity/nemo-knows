---
title: Chunk 16 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

## Chunk Context

This chunk, titled "8. Acknowledgements," serves as a comprehensive index of terms defined by various web standards (e.g., [CSP-NEXT], [FETCH], [HTML], [WEBIDL]) relevant to the Service Workers specification. It concludes with an IDL Index defining interfaces like `ServiceWorker`, `ServiceWorkerRegistration`, `ServiceWorkerContainer`, and event handlers such as `oninstall` and `onfetch`.

## Local Summary

The section lists specific terminology adopted from external specifications, categorizing them by their source standard (e.g., DOM, Fetch, ECMAScript). Following the term list, it provides normative references to these standards. Finally, it details the Interface Definitions Language (IDL) for key Service Worker APIs, including lifecycle states (`ServiceWorkerState`), registration management, client querying, and event handling within the `ServiceWorkerGlobalScope`.

## Key Claims

- The specification incorporates terms from diverse standards including [CSP-NEXT], [DOM], [ECMASCRIPT], [FETCH], [HTML], [INFRA], [STREAMS], [URL], and [WEBIDL].
- The `ServiceWorker` interface extends `EventTarget` and includes attributes like `scriptURL`, `state`, and event handlers such as `onstatechange`.
- The `ServiceWorkerRegistration` interface manages the lifecycle of a service worker, exposing properties like `installing`, `waiting`, `active`, and methods like `update()` and `unregister()`.
- Service Workers are exposed via the `serviceWorker` attribute on both `Navigator` and `WorkerNavigator`.
- The `ServiceWorkerContainer` interface allows for registration (`register`), retrieval of registrations, and listening to controller changes.
- The `NavigationPreloadManager` enables control over navigation preload headers via methods like `enable()`, `disable()`, and `setHeaderValue()`.
- Client management is handled by the `Clients` interface, which supports querying clients (`get`, `matchAll`) and opening windows (`openWindow`).
- Events within the service worker environment are extended (e.g., `ExtendableEvent`, `InstallEvent`, `FetchEvent`) to support lifecycle hooks like `waitUntil`.

## Entities And Concepts

- **Service Worker Lifecycle States**: `parsed`, `installing`, `installed`, `activating`, `activated`, `redundant` (defined in `ServiceWorkerState`).
- **Registration Update Cache**: Options for caching updates: `"imports"`, `"all"`, `"none"` (defined in `ServiceWorkerUpdateViaCache`).
- **Client Types**: Distinguish between `"window"`, `"worker"`, `"sharedworker"`, and `"all"`.
- **Frame Types**: Categorizes clients as `"auxiliary"`, `"top-level"`, `"nested"`, or `"none"`.
- **Running Status**: Indicates whether a service worker is `"running"` or `"not-running"`.
- **Router Sources**: Defines routing sources such as `"cache"`, `"fetch-event"`, `"network"`, and `"race-network-and-fetch-handler"`.
- **Navigation Preload**: A mechanism to preload navigation resources, managed by `NavigationPreloadManager`.

## Procedures And API Details

### ServiceWorker Registration
To register a service worker:
```webidl
Promise<ServiceWorkerRegistration> register((TrustedScriptURL or USVString) scriptURL, optional RegistrationOptions options = {});
```
- **Parameters**: `scriptURL` (the location of the worker script), `options` (containing `scope`, `type`, and `updateViaCache`).

### Client Querying
To query available clients:
```webidl
Promise<(Client or undefined)> get(DOMString id);
Promise<FrozenArray<Client>> matchAll(optional ClientQueryOptions options = {});
```
- **Options**: `includeUncontrolled` (boolean), `type` (e.g., `"window"`).

### Event Handling in Global Scope
The `ServiceWorkerGlobalScope` exposes event handlers:
- `oninstall`: Called when the service worker is being installed.
- `onactivate`: Called when a new service worker becomes active.
- `onfetch`: Called to handle network requests.
- `onmessage`: Called to handle messages from clients.

### Fetch Event Handling
```webidl
interface FetchEvent : ExtendableEvent {
  [SameObject] readonly attribute Request request;
  readonly attribute Promise<any> preloadResponse;
  DOMString clientId;
  DOMString resultingClientId;
  DOMString replacesClientId;
  Promise<undefined> handled;

  undefined respondWith(Promise<Response> r);
};
```
- **Methods**: `respondWith()` allows the handler to respond to the request.
- **Attributes**: Access to `request`, `clientId`, and `handled` promise.

### Navigation Preload Management
```webidl
interface NavigationPreloadManager {
  Promise<undefined> enable();
  Promise<undefined> disable();
  Promise<undefined> setHeaderValue(ByteString value);
  Promise<NavigationPreloadState> getState();
};
```
- **State**: `enabled` (boolean), `headerValue` (ByteString).

## Nuance Or Contradictions

- The specification distinguishes between different client types (`window`, `worker`) and frame types, which affects how clients are queried and interacted with.
- The `NavigationPreloadManager` is optional in some contexts but required for full navigation preload control, allowing developers to toggle this behavior dynamically.
- Service workers can be exposed via both `Navigator` and `WorkerNavigator`, indicating their availability across different browsing contexts.

## Candidate Wiki Hints

- **Service Worker Lifecycle**: Document the states (`parsed`, `installing`, etc.) and transitions between them.
- **Client Management**: Create a guide on querying and managing clients using the `Clients` interface.
- **Navigation Preload**: Explain how to enable, disable, and configure navigation preload headers.
- **Event Handlers**: Detail the lifecycle events (`oninstall`, `onactivate`, `onfetch`) and their typical use cases.
