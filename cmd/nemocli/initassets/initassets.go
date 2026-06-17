// Package initassets exposes the static templates that `nemocli init`
// copies into a freshly scaffolded $WIKI_ROOT.
//
// Templates are embedded with go:embed so the deployed binary needs no
// sibling files at runtime. The AGENTS.md template was authored during
// the R6 sync step and tracked at templates/wiki-agents.md; the
// .env.example template lives at templates/env.example and only contains
// blank placeholder values — never real secrets.
package initassets

import _ "embed"

//go:embed templates/wiki-agents.md
var wikiAgentsTemplate []byte

//go:embed templates/env.example
var envExampleTemplate []byte

// WikiAgentsTemplate returns the AGENTS.md contract written to
// $WIKI_ROOT/wiki/AGENTS.md during init.
func WikiAgentsTemplate() []byte {
	out := make([]byte, len(wikiAgentsTemplate))
	copy(out, wikiAgentsTemplate)
	return out
}

// EnvExampleTemplate returns the .env.example body written to
// $WIKI_ROOT/.env.example during init.
func EnvExampleTemplate() []byte {
	out := make([]byte, len(envExampleTemplate))
	copy(out, envExampleTemplate)
	return out
}
