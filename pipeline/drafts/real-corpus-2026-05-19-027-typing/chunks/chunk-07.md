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
