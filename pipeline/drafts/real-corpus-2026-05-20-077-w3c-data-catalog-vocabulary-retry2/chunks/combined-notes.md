## chunk-01

---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

# Chunk Context

**Heading path:** Document → W3C Data Catalog Vocabulary > Fetch Metadata / Retrieved Text > 1. Introduction / 2. Motivation for change
**Line range:** 1–473 (approx.)
**Source:** https://www.w3.org/TR/vocab-dcat-3/

This chunk introduces the **W3C Data Catalog Vocabulary (DCAT) Version 3**, a W3C Recommendation published on 22 August 2024. It outlines the purpose of DCAT, its evolution from earlier versions (DCAT 1 and DCAT 2), key design motivations for change, and high-level structural elements such as namespaces, conformance criteria, and vocabulary scope. The text emphasizes interoperability, federated search capabilities, and support for diverse data formats and access methods.

---

# Local Summary

DCAT 3 is an RDF-based vocabulary designed to describe datasets and data services in web-accessible catalogs. It aims to improve metadata interoperability across multiple organizations and platforms by standardizing how data resources are described. The vocabulary distinguishes between abstract datasets and their distributions (e.g., downloadable files or API endpoints). DCAT 3 introduces new features such as versioning, dataset series, checksums via SPDX, and inverse properties while maintaining backward compatibility with DCAT 2. It incorporates terms from external vocabularies like FOAF and DCTERMS where appropriate but defines its own namespace (`http://www.w3.org/ns/dcat#`) with a suggested prefix of `dcat`. The document is published as a W3C Recommendation under the Dataset Exchange Working Group and follows W3C patent policy and process guidelines.

---

# Key Claims

- DCAT 3 supersedes DCAT 2 but does not render it obsolete; both versions coexist with preserved backward compatibility.
- DCAT enables decentralized publishing of data catalogs and supports federated search across multiple sites using uniform query mechanisms.
- Aggregated DCAT metadata can function as manifest files in digital preservation workflows.
- DCAT 3 adds support for versioning (`dcat:version`, `dcat:previousVersion`, etc.), dataset series (`dcat:DatasetSeries`), and inverse properties.
- The vocabulary includes integration points with external standards such as SPDX (for checksums) and Schema.org (via alignment section).
- DCAT is non-prescriptive regarding deployment methods and supports serialization in multiple formats including Turtle, RDF/XML, JSON-LD, and N3.

---

# Entities And Concepts

- **DCAT**: Data Catalog Vocabulary – an RDF vocabulary for describing datasets and services.
- **DCAT 1 / VOCAB-DCAT-1**: Original version standardized in 2014 by the Government Linked Data Working Group.
- **DCAT 2 / VOCAB-DCAT-2**: Second revision addressing shortcomings identified through community use cases.
- **DCAT 3 / VOCAB-DCAT-3**: Current version (Recommendation as of August 2024), extending DCAT 2 with new classes and properties.
- **Namespace**: `http://www.w3.org/ns/dcat#` with prefix `dcat`.
- **Classes**: `Catalog`, `Cataloged Resource`, `Catalog Record`, `Dataset`, `Distribution`, `Data Service`, `Dataset Series`, `Checksum`, etc.
- **Properties**: e.g., `title`, `description`, `license`, `distribution`, `version`, `hasCurrentVersion`, `replaces`, `access URL`, `download URL`, `checksum`.
- **External vocabularies integrated**: FOAF (`foaf:homepage`), DCTERMS (`dcterms:title`), SPDX (for checksums).
- **Profiles**: DCAT profiles allow domain-specific extensions while remaining compatible with the base vocabulary.

---

# Procedures And API Details

No specific APIs are defined in this chunk; however, DCAT metadata may be exposed via:

- SPARQL endpoints
- HTML pages using RDFa (`[HTML-RDFa]`)
- Serializations such as Turtle, RDF/XML, JSON-LD, N3

Implementation guidance includes:

- Use of `spdx:checksum` and `spdx:Checksum` for digest representation.
- Representation of version chains using properties like `dcat:previousVersion`, `dcat:hasCurrentVersion`, and `dcat:replaces`.
- Specification of dataset series via `dcat:DatasetSeries` class and related properties.
- Inverse property usage patterns are documented in Section 7 of the full spec (not fully covered here).

---

# Nuance Or Contradictions

- **Backward Compatibility**: DCAT 3 preserves definitions from DCAT 2, allowing existing deployments to remain compliant without modification unless they adopt new features.
- **Non-normative Sections**: Introduction and Motivation sections are explicitly marked as non-normative; only Section 4 (Conformance) and the Vocabulary Specification (Section 6) define normative requirements.
- **External Term Definitions**: Informal summaries of externally defined terms (e.g., from DCTERMS or FOAF) are included for convenience, but authoritative definitions come from their original sources. Changes to those external definitions do not affect DCAT conformance.
- **Deployment Flexibility**: DCAT does not prescribe any particular method for deploying catalogs; implementations may choose among various RDF serialization formats and exposure mechanisms.

---

# Candidate Wiki Hints

Potential wiki pages or sections inspired by this chunk:

- **DCAT 3 Overview** – Summary of purpose, scope, and key enhancements over prior versions.
- **DCAT Version History** – Timeline from DCAT 1 (2014) through DCAT 2 (2020) to DCAT 3 (2024).
- **Namespace and Prefix Usage** – Details on `http://www.w3.org/ns/dcat#` and recommended prefix `dcat`.
- **Interoperability Benefits** – How DCAT enables federated search, aggregation, and cross-catalog discovery.
- **Backward Compatibility Strategy** – Explanation of how DCAT 3 maintains compatibility with DCAT 2 implementations.
- **External Vocabulary Integration** – Use of FOAF, DCTERMS, SPDX, and alignment with Schema.org.
- **Serialization Formats** – Supported formats for representing DCAT metadata (Turtle, JSON-LD, etc.).

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

Chunk Context
- **Source**: `raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md`
- **Heading Path**: 3. Namespaces through 5.4 Classifying datasets thematically
- **Line Range**: 474–867

Local Summary
This chunk details the namespace definitions for DCAT, distinguishing between normative and non-normative terms used in RDF representations. It outlines conformance criteria for data catalogs, defining what constitutes a compliant catalog versus a profile. The text introduces the core vocabulary scope (DCAT classes like `dcat:Catalog`, `dcat:Dataset`, `dcat:Distribution`) and provides a basic Turtle syntax example of a government catalog with multilingual labels. It also covers RDF considerations regarding blank nodes and concludes with examples on thematic classification using SKOS.

Key Claims
- DCAT defines a minimal set of classes and properties while extensively utilizing terms from external vocabularies like Dublin Core (DCTERMS).
- Conformance to DCAT requires organizing data access into datasets, distributions, data services, and dataset series, with an available RDF description of the catalog resources.
- DCAT profiles are specifications that add constraints to DCAT, such as cardinality limits or specific required metadata fields.
- Blank nodes are generally discouraged in Linked Data contexts because they limit collaborative annotation and interoperability; instances of main DCAT classes should have global identifiers (IRIs).
- The scope of DCAT currently focuses on datasets and data services, though it allows for extension via profiles for other resource types.

