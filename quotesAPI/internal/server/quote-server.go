package server

import (
	"encoding/json"
	"net/http"

	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/service"
)

type QuoteServer struct {
	quoteService *service.QuoteService
}

func NewQuoteServer(qtService *service.QuoteService) *QuoteServer {
	return &QuoteServer{
		quoteService: qtService,
	}
}

func (qws *QuoteServer) GetDiaryQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	quote, err := qws.quoteService.GetDiaryQuote()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(quote) 
}
