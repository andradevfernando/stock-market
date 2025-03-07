package Services

import (
	"api-stock-market/src/app/StockMarket.Domain/Models"
	Interface "api-stock-market/src/app/StockMarket.Infrastructure/Interfaces"
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Request"
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Response"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type StockAnalysisApplicationService struct {
	ExternalStocksRepository Interface.IExternalStocksRepository
	OpenAIRepository         Interface.IOpenApiRepository
	StockMarketRepository    Interface.IStockMarketRepository
}

func NewStockAnalysisApplicationService(externalStockRepo Interface.IExternalStocksRepository, openAiRepo Interface.IOpenApiRepository, stockRepository Interface.IStockMarketRepository) *StockAnalysisApplicationService {
	return &StockAnalysisApplicationService{externalStockRepo, openAiRepo, stockRepository}
}

func (s *StockAnalysisApplicationService) GetInvestmentRecommendation(ticker string) (Models.InvestmentRecommendation, error) {
	metrics, err := s.ExternalStocksRepository.GetKeyMetrics(ticker)
	if err != nil {
		metrics = Response.MetricsResponse{}
	}
	data, err := s.ExternalStocksRepository.GetRealtimeData(ticker)

	if err != nil {
		data = Response.RealTimeDataResponse{}
	}

	historicalPrice, err := s.ExternalStocksRepository.GetHistoricalFullPrice(ticker)

	if err != nil {
		historicalPrice = Response.HistoricalFullPriceResponse{}
	}

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

	return s.OpenAIRepository.GetInvestmentRecommendation(content)
}

func (s *StockAnalysisApplicationService) GetBestInvestments() ([]*Models.StockMarketModel, error) {

	stocks, err := s.StockMarketRepository.GetList(1, nil, nil, nil, "")
	if err != nil {
		return nil, err
	}
	if len(stocks) == 0 {
		return nil, fmt.Errorf("nenhuma ação encontrada")
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
