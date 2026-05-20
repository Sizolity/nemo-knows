package chunking

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const DefaultMaxChunkChars = 18000

var (
	markdownHeadingRE = regexp.MustCompile(`^(#{1,6})\s+(.+?)\s*$`)
	numberedHeadingRE = regexp.MustCompile(`^([0-9]+(?:\.[0-9]+)*)\.?\s+(.+)$`)
	namedHeadingRE    = regexp.MustCompile(`^(Appendix [A-Z]: .+|Application Porting Guide|Changes to .+|Auxiliary Function .+|Other Issues|Summary of .+)$`)
)

type Plan struct {
	SourcePath string  `json:"source_path"`
	Chunks     []Chunk `json:"chunks"`
}

type Chunk struct {
	Index               int        `json:"index"`
	HeadingPath         []string   `json:"heading_path"`
	SectionHeadingPaths [][]string `json:"heading_paths"`
	StartLine           int        `json:"start_line"`
	EndLine             int        `json:"end_line"`
	Text                string     `json:"-"`
}

type section struct {
	headingPath []string
	startLine   int
	endLine     int
	lines       []string
}

func PlanSource(sourcePath string, content string, maxChars int) Plan {
	if maxChars <= 0 {
		maxChars = DefaultMaxChunkChars
	}
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
	sections := sectionsFromLines(lines)
	chunks := chunksFromSections(sections, maxChars)
	return Plan{SourcePath: sourcePath, Chunks: chunks}
}

