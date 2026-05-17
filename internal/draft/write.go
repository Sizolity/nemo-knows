package draft

import "path/filepath"

type OutputPaths struct {
	Raw     string
	Cleaned string
}

// PathsFor derives the raw and cleaned draft output paths.
//
// The cleanedOutput path is the final Markdown draft path requested by the
// caller.
//
// The returned Raw path stores the complete llama.cpp output next to the
// cleaned draft, using the same base name with a .raw.txt suffix.
func PathsFor(cleanedOutput string) OutputPaths {
	ext := filepath.Ext(cleanedOutput)
	rawOutput := cleanedOutput + ".raw.txt"
	if ext != "" {
		rawOutput = cleanedOutput[:len(cleanedOutput)-len(ext)] + ".raw.txt"
	}

	return OutputPaths{
		Raw:     rawOutput,
		Cleaned: cleanedOutput,
	}
}
