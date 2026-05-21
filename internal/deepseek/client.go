package deepseek

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const defaultTimeout = 10 * time.Minute

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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint(c.BaseURL), bytes.NewReader(encoded))
	if err != nil {
		return "", fmt.Errorf("create deepseek request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultTimeout}
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("call deepseek API: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read deepseek response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("deepseek API returned %s: %s", resp.Status, truncate(responseBody, 1000))
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
