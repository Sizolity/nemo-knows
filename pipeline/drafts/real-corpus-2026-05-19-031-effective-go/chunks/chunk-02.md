---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

## Chunk Context
Lines 27-550 of `raw/web/corpus-2026-05-18/031-effective-go.md`. The text covers the "Effective Go" documentation, specifically focusing on formatting conventions (indentation, line length), naming conventions (packages, getters, interfaces), syntax details (semicolons, control structures like `if`, `for`, `switch`), and idiomatic code patterns (error handling, blank identifiers).

## Local Summary
This section of the document serves as a guide to writing clear, idiomatic Go code. It emphasizes that while formatting is often automated by `gofmt`, understanding its rules (tabs for indentation, no line length limits) is crucial. The text details naming conventions such as using MixedCaps for multi-word names, avoiding unnecessary getters (preferring capitalized field access), and adhering to standard interface suffixes like `-er`. It also explains the automatic semicolon insertion rule and how it influences code layout, particularly regarding control structures. Furthermore, it outlines the syntax of `if`, `for`, and `switch` statements, highlighting idioms like omitting `else` after error checks and using short declarations with `:=`.

## Key Claims
- Go programs should be formatted by `gofmt` to ensure consistency; manual formatting is discouraged unless necessary.
- Package names should be lowercase, single words, and concise (e.g., `bytes`, not `byteutil`).
- Getters should return capitalized field names (e.g., `Owner`) rather than using the prefix `Get`.
- Interface names often follow the pattern of `<Method>-er` (e.g., `Reader`, `Writer`).
- Semicolons are automatically inserted by the lexer unless they appear in specific contexts like `for` loop clauses.
- Control structures (`if`, `switch`) do not require parentheses around their expressions or conditions.
- The `else` clause is often omitted after an `if` statement that returns an error, allowing successful flow to continue naturally.
- Short declarations (`:=`) allow reassignment of existing variables in the same scope (e.g., `err`).
- The blank identifier `_` is used to ignore values in range loops when not needed.

## Entities And Concepts
- **gofmt**: A tool that automatically formats Go source code.
- **Package Name**: Should be short, concise, evocative, and lowercase (e.g., `bytes`).
- **Getter/Setter**: Methods for accessing/modifying fields; getters should not use the `Get` prefix.
- **Interface Naming**: Conventionally uses `-er` suffix for agent nouns (e.g., `Reader`).
- **Semicolon Insertion**: A lexer rule that inserts semicolons automatically at newlines under certain conditions.
- **Blank Identifier**: The underscore `_` used to discard values in loops or assignments.
- **Control Structures**: `if`, `for`, `switch`, `select`.
- **Short Declaration**: The `:=` operator for declaring and initializing variables.

## Procedures And API Details
- **Formatting Rules**:
  - Use tabs for indentation.
  - No line length limit; wrap lines if they feel too long.
  - Do not align comments manually; let `gofmt` handle it.
- **Naming Conventions**:
  - Exported names use uppercase first letter (e.g., `Owner`).
  - Avoid `Get` prefix for getters (use `Owner` instead of `GetOwner`).
  - Use MixedCaps for multi-word names (e.g., `MixedCaps`).
- **Code Patterns**:
  - Omit `else` when the `if` body returns an error.
  - Use `:=` to declare and assign variables in a single statement.
  - Use `_` in range loops to ignore unused values (e.g., `for _, value := range array`).
- **Control Structure Syntax**:
  - `if` statements do not require parentheses: `if condition { }`.
  - `switch` can be used as an if-else chain without an expression.
  - Labels are used with `break` to exit nested loops or switches (e.g., `break Loop`).

## Nuance Or Contradictions
- **Semicolon Visibility**: Although the formal grammar uses semicolons, they rarely appear in source code due to automatic insertion by the lexer. However, they are still required in specific places like `for` loop clauses and multiple statements on a line.
- **Getter Naming**: While getters are not idiomatic by default, providing them is allowed if appropriate; however, the naming convention should avoid the `Get` prefix.
- **Interface Names**: Standard interfaces like `Reader`, `Writer` have canonical signatures; creating methods with similar names but different signatures can cause confusion.

## Candidate Wiki Hints
- **Effective Go Formatting Guide**
- **Go Naming Conventions (Packages, Getters, Interfaces)**
- **Semicolon Insertion Rules in Go**
- **Idiomatic Control Structures in Go (`if`, `for`, `switch`)**
- **Blank Identifier Usage in Range Loops**
