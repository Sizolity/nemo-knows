---
title: Chunk 22 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk defines the "6. data: URLs" section of a web specification, listing all external standards and specifications referenced to define terms used within that context. It includes normative references (W3C Living Standards, RFCs) and non-normative references (security advisories). The chunk concludes with an IDL Index defining the `Headers` and `Request` interfaces relevant to fetching data.

# Local Summary

The document defines terminology for handling "data:" URLs by referencing a wide array of Web Platform APIs and standards, including cookies, content security policies, fetch metadata, file APIs, HTTP caching, mixed content, permissions policy, reporting, resource timing, secure contexts, subresource integrity, streams, service workers, upgrade-insecure-requests, URL parsing, web crypto, web driver bidi, web IDL, websockets, and web transport. It also lists specific normative and non-normative references for these standards. Finally, it provides an IDL Index defining the `Headers` interface (with methods like `append`, `delete`, `get`, `set`) and the `Request` interface (with attributes like `method`, `url`, `headers`), along with typedefs for `BodyInit` and `XMLHttpRequestRequestBodyInit`.

# Key Claims

*   The terms defined in section 6 ("data: URLs") are inherited from or defined by specific external specifications.
*   A comprehensive list of normative references (e.g., [HTML], [FETCH-METADATA], [HTTP-CACHING]) is provided to define the terminology used.
*   A list of non-normative references (e.g., security advisories like [HTTPVERBSEC1]) is provided for context.
*   The `Headers` interface allows manipulation of request/response headers via methods like `append`, `delete`, `get`, `has`, and `set`.
*   The `Request` interface represents a request, possessing attributes such as `method`, `url`, and `headers`.
*   Typedefs exist for `BodyInit` (representing the body of a fetch request) and `XMLHttpRequestRequestBodyInit`.

# Entities And Concepts

*   **data: URLs**: The subject of section 6.
*   **Normative References**: Standards that define behavior (e.g., [HTML], [RFC9110]).
*   **Non-Normative References**: Informational documents or advisories (e.g., [HTTPVERBSEC1]).
*   **Headers Interface**: An interface for managing HTTP headers.
    *   `append(name, value)`: Adds a header.
    *   `delete(name)`: Removes a header.
    *   `get(name)`: Retrieves a header value.
    *   `has(name)`: Checks if a header exists.
    *   `set(name, value)`: Sets a header value.
*   **Request Interface**: Represents an HTTP request.
    *   `method`: The HTTP method (e.g., GET, POST).
    *   `url`: The request URL.
    *   `headers`: A `Headers` object containing the request headers.
*   **BodyInit**: A type representing possible values for a fetch request body (`ReadableStream`, `XMLHttpRequestBodyInit`).
*   **XMLHttpRequestRequestBodyInit**: A type for XHR request bodies (Blob, BufferSource, FormData, etc.).

# Procedures And API Details

*   **Headers Construction**: The `Headers` interface can be constructed using an optional `HeadersInit` argument.
*   **Header Manipulation**:
    *   `append(ByteString name, ByteString value)`: Appends a header field with the given name and value.
    *   `delete(ByteString name)`: Deletes all headers whose name matches the provided string (case-insensitive).
    *   `get(ByteString name)`: Returns the value of the first header whose name matches the provided string, or null if no match is found.
    *   `has(ByteString name)`: Returns true if there are any headers whose name matches the provided string.
    *   `set(ByteString name, ByteString value)`: Sets the header's value to the given value, replacing any existing headers with that name.
*   **Request Construction**: The `Request` interface is constructed using a `RequestInfo` input (either a `Request` object or a USVString URL) and an optional `RequestInit` dictionary.
*   **Body Access**: The `Body` mixin provides methods to read the request/response body in various formats:
    *   `arrayBuffer()`: Returns a Promise resolving to an `ArrayBuffer`.
    *   `blob()`: Returns a Promise resolving to a `Blob`.
    *   `bytes()`: Returns a Promise resolving to a `Uint8Array`.
    *   `formData()`: Returns a Promise resolving to a `FormData` object.
    *   `json()`: Returns a Promise resolving to the parsed JSON value.
    *   `text()`: Returns a Promise resolving to a USVString.

# Nuance Or Contradictions

*   The document distinguishes between "Normative References" (which define the standard) and "Non-Normative References" (which provide context or security advisories).
*   The `Headers` interface supports both appending new headers (`append`) and setting/replacing existing ones (`set`).
*   The `get` method returns the *first* matching header value, implying multiple headers with the same name are possible.

# Candidate Wiki Hints

*   **data: URLs**: A brief overview of data URL schemes and their usage in web contexts.
*   **Web Platform References**: A page linking to key W3C standards (HTML, Fetch, Cookies, etc.) for developers needing precise definitions.
*   **Request/Response Headers API**: Documentation for the `Headers` interface methods (`append`, `delete`, `get`, `has`, `set`).
*   **Fetch Request Body Types**: Explaining `BodyInit` and how to provide different types of request bodies (streams, blobs, forms).
*   **Security References**: A curated list of security-related RFCs and advisories for web developers.
