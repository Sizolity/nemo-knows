---
title: Chunk 21 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---
Chunk Context
- Heading path: 6. data: URLs
- Line range: 6581–7175
- Scope: Covers the "data:" URL scheme definition and usage within the Fetch Standard, including associated attributes, headers, and security flags referenced across multiple sections (§ 2.x, § 3.x, § 4.x, § 5.x).

Local Summary
This chunk details the `data:` URL scheme's integration into the Fetch specification. It defines how such URLs are handled as request bodies or response content, referencing specific attributes like `url` and `type` on `Request` and `Response` objects. The text also links to related concepts such as CORS preflight checks (`use-CORS-preflight flag`), unsafe request handling (`unsafe-request flag`), and credential usage (`use-URL-credentials flag`). It notes associations with other schemes (e.g., `video`, `text`, `xslt`) and specific headers like `X-Content-Type-Options`.

Key Claims
- The `data:` URL scheme is a defined resource type within the Fetch standard.
- Attributes such as `url` and `type` exist on both `Request` (§ 5.4) and `Response` (§ 5.5) objects to describe data sources.
- Security flags like `use-CORS-preflight`, `unsafe-request`, and `use-URL-credentials` are relevant contexts for handling these requests.
- The standard references content types including `"text"`, `"video"`, and `"xslt"` in the context of data URLs (§ 5.4).

Entities And Concepts
- **data: URL**: A URI scheme that carries data within the URL itself, often used for embedding resources like images or text directly in HTML/JS.
- **Request / Response Attributes**: Specific properties (`url`, `type`) attached to Fetch API objects to identify source content.
- **Security Flags**: Mechanisms like `use-CORS-preflight` and `unsafe-request` that dictate how data URLs interact with cross-origin policies or caching.
- **Content Types**: MIME types associated with embedded data, explicitly mentioning `"text"`, `"video"`, and `"xslt"`.

Procedures And API Details
- **Attribute Access**: The `url` attribute is defined for both `Request` (§ 5.4) and `Response` (§ 5.5).
- **Type Definition**: The `type` attribute exists on the `Response` object (§ 5.5).
- **Flag Usage**: Flags such as `use-CORS-preflight` (§ 2.2.5) and `unsafe-request` (§ 2.2.5) are referenced in the context of processing these URLs.

Nuance Or Contradictions
- The chunk lists various MIME types (`text`, `video`, `xslt`) alongside the generic `data:` scheme, implying that the standard treats data URLs flexibly regarding content type identification, though specific handling rules for each type may be defined elsewhere in the full document.

Candidate Wiki Hints
- **Page: Data URL Scheme** – A dedicated page explaining the syntax and security implications of `data:` URLs in web development.
- **Concept: Fetch API Attributes** – Documentation covering the `url` and `type` attributes on Request/Response objects.
