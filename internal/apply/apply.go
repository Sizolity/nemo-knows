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

	wikischema "github.com/huic/nemo-knows/internal/wiki"
)

var (
	ErrApprovalRequired = errors.New("apply requires explicit approval")
	ErrEvalNotPassing   = errors.New("bundle eval score is not passing")
	ErrAlreadyApplied   = errors.New("bundle has already been applied")
	// ErrSlugConflict is returned when applying a bundle would put two distinct
	// wiki pages under the same slug. Slugs must be unique across the whole wiki
	// because the filename slug is each page's identity for de-duplication, index
	// catalogue entries, and the renderer's defensive wikilink fallback.
	ErrSlugConflict = errors.New("wiki slug conflict: slugs must be unique across the whole wiki")

	candidateLineRE = regexp.MustCompile("(?m)^- `([^`]+)` — (.+)$")
	duplicateOfRE   = regexp.MustCompile("possible duplicate of `([^`]+)`")
	frontmatterRE   = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*`)
	titleRE         = regexp.MustCompile(`(?m)^title:\s*(.+?)\s*$`)
	kindRE          = regexp.MustCompile(`(?m)^kind:\s*(.+?)\s*$`)
	headingRE       = regexp.MustCompile(`(?m)^#\s+(.+?)\s*$`)
	sourceRefRE     = regexp.MustCompile(`(?m)^\s*-\s*((?:pipeline/)?raw/[^ \n]+|wiki/sources/[^ \n]+)\s*$`)
	inlineSourceRE  = regexp.MustCompile(`(?:(?:pipeline/)?raw|wiki/sources)/[A-Za-z0-9._/-]*[A-Za-z0-9_-]\.md`)
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
	if err := checkSlugConflicts(root, writes); err != nil {
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
		cleanTarget, ok := cleanKnowledgePath(target)
		if !ok {
			skipped = append(skipped, fmt.Sprintf("%s — unsupported candidate target", candidate.path))
			continue
		}
		target = cleanTarget

		switch {
		case strings.HasPrefix(target, "wiki/sources/"):
			created := !wikiFileExists(root, target)
			needsIndex = needsIndex || created
			writes = append(writes, plannedWrite{target: target, draft: sourceDraft, created: created})
		case isCandidateDraftTarget(target):
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
	cleanPath, ok := cleanKnowledgePath(repoPath)
	if !ok {
		return fmt.Errorf("refuse to write outside wiki knowledge paths: %s", repoPath)
	}
	repoPath = cleanPath

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
	case strings.HasPrefix(target, "wiki/entities/"):
		wantKind = "entity"
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

func isCandidateDraftTarget(target string) bool {
	return strings.HasPrefix(target, "wiki/entities/") ||
		strings.HasPrefix(target, "wiki/concepts/") ||
		strings.HasPrefix(target, "wiki/topics/")
}

func cleanKnowledgePath(repoPath string) (string, bool) {
	if filepath.IsAbs(repoPath) {
		return "", false
	}
	clean := filepath.ToSlash(filepath.Clean(filepath.FromSlash(repoPath)))
	if clean != filepath.ToSlash(repoPath) {
		return "", false
	}
	if filepath.Ext(clean) != ".md" {
		return "", false
	}
	if strings.HasPrefix(clean, "wiki/sources/") ||
		strings.HasPrefix(clean, "wiki/entities/") ||
		strings.HasPrefix(clean, "wiki/concepts/") ||
		strings.HasPrefix(clean, "wiki/topics/") {
		return clean, true
	}
	return "", false
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

	// Catalog index entries by section + relative path, not by slug. A page is
	// already listed only when an entry points at the same wiki file. Pages that
	// merely share a slug across categories (for example a source and an entity
	// both named "sqlite") must each still appear in their own section, so a
	// slug-only check would silently drop the second page from the index.
	wantRel := wikischema.IndexRelativePath(target)
	for _, line := range strings.Split(string(content), "\n") {
		if existing, ok := wikischema.IndexEntryTarget(line); ok && wikischema.IndexRelativePath(existing) == wantRel {
			return false, nil
		}
	}

	title := candidateTitle(draft)
	if title == "" {
		return false, nil
	}
	entry := wikischema.FormatIndexEntry(target, title) + "\n"
	updated := wikischema.AppendIndexEntry(string(content), indexSection(target), entry)
	if err := os.WriteFile(indexPath, []byte(updated), 0o644); err != nil {
		return false, fmt.Errorf("write wiki index: %w", err)
	}

	return true, nil
}

// checkSlugConflicts fails the apply when it would create two distinct wiki
// pages that share a slug. It compares newly created pages against existing
// wiki pages at a different path and against one another, so a source page and
// an entity page can no longer both claim the slug "sqlite".
func checkSlugConflicts(root string, writes []plannedWrite) error {
	bySlug := existingWikiSlugPaths(root)
	for _, item := range writes {
		if !item.created {
			continue
		}
		slug := wikischema.SlugFromWikiPath(item.target)
		if other, ok := bySlug[slug]; ok && other != item.target {
			return fmt.Errorf("%w: slug %q is used by both %s and %s; rename one page so every slug is unique", ErrSlugConflict, slug, other, item.target)
		}
		bySlug[slug] = item.target
	}
	return nil
}

// existingWikiSlugPaths maps each maintained wiki page slug to its repo-relative
// path so callers can detect cross-category slug collisions.
func existingWikiSlugPaths(root string) map[string]string {
	slugs := map[string]string{}
	_ = filepath.WalkDir(filepath.Join(root, "wiki"), func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		rel, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return nil
		}
		repoPath := filepath.ToSlash(rel)
		if _, ok := cleanKnowledgePath(repoPath); !ok {
			return nil
		}
		slugs[wikischema.SlugFromWikiPath(repoPath)] = repoPath
		return nil
	})
	return slugs
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
	if strings.HasPrefix(target, "wiki/entities/") {
		return "## Entities"
	}
	if strings.HasPrefix(target, "wiki/topics/") {
		return "## Topics"
	}

	return "## Concepts"
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
	entry.WriteString("Source: ")
	entry.WriteString(firstSourceReference(sourceDraft))
	entry.WriteByte('\n')
	entry.WriteString("Applied bundle: ")
	entry.WriteString(displayBundle(root, bundleDir))
	entry.WriteByte('\n')
	entry.WriteString("Touched:\n")
	for _, touched := range result.Touched {
		entry.WriteString(fmt.Sprintf("- %s (%s)\n", touched.Path, touched.Action))
	}
	for _, skipped := range result.Skipped {
		entry.WriteString("- skipped: ")
		entry.WriteString(skipped)
		entry.WriteByte('\n')
	}
	entry.WriteString("Open: review skipped candidates before creating entity, concept, or topic pages.\n")

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
		b.WriteString("- ")
		b.WriteString(written)
		b.WriteByte('\n')
	}
	b.WriteString("\n## Skipped\n\n")
	if len(result.Skipped) == 0 {
		b.WriteString("(none)\n")
	}
	for _, skipped := range result.Skipped {
		b.WriteString("- ")
		b.WriteString(skipped)
		b.WriteByte('\n')
	}

	path := filepath.Join(bundleDir, "apply-report.md")
	if err := os.WriteFile(path, []byte(b.String()), 0o644); err != nil {
		return fmt.Errorf("write apply report: %w", err)
	}

	return nil
}
