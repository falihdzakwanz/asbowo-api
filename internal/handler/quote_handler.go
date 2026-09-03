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
type CreateQuoteRequest struct {
	Text string `json:"text" binding:"required"`
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

func (h *QuoteHandler) Create(c *gin.Context) {
	var req CreateQuoteRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON salah atau text kosong"})
		return
	}

	newQuote, err := h.svc.CreateNewQuote(req.Text)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terdapat kesalahan saat menyimpan quote baru. Silahkan coba lagi nanti"})
		return
	}

	c.JSON(http.StatusCreated, newQuote)
}

func (h *QuoteHandler) GetAll(c *gin.Context) {
	quotes, err := h.svc.GetAllQuotes()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terdapat kesalahan saat mengambil quotes. Silahkan coba lagi nanti"})
		return
	}

	c.JSON(http.StatusOK, quotes)
}