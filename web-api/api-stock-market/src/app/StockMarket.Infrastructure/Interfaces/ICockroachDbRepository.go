package Interface

import (
	models "api-stock-market/src/app/StockMarket.Domain/Models"
	"time"
)

type ICockroachDbRepository interface {
	CreateStocks(user []*models.StockMarketModel) error
	GetStockList(
		page int, limit *int,
		startDate, endDate *time.Time,
		companyName string,
	) ([]*models.StockMarketModel, error)
}
