package Interface

import (
	"api-stock-market/src/app/StockMarket.Domain/Models"
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Http/Request"
)

type IOpenAiHttpRepository interface {
	GetInvestmentRecommendation(content Request.Content) (Models.InvestmentRecommendation, error)
}
