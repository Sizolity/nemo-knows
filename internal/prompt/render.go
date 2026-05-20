package prompt

import "strings"

type Variables struct {
	RawSourcePath    string
	RawSourceContent string
	ConceptName      string
	SourceList       string
	SourceContent    string
	PageTitle        string
	PageKind         string
	TargetPath       string
	AllowedLinks     string
	ChunkContent     string
	ChunkNotes       string
	ChunkGroupNotes  string
	ChunkOutline     string
	ChunkIndex       string
}

// Render renders a prompt template with the given variables.
//
// The template is a string that contains the prompt template.
// The vars are the variables to substitute into the template.
//
// The returned string is the rendered prompt.
//
// The error is returned if the template is invalid or the variables are invalid.
func Render(template string, vars Variables) (string, error) {
	replacer := strings.NewReplacer(
		"{{RAW_SOURCE_PATH}}", vars.RawSourcePath,
		"{{RAW_SOURCE_CONTENT}}", vars.RawSourceContent,
		"{{CONCEPT_NAME}}", vars.ConceptName,
		"{{SOURCE_LIST}}", vars.SourceList,
		"{{SOURCE_CONTENT}}", vars.SourceContent,
		"{{PAGE_TITLE}}", vars.PageTitle,
		"{{PAGE_KIND}}", vars.PageKind,
		"{{TARGET_PATH}}", vars.TargetPath,
		"{{ALLOWED_LINKS}}", vars.AllowedLinks,
		"{{CHUNK_CONTENT}}", vars.ChunkContent,
		"{{CHUNK_NOTES}}", vars.ChunkNotes,
		"{{CHUNK_GROUP_NOTES}}", vars.ChunkGroupNotes,
		"{{CHUNK_OUTLINE}}", vars.ChunkOutline,
		"{{CHUNK_INDEX}}", vars.ChunkIndex,
		"RAW_SOURCE_CONTENT", vars.RawSourceContent,
		"RAW_SOURCE_PATH", vars.RawSourcePath,
		"CONCEPT_NAME", vars.ConceptName,
		"SOURCE_LIST", vars.SourceList,
		"SOURCE_CONTENT", vars.SourceContent,
		"PAGE_TITLE", vars.PageTitle,
		"PAGE_KIND", vars.PageKind,
		"TARGET_PATH", vars.TargetPath,
		"ALLOWED_LINKS", vars.AllowedLinks,
		"CHUNK_CONTENT", vars.ChunkContent,
		"CHUNK_NOTES", vars.ChunkNotes,
		"CHUNK_GROUP_NOTES", vars.ChunkGroupNotes,
		"CHUNK_OUTLINE", vars.ChunkOutline,
		"CHUNK_INDEX", vars.ChunkIndex,
	)

	return replacer.Replace(template), nil
}
