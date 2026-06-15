## chunk-01

---
title: Chunk 01 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---

Chunk Context
- **Chunk ID:** 1 of 10
- **Line Range:** 1-26
- **Heading Path:** Document > Alice's Adventures in Wonderland > Fetch Metadata
- **Source File:** raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md

Local Summary
This chunk establishes the metadata and acquisition context for the source document "Alice's Adventures in Wonderland" (Corpus item 104). It identifies the content as a Project Gutenberg plain-text edition retrieved from `https://www.gutenberg.org/files/11/11-0.txt`. The record notes a specific technical failure during initial retrieval where the ebook landing page failed TLS verification via `urllib`, necessitating a supplemental fetch of the direct file URL.

Key Claims
- The document is categorized under Project Gutenberg with tags including `project-gutenberg` and `web-corpus`.
- The content type is identified as `text/plain; charset=utf-8`.
- The retrieval status for the final file URL is reported as "ok via supplemental web fetch".
- The specific edition retrieved is described as a "Shorter narrative source."

Entities And Concepts
- **Alice's Adventures in Wonderland** (Title)
- **Project Gutenberg** (Source Category)
- **urllib** (Python Library referenced in error context)
- **TLS** (Transport Layer Security protocol involved in the fetch failure)
- **Corpus item 104** (Internal identifier)

Procedures And API Details
- **Initial Fetch Failure:** The landing page at `https://www.gutenberg.org/ebooks/11` failed TLS verification when accessed via `urllib`.
- **Resolution Procedure:** A supplemental acquisition was performed directly from the file URL: `https://www.gutenberg.org/files/11/11-0.txt`.

Nuance Or Contradictions
- The metadata describes the retrieved content as a "Shorter narrative source," which contrasts with the typical full-length nature of Project Gutenberg editions, though this may refer to a specific truncation or versioning within the corpus.

Candidate Wiki Hints
- **Acquisition Troubleshooting:** A note on handling TLS failures when fetching raw text files from legacy web archives using Python's `urllib`.
- **Project Gutenberg File Structure:** The distinction between the landing page (`/ebooks/<id>`) and the actual text file location (`/files/<id>/<id>-<version>.txt`).

## chunk-02

---
title: Chunk 2 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---
Chunk Context
- Source: `raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md`
- Chunk range: Lines 27–371 (Chapter I "Down the Rabbit-Hole" through the start of Chapter II "The Pool of Tears")
- Heading path: `Alice's Adventures in Wonderland > Retrieved Text`

Local Summary
- Alice falls down a rabbit-hole into a strange underground world. She discovers locked doors, a golden key, and a bottle labeled "DRINK ME" that shrinks her to ten inches. After failing to reach the key, she eats a cake marked "EAT ME" but remains the same size. Later, after drinking from another bottle (implied by her growth), she becomes over nine feet tall. She cries until a pool forms, then encounters the White Rabbit again and questions her identity while reciting incorrect arithmetic and geography.

Key Claims
- Alice enters Wonderland through a rabbit-hole that dips suddenly down like a tunnel.
- The well contains cupboards, book-shelves, maps, and pictures hung on pegs.
- A tiny golden key found on a glass table fits a small door behind a low curtain.
- Drinking from a bottle marked "DRINK ME" causes Alice to shrink to ten inches high.
- Eating a cake marked "EAT ME" does not change her size immediately, contrary to expectation.
- Alice's identity becomes uncertain; she questions whether she is still herself or someone else (e.g., Mabel).
- Arithmetic and geography facts are inverted in Wonderland (e.g., London is the capital of Paris).

Entities And Concepts
- **Alice**: The protagonist who falls into Wonderland and experiences size changes.
- **White Rabbit**: A character with pink eyes, a waistcoat-pocket, and a watch; serves as a guide figure.
- **Golden Key**: A tiny key that opens a small door leading to a garden passage.
- **Bottle ("DRINK ME")**: Causes Alice to shrink.
- **Cake ("EAT ME")**: Intended to change size but initially has no effect.
- **Dinah**: Alice's cat, mentioned as missing her.
- **Antipathies**: A term Alice considers when thinking about falling through the earth (later corrected to "Antipodes").
- **Multiplication Table**: Referenced with incorrect results (e.g., four times five is twelve).

