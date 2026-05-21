package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestMain clears NEMO_*, HTTPS_PROXY, HTTP_PROXY, and NO_PROXY environment
// variables before running the test suite. Without this, a developer shell that
// sourced .env would inject DeepSeek provider/credential vars into Defaults()
// and turn local fake-llama bundle tests into real DeepSeek API calls.
func TestMain(m *testing.M) {
	for _, key := range []string{
		"NEMO_MODEL_PROVIDER",
		"NEMO_LLAMA_CLI",
		"NEMO_LLAMA_MODEL",
		"NEMO_DEEPSEEK_API_KEY",
		"NEMO_DEEPSEEK_BASE_URL",
		"NEMO_DEEPSEEK_MODEL",
		"NEMO_DEEPSEEK_MAX_TOKENS",
		"NEMO_DEEPSEEK_THINKING",
		"NEMO_DEEPSEEK_REASONING_EFFORT",
		"NEMO_DEEPSEEK_TEMPERATURE",
		"NEMO_DEEPSEEK_TOP_P",
		"NEMO_DEEPSEEK_RESPONSE_FORMAT",
		"NEMO_DEEPSEEK_USER_ID",
		"NEMO_DEEPSEEK_SYSTEM_PROMPT",
		"NEMO_MAX_TOKENS",
		"NEMO_CHUNKED_THRESHOLD_CHARS",
		"NEMO_MAX_CHUNK_CHARS",
		"HTTPS_PROXY",
		"HTTP_PROXY",
		"NO_PROXY",
	} {
		_ = os.Unsetenv(key)
	}
	os.Exit(m.Run())
}

func TestRunGeneratesRawAndCleanedDraft(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	source := filepath.Join(dir, "source.md")
	prompt := filepath.Join(dir, "prompt.md")
	out := filepath.Join(dir, "draft.md")
	fake := filepath.Join(dir, "llama-cli")

	if err := os.WriteFile(source, []byte("# Source\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile(prompt, []byte("Path: RAW_SOURCE_PATH\nRAW_SOURCE_CONTENT"), 0o644); err != nil {
		t.Fatalf("write prompt: %v", err)
	}
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
cat <<'EOF'
[Start thinking]
private reasoning
[End thinking]
---
kind: source
title: Test
---
# Draft
[ Prompt: 1 t/s | Generation: 2 t/s ]
Exiting...
EOF
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}

	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-source", source, "-prompt", prompt, "-out", out}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}

	cleaned, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read cleaned draft: %v", err)
	}
	if !strings.HasPrefix(string(cleaned), "---\nkind: source") {
		t.Fatalf("unexpected cleaned draft:\n%s", cleaned)
	}
	if strings.Contains(string(cleaned), "private reasoning") {
		t.Fatalf("cleaned draft still contains thinking:\n%s", cleaned)
	}

	rawPath := filepath.Join(dir, "draft.raw.txt")
	if _, err := os.Stat(rawPath); err != nil {
		t.Fatalf("expected raw output at %s: %v", rawPath, err)
	}
}

func TestRunProviderFlagOverridesDotEnvProvider(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}

	if err := os.WriteFile(".env", []byte("NEMO_MODEL_PROVIDER=deepseek\n"), 0o600); err != nil {
		t.Fatalf("write .env: %v", err)
	}
	// loadDotEnv writes into the real process environment, not into a
	// t.Setenv-tracked override, so we have to unset the variable explicitly
	// to avoid leaking the deepseek provider into later tests in this package.
	t.Cleanup(func() { _ = os.Unsetenv("NEMO_MODEL_PROVIDER") })
	source := filepath.Join(dir, "source.md")
	promptPath := filepath.Join(dir, "prompt.md")
	out := filepath.Join(dir, "draft.md")
	fake := filepath.Join(dir, "llama-cli")
	if err := os.WriteFile(source, []byte("# Source\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile(promptPath, []byte("Path: RAW_SOURCE_PATH\nRAW_SOURCE_CONTENT"), 0o644); err != nil {
		t.Fatalf("write prompt: %v", err)
	}
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
cat <<'EOF'
---
kind: source
title: Test
---
# Draft
EOF
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}
	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-provider", "llama", "-source", source, "-prompt", promptPath, "-out", out}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("expected draft output: %v", err)
	}
}

func TestRunRequiresFlags(t *testing.T) {
	if code := run(nil); code == 0 {
		t.Fatal("expected non-zero exit code for missing flags")
	}
}

func TestRunGeneratesLocalIngestBundle(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}

	for _, path := range []string{"raw", "prompts"} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	if err := os.WriteFile("raw/source.md", []byte("# Source\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile("prompts/source-page.md", []byte("source {{RAW_SOURCE_CONTENT}}"), 0o644); err != nil {
		t.Fatalf("write source prompt: %v", err)
	}
	if err := os.WriteFile("prompts/ingest-plan.md", []byte("plan {{RAW_SOURCE_CONTENT}}"), 0o644); err != nil {
		t.Fatalf("write plan prompt: %v", err)
	}

	fake := filepath.Join(dir, "llama-cli")
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
cat <<'EOF'
---
kind: source
title: Test
---
# Draft
EOF
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}
	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-source", "raw/source.md", "-bundle-dir", "drafts/source"}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}

	for _, path := range []string{
		"drafts/source/source.md",
		"drafts/source/source.raw.txt",
		"drafts/source/ingest-plan.md",
		"drafts/source/ingest-plan.raw.txt",
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected bundle file %s: %v", path, err)
		}
	}
	sourceDraft, err := os.ReadFile("drafts/source/source.md")
	if err != nil {
		t.Fatalf("read source draft: %v", err)
	}
	for _, want := range []string{
		"title: Test",
		"kind: source",
		"sources:\n  - raw/source.md",
		"confidence: medium",
	} {
		if !strings.Contains(string(sourceDraft), want) {
			t.Fatalf("source draft missing %q:\n%s", want, sourceDraft)
		}
	}
}

