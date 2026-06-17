package main

import (
	"fmt"

	"github.com/huic/nemo-knows/internal/config"
)

// requireSecrets refuses to enter the LLM pipeline when the provider
// secrets are not configured. The error string explicitly tells the
// user where nemocli looked and how to fix it (run init, edit .env, or
// export the appropriate env var). Exit codes for ingest treat this as
// usage error (return 2).
//
// The same guard also lives in internal/config for safety-by-default,
// but this layer surfaces the message in a user-friendly form before
// any stage starts running.
func requireSecrets(cfg config.Config, wikiRoot string) error {
	switch cfg.Provider {
	case "deepseek":
		if cfg.DeepSeek.APIKey == "" {
			return fmt.Errorf(
				"ingest: missing DeepSeek API key\n"+
					"  looked in: env NEMO_DEEPSEEK_API_KEY, file %s/.env\n"+
					"  fix: run `nemocli init` to scaffold $WIKI_ROOT/.env.example,\n"+
					"       then `cp .env.example .env && $EDITOR .env`,\n"+
					"       or `set -x NEMO_DEEPSEEK_API_KEY <key>` (fish).",
				wikiRoot,
			)
		}
	case "llama":
		if cfg.LlamaCLI == "" || cfg.LlamaModel == "" {
			return fmt.Errorf(
				"ingest: llama provider requires both NEMO_LLAMA_CLI and NEMO_LLAMA_MODEL\n"+
					"  looked in: env vars, file %s/.env\n"+
					"  fix: `set -x NEMO_LLAMA_CLI /path/to/llama-cli` and\n"+
					"       `set -x NEMO_LLAMA_MODEL /path/to/model.gguf` (fish),\n"+
					"       or set both in %s/.env.",
				wikiRoot, wikiRoot,
			)
		}
	default:
		return fmt.Errorf("ingest: unsupported provider %q (set NEMO_MODEL_PROVIDER to deepseek or llama)", cfg.Provider)
	}
	return nil
}
