package main

import (
	"fmt"
	"html"
	"html/template"
	"io/fs"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/huic/nemo-knows/internal/config"
)

type webServer struct {
	cfg config.Config

	mu  sync.Mutex
	job *webJob
}

type webJob struct {
	ID         string
	Source     string
	BundleDir  string
	Profile    string
	Provider   string
	Status     string
	Step       string
	Error      string
	StartedAt  time.Time
	FinishedAt time.Time
	Files      []string
}

type sidebarItem struct {
	Title  string
	Path   string
	Active bool
}

type webPageData struct {
	Title        string
	CFG          config.Config
	Sources      []string
	WikiPages    []string
	WikiSources  []string
	WikiEntities []string
	WikiConcepts []string
	WikiTopics   []string
	Job          *webJob
	Path         string
	Content      template.HTML
	Graph        template.HTML
	Edges        []webGraphEdge
	GraphNote    string
	Error        string
	Success      string
	// Sidebar categories
	SidebarSources  []sidebarItem
	SidebarEntities []sidebarItem
	SidebarConcepts []sidebarItem
	SidebarTopics   []sidebarItem
}

func (s *webServer) makePageData(currentPath string, title string) webPageData {
	sources, entities, concepts, topics := s.getSidebarItems(currentPath)
	return webPageData{
		Title:           title,
		CFG:             s.cfg,
		Sources:         mustListTextFiles("raw"),
		WikiPages:       mustListKBFiles(),
		WikiSources:     mustListSubKBFiles("wiki/sources"),
		WikiEntities:    mustListSubKBFiles("wiki/entities"),
		WikiConcepts:    mustListSubKBFiles("wiki/concepts"),
		WikiTopics:      mustListSubKBFiles("wiki/topics"),
		Job:             s.currentJob(),
		SidebarSources:  sources,
		SidebarEntities: entities,
		SidebarConcepts: concepts,
		SidebarTopics:   topics,
	}
}

func (s *webServer) getSidebarItems(currentPath string) ([]sidebarItem, []sidebarItem, []sidebarItem, []sidebarItem) {
	sources := mustListSubKBFiles("wiki/sources")
	entities := mustListSubKBFiles("wiki/entities")
	concepts := mustListSubKBFiles("wiki/concepts")
	topics := mustListSubKBFiles("wiki/topics")

	makeItems := func(paths []string) []sidebarItem {
		items := make([]sidebarItem, len(paths))
		for i, p := range paths {
			items[i] = sidebarItem{
				Title:  titleForMarkdownFile(p),
				Path:   p,
				Active: p == currentPath,
			}
		}
		return items
	}
	return makeItems(sources), makeItems(entities), makeItems(concepts), makeItems(topics)
}

type webGraphNode struct {
	Slug  string
	Title string
	Path  string
	X     float64
	Y     float64
}

type webGraphEdge struct {
	From     string
	FromPath string
	To       string
	ToPath   string
	Known    bool
}

func runWeb(addr string, cfg config.Config) error {
	server := &webServer{cfg: cfg}
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.handleHome)
	mux.HandleFunc("/run", server.handleRun)
	mux.HandleFunc("/job", server.handleJob)
	mux.HandleFunc("/build", server.handleBuild)
	mux.HandleFunc("/graph", server.handleGraph)
	mux.HandleFunc("/view", server.handleView)
	mux.HandleFunc("/static/style.css", handleWebStyle)

	fmt.Fprintf(os.Stderr, "nemo web console listening on http://%s\n", addr)
	return http.ListenAndServe(addr, mux)
}

func (s *webServer) handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := s.makePageData("", "nemo-knows")
	data.Graph = renderWikiGraph(40)
	data.Edges = wikiGraphEdges(12)
	data.GraphNote = wikiGraphNote()
	s.render(w, "home", data)
}

func (s *webServer) handleRun(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	source, err := cleanProjectPath(r.FormValue("source"))
	if err != nil || !strings.HasPrefix(source, "raw/") {
		http.Error(w, "source must be a project-relative path under raw/", http.StatusBadRequest)
		return
	}
	if _, err := os.Stat(filepath.FromSlash(source)); err != nil {
		http.Error(w, fmt.Sprintf("source not found: %s", source), http.StatusBadRequest)
		return
	}

	profile := strings.TrimSpace(r.FormValue("profile"))
	if profile == "" {
		profile = s.cfg.Profile
	}
	provider := strings.TrimSpace(r.FormValue("provider"))
	cfg, err := config.ForProfileWithProvider(profile, provider)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jobID := webJobID(source)
	job := &webJob{
		ID:        jobID,
		Source:    source,
		BundleDir: filepath.ToSlash(filepath.Join("drafts", jobID)),
		Profile:   cfg.Profile,
		Provider:  cfg.Provider,
		Status:    "running",
		Step:      "queued",
		StartedAt: time.Now(),
	}

	s.mu.Lock()
	if s.job != nil && s.job.Status == "running" {
		s.mu.Unlock()
		http.Error(w, "another nemo run is already in progress", http.StatusConflict)
		return
	}
	s.job = job
	s.mu.Unlock()

	go s.runJob(job, cfg)
	http.Redirect(w, r, "/job", http.StatusSeeOther)
}

func (s *webServer) handleJob(w http.ResponseWriter, r *http.Request) {
	s.render(w, "job", s.makePageData("", "当前任务"))
}

