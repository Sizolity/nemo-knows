# Chunk Outline

Source: `raw/web/corpus-2026-05-18/061-service-workers.md`

## Chunk 01

- Lines: 1-497
- Primary heading path: Document
- Heading coverage:
  - Document
  - Service Workers
  - Service Workers > Fetch Metadata
  - Service Workers > Retrieved Text
  - 1. Motivations
  - 2. Model
  - 2. Model > 2.1. Service Worker
  - 2. Model > 2.1. Service Worker > 2.1.1. Lifetime
  - 2. Model > 2.1. Service Worker > 2.1.2. Events
  - 2. Model > 2.2. Service Worker Timing

## Chunk 02

- Lines: 498-958
- Primary heading path: 2. Model > 2.3. Service Worker Registration
- Heading coverage:
  - 2. Model > 2.3. Service Worker Registration
  - 2. Model > 2.3. Service Worker Registration > 2.3.1. Lifetime
  - 2. Model > 2.4. Service Worker Client
  - 2. Model > 2.5. Control and Use
  - 2. Model > 2.5. Control and Use > 2.5.1. The window client case
  - 2. Model > 2.5. Control and Use > 2.5.2. The worker client case
  - 2. Model > 2.6. Task Sources
  - 2. Model > 2.7. User Agent Shutdown
  - 3. Client Context
  - 3. Client Context > 3.1. ServiceWorker
  - 3. Client Context > 3.1. ServiceWorker > 3.1.1. Getting ServiceWorker instances
  - 3. Client Context > 3.1. ServiceWorker > 3.1.2. scriptURL
  - 3. Client Context > 3.1. ServiceWorker > 3.1.3. state
  - 3. Client Context > 3.1. ServiceWorker > 3.1.4. postMessage(message, transfer)
  - 3. Client Context > 3.1. ServiceWorker > 3.1.5. postMessage(message, options)
  - 3. Client Context > 3.1. ServiceWorker > 3.1.6. Event handler
  - 3. Client Context > 3.2. ServiceWorkerRegistration
  - 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.1. Getting ServiceWorkerRegistration instances
  - 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.2. installing
  - 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.3. waiting
  - 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.4. active
  - 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.5. navigationPreload
  - 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.6. scope
  - 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.7. updateViaCache

## Chunk 03

- Lines: 959-1432
- Primary heading path: 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.8. update()
- Heading coverage:
  - 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.8. update()
  - 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.9. unregister()
  - 3. Client Context > 3.2. ServiceWorkerRegistration > 3.2.10. Event handler
  - 3. Client Context > 3.3. navigator.serviceWorker
  - 3. Client Context > 3.4. ServiceWorkerContainer
  - 3. Client Context > 3.4. ServiceWorkerContainer > 3.4.1. controller
  - 3. Client Context > 3.4. ServiceWorkerContainer > 3.4.2. ready
  - 3. Client Context > 3.4. ServiceWorkerContainer > 3.4.3. register(scriptURL, options)
  - 3. Client Context > 3.4. ServiceWorkerContainer > 3.4.4. getRegistration(clientURL)
  - 3. Client Context > 3.4. ServiceWorkerContainer > 3.4.5. getRegistrations()
  - 3. Client Context > 3.4. ServiceWorkerContainer > 3.4.6. startMessages()
  - 3. Client Context > 3.4. ServiceWorkerContainer > 3.4.7. Event handlers
  - 3. Client Context > 3.5. Events
  - 3. Client Context > 3.6. NavigationPreloadManager
  - 3. Client Context > 3.6. NavigationPreloadManager > 3.6.1. enable()
  - 3. Client Context > 3.6. NavigationPreloadManager > 3.6.2. disable()
  - 3. Client Context > 3.6. NavigationPreloadManager > 3.6.3. setHeaderValue(value)
  - 3. Client Context > 3.6. NavigationPreloadManager > 3.6.4. getState()
  - 4. Execution Context
  - 4. Execution Context > 4.1. ServiceWorkerGlobalScope
  - 4. Execution Context > 4.1. ServiceWorkerGlobalScope > 4.1.1. clients
  - 4. Execution Context > 4.1. ServiceWorkerGlobalScope > 4.1.2. registration
  - 4. Execution Context > 4.1. ServiceWorkerGlobalScope > 4.1.3. serviceWorker
  - 4. Execution Context > 4.1. ServiceWorkerGlobalScope > 4.1.4. skipWaiting()
  - 4. Execution Context > 4.1. ServiceWorkerGlobalScope > 4.1.5. Event handlers

