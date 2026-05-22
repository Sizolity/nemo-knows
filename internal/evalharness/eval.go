package evalharness

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	candidateLineRE        = regexp.MustCompile("(?m)^- `([^`]+)`")
	sourceRefRE            = regexp.MustCompile(`(?m)^\s*-\s+(raw/[^\s]+)\s*$`)
	truncationMarkerRE     = regexp.MustCompile(`(?i)\[truncated at [^\]]+\]`)
	completenessClaimRE    = regexp.MustCompile(`(?i)\b(complete|entire|unabridged)\s+(work|source|text|novel|retrieval|document|file)\b|\b(full text|final chapters|without abridgment|all \d+ chapters|all chapters)\b`)
	truncationMentionRE    = regexp.MustCompile(`(?i)\b(truncated|truncation|incomplete|partial|excerpt|ends? (at|around|mid)|through (the )?(opening|beginning) of)\b`)
	completenessNegationRE = regexp.MustCompile(`(?i)\b(not|no|does not|do not|don't|without|absent|missing|forbid|forbidden|must not|should not)\b`)
)

type Result struct {
	Bundle string   `json:"bundle"`
	Scores Scores   `json:"scores"`
	Trace  []string `json:"trace"`
}

type Scores struct {
	Schema             string `json:"schema"`
	WikiSafety         string `json:"wiki_safety"`
	CandidatePaths     string `json:"candidate_paths"`
	DuplicateDetection string `json:"duplicate_detection"`
	SourceCompleteness string `json:"source_completeness"`
	ApplyPlanCoverage  string `json:"apply_plan_coverage"`
	ApplyReadiness     string `json:"apply_readiness"`
	Overall            string `json:"overall"`
}

// EvaluateBundle scores reviewed ingest artifacts without calling the model.
//
// The bundleDir directory must contain apply-plan.md. The score is intentionally
// coarse so it can act as a stable regression signal.
func EvaluateBundle(bundleDir string) (Result, error) {
	applyPlanPath := filepath.Join(bundleDir, "apply-plan.md")
	content, err := os.ReadFile(applyPlanPath)
	if err != nil {
		return Result{}, fmt.Errorf("read apply plan: %w", err)
	}

	applyPlan := string(content)
	result := Result{
		Bundle: bundleDir,
		Trace:  []string{},
	}
	result.Scores.Schema = scoreSchema(applyPlan, &result.Trace)
	result.Scores.WikiSafety = scoreWikiSafety(applyPlan, &result.Trace)
	result.Scores.CandidatePaths = scoreCandidatePaths(applyPlan, &result.Trace)
	result.Scores.DuplicateDetection = scoreDuplicateDetection(applyPlan, &result.Trace)
	result.Scores.SourceCompleteness = scoreSourceCompleteness(bundleDir, &result.Trace)
	result.Scores.ApplyPlanCoverage = scoreApplyPlanCoverage(bundleDir, applyPlan, &result.Trace)
	result.Scores.ApplyReadiness = scoreApplyReadiness(result.Scores)
	result.Scores.Overall = scoreOverall(result.Scores)

	return result, nil
}

func scoreSchema(applyPlan string, trace *[]string) string {
	required := []string{
		"`source.md` has YAML frontmatter",
		"`ingest-plan.md` has YAML frontmatter",
		"`source.md` frontmatter `kind` is `source`",
		"`ingest-plan.md` frontmatter `kind` is `topic`",
		"`source.md` includes required section `What It Is`",
		"`source.md` includes required section `Summary`",
		"`source.md` includes required section `Key Claims`",
		"`source.md` includes required section `Suggested Links`",
		"`ingest-plan.md` includes required section `Source Summary`",
		"`ingest-plan.md` includes required section `Candidate Wiki Pages`",
		"`ingest-plan.md` includes required section `Suggested Links`",
		"`ingest-plan.md` includes required section `Review Checklist`",
	}
	for _, needle := range required {
		if !strings.Contains(applyPlan, needle) {
			*trace = append(*trace, "schema: missing validation check "+needle)
			return "fail"
		}
	}
	*trace = append(*trace, "schema: all required review checks present")
	return "pass"
}

func scoreWikiSafety(applyPlan string, trace *[]string) string {
	if !strings.Contains(applyPlan, "Do not apply this plan automatically.") {
		*trace = append(*trace, "wiki_safety: missing no-auto-apply warning")
		return "fail"
	}
	*trace = append(*trace, "wiki_safety: no-auto-apply warning present")
	return "pass"
}