func TestRunGeneratesChunkedBundleForLongSource(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}

	for _, path := range []string{"raw", "prompts"} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	longSource := "# Long Source\n\n1. First Section\n\n" +
		strings.Repeat("alpha beta gamma delta.\n\n", 3000) +
		"2. Second Section\n\n" +
		strings.Repeat("epsilon zeta eta theta.\n\n", 3000)
	if err := os.WriteFile("raw/long.md", []byte(longSource), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile("prompts/source-page.md", []byte("source {{RAW_SOURCE_CONTENT}}"), 0o644); err != nil {
		t.Fatalf("write source prompt: %v", err)
	}
	if err := os.WriteFile("prompts/ingest-plan.md", []byte("plan {{RAW_SOURCE_CONTENT}}"), 0o644); err != nil {
		t.Fatalf("write plan prompt: %v", err)
	}
	if err := os.WriteFile("prompts/chunk-notes.md", []byte("Create concise source-backed notes\n{{CHUNK_CONTENT}}"), 0o644); err != nil {
		t.Fatalf("write chunk notes prompt: %v", err)
	}
	if err := os.WriteFile("prompts/chunk-group-notes.md", []byte("Create group-level notes\n{{CHUNK_NOTES}}"), 0o644); err != nil {
		t.Fatalf("write chunk group notes prompt: %v", err)
	}
	if err := os.WriteFile("prompts/chunk-source-page.md", []byte("Create a source summary page from chunk notes\n{{CHUNK_NOTES}}"), 0o644); err != nil {
		t.Fatalf("write chunk source prompt: %v", err)
	}
	if err := os.WriteFile("prompts/chunk-ingest-plan.md", []byte("Create a reviewed ingest plan draft from chunk notes\n{{CHUNK_NOTES}}"), 0o644); err != nil {
		t.Fatalf("write chunk ingest prompt: %v", err)
	}

	fake := filepath.Join(dir, "llama-cli")
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
prompt_file=""
while [ "$#" -gt 0 ]; do
  if [ "$1" = "-f" ]; then
    shift
    prompt_file="$1"
    break
  fi
  shift
done
prompt="$(cat "$prompt_file")"
case "$prompt" in
  *"Create concise source-backed notes"*)
    cat <<'EOF'
---
kind: topic
sources: [raw/long.md]
---

# Chunk Notes

## Chunk Context
Long source section.

## Local Summary
- Local summary.

## Key Claims
- Claim.

## Entities And Concepts
- Concept.

## Procedures And API Details
- Detail.

## Nuance Or Contradictions
- none

## Candidate Wiki Hints
- wiki/concepts/long-source.md
EOF
    ;;
  *"Create group-level notes"*)
    cat <<'EOF'
---
kind: topic
sources: [raw/long.md]
---

# Chunk Group Notes

## Group Context
Grouped long source section.

## Cross-Chunk Summary
- Group summary.

## Repeated Or Central Claims
- Claim.

## Important Local Details
- Detail.

## Candidate Wiki Hints
- wiki/concepts/long-source.md

## Gaps Or Cautions
- none
EOF
    ;;
  *"Create a source summary page from chunk notes"*)
    cat <<'EOF'
---
kind: source
sources:
  - raw/long.md
---

# Long Source

## What It Is
Long source summary.

## Summary
Summary.

## Key Claims
- Claim.

## Suggested Links
- none
EOF
    ;;
  *"Create a reviewed ingest plan draft from chunk notes"*)
    cat <<'EOF'
---
kind: topic
sources: [raw/long.md]
status: draft
---

# Ingest Plan

## Source Summary
- Summary.

## Candidate Wiki Pages
- wiki/sources/long-source.md — source page
- wiki/concepts/long-source.md — concept page

## Suggested Links
- none

## Review Checklist
- [ ] Review.
EOF
    ;;
  *) echo "unexpected prompt" >&2; exit 1 ;;
esac
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}
	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-source", "raw/long.md", "-bundle-dir", "drafts/long", "-profile", "stable"}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}

	for _, path := range []string{
		"drafts/long/source.md",
		"drafts/long/ingest-plan.md",
		"drafts/long/chunks/outline.md",
		"drafts/long/chunks/chunk-index.json",
		"drafts/long/chunks/combined-notes.md",
		"drafts/long/chunks/combined-group-notes.md",
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected chunked bundle file %s: %v", path, err)
		}
	}
	chunks, err := filepath.Glob("drafts/long/chunks/chunk-*.md")
	if err != nil {
		t.Fatalf("glob chunks: %v", err)
	}
	if len(chunks) < 2 {
		t.Fatalf("expected multiple chunk notes, got %d", len(chunks))
	}
	groups, err := filepath.Glob("drafts/long/chunks/group-*.md")
	if err != nil {
		t.Fatalf("glob chunk groups: %v", err)
	}
	if len(groups) == 0 {
		t.Fatalf("expected grouped chunk notes")
	}
}

func TestNormalizeSourceDraftDemotesRequiredSectionHeadings(t *testing.T) {
	got := normalizeSourceDraft(`---
title: Branches
kind: source
---

# What It Is
Git branches are pointers.

# Summary
Summary text.

# Key Claims
- Claim.

# Suggested Links
- none
`, "raw/web/branches.md")

	for _, want := range []string{
		"## What It Is",
		"## Summary",
		"## Key Claims",
		"## Suggested Links",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("normalized source missing %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, "\n# What It Is") {
		t.Fatalf("source section remained top-level:\n%s", got)
	}
}

func TestRunAcceptsProfileFlag(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	source := filepath.Join(dir, "source.md")
	prompt := filepath.Join(dir, "prompt.md")
	out := filepath.Join(dir, "draft.md")
	fake := filepath.Join(dir, "llama-cli")

	if err := os.WriteFile(source, []byte("# Source\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile(prompt, []byte("{{RAW_SOURCE_CONTENT}}"), 0o644); err != nil {
		t.Fatalf("write prompt: %v", err)
	}
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
case "$*" in
  *"--top-k 20"*"--min-p 0"*) ;;
  *) echo "missing qwen sampling args: $*" >&2; exit 1 ;;
esac
cat <<'EOF'
---
kind: source
---
# Draft
EOF
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}

	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-source", source, "-prompt", prompt, "-out", out, "-profile", "stable"}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}
}

