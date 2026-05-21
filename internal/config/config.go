package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	defaultLlamaCLI          = "/home/karo/src/llama.cpp/build/bin/llama-cli"
	defaultLlamaModel        = "/home/karo/models/qwen3.5-9b-q4_k_m.gguf"
	defaultProvider          = "llama"
	defaultDeepSeekBaseURL   = "https://api.deepseek.com"
	defaultDeepSeekModel     = "deepseek-v4-pro"
	defaultDeepSeekMaxTokens = 384000
	defaultGPULayers         = "all"
	defaultMaxTokens         = 8192
	defaultTemp              = 0.2
	defaultTopP              = 0.9
	defaultTopK              = 20
	defaultMinP              = 0.0

	// Local llama.cpp limits at the 24576-token context window: empirically a
	// single-shot prompt above ~90k source characters starts truncating frontmatter
	// or dropping mid-document detail. Above this size the bundle pipeline switches
	// to the structure-aware chunked path.
	defaultLlamaChunkedBundleCharThreshold = 90000
	defaultLlamaMaxChunkChars              = 18000

	// DeepSeek-V4 input window is 128K tokens (~460K ASCII chars). The pipeline
	// keeps a >50% safety margin for prompt scaffolding and thinking-mode
	// reasoning tokens, so single-shot is allowed up to ~300k source chars.
	// When chunked path still triggers, each chunk can be larger because the
	// model is far less sensitive to long prompts than local llama.cpp.
	defaultDeepSeekChunkedBundleCharThreshold = 300000
	defaultDeepSeekMaxChunkChars              = 60000
)

type Config struct {
	Profile                string
	Provider               string
	LlamaCLI               string
	LlamaModel             string
	DeepSeek               DeepSeekConfig
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

	// ChunkedBundleCharThreshold is the raw source size (in characters) above
	// which runBundle switches to the chunked, multi-stage synthesis path.
	// The default depends on the provider because local llama.cpp degrades on
	// long prompts far sooner than DeepSeek's 128K-token input window.
	ChunkedBundleCharThreshold int
	// MaxChunkChars caps a single chunk's size in the chunker. Lower values
	// produce more chunks (more API calls, finer audit granularity); higher
	// values keep each chunk's context tighter for capable models.
	MaxChunkChars int
}

type DeepSeekConfig struct {
	BaseURL         string
	APIKey          string
	Model           string
	MaxTokens       int
	Thinking        string
	ReasoningEffort string
	Temperature     *float64
	TopP            *float64
	ResponseFormat  string
	UserID          string
	SystemPrompt    string
}

