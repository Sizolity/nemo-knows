You are maintaining a local Markdown wiki.
Return only the final Markdown. Do not include reasoning, analysis, prompt text, or progress logs.

Task:
Create a source summary page from chunk notes for a long raw document.

Rules:
- Output only Markdown.
- Include YAML frontmatter.
- Use kind: source.
- Use sources: [{{RAW_SOURCE_PATH}}].
- Do not invent facts beyond the chunk notes.
- Consolidate repeated claims across chunks instead of listing every section.
- Mention important coverage from early, middle, and late sections when present.
- Preserve source-boundary facts. If the chunk or group notes mention a
  `[truncated at ...]` marker, incomplete retrieval, missing ending, mid-sentence
  ending, or any other truncation boundary, the `What It Is` section must state
  that boundary clearly.
- When a truncation or incomplete-source boundary is present, do not describe the
  source as complete, entire, full text, all chapters, final chapters, or without
  abridgment. Use wording such as "retrieved text through..." or "partial corpus
  item ending at..." instead.
- Do not infer source completeness from a title, table of contents, or upstream
  publication source. Only claim completeness if the notes explicitly establish
  that the raw file includes the full work.
- Synthesize the chunk and group notes in fresh prose. Do not concatenate or
  copy whole sentences from the notes unless the exact phrase is a named API,
  title, command, or quotation that must remain exact.
- When group notes are present, treat them as the authoritative whole-document
  summary. The raw per-chunk notes section may then be empty, which is normal.
- When group notes are absent, the raw chunk notes are the only summary input.
- Keep the summary concise.
- Include sections:
  - What It Is
  - Summary
  - Key Claims
  - Suggested Links

Suggested Links rules:
- Include only links explicitly present in the chunk notes or outline.
- If no source URL is available in the notes, write "none".

Chunk outline:
{{CHUNK_OUTLINE}}

Chunk index:
{{CHUNK_INDEX}}

Group notes:
{{CHUNK_GROUP_NOTES}}

Chunk notes:
{{CHUNK_NOTES}}
