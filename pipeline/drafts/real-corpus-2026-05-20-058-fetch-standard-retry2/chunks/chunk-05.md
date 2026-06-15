---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- Heading path: `2. Read a chunk from reader given readRequest. > 2.2.5. Requests`
- Line range: 1501–1595

Local Summary
This section defines request classifications (subresource vs non-subresource vs navigation), the algorithm to compute redirect-taint, origin serialization rules for reporting, request cloning mechanics (including WebDriver ID handling), Range header construction, and a security warning regarding partial response aggregation. It concludes with an algorithm to check Cross-Origin-Embedder-Policy credential allowance.

Key Claims
- Request destinations determine their classification: "audio", "audioworklet", "font", etc., are subresources; "document", "embed", etc., are non-subresources; a subset of the latter forms navigation requests.
- Redirect-taint is computed by iterating through a request’s URL list, comparing origins between consecutive URLs and the current origin to yield "same-origin", "same-site", or "cross-site".
- Origin serialization returns `null` if redirect-taint is not "same-origin" to avoid leaking redirect target information.
- Cloning a request copies all fields except body and WebDriver id; the new request receives a freshly generated random UUID as its WebDriver id.
- Range headers denote inclusive byte ranges where `first=0` and `last=500` implies 501 bytes.
- Features combining multiple responses into one logical resource are historically prone to security bugs and require security review.

Entities And Concepts
- **Subresource request**: Destination in `["audio", "audioworklet", "font", "image", "json", "manifest", "paintworklet", "script", "style", "text", "track", "video", "xslt", ""]`.
- **Non-subresource request**: Destination in `["document", "embed", "frame", "iframe", "object", "report", "serviceworker", "sharedworker", "worker"]`.
- **Navigation request**: Subset of non-subresources with destination in `["document", "embed", "frame", "iframe", "object"]`.
- **Redirect-taint**: A classification ("same-origin", "same-site", "cross-site") derived from origin comparisons across a request’s URL list.
- **Cross-Origin-Embedder-Policy**: A policy container value (`"credentialless"`) checked to determine if credentials are allowed for a given request.

Procedures And API Details
- `compute redirect-taint`: Iterates `request`’s URL list, tracking `lastURL`. Returns "cross-site" immediately if consecutive origins differ from the current origin in site terms; otherwise updates taint to "same-site" if origins differ but sites match, finally returning current taint.
- `serialize request origin`: Asserts origin is not `"client"`. Returns `"null"` if redirect-taint is not `"same-origin"`, otherwise returns serialized origin.
- `byte-serialize request origin`: Wraps the result of `serialize request origin` with isomorphic encoding.
- `clone request`: Creates a copy excluding body and WebDriver id; generates a new random UUID for the WebDriver id; clones the body if non-null.
- `add range header`: Asserts valid bounds, constructs `bytes=first-last` (or `bytes=first-` if last omitted), serializes components isomorphically, and appends to request headers.
- `check COEP allows credentials`: Asserts origin not `"client"`. Returns `true` unless mode is `"no-cors"` AND client exists AND policy container embedder policy value is `"credentialless"` AND redirect-taint is `"same-origin"`.

Nuance Or Contradictions
- The definition of "Range header denotes an inclusive byte range" includes a clarifying example: `first=0`, `last=500` represents 501 bytes, which aligns with standard HTTP range semantics (inclusive endpoints).
- The warning about features combining multiple responses into one logical resource explicitly flags them as historical sources of security bugs.

Candidate Wiki Hints
- Page: **Fetch Requests Classification** – Define subresource, non-subresource, and navigation request destinations.
- Page: **Request Redirect-Taint Computation** – Detail the algorithm for determining origin taint across URL lists.
- Page: **Origin Serialization Security** – Explain why `null` is returned when redirect-taint is not "same-origin" to prevent redirect leakage.
- Page: **Request Cloning Mechanics** – Document the exclusion of body and WebDriver id, and UUID generation.
- Page: **Range Header Construction** – Steps for adding Range headers and byte-range semantics.
- Page: **Cross-Origin-Embedder-Policy Credential Check** – Logic flow for determining credential allowance under COEP.
