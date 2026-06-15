---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/059-url-standard.md
confidence: medium
---
Chunk Context
The chunk covers the **URL** and **URLSearchParams** interfaces within the URL Standard. It details the API for constructing URLs, parsing strings (absolute vs. relative), accessing components (protocol, host, pathname, etc.), and manipulating query parameters. It also explains encoding differences between the `URL` object's serialization and `URLSearchParams`.

Local Summary
The **URL** interface allows creating objects from absolute or relative URLs. The constructor accepts a URL string and an optional base URL. Static methods `parse()` and `canParse()` handle parsing without throwing exceptions (returning null/false on failure). Accessors like `href`, `origin`, and setters for components (`protocol`, `username`, `host`, etc.) are defined, with specific behaviors regarding opaque paths and port handling. The **URLSearchParams** interface manages the query component of a URL using a list of name-value tuples, supporting methods like `append()`, `delete()`, `get()`, `getAll()`, `set()`, and `sort()`. Stringification follows the application/x-www-form-urlencoded format, which differs from standard URL percent-encoding (e.g., encoding spaces as `+`).

Key Claims
- A `URL` object has an associated `URL` record and a `query` object (a `URLSearchParams` instance).
- The `new URL(url)` constructor throws a `TypeError` if the input is a relative URL without a base.
- The `href` getter returns the serialization of the internal `URL` record; the setter re-parses the string and clears the query list before repopulating it.
- Setting the `host` property does not reset the port if the given value lacks one, which differs from the behavior implied by the `host` getter returning a port-inclusive string.
- `URLSearchParams` uses application/x-www-form-urlencoded encoding, where U+0020 SPACE is encoded as U+002B (+).
- The `sort()` method on `URLSearchParams` sorts tuples based on name code unit values to potentially increase cache hits.

Entities And Concepts
- **URL**: An object representing a URL with components like scheme, host, port, pathname, query, and hash.
- **URLSearchParams**: An object representing the query component of a URL as a list of name-value tuples.
- **Opaque path**: A condition where certain setters (e.g., `host`, `pathname`) are ignored or restricted.
- **application/x-www-form-urlencoded**: The encoding format used by `URLSearchParams`.
- **Percent-encode sets**: Distinct sets for general URLs vs. query strings affecting how characters like spaces or tildes are encoded.

Procedures And API Details
- **Constructing a URL**:
  - `new URL(url, base)`: Throws `TypeError` on failure.
  - `URL.parse(url, base)`: Returns `null` on failure.
  - `URL.canParse(url, base)`: Returns `false` on failure.
- **Setting components**:
  - `protocol`, `username`, `password`, `port`: May return early if the URL cannot have these (e.g., file URLs).
  - `host`: Ignores port if not provided in the setter value.
  - `pathname`: Clears the path before reparsing; ignores if opaque.
  - `search`: Removes leading `?` from input before parsing.
  - `hash`: Removes leading `#` from input before parsing.
- **URLSearchParams operations**:
  - `append(name, value)`: Adds a tuple to the list.
  - `set(name, value)`: Replaces existing tuples with that name; appends if none exist.
  - `delete(name, value)`: Removes specific or all tuples matching the name.
  - `sort()`: Sorts the internal list by name code units and updates the object.

Nuance Or Contradictions
- **Encoding Discrepancy**: The `URL` interface's serialization (used by `href`) uses a different percent-encode set than `URLSearchParams`. Specifically, `URLSearchParams` encodes U+0020 SPACE as U+002B (+), whereas standard URL encoding often uses `%20`. Additionally, `~` is encoded differently depending on the context (query vs. path).
- **Host Setter Surprise**: The `host` getter includes the port in its return value, but the `host` setter does not update the internal port if the provided string lacks a port number. This can lead to unexpected state where the visual host changes but the internal port remains static.

Candidate Wiki Hints
- URL vs. URLSearchParams encoding differences
- Handling relative URLs with the URL constructor
- Sorting URLSearchParams for cache optimization
- Understanding opaque paths in URL setters
