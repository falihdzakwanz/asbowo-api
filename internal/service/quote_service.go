package service

import (
	"errors"
	"math/rand/v2"
	"time"
	"asbowo-api/internal/model"
	"asbowo-api/internal/repository"
)

type QuoteService struct {
	repo repository.QuoteRepository
}

func NewQuoteService(repo repository.QuoteRepository) *QuoteService {
	return &QuoteService{repo: repo}
}

func (s *QuoteService) GetDailyQuote(now time.Time) (model.Quote, error) {
	total, err := s.repo.Count()

	if err != nil {
		return model.Quote{}, err
	}
	
	if total == 0 {
		return model.Quote{}, errors.New("belum ada asbun yang tersimpan")
	}

	dateSeed := int64(now.Year()*10000 + int(now.Month())*100 + now.Day())
	r := rand.New(rand.NewPCG(uint64(dateSeed), uint64(dateSeed)))
	offset := r.IntN(total)

	return s.repo.GetByOffset(offset)
}