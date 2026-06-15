---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---

## Chunk Context
This chunk covers sections 11.1.2 through 12.1 of the DCAT vocabulary specification, detailing versioning relationships (replacement), version information properties, resource life-cycle management, complementary versioning approaches (OWL/DCTERMS/PROV-O), and the concept of dataset series including hierarchical linking.

## Local Summary
The text explains how to model versions that replace others using `dcterms:replaces`. It defines properties for version identifiers (`dcat:version`) and release dates (`dcterms:issued`). The resource life-cycle is treated as orthogonal to versioning, utilizing statuses like "deprecated" or "withdrawn". Finally, it introduces `dcat:DatasetSeries` for grouping related datasets (e.g., yearly budgets) using properties like `dcat:inSeries`, `dcat:first`, and `dcat:last`.

## Key Claims
- **Replacement Relationships:** DCAT uses `dcterms:replaces` and its inverse `dcterms:isReplacedBy` to indicate if a version supersedes another. This does not automatically form a strict linear chain; a version may replace any prior one, not just the immediate predecessor.
- **Version Information:** Key properties include `dcat:version` (identifier), `dcterms:issued` (release date), and `adms:versionNotes` (description of changes). DCAT defers to [DWBP] Best Practice 7 for naming conventions.
- **Life-Cycle Orthogonality:** A resource's life-cycle status (e.g., development, approved, withdrawn) is distinct from its versioning history. Statuses are managed via `adms:status` and various date properties (`dcterms:created`, `dcterms:issued`, etc.).
- **Complementary Versioning:** DCAT properties coexist with OWL (`owl:versionInfo`) and PROV-O (`prov:generalizationOf`). These have different scopes (e.g., OWL for ontologies, DCTERMS for broad editions) and are not strictly equivalent.
- **Dataset Series:** DCAT treats dataset series as first-class citizens via the `dcat:DatasetSeries` class (subclass of `dcat:Dataset`). Datasets link to a series via `dcat:inSeries`. Series can be hierarchical, with relationships for first/last members (`dcat:first`, `dcat:last`) and sequential neighbors (`dcat:prev`, `dcat:next`).

## Entities And Concepts
- **dcterms:replaces / dcterms:isReplacedBy:** Properties defining replacement relationships between versions.
- **dcat:version:** Property for the version name or identifier.
- **dcterms:issued:** Property for the release date of a version.
- **adms:versionNotes:** Property for textual descriptions of changes and compatibility issues.
- **adms:status:** Property for specifying life-cycle statuses (e.g., deprecated, withdrawn).
- **dcat:DatasetSeries:** A new class in DCAT representing a collection of interrelated datasets.
- **dcat:inSeries:** Property linking individual datasets to a dataset series.
- **dcat:first / dcat:last:** Properties identifying the first and last members of a dataset series.
- **dcat:prev / dcat:next:** Properties linking datasets sequentially within a series.
- **ISO-19135 / ISO-19115:** Standards providing example life-cycle status vocabularies (e.g., accepted, experimental, stable).

## Procedures And API Details
- **Specifying Replaced Versions:** Use `dcterms:replaces` to point from the new version to the old one. The inverse property `dcterms:isReplacedBy` can be used on the older version for back-linking.
- **Linking Dataset Series:** Create a resource typed as `dcat:DatasetSeries`. Link child datasets using `dcat:inSeries`. Use `dcat:first`, `dcat:last`, `dcat:prev`, and `dcat:next` to define the structure of the series over time.
- **Combining Series and Versions:** A dataset in a series can have its own versions. Link versions using `dcat:hasVersion` (listing all) and `dcat:hasCurrentVersion` (pointing to the active one).

## Nuance Or Contradictions
- **Non-Equivalence of Properties:** Correspondence between DCAT properties and those in OWL or PROV-O does not imply equivalence. For instance, `prov:wasRevisionOf` is semantically similar to `dcat:previousVersion` but is not intended for building version chains. Similarly, `owl:versionIRI` differs from `dcat:hasCurrentVersion` in scope and usage.
- **Life-Cycle vs. Versioning:** While related, life-cycle evolution (creation, approval, publication) does not always result in a new version, and creating a new version does not necessarily change the life-cycle status.

## Candidate Wiki Hints
- **Page: DCAT Versioning Relationships** (Concept: Explaining `dcterms:replaces` vs. linear chains).
- **Page: Dataset Series in DCAT** (Concept: Using `dcat:DatasetSeries` and hierarchical linking properties).
