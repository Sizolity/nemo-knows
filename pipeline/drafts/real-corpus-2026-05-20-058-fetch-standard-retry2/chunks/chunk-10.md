---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context
This chunk details the **Main fetch** algorithm (steps 4.1) and introduces the **Override fetch** concept (step 4.2). It covers URL security upgrades, referrer logic, scheme switching to HTTPS via HSTS/SVCB, response tainting (basic, cors, opaque), filtering responses for CORS headers, integrity checks (SRI), timing information collection, and stream transformation for body handling. The chunk concludes by defining the `override fetch` mechanism which allows user agents to intercept requests and return synthetic responses or errors before the standard fetch proceeds.

## Local Summary
The fetch algorithm validates security constraints (local-only URLs, mixed content, CSP) and handles referrer logic. It upgrades HTTP requests to HTTPS for known HSTS hosts. The algorithm branches based on request mode (navigate, no-cors, same-origin) and scheme (data, http). For CORS requests, it manages preflight responses and exposes headers. Finally, it constructs a filtered response, checks integrity metadata, handles timing info, and pipes the body stream to notify completion. `override fetch` allows user agents to return synthetic responses for specific hosts or paths.

## Key Claims
- **Security Upgrades**: Requests are blocked if local-only flags conflict with non-local URLs, mixed content is unsafe, or CSP/Integrity policies block them. HTTP requests are automatically upgraded to HTTPS if the host matches HSTS rules (superdomain or congruent) and DNS resolves an HTTPS RR.
- **Response Tainting**: Responses are categorized into "basic", "cors", or "opaque". This determines how headers are exposed and whether the response body is accessible.
- **CORS Handling**: If `Access-Control-Expose-Headers` contains `*` and credentials mode is not "include", all headers are exposed. Preflight responses clear cache entries if they fail.
- **Integrity Checks**: If integrity metadata is set, the body must match; otherwise, a network error is returned.
- **Timing Info**: Server-timing headers are decoded from `Server-Timing` and reported via resource timing APIs for secure contexts.
- **Stream Transformation**: The response body stream is piped through an identity transform stream to trigger callbacks when the body ends or fails.

## Entities And Concepts
- **Fetch Params**: Object containing request, recursive flag, timing info, controller, etc.
- **Request Object**: Contains current URL, referrer, mode, credentials mode, headers, flags (local-URLs-only, mixed content).
- **Response Object**: Internal response, status, body, header list, tainting, cache state.
- **Filtered Response**: A wrapper exposing only allowed headers based on tainting ("basic", "cors", "opaque").
- **Override Fetch**: An algorithm entry point allowing interception to return a direct response or null.
- **Scheme Fetch / HTTP Fetch**: Sub-algorithms invoked by override fetch for specific schemes.
- **HSTS / SVCB**: Protocols used to determine if an HTTPS upgrade is valid based on DNS records and domain matching.
- **SRI (Subresource Integrity)**: Mechanism to verify response body integrity against metadata.

## Procedures And API Details
- **Main Fetch Steps**:
  1. Validate local URL constraints.
  2. Check CSP violations.
  3. Upgrade URLs if HSTS/SVCB conditions met.
  4. Apply referrer policy and determine referrer value.
  5. Switch scheme to "https" for eligible HTTP hosts.
  6. Branch based on `recursive` flag or initial response candidate (preload).
  7. Handle specific modes: navigate, websocket, webtransport, same-origin, no-cors.
  8. For CORS: manage preflight, set tainting to "cors", run HTTP fetch.
  9. Construct filtered response based on tainting.
 10. Set redirect taint and timing allow passed flag.
 11. Block if mixed content, CSP, MIME type, or nosniff policies fail.
 12. Handle range requests and null body statuses.
 13. Validate integrity metadata; abort if mismatch.
 14. Run fetch response handover: decode `Server-Timing`, update controller timing info.
 15. Queue tasks for end-of-body processing and resource timing reporting.
 16. Pipe body stream through a TransformStream to trigger flush callbacks.
- **Override Fetch**:
  - Input: type ("scheme-fetch" or "http-fetch"), fetchParams, makeCORSPreflight.
  - Logic: Execute potentially override response. If non-null, return immediately.
  - Implementation: Default returns null; user agents may inject synthetic responses (e.g., shims for unsafe.example/widget.js).

## Nuance Or Contradictions
- **DNS Timing**: DNS operations are implementation-defined and might occur earlier or later than traditionally expected to support HSTS upgrades, requiring potential logic unwinding.
- **Range Requests**: APIs traditionally accept ranged responses even if the range wasn't requested, preventing partial data leaks in service workers.
- **Header Exposure**: One of the header names in `Access-Control-Expose-Headers` can be `*`, which matches only headers named `*`, a specific edge case in CORS logic.
- **User Agent Intervention**: The `potentially override response` step is implementation-defined, meaning user agents can silently block or synthesize responses without following the standard fetch flow for certain domains.

## Candidate Wiki Hints
- **Fetch Algorithm Overview**: Explain the high-level steps of fetching resources, including security checks and mode handling.
- **CORS Filtering**: Detail how response headers are filtered based on tainting and `Access-Control-Expose-Headers`.
- **HSTS Upgrades**: Describe the conditions under which HTTP URLs are upgraded to HTTPS via HSTS or SVCB records.
- **Override Fetch Pattern**: Discuss the architecture allowing user agents to intercept fetch requests and return synthetic responses (shimming).
- **Response Tainting**: Define "basic", "cors", and "opaque" states and their impact on browser behavior.
