package wikimaint

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/huic/nemo-knows/internal/wikilint"
)

const (
	ModeReport  = "report"
	ModeSafe    = "safe"
	ModePropose = "propose"
	ModeAuto    = "auto"
)

var (
	frontmatterBlockRE = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*`)
	indexEntryRE       = regexp.MustCompile(`^- \[\[([^\]|#]+)(?:[|#][^\]]*)?\]\].*$`)
	fmLineRE           = regexp.MustCompile(`(?m)^([A-Za-z0-9_-]+):\s*(.+?)\s*$`)
	taskIDCleanRE      = regexp.MustCompile(`[^a-z0-9-]+`)
)

type Options struct {
	Mode           string
	OutDir         string
	Today          time.Time
	Generator      Generator
	PromptTemplate string
}

type Generator interface {
	Generate(ctx context.Context, prompt string) (string, error)
}

type Result struct {
	Mode     string           `json:"mode"`
	Changed  bool             `json:"changed"`
	Before   wikilint.Result  `json:"before"`
	After    *wikilint.Result `json:"after,omitempty"`
	Actions  []Action         `json:"actions"`
	Tasks    []Task           `json:"tasks"`
	Proposal *Proposal        `json:"proposal,omitempty"`
}

type Action struct {
	Type    string `json:"type"`
	Path    string `json:"path"`
	Message string `json:"message"`
}

type Task struct {
	ID             string `json:"id"`
	Kind           string `json:"kind"`
	Level          string `json:"level"`
	Path           string `json:"path"`
	Summary        string `json:"summary"`
	Recommendation string `json:"recommendation"`
	AutoSafe       bool   `json:"auto_safe"`
}

type Proposal struct {
	Notes   []string `json:"notes"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

type wikiPage struct {
	Path  string
	Slug  string
	Kind  string
	Title string
}

// Maintain runs a wiki-only maintenance pass. Report mode is read-only; safe
// mode applies deterministic repairs that do not require semantic judgment.
func Maintain(root string, opts Options) (Result, error) {
	mode := opts.Mode
	if mode == "" {
		mode = ModeReport
	}
	if mode != ModeReport && mode != ModeSafe && mode != ModePropose && mode != ModeAuto {
		return Result{}, fmt.Errorf("unsupported wiki maintenance mode %q", mode)
	}

	before, err := wikilint.LintWiki(root)
	if err != nil {
		return Result{}, fmt.Errorf("lint wiki before maintenance: %w", err)
	}
	result := Result{Mode: mode, Before: before}

	pages, err := readWikiPages(root)
	if err != nil {
		return Result{}, err
	}
	indexPath := filepath.Join(root, "wiki", "index.md")
	indexContent, err := os.ReadFile(indexPath)
	if err != nil {
		return Result{}, fmt.Errorf("read wiki index: %w", err)
	}
	updatedIndex, indexActions := syncIndex(string(indexContent), pages)
	result.Actions = append(result.Actions, indexActions...)
	if mode == ModeReport || mode == ModePropose {
		result.Tasks = buildTaskQueue(before.Issues, indexActions)
	}

	if (mode == ModeSafe || mode == ModeAuto) && len(indexActions) > 0 && updatedIndex != string(indexContent) {
		if err := os.WriteFile(indexPath, []byte(updatedIndex), 0o644); err != nil {
			return Result{}, fmt.Errorf("write wiki index: %w", err)
		}
		result.Changed = true
		result.Actions = append(result.Actions, Action{
			Type:    "write",
			Path:    "wiki/index.md",
			Message: "synchronized index entries with wiki pages",
		})
	}

	if result.Changed {
		if err := appendLintLog(root, opts.Today, result.Actions); err != nil {
			return Result{}, err
		}
		result.Actions = append(result.Actions, Action{
			Type:    "write",
			Path:    "wiki/log.md",
			Message: "recorded wiki maintenance lint entry",
		})
	}

	after, err := wikilint.LintWiki(root)
	if err != nil {
		return Result{}, fmt.Errorf("lint wiki after maintenance: %w", err)
	}
	result.After = &after
	if mode == ModeSafe || mode == ModeAuto {
		result.Tasks = buildTaskQueue(after.Issues, nil)
	}

	if mode == ModePropose || mode == ModeAuto {
		proposal, err := proposeSemanticMaintenance(root, opts, result.Tasks)
		if err != nil {
			return Result{}, err
		}
		result.Proposal = &proposal
		if mode == ModeAuto && len(proposal.Changes) > 0 {
			applied, err := applyProposal(root, opts.Today, proposal, result.Before.Summary.Total)
			if err != nil {
				return Result{}, err
			}
			result.Changed = result.Changed || applied
			if applied {
				result.Actions = append(result.Actions, Action{
					Type:    "semantic-apply",
					Path:    "wiki/",
					Message: "applied gated model-assisted wiki maintenance proposal",
				})
				afterAuto, err := wikilint.LintWiki(root)
				if err != nil {
					return Result{}, fmt.Errorf("lint wiki after semantic maintenance: %w", err)
				}
				result.After = &afterAuto
				result.Tasks = buildTaskQueue(afterAuto.Issues, nil)
			}
		}
	}

	if opts.OutDir != "" {
		if err := writeReport(opts.OutDir, result); err != nil {
			return Result{}, err
		}
	}

	return result, nil
}

func readWikiPages(root string) ([]wikiPage, error) {
	wikiRoot := filepath.Join(root, "wiki")
	pages := []wikiPage{}
	err := filepath.WalkDir(wikiRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if strings.HasPrefix(d.Name(), ".") && path != wikiRoot {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		repoPath := filepath.ToSlash(rel)
		if repoPath == "wiki/index.md" || repoPath == "wiki/log.md" {
			return nil
		}
		if !isKnowledgePage(repoPath) {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		frontmatter := frontmatter(string(content))
		kind := fmValue(frontmatter, "kind")
		title := fmValue(frontmatter, "title")
		if title == "" {
			title = titleFromSlug(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
		}
		pages = append(pages, wikiPage{
			Path:  repoPath,
			Slug:  strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
			Kind:  kind,
			Title: title,
		})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk wiki pages: %w", err)
	}
	sort.Slice(pages, func(i, j int) bool {
		return pages[i].Path < pages[j].Path
	})
	return pages, nil
}

func syncIndex(index string, pages []wikiPage) (string, []Action) {
	desired := desiredIndexEntries(pages)
	updated := index
	actions := []Action{}
	for _, section := range []struct {
		heading string
		kind    string
	}{
		{"## Sources", "source"},
		{"## Entities", "entity"},
		{"## Concepts", "concept"},
		{"## Topics", "topic"},
	} {
		var sectionActions []Action
		updated, sectionActions = syncIndexSection(updated, section.heading, desired[section.kind])
		actions = append(actions, sectionActions...)
	}
	return updated, actions
}

func desiredIndexEntries(pages []wikiPage) map[string]map[string]string {
	desired := map[string]map[string]string{
		"source":  {},
		"entity":  {},
		"concept": {},
		"topic":   {},
	}
	for _, page := range pages {
		if _, ok := desired[page.Kind]; !ok {
			continue
		}
		desired[page.Kind][page.Slug] = fmt.Sprintf("- [[%s]] — %s.", page.Slug, page.Title)
	}
	return desired
}

func syncIndexSection(index string, heading string, desired map[string]string) (string, []Action) {
	start := strings.Index(index, heading)
	if start == -1 {
		return appendMissingSection(index, heading, desired)
	}
	nextRel := strings.Index(index[start+len(heading):], "\n## ")
	end := len(index)
	if nextRel != -1 {
		end = start + len(heading) + nextRel
	}

	section := index[start:end]
	updatedSection, actions := syncSectionBody(heading, section, desired)
	if updatedSection == section {
		return index, actions
	}
	return index[:start] + updatedSection + index[end:], actions
}

func appendMissingSection(index string, heading string, desired map[string]string) (string, []Action) {
	var b strings.Builder
	b.WriteString(strings.TrimRight(index, "\n"))
	b.WriteString("\n\n")
	b.WriteString(heading)
	b.WriteString("\n")
	if len(desired) == 0 {
		b.WriteString("\n(none yet)\n")
	} else {
		b.WriteString("\n")
		for _, entry := range sortedEntries(desired) {
			b.WriteString(entry)
			b.WriteByte('\n')
		}
	}
	return b.String(), []Action{{
		Type:    "index-section",
		Path:    "wiki/index.md",
		Message: "added missing " + heading + " section",
	}}
}

func syncSectionBody(heading string, section string, desired map[string]string) (string, []Action) {
	lines := strings.Split(section, "\n")
	if len(lines) == 0 {
		return section, nil
	}

	kept := []string{lines[0]}
	entries := []string{}
	seen := map[string]bool{}
	actions := []Action{}

	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "(none yet)" && len(desired) > 0 {
			continue
		}
		match := indexEntryRE.FindStringSubmatch(line)
		if len(match) != 2 {
			kept = append(kept, line)
			continue
		}
		slug := match[1]
		if _, ok := desired[slug]; !ok {
			actions = append(actions, Action{
				Type:    "index-remove",
				Path:    "wiki/index.md",
				Message: "removed stale index entry [[" + slug + "]] from " + heading,
			})
			continue
		}
		if seen[slug] {
			actions = append(actions, Action{
				Type:    "index-dedupe",
				Path:    "wiki/index.md",
				Message: "removed duplicate index entry [[" + slug + "]] from " + heading,
			})
			continue
		}
		seen[slug] = true
		entries = append(entries, line)
	}

	for _, slug := range sortedSlugs(desired) {
		if seen[slug] {
			continue
		}
		entries = append(entries, desired[slug])
		actions = append(actions, Action{
			Type:    "index-add",
			Path:    "wiki/index.md",
			Message: "added missing index entry [[" + slug + "]] to " + heading,
		})
	}

	if len(entries) == 0 && !containsNoneYet(kept) {
		kept = append(trimTrailingBlankLines(kept), "", "(none yet)")
	}
	if len(entries) > 0 {
		kept = trimTrailingBlankLines(kept)
		kept = append(kept, entries...)
	}

	return strings.Join(kept, "\n"), actions
}

func sortedEntries(entries map[string]string) []string {
	slugs := sortedSlugs(entries)
	result := make([]string, 0, len(slugs))
	for _, slug := range slugs {
		result = append(result, entries[slug])
	}
	return result
}

func sortedSlugs(entries map[string]string) []string {
	slugs := make([]string, 0, len(entries))
	for slug := range entries {
		slugs = append(slugs, slug)
	}
	sort.Strings(slugs)
	return slugs
}

func trimTrailingBlankLines(lines []string) []string {
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func containsNoneYet(lines []string) bool {
	for _, line := range lines {
		if strings.TrimSpace(line) == "(none yet)" {
			return true
		}
	}
	return false
}

func appendLintLog(root string, today time.Time, actions []Action) error {
	if today.IsZero() {
		today = time.Now()
	}
	var b strings.Builder
	b.WriteString(fmt.Sprintf("\n## [%s] lint | wiki autonomous maintenance\n", today.Format("2006-01-02")))
	b.WriteString("Touched:\n")
	touched := map[string]bool{}
	for _, action := range actions {
		if action.Path == "" || touched[action.Path] {
			continue
		}
		touched[action.Path] = true
		fmt.Fprintf(&b, "- %s (updated)\n", action.Path)
	}
	b.WriteString("- wiki/log.md (updated)\n")
	b.WriteString("Open: none.\n")

	path := filepath.Join(root, "wiki", "log.md")
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open wiki log: %w", err)
	}
	defer file.Close()
	if _, err := file.WriteString(b.String()); err != nil {
		return fmt.Errorf("append wiki log: %w", err)
	}
	return nil
}

func buildTaskQueue(issues []wikilint.Issue, safeActions []Action) []Task {
	tasks := []Task{}
	for _, action := range safeActions {
		if !strings.HasPrefix(action.Type, "index-") {
			continue
		}
		tasks = append(tasks, Task{
			ID:             taskID(len(tasks)+1, "index-sync", action.Path),
			Kind:           "index-sync",
			Level:          "safe",
			Path:           action.Path,
			Summary:        action.Message,
			Recommendation: "Run `nemo -maintain-wiki -mode safe` to apply this bookkeeping repair.",
			AutoSafe:       true,
		})
	}
	for _, issue := range issues {
		tasks = append(tasks, taskFromIssue(len(tasks)+1, issue))
	}
	sort.SliceStable(tasks, func(i, j int) bool {
		if tasks[i].AutoSafe != tasks[j].AutoSafe {
			return tasks[i].AutoSafe
		}
		if tasks[i].Path == tasks[j].Path {
			return tasks[i].Kind < tasks[j].Kind
		}
		return tasks[i].Path < tasks[j].Path
	})
	return tasks
}

func proposeSemanticMaintenance(root string, opts Options, tasks []Task) (Proposal, error) {
	if opts.Generator == nil {
		return Proposal{}, fmt.Errorf("semantic wiki maintenance requires a model generator")
	}
	manualTasks := semanticTasks(tasks)
	if len(manualTasks) == 0 {
		return Proposal{Notes: []string{"no semantic maintenance tasks require model review"}}, nil
	}
	template := opts.PromptTemplate
	if strings.TrimSpace(template) == "" {
		content, err := os.ReadFile(filepath.Join(root, "prompts", "wiki-maintenance.md"))
		if err != nil {
			return Proposal{}, fmt.Errorf("read wiki maintenance prompt: %w", err)
		}
		template = string(content)
	}
	snapshot, err := wikiSnapshot(root)
	if err != nil {
		return Proposal{}, err
	}
	taskJSON, err := json.MarshalIndent(manualTasks, "", "  ")
	if err != nil {
		return Proposal{}, fmt.Errorf("encode maintenance tasks: %w", err)
	}
	rendered := strings.NewReplacer(
		"{{MAINTENANCE_TASKS}}", string(taskJSON),
		"{{WIKI_SNAPSHOT}}", snapshot,
	).Replace(template)
	raw, err := opts.Generator.Generate(context.Background(), rendered)
	if err != nil {
		return Proposal{}, fmt.Errorf("generate wiki maintenance proposal: %w", err)
	}
	proposal, err := parseProposal(raw)
	if err != nil {
		return Proposal{}, err
	}
	if err := validateProposal(root, proposal); err != nil {
		return Proposal{}, err
	}
	return proposal, nil
}

func semanticTasks(tasks []Task) []Task {
	result := []Task{}
	for _, task := range tasks {
		if task.AutoSafe {
			continue
		}
		result = append(result, task)
	}
	return result
}

func wikiSnapshot(root string) (string, error) {
	pages, err := readWikiPages(root)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	for _, page := range pages {
		content, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(page.Path)))
		if err != nil {
			return "", fmt.Errorf("read wiki page %s: %w", page.Path, err)
		}
		b.WriteString("\n--- FILE: ")
		b.WriteString(page.Path)
		b.WriteString(" ---\n")
		b.Write(content)
		if !bytes.HasSuffix(content, []byte("\n")) {
			b.WriteByte('\n')
		}
	}
	return b.String(), nil
}

func parseProposal(raw string) (Proposal, error) {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimPrefix(raw, "```json")
	raw = strings.TrimPrefix(raw, "```")
	raw = strings.TrimSuffix(raw, "```")
	raw = strings.TrimSpace(raw)
	var proposal Proposal
	if err := json.Unmarshal([]byte(raw), &proposal); err != nil {
		return Proposal{}, fmt.Errorf("parse wiki maintenance proposal JSON: %w", err)
	}
	return proposal, nil
}

func validateProposal(root string, proposal Proposal) error {
	for _, change := range proposal.Changes {
		if strings.TrimSpace(change.Path) == "" {
			return fmt.Errorf("proposal contains change with empty path")
		}
		if !isKnowledgePage(change.Path) {
			return fmt.Errorf("proposal change path outside wiki knowledge pages: %s", change.Path)
		}
		if strings.Contains(change.Path, "..") || filepath.IsAbs(change.Path) {
			return fmt.Errorf("proposal change path is unsafe: %s", change.Path)
		}
		if strings.TrimSpace(change.Content) == "" {
			return fmt.Errorf("proposal change %s has empty content", change.Path)
		}
		if err := validateProposedContent(change.Path, change.Content); err != nil {
			return err
		}
	}
	return nil
}

func validateProposedContent(path string, content string) error {
	fm := frontmatter(content)
	if fm == "" {
		return fmt.Errorf("proposal change %s is missing YAML frontmatter", path)
	}
	kind := fmValue(fm, "kind")
	if !validKindForPath(kind, path) {
		return fmt.Errorf("proposal change %s has kind %q that does not match path", path, kind)
	}
	if !strings.Contains(fm, "sources:") {
		return fmt.Errorf("proposal change %s is missing sources frontmatter", path)
	}
	confidence := fmValue(fm, "confidence")
	if confidence != "high" && confidence != "medium" && confidence != "low" {
		return fmt.Errorf("proposal change %s has invalid confidence %q", path, confidence)
	}
	return nil
}

func validKindForPath(kind string, path string) bool {
	switch {
	case strings.HasPrefix(path, "wiki/sources/"):
		return kind == "source"
	case strings.HasPrefix(path, "wiki/entities/"):
		return kind == "entity"
	case strings.HasPrefix(path, "wiki/concepts/"):
		return kind == "concept"
	case strings.HasPrefix(path, "wiki/topics/"):
		return kind == "topic"
	default:
		return false
	}
}

func applyProposal(root string, today time.Time, proposal Proposal, beforeIssues int) (bool, error) {
	if len(proposal.Changes) == 0 {
		return false, nil
	}
	backups := map[string][]byte{}
	existed := map[string]bool{}
	for _, change := range proposal.Changes {
		path := filepath.Join(root, filepath.FromSlash(change.Path))
		content, err := os.ReadFile(path)
		if err == nil {
			backups[change.Path] = content
			existed[change.Path] = true
		} else if !os.IsNotExist(err) {
			return false, fmt.Errorf("read backup for %s: %w", change.Path, err)
		}
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return false, fmt.Errorf("create wiki directory for %s: %w", change.Path, err)
		}
		if err := os.WriteFile(path, []byte(change.Content), 0o644); err != nil {
			return false, fmt.Errorf("write proposed wiki change %s: %w", change.Path, err)
		}
	}
	after, err := wikilint.LintWiki(root)
	if err != nil || after.Summary.ByLevel["error"] > 0 || after.Summary.Total > beforeIssues {
		rollbackProposal(root, proposal, backups, existed)
		if err != nil {
			return false, fmt.Errorf("lint proposed wiki changes: %w", err)
		}
		return false, fmt.Errorf("proposed wiki changes failed gates: before issues=%d after issues=%d errors=%d", beforeIssues, after.Summary.Total, after.Summary.ByLevel["error"])
	}
	actions := []Action{}
	for _, change := range proposal.Changes {
		actions = append(actions, Action{Path: change.Path})
	}
	if err := appendSemanticLog(root, today, actions, proposal.Notes); err != nil {
		return false, err
	}
	return true, nil
}

func rollbackProposal(root string, proposal Proposal, backups map[string][]byte, existed map[string]bool) {
	for _, change := range proposal.Changes {
		path := filepath.Join(root, filepath.FromSlash(change.Path))
		if existed[change.Path] {
			_ = os.WriteFile(path, backups[change.Path], 0o644)
			continue
		}
		_ = os.Remove(path)
	}
}

func appendSemanticLog(root string, today time.Time, actions []Action, notes []string) error {
	if today.IsZero() {
		today = time.Now()
	}
	var b strings.Builder
	b.WriteString(fmt.Sprintf("\n## [%s] lint | wiki semantic maintenance\n", today.Format("2006-01-02")))
	b.WriteString("Touched:\n")
	seen := map[string]bool{}
	for _, action := range actions {
		if action.Path == "" || seen[action.Path] {
			continue
		}
		seen[action.Path] = true
		fmt.Fprintf(&b, "- %s (updated)\n", action.Path)
	}
	b.WriteString("- wiki/log.md (updated)\n")
	if len(notes) == 0 {
		b.WriteString("Open: none.\n")
	} else {
		b.WriteString("Open: ")
		b.WriteString(strings.Join(notes, "; "))
		b.WriteString("\n")
	}
	path := filepath.Join(root, "wiki", "log.md")
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open wiki log: %w", err)
	}
	defer file.Close()
	if _, err := file.WriteString(b.String()); err != nil {
		return fmt.Errorf("append semantic maintenance log: %w", err)
	}
	return nil
}

func taskFromIssue(n int, issue wikilint.Issue) Task {
	task := Task{
		ID:       taskID(n, issue.Code, issue.Path),
		Kind:     issue.Code,
		Level:    issue.Level,
		Path:     issue.Path,
		Summary:  issue.Message,
		AutoSafe: false,
	}
	switch issue.Code {
	case "orphan-page":
		task.Recommendation = "Review whether the page should receive meaningful inbound links, be merged into a broader page, or remain intentionally standalone."
	case "missing-wikilink-target":
		task.Recommendation = "Repair the broken wikilink by retargeting it to an existing page or creating a sourced page if the concept is worth keeping."
	case "duplicate-index-entry":
		task.Recommendation = "Run safe maintenance to deduplicate the index entry."
		task.AutoSafe = true
	case "missing-frontmatter", "missing-kind", "invalid-kind", "missing-sources", "invalid-confidence":
		task.Recommendation = "Bring the page back into AGENTS.md schema before doing semantic edits."
	case "invalid-log-action":
		task.Recommendation = "Correct the log heading to one of the schema-approved actions without rewriting historical content."
	default:
		task.Recommendation = "Review this lint finding and decide whether to update the wiki, file a follow-up question, or leave it documented."
	}
	return task
}

func taskID(n int, kind string, path string) string {
	slug := strings.Trim(path, "/")
	slug = strings.ReplaceAll(slug, "/", "-")
	slug = strings.TrimSuffix(slug, ".md")
	slug = strings.ToLower(slug)
	slug = taskIDCleanRE.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	if slug == "" {
		slug = "wiki"
	}
	return fmt.Sprintf("task-%03d-%s-%s", n, kind, slug)
}

func writeReport(outDir string, result Result) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create maintenance output directory: %w", err)
	}
	encoded, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("encode maintenance result: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "wiki-maintain.json"), append(encoded, '\n'), 0o644); err != nil {
		return fmt.Errorf("write maintenance json: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "wiki-maintain.md"), []byte(renderReport(result)), 0o644); err != nil {
		return fmt.Errorf("write maintenance report: %w", err)
	}
	return nil
}

