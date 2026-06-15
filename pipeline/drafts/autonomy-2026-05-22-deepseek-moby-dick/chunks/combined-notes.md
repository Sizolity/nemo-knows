## chunk-01

---
title: Chunk 1 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

### Chunk Context
This is chunk 1 of 17 from `raw/web/corpus-2026-05-18/102-moby-dick.md`.
It covers lines 1–26, the YAML frontmatter and the “Moby-Dick > Fetch Metadata” section.
The chunk precedes any narrative content from the novel.

### Local Summary
The chunk records provenance metadata for the source document: corpus item 102, Project Gutenberg ebook 2701, source and final URLs, retrieval date (2026-05-18), content type (`text/plain; charset=utf-8`), and a fetch status.
The fetch succeeded via a supplemental curl request after the ebook landing page failed TLS when accessed with urllib.
A test value labels the content as “Long public-domain narrative text.”

### Key Claims
- The primary source is a Project Gutenberg plain-text edition of *Moby-Dick*, ebook 2701.
- The final URL used is `https://www.gutenberg.org/files/2701/2701-0.txt`.
- Retrieval occurred on 2026-05-18; content type is `text/plain; charset=utf-8`.
- Fetch status is “ok via supplemental curl fetch” because the landing page (`https://www.gutenberg.org/ebooks/2701`) failed TLS in urllib.
- The source is tagged `project-gutenberg` and `web-corpus` with confidence “medium.”

### Entities And Concepts
- **Moby-Dick** – Project Gutenberg etext 2701.
- **Project Gutenberg** – repository hosting the plain-text file.
- **Supplemental curl fetch** – fallback retrieval method after a TLS failure with urllib.
- **Corpus metadata** – item 102, category “Project Gutenberg”, test value “Long public-domain narrative text.”

### Procedures And API Details
- No API.
- The acquisition process:
  1. Attempt to fetch the landing page (`https://www.gutenberg.org/ebooks/2701`) via urllib – failed due to TLS.
  2. Supplemental curl fetch from the direct plain-text URL (`https://www.gutenberg.org/files/2701/2701-0.txt`) succeeded, yielding the source document.

### Nuance Or Contradictions
- The chunk is entirely metadata and contains no content from the novel.
- The fetch status documents a workaround for a TLS issue, which matters for reproducibility of the corpus build.

### Candidate Wiki Hints
- A page on **corpus acquisition methods** could describe the urllib‑TLS fallback pattern observed with Project Gutenberg items.
- A reference page for **Project Gutenberg source URLs** in the curated corpus could record the direct text‑path pattern (`/files/{ebook_id}/{ebook_id}-0.txt`).

## chunk-02

---
title: Chunk 02 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context

- Source: raw/web/corpus-2026-05-18/102-moby-dick.md
- Chunk: 2 of 17
- Lines: 27–1333
- Heading path: Moby-Dick > Retrieved Text
- The chunk covers the full “Retrieved Text” section from the Project Gutenberg edition, including the table of contents, the “Etymology” and “Extracts” prefatory materials, and the beginning of the narrative: Chapter 1 (Loomings), Chapter 2 (The Carpet-Bag), and the opening of Chapter 3 (The Spouter-Inn). It ends mid-scene in Chapter 3.

## Local Summary

The chunk opens with the header matter of *Moby-Dick*: a list of 135 chapters plus an epilogue, followed by the “Etymology” (a brief reflection on the word “whale” across languages) and the “Extracts” (Supplied by a Sub-Sub-Librarian). The extracts are a deliberately jumbled collection of over 80 quotations about whales from ancient and modern sources, ranging from the Bible to scientific and literary works. The narrator’s framing voice cautions that these excerpts are not “veritable gospel cetology” but a promiscuous bird’s‑eye view of what has been said of the whale.

The narrative then begins in the first person. The speaker, who calls himself Ishmael, explains his periodic need to go to sea as a remedy for melancholy and a substitute for suicide. He decides to sign onto a whaling voyage, preferring to sail as a simple sailor rather than a passenger, and insists on departing from Nantucket. Arriving in New Bedford on a cold Saturday night in December, he searches for cheap lodgings and eventually finds the Spouter‑Inn, whose sign advertises “Peter Coffin”. Inside, he examines a large smoky painting that he eventually interprets as a whale impaling itself on a ship’s masts, and he sees an array of old whaling weapons. The landlord tells him no private bed is available but he can share a harpooneer’s blanket. After a grim supper (including dumplings), the crew of the Grampus arrives noisily; among them is a tall, reserved man (later identified as Bulkington) who slips away. Ishmael grows uneasy about sleeping with the unnamed harpooneer and, before the harpooneer appears, resolves not to share the bed—ending the chunk with “Landlord! I’ve changed my mind about that harpooneer.—I shan’t sleep with him. I’ll try the bench here.”

## Key Claims

- The etymology section presents the word “whale” as derived from roundness or rolling, citing Webster and Richardson, and lists the term in Hebrew, Greek, Latin, Anglo‑Saxon, Danish, Dutch, Swedish, Icelandic, English, French, Spanish, Fegee, and Erromangoan.
- The extracts are described as the painstaking work of a “poor devil of a Sub‑Sub” who combed libraries for every mention of whales; they are offered as a panorama of human thought about Leviathan, not as authoritative cetology.
- Among the extracts are quotations from Genesis, Job, Jonah, Psalms, Isaiah, Plutarch, Pliny, Lucian, King Alfred, Montaigne, Rabelais, Bacon, Shakespeare, Hobbes, Milton, Dryden, Thomas Jefferson, Edmund Burke, Blackstone, Cowper, Cuvier, Darwin, and many others, covering subjects from the whale’s size and anatomy to whaling practices and shipwrecks.
- Ishmael goes to sea whenever he feels a “damp, drizzly November in my soul” and considers it his alternative to “pistol and ball.”
- He claims almost all men share a deep, magnetic attraction to water, using the image of crowds of New Yorkers gazing seaward and a meditation‑water connection.
- He always ships as a common sailor because passengers pay (a hardship) while sailors get paid, and because he abominates “honorable respectable toils.”
- He insists on sailing from Nantucket, regarding it as the original great whaling port.
- Arriving in New Bedford, he seeks cheap lodgings, rejects the “Crossed Harpoons” and “Sword‑Fish Inn,” stumbles into a negro church (“The Trap”), and finally chooses the dilapidated Spouter‑Inn.
- The inn’s smoky painting is eventually understood as a whale in the act of impaling itself on a ship’s mast‑heads; the entry is hung with clubs, spears, and old harpoons with stories attached.
- The landlord offers a shared bed with a harpooneer who is “dark complexioned” and eats only rare steaks.
- The crew of the Grampus (a whaler) enters and drinks heavily; one sober, tall Southerner (Bulkington) leaves early.
- Ishmael, increasingly anxious, retracts his agreement to share a bed and asks to sleep on a bench instead.

## Entities And Concepts

- **Ishmael** – first‑person narrator; decides to go whaling from Nantucket, travels to New Bedford.
- **The Sub‑Sub‑Librarian** – persona of the compiler of the extracts, portrayed as a hopeless, threadbare scholar.
- **New Bedford** – whaling port where Ishmael stays overnight; gradually monopolising the whaling business.
- **Nantucket** – original great American whaling island, source of the first dead American whale stranded and first whalemen.
- **Spouter‑Inn** – cheap lodging in New Bedford run by a landlord (Peter Coffin?); sign shows a jet of spray.
- **Peter Coffin** – name on the inn sign; common Nantucket name.
- **The painting in the inn** – a large, sooty, ambiguous image later interpreted as a Cape‑Horner in a hurricane with a whale impaling itself on the masts.
- **The weapons on the wall** – clubs, spears, lances, harpoons; some storied, e.g., Nathan Swain’s lance that killed fifteen whales in one day.
- **The bar** – made from a right whale’s jaw, inside which the old landlord (nicknamed Jonah) serves drinks.
- **The “dark complexioned” harpooneer** – yet to appear; described by the landlord as one who never eats dumplings, only rare steaks.
- **Bulkington** – tall, sober Southerner among the Grampus crew; slips away and will later become Ishmael’s shipmate.
- **The Grampus** – a whaling ship whose crew arrives at the inn that night.
- **Cape Horn, the Pacific** – Ishmael’s intended destination.
- **Euroclydon** – a tempestuous wind, contrasted between the view from within a warm house and the exposure of a poor man like Lazarus.

## Procedures And API Details

- None (no technical procedures or API commands).

## Nuance Or Contradictions

- The chunk ends abruptly mid‑scene in Chapter 3, with Ishmael’s line “Landlord! I’ve changed my mind about that harpooneer.—I shan’t sleep with him. I’ll try the bench here.” There is no truncation marker, but the raw source stops at that point before the harpooneer enters.
- The extracts are explicitly presented as “promiscuous” and not scientifically reliable, yet the sheer range of sources (including biblical, classical, and contemporary whaling narratives) establishes the cultural weight of the whale before the story begins.
- Ishmael’s statement “I abominate all honorable respectable toils” contrasts with his detailed description of the dignity of a simple sailor and his philosophical reflections; the text itself carries a highly literary, almost scholarly tone that belies a mere “simple sailor.”

## Candidate Wiki Hints

- A page on *Moby-Dick’s prefatory materials* could discuss the function of the “Etymology” and “Extracts” as a thematic overture, drawing on this chunk’s explicit framing (“not … veritable gospel cetology”).
- A page on *Ishmael’s motives for going to sea* would track his statements here about melancholy, water‑gazing, and the choice to ship as a common sailor.
- A page on *the Spouter‑Inn and its painting* could examine the hermeneutic game set up by the description of the indecipherable picture and its eventual interpretation.

## chunk-03

