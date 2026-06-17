package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/huic/nemo-knows/cmd/nemocli/initassets"
	"github.com/huic/nemo-knows/internal/cliroot"
)

// cmdInit scaffolds a fresh $WIKI_ROOT. It is idempotent: re-running on
// an initialized root only fills in any missing files unless --force is
// passed, in which case existing files are overwritten with the
// templates again.
//
// In tty mode the command also prompts for the DeepSeek API key and
// writes a 0600 .env when the user supplies one. The R3 contract in
// §7.2 of the v3 design covers the exact prompt strings and the
// non-tty fallback.
func cmdInit(args []string, g *globalConfig) int {
	fs := flag.NewFlagSet("nemocli init", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	force := fs.Bool("force", false, "overwrite existing init files instead of refusing")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	wikiRoot := g.WikiRoot
	wikiDir := cliroot.WikiContentDir(wikiRoot)

	if !*force {
		// Refuse to clobber a non-empty existing wiki content directory
		// unless force is set; this protects against accidentally
		// running init twice on a populated wiki.
		if hasContent, err := dirHasContent(wikiDir); err != nil {
			fmt.Fprintln(os.Stderr, "init:", err)
			return 1
		} else if hasContent {
			fmt.Fprintf(os.Stderr,
				"init: %s already has content; pass --force to overwrite or run nemocli ingest/lint instead\n",
				wikiDir)
			return 1
		}
	}

	if err := writeScaffold(wikiRoot, *force); err != nil {
		fmt.Fprintln(os.Stderr, "init:", err)
		return 1
	}

	wrote, msg := maybePromptAndWriteEnv(wikiRoot)

	fmt.Printf("$WIKI_ROOT initialized at %s\n", wikiRoot)
	if msg != "" {
		fmt.Fprintln(os.Stderr, msg)
	}
	fmt.Println()
	fmt.Println("Next:")
	if wrote {
		fmt.Println("  nemocli ingest <file>")
	} else {
		fmt.Printf("  cp %s/.env.example %s/.env\n", wikiRoot, wikiRoot)
		fmt.Printf("  $EDITOR %s/.env       # fill NEMO_DEEPSEEK_API_KEY\n", wikiRoot)
		fmt.Println("  nemocli ingest <file>")
	}
	return 0
}

func writeScaffold(wikiRoot string, force bool) error {
	if err := os.MkdirAll(wikiRoot, 0o755); err != nil {
		return fmt.Errorf("create wiki root: %w", err)
	}
	wikiDir := cliroot.WikiContentDir(wikiRoot)
	for _, sub := range []string{"sources", "entities", "concepts", "topics", "assets"} {
		if err := os.MkdirAll(filepath.Join(wikiDir, sub), 0o755); err != nil {
			return fmt.Errorf("create wiki subdir %s: %w", sub, err)
		}
	}

	today := time.Now().Format("2006-01-02")

	files := []scaffoldFile{
		{
			Path:   filepath.Join(wikiDir, "index.md"),
			Body:   defaultIndexBody(today),
			IfNone: true,
		},
		{
			Path:   filepath.Join(wikiDir, "log.md"),
			Body:   defaultLogBody(),
			IfNone: true,
		},
		{
			Path:   filepath.Join(wikiDir, "AGENTS.md"),
			Body:   string(initassets.WikiAgentsTemplate()),
			IfNone: true,
		},
		{
			Path:   filepath.Join(wikiRoot, ".env.example"),
			Body:   string(initassets.EnvExampleTemplate()),
			IfNone: true,
		},
	}
	for _, f := range files {
		if err := writeScaffoldFile(f, force); err != nil {
			return err
		}
	}
	return nil
}

type scaffoldFile struct {
	Path   string
	Body   string
	IfNone bool // if true and force is false, only write when path is missing
}

func writeScaffoldFile(f scaffoldFile, force bool) error {
	if !force && f.IfNone {
		if _, err := os.Stat(f.Path); err == nil {
			return nil // keep existing content
		}
	}
	if err := os.MkdirAll(filepath.Dir(f.Path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(f.Path, []byte(f.Body), 0o644)
}

func defaultIndexBody(date string) string {
	return fmt.Sprintf(`---
title: Index
kind: index
updated: %s
---

# Index

## Sources

(none yet)

## Entities

(none yet)

## Concepts

(none yet)

## Topics

(none yet)
`, date)
}

func defaultLogBody() string {
	return `---
title: Log
kind: log
---

# Log

Append-only record of every ingest, filed query answer, lint pass, and
schema change. Each entry begins with a heading in the canonical format
so it is greppable:

` + "```" + `
## [YYYY-MM-DD] <action> | <subject>
` + "```" + `

` + "`<action>`" + ` is one of: ` + "`ingest`, `query-filed`, `lint`, `schema-change`, `note`" + `.
The full format and conventions are defined in [` + "`AGENTS.md`" + `](AGENTS.md) §6.

To skim recent activity:

` + "```sh" + `
grep "^## \[" wiki/log.md | tail -10
` + "```" + `
`
}

// dirHasContent reports whether dir exists and contains any non-empty
// content. A bare directory or missing path is considered "empty" so
// init can fill it in.
func dirHasContent(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	for _, e := range entries {
		if e.IsDir() {
			has, err := dirHasContent(filepath.Join(dir, e.Name()))
			if err != nil {
				return false, err
			}
			if has {
				return true, nil
			}
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		if info.Size() > 0 {
			return true, nil
		}
	}
	return false, nil
}

// maybePromptAndWriteEnv handles the R3 tty/non-tty branches.
//
// Returns (wroteEnv, message). The message goes to stderr; the boolean
// drives the "Next:" hints printed on stdout.
func maybePromptAndWriteEnv(wikiRoot string) (bool, string) {
	envPath := filepath.Join(wikiRoot, ".env")
	if _, err := os.Stat(envPath); err == nil {
		return true, fmt.Sprintf("%s already exists; leaving it untouched", envPath)
	}

	if !isStdinTTY() {
		return false, "stdin is not a tty; .env not written. Edit .env.example then mv to .env, " +
			"or `set -x NEMO_DEEPSEEK_API_KEY <key>` (fish)."
	}

	fmt.Print("Enter NEMO_DEEPSEEK_API_KEY (or leave empty to skip): ")
	key, err := readLine(os.Stdin)
	if err != nil {
		return false, fmt.Sprintf("could not read key from stdin: %v; .env not written.", err)
	}
	key = strings.TrimSpace(key)
	if key == "" {
		return false, "no API key provided; .env not written. " +
			"Edit .env.example then mv to .env when you are ready."
	}
	if err := writeEnvFile(envPath, key); err != nil {
		return false, fmt.Sprintf("could not write %s: %v", envPath, err)
	}
	return true, fmt.Sprintf("wrote %s with provided NEMO_DEEPSEEK_API_KEY (mode 0600)", envPath)
}

// isStdinTTY reports whether nemocli was invoked from an interactive
// terminal. We use the standard "character device" mode bit instead of
// pulling in golang.org/x/term so the binary stays dep-free for V0.
func isStdinTTY() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeCharDevice != 0
}

// readLine reads a single line of input (LF or EOF terminated). It does
// NOT suppress echo: that would require x/term, which V0 deliberately
// avoids. Users in shared environments should set NEMO_DEEPSEEK_API_KEY
// via env instead, or `chmod 0600` $WIKI_ROOT/.env after writing.
func readLine(r io.Reader) (string, error) {
	br := bufio.NewReader(r)
	line, err := br.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimRight(line, "\n\r"), nil
}

func writeEnvFile(path string, apiKey string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	body := fmt.Sprintf("NEMO_MODEL_PROVIDER=deepseek\nNEMO_DEEPSEEK_API_KEY=%s\n", apiKey)
	return os.WriteFile(path, []byte(body), 0o600)
}
