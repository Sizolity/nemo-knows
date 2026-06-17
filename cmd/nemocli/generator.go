package main

import (
	"time"

	"github.com/huic/nemo-knows/internal/config"
	"github.com/huic/nemo-knows/internal/deepseek"
	"github.com/huic/nemo-knows/internal/llama"
)

// generatorFromCfg mirrors cmd/nemo's generatorFromConfig — it returns
// the model backend that internal/wikimaint expects for propose/auto
// modes. Keeping the mapping local to cmd/nemocli avoids any need to
// touch the existing cmd/nemo code while staying byte-compatible with
// the same configuration knobs.
func generatorFromCfg(cfg config.Config) llama.Generator {
	if cfg.Provider == "deepseek" {
		return deepseek.Client{
			BaseURL:         cfg.DeepSeek.BaseURL,
			APIKey:          cfg.DeepSeek.APIKey,
			Model:           cfg.DeepSeek.Model,
			MaxTokens:       cfg.DeepSeek.MaxTokens,
			Temperature:     cfg.DeepSeek.Temperature,
			TopP:            cfg.DeepSeek.TopP,
			Thinking:        cfg.DeepSeek.Thinking,
			ReasoningEffort: cfg.DeepSeek.ReasoningEffort,
			ResponseFormat:  cfg.DeepSeek.ResponseFormat,
			UserID:          cfg.DeepSeek.UserID,
			SystemPrompt:    cfg.DeepSeek.SystemPrompt,
			RetryMax:        cfg.DeepSeek.RetryMax,
			RetryBaseDelay:  time.Duration(cfg.DeepSeek.RetryBaseDelayMS) * time.Millisecond,
		}
	}
	return llama.CLI{
		Binary:                 cfg.LlamaCLI,
		Model:                  cfg.LlamaModel,
		GPULayers:              cfg.GPULayers,
		MaxTokens:              cfg.MaxTokens,
		CtxSize:                cfg.CtxSize,
		Temp:                   cfg.Temp,
		TopP:                   cfg.TopP,
		TopK:                   cfg.TopK,
		MinP:                   cfg.MinP,
		PresencePenalty:        cfg.PresencePenalty,
		RepeatPenalty:          cfg.RepeatPenalty,
		Reasoning:              cfg.Reasoning,
		ReasoningBudget:        cfg.ReasoningBudget,
		ReasoningBudgetMessage: cfg.ReasoningBudgetMessage,
		ChatTemplateKwargs:     cfg.ChatTemplateKwargs,
		Jinja:                  cfg.Jinja,
		NoContextShift:         cfg.NoContextShift,
	}
}