---
title: Chunk 03 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/102-moby-dick.md`
- Chunk: 3 of 17, lines 1335-2352
- Heading path: Moby-Dick > Retrieved Text
- Narrative covers Ishmael’s first night at the Spouter-Inn with Queequeg through the start of Father Mapple’s sermon in New Bedford (Chapters 3–9).

## Local Summary
Ishmael resolves to share a bed with the absent harpooneer. The landlord explains the harpooneer’s “head-peddling” as selling embalmed New Zealand heads. Queequeg returns late, a heavily tattooed “cannibal,” prays to a small idol, and climbs into bed with a tomahawk-pipe. After initial terror, Ishmael accepts him as a “sober cannibal” and sleeps well. They wake entangled; Queequeg’s tattooed arm blending with the patchwork quilt prompts a childhood memory of Ishmael’s phantom-hand experience. Queequeg dresses eccentrically (under the bed for boots, shaves with his harpoon). At breakfast a bashful crew of whalemen contrasts with Queequeg’s cool demeanor. Ishmael strolls through New Bedford, remarking on the town’s cosmopolitan whaling wealth, then visits the Whaleman’s Chapel with its marble memorials to dead seamen. The famous Father Mapple enters, climbs a rope ladder into a ship’s‑bow‑shaped pulpit, and begins his sermon on Jonah.

## Key Claims
- The harpooneer (Queequeg) is a South Sea “savage” who sells embalmed heads and is covered in dark‑purplish tattoos.
- He worships a small wooden idol (“Congo baby”) with burnt‑offering shavings and biscuit.
- Ishmael’s fear gives way to a pragmatic maxim: “Better sleep with a sober cannibal than a drunken Christian.”
- Queequeg’s tattooed arm against the patchwork counterpane appears indistinguishable from the quilt.
- Ishmael recalls a childhood memory of a supernatural hand that felt similarly strange but terrifying.
- Queequeg’s manners (dressing under the bed, harpoon‑shaving, restricting wash to torso) show a man “in the transition stage—neither caterpillar nor butterfly.”
- Whalemen at breakfast are bashful and silent despite their daring at sea.
- New Bedford’s opulent homes and gardens were “harpooned and dragged up” from the oceans by whaling wealth.
- The Whaleman’s Chapel contains marble tablets memorialising men killed by whales, loss overboard, or towed away boats.
- Father Mapple’s physical isolation in the pulpit (pulling up the ladder) symbolises spiritual withdrawal and a “self‑containing stronghold.”
- In his sermon, Father Mapple interprets Jonah as a lesson in disobedience, flight, punishment, repentance, and deliverance.

## Entities And Concepts
- **Queequeg**: Harpooneer from the South Seas, heavily tattooed, peddler of embalmed heads, idol‑worshipper, described as a “sober cannibal.”
- **The landlord (Peter Coffin)**: Proprietor of the Spouter-Inn; reveals the mystery of the “head” sales; mediates between Ishmael and Queequeg.
- **The New Zealand head**: Embalmed head Queequeg tries to sell; pushed into his seabag at night.
- **Queequeg’s idol (“Congo baby”)**: Small, black, hunch‑backed wooden figure prayed to with offerings.
- **Tomahawk‑pipe**: Queequeg’s weapon/implement, used for smoking and shaving.
- **Patchwork counterpane**: Quilt whose colours blend with Queequeg’s tattooed arm, blurring human and fabric.
- **Ishmael’s childhood phantom hand**: A memory of a terrifying, invisible hand felt while in bed, compared to Queequeg’s embrace.
- **New Bedford**: Whaling port full of cannibals, “bumpkin dandies,” lavish homes, and oil wealth.
- **Whaleman’s Chapel**: Contains marble cenotaphs to lost whalemen; pulpit shaped like a ship’s bow.
- **Father Mapple**: Former sailor‑harpooneer turned chaplain; ascends pulpit via rope ladder drawn up after him.
- **Jonah narrative**: Sermon text emphasising disobedience, flight, the “contradiction in the lamp” as crooked conscience, and the captain who tests Jonah’s purse.

## Procedures And API Details
- None.

## Nuance Or Contradictions
- The landlord’s cryptic claim that the harpooneer “can’t sell his head” is later resolved as a literal statement about unsold merchandise, not madness.
- Ishmael vacillates between fear and pragmatic acceptance of Queequeg; the text satirises “civilised” fear of the “savage” while pointing out Queequeg’s innate politeness.
- Queequeg’s manners are described simultaneously as “civilised” (giving Ishmael privacy to dress) and “savage” (dressing under the bed, harpoon‑shaving).

## Candidate Wiki Hints
- **Queequeg**: A reusable character page for the harpooneer, his origins, tattoos, idolatry, and relationship with Ishmael.
- **Whaleman’s Chapel**: A setting page covering the marble tablets, the pulpit-as-ship’s‑bow symbol, and Father Mapple’s entry ritual.
- **Father Mapple**: A topic page for the chaplain’s biography, preaching style, and his Jonah interpretation.
- **New Bedford in Moby-Dick**: A topic for the novel’s depiction of the town’s whaling wealth, demographics, and cosmopolitan docks.

## chunk-04

---
title: Chunk 04 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/102-moby-dick.md`, chunk 4 of 17, lines 2354‑3374.
- Heading path: Moby‑Dick > Retrieved Text.
- Continues Father Mapple’s sermon (climax and closing), then chapters 10‑15 and the opening of chapter 16 (“The Ship”).

## Local Summary
Father Mapple delivers a vivid retelling of Jonah’s ordeal, portraying him as a model of penitence, not of sin. After the sermon, Ishmael returns to the Spouter‑Inn, bonds with Queequeg, and the two share an intimate night of talk, pipe‑smoking, and a mutual acceptance that culminates in a “bosom friend” bond. Queequeg recounts his history as a Polynesian prince who left his island to learn from Christians but was disillusioned. He resolves to continue whaling and invites Ishmael to join him. They travel to Nantucket, where Queequeg’s physical prowess and quiet courage are displayed in a storm‑rescue. The narrative describes Nantucket as an isolated, sea‑soaked outpost, the Try Pots inn and its chowders, and the pair’s search for a whaleship. Ishmael selects the Pequod after a cryptic divine endorsement by Queequeg’s idol, Yojo. Captain Peleg, a part‑owner, tests Ishmael’s resolve with brusque, sceptical questions.

## Key Claims
- Jonah’s story is a lesson in true repentance: not demanding deliverance, but accepting punishment and turning toward God.
- Father Mapple applies Jonah’s lesson personally, declaring himself a greater sinner who must preach truth even when it appals.
- Queequeg’s “indifference” masks a simple, honest heart; his savage exterior contains a “Socratic wisdom” and a capacity for deep, unprejudiced friendship.
- Ishmael’s willingness to join Queequeg in idol‑worship is justified by redefining worship as doing God’s will, which includes loving one’s neighbour.
- Intimacy and comfort depend on contrast (e.g., a little coldness makes warmth more delightful).
- Queequeg is a king’s son who left Rokovoko to bring enlightenment to his people but ended up scornful of Christians’ hypocrisy.
- Nantucket is depicted as a sand‑heaped, lonely island whose inhabitants have conquered the sea, making the world’s oceans their plantation.
- The chowders at the Try Pots are a regional obsession; Hosea Hussey’s cow feeds on fish remnants.
- Queequeg’s idol Yojo insists that the choice of the ship be left to Ishmael alone, promising guidance.
- The Pequod is a grizzled, trophy‑laden whaler, melancholy yet noble, fitted with jawbone‑and‑ivory fixtures.
- Captain Peleg interrogates Ishmael about his motives, deriding the merchant service, and hints that Captain Ahab has lost a leg to a monstrous sperm whale.

## Entities And Concepts
- **Jonah** – Biblical prophet used as sermon subject; model of contrite acceptance.
- **Father Mapple** – Preacher who identifies with Jonah as a “pilot of the living God” called to speak hard truths.
- **Queequeg** – Cannibal‑heritage harpooneer from Rokovoko, son of a king, described as “George Washington cannibalistically developed.” Carries his own harpoon; rescues a greenhorn at sea.
- **Yojo** – Queequeg’s small black idol that prescribes the ship‑selection method.
- **Ishmael** – The narrator, who bonds with Queequeg and decides to ship on the Pequod.
- **Rokovoko** – Queequeg’s unmapped island kingdom “far away to the West and South.”
- **Nantucket** – Isolated sandy island, home port of the world’s whalemen, described as a sand‑hillock with no background.
- **Try Pots** – Inn kept by Hosea and Mrs. Hussey, famous for clam and cod chowder; decorated with whaling relics; harpoons forbidden in rooms.
- **Moss** – Packet schooner that carries Ishmael and Queequeg from New Bedford to Nantucket.
- **Pequod** – Whaleship of “old school,” owned by Peleg and Bildad, adorned with whale teeth and jaw‑bone tiller, captained by the missing‑legged Ahab.
- **Captain Peleg** – Retired whaleman, part‑owner and outfitter of the Pequod; gruff, distrustful of “aliens,” mocks merchant service.
- **Captain Ahab** – Mentioned only in absentia: captain of the Pequod, described as having one leg lost to a monstrous sperm whale.

## Procedures And API Details
- None.

## Nuance Or Contradictions
- The sermon’s interpretation of Jonah as a model of repentance stands in tension with the biblical text’s coercion (Jonah flees and is swallowed by a fish). The chunk presents it as a positive moral lesson.
- Ishmael’s religious accommodation (worshipping the idol by redefining worship) is a deliberate, pragmatic reinterpretation, not a doctrinal stance.
- Queequeg’s “indifference” is presented as both alienating and sublime, an ambiguous marker of noble savagery.
- The chunk ends mid‑chapter (chapter 16) without an explicit truncation marker; the final line is Peleg’s dismissive question: “Can’t ye see the world where you stand?”

## Candidate Wiki Hints
- **Queequeg** – A page collecting his biography, character traits, and relationship with Ishmael would be reusable throughout the novel.
- **Nantucket as whaling capital** – The hyperbolic descriptions of the island and its inhabitants can ground a topic on Moby‑Dick’s geographical and cultural setting.
- **The Pequod (ship description)** – The rich, symbol‑laden depiction of the ship and its owners could support a page on the Pequod’s material history and symbolism.
- **Father Mapple’s sermon** – Could be abstracted into a note on its theological‑narrative technique and the Jonah‑as‑pilot theme.
- **Yojo** – The role of Queequeg’s idol in guiding choice could form a small concept page about fate and decision in the novel.

## chunk-05

---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source: raw/web/corpus-2026-05-18/102-moby-dick.md
- Chunk: 5 of 17
- Lines: 3376–4501
- Heading path: Moby-Dick > Retrieved Text
- Coverage: Continuation of Chapter 16 through Chapter 22 (ending with “Merry Christmas”)
- No truncation marker present; the chunk concludes at the end of a chapter.

## Local Summary
Ishmael signs aboard the Pequod, meets part-owner and former captain Peleg, along with the pious yet contradictory Captain Bildad. After a heated negotiation over the lay (profit share), Ishmael is signed for the 300th lay. He asks about Captain Ahab, whom he has not yet seen, and receives an evasive, admiring, and ominous description from Peleg. Queequeg undergoes a prolonged Ramadan fast in silence, alarming the landlord and Ishmael, who eventually lectures him on religious folly. Queequeg impresses Peleg and Bildad with a harpoon demonstration and signs with his tattooed mark, securing the 90th lay. A mysterious ragged stranger named Elijah issues cryptic warnings about Ahab and the voyage. The Pequod’s outfitting proceeds under Bildad and his sister Charity. On sailing day, Ishmael and Queequeg board, find the crew mostly absent, and encounter the rigger who reveals Ahab came aboard the night before. Peleg and Bildad oversee departure while Ahab remains unseen in his cabin.

## Key Claims
- Nantucket’s Quaker heritage persists, but whalemen there are “fighting Quakers” who blend pious speech with violent pursuits, creating a paradoxical character type.
- Bildad embodies this contradiction: a strict Quaker who refuses to bear arms against humans yet has killed countless whales and driven crews to exhaustion.
- Ishmael speculates that such contradictions are reconciled by separating religion from the “practical world.”
- Peleg describes Ahab as a “grand, ungodly, god-like man,” educated and experienced beyond whales, with a mysterious, moody disposition following the loss of his leg.
- Peleg warns Ishmael never to repeat the biblical association of Ahab with a wicked king, claiming the name was his mother’s whim and denying any prophetic force.
- Ishmael reflects that all men—Presbyterians and pagans alike—are “dreadfully cracked” in matters of belief.
- Queequeg’s Ramadan is depicted as physically extreme but harmless; Ishmael later dismisses prolonged fasting as unhealthy and conducive to gloomy theology.
- The stranger Elijah hints at undisclosed events: Ahab’s three-day deathlike state off Cape Horn, a skirmish with a Spaniard, and the leg’s loss as prophecy fulfilled.
- Elijah’s parting remark suggests legal troubles (“unless it’s before the Grand Jury”), deepening the air of foreboding.
- The ship’s preparations highlight the complexity of a three-year whaling voyage, with Aunt Charity’s gifts blending domesticity and violence (oil-ladle and lance).

## Entities And Concepts
- **Captain Peleg**: Part-owner of the Pequod, retired whaleman, blusterous and irreverent, yet protective of Ahab’s reputation.
- **Captain Bildad**: Part-owner, strict Quaker, miserly, hard taskmaster, relies on biblical language.
- **Lay system**: Profit shares for whaling crew; green hands commonly receive a longer lay (e.g., 275th, 300th); Queequeg obtains the 90th lay as a skilled harpooneer.
- **Captain Ahab**: Never appears in person; described as moody, one-legged, college-educated, with a prophesied name and a recent unhinged period.
- **Queequeg’s Ramadan**: A day-long immobile fast with Yojo on his head, causing alarm; ends at sunrise without explanation.
- **Elijah**: A ragged prophet figure who intercepts Ishmael and Queequeg, hints at Ahab’s secrets, and trails them.
- **Aunt Charity**: Bildad’s sister, kindly and industrious, brings practical supplies plus weapons to the ship.
- **First Congregational Church**: Ishmael’s ironic retort to Bildad’s demand that Queequeg show Christian church membership; he argues all humanity belongs to one “ancient Catholic Church.”

## Procedures And API Details
No technical procedures or API details appear in this literary chunk.

## Nuance Or Contradictions
- Bildad’s piety and non-violence towards humans coexist with whaling violence and harsh treatment of sailors; Ishmael notes the inconsistency but says it seems resolved by compartmentalising religion from business.
- Ishmael both respects and mocks religious practice: he defends Queequeg’s Ramadan, yet later lectures him against it and calls dyspeptic religionists melancholy.
- Peleg warns against the biblical Ahab association, yet his description of Ahab emphasises a regal, almost supernatural quality, which undercuts his denial.
- Elijah’s cryptic statements are never clarified, leaving open whether genuine foreknowledge or lunacy is at work.
- The chunk ends with Ahab still absent, though the rigger indicates he came aboard the night before; the delay in his appearance to the narrator maintains tension.

