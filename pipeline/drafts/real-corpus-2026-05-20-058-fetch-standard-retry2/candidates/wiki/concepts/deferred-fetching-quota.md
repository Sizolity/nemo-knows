---
title: Deferred Fetching Quota
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Deferred Fetching Quota

The **Deferred Fetching Quota** is a memory management mechanism within the Fetch Standard that controls how much data can be buffered for HTTP request bodies when the source is a stream. It operates hierarchically to prevent excessive memory consumption while allowing for potential resends due to timeouts or network issues.

## Mechanism

When a fetch operation involves a stream, the system buffers incoming data up to **64 KiB** initially. This buffering supports features like incremental reading and cloning via `tee()`. The quota logic manages this buffer space by tracking limits across different levels:

-   **Top-level limit**: A global cap of approximately **640 KiB**.
-   **Concurrent per-origin limit**: A stricter cap of **64 KiB** for concurrent operations within the same origin.

If these limits are exceeded, the system may defer further fetching or trigger garbage collection to reclaim memory. This approach ensures that large payloads do not immediately block the thread or exhaust available heap space.

## Security Context

The quota mechanism interacts with broader security policies defined in the Fetch Standard:

-   It enforces strict **Security by Default** principles, ensuring that even legitimate large requests do not compromise system stability.
-   It complements **Connection Partitioning**, where connections are partitioned by a "Network Partition Key" (site origin + credentials) to prevent data leakage between sites.
-   It works alongside **Response Tainting** logic, which categorizes responses into `basic`, `cors`, and `opaque` states, influencing how headers and bodies are exposed or processed.

## Related Concepts

- [[fetch-standard-overview]]
- [[http-fetch-lifecycle]]
- [[request-response-objects]]
- [[fetch-body-streams]]
