---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/034-go-faq.md
confidence: medium
---
Chunk Context
This chunk covers language semantics and runtime behavior in Go, including type system rules (method receivers, value vs. pointer passing), numeric type handling, constant precision, built-in data structure design choices (maps, slices), library documentation tools, module versioning, memory allocation semantics, and concurrency primitives.

Local Summary
The text explains why Go lacks implicit numeric conversions, how constants behave with arbitrary precision, the rationale behind maps being built-in but not accepting slices as keys, the evolution of map/slice/channel reference semantics versus array value semantics, conventions for libraries and code style (gofmt), module versioning rules, pointer vs. value passing in functions, method receiver choices, differences between `new` and `make`, platform-specific integer sizes, heap/stack allocation via escape analysis, virtual memory usage by the Go allocator, atomic operations, and the principle of sharing memory by communicating rather than by sharing memory directly.

Key Claims
- Go enforces exact type matching for methods; covariant result types are not supported via interfaces.
- Implicit numeric conversions are omitted to improve clarity and portability; constants live in an ideal number space until assigned to a variable.
- `int` is generic; its size depends on the platform (32-bit or 64-bit).
- Maps are built-in due to their power and importance; slices cannot be map keys because equality is not well-defined for them.
- Maps, slices, and channels act as references (descriptors pointing to shared data), while arrays are values.
- Libraries should document via `go doc` or `pkg.go.dev`; there is no explicit style guide but `Effective Go` and `gofmt` enforce conventions.
- Modules were introduced in Go 1.11; backward compatibility is required for same import paths, enforced via semantic versioning and major-version suffixes.
- All values are passed by value; pointers to interfaces almost never need to be used directly.
- Method receivers should be pointers if the method modifies the receiver; otherwise, value receivers are preferred for small types.
- `new` allocates memory; `make` initializes slices, maps, and channels.
- Escape analysis determines heap vs. stack allocation; address-taking may trigger heap allocation but not always.
- The Go allocator reserves virtual memory locally to the process.
- Concurrency enables parallelism only for intrinsically parallel problems; sequential tasks cannot benefit from multiple CPUs.

Entities And Concepts
- `Value` type and its `Copy()` method
- Numeric types: `int`, `uint`, `float64`, `float32`
- Constants: literal numbers, `math.Pi`
- Built-in collections: maps, slices, channels, arrays
- Data structure semantics: reference vs. value behavior
- Interfaces and method sets: `io.Writer`, `fmt.Fprintf`
- Module system: `go mod init`, `go get`, semantic versioning
- Memory management: `new`, `make`, escape analysis, virtual memory
- Concurrency primitives: goroutines, channels, `sync`, `sync/atomic`
- Documentation tools: `go doc`, `godoc`, `pkg.go.dev`, `pkgsite`
- Style guidance: `Effective Go`, `gofmt`, `Go Code Review Comments`

Procedures And API Details
- To check documentation from the command line, run `go doc <declaration>`.
- To initialize a module, execute `go mod init example/project`.
- To add or upgrade a dependency, use `go get golang.org/x/text@v0.3.5`.
- To configure HTTPS authentication for git, add entries to `$HOME/.netrc` with machine, login, and password fields.
- To switch from HTTPS to SSH URLs in git, modify `~/.gitconfig` with:
  ```text
  [url "ssh://git@github.com/"]
    insteadOf = https://github.com/
  ```
- To set private modules when using a public proxy, configure `GOPRIVATE`.
- For floating-point literals, `foo := 3.0` creates a `float64`; to create a `float32`, write `var foo float32 = 3.0` or `foo := float32(3.0)`.

Nuance Or Contradictions
- Although everything is passed by value, slices and maps behave like references because their values are descriptors pointing to shared data.
- Pointers to interfaces are rarely needed; passing a pointer to an interface value typically causes a compile-time error unless assigning to `interface{}`.
- The Go allocator reserves virtual memory but this does not affect other processes' memory availability.
- Integer sizes are implementation-specific, whereas floating-point types are fixed (`float64` by default for untyped constants).

Candidate Wiki Hints
- Page: Go Numeric Types and Constants
  - Covers implicit conversion policy, constant precision, `int` platform dependency, and float literal typing.
- Page: Go Built-in Collections Semantics
  - Explains why maps are built-in, slice key restrictions, and reference/value behavior of slices/maps/channels vs arrays.
- Page: Go Method Receivers Best Practices
  - Details when to use pointer vs. value receivers, modifications visibility, and efficiency considerations.
- Page: Go Module Versioning and Compatibility
  - Summarizes module introduction, backward compatibility rules, semantic versioning, and major-version suffixes.
- Page: Go Memory Allocation and Escape Analysis
  - Describes heap/stack allocation decisions, address-taking implications, and virtual memory usage.
- Page: Go Concurrency Principles
  - Introduces goroutines, channels, atomic operations, and the "share memory by communicating" proverb.
