---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
This chunk covers Section 18, "Accessibility Considerations," followed by Appendix A (Acknowledgments), B (Alignment with Schema.org), and the beginning of C (Examples). It details accessibility mandates for non-text data, lists contributors to the DCAT vocabulary specification, maps DCAT concepts to Schema.org types, and provides RDF examples for loosely structured catalogs, dataset provenance, and linking datasets to publications.

Local Summary
The document section emphasizes enforcing alternative text for non-text resources within DCAT profiles to comply with accessibility guidelines like UNDERSTANDING-WCAG20. It acknowledges the working group members and external reviewers. Section B outlines the relationship between DCAT and Schema.org, noting that Google's Dataset Search relies on both. A detailed mapping table aligns DCAT elements (e.g., dcat:Dataset) with Schema.org types (e.g., sdo:Dataset). Section C begins with examples for loosely structured catalogs using dcterms:relation/dcterms:hasPart versus dcat:distribution, dataset provenance using the PROV ontology, and linking datasets to publications via dcterms:isReferencedBy.

Key Claims
- Data catalogs often contain non-text data; DCAT profiles should enforce alternative text for such resources to improve accessibility.
- Providing text alternatives for non-text content (large print, braille, speech, etc.) complies with UNDERSTANDING-WCAG20 guidelines.
- Schema.org types and properties are based on original DCAT work; Google's Dataset Search index relies on structured descriptions using both Schema.org and DCAT.
- A recommended mapping from revised DCAT to Schema.org version 3.4 uses predicates like rdfs:subClassOf, owl:equivalentClass, and sdo:domainIncludes.
- In loosely structured catalogs (e.g., legacy CKAN), datasets are often treated as a "bag of files" without distinguishing between distributions and other relationships.
- If the relationship nature is unknown, dcterms:relation or dcterms:hasPart should be used; if known and representing a proper representation, dcat:distribution should be used.
- Dataset provenance can be described using W3C Provenance Ontology elements like prov:wasGeneratedBy, prov:wasAttributedTo, and prov:wasDerivedFrom.
- Datasets are linked to publications (scholarly articles, reports) via the property dcterms:isReferencedBy.

Entities And Concepts
- **DCAT**: Vocabulary for describing data catalogs.
- **Schema.org**: Web vocabulary used by general-purpose search services; DCAT aligns with it for broader exposure.
- **UNDERSTANDING-WCAG20**: Accessibility guidelines referenced for text alternatives.
- **Prov-O (Provenance Ontology)**: Used to describe dataset provenance and business context.
- **Dublin Core (dcterms)**: Properties used in DCAT for metadata (e.g., title, description, creator).
- **RDF/Vocabularies**: Includes FOAF (foaf:Document), Geosciml ontology, OWL, RDF Schema.

Procedures And API Details
- **Mapping to Schema.org**: Use an available RDF file mapping DCAT revised elements to Schema.org 3.4 using specific predicates (rdfs:subClassOf, owl:equivalentClass, etc.).
- **Loosely Structured Catalogs**:
  - Use `dcterms:relation` or `dcterms:hasPart` when the nature of relationships between a dataset and component resources is unknown.
  - Use `dcat:distribution` when related resources are proper representations of the dataset.
- **Provenance Linking**:
  - Use `prov:wasGeneratedBy` to link a dataset to the project that generated it (formalized as a `prov:Activity`).
  - Use `prov:wasAttributedTo` for agents (sponsors, managers) not covered by standard Dublin Core creator/contributor/publisher roles.
  - Use `prov:wasDerivedFrom` for links to predecessor datasets.
- **Publication Linking**: Use `dcterms:isReferencedBy` to link a dataset to associated publications.

Nuance Or Contradictions
- The distinction between an abstract `Dataset` and a concrete `DataDownload` in DCAT maps to `dcat:Dataset / dcat:Distribution`, whereas Schema.org distinguishes `Dataset` from `DataDownload`.
- Legacy catalogs often lack the distinction between distributions (representations) and other relationships (documentation, schemas), requiring careful selection of DCAT properties to model this correctly.
- The mapping to Schema.org is described as non-normative in the text but provides a recommended alignment for interoperability with search indexes.

Candidate Wiki Hints
- **DCAT Accessibility Profile**: Guidelines for enforcing alternative text on non-text data resources within DCAT implementations.
- **DCAT vs. Schema.org Mapping**: A reference page detailing the axiomatized mapping between DCAT 3 and Schema.org 3.4 classes and properties.
- **Modeling Dataset Provenance**: Best practices for using PROV-O alongside DCAT to describe dataset origins, activities, and agents.
- **Linking Datasets to Publications**: Usage of `dcterms:isReferencedBy` to associate datasets with scholarly articles or reports.
