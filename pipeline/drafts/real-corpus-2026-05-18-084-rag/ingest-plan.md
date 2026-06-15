---
kind: topic
sources: [raw/web/corpus-2026-05-18/084-retrieval-augmented-generation.md]
status: draft
---

# Ingest Plan

## Source Summary
- The document describes Retrieval-Augmented Generation (RAG), a technique combining parametric and non-parametric memory for language generation.
- It details a fine-tuning recipe using pre-trained seq2seq models paired with a dense vector index of Wikipedia accessed via a neural retriever.
- The paper demonstrates state-of-the-art performance on open domain QA tasks and improved factual diversity in language generation compared to parametric-only baselines.
- Accepted at NeurIPS 2020, the work addresses limitations in large pre-trained models regarding knowledge access and provenance.

## Candidate Wiki Pages
- wiki/sources/rag-arxiv-2005-11401.md — Primary source document containing the full paper abstract, metadata, and author details.
- wiki/concepts/retrieval-augmented-generation.md — Conceptual explanation of RAG mechanics including parametric vs non-parametric memory usage.
- wiki/topics/knowledge-intensive-nlp-tasks.md — Discussion on specific application areas where RAG outperforms standard models.

## Suggested Links
- https://arxiv.org/abs/2005.11401

## Review Checklist
- [ ] Verify citation accuracy against arXiv metadata
- [ ] Ensure candidate pages follow wiki directory structure constraints
- [ ] Confirm no nested directories are created for candidate files
