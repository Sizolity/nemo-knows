## chunk-01

---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

Chunk Context
- Document: Effective Go (from https://go.dev/doc/effective_go)
- Chunk Range: Lines 1–26
- Heading Coverage: Document, Effective Go, Fetch Metadata
- Metadata Status: Retrieved 2026-05-18; content type text/html; charset=utf-8.

Local Summary
- This initial chunk serves as metadata and navigation for the "Effective Go" document, a style and language idiom source. It confirms retrieval success and tags the material under "go" and "web-corpus".

Key Claims
- The document is titled "Effective Go - The Go Programming Language".
- It functions as a canonical guide for Go style and idioms.
- Fetch status is "ok", indicating successful retrieval of the source content.

Entities And Concepts
- Effective Go (canonical style guide)
- Go Programming Language
- Style and language idiom source
- Metadata retrieval context

Procedures And API Details
- No specific Go code, APIs, or procedures are present in this chunk; it contains only document metadata and fetch information.

Nuance Or Contradictions
- None observed within this metadata-only chunk. The content aligns with the expected structure of a curated documentation source.

Candidate Wiki Hints
- Create or reference: `Effective Go` (canonical style guide page)
- Tag pages with: `go`, `style-guide`, `idioms`

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---
Chunk Context
- **Source:** Effective Go (Introduction to Formatting, Names, Control Structures).
- **Heading path:** Effective Go > Retrieved Text.
- **Line Range:** 27–550.

Local Summary
This chunk introduces the philosophy behind writing idiomatic Go code: simplicity, reliability, and efficiency. It emphasizes that Go programs should not be direct translations of C++ or Java but should follow specific conventions regarding formatting, naming, and control flow. The text covers how `gofmt` handles formatting, rules for naming (package names, getters, interface names), semicolon insertion logic, and the syntax of control structures like `if`, `for`, and `switch`.

Key Claims
- Go is designed to make it easy to build software at scale; straightforward translations from other languages are unlikely to succeed.
- Formatting issues are handled automatically by `gofmt` (or `go fmt`), which aligns columns and manages indentation using tabs.
- Package names should be short, concise, evocative, lower-case, and single-word (e.g., `bufio.Reader`, not `BufReader`).
- Getters and setters are manual; method names like `Owner()` distinguish exported methods from unexported fields better than `GetOwner()`.
- Interface names for one-method interfaces typically follow the pattern `<method>-er` (e.g., `Reader`, `Writer`).
- Multiword names use MixedCaps (or mixedCaps), not underscores.
- Semicolons are implicitly inserted by the lexer after specific tokens before newlines, except before closing braces or in specific control structure contexts.
- `if` and `switch` statements support optional initialization; bodies must always be brace-delimited.
- The `for` loop unifies C-style `for`, `while`, and infinite loops; range clauses manage iteration over arrays, slices, strings, and maps.
- `switch` is more flexible than C's, supporting non-constant expressions and omitting fall-through logic (cases run top-to-bottom).

Entities And Concepts
- **Effective Go:** A guide for writing clear, idiomatic Go code (originally written for the 2009 release).
- **gofmt / go fmt:** Tool for standardizing source formatting.
- **Package Naming Convention:** Lower-case, single-word names (e.g., `encoding/base64`).
- **Getters/Exported Methods:** Use PascalCase methods (e.g., `Owner()`) to expose fields; avoid `Get` prefixes.
- **Interface Naming:** Suffix `-er` for agent nouns (e.g., `Reader`, `Writer`).
- **MixedCaps:** Standard for multiword identifiers.
- **Semicolon Insertion:** Automatic insertion rules based on token type before newlines.
- **Control Structures:** `if`, `for`, `switch`, `select`.
- **Blank Identifier (`_`):** Used to discard unwanted values in range loops.

Procedures And API Details
- **Formatting with gofmt:** Run `gofmt` (or `go fmt`) to align comments and reformat code; do not work around its output unless filing a bug.
- **Naming a Getter:** For an unexported field `owner`, define the exported method as `Owner()`.
  - Example usage:
    ```go
    owner := obj.Owner()
    if owner != user {
        obj.SetOwner(user)
    }
    ```
- **Switching on True:** A switch without an expression evaluates to `true` and checks cases sequentially.
  - Example:
    ```go
    func unhex(c byte) byte {
        switch {
        case '0' <= c && c <= '9':
            return c - '0'
        // ... other cases
        }
        return 0
    }
    ```
- **Handling Errors in if Chains:** When an `if` body ends with `return`, omit the unnecessary `else`.
  - Example:
    ```go
    f, err := os.Open(name)
    if err != nil {
        return err
    }
    // code using f
    ```
- **Short Declaration in Loops:** Use `:=` to declare loop variables directly.
  - Example:
    ```go
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    ```
- **Range over Maps/Slices:** Iterate with key/value pairs or discard the unwanted value using `_`.
  - Example:
    ```go
    for key, value := range oldMap {
        newMap[key] = value
    }
    // Discard index
    for _, value := range array {
        sum += value
    }
    ```
- **Switch with Multiple Cases:** List cases comma-separated.
  - Example:
    ```go
    func shouldEscape(c byte) bool {
        switch c {
        case ' ', '?', '&', '=', '#', '+', '%':
            return true
        }
        return false
    }
    ```

Nuance Or Contradictions
- **Semicolon Visibility:** Unlike C, semicolons are not visible in source code because the lexer inserts them automatically. However, they must be explicitly written to separate multiple statements on one line or in specific loop clauses.
- **Brace Placement:** The opening brace of a control structure (`if`, `for`, etc.) cannot be placed on the next line after a newline following a statement-ending token; this would cause an implicit semicolon insertion before the brace, leading to syntax errors.
- **Redeclaration in Scope:** The `:=` operator allows redeclaring a variable in the same scope if it already exists (e.g., reusing `err`), provided another new variable is also declared in that statement.

Candidate Wiki Hints
- Page: **Formatting** (Subsection: `gofmt`, Indentation, Line Length)
- Page: **Naming Conventions** (Subsections: Package Names, Getters, Interface Names, MixedCaps)
- Page: **Control Structures** (Subsections: Semicolon Insertion, `if` Statements, `for` Loops, `switch` Statements)

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

Chunk Context
This chunk covers control flow (switch, defer), function design (multiple returns, named results), memory allocation primitives (`new` vs `make`), and collection types (arrays, slices, maps).

Local Summary
The text details idiomatic Go patterns for error handling via multiple return values, resource management using `defer`, and the distinction between `new` (pointer to zero value) and `make` (initialized slice/map/channel). It also explains the mechanics of slices, including appending and two-dimensional structures.

Key Claims
- Functions in Go can return multiple values; this replaces C-style error codes and reference parameters.
- Named result parameters act as documentation and simplify logic when using bare `return`.
- `defer` schedules a call to run before the function returns, evaluated at execution time for arguments (LIFO order).
- `new(T)` allocates zeroed storage and returns a pointer (`*T`). Since Go 1.26, it accepts an initial value expression.
- `make(T, args)` initializes slices, maps, and channels; it does not return a pointer.
- Arrays are values (copied on assignment); slices wrap arrays and pass references to underlying data.
- Slices can be sliced (`buf[0:32]`) without copying data; appending grows capacity or reallocates if necessary.

Entities And Concepts
- `continue` with labels (loop control)
- Type switch syntax (`switch t := t.(type)`)
- Multiple return values `(n int, err error)`
- Named result parameters `(value, nextPos int)`
- `defer` execution timing and LIFO order
- `new(T)` vs `make(T)` allocation semantics
- Zero value initialization strategy
- Composite literals for struct/map/slice/arrays
- Arrays (values) vs Slices (references)
- Slice capacity (`cap`) and length (`len`)
- Two-dimensional slices (array-of-slices or slice-of-slices)

Procedures And API Details
- **Compare byte slices**: Uses two switch statements to compare elements then lengths.
  ```go
  func Compare(a, b []byte) int {
      for i := 0; i < len(a) && i < len(b); i++ {
          switch {
          case a[i] > b[i]:
              return 1
          case a[i] < b[i]:
              return -1
          }
      }
      switch {
      case len(a) > len(b):
          return 1
      case len(a) < len(b):
          return -1
      }
      return 0
  }
  ```
- **Reading numbers**: `nextInt` returns value and next position; used in a loop to scan input.
  ```go
  func nextInt(b []byte, i int) (int, int) { ... }
  for i := 0; i < len(b); {
      x, i = nextInt(b, i)
      fmt.Println(x)
  }
  ```
- **File reading with defer**: `Contents` opens a file and defers its close.
  ```go
  func Contents(filename string) (string, error) {
      f, err := os.Open(filename)
      if err != nil { return "", err }
      defer f.Close()
      // read loop
      return string(result), nil
  }
  ```
- **Appending to a slice**: Manual implementation showing capacity check and reallocation.
  ```go
  func Append(slice, data []byte) []byte {
      l := len(slice)
      if l + len(data) > cap(slice) {
          newSlice := make([]byte, (l+len(data))*2)
          copy(newSlice, slice)
          slice = newSlice
      }
      slice = slice[0:l+len(data)]
      copy(slice[l:], data)
      return slice
  }
  ```

Nuance Or Contradictions
- `new` returns a pointer to zeroed memory; since Go 1.26, it can initialize with an expression (e.g., `new(int64(300))`).
- Returning the address of a composite literal is safe because the storage survives function return.
- `make([]int)` initializes the slice structure and underlying array; `new([]int)` returns a pointer to a nil slice structure.

Candidate Wiki Hints
- Multiple Return Values in Go
- Named Result Parameters
- Defer Mechanics and Evaluation Order
- new vs make: Allocation Primitives
- Slices: Underlying Arrays, Capacity, and Appending
- Two-Dimensional Slices Strategies

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---
## Chunk Context
Lines 1044–1546 of `effective-go.md`, covering "Maps", "Formatted printing in Go", "Append", "Initialization", "Constants", "Variables", "The init function", "Methods (Pointers vs. Values)", and "Interfaces".

## Local Summary
This chunk explains built-in maps, formatted output via `fmt`, slice modification with `append`, compile-time constants and the `iota` enumerator, variable initialization, package-level `init` functions, defining methods on various types including pointers vs. values, interface implementation (e.g., `io.Writer`, `sort.Interface`), and type conversions to safely call standard formatting or sorting routines.

## Key Claims
- Maps associate keys of any equality-definable type with values; slices cannot be keys.
- Map lookup returns the zero value for missing keys unless using the "comma ok" idiom.
- The blank identifier `_` discards a map's value in lookups, retaining only presence.
- `delete(map, key)` safely removes an entry even if absent.
- `fmt.Printf`, `Fprintf`, and `Sprintf` use `%v` for default formatting; `%+v` annotates struct fields; `%#v` prints full Go syntax.
- Custom types define a `String() string` method to control `%v` output; avoid recursive calls that print the receiver as a string.
- `append(slice, elements...)` appends values or expanded slices; use `...` at call sites to pass slices element-wise.
- Constants must be compile-time evaluatable expressions; use `iota` for enumerated constants.
- `init()` functions run after all variable initializers and imported packages, useful for state verification/repair.
- Value receivers work on both values and pointers; pointer receivers modify the original; addressable values allow calling pointer methods without explicit `&`.
- Interfaces specify behavior via method sets; a type can implement multiple interfaces (e.g., `sort.Interface`, custom printer).
- Converting a named slice type to its underlying slice type (`[]int(s)`) enables reuse of standard functions like `fmt.Sprint` or `sort.Sort`.

## Entities And Concepts
- `map[K]V`: built-in map type.
- `"comma ok" idiom`: `seconds, ok := timeZone[tz]`.
- `_`: blank identifier for ignoring values.
- `delete(map, key)`: removes a map entry.
- `fmt.Printf`, `Fprintf`, `Sprintf`, `Print`, `Println`, `Fprint`: formatted printing functions.
- `%v`, `%+v`, `%#v`, `%x`, `%q`, `%T`: format verbs.
- `append(slice, elements...)`: slice modification.
- `iota`: enumerator for constants.
- `init()`: package-level initialization function.
- Value vs. pointer receivers: rules and automatic address insertion for addressable values.
- Interfaces: e.g., `io.Writer`, `sort.Interface`.
- Type conversions: `[]int(s)` to access standard methods on slice types.

## Procedures And API Details
```go
// Map lookup with ok flag
seconds, ok := timeZone[tz]

// Delete a map entry
delete(timeZone, "PDT")

// Default formatting
fmt.Println(timeZone)           // equivalent to fmt.Printf("%v\n", timeZone)
fmt.Printf("%+v\n", t)          // struct fields annotated
fmt.Printf("%#v\n", t)          // full Go syntax

// Custom String method (safe from recursion)
func (b ByteSize) String() string {
    switch {
    case b >= YB: return fmt.Sprintf("%.2fYB", b/YB)
    // ... other cases
    }
    return fmt.Sprintf("%.2fB", b)
}

// Append values or a slice
x = append(x, 4, 5, 6)          // append multiple values
x = append(x, y...)             // append another slice

// Init function example
func init() {
    if user == "" { log.Fatal("$USER not set") }
    if home == "" { home = "/home/" + user }
    // ...
}

// Method on pointer receiver satisfying io.Writer
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // append logic here
    *p = slice
    return len(data), nil
}

// Interface implementation example
type Sequence []int
func (s Sequence) Len() int { return len(s) }
func (s Sequence) Less(i, j int) bool { return s[i] < s[j] }
func (s Sequence) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Conversion to underlying slice type
func (s Sequence) String() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
```

## Nuance Or Contradictions
- `fmt.Printf` does not accept signedness/size flags; it infers them from argument types.
- Value receivers work on both values and pointers; pointer receivers only on pointers, but addressable values allow calling pointer methods via automatic `&`.
- Recursive `String()` implementations that print the receiver as a string cause infinite recursion; convert to base `string` or use non-string format verbs like `%f`.
- Maps sort output lexicographically by key when printed with `%v`.

## Candidate Wiki Hints
- **Maps**: Key types, zero-value semantics, "comma ok" idiom, deletion.
- **Formatted Printing**: Format verbs, custom `String()` methods, recursion pitfalls.
- **Append & Slices**: Variadic usage, returning updated slices.
- **Initialization**: Constants, `iota`, variable initializers, `init()` lifecycle.
- **Methods**: Value vs. pointer receivers, automatic address insertion, satisfying interfaces.
- **Interfaces**: Method sets, implementing standard interfaces (`io.Writer`, `sort.Interface`).
- **Type Conversions**: Accessing underlying slice types for standard functions.

## chunk-05

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

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

Chunk Context
This chunk covers lines 2019–2488 of the source document, focusing on embedding types, concurrency models (goroutines and channels), parallelization strategies, resource management via leaky buffers, and error handling conventions including `panic`.

Local Summary
The text explains how Go handles type embedding, where embedded methods become accessible but retain their original receiver. It contrasts this with subclassing in other languages. The section transitions into concurrency, advocating for "sharing memory by communicating" using goroutines and channels rather than shared mutable state. It details goroutine creation, channel buffering, semaphore patterns, and the distinction between concurrency and parallelism. Finally, it covers error handling via `error` interfaces and `panic`.

Key Claims
- Embedding a type makes its methods available on the outer type, but the receiver remains the inner type unless explicitly changed.
- Go encourages avoiding shared mutable state; instead, use channels to synchronize goroutines ("share memory by communicating").
- Goroutines are lightweight functions that run concurrently within the same address space, multiplexed onto OS threads.
- Unbuffered channels provide synchronization because both sender and receiver block until data is exchanged.
- Channels of channels enable parallel demultiplexing and can be used to implement RPC-like systems without locks.
- Parallelism in Go is achieved by breaking work into independent pieces executed across multiple CPU cores, signaled via channels.
- Error values should follow the convention `type error interface { Error() string }`, with library routines returning detailed errors alongside results.
- `panic` stops program execution and is used for unrecoverable errors or impossible states.

Entities And Concepts
- Embedding: A Go idiom where a field of type T grants access to T's methods; receiver remains the embedded instance.
- Goroutine: Lightweight concurrent execution unit, created with the `go` keyword.
- Channel: Communication primitive for synchronizing goroutines; can be buffered or unbuffered.
- Semaphore pattern: Using a buffered channel to limit concurrency (e.g., limiting outstanding requests).
- Leaky buffer: A free-list managed via channels and garbage collection for memory reuse.
- Error interface: Built-in Go error handling mechanism requiring an `Error() string` method.
- Panic: Runtime function that halts execution, used for fatal errors.
- Parallelism vs Concurrency: Distinction between executing multiple computations simultaneously (parallelism) and structuring programs with independent components (concurrency).

Procedures And API Details
- Creating a goroutine: `go func() { ... }()`
- Sending on channel: `c <- value`
- Receiving from channel: `<-c`
- Buffered channel creation: `make(chan Type, capacity)`
- Unbuffered channel creation: `make(chan Type)`
- Checking CPU count: `runtime.NumCPU()` or `runtime.GOMAXPROCS(0)`
- Type assertion for errors: `if e, ok := err.(*os.PathError); ok { ... }`
- Calling panic: `panic(fmt.Sprintf("message"))`

Nuance Or Contradictions
- Embedding does not create a new type; it promotes methods but preserves the original receiver identity.
- While channels avoid data races by design, they are not always sufficient for all parallelization needs (Go is concurrent, not inherently parallel).
- A bug existed in Go versions before 1.22 where loop variables shared across goroutines could cause issues when using `range`.
- Leaky buffers rely on garbage collection to reclaim dropped buffers if the free list fills up.

Candidate Wiki Hints
- Type Embedding in Go
- Concurrency Model: Channels and Goroutines
- Semaphore Pattern with Buffered Channels
- Parallel Demultiplexing via Channels
- Error Handling Best Practices
- Panic Usage Guidelines

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---
Chunk Context
This chunk concludes the "Effective Go" section, focusing on error handling strategies using `panic` and `recover`, and demonstrates a complete web server implementation using `html/template`.

Local Summary
The text advises that library functions should avoid `panic` unless during initialization where setup failure is unrecoverable. It details how `recover` stops stack unwinding but only works within deferred functions. A pattern for handling parse errors in a regexp package is shown, converting internal panics into error values. Finally, a complete Go web server example is provided that generates QR codes via Google's chart API using HTML templates.

Key Claims
- Library functions should mask problems or work around them rather than taking down the whole program.
- `panic` unwinds the stack and runs deferred functions; `recover` stops this process and returns the panic argument.
- `recover` is only useful inside deferred functions because it must be called during stack unwinding.
- It is safe to call library routines that use `panic`/`recover` from within a deferred function handling a panic.
- Internal panics should be converted to error values before being exposed to clients; this prevents unwinding the caller's stack unexpectedly.
- The `html/template` package automatically escapes data, making it safe to display in HTML contexts.

Entities And Concepts
- `panic`: Stops execution and begins unwinding the goroutine stack.
- `recover`: Regains control of a goroutine during stack unwinding; returns the panic argument or nil.
- `defer`: Allows code to run before returning, essential for capturing `recover` calls.
- `os.Getenv`: Retrieves environment variables, used in the initialization example.
- `http.ListenAndServe`: Starts an HTTP server; blocks until shutdown.
- `html/template`: Package for executing HTML templates with data substitution and automatic escaping.
- QR Code: A matrix of boxes encoding text, generated via Google's chart API in the example.

Procedures And API Details
1. **Safe Initialization**: Use `os.Getenv` to check for required environment variables; panic if missing during `init()`.
   ```go
   var user = os.Getenv("USER")
   func init() {
     if user == "" {
       panic("no value for $USER")
     }
   }
   ```
2. **Recover Pattern**: Wrap risky code in a function with a deferred `recover` block to log errors and exit cleanly.
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
3. **Parse Error Handling**: Define a local `Error` type and use a deferred block to convert panics into returned errors.
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
4. **Web Server Example**: Create a server that accepts text input and renders an HTML page with a generated QR code image.
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

Nuance Or Contradictions
- While `recover` can catch panics, it should not be used to handle unexpected runtime errors (like index out of bounds) if the goal is to fail fast; the type assertion in the recovery block ensures only expected parse errors are caught.
- The re-panic idiom changes the panic value, but both original and new failures appear in crash reports, preserving root cause visibility.

Candidate Wiki Hints
- Topic: Panic and Recover Mechanics
- Topic: Error Handling Patterns in Go
- Topic: HTML Template Escaping