func TestRunBundlePersistsWebSourceBeforeGeneration(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}
	for _, path := range []string{"drafts/web", "prompts"} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	if err := os.WriteFile("drafts/web/qwen-llama-cpp.md", []byte("# Qwen llama.cpp\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile("prompts/source-page.md", []byte("{{RAW_SOURCE_PATH}}"), 0o644); err != nil {
		t.Fatalf("write source prompt: %v", err)
	}
	if err := os.WriteFile("prompts/ingest-plan.md", []byte("{{RAW_SOURCE_PATH}}"), 0o644); err != nil {
		t.Fatalf("write ingest prompt: %v", err)
	}
	fake := filepath.Join(dir, "llama-cli")
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
prompt_file=""
while [ "$#" -gt 0 ]; do
  if [ "$1" = "-f" ]; then
    shift
    prompt_file="$1"
    break
  fi
  shift
done
case "$(cat "$prompt_file")" in
  *"raw/web/qwen-llama-cpp.md"*) ;;
  *) echo "missing durable raw path in prompt" >&2; exit 1 ;;
esac
cat <<'EOF'
---
kind: source
sources:
  - raw/web/qwen-llama-cpp.md
---

# Draft
EOF
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}
	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-source", "drafts/web/qwen-llama-cpp.md", "-bundle-dir", "drafts/web-e2e-qwen", "-persist-raw-web"}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}

	persisted, err := os.ReadFile("raw/web/qwen-llama-cpp.md")
	if err != nil {
		t.Fatalf("read persisted raw source: %v", err)
	}
	if string(persisted) != "# Qwen llama.cpp\n" {
		t.Fatalf("persisted raw source changed content:\n%s", persisted)
	}
}

func TestRunRetriesWithFallbackProfileWhenCleanFails(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	source := filepath.Join(dir, "source.md")
	prompt := filepath.Join(dir, "prompt.md")
	out := filepath.Join(dir, "draft.md")
	fake := filepath.Join(dir, "llama-cli")
	attempts := filepath.Join(dir, "attempts")

	if err := os.WriteFile(source, []byte("# Source\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile(prompt, []byte("{{RAW_SOURCE_CONTENT}}"), 0o644); err != nil {
		t.Fatalf("write prompt: %v", err)
	}
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
attempts="$(cat "$NEMO_TEST_ATTEMPTS" 2>/dev/null || true)"
if [ -z "$attempts" ]; then
  echo 1 > "$NEMO_TEST_ATTEMPTS"
  echo "not markdown yet"
  exit 0
fi
case "$*" in
  *"-n 16384"*"--temp 0.2"*) ;;
  *) echo "missing fallback profile args: $*" >&2; exit 1 ;;
esac
cat <<'EOF'
---
kind: source
---
# Fallback Draft
EOF
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}

	t.Setenv("NEMO_TEST_ATTEMPTS", attempts)
	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-source", source, "-prompt", prompt, "-out", out, "-profile", "stable"}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}

	cleaned, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read cleaned draft: %v", err)
	}
	if !strings.Contains(string(cleaned), "# Fallback Draft") {
		t.Fatalf("expected fallback cleaned draft, got:\n%s", cleaned)
	}

	if _, err := os.Stat(filepath.Join(dir, "draft.fallback.raw.txt")); err != nil {
		t.Fatalf("expected fallback raw output: %v", err)
	}
}

func TestRunReviewsBundle(t *testing.T) {
	dir := t.TempDir()
	bundle := filepath.Join(dir, "bundle")
	if err := os.MkdirAll(bundle, 0o755); err != nil {
		t.Fatalf("mkdir bundle: %v", err)
	}
	if err := os.WriteFile(filepath.Join(bundle, "source.md"), []byte(`---
kind: source
---

# Source

## What It Is
Ok.

## Summary
Ok.

## Key Claims
Ok.

## Suggested Links
Ok.
`), 0o644); err != nil {
		t.Fatalf("write source draft: %v", err)
	}
	if err := os.WriteFile(filepath.Join(bundle, "ingest-plan.md"), []byte(`---
kind: topic
---

# Ingest Plan

## Source Summary
Ok.

## Candidate Wiki Pages
- wiki/sources/source.md — source page

## Suggested Links
Ok.

## Review Checklist
Ok.
`), 0o644); err != nil {
		t.Fatalf("write ingest plan: %v", err)
	}

	out := filepath.Join(bundle, "apply-plan.md")
	if code := run([]string{"-review-bundle", bundle, "-out", out}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}

	plan, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read apply plan: %v", err)
	}
	if !strings.Contains(string(plan), "# Reviewed Ingest Apply Plan") {
		t.Fatalf("unexpected apply plan:\n%s", plan)
	}
}

