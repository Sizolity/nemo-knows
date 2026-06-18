// Package rundir manages the per-ingest working directory for nemocli.
//
// Every ingest run gets a fresh runDir created via os.MkdirTemp. Stages
// inside the run treat that directory as the bundle workspace. On success
// the runDir is removed entirely, leaving wiki/ as the only persistent
// artifact; on failure the runDir is renamed (or copied, on EXDEV) to
// /tmp/nemo-checkpoints/<runID>/ so the user can inspect or resume it.
//
// This is the V0 "Go-native" lifecycle described in the v3 design §6.1.
// It deliberately does not introduce a Scope/VFS abstraction (R1 rollback).
package rundir

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// CheckpointBase is where failed runs are spilled.
//
// V0 keeps this as a single global location to align with the design
// (§6.1) and with the systemd flock path /tmp/nemo-ingest.lock; tests
// override it via NewSpiller.
var CheckpointBase = "/tmp/nemo-checkpoints"

// Open allocates a fresh runDir under os.TempDir() and returns it along
// with a runID. runID is intentionally independent of the runDir path so
// that downstream logs and wiki Run: anchors stay stable even if V2 ever
// switches to an in-memory backend.
func Open() (runDir string, runID string, err error) {
	dir, err := os.MkdirTemp("", "nemo-run-")
	if err != nil {
		return "", "", fmt.Errorf("create run dir: %w", err)
	}
	return dir, NewRunID(time.Now()), nil
}

// NewRunID renders a deterministic per-run identifier suitable for log
// anchors and checkpoint directory names.
//
// The id encodes a wall-clock timestamp and an 8-character random suffix
// so concurrent attempts at the same second still differ.
func NewRunID(now time.Time) string {
	nonce := make([]byte, 4)
	if _, err := rand.Read(nonce); err != nil {
		// Hex of zeros is acceptable; the timestamp prefix still makes
		// concurrent IDs distinguishable in practice for V0.
		return now.Format("20060102-150405") + "-00000000"
	}
	return now.Format("20060102-150405") + "-" + hex.EncodeToString(nonce)
}

// Close finalizes runDir. When success is true the directory is removed;
// otherwise it is moved to <CheckpointBase>/<runID>/, falling back to a
// copy + remove when rename hits EXDEV.
//
// The returned checkpoint path is empty when success is true. On failure
// it is the final on-disk location of the spilled data (the checkpoint
// directory when rename or copy succeeded; the original runDir when
// neither succeeded).
func Close(success bool, runDir, runID string) (checkpoint string, err error) {
	if success {
		if removeErr := os.RemoveAll(runDir); removeErr != nil {
			return "", fmt.Errorf("remove run dir: %w", removeErr)
		}
		return "", nil
	}
	return spillToCheckpoint(runDir, runID)
}

func spillToCheckpoint(runDir, runID string) (string, error) {
	if err := os.MkdirAll(CheckpointBase, 0o755); err != nil {
		return runDir, fmt.Errorf("create checkpoint base: %w", err)
	}
	dst := filepath.Join(CheckpointBase, runID)
	if err := os.Rename(runDir, dst); err == nil {
		return dst, nil
	} else if !errors.Is(err, errCrossDevice) && !isCrossDeviceErr(err) {
		// rename failed for some non-EXDEV reason; still try the copy
		// fallback so the data is not lost.
	}
	if err := copyDir(runDir, dst); err != nil {
		return runDir, fmt.Errorf("copy run dir to checkpoint: %w", err)
	}
	if err := os.RemoveAll(runDir); err != nil {
		// Copy succeeded; original leftover is non-fatal but worth noting
		// in the returned error so the caller can surface it.
		return dst, fmt.Errorf("copy to checkpoint succeeded but failed to remove source %s: %w", runDir, err)
	}
	return dst, nil
}

// errCrossDevice is a sentinel used only for documentation; the actual
// EXDEV check is done structurally by isCrossDeviceErr below since the
// concrete errno value differs across platforms.
var errCrossDevice = errors.New("rundir: cross-device link")

func isCrossDeviceErr(err error) bool {
	// On Linux rename(2) returns EXDEV (errno 18). Rather than depend on
	// syscall.EXDEV directly we match the unix error string; this is good
	// enough for V0 because the copy fallback runs unconditionally when
	// rename returns any error.
	if err == nil {
		return false
	}
	for _, marker := range []string{"cross-device", "invalid cross-device", "EXDEV"} {
		if containsErrText(err, marker) {
			return true
		}
	}
	return false
}

func containsErrText(err error, marker string) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	for i := 0; i+len(marker) <= len(msg); i++ {
		if msg[i:i+len(marker)] == marker {
			return true
		}
	}
	return false
}

// copyDir copies src into dst, preserving file modes. It exists so the
// rename fallback can succeed across filesystems (EXDEV).
func copyDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			return os.MkdirAll(target, info.Mode().Perm())
		}
		return copyFile(path, target)
	})
}

func copyFile(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, info.Mode().Perm())
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return nil
}
