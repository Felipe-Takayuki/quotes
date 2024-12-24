package entity

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type QuoteResponse struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

type QuoteAPI struct {
	GRQ func() (*Quote, error)
}

func (q *QuoteAPI) GetRandomQuote() (*Quote, error) {
	return q.GRQ()
}
