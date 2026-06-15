## group-01

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

# Group Context

**Source:** `raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md`
**Coverage:** Lines 1–6468 (Chunks 01 through 16)
**Primary Topic:** W3C Data Catalog Vocabulary (DCAT) Version 3 Specification

This group of notes synthesizes the full scope of the DCAT 3 specification, covering its evolution from DCAT 1 and 2, namespace definitions, core classes (Catalog, Dataset, Distribution, DataService), vocabulary for resource metadata (types, themes, relations), versioning strategies, dataset series, quality information, and considerations for security, privacy, and accessibility. The document serves as a W3C Recommendation published in August 2024, designed to enable interoperable data cataloging across federated systems using RDF.

---

# Cross-Chunk Summary

The document is structured into logical sections that build upon one another:

1.  **Introduction & Motivation (Chunks 01):** Establishes DCAT 3 as a W3C Recommendation, highlighting improvements in versioning, dataset series, and checksum support while maintaining backward compatibility with DCAT 2.
2.  **Namespaces & Scope (Chunks 02-03):** Defines normative and non-normative namespaces, outlines the core classes (`Catalog`, `Dataset`, `Distribution`, `DataService`), and introduces the concept of catalog records versus dataset instances. It also covers thematic classification using SKOS.
3.  **Vocabulary Specification - Resource Attributes (Chunks 04-05):** Details properties for describing resources, including language, publisher, identifiers, themes, types, relations, and versioning chains (`previousVersion`, `currentVersion`). It introduces the distinction between ordered series (`first`, `last`, `prev`) and version histories.
4.  **Vocabulary Specification - Distributions & Services (Chunk 06):** Focuses on the `Distribution` class, distinguishing between access URLs (services/landing pages) and download URLs. It defines `DataService` for API endpoints and introduces spatial/temporal resolution properties.
5.  **Auxiliary Classes & Relations (Chunks 07-08):** Covers auxiliary classes like `Relationship`, `Role`, `Period of Time`, `Location`, and `Checksum`. It also addresses dereferenceable identifiers, license/rights statements, and time/space modeling.
6.  **Versioning & Series (Chunks 09-10):** Elaborates on version life-cycles, conformance to standards, and dataset series metadata.
7.  **Quality & Conformance (Chunk 11):** Discusses quality information, conformance testing results, and qualified relations between datasets and agents/resources.
8.  **Profiles & Considerations (Chunks 12-16):** Introduces DCAT Profiles for domain-specific extensions and concludes with sections on Security/Privacy and Accessibility considerations.

---

# Repeated Or Central Claims

*   **Interoperability & Federation:** A core goal of DCAT is to enable decentralized publishing and federated search across multiple organizations using uniform query mechanisms (SPARQL, HTML-RDFa).
*   **Backward Compatibility:** DCAT 3 supersedes DCAT 2 but explicitly preserves definitions from previous versions, allowing existing implementations to remain compliant without mandatory modification.
*   **Resource vs. Distribution:** A fundamental distinction is maintained between the abstract `Dataset` (the collection of data) and its `Distribution` (specific serializations or access points).
*   **External Vocabulary Integration:** DCAT relies heavily on external vocabularies for specific semantics, particularly Dublin Core (`dcterms:`), FOAF (`foaf:`), SKOS (`skos:`), SPDX (for checksums), and Schema.org (`sdo:`). Changes to these external definitions do not affect DCAT conformance.
*   **Non-Prescriptive Deployment:** The specification does not mandate a specific deployment method; implementations may choose among various RDF serialization formats (Turtle, RDF/XML, JSON-LD, N3) and exposure mechanisms.
*   **Blank Node Avoidance:** While RDF allows blank nodes, the specification strongly discourages their use for DCAT instances to ensure global identifiers (IRIs) are used for collaborative annotation and interoperability.

---

# Important Local Details

*   **Namespace URI:** `http://www.w3.org/ns/dcat#` with a suggested prefix of `dcat`.
*   **Classes Hierarchy:**
    *   `Cataloged Resource`: Super-class for all catalog items.
    *   `Dataset`: Represents collections of data (numbers, text, pixels, imagery).
    *   `Distribution`: Represents accessible forms (files, APIs).
    *   `DataService`: Represents service endpoints (APIs).
    *   `Catalog`: A collection of metadata records; in DCAT 3, it is a subclass of `Dataset` allowing catalogs to contain other catalogs.
