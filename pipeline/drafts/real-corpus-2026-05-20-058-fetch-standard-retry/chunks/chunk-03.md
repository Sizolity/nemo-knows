---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context

This chunk details the internal structure of a `Body` object within the Fetch specification. It defines the composition of a body (stream, source, length), outlines the algorithm for cloning bodies, and describes procedures for incrementally reading and fully reading body content using readers and task queues. It also briefly touches on the concept of a `BodyWithType` tuple and handling content codings.

# Local Summary

The specification defines a `Body` as having three members: a `stream` (a ReadableStream), a `source` (initially null, potentially a byte sequence, Blob, or FormData), and a `length` (initially null). The chunk provides algorithms for cloning a body by teeing its stream. It further details the "incrementally-read loop" which processes chunks via callbacks (`processBodyChunk`, `processEndOfBody`, `processBodyError`) using a parallel queue. Finally, it describes the steps to fully read a body into a single byte sequence and introduces the concept of a `BodyWithType` tuple for content negotiation.

# Key Claims

- A `Body` object is structurally composed of a stream, an optional source, and an optional length.
- Cloning a body involves creating two streams from the original via `tee()`, assigning one to the original body and returning the other as the new body's stream, while copying other members.
- Incremental reading relies on a "read loop" that queues fetch tasks for chunk processing, closing, or error handling, utilizing a parallel queue or global object task destination.
- The incrementally-read loop expects chunks to be `Uint8Array` objects; non-Uint8Array chunks trigger an error callback.
- Fully reading a body involves reading all bytes from the stream and passing them to a provided processing algorithm via queued tasks.

# Entities And Concepts

- **Body**: An object representing the response or request body, consisting of a stream, source, and length.
- **ReadableStream**: The underlying stream interface used by bodies to provide data flow.
- **Source**: The initial content of a body (null, byte sequence, Blob, or FormData).
- **Incrementally-read loop**: An algorithmic loop that processes incoming chunks incrementally using callbacks.
- **Task Destination**: A parallel queue or global object where fetch tasks are queued to handle body processing logic.
- **BodyWithType**: A tuple consisting of a `Body` and a type header value (or null).
- **Content Codings**: Mechanisms for handling encoding transformations on the bytes received from a body.

# Procedures And API Details

**Cloning a Body:**
1.  Use `tee()` on the body's stream to get two streams (`out1`, `out2`).
2.  Set the original body's stream to `out1`.
3.  Return a new body object where the stream is `out2` and other members are copied from the original.

**Incrementally Reading a Body:**
-   **Inputs**: `processBodyChunk` (algorithm accepting byte sequence), `processEndOfBody` (algorithm accepting no arguments), `processBodyError` (algorithm accepting exception), optional `taskDestination`.
-   **Steps**:
    1.  Initialize `taskDestination` to a new parallel queue if null.
    2.  Get a reader for the body's stream.
    3.  Execute the incrementally-read loop with the reader, task destination, and processing algorithms.
-   **Loop Logic (Given `reader`, `taskDestination`, etc.)**:
    1.  Create a read request with "chunk steps", "close steps", and "error steps".
    2.  **Chunk Steps**:
        -   If chunk is not a `Uint8Array`, queue `processBodyError` with a `TypeError`.
        -   Otherwise, copy the chunk (implementation should avoid copying if possible), run `processBodyChunk`, then recursively perform the incrementally-read loop.
    3.  **Close Steps**: Queue `processEndOfBody`.
    4.  **Error Steps**: Queue `processBodyError` with the exception `e`.

**Fully Reading a Body:**
-   **Inputs**: `processBody`, `processBodyError`, optional `taskDestination`.
-   **Steps**:
    1.  Initialize `taskDestination` if null.
    2.  Define success steps: Queue a task to run `processBody` with the byte sequence.
    3.  Define error steps: Queue a task to run `processBodyError` with the exception.
    4.  Get a reader for the body's stream; if this throws, run error steps and return.
    5.  Read all bytes from the reader using the defined success and error steps.

# Nuance Or Contradictions

-   **Chunk Copying**: The spec explicitly states that implementations are "strongly encouraged to use an implementation strategy that avoids" copying the chunk (`bytes`) when possible, even though the algorithm step says "Let bytes be a copy of chunk." This highlights a tension between the defined abstract steps and performance optimization recommendations.
-   **Task Destination**: The default for `taskDestination` is `null`, which triggers the creation of a new parallel queue. This implies that if no destination is provided, the implementation must manage concurrency internally via a newly created queue.

# Candidate Wiki Hints

-   **Page: Fetch Body Structure** - Documenting the three members of a `Body` object (stream, source, length) and their initial states.
-   **Page: Cloning Fetch Bodies** - Explaining the mechanics of `tee()` and how to create duplicate body streams for parallel processing or cloning requests/responses.
-   **Page: Incremental Body Reading** - A guide on implementing the incrementally-read loop, handling `Uint8Array` chunks, and managing callbacks for data processing, completion, and errors.
-   **Page: Fetch Task Queues** - An overview of how fetch tasks are queued using parallel queues or global objects to handle body processing asynchronously.
