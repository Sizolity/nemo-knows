---
title: Chunk 17 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
This chunk details the "Acknowledgements" section of a technical specification regarding Service Workers. It provides the formal interface definitions for the `Cache` and `CacheStorage` APIs, along with a comprehensive compatibility matrix listing supported methods across various browsers (Firefox, Chrome, Safari, Edge, Opera) and mobile environments.

Local Summary
The section outlines the Web Cache API, defining how service workers interact with cached resources via the `Cache` interface and manage storage via `CacheStorage`. It lists specific methods for adding, deleting, matching, and retrieving cached responses. A significant portion of the text is dedicated to tracking browser compatibility for these APIs and related objects (like `Client`, `FetchEvent`, and `ExtendableEvent`), noting version requirements and support status for desktop and mobile browsers.

Key Claims
- The `Cache` interface is exposed in Secure Contexts within Window and Worker environments.
- Service workers utilize a unique processing model that differs from other web workers, specifically regarding the use of an environment settings object during script fetching.
- The standard `fetch` algorithms for classic and module worker scripts take `job’s client` as an argument, which is null when passed from the Soft Update algorithm.
- Browser support varies significantly by method; for example, `CacheStorage/match` requires Chrome 54+, while basic `CacheStorage` operations are supported in Chrome 43+.

Entities And Concepts
- **Cache**: An interface representing a cache object where requests can be stored and retrieved.
- **CacheStorage**: An interface providing access to the global Cache API, allowing creation of new caches or retrieving existing ones by name.
- **Service Worker**: A background script that enables caching and network interception.
- **FetchEvent**: An event type triggered when a request is intercepted by a service worker.
- **ExtendableEvent**: A base interface for events like FetchEvent and MessageEvent, allowing extension with custom properties.

Procedures And API Details
- **Cache.match**: Returns a promise resolving to a `Response` or `undefined` based on the provided request.
- **Cache.put**: Stores a response in the cache, returning a promise that resolves when complete.
- **CacheStorage.open**: Opens an existing cache named by the provided string and returns it as a `Promise<Cache>`.
- **CacheStorage.delete**: Permanently deletes a cache with the specified name.
- **CacheQueryOptions**: A dictionary containing flags like `ignoreSearch`, `ignoreMethod`, and `ignoreVary` to customize query behavior.

Nuance Or Contradictions
The document notes that certain behaviors are not fully specified yet and will be addressed in the HTML Standard via future issues and pull requests. Specifically, the use of a "to-be-created environment settings object" is highlighted as a temporary measure necessitated by the unique processing model of service workers compared to other web workers.

Candidate Wiki Hints
- **Web Cache API**: A page detailing how to cache resources using the `Cache` and `CacheStorage` interfaces.
- **Service Worker Lifecycle**: Explaining the interaction between `job’s client` and the script fetching algorithms during updates.
- **Browser Compatibility Matrix for Service Workers**: A reference table summarizing version support for specific API methods across major browsers.
