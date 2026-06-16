package wikimaint

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestReportModeDoesNotModifyWiki(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "wiki/index.md", indexWithConcepts("- [stale](concepts/stale.md) — Stale.\n"))
	writeFile(t, root, "wiki/log.md", logContent())
	writeFile(t, root, "wiki/concepts/known.md", conceptPage("Known"))
	writeFile(t, root, "raw/stale.md", "[[known]] should not be scanned from raw.\n")

	beforeIndex := readFile(t, root, "wiki/index.md")
	beforeLog := readFile(t, root, "wiki/log.md")
	result, err := Maintain(root, Options{
		Mode:   ModeReport,
		OutDir: filepath.Join(root, ".wiki-maintain"),
	})
	if err != nil {
		t.Fatalf("Maintain returned error: %v", err)
	}

	if result.Changed {
		t.Fatal("report mode must not mark wiki as changed")
	}
	if got := readFile(t, root, "wiki/index.md"); got != beforeIndex {
		t.Fatalf("report mode changed index:\n%s", got)
	}
	if got := readFile(t, root, "wiki/log.md"); got != beforeLog {
		t.Fatalf("report mode changed log:\n%s", got)
	}
	if _, err := os.Stat(filepath.Join(root, ".wiki-maintain", "wiki-maintain.md")); err != nil {
		t.Fatalf("expected markdown report: %v", err)
	}
	if !hasAction(result.Actions, "index-add") || !hasAction(result.Actions, "index-remove") {
		t.Fatalf("expected report actions for index sync, got %#v", result.Actions)
	}
	if !hasTask(result.Tasks, "index-sync") || !hasTask(result.Tasks, "orphan-page") {
		t.Fatalf("expected report task queue to include safe and semantic tasks, got %#v", result.Tasks)
	}
	report := readFile(t, root, ".wiki-maintain/wiki-maintain.md")
	if !strings.Contains(report, "## Maintenance Tasks") {
		t.Fatalf("expected task queue in report:\n%s", report)
	}
}

