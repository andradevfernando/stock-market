package Response

type MetricsResponse struct {
	PERatio       float64 `json:"peRatio"`
	PBRatio       float64 `json:"pbRatio"`
	DebtToEquity  float64 `json:"debtToEquity"`
	ROE           float64 `json:"roe"`
	DividendYield float64 `json:"dividendYield"`
	MarketCap     float64 `json:"marketCap"`
}
type RealTimeDataResponseList struct {
	RealTimeDataResponse []RealTimeDataResponse `json:"companiesPriceList"`
}
type RealTimeDataResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Volume int     `json:"volume"`
}
type HistoricalFullPriceResponse struct {
	Historical []HistoricalResponse `json:"historical"`
}
type HistoricalResponse struct {
	Close float64 `json:"close"`
}
