---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

### Chunk Context
**Source Path:** `raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md`
**Chunk Range:** Lines 3413–3879
**Covered Sections:**
- Section 9: License and rights statements (including ODRL policies)
- Section 10: Time and space (Temporal properties, Spatial properties)
- Section 11: Versioning (Relationships between versions, Version chains and hierarchies)

### Local Summary
This chunk details how to express legal conditions, temporal metadata, spatial boundaries, and version histories for data cataloged resources using DCAT and related vocabularies. It distinguishes between licensing, access rights, and copyright statements, recommending specific properties from DCTERMS and ODRL. It further defines five temporal properties (release time, revision time, update schedule, temporal resolution, and temporal extent) and two spatial properties (spatial resolution and spatial extent), illustrating usage with WKT geometries and bounding boxes. Finally, it outlines a versioning model supporting version chains and hierarchies using `dcat:previousVersion`, `dcat:hasVersion`, and `dcat:hasCurrentVersion`.

### Key Claims
- Selecting conditions for resource access is complex; legal advice is required before implementation.
- Three distinct situations exist for rights statements: explicit 'license', access rights only, and other rights (e.g., copyright).
- DCAT recommends using `dcterms:rights`, `dcterms:license`, and `dcterms:accessRights` to address the three scenarios above.
- For interoperability, canonical IRIs of well-known licenses (e.g., Creative Commons) are preferred.
- Access rights can be expressed via code lists/taxonomies (e.g., [EUV-AR], Eprints Access Rights Vocabulary).
- The Open Data Rights Statement Vocabulary (ODRS) offers a more sophisticated approach extending DCTERMS for copyright notices.
- When using ODRL policies, the `odrl:hasPolicy` property links the resource to the policy.
- Five temporal properties are supported: release time (`dcterms:issued`), revision time (`dcterms:modified`), update schedule (`dcterms:accrualPeriodicity`), temporal resolution (`dcat:temporalResolution`), and temporal extent (`dcterms:temporal`).
- Two spatial properties are supported: minimum separation (`dcat:spatialResolutionInMeters`) and spatial extent (`dcterms:spatial`).
- Spatial geometries (for `locn:geometry`, `dcat:bbox`, `dcat:centroid`) should be specified using WKT; omitting CRS implies default CRS84.
- DCAT versioning focuses on revisions resulting from a resource's life-cycle, building upon [PAV], [DCTERMS], and [OWL2-OVERVIEW].
- Version chains are navigated backward via `dcat:previousVersion`, while hierarchies link abstract resources to versions via `dcat:hasVersion`.

### Entities And Concepts
- **Rights Statements:**
  - `dcterms:rights`: General property for rights statements not covered by license or access rights.
  - `dcterms:license`: Refers specifically to licenses (e.g., CC-BY 4.0).
  - `dcterms:accessRights`: Expresses access restrictions (public vs. authorized).
  - ODRL (Open Digital Rights Language): Policy expression language for permissions, prohibitions, and obligations.
  - ODRS (Open Data Rights Statement Vocabulary): Extends DCTERMS for sophisticated rights specification.
- **Temporal Properties:**
  - `dcterms:issued`: Release time (usually `xsd:date`).
  - `dcterms:modified`: Revision/update time (usually `xsd:date`).
  - `dcterms:accrualPeriodicity`: Update schedule (controlled vocabulary).
  - `dcat:temporalResolution`: Minimum temporal separation (`xsd:duration`).
  - `dcterms:temporal`: Temporal extent (`dcterms:PeriodOfTime`).
- **Spatial Properties:**
  - `dcat:spatialResolutionInMeters`: Minimum spatial separation (decimal number).
  - `dcterms:spatial`: Spatial extent (`dcterms:Location`).
  - WKT (Well-Known Text): Encoding for geometries.
  - CRS84 / WGS84: Default Coordinate Reference System.
- **Versioning Concepts:**
  - Version chain/history: Linear progression of versions.
  - Version hierarchy: Relationship between an abstract resource and its versions.
  - `dcat:previousVersion`: Links a version to its predecessor.
  - `dcat:hasVersion`: Links an abstract resource to its versions.
  - `dcat:hasCurrentVersion`: Links an abstract resource to the current snapshot.
  - `dcat:isVersionOf`: Inverse of `hasVersion`; links a version back to its abstract resource.

### Procedures And API Details
- **Expressing Rights:**
  1. Use `dcterms:license` for licenses (prefer canonical IRIs like Creative Commons).
  2. Use `dcterms:accessRights` for access restrictions (e.g., using [EUV-AR] code list).
  3. Use `dcterms:rights` with `rdfs:label` for copyright statements or other rights.
  4. For ODRL policies, use `odrl:hasPolicy` to link the resource/distribution to the policy object.
- **Expressing Temporal Properties:**
  - Encode dates as `xsd:date`.
  - Encode durations as `xsd:duration` (e.g., "PT15M").
  - Use `dcterms:PeriodOfTime` for temporal extent; can be combined with `time:ProperInterval` and `time:Instant`.
  - Support various time representations: `xsd:date`, `xsd:gYear`, or geologic age via `time:TimePosition` and specific TRS (e.g., GeologicAge).
- **Expressing Spatial Properties:**
  - Encode geometries in WKT format within `geosparql:wktLiteral`.
  - Omit CRS specification to imply default CRS84.
  - Specify polygons for detailed coverage, centroids for points, or bounding boxes (`dcat:bbox`) for general areas.
- **Defining Version Hierarchy:**
  - Define an abstract resource (e.g., `mycity-bus:stops`).
  - Link specific versions to the abstract resource using `dcat:hasVersion`.
  - Identify the current version using `dcat:hasCurrentVersion`.
  - Link individual versions backward using `dcat:previousVersion`.
  - Link versions back to the abstract resource using `dcat:isVersionOf`.

### Nuance Or Contradictions
- **Legal Advice:** The specification explicitly states that implementers must seek legal advice before deciding on access conditions, implying these recommendations are not legally binding mandates.
- **Non-Normative Sections:** Sections 9 (License), 10 (Time and space), and 11 (Versioning) are marked as non-normative, suggesting they provide guidance rather than strict requirements.
- **Complementary Versioning:** DCAT's versioning approach is meant to complement existing domain-specific practices (e.g., [OWL2-OVERVIEW] for ontologies) rather than replace them entirely. Providers decide when to release a new version based on their policies and workflows; DCAT does not define rules for triggering version releases.
- **Geometry CRS Default:** While specific CRS can be specified, omitting it defaults to CRS84 (WGS84 with longitude/latitude axis order), which may differ from expectations in other contexts using latitude/longitude first.

### Candidate Wiki Hints
- **Page: DCAT Rights Statements**
  - Topic: How to structure rights information (License vs. Access Rights vs. Copyright).
  - Content: Usage of `dcterms:rights`, `dcterms:license`, `dcterms:accessRights`, and integration with ODRL/ODRS.
- **Page: DCAT Temporal Metadata**
  - Topic: Describing time-related properties of datasets.
  - Content: Definitions of release, revision, schedule, resolution, and extent; examples of time-series data representation.
- **Page: DCAT Spatial Coverage**
  - Topic: Defining geographic boundaries for datasets.
  - Content: Usage of `dcterms:spatial`, WKT encoding, CRS handling, and examples (polygon, centroid, bbox).
- **Page: DCAT Versioning Model**
  - Topic: Managing resource versions and lifecycles in catalogs.
  - Content: Properties for version chains (`previousVersion`) and hierarchies (`hasVersion`, `hasCurrentVersion`), abstract resources, and relationship to PAV ontology.
