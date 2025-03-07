package Models

import (
	"api-stock-market/src/app/StockMarket.Api/Response/Payload"
	"time"
)

type StockMarketModel struct {
	ID         int       `json:"id"`
	Ticker     string    `json:"ticker" binding:"required" example:"AAPL"`
	TargetFrom string    `json:"target_from" example:"$69.00"`
	TargetTo   string    `json:"target_to" example:"$74.00"`
	Company    string    `json:"company" example:"Apple Inc."`
	Action     string    `json:"action" example:"BUY"`
	Brokerage  string    `json:"brokerage" example:"NASDAQ"`
	RatingFrom string    `json:"rating_from" example:"4.5"`
	RatingTo   string    `json:"rating_to" example:"5.0"`
	Time       time.Time `json:"time" example:"2020-01-01T00:00:00Z"`
}

func (s StockMarketModel) ToResponse() Payload.StockResponse {

	return Payload.StockResponse{
		Ticker:     s.Ticker,
		TargetFrom: s.TargetFrom,
		TargetTo:   s.TargetTo,
		Company:    s.Company,
		Action:     s.Action,
		Brokerage:  s.Brokerage,
		RatingFrom: s.RatingFrom,
		RatingTo:   s.RatingTo,
		Time:       s.Time.UTC().Format("2006-01-02T15:04:05Z"),
	}
}

func MapToResponse(s []*StockMarketModel) []Payload.StockResponse {
	modelsSlice := make([]Payload.StockResponse, 0, len(s))
	for _, item := range s {
		var modelPtr Payload.StockResponse
		modelPtr = item.ToResponse()
		modelsSlice = append(modelsSlice, modelPtr)
	}
	return modelsSlice
}
