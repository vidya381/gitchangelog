package parser

import (
	"regexp"
	"strings"

	"github.com/google/go-github/v72/github"
)

var conventionalRe = regexp.MustCompile(`^(\w+)(\(.+\))?(!)?:\s(.+)`)

// ParsedCommit holds the structured data extracted from a single commit message.
type ParsedCommit struct {
	Type        string
	Scope       string
	Description string
	Breaking    bool
	SHA         string
}

// ParsedChangelog groups commits by category after parsing.
type ParsedChangelog struct {
	Features      []ParsedCommit
	BugFixes      []ParsedCommit
	Docs          []ParsedCommit
	Chores        []ParsedCommit
	Breaking      []ParsedCommit
	Uncategorized []ParsedCommit
}

// ParseCommits takes a list of commits from the GitHub API and sorts them
// into categories based on the Conventional Commits format.
// Commits that don't follow the format go into Uncategorized.
func ParseCommits(commits []*github.RepositoryCommit) ParsedChangelog {
	var result ParsedChangelog

	for _, c := range commits {
		fullMessage := c.Commit.GetMessage()
		subject := strings.SplitN(fullMessage, "\n", 2)[0]
		sha := c.GetSHA()

		matches := conventionalRe.FindStringSubmatch(subject)
		if matches == nil {
			result.Uncategorized = append(result.Uncategorized, ParsedCommit{
				Description: subject,
				SHA:         sha,
			})
			continue
		}

		pc := ParsedCommit{
			Type:        matches[1],
			Scope:       strings.Trim(matches[2], "()"),
			Breaking:    matches[3] == "!" || strings.Contains(fullMessage, "BREAKING CHANGE"),
			Description: matches[4],
			SHA:         sha,
		}

		if pc.Breaking {
			result.Breaking = append(result.Breaking, pc)
		}

		switch pc.Type {
		case "feat":
			result.Features = append(result.Features, pc)
		case "fix":
			result.BugFixes = append(result.BugFixes, pc)
		case "docs":
			result.Docs = append(result.Docs, pc)
		default:
			result.Chores = append(result.Chores, pc)
		}
	}

	return result
}
