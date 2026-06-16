You are maintaining a local Markdown wiki.
Return only the final Markdown. Do not include reasoning, analysis, prompt text, or progress logs.

Task:
Create an entity page from the source summary below.

Rules:
- Output only Markdown.
- Include YAML frontmatter.
- Use kind: entity.
- Use the provided entity name as the page title, but keep it concise: no more
  than 8 English words or 18 Chinese characters.
- Entity pages are for people, organisations, products, and places. Record only
  facts supported by the source material or Target Evidence.
- Use standard Markdown relative links only as semantic cross-references when
  the source material clearly supports the relationship. Do not use them as
  index/navigation entries.
- Link only to pages in the Allowed Links list, copying the exact
  `[Label](relative/path.md)` form shown there. Each path is already computed
  relative to this page's Target path (a sibling page in the same folder is
  `name.md`; a page in another wiki folder is `../folder/name.md`). Never emit
  Obsidian-style `[[wikilinks]]`.
- If a term is not in the Allowed Links list, or the source does not support the
  relationship, write it as plain text. Never invent a link to a page that may
  not exist.
- Cross-links are optional. Do not add a Related Concepts section just to include
  links.
- Do not invent biographies, dates, affiliations, product claims, locations, or
  roles.
- Prefer concise reference prose. Structure the body as at least four short
  paragraphs, each covering a distinct source-backed facet of the entity.
- Minimise bullet lists. Use running prose paragraphs as the primary form;
  reserve lists only for genuinely list-shaped content such as enumerations
  or comparison items.
- Do not copy whole sentences from the source material. Rephrase every
  source-backed claim in original reference prose.
- Use the Target Evidence section when present. If the source summary is broad
  but target evidence contains specific passages, ground the page in that
  evidence.
- If neither the source summary nor the target evidence supports the target
  title as an entity, do not pretend it does; write a short review note that the
  candidate needs more evidence.

Entity:
{{PAGE_TITLE}}

Target path:
{{TARGET_PATH}}

Sources:
{{SOURCE_LIST}}

Allowed Links:
{{ALLOWED_LINKS}}

Source material:
{{SOURCE_CONTENT}}

Target Evidence:
{{TARGET_EVIDENCE}}
