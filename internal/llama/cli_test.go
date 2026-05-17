package llama

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestCLIGenerateInvokesConfiguredBinary(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	fake := filepath.Join(dir, "llama-cli")
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
printf 'args:%s\n' "$*"
printf 'prompt:%s\n' "$6"
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}

	out, err := CLI{
		Binary:    fake,
		Model:     "model.gguf",
		GPULayers: "all",
		MaxTokens: 64,
		Temp:      0.2,
		TopP:      0.9,
		TopK:      20,
		MinP:      0,
	}.Generate(context.Background(), "hello")
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	for _, want := range []string{
		"-m model.gguf",
		"-p hello",
		"-n 64",
		"-ngl all",
		"--single-turn",
		"--simple-io",
		"--no-display-prompt",
		"--temp 0.2",
		"--top-p 0.9",
		"--top-k 20",
		"--min-p 0",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("output missing %q:\n%s", want, out)
		}
	}
}

func TestCLIGeneratePassesReasoningAndPenaltyControls(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script fake is Unix-specific")
	}

	dir := t.TempDir()
	fake := filepath.Join(dir, "llama-cli")
	if err := os.WriteFile(fake, []byte(`#!/usr/bin/env sh
printf 'args:%s\n' "$*"
`), 0o755); err != nil {
		t.Fatalf("write fake llama: %v", err)
	}
	reasoningBudget := 0

	out, err := CLI{
		Binary:                 fake,
		Model:                  "model.gguf",
		GPULayers:              "all",
		MaxTokens:              64,
		CtxSize:                40960,
		Temp:                   0.7,
		TopP:                   0.8,
		TopK:                   20,
		MinP:                   0,
		PresencePenalty:        1.5,
		RepeatPenalty:          1.0,
		Reasoning:              "off",
		ReasoningBudget:        &reasoningBudget,
		ChatTemplateKwargs:     `{"enable_thinking":false}`,
		Jinja:                  true,
		NoContextShift:         true,
		ReasoningBudgetMessage: "reasoning budget exceeded, now write the final Markdown",
	}.Generate(context.Background(), "hello")
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	for _, want := range []string{
		"-c 40960",
		"--presence-penalty 1.5",
		"--repeat-penalty 1",
		"--reasoning off",
		"--reasoning-budget 0",
		"--reasoning-budget-message reasoning budget exceeded, now write the final Markdown",
		"--chat-template-kwargs {\"enable_thinking\":false}",
		"--jinja",
		"--no-context-shift",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("output missing %q:\n%s", want, out)
		}
	}
}

func TestCLIGenerateErrorsOnMissingBinary(t *testing.T) {
	_, err := CLI{Binary: "does-not-exist"}.Generate(context.Background(), "hello")
	if err == nil {
		t.Fatal("expected error for missing binary")
	}
}