Procedures And API Details
- Shrinking procedure: Drink from a bottle labeled "DRINK ME" without checking for poison.
- Growth procedure: Implied by later context (likely another drink or action), causing Alice to grow to over nine feet tall.
- Identity verification attempt: Recite known facts (math, geography, poetry) to confirm self; results are incorrect in Wonderland.

Nuance Or Contradictions
- The cake marked "EAT ME" is expected to change size but initially does not, leading Alice to question reality.
- Arithmetic and geographic knowledge are reversed (e.g., four times five equals twelve; London is the capital of Paris).
- Alice's identity fluctuates between being herself and possibly Mabel or another person.
- The bottle labeled "DRINK ME" appears after Alice already drank from it earlier in the chapter, creating a temporal inconsistency in the narrative flow within this chunk.

Candidate Wiki Hints
- **Concept**: Size-changing objects in fantasy literature (bottles, cakes).
- **Theme**: Identity crisis and self-questioning in surreal environments.
- **Location**: Underground passages with locked doors and hidden rooms as common tropes in adventure fiction.

## chunk-03

---
title: Chunk 3 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---
Chunk Context
Lines 373-721 cover the transition from Alice's escape through the little door to her shrinking, falling into a pool of tears, and interacting with various creatures before the White Rabbit appears. This segment spans the end of Chapter II ("The Pool of Tears") and the beginning of Chapters III ("A Caucus-Race and a Long Tale") and IV ("The Rabbit Sends in a Little Bill").

Local Summary
Alice shrinks after finding herself unable to pass through the little door with her key. She falls into a pool of tears she wept while nine feet tall, where she meets a tiny mouse. After failing to communicate due to the mouse's hatred for cats and dogs, she joins a gathering of wet animals on the bank led by the Mouse. The group attempts to dry off via a Caucus-race organized by the Dodo, during which Alice shares comfits as prizes. Following the race, the Mouse refuses to continue his story about hating "C and D" (cats and dogs) after Alice mentions her cat, Dinah, causing most birds to flee. Alone again, Alice encounters the White Rabbit, who mistakes her for his housemaid, Mary Ann. Alice enters the Rabbit's home to retrieve gloves and a fan but accidentally drinks from an unlabeled bottle hoping to grow large again.

Key Claims
- Alice shrinks significantly after passing through the little door, making it impossible for her to reach the key on the table.
- The pool Alice falls into is actually the pool of tears she wept earlier when she was nine feet high.
- Animals such as a Duck, Dodo, Lory, and Eaglet have fallen into the same pool of tears.
- A Caucus-race is defined by its lack of rules: participants run whenever they like, and everyone is declared a winner regardless of performance.
- The Mouse claims to know the "driest thing" (a history lesson) but his attempt to dry the group fails.
- Alice's cat, Dinah, is described as excellent at catching mice but also capable of eating birds.
- Mentioning Dinah causes the bird creatures to leave immediately due to fear.
- The White Rabbit addresses Alice by her former name, Mary Ann, assuming she is his housemaid.

Entities And Concepts
- Alice: The protagonist who has recently shrunk.
- Mouse: A small creature in the pool who hates cats and dogs; claims authority among the wet animals.
- Dodo: An animal that organizes a Caucus-race to solve the drying problem.
- Lory: A colorful bird involved in the race and argument about age.
- Duck: Questions the meaning of pronouns during the Mouse's story.
- Eaglet: Points out that the other characters do not understand long words.
- Dinah: Alice's pet cat, known for catching mice and eating birds.
- White Rabbit: A character who believes Alice is his servant, Mary Ann.
- Caucus-race: A nonsensical race where everyone runs arbitrarily and everyone wins.
- Comfits: Candy-like treats Alice gives as prizes to the animals.
- Thimble: The only prize left after the comfits are distributed.

Procedures And API Details
- Caucus-race procedure:
  1. Mark out a race course in a circle (shape does not matter).
  2. Place all participants along the course randomly.
  3. Start running when ready and stop when ready.
  4. Declare the race over after approximately half an hour.
  5. Announce that everyone has won.
  6. Distribute prizes (e.g., comfits, thimble).
