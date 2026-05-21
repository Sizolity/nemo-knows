package apply

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	ErrApprovalRequired = errors.New("apply requires explicit approval")
	ErrEvalNotPassing   = errors.New("bundle eval score is not passing")
	ErrAlreadyApplied   = errors.New("bundle has already been applied")

	candidateLineRE = regexp.MustCompile("(?m)^- `([^`]+)` — (.+)$")
	duplicateOfRE   = regexp.MustCompile("possible duplicate of `([^`]+)`")
	frontmatterRE   = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*`)
	titleRE         = regexp.MustCompile(`(?m)^title:\s*(.+?)\s*$`)
	kindRE          = regexp.MustCompile(`(?m)^kind:\s*(.+?)\s*$`)
	headingRE       = regexp.MustCompile(`(?m)^#\s+(.+?)\s*$`)
	sourceRefRE     = regexp.MustCompile(`(?m)^\s*-\s*(raw/[^ \n]+|wiki/sources/[^ \n]+)\s*$`)
	inlineSourceRE  = regexp.MustCompile(`(?:raw|wiki/sources)/[A-Za-z0-9._/-]*[A-Za-z0-9_-]\.md`)
)

type Options struct {
	Approve bool
	Force   bool
}

type Result struct {
	Written []string
	Skipped []string
	Touched []Touched
}

type Touched struct {
	Path   string
	Action string
}

type scoresFile struct {
	Scores struct {
		Overall string `json:"overall"`
	} `json:"scores"`
}

// ApplyApproved applies a reviewed and evaluated bundle to wiki/.
//
// The function is intentionally conservative: it requires explicit approval,
// refuses failing evals, redirects duplicate source candidates to existing
// source pages, and records skipped candidates in an apply report.
func ApplyApproved(root string, bundleDir string, opts Options) (Result, error) {
	if !opts.Approve {
		return Result{}, ErrApprovalRequired
	}
	if err := requirePassingEval(bundleDir); err != nil {
		return Result{}, err
	}
	if !opts.Force {
		if err := requireNotAlreadyApplied(root, bundleDir); err != nil {
			return Result{}, err
		}
	}

	applyPlan, err := os.ReadFile(filepath.Join(bundleDir, "apply-plan.md"))
	if err != nil {
		return Result{}, fmt.Errorf("read apply plan: %w", err)
	}
	sourceDraft, err := os.ReadFile(filepath.Join(bundleDir, "source.md"))
	if err != nil {
		return Result{}, fmt.Errorf("read source draft: %w", err)
	}

	writes, skipped, err := planApprovedWrites(root, bundleDir, string(applyPlan), sourceDraft)
	if err != nil {
		return Result{}, err
	}

	result := Result{Skipped: skipped}
	for _, item := range writes {
		if err := writeWikiFile(root, item.target, item.draft); err != nil {
			return Result{}, err
		}
		result.Written = append(result.Written, item.target)
		result.Touched = append(result.Touched, Touched{Path: item.target, Action: createOrUpdate(item.created)})
		if item.created {
			written, err := updateIndexForCandidate(root, item.target, item.draft)
			if err != nil {
				return Result{}, err
			}
			if written && !hasWritten(result.Written, "wiki/index.md") {
				result.Written = append(result.Written, "wiki/index.md")
				result.Touched = append(result.Touched, Touched{Path: "wiki/index.md", Action: "updated"})
			}
		}
	}

	if len(result.Written) > 0 {
		if err := appendApplyLog(root, bundleDir, result); err != nil {
			return Result{}, err
		}
		result.Written = append(result.Written, "wiki/log.md")
	}
	if err := writeApplyReport(bundleDir, result); err != nil {
		return Result{}, err
	}

	return result, nil
}

type plannedWrite struct {
	target  string
	draft   []byte
	created bool
}

func planApprovedWrites(root string, bundleDir string, applyPlan string, sourceDraft []byte) ([]plannedWrite, []string, error) {
	writes := []plannedWrite{}
	skipped := []string{}
	needsIndex := false
	for _, candidate := range parseCandidates(applyPlan) {
		target := candidate.path
		if duplicate := candidate.duplicate; duplicate != "" {
			target = duplicate
		}

		switch {
		case strings.HasPrefix(target, "wiki/sources/"):
			created := !wikiFileExists(root, target)
			needsIndex = needsIndex || created
			writes = append(writes, plannedWrite{target: target, draft: sourceDraft, created: created})
		case strings.HasPrefix(target, "wiki/concepts/") || strings.HasPrefix(target, "wiki/topics/"):
			draft, existed, err := readCandidateDraft(bundleDir, target)
			if err != nil {
				return nil, nil, err
			}
			if !existed {
				skipped = append(skipped, fmt.Sprintf("%s — missing reviewed candidate draft", candidate.path))
				continue
			}
			if err := validateCandidateDraft(target, draft); err != nil {
				return nil, nil, err
			}
			created := !wikiFileExists(root, target)
			needsIndex = needsIndex || created
			writes = append(writes, plannedWrite{target: target, draft: draft, created: created})
		default:
			skipped = append(skipped, fmt.Sprintf("%s — unsupported candidate target", candidate.path))
		}
	}
	if needsIndex {
		if _, err := os.ReadFile(filepath.Join(root, "wiki", "index.md")); err != nil {
			return nil, nil, fmt.Errorf("preflight wiki index: %w", err)
		}
	}
	return writes, skipped, nil
}

