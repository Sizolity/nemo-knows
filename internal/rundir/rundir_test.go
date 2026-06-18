package rundir

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestNewRunID_Format(t *testing.T) {
	now := time.Date(2026, 6, 17, 17, 30, 42, 0, time.UTC)
	id := NewRunID(now)
	if !strings.HasPrefix(id, "20260617-173042-") {
		t.Fatalf("runID prefix wrong: %q", id)
	}
	// "<timestamp>-<8 hex>" = 15+1+8 = 24 chars
	if len(id) != len("20060102-150405-")+8 {
		t.Fatalf("runID length wrong: %q (%d)", id, len(id))
	}
}

func TestOpenClose_Success_RemovesRunDir(t *testing.T) {
	runDir, runID, err := Open()
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if _, err := os.Stat(runDir); err != nil {
		t.Fatalf("runDir should exist: %v", err)
	}
	// add a payload file
	payload := filepath.Join(runDir, "bundle", "source.md")
	if err := os.MkdirAll(filepath.Dir(payload), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(payload, []byte("hi"), 0o644); err != nil {
		t.Fatal(err)
	}
	checkpoint, err := Close(true, runDir, runID)
	if err != nil {
		t.Fatalf("Close(success): %v", err)
	}
	if checkpoint != "" {
		t.Fatalf("success path should return empty checkpoint, got %q", checkpoint)
	}
	if _, err := os.Stat(runDir); !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("runDir should be removed, stat err = %v", err)
	}
}

func TestOpenClose_Failure_RenamesToCheckpoint(t *testing.T) {
	tmp := t.TempDir()
	origBase := CheckpointBase
	CheckpointBase = filepath.Join(tmp, "nemo-checkpoints")
	t.Cleanup(func() { CheckpointBase = origBase })

	runDir, runID, err := Open()
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if err := os.WriteFile(filepath.Join(runDir, "marker.txt"), []byte("payload"), 0o644); err != nil {
		t.Fatal(err)
	}

	checkpoint, err := Close(false, runDir, runID)
	if err != nil {
		t.Fatalf("Close(fail): %v", err)
	}
	wantCheckpoint := filepath.Join(CheckpointBase, runID)
	// Rename may happen across filesystems in some sandboxes; accept
	// either the same path or the copied-to-checkpoint path.
	if checkpoint != wantCheckpoint {
		t.Fatalf("checkpoint should be %q, got %q", wantCheckpoint, checkpoint)
	}
	// Payload must be in the checkpoint
	got, err := os.ReadFile(filepath.Join(checkpoint, "marker.txt"))
	if err != nil {
		t.Fatalf("read checkpoint marker: %v", err)
	}
	if string(got) != "payload" {
		t.Fatalf("payload should match, got %q", got)
	}
	// Original runDir must be gone
	if _, err := os.Stat(runDir); !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("runDir should be removed after spill, stat err = %v", err)
	}
}

func TestCopyDir_FallbackPreservesFiles(t *testing.T) {
	tmp := t.TempDir()
	src := filepath.Join(tmp, "src")
	dst := filepath.Join(tmp, "dst")
	if err := os.MkdirAll(filepath.Join(src, "sub"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(src, "a.txt"), []byte("a"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(src, "sub", "b.txt"), []byte("b"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := copyDir(src, dst); err != nil {
		t.Fatalf("copyDir: %v", err)
	}
	a, err := os.ReadFile(filepath.Join(dst, "a.txt"))
	if err != nil || string(a) != "a" {
		t.Fatalf("dst/a.txt: %v %q", err, a)
	}
	b, err := os.ReadFile(filepath.Join(dst, "sub", "b.txt"))
	if err != nil || string(b) != "b" {
		t.Fatalf("dst/sub/b.txt: %v %q", err, b)
	}
}