## Chunk 04

- Lines: 1433-1903
- Primary heading path: 4. Execution Context > 4.2. Client
- Heading coverage:
  - 4. Execution Context > 4.2. Client
  - 4. Execution Context > 4.2. Client > 4.2.1. url
  - 4. Execution Context > 4.2. Client > 4.2.2. frameType
  - 4. Execution Context > 4.2. Client > 4.2.3. id
  - 4. Execution Context > 4.2. Client > 4.2.4. type
  - 4. Execution Context > 4.2. Client > 4.2.5. postMessage(message, transfer)
  - 4. Execution Context > 4.2. Client > 4.2.6. postMessage(message, options)
  - 4. Execution Context > 4.2. Client > 4.2.7. visibilityState
  - 4. Execution Context > 4.2. Client > 4.2.8. focused
  - 4. Execution Context > 4.2. Client > 4.2.9. ancestorOrigins
  - 4. Execution Context > 4.2. Client > 4.2.10. focus()
  - 4. Execution Context > 4.2. Client > 4.2.11. navigate(url)
  - 4. Execution Context > 4.3. Clients
  - 4. Execution Context > 4.3. Clients > 4.3.1. get(id)
  - 4. Execution Context > 4.3. Clients > 4.3.2. matchAll(options)
  - WindowClient objects whose browsing context has been
  - WindowClient objects whose browsing context has
  - Client objects whose associated service worker
  - Client objects whose associated service worker > 4.3.3. openWindow(url)
  - Client objects whose associated service worker > 4.3.3. openWindow(url) > 4.3.4. claim()
  - Client objects whose associated service worker > 4.4. ExtendableEvent

## Chunk 05

- Lines: 1904-2370
- Primary heading path: Client objects whose associated service worker > 4.4. ExtendableEvent > 4.4.1. event.waitUntil(f)
- Heading coverage:
  - Client objects whose associated service worker > 4.4. ExtendableEvent > 4.4.1. event.waitUntil(f)
  - Client objects whose associated service worker > 4.5. InstallEvent
  - Client objects whose associated service worker > 4.5. InstallEvent > 4.5.1. event.addRoutes(rules)
  - Client objects whose associated service worker > 4.6. FetchEvent
  - Client objects whose associated service worker > 4.6. FetchEvent > 4.6.1. event.request
  - Client objects whose associated service worker > 4.6. FetchEvent > 4.6.2. event.preloadResponse
  - Client objects whose associated service worker > 4.6. FetchEvent > 4.6.3. event.clientId
  - Client objects whose associated service worker > 4.6. FetchEvent > 4.6.4. event.resultingClientId
  - Client objects whose associated service worker > 4.6. FetchEvent > 4.6.5. event.replacesClientId
  - Client objects whose associated service worker > 4.6. FetchEvent > 4.6.6. event.handled
  - Client objects whose associated service worker > 4.6. FetchEvent > 4.6.7. event.respondWith(r)
  - 1. Set end-of-body to true.
  - 1. error newStream with a TypeError.
  - 1. error newStream with a TypeError. > 4.7. ExtendableMessageEvent
  - 1. error newStream with a TypeError. > 4.7. ExtendableMessageEvent > 4.7.1. event.data
  - 1. error newStream with a TypeError. > 4.7. ExtendableMessageEvent > 4.7.2. event.origin
  - 1. error newStream with a TypeError. > 4.7. ExtendableMessageEvent > 4.7.3. event.lastEventId
  - 1. error newStream with a TypeError. > 4.7. ExtendableMessageEvent > 4.7.4. event.source
  - 1. error newStream with a TypeError. > 4.7. ExtendableMessageEvent > 4.7.5. event.ports
  - 1. error newStream with a TypeError. > 4.8. Events
  - 5. Caches
  - 5. Caches > 5.1. Constructs
  - 5. Caches > 5.2. Understanding Cache Lifetimes
  - 5. Caches > 5.3. self.caches
  - 5. Caches > 5.3. self.caches > 5.3.1. caches

## Chunk 06