// Defaults returns the local default configuration for the nemo command.
//
// The returned configuration points at the verified local llama.cpp CUDA
// binary and GGUF model path by default.
//
// NEMO_MODEL_PROVIDER selects the generation backend ("llama" or "deepseek").
// NEMO_LLAMA_CLI overrides the llama.cpp executable path.
// NEMO_LLAMA_MODEL overrides the GGUF model path.
// NEMO_DEEPSEEK_API_KEY supplies the DeepSeek API key when using the DeepSeek backend.
// NEMO_DEEPSEEK_BASE_URL overrides the OpenAI-compatible DeepSeek base URL.
// NEMO_DEEPSEEK_MODEL overrides the DeepSeek model id.
// NEMO_DEEPSEEK_MAX_TOKENS overrides the DeepSeek output token budget.
// NEMO_DEEPSEEK_TEMPERATURE overrides the DeepSeek non-thinking temperature.
// NEMO_DEEPSEEK_TOP_P overrides the DeepSeek non-thinking top_p.
// NEMO_DEEPSEEK_RESPONSE_FORMAT can be "text" or "json_object".
// NEMO_DEEPSEEK_USER_ID sets DeepSeek's cache-isolation user_id.
// NEMO_DEEPSEEK_SYSTEM_PROMPT sends a system message before the rendered prompt.
// NEMO_MAX_TOKENS overrides the generation token budget.
// NEMO_CHUNKED_THRESHOLD_CHARS overrides the source-size threshold that triggers
// the chunked bundle path. NEMO_MAX_CHUNK_CHARS overrides the per-chunk size cap
// used by the chunker. Both default to provider-appropriate values.
func Defaults() Config {
	loadDotEnv(".env")

	cfg := Config{
		Profile:    "fast",
		Provider:   defaultProvider,
		LlamaCLI:   defaultLlamaCLI,
		LlamaModel: defaultLlamaModel,
		DeepSeek: DeepSeekConfig{
			BaseURL:         defaultDeepSeekBaseURL,
			Model:           defaultDeepSeekModel,
			MaxTokens:       defaultDeepSeekMaxTokens,
			Thinking:        "enabled",
			ReasoningEffort: "high",
			ResponseFormat:  "text",
		},
		GPULayers:                  defaultGPULayers,
		MaxTokens:                  defaultMaxTokens,
		Temp:                       defaultTemp,
		TopP:                       defaultTopP,
		TopK:                       defaultTopK,
		MinP:                       defaultMinP,
		RepeatPenalty:              1.0,
		Jinja:                      true,
		NoContextShift:             true,
		ChunkedBundleCharThreshold: defaultLlamaChunkedBundleCharThreshold,
		MaxChunkChars:              defaultLlamaMaxChunkChars,
	}

	if value := os.Getenv("NEMO_MODEL_PROVIDER"); value != "" {
		cfg.Provider = strings.ToLower(value)
	}
	applyProviderChunkDefaults(&cfg)
	if value := os.Getenv("NEMO_LLAMA_CLI"); value != "" {
		cfg.LlamaCLI = value
	}
	if value := os.Getenv("NEMO_LLAMA_MODEL"); value != "" {
		cfg.LlamaModel = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_API_KEY"); value != "" {
		cfg.DeepSeek.APIKey = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_BASE_URL"); value != "" {
		cfg.DeepSeek.BaseURL = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_MODEL"); value != "" {
		cfg.DeepSeek.Model = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_MAX_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.DeepSeek.MaxTokens = parsed
		}
	}
	if value := os.Getenv("NEMO_DEEPSEEK_THINKING"); value != "" {
		cfg.DeepSeek.Thinking = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_REASONING_EFFORT"); value != "" {
		cfg.DeepSeek.ReasoningEffort = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_TEMPERATURE"); value != "" {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			cfg.DeepSeek.Temperature = &parsed
		}
	}
	if value := os.Getenv("NEMO_DEEPSEEK_TOP_P"); value != "" {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			cfg.DeepSeek.TopP = &parsed
		}
	}
	if value := os.Getenv("NEMO_DEEPSEEK_RESPONSE_FORMAT"); value != "" {
		cfg.DeepSeek.ResponseFormat = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_USER_ID"); value != "" {
		cfg.DeepSeek.UserID = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_SYSTEM_PROMPT"); value != "" {
		cfg.DeepSeek.SystemPrompt = value
	}
	if value := os.Getenv("NEMO_MAX_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.MaxTokens = parsed
		}
	}
	applyChunkEnvOverrides(&cfg)

	return cfg
}

// ForProfile returns local generation settings tuned for a named task profile.
//
// Profiles keep model parameters explicit at the call site while preserving the
// verified local llama.cpp binary and GGUF model defaults.
func ForProfile(profile string) (Config, error) {
	return ForProfileWithProvider(profile, "")
}

// ForProfileWithProvider returns settings for profile and, when provider is
// non-empty, locks the generation backend after loading environment defaults.
// This gives CLI callers a stable per-invocation override that wins over .env.
func ForProfileWithProvider(profile string, provider string) (Config, error) {
	cfg := Defaults()
	if err := ApplyProviderOverride(&cfg, provider); err != nil {
		return Config{}, err
	}

	switch profile {
	case "", "fast":
		cfg.Profile = "fast"
		cfg.MaxTokens = 2048
		cfg.Temp = 0.2
		cfg.TopP = 0.9
		cfg.TopK = 20
		cfg.MinP = 0
		deepSeekTemp := 0.2
		deepSeekTopP := 0.9
		cfg.DeepSeek.Model = "deepseek-v4-flash"
		cfg.DeepSeek.MaxTokens = defaultDeepSeekMaxTokens
		cfg.DeepSeek.Thinking = "disabled"
		cfg.DeepSeek.ReasoningEffort = ""
		cfg.DeepSeek.Temperature = &deepSeekTemp
		cfg.DeepSeek.TopP = &deepSeekTopP
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
		cfg.DeepSeek.Model = "deepseek-v4-pro"
		cfg.DeepSeek.MaxTokens = defaultDeepSeekMaxTokens
		cfg.DeepSeek.Thinking = "enabled"
		cfg.DeepSeek.ReasoningEffort = "high"
		cfg.DeepSeek.Temperature = nil
		cfg.DeepSeek.TopP = nil
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
		cfg.DeepSeek.Model = "deepseek-v4-pro"
		cfg.DeepSeek.MaxTokens = defaultDeepSeekMaxTokens
		cfg.DeepSeek.Thinking = "enabled"
		cfg.DeepSeek.ReasoningEffort = "high"
		cfg.DeepSeek.Temperature = nil
		cfg.DeepSeek.TopP = nil
	case "fallback":
		cfg.Profile = "fallback"
		cfg.MaxTokens = 16384
		cfg.CtxSize = 24576
		cfg.Temp = 0.2
		cfg.TopP = 0.8
		cfg.TopK = 20
		cfg.MinP = 0
		cfg.PresencePenalty = 1.5
		deepSeekTemp := 0.2
		deepSeekTopP := 0.8
		cfg.DeepSeek.Model = "deepseek-v4-flash"
		cfg.DeepSeek.MaxTokens = defaultDeepSeekMaxTokens
		cfg.DeepSeek.Thinking = "disabled"
		cfg.DeepSeek.ReasoningEffort = ""
		cfg.DeepSeek.Temperature = &deepSeekTemp
		cfg.DeepSeek.TopP = &deepSeekTopP
		applyNonThinkingDefaults(&cfg)
	default:
		return Config{}, fmt.Errorf("unknown profile %q", profile)
	}

	if cfg.Provider != "llama" && cfg.Provider != "deepseek" {
		return Config{}, fmt.Errorf("unknown model provider %q", cfg.Provider)
	}
	if value := os.Getenv("NEMO_DEEPSEEK_MODEL"); value != "" {
		cfg.DeepSeek.Model = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_MAX_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.DeepSeek.MaxTokens = parsed
		}
	}
	if value := os.Getenv("NEMO_DEEPSEEK_THINKING"); value != "" {
		cfg.DeepSeek.Thinking = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_REASONING_EFFORT"); value != "" {
		cfg.DeepSeek.ReasoningEffort = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_TEMPERATURE"); value != "" {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			cfg.DeepSeek.Temperature = &parsed
		}
	}
	if value := os.Getenv("NEMO_DEEPSEEK_TOP_P"); value != "" {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			cfg.DeepSeek.TopP = &parsed
		}
	}
	if value := os.Getenv("NEMO_DEEPSEEK_RESPONSE_FORMAT"); value != "" {
		cfg.DeepSeek.ResponseFormat = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_USER_ID"); value != "" {
		cfg.DeepSeek.UserID = value
	}
	if value := os.Getenv("NEMO_DEEPSEEK_SYSTEM_PROMPT"); value != "" {
		cfg.DeepSeek.SystemPrompt = value
	}
	if value := os.Getenv("NEMO_MAX_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.MaxTokens = parsed
		}
	}
	applyChunkEnvOverrides(&cfg)

	return cfg, nil
}

