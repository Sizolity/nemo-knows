package main

import (
	"fmt"
	"os"
	"syscall"
)

// IngestLockPath is the cross-process mutex for the wiki writer. V0
// uses the same path as the legacy systemd timer to keep migration
// painless; .cursor/rules/nemo-single-instance.mdc and
// deploy/systemd/install-user-units.sh both reference this path.
const IngestLockPath = "/tmp/nemo-ingest.lock"

// acquireFlock takes a non-blocking flock(2) on path and returns a
// release function. The caller must defer release(); the file handle is
// also closed inside release.
//
// The lock file itself stays on disk; flock semantics tie the lock to
// the open file descriptor, not the inode, so leaving the file is
// fine — it just gets reused next time. When the lock is already held
// by another nemocli writer we return a clear error so the caller can
// surface "skipped" in the runs log.
func acquireFlock(path string) (func(), error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		return nil, fmt.Errorf("open lock file %s: %w", path, err)
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		f.Close()
		return nil, fmt.Errorf("acquire flock %s: %w (another nemocli writer is active)", path, err)
	}
	return func() {
		_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
		_ = f.Close()
	}, nil
}