func TestRunEvaluatesBundle(t *testing.T) {
	dir := t.TempDir()
	bundle := filepath.Join(dir, "bundle")
	outDir := filepath.Join(dir, "run")
	if err := os.MkdirAll(bundle, 0o755); err != nil {
		t.Fatalf("mkdir bundle: %v", err)
	}
	if err := os.WriteFile(filepath.Join(bundle, "apply-plan.md"), []byte(`# Reviewed Ingest Apply Plan

This is a review artifact. Do not apply this plan automatically.

## Validation

- [x] `+"`source.md`"+` has YAML frontmatter
- [x] `+"`ingest-plan.md`"+` has YAML frontmatter
- [x] `+"`source.md`"+` frontmatter `+"`kind`"+` is `+"`source`"+`
- [x] `+"`ingest-plan.md`"+` frontmatter `+"`kind`"+` is `+"`topic`"+`
- [x] `+"`source.md`"+` includes required section `+"`What It Is`"+`
- [x] `+"`source.md`"+` includes required section `+"`Summary`"+`
- [x] `+"`source.md`"+` includes required section `+"`Key Claims`"+`
- [x] `+"`source.md`"+` includes required section `+"`Suggested Links`"+`
- [x] `+"`ingest-plan.md`"+` includes required section `+"`Source Summary`"+`
- [x] `+"`ingest-plan.md`"+` includes required section `+"`Candidate Wiki Pages`"+`
- [x] `+"`ingest-plan.md`"+` includes required section `+"`Suggested Links`"+`
- [x] `+"`ingest-plan.md`"+` includes required section `+"`Review Checklist`"+`

## Candidate Changes

- `+"`wiki/sources/llm-wiki.md`"+` — update existing page.
`), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}

	if code := run([]string{"-eval-bundle", bundle, "-out-dir", outDir}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}
	for _, path := range []string{
		filepath.Join(outDir, "scores.json"),
		filepath.Join(outDir, "trace.md"),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected eval output %s: %v", path, err)
		}
	}
}

func TestRunEvaluatesCandidateDrafts(t *testing.T) {
	dir := t.TempDir()
	bundle := filepath.Join(dir, "bundle")
	outDir := filepath.Join(dir, "run")
	if err := os.MkdirAll(filepath.Join(bundle, "candidates", "wiki", "concepts"), 0o755); err != nil {
		t.Fatalf("mkdir bundle: %v", err)
	}
	if err := os.WriteFile(filepath.Join(bundle, "source.md"), []byte("---\nkind: source\nsources:\n  - raw/llm-wiki.md\n---\n\n# Source\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile(filepath.Join(bundle, "apply-plan.md"), []byte("## Candidate Changes\n\n- `wiki/concepts/llm-maintenance-pattern.md` — create new page.\n"), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}
	if err := os.WriteFile(filepath.Join(bundle, "candidates", "wiki", "concepts", "llm-maintenance-pattern.md"), []byte(`---
title: LLM Maintenance Pattern
kind: concept
sources:
  - source.md
  - raw/llm-wiki.md
confidence: medium
---

# LLM Maintenance Pattern

The [[LLM Wiki]] maintenance pattern describes how an LLM keeps a durable wiki current across ingest, query, and lint operations. It turns source summaries into maintained pages rather than treating every answer as temporary context.
`), 0o644); err != nil {
		t.Fatalf("write candidate: %v", err)
	}

	if code := run([]string{"-eval-candidates", bundle, "-out-dir", outDir}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}
	for _, path := range []string{
		filepath.Join(outDir, "candidate-scores.json"),
		filepath.Join(outDir, "candidate-trace.md"),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected candidate eval output %s: %v", path, err)
		}
	}
}

func TestRunReviewsCandidateDraftIssues(t *testing.T) {
	dir := t.TempDir()
	bundle := filepath.Join(dir, "bundle")
	outDir := filepath.Join(dir, "run")
	if err := os.MkdirAll(filepath.Join(bundle, "candidates", "wiki", "concepts"), 0o755); err != nil {
		t.Fatalf("mkdir bundle: %v", err)
	}
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}
	if err := os.MkdirAll(filepath.Join("wiki", "concepts"), 0o755); err != nil {
		t.Fatalf("mkdir wiki: %v", err)
	}
	if err := os.WriteFile(filepath.Join("wiki", "concepts", "unrelated.md"), []byte("---\ntitle: Unrelated\nkind: concept\nsources:\n  - raw/other.md\nconfidence: medium\n---\n\n# Unrelated\n"), 0o644); err != nil {
		t.Fatalf("write wiki page: %v", err)
	}
	if err := os.WriteFile(filepath.Join(bundle, "source.md"), []byte("---\nkind: source\nsources:\n  - raw/sqlite-wal.md\n---\n\n# Source\n\nSQLite WAL uses checkpointing.\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile(filepath.Join(bundle, "apply-plan.md"), []byte("## Candidate Changes\n\n- `wiki/concepts/checkpointing.md` — create new page.\n"), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}
	if err := os.WriteFile(filepath.Join(bundle, "candidates", "wiki", "concepts", "checkpointing.md"), []byte(`---
title: Checkpointing
kind: concept
sources:
  - source.md
  - raw/sqlite-wal.md
confidence: medium
---

# Checkpointing

SQLite WAL checkpointing copies committed frames back into the database while linking to [[Unrelated]], which exists but is not supported by this source. This line is long enough for the candidate length gate.
`), 0o644); err != nil {
		t.Fatalf("write candidate: %v", err)
	}

	if code := run([]string{"-review-candidates", bundle, "-out-dir", outDir}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}
	review, err := os.ReadFile(filepath.Join(outDir, "candidate-review.md"))
	if err != nil {
		t.Fatalf("read candidate review: %v", err)
	}
	content := string(review)
	for _, want := range []string{
		"# Candidate Review",
		"wikilinks: weak semantic targets: Unrelated",
		"Convert unsupported wikilinks to plain text",
		"Do not apply these suggestions automatically.",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("candidate review missing %q:\n%s", want, content)
		}
	}
}

func TestRunLintsWiki(t *testing.T) {
	dir := t.TempDir()
	outDir := filepath.Join(dir, "run")
	if err := os.MkdirAll(filepath.Join(dir, "wiki", "concepts"), 0o755); err != nil {
		t.Fatalf("mkdir wiki: %v", err)
	}
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}
	if err := os.WriteFile("wiki/index.md", []byte("---\ntitle: Index\nkind: index\n---\n\n- [[missing]] — Missing.\n"), 0o644); err != nil {
		t.Fatalf("write index: %v", err)
	}
	if err := os.WriteFile("wiki/log.md", []byte("---\ntitle: Log\nkind: log\n---\n\n## [2026-05-16] note | ok\n"), 0o644); err != nil {
		t.Fatalf("write log: %v", err)
	}
	if err := os.WriteFile("wiki/concepts/no-frontmatter.md", []byte("# No Frontmatter\n"), 0o644); err != nil {
		t.Fatalf("write page: %v", err)
	}

	if code := run([]string{"-lint-wiki", "-out-dir", outDir}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}
	for _, path := range []string{
		filepath.Join(outDir, "wiki-lint.json"),
		filepath.Join(outDir, "wiki-lint.md"),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected wiki lint output %s: %v", path, err)
		}
	}
}

