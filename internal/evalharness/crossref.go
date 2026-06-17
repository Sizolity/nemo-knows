package evalharness

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	wikischema "github.com/huic/nemo-knows/internal/wiki"
)

type CrosslinkResult struct {
	Bundle string           `json:"bundle"`
	Issues []CrosslinkIssue `json:"issues"`
	Graph  []CrosslinkEdge  `json:"graph"`
}

type CrosslinkIssue struct {
	Path    string `json:"path"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CrosslinkEdge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// EvaluateBundleCrosslinks checks the Markdown relative links among reviewed
// candidate drafts and against the existing wiki. It is read-only and meant to
// supplement candidate eval. Targets are reachable when they resolve to a
// sibling candidate or an existing wiki page on disk; residual Obsidian
// [[wikilink]] syntax is flagged so the bundle gate matches the lint rules.
func EvaluateBundleCrosslinks(root string, bundleDir string) (CrosslinkResult, error) {
	applyPlan, err := os.ReadFile(filepath.Join(bundleDir, "apply-plan.md"))
	if err != nil {
		return CrosslinkResult{}, fmt.Errorf("read apply plan: %w", err)
	}
	result := CrosslinkResult{Bundle: bundleDir}
	detectSlugConflicts(root, candidatePaths(string(applyPlan)), &result)
	detectSourceLinkSafety(bundleDir, &result)
	targets := candidateDraftPaths(string(applyPlan))
	candidateTargets := map[string]bool{}
	for _, target := range targets {
		candidateTargets[target] = true
	}
	inbound := map[string]int{}
	for _, target := range targets {
		path := filepath.Join(bundleDir, "candidates", filepath.FromSlash(target))
		content, err := os.ReadFile(path)
		if err != nil {
			result.Issues = append(result.Issues, CrosslinkIssue{Path: target, Code: "missing-candidate", Message: "candidate draft is missing"})
			continue
		}
		if refs := wikischema.WikilinkReferences(string(content)); len(refs) > 0 {
			result.Issues = append(result.Issues, CrosslinkIssue{Path: target, Code: "forbidden-wikilink", Message: "Obsidian [[wikilink]] syntax is no longer allowed; use Markdown relative links: " + strings.Join(refs, ", ")})
		}
		for _, link := range wikischema.UnsafeLinkTargets(target, string(content)) {
			result.Issues = append(result.Issues, CrosslinkIssue{Path: target, Code: "unsafe-link-target", Message: unsafeLinkMessage(link)})
		}
		for _, href := range wikischema.MarkdownLinkTargets(string(content)) {
			repoPath, ok := wikischema.RepoPathFromHref(target, href)
			if !ok || !strings.HasPrefix(repoPath, "wiki/") {
				continue
			}
			if candidateTargets[repoPath] {
				inbound[repoPath]++
				result.Graph = append(result.Graph, CrosslinkEdge{From: target, To: repoPath})
				continue
			}
			if _, statErr := os.Stat(filepath.Join(root, filepath.FromSlash(repoPath))); statErr == nil {
				result.Graph = append(result.Graph, CrosslinkEdge{From: target, To: repoPath})
				continue
			}
			result.Issues = append(result.Issues, CrosslinkIssue{Path: target, Code: "missing-target", Message: "relative link target does not exist in reviewed candidates or wiki: " + href})
		}
	}
	for _, target := range targets {
		if inbound[target] == 0 {
			result.Issues = append(result.Issues, CrosslinkIssue{Path: target, Code: "zero-inbound", Message: "candidate has no inbound links from sibling candidates"})
		}
	}
	sort.Slice(result.Issues, func(i int, j int) bool {
		if result.Issues[i].Path == result.Issues[j].Path {
			return result.Issues[i].Code < result.Issues[j].Code
		}
		return result.Issues[i].Path < result.Issues[j].Path
	})
	sort.Slice(result.Graph, func(i int, j int) bool {
		if result.Graph[i].From == result.Graph[j].From {
			return result.Graph[i].To < result.Graph[j].To
		}
		return result.Graph[i].From < result.Graph[j].From
	})
	return result, nil
}

// existingWikiSlugPaths maps each maintained wiki page slug to its repo-relative
// path so the harness can flag cross-category slug collisions.
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
		if !isWikiKnowledgePath(repoPath) {
			return nil
		}
		slugs[strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))] = repoPath
		return nil
	})
	return slugs
}

func isWikiKnowledgePath(repoPath string) bool {
	return strings.HasPrefix(repoPath, "wiki/sources/") ||
		strings.HasPrefix(repoPath, "wiki/entities/") ||
		strings.HasPrefix(repoPath, "wiki/concepts/") ||
		strings.HasPrefix(repoPath, "wiki/topics/")
}

// detectSlugConflicts reports cross-category slug collisions among the bundle's
// candidate pages and against the existing wiki. Slugs must remain unique across
// the whole wiki because they are the stable page identifiers (filenames); a
// collision is a structural problem even when every individual relative link
// resolves.
func detectSlugConflicts(root string, candidates []string, result *CrosslinkResult) {
	existing := existingWikiSlugPaths(root)
	seen := map[string]string{}
	for _, target := range candidates {
		if !isWikiKnowledgePath(target) {
			continue
		}
		slug := strings.TrimSuffix(filepath.Base(target), filepath.Ext(target))
		if other, ok := existing[slug]; ok && other != target {
			result.Issues = append(result.Issues, CrosslinkIssue{
				Path:    target,
				Code:    "slug-conflict",
				Message: "slug already used by existing wiki page at a different path: " + other,
			})
		}
		if other, ok := seen[slug]; ok && other != target {
			result.Issues = append(result.Issues, CrosslinkIssue{
				Path:    target,
				Code:    "slug-conflict",
				Message: "slug also used by a sibling candidate at a different path: " + other,
			})
		}
		seen[slug] = target
	}
}

// detectSourceLinkSafety flags forbidden cross-reference syntax and unsafe link
// targets in the bundle's source draft body. The crosslink pass otherwise only
// inspects generated candidate drafts, but the source page is applied to wiki/
// too and must satisfy the same link-safety invariants — this is the eval-side
// gate for the source-page body where the original Suggested Links defect lived.
func detectSourceLinkSafety(bundleDir string, result *CrosslinkResult) {
	content, err := os.ReadFile(filepath.Join(bundleDir, "source.md"))
	if err != nil {
		return
	}
	// Source pages live at wiki/sources/<slug>.md, so anchor relative-target
	// resolution at that depth. The reported path stays "source.md" because the
	// final slug is only resolved from the apply plan at write time.
	const fromPath = "wiki/sources/source.md"
	const display = "source.md"
	if refs := wikischema.WikilinkReferences(string(content)); len(refs) > 0 {
		result.Issues = append(result.Issues, CrosslinkIssue{Path: display, Code: "forbidden-wikilink", Message: "Obsidian [[wikilink]] syntax is no longer allowed; use Markdown relative links: " + strings.Join(refs, ", ")})
	}
	for _, link := range wikischema.UnsafeLinkTargets(fromPath, string(content)) {
		result.Issues = append(result.Issues, CrosslinkIssue{Path: display, Code: "unsafe-link-target", Message: unsafeLinkMessage(link)})
	}
}

// unsafeLinkMessage renders a human-readable explanation for an unsafe link
// finding so eval issues read the same way the lint unsafe-link-target rule does.
func unsafeLinkMessage(link wikischema.UnsafeLink) string {
	if link.Reason == wikischema.ReasonDangerousScheme {
		return "link uses a dangerous URI scheme (" + link.Scheme + "); only http/https/mailto are allowed: " + link.Target
	}
	return "relative link target escapes the repository root: " + link.Target
}
