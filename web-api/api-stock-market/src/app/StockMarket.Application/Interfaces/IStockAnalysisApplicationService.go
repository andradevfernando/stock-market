package ApplicationServiceInterface

import "api-stock-market/src/app/StockMarket.Domain/Models"

type IStockAnalysisApplicationService interface {
	GetInvestmentRecommendation(ticker string) (Models.InvestmentRecommendation, error)
	GetBestInvestments() ([]*Models.StockMarketModel, error)
}
