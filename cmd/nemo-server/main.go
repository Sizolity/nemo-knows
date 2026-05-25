// Command nemo-server is the dedicated backend API service for nemo-knows.
// It exposes a JSON API for programmatic access by other services (Shingo
// webhooks, automation scripts, monitoring) and runs as a long-lived systemd
// service.
//
// This is distinct from nemo-web (the interactive web console UI) and from the
// nemo CLI (the local pipeline tool). All three can coexist:
//
//	nemo          — local CLI for interactive pipeline operations
//	nemo-web      — HTML web console for browsing wiki + triggering ingest (port 8787)
//	nemo-server   — JSON API backend for webhooks + service integration (port 8788)
//
// Usage:
//
//	nemo-server                             # default 127.0.0.1:8788
//	nemo-server -addr 0.0.0.0:8788
//	nemo-server -nemo-binary .bin/nemo
//
// Environment variables:
//
//	NEMO_SERVER_ADDR          — listen address (overridden by -addr flag)
//	NEMO_BINARY               — path to nemo CLI binary
//	NEMO_KNOWS_WEBHOOK_TOKEN  — expected token for webhook authentication
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/huic/nemo-knows/internal/config"
	"github.com/huic/nemo-knows/internal/server"
)

func main() {
	addr := flag.String("addr", defaultAddr(), "address to listen on")
	nemoBinary := flag.String("nemo-binary", defaultNemoBinary(), "path to the nemo CLI binary")
	flag.Parse()

	cfg, err := config.ForProfile("")
	if err != nil {
		fmt.Fprintln(os.Stderr, "config:", err)
		os.Exit(1)
	}

	opts := server.Options{
		Config:       cfg,
		WebhookToken: os.Getenv("NEMO_KNOWS_WEBHOOK_TOKEN"),
		NemoBinary:   *nemoBinary,
	}

	if err := server.Run(*addr, opts); err != nil {
		fmt.Fprintln(os.Stderr, "server:", err)
		os.Exit(1)
	}
}

func defaultAddr() string {
	if v := os.Getenv("NEMO_SERVER_ADDR"); v != "" {
		return v
	}
	return "127.0.0.1:8788"
}

func defaultNemoBinary() string {
	if v := os.Getenv("NEMO_BINARY"); v != "" {
		return v
	}
	if _, err := os.Stat(".bin/nemo"); err == nil {
		return ".bin/nemo"
	}
	return "nemo"
}
