---
title: Chunk 15 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

## Chunk Context
This chunk covers sections 5.1 through 5.3 of the Fetch specification, detailing the `Headers` class, `BodyInit` unions, and the `Body` mixin interface. It also includes algorithmic steps for validating headers, handling CORS guards (specifically "request-no-cors"), extracting request/response bodies from various input types (Blob, BufferSource, FormData, etc.), and consuming body streams into different JavaScript value types.

## Local Summary
The text defines the `Headers` API with methods like `append`, `delete`, `get`, `set`, and `has`, emphasizing validation rules based on header guards ("immutable", "request", "response", "none", "request-no-cors"). It outlines specific logic for appending headers in CORS contexts, where privileged no-CORS request-headers are removed upon modification by unprivileged code. The chunk then transitions to body handling, defining `BodyInit` types and the extraction/consumption algorithms that convert byte sequences or streams into `ArrayBuffer`, `Blob`, `FormData`, `json()`, or `text`. Finally, it details the `Body` mixin methods (`arrayBuffer()`, `blob()`, `bytes()`, `formData()`, `json()`, `text()`) and their underlying consumption logic, including MIME type detection and multipart parsing rules.

## Key Claims
- A `Headers` object maintains an associated header list and a guard state ("immutable", "request", "response", "none", or "request-no-cors").
- Validation of headers throws a `TypeError` if the name/value is invalid or if the guard is "immutable".
- For "request-no-cors" guards, appending headers requires checking for no-CORS-safelisted status; privileged headers are removed if modified by unprivileged code.
- The `BodyInit` type union includes `Blob`, `BufferSource`, `FormData`, `URLSearchParams`, `USVString`, and `ReadableStream`.
- The `consume body` algorithm ensures a promise is returned, rejecting with `TypeError` if the object is unusable (body disturbed/locked).
- The `formData()` method handles "multipart/form-data" and "application/x-www-form-urlencoded" MIME types, parsing parts into entries based on `Content-Disposition` headers.

## Entities And Concepts
- **Headers Class**: Represents HTTP headers with associated guards for security and immutability.
- **HeadersGuard**: States including "immutable", "request", "response", "none", and "request-no-cors".
- **BodyInit**: A union type for request/response body sources (e.g., `Blob`, `FormData`, `ReadableStream`).
- **BodyMixin**: An interface mixin providing methods to read the body as different types (`arrayBuffer()`, `blob()`, `text()`, etc.).
- **CORS Guards**: Specific logic for "request-no-cors" headers involving safelisted checks and removal of privileged headers.
- **MIME Type Extraction**: Logic to determine content type from headers (e.g., `multipart/form-data`, `application/json`).

## Procedures And API Details
- **`Headers.append(name, value)`**: Normalizes the value, validates against guards, handles "request-no-cors" merging logic, and removes privileged headers if needed.
- **`Headers.delete(name)`**: Validates (with a dummy value), checks guard constraints, deletes from list, and cleans up privileged headers in "request-no-cors" mode.
- **`BodyInit` Extraction**: Converts inputs like `Blob` (via stream), `FormData` (multipart encoding), or `URLSearchParams` into a readable stream with an associated type.
- **`consume body` Algorithm**: Reads the body stream, resolves promises with converted data (e.g., `ArrayBuffer`, `Blob`), or rejects if the stream is disturbed/locked.
- **`formData()` Parsing**: Parses multipart parts using the boundary parameter; parts with `filename` become `File` objects, others become strings.

## Nuance Or Contradictions
- The specification notes that steps for "request-no-cors" are not shared in a simple way because a fake value cannot always succeed for CORS-safelisted headers.
- The `formData()` implementation is noted as a "rough approximation," with a more detailed parsing specification pending, particularly regarding the handling of `_charset_` parameters and default content types.
- The `bytes()` method can reject with a `RangeError`, while `arrayBuffer()` and `blob()` can also fail under specific conditions (e.g., if the stream is disturbed).

## Candidate Wiki Hints
- **Headers API**: Documenting the `Headers` interface, including initialization from objects or arrays, guard states, and CORS-specific restrictions.
- **Body Consumption**: Explaining how to convert request/response bodies into various JavaScript types using the `Body` mixin methods.
- **FormData Parsing**: Detailing the logic behind `formData()` for handling multipart and form-encoded data, including `File` object creation.
