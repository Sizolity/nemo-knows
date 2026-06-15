---
title: Go Idioms
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

# Go Idioms

Go idioms are the conventions, patterns, and language‑aware practices that distinguish effective Go code from a direct translation of other languages. They shape everything from formatting and naming to concurrency and error handling, and they help programs feel natural to other Go developers.

The foundation is mechanical formatting: using `gofmt` (or `go fmt`) to produce a single canonical layout eliminates stylistic debates. Names follow a distinctive scheme—package names are short, lower‑case, and single‑word; exported identifiers use `MixedCaps`, and getters drop the “Get” prefix (for example, `Owner` instead of `GetOwner`). One‑method interfaces often adopt the “‑er” suffix (`Reader`, `Writer`). Because the lexer inserts semicolons, opening braces must remain on the same line as the control keyword.

Control structures are streamlined. `if`, `for`, and `switch` can each embed a short initialisation statement. `switch` evaluates cases top‑to‑bottom without automatic fall‑through, and `for` is the only loop needed. Functions routinely return multiple values—typically a result and an error—which replaces in‑band error returns and address‑based output parameters. Named result parameters can clarify intent and enable bare `return`.

`defer` schedules a function call to run just before the enclosing function returns; it is the idiomatic way to pair operations such as open/close or lock/unlock. The allocation primitives `new` and `make` serve distinct roles: `new(T)` zeroes memory and returns a pointer, while `make` initialises slices, maps, and channels and returns a ready‑to‑use value. Slices, not arrays, are the workhorse sequence type—arrays are full values that copy on assignment, whereas a slice provides reference‑like behavior.

Go interfaces are satisfied implicitly; a type needs no declaration of intent. Constructors often return an interface value to conceal the concrete implementation. Instead of subclassing, struct and interface embeddings promote methods, offering a form of composition.

Concurrency idioms centre on the principle “do not communicate by sharing memory; instead, share memory by communicating.” Goroutines are lightweight threads of execution, and channels coordinate them safely. Unbuffered channels combine communication with synchronisation—a sender blocks until a receiver takes the value. A common pattern uses a channel to signal completion: a goroutine sends a token after finishing work, and the caller waits on that channel. Buffered channels decouple producers and consumers when asynchronous handoff is needed.

Errors are values, returned as the last return value, with a trivial `error` interface. The blank identifier (`_`) discards unwanted values, silences compiler complaints about unused imports during development, and enables compile‑time interface‑satisfaction checks. The `panic`/`recover` mechanism is reserved for truly exceptional situations and is often used inside libraries to translate panics into returned errors.

These idioms, originally described in the 2009 “Effective Go” document, remain central to the core language, though they do not address later additions such as generics.
