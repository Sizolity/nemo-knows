---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

# Group Context

This group of notes synthesizes the final sections of the W3C Data Catalog Vocabulary (DCAT) specification, specifically covering **Section 18 (Accessibility Considerations)** and the concluding appendices. The content spans lines 5088 through 6468 of the source document. Key themes include accessibility mandates for non-text data, the evolution of DCAT from Version 2 Recommendation to Version 3 Candidate Recommendation, detailed mappings to Schema.org, modeling of dataset series and provenance using PROV-O, and a comprehensive bibliography of related W3C and OGC standards (VOID, DQV, WFS, WMS).

# Cross-Chunk Summary

The selected chunks collectively address the non-functional requirements and ecosystem context of DCAT.
- **Accessibility:** Chunk 13 introduces the mandate for alternative text on non-text resources, referencing UNDERSTANDING-WCAG20. Chunks 14–16 continue this section, providing RDF examples for data services and distributions while noting that accessibility is achieved through structured metadata practices.
- **Ecosystem & Standards:** The group details DCAT's relationship with Schema.org (via a recommended mapping table) and lists the broader W3C vocabulary suite (VOID, DQV, ORG, SSN) and OGC standards (WFS, WMS, WSDL) referenced for interoperability.
- **Versioning & Series:** While earlier chunks covered general versioning, Chunks 15–16 specifically highlight the introduction of `dcat:DatasetSeries` as a first-class entity and the shift toward PAV-style version chains, noting the removal of backward compatibility predicates.
- **Change History:** A significant portion of this group documents the editorial evolution of the specification since May 2021 (Issue #1358), including terminology shifts (e.g., "item" to "resource") and namespace corrections (e.g., `dct:` to `dcterms:`).

# Repeated Or Central Claims

- **Accessibility Enforcement:** DCAT profiles should enforce alternative text for non-text resources to comply with accessibility guidelines like UNDERSTANDING-WCAG20.
- **Schema.org Alignment:** Google's Dataset Search relies on both Schema.org and DCAT; a recommended mapping exists between revised DCAT elements and Schema.org 3.4 using `rdfs:subClassOf` and `owl:equivalentClass`.
- **Dataset Series Model:** The vocabulary introduces `dcat:DatasetSeries` to treat series as first-class citizens, utilizing properties like `dcat:first`, `dcat:last`, `dcat:prev`, `dcat:next`, and `dcat:inSeries`.
- **Checksum Standards:** The specification adopts SPDX checksums (`spdx:checksum`, `spdx:algorithm`, `spdx:checksumValue`) for distribution integrity verification.
- **Terminology Shift:** The term "item" was replaced with "resource" to ensure consistent terminology across the vocabulary.
- **Inverse Properties:** The property `dcat:inCatalog` is defined as the inverse of `dcat:resource`, replacing older usage patterns like `dcterms:hasPart` for catalog-resource linking.

# Important Local Details

- **Data Service Modeling:** Data services are described using `rdf:type dcat:DataService`, with classifiers via `dcterms:type` (e.g., INSPIRE codes) and the actual endpoint identified by `dcat:endpointURL`.
- **Distribution Packaging:** Distributions support modeling for compressed files (GZIP) via `dcat:compressFormat` and packaged archives (TAR, ZIP) via `dcat:packageFormat`.
- **Provenance Ontology:** Dataset provenance is described using PROV-O elements (`prov:wasGeneratedBy`, `prov:wasAttributedTo`, `prov:wasDerivedFrom`) to link datasets to generating activities and agents.
- **Publication Linking:** Datasets are linked to scholarly articles or reports via the property `dcterms:isReferencedBy`.
- **Namespace Corrections:** The document updates replace the `[DCTERMS]` prefix `dct:` with `dcterms:` throughout, ensuring consistent use of "URI" vs "IRI".
- **Reference List:** The concluding section (Chunk 16) provides a bibliography including DCAT Version 2 (Feb 2020), DCAT Version 3 (Jan 2024 Candidate Recommendation), and external standards like WFS/WMS/WSDL.

# Candidate Wiki Hints

- **Page: DCAT Accessibility Profile**
  - *Content*: Guidelines for enforcing alternative text on non-text data resources, citing UNDERSTANDING-WCAG20 compliance.
- **Page: DCAT vs. Schema.org Mapping**
  - *Content*: Reference page detailing the axiomatized mapping between DCAT 3 and Schema.org 3.4 classes and properties (e.g., `dcat:Dataset` to `sdo:Dataset`).
- **Page: Modeling Dataset Provenance**
  - *Content*: Best practices for using PROV-O alongside DCAT to describe dataset origins, activities, and agents (`prov:Activity`, `prov:Agent`).
- **Page: Linking Datasets to Publications**
  - *Content*: Usage of `dcterms:isReferencedBy` to associate datasets with scholarly articles or reports.
- **Page: Dataset Series**
  - *Content*: Explanation of the `dcat:DatasetSeries` class and its role in linking datasets within a catalog using series-specific properties.
- **Page: Data Service Endpoints**
  - *Content*: Documentation of `dcat:endpointURL`, `dcterms:conformsTo`, and `dcat:endpointDescription` for defining service endpoints.
- **Page: Distribution Packaging Formats**
  - *Content*: Explanation of how to model downloadable, compressed (GZIP), and packaged (TAR/ZIP) distributions using specific DCAT properties.
- **Page: Related Vocabularies for Data Catalogs**
  - *Content*: Curated list of complementary ontologies including VOID (Linked Datasets), DQV (Data Quality), ORG, and SSN.

# Gaps Or Cautions

- **Normative Status:** The mapping to Schema.org is explicitly described as non-normative in the text, serving as a recommendation for interoperability rather than a binding specification.
- **Legacy Catalogs:** Older catalogs (e.g., legacy CKAN) often treat datasets as a "bag of files" without distinguishing distributions from other relationships; users must carefully select properties (`dcterms:relation` vs `dcat:distribution`) to model these correctly.
- **Removed Properties:** Support for `owl:backwardCompatibleWith` and `owl:incompatibleWith` has been dropped, and older inverse properties like `dcat:isVersionOf` were removed in the May 2021 draft.
- **Terminology Consistency:** Users must adhere to the updated terminology where "resource" replaces "item" for instances of `dcat:Resource`.
- **Implicit Accessibility:** While Section 18 addresses accessibility, the text notes that it is implicitly addressed through structured metadata practices; there are no explicit procedural checks beyond following the guidelines.
