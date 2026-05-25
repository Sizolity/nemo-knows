package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/huic/nemo-knows/internal/config"
)

func TestHealthEndpoint(t *testing.T) {
	srv := New(Options{Config: config.Config{Provider: "test", Profile: "stable"}})
	ts := httptest.NewServer(srv.Handler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body["status"] != "ok" {
		t.Fatalf("expected status ok, got %q", body["status"])
	}
}

func TestStatusEndpoint(t *testing.T) {
	srv := New(Options{Config: config.Config{Provider: "deepseek", Profile: "stable"}})
	ts := httptest.NewServer(srv.Handler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/status")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var body map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body["service"] != "nemo-knows" {
		t.Fatalf("expected service nemo-knows, got %v", body["service"])
	}
	if body["provider"] != "deepseek" {
		t.Fatalf("expected provider deepseek, got %v", body["provider"])
	}
}

func TestWebhookRequiresToken(t *testing.T) {
	srv := New(Options{
		Config:       config.Config{Provider: "test"},
		WebhookToken: "secret-token",
	})
	ts := httptest.NewServer(srv.Handler())
	defer ts.Close()

	// Without token
	resp, err := http.Post(ts.URL+"/hooks/content", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected 401 without token, got %d", resp.StatusCode)
	}

	// With correct token
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/hooks/content", nil)
	req.Header.Set("X-Webhook-Token", "secret-token")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("expected 202 with token, got %d", resp.StatusCode)
	}
}

func TestWebhookOpenWithoutToken(t *testing.T) {
	srv := New(Options{Config: config.Config{Provider: "test"}})
	ts := httptest.NewServer(srv.Handler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/hooks/content", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("expected 202 when no token configured, got %d", resp.StatusCode)
	}
}

func TestWikiPageInvalidPath(t *testing.T) {
	srv := New(Options{Config: config.Config{Provider: "test"}})
	ts := httptest.NewServer(srv.Handler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/wiki/page?path=../etc/passwd")
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 for path traversal, got %d", resp.StatusCode)
	}
}
