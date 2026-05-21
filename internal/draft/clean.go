package draft

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

var (
	ErrNoFrontmatter = errors.New("draft output does not contain complete frontmatter")

	thinkingBlockRE          = regexp.MustCompile(`(?s)\[Start thinking\].*?(?:\[End thinking\]|\z)`)
	markdownFenceRE          = regexp.MustCompile("(?s)```(?:markdown|md|yaml|yml)?\\s*\\n(.*?)\\n```")
	frontmatterLine          = regexp.MustCompile(`(?m)^---\s*$`)
	pseudoFrontmatterPrelude = regexp.MustCompile(`(?s)^\s*(title|kind|sources|confidence|path):.*?\n---\s*\n?`)
	backspaceRune            = "\b"
	performanceLine          = "[ Prompt:"
	exitingLine              = "Exiting..."
	memoryLogPrefix          = "common_memory_breakdown_print:"
)

// Clean converts raw llama.cpp output into a Markdown draft.
//
// The raw string may contain prompt echoes, thinking blocks, progress output,
// performance lines, and llama.cpp runtime logs.
//
// The returned string should contain only the cleaned Markdown page.
//
// The error is returned if the raw output cannot be interpreted as a usable
// Markdown draft.
func Clean(raw string) (string, error) {
	fencedCandidates := markdownFenceRE.FindAllStringSubmatch(raw, -1)
	withoutThinking := thinkingBlockRE.ReplaceAllString(raw, "\n")

	if cleaned, ok := cleanCandidate(withoutThinking); ok {
		return cleaned, nil
	}

	for i := len(fencedCandidates) - 1; i >= 0; i-- {
		if cleaned, ok := cleanCandidate(fencedCandidates[i][1]); ok {
			return cleaned, nil
		}
	}

	return "", ErrNoFrontmatter
}

func cleanCandidate(raw string) (string, bool) {
	text := stripNoise(raw)
	text = unwrapWholeDocumentFence(text)
	text = normalizePseudoFrontmatter(text)
	text = extractFromFrontmatter(text)
	if !hasCompleteFrontmatter(text) {
		return "", false
	}

	return strings.TrimSpace(text) + "\n", true
}

func unwrapWholeDocumentFence(text string) string {
	matches := markdownFenceRE.FindAllStringSubmatch(text, -1)
	if len(matches) != 1 || strings.TrimSpace(matches[0][0]) != strings.TrimSpace(text) {
		return text
	}
	return strings.TrimSpace(matches[0][1])
}

func normalizePseudoFrontmatter(text string) string {
	if strings.HasPrefix(strings.TrimLeftFunc(text, unicode.IsSpace), "---") || !pseudoFrontmatterPrelude.MatchString(text) {
		return text
	}
	return "---\n" + text
}

func stripNoise(raw string) string {
	text := strings.ReplaceAll(raw, backspaceRune, "")
	if idx := strings.Index(text, memoryLogPrefix); idx != -1 {
		text = text[:idx]
	}

	lines := strings.Split(text, "\n")
	kept := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, performanceLine) {
			continue
		}
		if trimmed == exitingLine {
			continue
		}
		if strings.HasPrefix(trimmed, memoryLogPrefix) {
			continue
		}
		kept = append(kept, strings.TrimRight(line, " \t\r"))
	}

	return strings.TrimSpace(strings.Join(kept, "\n"))
}

func extractFromFrontmatter(text string) string {
	loc := frontmatterLine.FindStringIndex(text)
	if loc == nil {
		return text
	}

	return text[loc[0]:]
}

func hasCompleteFrontmatter(text string) bool {
	matches := frontmatterLine.FindAllStringIndex(strings.TrimLeftFunc(text, unicode.IsSpace), 2)
	return len(matches) >= 2
}