- Alice's entry into White Rabbit's house:
  1. Follow directions given by the Rabbit to go home and fetch items.
  2. Enter a house labeled "W. RABBIT" without knocking.
  3. Retrieve fan and gloves from the table in the window.
  4. Uncork and drink from a bottle near the looking-glass hoping to restore size.

Nuance Or Contradictions
- The Mouse claims to be drying the animals with his story, yet Alice remains wet after listening for a while.
- The Lory refuses to state its age, leading to an impasse when Alice questions its authority based on age.
- Alice interprets the Mouse's silence as potential language barriers (English vs. French), showing her limited historical context regarding William the Conqueror.
- The Mouse is offended by any mention of cats or dogs, creating a conflict with Alice who wishes to be friendly.
- The animals interpret the Mouse's story about William the Conqueror as a method of drying, which is logically inconsistent.

Candidate Wiki Hints
- Page: "Caucus-race" (Explaining the specific type of nonsensical race in Wonderland).
- Page: "Pool of Tears" (Describing the phenomenon where Alice shrinks into her own tears).
- Page: "White Rabbit's House" (Details on the rabbit home and the unlabeled bottle).

## chunk-04

---
title: Chunk 4 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---
Chunk Context
This chunk covers Alice's second transformation (shrinking via a magic cake) and her subsequent escape from the house into a forest. It includes an encounter with an enormous puppy and the beginning of Chapter V, featuring a blue caterpillar smoking a hookah on top of a mushroom. The dialogue between Alice and the caterpillar initiates the logic of growing/shrinking using different sides of the mushroom.

Local Summary
After realizing she cannot grow larger inside the house, Alice finds that pebbles thrown by the Rabbit turn into cakes. She eats one and shrinks significantly. Escaping through the door, she finds a crowd of animals outside but flees into a wood to regain her size. There, she meets a large blue caterpillar sitting on a mushroom. The caterpillar challenges Alice's identity and confusion regarding her changing size. It reveals that the mushroom has two sides: one makes you grow taller, the other shorter.

Key Claims
- Consuming specific food items (cakes) or substances (mushroom sides) alters physical size.
- Identity and memory are fluid when physical size changes rapidly ("I can't remember things as I used").
- The caterpillar possesses knowledge of the mechanics of transformation that Alice does not yet understand.
- Specific verses from "You Are Old, Father William" must be recited correctly to satisfy the caterpillar's logic tests.

Entities And Concepts
- Alice: Protagonist experiencing rapid size changes.
- Rabbit: Owner of the house and the one throwing pebbles (which become cakes).
- Bill: A little lizard caught in the window frame, requiring rescue via a ladder/chimney.
- Dinah: Alice's cat, mentioned as a threat to the intruders.
- Mushroom: The object with two distinct sides affecting size; currently inhabited by a caterpillar.
- Blue Caterpillar: An entity sitting atop the mushroom who knows the rules of transformation.
- Hookah: A smoking device used by the caterpillar.

Procedures And API Details
- Shrinking Procedure: Eat a cake that has turned from a pebble. Result: Immediate shrinking until small enough to pass through the door.
- Growth/Shrinkage Mechanism: Use a mushroom with two sides. One side causes growth; the other causes shrinking. (Specific direction not yet determined in this chunk).
- Identity Verification Test: Recite the poem "You Are Old, Father William" accurately to demonstrate mental stability and knowledge.

Nuance Or Contradictions
- The caterpillar claims Alice's recitation was wrong "from beginning to end," though Alice admits only some words were altered, suggesting a high standard for correctness or a test of exact memory under stress.
- The caterpillar dismisses the confusion of changing sizes ("It isn't"), contrasting with Alice's experience of it being very confusing.
- The caterpillar asserts that turning into a chrysalis and butterfly is not "queer," contradicting Alice's expectation that such transformations are strange experiences.

Candidate Wiki Hints
- [Mushroom] - A magical object in Wonderland with dual properties for size manipulation.
- [Caterpillar] - A cryptic character who guards the secret of growth/shrinking.
- [You Are Old, Father William] - A poem used as a test of identity and memory in Wonderland.

## chunk-05

---
title: Chunk 5 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---
Chunk Context
- Source segment covering Alice's size manipulation via the mushroom, her descent to a garden, and the arrival at the Duchess's house.
- Scene transition from the "Curiouser and Curiouser" growth/shrinking sequence to the "Pig and Pepper" chapter.
- Key interactions include the Footmen's confusion, the Duchess's chaotic nursery, and the introduction of the Cheshire Cat.