## Candidate Wiki Hints
- **Lay (whaling share)** – concept of profit-sharing aboard whalers, with examples from the Pequod’s articles.
- **Peleg (Moby-Dick)** – character page for Captain Peleg.
- **Bildad (Moby-Dick)** – character page for Captain Bildad, the Quaker shipowner.
- **Queequeg’s Ramadan** – episode note on the fasting and Ishmael’s religious tolerance.
- **Elijah (prophet figure)** – note on this cryptic character and his warnings.
- **Ahab’s backstory and name** – summary of the early foreshadowing before his appearance.
- **Nantucket Quaker whalemen** – the cultural contradiction of pacifist sect and bloody industry.

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/102-moby-dick.md`
- Chunk: 6 of 17, lines 4503–5517
- Heading path: `Moby-Dick > Retrieved Text`
- Coverage: Continues directly from the previous chunk, covering the Pequod’s departure through the beginning of Cetology. The text picks up mid‑paragraph after a chapter break (“But there was not much chance to think…”) and ends part way through Chapter 32 (Cetology). No explicit marker of truncation.

## Local Summary
The Pequod gets under way with wildly contrasting pilots: pious Bildad sings psalms while Peleg swears and kicks. After a freezing night, the pilots depart emotionally. The ship plunges into the Atlantic. A brief elegy for Bulkington frames the lee shore as the soul’s treacherous temptation—better to perish in the open sea than cravenly cling to land. The narrator launches a long vindication of whaling’s dignity (The Advocate, Postscript), citing history, economics, exploration, and even royal coronation oil. The senior crew is introduced: Starbuck (earnest, careful, superstitious), Stubb (easy-going, perpetual pipe‑smoker), Flask (pugnacious, unconscious of danger). Their harpooneers are Queequeg (Starbuck’s), Tashtego the Gay‑Head Indian (Stubb’s), and Daggoo the gigantic African (Flask’s). The crew are described as “Isolatoes” federated along one keel. A brief mention of Pip, the Alabama cabin‑boy, prefigures his fate. Captain Ahab finally appears: a bronzed, branded figure with an ivory leg, a “crucifixion in his face,” and an unsurrenderable forward gaze. He withdraws, then gradually lives on deck. Stubb’s request for muffling the ivory leg provokes a savage rebuke (“Down, dog, and kennel!”). Stubb’s queer dream turns the kick into an honour, and he decides to leave Ahab alone. Ahab smokes his pipe, finds it no longer soothing, and casts it into the sea. Stubb recounts his dream to Flask. Ahab shouts from the mast‑head for a lookout, singling out a white whale. The chunk ends part‑way through Cetology, as the narrator begins a systematic classification of whales, quoting Scoresby and Beale on the utter confusion of the field.

## Key Claims
- Bildad piloted only vessels in which he had a financial interest, to save the Nantucket pilot‑fee.
- Bildad forbade profane songs, yet the crew sang a rowdy chorus under his psalmody.
- The lee shore metaphor: the land, which offers safety, is the ship’s direst jeopardy; true independence is found in the “landlessness” of the sea.
- Whaling is unjustly scorned as unpoetical; its perils and defilements are no worse than those of war, and it is a pioneer of exploration, commerce, and empire.
- Whaling opened the Pacific Spanish colonies, contributed to the liberation of Peru, Chile, and Bolivia, and first settled Australia.
- The whale is a “royal fish” by English statute. Sperm oil is suggested as the coronation anointing oil.
- Starbuck’s courage is practical and cautious; “I will have no man in my boat who is not afraid of a whale.”
- Starbuck’s fortitude could fail before “spiritual terrors” like an enraged man.
- Stubb’s pipe and constant smoking act as a disinfectant against mortal tribulations.
- Flask views whales as “magnified mice” requiring only a little circumvention.
- Most whalemen before the mast are foreign‑born; officers are native‑born Americans.
- Whalemen are “Isolatoes”—each living on a separate continent of his own, yet federated along one keel.
- Ahab is branded by a livid mark from grey hair to neck; an old Manxman superstitiously claims Ahab bears a birth‑mark “from crown to sole.”
- Ahab’s ivory leg is made from a sperm whale’s jaw; he has “a quiver of ’em.”
- Ahab’s pipe no longer soothes him, and he throws it into the sea.
- Ahab explicitly commands the lookout to cry out if a white whale is seen.
- Cetology is a field of “utter confusion” (Scoresby, Beale).

## Entities And Concepts
- **Pequod** – whaler commanded by Ahab, departing Nantucket.
- **Peleg and Bildad** – Quaker part‑owners and pilots; Bildad pious, Peleg profane.
- **Starbuck** – chief mate, Nantucket Quaker, thin, conscientious, superstitious, careful.
- **Stubb** – second mate, Cape‑Cod‑man, good‑humored, pipe‑smoker.
- **Flask (King‑Post)** – third mate, Tisbury native, pugnacious, ignorant of danger.
- **Queequeg** – Starbuck’s harpooneer.
- **Tashtego** – Stubb’s harpooneer, unmixed Gay‑Head Indian.
- **Daggoo** – Flask’s harpooneer, giant African with gold hoop earrings.
- **Pip** – “Black Little Pip,” Alabama cabin‑boy, mentioned proleptically.
- **Bulkington** – tall mariner, briefly seen at the helm; his “stoneless grave” is the six‑inch chapter.
- **Captain Ahab** – main figure, bronze complexion, ivory leg, livid scar, crucifixion‑like woe.
- **Lee Shore** – emblem of the soul’s dangerous temptation to abandon the open independence of the sea.
- **The Advocate** – narrator’s defence of whaling’s dignity, citing history, commerce, and famous figures.
- **Royal fish** – English statutory term for the whale.
- **Isolatoes** – term for crew members who exist as separate continents, united by the ship.
- **Cetology** – the systematic classification of whales, introduced at the chunk’s end.
- **White Whale** – first explicit command by Ahab to watch for a white whale.

## Procedures And API Details
None (literary narrative).

## Nuance Or Contradictions
- The chunk begins mid‑sentence (“But there was not much chance…”) from a prior chapter, illustrating the seamless narrative flow of the source.
- The source ends part‑way through Chapter 32 (Cetology), after quoting Scoresby and Beale; no explicit [truncated] marker is present, but the text stops abruptly mid‑discussion of taxonomic difficulties. The full argument of Cetology is not included in this chunk.
- The tone shifts from the poetic lee‑shore elegy to the rhetorical defence of whaling, then to character sketches and Ahab’s ominous presence, reflecting the novel’s encyclopaedic and digressive structure.
- Stubb’s dream recasts Ahab’s kick as an honour, foreshadowing the dynamic of capitulation and obsession among the crew.

## Candidate Wiki Hints
- **The Pequod (ship)**: its departure, officers, and initial crew.
- **Starbuck, Stubb, Flask**: character notes with defining traits and attitudes toward whaling.
- **Queequeg, Tashtego, Daggoo**: the harpooneers and their chiefs.
- **Captain Ahab**: initial appearance, physical description, ivory leg, and first command regarding the white whale.
- **The Lee Shore (symbol)**: Melville’s metaphor for the soul’s dangerous longing for safety.
- **The Advocate (Moby‑Dick chapter)**: arguments for the dignity and historical significance of whaling.
- **Cetology (Moby‑Dick chapter)**: beginnings of the classification of whales, with authorities cited.
- **Isolatoes**: concept of isolated selves united in a common venture.
- **Bulkington**: brief chapter as the “stoneless grave” and apotheosis.
- **Pip**: first mention, foreshadowing his later narrative role.

## chunk-07

---
title: Chunk 07 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context

- Source file: `raw/web/corpus-2026-05-18/102-moby-dick.md`
- Heading path: **Moby-Dick > Retrieved Text**
- Chunk lines: 5519–6487
- This chunk covers the later part of the cetology chapter (whale classification), Chapter 33 (The Specksnyder), Chapter 34 (The Cabin-Table), and Chapter 35 (The Mast-Head).

## Local Summary

The narrator offers a playful, self‑conscious classification of whales based on size (Folio, Octavo, Duodecimo), declares the sperm whale the true monarch of the seas, and defines “whale” as a spouting fish with a horizontal tail. He then describes the peculiar rank of the harpooneer (Specksnyder), the rigid and silent hierarchy at Ahab’s cabin table, the boisterous licence of the harpooneers at their meal, and finally the mast‑head vigil, warning against dreamy, philosophical look‑outs and praising the practical crow’s‑nest of Captain Sleet.

## Key Claims

- Most earlier whale authors never saw a living whale; only Captain Scoresby had real whaling experience, but he knew only the Greenland (right) whale, not the sperm whale.
- The sperm whale is the true monarch of the seas, not the Greenland whale; the sperm whale’s life remains unwritten.
- The classification proposed (Folio, Octavo, Duodecimo) is deliberately incomplete—“a draught of a draught”—and based on size, not anatomy.
- A whale is defined as **a spouting fish with a horizontal tail**; the narrator sides with Jonah and old‑fashioned usage, rejecting Linnæus’s separation of whales from fish.
- The “Specksnyder” (Chief Harpooneer) once held co‑equal command in Dutch whaling but is now merely a senior harpooneer.
- Aboard the Pequod, the cabin table is a place of oppressive silence and rigid precedence, while the harpooneers dine afterwards with savage, democratic freedom.
- Ahab is inaccessible, a soul shut in the caved trunk of his body.
- Mast‑head duty is a sublime monotony, but dangerous for dreamy young Platonists who miss whales; the Southern fishery lacks the Greenland crow’s‑nest comforts, though Captain Sleet’s crow’s‑nest included a case‑bottle he never mentions in his official account.

## Entities And Concepts

- **Sperm Whale** (Folio, Chapter I) – largest, most valuable, source of spermaceti; name considered philologically absurd.
- **Right Whale** (Greenland whale, Mysticetus) – first regularly hunted; yields baleen and common “whale oil.”
- **Fin‑Back, Hump‑Back, Razor Back, Sulphur Bottom** – other Folio whales, described briefly.
- **Grampus, Black Fish, Narwhale, Thrasher, Killer** – Octavo whales.
- **Huzza Porpoise, Algerine Porpoise, Mealy‑mouthed Porpoise** – Duodecimo whales.
- **Specksnyder** (Specksioneer) – the chief harpooneer, once co‑commander, now subordinate.
- **Ahab** – moody, mute captain; uses forms and usages as instruments of dictatorship.
- **Starbuck, Stubb, Flask** – the three mates, each introduced in the cabin‑table ritual.
- **Dough‑Boy** – the trembling steward.
- **Queequeg, Tashtego, Daggoo** – the harpooneers, whose meal is a raucous contrast.
- **Mast‑head** – the lookout station; criticized as uncosy; Southern whale‑ships lack crow’s‑nests.
- **Captain Sleet’s crow’s‑nest** – a patented, well‑stocked enclosure, described with ironic admiration.
- **Childe Harold** – quoted as a warning of the useless, meditative mast‑head stander.

## Procedures And API Details

No API or procedural content is present in this chunk.

## Nuance Or Contradictions

- The narrator’s classification proudly offers no completion; he prefers a “draught” and wishes God to keep him from ever completing anything.
- The whale is defined as a fish, contrary to Linnaean taxonomy, on the authority of Jonah and common‑sense external traits.
- The description of Captain Sleet’s crow’s‑nest notes a deliberate omission (the case‑bottle) that the narrator considers a faithful friend.
- The mast‑head chapter ends with a poetic quote from Childe Harold; the raw source does not show a truncation marker and the chapter conclusion appears intact.

## Candidate Wiki Hints

- **Cetology in Moby‑Dick** – the Folio/Octavo/Duodecimo system and the definition of a whale.
- **Sperm Whale** – its commercial value, supersedence of the Greenland whale, and the unwritten life.
- **Specksnyder** – history of the harpooneer rank and its transformation from co‑command to subaltern.
- **Cabin‑Table** – Ahab’s silent, authoritarian meal and the harpooneers’ savage post‑meal liberty.
- **Mast‑Head** – the philosophy of the lookout, dangers of reverie, and contrast with Captain Sleet’s crow’s‑nest.
- **Ahab** – his inaccessibility, sultanism, and the use of sea‑forms as instruments of absolute will.

## chunk-08

---
title: Chunk 08 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/102-moby-dick.md`
- Chunk: 8 of 17, lines 6489–7569
- Heading: Moby-Dick > Retrieved Text
- Covers the end of a philosophical reflection on the dreamy indifference of young philosophers at sea, then Chapter 36 through Chapter 41 and part of Chapter 42.
- Contains the famous Quarter-Deck scene, Ahab’s binding of the crew, the internal monologues of Ahab, Starbuck, and Stubb, a midnight forecastle revelry broken by a squall, Ishmael’s historical account of Moby Dick and Ahab’s monomania, and the opening of “The Whiteness of the Whale.”
- The chunk stops at the end of a paragraph with an asterisk (`*`) marking a footnote not present in the raw source. **Truncation marker:** The text ends with “white-shrouded bear or shark.*” — the raw source breaks there; the remainder of Chapter 42 is missing.

