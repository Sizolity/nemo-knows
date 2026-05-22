package evalharness

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type RegressionResult struct {
	Overall string                 `json:"overall"`
	Summary RegressionSummary      `json:"summary"`
	Cases   []RegressionCaseResult `json:"cases"`
}

type RegressionSummary struct {
	Total  int `json:"total"`
	Passed int `json:"passed"`
	Failed int `json:"failed"`
}

type RegressionCaseResult struct {
	Case          string            `json:"case"`
	Source        string            `json:"source"`
	Bundle        string            `json:"bundle"`
	Status        string            `json:"status"`
	MinimumScores map[string]string `json:"minimum_scores"`
	ActualScores  Scores            `json:"actual_scores"`
	Trace         []string          `json:"trace"`
	Failures      []string          `json:"failures"`
}

type regressionExpected struct {
	Case          string            `json:"case"`
	Source        string            `json:"source"`
	Bundle        string            `json:"bundle"`
	MinimumScores map[string]string `json:"minimum_scores"`
}

func EvaluateRegressionCases(casesDir string) (RegressionResult, error) {
	entries, err := os.ReadDir(casesDir)
	if err != nil {
		return RegressionResult{}, fmt.Errorf("read cases dir: %w", err)
	}

	result := RegressionResult{Overall: "pass"}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		caseDir := filepath.Join(casesDir, entry.Name())
		expected, err := readRegressionExpected(caseDir)
		if err != nil {
			return RegressionResult{}, err
		}
		caseResult, err := evaluateRegressionCase(caseDir, expected)
		if err != nil {
			return RegressionResult{}, err
		}
		result.Cases = append(result.Cases, caseResult)
	}
	sort.Slice(result.Cases, func(i, j int) bool {
		return result.Cases[i].Case < result.Cases[j].Case
	})
	result.Summary.Total = len(result.Cases)
	for _, item := range result.Cases {
		if item.Status == "pass" {
			result.Summary.Passed++
		} else {
			result.Summary.Failed++
			result.Overall = "fail"
		}
	}

	return result, nil
}

func readRegressionExpected(caseDir string) (regressionExpected, error) {
	content, err := os.ReadFile(filepath.Join(caseDir, "expected.json"))
	if err != nil {
		return regressionExpected{}, fmt.Errorf("read expected for %s: %w", caseDir, err)
	}
	var expected regressionExpected
	if err := json.Unmarshal(content, &expected); err != nil {
		return regressionExpected{}, fmt.Errorf("parse expected for %s: %w", caseDir, err)
	}
	if expected.Case == "" {
		expected.Case = filepath.Base(caseDir)
	}
	if expected.Bundle == "" {
		expected.Bundle = filepath.Join(caseDir, "bundle")
	}
	return expected, nil
}

func evaluateRegressionCase(caseDir string, expected regressionExpected) (RegressionCaseResult, error) {
	bundle := expected.Bundle
	if !filepath.IsAbs(bundle) {
		// Preserve repository-relative bundles, but make test temp bundles work too.
		if _, err := os.Stat(bundle); err != nil {
			bundle = filepath.Join(caseDir, "bundle")
		}
	}
	evaluated, err := EvaluateBundle(bundle)
	if err != nil {
		return RegressionCaseResult{}, fmt.Errorf("evaluate case %s: %w", expected.Case, err)
	}
	result := RegressionCaseResult{
		Case:          expected.Case,
		Source:        expected.Source,
		Bundle:        bundle,
		Status:        "pass",
		MinimumScores: expected.MinimumScores,
		ActualScores:  evaluated.Scores,
		Trace:         evaluated.Trace,
	}
	for scoreName, minimum := range expected.MinimumScores {
		actual := scoreByName(evaluated.Scores, scoreName)
		if !scoreMeetsMinimum(actual, minimum) {
			result.Status = "fail"
			result.Failures = append(result.Failures, fmt.Sprintf("%s: got %s, want at least %s", scoreName, actual, minimum))
		}
	}
	return result, nil
}

func scoreByName(scores Scores, name string) string {
	switch name {
	case "schema":
		return scores.Schema
	case "wiki_safety":
		return scores.WikiSafety
	case "candidate_paths":
		return scores.CandidatePaths
	case "duplicate_detection":
		return scores.DuplicateDetection
	case "source_completeness":
		return scores.SourceCompleteness
	case "apply_plan_coverage":
		return scores.ApplyPlanCoverage
	case "apply_readiness":
		return scores.ApplyReadiness
	case "overall":
		return scores.Overall
	default:
		return ""
	}
}

func scoreMeetsMinimum(actual string, minimum string) bool {
	return scoreRank(actual) >= scoreRank(minimum)
}

func scoreRank(score string) int {
	switch score {
	case "pass":
		return 3
	case "borderline":
		return 2
	case "fail":
		return 1
	default:
		return 0
	}
}
