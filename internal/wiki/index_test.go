package wiki

import (
	"strings"
	"testing"
)

func TestFormatIndexEntryUsesMarkdownRelativeLink(t *testing.T) {
	got := FormatIndexEntry("wiki/sources/sqlite-wal.md", "SQLite WAL")
	want := "- [sqlite-wal](sources/sqlite-wal.md) — SQLite WAL."
	if got != want {
		t.Fatalf("FormatIndexEntry = %q, want %q", got, want)
	}
}

func TestIndexEntrySlugParsesMarkdownAndLegacyEntries(t *testing.T) {
	for _, line := range []string{
		"- [sqlite-wal](sources/sqlite-wal.md) — SQLite WAL.",
		"- [SQLite WAL](wiki/sources/sqlite-wal.md) — SQLite WAL.",
		"- [[sqlite-wal]] — SQLite WAL.",
	} {
		got, ok := IndexEntrySlug(line)
		if !ok {
			t.Fatalf("IndexEntrySlug(%q) did not match", line)
		}
		if got != "sqlite-wal" {
			t.Fatalf("IndexEntrySlug(%q) = %q, want sqlite-wal", line, got)
		}
	}
}

func TestIndexEntryTargetReturnsMarkdownTargetOnly(t *testing.T) {
	if got, ok := IndexEntryTarget("- [sqlite-wal](sources/sqlite-wal.md) — SQLite WAL."); !ok || got != "sources/sqlite-wal.md" {
		t.Fatalf("IndexEntryTarget markdown = (%q, %v), want (sources/sqlite-wal.md, true)", got, ok)
	}
	if _, ok := IndexEntryTarget("- [[sqlite-wal]] — SQLite WAL."); ok {
		t.Fatal("IndexEntryTarget should not match legacy wikilink entries")
	}
	if _, ok := IndexEntryTarget("_People, organisations, products, places._"); ok {
		t.Fatal("IndexEntryTarget should not match prose lines")
	}
}

func TestAppendIndexEntryReplacesPlaceholderOnFirstEntry(t *testing.T) {
	index := "---\ntitle: Index\nkind: index\n---\n\n## Topics\n\n(none yet)\n"
	got := AppendIndexEntry(index, "## Topics", FormatIndexEntry("wiki/topics/wal.md", "WAL"))

	if strings.Contains(got, "(none yet)") {
		t.Fatalf("first entry should replace placeholder:\n%s", got)
	}
	if !strings.Contains(got, "## Topics\n\n- [wal](topics/wal.md) — WAL.\n") {
		t.Fatalf("expected canonical first entry under heading:\n%s", got)
	}
	if !strings.HasSuffix(got, "\n") {
		t.Fatalf("index should end with a newline:\n%q", got)
	}
}

func TestAppendIndexEntryAppendsAfterExistingEntriesAndKeepsOtherSections(t *testing.T) {
	index := "---\ntitle: Index\nkind: index\n---\n\n## Sources\n\n- [a](sources/a.md) — A.\n\n## Topics\n\n(none yet)\n"
	got := AppendIndexEntry(index, "## Sources", FormatIndexEntry("wiki/sources/b.md", "B"))

	if !strings.Contains(got, "## Sources\n\n- [a](sources/a.md) — A.\n- [b](sources/b.md) — B.\n") {
		t.Fatalf("expected new entry appended after existing one:\n%s", got)
	}
	if !strings.Contains(got, "## Topics\n\n(none yet)\n") {
		t.Fatalf("appending to one section must not disturb another:\n%s", got)
	}
}

func TestAppendIndexEntryCreatesMissingSection(t *testing.T) {
	index := "---\ntitle: Index\nkind: index\n---\n\n## Sources\n\n- [a](sources/a.md) — A.\n"
	got := AppendIndexEntry(index, "## Concepts", FormatIndexEntry("wiki/concepts/c.md", "C"))

	if !strings.Contains(got, "## Concepts\n\n- [c](concepts/c.md) — C.\n") {
		t.Fatalf("expected missing section to be appended with the entry:\n%s", got)
	}
	if !strings.Contains(got, "## Sources\n\n- [a](sources/a.md) — A.\n") {
		t.Fatalf("existing section should be preserved:\n%s", got)
	}
}