Local Summary
Alice finishes shrinking to her normal size using a mushroom she found earlier, allowing her to enter a garden. She arrives at a small house where two footmen in livery (described as fish-like) deliver an invitation from the Queen to play croquet for the Duchess. Inside, Alice finds the Duchess nursing a baby while a cook stirs a cauldron of soup containing excessive pepper. The atmosphere is chaotic, marked by constant sneezing and crashing dishes. The Duchess sings a lullaby advising harsh treatment of her child due to his tendency to sneeze. Alice nurses the baby until it turns into a pig, which she releases into the woods. Shortly after, she encounters the Cheshire Cat perched on a tree branch.

Key Claims
- Size alteration: Consuming specific parts of the mushroom allows Alice to shrink or grow at will.
- Garden entry: Shrinking is a prerequisite for entering the garden without causing alarm.
- Footmen description: The footmen are identified by their livery and powdered, curling hair; one is likened to a fish due to facial features.
- Chaos in the Duchess's kitchen: The environment is characterized by excessive pepper, constant sneezing from the Duchess and baby, and violent outbursts involving cookware.
- Cheshire Cat behavior: The cat grins constantly and engages Alice in dialogue regarding the nature of madness.

Entities And Concepts
- Alice: Protagonist undergoing physical transformations.
- Mushroom: Object used for size manipulation (right-hand side vs. left-hand side).
- Footmen: Servants in livery; one called "Fish-Footman," the other "Frog-Footman."
- Duchess: Mother of the baby, sings a song about beating children who sneeze.
- Baby: A child who sneezes constantly and transforms into a pig.
- Cheshire Cat: A feline character known for its grin and philosophical comments on madness.
- Croquet: The game the Queen invites the Duchess to play.

Procedures And API Details
- Size Adjustment Procedure: Nibble the right-hand bit of the mushroom to shrink; nibble the left-hand bit to grow.
- Nursery Song Structure: A lullaby sung by the Duchess with a chorus response ("Wow! wow! wow!") involving themes of physical punishment for sneezing.
- Baby Transformation: Holding the baby in a specific posture (twisted into a knot, holding right ear and left foot) prevents it from undoing itself during growth/shrinking processes.

Nuance Or Contradictions
- Footman Identity: Alice identifies them as footmen based on their clothing but notes their faces resemble fish, creating an ambiguity between professional title and physical description.
- Baby Nature: The baby is initially human-like but transforms into a pig after being nursed by Alice, challenging the stability of identity in this world.
- Madness Definition: The Cheshire Cat asserts that everyone in Wonderland is mad, yet defines madness pragmatically (e.g., dogs are not mad), leading to circular logic about Alice's sanity for visiting such a place.

Candidate Wiki Hints
- Page: Size Manipulation Methods in Wonderland
- Page: The Cheshire Cat and the Philosophy of Madness
- Page: Character Profiles: The Duchess

## chunk-06

---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---
### Chunk Context
This chunk spans the conclusion of Chapter VI ("The Mad Tea-Party") and the beginning of Chapter VIII ("The Queen's Croquet-Ground"). It features the famous riddle "Why is a raven like a writing-desk?", discussions on time mechanics at the tea party, the Dormouse's treacle-based story, Alice's departure from the party, her re-entry into the house via a tree door, and the arrival of gardeners painting white roses red to avoid the Queen's wrath.

### Local Summary
The narrative shifts from the absurdity of the Mad Tea-Party to Alice's journey back through the woods to the garden. The tea-party characters engage in circular logic regarding riddles and time, culminating in Alice leaving due to rudeness. She enters a tree with a door, finds the golden key, grows small enough to pass through a passage, and arrives at the Queen's Croquet-Ground where three gardeners are frantically painting white roses red before the Queen's arrival.

### Key Claims
- **Madness as Logic**: The Cat argues that its behavior (growling when pleased, wagging tail when angry) is evidence of madness because it contradicts dog logic.
- **Time Personification**: The Hatter claims Time is a person ("him") who went mad after quarreling with the March Hare at the Queen's concert, explaining why time is stuck at six o'clock.
- **Circular Riddles**: The tea-party characters equate "I mean what I say" with "I eat what I see," highlighting the nonsensical nature of their conversation.
- **Treacle Life**: The Dormouse claims three little sisters lived at the bottom of a well and subsisted on treacle, drawing things starting with 'M' (mouse-traps, moon, memory).
- **Rose Painting**: Three gardeners are painting white roses red because they were planted by mistake, fearing beheading if the Queen discovers the error.

