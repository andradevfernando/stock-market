package Payload

type StockAnalysisResponse struct {
	ShouldInvest   bool    `json:"shouldInvest" example:"true"`                                                          // Exemplo: true
	Reason         string  `json:"reason" example:"Example: High PE ratio compared to the sector average"`               // Example: "Example: High PE ratio compared to the sector average"
	Recommendation string  `json:"recommendation" example:"Example: Buy this asset"`                                     // Example: "Example: Buy this asset"
	Comment        string  `json:"comment" example:"Example: The asset has potential for appreciation despite the risk"` // Example: "Example: The asset has potential for appreciation despite the risk"
	Score          int     `json:"score" example:"85"`                                                                   // Exemplo: 85
	Ticker         string  `json:"ticker" example:"AAPL"`                                                                // Exemplo: "AAPL"
	PERatio        float64 `json:"pe_ratio" example:"15.23"`                                                             // Exemplo: 15.23
	PBRatio        float64 `json:"pb_ratio" example:"3.45"`                                                              // Exemplo: 3.45
	DividendYield  float64 `json:"dividend_yield" example:"2.15"`                                                        // Exemplo: 2.15
	Price          float64 `json:"price" example:"150.75"`                                                               // Exemplo: 150.75
	DebtToEquity   float64 `json:"debt_to_equity" example:"1.20"`                                                        // Exemplo: 1.20
	MarketCap      float64 `json:"market_cap" example:"250000000000.00"`                                                 // Exemplo: 250000000000.00
}
