---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- **Source**: `raw/web/corpus-2026-05-18/058-fetch-standard.md`
- **Heading**: 2. Read a chunk from reader given readRequest. > 2.2.5. Requests
- **Range**: Lines 1045–1499
- **Scope**: Detailed specification of the `request` input to the fetch algorithm, covering attributes, flags, modes, and their interactions with CORS, CSP, and navigation logic.

Local Summary
This section defines the structure and behavior of a `request` object in the Fetch standard. It details every attribute associated with a request (e.g., method, URL, headers, body, client) and explains their default states. The chunk further elaborates on specialized flags like `keepalive`, `unsafe-request`, and `use-URL-credentials`. A significant portion is dedicated to `mode` (same-origin, cors, no-cors, navigate), `cache mode` behaviors, and the relationship between `initiator`, `destination`, and Content Security Policy (CSP) directives.

Key Claims
- The default method for a request is `GET`.
- The default request mode is `no-cors`, which restricts methods to safe ones and returns an opaque response.
- The default credentials mode is `same-origin`, though this changes to `include` when the mode is `navigate`.
- A request's `origin` starts as `"client"` and resolves to a specific origin during fetching.
- `cache-mode: "default"` allows fetch to inspect the HTTP cache for fresh or stale responses before or alongside network requests.
- The `keepalive` flag allows requests (like those from `navigator.sendBeacon()` or `<img>`) to outlive their environment settings object.

Entities And Concepts
- **Request Attributes**: method, URL, header list, body, client, origin, referrer, credentials mode, cache mode, redirect mode, destination.
- **Modes**: `same-origin`, `cors`, `no-cors`, `navigate`, `websocket`, `webtransport`.
- **Cache Modes**: `default`, `no-store`, `reload`, `no-cache`, `force-cache`, `only-if-cached`.
- **Initiators**: Sources like `fetch`, `script`, `link`, `ping`, etc., mapped to specific CSP directives (e.g., `connect-src`).
- **Destination Types**: Categories like `document`, `script`, `style`, `image` used for CSP and routing.
- **Flags**: `keepalive`, `unsafe-request`, `use-CORS-preflight`, `timing allow failed`.

Procedures And API Details
- **Setting Up a Request**: Start by seeing "Setting up a request" (referenced at the start of the section).
- **CSP Hooks**: The `form-action` CSP directive must hook directly into HTML's navigate or form submission algorithms.
- **Cache Behavior**:
  - `default`: Inspects HTTP cache; returns fresh/stale responses or performs conditional/non-conditional network fetches.
  - `reload`: Ignores cache, creates a normal request, updates cache.
  - `only-if-cached`: Returns cached response or network error (requires same-origin mode).
- **CORS Preflight**: Triggered if event listeners are on an `XMLHttpRequestUpload` or a `ReadableStream` is used in a request.

Nuance Or Contradictions
- **Default Mode Safety**: Although `no-cors` is the default, standards are discouraged from using it for new features due to safety concerns (opaque responses).
- **Credentials vs URL**: Modern specs avoid setting the `use-URL-credentials` flag because putting credentials in URLs is discouraged.
- **Navigational Requests**: When mode is `navigate`, the credentials mode is assumed to be `include`, overriding other values.

Candidate Wiki Hints
- **Page: Fetch Request Modes** – Explain the differences between `cors`, `no-cors`, and `same-origin`, including their impact on response tainting and allowed methods.
- **Page: HTTP Cache Control in Fetch** – Detail how `cache-mode` affects interaction with the HTTP cache (freshness, staleness, revalidation).
- **Page: Content Security Policy and Fetch** – Map request initiators and destinations to CSP directives (e.g., `script-src`, `connect-src`).
- **Page: Request Attributes Reference** – A comprehensive list of all `request` attributes with their default values and usage contexts.
