---
title: Chunk NN Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

## Chunk Context
The chunk covers the `typing` module's advanced features introduced in Python 3.10–3.14, specifically focusing on type narrowing (`TypeIs`, `TypeGuard`), unpacking operators for generic types (`Unpack`), and the mechanics of building generic types using `Generic`, `TypeVar`, and `TypeVarTuple`.

## Local Summary
This section details how to refine type information during runtime checks, handle variadic generics via unpacking syntax, and construct reusable generic classes and functions. It distinguishes between bounded and constrained type variables, explains variance inference, and describes the attributes added in recent Python versions for introspecting type variable state.

## Key Claims
- `TypeIs` narrows a variable's type to an intersection of its original type and the guard type, whereas `TypeGuard` narrows it strictly to the guard type.
- The unpack operator `*` is semantically equivalent to `typing.Unpack` in contexts like `tuple[*Ts]` or `**kwargs: Unpack[Movie]`.
- Bounded type variables (`S: str`) are solved using the most specific type available, while constrained type variables (`A: (str, bytes)`) must match exactly one of the specified constraints.
- Manually created type variables default to invariant behavior unless explicitly marked covariant or contravariant.

## Entities And Concepts
- **TypeIs**: A function that narrows types by intersection when returning `True`.
- **TypeGuard**: A function that narrows types strictly to the guard's return type when returning `True`.
- **Unpack**: An operator (or `typing.Unpack`) used to mark objects as unpacked, allowing variadic generics and TypedDict usage.
- **Generic**: Abstract base class for declaring generic types.
- **TypeVar**: Represents a placeholder for a specific type in generic definitions.
- **TypeVarTuple**: Enables parameterization with an arbitrary number of types (variadic generics).
- **Bounded Type Variable**: Restricted to subtypes of a specified upper bound.
- **Constrained Type Variable**: Restricted to exactly one of several specific types.

## Procedures And API Details
- **Declaring Generics**: Use `class Mapping[KT, VT]:` syntax; brackets implicitly inherit from `typing.Generic`.
- **Creating TypeVars**:
  - Standard: `T = TypeVar('T')`
  - Bounded: `S = TypeVar('S', bound=str)`
  - Constrained: `A = TypeVar('A', str, bytes)`
- **Using Unpack**:
  ```python
  from typing import TypedDict, Unpack
  class Movie(TypedDict):
      name: str
      year: int

  def foo(**kwargs: Unpack[Movie]): ...
  ```
- **Variadic Generics**:
  ```python
  Ts = TypeVarTuple('Ts')
  def move_first_element_to_last[T, *Ts](tup: tuple[T, *Ts]) -> tuple[*Ts, T]:
      return (*tup[1:], tup[0])
  ```

## Nuance Or Contradictions
- **Syntax Evolution**: In Python <= 3.10, `*` could not be used directly in certain type contexts (e.g., `tuple[*Ts]`), requiring explicit use of `Unpack[Ts]`. From 3.11+, `*` is supported directly.
- **Evaluation Timing**: For type variables created via PEP 695 syntax, attributes like `__bound__`, `__constraints__`, and `__default__` are lazily evaluated (evaluated only when accessed).
- **Runtime Behavior**: Calling `isinstance(x, T)` on a manually created `TypeVar` raises `TypeError`.

## Candidate Wiki Hints
- **Topic: Type Narrowing Strategies** (Explaining differences between `TypeIs` and `TypeGuard`).
- **Topic: Variadic Generics in Python** (Usage of `TypeVarTuple` and `Unpack`).
- **Topic: Constructing Generic Types** (Using `Generic`, `TypeVar`, and syntax variations).