- Lines: 2371-2819
- Primary heading path: 5. Caches > 5.4. Cache
- Heading coverage:
  - 5. Caches > 5.4. Cache
  - 5. Caches > 5.4. Cache > 5.4.1. match(request, options)
  - 5. Caches > 5.4. Cache > 5.4.2. matchAll(request, options)
  - 5. Caches > 5.4. Cache > 5.4.3. add(request)
  - 5. Caches > 5.4. Cache > 5.4.4. addAll(requests)
  - 5. Caches > 5.4. Cache > 5.4.5. put(request, response)
  - 5. Caches > 5.4. Cache > 5.4.6. delete(request, options)
  - 5. Caches > 5.4. Cache > 5.4.7. keys(request, options)
  - 5. Caches > 5.5. CacheStorage
  - 5. Caches > 5.5. CacheStorage > 5.5.1. match(request, options)
  - 5. Caches > 5.5. CacheStorage > 5.5.2. has(cacheName)
  - 5. Caches > 5.5. CacheStorage > 5.5.3. open(cacheName)
  - 5. Caches > 5.5. CacheStorage > 5.5.4. delete(cacheName)
  - 5. Caches > 5.5. CacheStorage > 5.5.5. keys()
  - 6. Security Considerations

## Chunk 07

- Lines: 2820-3013
- Primary heading path: 6. Security Considerations > 6.1. Secure Context
- Heading coverage:
  - 6. Security Considerations > 6.1. Secure Context
  - 6. Security Considerations > 6.2. Content Security Policy
  - 6. Security Considerations > 6.3. Origin Relativity
  - 6. Security Considerations > 6.3. Origin Relativity > 6.3.1. Origin restriction
  - 6. Security Considerations > 6.3. Origin Relativity > 6.3.2. importScripts(urls)
  - 6. Security Considerations > 6.4. Cross-Origin Resources and CORS
  - 6. Security Considerations > 6.5. Path restriction
  - 6. Security Considerations > 6.6. Service worker script request
  - 6. Security Considerations > 6.7. Implementer Concerns
  - 6. Security Considerations > 6.8. Privacy
  - 7. Extensibility
  - 7. Extensibility > 7.1. Define API bound to Service Worker Registration
  - 7. Extensibility > 7.2. Define Functional Event
  - 7. Extensibility > 7.3. Define Event Handler
  - 7. Extensibility > 7.4. Firing Functional Events

## Chunk 08

- Lines: 3014-3350
- Primary heading path: 7. Extensibility > Appendix A: Algorithms
- Heading coverage:
  - 7. Extensibility > Appendix A: Algorithms

## Chunk 09

- Lines: 3352-3762
- Primary heading path: 7. Extensibility > Appendix A: Algorithms
- Heading coverage:
  - 7. Extensibility > Appendix A: Algorithms

## Chunk 10

- Lines: 3764-4232
- Primary heading path: 7. Extensibility > Appendix A: Algorithms
- Heading coverage:
  - 7. Extensibility > Appendix A: Algorithms

## Chunk 11

- Lines: 4234-4762
- Primary heading path: 7. Extensibility > Appendix A: Algorithms
- Heading coverage:
  - 7. Extensibility > Appendix A: Algorithms

## Chunk 12

- Lines: 4764-5240
- Primary heading path: 7. Extensibility > Appendix A: Algorithms
- Heading coverage:
  - 7. Extensibility > Appendix A: Algorithms

## Chunk 13

- Lines: 5242-5308
- Primary heading path: 7. Extensibility > Appendix B: Extended HTTP headers
- Heading coverage:
  - 7. Extensibility > Appendix B: Extended HTTP headers

## Chunk 14

- Lines: 5309-5393
- Primary heading path: 8. Acknowledgements
- Heading coverage:
  - 8. Acknowledgements

## Chunk 15

- Lines: 5395-5818
- Primary heading path: 8. Acknowledgements
- Heading coverage:
  - 8. Acknowledgements

## Chunk 16

- Lines: 5820-6594
- Primary heading path: 8. Acknowledgements
- Heading coverage:
  - 8. Acknowledgements

## Chunk 17

- Lines: 6596-7300
- Primary heading path: 8. Acknowledgements
- Heading coverage:
  - 8. Acknowledgements

## Chunk 18

- Lines: 7302-8020
- Primary heading path: 8. Acknowledgements
- Heading coverage:
  - 8. Acknowledgements

## Chunk 19

- Lines: 8022-8210
- Primary heading path: 8. Acknowledgements
- Heading coverage:
  - 8. Acknowledgements

