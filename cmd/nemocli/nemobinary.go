package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// findNemoBinary returns the absolute path of the sibling .bin/nemo
// binary that handles heavy draft / candidate generation. nemocli is
// designed to coexist with cmd/nemo in V0; the v3 design §10 V0 plan
// expects this binary to be built alongside .bin/nemocli.
//
// Resolution order:
//  1. Sibling: same directory as nemocli's own executable.
//  2. $PATH lookup: works for users who installed `nemo` system-wide.
//  3. Fallback: ./.bin/nemo (last-ditch, useful when nemocli is run
//     from the repo root via go run).
func findNemoBinary() (string, error) {
	if exe, err := os.Executable(); err == nil {
		sibling := filepath.Join(filepath.Dir(exe), "nemo")
		if isExecutable(sibling) {
			return sibling, nil
		}
	}
	if p, err := exec.LookPath("nemo"); err == nil {
		return p, nil
	}
	if cwd, err := os.Getwd(); err == nil {
		candidate := filepath.Join(cwd, ".bin", "nemo")
		if isExecutable(candidate) {
			return candidate, nil
		}
	}
	return "", fmt.Errorf("nemo binary not found alongside nemocli or in $PATH; build it with `go build -o .bin/nemo ./cmd/nemo`")
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if info.IsDir() {
		return false
	}
	return info.Mode()&0o111 != 0
}
