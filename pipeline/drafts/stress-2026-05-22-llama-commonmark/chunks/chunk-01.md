---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Chunk Context

**Heading path:** Document
**Line range:** 1–516
**Coverage:**
- Document overview and metadata
- Introduction to Markdown and the need for a specification
- Preliminaries: characters, lines, tabs, and insecure characters
- Backslash escapes and entity references

# Local Summary

This chunk introduces the CommonMark Specification (version 0.31.2) by John MacFarlane. It explains Markdown’s design goal of readability, contrasts Markdown with AsciiDoc, and justifies the need for an unambiguous spec due to historical ambiguities in John Gruber’s original description. The text covers character definitions, line endings, tab handling, and backslash escape rules.

# Key Claims

- Markdown is a plain text format for structured documents, prioritizing readability so that source text is publishable as-is.
- The original Markdown syntax description by John Gruber is ambiguous in several cases (e.g., list indentation, blank line requirements, code block rules).
- Without an unambiguous spec, implementations diverge, leading to inconsistent rendering across platforms.
- Tabs are not expanded to spaces by default but behave as four spaces in contexts defining block structure.
- Backslashes before ASCII punctuation characters escape their Markdown meaning; backslashes before other characters are literal.

# Entities And Concepts

- **CommonMark**: A specification for a Markdown-compatible markup language.
- **Markdown**: A lightweight markup language developed by John Gruber and Aaron Swartz.
- **AsciiDoc**: A markup language used for comparison to illustrate Markdown’s readability.
- **Backslash escapes**: Mechanism to treat punctuation characters literally.
- **Unicode code point**: The unit of character used in the spec.
- **Line ending**: LF, CR, or CRLF.
- **Blank line**: A line with no characters or only spaces/tabs.
- **Insecure characters**: Specifically U+0000, replaced with U+FFFD.

# Procedures And API Details

- **Tab handling**: Tabs are treated as four spaces in block structure contexts (e.g., indented code blocks, list item continuation).
- **Escape sequence**: `\` + ASCII punctuation character → literal character.
- **Test runner**: `python test/spec_tests.py --spec spec.txt --program PROGRAM` can run conformance tests.
- **Spec generation**: `tools/makespec.py` converts `spec.txt` (Markdown with test extensions) to HTML or CommonMark.

# Nuance Or Contradictions

- **Tabs vs. spaces**: Tabs are not expanded globally but act as four spaces in structural contexts. Internal tabs in code blocks remain literal.
- **Blank lines**: Some implementations require blank lines before block quotes or headings; the spec aims to clarify this.
- **List markers**: Ambiguities exist regarding indentation of sublists, right-aligned markers, and marker changes (numbers to bullets).
- **HTML rendering**: The spec uses HTML in examples but does not mandate all HTML features (e.g., percent-encoding of non-ASCII URLs).

# Candidate Wiki Hints

- **Page: CommonMark Specification** – Overview of the spec, its history, and design goals.
- **Page: Markdown vs. AsciiDoc** – Comparison of readability and syntax.
- **Page: Backslash Escapes** – Rules for escaping punctuation and literal backslashes.
- **Page: Tab Handling in Markdown** – How tabs interact with block structure.
- **Page: CommonMark Conformance Tests** – Using `spec_tests.py` for validation.
