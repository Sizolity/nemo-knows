package web

import (
	"strings"
	"testing"
	"time"
)

func TestValidateMarkdownUploadAcceptsBasicMarkdown(t *testing.T) {
	name, content, err := validateMarkdownUpload("my-note.md", "# My Note\n\nBody", time.Time{})
	if err != nil {
		t.Fatalf("validateMarkdownUpload returned error: %v", err)
	}
	if name != "my-note.md" {
		t.Fatalf("name = %q, want my-note.md", name)
	}
	if !strings.HasSuffix(content, "\n") {
		t.Fatalf("content should be newline-terminated: %q", content)
	}
}

func TestValidateMarkdownUploadDerivesNameFromTitleOrTimestamp(t *testing.T) {
	name, _, err := validateMarkdownUpload("", "# My New Note\n\nBody", time.Time{})
	if err != nil {
		t.Fatalf("validateMarkdownUpload returned error: %v", err)
	}
	if name != "my-new-note.md" {
		t.Fatalf("name = %q, want my-new-note.md", name)
	}

	now := time.Date(2026, 5, 24, 19, 58, 0, 0, time.UTC)
	name, _, err = validateMarkdownUpload("", "plain Markdown text without heading", now)
	if err != nil {
		t.Fatalf("validateMarkdownUpload returned error: %v", err)
	}
	if name != "source-20260524-195800.md" {
		t.Fatalf("name = %q, want timestamp source name", name)
	}
}

func TestValidateMarkdownUploadRejectsUnsafeInput(t *testing.T) {
	cases := []struct {
		name    string
		content string
	}{
		{"../evil.md", "# Title\n"},
		{"!!!.md", "# Title\n"},
		{"empty.md", ""},
		{"not-md.txt", "# Title\n"},
		{"script.md", "# Title\n\n<script>alert(1)</script>"},
	}
	for _, tc := range cases {
		if _, _, err := validateMarkdownUpload(tc.name, tc.content, time.Time{}); err == nil {
			t.Fatalf("expected %q to be rejected", tc.name)
		}
	}
}
