package Services

import (
	"api-stock-market/src/app/StockMarket.Domain/Models"
	Interface "api-stock-market/src/app/StockMarket.Infrastructure/Interfaces"
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Http/Request"
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Http/Response"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type StockAnalysisApplicationService struct {
	fmpStocksRepository Interface.IFmpStocksRepository
	openAIRepository    Interface.IOpenAiHttpRepository
	dbRepository        Interface.ICockroachDbRepository
}

func NewStockAnalysisApplicationService(externalStockRepo Interface.IFmpStocksRepository, openAiRepo Interface.IOpenAiHttpRepository, stockRepository Interface.ICockroachDbRepository) *StockAnalysisApplicationService {
	return &StockAnalysisApplicationService{externalStockRepo, openAiRepo, stockRepository}
}

func (s *StockAnalysisApplicationService) GetInvestmentRecommendation(ticker string) (Models.InvestmentRecommendation, error) {

	var wg sync.WaitGroup

	var metrics Response.MetricsResponse
	var data Response.RealTimeDataResponse
	var historicalPrice Response.HistoricalFullPriceResponse

	wg.Add(3)

	go func() {
		defer wg.Done()
		var err error
		metrics, err = s.fmpStocksRepository.GetKeyMetrics(ticker)
		if err != nil {
			metrics = Response.MetricsResponse{}
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		data, err = s.fmpStocksRepository.GetRealtimeData(ticker)
		if err != nil {
			data = Response.RealTimeDataResponse{}
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		historicalPrice, err = s.fmpStocksRepository.GetHistoricalFullPrice(ticker)
		if err != nil {
			historicalPrice = Response.HistoricalFullPriceResponse{}
		}
	}()
	wg.Wait()
	var movingAverage = calculateMovingAverage(historicalPrice)

	var content = Request.Content{
		Ticker:        ticker,
		DebtToEquity:  metrics.DebtToEquity,
		MarketCap:     metrics.MarketCap,
		ROE:           metrics.ROE,
		DividendYield: metrics.DividendYield,
		PBRatio:       metrics.PBRatio,
		PERatio:       metrics.PERatio,
		MovingAverage: movingAverage,
		Price:         data.Price,
		Volume:        data.Volume,
	}

	return s.openAIRepository.GetInvestmentRecommendation(content)
}

func (s *StockAnalysisApplicationService) GetBestInvestments() ([]*Models.StockMarketModel, error) {

	stocks, err := s.dbRepository.GetStockList(1, nil, nil, nil, "")
	if err != nil {
		return nil, err
	}
	if len(stocks) == 0 {
		return nil, fmt.Errorf("stocks not found")
	}

	type Recommendation struct {
		Stock     *Models.StockMarketModel
		Potential float64
	}
	var recs []Recommendation

	parsePrice := func(priceStr string) (float64, error) {
		priceStr = strings.TrimSpace(priceStr)
		priceStr = strings.TrimPrefix(priceStr, "$")
		return strconv.ParseFloat(priceStr, 64)
	}

	for _, stock := range stocks {
		targetFrom, err := parsePrice(stock.TargetFrom)
		if err != nil {
			continue
		}
		targetTo, err := parsePrice(stock.TargetTo)
		if err != nil {
			continue
		}
		targetPotential := targetTo - targetFrom

		combinedPotential := targetPotential

		recs = append(recs, Recommendation{
			Stock:     stock,
			Potential: combinedPotential,
		})
	}

	sort.Slice(recs, func(i, j int) bool {
		return recs[i].Potential > recs[j].Potential
	})

	limit := 5
	if len(recs) < limit {
		limit = len(recs)
	}
	bestRecs := recs[:limit]

	var responses []*Models.StockMarketModel
	for _, rec := range bestRecs {
		responses = append(responses, rec.Stock)
	}

	return responses, nil
}

func calculateMovingAverage(data Response.HistoricalFullPriceResponse) float64 {
	const period = 365
	var sum float64
	for i := 0; i < period && i < len(data.Historical); i++ {
		sum += data.Historical[i].Close
	}
	return sum / float64(period)
}
