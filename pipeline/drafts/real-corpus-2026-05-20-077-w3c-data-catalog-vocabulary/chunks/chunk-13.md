---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
# Chunk Context
**Heading**: 18. Accessibility Considerations (Lines 5088–5500)
**Scope**: Section 18 covers accessibility, acknowledgments, alignment with Schema.org, and catalog examples including provenance and publication linking.

# Local Summary
This chunk details how DCAT supports accessibility via alternative text for non-text resources, lists working group contributors, maps DCAT to Schema.org (DCAT → sdo), and provides RDF examples of loosely structured catalogs, dataset provenance using PROV-O, and linking datasets to publications.

# Key Claims
- Enforcing alternative text for non-text data improves accessibility; this aligns with WCAG 2.0 guidelines.
- DCAT backbone classes map closely to Schema.org (e.g., `dcat:Dataset` ↔ `sdo:Dataset`, `dcat:Distribution` ↔ `sdo:DataDownload`).
- A recommended mapping from revised DCAT to Schema.org v3.4 is available as an RDF file using RDFS/OWL/SKOS predicates.
- Legacy catalogs (e.g., CKAN) often treat datasets as bags of files; DCAT distinguishes distributions from other relations via `dcat:distribution` vs. `dcterms:relation`/`dcterms:hasPart`.
- Provenance can be expressed using PROV-O (`prov:wasGeneratedBy`, `prov:wasAttributedTo`, `prov:wasDerivedFrom`).
- Datasets should link to publications via `dcterms:isReferencedBy`.

# Entities And Concepts
- **DCAT vocabulary**: Data Catalog Vocabulary for describing catalogs and datasets.
- **WCAG 2.0**: Web Content Accessibility Guidelines referenced as [UNDERSTANDING-WCAG20].
- **Schema.org**: General-purpose metadata vocabulary used by search services; DCAT aligns with it for broader exposure.
- **PROV-O**: W3C Provenance Ontology for describing dataset origins and production agents.
- **Dublin Core (dcterms)**: Used in DCAT for titles, creators, identifiers, etc.
- **RDF/OWL/SKOS**: Semantic web technologies used in the DCAT-to-Schema.org mapping.
- **CKAN**: Example of a legacy catalog where datasets are loosely structured.
- **Dryad / Nature Scientific Data**: Examples of repositories and journals linked to datasets via publications.

# Procedures And API Details
- **Mapping DCAT → Schema.org**: Use an RDF file containing axiomatizations (`rdfs:subClassOf`, `owl:equivalentClass`, etc.) and annotation properties (`sdo:domainIncludes`, `sdo:rangeIncludes`).
- **Representing distributions**: Prefer `dcat:distribution` for proper representations; use `dcterms:relation`/`dcterms:hasPart` when relationship type is unknown.
- **Expressing provenance**: Link to generating projects with `prov:wasGeneratedBy`; link agents via `prov:wasAttributedTo`; link predecessor datasets via `prov:wasDerivedFrom`.
- **Linking to publications**: Use `dcterms:isReferencedBy` to connect datasets to scholarly articles or reports.
- **Example RDF graphs**: Provided in the DXWG DCAT 3 code repository (`csiro-dap-examples.ttl`, `csiro-stratchart_dcat3.ttl`).

# Nuance Or Contradictions
- Some legacy systems conflate distributions with arbitrary file relations; DCAT clarifies this distinction.
- PROV-O properties complement but do not replace Dublin Core properties (e.g., `prov:wasAttributedTo` vs. `dcterms:creator`).
- The mapping to Schema.org is non-normative and intended for interoperability with search indexes rather than strict equivalence.

# Candidate Wiki Hints
- **Page**: DCAT–Schema.org Mapping Guide (summarize the RDF mapping table and use cases).
- **Page**: Best Practices for Accessibility in DCAT Profiles (alternative text enforcement, WCAG alignment).
- **Page**: Using PROV-O with DCAT for Dataset Provenance (patterns for agents, activities, derived datasets).
- **Page**: Linking Datasets to Publications (using `dcterms:isReferencedBy` with examples from Dryad/Nature).
