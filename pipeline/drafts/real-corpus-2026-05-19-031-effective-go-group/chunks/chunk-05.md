---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---
Chunk Context
Effective Go > Retrieved Text (Lines 1548-2017)

Local Summary
This chunk explores advanced interface usage in Go, covering type assertions, the blank identifier `_` for compile-time checks and unused imports, interface embedding, and function receivers. It emphasizes that interfaces are sets of methods, allowing almost any type to satisfy them if it implements the required methods.

Key Claims
- Interfaces can be converted via type switches or type assertions (`value.(typeName)`). Type assertions may panic; use the "comma, ok" idiom for safety.
- Types need not export themselves if they only implement an interface; returning the interface from constructors promotes flexibility (e.g., `hash.Hash32`).
- The blank identifier `_` discards values, silences compiler warnings about unused imports/variables, and enables compile-time interface checks via global declarations like `var _ json.Marshaler = (*RawMessage)(nil)`.
- Interfaces can embed other interfaces (`type ReadWriter interface { Reader; Writer }`) or structs can embed pointers to structs to inherit methods without forwarding code.
- Functions can have methods attached (e.g., `HandlerFunc`), allowing ordinary functions to serve as HTTP handlers.

Entities And Concepts
- Type Switch: Converts an interface value by checking its concrete type or implemented interfaces.
- Type Assertion: Extracts a value of a specific type from an interface (`value.(string)`).
- Blank Identifier (`_`): Represents a discardable value; used for ignoring errors, unused imports, and compile-time interface verification.
- Interface Embedding: Combining multiple interfaces or structs within a new type to inherit methods automatically.
- HandlerFunc: An adapter type wrapping a function to satisfy the `http.Handler` interface.

Procedures And API Details
- Safe Type Assertion: `str, ok := value.(string)` checks if the value is of the specified type without panicking.
- Compile-Time Interface Check: `var _ json.Marshaler = (*RawMessage)(nil)` ensures a type satisfies an interface at compile time.
- Import for Side Effects: `import _ "net/http/pprof"` imports a package solely to trigger its initialization side effects.
- Creating ReadWriter via Embedding: `type ReadWriter struct { *Reader; *Writer }` combines reader and writer functionality without manual method forwarding.

Nuance Or Contradictions
- Discarding errors with `_` (e.g., `fi, _ := os.Stat(path)`) is flagged as terrible practice and can lead to crashes if the error indicates a missing file.
- Interface checks are typically static (compile-time), but some packages like `encoding/json` require run-time type assertions because the interface satisfaction isn't enforced statically.

Candidate Wiki Hints
- Go Type Assertions and Safety
- The Blank Identifier in Go
- Compile-Time Interface Checks with `_`
- Interface Embedding vs Explicit Method Forwarding
- Function Receivers and Adapter Patterns
