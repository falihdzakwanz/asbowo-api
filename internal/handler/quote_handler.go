package handler

import (
	"net/http"
	"time"

	"asbowo-api/internal/service"
	"github.com/gin-gonic/gin"
)

type QuoteHandler struct {
	svc *service.QuoteService
}

func NewQuoteHandler(svc *service.QuoteService) *QuoteHandler {
	return &QuoteHandler{svc: svc}
}

func (h *QuoteHandler) GetDaily(c *gin.Context) {
	quote, err := h.svc.GetDailyQuote(time.Now())
	if err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, quote)
}