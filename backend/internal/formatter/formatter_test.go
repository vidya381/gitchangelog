package formatter

import (
	"strings"
	"testing"

	"github.com/vidya381/gitchangelog/internal/parser"
)

func TestToMarkdown_sections(t *testing.T) {
	cl := parser.ParsedChangelog{
		Features: []parser.ParsedCommit{
			{Description: "add login page", SHA: "abc1234567"},
		},
		BugFixes: []parser.ParsedCommit{
			{Description: "fix nil pointer", Scope: "api", SHA: "def5678901"},
		},
	}

	out := ToMarkdown(cl, "https://github.com/owner/repo")

	if !strings.Contains(out, "### Features") {
		t.Error("expected Features section")
	}
	if !strings.Contains(out, "### Bug Fixes") {
		t.Error("expected Bug Fixes section")
	}
	if !strings.Contains(out, "add login page") {
		t.Error("expected feature description")
	}
	if !strings.Contains(out, "**api:**") {
		t.Error("expected scope in bug fix")
	}
	if !strings.Contains(out, "abc1234") {
		t.Error("expected short SHA")
	}
	if !strings.Contains(out, "https://github.com/owner/repo/commit/abc1234567") {
		t.Error("expected commit link")
	}
}

func TestToMarkdown_empty_sections_omitted(t *testing.T) {
	cl := parser.ParsedChangelog{
		Features: []parser.ParsedCommit{
			{Description: "add something", SHA: "abc1234567"},
		},
	}

	out := ToMarkdown(cl, "")

	if strings.Contains(out, "### Bug Fixes") {
		t.Error("empty Bug Fixes section should not appear")
	}
	if strings.Contains(out, "### Chores") {
		t.Error("empty Chores section should not appear")
	}
}

func TestToMarkdown_breaking_changes(t *testing.T) {
	cl := parser.ParsedChangelog{
		Breaking: []parser.ParsedCommit{
			{Description: "remove v1 endpoints", SHA: "bbb1234567"},
		},
	}

	out := ToMarkdown(cl, "")

	if !strings.Contains(out, "### Breaking Changes") {
		t.Error("expected Breaking Changes section")
	}
}

func TestToPlainText_sections(t *testing.T) {
	cl := parser.ParsedChangelog{
		Features: []parser.ParsedCommit{
			{Description: "add login page", SHA: "abc1234567"},
		},
		BugFixes: []parser.ParsedCommit{
			{Description: "fix nil pointer", Scope: "api", SHA: "def5678901"},
		},
	}

	out := ToPlainText(cl)

	if strings.Contains(out, "###") {
		t.Error("plain text should not contain markdown headers")
	}
	if strings.Contains(out, "**") {
		t.Error("plain text should not contain markdown bold")
	}
	if strings.Contains(out, "http") {
		t.Error("plain text should not contain links")
	}
	if !strings.Contains(out, "Features") {
		t.Error("expected Features section")
	}
	if !strings.Contains(out, "Bug Fixes") {
		t.Error("expected Bug Fixes section")
	}
	if !strings.Contains(out, "add login page") {
		t.Error("expected feature description")
	}
	if !strings.Contains(out, "api: fix nil pointer") {
		t.Error("expected scope without bold markers")
	}
	if !strings.Contains(out, "abc1234") {
		t.Error("expected short SHA")
	}
}

func TestToPlainText_empty_sections_omitted(t *testing.T) {
	cl := parser.ParsedChangelog{
		Features: []parser.ParsedCommit{
			{Description: "add something", SHA: "abc1234567"},
		},
	}

	out := ToPlainText(cl)

	if strings.Contains(out, "Bug Fixes") {
		t.Error("empty Bug Fixes section should not appear")
	}
}

func TestToMarkdown_no_repo_url(t *testing.T) {
	cl := parser.ParsedChangelog{
		Features: []parser.ParsedCommit{
			{Description: "add something", SHA: "abc1234567"},
		},
	}

	out := ToMarkdown(cl, "")

	if strings.Contains(out, "http") {
		t.Error("expected no links when repoURL is empty")
	}
	if !strings.Contains(out, "abc1234") {
		t.Error("expected short SHA without link")
	}
}
