// Package web provides the local nemo-knows browser console: an HTTP server
// that renders the wiki, exposes a knowledge graph, lets users import new raw
// sources, and triggers the AI ingest pipeline through an injected Pipeline.
//
// The package is self-contained — templates and CSS are embedded via
// embed.FS — and depends on the caller for the heavy ingest operations so the
// same web layer can be embedded into cmd/nemo (in-process) or driven from a
// standalone launcher binary (cmd/nemo-web).
package web

import (
	"embed"
	"fmt"
	"html"
	"html/template"
	"io"
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
	"unicode/utf8"

	"github.com/huic/nemo-knows/internal/config"
)

const maxMarkdownUploadBytes = 4 << 20

//go:embed templates/*.html
var templatesFS embed.FS

//go:embed static
var staticFS embed.FS

// Server holds the running web console state. Construct one with NewServer
// and either call Run (this package's helper) or wire its handlers into an
// existing mux.
type Server struct {
	cfg      config.Config
	pipeline Pipeline

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
	Title           string
	CFG             config.Config
	Sources         []string
	WikiPages       []string
	WikiSources     []string
	WikiEntities    []string
	WikiConcepts    []string
	WikiTopics      []string
	Job             *webJob
	Path            string
	Content         template.HTML
	Graph           template.HTML
	Edges           []webGraphEdge
	GraphNote       string
	Error           string
	Success         string
	SidebarSources  []sidebarItem
	SidebarEntities []sidebarItem
	SidebarConcepts []sidebarItem
	SidebarTopics   []sidebarItem
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

var webTemplates = template.Must(template.New("web").Funcs(template.FuncMap{
	"since": func(t time.Time) string {
		if t.IsZero() {
			return ""
		}
		return time.Since(t).Round(time.Second).String()
	},
	"titleForPath": titleForMarkdownFile,
}).ParseFS(templatesFS, "templates/*.html"))

var markdownFileNameRE = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]*\.md$`)

// NewServer constructs a Server with the given configuration and pipeline.
// pipeline may be nil; if it is, /run requests respond with 503.
func NewServer(cfg config.Config, pipeline Pipeline) *Server {
	return &Server{cfg: cfg, pipeline: pipeline}
}

// Handler returns an http.Handler that serves the entire web console under
// the root path. The caller can mount this directly on its preferred address.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleHome)
	mux.HandleFunc("/run", s.handleRun)
	mux.HandleFunc("/job", s.handleJob)
	mux.HandleFunc("/build", s.handleBuild)
	mux.HandleFunc("/graph", s.handleGraph)
	mux.HandleFunc("/view", s.handleView)

	// Static assets are served from the embedded FS under /static/.
	staticSub, err := fs.Sub(staticFS, "static")
	if err == nil {
		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticSub))))
	}
	return mux
}

// Run is a convenience wrapper that binds the console to the given TCP
// address. It blocks until the server exits.
func Run(addr string, cfg config.Config, pipeline Pipeline) error {
	srv := NewServer(cfg, pipeline)
	fmt.Fprintf(os.Stderr, "nemo web console listening on http://%s\n", addr)
	return http.ListenAndServe(addr, srv.Handler())
}

// ---------------------------------------------------------------------------
// Page-data assembly
// ---------------------------------------------------------------------------

func (s *Server) makePageData(currentPath string, title string) webPageData {
	sources, entities, concepts, topics := s.getSidebarItems(currentPath)
	return webPageData{
		Title:           title,
		CFG:             s.cfg,
		Sources:         mustListTextFiles("wiki/sources"),
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

func (s *Server) getSidebarItems(currentPath string) ([]sidebarItem, []sidebarItem, []sidebarItem, []sidebarItem) {
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

// ---------------------------------------------------------------------------
// HTTP handlers
// ---------------------------------------------------------------------------

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) handleRun(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if s.pipeline == nil {
		http.Error(w, "ingest pipeline is not available in this build", http.StatusServiceUnavailable)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	source, err := cleanProjectPath(r.FormValue("source"))
	if err != nil || !strings.HasPrefix(source, "wiki/sources/") {
		http.Error(w, "source must be a project-relative path under wiki/sources/", http.StatusBadRequest)
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
	if err := s.startJob(source, profile, provider); err != nil {
		status := http.StatusBadRequest
		if strings.Contains(err.Error(), "already in progress") {
			status = http.StatusConflict
		}
		http.Error(w, err.Error(), status)
		return
	}
	http.Redirect(w, r, "/job", http.StatusSeeOther)
}

func (s *Server) startJob(source string, profile string, provider string) error {
	if s.pipeline == nil {
		return fmt.Errorf("ingest pipeline is not available in this build")
	}
	cfg, err := config.ForProfileWithProvider(profile, provider)
	if err != nil {
		return err
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
		return fmt.Errorf("another nemo run is already in progress")
	}
	s.job = job
	s.mu.Unlock()

	go s.runJob(job, cfg)
	return nil
}

func (s *Server) handleJob(w http.ResponseWriter, r *http.Request) {
	s.render(w, "job", s.makePageData("", "当前任务"))
}

func (s *Server) handleGraph(w http.ResponseWriter, r *http.Request) {
	data := s.makePageData("", "知识图谱")
	data.Graph = renderWikiGraph(48)
	data.Edges = wikiGraphEdges(80)
	data.GraphNote = wikiGraphNote()
	s.render(w, "graph", data)
}

func (s *Server) handleView(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) handleBuild(w http.ResponseWriter, r *http.Request) {
	var errMsg string
	var successMsg string

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		if action == "save_source" {
			name, content, err := markdownUploadFromRequest(r, time.Now())
			if err != nil {
				errMsg = err.Error()
			} else {
				targetDir := filepath.Join("wiki", "sources")
				targetPath := filepath.Join(targetDir, name)
				if err := os.MkdirAll(targetDir, 0o755); err != nil {
					errMsg = "创建 wiki/sources/ 目录失败: " + err.Error()
				} else if _, err := os.Stat(targetPath); err == nil {
					errMsg = "同名 Markdown 已存在，请换一个文件名"
				} else if !os.IsNotExist(err) {
					errMsg = "检查文件失败: " + err.Error()
				} else if err := os.WriteFile(targetPath, []byte(content), 0o644); err != nil {
					errMsg = "保存 Markdown 失败: " + err.Error()
				} else if err := s.startJob(filepath.ToSlash(targetPath), "stable", ""); err != nil {
					errMsg = "Markdown 已保存，但启动编译失败: " + err.Error()
				} else {
					http.Redirect(w, r, "/job", http.StatusSeeOther)
					return
				}
			}
		}
	}

	data := s.makePageData("/build", "知识构建")
	data.Error = errMsg
	data.Success = successMsg
	s.render(w, "build", data)
}

func markdownUploadFromRequest(r *http.Request, now time.Time) (string, string, error) {
	if err := r.ParseMultipartForm(maxMarkdownUploadBytes); err != nil {
		return "", "", fmt.Errorf("解析上传内容失败: %w", err)
	}
	file, header, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		content, err := readLimitedMarkdown(file)
		if err != nil {
			return "", "", err
		}
		return validateMarkdownUpload(header.Filename, content, now)
	}
	if err != http.ErrMissingFile {
		return "", "", fmt.Errorf("读取上传文件失败: %w", err)
	}
	return validateMarkdownUpload("", r.FormValue("content"), now)
}

func readLimitedMarkdown(reader io.Reader) (string, error) {
	limited := io.LimitReader(reader, maxMarkdownUploadBytes+1)
	content, err := io.ReadAll(limited)
	if err != nil {
		return "", fmt.Errorf("读取 Markdown 失败: %w", err)
	}
	if len(content) > maxMarkdownUploadBytes {
		return "", fmt.Errorf("Markdown 文件不能超过 4 MiB")
	}
	return string(content), nil
}

func validateMarkdownUpload(name string, content string, now time.Time) (string, string, error) {
	name = strings.TrimSpace(name)
	content = strings.TrimSpace(content)
	if content == "" {
		return "", "", fmt.Errorf("请上传 .md 文件或粘贴 Markdown 内容")
	}
	if !utf8.ValidString(content) {
		return "", "", fmt.Errorf("内容必须是有效 UTF-8 文本")
	}
	if strings.ContainsRune(content, '\x00') {
		return "", "", fmt.Errorf("内容不能包含 NUL 控制字符")
	}
	lower := strings.ToLower(content)
	if strings.Contains(lower, "<script") {
		return "", "", fmt.Errorf("内容不能包含 script 标签")
	}
	name, err := normalizeMarkdownName(name, content, now)
	if err != nil {
		return "", "", err
	}
	return name, content + "\n", nil
}

func normalizeMarkdownName(name string, content string, now time.Time) (string, error) {
	if name == "" {
		if title := firstMarkdownTitle(content); title != "" {
			name = slugFromFilename(title) + ".md"
		} else {
			if now.IsZero() {
				now = time.Now()
			}
			name = "source-" + now.Format("20060102-150405") + ".md"
		}
	}
	if name != filepath.Base(name) || strings.ContainsAny(name, `/\`) {
		return "", fmt.Errorf("文件名不能包含目录路径")
	}
	name = filepath.Base(name)
	if name == "." || name == string(filepath.Separator) || name == "" {
		return "", fmt.Errorf("文件名不能为空")
	}
	if !strings.HasSuffix(strings.ToLower(name), ".md") {
		return "", fmt.Errorf("只支持 .md 文件")
	}
	slug := slugFromFilename(name)
	if slug == "" {
		return "", fmt.Errorf("无法从文件名生成有效标识")
	}
	name = slug + ".md"
	if !markdownFileNameRE.MatchString(name) {
		return "", fmt.Errorf("文件名只能包含英文字母、数字、点、下划线和连字符，并且必须是 .md 文件")
	}
	return name, nil
}

func firstMarkdownTitle(content string) string {
	trimmed := strings.TrimSpace(content)
	for _, line := range strings.Split(trimmed, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "# ") {
			return strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "# "))
		}
	}
	return ""
}

// ---------------------------------------------------------------------------
// Job execution
// ---------------------------------------------------------------------------

func (s *Server) runJob(job *webJob, cfg config.Config) {
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
	err = s.pipeline.RunBundle(job.Source, filepath.FromSlash(job.BundleDir), cfg)
	if err == nil {
		s.updateJob(job, func(j *webJob) {
			j.Step = "reviewing bundle"
		})
		err = s.pipeline.RunReviewBundle(filepath.FromSlash(job.BundleDir), filepath.Join(filepath.FromSlash(job.BundleDir), "apply-plan.md"))
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

// acquireNemoLock takes the same /tmp/nemo-ingest.lock used by the CLI so a
// single-instance invariant is preserved across both entry points.
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

func (s *Server) updateJob(job *webJob, update func(*webJob)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.job == job {
		update(s.job)
	}
}

func (s *Server) currentJob() *webJob {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.job == nil {
		return nil
	}
	cp := *s.job
	cp.Files = append([]string(nil), s.job.Files...)
	return &cp
}

func (s *Server) render(w http.ResponseWriter, name string, data webPageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := webTemplates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ---------------------------------------------------------------------------
// File listing helpers
// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------
// Knowledge graph
// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------
// Markdown / wikilink rendering
// ---------------------------------------------------------------------------

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

// wikilinkRE matches a [[slug]] or [[slug|label]] reference in markdown
// content. The raw [[...]] syntax is the on-disk source-of-truth used by the
// ingest pipeline; the web layer only strips the brackets at display time.
var wikilinkRE = regexp.MustCompile(`\[\[([^\]|#]+)(?:[|#][^\]]*)?\]\]`)

// candidateWikilinkRE is the broader pattern used for graph edge extraction.
// It must match the same shape as the renderer.
var candidateWikilinkRE = regexp.MustCompile(`\[\[([^\]|#]+)(?:[|#][^\]]*)?\]\]`)

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

// renderInline rewrites the inline portion of a line: it preserves the raw
// [[...]] tokens on disk but renders them in HTML as a clean link with just
// the label text (no visible brackets). Unresolved targets become a subtle
// dotted span so missing references are still discoverable.
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
			return fmt.Sprintf(`<a href="/view?path=%s" class="wikilink">%s</a>`, path, label)
		}
		return fmt.Sprintf(`<span class="wikilink-missing" title="缺少对应词条页面">%s</span>`, label)
	})
}

