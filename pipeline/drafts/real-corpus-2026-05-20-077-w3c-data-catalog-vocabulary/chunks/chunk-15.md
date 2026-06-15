---
title: Chunk 15 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/077-w3c-data-catalog-vocabulary.md
confidence: medium
---
Chunk Context
----------------
This chunk covers **Section 18. Accessibility Considerations** within the DCAT 2 Vocabulary document, though the provided text primarily details revision history (changes to versioning, dataset series, SPDX checksums, geometry usage, license examples, namespace prefixes, and URI/IRI terminology). It also lists changes since the W3C Recommendation of February 4, 2020, and provides a comprehensive list of normative and informative references used by the vocabulary.

Local Summary
----------------
The document has been updated to clarify versioning strategies using the [PAV] approach and to remove backward compatibility property specifications (`owl:backwardCompatibleWith`, `owl:incompatibleWith`). A new class `dcat:DatasetSeries` was introduced to treat dataset series as first-class citizens, alongside properties for linking series members (`dcat:inSeries`) and defining sequence order (`dcat:first`, `dcat:prev`, `dcat:next`, `dcat:last`). The vocabulary now includes SPDX checksums for distributions and updated geometry usage aligned with [LOCN]. Editorial changes include fixing inconsistent "URI"/"IRI" usage, replacing the `dct:` prefix with `dcterms:`, and clarifying that catalog resources are not limited to datasets. Changes since the 2020 Recommendation emphasize using specific subrelations over `dcterms:relation` and adding draft guidelines for versioning and dataset series.

Key Claims
----------------
- Versioning is now focused on versions derived from resource revision, following the [PAV] approach for version chains and hierarchies (previous, next, current, last).
- Support for specifying backward/incompatible compatibility between versions using `owl:backwardCompatibleWith` and `owl:incompatibleWith` has been dropped.
- Dataset series are first-class citizens; a new class `dcat:DatasetSeries` defines the relationship between datasets in a series.
- Properties `dcat:first`, `dcat:prev`, `dcat:next`, and `dcat:last` define the sequence of resources within a series.
- Property `spdx:checksum` was added to the Distribution class, introducing `spdx:Checksum` with properties `spdx:algorithm` and `spdx:checksumValue`.
- The range of `locn:geometry` has been revised to align with [LOCN], allowing usage with geometry literals or classes.
- Examples in the License section have been updated to address issues #676 and #1333.
- The namespace prefix for DCMI Metadata Terms was changed from `dct:` to `dcterms:` throughout the document.
- Inconsistent use of "URI" and "IRI" has been fixed to ensure terminology consistency.

Entities And Concepts
-----------------------
- **DCAT 2 Vocabulary**: The core data catalog vocabulary being maintained.
- **Versioning ([PAV])**: Approach for specifying version chains/hierarchies using `previous`, `next`, `current`, `last`.
- **Dataset Series**: A new concept where datasets are members of a series, represented by class `dcat:DatasetSeries`.
- **SPDX Checksum**: Property added to Distribution to specify checksums using the SPDX vocabulary.
- **LOCN Geometry**: Alignment with the ISA Programme Location Core Vocabulary for geometry representation.
- **W3C Recommendation (2020)**: The baseline document version from February 4, 2020.

Procedures And API Details
-----------------------------
- **Namespace Update**: Replace `dct:` with `dcterms:` in all references to DCMI Metadata Terms.
- **Geometry Usage**: Use `locn:geometry` with either geometry literals or classes, aligned with [LOCN].
- **Version Specification**: Use properties from the PAV ontology (e.g., `pav:previousVersion`) instead of OWL compatibility properties.
- **Checksum Specification**: Add `spdx:checksum` to a Distribution resource using `spdx:algorithm` and `spdx:checksumValue`.

Nuance Or Contradictions
---------------------------
- The chunk title mentions "Accessibility Considerations" (Section 18), but the content focuses on versioning, dataset series, and references. This suggests the provided text might be a fragment of a larger document where Section 18 follows or precedes these revision notes.
- The removal of `owl:backwardCompatibleWith` implies a shift away from explicitly modeling compatibility states in favor of the [PAV] approach for version relationships.

Candidate Wiki Hints
-----------------------
- **Dataset Series**: Create a page explaining `dcat:DatasetSeries`, `dcat:inSeries`, and sequence properties (`first`, `prev`, `next`, `last`).
- **Versioning Best Practices**: Summarize the shift to [PAV] for version management and the removal of OWL compatibility properties.
- **SPDX Integration**: Document how to use `spdx:checksum` within DCAT distributions.
- **Geometry Alignment**: Explain the updated usage of `locn:geometry`.