## Local Summary
Ahab assembles the ship’s company, nails a doubloon to the mainmast, and offers it as a reward for whoever raises the white whale Moby Dick, described by the harpooneers. He reveals that Moby Dick took his leg, and swears to chase the whale around the globe. Starbuck objects that vengeance on a dumb brute is blasphemy; Ahab retorts that all visible objects are pasteboard masks hiding a malicious power he must strike through. The crew drinks from harpoon sockets in a ritual oath: “Death to Moby Dick!” Ahab, alone at sunset, declares his will is iron-rails; Starbuck, at the mainmast, grieves his forced complicity; Stubb laughs and resigns himself to predestination. The midnight forecastle section shows sailors of many nations singing, dancing, arguing, then scattering before a squall. Ishmael narrates the history and terror of Moby Dick, detailing Ahab’s monomania—how the whale became the embodiment of all evil—and begins to explore the horror of the whale’s whiteness, but the source breaks at a footnote marker.

## Key Claims
- Ahab identifies Moby Dick by a white head, wrinkled brow, crooked jaw, and three holes in the starboard fluke.
- He announces the chase is the ship’s true purpose, not mercantile whaling.
- Starbuck argues that vengeance on a “dumb brute” that struck from blind instinct is madness and blasphemy.
- Ahab claims that all visible objects are pasteboard masks; behind them is an inscrutable, reasoning power, and the white whale is the mask thrust nearest him. He will strike through it.
- Ahab states, “Truth hath no confines,” and asserts his will is fixed and irresistible.
- Starbuck feels his reason has been blasted out of him, yet he is bound to help Ahab; he hopes God may “wedge aside” Ahab’s purpose.
- Stubb declares that a laugh is the wisest answer to all that is queer, comforted by predestination.
- The crew, a “mongrel” assembly of renegades and castaways, is described as morally enfeebled and seemingly picked by infernal fatality to aid Ahab’s revenge.
- Moby Dick is reported as ubiquitous, possibly immortal, and possesses “unexampled, intelligent malignity”; his treacherous retreats are especially dreaded.
- Ahab’s monomania is traced: his bodily and spiritual agonies interfused during his long convalescence, turning all the evil of the world into a personified target in Moby Dick; his sanity remains but is now a tool of his madness.
- The whiteness of the whale appalls Ishmael more than anything else; the color, despite its positive associations, amplifies terror when coupled with a terrible object (e.g., polar bear, white shark). The text is truncated mid-footnote.

## Entities And Concepts
- **Ahab**: captain of the _Pequod_, leg amputated by Moby Dick, driven by monomaniac revenge.
- **Moby Dick**: the white sperm whale with a wrinkled forehead, crooked jaw, and three harpoon marks; reputed ubiquitous, malicious, immortal in superstition; called “the monomaniac incarnation of all those malicious agencies” by Ahab.
- **Starbuck**: chief mate, a Nantucketer who sees the madness but feels bound; calls vengeance blasphemous.
- **Stubb**: second mate, laughs at horror and relies on predestination.
- **Flask**: third mate (mentioned as “pervading mediocrity”).
- **Harpooneers**: Tashtego, Daggoo, Queequeg—each recalls distinct details of Moby Dick’s appearance.
- **The doubloon**: a sixteen-dollar Spanish ounce nailed to the mast as reward.
- **Harpoon sockets as chalices**: used for the “murderous” oath-drinking.
- **Pasteboard mask**: Ahab’s metaphor for the visible world hiding a reasoning, malicious force.
- **Iron Crown of Lombardy**: Ahab’s metaphor for his tormenting, split inner crown.
- **Monomania**: Ahab’s concentrated, unrevealed madness that uses his intellect as an instrument.
- **Ubiquity and immortality** of Moby Dick: whalemen’s superstitious belief, supported by the whale’s ability to surface far apart within days.
- **Whiteness**: first introduced as a source of nameless horror, contrasted with its cultural associations of purity and majesty.

## Procedures And API Details
- **Nailing the doubloon**: Ahab asks Starbuck for the top-maul, rubs the coin for luster, then nails it to the mainmast as a visual pledge and reward.
- **Oath‑taking ritual**: the mates cross their lances at the centre; Ahab touches the crossed lances and then fills the upturned harpoon sockets with grog; the harpooneers drink from the barbed goblets; the crew passes the pewter, swearing “Death to Moby Dick!”
- **Call to general quarters**: Ahab orders “Send everybody aft,” an order rarely given except in emergencies.
- **Squall reaction**: the mate calls “Hands by the halyards! in top-gallant sails! Stand by to reef topsails!” and the crew scatters.

## Nuance Or Contradictions
- Ahab’s rhetoric shifts between acknowledging that the whale may be merely an agent (“be the white whale agent, or be the white whale principal”) and treating it as the wall he must strike through; there is ambiguity about whether he seeks to destroy a symbol or a literal animal.
- Starbuck’s objection (“Vengeance on a dumb brute! … seems blasphemous”) is never fully answered; Ahab deflects by declaring he would strike the sun if it insulted him, and then co-opts Starbuck.
- The narrative suggests Ahab’s monomania is a conscious instrument (“all my means are sane, my motive and my object mad”) yet also describes it as an involuntary contraction of his broader madness into a single channel.
- Ishmael admits he participated wholeheartedly in the oath, driven by a “wild, mystical, sympathetical feeling,” complicating his later analytical tone.
- The chapter “The Whiteness of the Whale” begins with an elaborate list of white’s positive connotations only to argue that a deeper panic lurks behind the hue; the argument is announced but the source truncates at the first footnote, leaving the full explanation incomplete in this chunk.
- **Truncation**: The raw source ends at the asterisk after “white-shrouded bear or shark.*” — the footnote content is not included, and the sentence breaks off. The remainder of Chapter 42 is missing from this chunk.

## Candidate Wiki Hints
- **Moby Dick (whale)** – physical description, reported capacities (ubiquity, malice), lore among whalemen.
- **Ahab’s monomania** – origin, expression, the pasteboard‑mask philosophy, the fusion of sanity and madness.
- **The Quarter‑Deck ritual** – the nailing of the doubloon, the oath with harpoon‑socket chalices.
- **Crew of the Pequod** – hierarchical structure (mates, harpooneers, hands), their responses to Ahab’s quest.
- **Symbolism of whiteness** – chapter introduction, contrast between cultural exaltation and primal dread, truncated footnote.
- **Ishmael’s narrative voice** – participation in the oath vs. retrospective analysis, limitations of understanding stated (“to dive deeper than Ishmael can go”).

## chunk-09

---
title: Chunk 09 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- **Source**: `raw/web/corpus-2026-05-18/102-moby-dick.md`
- **Chunk**: 9 of 17, lines 7571–8560
- **Heading path**: Moby-Dick > Retrieved Text
- **Coverage**: Continuation from the middle of Chapter 42 (“The Whiteness of the Whale”) through Chapter 47 (“The Mat-Maker”). The chunk begins partway through the whiteness meditation and ends with the close of Chapter 47.

## Local Summary
The chunk moves from a philosophical treatise on the terror of whiteness (Chapter 42), through a brief shipboard mystery (Chapter 43), to Ahab’s methodical plotting of whale migrations (Chapter 44), followed by Ishmael’s accumulation of documented evidence for individual whale vengeance and ship-sinking power (Chapter 45), Ahab’s calculated pragmatism to keep his crew engaged with ordinary whaling (Chapter 46), and finally a reflective scene of Ishmael and Queequeg weaving a sword-mat that becomes a metaphor for fate, free will, and chance (Chapter 47).

## Key Claims
- Whiteness, when stripped of positive associations, intensifies terror and appals the mind through its indefiniteness and suggestion of annihilation.
- Whiteness is both the visible absence of colour and the concrete of all colours, which evokes a “colourless, all-colour of atheism.”
- The albino whale is the symbol of these qualities, making Ahab’s hunt a fiery pursuit against a cosmic blankness.
- Sperm whales follow predictable migratory “veins” and can be charted by season, making it possible for Ahab to increase his chances of encountering Moby Dick.
- Individual sperm whales are long-lived, become famous for their dangerousness, and have been recognized across years and oceans.
- Sperm whales are capable of deliberate, calculated attacks on ships, as evidenced by multiple historical incidents (Essex, Union, Langsdorff’s account, Procopius’s sea-monster).
- Ahab consciously maintains ordinary commercial whaling activities to keep the crew compliant, prevent mutiny, and satisfy their “sordidness” for cash.
- The act of weaving a sword-mat prompts Ishmael to envision the “Loom of Time,” interweaving necessity (the warp), free will (the shuttle), and chance (Queequeg’s indifferent sword strokes) as jointly shaping events.

## Entities And Concepts
- **Ishmael**: The narrator, philosopher, and weaver of the mat.
- **Ahab**: Captain of the Pequod, methodically plotting whale migration on sea charts; consumed by monomania, yet shrewdly managing his crew.
- **Queequeg**: Harpooneer, partner in weaving the sword-mat; his casual sword strokes represent chance.
- **Archy & Cabaco**: Sailors who hear mysterious coughs below deck, suspecting a stowaway (Chapter 43).
- **Starbuck**: Chief mate, inwardly opposed to Ahab’s quest; Ahab manipulates him through ordinary whaling duties.
- **Moby Dick**: The White Whale, symbol of whiteness and terror, with distinctive physical marks (snow-white brow, hump, bored and scalloped fins).
- **Famous whales**: Timor Tom, New Zealand Jack, Morquan (King of Japan), Don Miguel — celebrated for their ferocity and individuality.
- **Whiteness examples**: Polar bear, white shark (*Requin*), albatross, White Steed of the Prairies, albino man, White Squall, White Hoods of Ghent, pallor of death, white shroud/ghosts, Whitsuntide, White Friar/Nun, White Tower of London, White Mountains, White Sea, tall pale man of the Hartz forest, Lima’s white veil, milky sea, snowy prairies, the Andean snow, Antarctic ice.
- **Season-on-the-Line**: The specific equatorial season and region where Moby Dick is most consistently sighted.
- **Propontis sea-monster**: Procopius’s historical account of a whale that destroyed ships for over 50 years, cited as evidence of whale longevity and malice.
- **Sword-mat / Loom of Time**: The woven mat as a model of necessity, free will, and chance combined.

## Procedures And API Details
- **Whale migration charting (Ahab’s method)**: Using sea charts, old logbooks, knowledge of tides, currents, and sperm whale feeding grounds to plot likely whale locations; marking “veins” — linear oceanic paths of migrating sperm whales of a few miles’ width.
- **Sword-mat weaving**: Described as passing filling (woof) between warp yarns with a hand shuttle, while a wooden sword beats down the weft; the process becomes a metaphysical analogy.

## Nuance Or Contradictions
- The whiteness meditation moves abruptly from poetic awe to philosophical dread, then shifts to anecdotal shipboard chatter and dense documentary citation. This juxtaposition may reflect Ishmael’s struggle to find a unified account of the whale.
- The chapter “The Affidavit” explicitly acknowledges the need to bolster truth with evidence, because the landsman’s ignorance might lead to disbelief in the story.
- No explicit truncation marker is present; the chunk ends at a natural chapter boundary.

## Candidate Wiki Hints
- **The Whiteness of the Whale (chapter)** — a dense philosophical set-piece on colour, terror, and symbolism.
- **Ahab’s Charts** — his methodical use of migration data and “veins” to hunt a single whale.
- **Famous Individualised Whales** — known named whales with histories (Timor Tom, etc.) and the evidence for their recognition over time.
- **Whale Attacks on Ships** — documented instances of sperm whales ramming and sinking vessels (Essex, Union, Langsdorff, Procopius).
- **The Sword-Mat and Fate** — Ishmael’s metaphor of weaving as the interplay of necessity, free will, and chance.

