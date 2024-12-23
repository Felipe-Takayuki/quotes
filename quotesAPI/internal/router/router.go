package router

import (
	"net/http"

	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/server"
	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Router(qtservice *service.QuoteService) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	corsConfig := cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	r.Use(cors.Handler(corsConfig))
	r.Get("/quote", server.NewQuoteServer(qtservice).GetDiaryQuote)
	return r
}
