---
title: Chunk 07 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---
Chunk Context
This chunk covers Section 6 (Security Considerations) and Section 7 (Extensibility) of the Service Workers specification. It details requirements for secure contexts, Content Security Policy enforcement, origin relativity restrictions, path limitations, CORS handling for cached resources, header requirements for script requests, implementer concerns regarding plugins and legacy code, privacy implications of persistent storage, and mechanisms for extending the API via functional events.

Local Summary
Service workers are restricted to secure contexts (typically HTTPS) and must execute within the registering client's origin. They enforce Content Security Policies on their scripts and restrict access to CDN-hosted resources unless specific CORS headers are present. Path restrictions prevent a single script from serving multiple unrelated scopes, though this can be overridden via `Service-Worker-Allowed` headers. The spec mandates headers like `Service-Worker` and specific MIME types for scripts. Privacy is addressed by requiring the clearing of persistent storage maps when users clear their data. Extensibility is supported through partial interface definitions for registrations, functional events, and event handlers.

Key Claims
- Service workers must execute in secure contexts; clients must also be secure contexts to register or interact with them.
- `localhost`, `127.0.0.0/8`, and `::1/128` are exceptions for development purposes.
- Content Security Policy (CSP) headers on the script resource are enforced or monitored by the user agent.
- Service workers execute in the registering client's origin, preventing hosting on CDNs directly.
- Resources from other origins can be fetched and cached only if appropriate CORS headers are set; stored responses are either CORS filtered or opaque filtered.
- Path restrictions limit a service worker script to its specific scope path unless overridden by `Service-Worker-Allowed`.
- Service worker scripts must include the `Service-Worker` header and be served with a JavaScript MIME type.
- Plugins should not load via service workers because the embedding worker cannot handle their security origins.
- Legacy networking stack code may require auditing for interactions with service workers.
- Persistent storages (registration map, cache name to map, script resource map) must be cleared when users purge data.
- Specifications can extend the API using partial interface definitions on `ServiceWorkerRegistration`, `ExtendableEvent`, and `ServiceWorkerGlobalScope`.

Entities And Concepts
- Secure Context
- Content Security Policy (CSP) / Content-Security-Policy-Report-Only
- Origin Relativity
- Service Worker Registration Scope
- Import Scripts
- Caches API
- CORS Filtered Response / Opaque Filtered Response
- Path Restriction
- Service-Worker-Allowed Header
- Service-Worker Header
- Persistent Storage (Registration Map, Cache Map, Script Resource Map)
- Extensibility
- Partial Interface Definition
- Functional Event
- ExtendableEvent

Procedures And API Details
- **Run Service Worker Algorithm**: Enforces CSP if `Content-Security-Policy` header matches policy; monitors if `Content-Security-Policy-Report-Only` matches.
- **importScripts(urls)**:
  - Checks worker state ("parsed" or "installing").
  - Validates script resource map against URL.
  - Sets service-workers mode to "none".
  - Determines cache mode based on registration update mode, force bypass flag, or staleness.
  - Fetches request and updates the response map if safe.
- **Event.respondWith(r)**: Can accept Response objects whose corresponding responses are basic filtered, CORS filtered, or opaque filtered, but cannot create them programmatically.
- **Fire Functional Event**: Invoked by specifications to dispatch a functional event to the active worker of a service worker registration.

Nuance Or Contradictions
- While path restriction offers some protection for multi-user content on the same origin, origins are the only hard security boundary; sites should use different origins for secure isolation.
- Unlike same-origin resources managed in Cache as basic filtered responses, off-origin resources stored in Cache are CORS or opaque filtered and cannot be meaningfully created programmatically.
- Plugins are explicitly discouraged from loading via service workers due to origin handling limitations in the Handle Fetch algorithm.

Candidate Wiki Hints
- Service Worker Security Requirements
- Secure Contexts and HTTPS
- Content Security Policy for Service Workers
- Origin Relativity and CDN Restrictions
- CORS and Cached Off-Origin Resources
- Path Restriction and Service-Worker-Allowed
- Privacy and Persistent Storage Clearing
- Extending the Service Worker API
