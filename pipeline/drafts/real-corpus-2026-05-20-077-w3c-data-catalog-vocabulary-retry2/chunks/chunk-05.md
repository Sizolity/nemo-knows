---
title: Chunk NN Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
### Chunk Context
This chunk covers sections 6.4.32 through 6.7 of the DCAT vocabulary specification, detailing properties for ordered collections (`last`, `previous`), the optional `Catalog Record` class, and extensive properties for the `Dataset` class including distributions, temporal/spatial coverage, and provenance. It concludes with the introduction of the `Dataset Series` class in DCAT 3.

### Local Summary
The text defines how to represent ordered sequences of datasets using `dcat:last`, `dcat:prev`, and `dcat:first`. It distinguishes between a dataset's intrinsic publication date and the date it was entered into a catalog via the optional `CatalogRecord` class. The `Dataset` class is defined broadly to include various media types, with specific properties for frequency, spatial/temporal resolution, and generation activities. Finally, it introduces `DatasetSeries` to group related datasets (e.g., time series or map series).

### Key Claims
- **Ordered Collections**: Properties `dcat:last`, `dcat:prev`, and `dcat:first` describe a dataset's position within an ordered collection or series.
- **Catalog Record Distinction**: The `dcat:CatalogRecord` class allows separating metadata about the dataset itself from metadata about its registration entry in the catalog (e.g., listing date vs. publication date).
- **Dataset Breadth**: The `dcat:Dataset` concept is broad and inclusive, accommodating numbers, text, pixels, imagery, sound, and other multimedia.
- **Temporal/Spatial Resolution**: Specific properties (`dcat:temporalResolution`, `dcat:spatialResolutionInMeters`) summarize the resolution of data distributions as single values, distinct from accrual frequency or coverage periods.
- **Dataset Series**: Introduced in DCAT 3, this class groups datasets that share characteristics but are published separately (e.g., periodic releases).

### Entities And Concepts
- `dcat:DatasetSeries`: A collection of datasets sharing characteristics.
- `dcat:CatalogRecord`: An optional record describing the registration of a single resource in a catalog.
- `dcat:last`, `dcat:prev`, `dcat:first`: Properties for ordered collections.
- `dcterms:accrualPeriodicity`: Frequency at which a dataset is published.
- `dcat:spatialResolutionInMeters`: Minimum spatial separation resolvable in meters.
- `dcat:temporalResolution`: Minimum time period resolvable in the dataset.
- `prov:wasGeneratedBy`: Links a dataset to the activity generating it.

### Procedures And API Details
- **Using `last` and `previous`**: Use `dcat:last` for the final resource in a series and `dcat:prev` for the immediately preceding resource. Note that `dcat:prev` refers to a distinct resource, not a previous version of the same resource (unlike `dcat:previousVersion`).
- **Catalog Entry Dates**: When both dataset publication date and catalog listing date exist, use `dcterms:issued` for the catalog entry date. Use `dcterms:modified` for updates to the catalog metadata description.
- **Spatial Resolution Encoding**: Values for `dcat:spatialResolutionInMeters` should be typed as `xsd:decimal`. While integers are semantically valid, validation tools like SHACL may require explicit datatype declarations.
- **Dataset Series Semantics**: Inherited properties (like `dcterms:type`) retain standard semantics but can have particular meanings when used with `dcat:DatasetSeries`.

### Nuance Or Contradictions
- **Date Distinction**: There is a clear distinction between the date a dataset was made available (`release date`) and the date it was added to the catalog (`listing date`). Confusing these can lead to metadata inaccuracies.
- **Version vs. Predecessor**: `dcat:prev` denotes a distinct resource in a sequence, whereas `dcat:previousVersion` denotes a version history of the same resource. These are not interchangeable.
- **Datatype Strictness**: While `xsd:integer` is a derived type of `xsd:decimal` and thus valid for resolution values, some syntactic validators (SHACL, ShEx) treat them as distinct. Explicit string formatting with `xsd:decimal` datatype is recommended to avoid validation errors.

### Candidate Wiki Hints
- **Page**: Dataset Series (`dcat:DatasetSeries`)
  - *Rationale*: The chunk introduces a new class in DCAT 3 specifically for grouping datasets (time series, map series), representing a significant conceptual expansion over previous versions.
- **Page**: Catalog Record vs. Dataset Metadata
  - *Rationale*: The distinction between `dcat:Dataset` and `dcat:CatalogRecord` is nuanced and important for data governance scenarios where cataloging metadata differs from the dataset's intrinsic content.
