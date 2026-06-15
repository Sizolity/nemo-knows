---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

Chunk Context
This chunk details the structure of a `Body` object, defining its components (stream, source, length) and cloning logic. It further elaborates on algorithms for incrementally reading a body, including handling chunks, closing, and errors via parallel queues. Finally, it covers fully reading a body and handling content codings.

Local Summary
A Body is composed of a ReadableStream, an optional source (null, byte sequence, Blob, or FormData), and an optional length. Cloning involves teeing the stream and duplicating properties. Incremental reading utilizes specific algorithms (`processBodyChunk`, `processEndOfBody`, `processBodyError`) queued as fetch tasks to handle incoming bytes, completion, or errors. Full reading aggregates these steps to consume the entire stream. Content codings are decoded if supported; otherwise, raw bytes are returned.

Key Claims
- A body consists of a stream, source (initially null), and length (initially null).
- Cloning a body tees its stream into two outputs (`out1`, `out2`); the original uses `out1`, while the clone uses `out2`.
- Incremental reading requires algorithms for processing chunks, end-of-body, and errors.
- If a chunk is not a Uint8Array, a TypeError is queued via `processBodyError`.
- Implementations are strongly encouraged to avoid copying byte sequences where possible during incremental reads.
- Fully reading a body queues tasks for success (byte sequence) and error (exception) handling.
- Content codings are decoded if supported; failure results in returning the raw bytes.

Entities And Concepts
- `Body`: An object representing the response body.
- `ReadableStream`: The underlying stream of data within a Body.
- `source`: The initial source of the body (null, byte sequence, Blob, FormData).
- `length`: The size of the body content.
- `processBodyChunk`: Algorithm to handle incoming byte sequences.
- `processEndOfBody`: Algorithm to handle the end of the stream.
- `processBodyError`: Algorithm to handle exceptions during reading.
- `taskDestination`: A parallel queue or global object for scheduling tasks.
- `Uint8Array`: The expected type for body chunks; deviation triggers a TypeError.

Procedures And API Details
- **Cloning a Body**:
  1. Tee the body's stream into `out1` and `out2`.
  2. Set the original body's stream to `out1`.
  3. Return a new body with stream `out2` and copied members.
- **Incrementally Reading a Body**:
  - Requires algorithms: `processBodyChunk(bytes)`, `processEndOfBody()`, `processBodyError(exception)`.
  - Initializes `taskDestination` to a new parallel queue if null.
  - Gets a reader for the body's stream.
  - Loops through chunks, queuing tasks for processing, closing, or errors.
- **Fully Reading a Body**:
  - Requires algorithms: `processBody(bytes)`, `processBodyError(exception)`.
  - Initializes `taskDestination` to a new parallel queue if null.
  - Reads all bytes from the reader, invoking success or error steps accordingly.

Nuance Or Contradictions
- The spec notes that getting a reader for a body's stream will not throw an exception, implying robustness in stream access.
- While copying byte sequences is part of the standard algorithm, implementations are strongly encouraged to avoid this copy to optimize performance.
- Content coding decoding fails if an error occurs during decoding, resulting in failure; otherwise, it returns the decoded bytes. If codings are unsupported, raw bytes are returned immediately.

Candidate Wiki Hints
- [Fetch API: Body Structure](https://wiki.example.com/fetch-api/body-structure)
- [Fetch API: Incremental Reading Algorithms](https://wiki.example.com/fetch-api/incremental-reading)
- [Fetch API: Handling Content Codings](https://wiki.example.com/fetch-api/content-codings)
