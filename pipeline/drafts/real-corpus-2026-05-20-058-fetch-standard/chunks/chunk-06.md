---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
This chunk covers sections 2.2.6 (Responses), 2.2.7 (Miscellaneous), 2.3 (Authentication entries), 2.4 (Fetch groups), and the beginning of 2.5 (Resolving domains) within the Fetch Standard. It details the structure, lifecycle flags, and filtering logic for `Response` objects, network errors, and fetch group state management.

Local Summary
A `fetch` call returns a `response` that evolves over time. Responses have specific types ("basic", "cors", "default", "error", "opaque", "opaqueredirect") and associated attributes like status, headers, body, and CORS-exposed header lists. The standard distinguishes between filtered responses (limited views for security) and internal responses. It also defines fetch groups to manage active requests and deferred invocations, and outlines the domain resolution process involving IP address sets.

Key Claims
- A response's type defaults to "default" unless specified otherwise.
- Responses over HTTP/2 connections always have an empty status message.
- `response.ok` may return useless results for filtered responses like opaque types.
- Cloning a filtered response involves cloning its internal response if the body is non-null.
- Fetch groups manage active `fetch records` and `deferred fetch records` until termination.
- Resolving an origin typically involves DNS queries and returns a set of IP addresses.

Entities And Concepts
- **Response**: The result of a fetch, evolving over time with fields like status, headers, body, and type.
- **Filtered Response**: A response offering a limited view (e.g., excluding headers) to prevent information leakage; includes "basic", "cors", "opaque", and "opaqueredirect" types.
- **Network Error**: A response with type "error", status 0, and null body.
- **Fetch Group**: Holds fetch records and deferred fetch records within an environment settings object.
- **Fetch Record**: Contains a `request` and a `controller`.
- **Deferred Fetch Record**: Maintains state to invoke a fetch later (e.g., on document unload).

Procedures And API Details
- **Cloning a Response**:
  1. If the response is filtered, return a new identical filtered response cloning the internal response.
  2. Copy the response except for its body.
  3. Clone the body if it is non-null.
  4. Return the new response.
- **Location URL Algorithm**: Used for redirect handling; returns null if status is not a redirect, extracts the `Location` header, parses the URL (handling fragments), and returns the location.
- **Resolving an Origin**:
  1. Returns host IPs if the host is an IP address.
  2. Returns `::1, 127.0.0.1` for "localhost".
  3. Performs implementation-defined operations (e.g., DNS queries) to return a set of IP addresses.
  4. Returns failure otherwise.

Nuance Or Contradictions
- **Filtered vs. Internal**: While filtered responses expose limited data, specification algorithms may access the internal response for legacy reasons (e.g., feeding image decoders), but new APIs should avoid this to prevent leaks.
- **Opaque Responses**: `opaque` and `opaqueredirect` filtered responses are nearly indistinguishable from network errors because they lack headers, status, and body info accessible to scripts.
- **Caching DNS Results**: The results of resolving an origin may be cached, but implementations vary on whether they can account for the partition key in local caching or rely solely on DNS server caching.

Candidate Wiki Hints
- [Fetch Standard / Response Types](https://wiki.local/fetch/response-types)
- [Filtered Responses Security Model](https://wiki.local/fetch/filtered-responses)
- [Fetch Group Lifecycle](https://wiki.local/fetch/groups)
