---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

Chunk Context
- Heading path: 12. Dataset series > 12.2 Dataset series metadata, 12.3 Dataset series in existing DCAT implementations, 13. Data citation, 14. Quality information
- Lines covered: 4243–4637 (spanning dataset series inheritance rules, implementation variants, citation requirements, and quality documentation patterns).

Local Summary
This chunk details how metadata for dataset series is derived from child datasets via upstream inheritance, outlines alternative modeling approaches for series in current DCAT implementations, defines core elements for data citation, and introduces patterns for documenting quality information and conformance to standards using DQV, PROV-O, and EARL.

Key Claims
- Dataset series metadata can be split into two groups: properties describing the series itself (e.g., accrualPeriodicity) and properties inherited from child datasets via upstream inheritance.
- For inherited properties, the series value is typically the union of child values (temporal/spatial coverage), with specific rules for dates: earliest creation/publication, latest update/modification.
- Existing DCAT implementations either type the series as dcat:Dataset with children as dcat:Distribution, or both as dcat:Dataset linked by dcterms:hasPart/isPartOf; soft-typing via dcterms:type is also used.
- Data citation requires dataset identifier, creator(s), title, publisher, and publication/release date; DCAT 2 added dereferenceable identifiers and creator indication to support this.
- Quality information can be expressed using dqv:QualityAnnotation (feedback/certificates), dqv:QualityPolicy, or dqv:QualityMeasurement (metrics). Quality dimensions are not normative in DQV but may follow ISO/IEC 25012.
- Conformance to standards is modeled with dcterms:conformsTo and dcterms:Standard; best practices recommend using canonical, persistent, non-versioned IRIs from reference registries.

Entities And Concepts
- Dataset series metadata properties (dcterms:accrualPeriodicity, dcat:temporalResolution, dcat:temporal, dcat:spatial).
- Upstream inheritance mechanisms for child dataset values.
- Data citation elements (identifier, creator, title, publisher, publication date).
- Quality dimensions and DQV classes: dqv:QualityAnnotation, dqv:QualityPolicy, dqv:QualityMeasurement.
- Standard conformance modeling: dcterms:conformsTo, dcterms:Standard.
- Reference standards: EU INSPIRE Regulation (Commission Regulation (EU) No 1089/2010), OGC CRS Registry (EPSG:28992).

Procedures And API Details
- Compute series temporal coverage as the union of child start/end dates.
- Compute series spatial coverage as the union of child bounding boxes; handle multiple spatial reference systems if children differ.
- Set series dcterms:created to earliest child creation date.
- Set series dcterms:issued to earliest child publication date.
- Set series dcterms:modified to latest child publication/update date.
- Use IRIs from reference registries (W3C TR, OGC Definitions Server, ISO OBP) for standard conformance.
- Couple DCAT with DQV and PROV-O to express quality measurements and activities (e.g., ex:myQualityChecking using ex:myQualityChecker).

Nuance Or Contradictions
- DCAT does not prescribe a specific strategy for implementing upstream inheritance; mechanisms are left to implementers.
- Existing implementations may soft-type series via dcterms:type without violating DCAT, allowing coexistence with dcat:DatasetSeries during upgrades.
- DQV is non-normative and does not define a mandatory list of quality dimensions; implementers choose dimensions fitting their needs.

Candidate Wiki Hints
- Dataset Series Metadata Inheritance Rules
- Data Citation Requirements in DCAT
- Quality Information Modeling with DQV
- Documenting Conformance to Standards (dcterms:conformsTo)
- Using Canonical IRIs for Standard References
