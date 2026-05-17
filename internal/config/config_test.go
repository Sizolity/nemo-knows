package config

import "testing"

func TestProfileAppliesQwenStableDefaults(t *testing.T) {
	cfg, err := ForProfile("stable")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.MaxTokens != 32768 {
		t.Fatalf("MaxTokens = %d, want 32768", cfg.MaxTokens)
	}
	if cfg.Temp != 0.7 {
		t.Fatalf("Temp = %v, want 0.7", cfg.Temp)
	}
	if cfg.TopP != 0.8 {
		t.Fatalf("TopP = %v, want 0.8", cfg.TopP)
	}
	if cfg.TopK != 20 {
		t.Fatalf("TopK = %d, want 20", cfg.TopK)
	}
	if cfg.MinP != 0 {
		t.Fatalf("MinP = %v, want 0", cfg.MinP)
	}
	if cfg.Reasoning != "off" {
		t.Fatalf("Reasoning = %q, want off", cfg.Reasoning)
	}
	if cfg.ReasoningBudget == nil || *cfg.ReasoningBudget != 0 {
		t.Fatalf("ReasoningBudget = %v, want 0", cfg.ReasoningBudget)
	}
	if cfg.ChatTemplateKwargs != `{"enable_thinking":false}` {
		t.Fatalf("ChatTemplateKwargs = %q, want enable_thinking false", cfg.ChatTemplateKwargs)
	}
	if cfg.PresencePenalty != 1.5 {
		t.Fatalf("PresencePenalty = %v, want 1.5", cfg.PresencePenalty)
	}
	if cfg.RepeatPenalty != 1.0 {
		t.Fatalf("RepeatPenalty = %v, want 1.0", cfg.RepeatPenalty)
	}
}

func TestProfileAppliesQwenDeepDefaults(t *testing.T) {
	cfg, err := ForProfile("deep")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.MaxTokens != 65536 {
		t.Fatalf("MaxTokens = %d, want 65536", cfg.MaxTokens)
	}
	if cfg.Temp != 0.6 {
		t.Fatalf("Temp = %v, want 0.6", cfg.Temp)
	}
	if cfg.TopP != 0.95 {
		t.Fatalf("TopP = %v, want 0.95", cfg.TopP)
	}
	if cfg.Reasoning != "on" {
		t.Fatalf("Reasoning = %q, want on", cfg.Reasoning)
	}
	if cfg.ReasoningBudget == nil || *cfg.ReasoningBudget != 2000 {
		t.Fatalf("ReasoningBudget = %v, want 2000", cfg.ReasoningBudget)
	}
	if cfg.ReasoningBudgetMessage == "" {
		t.Fatal("ReasoningBudgetMessage should be set for deep profile")
	}
}

func TestProfileAppliesQwenFallbackNonThinkingDefaults(t *testing.T) {
	cfg, err := ForProfile("fallback")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.Reasoning != "off" {
		t.Fatalf("Reasoning = %q, want off", cfg.Reasoning)
	}
	if cfg.MaxTokens != 16384 {
		t.Fatalf("MaxTokens = %d, want 16384", cfg.MaxTokens)
	}
	if cfg.ReasoningBudget == nil || *cfg.ReasoningBudget != 0 {
		t.Fatalf("ReasoningBudget = %v, want 0", cfg.ReasoningBudget)
	}
}

func TestUnknownProfileReturnsError(t *testing.T) {
	if _, err := ForProfile("unknown"); err == nil {
		t.Fatal("expected error for unknown profile")
	}
}
