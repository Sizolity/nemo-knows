---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
This chunk details the vocabulary properties for cataloging resources within DCAT specifications (DCAT 1, 2, and 3). It covers metadata attributes such as language, publisher, identifiers, themes, types, relations (generic and qualified), keywords, landing pages, rights/licenses, parts/policies, citations, versioning chains, lifecycle status, and series ordering. The text includes specific notes on domain relaxations from DCAT 1 to DCAT 2, the introduction of new properties in DCAT 3 (e.g., version chains), and distinctions between generic relations and specific ones like distribution or hasPart.

Local Summary
The section defines RDF properties for describing various facets of a cataloged resource. Key focus areas include linguistic identification using ISO 639 codes, entity attribution (publisher/creator) via `dcterms:publisher` and `prov:qualifiedAttribution`, categorization via `dcat:theme` and `dcterms:type`, and complex relationship modeling using `dcterms:relation` and the qualified relation pattern (`dcat:qualifiedRelation`). The chunk also addresses resource lifecycle management through properties for versions (`dcat:hasVersion`, `dcat:previousVersion`), replacement chains (`dcterms:replaces`), and series membership (`dcat:first`).

Key Claims
- Language identification should prioritize ISO 639-1 codes; if unavailable, ISO 639-2 is recommended. BCP47 tags are also noted as evolving standards for natural language in linked data.
- The domain of `dcat:theme` was relaxed in DCAT 2 to allow use beyond just datasets, unlike DCAT 1 where it was limited to `dcat:Dataset`.
- `dcterms:type` is a sub-property of `dc:type` and should draw from controlled vocabularies like DCMI Type or ISO 19115, though not all terms (e.g., Event, PhysicalObject) are strictly suitable for datasets.
- Generic relationships between resources should use `dcterms:relation`, but specific semantic links (e.g., distribution, isVersionOf) should take precedence to avoid over-generalization.
- DCAT 3 introduces specific properties for versioning lineages (`dcat:hasVersion`, `dcat:previousVersion`) and lifecycle status (`adms:status`).
- The qualified relation pattern (`dcat:qualifiedRelation`, `prov:qualifiedAttribution`) is used when the relationship nature is known but does not match standard DCTERMS or PROV-O properties.

Entities And Concepts
- **Vocabulary Standards**: DCAT 1, DCAT 2, DCAT 3, ISO 639-1, ISO 639-2, BCP47, DCMI Type, ISO 19115-1, DataCite.
- **Classes and Types**: `dcterms:LinguisticSystem`, `foaf:Agent`, `rdfs:Literal`, `skos:Concept`, `owl:ObjectProperty`.
- **Properties**: `dcterms:language`, `dcterms:publisher`, `dcterms:identifier`, `dcat:theme`, `dcterms:type`, `dcterms:relation`, `dcat:qualifiedRelation`, `dcat:keyword`, `dcat:landingPage`, `prov:qualifiedAttribution`, `dcterms:license`, `dcterms:rights`, `odrl:hasPolicy`, `dcterms:hasPart`, `dcat:isReferencedBy`, `dcat:previousVersion`, `dcat:hasVersion`, `dcat:currentVersion` (as `dcat:hasCurrentVersion`), `dcterms:replaces`, `dcat:version`, `adms:versionNotes`, `adms:status`, `dcat:first`.
- **Concepts**: Qualified relation pattern, version chains, resource lineage, dataset series.

Procedures And API Details
- **Language Usage**: If a resource is available in multiple languages, repeat the property. Values for catalog members override catalog-level values if conflicting. For separate language representations, define distinct `dcat:Distribution` instances.
- **Theme Organization**: Themes should be organized within a `skos:ConceptScheme`, `skos:Collection`, or similar ontology describing categories and relations.
- **Identifier Usage**: The identifier is a text string assigned for unambiguous reference; it may be part of the resource IRI but should still be explicitly represented.
- **Versioning Strategy**: Use `dcat:hasVersion` to link an abstract resource to versioned snapshots. Use `dcat:previousVersion` and `dcat:hasCurrentVersion` to specify lineage steps. Use `dcterms:replaces` for superseded resources.
- **Qualified Relations**: When linking to another resource where the relationship is known but not covered by standard properties, use `dcat:qualifiedRelation`. This entails that the context resource is a member of `prov:Entity`.

Nuance Or Contradictions
- DCAT 1 vs. DCAT 2: The domain of `dcat:theme` and `dcat:keyword` was restricted in DCAT 1 to `dcat:Dataset`, which limited their utility. DCAT 2 relaxed these domains to allow broader application.
- Property Semantics: `dcterms:type` is a sub-property of `dc:type`, but its range in DCAT 3 is adjusted to `rdfs:Class` for formal consistency, preventing automatic inference of objects as `skos:Concept`.
- Versioning Scope: `dcat:hasVersion` is more specific than `dcterms:hasVersion`; the former is limited to revisions during a resource's lifecycle (e.g., dataset updates), while the latter includes broader notions like editions or adaptations.

Candidate Wiki Hints
- **Data Catalog Vocabulary**: A structured set of RDF properties for describing resources in data catalogs, evolving from DCAT 1 through DCAT 3.
- **Qualified Relations**: A pattern using `prov:qualifiedInfluence` and `dcat:Relationship` to describe specific relationships that do not fit standard DCTERMS or PROV-O properties.
- **Versioning in DCAT**: Properties for managing resource versions, including lineage chains (`hasVersion`, `previousVersion`) and lifecycle status.
- **Dataset Series**: Use of `dcat:first` and related properties to define ordered collections or series of resources.
