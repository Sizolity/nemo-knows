---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

# Chunk Context
This chunk covers the functional and class-based syntaxes for `TypedDict`, inheritance rules, generic support, introspection attributes (`__total__`, `__required_keys__`, etc.), deprecation notices for older creation methods, and introduces protocols (such as `SupportsAbs`) and I/O type classes.

# Local Summary
The section explains how to define `TypedDict` using functional syntax versus class syntax, highlighting support for required/optional keys via `NotRequired`/`Required`, total/optional definitions via the `total` argument, and inheritance between TypedDicts. It details introspection attributes available since Python 3.9 and read-only/mutable key tracking added in 3.13. The chunk concludes with protocols for built-in methods (like `__abs__`) and I/O stream types.

# Key Claims
- By default, all keys in a `TypedDict` are required unless marked otherwise.
- Individual keys can be optional using `NotRequired`, or all keys can be optional by setting `total=False`.
- TypedDicts can inherit from other TypedDicts but not from non-TypedDict classes (except `Generic`).
- TypedDicts support generic types via inheritance from `Generic[T]` in Python 3.11 and lower.
- Introspection attributes like `__required_keys__` reflect the actual set of required/optional keys, accounting for `NotRequired` and inheritance.
- Since Python 3.13, `ReadOnly` qualifier is supported; `__readonly_keys__` and `__mutable_keys__` are provided.
- The functional syntax with missing or None fields is deprecated in Python 3.13 and will be removed in 3.15.

# Entities And Concepts
- `TypedDict`: A dictionary-like type definition for static typing.
- `NotRequired`: Marks a key as optional.
- `Required`: Marks a key as required (inferred when not marked otherwise).
- `total`: Boolean argument controlling whether all keys are required by default.
- `__total__`, `__required_keys__`, `__optional_keys__`: Introspection attributes on TypedDict.
- `Generic[T]`: Base class for generic TypedDicts.
- `SupportsAbs`, `SupportsInt`, etc.: Protocol classes for numeric methods.
- `IO[AnyStr]`, `TextIO`, `BinaryIO`: Generic types for I/O streams.

# Procedures And API Details
- Define a TypedDict with required keys:
  ```python
  class Point2D(TypedDict):
      x: int
      y: int
  ```
- Define a TypedDict with optional keys:
  ```python
  class Point2D(TypedDict, total=False):
      x: int
      y: int
  ```
- Mark specific keys as required within a `total=False` context:
  ```python
  class Point2D(TypedDict, total=False):
      x: Required[int]
      y: Required[int]
      label: str
  ```
- Create a generic TypedDict for Python 3.11 and below:
  ```python
  T = TypeVar("T")

  class Group(TypedDict, Generic[T]):
      key: T
      group: list[T]
  ```
- Use `typing.cast` to hint types at runtime without checking:
  ```python
  typing.cast(typ, val)
  ```
- Use `typing.assert_type` for static type checking assertions:
  ```python
  typing.assert_type(val, typ)
  ```

# Nuance Or Contradictions
- The attribute `__total__` reflects only the `total` argument value and does not fully capture semantic requirements (e.g., a class with `total=True` can still have optional keys via `NotRequired`).
- Inheritance allows declaring required/optional keys by inheriting from a TypedDict with a different `total` value.
- Using `from __future__ import annotations` or string annotations may break introspection of `__required_keys__` and `__optional_keys__`.

# Candidate Wiki Hints
- **TypedDict Syntax**: Compare functional vs class-based definitions.
- **Optional/Required Keys**: Use `NotRequired`, `Required`, and `total` arguments.
- **Generic TypedDicts**: Inherit from `Generic[T]` for older Python versions.
- **Introspection Attributes**: Understand differences between `__total__`, `__required_keys__`, and `__optional_keys__`.
- **Readonly Keys**: Use `ReadOnly` qualifier and inspect with `__readonly_keys__`.
- **Protocols**: List of built-in protocols for numeric and I/O operations.
