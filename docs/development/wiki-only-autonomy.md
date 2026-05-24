# Wiki-Only Autonomy

`nemo -maintain-wiki` runs the self-maintenance loop for the production
knowledge layer. It reads only `wiki/` as input and never selects sources from
`raw/`, `drafts/`, or `evals/`.

This follows the boundary described by `wiki/sources/llm-wiki.md`: humans
provide sources and intent, while the system keeps the wiki structured,
interlinked, indexed, logged, and explicit about maintenance work. It is not a
source-discovery daemon and it does not autonomously decide which external
material should enter the knowledge base.

The maintainer supports four modes:

- `report` mode scans the wiki and writes a maintenance report without editing
  wiki files.
- `safe` mode applies mechanical repairs that do not require semantic judgment,
  then appends a `lint` entry to `wiki/log.md` only when wiki files changed.
- `propose` mode sends the maintenance task queue and a wiki-only snapshot to
  the configured model, then records a proposed JSON patch without applying it.
- `auto` mode runs `safe`, asks the model for semantic maintenance, applies the
  proposed wiki changes, reruns lint, and rolls the proposal back if the gate
  fails.
- All modes classify lint findings into maintenance tasks, making the next LLM
  maintenance pass explicit instead of burying drift in raw lint output.

Current safe repairs are limited to keeping `wiki/index.md` consistent with
existing knowledge pages: missing entries are added, duplicate entries are
deduplicated, and stale entries for missing pages are removed.

Semantic tasks such as orphan pages, broken wikilinks, schema repairs, and log
anomalies are model-assisted in `propose` and `auto` modes. That matches the
LLM-wiki pattern: the LLM maintains summaries, links, concepts, topics, the
index, and the log, while humans still steer source choice and interpretation.

## Manual Runs

Build the CLI locally:

```sh
go build -o .bin/nemo ./cmd/nemo
```

Run a read-only report:

```sh
.bin/nemo -maintain-wiki -mode report -out-dir .wiki-maintain/latest
```

Run the safe maintainer:

```sh
.bin/nemo -maintain-wiki -mode safe -out-dir .wiki-maintain/latest
```

Generate a model-assisted proposal without applying it:

```sh
.bin/nemo -provider deepseek -profile stable -maintain-wiki -mode propose -out-dir .wiki-maintain/latest
```

Run gated semantic autonomy:

```sh
.bin/nemo -provider deepseek -profile stable -maintain-wiki -mode auto -out-dir .wiki-maintain/latest
```

Each run writes:

- `.wiki-maintain/latest/wiki-maintain.json`
- `.wiki-maintain/latest/wiki-maintain.md`

The report includes two sections:

- `Actions`: deterministic repairs that were applied or would be applied by
  `safe` mode.
- `Maintenance Tasks`: remaining wiki maintenance work derived from lint
  findings, such as orphan repair or broken wikilink review.
- `Proposal`: model-generated changes when `propose` or `auto` mode is used.

## Unattended Runs

Use a single-instance lock so two maintainers never rewrite the wiki at the
same time:

```sh
run_id=$(date +%Y%m%d-%H%M%S)
flock -n /tmp/nemo-wiki-maintain.lock \
  .bin/nemo -maintain-wiki -mode safe -out-dir ".wiki-maintain/$run_id"
```

For full LLM wiki autonomy, run `auto` mode under the same lock:

```sh
run_id=$(date +%Y%m%d-%H%M%S)
flock -n /tmp/nemo-wiki-maintain.lock \
  .bin/nemo -provider deepseek -profile stable -maintain-wiki -mode auto \
    -out-dir ".wiki-maintain/$run_id"
```

`auto` mode is suitable for background use only when the configured model is the
intended production maintainer. It never edits `raw/`, and it rolls back wiki
changes that increase lint issues or introduce lint errors.

## systemd User Units

Example web service:

```ini
[Unit]
Description=nemo-knows web console

[Service]
WorkingDirectory=/home/karo/workspace/huic/nemo-knows
ExecStart=/home/karo/workspace/huic/nemo-knows/.bin/nemo-web -addr 127.0.0.1:8787
Restart=on-failure

[Install]
WantedBy=default.target
```

Example autonomous maintainer service:

```ini
[Unit]
Description=nemo-knows wiki maintainer

[Service]
Type=oneshot
WorkingDirectory=/home/karo/workspace/huic/nemo-knows
ExecStart=/usr/bin/flock -n /tmp/nemo-wiki-maintain.lock /usr/bin/fish -lc 'set run_id (date +%%Y%%m%%d-%%H%%M%%S); .bin/nemo -provider deepseek -profile stable -maintain-wiki -mode auto -out-dir ".wiki-maintain/$run_id"'
```

Example timer:

```ini
[Unit]
Description=Run nemo-knows wiki maintainer periodically

[Timer]
OnBootSec=5min
OnUnitActiveSec=1h
Persistent=true

[Install]
WantedBy=timers.target
```

## Scope Rules

- The maintainer treats `wiki/` as the only production input.
- Hidden directories under `wiki/`, such as `wiki/.maintain/`, are ignored by
  the page scanner.
- Reports are operational artifacts; they do not become wiki pages.
- If a run changes no wiki files, it does not append to `wiki/log.md`.
