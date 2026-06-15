# Curated Web Corpus for Real Nemo-Knows Testing

Collected: 2026-05-18

Purpose: provide 120 stable, real-world web sources for testing `nemo-knows`
ingest, query, contradiction handling, cross-linking, and long-document
summarisation workflows.

Selection policy:
- Prefer official documentation, standards, open-access papers, public-domain
  literature, and public-sector resources.
- Store URLs and short notes instead of copying copyrighted page bodies.
- Use canonical landing pages where possible, so fetchers can retrieve the
  current version during tests.
- Mix technical, scientific, policy, and narrative sources to exercise different
  wiki page types.

Recommended test usage:
- Smoke test: ingest 5-10 sources across different categories.
- Breadth test: ingest 25-40 sources and run a lint pass for missing concepts.
- Scale test: ingest all 120 sources in batches of 10-15, updating
  `wiki/index.md` and `wiki/log.md` after each batch.

## Sources

| # | Category | Title | URL | Test value |
|---:|---|---|---|---|
| 1 | Git | Pro Git Book | https://git-scm.com/book/en/v2 | Broad technical source with many entities and concepts. |
| 2 | Git | Branches in a Nutshell | https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell | Existing corpus overlap for regression checks. |
| 3 | Git | Basic Branching and Merging | https://git-scm.com/book/en/v2/Git-Branching-Basic-Branching-and-Merging | Workflow-heavy source for procedural summaries. |
| 4 | Git | Rebasing | https://git-scm.com/book/en/v2/Git-Branching-Rebasing | Good for contrasting merge and rebase concepts. |
| 5 | Git | Remote Branches | https://git-scm.com/book/en/v2/Git-Branching-Remote-Branches | Tests distributed-system terminology. |
| 6 | Git | Git Hooks | https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks | Tests event-driven automation concepts. |
| 7 | Git | Git Internals - Plumbing and Porcelain | https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain | Deep implementation source for concept pages. |
| 8 | Git | Git Objects | https://git-scm.com/book/en/v2/Git-Internals-Git-Objects | Tests object-model extraction. |
| 9 | Git | Git References | https://git-scm.com/book/en/v2/Git-Internals-Git-References | Good for linking branches, refs, HEAD, and commits. |
| 10 | Git | githooks Reference | https://git-scm.com/docs/githooks | Reference-style source with many hook names. |
| 11 | SQLite | Write-Ahead Logging | https://www.sqlite.org/wal.html | Existing corpus overlap for regression checks. |
| 12 | SQLite | File Locking and Concurrency | https://www.sqlite.org/lockingv3.html | Tests concurrency and transaction concepts. |
| 13 | SQLite | Query Planning | https://www.sqlite.org/queryplanner.html | Dense technical explanation with examples. |
| 14 | SQLite | Query Optimizer Overview | https://www.sqlite.org/optoverview.html | Tests long technical source summarisation. |
| 15 | SQLite | Virtual Tables | https://www.sqlite.org/vtab.html | Tests extension mechanism extraction. |
| 16 | SQLite | FTS5 Extension | https://www.sqlite.org/fts5.html | Search-related source for internal dogfooding. |
| 17 | SQLite | JSON Functions and Operators | https://www.sqlite.org/json1.html | Tests structured-data concept extraction. |
| 18 | SQLite | Transactions | https://www.sqlite.org/lang_transaction.html | Compact reference source for transaction semantics. |
| 19 | SQLite | Isolation in SQLite | https://www.sqlite.org/isolation.html | Good for contradiction and nuance handling. |
| 20 | SQLite | Database File Format | https://www.sqlite.org/fileformat2.html | Tests low-level architecture summarisation. |
| 21 | Python | Python Tutorial | https://docs.python.org/3/tutorial/index.html | General programming tutorial with many concepts. |
| 22 | Python | Data Structures | https://docs.python.org/3/tutorial/datastructures.html | Tests extraction of lists, dicts, sets, and tuples. |
| 23 | Python | Modules | https://docs.python.org/3/tutorial/modules.html | Tests namespace and import concepts. |
| 24 | Python | Errors and Exceptions | https://docs.python.org/3/tutorial/errors.html | Tests control-flow and error-handling synthesis. |
| 25 | Python | Classes | https://docs.python.org/3/tutorial/classes.html | Tests object-oriented concept extraction. |
| 26 | Python | asyncio | https://docs.python.org/3/library/asyncio.html | Tests asynchronous programming concepts. |
| 27 | Python | typing | https://docs.python.org/3/library/typing.html | Long reference source with evolving semantics. |
| 28 | Python | pathlib | https://docs.python.org/3/library/pathlib.html | Practical API reference for path abstractions. |
| 29 | Python | packaging tutorial | https://packaging.python.org/en/latest/tutorials/packaging-projects/ | Tests procedural packaging workflow. |
| 30 | Python | pyproject.toml Specification | https://packaging.python.org/en/latest/specifications/pyproject-toml/ | Tests standards-like Python packaging content. |
| 31 | Go | Effective Go | https://go.dev/doc/effective_go | Style and language idiom source. |
| 32 | Go | Go Modules Reference | https://go.dev/ref/mod | Tests dependency-management concepts. |
| 33 | Go | Go Memory Model | https://go.dev/ref/mem | Tests precise concurrency claims. |
| 34 | Go | Go FAQ | https://go.dev/doc/faq | Question-answer structure for query tests. |
| 35 | Go | Go Code Review Comments | https://go.dev/wiki/CodeReviewComments | Style guidance with many small claims. |
| 36 | Go | Go Generics Tutorial | https://go.dev/doc/tutorial/generics | Tests language feature explanation. |
| 37 | Rust | Rust Book | https://doc.rust-lang.org/book/ | Broad language source. |
| 38 | Rust | Ownership | https://doc.rust-lang.org/book/ch04-00-understanding-ownership.html | Tests core concept extraction. |
| 39 | Rust | Error Handling | https://doc.rust-lang.org/book/ch09-00-error-handling.html | Tests result/error taxonomy. |
| 40 | Rust | Fearless Concurrency | https://doc.rust-lang.org/book/ch16-00-concurrency.html | Good cross-link with Go and SQLite concurrency. |
| 41 | Rust | Cargo Book | https://doc.rust-lang.org/cargo/ | Tests tooling and package manager concepts. |
| 42 | Rust | Rust API Guidelines | https://rust-lang.github.io/api-guidelines/ | Tests design guidance extraction. |
| 43 | Kubernetes | Kubernetes Concepts | https://kubernetes.io/docs/concepts/ | Broad cloud-native source. |
| 44 | Kubernetes | Pods | https://kubernetes.io/docs/concepts/workloads/pods/ | Tests workload abstraction extraction. |
| 45 | Kubernetes | Deployments | https://kubernetes.io/docs/concepts/workloads/controllers/deployment/ | Tests desired-state controller concepts. |
| 46 | Kubernetes | Services | https://kubernetes.io/docs/concepts/services-networking/service/ | Tests networking and discovery concepts. |
| 47 | Kubernetes | ConfigMaps | https://kubernetes.io/docs/concepts/configuration/configmap/ | Tests config vs secret distinction. |
| 48 | Kubernetes | Secrets | https://kubernetes.io/docs/concepts/configuration/secret/ | Tests sensitive-data handling concepts. |
| 49 | Kubernetes | Ingress | https://kubernetes.io/docs/concepts/services-networking/ingress/ | Tests edge routing terminology. |
| 50 | Kubernetes | Persistent Volumes | https://kubernetes.io/docs/concepts/storage/persistent-volumes/ | Tests storage lifecycle concepts. |
| 51 | Kubernetes | StatefulSets | https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/ | Tests stateful workload synthesis. |
| 52 | Kubernetes | Horizontal Pod Autoscaling | https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/ | Tests operational workflow extraction. |
| 53 | Web Standards | CommonMark Specification | https://spec.commonmark.org/0.31.2/ | Large formal spec useful for parser-style content. |
| 54 | Web Standards | GitHub Flavored Markdown Spec | https://github.github.io/gfm/ | Good overlap with CommonMark for comparisons. |
| 55 | Web Standards | Original Markdown Syntax | https://daringfireball.net/projects/markdown/syntax | Historical source for Markdown concept pages. |
| 56 | Web Standards | RFC 7764 Markdown Guidance | https://www.rfc-editor.org/rfc/rfc7764 | Tests RFC-style source extraction. |
| 57 | Web Standards | WHATWG HTML Standard | https://html.spec.whatwg.org/multipage/ | Very long living standard. |
| 58 | Web Standards | Fetch Standard | https://fetch.spec.whatwg.org/ | Tests modern web networking concepts. |
| 59 | Web Standards | URL Standard | https://url.spec.whatwg.org/ | Tests formal definitions and algorithms. |
| 60 | Web Standards | DOM Standard | https://dom.spec.whatwg.org/ | Tests API and data-model concepts. |
| 61 | Web Standards | Service Workers | https://www.w3.org/TR/service-workers/ | Tests offline and request interception concepts. |
| 62 | Web Standards | Indexed Database API | https://www.w3.org/TR/IndexedDB/ | Tests browser storage concepts. |
| 63 | Web Standards | WCAG 2.2 | https://www.w3.org/TR/WCAG22/ | Tests accessibility requirement extraction. |
| 64 | Web Standards | Web Content Accessibility Guidelines 3.0 | https://www.w3.org/TR/wcag-3.0/ | Good for standards evolution comparisons. |
| 65 | Security | OWASP Top 10 | https://owasp.org/www-project-top-ten/ | Security taxonomy source. |
| 66 | Security | OWASP ASVS | https://owasp.org/www-project-application-security-verification-standard/ | Tests control framework extraction. |
| 67 | Security | OWASP Cheat Sheet Series | https://cheatsheetseries.owasp.org/ | Broad practical security guidance. |
| 68 | Security | NIST Cybersecurity Framework 2.0 | https://www.nist.gov/cyberframework | Policy and risk-management source. |
| 69 | Security | NIST SP 800-53 Rev. 5 | https://csrc.nist.gov/publications/detail/sp/800-53/rev-5/final | Large security control catalogue. |
| 70 | Security | CISA Zero Trust Maturity Model | https://www.cisa.gov/resources-tools/resources/zero-trust-maturity-model | Tests government security guidance. |
| 71 | Security | Mozilla Web Security Guidelines | https://infosec.mozilla.org/guidelines/web_security | Practical security checklist source. |
| 72 | Security | Let's Encrypt Rate Limits | https://letsencrypt.org/docs/rate-limits/ | Tests operational limits and policy extraction. |
| 73 | Data | Data.gov Catalog API | https://resources.data.gov/catalog-api/ | Tests API documentation summarisation. |
| 74 | Data | DCAT-US 3.0 | https://resources.data.gov/resources/dcat-us3/ | Metadata schema source. |
| 75 | Data | Project Open Data Metadata Schema v1.1 | https://resources.data.gov/resources/dcat-us/ | Useful for comparing schema versions. |
| 76 | Data | Data.gov Open Data How-To | https://resources.data.gov/resources/data-gov-open-data-howto/ | Tests procedural public-sector guidance. |
| 77 | Data | W3C Data Catalog Vocabulary | https://www.w3.org/TR/vocab-dcat-3/ | Formal metadata vocabulary. |
| 78 | Data | FAIR Principles | https://www.go-fair.org/fair-principles/ | Tests research-data concept extraction. |
| 79 | Data | Open Definition | https://opendefinition.org/od/2.1/en/ | Tests legal/policy definition extraction. |
| 80 | Data | Creative Commons Licenses | https://creativecommons.org/share-your-work/cclicenses/ | Tests licensing concept pages. |
| 81 | AI/ML | Attention Is All You Need | https://arxiv.org/abs/1706.03762 | Landmark transformer paper. |
| 82 | AI/ML | BERT | https://arxiv.org/abs/1810.04805 | Landmark NLP pretraining paper. |
| 83 | AI/ML | GPT-3 Language Models are Few-Shot Learners | https://arxiv.org/abs/2005.14165 | Tests model-scaling concept extraction. |
| 84 | AI/ML | Retrieval-Augmented Generation | https://arxiv.org/abs/2005.11401 | Directly relevant to wiki-vs-RAG comparisons. |
| 85 | AI/ML | LoRA | https://arxiv.org/abs/2106.09685 | Tests fine-tuning concept extraction. |
| 86 | AI/ML | Chain-of-Thought Prompting | https://arxiv.org/abs/2201.11903 | Tests prompting technique concepts. |
| 87 | AI/ML | Constitutional AI | https://arxiv.org/abs/2212.08073 | Tests alignment and safety concepts. |
| 88 | AI/ML | Direct Preference Optimization | https://arxiv.org/abs/2305.18290 | Tests preference-learning concepts. |
| 89 | AI/ML | Llama 2 | https://arxiv.org/abs/2307.09288 | Tests model release paper extraction. |
| 90 | AI/ML | Mistral 7B | https://arxiv.org/abs/2310.06825 | Tests concise model architecture paper. |
| 91 | AI/ML | Qwen2 Technical Report | https://arxiv.org/abs/2407.10671 | Existing corpus topic overlap. |
| 92 | AI/ML | Llama 3 Herd of Models | https://arxiv.org/abs/2407.21783 | Tests long multi-model technical report. |
| 93 | Systems | Raft Paper | https://raft.github.io/raft.pdf | Consensus source for distributed systems. |
| 94 | Systems | Dynamo Paper | https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf | Tests eventually consistent storage concepts. |
| 95 | Systems | Spanner Paper | https://research.google/pubs/spanner-googles-globally-distributed-database/ | Distributed database source. |
| 96 | Systems | MapReduce Paper | https://research.google/pubs/mapreduce-simplified-data-processing-on-large-clusters/ | Classic distributed processing source. |
| 97 | Systems | Borg Paper | https://research.google/pubs/large-scale-cluster-management-at-google-with-borg/ | Cluster management source. |
| 98 | Systems | Site Reliability Engineering Book | https://sre.google/sre-book/table-of-contents/ | Operational practice source. |
| 99 | Systems | The Twelve-Factor App | https://12factor.net/ | App deployment methodology source. |
| 100 | Systems | Martin Fowler - Microservices | https://martinfowler.com/articles/microservices.html | Architecture source with trade-offs. |
| 101 | Project Gutenberg | Pride and Prejudice | https://www.gutenberg.org/ebooks/1342 | Public-domain narrative text. |
| 102 | Project Gutenberg | Moby-Dick | https://www.gutenberg.org/ebooks/2701 | Long public-domain narrative text. |
| 103 | Project Gutenberg | Frankenstein | https://www.gutenberg.org/ebooks/84 | Public-domain literature with entities and themes. |
| 104 | Project Gutenberg | Alice's Adventures in Wonderland | https://www.gutenberg.org/ebooks/11 | Shorter narrative source. |
| 105 | Project Gutenberg | The Adventures of Sherlock Holmes | https://www.gutenberg.org/ebooks/1661 | Multi-story narrative source. |
| 106 | Project Gutenberg | The Time Machine | https://www.gutenberg.org/ebooks/35 | Public-domain science fiction. |
| 107 | Project Gutenberg | Dracula | https://www.gutenberg.org/ebooks/345 | Public-domain novel with rich entities. |
| 108 | Project Gutenberg | A Tale of Two Cities | https://www.gutenberg.org/ebooks/98 | Long historical novel. |
| 109 | Project Gutenberg | The Republic | https://www.gutenberg.org/ebooks/1497 | Philosophical dialogue source. |
| 110 | Project Gutenberg | The Art of War | https://www.gutenberg.org/ebooks/132 | Short strategic text. |
| 111 | Public Policy | GDPR Text | https://gdpr-info.eu/ | Tests legal article extraction. |
| 112 | Public Policy | EU AI Act Overview | https://digital-strategy.ec.europa.eu/en/policies/regulatory-framework-ai | Tests current policy synthesis. |
| 113 | Public Policy | OECD AI Principles | https://oecd.ai/en/ai-principles | Tests governance concept extraction. |
| 114 | Public Policy | NIST AI Risk Management Framework | https://www.nist.gov/itl/ai-risk-management-framework | Tests risk taxonomy extraction. |
| 115 | Public Policy | White House Blueprint for an AI Bill of Rights | https://www.whitehouse.gov/ostp/ai-bill-of-rights/ | Tests policy document summarisation. |
| 116 | Public Policy | ISO/IEC 42001 Overview | https://www.iso.org/standard/81230.html | Tests standards metadata without full standard text. |
| 117 | Public Policy | UNESCO Recommendation on AI Ethics | https://www.unesco.org/en/artificial-intelligence/recommendation-ethics | Tests international governance concepts. |
| 118 | Public Policy | WHO Ethics and Governance of AI for Health | https://www.who.int/publications/i/item/9789240029200 | Tests health-policy source extraction. |
| 119 | Public Policy | NASA Open Data Portal | https://data.nasa.gov/ | Public-sector data portal source. |
| 120 | Public Policy | World Bank Open Data | https://data.worldbank.org/ | Data portal source with many entities and indicators. |

## Batch Suggestions

Batch 1, repo and local data systems:
1-20.

Batch 2, programming languages and cloud operations:
21-52.

Batch 3, standards, security, and open data:
53-80.

Batch 4, AI/ML and distributed systems:
81-100.

Batch 5, literature and public policy:
101-120.