Entities And Concepts
- **dcat:Catalog**: Represents a catalog where individual items are metadata records describing resources; scope includes collections of metadata about datasets, data services, or other resource types.
- **dcat:Resource**: A parent class for `dcat:Dataset`, `dcat:DataService`, and others; serves as an extension point for defining catalogs of any kind of resources.
- **dcat:Dataset**: Represents a collection of data published or curated by a single agent, available in one or more serializations (numbers, text, pixels, imagery, etc.).
- **dcat:Distribution**: Represents an accessible form of a dataset, such as a downloadable file.
- **dcat:DataService**: Represents a collection of operations accessible through an interface (API) providing access to datasets or processing functions.
- **dcat:DatasetSeries**: A dataset representing a collection of separately published datasets sharing characteristics.
- **dcat:CatalogRecord**: An optional metadata record concerning registration information (who added the record and when).
- **Namespace Prefixes**: Includes `adms`, `dc`, `dcterms`, `foaf`, `owl`, `prov`, `rdf`, `rdfs`, `skos`, `spdx`, `time`, `vcard`, `xsd` (normative) and `dqv`, `earl`, `geosparql`, `oa`, `pav`, `sdo`, `xhv` (non-normative).

Procedures And API Details
- **Namespace Declaration**: Normative namespaces are listed in a table with prefixes, IRI sources, and external references (e.g., `[VOCAB-DCAT]`). Non-normative namespaces include those used in examples like `sdo` for Schema.org.
- **Conformance Check**: A catalog conforms to DCAT if it organizes access into the specified resource types, provides an RDF description of the catalog and resources, includes metadata fields using appropriate DCAT classes/properties, and uses defined classes consistently.
- **Profile Definition**: Profiles add constraints such as minimum required metadata fields, sub-classes of standard properties, controlled vocabularies, or requirements for specific access mechanisms (RDF syntaxes, protocols).
- **Thematic Classification Procedure**: Use `dcat:themeTaxonomy` to link a catalog to a concept scheme (`skos:ConceptScheme`) which defines domains. Link individual datasets to specific concepts within this scheme using `dcat:theme`.

Nuance Or Contradictions
- **Blank Nodes vs. IRIs**: While RDF allows blank nodes, the specification recommends against them for DCAT instances because they undermine the "anyone can say anything anywhere" benefit of Linked Data and hinder collaborative annotation.
- **Normative vs. Non-normative**: The document explicitly marks sections like "Vocabulary overview" and diagrams as non-normative, distinguishing them from the core normative constraints defined in Section 4.
- **DCAT 1 vs. DCAT 2**: DCAT 1 was limited to catalogs of datasets. DCAT 2 extends this scope to include data services as members of a catalog, addressing use cases for Web services and service-based access previously out of scope.

Candidate Wiki Hints
- **Page: DCAT Namespaces**
  - *Content*: Summary table of normative vs. non-normative prefixes, IRIs, and external vocabularies referenced by DCAT.
- **Page: DCAT Conformance and Profiles**
  - *Content*: Criteria for catalog compliance, definition of profiles (application profiles), and allowed extensions (cardinality, sub-classes).
- **Page: Core DCAT Classes**
  - *Content*: Definitions and relationships between `dcat:Catalog`, `dcat:Dataset`, `dcat:Distribution`, `dcat:DataService`, and others.
- **Page: RDF Best Practices for DCAT**
  - *Content*: Guidelines on using IRIs over blank nodes, Turtle syntax usage, and handling multilingual labels (language tags).

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
This chunk details the DCAT vocabulary specifications for classifying dataset types, describing catalog record metadata, handling datasets available via Web pages or services, and defining the core classes and properties for Catalogs and Cataloged Resources. It covers sections 5.5 through 6.4.8 of the source document.

Local Summary
The text outlines how to classify datasets using `dcterms:type` with references to external vocabularies like DCMI or DataCite, allowing multiple type assignments. It distinguishes between catalog records (`dcat:CatalogRecord`) and raw dataset instances. The chunk further categorizes access patterns: datasets behind a landing page only, those available for download alongside a landing page, and those distributed via specific services (e.g., APIs). Finally, it defines the RDF structure for `dcat:Catalog` and `dcat:Resource`, listing their specific and inherited properties such as `homepage`, `themes`, `resource`, `dataset`, `service`, and various metadata fields like `creator` and `release date`.

Key Claims
- Dataset classification should utilize well-governed, recognized vocabularies (e.g., DCMI Type Vocabulary, MARC Genre/Terms Scheme).
- Multiple classifications can coexist on a single dataset description using multiple `dcterms:type` properties.
- `dcat:CatalogRecord` is used to describe the registration of a resource within a catalog, distinct from the resource itself.
- Access patterns are modeled via `dcat:landingPage`, `dcat:distribution`, and `dcat:accessURL`.
- Services are represented using `dcat:DataService`, characterized by type, conformance, and endpoint descriptions.
- `dcat:Catalog` is a sub-class of `dcat:Dataset` in DCAT 2/3, enabling catalogs to be composed of other catalogs.
- `dcat:Resource` serves as the super-class for all cataloged items (datasets, services, catalogs).

Entities And Concepts
- **Vocabularies**: dcterms:type, DCMI Type Vocabulary, MARC Genre/Terms Scheme, ISO-19115-1 MD_Scope codes, DataCite resource types, Re3data content-types.
- **Classes**: `dcat:Dataset`, `dcat:CatalogRecord`, `dcat:Distribution`, `dcat:DataService`, `dcat:Catalog`, `dcat:Resource`.
- **Properties**: `homepage`, `themes` (themeTaxonomy), `resource`, `dataset`, `service`, `catalog`, `record`, `access rights`, `conforms to`, `contact point`, `creator`, `description`, `title`, `release date`, `update/modification date`.

Procedures And API Details
- **Multiple Typing**: To assign multiple types, use multiple triples with `dcterms:type` pointing to different URIs (e.g., combining DCMI and DataCite values).
- **Service Distribution**: A distribution linked to a service uses `dcat:accessService` pointing to the specific `dcat:DataService` instance. The service instance must define `dcat:endpointURL`, `dcterms:type`, and `dcterms:conformsTo`.
- **Landing Page Handling**: If data is only behind a Web page, define `dcat:landingPage` and a `dcat:Distribution` with the same access URL. If downloadable, use `dcat:downloadURL` on the distribution instance.

Nuance Or Contradictions
- In DCAT 1, `dcat:Catalog` was strictly for datasets; in DCAT 2/3, it is generalized as a sub-class of `dcat:Dataset` and can contain other catalogs.
- The domain of `dcat:contactPoint` was relaxed from `dcat:Dataset` to `dcat:Resource` in DCAT 2 to allow broader usage.
- Definitions for terms outside the DCAT namespace (e.g., from DCTERMS or FOAF) are provided for convenience but are not normative; authoritative definitions remain in their original specifications.

Candidate Wiki Hints
- **Page**: Classifying Dataset Types – Covers using `dcterms:type` and external vocabularies.
- **Page**: Catalog Record Metadata – Explains the difference between a dataset instance and its catalog record.
- **Page**: Access Patterns – Details modeling datasets behind landing pages, downloadable files, and service-based access.
- **Page**: DCAT Class Definitions – Defines `dcat:Catalog` and `dcat:Resource` hierarchies and properties.

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

## Chunk Context
This chunk covers vocabulary specifications for the `Cataloged Resource` class in DCAT, spanning properties from **language** (6.4.9) through **first** (6.4.31). It details definitions, ranges, sub-properties, and usage notes for metadata describing resource attributes, relationships, versions, statuses, and series membership.

