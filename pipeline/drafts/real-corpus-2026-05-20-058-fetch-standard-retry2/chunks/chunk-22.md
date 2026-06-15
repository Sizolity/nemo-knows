---
title: Chunk 22 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
Heading path: 6. data: URLs
Line range: 7177-7884
This chunk defines a comprehensive glossary of terms used throughout the Fetch Standard, referencing external specifications such as HTML, DOM, ECMAScript, and various security policies. It also includes normative and non-normative references to RFCs and W3C standards, along with an IDL index for key interfaces like `Headers`, `Body`, and `Request`.

Local Summary
This section enumerates terms defined by multiple specifications (e.g., HTML, DOM, Fetch Metadata) that are relevant to the Fetch Standard. It also lists numerous normative references (RFCs, W3C specs) and non-normative resources. The chunk concludes with an IDL index defining interfaces like `Headers`, `Body`, and `Request`.

Key Claims
- The chunk does not make explicit claims but rather serves as a reference list for terms and specifications used in the Fetch Standard.
- Terms are defined by referencing external specifications (e.g., `[HTML]` defines "active document", `[DOM]` defines "Document").
- Normative references include RFCs (e.g., [RFC3986], [RFC9110]) and W3C standards (e.g., [HTML], [DOM]).
- Non-normative references cover security vulnerabilities (e.g., HTTP TRACE method) and older timing APIs (e.g., Navigation Timing).

Entities And Concepts
- **Terms Defined**: Includes concepts like "active document", "Document", "Blob", "File", "URL", "Request", "Response", "Headers", "Body", "AbortSignal", "Promise", "ReadableStream", "WritableStream", etc.
- **Specifications Referenced**: HTML, DOM, ECMAScript, Fetch Metadata, File API, High Resolution Time, HTTP, HTTP Caching, HTTP/1.1, HTTP/3, Infra, MIME Sniffing, Mixed Content, Permissions Policy, Referrer Policy, Reporting API, Resource Timing, Secure Contexts, Subresource Integrity, Streams, Service Workers, TLS, Upgrade Insecure Requests, URL, Web Crypto, WebDriver BiDi, Web IDL, WebSockets, WebTransport, XMLHttpRequest.
- **Interfaces**: `Headers`, `Body`, `Request`, `Response`, `ReadableStream`, `WritableStream`, `TransformStream`, `ServiceWorkerGlobalScope`.

Procedures And API Details
- The chunk does not describe procedures but lists terms and references.
- IDL definitions for `Headers` include methods like `append()`, `delete()`, `get()`, `has()`, `set()`.
- `Body` mixin includes attributes like `body`, `bodyUsed` and methods like `arrayBuffer()`, `blob()`, `bytes()`, `formData()`, `json()`, `text()`.
- `Request` interface constructor takes `RequestInfo` and optional `RequestInit`.

Nuance Or Contradictions
- The chunk is primarily a reference list and does not present contradictions or nuances in itself.
- Some terms are defined by reference to external specifications, which may evolve independently.

Candidate Wiki Hints
- **Fetch Standard Glossary**: A wiki page summarizing key terms defined across various specifications referenced in the Fetch Standard.
- **Normative References Overview**: A page listing and categorizing normative references (RFCs, W3C specs) used in web standards.
- **Non-Normative References**: A resource documenting non-normative references, including security advisories and legacy APIs.
- **IDL Index for Fetch API**: A technical reference detailing IDL definitions for `Headers`, `Body`, `Request`, and related interfaces.
