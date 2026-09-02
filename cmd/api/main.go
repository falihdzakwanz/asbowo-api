package main

import (
	"asbowo-api/internal/handler"
	"asbowo-api/internal/repository"
	"asbowo-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewMemoryQuoteRepository()

	svc := service.NewQuoteService(repo)

	h := handler.NewQuoteHandler(svc)

	r := gin.Default()

	r.GET("/api/v1/quotes/daily", h.GetDaily)

	r.Run(":8080")
}