func TestSafeModeSyncsIndexLogsAndIsIdempotent(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "wiki/index.md", indexWithConcepts(strings.Join([]string{
		"- [[known]] — Legacy entry.",
		"- [known](concepts/known.md) — Duplicate.",
		"- [stale](concepts/stale.md) — Stale.",
	}, "\n")+"\n"))
	writeFile(t, root, "wiki/log.md", logContent())
	writeFile(t, root, "wiki/concepts/known.md", conceptPage("Known"))
	writeFile(t, root, "wiki/concepts/missing-from-index.md", conceptPage("Missing From Index"))

	result, err := Maintain(root, Options{
		Mode:   ModeSafe,
		OutDir: filepath.Join(root, ".wiki-maintain", "run-1"),
		Today:  time.Date(2026, 5, 24, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("Maintain returned error: %v", err)
	}
	if !result.Changed {
		t.Fatal("safe mode should report changed wiki")
	}
	index := readFile(t, root, "wiki/index.md")
	if strings.Count(index, "[known](concepts/known.md)") != 1 {
		t.Fatalf("expected duplicate known entry to be removed:\n%s", index)
	}
	if strings.Contains(index, "stale.md") {
		t.Fatalf("expected stale entry to be removed:\n%s", index)
	}
	if !strings.Contains(index, "[missing-from-index](concepts/missing-from-index.md)") {
		t.Fatalf("expected missing page to be added:\n%s", index)
	}
	log := readFile(t, root, "wiki/log.md")
	if !strings.Contains(log, "## [2026-05-24] lint | wiki autonomous maintenance") {
		t.Fatalf("expected lint log entry:\n%s", log)
	}

	indexAfterFirst := readFile(t, root, "wiki/index.md")
	logAfterFirst := readFile(t, root, "wiki/log.md")
	second, err := Maintain(root, Options{
		Mode:   ModeSafe,
		OutDir: filepath.Join(root, ".wiki-maintain", "run-2"),
		Today:  time.Date(2026, 5, 24, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("second Maintain returned error: %v", err)
	}
	if second.Changed {
		t.Fatalf("expected second safe run to be idempotent, actions: %#v", second.Actions)
	}
	if got := readFile(t, root, "wiki/index.md"); got != indexAfterFirst {
		t.Fatalf("second run changed index:\n%s", got)
	}
	if got := readFile(t, root, "wiki/log.md"); got != logAfterFirst {
		t.Fatalf("second run changed log:\n%s", got)
	}
}

func TestSafeModeIgnoresHiddenMaintenanceReports(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "wiki/index.md", indexWithConcepts("- [known](concepts/known.md) — Known.\n"))
	writeFile(t, root, "wiki/log.md", logContent())
	writeFile(t, root, "wiki/concepts/known.md", conceptPage("Known"))
	writeFile(t, root, "wiki/.maintain/report.md", "# Not a wiki page\n")

	result, err := Maintain(root, Options{Mode: ModeSafe})
	if err != nil {
		t.Fatalf("Maintain returned error: %v", err)
	}
	if result.Changed {
		t.Fatalf("hidden report directory should not trigger index changes: %#v", result.Actions)
	}
}

func TestSafeModeNormalizesLegacyProseAndStrandedPlaceholder(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "wiki/index.md", `---
title: Index
kind: index
---

# Index

## Sources

(none yet)

## Entities

(none yet)

## Concepts

(none yet)

## Topics

_Cross-cutting syntheses, comparisons, derived insights — including
high-value query answers filed back from chat._

(none yet)
- [persistent-wiki](topics/persistent-wiki.md) — Persistent wiki architecture.
`)
	writeFile(t, root, "wiki/log.md", logContent())
	writeFile(t, root, "wiki/topics/persistent-wiki.md", `---
title: Persistent Wiki Architecture
kind: topic
sources:
  - wiki/sources/example.md
confidence: medium
---

# Persistent Wiki Architecture
`)

	result, err := Maintain(root, Options{
		Mode:   ModeSafe,
		Today:  time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC),
		OutDir: filepath.Join(root, ".wiki-maintain"),
	})
	if err != nil {
		t.Fatalf("Maintain returned error: %v", err)
	}
	if !result.Changed {
		t.Fatal("safe mode should normalize legacy prose and placeholder")
	}
	if !hasAction(result.Actions, "index-strip-prose") {
		t.Fatalf("expected prose-stripping action, got %#v", result.Actions)
	}
	index := readFile(t, root, "wiki/index.md")
	if strings.Contains(index, "Cross-cutting syntheses") {
		t.Fatalf("expected per-section prose to be removed:\n%s", index)
	}
	if strings.Contains(index, "## Topics\n\n(none yet)") {
		t.Fatalf("expected stranded placeholder to be removed once a real entry exists:\n%s", index)
	}
	if !strings.Contains(index, "## Topics\n\n- [persistent-wiki](topics/persistent-wiki.md) — Persistent wiki architecture.\n") {
		t.Fatalf("expected curated topic entry to survive normalization:\n%s", index)
	}

	second, err := Maintain(root, Options{
		Mode:  ModeSafe,
		Today: time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("second Maintain returned error: %v", err)
	}
	if second.Changed {
		t.Fatalf("expected normalized index to be idempotent, actions: %#v", second.Actions)
	}
}

func TestSafeModeRestoresPlaceholderWhenSectionEmpties(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "wiki/index.md", indexWithConcepts("- [gone](concepts/gone.md) — Removed page.\n"))
	writeFile(t, root, "wiki/log.md", logContent())

	result, err := Maintain(root, Options{Mode: ModeSafe, Today: time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC)})
	if err != nil {
		t.Fatalf("Maintain returned error: %v", err)
	}
	if !result.Changed {
		t.Fatal("safe mode should remove the stale entry")
	}
	index := readFile(t, root, "wiki/index.md")
	if strings.Contains(index, "gone") {
		t.Fatalf("expected stale entry to be removed:\n%s", index)
	}
	if !strings.Contains(index, "## Concepts\n\n(none yet)\n") {
		t.Fatalf("expected placeholder to be restored for the emptied section:\n%s", index)
	}
}

func TestTaskQueueMapsLintFindingsToLLMMaintenanceWork(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "wiki/index.md", indexWithConcepts("- [known](concepts/known.md) — Known.\n"))
	writeFile(t, root, "wiki/log.md", logContent())
	writeFile(t, root, "wiki/concepts/known.md", `---
title: Known
kind: concept
sources:
  - wiki/sources/example.md
tags: [test]
confidence: medium
---

# Known

This page links to [[missing-target]].
`)
	writeFile(t, root, "wiki/concepts/lonely.md", conceptPage("Lonely"))

	result, err := Maintain(root, Options{Mode: ModeReport})
	if err != nil {
		t.Fatalf("Maintain returned error: %v", err)
	}
	for _, kind := range []string{"missing-wikilink-target", "orphan-page"} {
		if !hasTask(result.Tasks, kind) {
			t.Fatalf("expected task kind %q in %#v", kind, result.Tasks)
		}
	}
	for _, task := range result.Tasks {
		if task.Kind == "missing-wikilink-target" && !strings.Contains(task.Recommendation, "retargeting") {
			t.Fatalf("missing wikilink task has weak recommendation: %#v", task)
		}
	}
}

func TestProposeModeGeneratesProposalWithoutApplying(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "wiki/index.md", indexWithConcepts("- [known](concepts/known.md) — Known.\n"))
	writeFile(t, root, "wiki/log.md", logContent())
	writeFile(t, root, "wiki/concepts/known.md", `---
title: Known
kind: concept
sources:
  - wiki/sources/example.md
tags: [test]
confidence: medium
---

# Known

This page links to [[missing-target]].
`)

	proposed := strings.ReplaceAll(conceptPage("Known"), "# Known\n", "# Known\n\nThis page mentions missing target without a broken wikilink.\n")
	result, err := Maintain(root, Options{
		Mode:           ModePropose,
		PromptTemplate: "tasks {{MAINTENANCE_TASKS}}\nwiki {{WIKI_SNAPSHOT}}",
		Generator: fakeGenerator{Output: `{
  "notes": ["retargeted broken link"],
  "changes": [{"path": "wiki/concepts/known.md", "content": ` + strconvQuote(proposed) + `}]
}`},
	})
	if err != nil {
		t.Fatalf("Maintain returned error: %v", err)
	}
	if result.Proposal == nil || len(result.Proposal.Changes) != 1 {
		t.Fatalf("expected proposal change, got %#v", result.Proposal)
	}
	if got := readFile(t, root, "wiki/concepts/known.md"); strings.Contains(got, "without a broken wikilink") {
		t.Fatalf("propose mode applied content unexpectedly:\n%s", got)
	}
}

