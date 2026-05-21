You are maintaining a local Markdown wiki.
Return only the final Markdown. Do not include reasoning, analysis, prompt text, or progress logs.

Task:
Create a reviewed ingest plan draft from chunk notes for a long raw document.

Rules:
- Output Markdown only.
- Begin with exactly `---`.
- Use `kind: topic`.
- Use `sources: [{{RAW_SOURCE_PATH}}]`.
- Do not propose wiki/index.md, wiki/log.md, AGENTS.md, or schema files as candidate pages.
- Candidate pages must be immediate children of `wiki/sources/`, `wiki/concepts/`, or `wiki/topics/`.
- Valid examples: `wiki/sources/go-modules-reference.md`,
  `wiki/concepts/go-mod-directives.md`, `wiki/topics/go-workspaces.md`.
- Invalid examples: `wiki/tools/go-mod-edit-json.md`,
  `wiki/sources/go/modules.md`, `wiki/apis/foo.md`.
- Do not create nested directories.
- If a candidate is a tool/API page, place it under `wiki/concepts/` or
  `wiki/topics/`; never invent another directory.
- First infer the source kind before proposing non-source pages. Technical
  specifications and APIs usually produce concept pages; tutorials and
  comparisons usually produce topic pages; literary, historical, or narrative
  sources should usually produce topic pages for themes, motifs, or practices
  and only rarely concept pages.
- Consolidate repeated candidate hints across chunks. Prefer a small set of broadly useful pages.
- When group notes are present, treat them as the authoritative whole-document
  summary. The raw per-chunk notes section may then be empty, which is normal.
- When group notes are absent, the raw chunk notes are the only summary input.
- Include only these sections: Source Summary, Candidate Wiki Pages, Suggested Links, Review Checklist.

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
- <links explicitly present in the notes, or "none">

## Review Checklist
- [ ] <review item>

Chunk outline:
{{CHUNK_OUTLINE}}

Chunk index:
{{CHUNK_INDEX}}

Group notes:
{{CHUNK_GROUP_NOTES}}

Chunk notes:
{{CHUNK_NOTES}}
