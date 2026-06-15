---
title: Cors Preflight Cache
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Cors Preflight Cache

The **CORS Preflight Cache** is a mechanism within the Fetch Standard that optimizes Cross-Origin Resource Sharing (CORS) interactions by storing the results of preflight requests. When a browser receives a `GET` or `POST` request to a cross-origin resource, it first checks if the response headers are "CORS-safelisted." If they are not, the browser must perform a preflight check using an `OPTIONS` request to verify that the server permits the actual operation.

The Fetch Standard ensures that this preflight logic is handled efficiently by caching the results of successful `OPTIONS` requests. This prevents the need to repeat expensive network round-trips for identical requests. The standard categorizes responses into three states—`basic`, `cors`, and `opaque`—where CORS responses rely on specific header exposure rules managed during these cached interactions.

This mechanism operates within the broader **http-fetch-lifecycle**, interacting with connection management and security policies to ensure that data does not leak between different origins. The architecture supports strict MIME type extraction and enforces security constraints like Cross-Origin-Embedder-Policy (COEP) alongside standard CORS checks. By utilizing deferred fetching quotas and managing request-response objects effectively, the system maintains memory efficiency while handling incremental reading of body streams.
