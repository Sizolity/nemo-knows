---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

## Chunk Context
- **Heading**: `baz`
- **Line Range**: 1979–2982
- **Source Section**: Covers code block info strings, HTML blocks (types 1–7), and the beginning of link reference definitions.

## Local Summary
This chunk details the handling of info strings in code blocks, distinguishing between backtick and tilde fences regarding allowed characters. It then extensively defines HTML blocks, categorizing them into seven types based on start/end tags (e.g., `<pre>`, comments, CDATA). It explains how HTML blocks interrupt paragraphs, handle indentation, and interact with Markdown syntax. Finally, it introduces link reference definitions, outlining their structure, precedence rules, and constraints regarding indentation and placement.

## Key Claims
- Info strings for backtick code blocks cannot contain backticks; tilde code blocks can.
- HTML blocks are treated as raw HTML and are not escaped in output.
- There are seven specific kinds of HTML blocks defined by start/end conditions.
- HTML blocks of types 1–6 can interrupt a paragraph; type 7 cannot.
- Link reference definitions consist of a label, colon, destination, and optional title.
- Link reference definitions do not correspond to structural elements but define labels for later use.
- Matching of link labels is case-insensitive.

## Entities And Concepts
- **Info String**: Optional text after a code fence indicating language or metadata.
- **HTML Block**: A group of lines treated as raw HTML.
- **Link Reference Definition**: A declaration defining a label for reference links.
- **Block-level HTML Elements**: Elements like `<div>`, `<table>`, `<pre>`.
- **Inline HTML**: Tags not on their own line (e.g., `<del>text</del>`).

## Procedures And API Details
- **HTML Block Start Conditions**:
  1. `<pre`, `<script`, `<style`, or `<textarea` followed by space/tab/>.
  2. `<!--` (comment).
  3. `<?` (processing instruction).
  4. `<!` followed by an ASCII letter (declaration).
  5. `<![CDATA[` (CDATA).
  6. Block-level tags (e.g., `<div>`, `<table>`) followed by space/tab/>.
  7. Any complete open/closing tag (except types 1–4) followed by space/tab/end of line.
- **HTML Block End Conditions**:
  1. Matching end tag (`</pre>`, etc.).
  2. `-->` (comment).
  3. `?>` (processing instruction).
  4. `>` (declaration).
  5. `]]>` (CDATA).
  6. Blank line.
  7. Blank line.
- **Link Reference Definition Syntax**: `[label]: destination "title"` (with optional spaces/tabs/line breaks).

## Nuance Or Contradictions
- **Gruber’s Original Markdown vs. CommonMark**: Gruber’s specification required blank lines before HTML blocks and matching end tags, whereas CommonMark allows HTML blocks to interrupt paragraphs and does not require matching end tags for types 1–6.
- **Blank Lines in HTML Blocks**: CommonMark disallows blank lines inside HTML blocks (except types 1–5) to avoid expensive parsing of balanced tags, whereas Gruber’s rule allowed them.
- **Indentation Sensitivity**: HTML blocks can be preceded by up to three spaces of indentation; four spaces trigger a code block instead.

## Candidate Wiki Hints
- **HTML Blocks in Markdown**: A guide to understanding how HTML blocks are parsed, including the seven types and their interaction with Markdown syntax.
- **Link Reference Definitions**: Rules for defining and using reference links in CommonMark.
- **Code Block Info Strings**: Best practices for specifying language metadata in code blocks.
