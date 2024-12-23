package service

import (
	"fmt"

	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/db"
	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/entity"
)

type QuoteService struct {
	quoteDB *db.QuoteDB
}

func NewQuoteService(qtDB *db.QuoteDB) *QuoteService {
	return &QuoteService{
		quoteDB: qtDB,
	}
}

func (qs *QuoteService) SaveDiaryQuote() () {
	err := qs.quoteDB.SaveDiaryQuote()
	if err != nil {
		return
	}
	fmt.Println("Diary Quote Generated!")
	return
}

func (qs *QuoteService) GetDiaryQuote() (*entity.Quote, error) {
	quote, err := qs.quoteDB.GetDiaryQuote()
	if err != nil {
		return nil, err 
	}
	return quote, nil 
}