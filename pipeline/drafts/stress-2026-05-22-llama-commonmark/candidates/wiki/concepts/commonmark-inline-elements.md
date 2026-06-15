---
title: Commonmark Inline Elements
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Inline Elements

CommonMark defines a two-phase parsing strategy where block structures are established first, followed by the analysis of raw text contents into inline elements. This phase handles emphasis, links, and code spans within the text.

## Parsing Precedence

Inline elements follow a strict hierarchy of precedence. Code spans, links, and HTML tags interrupt or take precedence over emphasis markers. Within this hierarchy, links bind more tightly than brackets, and brackets bind more tightly than emphasis markers.

## Syntax Rules

The specification details the matching logic for emphasis delimiters, such as asterisks and underscores. These delimiters must form runs that are left-flanking to open and right-flanking to close. While intraword emphasis is generally forbidden for underscores, it is permitted for asterisks under specific conditions.

## Escaping and Entities

Special characters can be escaped using backslashes in most contexts. However, escaping is invalid within code blocks, code spans, autolinks, and raw HTML. Additionally, HTML entity references and numeric character references are treated as literal text inside code spans and code blocks, rather than being interpreted as their corresponding Unicode characters.

## Autolinks

The specification accepts strings resembling `a+b+c:d` as autolinks. These may not strictly conform to URI standards defined by RFCs. Within autolinks, backslash escapes do not function as intended.
