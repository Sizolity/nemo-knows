---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
- Location: Section 3. HTTP extensions, subsections 3.2 (`Origin` header) and 3.3 (CORS protocol), followed by a brief note on `Content-Length`.
- Scope: Defines the `Origin` header serialization rules, conditions for including it in requests, CORS preflight mechanics, response headers for cross-origin sharing, credential handling, and specific exceptions for certain content types.

Local Summary
This chunk details how the `Origin` header is serialized and when it must be included in fetch requests based on response tainting, request method, and referrer policy. It then introduces the CORS protocol as an opt-in mechanism to allow cross-origin resource sharing while preventing data leakage. The text covers preflight requests (`OPTIONS`), relevant headers like `Access-Control-Allow-Origin`, handling of credentials, and specific rules for header exposure. Finally, it notes limited exceptions for certain non-safelisted content types in the CORS context.

Key Claims
- The `Origin` header indicates where a fetch originates and is used for all HTTP fetches with response tainting "cors" or methods other than `GET`/`HEAD`.
- The `Origin` header serialization is more constrained than RFC 3986: it uses lower-case ASCII without percent encoding, forbids leading zeros in IPv4, requires lowercase hex for IPv6, and limits the use of `::` to at most 6 blocks.
- CORS is an opt-in protocol necessary to prevent leaking data from responses behind firewalls and sensitive data with credentials.
- A successful CORS response can have any status code if it includes the appropriate headers; a preflight response must be "ok" (e.g., 200 or 204).
- When `Access-Control-Allow-Credentials` is present, `Access-Control-Allow-Origin` cannot be `*` and must match the request origin exactly.
- The `Content-Length` header extraction logic validates that values are consistent ASCII digits; inconsistent values result in failure.

Entities And Concepts
- **Origin Header**: Indicates fetch source; serialized as scheme + host + optional port.
- **Serialized Origin**: Strict ABNF grammar for IPv4, IPv6, and domain names (lowercase only).
- **Response Tainting**: Determines if CORS headers are required ("cors").
- **CORS Protocol**: Mechanism allowing cross-origin resource sharing via specific headers.
- **Preflight Request**: An `OPTIONS` request checking server CORS support.
- **Access-Control-Allow-Origin**: Response header specifying which origins can access the resource.
- **Credentials Mode**: "include" or "same-origin"; impacts whether cookies/authorization headers are sent and how CORS rules apply.
- **Referrer Policy**: Affects whether an `Origin` header is included (e.g., "no-referrer").

Procedures And API Details
**Appending `Origin` Header Steps:**
1. Assert request origin is not "client".
2. Byte-serialize the request origin.
3. If response tainting is "cors" or mode is "websocket"/"webtransport", append (`Origin`, serializedOrigin).
4. Otherwise, if method is not `GET`/`HEAD`:
   - If mode is not "cors", apply referrer policy rules:
     - "no-referrer": Set origin to `null`.
     - "no-referrer-when-downgrade", "strict-origin", "strict-origin-when-cross-origin": Set origin to `null` if scheme downgrades from HTTPS to non-HTTPS.
     - "same-origin": Set origin to `null` if not same origin as current URL.
   - Append (`Origin`, serializedOrigin).

**CORS Response Headers:**
- **Preflight**: Must return 2xx/3xx status with `Access-Control-Allow-Origin`, `Allow-Methods`, `Allow-Headers`, and optionally `Max-Age`.
- **Actual Request**: Can return any status with `Access-Control-Allow-Origin` (matching origin or `*`), `Allow-Credentials` (if needed), and `Expose-Headers` (list of names or `*` if no credentials).

**Credential Logic:**
- If `credentials` mode is "include", `Access-Control-Allow-Origin` must not be `*`.
- `Access-Control-Allow-Credentials: true` is required when sharing with credentials.
- `Set-Cookie` headers are functional only if the request includes credentials and CORS allows it.

Nuance Or Contradictions
- **Origin Header Ambiguity**: The `Origin` header is present in all non-simple requests (non-GET/HEAD) regardless of whether they participate in the CORS protocol, making it unreliable as a sole indicator of CORS participation without checking other headers or context.
- **Status Code Flexibility**: A "successful" response to a non-preflight CORS request can technically be 403 if the server intends to share it (by including headers), which contradicts typical HTTP semantics where 4xx implies failure, though the spec clarifies that side channels might still leak data.
- **Case Sensitivity**: `Access-Control-Allow-Credentials: true` is byte-case-sensitive; "True" or "TRUE" are invalid.

Candidate Wiki Hints
- **CORS Protocol Overview**: Explain the opt-in nature, preflight requests, and key headers (`Allow-Origin`, `Allow-Credentials`, `Expose-Headers`).
- **Origin Header Serialization**: Detail the strict ABNF constraints compared to RFC 3986 (lowercase, IPv6 limits).
- **Referrer Policy and Origin**: How policies like "no-referrer" affect the `Origin` header inclusion.
- **Credentials in CORS**: Rules governing `Access-Control-Allow-Credentials` and why `*` is forbidden with credentials.
- **CORS Exceptions**: List of allowed non-safelisted content types (`application/csp-report`, etc.).
