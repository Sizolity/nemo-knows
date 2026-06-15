---
title: Knowledge Intensive Nlp Tasks
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/084-retrieval-augmented-generation.md
confidence: medium
---

# Knowledge Intensive Nlp Tasks

Large pre-trained language models store factual knowledge within their parameters but often struggle with accessing and manipulating that knowledge precisely. This limitation leads to performance deficits on **[[knowledge-intensive-nlp-tasks]]** compared to task-specific architectures. Standard approaches fail to provide provenance for model decisions or easily update world knowledge.

## Retrieval-Augmented Generation (RAG)

Retrieval-Augmented Generation (**[[retrieval-augmented-generation]]**) is a general-purpose fine-tuning recipe designed to address these specific challenges. It combines pre-trained parametric memory with non-parametric explicit memory. In this architecture:
- **Parametric Memory**: Consists of a pre-trained sequence-to-sequence model.
- **Non-parametric Memory**: A dense vector index of Wikipedia accessed via a pre-trained neural retriever.

This approach enables differentiable access to explicit external knowledge, allowing models to overcome the limited ability of standard large models to precisely manipulate information.

## Performance and Quality

Evaluation studies show that RAG models outperform parametric-only seq2seq models and task-specific retrieve-and-extract architectures on open domain QA tasks. Specifically:
- **Performance**: RAG sets the state-of-the-art on three open domain QA tasks.
- **Language Generation**: Models produce more specific, diverse, and factual language than state-of-the-art parametric baselines.

## Research Context

The methodology was explored in a paper accepted at NeurIPS 2020. The research highlighted that while large models have vast internal knowledge, they lack the mechanism to access it precisely when needed for complex reasoning or fact retrieval. RAG provides a solution by conditioning on retrieved passages either across the whole generated sequence or per token.
