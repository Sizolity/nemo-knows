---
title: Service Worker Security Constraints
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

# Service Worker Security Constraints

Service Workers operate within strict security boundaries designed to prevent unauthorized access and ensure the integrity of offline functionality. These constraints are enforced by the browser engine before a worker is allowed to execute or intercept requests.

## Secure Context Requirement

The primary security gatekeeper for Service Workers is the **secure context** requirement. A Service Worker must be registered and executed within a secure environment, typically defined as:

- HTTPS connections
- `localhost` or IP addresses in the `127.0.0.0/8` and `::1/128` ranges (development exceptions)

Both the Service Worker script itself and the client initiating the registration must be served from a secure context. If either is accessed via an insecure HTTP connection, the browser will reject the registration attempt. This ensures that network requests intercepted by the worker cannot be modified by third parties on the same network.

## Origin Relativity and Scope

Security is further reinforced through **origin relativity**. A Service Worker's scope defaults to the directory containing its script file (e.g., a script at `/js/sw.js` controls `/js/`). This prefix-based matching ensures that the worker can only intercept requests intended for its registered origin.

The scope can be explicitly expanded using the `Service-Worker-Allowed` HTTP response header, but this expansion is still bound by the originating secure context and CORS policies.

## Request Handling Restrictions

When handling network requests, Service Workers are subject to specific restrictions:

- **CORS Enforcement**: Cross-Origin Resource Sharing (CORS) rules apply strictly to resources accessed from within the worker. The worker cannot read response bodies from cross-origin requests unless explicitly permitted by headers.
- **Path Restrictions**: Requests are validated against the worker's scope and allowed patterns.
- **Script Loading**: The `importScripts()` API is subject to origin restrictions, preventing the execution of scripts from untrusted or insecure origins.

## Caching Security

The Cache API (`self.caches`) provides an isolated storage layer distinct from the browser's HTTP cache. This isolation enhances security by:

- Preventing accidental overwrites of cached resources by other applications.
- Allowing atomic writes with rollback on failure, ensuring data integrity.
- Supporting different response filters (Basic, CORS Filtered, Opaque Filtered) to manage access control per stored resource.

## Error Handling and State Recovery

To maintain system stability under attack or failure, Service Workers implement robust error recovery mechanisms. Unlike older application caching models that could leave the browser in an unrecoverable broken state, Service Workers handle installation or activation errors by moving to a redundant state rather than failing completely. This ensures that even if a worker script fails to load or install correctly, the user experience remains functional via fallback mechanisms.
