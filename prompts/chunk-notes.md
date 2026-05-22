You are maintaining a local Markdown wiki.
Return only the final Markdown. Do not include reasoning, analysis, prompt text, or progress logs.

Task:
Create concise source-backed notes for one chunk of a longer raw document.

Rules:
- Output Markdown only.
- Begin with this exact YAML frontmatter shape:
  ---
  title: Chunk NN Notes
  kind: topic
  sources:
    - {{RAW_SOURCE_PATH}}
  confidence: medium
  ---
- Use kind: topic.
- Use sources: [{{RAW_SOURCE_PATH}}].
- Do not invent facts.
- Preserve the chunk's heading context and line range.
- If the chunk contains a marker like `[truncated at ...]`, preserve that exact
  marker and explain that the raw source stops there.
- If a chunk ends mid-sentence, mid-section, or with an explicit truncation
  marker, state that in `Nuance Or Contradictions`.
- Do not propose one wiki page per chunk. Suggest pages only when the chunk clearly supports a reusable source, concept, or topic.
- Prefer original wording. Do not copy long phrases from the source except for exact API names, command names, or quoted terms.
- Include only these sections:
  - Chunk Context
  - Local Summary
  - Key Claims
  - Entities And Concepts
  - Procedures And API Details
  - Nuance Or Contradictions
  - Candidate Wiki Hints

Raw source path:
{{RAW_SOURCE_PATH}}

Chunk:
{{CHUNK_CONTENT}}
