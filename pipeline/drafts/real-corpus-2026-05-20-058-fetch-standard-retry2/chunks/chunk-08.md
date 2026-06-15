---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context

This chunk covers sections **3.2** and **3.4**, focusing on the `Origin` request header, its serialization grammar, and its usage rules within the Fetch standard. It also details **3.3**, which defines the CORS protocol, including preflight requests, response headers, credential handling, and specific exceptions for certain `Content-Type` values. Section **3.4** briefly outlines the processing model for the `Content-Length` header.

## Local Summary

The `Origin` header indicates the source of a fetch but omits the path, differing from the legacy `Referer`. It is sent based on response tainting ("cors"), request method (non-GET/HEAD), and referrer policy. The chunk provides strict ABNF for serializing origins, enforcing lowercase ASCII schemes/domains and specific IPv6 formatting rules.

The CORS protocol allows cross-origin resource sharing via opt-in headers (`Access-Control-Allow-Origin`, `Allow-Credentials`, etc.). It distinguishes between preflight requests (using `OPTIONS`) and actual requests. The section details how credentials impact header values (e.g., `*` is forbidden for `Allow-Origin` when credentials are included) and lists specific `Content-Type` exceptions for non-preflighted cross-origin requests.

## Key Claims

- The `Origin` header is a path-less version of `Referer`, used for CORS-tainted responses and non-GET/HEAD methods.
- Origin serialization is stricter than RFC 3986: schemes/domains are lowercase ASCII, IPv6 addresses cannot be elided to single zero blocks, and leading zeros are forbidden.
- A CORS request includes an `Origin` header; a CORS-preflight request specifically uses the `OPTIONS` method with `Access-Control-Request-Method` and `Access-Control-Request-Headers`.
- Successful responses to CORS requests can use any status code if they include appropriate headers, while preflight responses are restricted to "ok" statuses (e.g., 200).
- When credentials mode is "include", `Access-Control-Allow-Origin` cannot be `*`, and `Access-Control-Allow-Credentials` must be present with the value `true`.
- The protocol includes exceptions for specific `Content-Type` headers (`application/csp-report`, etc.) to allow non-preflighted cross-origin requests.

## Entities And Concepts

- **Origin Header**: Indicates fetch origin; does not reveal path.
- **CORS Protocol**: Mechanism for cross-origin resource sharing.
- **Preflight Request**: `OPTIONS` request checking CORS support.
- **Credentials Mode**: Controls whether cookies/auth headers are sent ("omit", "same-origin", "include").
- **Access-Control-Allow-Origin**: Header declaring allowed origins (`*` or specific origin).
- **Access-Control-Allow-Credentials**: Header permitting shared responses with credentials (`true`).
- **Access-Control-Expose-Headers**: Header listing response headers exposed to client JavaScript.
- **Referrer Policy**: Determines when `Origin` is sent based on privacy settings (e.g., "no-referrer").
- **Content-Type Exceptions**: Specific MIME types allowed without preflight for security reasons.

## Procedures And API Details

**Appending the `Origin` Header:**
1. Assert request origin is not "client".
2. Byte-serialize the request origin to get `serializedOrigin`.
3. If response tainting is "cors" or mode is "websocket"/"webtransport", append (`Origin`, `serializedOrigin`).
4. Otherwise, if method is neither `GET` nor `HEAD`:
   - Check referrer policy:
     - `"no-referrer"`: Set origin to `null`.
     - `"no-referrer-when-downgrade"`, `"strict-origin"`, `"strict-origin-when-cross-origin"`: If origin scheme is "https" and current URL scheme is not, set origin to `null`.
     - `"same-origin"`: If origins differ, set origin to `null`.
   - Append (`Origin`, `serializedOrigin`) if applicable.

**CORS Header Values (ABNF):**
- `Access-Control-Allow-Origin`: `origin-or-null` or `*`.
- `Access-Control-Allow-Credentials`: `%s"true"` (case-sensitive).
- `Access-Control-Expose-Headers`: List of field names or `*` (only without credentials).
- `Access-Control-Max-Age`: Delta seconds.

## Nuance Or Contradictions

- The `Origin` header is technically a variant of `Referer` but omits the path to prevent leaking specific resource locations.
- A response can technically be "successful" with status 403 if it includes the necessary CORS headers, distinguishing between HTTP success and CORS permission logic.
- `Access-Control-Allow-Origin` values are case-sensitive; `"True"` is invalid even though `"true"` is valid.
- The `Allow` header is explicitly irrelevant for CORS protocol compliance.

## Candidate Wiki Hints

- **Origin Header**: Explain the difference from `Referer`, serialization rules, and when it is sent.
- **CORS Protocol**: Overview of headers (`Allow-Origin`, `Allow-Credentials`, `Expose-Headers`) and preflight logic.
- **Credentials in CORS**: Detail the restrictions on using `*` for origins and the requirement for `Allow-Credentials: true`.
- **Content-Type Exceptions**: List the specific MIME types exempt from preflight requirements.
