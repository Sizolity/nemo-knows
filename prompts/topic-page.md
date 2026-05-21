You are maintaining a local Markdown wiki.
Return only the final Markdown. Do not include reasoning, analysis, prompt text, or progress logs.

Task:
Create a topic page from the source summary below.

Rules:
- Output only Markdown.
- Include YAML frontmatter.
- Use kind: topic.
- Use Obsidian-style wikilinks only when the source material clearly supports
  the relationship.
- Only use wikilinks from the Allowed Links list. If a term is not listed or the
  source does not support the relationship, write it as plain text.
- Wikilinks are optional. Do not add a Related Concepts section just to include
  links.
- Do not invent facts.
- Prefer concise synthesis prose.
- Do not copy whole sentences from the source material. Rephrase every
  source-backed claim in original reference prose.
- A paragraph that merely restates the source summary belongs in the source
  page, not in a topic page.

Topic:
{{PAGE_TITLE}}

Target path:
{{TARGET_PATH}}

Sources:
{{SOURCE_LIST}}

Allowed Links:
{{ALLOWED_LINKS}}

Source material:
{{SOURCE_CONTENT}}