func scoreCandidatePaths(applyPlan string, trace *[]string) string {
	candidates := candidatePaths(applyPlan)
	if len(candidates) == 0 {
		*trace = append(*trace, "candidate_paths: no candidate paths found")
		return "fail"
	}
	for _, candidate := range candidates {
		if !strings.HasPrefix(candidate, "wiki/sources/") &&
			!strings.HasPrefix(candidate, "wiki/concepts/") &&
			!strings.HasPrefix(candidate, "wiki/topics/") {
			*trace = append(*trace, "candidate_paths: invalid candidate "+candidate)
			return "fail"
		}
	}
	*trace = append(*trace, fmt.Sprintf("candidate_paths: %d legal candidates", len(candidates)))
	return "pass"
}

func scoreDuplicateDetection(applyPlan string, trace *[]string) string {
	candidateLines := candidateLines(applyPlan)
	for _, line := range candidateLines {
		if strings.Contains(line, "possible duplicate") ||
			strings.Contains(line, "update existing page") {
			continue
		}
		if strings.Contains(line, "llm-wiki-pattern") {
			*trace = append(*trace, "duplicate_detection: likely duplicate lacks duplicate label")
			return "borderline"
		}
	}
	*trace = append(*trace, "duplicate_detection: duplicates are labeled or no obvious duplicate")
	return "pass"
}

func scoreApplyReadiness(scores Scores) string {
	if scores.Schema == "fail" || scores.WikiSafety == "fail" || scores.CandidatePaths == "fail" || scores.SourceCompleteness == "fail" || scores.ApplyPlanCoverage == "fail" {
		return "fail"
	}
	if scores.DuplicateDetection == "borderline" || scores.SourceCompleteness == "borderline" || scores.ApplyPlanCoverage == "borderline" {
		return "borderline"
	}
	return "pass"
}

func scoreOverall(scores Scores) string {
	for _, score := range []string{
		scores.Schema,
		scores.WikiSafety,
		scores.CandidatePaths,
		scores.SourceCompleteness,
		scores.ApplyPlanCoverage,
		scores.ApplyReadiness,
	} {
		if score == "fail" {
			return "fail"
		}
	}
	if scores.DuplicateDetection == "borderline" || scores.SourceCompleteness == "borderline" || scores.ApplyPlanCoverage == "borderline" || scores.ApplyReadiness == "borderline" {
		return "borderline"
	}
	return "pass"
}

func scoreApplyPlanCoverage(bundleDir string, applyPlan string, trace *[]string) string {
	paths := candidatePaths(applyPlan)
	if len(paths) == 0 {
		*trace = append(*trace, "apply_plan_coverage: no candidate paths to inspect")
		return "fail"
	}

	sourceCandidates := 0
	missing := []string{}
	for _, path := range paths {
		switch {
		case strings.HasPrefix(path, "wiki/sources/"):
			sourceCandidates++
		case strings.HasPrefix(path, "wiki/concepts/") || strings.HasPrefix(path, "wiki/topics/"):
			if _, err := os.Stat(filepath.Join(bundleDir, "candidates", filepath.FromSlash(path))); err != nil {
				missing = append(missing, path)
			}
		}
	}

	if sourceCandidates > 0 {
		sourceDraft, err := os.ReadFile(filepath.Join(bundleDir, "source.md"))
		if err != nil {
			*trace = append(*trace, "apply_plan_coverage: source candidates listed but source.md is missing")
			return "fail"
		}
		frontmatter, _ := splitCandidateFrontmatter(string(sourceDraft))
		if !strings.Contains(frontmatter, "kind: source") {
			*trace = append(*trace, "apply_plan_coverage: source candidates listed but source.md is not kind: source")
			return "fail"
		}
	}

	if len(missing) > 0 {
		*trace = append(*trace, "apply_plan_coverage: missing generated candidates: "+strings.Join(missing, ", "))
		return "fail"
	}
	if sourceCandidates > 1 {
		*trace = append(*trace, fmt.Sprintf("apply_plan_coverage: %d source candidates share one source.md artifact", sourceCandidates))
		return "borderline"
	}
	*trace = append(*trace, "apply_plan_coverage: planned pages are represented by generated artifacts")
	return "pass"
}

