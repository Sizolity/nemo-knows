package llama

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
)

var ErrNotImplemented = errors.New("llama CLI generation is not implemented")

// Generator produces text from a rendered prompt.
//
// Implementations may call a local CLI process, an HTTP model server, or a
// future orchestration layer.
type Generator interface {
	// Generate sends prompt to the model backend and returns the raw model output.
	//
	// The ctx controls cancellation and timeouts for the generation call.
	//
	// The returned string is intentionally raw so the draft layer can preserve
	// and clean llama.cpp output separately.
	Generate(ctx context.Context, prompt string) (string, error)
}

type CLI struct {
	Binary                 string
	Model                  string
	GPULayers              string
	MaxTokens              int
	CtxSize                int
	Temp                   float64
	TopP                   float64
	TopK                   int
	MinP                   float64
	PresencePenalty        float64
	RepeatPenalty          float64
	Reasoning              string
	ReasoningBudget        *int
	ReasoningBudgetMessage string
	ChatTemplateKwargs     string
	Jinja                  bool
	NoContextShift         bool
}

// Generate invokes a llama.cpp CLI backend with the configured model settings.
//
// The prompt argument is the fully rendered prompt text.
//
// The returned string should include the raw combined output from llama.cpp so
// callers can store it for debugging before cleaning it into Markdown.
//
// The error is returned if the command cannot start, exits unsuccessfully, or
// the context is cancelled.
func (c CLI) Generate(ctx context.Context, prompt string) (string, error) {
	if c.Binary == "" {
		return "", errors.New("llama binary path is required")
	}
	if c.Model == "" {
		return "", errors.New("llama model path is required")
	}

	maxTokens := c.MaxTokens
	if maxTokens == 0 {
		maxTokens = 2048
	}
	gpuLayers := c.GPULayers
	if gpuLayers == "" {
		gpuLayers = "all"
	}
	temp := c.Temp
	topP := c.TopP
	topK := c.TopK
	if topK == 0 {
		topK = 20
	}
	minP := c.MinP

	args := []string{
		"-m", c.Model,
		"-p", prompt,
		"-n", strconv.Itoa(maxTokens),
		"-ngl", gpuLayers,
		"--single-turn",
		"--simple-io",
		"--no-display-prompt",
		"--temp", strconv.FormatFloat(temp, 'f', -1, 64),
		"--top-p", strconv.FormatFloat(topP, 'f', -1, 64),
		"--top-k", strconv.Itoa(topK),
		"--min-p", strconv.FormatFloat(minP, 'f', -1, 64),
	}
	if c.CtxSize > 0 {
		args = append(args, "-c", strconv.Itoa(c.CtxSize))
	}
	if c.PresencePenalty != 0 {
		args = append(args, "--presence-penalty", strconv.FormatFloat(c.PresencePenalty, 'f', -1, 64))
	}
	if c.RepeatPenalty != 0 {
		args = append(args, "--repeat-penalty", strconv.FormatFloat(c.RepeatPenalty, 'f', -1, 64))
	}
	if c.Reasoning != "" {
		args = append(args, "--reasoning", c.Reasoning)
	}
	if c.ReasoningBudget != nil {
		args = append(args, "--reasoning-budget", strconv.Itoa(*c.ReasoningBudget))
	}
	if c.ReasoningBudgetMessage != "" {
		args = append(args, "--reasoning-budget-message", c.ReasoningBudgetMessage)
	}
	if c.ChatTemplateKwargs != "" {
		args = append(args, "--chat-template-kwargs", c.ChatTemplateKwargs)
	}
	if c.Jinja {
		args = append(args, "--jinja")
	}
	if c.NoContextShift {
		args = append(args, "--no-context-shift")
	}

	cmd := exec.CommandContext(ctx, c.Binary, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("run llama CLI: %w", err)
	}

	return string(output), nil
}
