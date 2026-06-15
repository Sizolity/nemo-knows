---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

## Chunk Context
The chunk covers special forms in the `typing` module, starting with `Union`, `Optional`, and higher-order function annotations like `Concatenate`. It continues through literal types (`Literal`), class variables (`ClassVar`), immutability markers (`Final`), TypedDict modifiers (`Required`, `NotRequired`, `ReadOnly`), metadata handling (`Annotated`), and type predicate functions (`TypeIs`, `TypeGuard`).

## Local Summary
This section details advanced typing features introduced across Python versions 3.5 to 3.14. It explains how to construct unions, handle optional values without default arguments, annotate decorators with parameter transformation, define restricted value sets, mark class-level attributes, enforce immutability at the type-checking level, specify TypedDict key requirements, attach arbitrary metadata to types, and implement precise type narrowing for conditional logic.

## Key Claims
- **Union Syntax**: `Union[X, Y]` is equivalent to `X | Y`. Unions of unions are flattened automatically, but references through a type alias (`type A = Union[int, str]`) prevent flattening to avoid evaluating the underlying `TypeAliasType`.
- **Optional vs Default**: `Optional[X]` (or `X | None`) indicates that `None` is an allowed value. This is distinct from an optional argument with a default value; a parameter with a default does not automatically imply it can be `None`.
- **Concatenate Usage**: Used with `Callable` and `ParamSpec` to annotate decorators that modify function signatures (e.g., adding a lock). The last argument must be a `ParamSpec` or ellipsis.
- **Literal Types**: Define specific allowed values. Nested literals are flattened, but again, type alias references prevent flattening. Equality comparisons ignore order.
- **ClassVar and Final**: `ClassVar` marks attributes intended for the class body, not instances. `Final` prevents reassignment in any scope. Both can now be nested in each other since version 3.13.
- **Annotated Metadata**: Adds context-specific metadata to a type. Metadata is stored in `__metadata__`. Order matters for equality checks. Nested annotations are flattened unless referenced via a type alias.
- **TypeIs vs TypeGuard**: Both mark user-defined type predicate functions. `TypeIs` implies the argument type is the intersection of the original and narrowed type (if True) or excludes the narrowed type (if False). `TypeGuard` narrows to the exact type inside if True.

## Entities And Concepts
- **Special Forms**: `Union`, `Optional`, `Concatenate`, `Literal`, `ClassVar`, `Final`, `Required`, `NotRequired`, `ReadOnly`, `Annotated`, `TypeIs`, `TypeGuard`.
- **PEP References**: PEP 613 (TypeAlias), PEP 612 (ParamSpec/Concatenate), PEP 586 (Literal), PEP 526 (ClassVar), PEP 591 (Final), PEP 647 (TypeGuard), PEP 655 (TypedDict total=False), PEP 705 (ReadOnly), PEP 593 (Annotated), PEP 742 (TypeIs).
- **Attributes**: `__metadata__` (stores metadata in `Annotated`), `__origin__` (returns the underlying type in `Annotated`).

## Procedures And API Details
- **Defining a Union**: Use `Union[X, Y]` or shorthand `X | Y`. Avoid writing `Union[X][Y]`.
- **Creating a Literal**: Use `Literal['value']` or `Literal[1, 2]`. Nested literals like `Literal[Literal[1, 2], 3]` are flattened to `Literal[1, 2, 3]`.
- **Marking Class Variables**: `class Starship: stats: ClassVar[dict[str, int]] = {}`.
- **Final Assignment**: `MAX_SIZE: Final = 9000`. Cannot be reassigned or overridden in subclasses.
- **TypedDict Modifiers**:
  - Required: `key: ReadOnly[str]` (actually `Required` marks keys as required, `ReadOnly` marks items).
  - NotRequired: Marks keys as potentially missing.
  - ReadOnly: Marks items of a TypedDict as read-only.
- **Annotated Usage**: `Annotated[int, ValueRange(-10, 5)]`. Retrieve metadata via `. __metadata__`. Get origin type via `. __origin__` or `get_origin()`.
- **Type Predicate Functions**:
  - Signature: `def func(arg: TypeA) -> TypeIs[TypeB]` or `-> TypeGuard[TypeC]`.
  - Behavior: If function returns True, the type checker narrows the argument type accordingly.

## Nuance Or Contradictions
- **Type Alias Evaluation**: Flattening rules for `Union`, `Literal`, and `Annotated` do not apply when these types are referenced through a type alias (e.g., `type A = Union[int, str]`). This is to prevent forcing the evaluation of the underlying `TypeAliasType`.
- **Order Sensitivity**: For `Annotated`, the order of metadata elements matters for equality checks (`Annotated[int, A, B] != Annotated[int, B, A]`).
- **Runtime vs Static**: These constructs (e.g., `Final`, `ClassVar`, `ReadOnly`) do not change Python runtime behavior; they are solely for static type checkers.
- **get_origin() Behavior**: For `Annotated` types, `get_origin()` returns `typing.Annotated` itself, not the underlying type (unlike other generic types).

## Candidate Wiki Hints
- **Union Type Expressions**: Documenting the shorthand `X | Y` and version history.
- **Type Predicates**: A dedicated page comparing `TypeIs` vs `TypeGuard`.
- **TypedDict Modifiers**: Explaining `Required`, `NotRequired`, and `ReadOnly`.
- **Annotated Types**: Deep dive into metadata storage, flattening rules, and usage with generics.