// ---------------------------------------------------------------------------
// Slug / title helpers (kept local to avoid coupling the web layer to
// cmd/nemo. They mirror the equivalents in cmd/nemo/main.go that the ingest
// pipeline uses on disk.)
// ---------------------------------------------------------------------------

func titleFromSlug(slug string) string {
	parts := strings.Split(slug, "-")
	for i, part := range parts {
		if acronym := knownTitleAcronym(part); acronym != "" {
			parts[i] = acronym
			continue
		}
		if part != "" {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, " ")
}

func knownTitleAcronym(part string) string {
	switch strings.ToLower(part) {
	case "llm":
		return "LLM"
	case "rag":
		return "RAG"
	case "mvp":
		return "MVP"
	default:
		return ""
	}
}

func slugFromWikiPath(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

func slugFromLink(link string) string {
	slug := strings.TrimSpace(link)
	slug = strings.TrimSuffix(slug, filepath.Ext(slug))
	slug = strings.ToLower(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}

func slugFromFilename(name string) string {
	base := strings.TrimSuffix(name, filepath.Ext(name))
	base = strings.ToLower(base)
	var b strings.Builder
	lastHyphen := false
	for _, r := range base {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
			lastHyphen = false
		default:
			if !lastHyphen && b.Len() > 0 {
				b.WriteByte('-')
				lastHyphen = true
			}
		}
	}
	return strings.TrimRight(b.String(), "-")
}
