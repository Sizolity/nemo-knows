package deepseek

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientGenerateCallsChatCompletions(t *testing.T) {
	var requestPath string
	var authHeader string
	var request chatRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		authHeader = r.Header.Get("Authorization")
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"choices":[{"message":{"role":"assistant","content":"---\nkind: source\n---\n"}}]}`))
	}))
	defer server.Close()

	out, err := Client{
		BaseURL:         server.URL,
		APIKey:          "test-key",
		Model:           "deepseek-v4-pro",
		MaxTokens:       4096,
		Thinking:        "enabled",
		ReasoningEffort: "high",
	}.Generate(context.Background(), "hello")
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if out != "---\nkind: source\n---\n" {
		t.Fatalf("out = %q", out)
	}
	if requestPath != "/chat/completions" {
		t.Fatalf("request path = %q, want /chat/completions", requestPath)
	}
	if authHeader != "Bearer test-key" {
		t.Fatalf("auth header = %q", authHeader)
	}
	if request.Model != "deepseek-v4-pro" {
		t.Fatalf("model = %q", request.Model)
	}
	if len(request.Messages) != 1 || request.Messages[0].Role != "user" || request.Messages[0].Content != "hello" {
		t.Fatalf("messages = %#v", request.Messages)
	}
	if request.MaxTokens != 4096 {
		t.Fatalf("max tokens = %d", request.MaxTokens)
	}
	if request.Thinking == nil || request.Thinking.Type != "enabled" {
		t.Fatalf("thinking = %#v", request.Thinking)
	}
	if request.ReasoningEffort != "high" {
		t.Fatalf("reasoning effort = %q", request.ReasoningEffort)
	}
	if request.Temperature != nil {
		t.Fatalf("temperature = %v, want omitted in thinking mode", *request.Temperature)
	}
	if request.TopP != nil {
		t.Fatalf("top_p = %v, want omitted in thinking mode", *request.TopP)
	}
}

func TestClientGenerateSendsDeepSeekSpecificRequestFields(t *testing.T) {
	var request chatRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"choices":[{"message":{"role":"assistant","content":"{}"}}],"usage":{"completion_tokens":1,"prompt_tokens":2,"prompt_cache_hit_tokens":1,"prompt_cache_miss_tokens":1,"total_tokens":3}}`))
	}))
	defer server.Close()

	out, err := Client{
		BaseURL:        server.URL,
		APIKey:         "test-key",
		Model:          "deepseek-v4-pro",
		MaxTokens:      384000,
		Thinking:       "enabled",
		ResponseFormat: "json_object",
		UserID:         "nemo-test",
		SystemPrompt:   "Return JSON.",
	}.Generate(context.Background(), "hello")
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if out != "{}" {
		t.Fatalf("out = %q, want {}", out)
	}
	if len(request.Messages) != 2 {
		t.Fatalf("message count = %d, want 2", len(request.Messages))
	}
	if request.Messages[0].Role != "system" || request.Messages[0].Content != "Return JSON." {
		t.Fatalf("system message = %#v", request.Messages[0])
	}
	if request.Messages[1].Role != "user" || request.Messages[1].Content != "hello" {
		t.Fatalf("user message = %#v", request.Messages[1])
	}
	if request.ResponseFormat == nil || request.ResponseFormat.Type != "json_object" {
		t.Fatalf("response format = %#v", request.ResponseFormat)
	}
	if request.UserID != "nemo-test" {
		t.Fatalf("user_id = %q, want nemo-test", request.UserID)
	}
}

func TestClientGenerateSendsSamplingParamsWhenThinkingDisabled(t *testing.T) {
	var request chatRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"choices":[{"message":{"role":"assistant","content":"ok"}}]}`))
	}))
	defer server.Close()

	temp := 0.2
	topP := 0.9
	_, err := Client{
		BaseURL:     server.URL,
		APIKey:      "test-key",
		Model:       "deepseek-v4-flash",
		MaxTokens:   384000,
		Temperature: &temp,
		TopP:        &topP,
		Thinking:    "disabled",
	}.Generate(context.Background(), "hello")
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if request.MaxTokens != 384000 {
		t.Fatalf("max tokens = %d, want 384000", request.MaxTokens)
	}
	if request.Thinking == nil || request.Thinking.Type != "disabled" {
		t.Fatalf("thinking = %#v", request.Thinking)
	}
	if request.Temperature == nil || *request.Temperature != 0.2 {
		t.Fatalf("temperature = %v, want 0.2", request.Temperature)
	}
	if request.TopP == nil || *request.TopP != 0.9 {
		t.Fatalf("top_p = %v, want 0.9", request.TopP)
	}
}

func TestClientGenerateRequiresAPIKey(t *testing.T) {
	_, err := Client{BaseURL: "https://api.deepseek.com", Model: "deepseek-v4-pro"}.Generate(context.Background(), "hello")
	if err == nil {
		t.Fatal("expected missing API key error")
	}
}
