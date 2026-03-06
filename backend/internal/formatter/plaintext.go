package formatter

import (
	"fmt"
	"strings"

	"github.com/vidya381/gitchangelog/internal/parser"
)

// ToPlainText formats a parsed changelog into a plain text string.
// Same structure as ToMarkdown but without headers, links, or formatting.
func ToPlainText(cl parser.ParsedChangelog) string {
	var b strings.Builder

	b.WriteString("What's Changed\n")

	writePlainSection(&b, "Breaking Changes", cl.Breaking)
	writePlainSection(&b, "Features", cl.Features)
	writePlainSection(&b, "Bug Fixes", cl.BugFixes)
	writePlainSection(&b, "Documentation", cl.Docs)
	writePlainSection(&b, "Chores", cl.Chores)
	writePlainSection(&b, "Other", cl.Uncategorized)

	return b.String()
}

func writePlainSection(b *strings.Builder, title string, commits []parser.ParsedCommit) {
	if len(commits) == 0 {
		return
	}

	fmt.Fprintf(b, "\n%s\n", title)

	for _, c := range commits {
		line := "- "
		if c.Scope != "" {
			line += c.Scope + ": "
		}
		line += c.Description

		if c.SHA != "" {
			short := c.SHA
			if len(short) > 7 {
				short = short[:7]
			}
			line += fmt.Sprintf(" (%s)", short)
		}

		b.WriteString(line + "\n")
	}
}
