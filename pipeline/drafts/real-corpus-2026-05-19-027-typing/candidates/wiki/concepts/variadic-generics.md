---
title: Variadic Generics
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

# Variadic Generics

Variadic generics are a feature in Python's type system that allow generic classes or functions to accept a variable number of type parameters. This capability enables the definition of complex generic structures where the arity (number of type arguments) is not fixed at compile time.

## Definition and Syntax

In Python, variadic generics are primarily implemented using `TypeVarTuple` combined with the unpacking operator (`*`) or the `Unpack[Ts]` syntax. This allows users to define types that behave similarly to tuple types but with dynamic length constraints.

The standard library provides specific support for this through:
- **`tuple[*Ts]`**: Defines a generic tuple where `Ts` is a tuple of type variables representing each element's type.
- **`Unpack[Ts]`**: Used within generics to unpack a tuple of types into the current context, allowing variable-length type arguments.

## Usage Context

Variadic generics are essential for:
1.  **Generic Tuples**: Creating generic types that preserve the specific structure and types of elements regardless of how many elements are present.
2.  **Advanced Decorators**: Supporting decorators that need to inspect or manipulate arbitrary numbers of type parameters passed by a class or function.
3.  **Library Compatibility**: Enabling libraries to define APIs that accept variable-length generic arguments without sacrificing type safety.

## Implementation Details

Under the hood, variadic generics rely on `TypeVarTuple` objects. When defining a variadic generic, the syntax typically involves:
- Declaring a `TypeVarTuple` (e.g., `Ts = TypeVarTuple("Ts")`).
- Using this tuple in the type definition alongside the unpacking operator.

This mechanism ensures that the type checker can correctly infer and validate the structure of the generic arguments during static analysis, even when the exact number of types is unknown at the point of definition.

## Static Analysis Support

Static type checkers like `mypy` enforce these constraints. While the Python runtime ignores type hints, tools verify that variadic generics are used correctly, ensuring that the unpacked types align with the expected structure of the generic container (e.g., a tuple).
