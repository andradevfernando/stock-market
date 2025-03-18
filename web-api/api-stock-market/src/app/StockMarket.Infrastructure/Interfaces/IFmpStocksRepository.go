package Interface

import (
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Http/Response"
)

type IFmpStocksRepository interface {
	GetRealtimeData(ticker string) (realTimeData Response.RealTimeDataResponse, err error)
	GetHistoricalFullPrice(ticker string) (Response.HistoricalFullPriceResponse, error)
	GetKeyMetrics(ticker string) (metrics Response.MetricsResponse, err error)
}
