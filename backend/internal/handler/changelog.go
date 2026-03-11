package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidya381/gitchangelog/internal/formatter"
	githubclient "github.com/vidya381/gitchangelog/internal/github"
	"github.com/vidya381/gitchangelog/internal/parser"
)

type changelogRequest struct {
	RepoURL string `json:"repo_url" binding:"required"`
	Token   string `json:"token"`
	FromTag string `json:"from_tag"`
	ToTag   string `json:"to_tag"`
	Format  string `json:"format"`
}

type changelogResponse struct {
	Changelog   string `json:"changelog"`
	CommitCount int    `json:"commit_count"`
	From        string `json:"from"`
	To          string `json:"to"`
}

func Changelog(c *gin.Context) {
	var req changelogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repo_url is required"})
		return
	}

	owner, repo, err := githubclient.ParseRepoURL(req.RepoURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := githubclient.NewClient(req.Token)
	commits, err := githubclient.FetchCommits(context.Background(), client, owner, repo, req.ToTag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	parsed := parser.ParseCommits(commits)
	changelog := formatter.ToMarkdown(parsed, req.RepoURL)

	c.JSON(http.StatusOK, changelogResponse{
		Changelog:   changelog,
		CommitCount: len(commits),
		From:        req.FromTag,
		To:          req.ToTag,
	})
}
