package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/db"
	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/diary"
	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/router"
	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	qdb, err := sql.Open("mysql", "root:root@tcp(192.168.0.12:3306)/quotes_db")
	if err != nil {
		panic(err)
	}

	quoteDB := db.NewQuoteDB(qdb)
	quoteService := service.NewQuoteService(quoteDB)
	cron := diary.NewDiaryQuote(quoteService)

 	cron.GenerateDiaryQuote()
	routes := router.Router(quoteService)
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", routes)

}
