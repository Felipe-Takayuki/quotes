package entity

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type QuoteKenyeAPI struct {
	GRQ func() (*Quote, error)
}

func (q *QuoteKenyeAPI) GetRandomQuote() (*Quote, error) {
	return q.GRQ()
}
