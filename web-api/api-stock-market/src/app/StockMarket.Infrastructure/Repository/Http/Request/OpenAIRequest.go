package Request

type OpenAIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Content struct {
	Ticker        string
	DebtToEquity  float64
	MarketCap     float64
	ROE           float64
	DividendYield float64
	PBRatio       float64
	PERatio       float64
	MovingAverage float64
	Price         float64
	Volume        int
}
