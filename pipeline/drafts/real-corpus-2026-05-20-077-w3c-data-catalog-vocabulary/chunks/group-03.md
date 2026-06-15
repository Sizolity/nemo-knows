---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
# Group Context

This group of notes synthesizes the final chapters of the W3C Data Catalog Vocabulary (DCAT) specification, specifically focusing on **Section 18: Accessibility Considerations**, **Versioning**, **Dataset Series**, and the **Reference Bibliography**. The content bridges the gap between technical vocabulary definitions and practical interoperability concerns such as accessibility compliance (WCAG), spatial data service types (INSPIRE), and provenance tracking (PROV-O). It also documents the evolution of the specification from DCAT 2 (Recommendation) to DCAT 3 (Candidate Recommendation), highlighting changes in property usage, geometry handling, and checksum standards.

# Cross-Chunk Summary

The combined chunks cover the transition from core vocabulary definitions to advanced implementation considerations:

*   **Accessibility & Interoperability**: Section 18 introduces explicit guidelines for accessibility, mapping DCAT concepts to Schema.org for search engine compatibility, and enforcing alternative text. It details how to describe spatial data services (CSW, WFS, WMS) using INSPIRE classification codes.
*   **Versioning & Series Management**: The specification evolves to treat dataset series as first-class citizens (`dcat:DatasetSeries`). Versioning strategies shift from OWL-based compatibility predicates (`owl:backwardCompatibleWith`) to the [PAV] (Persistent Identifiers for Versions) approach, utilizing properties like `pav:previousVersion`, `pav:currentVersion`, and sequence indicators (`first`, `last`).
*   **Technical Refinements**: Updates include the adoption of SPDX checksums for data integrity, alignment of geometry properties with [LOCN] (Location Core Vocabulary), and corrections to namespace prefixes (e.g., `dcterms` replacing `dct:`).
*   **References & Bibliography**: The final chunk serves as a comprehensive bibliography, linking DCAT 2 and 3 documents to related W3C ontologies (DQV, ORG, VOID), OGC standards (WFS, WMS), and academic literature on Linked Data quality.

# Repeated Or Central Claims

*   **DCAT 3 Evolution**: The specification is actively evolving from the DCAT 2 Recommendation (Feb 2020) to a DCAT 3 Candidate Recommendation (Jan 2024). This evolution involves removing legacy properties (e.g., `dcat:isVersionOf`, `dcterms:hasPart` replaced by `dcat:resource`) and adding new structural elements like dataset series.
*   **Accessibility is Mandatory**: Section 18 emphasizes that accessibility considerations are not optional; they require aligning with WCAG 2.0 guidelines, particularly regarding alternative text for non-text resources and the classification of spatial data services to ensure discoverability and usability.
*   **Schema.org Alignment**: To improve interoperability with general search engines (like Google), DCAT backbone classes are mapped to Schema.org equivalents (e.g., `dcat:Dataset` ↔ `sdo:Dataset`). This mapping is non-normative but recommended for broader exposure.
*   **Provenance via PROV-O**: Dataset origins and production agents should be expressed using the W3C Provenance Ontology (PROV-O), specifically properties like `prov:wasGeneratedBy`, rather than relying solely on Dublin Core creator fields.
*   **Checksums via SPDX**: Data integrity is enforced by adding `spdx:checksum` to distributions, utilizing standardized algorithms and checksum values defined in the SPDX specification.

# Important Local Details

*   **Property Replacements**:
    *   `dcterms:hasPart` has been deprecated in favor of `dcat:resource`.
    *   The inverse property `dcat:inCatalog` was added to complement `dcat:resource`.
    *   `dcat:isVersionOf` and `dcat:next` were removed or modified in later revisions.
*   **Geometry Handling**: The property `locn:geometry` now supports geometry literals as well as classes, aligning with the Location Core Vocabulary ([LOCN]).
*   **Service Classification**: Data services (Catalog Search Web, Web Feature Service, Web Map Service) are described using `dcterms:type` values drawn from the INSPIRE Spatial Data Service Type classification (e.g., "discovery", "download", "view").
*   **Compression & Packaging**: Distributions can specify both compression formats (`dcat:compressFormat`, e.g., GZIP) and packaging formats (`dcat:packageFormat`, e.g., TAR). These are combined for archives like `.tar.gz`.
*   **Reference Examples**: The notes include references to specific RDF examples (e.g., `csiro-dap-examples.ttl`) and legacy catalog behaviors (e.g., CKAN treating datasets as bags of files) that DCAT aims to clarify.

# Candidate Wiki Hints

*   **DCAT 3 Accessibility Guide**: A page documenting Section 18, detailing best practices for enforcing alternative text and mapping spatial data services to INSPIRE types.
*   **DCAT Versioning & Series**: A guide explaining the shift to [PAV] versioning, the introduction of `dcat:DatasetSeries`, and the correct usage of sequence properties (`first`, `prev`, `next`, `last`).
*   **Schema.org Mapping Reference**: A table summarizing the non-normative mapping between DCAT classes/properties and Schema.org equivalents for search engine optimization.
*   **PROV-O Integration Patterns**: Examples of how to link datasets to generating projects, activities, and agents using PROV-O predicates within a DCAT graph.
*   **Checksum Specification**: Instructions on implementing SPDX checksums (`spdx:algorithm`, `spdx:checksumValue`) for distribution integrity verification.

# Gaps Or Cautions

*   **Content Discrepancy in Chunk 15**: While the heading indicates "Accessibility Considerations," the content of Chunk 15 primarily focuses on versioning, dataset series, and revision history. Users should not assume this section contains accessibility-specific text; that content is found in Chunks 13-14 and 16's bibliography.
*   **Non-Normative Mapping**: The DCAT-to-Schema.org mapping is explicitly non-normative. It is intended for interoperability with search indexes rather than asserting strict semantic equivalence between the two vocabularies.
*   **Legacy Property Usage**: Implementers must avoid using deprecated properties like `dcterms:hasPart` or removed predicates like `dcat:isVersionOf`, as these may lead to validation errors in DCAT 3 processors.
*   **Missing Implementation Details**: The bibliography section (Chunk 16) lists numerous external standards (OGC WFS, WSDL 2.0, DQV) but does not detail the specific syntax or RDF serialization rules for integrating them beyond the property names provided.
