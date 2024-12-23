package client

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/entity"
)

func GetRandomQuote() (*entity.Quote, error) {
	url := "https://api.kanye.rest/"
	response, err := http.Get(url)
	if err != nil {
		return nil, err 
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err 
	}
	var quote *entity.Quote
	json.Unmarshal(body, &quote)
	quote.Author = "Kenye West"

	return quote, nil 
}


