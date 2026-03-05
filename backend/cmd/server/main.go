package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vidya381/gitchangelog/internal/handler"
)

func main() {
	r := gin.Default()

	r.GET("/health", handler.Health)
	r.POST("/api/changelog", handler.Changelog)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