## Local Summary
The section defines RDF properties used to describe cataloged resources, ranging from basic attributes like language and publisher to complex relationship patterns involving qualified relations and versioning. It highlights transitions between DCAT 1, 2, and 3, noting domain relaxations and the introduction of new properties for version chains and resource series.

## Key Claims
- **Language**: Use ISO 639-1 codes; fall back to ISO 639-2 if ISO 639-1 is undefined. Values for members override catalog-level values.
- **Publisher**: Corresponds to general attribution with the role 'publisher'; `foaf:Agent` resources are recommended.
- **Identifier**: Represents a unique text string assigned to a resource; may be used in the IRI but explicitly represented as an `rdfs:Literal`.
- **Theme/Category**: Organized within a `skos:ConceptScheme`; DCAT 3 treats it as an OWL object property to avoid automatic inference as `skos:Concept`.
- **Type/Genre**: Values should come from controlled vocabularies (e.g., DCMI, ISO 19115). File formats and physical media should use `dcterms:format` instead.
- **Relation**: Generic link when relationship nature is unknown; specific sub-properties (e.g., `dcat:distribution`) are preferred if the relationship type is known.
- **Qualified Relation**: Used for relationships not covered by standard DCTERMS or PROV-O properties, linking to a `dcat:Relationship`.
- **Versioning**: DCAT 3 introduces specific properties (`previousVersion`, `hasVersion`, `currentVersion`) limited to revisions in a resource's life-cycle, distinct from broader notions of edition/adaptation.
- **Series**: Properties like `first` are used for resources belonging to a `dcat:DatasetSeries`.

## Entities And Concepts
- **Classes**: `Cataloged Resource`, `Dataset`, `Distribution`, `Agent`, `Organization/Person`.
- **Vocabularies**: DCMI Type, ISO 19115, DataCite, PARSE.Insight (re3data), MARC.
- **Standards**: ISO 639-1, ISO 639-2, ISO 639-3, BCP47, PROV-O, ODRL, VOCAB-ADMS.
- **Properties**: `dcterms:language`, `dcterms:publisher`, `dcterms:identifier`, `dcat:theme`, `dcterms:type`, `dcterms:relation`, `dcat:qualifiedRelation`, `prov:qualifiedAttribution`, `dcterms:license`, `dcterms:rights`, `dcterms:hasPart`, `odrl:hasPolicy`, `dcterms:isReferencedBy`, `dcat:previousVersion`, `dcat:hasVersion`, `dcat:currentVersion`, `dcterms:replaces`, `dcat:version`, `adms:versionNotes`, `adms:status`, `dcat:first`.

## Procedures And API Details
- **Language Selection**: Check for ISO 639-1 code first; if absent, use the corresponding ISO 639-2 IRI.
- **Distribution Mapping**: If a dataset has separate representations per language, define distinct `dcat:Distribution` instances with specific `dcterms:language` values. Multilingual distributions have multiple language values.
- **Qualified Relation Usage**: Construct a link to another resource where the relationship nature is known but does not match standard properties. This entails using `dcat:qualifiedRelation` and linking to a `dcat:Relationship`.
- **Version Chain Construction**: Use `dcat:previousVersion`, `dcat:hasVersion`, and `dcat:currentVersion` to specify version chains consisting of snapshots resulting from revisions.
- **Series Membership**: Use `dcat:first` to identify the first resource in an ordered collection or series (e.g., `dcat:DatasetSeries`).

## Nuance Or Contradictions
- **Theme Range Change**: In DCAT 1, `dcat:theme` domain was limited to `dcat:Dataset`. DCAT 2 relaxed this, and DCAT 3 drops the range entirely to prevent automatic inference of objects as `skos:Concept`, ensuring formalization consistency.
- **Version Semantics**: `dcat:hasVersion` is a sub-property of `dcterms:hasVersion` but is more specific, limited to revisions in the life-cycle, whereas `dcterms:hasVersion` covers broader notions like editions and adaptations.
- **Identifier vs IRI**: While the identifier might be part of the resource's IRI, explicitly representing it as an `rdfs:Literal` is still considered useful.

## Candidate Wiki Hints
- Page on **DCAT Vocabulary Properties** covering attributes like language, publisher, and identifiers.
- Guide on **Managing Resource Versions** in DCAT 3, focusing on revision chains vs. editions.
- Article on **Qualified Relations** explaining the pattern for non-standard relationships using PROV-O and DCTERMS sub-properties.
- Documentation on **Dataset Series** utilizing `dcat:first` and related ordering properties.

## chunk-05

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

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
- Heading path: 6. Vocabulary specification > 6.8 Class: Distribution
- Line range: 2409–2920
- Covers properties of the `dcat:Distribution` class, definitions for `dcat:DataService`, and brief notes on `skos:ConceptScheme`, `skos:Concept`, and `foaf:Agent` classes.

Local Summary
- Defines the `dcat:Distribution` class as a specific representation (serialization) of a dataset, distinguishing it from the dataset itself.
- Lists 20 specific properties for describing distributions (e.g., title, license, access URLs, byte size).
- Clarifies the distinction between `accessURL` (service/location) and `downloadURL` (direct file link).
- Introduces new DCAT 2/3 properties like `spatialResolutionInMeters`, `temporalResolution`, `compressFormat`, `packageFormat`, and `checksum`.
- Defines the `dcat:DataService` class for describing API endpoints or services that serve datasets.
- Briefly introduces classes for categorization (`skos:ConceptScheme`, `skos:Concept`) and agents (`foaf:Person`, `foaf:Organization`).

Key Claims
- A dataset may have multiple distributions differing by format, resolution, or profile; they are not always fully informationally equivalent.
- `dcat:Distribution` implies general availability but does not specify the access method (download vs. API) without using `dcat:downloadURL` or `dcat:accessService`.
- License and rights information should be provided at the Distribution level; applying different license info to a Dataset than its Distributions creates legal conflicts.
- `dcat:mediaType` is preferred over generic `dcterms:format` when the media type is defined by IANA.
- `spatialResolutionInMeters` and `temporalResolution` are intended as single-value summaries; complex precision data belongs in the Data Quality Vocabulary.
- `checksum` (DCAT 3) links to a download URL to verify file integrity.

Entities And Concepts
- **dcat:Distribution**: A specific representation of a dataset (e.g., CSV, netCDF, JSON).
- **dcat:DataService**: A collection of operations providing access to datasets or processing functions.
- **skos:ConceptScheme**: A knowledge organization system for themes/categories.
- **skos:Concept**: A category used to describe datasets in a catalog.
- **foaf:Person / foaf:Organization**: Agents describing people or entities.

Procedures And API Details
- **Access URL (`dcat:accessURL`)**: Used for services, landing pages, or SPARQL endpoints. If only accessible via a landing page, duplicate the dataset's landing page URL here.
- **Download URL (`dcat:downloadURL`)**: Used for direct HTTP GET requests to downloadable files.
- **Service Linking**: Use `dcat:accessService` to link a distribution to a `dcat:DataService` description (e.g., OpenAPI, WSDL).
- **Resolution Properties**: Set `dcat:spatialResolutionInMeters` for grid/image spacing or minimum distance; set `dcat:temporalResolution` as an ISO 8601 duration for time-series spacing.

