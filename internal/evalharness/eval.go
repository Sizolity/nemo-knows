package evalharness

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var candidateLineRE = regexp.MustCompile("(?m)^- `([^`]+)`")

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
	if scores.Schema == "fail" || scores.WikiSafety == "fail" || scores.CandidatePaths == "fail" {
		return "fail"
	}
	if scores.DuplicateDetection == "borderline" {
		return "borderline"
	}
	return "pass"
}

func scoreOverall(scores Scores) string {
	for _, score := range []string{
		scores.Schema,
		scores.WikiSafety,
		scores.CandidatePaths,
		scores.ApplyReadiness,
	} {
		if score == "fail" {
			return "fail"
		}
	}
	if scores.DuplicateDetection == "borderline" || scores.ApplyReadiness == "borderline" {
		return "borderline"
	}
	return "pass"
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

func frontmatterValue(frontmatter string, re *regexp.Regexp) string {
	match := re.FindStringSubmatch(frontmatter)
	if len(match) != 2 {
		return ""
	}

	return strings.Trim(strings.TrimSpace(match[1]), `"'`)
}