func renderReport(result Result) string {
	var b strings.Builder
	b.WriteString("# Wiki Maintenance Report\n\n")
	b.WriteString(fmt.Sprintf("- mode: `%s`\n", result.Mode))
	b.WriteString(fmt.Sprintf("- changed: `%t`\n", result.Changed))
	b.WriteString(fmt.Sprintf("- issues before: `%d`\n", result.Before.Summary.Total))
	if result.After != nil {
		b.WriteString(fmt.Sprintf("- issues after: `%d`\n", result.After.Summary.Total))
	}
	b.WriteString("\n## Actions\n\n")
	if len(result.Actions) == 0 {
		b.WriteString("(none)\n")
	} else {
		for _, action := range result.Actions {
			b.WriteString(fmt.Sprintf("- `%s` %s: %s\n", action.Type, action.Path, action.Message))
		}
	}
	b.WriteString("\n## Maintenance Tasks\n\n")
	if len(result.Tasks) == 0 {
		b.WriteString("(none)\n")
		return b.String()
	}
	for _, task := range result.Tasks {
		auto := "manual"
		if task.AutoSafe {
			auto = "safe"
		}
		b.WriteString(fmt.Sprintf("### `%s`\n\n", task.ID))
		b.WriteString(fmt.Sprintf("- kind: `%s`\n", task.Kind))
		b.WriteString(fmt.Sprintf("- level: `%s`\n", task.Level))
		b.WriteString(fmt.Sprintf("- mode: `%s`\n", auto))
		b.WriteString(fmt.Sprintf("- path: `%s`\n", task.Path))
		b.WriteString(fmt.Sprintf("- summary: %s\n", task.Summary))
		b.WriteString(fmt.Sprintf("- recommendation: %s\n\n", task.Recommendation))
	}
	return b.String()
}

func frontmatter(content string) string {
	match := frontmatterBlockRE.FindStringSubmatch(content)
	if len(match) != 2 {
		return ""
	}
	return match[1]
}

func fmValue(frontmatter string, key string) string {
	for _, match := range fmLineRE.FindAllStringSubmatch(frontmatter, -1) {
		if len(match) == 3 && match[1] == key {
			return strings.Trim(strings.TrimSpace(match[2]), `"'`)
		}
	}
	return ""
}

func isKnowledgePage(path string) bool {
	return strings.HasPrefix(path, "wiki/sources/") ||
		strings.HasPrefix(path, "wiki/entities/") ||
		strings.HasPrefix(path, "wiki/concepts/") ||
		strings.HasPrefix(path, "wiki/topics/")
}

func titleFromSlug(slug string) string {
	parts := strings.Split(slug, "-")
	for i, part := range parts {
		if part == "" {
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, " ")
}