Nuance Or Contradictions
- **Format vs. Media Type**: `dcat:mediaType` is a sub-property of `dcterms:format`. Use `dcat:mediaType` for IANA-defined types; otherwise, use `dcterms:format`.
- **License vs. Rights**: `dcterms:license` links to a specific document (sub-property of rights). `dcterms:rights` allows broader statements including attribution. Both should generally be on the Distribution, not the Dataset, unless consistent.
- **Equivalence of Distributions**: While some distributions are losslessly transformable (e.g., RDF/XML vs. Turtle), others (e.g., CSV summary) may lose information but still represent distributions of the same dataset. Judgment on equivalence is application-specific.

Candidate Wiki Hints
- Page: DCAT Distribution Class Properties
  - Summary: Reference for all 20 properties of `dcat:Distribution`, including usage notes on license, resolution, and format.
- Page: DCAT Data Services
  - Summary: Overview of `dcat:DataService` and its relation to datasets via `dcat:servesDataset`.

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
This chunk covers the vocabulary specification for relationship, role, time, and location classes in DCAT 2 (and 3), plus sections on inverse properties and dereferenceable identifiers.

Local Summary
The document defines new RDF classes and properties introduced in DCAT 2 for expressing relationships between datasets, roles of agents or entities, temporal intervals, spatial locations, and file integrity checksums. It also clarifies how to use inverse properties, manage various identifier types (including non-HTTP IDs), and link external metadata.

Key Claims
- Class dcat:Relationship is an association class for attaching additional information to a relationship between DCAT resources; it subclasses prov:EntityInfluence.
- dcat:relation points to another dataset or cataloged resource in the context of dcat:Relationship.
- dcat:hadRole expresses the function of an entity or agent with respect to another entity or resource; recommended values come from controlled vocabularies like ISO-19115 CI_RoleCode, DS_AssociationTypeCode, IANA RELATIONS, DataCite, or MARC relators.
- dcat:Role is a subclass of skos:Concept used in qualified-attribution or qualified-relation; values should be managed as controlled vocabularies.
- dcterms:PeriodOfTime represents an interval defined by start and end; DCAT 2 uses dcat:startDate/dcat:endDate (ISO 8601) or time:hasBeginning/time:hasEnd (OWL-TIME).
- dcat:bbox, dcat:centroid, locn:geometry are used for spatial coverage; WKT literals support non-WGS84 coordinate systems.
- spdx:Checksum and its properties algorithm/checksumValue allow verification of file integrity using SPDX-defined algorithms (MD5, SHA families, etc.).
- DCAT intentionally omits inverse properties in the core vocabulary but allows them as additions, not replacements.
- HTTP IRIs are preferred for dereferenceable identifiers; legacy or non-HTTP identifiers can be expressed via dcterms:identifier, adms:identifier, and related metadata (adms:schemaAgency, skos:notation).
- owl:sameAs may link to external RDF descriptions (e.g., DOIs) that return machine-readable metadata.

Entities And Concepts
- dcat:Relationship, dcat:Role, dcterms:PeriodOfTime, dcterms:Location, spdx:Checksum
- Properties: dcat:relation, dcat:hadRole, dcat:startDate, dcat:endDate, time:hasBeginning, time:hasEnd, locn:geometry, dcat:bbox, dcat:centroid, spdx:algorithm, spdx:checksumValue
- Inverse properties: dcat:prev/dcat:next, dcat:distribution/dcat:isDistributionOf, dcterms:hasPart/dcterms:isPartOf, etc.
- Identifier types: HTTP IRIs, DOI, ORCID, VIAF, ISNI, locally minted identifiers, legacy identifiers
- Controlled vocabularies: ISO-19115 CI_RoleCode, DS_AssociationTypeCode, IANA RELATIONS, DataCite, MARC relators

Procedures And API Details
- Use dcat:startDate/dcat:endDate with ISO 8601 datatypes (xsd:gYear, xsd:date, xsd:dateTime) for temporal intervals.
- Use time:hasBeginning/time:hasEnd when the value is a time:Instant from OWL-TIME; implies dcterms:PeriodOfTime can be treated as time:ProperInterval.
- Encode geometry as WKT literals (geosparql:wktLiteral) within rdfs:Literal ranges for bbox/centroid; supports non-WGS84 CRS.
- For checksums, pair spdx:algorithm with spdx:checksumValue (xsd:hexBinary) to verify file integrity.
- Represent non-HTTP identifiers using adms:identifier with skos:notation and optionally adms:schemaAgency; use dcterms:publisher/creator for authority metadata.
- When an external IRI returns RDF/OWL, consider owl:sameAs to link the dataset to that description.

Nuance Or Contradictions
- DCAT 2 introduces dcat:hadRole and dcat:Role; spdx:Checksum is new in DCAT 3.
- Inverse properties are allowed only as additions to the core vocabulary, never as replacements.
- For DOIs, the authority (e.g., DOI Foundation) is recorded via adms:schemaAgency or dcterms:creator, not by naming the registrant organization, to preserve DOI stability across organizational changes.
- Identifier schemes registered as IRIs (e.g., info:doi/) do not need a separate 'type' annotation per RFC 3986.

Candidate Wiki Hints
- Page on DCAT relationship modeling (dcat:Relationship, dcat:relation, dcat:hadRole)
- Page on role vocabularies and controlled lists (ISO-19115, DataCite, MARC relators)
- Guide to temporal properties in DCAT (start/end dates vs time:hasBeginning/hasEnd)
- Spatial coverage patterns (geometry, bbox, centroid; WKT literals)
- File integrity with SPDX checksums (algorithm + hex digest)
- Identifier management strategies (HTTP IRIs, legacy IDs, adms:identifier, schemaAgency)
- Using owl:sameAs for external RDF metadata enrichment

## chunk-08

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

## chunk-09

---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

## Chunk Context
This chunk covers sections **11.1.2** through **12.1** of the DCAT vocabulary specification, focusing on version relationships, version information properties, resource life-cycles, complementary versioning approaches (OWL/PROV-O), and dataset series organization. It includes specific OWL examples illustrating how to model replaced versions, version notes, life-cycle statuses, and hierarchical series with internal versioning.

## Local Summary
The text defines how to express that one dataset version replaces another using `dcterms:replaces` and `dcat:isVersionOf`. It details properties for storing version identifiers (`dcat:version`), release dates (`dcterms:issued`), and change descriptions (`adms:versionNotes`). The section on resource life-cycles distinguishes versioning from the administrative status of a resource (e.g., accepted, deprecated) and lists various standard vocabularies (ISO-19135, ISO-19115, ADMS-SKOS) for these statuses. Finally, it introduces `dcat:DatasetSeries` to group related datasets (like yearly budgets), linking them via `dcat:inSeries` and managing the series timeline with `dcat:first`, `dcat:last`, `dcat:prev`, and `dcat:next`.

## Key Claims
- **Replacement Logic**: The relationship where a version supersedes another is explicitly modeled using `dcterms:replaces` (forward) and `dcterms:isReplacedBy` (inverse). A version does not automatically replace its immediate predecessor unless explicitly stated.
- **Version Metadata**: DCAT utilizes `dcat:version` for the name/identifier, `dcterms:issued` for release dates, and `adms:versionNotes` for textual descriptions of changes or compatibility issues.
- **Life-Cycle Orthogonality**: A resource's life-cycle status (e.g., stable, deprecated) is distinct from its versioning history; a new version does not always imply a status change, nor does a status change always imply a new version.
- **Complementary Ontologies**: DCAT properties can coexist with OWL (`owl:versionIRI`, `owl:priorVersion`) and PROV-O (`prov:generalizationOf`) properties. While they share semantic ground, they are not equivalent due to differing scopes (e.g., OWL for ontologies vs. DCAT for general catalog resources).
- **Dataset Series**: `dcat:DatasetSeries` is a subclass of `dcat:Dataset`. Individual datasets in a series are linked via `dcat:inSeries`. A series can be hierarchical, and datasets within it can still have their own version chains.