func (p Plan) OutlineMarkdown() string {
	var b strings.Builder
	b.WriteString("# Chunk Outline\n\n")
	b.WriteString("Source: `")
	b.WriteString(p.SourcePath)
	b.WriteString("`\n\n")
	for _, chunk := range p.Chunks {
		b.WriteString(fmt.Sprintf("## Chunk %02d\n\n", chunk.Index))
		b.WriteString(fmt.Sprintf("- Lines: %d-%d\n", chunk.StartLine, chunk.EndLine))
		b.WriteString("- Primary heading path: ")
		if len(chunk.HeadingPath) == 0 {
			b.WriteString("(document root)\n")
		} else {
			b.WriteString(strings.Join(chunk.HeadingPath, " > "))
			b.WriteString("\n")
		}
		b.WriteString("- Heading coverage:\n")
		for _, path := range chunkHeadingCoverage(chunk) {
			b.WriteString("  - ")
			if len(path) == 0 {
				b.WriteString("(document root)")
			} else {
				b.WriteString(strings.Join(path, " > "))
			}
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (p Plan) IndexJSON() ([]byte, error) {
	type indexChunk struct {
		Index               int        `json:"index"`
		HeadingPath         []string   `json:"heading_path"`
		SectionHeadingPaths [][]string `json:"heading_paths"`
		StartLine           int        `json:"start_line"`
		EndLine             int        `json:"end_line"`
		Chars               int        `json:"chars"`
	}
	out := struct {
		SourcePath string       `json:"source_path"`
		Chunks     []indexChunk `json:"chunks"`
	}{SourcePath: p.SourcePath}
	for _, chunk := range p.Chunks {
		out.Chunks = append(out.Chunks, indexChunk{
			Index:               chunk.Index,
			HeadingPath:         append([]string(nil), chunk.HeadingPath...),
			SectionHeadingPaths: cloneHeadingPaths(chunkHeadingCoverage(chunk)),
			StartLine:           chunk.StartLine,
			EndLine:             chunk.EndLine,
			Chars:               len(chunk.Text),
		})
	}
	return json.MarshalIndent(out, "", "  ")
}

func sectionsFromLines(lines []string) []section {
	var sections []section
	stack := []string{"Document"}
	current := section{headingPath: append([]string(nil), stack...), startLine: 1}

	flush := func(endLine int) {
		current.endLine = endLine
		if len(current.lines) > 0 || len(sections) == 0 {
			sections = append(sections, current)
		}
	}

	for i, line := range lines {
		lineNo := i + 1
		if level, title, ok := headingAt(lines, i); ok && lineNo != 1 {
			flush(lineNo - 1)
			if level < 1 {
				level = 1
			}
			if level > len(stack) {
				level = len(stack) + 1
			}
			stack = stack[:level-1]
			stack = append(stack, title)
			current = section{headingPath: append([]string(nil), stack...), startLine: lineNo}
		}
		current.lines = append(current.lines, line)
	}
	flush(len(lines))
	return sections
}

func headingAt(lines []string, index int) (int, string, bool) {
	line := lines[index]
	trimmed := strings.TrimSpace(line)
	if trimmed == "" || len(trimmed) > 140 {
		return 0, "", false
	}
	if match := markdownHeadingRE.FindStringSubmatch(trimmed); len(match) == 3 {
		return len(match[1]), cleanHeading(match[2]), true
	}
	if match := numberedHeadingRE.FindStringSubmatch(trimmed); len(match) == 3 {
		if !surroundedByBlank(lines, index) {
			return 0, "", false
		}
		level := strings.Count(match[1], ".") + 1
		return level, cleanHeading(trimmed), true
	}
	if namedHeadingRE.MatchString(trimmed) {
		if !surroundedByBlank(lines, index) {
			return 0, "", false
		}
		return 2, cleanHeading(trimmed), true
	}
	return 0, "", false
}

func surroundedByBlank(lines []string, index int) bool {
	prevBlank := index == 0 || strings.TrimSpace(lines[index-1]) == ""
	nextBlank := index == len(lines)-1 || strings.TrimSpace(lines[index+1]) == ""
	return prevBlank && nextBlank
}

func cleanHeading(title string) string {
	return strings.Trim(strings.TrimSpace(title), "# ")
}

func chunksFromSections(sections []section, maxChars int) []Chunk {
	var chunks []Chunk
	var pending []section
	pendingChars := 0

	flush := func() {
		if len(pending) == 0 {
			return
		}
		chunks = append(chunks, makeChunk(len(chunks)+1, pending))
		pending = nil
		pendingChars = 0
	}

	for _, sec := range sections {
		secText := strings.Join(sec.lines, "\n")
		if len(secText) > maxChars {
			flush()
			chunks = append(chunks, splitOversizedSection(sec, maxChars, len(chunks)+1)...)
			continue
		}
		if pendingChars > 0 && pendingChars+len(secText) > maxChars {
			flush()
		}
		pending = append(pending, sec)
		pendingChars += len(secText)
	}
	flush()
	return chunks
}

func makeChunk(index int, sections []section) Chunk {
	first := sections[0]
	last := sections[len(sections)-1]
	var b strings.Builder
	for i, sec := range sections {
		if i > 0 {
			b.WriteString("\n\n")
		}
		b.WriteString(strings.Join(sec.lines, "\n"))
	}
	return Chunk{
		Index:               index,
		HeadingPath:         append([]string(nil), first.headingPath...),
		SectionHeadingPaths: headingPathsFromSections(sections),
		StartLine:           first.startLine,
		EndLine:             last.endLine,
		Text:                strings.TrimSpace(b.String()),
	}
}

func headingPathsFromSections(sections []section) [][]string {
	var paths [][]string
	for _, sec := range sections {
		paths = appendUniqueHeadingPath(paths, sec.headingPath)
	}
	return paths
}

func chunkHeadingCoverage(chunk Chunk) [][]string {
	if len(chunk.SectionHeadingPaths) > 0 {
		return chunk.SectionHeadingPaths
	}
	if len(chunk.HeadingPath) == 0 {
		return nil
	}
	return [][]string{append([]string(nil), chunk.HeadingPath...)}
}

func cloneHeadingPaths(paths [][]string) [][]string {
	out := make([][]string, 0, len(paths))
	for _, path := range paths {
		out = append(out, append([]string(nil), path...))
	}
	return out
}

func appendUniqueHeadingPath(paths [][]string, candidate []string) [][]string {
	for _, existing := range paths {
		if sameStringSlice(existing, candidate) {
			return paths
		}
	}
	return append(paths, append([]string(nil), candidate...))
}

func sameStringSlice(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func splitOversizedSection(sec section, maxChars int, startIndex int) []Chunk {
	paragraphs := paragraphsWithLines(sec)
	var chunks []Chunk
	var current []paragraph
	currentChars := 0
	flush := func() {
		if len(current) == 0 {
			return
		}
		chunks = append(chunks, chunkFromParagraphs(startIndex+len(chunks), sec.headingPath, current))
		current = nil
		currentChars = 0
	}
	for _, para := range paragraphs {
		if len(para.text) > maxChars {
			flush()
			chunks = append(chunks, splitParagraph(sec.headingPath, para, maxChars, startIndex+len(chunks))...)
			continue
		}
		if currentChars > 0 && currentChars+len(para.text) > maxChars {
			flush()
		}
		current = append(current, para)
		currentChars += len(para.text)
	}
	flush()
	return chunks
}

type paragraph struct {
	startLine int
	endLine   int
	text      string
}

func paragraphsWithLines(sec section) []paragraph {
	var out []paragraph
	var buf []string
	start := sec.startLine
	flush := func(end int) {
		if len(buf) == 0 {
			return
		}
		out = append(out, paragraph{startLine: start, endLine: end, text: strings.Join(buf, "\n")})
		buf = nil
	}
	for i, line := range sec.lines {
		lineNo := sec.startLine + i
		if strings.TrimSpace(line) == "" {
			flush(lineNo - 1)
			start = lineNo + 1
			continue
		}
		if len(buf) == 0 {
			start = lineNo
		}
		buf = append(buf, line)
	}
	flush(sec.endLine)
	return out
}

func chunkFromParagraphs(index int, headingPath []string, paragraphs []paragraph) Chunk {
	first := paragraphs[0]
	last := paragraphs[len(paragraphs)-1]
	parts := make([]string, 0, len(paragraphs)+1)
	parts = append(parts, "Context: "+strings.Join(headingPath, " > "))
	for _, para := range paragraphs {
		parts = append(parts, para.text)
	}
	return Chunk{
		Index:               index,
		HeadingPath:         append([]string(nil), headingPath...),
		SectionHeadingPaths: [][]string{append([]string(nil), headingPath...)},
		StartLine:           first.startLine,
		EndLine:             last.endLine,
		Text:                strings.TrimSpace(strings.Join(parts, "\n\n")),
	}
}

func splitParagraph(headingPath []string, para paragraph, maxChars int, startIndex int) []Chunk {
	var chunks []Chunk
	parts := splitTextBySoftBoundaries(para.text, maxChars)
	for i, part := range parts {
		chunks = append(chunks, Chunk{
			Index:               startIndex + len(chunks),
			HeadingPath:         append([]string(nil), headingPath...),
			SectionHeadingPaths: [][]string{append([]string(nil), headingPath...)},
			StartLine:           para.startLine,
			EndLine:             para.endLine,
			Text: strings.TrimSpace(fmt.Sprintf(
				"Context: %s\nOversized paragraph segment: %d of %d\n\n%s",
				strings.Join(headingPath, " > "),
				i+1,
				len(parts),
				part,
			)),
		})
	}
	return chunks
}

func splitTextBySoftBoundaries(text string, maxChars int) []string {
	// Truncation strategy for pathological long inputs:
	// 1. Preserve document structure first (handled before this function by
	//    section and paragraph splitting).
	// 2. If a paragraph is still too large, preserve local coherence by
	//    cutting at line boundaries, then sentence boundaries, then word
	//    boundaries.
	// 3. Use a character cut only as the final safety valve when the source has
	//    no usable boundaries, such as a generated blob or a very long token.
	//
	// Nothing is discarded here. The model summarizes each resulting segment,
	// and later synthesis combines those local notes with the outline/index.
	if maxChars <= 0 || len(text) <= maxChars {
		return []string{strings.TrimSpace(text)}
	}

	var parts []string
	var current strings.Builder
	flush := func() {
		part := strings.TrimSpace(current.String())
		if part != "" {
			parts = append(parts, part)
		}
		current.Reset()
	}
	for _, line := range strings.SplitAfter(text, "\n") {
		if len(line) > maxChars {
			flush()
			parts = append(parts, splitSentenceOrWord(line, maxChars)...)
			continue
		}
		if current.Len() > 0 && current.Len()+len(line) > maxChars {
			flush()
		}
		current.WriteString(line)
	}
	flush()
	return parts
}

func splitSentenceOrWord(text string, maxChars int) []string {
	sentences := sentenceUnits(text)
	var parts []string
	var current strings.Builder
	flush := func() {
		part := strings.TrimSpace(current.String())
		if part != "" {
			parts = append(parts, part)
		}
		current.Reset()
	}
	for _, sentence := range sentences {
		if len(sentence) > maxChars {
			flush()
			parts = append(parts, splitAtWordBoundary(sentence, maxChars)...)
			continue
		}
		if current.Len() > 0 && current.Len()+len(sentence) > maxChars {
			flush()
		}
		current.WriteString(sentence)
	}
	flush()
	return parts
}

func sentenceUnits(text string) []string {
	var out []string
	start := 0
	for i, r := range text {
		if !isSentenceTerminal(r) {
			continue
		}
		end := i + len(string(r))
		if end < len(text) {
			next, _ := nextRune(text[end:])
			if next != 0 && !isBoundaryAfterSentence(next) {
				continue
			}
		}
		out = append(out, text[start:end])
		start = end
	}
	if start < len(text) {
		out = append(out, text[start:])
	}
	if len(out) == 0 {
		return []string{text}
	}
	return out
}

func nextRune(text string) (rune, int) {
	for i, r := range text {
		return r, i
	}
	return 0, 0
}

func isSentenceTerminal(r rune) bool {
	return r == '.' || r == '!' || r == '?' || r == ';' || r == '。' || r == '！' || r == '？' || r == '；'
}

func isBoundaryAfterSentence(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t' || r == '"' || r == '\'' || r == ')' || r == ']' || r == '}'
}

func splitAtWordBoundary(text string, maxChars int) []string {
	var parts []string
	remaining := strings.TrimSpace(text)
	for len(remaining) > maxChars {
		cut := lastWhitespaceBefore(remaining, maxChars)
		if cut < maxChars/2 {
			cut = maxChars
		}
		parts = append(parts, strings.TrimSpace(remaining[:cut]))
		remaining = strings.TrimSpace(remaining[cut:])
	}
	if remaining != "" {
		parts = append(parts, remaining)
	}
	return parts
}

func lastWhitespaceBefore(text string, limit int) int {
	if limit > len(text) {
		limit = len(text)
	}
	for i := limit - 1; i >= 0; i-- {
		if text[i] == ' ' || text[i] == '\n' || text[i] == '\t' {
			return i
		}
	}
	return -1
}

func FormatChunkForPrompt(sourcePath string, chunk Chunk, total int) string {
	var b strings.Builder
	b.WriteString("Source: ")
	b.WriteString(sourcePath)
	b.WriteString("\nChunk: ")
	b.WriteString(strconv.Itoa(chunk.Index))
	b.WriteString(" of ")
	b.WriteString(strconv.Itoa(total))
	b.WriteString("\nLines: ")
	b.WriteString(strconv.Itoa(chunk.StartLine))
	b.WriteString("-")
	b.WriteString(strconv.Itoa(chunk.EndLine))
	b.WriteString("\nHeading path: ")
	if len(chunk.HeadingPath) == 0 {
		b.WriteString("(document root)")
	} else {
		b.WriteString(strings.Join(chunk.HeadingPath, " > "))
	}
	b.WriteString("\nHeading coverage:")
	for _, path := range chunkHeadingCoverage(chunk) {
		b.WriteString("\n- ")
		if len(path) == 0 {
			b.WriteString("(document root)")
		} else {
			b.WriteString(strings.Join(path, " > "))
		}
	}
	b.WriteString("\n\nChunk text:\n")
	b.WriteString(chunk.Text)
	return b.String()
}
