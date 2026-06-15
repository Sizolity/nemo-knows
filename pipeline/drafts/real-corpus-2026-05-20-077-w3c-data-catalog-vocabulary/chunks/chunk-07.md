---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

### Chunk Context
This chunk details the vocabulary specifications for **DCAT 2** and **DCAT 3**, covering relationships, roles, temporal intervals, spatial locations, and checksums. It further defines the use of inverse properties, dereferenceable identifiers (including HTTP IRIs, proxies, and legacy IDs), and strategies for indicating common identifier types using RDF datatypes or OWL datatypes when not HTTP-dereferenceable.

### Local Summary
The specification introduces specific classes (`Relationship`, `Role`, `PeriodOfTime`, `Location`, `Checksum`) to characterize associations between datasets that are not covered by standard DCTERMS or PROV-O properties. It defines temporal and spatial coverage using ISO 8601 literals or GeoSPARQL geometries. The section on identifiers emphasizes the use of persistent HTTP IRIs (`dcterms:identifier`) while allowing for proxy dereferenceable IRIs, legacy IDs, and locally minted identifiers via `adms:identifier`. It concludes with guidelines for encoding non-HTTP identifiers using custom datatypes.

### Key Claims
- **Relationship Class**: Added in DCAT 2 to attach additional information to relationships between DCAT resources where standard properties (DCTERMS/PROV-O) are insufficient.
- **Role Class**: Added in DCAT 2 to specify the function of a resource or agent with respect to another, recommending controlled vocabularies like ISO-19115.
- **Temporal Properties**: `dcat:startDate` and `dcat:endDate` use ISO 8601 strings; `time:hasBeginning`/`time:hasEnd` use `time:Instant` for non-Gregorian or numeric time positions.
- **Spatial Properties**: `locn:geometry` is preferred for extensive geometries, while `dcat:bbox` and `dcat:centroid` are used for bounding boxes and centers respectively.
- **Checksum Class**: Added in DCAT 3; uses `spdx:algorithm` and `spdx:checksumValue` to verify file integrity (e.g., MD5, SHA-256).
- **Inverse Properties**: Supported only in addition to standard properties, never as replacements, to ensure interoperability without OWL reasoning.
- **Identifiers**: Prefer persistent HTTP IRIs; use `owl:sameAs` for dereferenceable IDs returning RDF/OWL descriptions; distinguish primary vs. alternative identifiers based on application context (e.g., DCAT-AP).

### Entities And Concepts
- **dcat:Relationship**: An association class for relationships between datasets/resources.
- **dcat:hadRole**: Specifies the function of an entity/agent with respect to another.
- **dcat:Role**: A subclass of `skos:Concept` representing a role in attribution or relationships.
- **dcterms:PeriodOfTime**: An interval of time defined by start and end.
- **locn:geometry**: Associates a spatial thing with a geometry (WKT, GeoSPARQL).
- **spdx:Checksum**: Represents checksum algorithms and values for file integrity.
- **adms:identifier**: Used for locally minted or external identifiers (DOI, ORCID).
- **owl:sameAs**: Links resources that share the same identity when dereferenced.

### Procedures And API Details
- **Temporal Encoding**: Use `xsd:gYear`, `xsd:date`, or `xsd:dateTime` for `dcat:startDate`/`endDate`; use `time:Instant` for `time:hasBeginning`/`time:hasEnd`.
- **Spatial Encoding**: Encode geometries as literals (WKT) or classes (`geosparql:Geometry`). WKT supports non-WGS84 coordinate systems.
- **Identifier Types**:
  - HTTP IRIs: Use `dcterms:identifier` with `xsd:anyURI`.
  - Non-HTTP: Use RDF datatypes (`ex:type`) or OWL datatypes.
  - DOI Encoding: Use full URL form (e.g., `https://doi.org/10.xxxx/xxxxx/`).
- **Agency Representation**: Use `adms:schemaAgency` for the authority defining the scheme; use `dcterms:creator` if the agency has an IRI.

### Nuance Or Contradictions
- **Inverse Properties**: The spec intentionally excludes inverses in the core vocabulary to support systems not using OWL reasoning, but allows them as additions (e.g., `dcat:prev`/`dcat:next`).
- **Identifier Philosophy**: DOIs abstract sub-spaces from the registering organization to ensure stability if responsibility changes; registrants should not represent the assigning authority in the identifier value itself.
- **Primary vs. Alternative IDs**: Distinguishing between primary and legacy identifiers is application-specific and better addressed in DCAT profiles (e.g., DCAT-AP) rather than a general mandate.

### Candidate Wiki Hints
- **Page: DCAT Relationship Class** - Covers `dcat:Relationship`, `hadRole`, and usage examples from qualified relations.
- **Page: DCAT Temporal Properties** - Documents `PeriodOfTime`, `startDate`, `endDate`, and handling of non-Gregorian time scales.
- **Page: DCAT Spatial Properties** - Details geometry, bounding box, and centroid specifications using GeoSPARQL/WKT.
- **Page: DCAT Checksums** - Explains the `spdx:Checksum` class for file integrity verification in DCAT 3.
- **Page: Managing Identifiers in DCAT** - Guides on HTTP IRIs, proxy IDs, legacy identifiers, and using `adms:identifier` with schema agencies.
