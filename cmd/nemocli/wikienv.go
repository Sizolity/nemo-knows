package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// loadWikiEnv reads $WIKI_ROOT/.env if it exists and exports any keys
// that are not already set in the process environment. Missing files
// are silently ignored — fresh installs run init before any .env exists.
//
// The parser is intentionally minimal (KEY=VALUE per line, # comments,
// optional "export " prefix, optional surrounding quotes). It mirrors
// internal/config.loadDotEnv so behaviour is consistent across the two
// callers; the function deliberately does not modify any existing env
// var, which matches the v3 design §9.2 "env-wins" rule.
func loadWikiEnv(wikiRoot string) error {
	if wikiRoot == "" {
		return nil
	}
	path := filepath.Join(wikiRoot, ".env")
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimSpace(strings.TrimPrefix(line, "export "))
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		if key == "" || strings.ContainsAny(key, " \t") {
			continue
		}
		if _, exists := os.LookupEnv(key); exists {
			continue
		}
		_ = os.Setenv(key, trimDotEnvValue(value))
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	return nil
}

// trimDotEnvValue strips one matching pair of surrounding quotes and
// trims whitespace, mirroring internal/config.trimDotEnvValue.
func trimDotEnvValue(value string) string {
	value = strings.TrimSpace(value)
	if len(value) < 2 {
		return value
	}
	q := value[0]
	if q != '"' && q != '\'' {
		return value
	}
	if value[len(value)-1] != q {
		return value
	}
	return value[1 : len(value)-1]
}
