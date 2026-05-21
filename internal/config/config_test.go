package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// isolateConfigEnv clears DeepSeek- and provider-related environment variables
// so tests are not contaminated by the developer's shell or a project .env file.
// Anything the test wants to assert must be set explicitly inside the test.
func isolateConfigEnv(t *testing.T) {
	t.Helper()
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
		"NEMO_DEEPSEEK_RETRY_MAX",
		"NEMO_DEEPSEEK_RETRY_BASE_MS",
		"NEMO_MAX_TOKENS",
		"NEMO_MODEL_CONTEXT_TOKENS",
		"NEMO_CHARS_PER_TOKEN",
		"NEMO_CONTEXT_RESERVE_TOKENS",
		"NEMO_CONTEXT_OUTPUT_RESERVE_TOKENS",
		"NEMO_CONTEXT_SAFETY_MARGIN",
		"NEMO_CHUNKED_THRESHOLD_CHARS",
		"NEMO_MAX_CHUNK_CHARS",
	} {
		t.Setenv(key, "")
	}
}

func TestProfileAppliesQwenStableDefaults(t *testing.T) {
	isolateConfigEnv(t)
	cfg, err := ForProfile("stable")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.MaxTokens != 32768 {
		t.Fatalf("MaxTokens = %d, want 32768", cfg.MaxTokens)
	}
	if cfg.CtxSize != 24576 {
		t.Fatalf("CtxSize = %d, want 24576", cfg.CtxSize)
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
	isolateConfigEnv(t)
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
	isolateConfigEnv(t)
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
	if cfg.CtxSize != 24576 {
		t.Fatalf("CtxSize = %d, want 24576", cfg.CtxSize)
	}
	if cfg.ReasoningBudget == nil || *cfg.ReasoningBudget != 0 {
		t.Fatalf("ReasoningBudget = %v, want 0", cfg.ReasoningBudget)
	}
}

func TestDeepSeekProviderConfigFromEnvironment(t *testing.T) {
	isolateConfigEnv(t)
	t.Setenv("NEMO_MODEL_PROVIDER", "deepseek")
	t.Setenv("NEMO_DEEPSEEK_API_KEY", "test-key")
	t.Setenv("NEMO_DEEPSEEK_BASE_URL", "https://example.test")
	t.Setenv("NEMO_DEEPSEEK_MODEL", "deepseek-v4-pro")
	t.Setenv("NEMO_DEEPSEEK_MAX_TOKENS", "384000")
	t.Setenv("NEMO_DEEPSEEK_THINKING", "enabled")
	t.Setenv("NEMO_DEEPSEEK_REASONING_EFFORT", "high")
	t.Setenv("NEMO_DEEPSEEK_RESPONSE_FORMAT", "json_object")
	t.Setenv("NEMO_DEEPSEEK_USER_ID", "nemo-test")
	t.Setenv("NEMO_DEEPSEEK_SYSTEM_PROMPT", "Return JSON.")

	cfg, err := ForProfile("stable")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.Provider != "deepseek" {
		t.Fatalf("Provider = %q, want deepseek", cfg.Provider)
	}
	if cfg.DeepSeek.APIKey != "test-key" {
		t.Fatalf("DeepSeek.APIKey = %q, want test-key", cfg.DeepSeek.APIKey)
	}
	if cfg.DeepSeek.BaseURL != "https://example.test" {
		t.Fatalf("DeepSeek.BaseURL = %q, want https://example.test", cfg.DeepSeek.BaseURL)
	}
	if cfg.DeepSeek.Model != "deepseek-v4-pro" {
		t.Fatalf("DeepSeek.Model = %q, want deepseek-v4-pro", cfg.DeepSeek.Model)
	}
	if cfg.DeepSeek.MaxTokens != 384000 {
		t.Fatalf("DeepSeek.MaxTokens = %d, want 384000", cfg.DeepSeek.MaxTokens)
	}
	if cfg.DeepSeek.Thinking != "enabled" {
		t.Fatalf("DeepSeek.Thinking = %q, want enabled", cfg.DeepSeek.Thinking)
	}
	if cfg.DeepSeek.ReasoningEffort != "high" {
		t.Fatalf("DeepSeek.ReasoningEffort = %q, want high", cfg.DeepSeek.ReasoningEffort)
	}
	if cfg.DeepSeek.ResponseFormat != "json_object" {
		t.Fatalf("DeepSeek.ResponseFormat = %q, want json_object", cfg.DeepSeek.ResponseFormat)
	}
	if cfg.DeepSeek.UserID != "nemo-test" {
		t.Fatalf("DeepSeek.UserID = %q, want nemo-test", cfg.DeepSeek.UserID)
	}
	if cfg.DeepSeek.SystemPrompt != "Return JSON." {
		t.Fatalf("DeepSeek.SystemPrompt = %q, want Return JSON.", cfg.DeepSeek.SystemPrompt)
	}
	if cfg.DeepSeek.RetryMax != 2 {
		t.Fatalf("DeepSeek.RetryMax = %d, want 2", cfg.DeepSeek.RetryMax)
	}
	if cfg.DeepSeek.RetryBaseDelayMS != 1000 {
		t.Fatalf("DeepSeek.RetryBaseDelayMS = %d, want 1000", cfg.DeepSeek.RetryBaseDelayMS)
	}
}

func TestDeepSeekFastProfileUsesFlashModel(t *testing.T) {
	isolateConfigEnv(t)
	t.Setenv("NEMO_MODEL_PROVIDER", "deepseek")

	cfg, err := ForProfile("fast")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.DeepSeek.Model != "deepseek-v4-flash" {
		t.Fatalf("DeepSeek.Model = %q, want deepseek-v4-flash", cfg.DeepSeek.Model)
	}
	if cfg.DeepSeek.MaxTokens != 384000 {
		t.Fatalf("DeepSeek.MaxTokens = %d, want 384000", cfg.DeepSeek.MaxTokens)
	}
	if cfg.DeepSeek.Thinking != "disabled" {
		t.Fatalf("DeepSeek.Thinking = %q, want disabled", cfg.DeepSeek.Thinking)
	}
	if cfg.DeepSeek.ReasoningEffort != "" {
		t.Fatalf("DeepSeek.ReasoningEffort = %q, want empty", cfg.DeepSeek.ReasoningEffort)
	}
	if cfg.DeepSeek.Temperature == nil || *cfg.DeepSeek.Temperature != 0.2 {
		t.Fatalf("DeepSeek.Temperature = %v, want 0.2", cfg.DeepSeek.Temperature)
	}
	if cfg.DeepSeek.TopP == nil || *cfg.DeepSeek.TopP != 0.9 {
		t.Fatalf("DeepSeek.TopP = %v, want 0.9", cfg.DeepSeek.TopP)
	}
}

func TestDeepSeekThinkingProfileOmitsSamplingConfig(t *testing.T) {
	isolateConfigEnv(t)
	t.Setenv("NEMO_MODEL_PROVIDER", "deepseek")

	cfg, err := ForProfile("stable")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.DeepSeek.Temperature != nil {
		t.Fatalf("DeepSeek.Temperature = %v, want nil", *cfg.DeepSeek.Temperature)
	}
	if cfg.DeepSeek.TopP != nil {
		t.Fatalf("DeepSeek.TopP = %v, want nil", *cfg.DeepSeek.TopP)
	}
}

func TestUnknownProfileReturnsError(t *testing.T) {
	isolateConfigEnv(t)
	if _, err := ForProfile("unknown"); err == nil {
		t.Fatal("expected error for unknown profile")
	}
}

func TestUnknownProviderReturnsError(t *testing.T) {
	isolateConfigEnv(t)
	t.Setenv("NEMO_MODEL_PROVIDER", "unknown")

	if _, err := ForProfile("stable"); err == nil {
		t.Fatal("expected error for unknown provider")
	}
}

func TestLlamaProviderUsesConservativeChunkThresholds(t *testing.T) {
	isolateConfigEnv(t)
	t.Setenv("NEMO_MODEL_PROVIDER", "llama")

	cfg, err := ForProfile("stable")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.ChunkedBundleCharThreshold != 90000 {
		t.Fatalf("ChunkedBundleCharThreshold = %d, want 90000 for llama", cfg.ChunkedBundleCharThreshold)
	}
	if cfg.MaxChunkChars != 18000 {
		t.Fatalf("MaxChunkChars = %d, want 18000 for llama", cfg.MaxChunkChars)
	}
}

func TestDeepSeekProviderDerivesChunkThresholdFromModelContext(t *testing.T) {
	isolateConfigEnv(t)
	t.Setenv("NEMO_MODEL_PROVIDER", "deepseek")

	cfg, err := ForProfile("stable")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.ModelContextTokens != 1000000 {
		t.Fatalf("ModelContextTokens = %d, want 1000000 for deepseek", cfg.ModelContextTokens)
	}
	if cfg.ChunkedBundleCharThreshold != 1083600 {
		t.Fatalf("ChunkedBundleCharThreshold = %d, want model-derived 1083600 for deepseek", cfg.ChunkedBundleCharThreshold)
	}
	if cfg.MaxChunkChars != 60000 {
		t.Fatalf("MaxChunkChars = %d, want 60000 for deepseek", cfg.MaxChunkChars)
	}
}

func TestModelContextEnvOverridesDeriveChunkThreshold(t *testing.T) {
	isolateConfigEnv(t)
	t.Setenv("NEMO_MODEL_PROVIDER", "llama")
	t.Setenv("NEMO_MODEL_CONTEXT_TOKENS", "200000")
	t.Setenv("NEMO_CHARS_PER_TOKEN", "3")
	t.Setenv("NEMO_CONTEXT_RESERVE_TOKENS", "50000")
	t.Setenv("NEMO_CONTEXT_OUTPUT_RESERVE_TOKENS", "20000")
	t.Setenv("NEMO_CONTEXT_SAFETY_MARGIN", "0.5")

	cfg, err := ForProfile("stable")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.ChunkedBundleCharThreshold != 195000 {
		t.Fatalf("ChunkedBundleCharThreshold = %d, want 195000 from model context", cfg.ChunkedBundleCharThreshold)
	}
	if cfg.MaxChunkChars != 18000 {
		t.Fatalf("MaxChunkChars = %d, want llama default 18000", cfg.MaxChunkChars)
	}
}

func TestChunkEnvOverridesWinOverProviderDefaults(t *testing.T) {
	isolateConfigEnv(t)
	t.Setenv("NEMO_MODEL_PROVIDER", "deepseek")
	t.Setenv("NEMO_CHUNKED_THRESHOLD_CHARS", "150000")
	t.Setenv("NEMO_MAX_CHUNK_CHARS", "40000")

	cfg, err := ForProfile("stable")
	if err != nil {
		t.Fatalf("ForProfile returned error: %v", err)
	}

	if cfg.ChunkedBundleCharThreshold != 150000 {
		t.Fatalf("ChunkedBundleCharThreshold = %d, want 150000 from env override", cfg.ChunkedBundleCharThreshold)
	}
	if cfg.MaxChunkChars != 40000 {
		t.Fatalf("MaxChunkChars = %d, want 40000 from env override", cfg.MaxChunkChars)
	}
}

func TestForProfileWithProviderOverridesEnvironmentProvider(t *testing.T) {
	isolateConfigEnv(t)
	t.Setenv("NEMO_MODEL_PROVIDER", "deepseek")

	cfg, err := ForProfileWithProvider("stable", "llama")
	if err != nil {
		t.Fatalf("ForProfileWithProvider returned error: %v", err)
	}

	if cfg.Provider != "llama" {
		t.Fatalf("Provider = %q, want llama", cfg.Provider)
	}
	if cfg.ChunkedBundleCharThreshold != 90000 {
		t.Fatalf("ChunkedBundleCharThreshold = %d, want llama default 90000", cfg.ChunkedBundleCharThreshold)
	}
	if cfg.MaxChunkChars != 18000 {
		t.Fatalf("MaxChunkChars = %d, want llama default 18000", cfg.MaxChunkChars)
	}
}

func TestForProfileWithProviderRejectsUnknownProvider(t *testing.T) {
	isolateConfigEnv(t)

	if _, err := ForProfileWithProvider("stable", "other"); err == nil {
		t.Fatal("expected error for unknown provider override")
	}
}

func TestLoadDotEnvSetsUnsetValues(t *testing.T) {
	isolateConfigEnv(t)
	path := filepath.Join(t.TempDir(), ".env")
	if err := os.WriteFile(path, []byte(strings.Join([]string{
		"# local DeepSeek config",
		"NEMO_TEST_DOTENV_VALUE=from-file",
		"export NEMO_TEST_DOTENV_QUOTED=\"quoted value\"",
		"NEMO_TEST_DOTENV_EXISTING=from-file",
		"",
	}, "\n")), 0o600); err != nil {
		t.Fatalf("write .env: %v", err)
	}
	t.Setenv("NEMO_TEST_DOTENV_EXISTING", "from-env")
	t.Cleanup(func() {
		_ = os.Unsetenv("NEMO_TEST_DOTENV_VALUE")
		_ = os.Unsetenv("NEMO_TEST_DOTENV_QUOTED")
	})

	loadDotEnv(path)

	if got := os.Getenv("NEMO_TEST_DOTENV_VALUE"); got != "from-file" {
		t.Fatalf("NEMO_TEST_DOTENV_VALUE = %q, want from-file", got)
	}
	if got := os.Getenv("NEMO_TEST_DOTENV_QUOTED"); got != "quoted value" {
		t.Fatalf("NEMO_TEST_DOTENV_QUOTED = %q, want quoted value", got)
	}
	if got := os.Getenv("NEMO_TEST_DOTENV_EXISTING"); got != "from-env" {
		t.Fatalf("NEMO_TEST_DOTENV_EXISTING = %q, want from-env", got)
	}
}
