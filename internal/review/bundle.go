package review

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/huic/nemo-knows/internal/wiki"
)

var (
	candidateWikiPathRE = regexp.MustCompile(`wiki/[A-Za-z0-9/_-]+\.md`)
	allowedWikiPathRE   = regexp.MustCompile(`^wiki/(sources|concepts|topics)/[a-z0-9][a-z0-9-]*\.md$`)
)

type candidateChange struct {
	Path      string
	Action    string
	Duplicate string
}

// ReviewBundle validates a local ingest draft bundle and renders an apply plan.
//
// The bundleDir argument points at a directory containing source.md and
// ingest-plan.md draft files.
//
// The returned Markdown is a review artifact for a human or agent. It does not
// apply changes to wiki/.
func ReviewBundle(bundleDir string) (string, error) {
	return ReviewBundleWithRoot(".", bundleDir)
}

// ReviewBundleWithRoot validates a bundle against a repository root.
//
// The root is used only to check whether candidate wiki pages already exist or
// look like duplicates of existing pages.
func ReviewBundleWithRoot(root string, bundleDir string) (string, error) {
	sourcePath := filepath.Join(bundleDir, "source.md")
	ingestPlanPath := filepath.Join(bundleDir, "ingest-plan.md")

	source, err := readDraft(sourcePath)
	if err != nil {
		return "", err
	}
	ingestPlan, err := readDraft(ingestPlanPath)
	if err != nil {
		return "", err
	}

	checks := make([]string, 0, 16)
	if err := requireFrontmatter("source.md", source, &checks); err != nil {
		return "", err
	}
	if err := requireFrontmatter("ingest-plan.md", ingestPlan, &checks); err != nil {
		return "", err
	}
	if err := requireFrontmatterValue("source.md", source, "kind", "source", &checks); err != nil {
		return "", err
	}
	if err := requireFrontmatterValue("ingest-plan.md", ingestPlan, "kind", "topic", &checks); err != nil {
		return "", err
	}
	for _, section := range []string{"What It Is", "Summary", "Key Claims", "Suggested Links"} {
		if err := requireSection("source.md", source, section, &checks); err != nil {
			return "", err
		}
	}
	for _, section := range []string{"Source Summary", "Candidate Wiki Pages", "Suggested Links", "Review Checklist"} {
		if err := requireSection("ingest-plan.md", ingestPlan, section, &checks); err != nil {
			return "", err
		}
	}

	candidates, err := extractCandidatePaths(ingestPlan)
	if err != nil {
		return "", err
	}
	changes := classifyCandidates(root, candidates)

	return renderApplyPlan(bundleDir, checks, changes), nil
}

func readDraft(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read draft %s: %w", path, err)
	}

	return string(content), nil
}

func requireFrontmatter(name string, markdown string, checks *[]string) error {
	ok, err := wiki.HasFrontmatter(markdown)
	if err != nil {
		return fmt.Errorf("validate %s frontmatter: %w", name, err)
	}
	if !ok {
		return fmt.Errorf("%s is missing YAML frontmatter", name)
	}

	*checks = append(*checks, fmt.Sprintf("- [x] `%s` has YAML frontmatter", name))
	return nil
}

func requireFrontmatterValue(name string, markdown string, key string, want string, checks *[]string) error {
	values := frontmatterValues(markdown)
	got := values[key]
	if got != want {
		return fmt.Errorf("%s frontmatter %q = %q, want %q", name, key, got, want)
	}

	*checks = append(*checks, fmt.Sprintf("- [x] `%s` frontmatter `%s` is `%s`", name, key, want))
	return nil
}

func frontmatterValues(markdown string) map[string]string {
	values := map[string]string{}
	lines := strings.Split(strings.TrimLeft(markdown, "\t\n\r "), "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return values
	}

	for _, line := range lines[1:] {
		trimmed := strings.TrimSpace(line)
		if trimmed == "---" {
			return values
		}
		key, value, ok := strings.Cut(trimmed, ":")
		if !ok {
			continue
		}
		values[strings.TrimSpace(key)] = strings.Trim(strings.TrimSpace(value), `"'`)
	}

	return values
}

func requireSection(name string, markdown string, section string, checks *[]string) error {
	heading := "## " + section
	if !strings.Contains(markdown, heading) {
		return fmt.Errorf("%s is missing required section %q", name, section)
	}

	*checks = append(*checks, fmt.Sprintf("- [x] `%s` includes required section `%s`", name, section))
	return nil
}

