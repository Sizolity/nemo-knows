---
title: Python typing module summary
kind: source
sources:
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

## What It Is

The `typing` module in Python provides runtime support for type hints, enabling developers to annotate function arguments, return values, and variables. These hints assist static type checkers, IDEs, and linters to catch errors early without enforcing constraints at the Python runtime itself. The documentation covers features introduced across Python versions 3.5 through 3.14, including new syntax for generics (PEP 695), type aliases (`type` statement), protocols (structural subtyping via PEP 544), and special directives like `@override`, `@final`, and `assert_never`. While the module historically provided aliases for built-in types (e.g., `typing.Dict`), these are deprecated since Python 3.9 in favor of direct subscripting (`dict[str, int]`).

## Summary

The `typing` module enables precise type checking by allowing developers to define complex types such as generics, unions, protocols, and typed dictionaries. It distinguishes between runtime behavior and static analysis: while the interpreter ignores annotations, tools like `mypy` enforce them. Key capabilities include creating distinct subtypes with `NewType`, defining generic classes with `Generic[T]`, using `ParamSpec` for decorators, and implementing structural subtyping via protocols. The module also provides utilities for type narrowing (`TypeIs`, `TypeGuard`), exhaustiveness checking (`assert_never`), and introspection (`get_type_hints`, `reveal_type`). Since Python 3.12, features like the `type` statement for aliases and PEP 695 syntax have streamlined type definitions. Deprecated features, such as functional syntax for `TypedDict` (removal in 3.15) and legacy collection ABCs, are phased out in favor of modern standards.

## Key Claims

- **Runtime vs. Static Enforcement**: Type hints are ignored by the Python runtime; enforcement is handled entirely by third-party static type checkers and IDEs.
- **Type Aliases and NewType**: The `type` statement (Python 3.12+) or `TypeAlias` creates equivalent types, while `NewType` creates subtypes to prevent logical errors without runtime overhead.
- **Generics and Variance**: Generics like `Generator[YieldType, SendType, ReturnType]` support variance rules; `SendType` is contravariant. User-defined generics inherit from `Generic[T]` or use PEP 695 syntax (`def first[T](...)`).
- **Protocols and Structural Subtyping**: Protocols enable duck-typing via structural subtyping (PEP 544), allowing implicit ABC conformance. The `runtime_checkable()` decorator enables `isinstance()` checks based on attribute presence rather than signatures.
- **TypedDict**: Typed dictionaries support required/optional keys via `Required`, `NotRequired`, and the `total` argument. They inherit from `Generic[T]` for generic support in Python 3.11 and lower.
- **Special Forms**: Union types use `X | Y`; `Optional[X]` is equivalent to `X | None`. `Literal` restricts values, `Final` prevents reassignment, and `ClassVar` marks class attributes.
- **Type Narrowing**: `TypeIs` narrows types by intersection; `TypeGuard` narrows strictly to the guard type. Both are used with predicate functions returning booleans.
- **Decorators and Utilities**: `@overload` supports multiple signatures for type checking; `@override` ensures method overrides; `@final` prevents subclassing; `assert_never` flags unreachable code.
- **Unpacking and Variadic Generics**: The `*` operator or `Unpack[Ts]` enables variadic generics like `tuple[*Ts]` and function arguments with variable-length type parameters (`TypeVarTuple`).
- **Deprecated Features**: Aliases for built-ins (`typing.Dict`) are deprecated since Python 3.9; functional `TypedDict` syntax is deprecated in 3.13; `@no_type_check_decorator` will be removed in 3.15.

## Suggested Links

- raw/web/corpus-2026-05-18/027-typing.md
- none
