---
title: Retrieval Augmented Generation
kind: concept
sources:
  - source.md
  - raw/web/corpus-2026-05-18/084-retrieval-augmented-generation.md
confidence: medium
---

# Retrieval Augmented Generation

Retrieval-Augmented Generation (RAG) is a general-purpose fine-tuning recipe for language models that combines pre-trained parametric memory with non-parametric explicit memory. In this architecture, the parametric memory consists of a pre-trained sequence-to-sequence model, while the non-parametric memory is a dense vector index accessed via a pre-trained neural retriever.

This approach addresses limitations in large pre-trained language models regarding their ability to access and precisely manipulate knowledge, particularly on knowledge-intensive tasks. RAG overcomes these issues by enabling differentiable access to explicit non-parametric memory.

## Performance and Quality

RAG models set the state-of-the-art on open domain QA tasks, outperforming both parametric seq2seq models and task-specific retrieve-and-extract architectures. On language generation tasks, RAG models produce more specific, diverse, and factual language compared to state-of-the-art parametric baselines.

## Advantages

By combining pre-trained parametric and non-parametric memory, RAG allows models to overcome the limited ability of standard large models to access and manipulate knowledge precisely. Additionally, this approach provides a solution for providing provenance for model decisions and updating world knowledge, which remain open research problems for standard pre-trained models.
