package db

import (
	"database/sql"

	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/client"
	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/entity"
)

type QuoteDB struct {
	database *sql.DB
}

func NewQuoteDB(db *sql.DB) *QuoteDB {
	return &QuoteDB{
		database: db,
	}
} 

func (qdb *QuoteDB) SaveDiaryQuote() ( error) {
	quote, err := client.GetRandomQuote()
	if err != nil {
		return err 
	}
	_, err = qdb.database.Exec("INSERT INTO QUOTE(quote_text, author) VALUES (?, ?)", quote.Quote, quote.Author)
	if err != nil {
		return err 
	}
	return nil 
}

func (qdb *QuoteDB) GetDiaryQuote() (*entity.Quote, error) {
	var quote entity.Quote
	err := qdb.database.QueryRow("SELECT quote_text, author FROM QUOTE WHERE DATE(saved_at) = CURDATE()").Scan(&quote.Quote, &quote.Author)
	if err != nil {
		return nil, err 
	}
	return &quote, nil 
}