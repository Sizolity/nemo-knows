---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
This chunk details the `main fetch` algorithm (section 4.1) and the `override fetch` algorithm (section 4.2) within the Fetch Standard. It covers security checks, URL upgrades, response tainting logic, filtering, timing information handling, and specific behaviors for preloaded responses or data URLs.

Local Summary
The `main fetch` algorithm initializes a request, performs various security and policy checks (local-only flags, CSP, mixed content, referrer policy, HSTS), and determines the appropriate response type based on the URL scheme and CORS mode. It handles parallel execution for non-recursive requests, applies filtering to expose headers appropriately, manages timing information, validates integrity metadata, and processes the response body via a TransformStream. The `override fetch` algorithm allows user agents to intervene before standard fetching occurs, enabling features like content shimming or blocking unsafe domains.

Key Claims
- **Security Checks**: Before fetching, the algorithm checks for local-only flags, Content Security Policy violations, mixed content risks, bad ports, and Integrity Policy blocks.
- **URL Upgrades**: Requests are upgraded to HTTPS if the host matches an HSTS domain or has a matching SVCB HTTPS RR.
- **Response Tainting**: The response is tainted as "basic", "cors", or "opaque" depending on the request's mode, URL scheme, and redirect settings.
- **Header Exposure**: For CORS requests, `Access-Control-Expose-Headers` determines which headers are exposed; if credentials are not included but headers contain `*`, all headers are exposed.
- **Integrity Validation**: If integrity metadata is present, the response body must match it; otherwise, a network error is returned.
- **Override Fetch**: Allows user agents to intercept requests and return custom responses (e.g., shimming resources) or errors before standard fetching proceeds.

Entities And Concepts
- **Fetch Params**: The structured data containing the request, client info, and callbacks for handling response end-of-body.
- **Response Tainting**: A state (`basic`, `cors`, `opaque`) indicating how a response is exposed to JavaScript contexts.
- **HSTS (HTTP Strict Transport Security)**: Mechanism to enforce HTTPS connections via known host domain matching.
- **SVCB (Service Binding)**: Protocol for specifying alternative server bindings, including HTTPS requirements.
- **CORS (Cross-Origin Resource Sharing)**: Security mechanism controlling cross-origin resource access via headers like `Access-Control-Allow-Origin`.
- **TransformStream**: Used to wrap the response body stream to notify when the end is reached (`processResponseEndOfBody`).

Procedures And API Details
- **Main Fetch Steps**:
  1. Initialize `response` as null.
  2. Check local-only flag; if set and URL is not local, return network error.
  3. Run CSP violation reports.
  4. Upgrade request to trustworthy URL if applicable (HSTS/SVCB).
  5. Handle referrer policy and determination.
  6. Set scheme to "https" for HSTS/SVCB matches.
  7. If recursive is false, run steps in parallel; otherwise, return current response.
  8. Determine response source (preloaded, data URL, same-origin, no-cors, HTTP(S), etc.).
  9. Apply CORS filtering based on `Access-Control-Expose-Headers`.
  10. Set internal response redirect taint and timing flags.
  11. Block responses if blocked by mixed content, CSP, MIME type, or nosniff.
  12. Handle ranged responses (return network error if range requested but header missing).
  13. Nullify body for HEAD/CONNECT methods or null body status.
  14. Validate integrity metadata; abort on mismatch.
- **Fetch Response Handover**:
  - Extract `Server-Timing` headers if in a secure context.
  - Update timing info and mark resource timing.
  - Queue tasks for reporting timing and processing response end-of-body.
  - Pipe response body through a TransformStream to handle completion notifications.
- **Override Fetch**:
  - Executes `potentially override response` which returns null by default or a custom response/error.
  - Supports "scheme-fetch" (runs `scheme fetch`) and "http-fetch" (runs `HTTP fetch`).

Nuance Or Contradictions
- **DNS Operations Timing**: DNS resolution for HTTPS RRs may happen before connection attempts, requiring implementation-defined logic to unwind earlier steps if the scheme needs upgrading.
- **Ranged Responses**: Traditionally, APIs accept ranged responses even without a range request header, but the standard prevents partial responses from being provided to APIs that didn't explicitly request a range.
- **User Agent Intervention**: The default behavior of `potentially override response` is trivial (return null), but user agents often implement complex logic like blocking unsafe domains while shimming specific resources.

Candidate Wiki Hints
- **Fetch Algorithm Overview**: Summarize the high-level flow of fetching resources, including security checks and response handling.
- **Response Tainting Explained**: Detail the differences between "basic", "cors", and "opaque" responses and their impact on JavaScript access.
- **HSTS and URL Upgrades**: Explain how HSTS and SVCB protocols force HTTPS upgrades based on host matching.
- **CORS Header Exposure**: Describe how `Access-Control-Expose-Headers` interacts with credentials mode to determine visible headers.
- **Fetch Response Body Handling**: Outline the use of TransformStreams for monitoring response body completion and integrity validation.
