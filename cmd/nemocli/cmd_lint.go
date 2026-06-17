package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/huic/nemo-knows/internal/cliroot"
	"github.com/huic/nemo-knows/internal/runslog"
	"github.com/huic/nemo-knows/internal/wikilint"
)

// cmdLint is the V0 read-only audit pass. It calls internal/wikilint
// against $WIKI_ROOT and either prints the rendered report to stderr
// (when --out-dir is empty) or writes both Markdown and JSON forms into
// the given directory.
//
// Exit code: 1 if any "error"-level issues exist.
func cmdLint(args []string, g *globalConfig) int {
	fs := flag.NewFlagSet("nemocli lint", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	outDir := fs.String("out-dir", "", "directory to write wiki-lint.json/wiki-lint.md")
	format := fs.String("format", "md", "report format when writing to --out-dir (md|json)")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if err := cliroot.RequireInitialized(g.WikiRoot); err != nil {
		fmt.Fprintln(os.Stderr, "lint:", err)
		return 2
	}

	start := time.Now()
	result, err := wikilint.LintWiki(g.WikiRoot)
	if err != nil {
		fmt.Fprintln(os.Stderr, "lint:", err)
		logLintResult(g, start, "fail", err.Error())
		return 1
	}

	md := renderLintMarkdown(result)
	if *outDir == "" {
		fmt.Fprint(os.Stderr, md)
	} else {
		if err := os.MkdirAll(*outDir, 0o755); err != nil {
			fmt.Fprintln(os.Stderr, "lint: create out dir:", err)
			return 1
		}
		switch strings.ToLower(*format) {
		case "json":
			payload, _ := json.MarshalIndent(result, "", "  ")
			if err := os.WriteFile(filepath.Join(*outDir, "wiki-lint.json"), append(payload, '\n'), 0o644); err != nil {
				fmt.Fprintln(os.Stderr, "lint: write json:", err)
				return 1
			}
		default:
			if err := os.WriteFile(filepath.Join(*outDir, "wiki-lint.md"), []byte(md), 0o644); err != nil {
				fmt.Fprintln(os.Stderr, "lint: write md:", err)
				return 1
			}
		}
	}
	hadErrors := result.Summary.ByLevel["error"] > 0
	status := "ok"
	if hadErrors {
		status = "fail"
	}
	logLintResult(g, start, status, "")
	if hadErrors {
		return 1
	}
	return 0
}

func renderLintMarkdown(result wikilint.Result) string {
	var b strings.Builder
	b.WriteString("# Wiki Lint Report\n\n")
	b.WriteString(fmt.Sprintf("- total issues: %d\n", result.Summary.Total))
	b.WriteString(fmt.Sprintf("- pages checked: %d\n", result.Summary.PageCount))
	if result.Summary.ByLevel != nil {
		b.WriteString(fmt.Sprintf("- by level: %v\n", result.Summary.ByLevel))
	}
	b.WriteString("\n## Issues\n\n")
	if len(result.Issues) == 0 {
		b.WriteString("(none)\n")
		return b.String()
	}
	for _, issue := range result.Issues {
		b.WriteString(fmt.Sprintf("- `%s` `%s` %s: %s\n", issue.Level, issue.Code, issue.Path, issue.Message))
	}
	return b.String()
}

func logLintResult(g *globalConfig, start time.Time, status, errMsg string) {
	entry := runslog.RunEntry{
		RunID:      newRunIDForCmd(start),
		Subcommand: "lint",
		Start:      start,
		End:        time.Now(),
		WikiRoot:   g.WikiRoot,
		Status:     status,
		Error:      errMsg,
	}
	if err := runslog.AppendRun(entry); err != nil && g.Verbose {
		fmt.Fprintln(os.Stderr, "warn: write runs log:", err)
	}
}