func (s *webServer) handleGraph(w http.ResponseWriter, r *http.Request) {
	data := s.makePageData("", "知识图谱")
	data.Graph = renderWikiGraph(48)
	data.Edges = wikiGraphEdges(80)
	data.GraphNote = wikiGraphNote()
	s.render(w, "graph", data)
}

func (s *webServer) handleView(w http.ResponseWriter, r *http.Request) {
	path, err := cleanProjectPath(r.URL.Query().Get("path"))
	if err != nil || !isWebViewPath(path) {
		http.Error(w, "invalid view path", http.StatusBadRequest)
		return
	}
	content, err := os.ReadFile(filepath.FromSlash(path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	slugToPath := getSlugToPathMap()
	data := s.makePageData(path, titleForMarkdownFile(path))
	data.Path = path
	data.Content = renderWebContent(path, content, slugToPath)
	s.render(w, "view", data)
}

func (s *webServer) handleBuild(w http.ResponseWriter, r *http.Request) {
	var errMsg string
	var successMsg string

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		if action == "save_source" {
			name := strings.TrimSpace(r.FormValue("name"))
			content := r.FormValue("content")
			if name == "" || content == "" {
				errMsg = "文献名称和内容不能为空"
			} else {
				if !strings.HasSuffix(name, ".md") && !strings.HasSuffix(name, ".txt") {
					name += ".md"
				}
				// sanitize file name to avoid directory traversal
				name = filepath.Base(name)
				targetPath := filepath.Join("raw", name)
				err := os.WriteFile(targetPath, []byte(content), 0644)
				if err != nil {
					errMsg = "保存文献失败: " + err.Error()
				} else {
					successMsg = fmt.Sprintf("新文献 %s 导入成功！现在可以在下方选择并开始构建知识库。", name)
				}
			}
		}
	}

	data := s.makePageData("/build", "知识构建")
	data.Error = errMsg
	data.Success = successMsg
	s.render(w, "build", data)
}

func (s *webServer) runJob(job *webJob, cfg config.Config) {
	lockFile, err := acquireNemoLock()
	if err != nil {
		s.updateJob(job, func(j *webJob) {
			j.FinishedAt = time.Now()
			j.Status = "failed"
			j.Step = "lock"
			j.Error = err.Error()
		})
		return
	}
	defer releaseNemoLock(lockFile)

	s.updateJob(job, func(j *webJob) {
		j.Step = "generating source draft and ingest plan"
	})
	err = runBundle(job.Source, filepath.FromSlash(job.BundleDir), cfg)
	if err == nil {
		s.updateJob(job, func(j *webJob) {
			j.Step = "reviewing bundle"
		})
		err = runReviewBundle(filepath.FromSlash(job.BundleDir), filepath.Join(filepath.FromSlash(job.BundleDir), "apply-plan.md"))
	}

	s.updateJob(job, func(j *webJob) {
		j.FinishedAt = time.Now()
		j.Files = mustListMarkdownFiles(filepath.FromSlash(job.BundleDir))
		if err != nil {
			j.Status = "failed"
			j.Error = err.Error()
			return
		}
		j.Status = "completed"
		j.Step = "done"
	})
}

func acquireNemoLock() (*os.File, error) {
	file, err := os.OpenFile("/tmp/nemo-ingest.lock", os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, fmt.Errorf("open ingest lock: %w", err)
	}
	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		_ = file.Close()
		return nil, fmt.Errorf("another nemo ingest is already running")
	}
	return file, nil
}

func releaseNemoLock(file *os.File) {
	_ = syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	_ = file.Close()
}

func (s *webServer) updateJob(job *webJob, update func(*webJob)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.job == job {
		update(s.job)
	}
}

func (s *webServer) currentJob() *webJob {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.job == nil {
		return nil
	}
	copy := *s.job
	copy.Files = append([]string(nil), s.job.Files...)
	return &copy
}

func (s *webServer) render(w http.ResponseWriter, name string, data webPageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := webTemplates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleWebStyle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	_, _ = w.Write([]byte(webCSS))
}

func mustListMarkdownFiles(root string) []string {
	return mustListFiles(root, func(path string, d fs.DirEntry) bool {
		return !d.IsDir() && strings.EqualFold(filepath.Ext(path), ".md")
	})
}

func mustListSubKBFiles(dir string) []string {
	if _, err := os.Stat(dir); err != nil {
		return nil
	}
	return mustListMarkdownFiles(dir)
}

func mustListKBFiles() []string {
	var files []string
	for _, sub := range []string{"wiki/sources", "wiki/entities", "wiki/concepts", "wiki/topics"} {
		files = append(files, mustListSubKBFiles(sub)...)
	}
	return files
}

func getSlugToPathMap() map[string]string {
	pages := mustListKBFiles()
	m := make(map[string]string, len(pages))
	for _, p := range pages {
		slug := slugFromWikiPath(p)
		m[slug] = p
	}
	return m
}

func mustListTextFiles(root string) []string {
	return mustListFiles(root, func(path string, d fs.DirEntry) bool {
		if d.IsDir() {
			return false
		}
		switch strings.ToLower(filepath.Ext(path)) {
		case ".md", ".txt":
			return true
		default:
			return false
		}
	})
}

func mustListFiles(root string, include func(string, fs.DirEntry) bool) []string {
	var files []string
	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if include(path, d) {
			files = append(files, filepath.ToSlash(path))
		}
		return nil
	})
	sort.Strings(files)
	return files
}

