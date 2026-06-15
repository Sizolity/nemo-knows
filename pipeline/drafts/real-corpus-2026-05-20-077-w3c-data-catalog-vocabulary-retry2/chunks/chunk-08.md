---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
## Chunk Context
Lines 3413–3879 cover sections **9. License and rights statements**, **10. Time and space** (temporal and spatial properties), and the beginning of **11. Versioning**. The chunk details how to express licensing conditions, access rights, copyright notices, and ODRL policies using DCAT properties (`dcterms:rights`, `dcterms:license`, `dcterms:accessRights`, `odrl:hasPolicy`). It also defines temporal metadata (release time, revision, accrual periodicity, resolution, extent) and spatial metadata (resolution, polygon/centroid/bbox coverage). Finally, it introduces versioning relationships (`previousVersion`, `hasVersion`, `hasCurrentVersion`) to manage resource life-cycles.

## Local Summary
This section provides guidance on declaring rights statements, temporal metadata, and spatial coverage for datasets within DCAT. It distinguishes between licenses, access rights, and other rights (e.g., copyright) using specific Dublin Core properties. For time-based data, it defines properties for release time, update schedules, temporal resolution, and coverage intervals. Spatial properties cover resolution in meters and geometric representations (polygons, centroids, bounding boxes). Versioning is introduced to track resource revisions using a chain of previous/current versions linked to an abstract resource.

## Key Claims
- Selecting access conditions requires legal advice; DCAT distinguishes three scenarios: licenses, access rights only, and other rights statements.
- `dcterms:license` is for licenses (prefer canonical IRIs like Creative Commons); `dcterms:accessRights` for access restrictions; `dcterms:rights` for others (e.g., copyright).
- ODRL policies can be linked via `odrl:hasPolicy` for complex usage rules.
- Five temporal properties exist: release time (`dcterms:issued`), revision time (`dcterms:modified`), accrual periodicity, temporal resolution, and temporal extent.
- Spatial coverage uses `dcat:spatialResolutionInMeters` and `dcterms:spatial` with geometry (WKT) in CRS84 or specific CRS.
- Versioning uses `dcat:previousVersion`, `dcat:hasVersion`, and `dcat:hasCurrentVersion` to build version chains and hierarchies aligned with PAV.

## Entities And Concepts
- **Rights Statement Types**: License, Access Rights, Copyright (via `dcterms:rights`).
- **Vocabularies**: Creative Commons, EU-AR (`http://publications.europa.eu/resource/authority/access-right/`), ODRL.
- **Temporal Properties**: `dcterms:issued`, `dcterms:modified`, `dcterms:accrualPeriodicity`, `dcat:temporalResolution`, `dcterms:temporal`.
- **Spatial Properties**: `dcat:spatialResolutionInMeters`, `dcterms:spatial`, `locn:geometry`, `dcat:bbox`, `dcat:centroid`.
- **Versioning Properties**: `dcat:previousVersion`, `dcat:hasVersion`, `dcat:hasCurrentVersion`, `dcat:isVersionOf`.
- **Resources**: DCAT Dataset, Catalog, Distribution.

## Procedures And API Details
1. **Declaring Rights**: Use `dcterms:rights` with sub-properties `dcterms:license` or `dcterms:accessRights`. For copyright, use `rdfs:label` within a `dcterms:RightsStatement`. Link ODRL policies via `odrl:hasPolicy`.
2. **Defining Time**: Encode dates as `xsd:date`, durations as `xsd:duration`, and intervals using `time:ProperInterval` or `gYear`. Combine accrual periodicity with temporal resolution for time-series data.
3. **Defining Space**: Specify geometry in WKT format (default CRS84). Use `dcat:bbox` for bounding boxes, `locn:geometry` for polygons, and `dcat:centroid` for points.
4. **Managing Versions**: Link an abstract resource to versions using `dcat:hasVersion`. Use `dcat:previousVersion` for backward chaining and `dcat:hasCurrentVersion` to mark the latest version.

## Nuance Or Contradictions
- The specification notes that `dcterms:rights` is non-normative and advises legal consultation before applying license conditions.
- ODRL policies assume the enclosing entity is the subject unless an ODRL Asset is explicitly defined (per § 2.2.3 of ODRL-MODEL).
- Spatial examples default to CRS84 (WGS84) unless a specific CRS (e.g., EPSG:28992) is provided in the WKT literal.
- Versioning rules are left to data providers; DCAT does not dictate when a change constitutes a new release, referring instead to DWBP guidelines.

## Candidate Wiki Hints
- **Page: Rights Statements in DCAT** – Cover `dcterms:rights`, `dcterms:license`, `dcterms:accessRights`, and ODRL integration.
- **Page: Temporal Metadata for Datasets** – Explain temporal properties, time-series examples, and interval encoding.
- **Page: Spatial Coverage Encoding** – Detail WKT geometry usage, CRS handling, and spatial property choices.
- **Page: Versioning Strategies in DCAT** – Describe version chains, hierarchies, and use of PAV-aligned properties.
