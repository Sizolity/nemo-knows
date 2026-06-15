---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
This chunk details the HTTP Fetch Standard's mechanisms for connection management, specifically focusing on how user agents obtain connections from a pool, handle network partition keys, manage HTTP cache partitions, block bad ports, and process cookies. It also covers timing information for connections to prevent leaking reuse details and defines infrastructure for SameSite cookie modes.

Local Summary
A user agent maintains an ordered set of connections identified by a network partition key, origin, and credentials. Obtaining a connection involves checking the pool, finding proxies (including non-standard ones like WPAD/PAC), and creating new connections if necessary. Connection timing is recorded and coarsened to prevent exposing reuse details. Network partition keys are derived from the top-level site. Bad ports (like echo or daytime) trigger blocking. Cookies are managed via specific headers with infrastructure for SameSite modes, HttpOnly flags, and path serialization.

Key Claims
- Connections are keyed by a network partition key, origin, and credentials; TLS session identifiers are not reused across connections with different credentials.
- Obtaining a connection involves a sequence: check pool, find proxies, resolve hosts (racing IPv4/IPv6), create connection, and potentially add to the pool.
- Connection timing info includes domain lookup, connection start/end, secure connection start, and ALPN protocol, which are clamped and coarsened to hide reuse details.
- A port is considered "bad" if listed in a specific table (e.g., 0, 1, 7, 9), leading to request blocking for HTTP(S) requests targeting those ports.
- MIME types like `audio/*`, `image/*`, `video/*`, and `text/csv` are blocked when the destination is script-like.
- The `Cookie` header is appended based on retrieved cookies filtered by security context, host, path, and SameSite mode.
- The `Set-Cookie` header is parsed and stored independently for each occurrence, with garbage collection of cookies.

Entities And Concepts
- **Connection Pool**: An ordered set of connections associated with a user agent.
- **Network Partition Key**: A tuple consisting of a site and an optional implementation-defined value (second key).
- **Connection Timing Info**: A struct tracking domain lookup times, connection start/end times, secure connection start time, and ALPN negotiated protocol.
- **Bad Port**: Ports listed in a standard table (e.g., 0, 1, 7) that should block fetching requests.
- **HTTP Cache Partition**: The unique HTTP cache associated with a network partition key.
- **SameSite Mode**: Determined for requests to handle cookie security contexts (strict-or-less, lax-or-less, unset-or-less).

Procedures And API Details
- **Obtain Connection**:
  1. Check pool for existing connection matching key, origin, and credentials.
  2. If none or `requireUnreliable` is true, find proxies (defaulting to "DIRECT").
  3. Resolve hosts (handling failures).
  4. Create connection in parallel if multiple proxies exist; select one return value.
  5. Add to pool unless `new` setting is "yes-and-dedicated".
- **Create Connection**:
  - Set connection start time.
  - Establish transport (TCP/UDP/TLS), handling proxy tunnels and ALPN.
  - Handle WebTransport requirements (SETTINGS_ENABLE_WEBTRANSPORT, H3_DATAGRAM).
  - Manage TLS certificates (client cert if credentials=true, custom verification via webTransportHashes).
- **Record Timing**:
  - Set connection end time immediately after establishing transport/TLS handshake sufficient to request resource.
  - Secure connection start time set before handshake.
  - For HTTP/3, connection start and secure connection start times must be equal.
- **Determine Network Partition Key**:
  - Get top-level origin from environment or creation URL.
  - Obtain site for the origin.
  - Return tuple (topLevelSite, secondKey).
- **Append Cookie Header**:
  - Check cookie disable configuration.
  - Determine SameSite mode and isSecure.
  - Retrieve cookies based on host, path, httpOnlyAllowed, and SameSite.
  - Serialize and append to header list.

Nuance Or Contradictions
- Connection management details are intentionally vague regarding nuances like IP address selection (favoring IPv6) or retry logic, leaving discretion to implementers.
- The "second key" in network partition keys is noted as evolving (referencing issue #1035).
- Timing info clamping ensures reused connections do not expose timing details; specific rules apply for TLS False Start and early data scenarios.
- Proxy handling allows non-standard technologies like WPAD/PAC to influence the "DIRECT" or proxy value.

Candidate Wiki Hints
- **Connection Pooling Strategy**: How browsers manage connection reuse, pooling, and lifecycle based on credentials and origins.
- **Network Partition Key Definition**: The tuple structure (site, secondKey) used to isolate network resources for privacy and caching.
- **Connection Timing Coarsening**: Techniques to hide connection reuse patterns by clamping timestamps.
- **Bad Port Blocking List**: A curated list of historically dangerous or obsolete ports (e.g., echo, daytime) that fetch should block.
- **Cookie Infrastructure in Fetch**: How the Fetch API handles `Cookie` and `Set-Cookie` headers, including SameSite logic and path serialization.