func wikiGraphNote() string {
	pages := mustListKBFiles()
	edges := wikiGraphEdges(0)
	return fmt.Sprintf("%d 个知识节点，%d 条链接关系", len(pages), len(edges))
}

func renderWikiGraph(limit int) template.HTML {
	nodes, edges := buildWikiGraph(limit)
	if len(nodes) == 0 {
		return template.HTML(`<p class="muted">还没有可展示的 wiki 页面。</p>`)
	}

	var b strings.Builder
	b.WriteString(`<svg class="knowledge-graph" viewBox="0 0 860 520" role="img" aria-label="wiki 知识图谱">`)
	b.WriteString(`<defs><marker id="arrow" markerWidth="8" markerHeight="8" refX="7" refY="3" orient="auto"><path d="M0,0 L0,6 L7,3 z"/></marker></defs>`)
	for _, edge := range edges {
		if edge.FromPath == "" || edge.ToPath == "" {
			continue
		}
		from, okFrom := nodeByPath(nodes, edge.FromPath)
		to, okTo := nodeByPath(nodes, edge.ToPath)
		if !okFrom || !okTo {
			continue
		}
		fmt.Fprintf(&b, `<line class="graph-edge" x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f"/>`, from.X, from.Y, to.X, to.Y)
	}
	for _, node := range nodes {
		fmt.Fprintf(&b, `<a href="/view?path=%s"><g class="graph-node">`, html.EscapeString(node.Path))
		fmt.Fprintf(&b, `<circle cx="%.1f" cy="%.1f" r="26"/>`, node.X, node.Y)
		fmt.Fprintf(&b, `<text x="%.1f" y="%.1f">%s</text>`, node.X, node.Y+44, html.EscapeString(shortGraphLabel(node.Title)))
		b.WriteString(`</g></a>`)
	}
	b.WriteString(`</svg>`)
	return template.HTML(b.String())
}

func nodeByPath(nodes []webGraphNode, path string) (webGraphNode, bool) {
	for _, node := range nodes {
		if node.Path == path {
			return node, true
		}
	}
	return webGraphNode{}, false
}

func buildWikiGraph(limit int) ([]webGraphNode, []webGraphEdge) {
	pages := mustListKBFiles()
	if limit > 0 && len(pages) > limit {
		pages = pages[:limit]
	}

	nodes := make([]webGraphNode, 0, len(pages))
	pathBySlug := map[string]string{}
	titleByPath := map[string]string{}
	for _, path := range pages {
		title := titleForMarkdownFile(path)
		slug := slugFromWikiPath(path)
		pathBySlug[slug] = path
		titleByPath[path] = title
		nodes = append(nodes, webGraphNode{Slug: slug, Title: title, Path: path})
	}

	const centerX = 430.0
	const centerY = 245.0
	const radius = 175.0
	for i := range nodes {
		angle := (2 * math.Pi * float64(i) / math.Max(float64(len(nodes)), 1)) - math.Pi/2
		nodes[i].X = centerX + radius*math.Cos(angle)
		nodes[i].Y = centerY + radius*math.Sin(angle)
	}

	edges := make([]webGraphEdge, 0)
	seen := map[string]bool{}
	for _, from := range pages {
		content, err := os.ReadFile(filepath.FromSlash(from))
		if err != nil {
			continue
		}
		matches := candidateWikilinkRE.FindAllStringSubmatch(string(content), -1)
		for _, match := range matches {
			targetSlug := slugFromLink(match[1])
			toPath := pathBySlug[targetSlug]
			key := from + "->" + targetSlug
			if seen[key] {
				continue
			}
			seen[key] = true
			edges = append(edges, webGraphEdge{
				From:     titleByPath[from],
				FromPath: from,
				To:       titleFromSlug(targetSlug),
				ToPath:   toPath,
				Known:    toPath != "",
			})
		}
	}
	return nodes, edges
}

func wikiGraphEdges(limit int) []webGraphEdge {
	_, edges := buildWikiGraph(0)
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].From == edges[j].From {
			return edges[i].To < edges[j].To
		}
		return edges[i].From < edges[j].From
	})
	if limit > 0 && len(edges) > limit {
		return edges[:limit]
	}
	return edges
}

func titleForMarkdownFile(path string) string {
	content, err := os.ReadFile(filepath.FromSlash(path))
	if err != nil {
		return titleFromSlug(slugFromWikiPath(path))
	}
	for _, line := range strings.Split(string(content), "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "title:") {
			return strings.Trim(strings.TrimSpace(strings.TrimPrefix(trimmed, "title:")), `"'`)
		}
		if strings.HasPrefix(trimmed, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(trimmed, "# "))
		}
	}
	return titleFromSlug(slugFromWikiPath(path))
}

func shortGraphLabel(label string) string {
	const maxRunes = 18
	runes := []rune(label)
	if len(runes) <= maxRunes {
		return label
	}
	return string(runes[:maxRunes-1]) + "…"
}

func cleanProjectPath(path string) (string, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return "", fmt.Errorf("empty path")
	}
	path = filepath.ToSlash(path)
	if strings.HasPrefix(path, "/") {
		return "", fmt.Errorf("absolute paths are not allowed")
	}
	cleaned := filepath.ToSlash(filepath.Clean(path))
	if cleaned == "." || strings.HasPrefix(cleaned, "../") || cleaned == ".." {
		return "", fmt.Errorf("path escapes project root")
	}
	return cleaned, nil
}

