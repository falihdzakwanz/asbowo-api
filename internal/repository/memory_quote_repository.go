package repository

import (
	"errors"
	"time"
	"asbowo-api/internal/model"
)

type MemoryQuoteRepository struct {
	quotes []model.Quote
}

func NewMemoryQuoteRepository() *MemoryQuoteRepository {
	now := time.Now()
	return &MemoryQuoteRepository{
		quotes: []model.Quote{
			{ID: 1, Text: "Ndhasmu etik!", CreatedAt: now},
			{ID: 2, Text: "Masak air biar mateng!", CreatedAt: now},
			{ID: 3, Text: "Omon-omon saja.", CreatedAt: now},
		},
	}
}

func (m *MemoryQuoteRepository) Count() (int, error) {
	return len(m.quotes), nil
}

func (m *MemoryQuoteRepository) GetByOffset(offset int) (model.Quote, error) {
	if offset < 0 || offset >= len(m.quotes) {
		return model.Quote{}, errors.New("offset out of range")
	}
	return m.quotes[offset], nil
}

func (m *MemoryQuoteRepository) GetAll() ([]model.Quote, error) {
	return m.quotes, nil
}

func (m *MemoryQuoteRepository) Create(quote model.Quote) error {
	m.quotes = append(m.quotes, quote)
	return nil
}