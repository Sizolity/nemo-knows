## chunk-01

---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

## Chunk Context
This chunk covers the initial introduction to Python's `typing` module within the 3.14.5 documentation, specifically lines 1–482 of the raw corpus. It details how the module provides runtime support for type hints (though enforcement is handled by static checkers), introduces core concepts like type aliases and `NewType`, and explains annotations for callables and generic types.

## Local Summary
The `typing` module enables developers to add type hints to function arguments, return values, and variables to assist static type checkers, IDEs, and linters. While the Python runtime does not enforce these hints, they provide significant utility for catching errors early. The text distinguishes between simple built-in types (e.g., `float`, `str`) and complex hints provided by the module. It highlights that new features often appear in `typing` first, with `typing_extensions` offering backports for older Python versions.

## Key Claims
- **Runtime vs. Static**: The Python runtime does not enforce function or variable type annotations; enforcement is left to third-party tools like type checkers and IDEs.
- **Type Aliases**: Defined using the `type` statement (new in Python 3.12) or assignment, these make two types equivalent for static checkers. For example, `type Vector = list[float]` treats `Vector` as exactly equivalent to `list[float]`.
- **NewType**: Creates distinct subtypes of existing types (e.g., a specific ID type based on `int`) to help catch logical errors without runtime overhead. A `Derived` type is a subtype of its base, preventing usage of the base where the derived is expected.
- **Callable Annotations**: Functions are annotated using `collections.abc.Callable` or `typing.Callable`. The signature requires a list of argument types and a single return type (e.g., `Callable[[int], str]`).
- **Generics**: Container classes support subscription to denote element types (e.g., `Sequence[Employee]`). Functions can be parameterized using PEP 695 syntax (`def first[T](...)`) or `TypeVar`.

## Entities And Concepts
- **typing module**: Provides runtime support for type hints.
- **type alias**: A name that refers to another type, created via `type` statement or assignment, treated as equivalent by static checkers.
- **NewType**: Helper function/class to create a distinct subtype of an existing type for logical error prevention.
- **Callable**: Interface for annotating functions; supports argument lists and return types.
- **Protocol**: Used to define callables with complex signatures (e.g., variadic arguments, keyword-only parameters) that standard `Callable` cannot express.
- **ParamSpec / Concatenate**: Operators introduced in Python 3.10 to handle callables dependent on each other or those adding/removing arguments.
- **TypeVar**: Factory for declaring type variables used in generic functions and classes.

## Procedures And API Details
- **Creating a Type Alias**:
  ```python
  type Vector = list[float]
  # Or with explicit marker:
  from typing import TypeAlias
  Vector: TypeAlias = list[float]
  ```
- **Using NewType**:
  ```python
  from typing import NewType
  UserId = NewType('UserId', int)
  # Usage: get_user_name(user_id: UserId) -> str
  ```
- **Annotating a Function with Callable**:
  ```python
  from collections.abc import Callable, Awaitable
  def feeder(get_next_item: Callable[[], str]) -> None: ...
  ```
- **Defining Generics**:
  ```python
  # PEP 695 syntax (Python 3.12+)
  def first[T](l: Sequence[T]) -> T: return l[0]

  # TypeVar syntax
  from typing import TypeVar
  U = TypeVar('U')
  def second(l: Sequence[U]) -> U: return l[1]
  ```
- **Tuple Annotations**:
  - Fixed length: `tuple[int, str]`
  - Variable length (same type): `tuple[T, ...]`
  - Empty tuple: `tuple[()]`
  - Any tuple: `tuple` (equivalent to `tuple[Any, ...]`)

## Nuance Or Contradictions
- **Equivalence vs. Subtyping**: A `type alias` declares equivalence (`Alias` == `Original`), whereas `NewType` declares subtyping (`Derived` is a subclass of `Original`). Confusing these can lead to type checking errors where a base value is used where a derived one is expected.
- **Runtime Behavior**: `NewType` definitions are runtime functions that immediately return their argument; they do not create new classes at runtime, keeping overhead low. However, creating a class inheriting from a `NewType` (e.g., `class AdminUserId(UserId): pass`) fails at runtime and does not pass type checking.
- **Deprecated Aliases**: The documentation notes the existence of deprecated aliases like `typing.Callable` in favor of `collections.abc.Callable`.

