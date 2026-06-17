package main

import (
	"fmt"
	"io"
)

const usageText = `nemocli — single-instance wiki automation (V0)

Usage:
  nemocli init                                   # create $WIKI_ROOT scaffold; in tty mode prompt for API key
  nemocli ingest <file>                          # end-to-end ingest, auto-apply into $WIKI_ROOT/wiki/
  nemocli ingest --resume <checkpoint-path>      # resume a previously failed run
  nemocli lint                                   # read-only audit of $WIKI_ROOT/wiki/
  nemocli maintain --mode <name>                 # maintenance pass (report | safe | propose | auto)
  nemocli once [--mode <name>]                   # lint + maintain (systemd timer entry point)
  nemocli logs [--follow] [--gc] [--dry-run]     # read runs log; --gc removes daily files older than retention
  nemocli gc [--days <N>] [--dry-run]            # manually remove old /tmp/nemo-checkpoints/ directories
  nemocli help                                   # this message
  nemocli version

Global flags (must precede the subcommand):
  --wiki-root <path>       override $NEMO_WIKI_ROOT (default ~/.wiki)
  --provider <name>        deepseek | llama (default: env or deepseek)
  --profile <name>         fast | stable | deep | fallback (default stable)
  --verbose, -v            extra stderr logging

V1+ commands (interface placeholders only):
  nemocli serve [--watch <dir>] [--every <duration>]
  nemocli query <q> [--file] [--approve] [--llm-review]

Environment:
  NEMO_WIKI_ROOT                  override default wiki root
  NEMO_MODEL_PROVIDER             deepseek (default) | llama
  NEMO_DEEPSEEK_API_KEY           required for the DeepSeek backend
  NEMO_LLAMA_CLI, NEMO_LLAMA_MODEL required for the llama backend
  NEMO_ONCE_MODE                  default mode for 'nemocli once' (default: report)
  NEMO_RUNS_RETENTION_DAYS        days to keep ~/.local/state/nemo/runs/ (default 30)
  NEMO_CHECKPOINT_RETENTION_DAYS  days for 'nemocli gc' (default 7)
  XDG_STATE_HOME                  honors XDG layout for runs log
`

func printUsage(w io.Writer) {
	fmt.Fprint(w, usageText)
}
