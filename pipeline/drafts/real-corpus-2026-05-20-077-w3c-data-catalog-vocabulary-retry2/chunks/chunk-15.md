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
