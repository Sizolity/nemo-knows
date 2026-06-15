---
title: Direct Preference Optimization
kind: source
created: 2026-05-18
updated: 2026-05-18
sources:
  - raw/web/curated-web-corpus-2026-05-18.md
  - https://arxiv.org/abs/2305.18290
tags: [ai-ml, web-corpus]
confidence: medium
---

# Direct Preference Optimization

## Fetch Metadata

- Corpus item: 88
- Category: AI/ML
- Source URL: https://arxiv.org/abs/2305.18290
- Final URL: https://arxiv.org/abs/2305.18290
- Retrieved: 2026-05-18
- Content-Type: text/html; charset=utf-8
- Fetch status: ok
- Test value: Tests preference-learning concepts.
- Fetched page title: [2305.18290] Direct Preference Optimization: Your Language Model is Secretly a Reward Model

## Retrieved Text

Skip to main content
Cornell University
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member
institutions, and all contributors. Donate
arxiv logo > cs > arXiv:2305.18290
____________________

Help | Advanced Search
[All fields________]
(BUTTON) Search
arXiv logo
Cornell University Logo
(BUTTON) open search
____________________ (BUTTON) GO
(BUTTON) open navigation menu

quick links

* Login
* Help Pages
* About

Computer Science > Machine Learning

arXiv:2305.18290 (cs)
[Submitted on 29 May 2023 (v1), last revised 29 Jul 2024 (this version,
v3)]

Title:Direct Preference Optimization: Your Language Model is Secretly a
Reward Model

Authors:Rafael Rafailov, Archit Sharma, Eric Mitchell, Stefano Ermon,
Christopher D. Manning, Chelsea Finn
View a PDF of the paper titled Direct Preference Optimization: Your
Language Model is Secretly a Reward Model, by Rafael Rafailov and 5
other authors
View PDF HTML (experimental)

Abstract:While large-scale unsupervised language models (LMs) learn
broad world knowledge and some reasoning skills, achieving precise
control of their behavior is difficult due to the completely
unsupervised nature of their training. Existing methods for gaining
such steerability collect human labels of the relative quality of
model generations and fine-tune the unsupervised LM to align with
these preferences, often with reinforcement learning from human
feedback (RLHF). However, RLHF is a complex and often unstable
procedure, first fitting a reward model that reflects the human
preferences, and then fine-tuning the large unsupervised LM using
reinforcement learning to maximize this estimated reward without
drifting too far from the original model. In this paper we introduce
a new parameterization of the reward model in RLHF that enables
extraction of the corresponding optimal policy in closed form,
allowing us to solve the standard RLHF problem with only a simple
classification loss. The resulting algorithm, which we call Direct
Preference Optimization (DPO), is stable, performant, and
computationally lightweight, eliminating the need for sampling from
the LM during fine-tuning or performing significant hyperparameter
tuning. Our experiments show that DPO can fine-tune LMs to align
with human preferences as well as or better than existing methods.
Notably, fine-tuning with DPO exceeds PPO-based RLHF in ability to
control sentiment of generations, and matches or improves response
quality in summarization and single-turn dialogue while being
substantially simpler to implement and train.

Subjects: Machine Learning (cs.LG); Artificial Intelligence (cs.AI);
Computation and Language (cs.CL)
Cite as: arXiv:2305.18290 [cs.LG]
(or arXiv:2305.18290v3 [cs.LG] for this version)
https://doi.org/10.48550/arXiv.2305.18290
(BUTTON) Focus to learn more
arXiv-issued DOI via DataCite

Submission history

From: Archit Sharma [view email]
[v1] Mon, 29 May 2023 17:57:46 UTC (982 KB)
[v2] Wed, 13 Dec 2023 18:48:48 UTC (983 KB)
[v3] Mon, 29 Jul 2024 22:26:36 UTC (999 KB)
Full-text links:

Access Paper:

View a PDF of the paper titled Direct Preference Optimization: Your
Language Model is Secretly a Reward Model, by Rafael Rafailov and 5
other authors
* View PDF
* HTML (experimental)
* TeX Source

license icon view license
Current browse context:
cs.LG
< prev | next >
new | recent | 2023-05
Change to browse by:
cs
cs.AI
cs.CL

References & Citations

* NASA ADS
* Google Scholar
* Semantic Scholar

2 blog links

(what is this?)
export BibTeX citation Loading...

BibTeX formatted citation

×

loading...__________________________________________________
____________________________________________________________
____________________________________________________________
____________________________________________________________
Data provided by:

Bookmark

BibSonomy logo Reddit logo
(*) Bibliographic Tools

Bibliographic and Citation Tools

[ ] Bibliographic Explorer Toggle
Bibliographic Explorer (What is the Explorer?)
[ ] Connected Papers Toggle
Connected Papers (What is Connected Papers?)
[ ] Litmaps Toggle
Litmaps (What is Litmaps?)
[ ] scite.ai Toggle
scite Smart Citations (What are Smart Citations?)
( ) Code, Data, Media

Code, Data and Media Associated with this Article

[ ] alphaXiv Toggle
alphaXiv (What is alphaXiv?)
[ ] Links to Code Toggle
CatalyzeX Code Finder for Papers (What is CatalyzeX?)
[ ] DagsHub Toggle
DagsHub (What is DagsHub?)
[ ] GotitPub Toggle
Gotit.pub (What is GotitPub?)
[ ] Huggingface Toggle
Hugging Face (What is Huggingface?)
[ ] Links to Code Toggle
Papers with Code (What is Papers with Code?)
[ ] ScienceCast Toggle
ScienceCast (What is ScienceCast?)
( ) Demos

Demos

[ ] Replicate Toggle
Replicate (What is Replicate?)
[ ] Spaces Toggle
Hugging Face Spaces (What is Spaces?)
[ ] Spaces Toggle
TXYZ.AI (What is TXYZ.AI?)
( ) Related Papers

Recommenders and Search Tools

[ ] Link to Influence Flower
Influence Flower (What are Influence Flowers?)
[ ] Core recommender toggle
CORE Recommender (What is CORE?)
[ ] IArxiv recommender toggle
IArxiv Recommender (What is IArxiv?)
* Author
* Venue
* Institution
* Topic

( ) About arXivLabs

arXivLabs: experimental projects with community collaborators

arXivLabs is a framework that allows collaborators to develop and share
new arXiv features directly on our website.

Both individuals and organizations that work with arXivLabs have
embraced and accepted our values of openness, community, excellence,
and user data privacy. arXiv is committed to these values and only
works with partners that adhere to them.

Have an idea for a project that will add value for arXiv's community?
Learn more about arXivLabs.

Which authors of this paper are endorsers? | Disable MathJax (What is
MathJax?)

* About
* Help

* Click here to contact arXiv Contact
* Click here to subscribe Subscribe

* Copyright
* Privacy Policy

* Web Accessibility Assistance
* arXiv Operational Status
