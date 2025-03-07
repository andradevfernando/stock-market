package Interface

import (
	models "api-stock-market/src/app/StockMarket.Domain/Models"
	"time"
)

type IStockMarketRepository interface {
	Create(user []*models.StockMarketModel) error
	GetList(
		page int, limit *int,
		startDate, endDate *time.Time,
		companyName string,
	) ([]*models.StockMarketModel, error)
}
