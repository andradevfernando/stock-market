package Response

import (
	. "api-stock-market/src/app/StockMarket.Domain/Models"
	"time"
)

type StockItemResponse struct {
	Ticker     string    `json:"ticker"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Company    string    `json:"company"`
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	Time       time.Time `json:"time"`
}

type StockListResponse struct {
	Items    []StockItemResponse `json:"items"`
	NextPage string              `json:"next_page"`
}

func (s StockItemResponse) ToStockMarketModel() StockMarketModel {
	return StockMarketModel{
		Ticker:     s.Ticker,
		TargetFrom: s.TargetFrom,
		TargetTo:   s.TargetTo,
		Company:    s.Company,
		Action:     s.Action,
		Brokerage:  s.Brokerage,
		RatingFrom: s.RatingFrom,
		RatingTo:   s.RatingTo,
		Time:       s.Time,
	}
}

func MapStockItemsToModels(responses []StockItemResponse) []StockMarketModel {
	modelsSlice := make([]StockMarketModel, 0, len(responses))
	for _, item := range responses {
		model := item.ToStockMarketModel()
		modelsSlice = append(modelsSlice, model)
	}
	return modelsSlice
}
func MapStockItemsToModelsPtr(responses []StockItemResponse) []*StockMarketModel {
	modelsSlice := make([]*StockMarketModel, 0, len(responses))
	for _, item := range responses {
		modelPtr := new(StockMarketModel)
		*modelPtr = item.ToStockMarketModel()
		modelsSlice = append(modelsSlice, modelPtr)
	}
	return modelsSlice
}
