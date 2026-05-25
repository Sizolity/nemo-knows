// Package server provides the nemo-knows backend API service. Unlike the web
// console (internal/web) which renders HTML for interactive browsing, this
// package exposes a JSON API intended for programmatic access by other services
// (e.g. Shingo webhooks, automation scripts, monitoring).
//
// The server is designed to run as a long-lived systemd service alongside or
// instead of nemo-web, listening on a separate port.
package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/huic/nemo-knows/internal/config"
	"github.com/huic/nemo-knows/internal/wikilint"
)

// Server is the nemo-knows backend API service.
type Server struct {
	cfg          config.Config
	webhookToken string
	nemoPath     string

	mu  sync.Mutex
	job *apiJob
}

type apiJob struct {
	ID        string    `json:"id"`
	Source    string    `json:"source"`
	Status    string    `json:"status"`
	Step      string    `json:"step,omitempty"`
	Error     string    `json:"error,omitempty"`
	StartedAt time.Time `json:"started_at"`
}

// Options configures the server at construction time.
type Options struct {
	Config       config.Config
	WebhookToken string // expected token for webhook authentication
	NemoBinary   string // path to the nemo CLI binary for triggering pipelines
}

// New constructs a Server with the given options.
func New(opts Options) *Server {
	return &Server{
		cfg:          opts.Config,
		webhookToken: opts.WebhookToken,
		nemoPath:     opts.NemoBinary,
	}
}

// Handler returns an http.Handler with all API routes mounted.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/api/status", s.handleStatus)
	mux.HandleFunc("/api/wiki/pages", s.handleWikiPages)
	mux.HandleFunc("/api/wiki/page", s.handleWikiPage)
	mux.HandleFunc("/api/wiki/lint", s.handleWikiLint)
	mux.HandleFunc("/api/wiki/search", s.handleWikiSearch)
	mux.HandleFunc("/hooks/content", s.handleContentWebhook)
	return mux
}

// Run starts the API server on the given address. It blocks until the server
// exits or encounters an error.
func Run(addr string, opts Options) error {
	srv := New(opts)
	fmt.Fprintf(os.Stderr, "nemo-server listening on http://%s\n", addr)
	return http.ListenAndServe(addr, srv.Handler())
}

// --- Health & Status ---

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	s.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	pages := countWikiPages()
	s.mu.Lock()
	job := s.job
	s.mu.Unlock()

	resp := map[string]any{
		"service":    "nemo-knows",
		"provider":   s.cfg.Provider,
		"profile":    s.cfg.Profile,
		"wiki_pages": pages,
	}
	if job != nil {
		resp["current_job"] = job
	}
	s.writeJSON(w, http.StatusOK, resp)
}

// --- Wiki API ---

func (s *Server) handleWikiPages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.methodNotAllowed(w)
		return
	}
	kind := r.URL.Query().Get("kind")
	pages := listWikiPages(kind)
	s.writeJSON(w, http.StatusOK, map[string]any{"pages": pages})
}

func (s *Server) handleWikiPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.methodNotAllowed(w)
		return
	}
	path := r.URL.Query().Get("path")
	if path == "" || !isValidWikiPath(path) {
		s.writeError(w, http.StatusBadRequest, "invalid or missing path parameter")
		return
	}
	content, err := os.ReadFile(filepath.FromSlash(path))
	if err != nil {
		s.writeError(w, http.StatusNotFound, "page not found")
		return
	}
	s.writeJSON(w, http.StatusOK, map[string]any{
		"path":    path,
		"content": string(content),
	})
}

func (s *Server) handleWikiLint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		s.methodNotAllowed(w)
		return
	}
	result, err := wikilint.LintWiki(".")
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, "lint failed: "+err.Error())
		return
	}
	s.writeJSON(w, http.StatusOK, result)
}

func (s *Server) handleWikiSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.methodNotAllowed(w)
		return
	}
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		s.writeError(w, http.StatusBadRequest, "missing q parameter")
		return
	}
	results := searchWikiPages(query)
	s.writeJSON(w, http.StatusOK, map[string]any{"query": query, "results": results})
}

// --- Webhook ---

