package diary

import (
	"github.com/Felipe-Takayuki/quotes/quotesAPI/internal/service"
	"github.com/robfig/cron/v3"
)

type DiaryQuote struct {
	quoteService *service.QuoteService
}

func NewDiaryQuote(qts *service.QuoteService) *DiaryQuote {
	return &DiaryQuote{
		quoteService: qts,
	}
}

func (dq *DiaryQuote) GenerateDiaryQuote() {
	c := cron.New()
	id, _ := c.AddFunc("@every 24h", dq.quoteService.SaveDiaryQuote)

	c.Entry(id).Job.Run()

	go c.Start()
}
