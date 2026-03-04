package formatter

import (
	"fmt"
	"strings"

	"github.com/vidya381/gitchangelog/internal/parser"
)

// ToMarkdown formats a parsed changelog into a Markdown string.
// Each commit links back to its commit page on GitHub using repoURL.
// Sections with no commits are omitted.
func ToMarkdown(cl parser.ParsedChangelog, repoURL string) string {
	var b strings.Builder

	b.WriteString("## What's Changed\n")

	writeSection(&b, "Breaking Changes", cl.Breaking, repoURL)
	writeSection(&b, "Features", cl.Features, repoURL)
	writeSection(&b, "Bug Fixes", cl.BugFixes, repoURL)
	writeSection(&b, "Documentation", cl.Docs, repoURL)
	writeSection(&b, "Chores", cl.Chores, repoURL)
	writeSection(&b, "Other", cl.Uncategorized, repoURL)

	return b.String()
}

func writeSection(b *strings.Builder, title string, commits []parser.ParsedCommit, repoURL string) {
	if len(commits) == 0 {
		return
	}

	fmt.Fprintf(b, "\n### %s\n", title)

	for _, c := range commits {
		line := "- "
		if c.Scope != "" {
			line += "**" + c.Scope + ":** "
		}
		line += c.Description

		if c.SHA != "" {
			short := c.SHA
			if len(short) > 7 {
				short = short[:7]
			}
			if repoURL != "" {
				line += fmt.Sprintf(" ([%s](%s/commit/%s))", short, strings.TrimRight(repoURL, "/"), c.SHA)
			} else {
				line += fmt.Sprintf(" (%s)", short)
			}
		}

		b.WriteString(line + "\n")
	}
}