## chunk-10

---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- **Source path:** `raw/web/corpus-2026-05-18/102-moby-dick.md`
- **Chunk:** 10 of 17
- **Lines:** 8562–9576
- **Heading path:** Moby-Dick > Retrieved Text
- **Content span:** From “Thus we were weaving and weaving away…” through Chapter 53 and into the opening of Chapter 54 (“The Town-Ho’s Story”)
- **Chapters covered:** 48 (The First Lowering), 49 (The Hyena), 50 (Ahab’s Boat and Crew. Fedallah), 51 (The Spirit-Spout), 52 (The Albatross), 53 (The Gam), start of 54.

## Local Summary
Tashtego sights a school of sperm whales, triggering the first lowering of the boats. Ahab reveals his secret boat crew—Fedallah and five “tiger-yellow” men—stowed aboard. The boats give chase; Starbuck’s boat gets fast to a whale but is swamped by a squall, narrowly escaping. Ishmael, in a detached “hyena” mood, writes his will and views the voyage as a cosmic joke. Ahab’s personal boat and its modifications are explained; Fedallah remains an enigmatic figure. A nocturnal “spirit-spout,” possibly Moby Dick, repeatedly appears and lures the ship onward. The Pequod encounters the bleached, spectral whaler *Goney* (Albatross); Ahab’s attempt to ask about the White Whale is frustrated. The custom of the **Gam**—a social meeting between whaleships—is defined. The chunk closes with the first sentence of “The Town-Ho’s Story.”

## Key Claims
- Sperm whales blow with “undeviating and reliable uniformity,” like a clock tick, a trait whalemen use to distinguish the species.
- Ahab’s own boat crew (Fedallah and five Manilla‑region men) were hidden stowaways, brought aboard in secret before the ship sailed.
- During the chase, Starbuck’s boat harpoons a whale but is swamped in a squall; the crew survives by clinging to the swamped boat until rescued by the *Pequod*.
- Ishmael’s “hyena” philosophy transforms peril and death into a joke; he writes his will and feels “supplementary clean gain” of life after each brush with death.
- Ahab had privately outfitted his boat with a thigh‑cleat and extra sheathing, anticipating his active role in the hunt, despite knowing the owners would not have approved.
- Fedallah’s origin and his uncanny link to Ahab remain unexplained; the crew views him as a phantom‑like figure from “primal generations.”
- The “Spirit‑Spout” appears by moonlight on multiple nights, always ahead of the ship; some seamen believe it is Moby Dick treacherously luring them onward.
- In a storm, Ahab resists sleep, sitting erect with a chart and the tell‑tale compass; Starbuck observes him seemingly asleep yet still fixed on his purpose.
- A **Gam** is a formal, friendly meeting between whaleships where captains and chief mates exchange visits; the captain must stand in the whaleboat (no seat, no tiller) and maintain “leg‑dignity.”
- When the *Pequod* encounters the whaler *Goney* (Albatross), Ahab’s speaker‑trumpet falls into the sea before he can receive an answer about the White Whale; he later cries “Round the world!”

## Entities And Concepts
- **Tashtego** – Gay‑Head Indian, harpooneer; his wild cadence while calling a sighting.
- **Fedallah** – Tall, swart, white‑turbaned Parsee; leader of Ahab’s secret crew; remains a “muffled mystery.”
- **Ahab** – Fits his own boat, pushes into the chase, obsessed with the White Whale, sleepless during storms.
- **Starbuck** – First mate, cautious yet drives on; his boat swamped.
- **Stubb** – Second mate; peculiar, jocular‑furious style of coaxing his crew, “religion of rowing.”
- **Flask (King‑Post)** – Third mate; stands on loggerhead, then on Daggoo’s shoulders; excitable.
- **Queequeg** – Harpooneer in Starbuck’s boat; stands at bow.
- **Daggoo** – Gigantic negro harpooneer; offers his shoulders as a lookout platform for Flask.
- **Ishmael** – Narrator; adopts “hyena” mood, writes his will.
- **Archy / Cabaco** – Sailors who suspected stowaways.
- **The *Pequod*** – Whale‑ship; rescue after squall; ghostly journey.
- **The *Goney* (Albatross)** – Bleached, spectral Nantucket whaler met near the Crozetts.
- **White Whale (Moby Dick)** – Object of Ahab’s quest; believed by some to be the Spirit‑Spout.
- **Gam** – Formal social meeting between whaleships; defined lexically.
- **Spirit‑Spout** – Mysterious silvery jet appearing repeatedly by night, pulling the ship onward.
- **Hyena mood** – Ishmael’s philosophical state: viewing life, death, and the voyage as a vast practical joke.
- **Captain’s spare boat** – Ahab’s extra boat, secretly prepared.
- **Tell‑tale compass** – Cabin compass enabling captain to know the ship’s course without going on deck.
- **Loggerhead** – A stout post in the boat’s stern; used as a precarious stand‑point.
- **Thigh‑board / cleat** – Horizontal piece in the bow for bracing the knee when darting the harpoon.

## Procedures And API Details
- **Sperm whale spouting** – Blows as regularly as a clock tick; used as a field mark by whalemen.
- **Lowering the boats** – Tubs fixed, cranes out, mainyard backed; boats swing over the side and crews leap in as the boat drops.
- **Rowing technique** – Stubb’s “religion of rowing”: a mixture of barking commands and joking fury that compels oarsmen to pull harder.
- **Standing in the boat** – Harpooneers stand on the triangular raised box in the bow; the mate balances on the stern platform; Flask uses the loggerhead and later Daggoo’s shoulders as a lookout perch.
- **Harpooning** – At the critical moment the harpooneer stands ready; the mate whispers “There, there, give it to him!”
- **Gam protocol** – After exchanging hails, a boat’s crew rows the captain to the other ship; the captain must stand throughout because the whaleboat has no seat or tiller, and must maintain balance without holding on, often with hands in pockets for ballast.
- **Tell‑tale compass** – A cabin‑mounted compass that allows the captain to check the ship’s heading without returning to deck.

## Nuance Or Contradictions
- The “Spirit‑Spout” is an ambiguous phenomenon: some crew believe it is Moby Dick luring them to destruction; its repeated vanishings and reappearances remain unexplained, blending superstition with natural possibility.
- Ishmael’s “hyena” outlook—treating death as a joke and writing a will for “a supplementary clean gain” after surviving—seems at odds with the mortal peril he has just faced, creating a jarring, gallows‑humour philosophy.
- Ahab’s secret crew contradicts the owners’ expectations and normal whaling practice; his private modifications to the boat suggest careful, rational preparation, yet his ultimate goal is mad.
- Despite the extreme danger, Ahab’s physical disability is framed as a debated question among “whale‑wise” people; Flask even argues that one remaining knee makes it feasible.
- The end of the chunk is the opening sentence of Chapter 54, which is a complete thought; no explicit truncation marker appears, and the narrative does not cut off mid‑sentence.

## Candidate Wiki Hints
- **Fedallah** – Ahab’s mysterious Parsee harpooneer and his secret crew.
- **The Spirit‑Spout** – The uncanny night‑jet that seems to lure the Pequod onward.
- **The Gam** – The custom of social meetings between whaleships, its rules and etiquette.
- **Ishmael’s Hyena Philosophy** – The mood of cosmic jest and will‑making in the face of death.
- **The First Lowering** – Ahab’s first whale chase and the revelation of his hidden boat crew.

## chunk-11

---
title: Chunk 11 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context

- Source: raw/web/corpus-2026-05-18/102-moby-dick.md
- Lines: 9578-10613
- Heading: Moby-Dick > Retrieved Text
- This chunk begins with the aftermath of the Goney encounter and the introduction of the Town-Ho’s story, continues with the full narrative of the Town-Ho (as told by Ishmael in Lima), and concludes with Chapters 55 (“Of the Monstrous Pictures of Whales”) and 56 (“Of the Less Erroneous Pictures of Whales, and the True Pictures of Whaling Scenes”). The chunk ends naturally with the close of Chapter 56; no truncation marker is present.

## Local Summary

The narrator recounts how the Pequod learns of Moby Dick from the Town-Ho, a whaler manned mostly by Polynesians. A secret part of the Town-Ho’s story—unknown to the captain and mates—spreads among the Pequod’s crew but never reaches Ahab. Ishmael then presents the full Town-Ho account as he once narrated it to Spanish friends in Lima. The tale concerns a conflict between the Lakeman Steelkilt and the Vineyarder mate Radney aboard the leaky Town-Ho, culminating in a violent mutiny, a near-fatal beating, and Radney’s death in the jaws of Moby Dick. Afterwards, a lengthy digression critiques historical and scientific depictions of whales (Chapters 55–56), arguing that virtually all existing pictures are monstrously inaccurate, and that the only true impression of a living whale comes from witnessing one at sea.

## Key Claims

- The Town-Ho encountered Moby Dick; the whale killed the mate Radney during a boat chase.
- A secret aspect of the Town-Ho affair involved a “certain wondrous, inverted visitation of one of those so called judgments of God.”
- That secret was kept among a few Pequod seamen and never reached Ahab or his mates.
- Steelkilt, a Lakeman, and Radney, a Vineyarder, were natural antagonists; their conflict stemmed from Radney’s overbearing command to sweep and shovel after exhausting pump work.
- Steelkilt’s refusal led to Radney striking him with a hammer, whereupon Steelkilt stove in Radney’s jaw.
- The subsequent mutiny was contained when Steelkilt and his allies were locked in the forecastle; after days of starvation, most surrendered, and Steelkilt was betrayed by his two Canaller comrades.
- Radney attempted to flog Steelkilt, but Steelkilt whispered an unstated threat that made the captain relent.
- Later, Radney was tossed from a whaleboat onto Moby Dick’s back and was seized and drowned by the whale.
- Steelkilt afterwards deserted with most of the crew, seized a war-canoe, and eventually sailed to Tahiti and France.
- The narrator swears on the Holy Evangelists that the story is true in substance.
- All extant pictures of whales—from ancient Hindu sculptures to 19th-century scientific plates—are grossly inaccurate.
- The living whale cannot be accurately portrayed because it cannot be hoisted whole from the water, and the skeleton does not convey the true shape.
- The only way to gain a tolerable idea of the whale’s living contour is to go whaling, at the risk of being killed by the whale.
- Among published Sperm Whale outlines, Beale’s are the best; Garnery’s French engravings are the finest whaling scenes overall.

## Entities And Concepts

- **Town-Ho**: Nantucket sperm whaler, leaky, scene of mutiny and Moby Dick encounter.
- **Moby Dick**: The white whale, kills Radney; already famous among whalemen.
- **Steelkilt**: A “Lakeman” (from the Great Lakes region, Buffalo), desperado, tall, golden-bearded, leader of the mutineers, bent on revenge against Radney.
- **Radney**: Mate of the Town-Ho, a Vineyarder (Martha’s Vineyard), described as ugly, malicious, and a part-owner; killed by Moby Dick.
- **Canallers**: Boatmen of the Erie Canal, portrayed as lawless, picturesque rogues; two of them betray Steelkilt.
- **Tashtego**: Pequod harpooneer who receives and then sleep-talks the secret of the Town-Ho.
- **Lima frame narrative**: Ishmael tells the story to Spanish friends (Don Pedro, Don Sebastian) at the Golden Inn.
- **Matse Avatar**: Hindu sculpture in Elephanta depicting Vishnu as half-man, half-whale; cited as the oldest purported whale portrait, but inaccurate.
- **Guido, Hogarth, Sibbald, Goldsmith, Lacépède, Cuvier, Scoresby, Beale, Garnery**: Authors/artists whose whale depictions are evaluated.
- **Right Whale and Sperm Whale**: Distinct species; most pictures misrepresent both.

## Procedures And API Details

None.

## Nuance Or Contradictions

- The secret part of the Town-Ho story is said to involve a “wondrous, inverted visitation” of a judgment of God, but the text does not explicitly state what that judgment was; it remains obscure.
- The narrator’s sworn oath on the Evangelists in Lima frames the tale as factual, yet the entire work is a novel—blurring truth and fiction.
- The chunk ends mid-discussion on whale pictures (Chapter 56), with the admission that even the best anatomical details could be faulted, but the narrator cannot draw better.

