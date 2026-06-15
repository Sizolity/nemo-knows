---
title: Chunk 02 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

## Chunk Context

The chunk covers advanced type system features in Python, specifically focusing on generic classes (`type[BasicUser | ProUser]`), annotations for generators and coroutines, user-defined generics (including `ParamSpec`), the `Any` type versus `object`, structural subtyping via PEP 544, and special typing primitives like `LiteralString`, `Never`, `Self`, and `TypeAlias`.

## Local Summary

This section of the documentation details how to annotate generators (`Generator`, `AsyncGenerator`) and coroutines (`Coroutine`), explaining variance rules for `SendType`. It introduces user-defined generic classes, their inheritance from `Generic[T]`, and handling multiple type variables or parameter specification variables (`**P`). The text distinguishes between `Any` (dynamic compatibility) and `object` (structural subtyping). It also covers the shift from nominal to structural subtyping (PEP 544) allowing implicit ABC support. Finally, it documents special typing primitives: `AnyStr`, `LiteralString`, `Never/NoReturn`, `Self`, and `TypeAlias`.

## Key Claims

- Generators are annotated using `Generator[YieldType, SendType, ReturnType]`, where `SendType` behaves contravariantly.
- Async generators (`AsyncGenerator`) omit the return type argument; `AsyncIterable` and `AsyncIterator` variants exist.
- Coroutines use `Coroutine[YieldType, SendType, ReturnType]`.
- User-defined generic classes can inherit from `Generic[T]` or other generic classes like `Mapping[str, T]`.
- Generic types can be parameterized at runtime via `__class_getitem__()`.
- `Any` is compatible with every type and allows any operation without static checking; it defaults functions lacking annotations.
- `object` is a supertype of all types but does not allow arbitrary operations like `Any`.
- PEP 544 enables structural subtyping, allowing classes to be implicitly considered subtypes of ABCs without explicit inheritance.
- `LiteralString` restricts inputs to literal strings only (PEP 675).
- `Never` and `NoReturn` represent the bottom type; they are equivalent in static checkers.
- `Self` ensures return types reflect the actual class instance, including subclasses.

## Entities And Concepts

- **Generators**: Annotated with `Generator[YieldType, SendType, ReturnType]`.
- **Coroutines**: Annotated with `Coroutine[YieldType, SendType, ReturnType]`.
- **Async Generators**: Use `AsyncGenerator[YieldType, SendType]` (no return type).
- **Generic Classes**: Defined via `class Name[T]: ...` or explicit `Generic[T]` inheritance.
- **ParamSpec**: Used for parameter expressions with syntax `[**P]`.
- **Any**: A special type indicating unconstrained compatibility.
- **object**: The root of the nominal hierarchy; used for typesafe dynamic typing.
- **Structural Subtyping**: Implicitly checking ABC conformance (PEP 544).
- **Protocol**: Base class for defining custom protocols enabling structural subtyping.
- **LiteralString**: Special type for literal strings only (PEP 675).
- **Never/NoReturn**: Bottom types indicating non-returning functions or invalid calls.
- **Self**: Represents the current enclosed class instance in return annotations.
- **TypeAlias**: Annotation for explicitly declaring type aliases.

## Procedures And API Details

### Annotating Generators
```python
from typing import Generator

def echo_round() -> Generator[int, float, str]:
    sent = yield 0
    while sent >= 0:
        sent = yield round(sent)
    return 'Done'

# Default types for SendType and ReturnType are None
def infinite_stream(start: int) -> Generator[int]:
    while True:
        yield start
        start += 1
```

### Annotating Async Generators
```python
from typing import AsyncGenerator

async def infinite_stream(start: int) -> AsyncGenerator[int]:
    while True:
        yield start
        start = await increment(start)
```

### User-Defined Generic Classes
```python
from logging import Logger
from typing import TypeVar, Generic

T = TypeVar('T')

class LoggedVar(Generic[T]):
    def __init__(self, value: T, name: str, logger: Logger) -> None:
        self.value = value

def zero_all_vars(vars):
    for var in vars:
        var.set(0)
```

### Using ParamSpec
```python
from typing import ParamSpec, Generic

P = ParamSpec('P')

class Z(Generic[P]):
    pass

# Equivalent forms accepted by type checkers
X[int, str]  # Internally converted to X[[int, str]]
X[[int, str]]
```

### Structural Subtyping Example
```python
from collections.abc import Sized, Iterable

class Bucket:  # No explicit base classes needed
    def __len__(self) -> int: ...
    def __iter__(self): ...

def collect(items: Iterable[int]) -> int: ...
result = collect(Bucket())  # Passes type check
```

### Special Type Usage Examples
```python
from typing import Any, LiteralString, Never, Self, TypeAlias

a: Any = None
a = []  # OK
a = 2   # OK

def run_query(sql: LiteralString) -> None: ...
run_query("SELECT * FROM students")  # OK
run_query(f"SELECT {user_input}")     # Error

def stop() -> Never:
    raise RuntimeError('no way')

class Foo:
    def return_self(self) -> Self:
        return self

Factors: TypeAlias = list[int]
```

## Nuance Or Contradictions

- **Any vs object**: While both allow dynamic typing, `Any` permits any operation without checking, whereas `object` rejects most operations except those defined on the base class.
- **SendType Variance**: Unlike most generic classes in the standard library, `Generator`'s `SendType` behaves contravariantly.
- **Deprecated Primitives**: `AnyStr` is deprecated since Python 3.13 and will be removed in 3.18; use parameter specification syntax instead.
- **Runtime vs Static Checking**: Some generic classes with `ParamSpec` may not have correct `__parameters__` after substitution because they are primarily intended for static type checking.
- **Version Changes**: PEP 695 introduces new type parameter syntax; previously, explicit inheritance from `Generic` or containing a type variable was required.

## Candidate Wiki Hints

- **Generators and Coroutines Annotations** – Guide on using `Generator`, `AsyncGenerator`, and `Coroutine` with correct variance rules.
- **User-Defined Generics** – How to define generic classes, inherit from `Generic[T]`, and use `ParamSpec`.
- **Any vs object** – Clarifying differences between dynamic compatibility (`Any`) and structural subtyping (`object`).
- **Structural Subtyping with Protocols** – Explaining PEP 544 and implicit ABC conformance.
- **Special Typing Primitives** – Usage of `LiteralString`, `Never`, `Self`, and `TypeAlias`.
