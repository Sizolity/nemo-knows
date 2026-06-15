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
