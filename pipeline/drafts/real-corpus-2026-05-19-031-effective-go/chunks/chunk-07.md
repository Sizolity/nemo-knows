---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## Chunk Context
This chunk concludes the "Effective Go" chapter, focusing on advanced error handling patterns involving `panic` and `recover`, followed by a complete example of a QR code web server using the `html/template` package. The text ends with footer navigation links from go.dev.

## Local Summary
The section first advises that library functions should avoid panicking unless during initialization, preferring to mask or work around problems to keep programs running. It then details how to use `recover` inside deferred functions to catch panics and allow a goroutine to exit cleanly without crashing the whole server. A specific pattern is shown where `panic` reports parse errors, which are caught by a recovery block that converts them into return values (`nil` and an error). Finally, a complete Go program is presented that acts as a web re-server for generating QR codes from text inputs using a template engine.

## Key Claims
- Library functions should avoid panicking; if a problem can be masked or worked around, it is better to let the program continue running rather than taking it down.
- `recover` is only useful inside deferred functions because it stops stack unwinding and returns the argument passed to `panic`.
- A recovery pattern allows a failing goroutine to log its failure and exit cleanly without disturbing other executing goroutines.
- Library routines that use `panic` and `recover` can be called from deferred code without failing, provided `recover` is not invoked directly outside of a defer block.
- Panics should generally be confined within a package; internal panics can be converted into error values for external consumers (e.g., the `Compile` function).
- The `html/template` package allows dynamic rewriting of HTML text by substituting data items, with automatic escaping to ensure safety.

## Entities And Concepts
- **panic**: Immediately stops execution of the current function and begins unwinding the stack; used for run-time errors like index out of bounds or failed type assertions.
- **recover**: Built-in function that regains control of a goroutine when called from a deferred function, returning the value passed to `panic` or `nil`.
- **deferred functions**: Only code running during stack unwinding; the only place where `recover` is effective.
- **Error interface**: Satisfied by custom types like `type Error string` to report parse errors via panic recovery.
- **html/template**: Package for executing templates with data items, providing automatic escaping and conditional rendering (e.g., `{{if .}}`).
- **QR code generation**: A web service example using `chart.apis.google.com` to generate QR codes from text input via a form handler.

## Procedures And API Details
### Panic Recovery Pattern
1. Define a custom error type satisfying the `error` interface (e.g., `type Error string`).
2. Create an internal method that panics with this error type (e.g., `func (regexp *Regexp) error(err string)`).
3. Use a deferred function in functions like `Compile` to catch the panic:
   ```go
   defer func() {
       if e := recover(); e != nil {
           regexp = nil // Clear return value.
           err = e.(Error) // Re-panic if not a parse error.
       }
   }()
   ```
4. Check the type of recovered error to distinguish between expected parse errors and unexpected run-time failures.

### Web Server Implementation
1. Define a `main` function that parses flags, binds a handler (`QR`) to the root path, and starts `http.ListenAndServe`.
2. Use `flag.String` to set default ports (e.g., `:1718`).
3. Load an HTML template using `template.Must(template.New("qr").Parse(templateStr))`.
4. Implement the handler `QR(w http.ResponseWriter, req *http.Request)` which executes the template with the form value `req.FormValue("s")`.
5. Define `templateStr` containing HTML and Go expressions (`{{if .}}`, `{{.}}`) to conditionally render QR images based on input data.

## Nuance Or Contradictions
- While `panic` is used for exceptional conditions, the text notes that if an unexpected failure occurs (e.g., index out of bounds) within a recovery block, the type assertion `err = e.(Error)` will fail, causing a new panic and continuing stack unwinding. This ensures that only intended parse errors are handled gracefully.
- The re-panic idiom changes the panic value if an actual error occurs; both original and new failures appear in crash reports, which is usually sufficient but can be filtered if only the original value is desired.

## Candidate Wiki Hints
- **Topic: Panic Recovery Pattern** – A reusable source for implementing robust error handling in Go libraries where panics are internal signals for specific conditions (like parse errors) and external consumers expect `error` values.
- **Topic: Deferred Recover Usage** – Notes on the constraint that `recover` must be called from a deferred function to effectively stop stack unwinding.
- **Topic: Internal vs External Errors** – Guidance on keeping panic handling within a package and converting internal panics to returnable errors for public APIs.
