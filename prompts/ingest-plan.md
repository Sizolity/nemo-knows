You are maintaining a local Markdown wiki.
Return only the final Markdown. Do not include reasoning, analysis, prompt text, or progress logs.

Task:
Create a reviewed ingest plan draft from the raw source below. Start your
response with YAML frontmatter. Keep the whole answer short.

Rules:
- Output Markdown only.
- Begin with exactly `---`.
- Use `kind: topic`.
- Use `sources: [{{RAW_SOURCE_PATH}}]`.
- Do not propose wiki/index.md, wiki/log.md, AGENTS.md, or schema files as candidate pages.
- Candidate pages must be under wiki/sources/, wiki/concepts/, or wiki/topics/.
- Candidate pages must be immediate children of those directories: use
  `wiki/sources/<slug>.md`, `wiki/concepts/<slug>.md`, or
  `wiki/topics/<slug>.md`. Do not create nested directories.
- First infer the source kind before proposing non-source pages. Technical
  specifications and APIs usually produce concept pages; tutorials and
  comparisons usually produce topic pages; literary, historical, or narrative
  sources should usually produce topic pages for themes, motifs, or practices
  and only rarely concept pages.
- Include only these sections: Source Summary, Candidate Wiki Pages, Suggested Links, Review Checklist.

Raw source path:
{{RAW_SOURCE_PATH}}

Use this output shape:
---
kind: topic
sources: [{{RAW_SOURCE_PATH}}]
status: draft
---

# Ingest Plan

## Source Summary
<2-4 bullets>

## Candidate Wiki Pages
- wiki/sources/<slug>.md — <why>
- wiki/concepts/<slug>.md — <why>
- wiki/topics/<slug>.md — <why>

## Suggested Links
- <links explicitly present in the source, or "none">

## Review Checklist
- [ ] <review item>

Raw document:
{{RAW_SOURCE_CONTENT}}
