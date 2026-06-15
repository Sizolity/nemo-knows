---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/061-service-workers.md
confidence: medium
---

Chunk Context
Section 8 (Acknowledgements) and the subsequent "Conformance" preamble, covering lines 5309–5393. The chunk lists individuals who contributed to the specification's development via workshops, design discussions, feedback, and tooling. It then defines how conformance requirements are expressed (using RFC 2119 terminology), distinguishes normative text from examples/notes, and clarifies that imperative algorithm steps inherit their modality ("must", "should") from the surrounding context.

Local Summary
The chunk acknowledges contributors to the service worker specification and outlines the document's conventions for conformance: normative vs. informative content, interpretation of RFC 2119 key words, and how algorithmic imperatives map to modal verbs.

Key Claims
- Andrew Betts organized a workshop that advanced the work; EdgeConf sessions on "Offline" created connections enabling progress.
- Anne van Kesteren's prior work on URLs, HTTP Fetch, Promises, and DOM is foundational; Ian Hickson's Web Worker spec is also essential.
- A long list of individuals provided design guidance and discussion (e.g., Domenic Denicola, Jake Archibald, etc.).
- Jason Weber, Chris Wilson, Paul Kinlan, Ehsan Akhgari, and Daniel Austin gave well-timed feedback on requirements and standardization.
- Dimitri Glazkov's scripts and formatting tools were essential for producing the specification.
- Professional support was provided by Vivian Cromwell, Greg Simon, Alex Komoroske, Wonsuk Lee, and Seojin Kim.
- Conformance requirements use RFC 2119 key words ("MUST", "SHOULD", etc.), though they appear in mixed case for readability.
- All text is normative except sections explicitly marked non-normative, examples, and notes.
- Examples are introduced with "for example" or set apart via class="example".
- Informative notes begin with "Note" and use class="note".
- Imperative steps in algorithms inherit their modality from the introducing key word; implementations may optimize beyond the prescribed algorithms.

Entities And Concepts
- RFC 2119 (key words: MUST, MUST NOT, REQUIRED, SHALL, SHALL NOT, SHOULD, SHOULD NOT, RECOMMENDED, MAY, OPTIONAL)
- Normative text vs. informative examples/notes
- Class attributes: class="example", class="note"
- Conformance algorithms (intended for clarity, not performance; implementers encouraged to optimize)

Procedures And API Details
None in this chunk.

Nuance Or Contradictions
Readability convention: RFC 2119 key words are written in mixed case rather than all caps throughout the specification. Algorithms are defined for ease of understanding and equivalence, not for performance; implementers should optimize as needed.

Candidate Wiki Hints
- Page: Service Worker Conformance Model (normative text, examples, notes, algorithmic modality)
- Topic: Contributors to the Service Worker Specification (acknowledgements summary)
