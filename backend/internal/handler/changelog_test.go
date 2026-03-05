package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/api/changelog", Changelog)
	return r
}

func TestChangelog_MissingRepoURL(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/api/changelog", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestChangelog_InvalidRepoURL(t *testing.T) {
	r := setupRouter()

	body := `{"repo_url": "not-a-valid-url"}`
	req := httptest.NewRequest(http.MethodPost, "/api/changelog", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}