func TestRunEvaluatesRegressionCases(t *testing.T) {
	dir := t.TempDir()
	casesDir := filepath.Join(dir, "cases")
	outDir := filepath.Join(dir, "run")
	makeCLIRegressionCase(t, casesDir, "article", "wiki/concepts/article-pattern.md")
	makeCLIRegressionCase(t, casesDir, "meeting-notes", "wiki/topics/meeting-decisions.md")

	if code := run([]string{"-eval-regression", casesDir, "-out-dir", outDir}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}
	for _, path := range []string{
		filepath.Join(outDir, "regression-summary.json"),
		filepath.Join(outDir, "regression-summary.md"),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected regression output %s: %v", path, err)
		}
	}
}

func makeCLIRegressionCase(t *testing.T, casesDir string, name string, candidate string) {
	t.Helper()
	caseDir := filepath.Join(casesDir, name)
	if err := os.MkdirAll(filepath.Join(caseDir, "bundle"), 0o755); err != nil {
		t.Fatalf("mkdir case: %v", err)
	}
	if err := os.WriteFile(filepath.Join(caseDir, "expected.json"), []byte(`{
  "case": "`+name+`",
  "source": "fixture:`+name+`",
  "bundle": "`+filepath.ToSlash(filepath.Join(caseDir, "bundle"))+`",
  "minimum_scores": {
    "schema": "pass",
    "wiki_safety": "pass",
    "candidate_paths": "pass",
    "duplicate_detection": "pass",
    "apply_readiness": "pass",
    "overall": "pass"
  }
}
`), 0o644); err != nil {
		t.Fatalf("write expected: %v", err)
	}
	if err := os.WriteFile(filepath.Join(caseDir, "bundle", "source.md"), []byte("---\nkind: source\n---\n# Source\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile(filepath.Join(caseDir, "bundle", "ingest-plan.md"), []byte("---\nkind: topic\n---\n# Ingest Plan\n"), 0o644); err != nil {
		t.Fatalf("write ingest plan: %v", err)
	}
	if err := os.WriteFile(filepath.Join(caseDir, "bundle", "apply-plan.md"), []byte("# Reviewed Ingest Apply Plan\n\n"+
		"This is a review artifact. Do not apply this plan automatically.\n\n"+
		"## Validation\n\n"+
		"- [x] `source.md` has YAML frontmatter\n"+
		"- [x] `ingest-plan.md` has YAML frontmatter\n"+
		"- [x] `source.md` frontmatter `kind` is `source`\n"+
		"- [x] `ingest-plan.md` frontmatter `kind` is `topic`\n"+
		"- [x] `source.md` includes required section `What It Is`\n"+
		"- [x] `source.md` includes required section `Summary`\n"+
		"- [x] `source.md` includes required section `Key Claims`\n"+
		"- [x] `source.md` includes required section `Suggested Links`\n"+
		"- [x] `ingest-plan.md` includes required section `Source Summary`\n"+
		"- [x] `ingest-plan.md` includes required section `Candidate Wiki Pages`\n"+
		"- [x] `ingest-plan.md` includes required section `Suggested Links`\n"+
		"- [x] `ingest-plan.md` includes required section `Review Checklist`\n\n"+
		"## Candidate Changes\n\n"+
		"- `"+candidate+"` — create new page.\n"), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}
}

func TestRunApplyApprovedRequiresApproveFlag(t *testing.T) {
	dir := t.TempDir()
	bundle := filepath.Join(dir, "bundle")
	if err := os.MkdirAll(bundle, 0o755); err != nil {
		t.Fatalf("mkdir bundle: %v", err)
	}

	if code := run([]string{"-apply-approved", bundle}); code == 0 {
		t.Fatal("expected apply command to fail without -approve")
	}
}

func TestRunGeneratesReviewedCandidateDrafts(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}

	for _, path := range []string{"drafts/bundle", "prompts"} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	if err := os.WriteFile("drafts/bundle/source.md", []byte("---\nkind: source\nsources:\n  - raw/llm-wiki.md\n---\n\n# Source Summary\n"), 0o644); err != nil {
		t.Fatalf("write source draft: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/apply-plan.md", []byte("# Apply Plan\n\n## Candidate Changes\n\n- `wiki/concepts/llm-maintenance-pattern.md` — create new page.\n- `wiki/sources/llm-wiki-pattern.md` — create new page.\n- `wiki/topics/persistent-wiki-architecture.md` — create new page.\n"), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}
	if err := os.WriteFile("prompts/concept-page.md", []byte("concept {{PAGE_TITLE}} {{PAGE_KIND}} {{TARGET_PATH}}\n{{SOURCE_CONTENT}}"), 0o644); err != nil {
		t.Fatalf("write concept prompt: %v", err)
	}
	if err := os.WriteFile("prompts/topic-page.md", []byte("topic {{PAGE_TITLE}} {{PAGE_KIND}} {{TARGET_PATH}}\n{{SOURCE_CONTENT}}"), 0o644); err != nil {
		t.Fatalf("write topic prompt: %v", err)
	}
	staleFallback := "drafts/bundle/candidates/wiki/concepts/llm-maintenance-pattern.fallback.raw.txt"
	if err := os.MkdirAll(filepath.Dir(staleFallback), 0o755); err != nil {
		t.Fatalf("mkdir stale fallback parent: %v", err)
	}
	if err := os.WriteFile(staleFallback, []byte("[Start thinking]\nstale"), 0o644); err != nil {
		t.Fatalf("write stale fallback raw: %v", err)
	}
	fake := filepath.Join(dir, "llama-cli")
	attempts := filepath.Join(dir, "candidate-attempts")
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
attempts="$(cat "$NEMO_TEST_CANDIDATE_ATTEMPTS" 2>/dev/null || true)"
if [ -z "$attempts" ]; then
  echo 1 > "$NEMO_TEST_CANDIDATE_ATTEMPTS"
  kind=concept
  title="LLM Maintenance Pattern"
else
  kind=topic
  title="Persistent Wiki Architecture"
fi
cat <<EOF
---
title: $title
kind: $kind
sources:
  - raw/llm-wiki.md
---

# $title

Generated candidate draft.
EOF
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}
	t.Setenv("NEMO_TEST_CANDIDATE_ATTEMPTS", attempts)
	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-generate-candidates", "drafts/bundle"}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}

	for _, path := range []string{
		"drafts/bundle/candidates/wiki/concepts/llm-maintenance-pattern.md",
		"drafts/bundle/candidates/wiki/concepts/llm-maintenance-pattern.raw.txt",
		"drafts/bundle/candidates/wiki/topics/persistent-wiki-architecture.md",
		"drafts/bundle/candidates/wiki/topics/persistent-wiki-architecture.raw.txt",
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected generated candidate artifact %s: %v", path, err)
		}
	}
	if _, err := os.Stat("drafts/bundle/candidates/wiki/sources/llm-wiki-pattern.md"); err == nil {
		t.Fatal("source candidate draft should not be generated")
	}
	if _, err := os.Stat(staleFallback); !os.IsNotExist(err) {
		t.Fatalf("stale fallback raw should be removed after primary success, err=%v", err)
	}
}

