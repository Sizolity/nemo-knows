---
title: Chunk 15 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk details the specification for the **Headers** and **Body** classes within the Fetch API. It covers the `Headers` interface (construction, mutation methods like `append`, `delete`, `get`, `set`, and guard states), the `BodyInit` union types for request bodies, and the `Body` mixin methods (`arrayBuffer`, `blob`, `json`, `text`, etc.) along with MIME type extraction logic.

# Local Summary

The document defines the `Headers` interface as a container for header name-value pairs with associated guards ("immutable", "request", "response", "request-no-cors", "none") to control mutation and CORS safety. It specifies validation steps, normalization, and the removal of privileged headers in "request-no-cors" contexts. The chunk also defines `BodyInit` unions (Blob, BufferSource, FormData, URLSearchParams, USVString, ReadableStream) and the `Body` mixin, which provides methods to read response/request bodies as streams or parsed types (JSON, text, etc.).

# Key Claims

- **Headers Object**: Has an associated header list and a guard state.
- **Guard States**: "immutable", "request", "response", "request-no-cors", "none".
- **Mutation Logic**: `set` replaces the first header with the given name; `append` adds to the list.
- **CORS Safety**: In "request-no-cors" guards, privileged no-CORS request headers are removed upon modification.
- **BodyInit Unions**: Supports Blob, BufferSource, FormData, URLSearchParams, USVString, and ReadableStream.
- **MIME Types**: `FormData` maps to `multipart/form-data` or `application/x-www-form-urlencoded`.
- **Consumption**: The `consume body` algorithm handles reading streams into promises (ArrayBuffer, Blob, etc.).

# Entities And Concepts

- **Headers**: Interface for managing HTTP headers.
- **HeadersInit**: Union of record or sequence for initializing headers.
- **Guard**: A state controlling header mutation and CORS behavior.
- **BodyInit**: Union type for request body sources (Blob, FormData, string, etc.).
- **ReadableStream**: Used to stream body data.
- **FormData**: Represents form data with multipart or URL-encoded boundaries.
- **BodyMixin**: Provides access methods (`arrayBuffer`, `blob`, `json`, `text`) for response/request bodies.
- **MIME Type Extraction**: Logic to derive content types from header lists.

# Procedures And API Details

### Headers Initialization
```javascript
const meta = { "Content-Type": "text/xml", "Breaking-Bad": "<3" };
new Headers(meta);
// Equivalent to:
const meta2 = [ ["Content-Type", "text/xml"], ["Breaking-Bad", "<3"] ];
new Headers(meta2);
```

### Mutation Methods
- `headers.append(name, value)`: Appends a header.
- `headers.delete(name)`: Removes a header (validates name first).
- `headers.set(name, value)`: Replaces existing headers with the same name.
- `headers.get(name)`: Returns comma-space separated values.
- `headers.getSetCookie()`: Returns list of Set-Cookie values.
- `headers.has(name)`: Checks existence.

### Validation Steps for Headers
1. Validate name/value types (throw TypeError if invalid).
2. Check guard state (throw TypeError if "immutable").
3. Check forbidden headers based on guard ("request" or "response").
4. Return true/false based on safety.

### BodyInit Processing
- **Blob**: Sets source to object, length to size, type to attribute value.
- **FormData**: Encodes as multipart/form-data with boundary.
- **URLSearchParams**: Serializes as `application/x-www-form-urlencoded`.
- **USVString**: Encodes as UTF-8 text/plain.

### Body Mixin Methods
- `arrayBuffer()`: Returns Promise<ArrayBuffer>.
- `blob()`: Returns Promise<Blob> with MIME type.
- `bytes()`: Returns Promise<Uint8Array>.
- `formData()`: Parses multipart or urlencoded data.
- `json()`: Parses JSON (rejects SyntaxError).
- `text()`: UTF-8 decodes body to string.

# Nuance Or Contradictions

- **Forbidden Headers**: The spec distinguishes between "forbidden request-header" and "forbidden response-header" based on the guard state, returning false rather than throwing an error for forbidden headers in mutable contexts.
- **CORS Privileged Headers**: In "request-no-cors" mode, privileged no-CORS request headers are automatically removed from the header list when modified by unprivileged code.
- **FormData Parsing**: The spec notes that a detailed parsing specification for `multipart/form-data` is to be written later, suggesting the current implementation is a rough approximation.
- **Stream Safety**: Reading a ReadableStream body throws a TypeError if the stream is disturbed, locked, or if `keepalive` is true.

# Candidate Wiki Hints

1. **Fetch Headers Guard States** (Concept)
2. **Headers Mutation Rules** (Concept)
3. **BodyInit Types and MIME Mapping** (Concept)
4. **Consuming Fetch Body Streams** (Procedure)
5. **FormData Parsing Logic** (Concept)
