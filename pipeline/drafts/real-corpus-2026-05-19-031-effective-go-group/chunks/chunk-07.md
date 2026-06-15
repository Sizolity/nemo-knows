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