func TestRunGenerateCandidatesNormalizesCandidateFrontmatter(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}
	for _, path := range []string{"drafts/bundle", "prompts"} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	if err := os.WriteFile("drafts/bundle/source.md", []byte("---\nkind: source\nsources:\n  - raw/llm-wiki.md\n---\n\n# Source Summary\n"), 0o644); err != nil {
		t.Fatalf("write source draft: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/apply-plan.md", []byte("## Candidate Changes\n\n- `wiki/concepts/llm-maintenance-pattern.md` — create new page.\n"), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}
	if err := os.WriteFile("prompts/concept-page.md", []byte("concept {{PAGE_TITLE}}"), 0o644); err != nil {
		t.Fatalf("write concept prompt: %v", err)
	}
	fake := filepath.Join(dir, "llama-cli")
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
cat <<'EOF'
---
kind: concept
title: LLM Maintenance Pattern
path: wiki/concepts/llm-maintenance-pattern.md
---

# LLM Maintenance Pattern

Generated body.
EOF
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}
	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-generate-candidates", "drafts/bundle"}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}

	candidate, err := os.ReadFile("drafts/bundle/candidates/wiki/concepts/llm-maintenance-pattern.md")
	if err != nil {
		t.Fatalf("read candidate draft: %v", err)
	}
	content := string(candidate)
	for _, want := range []string{
		"title: LLM Maintenance Pattern",
		"kind: concept",
		"sources:\n  - source.md\n  - raw/llm-wiki.md",
		"# LLM Maintenance Pattern",
		"Generated body.",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("normalized candidate missing %q:\n%s", want, content)
		}
	}
	if strings.Contains(content, "path:") {
		t.Fatalf("normalized candidate should remove non-schema path field:\n%s", content)
	}
}

func TestSourceRefsForCandidateExtractsInlineRawSources(t *testing.T) {
	refs := sourceRefsForCandidate([]byte(`---
kind: source
sources: [raw/web/qwen-llama-cpp.md]
confidence: medium
---

# Source Summary
`))

	want := []string{"source.md", "raw/web/qwen-llama-cpp.md"}
	if strings.Join(refs, ",") != strings.Join(want, ",") {
		t.Fatalf("refs = %v, want %v", refs, want)
	}
}

func TestNormalizeCandidateDraftFixesHeadingWithoutAddingWeakLink(t *testing.T) {
	target := candidateDraftTarget{
		Path:  "wiki/concepts/tooling-stack.md",
		Kind:  "concept",
		Title: "Tooling Stack",
	}
	allowedLinks := map[string]bool{
		"ingest":        true,
		"tooling-stack": true,
	}

	content := normalizeCandidateDraft(`---
title: Wrong
kind: concept
---

# [[tooling-stack]]

Ingest processing turns sources into wiki maintenance drafts.`, target, []string{"source.md", "raw/llm-wiki.md"}, allowedLinks)

	if !strings.Contains(content, "# Tooling Stack\n") {
		t.Fatalf("heading should be normalized to target title:\n%s", content)
	}
	if strings.Contains(content, "# [[tooling-stack]]") {
		t.Fatalf("heading wikilink should not remain as the H1:\n%s", content)
	}
	if strings.Contains(content, "[[ingest|Ingest]]") {
		t.Fatalf("body should not auto-add wikilinks just to satisfy structure:\n%s", content)
	}
}

