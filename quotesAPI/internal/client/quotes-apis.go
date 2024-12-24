package client

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/entity"
)

func GetRandomQuote() (*entity.Quote, error) {
	url := "https://zenquotes.io/api/random"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var quoteResp []*entity.QuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, err 
	}

	var quote entity.Quote
	quote.Quote = quoteResp[0].Quote
	quote.Author = quoteResp[0].Author



	return &quote, nil
}