func isWebViewPath(path string) bool {
	return strings.HasPrefix(path, "raw/") ||
		strings.HasPrefix(path, "wiki/") ||
		strings.HasPrefix(path, "drafts/")
}

func webJobID(source string) string {
	slug := slugFromFilename(filepath.Base(source))
	if slug == "" {
		slug = "source"
	}
	return "web-" + time.Now().Format("20060102-150405") + "-" + slug
}

func renderWebContent(path string, content []byte, slugToPath map[string]string) template.HTML {
	if strings.EqualFold(filepath.Ext(path), ".md") {
		return renderMarkdown(content, slugToPath)
	}
	return template.HTML("<pre>" + html.EscapeString(string(content)) + "</pre>")
}

var wikilinkRE = regexp.MustCompile(`\[\[([^\]|#]+)(?:[|#][^\]]*)?\]\]`)

func renderMarkdown(content []byte, slugToPath map[string]string) template.HTML {
	lines := strings.Split(stripFrontmatter(string(content)), "\n")
	var b strings.Builder
	inCode := false
	inList := false
	var paragraph []string

	flushParagraph := func() {
		if len(paragraph) == 0 {
			return
		}
		b.WriteString("<p>")
		b.WriteString(renderInline(strings.Join(paragraph, " "), slugToPath))
		b.WriteString("</p>\n")
		paragraph = nil
	}
	closeList := func() {
		if inList {
			b.WriteString("</ul>\n")
			inList = false
		}
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "```") {
			flushParagraph()
			closeList()
			if inCode {
				b.WriteString("</code></pre>\n")
				inCode = false
			} else {
				b.WriteString("<pre><code>")
				inCode = true
			}
			continue
		}
		if inCode {
			b.WriteString(html.EscapeString(line))
			b.WriteByte('\n')
			continue
		}
		if trimmed == "" {
			flushParagraph()
			closeList()
			continue
		}
		if level, text, ok := markdownHeading(trimmed); ok {
			flushParagraph()
			closeList()
			fmt.Fprintf(&b, "<h%d>%s</h%d>\n", level, renderInline(text, slugToPath), level)
			continue
		}
		if strings.HasPrefix(trimmed, "- ") {
			flushParagraph()
			if !inList {
				b.WriteString("<ul>\n")
				inList = true
			}
			b.WriteString("<li>")
			b.WriteString(renderInline(strings.TrimSpace(strings.TrimPrefix(trimmed, "- ")), slugToPath))
			b.WriteString("</li>\n")
			continue
		}
		paragraph = append(paragraph, trimmed)
	}
	flushParagraph()
	closeList()
	if inCode {
		b.WriteString("</code></pre>\n")
	}
	return template.HTML(b.String())
}

func stripFrontmatter(content string) string {
	if !strings.HasPrefix(content, "---\n") {
		return content
	}
	rest := content[len("---\n"):]
	end := strings.Index(rest, "\n---")
	if end == -1 {
		return content
	}
	after := rest[end+len("\n---"):]
	return strings.TrimPrefix(after, "\n")
}

func markdownHeading(line string) (int, string, bool) {
	level := 0
	for level < len(line) && line[level] == '#' {
		level++
	}
	if level == 0 || level > 6 || level >= len(line) || line[level] != ' ' {
		return 0, "", false
	}
	return level, strings.TrimSpace(line[level+1:]), true
}

func renderInline(text string, slugToPath map[string]string) string {
	escaped := html.EscapeString(text)
	return wikilinkRE.ReplaceAllStringFunc(escaped, func(match string) string {
		content := strings.Trim(match, "[]")
		slug := content
		label := content
		if parts := strings.SplitN(content, "|", 2); len(parts) == 2 {
			slug = parts[0]
			label = parts[1]
		}
		resolvedSlug := slugFromLink(html.UnescapeString(slug))
		if path, ok := slugToPath[resolvedSlug]; ok {
			return fmt.Sprintf(`<a href="/view?path=%s" class="wikilink">[[%s]]</a>`, path, label)
		}
		return fmt.Sprintf(`<span class="wikilink missing-wikilink">[[%s]]</span>`, label)
	})
}