func scoreSourceCompleteness(bundleDir string, trace *[]string) string {
	sourceDraft, err := os.ReadFile(filepath.Join(bundleDir, "source.md"))
	if err != nil {
		*trace = append(*trace, "source_completeness: cannot read source.md")
		return "fail"
	}
	ingestPlan, err := os.ReadFile(filepath.Join(bundleDir, "ingest-plan.md"))
	if err != nil {
		*trace = append(*trace, "source_completeness: cannot read ingest-plan.md")
		return "fail"
	}

	rawRefs := rawSourceRefs(string(sourceDraft))
	if len(rawRefs) == 0 {
		*trace = append(*trace, "source_completeness: no raw source references to inspect")
		return "pass"
	}

	hasTruncatedRaw := false
	for _, ref := range rawRefs {
		content, ok := readReferencedRaw(bundleDir, ref)
		if !ok {
			continue
		}
		if truncationMarkerRE.MatchString(content) {
			hasTruncatedRaw = true
			break
		}
	}
	if !hasTruncatedRaw {
		*trace = append(*trace, "source_completeness: no raw truncation markers found")
		return "pass"
	}

	generated := string(sourceDraft) + "\n" + string(ingestPlan)
	hasClaim := hasUnsupportedCompletenessClaim(generated)
	hasMention := truncationMentionRE.MatchString(generated)
	switch {
	case hasClaim:
		*trace = append(*trace, "source_completeness: raw source is truncated but generated drafts make completeness claims")
		return "fail"
	case !hasMention:
		*trace = append(*trace, "source_completeness: raw source is truncated but generated drafts do not mention the truncation boundary")
		return "borderline"
	default:
		*trace = append(*trace, "source_completeness: generated drafts acknowledge raw truncation")
		return "pass"
	}
}

func hasUnsupportedCompletenessClaim(generated string) bool {
	for _, line := range strings.Split(generated, "\n") {
		if !completenessClaimRE.MatchString(line) {
			continue
		}
		if truncationMentionRE.MatchString(line) || completenessNegationRE.MatchString(line) {
			continue
		}
		return true
	}
	return false
}

func candidatePaths(applyPlan string) []string {
	lines := candidateLines(applyPlan)
	paths := make([]string, 0, len(lines))
	for _, line := range lines {
		match := candidateLineRE.FindStringSubmatch(line)
		if len(match) == 2 {
			paths = append(paths, match[1])
		}
	}
	return paths
}

func candidateLines(applyPlan string) []string {
	section := applyPlan
	if start := strings.Index(applyPlan, "## Candidate Changes"); start != -1 {
		section = applyPlan[start:]
	}
	if end := strings.Index(section, "\n## "); end != -1 {
		section = section[:end]
	}
	lines := []string{}
	for _, line := range strings.Split(section, "\n") {
		if candidateLineRE.MatchString(line) {
			lines = append(lines, line)
		}
	}
	return lines
}

func rawSourceRefs(sourceDraft string) []string {
	matches := sourceRefRE.FindAllStringSubmatch(sourceDraft, -1)
	refs := make([]string, 0, len(matches))
	seen := map[string]bool{}
	for _, match := range matches {
		if len(match) != 2 {
			continue
		}
		ref := strings.TrimSpace(match[1])
		if ref == "" || seen[ref] {
			continue
		}
		seen[ref] = true
		refs = append(refs, ref)
	}
	return refs
}

func readReferencedRaw(bundleDir string, ref string) (string, bool) {
	candidates := []string{ref}
	if cwd, err := os.Getwd(); err == nil {
		candidates = append(candidates, filepath.Join(cwd, filepath.FromSlash(ref)))
	}
	for dir := filepath.Clean(bundleDir); ; dir = filepath.Dir(dir) {
		candidates = append(candidates, filepath.Join(dir, filepath.FromSlash(ref)))
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
	}
	for _, candidate := range candidates {
		content, err := os.ReadFile(candidate)
		if err == nil {
			return string(content), true
		}
	}
	return "", false
}

func frontmatterValue(frontmatter string, re *regexp.Regexp) string {
	match := re.FindStringSubmatch(frontmatter)
	if len(match) != 2 {
		return ""
	}

	return strings.Trim(strings.TrimSpace(match[1]), `"'`)
}
