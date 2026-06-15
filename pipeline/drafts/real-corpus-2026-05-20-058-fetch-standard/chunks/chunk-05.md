---
title: Chunk 05 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- **Heading**: 2. Read a chunk from reader given readRequest. > 2.2.5. Requests
- **Line Range**: 1501–1595 of the source document `raw/web/corpus-2026-05-18/058-fetch-standard.md`.

Local Summary
This chunk defines request classification (subresource, non-subresource, navigation), specifies the algorithm to compute `redirect-taint` and serialize request origins, details how to clone a request while handling bodies and WebDriver IDs, describes adding Range headers for partial responses, warns about security implications of combining multiple responses, and outlines serialization for response URLs to prevent leaking redirect targets. It also includes logic for checking Cross-Origin-Embedder-Policy credential allowance.

Key Claims
- A subresource request has destinations like "audio", "font", "image", "script", etc., while a non-subresource request includes "document", "embed", "worker", etc.
- A navigation request is restricted to "document", "embed", "frame", "iframe", or "object".
- `redirect-taint` computation iterates through a request's URL list, comparing origins of subsequent URLs against the previous one and the request's origin to determine if the taint is "same-origin", "same-site", or "cross-site".
- Serializing a request origin returns "null" if the redirect-taint is not "same-origin"; otherwise, it returns the serialized origin.
- Cloning a request creates a new copy excluding the original body and WebDriver id, then generates a random UUID for the new WebDriver id and clones the body if it exists.
- Adding a Range header constructs a value like `bytes=<first>-<last>` (or just `<first>-` if last is omitted) and appends it to the request's header list.
- Combining multiple responses into one logical resource is historically a source of security bugs and should undergo security review.
- Serializing a response URL for reporting uses the first URL in the list to avoid leaking information about redirect targets.
- Cross-Origin-Embedder-Policy allows credentials if the request mode is not "no-cors", the client is null, the embedder policy is not "credentialless", or specific origin/redirect-taint conditions are met.

Entities And Concepts
- **Request Classification**: subresource request, non-subresource request, navigation request.
- **Redirect Taint**: same-origin, same-site, cross-site.
- **Serialization**: request origin serialization, byte-serialization, response URL serialization for reporting.
- **Request Cloning**: handling body cloning and WebDriver id generation.
- **Range Headers**: inclusive byte range notation (`bytes=<first>-<last>`).
- **Security Concepts**: redirect target leakage prevention, Cross-Origin-Embedder-Policy credential allowance, partial response security bugs.

Procedures And API Details
- **Compute Redirect Taint**:
  1. Assert request's origin is not "client".
  2. Initialize `lastURL` to null and `taint` to "same-origin".
  3. Iterate through each URL in the request's URL list:
     - If `lastURL` is null, set it to the current URL and continue.
     - If current URL's origin is not same site with `lastURL`'s origin AND request's origin is not same site with `lastURL`'s origin, return "cross-site".
     - If current URL's origin is not same origin with `lastURL`'s origin AND request's origin is not same origin with `lastURL`'s origin, set `taint` to "same-site".
     - Set `lastURL` to the current URL.
  4. Return the final `taint`.

- **Serialize Request Origin**:
  1. Assert request's origin is not "client".
  2. If redirect-taint is not "same-origin", return "null".
  3. Return the serialized request's origin.

- **Clone a Request**:
  1. Create a copy of the request excluding body and WebDriver id.
  2. Generate a random UUID for the new request's WebDriver id.
  3. If the original body is non-null, clone the body for the new request.
  4. Return the new request.

- **Add Range Header**:
  1. Assert `last` is not given or `first` <= `last`.
  2. Initialize `rangeValue` with `bytes=`.
  3. Serialize and isomorphic encode `first`, append to `rangeValue`.
  4. Append `-` (0x2D) to `rangeValue`.
  5. If `last` is given, serialize and isomorphic encode it, append to `rangeValue`.
  6. Append (`Range`, `rangeValue`) to the request's header list.

- **Serialize Response URL for Reporting**:
  1. Assert response's URL list is not empty.
  2. Copy the first URL from the response's URL list (not the current response URL) to avoid leaking redirect targets.
  3. Set username and password to empty strings for that URL.
  4. Return the serialization of the URL with the fragment excluded.

- **Check Cross-Origin-Embedder-Policy Credential Allowance**:
  1. Assert request's origin is not "client".
  2. If request mode is not "no-cors", return true.
  3. If request's client is null, return true.
  4. If request's client's policy container's embedder policy value is not "credentialless", return true.
  5. If request's origin is same origin with current URL's origin AND redirect-taint is not "same-origin", return true.
  6. Return false.

Nuance Or Contradictions
- The definition of a navigation request overlaps significantly with non-subresource requests (both include "document", "embed"), suggesting navigation is a subset of non-subresource requests in terms of destination types, but distinguished by specific intent or handling not fully detailed here.
- The `redirect-taint` computation relies on comparing origins across the URL list; if any step triggers a "cross-site" condition, it immediately returns that status, prioritizing strict cross-origin detection over cumulative site-level analysis.
- Serializing a response URL specifically avoids using the current response URL to prevent leaking redirect targets, implying that the current URL might differ from the final resolved URL after redirects.

Candidate Wiki Hints
- **Request Classification**: Define and differentiate subresource, non-subresource, and navigation requests with their respective destination types.
- **Redirect Taint Algorithm**: Document the step-by-step logic for computing `redirect-taint` values ("same-origin", "same-site", "cross-site").
- **Request Origin Serialization**: Explain conditions under which a request origin serializes to "null" versus returning the actual origin, focusing on redirect-taint implications.
- **Request Cloning**: Detail the process of cloning a request, specifically handling body cloning and WebDriver id regeneration.
- **Range Headers**: Describe the format and construction of Range headers for partial content retrieval.
- **Response URL Serialization**: Explain the security rationale for using the first URL in the list rather than the current response URL when reporting.
- **Cross-Origin-Embedder-Policy**: Outline the conditions under which credentials are allowed or disallowed based on request mode, client state, and policy settings.