*   **Versioning Properties:**
    *   `dcat:previousVersion`, `dcat:hasCurrentVersion`, `dcat:replaces`: Used for version chains (snapshots).
    *   `dcat:first`, `dcat:last`, `dcat:prev`: Used for ordered collections (series), distinct from version history.
*   **Access Modeling:**
    *   `dcat:landingPage`: For datasets accessible only via a web page.
    *   `dcat:downloadURL`: For direct file downloads.
    *   `dcat:accessService` / `dcat:accessURL`: Link to services or service endpoints.
*   **Quality & Conformance:** The spec allows linking to quality measurement activities (`ex:myQualityChecker`) and documenting conformance results (manual, automatic, conformant, non-conformant) using SKOS concept schemes and EARL vocabulary patterns.
*   **Checksums:** DCAT 3 integrates `spdx:checksum` to represent file digests, linking the checksum to a specific distribution.

---

# Candidate Wiki Hints

*   **DCAT 3 Overview**: High-level summary of purpose, scope, and evolution from DCAT 1/2.
*   **Namespace & Prefix Guide**: Table of normative vs. non-normative prefixes (dcat, dcterms, foaf, skos, spdx, etc.).
*   **Core Classes Reference**: Definitions and relationships for `Catalog`, `Dataset`, `Distribution`, `DataService`, `CatalogRecord`.
*   **Versioning Strategies**: Guide on distinguishing between version chains (`previousVersion`) and series ordering (`first`, `last`).
*   **Access Patterns**: How to model different access methods (landing page, download, service) using distributions.
*   **Thematic Classification**: Using `dcat:theme` with `skos:ConceptScheme` for categorization.
*   **Quality & Conformance Documentation**: How to represent quality metrics and conformance test results in DCAT graphs.
*   **Profiles Guide**: Explanation of how to create DCAT profiles (application profiles) with constraints or sub-classes.

---

# Gaps Or Cautions

*   **Blank Nodes:** Be careful not to use blank nodes for instances of main DCAT classes; always use IRIs.
*   **Identifier vs. IRI:** While an identifier might be part of the resource's IRI, it should still be explicitly represented as an `rdfs:Literal` (e.g., using `dcterms:identifier`) to avoid ambiguity.
*   **Theme Inference:** In DCAT 3, `dcat:theme` is treated as an OWL object property. Avoid inferring that the target of a theme link is automatically a `skos:Concept`; use explicit typing if needed.
*   **License Placement:** Legal license information should generally be placed on the `Distribution` level rather than the `Dataset` level to avoid conflicts when different distributions have different licenses.
*   **Resolution Datatypes:** When using `dcat:spatialResolutionInMeters`, ensure values are typed as `xsd:decimal`. While integers are semantically valid, some validators (SHACL) may require explicit datatype declarations.
*   **External Definitions:** Do not assume definitions of external terms (e.g., from DCTERMS or FOAF) are static; refer to their original specifications for authoritative updates.

## group-02

---
title: Chunk Group Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

# Group Context

This group synthesizes notes from Chunks 7 through 16, covering the complete **Vocabulary Specification** (Section 6) and subsequent sections on License, Rights, Time/Space, Versioning, Dataset Series, Quality Information, Conformance Testing, Qualified Relations, Security, and Accessibility. The content details how to model complex relationships between datasets and agents, manage version chains and hierarchies, express spatial/temporal coverage, document rights statements, handle quality metrics, and address security considerations.

# Cross-Chunk Summary

The specification defines a rich set of RDF classes and properties for the DCAT 2/3 vocabulary:

*   **Relationships & Roles:** Introduced `dcat:Relationship` and `dcat:Role` to model associations between resources and agents without property explosion, using qualified relations (`dcat:qualifiedRelation`, `prov:qualifiedAttribution`).
*   **Temporal & Spatial Metadata:** Defines properties for time intervals (`dcterms:issued`, `dcat:startDate`, `time:hasBeginning`) and spatial coverage (`locn:geometry`, `dcat:bbox`, `dcat:centroid`), supporting both WGS84 and non-WGS84 coordinate systems via WKT literals.
*   **Versioning & Series:** Distinguishes between version metadata (using `dcat:previousVersion`, `dcat:hasCurrentVersion`) and resource life-cycle status (`adms:status`). Introduces `dcat:DatasetSeries` for grouping related datasets, with rules for aggregating series-level temporal/spatial coverage.
*   **Rights & Licenses:** Differentiates between licenses (`dcterms:license`), access rights (`dcterms:accessRights`), and other rights like copyright (`dcterms:rights`). Supports ODRL policies via `odrl:hasPolicy`.
*   **Quality & Conformance:** Models quality metrics using VOCAB-DQV (`dqv:QualityMeasurement`) and conformance testing using PROV-O and EARL (`earl:Assertion`, `prov:Plan`), allowing for degrees of conformance (full, partial, non-conformant).
*   **Security & Integrity:** Recommends using `spdx:Checksum` for data integrity verification, emphasizing the need to separate checksum delivery channels from the data to prevent tampering.

# Repeated Or Central Claims

*   **Inverse Properties:** DCAT intentionally omits inverse properties in the core vocabulary but allows them as additions (e.g., `dcat:isDistributionOf`), never as replacements for normative properties.
*   **Identifier Management:** HTTP IRIs are preferred for dereferenceable identifiers. Non-HTTP identifiers should use `adms:identifier` with `skos:notation`. External RDF descriptions (e.g., DOIs) can be linked via `owl:sameAs`.
*   **Spatial Encoding:** Geometry is encoded as WKT literals (`geosparql:wktLiteral`) within `rdfs:Literal` ranges. Spatial properties default to CRS84 unless a specific CRS (e.g., EPSG:28992) is provided.
*   **Versioning vs. Life-Cycle:** A resource's life-cycle status (e.g., deprecated, withdrawn) is distinct from its versioning history. New versions do not automatically imply a status change.
*   **Rights Distinction:** Three scenarios exist: licenses, access rights only, and other rights statements (e.g., copyright). Legal advice is recommended for selecting license conditions.

# Important Local Details

*   **Vocabulary Classes:** `dcat:Relationship`, `dcat:Role`, `dcterms:PeriodOfTime`, `spdx:Checksum`, `dcat:DatasetSeries`.
*   **Key Properties:**
    *   *Relationships:* `dcat:relation`, `dcat:hadRole`.
    *   *Temporal:* `dcat:startDate`, `dcat:endDate`, `time:hasBeginning`, `time:hasEnd`.
    *   *Spatial:* `locn:geometry`, `dcat:bbox`, `dcat:centroid`, `dcat:spatialResolutionInMeters`.
    *   *Versioning:* `dcterms:replaces`, `dcat:isVersionOf`, `adms:versionNotes`.
    *   *Series:* `dcat:first`, `dcat:last`, `dcat:prev`, `dcat:next`.
    *   *Quality:* `dqv:QualityMeasurement`, `dqv:Metric`.
    *   *Conformance:* `earl:Assertion`, `prov:Activity`, `prov:Plan`.
*   **Vocabularies & Standards:** ISO-19115 CI_RoleCode, DataCite relators, MARC relators, ODRL, SPDX, VOCAB-DQV, EARL, INSPIRE DoC.
*   **Checksum Requirements:** Checksums must be provided via a route separate from the data they sum to ensure integrity and authenticity.

# Gaps Or Cautions

*   **Legal Disclaimer:** Selecting access conditions (licenses) requires legal advice; DCAT distinguishes scenarios but does not define specific license content.
*   **Implementation Flexibility:** Existing DCAT implementations may use alternative patterns (e.g., series as `dcat:Dataset` with distributions) which are not formally incompatible but differ from the recommended `dcat:DatasetSeries` pattern.
*   **OWD Limitations:** PROV-O is activity-centric and lacks direct Entity-Entity relations; DCAT elements fill this gap for cataloging purposes.
*   **Test Mode Ambiguity:** It is unclear in some examples whether a conformance test was manual or automatic, suggesting metadata should be explicit about test execution modes.
*   **Security Scope:** Detailed web security mechanisms (authentication, specific encryption standards) are out of scope; implementers must rely on external policies and HTTPS origins for trust.

## group-03

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

