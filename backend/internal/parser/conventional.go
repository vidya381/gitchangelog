package parser

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