func (s *Server) handleContentWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.methodNotAllowed(w)
		return
	}
	if !s.authenticateWebhook(r) {
		s.writeError(w, http.StatusUnauthorized, "invalid or missing token")
		return
	}

	// Acknowledge receipt immediately; heavy work would be async via the CLI.
	s.writeJSON(w, http.StatusAccepted, map[string]string{
		"status":  "accepted",
		"message": "content webhook received",
	})
}

func (s *Server) authenticateWebhook(r *http.Request) bool {
	if s.webhookToken == "" {
		return true // no token configured = open (should only be on localhost)
	}
	token := r.Header.Get("X-Webhook-Token")
	if token == "" {
		token = r.URL.Query().Get("token")
	}
	return token == s.webhookToken
}

// --- Helpers ---

func (s *Server) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
}

func (s *Server) writeError(w http.ResponseWriter, status int, msg string) {
	s.writeJSON(w, status, map[string]string{"error": msg})
}

func (s *Server) methodNotAllowed(w http.ResponseWriter) {
	s.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
}

func isValidWikiPath(path string) bool {
	clean := filepath.ToSlash(filepath.Clean(path))
	return strings.HasPrefix(clean, "wiki/") && !strings.Contains(clean, "..")
}

type wikiPageInfo struct {
	Path  string `json:"path"`
	Title string `json:"title"`
	Kind  string `json:"kind"`
}

func listWikiPages(kind string) []wikiPageInfo {
	dirs := []string{"wiki/sources", "wiki/entities", "wiki/concepts", "wiki/topics"}
	if kind != "" {
		dirs = []string{"wiki/" + kind}
	}
	var pages []wikiPageInfo
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		dirKind := filepath.Base(dir)
		for _, entry := range entries {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
				continue
			}
			path := filepath.ToSlash(filepath.Join(dir, entry.Name()))
			pages = append(pages, wikiPageInfo{
				Path:  path,
				Title: titleFromPath(path),
				Kind:  dirKind,
			})
		}
	}
	return pages
}

func countWikiPages() int {
	return len(listWikiPages(""))
}

func titleFromPath(path string) string {
	content, err := os.ReadFile(filepath.FromSlash(path))
	if err != nil {
		return slugToTitle(path)
	}
	for _, line := range strings.SplitN(string(content), "\n", 20) {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "title:") {
			return strings.Trim(strings.TrimSpace(strings.TrimPrefix(trimmed, "title:")), `"'`)
		}
		if strings.HasPrefix(trimmed, "# ") {
			return strings.TrimPrefix(trimmed, "# ")
		}
	}
	return slugToTitle(path)
}

func slugToTitle(path string) string {
	slug := strings.TrimSuffix(filepath.Base(path), ".md")
	parts := strings.Split(slug, "-")
	for i, p := range parts {
		if p != "" {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}

type searchResult struct {
	Path    string `json:"path"`
	Title   string `json:"title"`
	Excerpt string `json:"excerpt,omitempty"`
}

func searchWikiPages(query string) []searchResult {
	lower := strings.ToLower(query)
	pages := listWikiPages("")
	var results []searchResult
	for _, page := range pages {
		content, err := os.ReadFile(filepath.FromSlash(page.Path))
		if err != nil {
			continue
		}
		text := string(content)
		if !strings.Contains(strings.ToLower(text), lower) {
			continue
		}
		excerpt := extractExcerpt(text, lower, 200)
		results = append(results, searchResult{
			Path:    page.Path,
			Title:   page.Title,
			Excerpt: excerpt,
		})
	}
	return results
}

func extractExcerpt(text string, query string, maxLen int) string {
	idx := strings.Index(strings.ToLower(text), query)
	if idx == -1 {
		return ""
	}
	start := idx - maxLen/4
	if start < 0 {
		start = 0
	}
	end := idx + len(query) + maxLen*3/4
	if end > len(text) {
		end = len(text)
	}
	excerpt := strings.TrimSpace(text[start:end])
	excerpt = strings.ReplaceAll(excerpt, "\n", " ")
	if start > 0 {
		excerpt = "..." + excerpt
	}
	if end < len(text) {
		excerpt = excerpt + "..."
	}
	return excerpt
}
