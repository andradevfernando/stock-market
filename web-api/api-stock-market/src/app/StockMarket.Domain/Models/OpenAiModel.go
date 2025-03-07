package Models

import (
	"api-stock-market/src/app/StockMarket.Api/Response/Payload"
)

type InvestmentRecommendation struct {
	ShouldInvest   bool    `json:"shouldInvest"`
	Reason         string  `json:"reason"`
	Recommendation string  `json:"recommendation"`
	Comment        string  `json:"comment"`
	Score          int     `json:"score"`
	Ticker         string  `json:"ticker"`
	PERatio        float64 `json:"pe_ratio"`
	PBRatio        float64 `json:"pb_ratio"`
	DividendYield  float64 `json:"dividend_yield"`
	Price          float64 `json:"price"`
	DebtToEquity   float64 `json:"debt_to_equity"`
	MarketCap      float64 `json:"market_cap"`
}

func (rec *InvestmentRecommendation) ToResponse() Payload.StockAnalysisResponse {
	return Payload.StockAnalysisResponse{
		ShouldInvest:   rec.ShouldInvest,
		Reason:         rec.Reason,
		Recommendation: rec.Recommendation,
		Comment:        rec.Comment,
		Score:          rec.Score,
		Ticker:         rec.Ticker,
		PERatio:        rec.PERatio,
		PBRatio:        rec.PBRatio,
		DividendYield:  rec.DividendYield,
		Price:          rec.Price,
		DebtToEquity:   rec.DebtToEquity,
		MarketCap:      rec.MarketCap,
	}
}
