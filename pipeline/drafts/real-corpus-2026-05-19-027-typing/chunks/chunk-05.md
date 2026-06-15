---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

## Chunk Context
This chunk covers advanced type system features in Python, specifically focusing on **type variable tuples** (PEP 646), **parameter specification variables** (`ParamSpec`, PEP 612), and **protocol handling** (PEP 544). It also details attributes for these types (e.g., `__default__`, `args`, `kwargs`) and introduces the `typing.TypeAliasType` and `typing.TypedDict`. The text spans lines 1925–2531 of the source document.

## Local Summary
The document explains how to use unpacked type variable tuples (`*Ts`) to distinguish them from regular type variables, ensuring correct usage in class definitions and function arguments. It details the `ParamSpec` class for forwarding parameter types between callables (decorators) and its runtime attributes (`args`, `kwargs`). The chunk also covers `typing.TypeAliasType` created via the `type` statement, including star unpacking support added in version 3.14. Finally, it describes `typing.Protocol`, `runtime_checkable()`, and `typing.TypedDict` for structural typing and dictionary type hints.

## Key Claims
- Type variable tuples must always be unpacked using the `*` syntax (e.g., `tuple[*Ts]`) to distinguish them from normal type variables.
- At most one type variable tuple may appear in a single list of type arguments or parameters; multiple unpacked tuples are invalid.
- `ParamSpec` is used to forward parameter types in decorators and higher-order functions, capturing both positional (`args`) and keyword (`kwargs`) arguments.
- `typing.TypeAliasType` represents aliases created via the `type` statement (PEP 617) and supports lazy evaluation of its value.
- `typing.TypedDict` is a special construct for adding type hints to dictionaries; instances are regular `dict`s at runtime, but key existence and types are enforced by static checkers.
- Runtime-checkable protocols use `inspect.getattr_static()` (since Python 3.12) instead of `hasattr()` for attribute lookup, potentially changing isinstance behavior.

## Entities And Concepts
- **Type Variable Tuples**: Introduced in PEP 646, represented as `tuple[T, *Ts]` or `*Ts`. Used to represent variable-length type parameters.
- **ParamSpec**: A specialized type variable (PEP 612) for capturing function signatures in decorators. Attributes include `args`, `kwargs`, `__name__`, and `__default__`.
- **TypeAliasType**: The runtime class for types created with the `type` statement (Python 3.12+). Supports lazy evaluation via `__value__` and star unpacking (`*Alias`).
- **Protocol**: A base class for structural typing (duck-typing) recognized by static type checkers. Can be generic.
- **runtime_checkable**: Decorator to make a protocol usable with `isinstance()` and `issubclass()`. Checks only attribute presence, not signatures.
- **TypedDict**: A dictionary subclass with typed keys enforced by static checkers. Supports class-based and functional syntaxes.

## Procedures And API Details
### Using Type Variable Tuples
```python
x: tuple[*Ts]  # Correct usage
class Array[DType, *Shape]: pass
```
- Invalid: `x: Ts`, `x: tuple[Ts]`, `x: tuple[*Ts, *Ts]`.

### ParamSpec Usage
```python
from collections.abc import Callable
import logging

def add_logging[T, **P](f: Callable[P, T]) -> Callable[P, T]:
    def inner(*args: P.args, **kwargs: P.kwargs) -> T:
        logging.info(f'{f.__name__} was called')
        return f(*args, **kwargs)
    return inner

@add_logging
def add_two(x: float, y: float) -> float:
    return x + y
```
- Access `P.args` and `P.kwargs` to annotate `*args` and `**kwargs`.
- `ParamSpecArgs` and `ParamSpecKwargs` are runtime types for `P.args` and `P.kwargs`.

### TypeAliasType Examples
```python
type Alias = int
type ListOrSet[T] = list[T] | set[T]
>>> Alias.__value__  # Lazily evaluated
>>> call_evaluate_function(Alias.evaluate_value, Format.FORWARDREF)
ForwardRef('undefined')
```
- Star unpacking: `type Unpacked = tuple[bool, *Alias]`.

### Protocol and TypedDict
```python
@runtime_checkable
class Closable(Protocol):
    def close(self): ...

assert isinstance(open('/some/file'), Closable)

class Point2D(TypedDict):
    x: int
    y: int
```

## Nuance Or Contradictions
- **Version Changes**:
  - Python 3.12 changed `isinstance()` checks for runtime-checkable protocols to use `inspect.getattr_static()`, which may alter results compared to Python 3.11.
  - Python 3.14 added star unpacking support for type aliases (`*Alias`).
- **Runtime Behavior**: `NewType` is distinct to static checkers but returns the argument unchanged at runtime. `TypedDict` instances are plain dicts at runtime.
- **Deprecated Syntax**: Functional syntax for `NamedTuple` without fields (`NT = NamedTuple("NT")`) is deprecated since 3.13. Keyword argument syntax for `NamedTuple` creation is also deprecated.

## Candidate Wiki Hints
- **Page: Type Variable Tuples** – Covers PEP 646, unpacking syntax, and usage constraints.
- **Page: ParamSpec** – Explains decorator patterns, `args`/`kwargs` splitting, and runtime introspection.
- **Page: Protocol Typing** – Details structural typing, `runtime_checkable`, and performance considerations for `isinstance()`.
- **Page: TypedDict** – Usage with class and functional syntaxes, key constraints, and runtime behavior.