func createOrUpdate(created bool) string {
	if created {
		return "created"
	}
	return "updated"
}

func requireNotAlreadyApplied(root string, bundleDir string) error {
	content, err := os.ReadFile(filepath.Join(root, "wiki", "log.md"))
	if err != nil {
		return fmt.Errorf("read wiki log: %w", err)
	}
	subject := displayBundle(root, bundleDir)
	needle := " ingest | " + subject
	appliedNeedle := "Applied bundle: " + subject
	if strings.Contains(string(content), needle) || strings.Contains(string(content), appliedNeedle) {
		return fmt.Errorf("%w: %s", ErrAlreadyApplied, subject)
	}

	return nil
}

func requirePassingEval(bundleDir string) error {
	content, err := os.ReadFile(filepath.Join(bundleDir, "scores.json"))
	if err != nil {
		return fmt.Errorf("read scores: %w", err)
	}

	var scores scoresFile
	if err := json.Unmarshal(content, &scores); err != nil {
		return fmt.Errorf("parse scores: %w", err)
	}
	if scores.Scores.Overall != "pass" {
		return fmt.Errorf("%w: overall=%s", ErrEvalNotPassing, scores.Scores.Overall)
	}

	return nil
}

type candidate struct {
	path      string
	duplicate string
}

func parseCandidates(applyPlan string) []candidate {
	matches := candidateLineRE.FindAllStringSubmatch(applyPlan, -1)
	candidates := make([]candidate, 0, len(matches))
	for _, match := range matches {
		item := candidate{path: match[1]}
		if duplicate := duplicateOfRE.FindStringSubmatch(match[2]); len(duplicate) == 2 {
			item.duplicate = duplicate[1]
		}
		candidates = append(candidates, item)
	}

	return candidates
}

func writeWikiFile(root string, repoPath string, content []byte) error {
	if !strings.HasPrefix(repoPath, "wiki/sources/") &&
		!strings.HasPrefix(repoPath, "wiki/concepts/") &&
		!strings.HasPrefix(repoPath, "wiki/topics/") {
		return fmt.Errorf("refuse to write outside wiki knowledge paths: %s", repoPath)
	}

	path := filepath.Join(root, filepath.FromSlash(repoPath))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create wiki directory: %w", err)
	}
	if err := os.WriteFile(path, content, 0o644); err != nil {
		return fmt.Errorf("write wiki file %s: %w", repoPath, err)
	}

	return nil
}

func readCandidateDraft(bundleDir string, target string) ([]byte, bool, error) {
	path := filepath.Join(bundleDir, "candidates", filepath.FromSlash(target))
	content, err := os.ReadFile(path)
	if err == nil {
		return content, true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return nil, false, nil
	}

	return nil, false, fmt.Errorf("read candidate draft %s: %w", target, err)
}

func validateCandidateDraft(target string, content []byte) error {
	frontmatter, err := candidateFrontmatter(content)
	if err != nil {
		return fmt.Errorf("validate candidate draft %s: %w", target, err)
	}

	wantKind := ""
	switch {
	case strings.HasPrefix(target, "wiki/concepts/"):
		wantKind = "concept"
	case strings.HasPrefix(target, "wiki/topics/"):
		wantKind = "topic"
	default:
		return fmt.Errorf("unsupported candidate target: %s", target)
	}
	if got := frontmatterValue(frontmatter, kindRE); got != wantKind {
		return fmt.Errorf("candidate draft %s has kind %q, want %q", target, got, wantKind)
	}
	if !strings.Contains(frontmatter, "sources:") {
		return fmt.Errorf("candidate draft %s missing sources frontmatter", target)
	}

	return nil
}

func candidateFrontmatter(content []byte) (string, error) {
	match := frontmatterRE.FindSubmatch(content)
	if len(match) != 2 {
		return "", errors.New("missing YAML frontmatter")
	}

	return string(match[1]), nil
}

func frontmatterValue(frontmatter string, re *regexp.Regexp) string {
	match := re.FindStringSubmatch(frontmatter)
	if len(match) != 2 {
		return ""
	}

	return strings.Trim(strings.TrimSpace(match[1]), `"'`)
}

func wikiFileExists(root string, repoPath string) bool {
	_, err := os.Stat(filepath.Join(root, filepath.FromSlash(repoPath)))
	return err == nil
}

func hasWritten(written []string, target string) bool {
	for _, item := range written {
		if item == target {
			return true
		}
	}

	return false
}

