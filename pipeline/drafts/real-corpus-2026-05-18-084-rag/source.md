---
title: Retrieval-Augmented Generation (RAG)
kind: source
sources:
  - raw/web/corpus-2026-05-18/084-retrieval-augmented-generation.md
confidence: medium
---

# Retrieval-Augmented Generation (RAG)

## What It Is
Retrieval-Augmented Generation (RAG) is a general-purpose fine-tuning recipe for language models that combines pre-trained parametric memory with non-parametric explicit memory. In this architecture, the parametric memory consists of a pre-trained sequence-to-sequence model, while the non-parametric memory is a dense vector index of Wikipedia accessed via a pre-trained neural retriever. This approach addresses limitations in large pre-trained language models regarding their ability to access and precisely manipulate knowledge, particularly on knowledge-intensive tasks.

## Summary
Large pre-trained language models store factual knowledge within their parameters but often struggle with accessing and manipulating that knowledge precisely, leading to performance deficits on knowledge-intensive tasks compared to task-specific architectures. RAG overcomes these issues by enabling differentiable access to explicit non-parametric memory. The methodology was explored in a paper accepted at NeurIPS 2020, which introduced models conditioning on retrieved passages either across the whole generated sequence or per token. Evaluation showed that RAG models outperform parametric-only seq2seq models and task-specific retrieve-and-extract architectures on open domain QA tasks. Furthermore, for language generation tasks, RAG models produce more specific, diverse, and factual language than state-of-the-art parametric baselines.

## Key Claims
- **Performance Improvement**: RAG models set the state-of-the-art on three open domain QA tasks, outperforming both parametric seq2seq models and task-specific retrieve-and-extract architectures.
- **Language Quality**: On language generation tasks, RAG models generate more specific, diverse, and factual language compared to a state-of-the-art parametric-only seq2seq baseline.
- **Architectural Advantage**: Combining pre-trained parametric and non-parametric memory allows models to overcome the limited ability of standard large models to access and manipulate knowledge precisely.
- **Research Problems Addressed**: The approach provides a solution for providing provenance for model decisions and updating world knowledge, which remain open research problems for standard pre-trained models.

## Suggested Links
- https://arxiv.org/abs/2005.11401
