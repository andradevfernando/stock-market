package ApplicationServiceInterface

import (
	models "api-stock-market/src/app/StockMarket.Domain/Models"
	"time"
)

type IStockMarketApplicationService interface {
	GetStockMarketList(page int, limit *int,
		startDate, endDate *time.Time,
		companyName string) ([]*models.StockMarketModel, error)
	FetchAndSaveStocks() error
}
