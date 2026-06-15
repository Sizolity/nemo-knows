---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context

This chunk covers sections 3.5 through 4 of the Fetch Standard, detailing HTTP header processing logic and the core fetching algorithm. It includes specifications for `Content-Type`, `X-Content-Type-Options`, `Cross-Origin-Resource-Policy`, and `Sec-Purpose` headers, followed by the main steps to initiate a fetch operation.

## Local Summary

The section defines algorithms for extracting MIME types from header lists, handling charset parameters, and legacy encoding extraction. It specifies logic for determining if a response should be blocked based on `nosniff` directives relative to script or style destinations. The `Cross-Origin-Resource-Policy` (CORP) header is detailed with its ABNF, internal check mechanisms, and violation reporting workflows. Finally, the main "To fetch" algorithm outlines request population, early hint processing, preloaded resource handling, header insertion (Accept, Accept-Language), priority setting, and the invocation of the main fetch controller.

## Key Claims

- The `Content-Type` header extraction model differs from standard HTTP models when applied to web content; it prioritizes safety over strict compatibility with legacy features that have caused security vulnerabilities.
- MIME type parameters can typically be safely ignored during extraction unless they define the essence or charset, which are critical for format validation.
- The `X-Content-Type-Options: nosniff` directive triggers a blocking mechanism only for "script-like" and "style" destinations; image destinations are explicitly noted as incompatible with this exploit pattern in current deployments.
- The `Cross-Origin-Resource-Policy` header allows origins to enforce restrictions on resources loaded via "no-cors" requests, supporting values like `same-origin`, `same-site`, and `cross-origin`.
- Fetching algorithms support suspension and resumption of ongoing operations, with specific constraints regarding HTTP cache updates for "no-store" responses.

## Entities And Concepts

- **MIME Type Extraction**: Algorithm to parse `Content-Type` headers, handling multiple values, essence matching, and charset parameters.
- **Legacy Encoding Extraction**: A fallback method to retrieve encoding from a MIME type, returning a fallback encoding if the type is failure or lacks a charset parameter.
- **Nosniff Check**: Logic to determine if a response's `Content-Type` matches the request destination (script/style) when `nosniff` is present.
- **Cross-Origin Resource Policy (CORP)**: Mechanism to check request origin against response URL origin for "no-cors" requests, supporting violation reporting.
- **Fetch Params**: An object encapsulating request details, timing info, callbacks, and controller references during the fetch lifecycle.
- **Preloaded Response Candidate**: A state indicating whether a preloaded resource is available or pending for a specific request.
- **Accept Header Insertion**: Logic to populate `Accept` headers based on destination type (document, image, json, style, text) if missing.

## Procedures And API Details

### Extract MIME Type
1. Initialize `charset`, `essence`, and `mimeType` to null.
2. Decode and split the `Content-Type` header from the list.
3. Iterate through values, parsing each; skip failures or "*/*" essences.
4. Set `mimeType` on the first valid parse.
5. If subsequent parses have a different essence, reset `charset` to null and update `essence`.
6. If `mimeType` parameters lack "charset" but `charset` is non-null, assign it to the parameters.
7. Return failure if no valid `mimeType` was found; otherwise, return the final `mimeType`.

### Determine Nosniff
1. Decode and split `X-Content-Type-Options`.
2. If null, return false.
3. If the first value matches "nosniff" (case-insensitive), return true.
4. Return false.

### Cross-Origin Resource Policy Internal Check
1. If `forNavigation` is true and embedder policy is "unsafe-none", return allowed.
2. Retrieve `Cross-Origin-Resource-Policy` from the response header list.
3. Normalize policy: if not one of `same-origin`, `same-site`, or `cross-origin`, treat as null.
4. Switch on embedder policy value:
   - "unsafe-none": Do nothing (policy remains null).
   - "credentialless": Set policy to `same-origin` if request includes credentials or is navigation.
   - "require-corp": Set policy to `same-origin`.
5. Evaluate based on normalized policy:
   - `null`: Return allowed.
   - `cross-origin`: Return allowed.
   - `same-origin`: Return allowed only if origins match; otherwise blocked.
   - `same-site`: Return allowed only if schemes match (HTTPS/HTTP) and sites are schemelessly same; otherwise blocked.

### To Fetch
1. Assert request mode is "navigate" or early hints handler is null.
2. Initialize task destination and cross-origin isolated capability.
3. Populate request from client details.
4. If client exists, set task destination to client's global object and capture cross-origin isolated capability.
5. If parallel queue is used, start a new one for the task destination.
6. Create fetch timing info with coarsened current time.
7. Initialize fetch params with request, callbacks, and controller references.
8. Convert byte sequence body to a body object if present.
9. Run WebDriver BiDi clone network request body steps.
10. Check for preloaded resources if URL is HTTP(S), mode is safe, client is Window, method is GET, and unsafe flag is unset:
    - Define `onPreloadedResponseAvailable` callback.
    - Invoke consume a preloaded resource algorithm.
    - Update fetch params preloaded response candidate state if found.
11. If `Accept` header is missing:
    - Default to `*/*`.
    - Override for "prefetch" initiator with document value.
    - Otherwise, set based on destination (document, image, json, style, text).
    - Append (`Accept`, value) to header list.
12. If `Accept-Language` is missing and client exists:
    - Use emulated language from WebDriver BiDi if available.
    - Encode and append (`Accept-Language`, encodedEmulatedLanguage).
13. If `Accept-Language` still missing, suggest appending an appropriate value.
14. Set request's internal priority using implementation-defined logic based on priority, initiator, destination, and render-blocking.
15. For subresource requests, create a fetch record and append to client's fetch group.
16. Run main fetch given fetch params.
17. Return fetchParams’s controller.

## Nuance Or Contradictions

- **MIME Type Safety**: The standard explicitly notes that existing web platform features have not always followed the strict MIME type extraction pattern, leading to security vulnerabilities. This suggests a divergence from historical behavior toward a more secure, albeit stricter, model.
- **Image Destination and Nosniff**: The text notes that considering "image" as a target for `nosniff` blocking was not compatible with deployed content, implying that current implementations likely ignore this check for images despite the general algorithmic possibility.
- **Policy Header Matching**: Multiple `Cross-Origin-Resource-Policy` headers or values like `same-site, same-origin` are noted to have specific matching behaviors (e.g., `same-site` ignoring secure transport mismatches) which can result in unexpected "allowed" states if not carefully parsed.

## Candidate Wiki Hints

- **MIME Type Extraction Algorithm**: A reusable procedure for parsing and validating MIME types with charset handling, useful for browser engine documentation.
- **X-Content-Type-Options Security**: Notes on the `nosniff` header's scope (script/style only) and its role in preventing content-type sniffing attacks.
- **Cross-Origin Resource Policy (CORP)**: Detailed breakdown of CORP header values, internal check logic, and interaction with Embedder Policies.
- **Fetch Algorithm Initialization**: The "To fetch" algorithm steps serve as a foundational reference for implementing network request lifecycles in web engines.
