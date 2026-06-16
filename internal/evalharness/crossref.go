package evalharness

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
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

// EvaluateBundleCrosslinks checks wikilinks among reviewed candidate drafts and
// the existing wiki. It is read-only and meant to supplement candidate eval.
func EvaluateBundleCrosslinks(root string, bundleDir string) (CrosslinkResult, error) {
	applyPlan, err := os.ReadFile(filepath.Join(bundleDir, "apply-plan.md"))
	if err != nil {
		return CrosslinkResult{}, fmt.Errorf("read apply plan: %w", err)
	}
	result := CrosslinkResult{Bundle: bundleDir}
	detectSlugConflicts(root, candidatePaths(string(applyPlan)), &result)
	targets := candidateDraftPaths(string(applyPlan))
	candidateSlugs := map[string]string{}
	for _, target := range targets {
		candidateSlugs[strings.TrimSuffix(filepath.Base(target), filepath.Ext(target))] = target
	}
	wikiSlugs := existingWikiSlugs(root)
	inbound := map[string]int{}
	for _, target := range targets {
		path := filepath.Join(bundleDir, "candidates", filepath.FromSlash(target))
		content, err := os.ReadFile(path)
		if err != nil {
			result.Issues = append(result.Issues, CrosslinkIssue{Path: target, Code: "missing-candidate", Message: "candidate draft is missing"})
			continue
		}
		for _, match := range wikilinkRE.FindAllStringSubmatch(string(content), -1) {
			if len(match) != 2 {
				continue
			}
			slug := slugFromCandidateLink(match[1])
			if candidateTarget, ok := candidateSlugs[slug]; ok {
				inbound[slug]++
				result.Graph = append(result.Graph, CrosslinkEdge{From: target, To: candidateTarget})
				continue
			}
			if wikiSlugs[slug] {
				result.Graph = append(result.Graph, CrosslinkEdge{From: target, To: "wiki/" + slug})
				continue
			}
			result.Issues = append(result.Issues, CrosslinkIssue{Path: target, Code: "missing-target", Message: "wikilink target does not exist in reviewed candidates or wiki: " + match[1]})
		}
	}
	for slug, target := range candidateSlugs {
		if inbound[slug] == 0 {
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

func existingWikiSlugs(root string) map[string]bool {
	slugs := map[string]bool{}
	_ = filepath.WalkDir(filepath.Join(root, "wiki"), func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		slugs[strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))] = true
		return nil
	})
	return slugs
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
// candidate pages and against the existing wiki. Slugs must be unique across the
// whole wiki so [[slug]] wikilinks resolve to a single page; a collision is a
// structural problem even when every individual wikilink target exists.
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
