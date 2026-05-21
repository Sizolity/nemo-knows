package deepseek

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"time"
)

const defaultTimeout = 10 * time.Minute
const defaultRetryMax = 2
const defaultRetryBaseDelay = time.Second

// Client calls DeepSeek's OpenAI-compatible chat completions API.
type Client struct {
	BaseURL         string
	APIKey          string
	Model           string
	MaxTokens       int
	Temperature     *float64
	TopP            *float64
	Thinking        string
	ReasoningEffort string
	ResponseFormat  string
	UserID          string
	SystemPrompt    string
	HTTPClient      *http.Client
	RetryMax        int
	RetryBaseDelay  time.Duration
}

type chatRequest struct {
	Model           string          `json:"model"`
	Messages        []chatMessage   `json:"messages"`
	MaxTokens       int             `json:"max_tokens,omitempty"`
	Temperature     *float64        `json:"temperature,omitempty"`
	TopP            *float64        `json:"top_p,omitempty"`
	Thinking        *thinkingBlock  `json:"thinking,omitempty"`
	ReasoningEffort string          `json:"reasoning_effort,omitempty"`
	ResponseFormat  *responseFormat `json:"response_format,omitempty"`
	UserID          string          `json:"user_id,omitempty"`
	Stream          bool            `json:"stream"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type thinkingBlock struct {
	Type string `json:"type"`
}

type responseFormat struct {
	Type string `json:"type"`
}

type chatResponse struct {
	Choices []struct {
		Message chatMessage `json:"message"`
	} `json:"choices"`
	Usage *usage `json:"usage,omitempty"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

type usage struct {
	CompletionTokens       int `json:"completion_tokens"`
	PromptTokens           int `json:"prompt_tokens"`
	PromptCacheHitTokens   int `json:"prompt_cache_hit_tokens"`
	PromptCacheMissTokens  int `json:"prompt_cache_miss_tokens"`
	TotalTokens            int `json:"total_tokens"`
	CompletionTokenDetails *struct {
		ReasoningTokens int `json:"reasoning_tokens"`
	} `json:"completion_tokens_details,omitempty"`
}

// Generate sends the rendered prompt using DeepSeek chat messages and returns
// the assistant content exactly as received so the draft cleaner can preserve it.
func (c Client) Generate(ctx context.Context, prompt string) (string, error) {
	if strings.TrimSpace(c.BaseURL) == "" {
		return "", errors.New("deepseek base URL is required")
	}
	if strings.TrimSpace(c.APIKey) == "" {
		return "", errors.New("deepseek API key is required")
	}
	if strings.TrimSpace(c.Model) == "" {
		return "", errors.New("deepseek model is required")
	}

	messages := []chatMessage{}
	if strings.TrimSpace(c.SystemPrompt) != "" {
		messages = append(messages, chatMessage{Role: "system", Content: c.SystemPrompt})
	}
	messages = append(messages, chatMessage{Role: "user", Content: prompt})

	body := chatRequest{
		Model:     c.Model,
		Messages:  messages,
		MaxTokens: c.MaxTokens,
		UserID:    c.UserID,
		Stream:    false,
	}
	if c.Thinking != "" {
		body.Thinking = &thinkingBlock{Type: c.Thinking}
	}
	if c.Thinking != "enabled" {
		body.Temperature = c.Temperature
		body.TopP = c.TopP
	}
	if c.ReasoningEffort != "" {
		body.ReasoningEffort = c.ReasoningEffort
	}
	if c.ResponseFormat != "" && c.ResponseFormat != "text" {
		body.ResponseFormat = &responseFormat{Type: c.ResponseFormat}
	}

	encoded, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("encode deepseek request: %w", err)
	}

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultTimeout}
	}

	responseBody, err := c.doWithRetries(ctx, httpClient, encoded)
	if err != nil {
		return "", err
	}

	var decoded chatResponse
	if err := json.Unmarshal(responseBody, &decoded); err != nil {
		return "", fmt.Errorf("decode deepseek response: %w", err)
	}
	if decoded.Error != nil {
		return "", fmt.Errorf("deepseek API error %s: %s", decoded.Error.Type, decoded.Error.Message)
	}
	if len(decoded.Choices) == 0 {
		return "", errors.New("deepseek API returned no choices")
	}
	content := decoded.Choices[0].Message.Content
	if content == "" {
		return "", errors.New("deepseek API returned empty content")
	}

	return content, nil
}

func (c Client) doWithRetries(ctx context.Context, httpClient *http.Client, encoded []byte) ([]byte, error) {
	maxRetries := c.RetryMax
	if maxRetries < 0 {
		maxRetries = defaultRetryMax
	}
	baseDelay := c.RetryBaseDelay
	if baseDelay <= 0 {
		baseDelay = defaultRetryBaseDelay
	}

	var lastErr error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		responseBody, retryable, err := c.doOnce(ctx, httpClient, encoded)
		if err == nil {
			return responseBody, nil
		}
		lastErr = err
		if !retryable || attempt == maxRetries {
			break
		}
		if err := sleepBeforeRetry(ctx, baseDelay, attempt); err != nil {
			return nil, err
		}
	}
	return nil, lastErr
}

func (c Client) doOnce(ctx context.Context, httpClient *http.Client, encoded []byte) ([]byte, bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint(c.BaseURL), bytes.NewReader(encoded))
	if err != nil {
		return nil, false, fmt.Errorf("create deepseek request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, isRetryableRequestError(err), fmt.Errorf("call deepseek API: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, true, fmt.Errorf("read deepseek response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		retryable := resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500
		return nil, retryable, fmt.Errorf("deepseek API returned %s: %s", resp.Status, truncate(responseBody, 1000))
	}
	return responseBody, false, nil
}

func isRetryableRequestError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) {
		return false
	}
	text := strings.ToLower(err.Error())
	return strings.Contains(text, "eof") ||
		strings.Contains(text, "connection reset") ||
		strings.Contains(text, "broken pipe") ||
		strings.Contains(text, "timeout") ||
		strings.Contains(text, "temporary")
}

func sleepBeforeRetry(ctx context.Context, baseDelay time.Duration, attempt int) error {
	multiplier := math.Pow(2, float64(attempt))
	delay := time.Duration(float64(baseDelay) * multiplier)
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func endpoint(baseURL string) string {
	return strings.TrimRight(baseURL, "/") + "/chat/completions"
}

func truncate(body []byte, max int) string {
	text := strings.TrimSpace(string(body))
	if len(text) <= max {
		return text
	}
	return text[:max] + "..."
}