func updateIndexForCandidate(root string, target string, draft []byte) (bool, error) {
	indexPath := filepath.Join(root, "wiki", "index.md")
	content, err := os.ReadFile(indexPath)
	if err != nil {
		return false, fmt.Errorf("read wiki index: %w", err)
	}

	slug := strings.TrimSuffix(filepath.Base(target), filepath.Ext(target))
	link := "[[" + slug + "]]"
	if strings.Contains(string(content), link) {
		return false, nil
	}

	title := candidateTitle(draft)
	if title == "" {
		return false, nil
	}
	entry := fmt.Sprintf("- %s — %s.\n", link, title)
	updated := appendIndexEntry(string(content), indexSection(target), entry)
	if err := os.WriteFile(indexPath, []byte(updated), 0o644); err != nil {
		return false, fmt.Errorf("write wiki index: %w", err)
	}

	return true, nil
}

func candidateTitle(draft []byte) string {
	if frontmatter, err := candidateFrontmatter(draft); err == nil {
		if title := frontmatterValue(frontmatter, titleRE); title != "" {
			return title
		}
	}
	match := headingRE.FindSubmatch(draft)
	if len(match) == 2 {
		return strings.TrimSpace(string(match[1]))
	}

	return ""
}

func indexSection(target string) string {
	if strings.HasPrefix(target, "wiki/sources/") {
		return "## Sources"
	}
	if strings.HasPrefix(target, "wiki/topics/") {
		return "## Topics"
	}

	return "## Concepts"
}

func appendIndexEntry(index string, section string, entry string) string {
	if !strings.Contains(index, section) {
		if !strings.HasSuffix(index, "\n") {
			index += "\n"
		}
		return index + "\n" + section + "\n" + entry
	}

	lines := strings.Split(index, "\n")
	sectionLine := -1
	for i, line := range lines {
		if strings.TrimSpace(line) == section {
			sectionLine = i
			break
		}
	}
	if sectionLine == -1 {
		return index
	}

	end := len(lines)
	for i := sectionLine + 1; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "## ") {
			end = i
			break
		}
	}

	insertAt := end
	for insertAt > sectionLine+1 && strings.TrimSpace(lines[insertAt-1]) == "" {
		insertAt--
	}

	entryLine := strings.TrimSuffix(entry, "\n")
	updated := make([]string, 0, len(lines)+2)
	updated = append(updated, lines[:insertAt]...)
	updated = append(updated, entryLine)
	if end < len(lines) {
		updated = append(updated, "")
	}
	updated = append(updated, lines[insertAt:]...)

	return strings.Join(updated, "\n")
}

func appendApplyLog(root string, bundleDir string, result Result) error {
	path := filepath.Join(root, "wiki", "log.md")
	date := time.Now().Format("2006-01-02")
	subject := displayBundle(root, bundleDir)
	sourceDraft, _ := os.ReadFile(filepath.Join(bundleDir, "source.md"))
	if title := candidateTitle(sourceDraft); title != "" {
		subject = title
	}
	entry := strings.Builder{}
	entry.WriteString(fmt.Sprintf("\n## [%s] ingest | %s\n", date, subject))
	entry.WriteString("Source: " + firstSourceReference(sourceDraft) + "\n")
	entry.WriteString("Applied bundle: " + displayBundle(root, bundleDir) + "\n")
	entry.WriteString("Touched:\n")
	for _, touched := range result.Touched {
		entry.WriteString(fmt.Sprintf("- %s (%s)\n", touched.Path, touched.Action))
	}
	for _, skipped := range result.Skipped {
		entry.WriteString("- skipped: " + skipped + "\n")
	}
	entry.WriteString("Open: review skipped candidates before creating concept or topic pages.\n")

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open wiki log: %w", err)
	}
	defer file.Close()
	if _, err := file.WriteString(entry.String()); err != nil {
		return fmt.Errorf("append wiki log: %w", err)
	}

	return nil
}

func firstSourceReference(sourceDraft []byte) string {
	if match := sourceRefRE.FindSubmatch(sourceDraft); len(match) == 2 {
		return string(match[1])
	}
	if match := inlineSourceRE.Find(sourceDraft); len(match) > 0 {
		return string(match)
	}
	return "unknown"
}

func displayBundle(root string, bundleDir string) string {
	rel, err := filepath.Rel(root, bundleDir)
	if err != nil || strings.HasPrefix(rel, "..") {
		return filepath.ToSlash(bundleDir)
	}
	return filepath.ToSlash(rel)
}

func writeApplyReport(bundleDir string, result Result) error {
	var b strings.Builder
	b.WriteString("# Approved Apply Report\n\n")
	b.WriteString("## Written\n\n")
	if len(result.Written) == 0 {
		b.WriteString("(none)\n")
	}
	for _, written := range result.Written {
		b.WriteString("- " + written + "\n")
	}
	b.WriteString("\n## Skipped\n\n")
	if len(result.Skipped) == 0 {
		b.WriteString("(none)\n")
	}
	for _, skipped := range result.Skipped {
		b.WriteString("- " + skipped + "\n")
	}

	path := filepath.Join(bundleDir, "apply-report.md")
	if err := os.WriteFile(path, []byte(b.String()), 0o644); err != nil {
		return fmt.Errorf("write apply report: %w", err)
	}

	return nil
}