var webTemplates = template.Must(template.New("web").Funcs(template.FuncMap{
	"since": func(t time.Time) string {
		if t.IsZero() {
			return ""
		}
		return time.Since(t).Round(time.Second).String()
	},
	"titleForPath": titleForMarkdownFile,
}).Parse(`
{{define "layout-start"}}
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
  <link rel="stylesheet" href="/static/style.css">
</head>
<body>
<div class="app-container">
  <header>
    <a class="brand" href="/">nemo-knows Wiki</a>
    <nav>
      <a href="/graph">知识图谱</a>
      <a href="https://github.com/Sizolity/nemo-knows" target="_blank" rel="noopener noreferrer">原仓库</a>
    </nav>
  </header>
  <div class="app-body">
    <aside class="app-sidebar">
      <div class="sidebar-category">
        <div class="sidebar-cat-header badge-general">
          <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path><circle cx="12" cy="10" r="3"></circle></svg>
          快速导航
        </div>
        <ul class="sidebar-list">
          <li class="{{if eq .Path "wiki/index.md"}}active{{end}}">
            <a href="/view?path=wiki/index.md">
              <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"></path><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"></path></svg>
              内容索引
            </a>
          </li>
          <li class="{{if eq .Path "wiki/log.md"}}active{{end}}">
            <a href="/view?path=wiki/log.md">
              <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
              变更日志
            </a>
          </li>
          <li class="{{if eq .Path "/build"}}active{{end}}">
            <a href="/build">
              <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
              知识构建
            </a>
          </li>
        </ul>
      </div>

      <div class="sidebar-category">
        <div class="sidebar-cat-header badge-source">
          <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path><polyline points="14 2 14 8 20 8"></polyline><line x1="16" y1="13" x2="8" y2="13"></line><line x1="16" y1="17" x2="8" y2="17"></line><polyline points="10 9 9 9 8 9"></polyline></svg>
          来源文献 ({{len .SidebarSources}})
        </div>
        <ul class="sidebar-list">
          {{range .SidebarSources}}
            <li class="{{if .Active}}active{{end}}"><a href="/view?path={{.Path}}">{{.Title}}</a></li>
          {{else}}
            <li class="empty-list">暂无来源</li>
          {{end}}
        </ul>
      </div>

      <div class="sidebar-category">
        <div class="sidebar-cat-header badge-entity">
          <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
          核心实体 ({{len .SidebarEntities}})
        </div>
        <ul class="sidebar-list">
          {{range .SidebarEntities}}
            <li class="{{if .Active}}active{{end}}"><a href="/view?path={{.Path}}">{{.Title}}</a></li>
          {{else}}
            <li class="empty-list">暂无实体</li>
          {{end}}
        </ul>
      </div>

      <div class="sidebar-category">
        <div class="sidebar-cat-header badge-concept">
          <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 18h6"></path><path d="M10 22h4"></path><path d="M15.09 14c.18-.19.33-.42.49-.61C16.85 11.8 17.5 10 17.5 8a5.5 5.5 0 0 0-11 0c0 2 .65 3.8 1.92 5.39.16.19.31.42.49.61c.42.43.59.83.59 1.41h5c0-.58.17-.98.59-1.41z"></path></svg>
          技术概念 ({{len .SidebarConcepts}})
        </div>
        <ul class="sidebar-list">
          {{range .SidebarConcepts}}
            <li class="{{if .Active}}active{{end}}"><a href="/view?path={{.Path}}">{{.Title}}</a></li>
          {{else}}
            <li class="empty-list">暂无概念</li>
          {{end}}
        </ul>
      </div>

      <div class="sidebar-category">
        <div class="sidebar-cat-header badge-topic">
          <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="12 2 2 7 12 12 22 7 12 2"></polygon><polyline points="2 17 12 22 22 17"></polyline><polyline points="2 12 12 17 22 12"></polyline></svg>
          综合主题 ({{len .SidebarTopics}})
        </div>
        <ul class="sidebar-list">
          {{range .SidebarTopics}}
            <li class="{{if .Active}}active{{end}}"><a href="/view?path={{.Path}}">{{.Title}}</a></li>
          {{else}}
            <li class="empty-list">暂无主题</li>
          {{end}}
        </ul>
      </div>
    </aside>
    <article class="app-content">
{{end}}

{{define "layout-end"}}
    <footer class="app-footer">
      <p>© 2026 <a href="https://github.com/Sizolity" target="_blank" rel="noopener noreferrer">Sizolity</a>. Nemo Knows. All rights reserved.</p>
    </footer>
    </article>
  </div>
</div>
</body>
</html>
{{end}}

{{define "home"}}
{{template "layout-start" .}}
<section class="hero">
  <p class="eyebrow">领域知识库</p>
  <h1>浏览并探索您构建和凝练的领域知识体系。</h1>
  <p>这是一个完全基于双链（wikilink）架构的知识网络。请使用左侧侧边栏导航到各个词条，或者直接全屏探索下方的知识图谱。</p>
  <div class="actions">
    <a class="button-link" href="/graph">全屏探索图谱</a>
    <a class="secondary-link" href="/view?path=wiki/index.md">分类索引 (index)</a>
    <a class="secondary-link" href="/view?path=wiki/log.md">变更日志 (log)</a>
  </div>
</section>

<section class="card graph-card">
  <div class="section-heading">
    <div>
      <p class="eyebrow">Wiki 关联</p>
      <h2>知识图谱</h2>
      <p class="muted">{{.GraphNote}}。可点击下方图中任意圆形节点以查阅该条目的详细正文内容。</p>
    </div>
  </div>
  {{.Graph}}
</section>
{{template "layout-end" .}}
{{end}}

{{define "job"}}
{{template "layout-start" .}}
<section class="card">
  <h1>任务监控</h1>
  {{if .Job}}
    <dl class="details">
      <dt>任务状态</dt><dd>{{.Job.Status}}</dd>
      <dt>当前步骤</dt><dd>{{.Job.Step}}</dd>
      <dt>输入文献</dt><dd><a href="/view?path={{.Job.Source}}">{{.Job.Source}}</a></dd>
      <dt>草稿输出</dt><dd>{{.Job.BundleDir}}</dd>
      <dt>模型后端</dt><dd>{{.Job.Provider}}</dd>
      <dt>运行强度</dt><dd>{{.Job.Profile}}</dd>
      <dt>启动时间</dt><dd>{{.Job.StartedAt}}</dd>
      {{if .Job.FinishedAt.IsZero}}{{else}}<dt>结束时间</dt><dd>{{.Job.FinishedAt}}</dd>{{end}}
      {{if .Job.Error}}<dt>错误详情</dt><dd class="error">{{.Job.Error}}</dd>{{end}}
    </dl>
    {{if .Job.Files}}
    <h2>生成的草稿文件</h2>
    <ul class="file-list">
      {{range .Job.Files}}<li><a href="/view?path={{.}}">{{.}}</a></li>{{end}}
    </ul>
    {{end}}
    {{if eq .Job.Status "running"}}<p class="hint">异步任务运行中，刷新本页即可更新进度。</p>{{end}}
  {{else}}
    <p>当前无运行中的 Ingest 任务。</p>
  {{end}}
</section>
{{template "layout-end" .}}
{{end}}

{{define "graph"}}
{{template "layout-start" .}}
<section class="card graph-card">
  <div class="section-heading">
    <div>
      <p class="eyebrow">Wiki 关联</p>
      <h1>知识图谱</h1>
      <p class="muted">{{.GraphNote}}。点击任意圆圈节点可以直接跳转查阅对应页面的全部正文。</p>
    </div>
    <a class="secondary-link" href="/">返回主页</a>
  </div>
  {{.Graph}}
</section>
{{template "layout-end" .}}
{{end}}

{{define "view"}}
{{template "layout-start" .}}
<div class="card content">
  <p class="eyebrow" style="margin-bottom: 24px; border-bottom: 1px solid var(--line); padding-bottom: 12px; color: var(--muted);">{{.Path}}</p>
  {{.Content}}
</div>
{{template "layout-end" .}}
{{end}}

{{define "build"}}
{{template "layout-start" .}}
<section class="hero" style="margin-bottom: 32px;">
  <p class="eyebrow">数据导入与编译</p>
  <h1>添加新文献并构建为大模型知识库</h1>
  <p>在此处直接输入或粘贴全新的原始文献材料。保存后，文献将以 Markdown 格式持久化存储，随后您便可以一键调度 nemo-knows 的 AI 流水线进行知识提炼、关系编译并同步汇入您的双链知识图谱中。</p>
</section>

{{if .Error}}
  <div class="card" style="margin-bottom: 24px; border-color: #fca5a5; background: #fef2f2; color: #b91c1c; padding: 16px 20px; border-radius: 12px;">
    <p style="margin: 0; display: flex; align-items: center; gap: 8px;">
      <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="width: 18px; height: 18px;"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="8" x2="12" y2="12"></line><line x1="12" y1="16" x2="12.01" y2="16"></line></svg>
      <strong>保存失败：</strong>{{.Error}}
    </p>
  </div>
{{end}}

{{if .Success}}
  <div class="card" style="margin-bottom: 24px; border-color: #86efac; background: #f0fdf4; color: #166534; padding: 16px 20px; border-radius: 12px;">
    <p style="margin: 0; display: flex; align-items: center; gap: 8px;">
      <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="width: 18px; height: 18px;"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path><polyline points="22 4 12 14.01 9 11.01"></polyline></svg>
      <strong>导入成功：</strong>{{.Success}}
    </p>
  </div>
{{end}}

<div class="grid" style="display: grid; gap: 24px; grid-template-columns: repeat(auto-fit, minmax(380px, 1fr)); margin-bottom: 40px;">
  <section class="card" style="padding: 28px;">
    <h2 style="margin-top: 0; font-size: 20px; display: flex; align-items: center; gap: 10px;">
      <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="width: 20px; height: 20px; color: var(--accent);"><path d="M12 5v14M5 12h14"></path></svg>
      📥 导入新文献数据
    </h2>
    <p class="muted" style="font-size: 13.5px; margin-bottom: 20px; line-height: 1.6;">输入文献英文唯一标识文件名并粘贴内容。提交后将生成并存储在 <code>raw/</code> 目录下，充当知识库的根事实（Sources）来源。</p>
    
    <form method="POST" action="/build">
      <input type="hidden" name="action" value="save_source">
      <label style="display: block; margin-bottom: 18px;">
        <span style="font-size: 13.5px; font-weight: 600; display: block; margin-bottom: 8px; color: var(--text);">文献路径文件名 (例如: <code>my-essay.md</code>)</span>
        <input type="text" name="name" placeholder="请输入文件名，如 test-go-faq.md" required style="padding: 11px 14px; border: 1px solid var(--line); border-radius: 8px; width: 100%; font-family: inherit; font-size: 14px; background: #fff;">
      </label>
      <label style="display: block; margin-bottom: 20px;">
        <span style="font-size: 13.5px; font-weight: 600; display: block; margin-bottom: 8px; color: var(--text);">文献正文内容 (Markdown / 纯文本)</span>
        <textarea name="content" placeholder="请在此处直接粘贴或输入您的根文献内容..." required style="min-height: 240px; padding: 12px 14px; border: 1px solid var(--line); border-radius: 8px; width: 100%; font-family: inherit; font-size: 14px; resize: vertical; line-height: 1.6; background: #fff;"></textarea>
      </label>
      <button type="submit" class="button-link" style="border: 0; min-height: 42px; display: flex; justify-content: center; align-items: center; width: 100%; font-weight: 700; font-size: 14px;">📥 保存文献至本地 raw/</button>
    </form>
  </section>

  <section class="card" style="padding: 28px;">
    <h2 style="margin-top: 0; font-size: 20px; display: flex; align-items: center; gap: 10px;">
      <svg class="sidebar-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="width: 20px; height: 20px; color: var(--accent);"><path d="M12 2v4M12 18v4M4.93 4.93l2.83 2.83M16.24 16.24l2.83 2.83M2 12h4M18 12h4M4.93 19.07l2.83-2.83M16.24 7.76l2.83-2.83"></path></svg>
      ⚙️ 开始编译知识库 (Compile)
    </h2>
    <p class="muted" style="font-size: 13.5px; margin-bottom: 20px; line-height: 1.6;">大语言模型将提取文献中涉及的所有来源（Sources）、实体（Entities）、概念（Concepts）与主题（Topics），生成结构化草稿，并在审批通过后将其汇编入正式知识库。</p>
    
    <form method="POST" action="/run">
      <label style="display: block; margin-bottom: 18px;">
        <span style="font-size: 13.5px; font-weight: 600; display: block; margin-bottom: 8px; color: var(--text);">选择参考文献来源</span>
        <select name="source" required style="padding: 11px 14px; border: 1px solid var(--line); border-radius: 8px; width: 100%; font-family: inherit; font-size: 14px; background: #fff;">
          <option value="">-- 请选择要编译构建的原始文献 --</option>
          {{range .Sources}}
            <option value="{{.}}">{{.}}</option>
          {{end}}
        </select>
      </label>
      <label style="display: block; margin-bottom: 18px;">
        <span style="font-size: 13.5px; font-weight: 600; display: block; margin-bottom: 8px; color: var(--text);">生成强度配置 (Profile)</span>
        <select name="profile" style="padding: 11px 14px; border: 1px solid var(--line); border-radius: 8px; width: 100%; font-family: inherit; font-size: 14px; background: #fff;">
          <option value="stable" {{if eq .CFG.Profile "stable"}}selected{{end}}>Stable (推荐，稳定性最佳)</option>
          <option value="fast" {{if eq .CFG.Profile "fast"}}selected{{end}}>Fast (快速测试)</option>
          <option value="deep" {{if eq .CFG.Profile "deep"}}selected{{end}}>Deep (深度推理提取)</option>
        </select>
      </label>
      <button type="submit" class="button-link" style="border: 0; min-height: 42px; display: flex; justify-content: center; align-items: center; width: 100%; font-weight: 700; font-size: 14px;">⚡ 启动 AI 汇编与提炼</button>
    </form>
  </section>
</div>
{{template "layout-end" .}}
{{end}}
`))

