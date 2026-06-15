---
title: Chunk NN Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
## Chunk Context
This chunk covers sections 6.4.32 through 6.7 of the DCAT vocabulary specification. It details properties for ordered resource collections (`last`, `previous`), defines the optional `Catalog Record` class to distinguish between dataset metadata and catalog entry metadata, specifies core `Dataset` properties including distribution, frequency, spatial/temporal coverage, and provenance, and introduces the new `Dataset Series` class in DCAT 3.

## Local Summary
The text outlines specific vocabulary extensions for handling resource series (first/last/previous), clarifies the distinction between a dataset's intrinsic metadata and its catalog entry via the `Catalog Record` class, lists standard properties for datasets including spatial/temporal resolution and provenance generation, and defines the `Dataset Series` class to group related datasets sharing characteristics like time or theme.

## Key Claims
- The `last` and `previous` properties were added in DCAT 3 to manage ordered collections of resources within a series.
- The `dcat:CatalogRecord` class is optional and serves to distinguish metadata about the catalog entry itself (e.g., listing date) from metadata about the underlying dataset (e.g., release date).
- Information regarding licenses and rights should primarily be provided at the `Distribution` level, though it may also appear on a `Dataset` if necessary, but differing values between them should be avoided to prevent legal conflicts.
- The `dcat:DatasetSeries` class was added in DCAT 3 to represent collections of separately published datasets that share characteristics (e.g., time series, map series).

## Entities And Concepts
- **Classes**: `dcat:CatalogRecord`, `dcat:Dataset`, `dcat:DatasetSeries`.
- **Properties**: `dcat:last`, `dcat:prev` (previous), `dcterms:title`, `dcterms:description`, `dcterms:issued` (listing date), `dcterms:modified` (update/modification date), `foaf:primaryTopic`, `dcterms:conformsTo`, `dcat:distribution`, `dcterms:accrualPeriodicity` (frequency), `dcat:inSeries`, `dcterms:spatial`, `dcat:spatialResolutionInMeters`, `dcterms:temporal`, `dcat:temporalResolution`, `prov:wasGeneratedBy`.
- **Concepts**: Dataset Series, Catalog Entry vs. Dataset Metadata, Spatial Resolution, Temporal Resolution, Provenance Activity.

## Procedures And API Details
- **Usage of `last` and `previous`**: These properties link a resource to its predecessor or successor in an ordered series (`dcat:DatasetSeries`). They are distinct from versioning properties like `previousVersion`.
- **Catalog Record vs. Dataset**: When listing dates differ (dataset publication vs. catalog addition), the date of the catalog entry should be specified on the `CatalogRecord`, while the dataset's original availability date remains on the `Dataset`.
- **Spatial Resolution Encoding**: Values for `dcat:spatialResolutionInMeters` must be typed as `xsd:decimal`. Note that JSON-LD may convert numbers to `xsd:double` or `xsd:integer`; validation schemas should accept `xsd:integer` as valid for this property.
- **Temporal Resolution vs. Frequency**: `dcterms:accrualPeriodicity` describes the rate at which the dataset is published/updated, whereas `dcat:temporalResolution` indicates the minimum time period resolvable within the data (e.g., spacing of items in a time series).

## Nuance Or Contradictions
- **Sub-class Removal**: In DCAT 1, `dcat:Dataset` was a sub-class of `dctype:Dataset`. This relationship was removed in DCAT 2 to broaden the scope of `dcat:Dataset` to include multimedia and text from the DCMI Types vocabulary.
- **Datatype Nuance**: While `xsd:integer` is semantically valid for `dcat:spatialResolutionInMeters` (as it derives from `xsd:decimal`), strict validation tools like SHACL or ShEx may treat them as distinct datatypes, requiring explicit datatype handling in schema definitions.
- **Inverse Properties**: The document references the use of inverse properties (Section 7) and specific guidance on series usage (Section 12), implying that relationships like `inSeries` might be traversed inversely depending on implementation patterns.

## Candidate Wiki Hints
- **Page: DCAT Dataset Series** - Covers the definition, usage scenarios (time series, map series), and related properties (`first`, `last`, `previous`, `inSeries`).
- **Page: DCAT Catalog Record** - Explains the optional nature of this class and the distinction between catalog entry metadata and dataset metadata.
- **Page: DCAT Spatial and Temporal Properties** - Details resolution vs. coverage, frequency vs. accrual periodicity, and encoding standards for spatial/temporal data.
