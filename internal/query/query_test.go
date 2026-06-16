package query

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestExecuteAnswersFromRelevantWikiPages(t *testing.T) {
	root := makeQueryWiki(t)

	result, err := Execute(Options{
		Root:     root,
		Question: "How does WAL help SQLite readers?",
	})
	if err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}
	if !strings.Contains(result.Answer, "wiki/sources/sqlite-wal.md") {
		t.Fatalf("answer missing wiki citation:\n%s", result.Answer)
	}
	if len(result.Written) != 0 {
		t.Fatalf("read-only query should not write files, got %#v", result.Written)
	}
}

func TestExecuteFileQueryRequiresApprovalForWikiWrite(t *testing.T) {
	root := makeQueryWiki(t)

	_, err := Execute(Options{
		Root:     root,
		Question: "How does WAL help SQLite readers?",
		File:     true,
		Out:      "wiki/topics/how-does-wal-help-sqlite-readers.md",
	})
	if !errors.Is(err, ErrApprovalRequired) {
		t.Fatalf("error = %v, want ErrApprovalRequired", err)
	}
	if _, statErr := os.Stat(filepath.Join(root, "wiki", "topics", "how-does-wal-help-sqlite-readers.md")); !errors.Is(statErr, os.ErrNotExist) {
		t.Fatalf("unapproved query should not create topic, stat err=%v", statErr)
	}
}

func TestExecuteWritesReviewDraftWithoutApproval(t *testing.T) {
	root := makeQueryWiki(t)

	result, err := Execute(Options{
		Root:     root,
		Question: "How does WAL help SQLite readers?",
		File:     true,
		Now:      time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}
	if len(result.Written) != 1 || result.Written[0] != "tmp/query-drafts/wal-help-sqlite-readers.md" {
		t.Fatalf("expected review draft write, got %#v", result.Written)
	}
	draft, err := os.ReadFile(filepath.Join(root, "tmp", "query-drafts", "wal-help-sqlite-readers.md"))
	if err != nil {
		t.Fatalf("read draft: %v", err)
	}
	if !strings.Contains(string(draft), "kind: topic") || !strings.Contains(string(draft), "wiki/sources/sqlite-wal.md") {
		t.Fatalf("draft missing topic frontmatter or citations:\n%s", draft)
	}
	if _, statErr := os.Stat(filepath.Join(root, "wiki", "topics")); !errors.Is(statErr, os.ErrNotExist) {
		t.Fatalf("review draft should not create wiki topic dir, stat err=%v", statErr)
	}
}

func TestExecuteApprovedQueryWritesTopicIndexAndLog(t *testing.T) {
	root := makeQueryWiki(t)

	result, err := Execute(Options{
		Root:     root,
		Question: "How does WAL help SQLite readers?",
		File:     true,
		Approve:  true,
		Out:      "wiki/topics/wal-reader-concurrency.md",
		Now:      time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}
	for _, want := range []string{"wiki/topics/wal-reader-concurrency.md", "wiki/index.md", "wiki/log.md"} {
		if !contains(result.Written, want) {
			t.Fatalf("written files missing %s: %#v", want, result.Written)
		}
	}
	topic, err := os.ReadFile(filepath.Join(root, "wiki", "topics", "wal-reader-concurrency.md"))
	if err != nil {
		t.Fatalf("read topic: %v", err)
	}
	if !strings.Contains(string(topic), "source: wiki/sources/sqlite-wal.md") {
		t.Fatalf("topic missing citation:\n%s", topic)
	}
	index, err := os.ReadFile(filepath.Join(root, "wiki", "index.md"))
	if err != nil {
		t.Fatalf("read index: %v", err)
	}
	if !strings.Contains(string(index), "[wal-reader-concurrency](topics/wal-reader-concurrency.md)") {
		t.Fatalf("index missing filed query topic:\n%s", index)
	}
	if strings.Contains(string(index), "(none yet)") {
		t.Fatalf("filing the first topic should replace the Topics placeholder:\n%s", index)
	}
	log, err := os.ReadFile(filepath.Join(root, "wiki", "log.md"))
	if err != nil {
		t.Fatalf("read log: %v", err)
	}
	if !strings.Contains(string(log), "query-filed | How does WAL help SQLite readers") {
		t.Fatalf("log missing query-filed entry:\n%s", log)
	}
}

func TestExecuteRejectsApprovedQueryOutsideTopics(t *testing.T) {
	root := makeQueryWiki(t)

	_, err := Execute(Options{
		Root:     root,
		Question: "How does WAL help SQLite readers?",
		File:     true,
		Approve:  true,
		Out:      "wiki/entities/not-a-topic.md",
	})
	if err == nil {
		t.Fatal("expected unsafe target error")
	}
	if _, statErr := os.Stat(filepath.Join(root, "wiki", "entities", "not-a-topic.md")); !errors.Is(statErr, os.ErrNotExist) {
		t.Fatalf("unsafe approved query should not write entity, stat err=%v", statErr)
	}
}

func makeQueryWiki(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeQueryFile(t, root, "wiki/index.md", "---\ntitle: Index\nkind: index\n---\n\n## Sources\n\n- [sqlite-wal](sources/sqlite-wal.md) — SQLite WAL notes.\n\n## Topics\n\n(none yet)\n")
	writeQueryFile(t, root, "wiki/log.md", "# Log\n")
	writeQueryFile(t, root, "wiki/sources/sqlite-wal.md", `---
title: SQLite Write-Ahead Logging Notes
kind: source
sources:
  - pipeline/raw/sqlite-wal.md
confidence: medium
---

# SQLite Write-Ahead Logging Notes

SQLite WAL lets readers keep using a stable database snapshot while writers append committed frames to the write-ahead log. Checkpointing later copies frames back into the main database file.
`)
	return root
}

func writeQueryFile(t *testing.T, root string, repoPath string, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(repoPath))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir for %s: %v", repoPath, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", repoPath, err)
	}
}