func TestAutoModeAppliesGatedProposalAndLogs(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "wiki/index.md", indexWithConcepts("- [known](concepts/known.md) — Known.\n"))
	writeFile(t, root, "wiki/log.md", logContent())
	writeFile(t, root, "wiki/concepts/known.md", `---
title: Known
kind: concept
sources:
  - wiki/sources/example.md
tags: [test]
confidence: medium
---

# Known

This page links to [[missing-target]].
`)

	proposed := strings.ReplaceAll(conceptPage("Known"), "# Known\n", "# Known\n\nThis page mentions missing target without a broken wikilink.\n")
	result, err := Maintain(root, Options{
		Mode:           ModeAuto,
		PromptTemplate: "tasks {{MAINTENANCE_TASKS}}\nwiki {{WIKI_SNAPSHOT}}",
		Today:          time.Date(2026, 5, 24, 0, 0, 0, 0, time.UTC),
		Generator: fakeGenerator{Output: `{
  "notes": ["removed unsupported link"],
  "changes": [{"path": "wiki/concepts/known.md", "content": ` + strconvQuote(proposed) + `}]
}`},
	})
	if err != nil {
		t.Fatalf("Maintain returned error: %v", err)
	}
	if !result.Changed {
		t.Fatal("auto mode should apply the gated proposal")
	}
	page := readFile(t, root, "wiki/concepts/known.md")
	if !strings.Contains(page, "without a broken wikilink") || strings.Contains(page, "[[missing-target]]") {
		t.Fatalf("expected proposed page content to be applied:\n%s", page)
	}
	log := readFile(t, root, "wiki/log.md")
	if !strings.Contains(log, "lint | wiki semantic maintenance") {
		t.Fatalf("expected semantic maintenance log entry:\n%s", log)
	}
}

func TestAutoModeRollsBackProposalThatFailsLintGate(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "wiki/index.md", indexWithConcepts("- [known](concepts/known.md) — Known.\n"))
	writeFile(t, root, "wiki/log.md", logContent())
	original := `---
title: Known
kind: concept
sources:
  - wiki/sources/example.md
tags: [test]
confidence: medium
---

# Known

This page links to [[missing-target]].
`
	writeFile(t, root, "wiki/concepts/known.md", original)

	bad := strings.ReplaceAll(conceptPage("Known"), "confidence: medium", "confidence: impossible")
	_, err := Maintain(root, Options{
		Mode:           ModeAuto,
		PromptTemplate: "tasks {{MAINTENANCE_TASKS}}\nwiki {{WIKI_SNAPSHOT}}",
		Generator: fakeGenerator{Output: `{
  "notes": ["bad edit"],
  "changes": [{"path": "wiki/concepts/known.md", "content": ` + strconvQuote(bad) + `}]
}`},
	})
	if err == nil {
		t.Fatal("expected auto mode to reject invalid proposal")
	}
	if got := readFile(t, root, "wiki/concepts/known.md"); got != original {
		t.Fatalf("expected rollback to restore original content:\n%s", got)
	}
}

func writeFile(t *testing.T, root string, rel string, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func readFile(t *testing.T, root string, rel string) string {
	t.Helper()
	content, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatalf("read %s: %v", rel, err)
	}
	return string(content)
}

func indexWithConcepts(entries string) string {
	return `---
title: Index
kind: index
---

# Index

## Sources

(none yet)

## Entities

(none yet)

## Concepts

` + entries + `
## Topics

(none yet)
`
}

func logContent() string {
	return `---
title: Log
kind: log
---

# Log

## [2026-05-16] note | ok
Open: none.
`
}

func conceptPage(title string) string {
	return `---
title: ` + title + `
kind: concept
sources:
  - wiki/sources/example.md
tags: [test]
confidence: medium
---

# ` + title + `
`
}

func hasAction(actions []Action, actionType string) bool {
	for _, action := range actions {
		if action.Type == actionType {
			return true
		}
	}
	return false
}

func hasTask(tasks []Task, kind string) bool {
	for _, task := range tasks {
		if task.Kind == kind {
			return true
		}
	}
	return false
}

type fakeGenerator struct {
	Output string
}

func (g fakeGenerator) Generate(ctx context.Context, prompt string) (string, error) {
	return g.Output, nil
}

func strconvQuote(s string) string {
	quoted, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(quoted)
}
