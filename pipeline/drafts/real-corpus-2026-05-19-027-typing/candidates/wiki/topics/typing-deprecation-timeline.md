---
title: Typing Deprecation Timeline
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/027-typing.md
confidence: medium
---

# Typing Deprecation Timeline

The Python `typing` module has undergone several deprecations to align with evolving language standards and modern syntax. These changes generally move the ecosystem away from legacy aliases and functional constructs toward direct subscripting, standardized statements, and cleaner definitions.

## Deprecated Built-in Type Aliases
Starting with Python 3.9, explicit aliases for built-in collection types were deprecated in favor of direct subscripting. Developers are encouraged to replace `typing.Dict`, `typing.List`, and similar legacy names with the built-in `dict`, `list`, and others used directly as generics (e.g., `dict[str, int]`).

## TypedDict Syntax Changes
The functional syntax for defining `TypedDict` classes was deprecated in Python 3.13. This shift encourages the use of class-based definitions to better support generic parameters inherited from `Generic[T]`, which became available in Python 3.11 and earlier versions as well. The functional approach, specifically the removal of features planned for version 3.15, is being phased out to streamline type definitions.

## Special Decorator Removals
Certain decorators intended for specific legacy behaviors are scheduled for removal. The `@no_type_check_decorator` directive is set to be removed in Python 3.15, reflecting the module's move toward stricter and more consistent static analysis standards without the need for opt-out mechanisms.