## Entities And Concepts
- **dcterms:replaces**: Property indicating a version supersedes another.
- **dcat:isVersionOf**: Inverse of `dcat:hasVersion`; used to link versions.
- **adms:versionNotes**: Textual description of changes in a resource version.
- **Resource Life-Cycle Status**: Administrative states like "accepted", "deprecated", "withdrawn", "under development".
- **ISO-19135**: Standard for item registration statuses (e.g., valid, superseded).
- **ISO-19115**: Standard for geospatial data progress codes.
- **ADMS-SKOS**: Vocabulary containing statuses like "completed", "deprecated", "withdrawn".
- **dcat:DatasetSeries**: Class representing a collection of interrelated datasets (e.g., time series).
- **dcat:inSeries**: Property linking a dataset to its parent series.
- **dcat:first / dcat:last**: Properties identifying the first and last datasets in a series.
- **dcat:prev / dcat:next**: Properties defining the sequential order within a series.

## Procedures And API Details
- **Modeling Replacement**:
  ```turtle
  :dataset-2015-05 a dcat:Dataset ;
      dcterms:replaces :dataset-2015-01 .
  ```
- **Specifying Version Info**:
  ```turtle
  :stops-2015-01-01 a dcat:Dataset ;
      dcterms:issued "2015-01-01"^^xsd:date ;
      dcat:version "1.0" ;
      adms:versionNotes "First version"@en .
  ```
- **Defining a Dataset Series**:
  ```turtle
  :budgetSeries a dcat:DatasetSeries ;
      dcterms:title "Budget data"@en .

  :budget-2018 a dcat:Dataset ;
      dcat:inSeries :budgetSeries .
  ```
- **Linking Series Timeline**:
  ```turtle
  :budgetSeries a dcat:DatasetSeries ;
      dcat:first :budget-2018 ;
      dcat:last :budget-2020 .

  :budget-2019 a dcat:Dataset ;
      dcat:prev :budget-2018 ;
      dcat:next :budget-2020 .
  ```
- **Combining Series and Versions**:
  ```turtle
  :budget-2019 a dcat:Dataset ;
      dcat:hasVersion :budget-2019-rev0, :budget-2019-rev1 ;
      dcat:hasCurrentVersion :budget-2019-rev1 .

  :budget-2019-rev1 a dcat:Dataset ;
      dcat:previousVersion :budget-2019-rev0 .
  ```