func extractCandidatePaths(markdown string) ([]string, error) {
	section, err := sectionBody(markdown, "Candidate Wiki Pages")
	if err != nil {
		return nil, err
	}
	matches := candidateWikiPathRE.FindAllString(section, -1)
	seen := map[string]bool{}
	paths := make([]string, 0, len(matches))
	for _, match := range matches {
		if !allowedWikiPathRE.MatchString(match) {
			return nil, fmt.Errorf("candidate wiki path %q must be under wiki/sources/, wiki/concepts/, or wiki/topics/ with lowercase hyphenated filename", match)
		}
		if seen[match] {
			continue
		}
		seen[match] = true
		paths = append(paths, match)
	}
	sort.Strings(paths)

	if len(paths) == 0 {
		return nil, fmt.Errorf("ingest-plan.md does not contain candidate wiki paths")
	}

	return paths, nil
}

func sectionBody(markdown string, section string) (string, error) {
	heading := "## " + section
	start := strings.Index(markdown, heading)
	if start == -1 {
		return "", fmt.Errorf("missing required section %q", section)
	}
	bodyStart := start + len(heading)
	rest := markdown[bodyStart:]
	if next := strings.Index(rest, "\n## "); next != -1 {
		rest = rest[:next]
	}

	return rest, nil
}

func classifyCandidates(root string, candidates []string) []candidateChange {
	existing := existingWikiPaths(root)
	changes := make([]candidateChange, 0, len(candidates))
	for _, candidate := range candidates {
		change := candidateChange{
			Path:   candidate,
			Action: "create",
		}
		if existing[candidate] {
			change.Action = "update"
		} else if duplicate := possibleDuplicate(candidate, existing); duplicate != "" {
			change.Duplicate = duplicate
		}
		changes = append(changes, change)
	}

	return changes
}

func existingWikiPaths(root string) map[string]bool {
	paths := map[string]bool{}
	for _, dir := range []string{"sources", "concepts", "topics"} {
		base := filepath.Join(root, "wiki", dir)
		entries, err := os.ReadDir(base)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
				continue
			}
			paths[filepath.ToSlash(filepath.Join("wiki", dir, entry.Name()))] = true
		}
	}

	return paths
}

func possibleDuplicate(candidate string, existing map[string]bool) string {
	candidateTokens := slugTokens(candidate)
	best := ""
	bestScore := 0
	for path := range existing {
		if filepath.Dir(path) != filepath.Dir(candidate) {
			continue
		}
		score := overlapCount(candidateTokens, slugTokens(path))
		if score > bestScore {
			best = path
			bestScore = score
		}
	}
	if bestScore < 2 {
		return ""
	}

	return best
}

func slugTokens(path string) map[string]bool {
	base := strings.TrimSuffix(filepath.Base(path), ".md")
	parts := strings.Split(base, "-")
	tokens := map[string]bool{}
	for _, part := range parts {
		if part == "" {
			continue
		}
		tokens[part] = true
	}

	return tokens
}

func overlapCount(a map[string]bool, b map[string]bool) int {
	count := 0
	for token := range a {
		if b[token] {
			count++
		}
	}

	return count
}

func renderApplyPlan(bundleDir string, checks []string, candidates []candidateChange) string {
	var b strings.Builder
	b.WriteString("# Reviewed Ingest Apply Plan\n\n")
	b.WriteString(fmt.Sprintf("Bundle: `%s`\n\n", bundleDir))
	b.WriteString("This is a review artifact. Do not apply this plan automatically.\n\n")
	b.WriteString("## Validation\n\n")
	for _, check := range checks {
		b.WriteString(check)
		b.WriteByte('\n')
	}
	b.WriteString("\n## Candidate Changes\n\n")
	for _, candidate := range candidates {
		switch {
		case candidate.Action == "update":
			b.WriteString(fmt.Sprintf("- `%s` — update existing page.\n", candidate.Path))
		case candidate.Duplicate != "":
			b.WriteString(fmt.Sprintf("- `%s` — create new page; possible duplicate of `%s`.\n", candidate.Path, candidate.Duplicate))
		default:
			b.WriteString(fmt.Sprintf("- `%s` — create new page.\n", candidate.Path))
		}
	}
	b.WriteString("\n## Required Manual Steps\n\n")
	b.WriteString("1. Compare each candidate page against the raw source and cleaned drafts.\n")
	b.WriteString("2. Create or update approved `wiki/sources/`, `wiki/concepts/`, and `wiki/topics/` pages.\n")
	b.WriteString("3. Update `wiki/index.md` after accepted page changes.\n")
	b.WriteString("4. Append an `ingest` entry to `wiki/log.md` after accepted page changes.\n")
	b.WriteString("5. Re-run wiki lint checks before committing.\n")

	return b.String()
}