### Entities And Concepts
- **Characters**: Alice, The Cat, The March Hare, The Hatter, The Dormouse, Five, Seven, Two (Gardeners).
- **Objects**: Mushroom (for growing/shrinking), Golden Key, Writing-desk, Raven, Tea-things, Treacle-well.
- **Concepts**: Madness, Time personification, Circular logic, Riddles, Croquet-Ground, Painting roses red.

### Procedures And API Details
- **Growth Mechanism**: Alice nibbles the lefthand bit of mushroom to raise herself to about two feet high; later she keeps a piece in her pocket to grow to a foot high again.
- **Party Dynamics**: The tea-party involves moving around the table as items are used up, with the March Hare and Hare resting elbows on the sleeping Dormouse.
- **Storytelling Protocol**: When Alice asks for a story, the Dormouse is pinched awake by both the Hare and Hatter to deliver a hurried tale about three sisters living in a treacle-well.

### Nuance Or Contradictions
- **Riddle Resolution**: The Hatter admits he has no idea why a raven is like a writing-desk, yet the conversation treats it as a solvable puzzle that Alice can guess.
- **Time Perception**: While Alice views time as linear (staying the same year for a long time), the Hatter treats Time as a volatile entity who refuses to be beaten or managed.
- **Garden Logic**: The gardeners paint roses red not because they are red, but because the white one was put in by mistake, creating a surreal task of altering nature to fit political (Queen's) expectations.

### Candidate Wiki Hints
- Page: `Mad Tea-Party` (covering riddles, time mechanics, and character interactions).
- Page: `Time Personification in Wonderland` (exploring the Hatter's lore about Time going mad).
- Page: `Treacle-Well` (documenting the absurd geography of the Dormouse's story).

## chunk-07

---
title: Chunk 7 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---
Chunk Context
- Source path: raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
- Heading path: Alice's Adventures in Wonderland > Retrieved Text
- Line range: 1927-2320
- Narrative segment: The trial scene at the croquet grounds, featuring the procession of cards, the Queen's erratic execution orders, the Cheshire Cat's appearance and disappearance, and the transition to the Duchess.

Local Summary
Alice observes a grand procession led by the King and Queen of Hearts. The crowd includes soldiers, courtiers, royal children, and the White Rabbit. Alice intervenes when the Queen orders the beheading of three gardeners hiding under a rose-tree, placing them in a flower-pot for safety. The scene shifts to a chaotic croquet game where hedgehogs are balls, flamingoes are mallets, and soldiers serve as arches. The Cheshire Cat appears, engaging Alice in conversation while its body gradually fades away. Later, the Duchess joins Alice, offering cynical musings on morals and temperaments. The Queen interrupts with threats of execution until the Duchess vanishes.

Key Claims
- The procession order: ten soldiers (oblong/flat), ten courtiers (diamonds), ten royal children (hearts), guests (Kings/Queens including the White Rabbit), the Knave carrying the King's crown, and finally the King and Queen of Hearts.
- Alice questions the utility of a procession if everyone lies face down like the gardeners; she stands still to watch.
- The Queen demands names and threatens execution ("Off with her head!"), but Alice retorts "Nonsense!" loudly.
- The King attempts to intervene gently, calling Alice a child, but the Queen ignores him.
- Gardeners are turned over by the Knave; soldiers execute them after confirming their heads are gone.
- The croquet game features live animals (hedgehogs as balls, flamingoes as mallets) and soldiers doubling up as arches.
- Players quarrel constantly; the Queen shouts execution orders frequently.
- Alice perceives a grin in the air, identifying it as the Cheshire Cat.
- The King demands the Cat's removal; the Queen orders its beheading without seeing it fully.
- The executioner debates whether to execute a headless body; the King insists anything with a head can be beheaded.
- Alice suggests asking the Duchess, who is imprisoned.
- The Cheshire Cat fades away as the executioner leaves.
- The Duchess claims everything has a moral and offers sharp, often contradictory wisdom (e.g., "Birds of a feather flock together," "Take care of the sense").
- The Queen threatens execution unless the Duchess or her head is removed; the Duchess vanishes instantly.

Entities And Concepts
- Alice: Protagonist, polite but courageous toward card figures.
- White Rabbit: Nervous character, warns about the Duchess's execution for boxing the Queen's ears.
- Knave of Hearts: Carries the crown, executes orders carefully with one foot.
- King and Queen of Hearts: Antagonists; the Queen is volatile, the King timid but powerless against her.
- Cheshire Cat: Mysterious entity appearing as a grin/head then fully, capable of fading away; engages Alice in dialogue.
- Duchess: Ugly, sharp-chinned figure linked to Alice; offers cynical morals and vanishes under threat.
- Executioner: Debates logic of beheading headless bodies; acts on orders.
- Gardeners (Two): Lie under rose-tree, accused of unknown activity; protected by Alice.
- Soldiers/Courtiers/Royal Children: Card figures forming the procession.
- Croquet Game Elements: Live hedgehogs (balls), live flamingoes (mallets), soldiers as arches.

Procedures And API Details
- Execution Procedure: Queen shouts "Off with [his/her] head!" -> Knave/Executioner acts -> Soldiers confirm heads gone.
- Cat Disappearance: King calls for removal -> Executioner leaves -> Cat fades entirely upon his return.
- Duchess Vanishing: Queen threatens execution -> Duchess takes choice to leave -> Disappears instantly.

Nuance Or Contradictions
- Identity Ambiguity: Gardeners appear identical to soldiers/courtiers when lying face down, making classification difficult for Alice.
- Logic vs. Authority: Executioner argues logically (no body = no head removal), but the King and Queen override reason with arbitrary decrees.
- Morality vs. Reality: Duchess insists everything has a moral, yet her own actions (vanishing under threat) contradict stable morality.
- Game Chaos: Croquet is played without rules; players quarrel constantly; arches move independently of turns.
- Appearance vs. Existence: The Cheshire Cat exists partially (grin first, then head/body), defying normal physical continuity.

Candidate Wiki Hints
- "The Queen of Hearts as an Antagonist": Explores her volatile nature and execution decrees.
- "The Cheshire Cat's Metaphysical Presence": Discusses its ability to appear/disappear independently of physical laws.
- "Morals in Nonsense": Analyzes the Duchess's habit of extracting morals from random statements.
- "Card Soldiers and Gardeners": Examines identity confusion among card figures based on posture and pattern.

## chunk-08

---
title: Chunk 8 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---

### Chunk Context
This chunk covers the transition from the Queen of Hearts' chaotic executions to the meeting with the Mock Turtle and the Gryphon. It details the famous "Mock Turtle" lesson on sea education, the description of the Lobster Quadrille dance, and concludes with Alice reciting poems ("The Lobster," "The Owl and the Panther") that are presented in a fragmented state between original text and later edition additions within this specific document file.

### Local Summary
After the Queen orders more executions, she leaves Alice with the Gryphon to visit the Mock Turtle. The Mock Turtle recounts his education under a master named Tortoise (who taught "Reeling and Writhing" and branches of Arithmetic like "Ambition" and "Derision"). He then describes the Lobster Quadrille, a dance involving lobsters as partners that ends with dancers being thrown out to sea. Alice recites verses about a whiting getting its tail stuck in its mouth and boots made of soles and eels. The section ends with Alice attempting to repeat poems about a lobster trimming his nose and an owl/panther sharing a pie, which the Gryphon interrupts as confusing nonsense.

### Key Claims
- **Education under the Sea:** The curriculum included Reeling and Writhing, Ambition, Distraction, Uglification, Derision, Mystery, Ancient and Modern, Seaography, Drawling, Stretching, Fainting in Coils, Laughing, and Grief.
- **Lesson Duration:** Lessons lasted ten hours on the first day and decreased daily ("they lessen from day to day").
- **The Lobster Quadrille:** A dance performed in two lines along the sea-shore involving partners (lobsters), advancing, changing partners, throwing lobsters out to sea, swimming after them, somersaulting, and returning to land.
- **Whiting Anatomy/History:** Whitings have their tails fast in their mouths because they were thrown out to sea by lobsters; they are made of soles and eels.
- **Poetic Content:** The text provided includes specific verses for "The Lobster" (trimming nose/belt) and "The Owl and the Panther" (sharing a pie), noting that these specific lines appear in later editions.

### Entities And Concepts
- **Mock Turtle:** A creature who claims to have been a real turtle; serves as an educator.
- **Gryphon:** A winged lion-eagle creature acting as Alice's guide and chaperone.
- **Tortoise:** The master of the sea school (distinct from the Mock Turtle).
- **Lobster Quadrille:** A specific dance choreography involving lobsters.
- **Whiting:** A fish characterized by having its tail stuck in its mouth due to being thrown into deep water.
- **Seaography:** Listed as a subject taught at the sea school.
- **Drawling:** An art form taught by an old conger-eel, involving drawling, stretching, and fainting in coils.

### Procedures And API Details
None. The text describes fictional procedures (dancing, schooling) rather than technical APIs or real-world commands.

### Nuance Or Contradictions
- **Textual Integrity:** The chunk explicitly contains markers for "[later editions continued as follows]," indicating that the source document merges original Lewis Carroll text with additions from later printings of *Alice's Adventures in Wonderland*. This creates a hybrid narrative where Alice recites lines (e.g., "I heard him declare," "The Panther took pie-crust") that were not present in the original 1865 edition.
- **Character Logic:** The Gryphon and Mock Turtle frequently contradict each other or Alice regarding facts (e.g., whether crumbs would wash off, the definition of uglifying), reflecting the nonsensical nature of Wonderland logic.
- **Lesson Reduction:** The concept that lessons "lessen" from day to day is a pun on the word "lesson," which is treated as a factual rule in this context.

### Candidate Wiki Hints
- **Page: Mock Turtle** (Concept: Fictional educator and character analysis)
- **Page: Lobster Quadrille** (Concept: Descriptive entry for the fictional dance)
- **Page: Sea School Curriculum** (Concept: List of absurd subjects like Ambition, Derision, and Seaography)

## chunk-09

---
title: Chunk 9 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---

### Chunk Context
This chunk covers the transition from the Mock Turtle's song to the beginning of Chapter XI, "Who Stole the Tarts?", and continues through the trial proceedings involving the Knave. It includes the arrival of the Hatter, March Hare, and Dormouse; Alice's growth back to normal size; the chaotic cross-examination of the Hatter; the Duchess's cook's testimony; and the sudden appearance of Alice as the next witness in Chapter XII. The text ends with the reading of a set of verses allegedly written by the prisoner (the Knave).

### Local Summary
The scene shifts from the Jabberwocky party to the courtroom where the trial for stealing tarts begins. The King presides, looking uncomfortable under his crown and wig, while twelve jurors—various animals and birds—write names on slates to avoid forgetting them. Alice observes the proceedings, noting the absurdity of the court. When the Hatter is called, he claims ignorance about dates (offering conflicting answers from the March Hare and Dormouse) and admits to having no hat because he sells stolen ones. The Queen interrupts frequently, demanding executions or suppressing applause from guinea-pigs. Alice grows large again, knocking over the jury box with her skirt. After restoring the jurors, she is called as a witness. She claims ignorance of the case until the King cites "Rule Forty-two: All persons more than a mile high to leave the court." The Knave denies writing the verses found in his possession, and the King orders them read aloud.

### Key Claims
- **The Trial**: The Queen of Hearts made tarts on a summer day, and the Knave stole them.
- **Juror Behavior**: Jurors write their names down immediately to prevent forgetting them; one juror (Bill the Lizard) cannot spell "stupid" and writes with his finger after Alice takes his pencil.
- **Rule Forty-two**: A newly invented rule states that anyone over a mile high must leave the court.
- **The Verses**: A set of verses is found in the Knave's possession, but it is not in his handwriting, nor signed. The King assumes this implies guilt ("You _must_ have meant some mischief").

### Entities And Concepts
- **Characters**: Alice (growing large), King of Hearts (presiding judge), Queen of Hearts (accuser/judge), Knave of Hearts (defendant), White Rabbit (timekeeper/clerk), Gryphon, Mock Turtle (previous scene), Hatter, March Hare, Dormouse, Duchess's cook, Bill the Lizard (juror).
- **Objects**: Slates, pencils, tarts, teacup, bread-and-butter, trumpet, parchment scroll, canvas bag (for suppressing applause).
- **Concepts**: Court of justice, jury box, verdict, cross-examination, execution, Rule Forty-two.

### Procedures And API Details
- **Suppressing Applause**: Officers place a victim in a large canvas bag tied at the mouth and sit upon it to suppress cheering (demonstrated with guinea-pigs).
- **Trial Procedure**: The King reads accusations from a scroll, calls witnesses via trumpet blasts, asks questions, and orders verdicts. Witnesses are cross-examined by the King or Queen.
- **Rule Application**: When Alice refuses to leave despite Rule Forty-two, the King cites it as an "oldest rule in the book," though Alice notes it was just invented.

### Nuance Or Contradictions
- **Time and Dates**: The Hatter gives conflicting dates for starting tea (Fourteenth, Fifteenth, Sixteenth), which the jury adds up mathematically despite the absurdity.
- **Guilt vs. Evidence**: The King equates the lack of a signature on the verses with an admission of guilt ("You _must_ have meant some mischief"), a logical fallacy highlighted by Alice.
- **Rule Validity**: The King insists Rule Forty-two is the "oldest rule in the book," yet Alice correctly identifies it as something invented just moments prior.
- **Juror Orientation**: Alice accidentally puts the Lizard juror in head-downwards; she argues it makes no difference to the trial, contrasting with the court's rigid adherence to form.

### Candidate Wiki Hints
- **Rule Forty-two**: A fictional rule from *Alice's Adventures in Wonderland* illustrating arbitrary authority and logical absurdity.
- **The Mock Turtle Song**: Features the nonsense phrase "Beautiful Soup," a recurring motif of food-related absurdity.
- **Courtroom Absurdity**: The depiction of a trial where logic is inverted, rules are made up on the spot, and evidence is interpreted arbitrarily.

## chunk-10

---
title: Chunk 10 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/104-alice-s-adventures-in-wonderland.md
confidence: medium
---

Chunk Context
This chunk covers the conclusion of *Alice's Adventures in Wonderland*. It includes the climax where Alice wakes up after being buried under playing cards, her reunion with her sister by the riverbank, and the final narrative section detailing her sister's dream about Alice. The text ends with the story's closing remarks and the standard Project Gutenberg footer.

Local Summary
The Queen orders Alice's execution, but Alice retorts that they are merely a pack of cards. The cards fall on her, and she wakes up beside her sister. Alice recounts her adventures to her sister, who then drifts off into a dream about Alice. The dream features various characters from the story (White Rabbit, March Hare, Mock Turtle) and transitions back to reality, ending with the sister imagining Alice as an adult woman sharing these tales with future children.

Key Claims
- The entire sequence of events described in the book is revealed to be a dream experienced by Alice while resting by the riverbank.
- The characters (White Rabbit, March Hare, Queen, etc.) are products of Alice's imagination during her nap.
- Alice's sister projects herself into Alice's dream world and then imagines Alice's future life as an adult storyteller.

Entities And Concepts
- **Alice**: The protagonist who falls asleep and dreams the adventure; later envisioned as a grown woman.
- **Sister**: Alice's caretaker who finds her after the card attack; the actual narrator of the dream sequence.
- **The Queen**: Antagonist whose order to execute Alice triggers the waking event.
- **The Cards**: Represented as a physical barrier that covers Alice, symbolizing the end of the fantasy world.
- **Wonderland**: The dream realm containing the trial, executions, and nonsensical logic.

Procedures And API Details
N/A (Narrative fiction contains no technical procedures or APIs).

Nuance Or Contradictions
The story explicitly contradicts its own reality by revealing that the previous chapters were a dream. This creates a meta-fictional layer where the "real world" is mundane nature (riverbank, tea time), contrasting sharply with the chaotic fantasy of Wonderland. The text resolves the absurdity of the trial and executions as mere figments of imagination.

Candidate Wiki Hints
- **Dream Within a Dream**: A literary technique illustrated by Alice's awakening and her sister's subsequent dream.
- **Alice's Growth and Legacy**: The thematic conclusion focusing on childhood memories persisting into adulthood and the role of storytelling in preserving wonder.

