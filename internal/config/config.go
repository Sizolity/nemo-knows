package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defaultLlamaCLI   = "/home/karo/src/llama.cpp/build/bin/llama-cli"
	defaultLlamaModel = "/home/karo/models/qwen3.5-9b-q4_k_m.gguf"
	defaultGPULayers  = "all"
	defaultMaxTokens  = 8192
	defaultTemp       = 0.2
	defaultTopP       = 0.9
	defaultTopK       = 20
	defaultMinP       = 0.0
)

type Config struct {
	Profile                string
	LlamaCLI               string
	LlamaModel             string
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

// Defaults returns the local default configuration for the nemo command.
//
// The returned configuration points at the verified local llama.cpp CUDA
// binary and GGUF model path by default.
//
// NEMO_LLAMA_CLI overrides the llama.cpp executable path.
// NEMO_LLAMA_MODEL overrides the GGUF model path.
// NEMO_MAX_TOKENS overrides the generation token budget.
func Defaults() Config {
	cfg := Config{
		Profile:        "fast",
		LlamaCLI:       defaultLlamaCLI,
		LlamaModel:     defaultLlamaModel,
		GPULayers:      defaultGPULayers,
		MaxTokens:      defaultMaxTokens,
		Temp:           defaultTemp,
		TopP:           defaultTopP,
		TopK:           defaultTopK,
		MinP:           defaultMinP,
		RepeatPenalty:  1.0,
		Jinja:          true,
		NoContextShift: true,
	}

	if value := os.Getenv("NEMO_LLAMA_CLI"); value != "" {
		cfg.LlamaCLI = value
	}
	if value := os.Getenv("NEMO_LLAMA_MODEL"); value != "" {
		cfg.LlamaModel = value
	}
	if value := os.Getenv("NEMO_MAX_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.MaxTokens = parsed
		}
	}

	return cfg
}

// ForProfile returns local generation settings tuned for a named task profile.
//
// Profiles keep model parameters explicit at the call site while preserving the
// verified local llama.cpp binary and GGUF model defaults.
func ForProfile(profile string) (Config, error) {
	cfg := Defaults()

	switch profile {
	case "", "fast":
		cfg.Profile = "fast"
		cfg.MaxTokens = 2048
		cfg.Temp = 0.2
		cfg.TopP = 0.9
		cfg.TopK = 20
		cfg.MinP = 0
		applyNonThinkingDefaults(&cfg)
	case "stable":
		cfg.Profile = "stable"
		cfg.MaxTokens = 32768
		cfg.CtxSize = 24576
		cfg.Temp = 0.7
		cfg.TopP = 0.8
		cfg.TopK = 20
		cfg.MinP = 0
		cfg.PresencePenalty = 1.5
		applyNonThinkingDefaults(&cfg)
	case "deep":
		cfg.Profile = "deep"
		cfg.MaxTokens = 65536
		cfg.Temp = 0.6
		cfg.TopP = 0.95
		cfg.TopK = 20
		cfg.MinP = 0
		cfg.Reasoning = "on"
		budget := 2000
		cfg.ReasoningBudget = &budget
		cfg.ReasoningBudgetMessage = "reasoning budget exceeded, now write the final Markdown"
	case "fallback":
		cfg.Profile = "fallback"
		cfg.MaxTokens = 16384
		cfg.CtxSize = 24576
		cfg.Temp = 0.2
		cfg.TopP = 0.8
		cfg.TopK = 20
		cfg.MinP = 0
		cfg.PresencePenalty = 1.5
		applyNonThinkingDefaults(&cfg)
	default:
		return Config{}, fmt.Errorf("unknown profile %q", profile)
	}

	if value := os.Getenv("NEMO_MAX_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.MaxTokens = parsed
		}
	}

	return cfg, nil
}

func applyNonThinkingDefaults(cfg *Config) {
	cfg.Reasoning = "off"
	budget := 0
	cfg.ReasoningBudget = &budget
	cfg.ChatTemplateKwargs = `{"enable_thinking":false}`
}
