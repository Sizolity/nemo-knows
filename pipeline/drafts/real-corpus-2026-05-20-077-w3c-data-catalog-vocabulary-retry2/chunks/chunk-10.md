---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

# Chunk Context

**Heading path:** `12. Dataset series > 12.2 Dataset series metadata`
**Line range:** 4243–4637
**Coverage:** This chunk details how to model dataset series metadata, specifically properties inherited from child datasets (upstream inheritance), strategies for calculating series-level temporal/spatial coverage, and handling of dates. It also covers alternative implementation patterns in existing DCAT versions, data citation requirements, and quality information modeling using DQV.

# Local Summary

This section explains how to aggregate metadata properties from individual child datasets into a parent dataset series. Properties are categorized into those describing the series itself (e.g., accrual periodicity) and those reflecting dimensions described in children via inheritance. For inherited properties like temporal and spatial coverage, the series metadata typically takes the union of child values (e.g., a time range spanning all child start/end dates, or a bounding box encompassing all child locations). Date handling follows specific rules: the series creation date is the earliest child creation date, publication date is the earliest child publication date, and update date is the latest child update date. The chunk provides an example (Example 40) of yearly budget data where the series covers multiple countries and years. It notes that while mechanisms can automate this inheritance, DCAT does not mandate a specific strategy. Subsequent sections discuss non-normative guidance on data citation (referencing identifiers, creators, titles, etc.) and quality information (using DQV for annotations, measurements, and policies), including examples of conformance to standards like INSPIRE regulations.

# Key Claims

- Dataset series metadata properties fall into two groups: those describing the series itself and those reflecting child dataset dimensions via upstream inheritance.
- For inherited properties, the dataset series value is typically the union of values specified in child datasets (e.g., temporal coverage spans all child years; spatial coverage unites bounding boxes).
- Date handling for series follows specific rules: creation date = earliest child creation; publication date = earliest child publication; update date = latest child update/publication.
- DCAT does not recommend a specific strategy for implementing upstream inheritance automatically, though mechanisms can be put in place.
- Existing DCAT implementations may use alternative typing strategies (e.g., series as `dcat:Dataset` with children as `dcat:Distribution`, or both as `dcat:Dataset` linked via `dcterms:hasPart`).
- Data citation requires including dataset identifier, creator(s), title, publisher, and publication/release date.
- Quality information can be documented using DQV, distinguishing between user feedback (`dqv:QualityAnnotation`) and automated measurements (`dqv:QualityMeasurement`).
- Conformance to standards should use IRIs from reference registries (e.g., W3C TR, OGC CRS Registry) rather than namespace IRIs.

# Entities And Concepts

- **Dataset Series (`dcat:DatasetSeries`)**: A collection of related child datasets.
- **Upstream Inheritance**: The process where properties of child datasets are inherited or aggregated by their parent series.
- **Temporal Resolution (`dcat:temporalResolution`)**: Frequency of new data addition (e.g., annual).
- **Temporal Coverage (`dcat:temporal`)**: Time period covered by the series, derived from child datasets.
- **Spatial Coverage (`dcat:spatial`)**: Geographic area covered, represented as a union of bounding boxes or multiple reference systems.
- **Data Citation**: Practice of referencing data with bibliographic-style metadata (identifier, creator, title, etc.).
- **Quality Information**: Documentation of dataset quality using DQV, including annotations, policies, and measurements.
- **DQV (`dqv:`)**: Data Quality Vocabulary for modeling quality aspects like completeness, accuracy, and precision.
- **Conformance to Standards**: Representation of adherence to standards (e.g., INSPIRE Regulation) using `dcterms:conformsTo`.

# Procedures And API Details

**Aggregating Temporal Coverage:**
If child datasets cover different years (e.g., 2018, 2019, 2020), the series temporal coverage is the period between the earliest start date and latest end date of the children.

**Aggregating Spatial Coverage:**
If child datasets have different geographic bounding boxes, the series spatial coverage is the union of these boxes. If each child uses a different spatial reference system, the series includes multiple systems.

**Date Calculation Rules:**
- `dcterms:created` (series) = Earliest `dcterms:created` among children.
- `dcterms:issued` (series) = Earliest `dcterms:issued` among children.
- `dcterms:modified` (series) = Latest `dcterms:issued` or `dcterms:modified` among children.

**Example Code Snippet (Series Metadata):**
```turtle
ex:budget a dcat:DatasetSeries ;
dcterms:title "Budget data"@en ;
dcterms:issued "2019-01-01"^^xsd:date ,
dcterms:modified "2021-01-01"^^xsd:date ;
dcterms:accrualPeriodicity <http://purl.org/cld/freq/annual> ;
dcat:temporalResolution "P1Y"^^xsd:duration ;
dcterms:temporal [ a dcterms:PeriodOfTime ;
dcat:startDate "2018-01-01"^^xsd:date ;
dcat:endDate "2020-12-31"^^xsd:date ;
] ;
dcterms:spatial <http://publications.europa.eu/resource/dataset/country/BEL> ,
<http://publications.europa.eu/resource/dataset/country/FRA> ,
<http://publications.europa.eu/resource/dataset/country/ITA> ,
... ;
.
```

**Quality Measurement Example:**
```turtle
ex:genoaBusStopsDatasetCompletenessMeasurement a dqv:QualityMeasurement ;
dqv:computedOn ex:genoaBusStopsDataset ;
dqv:isMeasurementOf ex:completenessWRTExpectedNumberOfEntities ;
dqv:value "0.6833333"^^xsd:decimal ;
prov:wasAttributedTo ex:myQualityChecker ;
prov:generatedAtTime "2018-05-27T02:52:02Z"^^xsd:dateTime ;
prov:wasGeneratedBy ex:myQualityChecking .
```

**Standard Conformance Example:**
```turtle
ex:Dataset-1 a dcat:Dataset;
dcterms:conformsTo <http://data.europa.eu/eli/reg/2014/1312/oj> .
```

# Nuance Or Contradictions

- **Implementation Variations:** Existing DCAT implementations may not strictly use `dcat:DatasetSeries` but instead type the series as a `dcat:Dataset` with children as distributions, or link them via `dcterms:hasPart`. These are not formally incompatible with DCAT and can coexist during upgrades.
- **Quality Information Scope:** The examples provided do not specify where quality information resides or how it is managed, leaving implementation choices to data portals (e.g., UI annotations vs. third-party services).
- **Standard IRI Selection:** While using IRIs from reference registries is recommended, there is flexibility in choosing between un-versioned and versioned IRIs depending on whether conformance to a specific version needs to be stated.

# Candidate Wiki Hints

- **Dataset Series Metadata Aggregation**: How to combine temporal/spatial coverage and dates from child datasets into a series.
- **Data Citation Requirements**: Essential metadata fields for supporting data citation (identifier, creator, title, publisher, date).
- **Quality Information Modeling**: Using DQV to document quality annotations, measurements, and policies.
- **Standard Conformance Patterns**: Representing adherence to standards like INSPIRE or coordinate reference systems using `dcterms:conformsTo`.
- **DCAT Implementation Strategies**: Alternatives for modeling dataset series in existing DCAT versions.
