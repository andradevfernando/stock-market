package Payload

type StockResponse struct {
	Ticker     string `json:"ticker" binding:"required" example:"AAPL"`
	TargetFrom string `json:"target_from" example:"$69.00"`
	TargetTo   string `json:"target_to" example:"$74.00"`
	Company    string `json:"company" example:"Apple Inc."`
	Action     string `json:"action" example:"BUY"`
	Brokerage  string `json:"brokerage" example:"NASDAQ"`
	RatingFrom string `json:"rating_from" example:"4.5"`
	RatingTo   string `json:"rating_to" example:"5.0"`
	Time       string `json:"time" example:"2020-01-01T00:00:00Z"`
}
