package config

import (
	"bufio"
	"fmt"
	"math"
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

	// DeepSeek hosted models expose much larger contexts than the local
	// llama.cpp backend. The default chunk threshold is computed from these
	// explicit model-capability defaults, then can be overridden by env.
	defaultDeepSeekContextTokens = 1000000
	defaultDeepSeekCharsPerToken = 3.5
	// Reserve output/reasoning/system prompt space before considering raw source.
	defaultDeepSeekContextReserveTokens      = 100000
	defaultDeepSeekContextSafetyMargin       = 0.60
	defaultDeepSeekQualityChunkCharThreshold = 600000
	defaultDeepSeekMaxChunkChars             = 60000

	defaultDeepSeekRetryMax         = 2
	defaultDeepSeekRetryBaseDelayMS = 1000
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
	// The default may be provider-empirical or computed from model context
	// capability, with NEMO_CHUNKED_THRESHOLD_CHARS as the final override.
	ChunkedBundleCharThreshold int
	// MaxChunkChars caps a single chunk's size in the chunker. Lower values
	// produce more chunks (more API calls, finer audit granularity); higher
	// values keep each chunk's context tighter for capable models.
	MaxChunkChars int
	// ModelContextTokens describes the configured model's input window. When
	// positive, nemo can derive a chunk threshold from actual model capability.
	ModelContextTokens int
	// CharsPerToken estimates prompt characters per model token for threshold
	// calculation. It is intentionally conservative and can be overridden.
	CharsPerToken float64
	// ContextReserveTokens is held back for prompt scaffolding, output, and
	// reasoning overhead before raw-source capacity is estimated.
	ContextReserveTokens int
	// ContextOutputReserveTokens is held back for the requested generation
	// budget. When negative, nemo uses the provider's active max_tokens setting.
	ContextOutputReserveTokens int
	// ContextSafetyMargin is applied after reserve subtraction to keep long
	// single-shot prompts away from the hard model context limit.
	ContextSafetyMargin float64
	// QualityChunkCharThreshold is an optional source-size ceiling for quality,
	// independent of the hard model context budget. Positive values are folded
	// into the final chunk trigger with the model-derived threshold.
	QualityChunkCharThreshold int
}

type DeepSeekConfig struct {
	BaseURL          string
	APIKey           string
	Model            string
	MaxTokens        int
	Thinking         string
	ReasoningEffort  string
	Temperature      *float64
	TopP             *float64
	ResponseFormat   string
	UserID           string
	SystemPrompt     string
	RetryMax         int
	RetryBaseDelayMS int
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
// NEMO_DEEPSEEK_RETRY_MAX and NEMO_DEEPSEEK_RETRY_BASE_MS tune transient-error retries.
// NEMO_MAX_TOKENS overrides the generation token budget.
// NEMO_MODEL_CONTEXT_TOKENS, NEMO_CHARS_PER_TOKEN,
// NEMO_CONTEXT_RESERVE_TOKENS, NEMO_CONTEXT_OUTPUT_RESERVE_TOKENS, and
// NEMO_CONTEXT_SAFETY_MARGIN can derive the chunked-bundle threshold from a
// specific model's configured context window.
// NEMO_QUALITY_CHUNK_THRESHOLD_CHARS sets a quality-oriented ceiling that can
// force chunking before the model context is mechanically full.
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
			BaseURL:          defaultDeepSeekBaseURL,
			Model:            defaultDeepSeekModel,
			MaxTokens:        defaultDeepSeekMaxTokens,
			Thinking:         "enabled",
			ReasoningEffort:  "high",
			ResponseFormat:   "text",
			RetryMax:         defaultDeepSeekRetryMax,
			RetryBaseDelayMS: defaultDeepSeekRetryBaseDelayMS,
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
		ContextOutputReserveTokens: -1,
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
	applyDeepSeekRetryEnvOverrides(&cfg)
	if value := os.Getenv("NEMO_MAX_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.MaxTokens = parsed
		}
	}
	applyModelCapabilityEnvOverrides(&cfg)
	applyModelAwareChunkDefaults(&cfg)
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
	applyDeepSeekRetryEnvOverrides(&cfg)
	if value := os.Getenv("NEMO_MAX_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.MaxTokens = parsed
		}
	}
	applyModelCapabilityEnvOverrides(&cfg)
	applyModelAwareChunkDefaults(&cfg)
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
	applyModelCapabilityEnvOverrides(cfg)
	applyModelAwareChunkDefaults(cfg)
	applyChunkEnvOverrides(cfg)
	return nil
}