const webCSS = `
:root {
  color-scheme: light dark;
  --bg: #f6f2ea;
  --surface: #fffdf8;
  --card: #ffffff;
  --text: #1f2933;
  --muted: #667085;
  --line: #e5ded1;
  --accent: #2563eb;
  --accent-strong: #1d4ed8;
  --warm: #fff3d8;
  --shadow: 0 8px 30px rgba(50, 43, 31, 0.04);
  --sidebar-width: 280px;
  --header-height: 60px;
}
* { box-sizing: border-box; }
body {
  margin: 0;
  background: var(--bg);
  color: var(--text);
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", sans-serif;
  font-size: 15px;
  line-height: 1.65;
}
a { color: var(--accent); text-decoration: none; }
a:hover { text-decoration: underline; }

.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

header {
  align-items: center;
  background: rgba(255, 253, 248, .92);
  backdrop-filter: blur(16px);
  border-bottom: 1px solid var(--line);
  display: flex;
  height: var(--header-height);
  justify-content: space-between;
  padding: 0 28px;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100;
}
nav { display: flex; gap: 16px; }
.brand { color: var(--text); font-weight: 700; }

.app-body {
  display: flex;
  flex: 1;
  margin-top: var(--header-height);
}

.app-sidebar {
  background: #faf7f0;
  border-right: 1px solid var(--line);
  width: var(--sidebar-width);
  position: fixed;
  top: var(--header-height);
  bottom: 0;
  left: 0;
  overflow-y: auto;
  padding: 24px 16px;
  z-index: 90;
}

.app-content {
  flex: 1;
  margin-left: var(--sidebar-width);
  padding: 40px 48px 64px;
  max-width: calc(100% - var(--sidebar-width));
  background: var(--bg);
}

.sidebar-category {
  margin-bottom: 24px;
}
.sidebar-cat-header {
  font-size: 11px;
  font-weight: 700;
  letter-spacing: .08em;
  padding: 6px 12px;
  border-radius: 6px;
  margin-bottom: 8px;
  display: flex;
  align-items: center;
  gap: 8px;
  text-transform: uppercase;
}
.sidebar-icon {
  width: 14px;
  height: 14px;
  flex-shrink: 0;
}
.sidebar-list {
  list-style: none;
  margin: 0;
  padding: 0;
}
.sidebar-list li {
  margin-bottom: 4px;
}
.sidebar-list li a {
  color: var(--text);
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  padding: 6px 12px;
  border-radius: 6px;
  text-decoration: none;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  transition: all .15s ease;
}
.sidebar-list li a:hover {
  background: rgba(37, 99, 235, .05);
  color: var(--accent);
}
.sidebar-list li.active a {
  background: rgba(37, 99, 235, .08);
  color: var(--accent);
  font-weight: 700;
}

.hero {
  margin-bottom: 22px;
  padding: 12px 0 8px;
}
.hero h1 { font-size: clamp(28px, 4vw, 42px); letter-spacing: -.035em; line-height: 1.1; margin: 0 0 14px; }
.hero p { color: var(--muted); font-size: 16px; max-width: 780px; }
.card {
  background: var(--card);
  border: 1px solid var(--line);
  border-radius: 18px;
  padding: 24px;
  box-shadow: var(--shadow);
}
.eyebrow { color: #9a6b16; font-size: 12px; font-weight: 700; letter-spacing: .08em; text-transform: uppercase; margin: 0; }
.actions { display: flex; flex-wrap: wrap; gap: 12px; margin-top: 20px; }
.button-link, .secondary-link {
  align-items: center;
  border-radius: 999px;
  display: inline-flex;
  font-weight: 700;
  min-height: 38px;
  padding: 8px 16px;
  font-size: 14px;
}
.button-link {
  background: var(--accent);
  color: #fff;
}
.button-link:hover { background: var(--accent-strong); text-decoration: none; }
.secondary-link {
  background: var(--surface);
  border: 1px solid var(--line);
  color: var(--text);
}
.secondary-link:hover { border-color: #bfdbfe; color: var(--accent); text-decoration: none; }

input[type="text"], textarea, select {
  transition: border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
}
input[type="text"]:focus, textarea:focus, select:focus {
  border-color: var(--accent) !important;
  outline: 0;
  box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.15);
}

.badge {
  border-radius: 6px;
  font-size: 11px;
  font-weight: 700;
  padding: 2px 8px;
}
.badge-general { background: #e2e8f0; color: #475569; }
.badge-source { background: #e0f2fe; color: #0369a1; }
.badge-entity { background: #f3e8ff; color: #6b21a8; }
.badge-concept { background: #e0faf2; color: #039855; }
.badge-topic { background: #fef3c7; color: #b45309; }

.empty-list {
  color: var(--muted);
  font-style: italic;
  font-size: 13px;
  list-style-type: none;
  padding: 6px 12px;
}

.section-heading {
  align-items: flex-start;
  display: flex;
  gap: 18px;
  justify-content: space-between;
  margin-bottom: 16px;
}
.section-heading h1, .section-heading h2 { margin: 0 0 8px; }

.file-list { list-style: none; margin: 0; max-height: 360px; overflow: auto; padding: 0; }
.file-list li { border-top: 1px solid var(--line); padding: 8px 0; }
.file-list li:first-child { border-top: 0; }
.knowledge-graph {
  background:
    radial-gradient(circle at center, rgba(37, 99, 235, .04), transparent 20rem),
    #fffdf8;
  border: 1px solid var(--line);
  border-radius: 16px;
  display: block;
  height: auto;
  max-height: 520px;
  width: 100%;
}
.graph-edge {
  stroke: #c9b99a;
  stroke-width: 1.5;
  marker-end: url(#arrow);
}
.graph-node circle {
  fill: #ffffff;
  stroke: var(--accent);
  stroke-width: 2.5;
  filter: drop-shadow(0 8px 14px rgba(37, 99, 235, .12));
}
.graph-node text {
  fill: var(--text);
  font-size: 12px;
  text-anchor: middle;
}
.graph-node:hover circle { fill: #eff6ff; stroke: var(--accent-strong); }
.edge-table { display: grid; gap: 0; }
.edge-row {
  align-items: center;
  border-top: 1px solid var(--line);
  display: grid;
  gap: 12px;
  grid-template-columns: minmax(0, 1fr) auto minmax(0, 1fr);
  padding: 10px 0;
}
.edge-row:first-child { border-top: 0; }
.edge-row a, .edge-row span { overflow-wrap: anywhere; }
.missing-link { color: #b45309; }
.details { display: grid; gap: 8px 18px; grid-template-columns: 120px minmax(0, 1fr); }
.details dt { color: var(--muted); }
.details dd { margin: 0; overflow-wrap: anywhere; }

.content {
  background: var(--card);
}
.content h1, .content h2, .content h3 { line-height: 1.2; margin-top: 1.6em; }
.content p { line-height: 1.8; margin-bottom: 1.2em; }
.content pre { background: #f8fafc; border: 1px solid var(--line); border-radius: 12px; overflow: auto; padding: 16px; }
.content code, code { background: #f8fafc; border: 1px solid #eef2f7; border-radius: 6px; color: #334155; padding: 2px 6px; font-size: 14px; }
.content pre code { background: transparent; border: 0; padding: 0; font-size: 14px; }

.wikilink {
  color: var(--accent);
  font-weight: 600;
  border-bottom: 1px dashed var(--accent);
  padding: 0 2px;
}
.wikilink:hover {
  background: rgba(37, 99, 235, .05);
  border-bottom-style: solid;
  text-decoration: none;
}
.missing-wikilink {
  color: #b45309;
  border-bottom-color: #b45309;
  cursor: not-allowed;
}

.error { color: #b42318; }
.hint, .muted { color: var(--muted); }

.app-footer {
  margin-top: 48px;
  border-top: 1px solid var(--line);
  padding-top: 16px;
  text-align: center;
  font-size: 13px;
  color: var(--muted);
}

@media (max-width: 900px) {
  .app-body {
    flex-direction: column;
  }
  .app-sidebar {
    position: static;
    width: 100%;
    height: auto;
    border-right: none;
    border-bottom: 1px solid var(--line);
  }
  .app-content {
    margin-left: 0;
    max-width: 100%;
    padding: 24px 16px;
  }
}
`
