---
title: Commonmark Html Blocks
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/053-commonmark-specification.md
confidence: medium
---

# Commonmark Html Blocks

In the CommonMark specification, HTML blocks are treated as raw text segments that can interrupt the flow of paragraphs. These elements are recognized based on specific start and end conditions, allowing them to coexist with the surrounding Markdown structure.

The parser identifies these blocks by looking for opening tags that match the syntax of HTML. Once a valid start tag is detected, the content within the block is preserved verbatim, including internal spaces. This raw text preservation ensures that the original formatting of the HTML is maintained without interpretation by the Markdown engine.

## Parsing Rules

The handling of HTML blocks follows a rigorous algorithmic definition. The parser distinguishes between block-level and inline-level syntax, ensuring that HTML blocks are processed during the initial phase of constructing the block structure.

- **Start Conditions**: A block begins when a line contains a valid HTML start tag.
- **End Conditions**: The block concludes when the parser encounters a line that does not match the continuation rules for the specific HTML block type.
- **Indentation**: Indentation rules apply relative to the start of the block. While 4 spaces typically trigger a code block in other contexts, specific indentation logic governs HTML blocks to determine their boundaries.
- **Raw Text Preservation**: Unlike code spans or inline elements, the text inside an HTML block is not parsed for Markdown syntax like links or emphasis.

## Interaction with Other Elements

HTML blocks interact with the broader document structure in specific ways:

- **Paragraph Interruption**: An HTML block can terminate a paragraph, effectively breaking the flow of text.
- **Container Context**: These blocks exist within the tree of blocks, which includes document roots, block quotes, and lists.
- **Escaping Limitations**: While backslashes escape punctuation in most contexts, this mechanism is invalid within raw HTML blocks, code blocks, and code spans.

## Related Concepts

- [[commonmark-parsing-phases]]
- [[commonmark-indentation-rules]]
- [[commonmark-code-blocks]]
