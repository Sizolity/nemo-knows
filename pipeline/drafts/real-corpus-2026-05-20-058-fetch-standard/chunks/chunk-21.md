---
title: Chunk 21 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk covers **Section 6: data: URLs** within the Fetch Standard. It defines how fetch algorithms handle requests where the URL is a `data:` URI, detailing the parsing of the MIME type and optional parameters from the fragment identifier, the creation of an empty body, and the interaction with CORS (Cross-Origin Resource Sharing) settings.

# Local Summary

The section establishes that when a request targets a `data:` URL, the algorithm must parse the URL to extract the media type and any associated metadata. It specifies that such requests have an empty body by default. The handling of `cross-origin` resources via `data:` URLs is restricted; specifically, a fetch from one origin for a `data:` URL from another origin (conceptually) is not permitted in the same way as network resources, often resulting in a CORS error if the resource is treated as cross-origin or if specific headers like `Cross-Origin-Opener-Policy` are violated. The text also notes that `data:` URLs cannot be used to load resources that require specific network-level behaviors (like redirects) and that they are generally processed synchronously within the context of the initiator unless explicitly handled otherwise.

# Key Claims

*   **Parsing**: The algorithm parses the `data:` URL to separate the media type from any optional parameters in the fragment.
*   **Body Content**: Requests for `data:` URLs result in an empty body unless specific logic dictates otherwise (though typically, `data:` URIs contain the data in the fragment). *Correction based on standard behavior implied*: The text implies the data is read from the URL's fragment as the body.
*   **CORS Restrictions**: Fetching a `data:` URL from a different origin than the initiator is generally not allowed or treated as an error depending on the specific CORS context (specifically regarding `Cross-Origin-Opener-Policy`).
*   **Redirection**: `data:` URLs do not support redirection; any attempt to redirect would fail.

# Entities And Concepts

*   **data: URL**: A Uniform Resource Identifier that contains data directly in the URI rather than pointing to a remote resource.
*   **MIME Type**: The media type identifier extracted from the `data:` URL fragment (e.g., `text/plain`, `image/png`).
*   **Cross-Origin-Opener-Policy**: A security policy mentioned in relation to `data:` URLs, restricting how they can be accessed across different browsing contexts.
*   **Fetch Algorithm**: The set of steps defined by the Fetch Standard for handling network requests and `data:` URLs.

# Procedures And API Details

*   **Parse Data URL**: The algorithm involves parsing the URL string to identify the media type and parameters.
*   **Create Request**: A new request object is created with an empty body initially, which is then populated if the data is read from the URL fragment.
*   **CORS Check**: Before processing the `data:` URL, the system checks for CORS violations, particularly regarding opener policies.

# Nuance Or Contradictions

The text implies a distinction between loading resources that require network behavior versus those that are self-contained. There is a potential nuance in how "cross-origin" applies to `data:` URLs compared to HTTP resources; since `data:` URLs are local to the document, the concept of "origin" for the resource itself differs from the initiator's origin in network contexts.

# Candidate Wiki Hints

*   **data: URL Structure**: How to parse and use `data:` URIs in web applications.
*   **CORS with data: URLs**: Security implications and restrictions when using `data:` URIs across origins.
*   **Fetch API and Local Resources**: Differences between fetching network resources and local `data:` resources.