## Candidate Wiki Hints

- A page on the **Town-Ho mutiny** could distill the Steelkilt–Radney conflict, the mutiny, and its connection to Moby Dick.
- A page on **Cetacean Depiction in Art and Science** could collect the criticisms from Chapters 55–56.
- A concept note on **Steelkilt** as a character type (Lakeman, Canaller, desperado).
- A note on the **frame narratives in Moby-Dick** (Lima inset tale) might be reusable.

## chunk-12

---
title: Chunk 12 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source path: `raw/web/corpus-2026-05-18/102-moby-dick.md`
- Heading path: Moby-Dick > Retrieved Text
- Lines: 10615–11666
- Covers chapters 56 through 64 in the retrieved text (from the end of “The Dutch and the French Pictures” through “Stubb’s Supper”). The chunk ends naturally at the close of Chapter 64.

## Local Summary
The chunk opens by describing two French engravings of whaling (by Garnery and “H. Durand”) and praising French artists for capturing the spirit of the hunt better than English or American draftsmen. Chapter 57 surveys whales depicted in many media: scrimshaw, wood, brass, sheet‑iron, natural rock formations, and constellations. Chapter 58 introduces vast meadows of brit upon which Right Whales feed, then broadens into a meditation on the terror and alienness of the sea compared with the land. Chapter 59 recounts a rare sighting of a giant squid, initially mistaken for Moby Dick; the narrator notes its importance as the sperm whale’s food. Chapters 60–63 give a meticulous technical account of the whale‑line, its arrangement, the harpooneer’s dangerous role, and the crotch that holds harpoons. Chapter 61 narrates Stubb’s chase and killing of a sperm whale. Chapter 64 shows the crew towing the carcass, Stubb ordering a whale‑steak for himself, and his comical, profane “sermon” to the sharks through the old black cook, Fleece.

## Key Claims
- French artists (Garnery, Durand) provide “the only finished sketches at all capable of conveying the real spirit of the whale hunt”; English and American drawings tend toward the mechanical and profile‑like.
- Sailors create scrimshaw (skrimshander) articles with jack‑knives during idle hours; a whale‑hunter is as much a “savage” as an Iroquois.
- The sea is a “fiend to its own off‑spring,” a place of cannibalism and hidden terror, while the soul of man contains an “insular Tahiti” surrounded by the “horrors of the half known life.”
- The great live squid is rarely seen and is believed to be the sperm whale’s only food; the Bishop Pontoppidan’s Kraken may partially correspond to the squid.
- The whale‑line, folded in complex coils, endangers every man in the boat; “All men live enveloped in whale‑lines.”
- The standard practice of having the harpooneer row strenuously and then throw the harpoon is described as “foolish and unnecessary”; exhaustion is a major cause of missed darts.
- The crotch holds two harpoons connected to the line, increasing the chance of a hold, but a second iron left dangling can become a “sharp‑edged terror.”
- Sharks gather in the greatest numbers around a dead sperm whale moored at night; Stubb’s address to the sharks, delivered by Fleece, mocks religious exhortation.

## Entities And Concepts
- **Garnery** – French painter of whaling scenes; praised for conveying motion.
- **H. Durand** – French engraver of two whaling pieces (one a calm Pacific anchorage, the other a cutting‑in scene).
- **Scoresby** – Right whaleman and author, depicted mechanical details but lacked picturesqueness.
- **Skrimshander** – Scrimshaw; carved sperm‑whale teeth, whale‑bone busks, and other items made by sailors.
- **Brit** – Minute yellow substance on which the Right Whale feeds; forms vast “meadows” at sea.
- **Squid (great live squid)** – A huge, formless, cream‑coloured mass; few ships see it and live to tell; reputed food of the sperm whale.
- **Kraken (Bishop Pontoppidan)** – Legendary sea‑monster partly equated with the squid.
- **Whale‑line** – Hemp or Manilla rope, 2/3 inch thick, over 200 fathoms long, coiled in a tub; attached to the harpoons and run around the loggerhead.
- **Loggerhead** – Post in the boat around which the line takes turns to slow the whale.
- **Headsman / boatheader** – The officer (usually the mate) who kills the whale with the lance; temporarily steers at the start of a chase.
- **Harpooneer / boat‑fastener** – The man who rows the foremost oar and throws the harpoon; then changes places with the headsman.
- **Crotch** – Notched stick about two feet long, inserted in the starboard gunwale, holding two harpoons ready for use.
- **Stubb** – Second mate of the *Pequod*; kills a sperm whale and later eats a steak at the capstan.
- **Fleece (the cook)** – Old black cook, called Fleece, who delivers Stubb’s mock sermon to the sharks.

## Procedures And API Details
- **Whale‑line arrangement**
  - Lower end: an eye‑splice hangs free at the tub’s side to allow connection to a neighbour boat’s line; it must never be attached to the boat.
  - Upper end: passes aft around the loggerhead, runs forward along the oars, through a groove at the prow, hangs in a festoon, then is coiled as box‑line and finally attached to the short‑warp and harpoon.
  - Coiling: line is carefully spiralled in the tub in concentric “sheaves” around a central “heart”; any tangle could sever a limb.
- **Harpooning sequence**
  - Boat pushes off with the headsman temporarily steering and the harpooneer pulling the foremost oar.
  - On command (“Stand up, and give it to him!”), the harpooneer drops his oar, turns, grabs a harpoon from the crotch, and darts it.
  - If the dart is successful, the headsman and harpooneer change places while the whale runs.
  - The line is wetted to prevent burning; extra turns around the loggerhead are taken.
- **Mooring a dead whale alongside**
  - The corpse is moored head to stern and tail to bows with chains.
  - A small line with a wooden float and weight is used to girdle the tail, allowing the chain to follow and be locked around the narrowest part.
- **Cutting a whale‑steak**
  - Stubb orders Daggoo to cut a steak “from his small” (the tapering extremity of the body); the meat is cooked and eaten at the capstan.

## Nuance Or Contradictions
- The chunk ends cleanly at the close of Chapter 64; there is no truncation or missing text.
- The narrator explicitly opposes the established fishery practice: he argues the headsman should stay in the bows and both dart the harpoon and the lance, eliminating the exhausted harpooneer’s rowing (Chapter 62). This contradicts the standard custom described in the same chapters.
- The description of the squid as “the largest animated thing in the ocean” yet “formless” and lacking a face is acknowledged as based on vague seamen’s reports; only detached arms have been directly observed.

## Candidate Wiki Hints
- **Whale‑line and whaling‑boat gear** – detailed description of line layout, tub, loggerhead, and safety practices.
- **Scrimshaw (skrimshander)** – sailor‑made carvings on teeth and bone, with notes on materials and tools.
- **French whaling art** – separate page on Garnery, Durand, and the contrast with English/American representations.
- **Brit and Right Whale feeding** – the “Brazil Banks” and the feeding method using baleen.
- **Sperm whale diet and the giant squid** – link to the Kraken legend and sparse observations.
- **Role of the harpooneer and headsman** – critique of traditional boat‑crew configuration.
- **Stubb’s character and Stubb’s supper** – comic episode with the cook and the sharks, illustrating his high spirits.

## chunk-13

---
title: Chunk 13 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- Source: `raw/web/corpus-2026-05-18/102-moby-dick.md`
- Lines: 11668–12726 (chunk 13 of 17)
- Heading path: **Moby‑Dick > Retrieved Text**
- Coverage: concludes the dialogue between Stubb and the cook Fleece from late in “Stubb’s Supper” (Chapter 64), then continues through Chapters 65–73: “The Whale as a Dish”, “The Shark Massacre”, “Cutting In”, “The Blanket”, “The Funeral”, “The Sphynx”, “The Jeroboam’s Story”, “The Monkey‑Rope”, and part of “Stubb and Flask kill a Right Whale; and Then Have a Talk over Him”.

## Local Summary
The chunk moves from comic shipboard talk (Stubb ordering the cook and teasing him about heaven) to a series of discursive chapters on eating whale, the gory mechanics of cutting‑in, the physiology and symbolism of blubber, the aftermath of the whale’s death, and a cluster of narrative episodes—the meeting with the plague‑stricken Jeroboam and its crazed prophet Gabriel, the mutual dependence symbolised by the monkey‑rope, and the killing of a Right Whale while Stubb and Flask discuss Fedallah’s diabolical aspect.

## Key Claims
- In the 16th century, Right Whale tongue was a delicacy in France, and a court cook was rewarded for a porpoise sauce under Henry VIII.
- Porpoises (a small whale) are still eaten; their meat is made into seasoned balls.
- Among hunters, whale would be a noble dish if portion sizes were not so vast; only “unprejudiced” men like Stubb eat it today, though the Esquimaux live on whale and train oil.
- Sperm‑whale brains, when cooked with flour, resemble calves’ head and are eaten by some.
- Eating a whale “by its own light” is considered outlandish; the narrator connects this to a general hypocrisy about eating animals.
- Sharks are attracted in “incalculable hosts” to a moored whale carcass; whalemen sometimes stir them with spades but often only increase their frenzy.
- During cutting‑in, the whale is rolled by the windlass while a long strip of blubber (“blanket‑piece”) is peeled off, hoisted to the main‑top, then severed and lowered into the blubber‑room.
- The skin of the whale is debated; Ishmael argues the blubber is the true skin, while a thin transparent outer layer is only “the skin of the skin”.
- The visible surface of the Sperm Whale carries oblique linear marks like Italian engravings, and sometimes “hieroglyphical” patterns that remain undecipherable.
- The whale’s blubber acts as a blanket, enabling it to keep warm even in icy seas; its blood is warmer than that of a “Borneo negro in summer”.
- After stripping the blubber, the headless white body is cast adrift, becoming a “most doleful and most mocking funeral” besieged by sharks and seabirds; sometimes mistaken for shoals and logged as a danger, perpetuating false beliefs.
- Ahab addresses the severed Sperm Whale head as a Sphinx, demanding it speak of the ocean’s secrets, but gets no answer.
- The Jeroboam carries a malignant epidemic; its captain refuses to board the Pequod. Among its crew is Gabriel, a former Shaker prophet who claims to be the archangel and has gained fearful sway over the crew.
- Gabriel declares Moby Dick to be the incarnated Shaker God and warns against attacking him; a mate named Macey is killed by a blow from the whale, confirming Gabriel’s terrifying authority.
- Ahab tries to pass a letter for the dead Macey; Gabriel intercepts it, impales it on a boat‑knife, and hurls it back.
- The “monkey‑rope” ties the bowsman (Ishmael) to Queequeg while the harpooneer works on the whale’s back, so that if Queequeg sinks, Ishmael would be dragged down as well—a symbol of “the precise situation of every mortal”.
- Stubb’s innovation on the Pequod is to tie the monkey‑rope to the holder’s belt as well, making them a “Siamese ligature”.
- After the Sperm Whale head is hoisted, a Right Whale is spotted and, against custom, the mates kill it so that both a Sperm Whale head on the starboard and a Right Whale head on the larboard may provide a charm against capsizing, as allegedly stated by Fedallah.
- Stubb and Flask distrust Fedallah; Stubb speculates he is the devil, hiding his tail in his pocket, and that Ahab is bargaining his soul for Moby Dick.

## Entities And Concepts
- **Stubb** – second mate; delivers comic and philosophical lines to Fleece and later to Flask.
- **Fleece (the cook)** – old black cook; his ideas of heaven and the whale steak are mocked.
- **Whale as a dish** – historical and contemporary consumption of whale, porpoise, and blubber.
- **Cutting‑in** – process of stripping blubber from the carcass using tackles, windlass, boarding‑swords, and spades.
- **Blanket‑piece** – the long upper strip of blubber peeled off; also the term for the blubber as the whale’s insulating coat.
- **Skin of the whale** – debate between blubber as skin vs. a thin outer isinglass‑like layer.
- **Hieroglyphics on the whale** – mysterious linear marks and scratches on the sperm whale’s surface.
- **Whale’s funeral** – the drifting, headless carcass surrounded by predators, sometimes mistaken for rocks.
- **Sphynx** – Ahab’s apostrophe to the severed head as a silent oracle.
- **Jeroboam** – Nantucket whaler with an epidemic; carries Gabriel the archangel.
- **Gabriel** – crazed Shaker prophet who dominates the Jeroboam’s crew, warns against Moby Dick.
- **Macey** – Jeroboam’s chief mate killed by the White Whale.
- **Monkey‑rope** – safety line that binds Ishmael to Queequeg; metaphor for mutual human dependency.
- **Queequeg** – harpooneer working on the whale’s back, protected by monkey‑rope and spades.
- **Right Whale** – captured alongside the Sperm Whale, supposedly for a charm.
- **Fedallah** – mysterious personage on the Pequod; Stubb claims he is the devil and that Ahab is bartering with him.

