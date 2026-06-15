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
