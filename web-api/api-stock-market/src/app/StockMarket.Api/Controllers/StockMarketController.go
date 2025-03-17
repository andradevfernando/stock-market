package controllers

import (
	errorResponse "api-stock-market/src/app/StockMarket.Api/Response/Errors"
	services "api-stock-market/src/app/StockMarket.Application/Interfaces"
	models "api-stock-market/src/app/StockMarket.Domain/Models"
	"encoding/json"
	"strings"

	"log"
	"net/http"
	"strconv"
	"time"
)

type (
	StockMarketController struct {
		StockMarketService         services.IStockMarketApplicationService
		StockMarketAnalysisService services.IStockAnalysisApplicationService
	}
)

// GetStockMarkets @Summary List of stocks
// @Description Returns a paginated list of stocks.
// @Tags Stocks
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param items query int false "Items per page" default(10)
// @Param startDate query string false "Date in format RFC3339" default(2024-03-05T15:04:05Z)
// @Param endDate query string false "Date in format RFC3339" default(2025-03-15T15:04:05Z)
// @Param company query string false "Name of company" default(Apple)
// @Success 200 {array} Payload.StockResponse "List of stocks"
// @Failure 400 {object} errorResponse.ErrorBadRequest "Bad request"
// @Failure 404 {object} errorResponse.ErrorNotFound "Stocks not found"
// @Failure 500 {object} errorResponse.ErrorInternal "Internal Server Error"
// @Router /stocks [get]
func (c *StockMarketController) GetStockMarkets(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	pageStr := query.Get("page")
	items := query.Get("items")
	startDateStr := query.Get("startDate")
	endDateStr := query.Get("endDate")
	company := query.Get("company")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	page := 1
	limit := 10
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			log.Printf("Page must be greater than 0: %v", err)
			badRequestError := errorResponse.ErrorNotFound{
				Code:    http.StatusBadRequest,
				Error:   "Bad request",
				Message: "Page must be greater than 0",
			}
			w.WriteHeader(http.StatusBadRequest)
			if err := encoder.Encode(badRequestError); err != nil {
				panic(err)
			}
			return
		}
	}

	if items != "" {
		limit, err = strconv.Atoi(items)
		if err != nil || limit < 1 {
			log.Printf("Item quantity must be greater than 0: %v", err)
			badRequestError := errorResponse.ErrorNotFound{
				Code:    http.StatusBadRequest,
				Error:   "Bad request",
				Message: "Item quantity must be greater than 0",
			}
			w.WriteHeader(http.StatusBadRequest)
			if err := encoder.Encode(badRequestError); err != nil {
				panic(err)
			}
			return
		}
	}

	var startDate, endDate *time.Time
	if startDateStr != "" || endDateStr != "" {

		if startDateStr == "" || endDateStr == "" {
			log.Printf("Must send both dates: %v", err)
			badRequestError := errorResponse.ErrorNotFound{
				Code:    http.StatusBadRequest,
				Error:   "Bad request",
				Message: "Must send both dates",
			}
			w.WriteHeader(http.StatusBadRequest)
			if err := encoder.Encode(badRequestError); err != nil {
				panic(err)
			}
			return
		}

		parsedStart, err := time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			log.Printf("startDate requestData format invalid (use RFC3339): %v", err)
			badRequestError := errorResponse.ErrorNotFound{
				Code:    http.StatusBadRequest,
				Error:   "Bad request",
				Message: "startDate requestData format invalid (use RFC3339)",
			}
			w.WriteHeader(http.StatusBadRequest)
			if err := encoder.Encode(badRequestError); err != nil {
				panic(err)
			}
			return
		}

		parsedEnd, err := time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			log.Printf("endDate requestData format invalid (use RFC3339): %v", err)
			badRequestError := errorResponse.ErrorNotFound{
				Code:    http.StatusBadRequest,
				Error:   "Bad request",
				Message: "endDate requestData format invalid (use RFC3339)",
			}
			w.WriteHeader(http.StatusBadRequest)
			if err := encoder.Encode(badRequestError); err != nil {
				panic(err)
			}
			return
		}

		startDate = &parsedStart
		endDate = &parsedEnd
	}

	stockMarkets, err := c.StockMarketService.GetStockMarketList(
		page,
		&limit,
		toMidnightUTC(startDate),
		toMidnightUTC(endDate),
		company,
	)
	if err != nil {
		log.Printf("Error obtaining stock requestData: %v", err)
		notFoundError := errorResponse.ErrorNotFound{
			Code:    http.StatusNotFound,
			Error:   "Not found",
			Message: "Stock markets not found in database",
		}
		w.WriteHeader(http.StatusNotFound)
		if err := encoder.Encode(notFoundError); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := encoder.Encode(models.MapToResponse(stockMarkets)); err != nil {
		panic(err)
	}
}

// FetchAndSaveStocks
// @Router /fetchstocks [post]
func (c *StockMarketController) FetchAndSaveStocks(w http.ResponseWriter, r *http.Request) {
	err := c.StockMarketService.FetchAndSaveStocks()
	if err != nil {
		panic(err.Error())
	}
	w.WriteHeader(http.StatusCreated)
}
func toMidnightUTC(t *time.Time) *time.Time {
	midnight := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		0, // hour
		0, // min
		0, // sec
		0, // nsec
		time.UTC,
	)
	return &midnight
}

// StockAnalysis performs stock analysis based on the provided requestData.
//
// @Summary Stock Analysis
// @Description Performs a stock analysis based on the provided requestData
// @Tags StockAnalysis
// @Accept json
// @Produce json
// @Param        ticker  path      string  true  "Ticker stock symbol"
// @Success 200 {object} Payload.StockAnalysisResponse
// @Failure 400 {object} errorResponse.ErrorBadRequest
// @Failure 500 {object} errorResponse.ErrorInternal
// @Router /analysis/{ticker} [post]
func (c *StockMarketController) StockAnalysis(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	ticker := strings.TrimPrefix(r.URL.Path, "/analysis/")
	if ticker == "" {
		log.Println("Ticker not provided in the route")
		w.WriteHeader(http.StatusBadRequest)
		if err := encoder.Encode(errorResponse.ErrorBadRequest{
			Code:    http.StatusBadRequest,
			Error:   "Ticker not provided",
			Message: "'ticker' parameter is required in the route",
		}); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
		return
	}

	output, err := c.StockMarketAnalysisService.GetInvestmentRecommendation(ticker)

	if err != nil {
		log.Printf("Error obtaining AI analysis: %v", err)
		internalServerError := errorResponse.ErrorInternal{
			Code:    http.StatusInternalServerError,
			Error:   "Internal server error",
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		if err := encoder.Encode(internalServerError); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := encoder.Encode(output.ToResponse()); err != nil {
		panic(err)
	}
}

// GetBestInvestments @Summary Returns the top 5 stocks for investment
// @Description This endpoint calls the service responsible for identifying the best stocks for investment and returns the results in JSON.
// @Tags Stocks
// @Accept json
// @Produce json
// @Success 200 {array} Payload.StockResponse "Successful operation"
// @Failure 500 {object} errorResponse.ErrorInternal "Internal Server Error"
// @Router /best-investments [get]
func (c *StockMarketController) GetBestInvestments(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	bestInvestments, err := c.StockMarketAnalysisService.GetBestInvestments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := encoder.Encode(errorResponse.ErrorInternal{Code: http.StatusInternalServerError, Message: err.Error()}); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := encoder.Encode(models.MapToResponse(bestInvestments)); err != nil {
		panic(err)
	}
}
