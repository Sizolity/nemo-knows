---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/058-fetch-standard.md
confidence: medium
---

# Chunk Context
This chunk details the structural definition of a `body` object within the Fetch Standard, specifically covering sections 2.2.4 (Bodies) and associated algorithms for cloning, incrementally reading, fully reading, and handling content codings. It outlines the lifecycle of a body's stream, source, and length, and defines the microsteps required to process data chunks via parallel queues.

# Local Summary
A `body` is defined by three members: a `stream` (a ReadableStream), a `source` (initially null, representing raw data or a Blob/FormData), and a `length` (initially null). The chunk describes the procedure to clone a body by teeing its stream. It further defines two primary reading algorithms: "incrementally read," which processes chunks as they arrive using provided callback algorithms (`processBodyChunk`, `processEndOfBody`, `processBodyError`), and "fully read," which consumes the entire stream before invoking a final processing algorithm. Both reading mechanisms utilize parallel queues or global objects to manage asynchronous tasks. Finally, it introduces the concept of handling content codings by decoding bytes according to HTTP standards if supported.

# Key Claims
- A body is composed of a `stream`, a `source`, and a `length`.
- Cloning a body involves teeing the stream into two outputs (`out1` and `out2`) and assigning them to the cloned body's members.
- Incremental reading requires specific algorithms for handling chunks, end-of-body signals, and errors.
- The incrementally-read loop queues fetch tasks based on the state of the reader (chunk received, close requested, or error encountered).
- Fully reading a body involves reading all bytes from the stream and invoking a success callback with the full byte sequence or an error callback if reading fails.
- Content coding handling attempts to decode bytes using provided codings; if unsupported or decoding fails, the raw bytes are returned.

# Entities And Concepts
- **Body**: An object representing the payload of a request or response.
- **ReadableStream**: The underlying stream mechanism used by a body.
- **Source**: The origin of the data (null, byte sequence, Blob, or FormData).
- **Cloning**: Creating a new body instance sharing the same underlying stream via `tee()`.
- **Parallel Queue**: An execution context used to queue fetch tasks without blocking the main thread.
- **Content Codings**: Headers indicating compression or encoding applied to the payload (e.g., gzip, deflate).
- **Fetch Task**: A microtask queued for asynchronous execution.

# Procedures And API Details
**Cloning a Body**
1.  Let `« out1, out2 »` be the result of `tee()` on the body's stream.
2.  Set the body's stream to `out1`.
3.  Return a new body object where the stream is `out2` and other members are copied from the original.

**Incrementally Reading a Body**
Inputs: `body`, `processBodyChunk`, `processEndOfBody`, `processBodyError`, optional `taskDestination`.
Steps:
1.  Initialize `taskDestination` to a new parallel queue if null.
2.  Get a reader for the body's stream.
3.  Execute the "incrementally-read loop" with the reader and callbacks.

**The Incrementally-Read Loop Logic**
Inputs: `reader`, `taskDestination`, `processBodyChunk`, `processEndOfBody`, `processBodyError`.
Loop Actions:
-   **Chunk Steps**:
    -   If chunk is not a `Uint8Array`, run `processBodyError` with a `TypeError`.
    -   Otherwise, copy the chunk (implementation should minimize this copy).
    -   Run `processBodyChunk` with the bytes.
    -   Recursively perform the incrementally-read loop.
-   **Close Steps**: Queue a fetch task to run `processEndOfBody`.
-   **Error Steps**: Queue a fetch task to run `processBodyError` with the exception.

**Fully Reading a Body**
Inputs: `body`, `processBody`, `processBodyError`, optional `taskDestination`.
Steps:
1.  Initialize `taskDestination` if null.
2.  Define success steps: Queue a fetch task to run `processBody` with bytes.
3.  Define error steps: Queue a fetch task to run `processBodyError` with the exception.
4.  Get a reader for the body's stream; if getting the reader throws, run error steps and return.
5.  Read all bytes from the reader using success and error steps.

**Handling Content Codings**
Inputs: `codings`, `bytes`.
Steps:
1.  If `codings` are not supported, return `bytes`.
2.  Return the result of decoding `bytes` with `codings` per HTTP specifications, or failure if decoding errors occur.

# Nuance Or Contradictions
- **Stream Sharing**: Cloning a body does not duplicate the stream data but splits the single underlying stream into two distinct readable streams (`out1` and `out2`). Consuming one will affect the other.
- **Copy Minimization**: The spec explicitly encourages implementations to avoid copying the chunk bytes when possible, suggesting that the internal buffer might be used directly if safe.
- **Task Destination**: Both incremental and full reading default to creating a new parallel queue if none is provided, ensuring non-blocking execution of callbacks.

# Candidate Wiki Hints
- **Body Object Structure**: A page defining the members of a `body` (stream, source, length) and their initial states.
- **Body Cloning**: A guide on how to duplicate a request/response body while maintaining stream integrity using `tee()`.
- **Reading Algorithms**: A comparative overview of "incrementally reading" vs. "fully reading," detailing when to use callbacks for chunks versus the full payload.
- **Content Decoding**: An explanation of the content coding handling process, including fallback behavior when codings are unsupported.
