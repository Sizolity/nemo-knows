package draft

import "testing"

func TestPathsForMarkdownOutput(t *testing.T) {
	paths := PathsFor("drafts/llm-wiki-source.md")

	if paths.Cleaned != "drafts/llm-wiki-source.md" {
		t.Fatalf("unexpected cleaned path: %s", paths.Cleaned)
	}
	if paths.Raw != "drafts/llm-wiki-source.raw.txt" {
		t.Fatalf("unexpected raw path: %s", paths.Raw)
	}
}
