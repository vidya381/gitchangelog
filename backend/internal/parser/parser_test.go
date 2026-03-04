package parser

import (
	"testing"

	"github.com/google/go-github/v72/github"
)

func makeCommit(sha, message string) *github.RepositoryCommit {
	return &github.RepositoryCommit{
		SHA: &sha,
		Commit: &github.Commit{
			Message: &message,
		},
	}
}

func TestParseCommits(t *testing.T) {
	tests := []struct {
		name    string
		message string
		check   func(t *testing.T, result ParsedChangelog)
	}{
		{
			name:    "feat commit goes to Features",
			message: "feat: add login page",
			check: func(t *testing.T, result ParsedChangelog) {
				if len(result.Features) != 1 {
					t.Fatalf("expected 1 feature, got %d", len(result.Features))
				}
				if result.Features[0].Description != "add login page" {
					t.Errorf("unexpected description: %s", result.Features[0].Description)
				}
			},
		},
		{
			name:    "fix commit with scope goes to BugFixes",
			message: "fix(api): handle nil response",
			check: func(t *testing.T, result ParsedChangelog) {
				if len(result.BugFixes) != 1 {
					t.Fatalf("expected 1 bug fix, got %d", len(result.BugFixes))
				}
				if result.BugFixes[0].Scope != "api" {
					t.Errorf("expected scope 'api', got '%s'", result.BugFixes[0].Scope)
				}
			},
		},
		{
			name:    "docs commit goes to Docs",
			message: "docs: update README",
			check: func(t *testing.T, result ParsedChangelog) {
				if len(result.Docs) != 1 {
					t.Fatalf("expected 1 doc, got %d", len(result.Docs))
				}
			},
		},
		{
			name:    "chore commit goes to Chores",
			message: "chore: update dependencies",
			check: func(t *testing.T, result ParsedChangelog) {
				if len(result.Chores) != 1 {
					t.Fatalf("expected 1 chore, got %d", len(result.Chores))
				}
			},
		},
		{
			name:    "breaking change using ! goes to Breaking and Features",
			message: "feat!: remove deprecated endpoint",
			check: func(t *testing.T, result ParsedChangelog) {
				if len(result.Breaking) != 1 {
					t.Fatalf("expected 1 breaking change, got %d", len(result.Breaking))
				}
				if len(result.Features) != 1 {
					t.Fatalf("expected 1 feature, got %d", len(result.Features))
				}
			},
		},
		{
			name:    "breaking change using footer goes to Breaking and Features",
			message: "feat: remove deprecated endpoint\n\nBREAKING CHANGE: /v1 routes removed",
			check: func(t *testing.T, result ParsedChangelog) {
				if len(result.Breaking) != 1 {
					t.Fatalf("expected 1 breaking change, got %d", len(result.Breaking))
				}
				if len(result.Features) != 1 {
					t.Fatalf("expected 1 feature, got %d", len(result.Features))
				}
			},
		},
		{
			name:    "non-conventional commit goes to Uncategorized",
			message: "fixed a thing",
			check: func(t *testing.T, result ParsedChangelog) {
				if len(result.Uncategorized) != 1 {
					t.Fatalf("expected 1 uncategorized, got %d", len(result.Uncategorized))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commits := []*github.RepositoryCommit{makeCommit("abc123", tt.message)}
			result := ParseCommits(commits)
			tt.check(t, result)
		})
	}
}
