You are maintaining a local Markdown wiki.
Return only the final Markdown. Do not include reasoning, analysis, prompt text, or progress logs.

Task:
Create a source summary page from the raw document below.

Rules:
- Output only Markdown.
- Include YAML frontmatter.
- Use kind: source.
- Use sources: [{{RAW_SOURCE_PATH}}].
- Do not invent facts.
- Do not add URLs that are not explicitly present in the raw document.
- Condense and synthesize the source; do not quote or concatenate long
  passages from it.
- Keep the summary concise.
- Include sections:
  - What It Is
  - Summary
  - Key Claims
  - Suggested Links

Suggested Links rules:
- Include only links explicitly present in the raw document.
- If a relevant tool or concept is mentioned without a URL, mention it as plain text, not as a Markdown link.

Raw source path:
{{RAW_SOURCE_PATH}}

Raw document:
{{RAW_SOURCE_CONTENT}}