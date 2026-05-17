package wiki

import "testing"

func TestIsRawPath(t *testing.T) {
	ok, err := IsRawPath("raw/llm-wiki.md")
	if err != nil {
		t.Fatalf("IsRawPath returned error: %v", err)
	}
	if !ok {
		t.Fatal("expected raw path to be valid")
	}

	ok, err = IsRawPath("wiki/sources/llm-wiki.md")
	if err != nil {
		t.Fatalf("IsRawPath returned error: %v", err)
	}
	if ok {
		t.Fatal("expected wiki path not to be a raw path")
	}
}

func TestIsWikiPath(t *testing.T) {
	ok, err := IsWikiPath("wiki/sources/llm-wiki.md")
	if err != nil {
		t.Fatalf("IsWikiPath returned error: %v", err)
	}
	if !ok {
		t.Fatal("expected wiki path to be valid")
	}

	ok, err = IsWikiPath("raw/llm-wiki.md")
	if err != nil {
		t.Fatalf("IsWikiPath returned error: %v", err)
	}
	if ok {
		t.Fatal("expected raw path not to be a wiki path")
	}
}

func TestHasFrontmatter(t *testing.T) {
	ok, err := HasFrontmatter("---\ntitle: Test\n---\n# Body\n")
	if err != nil {
		t.Fatalf("HasFrontmatter returned error: %v", err)
	}
	if !ok {
		t.Fatal("expected frontmatter")
	}

	ok, err = HasFrontmatter("# Body\n")
	if err != nil {
		t.Fatalf("HasFrontmatter returned error: %v", err)
	}
	if ok {
		t.Fatal("did not expect frontmatter")
	}
}