## Procedures And API Details
- **Cutting‑in** (Ch. 67):
  - The enormous lower blubber‑hook (≈100 lbs) is attached to a tackle swung over the whale.
  - A hole is cut above a side‑fin; a semicircular “scarf” line is cut round it.
  - The windlass heaves the strip; the whale rolls, and the blubber peels off like spiralling an orange rind.
  - The “blanket‑piece” is hoisted to the main‑top, then severed by a boarding‑sword and lowered into the blubber‑room to be coiled.
  - Two tackles work alternately: one hoists a new strip, the other lowers the previous strip.
- **Beheading** (Ch. 70):
  - Sperm Whale decapitation is a feat where the surgeon cuts through the thickest part of the body from above, steering clear of vital parts to divide the spine at the critical point near the skull.
  - Stubb boasted he needed only ten minutes to behead a sperm whale.
- **Monkey‑rope** (Ch. 72):
  - A rope is fastened at both ends: one end to a strong canvas belt around the harpooneer, the other to the holder’s belt.
  - In the Pequod, Stubb’s improvement ties the rope to the holder as well, ensuring mutual entanglement.

## Nuance Or Contradictions
- Ishmael admits that his conclusion about the whale’s skin (blubber = true skin) is “only an opinion”.
- The thin outer layer is described as transparent and flexible like isinglass, yet the chapter dismisses it as “skin of the skin” in an almost paradoxical argument against its delicacy.
- The Jeroboam’s captain refuses direct contact due to the epidemic, yet communication continues via a boat kept at a distance—a land‑style quarantine undercut by the open air.
- Stubb and Flask’s conversation implies a supernatural pact (Fedallah as devil, Ahab’s soul in exchange for Moby Dick) that remains ambiguous, introduced as speculation.

## Candidate Wiki Hints
- **Whale as a Dish** – historical notes on eating whale, porpoise, blubber, and train oil; the philosophical point about cannibalism and hypocrisy.
- **Cutting‑In** – the staged process, equipment (tackles, boarding‑sword, blubber‑hook), and roles.
- **The Blanket** – the debate over whale skin, insulating properties of blubber, and the hieroglyphics on the sperm whale.
- **The Funeral** – the imagery of the abandoned carcass, vultureism of scavengers, and the mistaken log entry as false precedent.
- **The Jeroboam’s Story** – Gabriel’s rise, Macey’s death, the epidemic, and the letter episode.
- **Monkey‑Rope** – equipment, Stubb’s improvement, and the philosophical reading of human interdependence.
- **Fedallah as Devil** – Stubb’s suspicions and the idea of a supernatural bargain.

## chunk-14

---
title: Chunk 14 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
- **Source:** `raw/web/corpus-2026-05-18/102-moby-dick.md`
- **Chunk:** 14 of 17
- **Lines:** 12728–13761
- **Heading path:** Moby-Dick > Retrieved Text
- **Coverage:** The end of Chapter 73 (dialogue between Stubb and Flask about Fedallah/the devil, the two whale heads hoisted) and full Chapters 74–81 (through the early stage of the chase in Chapter 81).

## Local Summary
The chunk opens with Stubb telling Flask a story of the devil visiting an old governor, then swerves to a discussion of Fedallah’s possible diabolical nature and age. The two whale heads are hung from the Pequod, giving Ishmael the opportunity to contrast the sperm whale’s and right whale’s heads in minute anatomical and philosophical detail. He examines the eyes (two independent fields of vision), ears (almost invisible), the sperm whale’s battering‑ram of a forehead, and the internal structure of the head: the upper “Case” (the Heidelburgh Tun) containing pure spermaceti, and the lower “junk.” The operation of baling the case is described, during which Tashtego falls into the nearly empty tun and is rescued by Queequeg’s “obstetric” diving. The narrative then moves to physiognomy and phrenology of the sperm whale (the forehead, the tiny brain, the spinal‑cord theory of character) before turning to the encounter with the German whaler Jungfrau. Captain Derick De Deer comes begging for lamp oil; a pod of eight whales is sighted, and a fierce race for a sick, old, fin‑less bull ensues among the Pequod’s three boats and Derick’s. The chunk ends mid‑chase, with the four boats in the whale’s immediate wake.

## Key Claims
- In Stubb’s yarn the devil kidnaps a man named John and gives him Asiatic cholera; the old governor lets the devil take him.
- Stubb asserts that Fedallah is ancient (older than all the hoops on the Pequod) and possibly the devil, always present with a latch‑key.
- The sperm whale’s head has “mathematical symmetry” and dignity; the right whale’s head is ungainly, like a gigantic shoe or shoemaker’s last.
- A whale’s eyes are placed so far back and low that it cannot see forward or aft; each eye gives a separate picture, leaving a “profound darkness” directly ahead and behind.
- The whale’s ear is a minuscule hole with no external flap; in the right whale the ear is covered by a membrane.
- The sperm whale’s front is a “dead, blind wall” of boneless, extremely tough tissue that acts as a battering‑ram, possibly aided by compressible internal honeycombs that connect to the outer air.
- The sperm whale’s head is divided into the upper “Case” (Heidelburgh Tun), full of pure spermaceti, and the lower “junk,” a honeycomb of oil‑filled cells.
- When spermaceti is bailed, Tashtego accidentally falls headfirst into the nearly empty case; Queequeg dives, cuts a hole in the head, hauls Tashtego out by the head, in a “running delivery.”
- The sperm whale has no face—no proper nose, no visible mouth from the front—only a vast, riddled brow; its expression is one of sublime genius and god‑like silence.
- The whale’s actual brain is a mere handful, hidden twenty feet behind the apparent forehead; its cranial cavity is tiny, but the spinal cord is immense, possibly the true seat of its power and character.
- The German ship Jungfrau (“Virgin”) has no oil and begs a lamp‑feeder from the Pequod; a race for a jaundiced, slow‑moving old bull ensues, with Derick mocking the Pequod’s boats.

## Entities And Concepts
- **Fedallah (the Parsee):** Ahab’s shadow‑like harpooneer, suspected by the crew of being the devil.
- **Stubb, Flask:** Second and third mates; their banter about the devil and Fedallah.
- **Sperm Whale head vs. Right Whale head:** Contrasted in shape, dignity, mouth structures (teeth vs. baleen), spout‑holes, tongue, lips.
- **Whale vision:** Two independent, side‑facing eyes; no binocular overlap; possible brain‑side integration problem.
- **Whale ear:** Minute, ear‑hole only; covered by membrane in right whale.
- **Battering‑ram (Chapter 76):** The sperm whale’s forehead as an un‑impaleable, elastic, boneless wall for ramming.
- **Case (Heidelburgh Tun):** Upper chamber of the sperm whale’s head containing highly prized pure spermaceti.
- **Junk:** Lower honeycomb of oil‑filled cells in the sperm whale’s head.
- **Spermaceti:** Clear, limpid oil that crystallizes when exposed to air; yields about 500 gallons from a large whale.
- **Tashtego’s accident and rescue:** Tashtego falls into the emptied case; Queequeg cuts a hole and pulls him out, described as a running obstetrical delivery.
- **Physiognomy of the Leviathan:** The sperm whale’s brow as a sublime, blank firmament; no proper nose; genius declared by pyramidical silence.
- **Phrenology and spinal theory:** The whale’s real brain is tiny; some whalemen deny it exists; the spinal cord is huge; backbone might reflect character; hump corresponds to the organ of firmness.
- **Jungfrau:** German whaler, master Derick De Deer, out of oil, “clean” (empty) ship.
- **The old bull whale:** Huge, humped, yellowish incrustations; missing starboard fin; moves slowly, labored spout, possible illness.
- **Boat race:** Pequod’s three boats vs. Derick’s boat for the old whale; German taunts; crab (oar mis‑stroke) slows German boat.

## Procedures And API Details
- **Securing a right whale body:** Lips and tongue are removed separately with the attached “crown‑piece” black bone, unlike sperm whale where the head is cut off whole.
- **Extracting jaw teeth (sperm whale):** Jaw unhinged, hoisted; Queequeg, Daggoo, Tashtego lance gums; jaw lashed to ringbolts; tackle rigged from aloft drags teeth out. Generally forty‑two teeth. Jaw sawn into slabs.
- **Tapping the Heidelburgh Tun (baling spermaceti):** A light tackle (whip) is secured to the yard‑arm. Tashtego uses a spade to break into the case, then guides an iron‑bound bucket with a long pole; deck hands hoist the full bucket, empty into a tub. Repeated until empty.
- **Rescue of Tashtego:** Queequeg dives, scuttles a hole in the descending head with his sword, reaches in upwards, and hauls Tashtego out by the head, deliberately turning him so he comes out headfirst.
- **Lamp‑feeder begging:** Derick comes in his boat holding a lamp‑feeder (and later oil‑can) to ask for oil; the Pequod supplies him.

## Nuance Or Contradictions
- The chunk ends mid‑scene in Chapter 81 during the boat chase, with the four boats “diagonically in the whale’s immediate wake.” The sentence is complete, and no explicit truncation marker appears (e.g., “[truncated at …]”). The surrounding text continues the hunt in the following chapter.
- The narrator’s anatomical‑philosophical excursions are speculative and heavily laced with satire (compare the “Kant” and “Locke” heads analogy to the ship’s list).
- The description of the whale’s brain being a mere handful is qualified; some whalemen deny there is any brain beyond the sperm magazine. The spinal‑cord theory is presented as a “hint.”
- The text moves freely between scientific observation and literary‑philosophical riff (e.g., the brow as a “great golden seal” of the German emperors, or the nose as a “pestilent conceit”).

## Candidate Wiki Hints
- **Sperm Whale anatomy** – Head structure (case, junk, battering‑ram, eyes, ears, brain/spine).
- **Right Whale anatomy** – Baleen, bonnet/crown, lower lip, tongue, and spout‑holes.
- **Whale vision and perception** – Two independent visual fields, implications for behaviour.
- **Spermaceti extraction** – The baling process and the “Heidelburgh Tun” metaphor.
- **Cistern and Buckets incident (Tashtego’s fall and rescue)** – Detailed sequence with Queequeg’s “obstetric” dive.
- **Jungfrau encounter** – Oil‑begging and the race for the sick bull; German whaler in the Pacific.
- **Physiognomy and phrenology of the Sperm Whale** – The brow, the absence of a face, and the spinal‑phrenology theory.

## chunk-15

---
title: Chunk 15 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context
Source: `Moby-Dick` > Retrieved Text, lines 13763-14749.
Chunk 15 of 17 covers chapters 82 (The Honor and Glory of Whaling) through the beginning of chapter 87 (The Grand Armada). The Pequod continues its voyage, with narrative digressions on whaling’s heritage, techniques, natural history, and the approach to the Sunda straits.

## Local Summary
Chapter 82 constructs a mythic genealogy of whalemen: Perseus, St. George, Hercules, Jonah, and Vishnoo. Chapter 83 defends the biblical Jonah against skeptics. Chapter 84 details the technique of pitchpoling—a long-distance lance throw. Chapter 85 examines the enigmatic whale spout (vapor or water) and respiratory physiology. Chapter 86 rhapsodises the tail’s anatomy and five motions. Chapter 87 introduces a massive sperm‑whale herd sighted near the Straits of Sunda.

## Key Claims
- Whaling has ancient, divine lineage extending from Greek myth to Hindu scripture.
- The story of Jonah can be rationalized (dead whale, life‑preserver, spacious mouth) to refute skeptics.
- Pitchpoling is the ultimate whaling feat: throwing a long, light pine lance in an arc from a bucking boat to strike a fast‑swimming whale.
- Despite centuries of observation, the composition of the sperm‑whale spout remains unknown; it is acrid/poisonous and can blister skin or blind.
- Whales breathe only through the spiracle; they store oxygenated blood in a “labyrinth” of vessels, allowing hour‑long dives.
- The tail’s tri‑layered structure (long horizontal fibres sandwiching short cross‑fibres) yields tremendous power; its five distinct motions are progression, mace‑like blow, sweeping, lobtailing, and peaking flukes.
- Whaling grounds are now characterised by vast aggregations of sperm whales, interspersed with empty expanses.