// ApplyProviderOverride locks cfg to provider when provider is non-empty.
// Provider-specific chunk defaults are recomputed before explicit chunk env
// overrides are re-applied, preserving NEMO_CHUNKED_THRESHOLD_CHARS and
// NEMO_MAX_CHUNK_CHARS as the final word.
func ApplyProviderOverride(cfg *Config, provider string) error {
	provider = strings.ToLower(strings.TrimSpace(provider))
	if provider == "" {
		return nil
	}
	if provider != "llama" && provider != "deepseek" {
		return fmt.Errorf("unknown model provider %q", provider)
	}
	cfg.Provider = provider
	applyProviderChunkDefaults(cfg)
	applyChunkEnvOverrides(cfg)
	return nil
}

// applyProviderChunkDefaults selects chunked-bundle thresholds based on the
// active provider. DeepSeek's 128K input window is much larger than local
// llama.cpp's 24576-token context, so the chunked path triggers far later and
// each chunk can be larger.
func applyProviderChunkDefaults(cfg *Config) {
	switch cfg.Provider {
	case "deepseek":
		cfg.ChunkedBundleCharThreshold = defaultDeepSeekChunkedBundleCharThreshold
		cfg.MaxChunkChars = defaultDeepSeekMaxChunkChars
	default:
		cfg.ChunkedBundleCharThreshold = defaultLlamaChunkedBundleCharThreshold
		cfg.MaxChunkChars = defaultLlamaMaxChunkChars
	}
}

func applyChunkEnvOverrides(cfg *Config) {
	if value := os.Getenv("NEMO_CHUNKED_THRESHOLD_CHARS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.ChunkedBundleCharThreshold = parsed
		}
	}
	if value := os.Getenv("NEMO_MAX_CHUNK_CHARS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.MaxChunkChars = parsed
		}
	}
}

func applyNonThinkingDefaults(cfg *Config) {
	cfg.Reasoning = "off"
	budget := 0
	cfg.ReasoningBudget = &budget
	cfg.ChatTemplateKwargs = `{"enable_thinking":false}`
}

func loadDotEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimSpace(strings.TrimPrefix(line, "export "))
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		if key == "" || strings.ContainsAny(key, " \t") {
			continue
		}
		if _, exists := os.LookupEnv(key); exists {
			continue
		}
		_ = os.Setenv(key, trimDotEnvValue(value))
	}
}

func trimDotEnvValue(value string) string {
	value = strings.TrimSpace(value)
	if len(value) < 2 {
		return value
	}
	quote := value[0]
	if quote != '"' && quote != '\'' {
		return value
	}
	if value[len(value)-1] != quote {
		return value
	}
	return value[1 : len(value)-1]
}
