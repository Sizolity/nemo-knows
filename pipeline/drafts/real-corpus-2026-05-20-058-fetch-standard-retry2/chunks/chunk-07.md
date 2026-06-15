---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
This chunk covers the Fetch Standard's sections on connection management (2.6), network partition keys (2.7), HTTP cache partitions (2.8), port blocking rules (2.9), MIME type restrictions for scripts (2.10), and cookie infrastructure including `Cookie`/`Set-Cookie` headers (3.1).

Local Summary
The specification defines how a user agent manages a pool of connections identified by network partition keys, origins, and credentials. It details the process of obtaining or creating connections, handling proxies, negotiating ALPN protocols, and recording timing information to avoid exposing reused connection details. The chunk also outlines logic for determining cache partitions based on these keys, blocking requests from bad ports (e.g., telnet, ftp), restricting certain MIME types for script destinations, and managing cookies via secure headers and infrastructure algorithms.

Key Claims
- A user agent maintains an ordered set of connections identified by a network partition key, origin, and credentials boolean.
- Connection timing info includes timestamps for domain lookup, connection start/end, and secure connection start, along with the ALPN negotiated protocol.
- The "clamp and coarsen" algorithm ensures reused connection details are not exposed by resetting times to a default or coarse value.
- Connections can be created with specific settings regarding unreliable transport (e.g., HTTP/3) and custom certificate verification for WebTransport.
- Network partition keys consist of a site and an optional implementation-defined second key, used to determine HTTP cache partitions.
- Fetching is blocked if the URL scheme is HTTP(S) and the port is listed as a "bad port" (e.g., 23 telnet).
- Responses with MIME types like `audio/`, `image/`, `video/`, or `text/csv` are blocked for script-like destinations.
- Cookie infrastructure includes logic to determine same-site modes (`lax-or-less`, `strict-or-less`) and serialize cookies for the `Cookie` header.

Entities And Concepts
- **Connection Pool**: An ordered set of connections maintained by a user agent.
- **Network Partition Key**: A tuple (site, secondKey) used to identify connections and cache partitions.
- **Connection Timing Info**: A struct tracking timing metrics like domain lookup duration and ALPN protocol.
- **Bad Port**: Ports listed in a table (e.g., 21 ftp, 23 telnet) that trigger request blocking.
- **Same-Site Mode**: Values such as `lax-or-less` or `strict-or-less` determining cookie behavior across origins.
- **WebTransport**: A feature requiring specific ALPN settings (`SETTINGS_ENABLE_WEBTRANSPORT`) on HTTP/3 connections.

Procedures And API Details
- **Obtain a Connection**: Involves checking the pool, finding proxies (defaulting to "DIRECT"), resolving hosts, and racing connection attempts.
- **Create a Connection**: Sets timing info, establishes transport (handling TLS certificates and ALPN), and returns the connection or failure.
- **Record Connection Timing Info**: Defines when `connection end time` and `secure connection start time` are recorded relative to handshake completion and early data usage.
- **Append Cookie Header**: Retrieves cookies based on security context, host/path, and same-site mode, then serializes them into the header value.
- **Parse Set-Cookie Headers**: Iterates through response headers, parsing each `Set-Cookie`, storing it, and performing garbage collection for the host.

Nuance Or Contradictions
- The specification notes that connection management details are intentionally vague to allow implementer discretion, particularly regarding proxy auto-config (PAC) and IP address racing.
- Timing info clamping logic differs based on whether a cross-origin isolated capability is present, affecting how times are coarsened.
- ALPN Protocol ID identification must account for tunnelled protocols when a proxy is configured versus the first hop to the proxy.

Candidate Wiki Hints
- Connection Pooling and Partition Keys
- HTTP/3 and WebTransport Setup
- Port Blocking List Reference
- Same-Site Cookie Logic