func TestAllowedLinkSlugsOnlyIncludesCandidatesAndSourceSupportedWikiPages(t *testing.T) {
	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}
	for _, path := range []string{"wiki/concepts"} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	for _, path := range []string{
		"wiki/concepts/persistent-wiki.md",
		"wiki/concepts/unrelated.md",
	} {
		if err := os.WriteFile(path, []byte("# Page\n"), 0o644); err != nil {
			t.Fatalf("write wiki page: %v", err)
		}
	}
	allowed := allowedLinkSlugs("The source explicitly mentions a persistent wiki.", []candidateDraftTarget{
		{Path: "wiki/concepts/llama-cli-arguments.md"},
	})

	if !allowed["llama-cli-arguments"] {
		t.Fatalf("candidate target should be allowed: %v", allowed)
	}
	if !allowed["persistent-wiki"] {
		t.Fatalf("source-supported existing page should be allowed: %v", allowed)
	}
	if allowed["unrelated"] {
		t.Fatalf("source-unsupported existing page should not be allowed: %v", allowed)
	}
}

func TestRunGenerateCandidatesRetriesWithFallbackWhenCleanFails(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}
	for _, path := range []string{"drafts/bundle", "prompts"} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	if err := os.WriteFile("drafts/bundle/source.md", []byte("---\nkind: source\nsources:\n  - raw/llm-wiki.md\n---\n\n# Source Summary\n"), 0o644); err != nil {
		t.Fatalf("write source draft: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/apply-plan.md", []byte("## Candidate Changes\n\n- `wiki/concepts/tooling-stack.md` — create new page.\n"), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}
	if err := os.WriteFile("prompts/concept-page.md", []byte("concept {{PAGE_TITLE}}"), 0o644); err != nil {
		t.Fatalf("write concept prompt: %v", err)
	}
	fake := filepath.Join(dir, "llama-cli")
	attempts := filepath.Join(dir, "fallback-attempts")
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
attempts="$(cat "$NEMO_TEST_FALLBACK_ATTEMPTS" 2>/dev/null || true)"
if [ -z "$attempts" ]; then
  echo 1 > "$NEMO_TEST_FALLBACK_ATTEMPTS"
  echo "thinking forever without markdown"
  exit 0
fi
case "$*" in
  *"--temp 0.2"*) ;;
  *) echo "fallback profile not used: $*" >&2; exit 1 ;;
esac
cat <<'EOF'
---
title: Tooling Stack
kind: concept
sources:
  - raw/llm-wiki.md
---

# Tooling Stack

Generated by fallback.
EOF
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}
	t.Setenv("NEMO_TEST_FALLBACK_ATTEMPTS", attempts)
	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-generate-candidates", "drafts/bundle", "-profile", "stable"}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}
	candidate, err := os.ReadFile("drafts/bundle/candidates/wiki/concepts/tooling-stack.md")
	if err != nil {
		t.Fatalf("read candidate draft: %v", err)
	}
	if !strings.Contains(string(candidate), "Generated by fallback.") {
		t.Fatalf("candidate did not use fallback output:\n%s", candidate)
	}
	if _, err := os.Stat("drafts/bundle/candidates/wiki/concepts/tooling-stack.fallback.raw.txt"); err != nil {
		t.Fatalf("expected fallback raw output: %v", err)
	}
}

func TestRunGenerateCandidatesFailsWhenPrimaryAndFallbackCannotBeCleaned(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}
	for _, path := range []string{"drafts/bundle", "prompts"} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	if err := os.WriteFile("drafts/bundle/source.md", []byte("---\nkind: source\nsources:\n  - raw/llm-wiki.md\n---\n\n# Source Summary\n"), 0o644); err != nil {
		t.Fatalf("write source draft: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/apply-plan.md", []byte("## Candidate Changes\n\n- `wiki/concepts/tooling-stack.md` — create new page.\n"), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}
	if err := os.WriteFile("prompts/concept-page.md", []byte("concept {{PAGE_TITLE}}"), 0o644); err != nil {
		t.Fatalf("write concept prompt: %v", err)
	}
	fake := filepath.Join(dir, "llama-cli")
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
echo "thinking forever without markdown"
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}
	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-generate-candidates", "drafts/bundle", "-profile", "stable"}); code == 0 {
		t.Fatal("expected candidate generation to fail when primary and fallback are uncleanable")
	}
	candidatePath := "drafts/bundle/candidates/wiki/concepts/tooling-stack.md"
	if _, err := os.Stat(candidatePath); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected no deterministic boilerplate candidate at %s, stat err=%v", candidatePath, err)
	}
}

func TestTargetEvidenceExcerptsFindTitleTermsInRawSource(t *testing.T) {
	content := strings.Repeat("opening sea material\n", 20) +
		"Queequeg shares a room with Ishmael and becomes central to the novel's cross-cultural friendship.\n" +
		strings.Repeat("middle sea material\n", 20) +
		"The whiteness of the whale gives Ishmael a language for terror and ambiguity.\n"

	excerpts := targetEvidenceExcerpts(content, "Queequeg And Cosmopolitanism", 2, 300)
	if len(excerpts) == 0 {
		t.Fatal("expected target evidence for Queequeg")
	}
	if !strings.Contains(excerpts[0], "Queequeg") {
		t.Fatalf("excerpt does not contain target term:\n%s", excerpts[0])
	}
}

func TestRunGenerateCandidatesRemovesNestedFrontmatterFromBody(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}
	for _, path := range []string{"drafts/bundle", "prompts"} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	if err := os.WriteFile("drafts/bundle/source.md", []byte("---\nkind: source\nsources:\n  - raw/llm-wiki.md\n---\n\n# Source Summary\n"), 0o644); err != nil {
		t.Fatalf("write source draft: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/apply-plan.md", []byte("## Candidate Changes\n\n- `wiki/topics/ingest-workflow.md` — create new page.\n"), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}
	if err := os.WriteFile("prompts/topic-page.md", []byte("topic {{PAGE_TITLE}}"), 0o644); err != nil {
		t.Fatalf("write topic prompt: %v", err)
	}
	fake := filepath.Join(dir, "llama-cli")
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
cat <<'EOF'
---
title: Ingest Workflow
kind: topic
sources:
  - raw/llm-wiki.md
---

title: Ingest Workflow
kind: topic
---

# Ingest Workflow

Generated body with [[RAG]] and [[Query]] links.
EOF
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}
	t.Setenv("NEMO_LLAMA_CLI", fake)
	t.Setenv("NEMO_LLAMA_MODEL", "model.gguf")

	if code := run([]string{"-generate-candidates", "drafts/bundle"}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}
	candidate, err := os.ReadFile("drafts/bundle/candidates/wiki/topics/ingest-workflow.md")
	if err != nil {
		t.Fatalf("read candidate draft: %v", err)
	}
	content := string(candidate)
	if strings.Contains(content, "\ntitle: Ingest Workflow\nkind: topic\n---") {
		t.Fatalf("nested frontmatter was not removed:\n%s", content)
	}
	if strings.Contains(content, "[[RAG]]") {
		t.Fatalf("unknown wikilink should be normalized to plain text:\n%s", content)
	}
}

