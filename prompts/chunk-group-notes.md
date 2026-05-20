You are maintaining a local Markdown wiki.
Return only the final Markdown. Do not include reasoning, analysis, prompt text, or progress logs.

Task:
Create group-level notes from several adjacent chunk notes of a longer raw document.

Rules:
- Output Markdown only.
- Begin with this exact YAML frontmatter shape:
  ---
  title: Chunk Group Notes
  kind: topic
  sources:
    - {{RAW_SOURCE_PATH}}
  confidence: medium
  ---
- Do not invent facts beyond the supplied chunk notes, outline, and index.
- Preserve the covered chunk range and major heading coverage.
- Merge repeated claims and identify themes that span multiple chunks.
- Prefer synthesis over concatenation; keep enough detail for final source and ingest synthesis.
- Include only these sections:
  - Group Context
  - Cross-Chunk Summary
  - Repeated Or Central Claims
  - Important Local Details
  - Candidate Wiki Hints
  - Gaps Or Cautions

Chunk outline:
{{CHUNK_OUTLINE}}

Chunk index:
{{CHUNK_INDEX}}

Chunk notes in this group:
{{CHUNK_NOTES}}