## Candidate Wiki Hints
- **Page: Python Type System Overview**
  - Covers the distinction between static hints and runtime enforcement.
  - Explains `type alias` vs. `NewType` with examples.
- **Page: Advanced Callable Annotations**
  - Details `Callable`, `Protocol`, `ParamSpec`, and `Concatenate`.
- **Page: Generic Types in Python 3.12+**
  - Focuses on PEP 695 syntax for generic functions and classes.

## chunk-02

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

## chunk-03

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

## chunk-04

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

## chunk-05

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

## chunk-06

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

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
    - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

# Chunk Context

This chunk covers the `typing` module's utilities for static type checking, including exhaustiveness verification (`assert_never`), type revelation (`reveal_type`), dataclass-like behavior decoration (`dataclass_transform`), function overloading (`overload`, `get_overloads`), finality enforcement (`final`), disabling type checks (`no_type_check`, `no_type_check_decorator`), override enforcement (`override`), runtime-unavailable marking (`type_check_only`), and introspection helpers like `get_type_hints`, `get_origin`, `get_args`, `is_protocol`, `is_typeddict`, and `get_protocol_members`.

# Local Summary

The chunk documents several typing utilities added in Python 3.11, 3.12, and 3.13 that aid static type checkers in verifying exhaustiveness, overloading, finality, and other advanced typing patterns. It also covers decorators like `@dataclass_transform` for simulating dataclasses without runtime overhead, `@override` to ensure subclass methods override parent methods, and introspection tools to inspect types at runtime or statically.

# Key Claims

- `typing.assert_never(arg)` is used to assert that a branch of code is unreachable; if the type checker determines it is reachable, an error is emitted.
- `reveal_type(obj)` emits a diagnostic with the inferred type of an expression for debugging purposes.
- `@dataclass_transform` allows marking objects (classes, metaclasses, or decorators) as providing dataclass-like behavior to static type checkers.
- `@overload` enables defining multiple function signatures for type checking; only the non-decorated definition runs at runtime.
- `get_overloads(func)` returns a sequence of overloaded definitions for introspection.
- `@final` indicates that a method or class cannot be overridden or subclassed, respectively.
- `@no_type_check` tells type checkers to ignore annotations in the decorated function or class.
- `@override` ensures that a method in a subclass overrides a method from its parent.
- `get_type_hints(obj)` returns resolved type hints for an object, handling forward references and merging base class annotations.
- `get_origin(tp)` and `get_args(tp)` extract the origin and arguments of generic types.

# Entities And Concepts

- **Exhaustiveness Checking**: Ensures all cases in a match statement are covered; uses `assert_never` to assert unreachable branches.
- **Type Revelation**: Uses `reveal_type` to inspect inferred types during development.
- **Dataclass-like Behavior**: Simulates dataclasses using `@dataclass_transform`.
- **Function Overloading**: Supports multiple signatures via `@overload` and introspection via `get_overloads`.
- **Finality Enforcement**: Prevents overriding or subclassing with `@final`.
- **Type Check Disabling**: Ignores annotations in decorated functions/classes with `@no_type_check`.
- **Override Enforcement**: Ensures method overrides using `@override`.
- **Runtime-Unavailable Classes**: Marks classes unavailable at runtime with `@type_check_only`.
- **Type Introspection**: Tools like `get_type_hints`, `get_origin`, `get_args` for inspecting types.

# Procedures And API Details

### assert_never(arg, /)
- Asserts that a line of code is unreachable.
- Emits an error if the type checker finds it reachable.
- Throws an exception at runtime.

### reveal_type(obj, /)
- Reveals the inferred type of an expression to a static type checker.
- Prints the runtime type to `sys.stderr` and returns the argument unchanged.

### @dataclass_transform(*, eq_default=True, order_default=False, kw_only_default=False, frozen_default=False, field_specifiers=(), **kwargs)
- Decorator to mark objects as providing dataclass-like behavior.
- Can be applied to classes, metaclasses, or decorator functions.
- Accepts parameters like `eq`, `order`, `frozen`, etc., which affect synthesized methods.
- Records arguments in `__dataclass_transform__` attribute at runtime.

