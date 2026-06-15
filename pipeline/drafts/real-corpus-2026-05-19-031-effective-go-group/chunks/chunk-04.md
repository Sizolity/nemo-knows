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
