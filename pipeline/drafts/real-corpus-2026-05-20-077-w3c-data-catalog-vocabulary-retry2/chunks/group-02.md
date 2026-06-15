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
