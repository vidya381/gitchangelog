package github

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/gofri/go-github-pagination/githubpagination"
	"github.com/google/go-github/v72/github"
)

// NewClient returns a GitHub API client. Pass a personal access token to
// authenticate — this raises the rate limit and allows access to private repos.
// Pass an empty string to use the API without authentication (60 req/hour limit).
func NewClient(token string) *github.Client {
	paginator := githubpagination.NewClient(nil,
		githubpagination.WithPerPage(100),
	)
	if token != "" {
		return github.NewClient(paginator).WithAuthToken(token)
	}
	return github.NewClient(paginator)
}

// ParseRepoURL takes a full GitHub URL like https://github.com/owner/repo
// and returns the owner and repo name separately.
func ParseRepoURL(repoURL string) (owner, repo string, err error) {
	u, err := url.Parse(repoURL)
	if err != nil {
		return "", "", fmt.Errorf("invalid URL: %w", err)
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("URL must be in the format https://github.com/owner/repo")
	}
	return parts[0], parts[1], nil
}

// FetchCommits returns all commits in the repository up to the given ref (tag or SHA).
// If to is empty, it starts from HEAD. Pagination is handled automatically.
func FetchCommits(ctx context.Context, client *github.Client, owner, repo, to string) ([]*github.RepositoryCommit, error) {
	opts := &github.CommitsListOptions{
		SHA: to,
	}
	commits, _, err := client.Repositories.ListCommits(ctx, owner, repo, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch commits: %w", err)
	}
	return commits, nil
}