### @overload
- Decorator for overloaded functions; only the non-decorated definition runs at runtime.
- Followed by exactly one non-decorated definition.

### get_overloads(func)
- Returns a sequence of `@overload`-decorated definitions for introspection.
- Returns an empty sequence if no overloads exist.

### clear_overloads()
- Clears all registered overloads in the internal registry.

### @final
- Marks methods or classes as final (cannot be overridden or subclassed).
- Sets `__final__` attribute to `True` at runtime.

### @no_type_check
- Tells type checkers to ignore annotations in decorated functions/classes.
- Mutates the decorated object in place.

### @no_type_check_decorator
- Wraps a decorator to apply `@no_type_check` effect.
- Deprecated since Python 3.13; will be removed in 3.15.

### @override
- Indicates that a method is intended to override a superclass method.
- Sets `__override__` attribute to `True` at runtime.

### @type_check_only
- Marks a class or function as unavailable at runtime.
- Intended for private classes in type stubs.

### get_type_hints(obj, globalns=None, localns=None, include_extras=False, *, format=Format.VALUE)
- Returns resolved type hints for an object.
- Handles forward references, merges base class annotations, and replaces `None` with `types.NoneType`.
- May execute arbitrary code in annotations (security risk).
- Not supported for instances since Python 3.14.

### get_origin(tp)
- Returns the unsubscripted version of a type (e.g., `dict` from `Dict[str, int]`).
- Normalizes typing-module aliases to original classes.

### get_args(tp)
- Returns type arguments with substitutions performed (e.g., `(int, str)` from `Dict[int, str]`).
- Returns an empty tuple for unsupported objects.

### is_protocol(tp)
- Determines if a type is a `Protocol`.
- Returns `False` for generic aliases of protocols.

### is_typeddict(tp)
- Checks if a type is a `TypedDict`.

### get_protocol_members(tp)
- Returns the set of members defined in a `Protocol`.
- Raises `TypeError` for non-protocol arguments.

# Nuance Or Contradictions

- **Runtime vs Static Types**: The runtime type of an expression may differ from its statically inferred type.
- **Deprecated Decorators**: `@no_type_check_decorator` is deprecated and will be removed in Python 3.15.
- **Security Risk**: `get_type_hints()` may execute arbitrary code in annotations, posing a security risk.
- **Instance Support**: `get_type_hints()` no longer supports instances since Python 3.14.

# Candidate Wiki Hints

- **Exhaustiveness Checking with Static Typing**
- **Type Revelation for Debugging**
- **Simulating Dataclasses with @dataclass_transform**
- **Function Overloading and Introspection**
- **Finality Enforcement in Type Checkers**
- **Disabling Type Checks Locally**
- **Ensuring Method Overrides with @override**
- **Marking Runtime-Unavailable Classes**
- **Type Introspection Utilities**

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---
Chunk Context
This chunk documents the `typing` module in Python, focusing on deprecated aliases for built-in types (PEP 585), internal typing representations like `ForwardRef`, sentinel objects like `NoDefault`, and the timeline of major deprecations. It covers standard collection ABCs, asynchronous ABCs, context managers, and special forms.

Local Summary
The `typing` module historically provided aliases for built-in types (e.g., `typing.Dict`) to support generic parameterization before Python 3.9. Since PEP 585 was implemented in Python 3.9, built-ins like `dict`, `list`, and `set` now support subscripting directly, rendering the `typing` aliases redundant. These aliases are deprecated but not yet removed (except for specific cases like `ByteString`). The module also provides tools for handling forward references (`ForwardRef`) and managing type annotations safely using `TYPE_CHECKING`.

Key Claims
- Built-in types (`dict`, `list`, `set`, etc.) now support subscripting (`[]`) in Python 3.9+, making `typing.Dict`, `typing.List`, etc., deprecated aliases.
- `is_typeddict()` returns `True` only for `TypedDict` classes, not generic aliases.
- `ForwardRef` represents string forward references (e.g., `List["SomeClass"]`) and should not be instantiated by users.
- `typing.evaluate_forward_ref()` recursively evaluates nested forward references but may execute arbitrary code contained in annotations.
- `typing.TYPE_CHECKING` is a special constant assumed `True` by static type checkers to allow safe imports of expensive modules used only for type hints.
- Deprecated aliases like `typing.Text` and `typing.Hashable` have specific projected removal dates or remain with warnings pending further decisions.

