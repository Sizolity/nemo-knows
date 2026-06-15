---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Group Context

This document synthesizes the Fetch Standard specification covering the complete lifecycle of a network request and response. It spans from initial infrastructure (URLs, HTTP methods, headers) through connection management, security policies (CORS, CORP), the core fetching algorithms (scheme fetches, redirects, caching), and finally the Fetch API types (`Headers`, `Request`, `Response`). The notes bridge low-level network mechanics with high-level API usage patterns.

# Cross-Chunk Summary

The specification describes a request as an interaction between a user agent and a remote server. This process begins with **Infrastructure** (defining URLs, HTTP methods, and headers), moves to **Connection Management** (partitioning connections by origin/credentials, handling bad ports), and applies security protocols like **CORS** and **Content-Type** safety checks. The core logic resides in the **Fetching** algorithms, which handle scheme resolution (HTTP, HTTPS, Data, Blob), caching strategies (validation, negative caching), redirect logic, and service worker integration. Finally, the **Fetch API** section defines the JavaScript classes (`Headers`, `Request`, `Response`) that expose these low-level operations to developers.

# Repeated Or Central Claims

- **Security by Default**: The standard prioritizes security over legacy compatibility. This is evident in strict MIME type extraction (ignoring dangerous parameters), blocking requests on "bad ports" (e.g., telnet, ftp), enforcing CORS preflight checks for unsafe methods, and using `nosniff` directives to prevent content-type sniffing attacks.
- **Connection Pooling**: Connections are not treated as generic resources but are strictly partitioned by a "Network Partition Key" derived from the site origin and credentials (secure/insecure). This ensures that cookies and other sensitive data do not leak between different origins or security contexts.
- **Response Tainting**: Responses are categorized into three distinct states: `basic`, `cors`, and `opaque`. This tainting mechanism dictates whether headers are exposed to client JavaScript, whether the body is readable, and how redirects are handled.
- **CORS as an Exception**: Cross-Origin Resource Sharing is not a default behavior; it requires explicit opt-in via headers (`Access-Control-Allow-Origin`). The protocol distinguishes between preflight requests (checking support) and actual requests, with specific rules for credentials and exposed headers.
- **Caching Complexity**: Caching is sophisticated but constrained by browser limitations. It supports standard modes (`default`, `no-store`, `reload`) and advanced strategies like "stale-while-revalidate" and negative caching (caching network errors). Request bodies must be cloned to support redirects or authentication retries without stream exhaustion.
- **Fetch API Abstraction**: The JavaScript API provides a high-level interface that abstracts the underlying fetch algorithms. However, the classes (`Request`, `Response`) still maintain internal state regarding timing info, controller references, and headers, reflecting the complexity of the underlying network operations.

# Important Local Details

## Connection and Network Logic
- **Network Partition Key**: Composed of a "site" (scheme + host) and an optional implementation-defined second key. This tuple identifies connections and determines HTTP cache partitions.
- **Bad Ports**: Fetching is explicitly blocked if the URL scheme is HTTP(S) and the port matches a list of "bad ports" (e.g., 21, 23).
- **Timing Info**: Connections track timestamps for domain lookup, connection start/end, and secure connection establishment. A "clamp and coarsen" algorithm resets these times to default values to prevent exposing reused connection details to the client.
- **WebTransport**: For WebTransport connections, specific ALPN settings (`SETTINGS_ENABLE_WEBTRANSPORT`) are required on HTTP/3 transports.

## Header Processing
- **Origin Header**: Unlike `Referer`, the `Origin` header omits the path. It is appended based on response tainting ("cors") or request method (non-GET/HEAD) combined with referrer policy. Serialization rules enforce lowercase ASCII schemes and forbid leading zeros in IPv6 addresses.
- **Content-Type Extraction**: The algorithm parses headers to extract MIME types, prioritizing safety. It ignores parameters unless they define the essence or charset. Legacy encodings are extracted as a fallback.
- **CORS Headers**: `Access-Control-Allow-Origin` accepts an origin string or `*`. When credentials are included (`Access-Control-Allow-Credentials: true`), `*` is forbidden for the allow-origin value. `Access-Control-Expose-Headers` can list specific headers or `*` (though `*` only matches a header literally named `*`).

## Fetch Algorithms
- **Main Fetch**: Orchestrates security upgrades (HSTS/SVCB to HTTPS), referrer logic, and response filtering. It branches based on request mode (`navigate`, `same-origin`, `no-cors`) and scheme.
- **Scheme Fetch**: Handles specific URL schemes:
  - `about:blank`: Returns empty body with `OK`.
  - `data:` URLs: Processed via a specific processor returning the serialized MIME type.
  - `blob:` URLs: Only allow `GET`; range requests calculate byte ranges for slicing.
- **Redirect Handling**: Supports manual, follow, and error modes. Redirects update timing info, strip non-wildcard CORS headers on cross-origin jumps, and normalize methods (e.g., POST to GET on 301/302). Max redirect count is 20.
- **CORS Preflight Cache**: Maintains a cache for preflight responses (`OPTIONS` requests) to avoid repeated network calls. Entries are appended if the method/headers match cached entries, subject to TAO (Token Authentication Object) checks.

## API Classes
- **Headers**: Represents a collection of header entries. Supports iteration and appending.
- **Request**: Encapsulates request details including timing info, callbacks, controller references, and body. Handles `BodyInit` unions for body data.
- **Response**: Wraps the internal response object, exposing status, headers, and body. Supports integrity checks (SRI) where the body must match metadata.

# Gaps Or Cautions

- **Implementation Discretion**: The specification intentionally leaves certain details vague to allow implementer discretion. Notable areas include Proxy Auto-Config (PAC) handling, IP address racing logic for connection establishment, and specific DNS timing operations.
- **File Scheme Ambiguity**: The document explicitly notes that `file:` URLs are left as an exercise for the reader and their behavior is "unfortunate" or undefined in the current standard context.
- **MIME Type Safety vs. Legacy**: While the spec defines a strict MIME type extraction model for safety, it acknowledges that existing web platform features have not always followed this pattern, implying potential divergence from historical browser behavior.
- **Partial Content Limitations**: Although HTTP standards allow partial content caching, the text notes that browser implementations do not widely support this, which may lead to unexpected behavior regarding range requests.
- **Referrer Policy Nuance**: The handling of `Origin` headers is heavily dependent on Referrer Policy settings (`no-referrer`, `strict-origin`, etc.). Misconfiguring these policies can inadvertently leak or hide origin information contrary to security expectations.
