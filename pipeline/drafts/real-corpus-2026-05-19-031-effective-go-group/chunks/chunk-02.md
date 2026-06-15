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
