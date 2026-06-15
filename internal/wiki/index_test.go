package wiki

import "testing"

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
