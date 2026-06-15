---
title: Chunk 16 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
The chunk details the **Request class** within the Fetch API specification. It defines the `Request` interface, its associated attributes (such as `url`, `method`, `headers`, and various modes like `mode`, `credentials`, and `cache`), the `RequestInit` dictionary for initialization options, enumeration types for attribute values, and the internal steps for constructing a new `Request` object via the `new Request(input, init)` constructor. It also covers getters for attributes, specific handling of referrers, and the logic for cloning a request.

Local Summary
This section explains how to create a `Request` object, detailing the properties available in the `init` argument (like `method`, `headers`, `body`, `referrer`, etc.). It outlines the internal constructor steps, including URL parsing, mode validation, referrer processing, header sanitization, and body handling. Getters for standard attributes are described, along with the specific logic for the `clone()` method which creates a new request object linked to the original signal.

Key Claims
- The `Request` interface represents a resource that is being fetched.
- The `new Request(input, init)` constructor accepts a URL string or an existing `Request` object as input.
- If `input` is a string, it must be parsed; if parsing fails or includes credentials, a `TypeError` is thrown.
- The default mode for a new request from a string is `"cors"`.
- Specific referrer values like `"about:client"` and `"no-referrer"` are handled internally during construction.
- The `clone()` method returns a clone of the request but requires the signal to be non-null.
- Headers added in the network layer (e.g., `Host`) are not included in the returned `Headers` object.

Entities And Concepts
- **Request**: The main interface for an HTTP request.
- **RequestInfo**: A union type representing either a `Request` object or a URL string (`USVString`).
- **RequestInit**: A dictionary containing optional initialization parameters for the `Request` constructor.
- **Headers**: An object used to manage HTTP headers associated with the request.
- **AbortSignal**: Used to cancel an ongoing fetch operation.
- **Referrer Policy**: Controls how referrer information is sent with requests.
- **CORS**: Cross-Origin Resource Sharing mode, distinct from `"same-origin"` or `"navigate"`.
- **SRI (Subresource Integrity)**: Cryptographic hash metadata for verifying resource integrity.

Procedures And API Details
- **Creating a Request**:
  - Pass a URL string to `new Request(url)`; default method is `"GET"`, default mode is `"cors"`.
  - Pass an existing `Request` object to copy its properties, optionally overriding them via the second argument.
  - Set `init.method` to change the HTTP verb (e.g., `"POST"`).
  - Set `init.headers` to provide a `Headers` object, literal, or array of arrays.
  - Set `init.body` to provide the request body content.
- **Referrer Handling**:
  - If `init.referrer` is an empty string, it sets the internal referrer to `"no-referrer"`.
  - If `init.referrer` is `"about:client"`, it sets the internal referrer to `"about:client"`.
  - If the parsed referrer's origin is not same-origin or scheme is `"about"` with path `"client"`, it defaults to `"client"`.
- **Cloning**:
  - Call `request.clone()` to get a copy.
  - The clone shares the abort signal dependency.
  - Throws a `TypeError` if the request is unusable.

Nuance Or Contradictions
- **Service Worker Origins**: A request can have an origin different from the current client when handled by a service worker, specifically for navigation requests.
- **Window Disassociation**: The `window` property in `RequestInit` can only be set to `null` to disassociate the request from any Window; setting it to a non-null value throws a `TypeError`.
- **Duplex Mode**: Currently, only `"half"` is valid for `duplex`. `"full"` is reserved for future use where response processing occurs before the full request body is sent.
- **No-CORS Restrictions**: If mode is `"no-cors"`, the method must be a CORS-safelisted method; otherwise, a `TypeError` is thrown during construction.

Candidate Wiki Hints
- [Request API](https://wiki.local/Request_API) - Core interface and attributes.
- [Request Constructor Options](https://wiki.local/Request_Constructor_Options) - Detailed breakdown of `RequestInit`.
- [Cloning Requests](https://wiki.local/Cloning_Requests) - Usage of the `clone()` method.
- [Referrer Policies in Fetch](https://wiki.local/Referrer_Policies_in_Fetch) - Handling of referrer strings and policies.