func TestRunApplyApprovedAppliesBundle(t *testing.T) {
	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}

	for _, path := range []string{
		"wiki/sources",
		"wiki/concepts",
		"drafts/bundle/candidates/wiki/concepts",
	} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	if err := os.WriteFile("wiki/index.md", []byte("---\ntitle: Index\nkind: index\n---\n\n## Sources\n- [[llm-wiki]] — Existing source.\n"), 0o644); err != nil {
		t.Fatalf("write index: %v", err)
	}
	if err := os.WriteFile("wiki/log.md", []byte("# Log\n"), 0o644); err != nil {
		t.Fatalf("write log: %v", err)
	}
	if err := os.WriteFile("wiki/sources/llm-wiki.md", []byte("---\nkind: source\n---\n# Old\n"), 0o644); err != nil {
		t.Fatalf("write existing source: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/source.md", []byte("---\nkind: source\n---\n# Applied\n"), 0o644); err != nil {
		t.Fatalf("write source draft: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/apply-plan.md", []byte("- `wiki/sources/llm-wiki-pattern.md` — create new page; possible duplicate of `wiki/sources/llm-wiki.md`.\n- `wiki/concepts/llm-maintenance-pattern.md` — create new page.\n"), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/candidates/wiki/concepts/llm-maintenance-pattern.md", []byte("---\ntitle: LLM Maintenance Pattern\nkind: concept\nsources:\n  - raw/llm-wiki.md\n---\n\n# LLM Maintenance Pattern\n\nReviewed concept.\n"), 0o644); err != nil {
		t.Fatalf("write concept draft: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/scores.json", []byte("{\"scores\":{\"overall\":\"pass\"}}\n"), 0o644); err != nil {
		t.Fatalf("write scores: %v", err)
	}

	if code := run([]string{"-apply-approved", "drafts/bundle", "-approve"}); code != 0 {
		t.Fatalf("run returned exit code %d", code)
	}
	applied, err := os.ReadFile("wiki/sources/llm-wiki.md")
	if err != nil {
		t.Fatalf("read applied source: %v", err)
	}
	if !strings.Contains(string(applied), "# Applied") {
		t.Fatalf("source was not applied:\n%s", applied)
	}
	concept, err := os.ReadFile("wiki/concepts/llm-maintenance-pattern.md")
	if err != nil {
		t.Fatalf("read applied concept: %v", err)
	}
	if !strings.Contains(string(concept), "Reviewed concept.") {
		t.Fatalf("concept was not applied:\n%s", concept)
	}
	index, err := os.ReadFile("wiki/index.md")
	if err != nil {
		t.Fatalf("read index: %v", err)
	}
	if !strings.Contains(string(index), "[[llm-maintenance-pattern]]") {
		t.Fatalf("index was not updated:\n%s", index)
	}
	if _, err := os.Stat("drafts/bundle/apply-report.md"); err != nil {
		t.Fatalf("expected apply report: %v", err)
	}
	if code := run([]string{"-lint-wiki", "-out-dir", "evals/apply-lint"}); code != 0 {
		t.Fatalf("post-apply lint returned exit code %d", code)
	}
	if _, err := os.Stat("evals/apply-lint/wiki-lint.json"); err != nil {
		t.Fatalf("expected post-apply lint output: %v", err)
	}
}

func TestRunApplyApprovedSupportsForceApply(t *testing.T) {
	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get wd: %v", err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Fatalf("restore wd: %v", err)
		}
	}()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir temp repo: %v", err)
	}

	for _, path := range []string{"wiki/sources", "drafts/bundle"} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", path, err)
		}
	}
	if err := os.WriteFile("wiki/log.md", []byte("# Log\n\n## [2026-05-16] ingest | drafts/bundle\n"), 0o644); err != nil {
		t.Fatalf("write log: %v", err)
	}
	if err := os.WriteFile("wiki/sources/llm-wiki.md", []byte("---\nkind: source\n---\n# Old\n"), 0o644); err != nil {
		t.Fatalf("write existing source: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/source.md", []byte("---\nkind: source\n---\n# Applied\n"), 0o644); err != nil {
		t.Fatalf("write source draft: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/apply-plan.md", []byte("- `wiki/sources/llm-wiki-pattern.md` — create new page; possible duplicate of `wiki/sources/llm-wiki.md`.\n"), 0o644); err != nil {
		t.Fatalf("write apply plan: %v", err)
	}
	if err := os.WriteFile("drafts/bundle/scores.json", []byte("{\"scores\":{\"overall\":\"pass\"}}\n"), 0o644); err != nil {
		t.Fatalf("write scores: %v", err)
	}

	if code := run([]string{"-apply-approved", "drafts/bundle", "-approve"}); code == 0 {
		t.Fatal("expected duplicate apply to fail without force")
	}
	if code := run([]string{"-apply-approved", "drafts/bundle", "-approve", "-force-apply"}); code != 0 {
		t.Fatalf("force apply returned exit code %d", code)
	}
}
