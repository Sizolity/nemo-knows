---
title: Typeddict Syntax And Modifiers
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

# Typeddict Syntax And Modifiers

The `typing.TypedDict` class in Python enables the definition of dictionary types with specific key requirements, allowing developers to distinguish between required and optional fields. This capability is essential for maintaining data integrity within complex data structures while leveraging static type checkers like `mypy`.

## Core Syntax

Typed dictionaries are defined by inheriting from `typing.TypedDict` (or simply `TypedDict` in Python 3.12+). The primary mechanism for controlling field behavior involves the `total` argument and specific key modifiers.

- **Total Mode**: By default, TypedDicts are "total", meaning all keys listed in the definition must be present in an instance. This is controlled by the `total=True` argument (implicit by default) or explicitly set via `@dataclass(frozen=True)` logic where applicable.
- **Optional Keys**: To allow a key to be missing, use the `NotRequired` modifier from the `typing` module. For example:

  ```python
  class User(TypedDict):
      name: str
      age: NotRequired[int]
  ```

- **Required Keys**: Keys without modifiers are implicitly required in total mode. To explicitly enforce requirement in a non-total context, use the `Required` modifier.

## Key Modifiers

The standard library provides specific types to annotate keys based on their presence constraints:

### Required

Use `typing.Required[T]` to denote a key that must be present. This is particularly useful when defining a TypedDict with `total=False`, ensuring that specific fields are mandatory even if others are optional.

```python
class Config(TypedDict, total=False):
    debug: NotRequired[bool]
    host: Required[str]  # Must always be present
```

### NotRequired

Use `typing.NotRequired[T]` to denote a key that may be absent. This modifier effectively creates an optional field without needing to rely on the `total` flag alone, making the intent explicit for readers and type checkers.

## Runtime Behavior

It is important to distinguish between static analysis and runtime enforcement. Type hints, including TypedDict definitions, are ignored by the Python interpreter at runtime. Enforcement relies entirely on third-party tools like `mypy`. While a standard dictionary can be passed where a TypedDict is expected at runtime (due to duck typing), type checkers will flag violations of key requirements or types defined in the TypedDict.

## Deprecated Features

Developers should avoid using the functional syntax for defining TypedDicts, such as:

```python
User = TypedDict("User", {"name": str})
```

This syntax was deprecated in Python 3.13 and removed in favor of the class-based definition pattern introduced in PEP 589. Similarly, legacy collection ABCs and functional aliases like `typing.Dict` are deprecated since Python 3.9.
