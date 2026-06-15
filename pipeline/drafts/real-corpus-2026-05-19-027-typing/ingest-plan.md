---
kind: topic
sources: [raw/web/corpus-2026-05-18/027-typing.md]
status: draft
---

# Ingest Plan

## Source Summary
- The source is the Python 3.14.5 `typing` module documentation, split into nine chunks covering introduction to type hints, generics, unions, special forms like `Annotated` and `Literal`, advanced features like `ParamSpec` and `TypeVarTuple`, and utility decorators for static checking.
- The content distinguishes between runtime behavior (mostly non-enforcing) and static type checking capabilities provided by third-party tools.
- Key themes include the evolution from nominal to structural subtyping, handling variadic generics via unpacking, and managing deprecated aliases versus modern built-in subscripting.

## Candidate Wiki Pages
- wiki/topics/python-type-system-overview.md — Covers runtime vs. static enforcement, type aliases, `NewType`, and basic generics.
- wiki/concepts/type-narrowing-strategies.md — Compares `TypeIs` vs. `TypeGuard` and explains intersection vs. strict narrowing.
- wiki/concepts/variadic-generics.md — Explains `TypeVarTuple`, unpacking syntax (`*Ts`), and constraints on variadic usage.
- wiki/topics/typeddict-syntax-and-modifiers.md — Details functional/class syntax, `Required`, `NotRequired`, `total`, `ReadOnly`, and introspection attributes.
- wiki/concepts/static-checking-utilities.md — Documents `assert_never`, `reveal_type`, `@override`, `@final`, `@dataclass_transform`, and `get_type_hints`.
- wiki/topics/typing-deprecation-timeline.md — Tracks deprecated aliases (`Dict` -> `dict`), removed features, and removal schedules.

## Suggested Links
- none

## Review Checklist
- [ ] Verify that all candidate pages are immediate children of the allowed directories.
- [ ] Ensure no nested directories are created in the output structure.
- [ ] Confirm that `sources` array contains exactly the provided raw path.
- [ ] Check that repeated concepts (e.g., generics, protocols) are consolidated into single broad topics rather than split by chunk number.
