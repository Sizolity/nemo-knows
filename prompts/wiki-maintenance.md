You are maintaining a local LLM wiki. Work only from the wiki snapshot and
maintenance tasks below. Do not read or invent raw source material. Preserve the
schema described by AGENTS.md: frontmatter, wikilinks, confidence, and concise
plain prose.

Return JSON only, with this shape:

{
  "notes": ["short note"],
  "changes": [
    {
      "path": "wiki/concepts/example.md",
      "content": "complete replacement file content"
    }
  ]
}

Rules:
- Only write files under wiki/sources/, wiki/entities/, wiki/concepts/, or
  wiki/topics/.
- Prefer repairing existing pages over creating new pages.
- For orphan pages, add natural inbound links from related existing pages when
  the relationship is supported by the wiki snapshot.
- For broken wikilinks, retarget to an existing page if the target is obvious;
  otherwise convert the unsupported wikilink to plain text.
- Do not edit wiki/index.md or wiki/log.md. The maintainer handles bookkeeping.
- If no safe semantic change is justified, return an empty changes array and a
  note explaining why.

## Maintenance Tasks

{{MAINTENANCE_TASKS}}

## Wiki Snapshot

{{WIKI_SNAPSHOT}}
