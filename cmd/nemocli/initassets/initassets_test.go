package initassets

import (
	"bytes"
	"strings"
	"testing"
)

func TestWikiAgentsTemplateNonEmpty(t *testing.T) {
	body := WikiAgentsTemplate()
	if len(body) == 0 {
		t.Fatal("wiki-agents.md template is empty")
	}
	// Sanity: header line should exist so we know the right file was embedded.
	want := []byte("# AGENTS.md")
	if !bytes.Contains(body, want) {
		t.Fatalf("template missing %q header; first 200 bytes:\n%s", want, body[:min(200, len(body))])
	}
}

func TestWikiAgentsTemplateContainsCoreSections(t *testing.T) {
	body := string(WikiAgentsTemplate())
	for _, marker := range []string{
		"## 0. Mental model",
		"## 1. Directory conventions",
		"## 6. The log format",
	} {
		if !strings.Contains(body, marker) {
			t.Errorf("wiki-agents.md missing section %q", marker)
		}
	}
}

func TestEnvExampleTemplate(t *testing.T) {
	body := string(EnvExampleTemplate())
	if len(body) == 0 {
		t.Fatal("env.example template is empty")
	}
	for _, want := range []string{
		"NEMO_MODEL_PROVIDER=deepseek",
		"NEMO_DEEPSEEK_API_KEY=",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("env.example missing %q", want)
		}
	}
	// Hard guard: the committed template must never contain a populated key.
	for _, leak := range []string{
		"sk-",
		"NEMO_DEEPSEEK_API_KEY=sk",
	} {
		if strings.Contains(body, leak) {
			t.Errorf("env.example contains suspicious value %q", leak)
		}
	}
}

func TestTemplatesAreCopies(t *testing.T) {
	a := WikiAgentsTemplate()
	b := WikiAgentsTemplate()
	if &a[0] == &b[0] {
		t.Fatal("WikiAgentsTemplate must return a fresh copy each call")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