Entities And Concepts
- **TypedDict**: A class used for defining typed dictionaries; generic aliases are not considered typed dicts.
- **ForwardRef**: Internal representation of string forward references in type hints.
- **NoDefault**: Sentinel object indicating a type parameter has no default value.
- **TYPE_CHECKING**: Constant for conditionally importing types only during static analysis.
- **Deprecated Aliases**: Classes like `typing.Dict`, `typing.List`, etc., now pointing to built-ins.
- **Collection ABCs**: Interfaces in `collections.abc` with corresponding deprecated aliases in `typing`.
- **Context Managers**: Abstract base classes for synchronous and asynchronous context management.

Procedures And API Details
- Use `annotationlib.get_annotations()` with `annotationlib.Format.STRING` or `annotationlib.Format.FORWARDREF` to safely inspect annotations containing undefined symbols.
- Prefer abstract collection types (e.g., `Sequence`, `Mapping`) over concrete built-ins for annotations.
- For runtime checks on buffer protocols, use `isinstance(obj, collections.abc.Buffer)` instead of `ByteString`.

Nuance Or Contradictions
- While deprecated aliases are not currently planned for removal, they may be removed in future versions with at least two releases of deprecation warnings prior to removal.
- `typing.evaluate_forward_ref()` differs from `annotationlib.ForwardRef.evaluate()` by recursively evaluating nested references.
- Some deprecated types like `ByteString` were intended as supertypes but never provided useful structural information, leading to their deprecation and planned removal in Python 3.17.

Candidate Wiki Hints
- **Generic Alias Type**: Explains the shift from `typing.Dict` to `dict[str, int]`.
- **Forward References in Type Hints**: Discusses handling of string-based forward references using `ForwardRef`.
- **TYPE_CHECKING Best Practices**: Guidelines for conditionally importing expensive modules.
- **Deprecation Timeline of typing Features**: Overview of deprecated features and their projected removal dates.

## chunk-09

---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

## Chunk Context
This chunk represents the final segment (Chunk 9 of 9) of the `typing` module documentation, specifically covering "Other special directives" and navigation elements. The content lists various alias categories for built-in types, collections, contextlib, and ABCs, followed by a deprecation timeline note. It includes standard Sphinx navigation links (Previous/Next topic), footer information regarding the Python Software Foundation license, copyright date (May 17, 2026), and build tools (Sphinx 8.2.3).

## Local Summary
The text serves as a structural outline for advanced typing features involving special directives, type aliases across standard libraries, and version deprecation schedules. It concludes the documentation section with administrative links and licensing metadata.

## Key Claims
- The `typing` module provides aliases to built-in types, types within collections, concrete types, container ABCs in `collections.abc`, asynchronous ABCs in `collections.abc`, other ABCs in `collections.abc`, and contextlib ABCs.
- A "Deprecation Timeline of Major Features" exists as a distinct section or topic within this documentation scope.
- The documentation is part of Python 3.14.5.

## Entities And Concepts
- **typing**: Support for type hints in Python.
- **collections.abc**: Abstract Base Classes (ABCs) for container and asynchronous types.
- **contextlib**: Context manager ABCs.
- **Sphinx**: The documentation generator used to create this page.
- **Python Software Foundation License Version 2**: The license governing the documentation content.

## Procedures And API Details
- No specific procedural steps or API usage examples are detailed in this chunk; it functions as a directory of available alias categories and navigation paths.
- Navigation path indicates transition from "Development Tools" to `pydoc` (Documentation generator) via "Next topic".

## Nuance Or Contradictions
- The chunk contains no explicit contradictions, but the presence of a "Deprecation Timeline" suggests evolving standards regarding type hints that may impact backward compatibility for specific features listed under aliases.

## Candidate Wiki Hints
- **Page: `typing_aliases`**: A dedicated page summarizing the various alias categories (built-in, collections, contextlib) found in this chunk.
- **Page: `typing_deprecation_schedule`**: A resource detailing the "Deprecation Timeline of Major Features" mentioned.