// applyProviderChunkDefaults selects provider defaults, then model-aware
// calculation can refine the chunk threshold for backends that expose context
// capability. Local llama keeps its empirical threshold unless explicitly
// configured with NEMO_MODEL_CONTEXT_TOKENS.
func applyProviderChunkDefaults(cfg *Config) {
	switch cfg.Provider {
	case "deepseek":
		cfg.ModelContextTokens = defaultDeepSeekContextTokens
		cfg.CharsPerToken = defaultDeepSeekCharsPerToken
		cfg.ContextReserveTokens = defaultDeepSeekContextReserveTokens
		cfg.ContextOutputReserveTokens = -1
		cfg.ContextSafetyMargin = defaultDeepSeekContextSafetyMargin
		cfg.QualityChunkCharThreshold = defaultDeepSeekQualityChunkCharThreshold
		cfg.ChunkedBundleCharThreshold = defaultDeepSeekQualityChunkCharThreshold
		cfg.MaxChunkChars = defaultDeepSeekMaxChunkChars
	default:
		cfg.ModelContextTokens = 0
		cfg.CharsPerToken = 0
		cfg.ContextReserveTokens = 0
		cfg.ContextOutputReserveTokens = -1
		cfg.ContextSafetyMargin = 0
		cfg.QualityChunkCharThreshold = 0
		cfg.ChunkedBundleCharThreshold = defaultLlamaChunkedBundleCharThreshold
		cfg.MaxChunkChars = defaultLlamaMaxChunkChars
	}
}

func applyModelAwareChunkDefaults(cfg *Config) {
	if threshold := finalChunkThreshold(*cfg); threshold > 0 {
		cfg.ChunkedBundleCharThreshold = threshold
	}
}

func finalChunkThreshold(cfg Config) int {
	return minPositive(derivedChunkThreshold(cfg), cfg.QualityChunkCharThreshold)
}

func derivedChunkThreshold(cfg Config) int {
	if cfg.ModelContextTokens <= 0 || cfg.CharsPerToken <= 0 {
		return 0
	}
	availableTokens := cfg.ModelContextTokens - cfg.ContextReserveTokens - contextOutputReserveTokens(cfg)
	if availableTokens <= 0 {
		return 0
	}
	margin := cfg.ContextSafetyMargin
	if margin <= 0 || margin > 1 {
		margin = 0.75
	}
	threshold := int(math.Floor(float64(availableTokens) * cfg.CharsPerToken * margin))
	if threshold <= 0 {
		return 0
	}
	return threshold
}

func applyModelCapabilityEnvOverrides(cfg *Config) {
	if value := os.Getenv("NEMO_MODEL_CONTEXT_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			cfg.ModelContextTokens = parsed
		}
	}
	if value := os.Getenv("NEMO_CHARS_PER_TOKEN"); value != "" {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil && parsed > 0 {
			cfg.CharsPerToken = parsed
		}
	}
	if value := os.Getenv("NEMO_CONTEXT_RESERVE_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed >= 0 {
			cfg.ContextReserveTokens = parsed
		}
	}
	if value := os.Getenv("NEMO_CONTEXT_OUTPUT_RESERVE_TOKENS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed >= 0 {
			cfg.ContextOutputReserveTokens = parsed
		}
	}
	if value := os.Getenv("NEMO_CONTEXT_SAFETY_MARGIN"); value != "" {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil && parsed > 0 && parsed <= 1 {
			cfg.ContextSafetyMargin = parsed
		}
	}
	if value := os.Getenv("NEMO_QUALITY_CHUNK_THRESHOLD_CHARS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed >= 0 {
			cfg.QualityChunkCharThreshold = parsed
		}
	}
}

func contextOutputReserveTokens(cfg Config) int {
	if cfg.ContextOutputReserveTokens >= 0 {
		return cfg.ContextOutputReserveTokens
	}
	if cfg.Provider == "deepseek" {
		return cfg.DeepSeek.MaxTokens
	}
	return cfg.MaxTokens
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

func minPositive(values ...int) int {
	min := 0
	for _, value := range values {
		if value <= 0 {
			continue
		}
		if min == 0 || value < min {
			min = value
		}
	}
	return min
}

func applyDeepSeekRetryEnvOverrides(cfg *Config) {
	if value := os.Getenv("NEMO_DEEPSEEK_RETRY_MAX"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed >= 0 {
			cfg.DeepSeek.RetryMax = parsed
		}
	}
	if value := os.Getenv("NEMO_DEEPSEEK_RETRY_BASE_MS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed >= 0 {
			cfg.DeepSeek.RetryBaseDelayMS = parsed
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
