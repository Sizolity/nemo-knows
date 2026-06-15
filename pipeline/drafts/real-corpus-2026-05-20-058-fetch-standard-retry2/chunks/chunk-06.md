---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
The chunk covers sections 2.2.6 (Responses), 2.2.7 (Miscellaneous), 2.3 (Authentication entries), 2.4 (Fetch groups), and 2.5 (Resolving domains). It details the structure of a `Response`, its lifecycle states, filtering mechanisms, cloning procedures, caching classifications, authentication entry structures, fetch group state management, and origin resolution logic.

Local Summary
A `Response` evolves over time with fields becoming available sequentially. It possesses an associated type ("basic", "cors", "default", "error", "opaque", "opaqueredirect"), a status code (default 200), headers, body, cache state, and various flags like `aborted`, `timing allow passed`, and `range-requested`. Filtered responses provide limited views to prevent information leakage. The chunk also defines `fetch` destinations, authentication entry tuples, fetch group records for managing concurrent requests, and the `resolve an origin` algorithm which maps origins to IP addresses via DNS or implementation-defined operations.

Key Claims
- A response's type defaults to "default" unless specified otherwise.
- Responses over HTTP/2 always have an empty status message.
- Filtered responses expose a limited view; their internal response is accessible only for legacy reasons like feeding image data to decoders.
- New specifications should not build on opaque filtered responses or opaque-redirect filtered responses due to architectural limitations.
- The `location` URL algorithm returns null if the status is not a redirect status.
- A fetch group holds lists of fetch records and deferred fetch records to manage state during document unloading or inactivity.

Entities And Concepts
- Response (type, status, headers, body, cache state)
- Filtered response (basic, CORS, opaque, opaque-redirect)
- Network error
- Aborted network error
- Fetch destination ("fetch" or string)
- Authentication entry (username, password, realm)
- Proxy-authentication entry
- Fetch group (fetch records, deferred fetch records)
- Fetch record (request, controller)
- Deferred fetch record (request, notify invoked, invoke state)
- Resolve an origin (network partition key, origin, IP addresses)

Procedures And API Details
- Cloning a response:
  1. If filtered, return a new identical filtered response with a cloned internal response.
  2. Copy the response except for its body.
  3. If the original body is non-null, clone it and set it on the new response.
  4. Return the new response.
- Translating a potential destination:
  1. If "fetch", return the empty string.
  2. Assert it is a destination.
  3. Return the destination.
- Resolving an origin:
  1. If host is an IP, return the host as a set.
  2. If host is "localhost" or "localhost.", return « ::1, 127.0.0.1 ».
  3. Perform implementation-defined operation (e.g., DNS query) to get a set of IP addresses.
  4. Return failure if unsuccessful.

Nuance Or Contradictions
- The `Location` header parsing requires an absolute URL with fragment if the response was constructed via the Response constructor (null URL).
- Opaque and opaque-redirect filtered responses are nearly indistinguishable from network errors, making properties like `response.ok` useless for them.
- Origin resolution order of IP addresses can differ between invocations.

Candidate Wiki Hints
- Fetch API: Response Object Structure and Types
- Security: Filtered Responses and Information Leakage Prevention
- Networking: Origin Resolution and DNS Caching in Fetch
