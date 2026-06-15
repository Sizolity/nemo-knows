## group-01

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

Group Context
- **Source Document**: Effective Go (from https://go.dev/doc/effective_go).
- **Scope**: Covers the complete style and language idiom guide, ranging from initial metadata retrieval to detailed discussions on formatting, control structures, memory allocation, maps, interfaces, concurrency models, and error handling.
- **Chunk Range**: Lines 1–2730 (Chunks 01–07).
- **Primary Heading Path**: Effective Go > Retrieved Text (with introductory sections in Chunk 01).

Cross-Chunk Summary
The document serves as a canonical guide for writing idiomatic Go code, emphasizing simplicity, reliability, and efficiency. It explicitly advises against direct translations of C++ or Java. The content flows logically from basic formatting conventions (`gofmt`) and naming rules to the mechanics of control flow (`if`, `for`, `switch`), memory management (`new` vs `make`, slices, maps), type system nuances (interfaces, embedding, type assertions), and finally concurrency patterns ("sharing memory by communicating") and robust error handling.

Repeated Or Central Claims
- **Idiomatic Style**: Go is designed for building software at scale; code should be simple and avoid complex idioms from other languages. Formatting is automatic via `gofmt`.
- **Naming Conventions**: Package names are short, lower-case, and single-word. Interface names follow the `<method>-er` pattern (e.g., `Reader`). Multiword identifiers use MixedCaps. Getters should be methods (e.g., `Owner()`) rather than prefixed fields (`GetOwner()`).
- **Control Structures**: Semicolons are implicit; braces must always delimit bodies. `switch` supports non-constant expressions and omitting fall-through logic. `for` loops unify C-style and range iterations.
- **Memory & Collections**: Use `make` for slices/maps/channels and `new` for pointers to zero values. Slices wrap arrays and pass references; use `_` to discard unwanted range values.
- **Interfaces**: Interfaces are sets of methods. A type implements an interface if it has the required methods (duck typing). The blank identifier `_` is used for compile-time checks (`var _ Type = value`) and discarding values/errors, though discarding errors in practice is discouraged.
- **Concurrency**: Avoid shared mutable state; use channels to synchronize goroutines ("sharing memory by communicating"). Channels provide synchronization primitives (blocking send/receive) that replace locks.

Important Local Details
- **Formatting**: `gofmt` aligns columns and uses tabs for indentation. Comments should be aligned vertically if they are part of a block.
- **Allocation**: `new(T)` returns `*T` with zeroed storage. `make(T, args)` initializes the underlying array/map/channel. Since Go 1.26, `new` accepts an initial value expression.
- **Maps**: Map lookups return `(value, ok)`. Use the "comma ok" idiom to check for existence. `delete(map, key)` removes entries safely even if absent. Keys must be comparable types; slices cannot be keys.
- **Printing**: `fmt.Printf` uses `%v`, `%+v` (struct fields), and `%#v` (Go syntax). Custom types should define a `String()` method to control output, avoiding recursion that prints the receiver as a string.
- **Methods & Receivers**: Value receivers work on both values and pointers; pointer receivers modify the original. Addressable values allow calling pointer methods without explicit `&`. Embedding promotes methods but retains the inner receiver type unless changed.
- **Errors & Panic**: Errors follow the convention `type error interface { Error() string }`. `panic` halts execution for unrecoverable states. Do not recover from panics in production code typically.
- **Concurrency Primitives**: Goroutines are lightweight threads. Channels can be buffered or unbuffered. Semaphore patterns use buffered channels to limit concurrency. Parallelism is achieved by breaking work into independent pieces executed across cores.

Candidate Wiki Hints
- **Style Guide**: `Effective Go` (Main Page)
- **Formatting**: `gofmt`, Indentation, Line Length, Comments
- **Naming**: Package Names, Getters/Setters, Interface Naming (`-er`), MixedCaps
- **Control Flow**: Semicolon Insertion, `if`/`switch`/`for` Logic, Blank Identifier Usage
- **Memory**: `new` vs `make`, Slices (Capacity, Length, Appending), Maps (Zero Value, Deletion)
- **Types & Interfaces**: Type Assertions, Embedding, Method Sets, Compile-Time Checks (`var _`)
- **Concurrency**: Goroutines, Channels (Buffered/Unbuffered), Semaphore Pattern, Parallelism vs Concurrency
- **Error Handling**: Error Interface Convention, `panic` Usage, Recovering from Errors

Gaps Or Cautions
- **Metadata Only Chunk**: Chunk 01 contains only retrieval metadata and navigation headers; it does not contain substantive code or style rules.
- **Implicit Semicolons**: Be cautious of implicit semicolon insertion rules (e.g., after `if`, `return`, `break`), especially when writing multi-statement lines or placing braces on new lines.
- **Error Discarding**: Using `_` to discard errors (e.g., in file operations) is a dangerous practice that can lead to runtime crashes; always check error conditions unless the context guarantees safety.
- **Interface Satisfaction**: While most interface satisfaction is checked at compile time, some packages (like `encoding/json`) may require run-time checks or specific type assertions due to how they define their interfaces.
- **Concurrency Bugs**: Historically, loop variables in goroutines (`range`) could cause data races; ensure variables are captured correctly (e.g., using an index variable) to avoid shared state issues.

## group-02

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Group Context

This group of notes covers the latter portion of the "Effective Go" document, specifically focusing on robust error handling strategies and a complete web server implementation. The content transitions from general advice on using `panic` and `recover` to concrete patterns for masking internal errors, and concludes with an example application that generates QR codes using HTML templates.

# Cross-Chunk Summary

The earlier chunks of the document (not included in this specific group's raw text but referenced by the source path) likely established the foundational concepts of Go idioms and error handling. This group synthesizes those concepts into actionable patterns:
1.  **Error Masking:** Moving away from letting unexpected panics crash the entire application.
2.  **Stack Unwinding Control:** Utilizing `defer` and `recover` to intercept panic values and convert them into standard error returns.
3.  **Safe Rendering:** Using `html/template` to prevent XSS vulnerabilities during data display.

The progression moves from abstract principles (library functions should not take down the whole program) to specific implementation details (the `Compile` function regex pattern), and finally to a full-stack example (QR code generator).

# Repeated Or Central Claims

- **Library Responsibility:** Library functions should prioritize masking problems or working around them rather than causing the entire program to halt.
- **The Panic/Recover Mechanism:** `panic` stops execution and begins unwinding the goroutine stack; `recover` is only effective within deferred functions during this unwinding phase.
- **Error Conversion:** Internal panics should be caught, converted into error values, and returned to callers to prevent unexpected stack unwinding in client code.
- **Safety of Templates:** The `html/template` package automatically escapes data, ensuring it is safe for display in HTML contexts without additional manual escaping.

# Important Local Details

### Panic and Recover Mechanics
- **Stack Unwinding:** When a panic occurs, the runtime unwinds the stack, executing any deferred functions before terminating the goroutine.
- **Recovery Scope:** `recover()` returns the panic value if called during unwinding; otherwise, it returns `nil`.
- **Deferred Functions:** Code to handle panics must be placed in a deferred function (or directly inside the function being panicked from) to be effective.

### Error Handling Patterns
1.  **Safe Initialization:** Check for required environment variables (e.g., `$USER`) during `init()`. Panic if missing, as this represents an unrecoverable setup failure.
    ```go
    var user = os.Getenv("USER")
    func init() {
      if user == "" {
        panic("no value for $USER")
      }
    }
    ```
2.  **Recover Pattern (Logging/Exit):** Wrap risky code in a function with a deferred `recover` block to log errors and exit cleanly without crashing the server.
    ```go
    func safelyDo(work *Work) {
      defer func() {
        if err := recover(); err != nil {
          log.Println("work failed:", err)
        }
      }()
      do(work)
    }
    ```
3.  **Parse Error Handling (Internal vs External):** Define a local `Error` type. In the `Compile` function, catch panics to convert them into returned errors. Crucially, re-panic if the recovered value is not an instance of the internal error type to ensure unexpected crashes are still reported.
    ```go
    type Error string
    func (e Error) Error() string { return string(e) }

    func Compile(str string) (*Regexp, error) {
      regexp := new(Regexp)
      defer func() {
        if e := recover(); e != nil {
          regexp = nil
          err = e.(Error) // Re-panics if not a parse error.
        }
      }()
      return regexp.doParse(str), nil
    }
    ```

### Web Server Implementation
- **Server Setup:** Use `http.ListenAndServe` to block until shutdown. Handle errors by logging and fatal-ing (e.g., in `main`).
- **Handler Logic:** Define a handler function that accepts text input and renders an HTML page.
- **QR Code Generation:** The example uses Google's chart API to generate QR codes based on the input string.
- **Template Execution:** Use `html/template`'s `Execute` method, passing the form value directly; automatic escaping handles security.

```go
func main() {
  flag.Parse()
  http.Handle("/", http.HandlerFunc(QR))
  err := http.ListenAndServe(*addr, nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}

func QR(w http.ResponseWriter, req *http.Request) {
  templ.Execute(w, req.FormValue("s"))
}
```

# Gaps Or Cautions

- **Unexpected Runtime Errors:** While `recover` can catch panics, it should not be used to handle unexpected runtime errors (like index out of bounds) if the goal is "fail fast." Such errors usually indicate a bug that should crash the program for debugging.
- **Type Assertions in Recovery:** The idiom of re-panicking inside a recovery block relies on type assertions (`e.(Error)`). If the assertion fails, it triggers a new panic. This preserves visibility into root causes in crash reports while masking expected internal errors.
- **Scope of `recover`:** Remember that `recover` outside of a deferred function or during stack unwinding is useless (it will always return `nil`).