## Nuance Or Contradictions
- **Non-Equivalence**: Correspondence between DCAT properties and OWL/DCTERMS/PROV-O properties does not imply semantic equivalence. For instance, `prov:wasRevisionOf` is not meant for building strict version chains in the same way `dcat:previousVersion` is.
- **Optional Statuses**: DCAT does not prescribe a specific set of life-cycle statuses but refers to external standards (ISO, ADMS, EU Vocabularies) and community practices.
- **Series Membership**: The distinction between using distributions or datasets as members of a series was noted as a topic of discussion in GitHub issues (#868, #1429), implying the specification allows flexibility or acknowledges ongoing debate on implementation details.

## Candidate Wiki Hints
- **Page: DCAT Versioning Properties**
  - Focus: Mapping `dcat:version`, `dcterms:issued`, and `adms:versionNotes`.
  - Content: Explain the difference between version identifiers and life-cycle statuses.

- **Page: Dataset Series in DCAT**
  - Focus: Using `dcat:DatasetSeries`, `dcat:inSeries`, and temporal linking properties (`first`, `last`, `prev`, `next`).
  - Content: Examples of grouping yearly budget data or time-series datasets.

- **Page: Resource Life-Cycle Management**
  - Focus: Integrating life-cycle statuses (deprecated, withdrawn) with versioning strategies.
  - Content: List compatible vocabularies (ISO-19135, ADMS-SKOS) and best practices for status transitions.

## chunk-10

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

## chunk-11

---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

Chunk Context
- Location: Section 14.2.2 "Degree of conformance" through 14.2.3 "Conformance test results", followed by Section 15 "Qualified relations" and the start of Section 16 "DCAT Profiles".
- Scope: Describes modeling conformance test outcomes (full, partial, or non-compliance), uses of PROV-O, EARL, and VOCAB-DQV for quality metrics, and introduces qualified relationships between datasets, agents, and other resources.

Local Summary
The document details how to express the degree of conformance for datasets using controlled vocabularies and provenance modeling. It illustrates patterns for reporting full compliance, partial compliance (e.g., 50%), and non-conformance using `prov:Entity`, `prov:Activity`, `prov:Plan`, and EARL classes. It also covers qualified relations to assign specific roles (e.g., distributor, funder) to agents or other resources without creating an explosion of properties.

Key Claims
- Legal contexts like INSPIRE require specifying degrees of conformance beyond simple "conformant/not conformant".
- Conformance tests can be modeled using `prov:Entity` for results, `prov:Activity` for the test execution, and `prov:Plan` derived from standards like DWBP.
- Partial compliance (e.g., 50% of tests passed) is expressible via quality metrics in VOCAB-DQV.
- EARL provides specific classes (`earl:Assertion`, `earl:TestResult`) to describe testing activities, complementing PROV-O.
- Qualified relations use an additional resource (e.g., `prov:Attribution`, `dcat:Relationship`) to carry parameters like roles, avoiding property explosion.

Entities And Concepts
- **Degree of Conformance**: Represented by SKOS concepts (e.g., `ex:conformant`, `ex:notConformant`, `ex:nonConformant`) or INSPIRE vocabulary (`http://inspire.ec.europa.eu/metadata-codelist/DegreeOfConformity`).
- **Provenance Modeling**: Uses `prov:Entity` (results), `prov:Activity` (testing), `prov:Plan` (test suite derived from DWBP), and `prov:Association` to bind activities to plans/agents.
- **Quality Metrics**: Defined in VOCAB-DQV (`dqv:Metric`, `dqv:Dimension`, `dqv:Category`, `dqv:QualityMeasurement`) to measure compliance percentages.
- **EARL Testing**: Uses `earl:Assertion`, `earl:TestRequirement`, `earl:TestResult`, `earl:mode` (manual/automatic), and `earl:outcome`.
- **Qualified Relations**: Pattern using `prov:qualifiedAttribution` and `dcat:qualifiedRelation` with `dcat:hadRole` to specify roles like distributor, funder, or stereo-mate.
- **Vocabularies Referenced**: DWBP (Data on the Web Best Practices), ISO/IEC 25012 (quality dimensions), INSPIRE DoC, VOCAB-DQV, EARL, PROV-O, DCTERMS, CI_RoleCode, DS_AssociationTypeCodes.

Procedures And API Details
- **Modeling Conformance Test**:
  1. Define a `prov:Plan` (e.g., `ex:conformanceTest`) derived from a standard (e.g., DWBP).
  2. Create a `prov:Activity` (`ex:testingActivity`) using the plan.
  3. Generate an `earl:Assertion` or `prov:Entity` (`ex:testResult`) as output.
  4. Bind the activity to the plan via `prov:qualifiedAssociation` (or EARL's direct assertion).
- **Reporting Partial Compliance**:
  1. Define a `dqv:Metric` (e.g., `levelOfComplianceToDWBP`).
  2. Create a `dqv:QualityMeasurement` with a numeric value (e.g., "50"^^xsd:double) and unit (Percentage).
- **Reporting Non-Conformance**:
  1. Use `earl:Assertion` with `earl:result` containing an `earl:TestResult`.
  2. Set `dcterms:type` to `ex:nonConformant` or use INSPIRE's non-conformant URI.
  3. Include human-readable errors via `dcterms:description` (HTML/XML) and technical details in `earl:info`.
- **Assigning Agent Roles**:
  1. Use `prov:qualifiedAttribution` with a nested `prov:Attribution`.
  2. Link an agent (`prov:agent`) to the resource via `dcat:hadRole` inside the attribution.
- **Relating Datasets**:
  1. Use `dcat:qualifiedRelation` with a nested `dcat:Relationship`.
  2. Specify the target resource via `dcterms:relation` and the role via `dcat:hadRole`.

Nuance Or Contradictions
- **PROV-O Limitations**: PROV-O is activity-centric and lacks support for direct Entity-Entity relations (except "was derived from"), necessitating DCAT's new elements (`dcat:qualifiedRelation`, `dcat:Relationship`).
- **Role Property Domain**: In PROV-O, `prov:hadRole` relates to activities, not entities; thus, DCAT introduces `dcat:hadRole` on the association-class `prov:Attribution`.
- **Vocabulary Choices**: The text notes that while IRIs from non-normative codelists (e.g., CI_RoleCode via URNs) are used in examples, linked data dereferenceable and normative representations should be preferred when available.
- **Test Mode Ambiguity**: Example 51 notes uncertainty about whether a test was manual or automatic (`earl:mode earl:automatic.` with a comment "we do not know..."), highlighting the need for explicit metadata to clarify test execution modes.

Candidate Wiki Hints
- **Page: Modeling Conformance Tests in DCAT**
  - Focus: How to use PROV-O and EARL to model conformance testing activities, plans, and results.
  - Key content: Examples of `prov:Plan`, `prov:Activity`, `earl:Assertion`, and handling partial/non-conformance.
- **Page: Qualified Relations in DCAT**
  - Focus: Using `dcat:qualifiedRelation` and `prov:qualifiedAttribution` to assign roles (e.g., distributor, funder) without property explosion.
  - Key content: Patterns for relating datasets to agents and other resources with specific semantic roles.
- **Page: Quality Metrics for Data Compliance**
  - Focus: Using VOCAB-DQV to define metrics for compliance percentages and data quality dimensions.
  - Key content: Defining `dqv:Metric`, `dqv:QualityMeasurement`, and linking to ISO/IEC 25012 dimensions.

## chunk-12

---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
- Heading: 17. Security and Privacy Considerations
- Line Range: 5050–5087
- Scope: Guidelines for handling personal data within DCAT, integrity verification via checksums, and trust model requirements.

Local Summary
This section addresses how the Data Catalog Vocabulary (DCAT) handles datasets containing personal information and metadata that might reveal private details about creators or publishers. It outlines responsibilities for implementers to ensure secure storage and access control, noting that detailed web security mechanisms are out of scope. The text emphasizes using checksums from SPDX to verify data integrity, requiring separate delivery channels for checksums to prevent tampering. Ultimately, the trustworthiness of DCAT data relies on securing API endpoints and serving content over authoritative HTTPS origins.

Key Claims
- DCAT vocabulary supports datasets containing personal or private information.
- Metadata expressed with DCAT may itself contain personal/private information (e.g., resource creators, publishers).
- Implementers must take steps to ensure security and privacy considerations are addressed for sensitive data.
- Sensitive data and metadata must be stored securely and made available only to authorized parties per legal/functional requirements.
- Detailing how to secure web content and authenticate users is beyond the scope of DCAT.
- Checksums can serve as verification for datasets requiring assurances of integrity and authenticity (e.g., software vulnerabilities).
- DCAT borrows `spdx:Checksum` from SPDX to ensure integrity and authenticity of DCAT distributions.
- A checksum must be provided via a route separate from the data it sums.
- If a checksum is included in metadata with the data, a checksum for that metadata must also be provided separately to foil attackers.
- Integrity and authenticity of DCAT metadata depend on the trustworthiness of the source.
- Providers should address integrity/authenticity at application and transport levels (e.g., HTTPS origins).

Entities And Concepts
- DCAT (Data Catalog Vocabulary)
- Personal/private information
- Resource creators, publishers, agents
- Sensitive data/metadata
- Checksums (hashes), algorithms
- `spdx:Checksum` class
- Integrity and authenticity assurances
- API endpoints, download endpoints
- Authoritative HTTPS origins
- Application level security
- Transport level security

Procedures And API Details
- Store sensitive data securely; restrict access to authorized parties.
- Provide a checksum value (hash) and the generating algorithm for each resource in a distribution.
- Deliver checksums via a separate channel from the data they represent.
  - Example: A tarfile containing a distribution file and a metadata file with the distribution's checksum.
  - Requirement: Also provide a checksum for the metadata file itself if embedded with the data.
- Ensure integrity and authenticity of API and download endpoints.
- Make DCAT data/metadata files downloadable from authoritative HTTPS origins.

Nuance Or Contradictions
- No direct contradictions found; recommendations align with standard security practices (separation of data and verification, use of HTTPS).
- The distinction between verifying the data checksum and the metadata checksum is critical to prevent manipulation attacks.
- DCAT explicitly defers detailed web security implementation details to external standards or organizational policies.

Candidate Wiki Hints
- Security best practices for DCAT implementations
- Handling privacy in catalog metadata
- Integrity verification using SPDX checksums in DCAT distributions

## chunk-13

---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
This chunk covers Section 18, "Accessibility Considerations," followed by Appendix A (Acknowledgments), B (Alignment with Schema.org), and the beginning of C (Examples). It details accessibility mandates for non-text data, lists contributors to the DCAT vocabulary specification, maps DCAT concepts to Schema.org types, and provides RDF examples for loosely structured catalogs, dataset provenance, and linking datasets to publications.

Local Summary
The document section emphasizes enforcing alternative text for non-text resources within DCAT profiles to comply with accessibility guidelines like UNDERSTANDING-WCAG20. It acknowledges the working group members and external reviewers. Section B outlines the relationship between DCAT and Schema.org, noting that Google's Dataset Search relies on both. A detailed mapping table aligns DCAT elements (e.g., dcat:Dataset) with Schema.org types (e.g., sdo:Dataset). Section C begins with examples for loosely structured catalogs using dcterms:relation/dcterms:hasPart versus dcat:distribution, dataset provenance using the PROV ontology, and linking datasets to publications via dcterms:isReferencedBy.

Key Claims
- Data catalogs often contain non-text data; DCAT profiles should enforce alternative text for such resources to improve accessibility.
- Providing text alternatives for non-text content (large print, braille, speech, etc.) complies with UNDERSTANDING-WCAG20 guidelines.
- Schema.org types and properties are based on original DCAT work; Google's Dataset Search index relies on structured descriptions using both Schema.org and DCAT.
- A recommended mapping from revised DCAT to Schema.org version 3.4 uses predicates like rdfs:subClassOf, owl:equivalentClass, and sdo:domainIncludes.
- In loosely structured catalogs (e.g., legacy CKAN), datasets are often treated as a "bag of files" without distinguishing between distributions and other relationships.
- If the relationship nature is unknown, dcterms:relation or dcterms:hasPart should be used; if known and representing a proper representation, dcat:distribution should be used.
- Dataset provenance can be described using W3C Provenance Ontology elements like prov:wasGeneratedBy, prov:wasAttributedTo, and prov:wasDerivedFrom.
- Datasets are linked to publications (scholarly articles, reports) via the property dcterms:isReferencedBy.

Entities And Concepts
- **DCAT**: Vocabulary for describing data catalogs.
- **Schema.org**: Web vocabulary used by general-purpose search services; DCAT aligns with it for broader exposure.
- **UNDERSTANDING-WCAG20**: Accessibility guidelines referenced for text alternatives.
- **Prov-O (Provenance Ontology)**: Used to describe dataset provenance and business context.
- **Dublin Core (dcterms)**: Properties used in DCAT for metadata (e.g., title, description, creator).
- **RDF/Vocabularies**: Includes FOAF (foaf:Document), Geosciml ontology, OWL, RDF Schema.

Procedures And API Details
- **Mapping to Schema.org**: Use an available RDF file mapping DCAT revised elements to Schema.org 3.4 using specific predicates (rdfs:subClassOf, owl:equivalentClass, etc.).
- **Loosely Structured Catalogs**:
  - Use `dcterms:relation` or `dcterms:hasPart` when the nature of relationships between a dataset and component resources is unknown.
  - Use `dcat:distribution` when related resources are proper representations of the dataset.
- **Provenance Linking**:
  - Use `prov:wasGeneratedBy` to link a dataset to the project that generated it (formalized as a `prov:Activity`).
  - Use `prov:wasAttributedTo` for agents (sponsors, managers) not covered by standard Dublin Core creator/contributor/publisher roles.
  - Use `prov:wasDerivedFrom` for links to predecessor datasets.
- **Publication Linking**: Use `dcterms:isReferencedBy` to link a dataset to associated publications.

Nuance Or Contradictions
- The distinction between an abstract `Dataset` and a concrete `DataDownload` in DCAT maps to `dcat:Dataset / dcat:Distribution`, whereas Schema.org distinguishes `Dataset` from `DataDownload`.
- Legacy catalogs often lack the distinction between distributions (representations) and other relationships (documentation, schemas), requiring careful selection of DCAT properties to model this correctly.
- The mapping to Schema.org is described as non-normative in the text but provides a recommended alignment for interoperability with search indexes.

Candidate Wiki Hints
- **DCAT Accessibility Profile**: Guidelines for enforcing alternative text on non-text data resources within DCAT implementations.
- **DCAT vs. Schema.org Mapping**: A reference page detailing the axiomatized mapping between DCAT 3 and Schema.org 3.4 classes and properties.
- **Modeling Dataset Provenance**: Best practices for using PROV-O alongside DCAT to describe dataset origins, activities, and agents.
- **Linking Datasets to Publications**: Usage of `dcterms:isReferencedBy` to associate datasets with scholarly articles or reports.

## chunk-14

---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

## Chunk Context
**Heading:** 18. Accessibility Considerations
**Line Range:** 5502–5860
**Scope:** This chunk covers the introduction of accessibility considerations, followed by detailed examples of data services (C.4), compressed/packaged distributions (C.5), and a comprehensive change history detailing revisions from the Candidate Recommendation Snapshot through multiple public working drafts up to May 2021.

## Local Summary
The text introduces **Section 18**, noting its addition in changes since the second public working draft of 4 May 2021 (Issue #1358). It provides concrete RDF examples for describing data services using DCAT properties like `dcterms:type`, `dcat:endpointDescription`, and `dcat:endpointURL`. The section on distributions demonstrates how to model compressed files (GZIP) and packaged archives (TAR, ZIP) using `dcat:compressFormat` and `dcat:packageFormat`. Finally, the change history logs specific editorial fixes, definition revisions (e.g., replacing "item" with "resource"), and removal of properties in response to implementation feedback.

## Key Claims
- **Section 18 (Accessibility Considerations)** was added to the specification following Issue #1358 during updates since May 2021.
- Data services can be described using classifiers such as `dcterms:type`, `dcterms:conformsTo`, and `dcat:endpointDescription` to provide progressive detail about a service.
- The actual endpoint of a service is identified via the property **`dcat:endpointURL`**.
- Distributions support modeling for compressed files (e.g., GZIP) using **`dcat:compressFormat`** and packaged archives (e.g., TAR) using **`dcat:packageFormat`**.
- Property **`dcat:resource`** was introduced to link a `dcat:Catalog` to a `dcat:Resource`, replacing the usage of `dcterms:hasPart` from DCAT 2.
- The property **`dcat:inCatalog`** serves as the inverse of `dcat:resource`.

## Entities And Concepts
- **Accessibility Considerations**: A new section addressing accessibility in data catalogs.
- **Data Services**: Described using DCAT, including examples for the European Environment Agency (EEA) and Geoscience Australia.
- **Distributions**: Models for downloadable files, compressed archives (GZIP), and packaged files (TAR).
- **DCAT Properties**:
  - `dcat:endpointURL`: Identifies the actual endpoint of a service.
  - `dcat:compressFormat`: Specifies compression types (e.g., GZIP).
  - `dcat:packageFormat`: Specifies packaging formats (e.g., TAR, ZIP).
  - `dcat:resource`: Links catalog to resource instances.
  - `dcat:inCatalog`: Inverse property for resources.
- **INSPIRE Classification**: Referenced for spatial data service types (`discovery`, `download`, `view`).

## Procedures And API Details
- **Modeling a Data Service**: Use `rdf:type dcat:DataService` and set `dcterms:type` (e.g., INSPIRE codes) and `dcat:endpointURL`.
- **Modeling Distributions**:
  - For compressed files: Set `dcat:downloadURL` and add `dcat:compressFormat`.
  - For packaged files: Set `dcat:packageFormat`.
  - Combined compression/packaging: Use both properties.
- **Service Types**: Examples use `dcterms:type` values from the INSPIRE Spatial Data Service Category list (e.g., `infoCatalogueService`, `download`, `view`).

## Nuance Or Contradictions
- **Terminology Shift**: The term "item" was replaced with "resource" when referring to instances of `dcat:Resource` to ensure consistent terminology throughout the document.
- **Property Replacement**: `dcterms:hasPart` (from DCAT 2) is no longer recommended for linking a catalog to a resource; `dcat:resource` is now defined for this purpose.
- **Inverse Properties**: Section 7 defines inverse properties like `dcat:inCatalog`, while older properties like `dcat:isVersionOf` and `dcterms:isReplacedBy` were removed in the May 2021 draft.

## Candidate Wiki Hints
- **Page: DCAT Accessibility**
  - *Content*: Summarize Section 18, explaining why accessibility considerations were added and how they integrate into catalog metadata standards.
- **Page: Data Service Endpoints**
  - *Content*: Document the use of `dcat:endpointURL` and related descriptors (`dcterms:conformsTo`, `dcat:endpointDescription`) for defining service endpoints in DCAT.
- **Page: Distribution Packaging Formats**
  - *Content*: Explain how to model downloadable, compressed (GZIP), and packaged (TAR/ZIP) distributions using specific DCAT properties.

## chunk-15

---
title: Chunk 15 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

## Chunk Context

**Heading:** 18. Accessibility Considerations
**Line Range:** 5862–6377
**Source Document:** `raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md`

## Local Summary

This chunk documents the evolution of the DCAT vocabulary since the W3C Recommendation of 4 February 2020. It details significant revisions in versioning strategies, the introduction of dataset series as first-class entities, and the integration of SPDX checksum properties. The text also lists various editorial updates, namespace corrections (e.g., `dcterms:`), and clarifications regarding resource types. The latter half of the chunk provides an extensive list of normative and informative references used to define the DCAT vocabulary.

## Key Claims

- Versioning specifications have shifted to focus on version chains derived from resource revisions, utilizing the [PAV] approach for hierarchies.
- Support for specifying backward compatibility using `owl:backwardCompatibleWith` and `owl:incompatibleWith` has been dropped.
- The `dcat:DatasetSeries` class was introduced to treat dataset series as first-class citizens, alongside new properties like `dcat:inSeries`.
- Properties `dcat:first`, `dcat:prev`, `dcat:next`, and `dcat:last` were added to the `Cataloged Resource` class.
- The property `spdx:checksum` and the class `spdx:Checksum` (with properties `spdx:algorithm` and `spdx:checksumValue`) were added to distributions.
- The range of `locn:geometry` was revised to align with its definition in [LOCN], allowing usage with geometry literals or classes.
- The document has been updated to clarify that catalog resources are not limited to datasets and data services.

## Entities And Concepts

- **Versioning Strategy:** Focuses on versions derived from resource revision; utilizes [PAV] approach for chains/hierarchies (previous, next, current, last).
- **Dataset Series:** New class `dcat:DatasetSeries`; properties `dcat:inSeries`, `dcat:first`, `dcat:prev`, `dcat:next`, `dcat:last`.
- **Checksums:** Class `spdx:Checksum` with properties `spdx:algorithm` and `spdx:checksumValue` added to distributions.
- **Geography:** Property `locn:geometry` revised for alignment with [LOCN].
- **Namespace Corrections:** Replacement of `[DCTERMS]` prefix `dct:` with `dcterms:`; consistent use of "URI" vs "IRI".
- **References:** Includes normative references (e.g., [DC11], [DCTERMS], [OWL-TIME], [SPDX]) and informative references (e.g., [GeoDCAT-AP], [FAIR], [ODRL-VOCAB]).

## Procedures And API Details

- **Namespace Updates:** Replace `dct:` with `dcterms:` throughout the document.
- **Property Additions:**
  - Add `spdx:checksum` to distribution class.
  - Define class `spdx:Checksum` and add properties `spdx:algorithm`, `spdx:checksumValue`.
  - Add `dcat:inSeries` to dataset class.
  - Add `dcat:first`, `dcat:prev`, `dcat:next`, `dcat:last` to cataloged resource class.
- **Removals:** Remove section on version types; remove support for `owl:backwardCompatibleWith` and `owl:incompatibleWith`.

## Nuance Or Contradictions

- The chunk notes that previous sections included only editorial changes, while specific substantive changes (versioning, dataset series) are highlighted separately.
- There is a distinction between normative references (binding specifications) and informative references (guidelines or related standards).
- The text explicitly mentions dropping support for certain OWL properties previously included in version information, indicating a deliberate simplification or standardization shift.

## Candidate Wiki Hints

- **Dataset Series:** Create a page explaining the concept of `dcat:DatasetSeries` and its role in linking datasets within a catalog.
- **Versioning Guidelines:** Document the new approach to version chains using [PAV] and the removal of backward compatibility properties.
- **Checksums:** Explain the integration of SPDX checksums for distribution integrity verification.
- **Reference Catalog:** Maintain a page listing normative and informative references relevant to DCAT extensions (e.g., GeoDCAT, ODRL).

## chunk-16

---
title: Chunk 16 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

## Chunk Context

This chunk (lines 6379–6468) concludes the source document, focusing on **Section 18: Accessibility Considerations**. It serves as an appendix or reference list containing a curated collection of related W3C vocabularies and standards relevant to data catalogs, semantic interoperability, and accessibility. The text lists multiple versions of DCAT (from Version 2 Recommendation in Feb 2020 to Version 3 Candidate Recommendation in Jan 2024), alongside other key ontologies like VOID, SSN, DQV, ORG, and external standards such as WFS, WMS, and WSDL.

## Local Summary

The section provides a bibliography of normative and informative documents that support the Data Catalog Vocabulary (DCAT) ecosystem. It emphasizes the evolution of DCAT through various working drafts leading to candidate recommendation status and includes references to foundational vocabularies for linked datasets (VOID), quality assessment (DQV), organizations (ORG), sensor networks (SSN), geospatial data (W3C-BASIC-GEO), and service interfaces (WFS, WMS). Accessibility considerations are implied through the inclusion of standards that facilitate structured description and discovery, which underpins accessible data access.

## Key Claims

- DCAT has evolved from a 2020 Recommendation (Version 2) to a Candidate Recommendation (Version 3) as of January 2024.
- The W3C maintains a suite of complementary vocabularies including VOID for describing linked datasets and DQV for data quality.
- External standards like WFS, WMS, and WSDL are referenced to ensure interoperability with existing geospatial and web service infrastructures.
- Accessibility considerations are addressed by incorporating structured metadata practices that enable consistent and machine-readable descriptions of data resources.

## Entities And Concepts

- **DCAT**: Data Catalog Vocabulary; versions 2 (Recommendation) and 3 (Candidate Recommendation).
- **VOID**: Vocabulary for describing Linked Datasets.
- **DQV**: Data Quality Vocabulary.
- **ORG**: Organization Ontology.
- **SSN**: Semantic Sensor Network Ontology.
- **W3C-BASIC-GEO**: Basic Geo vocabulary using WGS84 lat/long coordinates.
- **WFS**: Web Feature Service 2.0 Interface Standard (OGC).
- **WMS**: Web Map Service Implementation Specification (OGC).
- **WSDL**: Web Services Description Language Version 2.0.
- **Accessibility Considerations**: Implicitly addressed through structured metadata and interoperable standards enabling inclusive data access.

## Procedures And API Details

No specific procedures or API details are described in this chunk. The text functions as a reference list with versioned URLs for various W3C recommendations, working drafts, and OGC interface standards. Each entry includes authors, publication dates, status (e.g., Recommendation, Working Draft, Candidate Recommendation), and canonical URLs.

## Nuance Or Contradictions

The document reflects the iterative nature of W3C standards development, particularly in DCAT, where multiple versions are listed with changing statuses—from Working Draft to Candidate Recommendation. There is no explicit contradiction, but the progression suggests active refinement and community feedback integration prior to final recommendation. The inclusion of non-W3C standards (e.g., OGC WFS/WMS) highlights cross-ecosystem interoperability goals without implying endorsement by W3C beyond citation.

## Candidate Wiki Hints

- **Page: DCAT Version History** – Track the evolution of DCAT from 2020 Recommendation to 2024 Candidate Recommendation with links to each version.
- **Page: Related Vocabularies for Data Catalogs** – Curate VOID, DQV, ORG, SSN as complementary ontologies used alongside DCAT.
- **Page: Interoperability Standards in Data Catalogs** – Include WFS, WMS, and WSDL as external standards enabling service discovery and geospatial integration.
- **Page: Accessibility in Linked Data Descriptions** – Explore how structured metadata supports accessible data access and discoverability.

