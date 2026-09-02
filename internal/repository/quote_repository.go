package repository

import (
	"asbowo-api/internal/model"
)

type QuoteRepository interface {
	Count() (int, error)
	GetByOffset(offset int) (model.Quote, error)
	GetAll() ([]model.Quote, error)
	Create(quote model.Quote) error
}