func contains(items []string, want string) bool {
	for _, item := range items {
		if item == want {
			return true
		}
	}
	return false
}

func TestStripInlineMarkdownLinks(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"plain", "plain text", "plain text"},
		{"same-level link", "[Llama Cpp](llama-cpp.md) runs fast", "Llama Cpp runs fast"},
		{"parent link", "see [Georgi](../entities/georgi-gerganov.md) here", "see Georgi here"},
		{"anchor link", "[Sec](page.md#anchor) ref", "Sec ref"},
		{"image", "diagram ![alt text](../assets/x.png) below", "diagram alt text below"},
		{"multiple links", "[A](a.md) and [B](b.md)", "A and B"},
		{"empty alt image", "![](only.png)", ""},
		{"url target", "[label](http://example.com/path?q=1)", "label"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := stripInlineMarkdownLinks(tc.in); got != tc.want {
				t.Fatalf("stripInlineMarkdownLinks(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestTopicExcerptFlattensAndTruncates(t *testing.T) {
	const fm = "---\ntitle: X\nkind: entity\n---\n\n# X\n\n"

	flattened := topicExcerpt(fm + "Created by [Georgi Gerganov](georgi-gerganov.md) in 2023.\n")
	if want := "Created by Georgi Gerganov in 2023."; flattened != want {
		t.Fatalf("flattened excerpt = %q, want %q", flattened, want)
	}

	long := topicExcerpt(fm + "[Doc](x.md) " + strings.Repeat("a", 240) + "\n")
	if !strings.HasSuffix(long, "...") {
		t.Fatalf("long excerpt should be truncated with ellipsis: %q", long)
	}
	if strings.Contains(long, "](") {
		t.Fatalf("truncated excerpt still contains a link target: %q", long)
	}
	if len(long) > 223 {
		t.Fatalf("truncated excerpt too long (%d bytes): %q", len(long), long)
	}

	if got := topicExcerpt("---\ntitle: X\nkind: entity\n---\n\n# X\n"); got != "No short summary is available." {
		t.Fatalf("empty-body excerpt = %q, want fallback text", got)
	}
}

func TestRenderTopicDraftOmitsInlineLinkSyntax(t *testing.T) {
	const fmE = "---\ntitle: X\nkind: entity\n---\n\n"
	pages := []Page{
		{
			Path:    "wiki/entities/llama-cpp.md",
			Title:   "llama.cpp",
			Content: fmE + "# llama.cpp\n\nCreated by [Georgi Gerganov](georgi-gerganov.md), released in 2023.\n",
		},
		{
			Path:    "wiki/entities/georgi-gerganov.md",
			Title:   "Georgi Gerganov",
			Content: fmE + "# Georgi Gerganov\n\nInitiated [llama.cpp](llama-cpp.md); see ![logo](../assets/logo.png).\n",
		},
	}

	draft := renderTopicDraft("llama.cpp project history", pages, time.Date(2026, 6, 16, 0, 0, 0, 0, time.UTC))

	// The filed draft is written into wiki/topics/, so any inline link/image
	// syntax copied from a source page's directory would dangle there.
	for _, bad := range []string{"](", "![", "[["} {
		if strings.Contains(draft, bad) {
			t.Fatalf("filed topic draft must not contain %q (dangling-link risk):\n%s", bad, draft)
		}
	}
	for _, want := range []string{
		"Created by Georgi Gerganov, released in 2023.",
		"Initiated llama.cpp; see logo.",
		"(source: wiki/entities/llama-cpp.md)",
		"(source: wiki/entities/georgi-gerganov.md)",
	} {
		if !strings.Contains(draft, want) {
			t.Fatalf("filed topic draft missing %q:\n%s", want, draft)
		}
	}
}