## Entities And Concepts
- **Perseus & Andromeda**: first whaleman, rescuing the princess by harpooning the sea monster.
- **St. George**: patron saint of England reinterpreted as a whaleman; dragon = whale.
- **Hercules**: swallowed by a whale in Greek myth, an “involuntary whaleman.”
- **Jonah**: Hebrew prophet, whose whale may have been a Right Whale (large mouth, toothless) or a dead whale acting as a life‑raft.
- **Vishnoo**: Hindu deity who incarnated as a whale to recover the Vedas.
- **Pitchpoling**: lance ~10–12 ft, pine staff, small warp line; balanced upright then arced into the whale.
- **Spout physiology**: spiracle, no connection between mouth and windpipe; labyrinth vessels for oxygen storage; non‑valvular blood‑vessels causing rapid bleed‑out.
- **Tail anatomy**: upper/middle/lower layers, crosswise middle fibres; surface area ≥50 sq. ft., width over 20 ft.
- **Five tail motions**: fin for progression, mace in battle, surface sweeping, lobtailing (playful thunderous smiting), peaking flukes (pre‑dive vertical display).
- **Sunda Straits**: between Sumatra and Java, entry to Java Sea, past piracy and aromatic cinnamon shores; Pequod’s route toward the Pacific.
- **Grand Armada**: immense herd of sperm whales, described as if nations had formed a mutual‑defense pact.

## Procedures And API Details
- **Pitchpoling method**: after the whale is fast, the harpooneer stands in the bow, levels the lance, depresses the butt to raise the point ~15 feet, then throws in a high arch to the whale’s life‑spot; the warp ensures retrieval.
- **Boat greasing**: Queequeg anoints the boat’s bottom with oil to improve slip, performing the act as a presentiment before a successful chase.

## Nuance Or Contradictions
- The chunk ends with “thousands on thousands.” — a complete sentence; no truncation marker is present.
- The defence of Jonah’s story is playfully paradoxical: it mocks skeptics while presenting their own arguments and then dismissing them as impious, yet the narrator labels Sag‑Harbor’s geographic objection “foolish pride.”
- The spout remains an “undecidable” matter; the narrator admits that even close observation yields no certainty, then proposes a semi‑comic hypothesis that all profound beings emit a mist.
- The tail is said to have no prehensile power, yet the sweeping motion is so delicate it could detect a sailor’s whisker.
- The whale’s seeming “mystic gestures” lead some hunters to compare them to Free‑Mason signs; the narrator says the whale’s face is unknowable.

## Candidate Wiki Hints
- Pitchpoling (whaling technique)
- Sperm‑whale spout (composition and toxicity)
- Whale respiratory and diving adaptations
- Cetacean tail anatomy and force
- Mythological and historical “whalemen” (Perseus, St. George, Jonah, Vishnoo)
- Grand Armada (sperm‑whale aggregation phenomenon)

## chunk-16

---
title: Chunk 16 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

Chunk Context:
This chunk corresponds to Chapters 87–92 of *Moby-Dick*, covering the Pequod’s encounter with a vast sperm‑whale aggregation, the chaos of a “gallied” herd, an interlude with the French whaler *Bouton de Rose*, and essays on whaling law, whale schools, royal perquisites and ambergris. The raw source spans lines 14751‑15806 under the heading “Moby-Dick > Retrieved Text”. The chunk ends un‑truncated at the close of Chapter 92; the narrative is complete for these chapters.

Local Summary:
The Pequod drives a crescent‑shaped host of sperm whales through the Straits of Sunda, is chased by Malay pirates, and then lowers boats into the “grand armada.” The boats become trapped in the inner calm, witnessing nursing whales, panic, and the destructive frenzy of a wounded whale. Stubb later tricks the captain of the aromatic French whaler *Rose‑Bud* into abandoning two rotting whales, from one of which he extracts ambergris. Interleaved chapters explain the sociological structure of whale schools, the unwritten whaling laws of Fast‑Fish and Loose‑Fish, the English crown’s claim to whale heads and tails, and the nature of ambergris.

Key Claims:
- Sperm whales can gather in herds covering several square miles, moving in crescent formations.
- When panicked (“gallied”) the herd can break into chaotic, paralyzed, or self‑destructive motion, dangerous to whalers.
- The *drugg*—a cross‑block fastened to a harpoon line—slows and marks a whale so it can be killed later.
- Whale schools are either harems (females with one attending male, the “schoolmaster”) or bachelor bands of young bulls (forty‑barrel‑bulls).
- The harem‑master fights rival males fiercely, but in old age becomes solitary and “admonitory.”
- The universal whaling law: I. A Fast‑Fish belongs to the party fast to it. II. A Loose‑Fish is fair game for anyone.
- English law reserves the whale’s head for the King and tail for the Queen (supposedly for whalebone), a custom still enforced in the narrator’s time.
- Ambergris is a soft, fragrant, waxy substance from the intestines of sick sperm whales, used in perfumery and cooking; it is distinct from amber.
- Stubb’s recovery of ambergris from a “dried” whale abandoned by the Rose‑Bud exemplifies Loose‑Fish doctrine.

Entities And Concepts:
- **Gallied**: a state of panic or confusion in whales when pursued.
- **Drugg**: a wooden drag (two squared boards crossed at right angles, attached by line) used to slow and mark whales in a crowd.
- **Waif**: a pennoned pole stuck into a dead whale to mark possession and location.
- **Fast‑Fish**: a whale connected to an occupied vessel or bearing a recognized mark of possession; belongs to that party.
- **Loose‑Fish**: an unmarked, unattached whale that can be claimed by the first catcher.
- **School**: a band of 20‑50 sperm whales; two types: harem (females + one male) and bachelor schools.
- **Schoolmaster**: the dominant male with a harem; in old age a solitary “hermit.”
- **Forty‑barrel‑bulls**: young, pugnacious male whales.
- **Royal fish**: whale and sturgeon, claimed by the English crown under *Bracton*.
- **Ambergris**: grey‑amber, a scented morbific secretion of the sperm whale, unrelated to fossilized amber.
- **Bouton de Rose** (Rose‑Bud): a French whaler with a green, rose‑like figurehead, captained by an inexperienced former cologne‑maker.

Procedures And API Details:
- **Drugging**: fasten a line between a harpoon and the drugg block; dart at a whale to create drag and mark it for later lancing.
- **Waifing**: after a kill, thrust a waif pole into the floating carcass to secure possession and location.
- **Cutting‑spade hamstringing**: dart a short‑handled spade attached to a retrieval rope at a whale’s tail tendon to disable it.
- **Ambergris extraction**: cut behind the side fin of a “dry” whale, reach into the cavity and pull out the wax‑like lumps.

Nuance Or Contradictions:
- The Fast‑Fish/Loose‑Fish passages blend genuine whaling custom with satire on property, colonialism, and human rights; they are not formal legislation.
- Lord Ellenborough’s ruling analogises a lost whale to an abandoned wife, a comically strained legal argument that the narrator presents as foundational.
- The English royal‑fish law rests on the mistaken‑but‑comedic assumption that whalebone comes from the tail, yet is still enforced with “a reason in all things, even in law.”
- Ambergris’s exact origin was debated (cause or effect of whale “dyspepsia”); the novel offers no final resolution.
- The chunk ends naturally at Chapter 92 with the remark on Cologne‑water’s foul beginnings; no truncation marker is present.

Candidate Wiki Hints:
- Whaling tools and techniques: drugg, waif, cutting‑spade, gallied.
- Fast‑Fish / Loose‑Fish as a philosophic‑legal concept.
- Social structure of sperm whales in *Moby‑Dick* (schools, schoolmaster, bachelor bulls).
- Ambergris: historical trade, properties, literary role.
- English royal prerogatives over whale catches (head and tail).
- The *Bouton de Rose* episode as a study of greed, trickery, and salvage.

## chunk-17

---
title: Chunk 17 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/102-moby-dick.md
confidence: medium
---

## Chunk Context

- Source lines: 15808–15995
- Heading path: Moby-Dick > Retrieved Text
- The chunk covers the closing appeal of Chapter 92 (The Whale as a Dish) and the beginning of Chapter 93 (The Castaway), then truncates partway through the Pip episode.

## Local Summary

The narrator rebuts the charge that all whales smell bad. He traces the origin of the stigma to the practices of Greenland whalers, who had to bring blubber home un-tried, and to the on-shore try‑works at the Dutch village Smeerenberg. He argues that sperm whales, when properly handled, are nearly scentless and even fragrant. Then the narrative shifts to the story of Pip, the Pequod’s young black ship-keeper, who is temporarily placed as an oarsman. During a second lowering Pip jumps again from the boat and is left behind in the open sea. The chunk ends mid‑sentence with Pip not yet rescued.

## Key Claims

- The belief that all whales smell bad comes from the un-tried blubber brought home by early Greenland whalers and from the on‑shore boiling operations at Smeerenberg.
- Sperm‑whale oil, after proper boiling and casking, is “nearly scentless.”
- Well‑treated living or dead whales are not creatures of ill odor; the motion of a sperm whale’s flukes even gives off a musk‑like perfume.
- Ship‑keepers are often those deemed too slight or timid for the boats; Pip is made a ship‑keeper.
- Pip is intelligent and genial but terrified of the whale‑hunt. His first accidental overboard incident ends with Stubb cutting the line; Stubb warns him sternly not to jump again.
- On the second lowering Pip jumps again, is abandoned by Stubb, and is left far astern. The narrative says Stubb expected other boats to pick him up, but they chase nearby whales instead.
- The chunk explicitly ends with a truncation marker: `[truncated at 900000 characters]`.

## Entities And Concepts

- **Greenland whaling ships**: Brought blubber home in casks without trying out the oil at sea, leading to foul smell.
- **Smeerenberg (Schmerenburgh)**: Dutch on‑shore try‑works village on the Greenland coast; name means “fat‑mountain”.
- **Fogo Von Slack**: (Fictitious) author of a great work on smells, cited for the name Smeerenberg.
- **Sperm Whale fragrance**: Compared to a musk‑scented lady’s dress and to a myrrh‑bearing elephant led out for Alexander the Great.
- **Pip (Pippin)**: Young black ship‑keeper from Tolland County, Connecticut; originally a lively tambourine player; forced into a boat after the after‑oarsman’s injury.
- **Stubb**: Second mate; cuts the line to save Pip the first time, but strictly warns him; true to his word when Pip jumps again.
- **Tashtego**: Harpooneer; calls “Cut?” with knife ready when Pip is tangled.
- **Pequod**: The whaling ship.

## Procedures And API Details

- **Greenland whalers’ method**: Cut fresh blubber into small bits, thrust it through bung‑holes of large casks, and carry it home without trying out at sea because of short seasons and violent storms.
- **Southern whalers’ method**: Try out oil at sea; a sperm‑whaling voyage of four years may consume only about fifty days in boiling; the casked oil is nearly scentless.
- **Smeerenberg operations**: A collection of furnaces, fat‑kettles, and oil sheds for trying out Dutch fleet blubber on‑shore.

## Nuance Or Contradictions

- The narrator first claims whales “as a species” are not ill‑smelling, then focuses only on sperm whales; the Greenland/right‑whale blubber example suggests some whales **do** stink if not tried out promptly.
- Stubb’s advice is deliberately contradictory: _Stick to the boat_ vs. _Leap from the boat_ depending on the situation, followed by an absolute command not to jump.
- The chunk ends mid‑episode with a hard truncation: `[truncated at 900000 characters]`. The raw source stops in the middle of a sentence as the ship is about to rescue Pip. The full outcome is not present in this chunk.

## Candidate Wiki Hints

- **Pip (Moby-Dick)**: The chunk provides detailed backstory, personality, and the pivotal “castaway” event; a character page could track his transformation.
- **Whale‑oil rendering at sea vs. on shore**: Contrast between Greenland and South Sea methods; could be extracted as a whaling‑history note.
- **Smeerenberg (Schmerenburgh)**: A historical/tale note on the Dutch try‑works village, its name and role.

