package Repository

import (
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Response"
	"encoding/json"
	"fmt"
	"net/http"
)

type ExternalStocksRepository struct {
	fmpBaseURL string
	fmpAPIKey  string
}

func NewExternalStocksRepository(apiKey string) *ExternalStocksRepository {
	return &ExternalStocksRepository{
		fmpBaseURL: "https://financialmodelingprep.com/api/v3",
		fmpAPIKey:  apiKey,
	}
}

func (e ExternalStocksRepository) GetRealtimeData(ticker string) (realTimeData Response.RealTimeDataResponse, err error) {
	url := fmt.Sprintf("%s/stock/real-time-price/%s?apikey=%s",
		e.fmpBaseURL, ticker, e.fmpAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return Response.RealTimeDataResponse{}, err
	}
	defer resp.Body.Close()

	var data Response.RealTimeDataResponseList
	json.NewDecoder(resp.Body).Decode(&data)

	if len(data.RealTimeDataResponse) > 0 {
		return data.RealTimeDataResponse[0], nil
	}

	return realTimeData, nil
}
func (e ExternalStocksRepository) GetKeyMetrics(ticker string) (metrics Response.MetricsResponse, err error) {
	url := fmt.Sprintf("%s/key-metrics/%s?apikey=%s",
		e.fmpBaseURL, ticker, e.fmpAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return metrics, err
	}
	defer resp.Body.Close()

	var data []Response.MetricsResponse
	json.NewDecoder(resp.Body).Decode(&data)

	if len(data) > 0 {
		return data[0], nil
	}
	return metrics, fmt.Errorf("no data found")
}

func (e ExternalStocksRepository) GetHistoricalFullPrice(ticker string) (Response.HistoricalFullPriceResponse, error) {

	const period int = 365
	url := fmt.Sprintf("%s/historical-price-full/%s?timeseries=%d&apikey=%s",
		e.fmpBaseURL, ticker, period, e.fmpAPIKey)

	resp, _ := http.Get(url)
	var data Response.HistoricalFullPriceResponse
	json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}
