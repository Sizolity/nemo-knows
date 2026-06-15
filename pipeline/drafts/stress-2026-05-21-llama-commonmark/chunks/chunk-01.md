---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
- **Source**: CommonMark Specification (Version 0.31.2, 2024-01-28)
- **Scope**: Introduction to Markdown, rationale for the specification, and preliminary definitions regarding characters, lines, and tabs.
- **Key Sections**:
  - 1.1 What is Markdown?
  - 1.2 Why is a spec needed?
  - 1.3 About this document
  - 2.1 Characters and lines
  - 2.2 Tabs
  - 2.3 Insecure characters
  - 2.4 Backslash escapes

## Local Summary
This chunk introduces the CommonMark specification, contrasting its readability with the ambiguity of John Gruber's original Markdown description. It explains the necessity of a formal spec to resolve implementation divergences (e.g., list indentation, blank line requirements). The text defines fundamental character classes (whitespace, punctuation, control characters) and establishes rules for handling tabs in block structure contexts.

## Key Claims
- **Readability**: Markdown's primary design goal is readability; documents should be publishable as plain text without looking marked up.
- **Ambiguity of Original Syntax**: John Gruber's canonical description does not unambiguously specify syntax for edge cases (e.g., sublist indentation, blank lines before block quotes, code block indentation).
- **Implementation Divergence**: Without a spec, implementations diverge significantly, causing documents to render differently across systems (e.g., GitHub wiki vs. pandoc).
- **Tab Handling**: Tabs are not expanded to spaces generally but behave as if replaced by spaces with a 4-character tab stop in contexts defining block structure (e.g., indented code blocks).
- **Security**: The Unicode character U+0000 must be replaced with the REPLACEMENT CHARACTER (U+FFFD).

## Entities And Concepts
- **Markdown**: A plain text format for structured documents based on email/usenet conventions.
- **CommonMark**: A specification for Markdown syntax intended to be unambiguous.
- **John Gruber**: Developer of the original Markdown syntax.
- **Markdown.pl**: The original Perl script for converting Markdown to HTML; noted as buggy and insufficient as a spec.
- **Unicode Code Point**: The unit of character definition used in the spec.
- **Line Ending**: Defined as U+000A, U+000D (not followed by U+000A), or U+000D followed by U+000A.
- **Blank Line**: A line with no characters or only spaces/tabs.
- **Backslash Escape**: Mechanism to treat ASCII punctuation characters as literals, stripping their Markdown meaning.

## Procedures And API Details
- **Tab Expansion Rule**: In contexts where spaces define block structure, tabs behave as if replaced by spaces with a tab stop of 4 characters.
- **Backslash Escape Rule**: Any ASCII punctuation character may be backslash-escaped to become a literal character. Backslashes before non-punctuation characters are treated as literal backslashes.
- **Test Execution**: The spec includes a script `spec_tests.py` to run conformance tests against any Markdown program using the `spec.txt` source file.

## Nuance Or Contradictions
- **Tabs vs. Spaces**: While tabs are not expanded globally, they are treated as 4 spaces specifically for block structure definition (e.g., indented code blocks). Internal tabs within code blocks are passed through as literal tabs.
- **Blank Lines**: Most implementations do not require blank lines before block quotes or headings, though the original spec was ambiguous, leading to parsing ambiguities in hard-wrapped text.
- **Original vs. Spec**: The original Markdown description suggests certain behaviors (e.g., two lists when markers change from numbers to bullets), but the CommonMark spec and many implementations produce one list in such cases.

## Candidate Wiki Hints
- **Page**: CommonMark Specification Overview
- **Topic**: Markdown Syntax Ambiguities
- **Topic**: Handling Tabs in Markdown
- **Topic**: Backslash Escapes in Markdown
