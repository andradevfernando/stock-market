package Services

import (
	models "api-stock-market/src/app/StockMarket.Domain/Models"
	repo "api-stock-market/src/app/StockMarket.Infrastructure/Interfaces"
	response "api-stock-market/src/app/StockMarket.Infrastructure/Repository/Http/Response"
	"encoding/json"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

type StockMarketApplicationService struct {
	dbRepository               repo.ICockroachDbRepository
	truoraStocksHttpRepository repo.ITruoraStocksHttpRepository
}

func NewStockMarketApplicationService(
	stockRepo repo.ICockroachDbRepository,
	httpRepo repo.ITruoraStocksHttpRepository,
) *StockMarketApplicationService {

	return &StockMarketApplicationService{
		dbRepository:               stockRepo,
		truoraStocksHttpRepository: httpRepo,
	}
}

func (s StockMarketApplicationService) GetStockMarketList(page int, limit *int,
	startDate, endDate *time.Time,
	companyName string) ([]*models.StockMarketModel, error) {

	var stockMarketList, err = s.dbRepository.GetStockList(page, limit, startDate, endDate, companyName)

	if err != nil {
		return stockMarketList, err
	}
	return stockMarketList, nil
}
func (s StockMarketApplicationService) FetchAndSaveStocks() error {

	viper.SetConfigName("appsettings")
	viper.SetConfigType("json")
	viper.AddConfigPath("src/app/StockMarket.Api/")

	if err := viper.ReadInConfig(); err != nil {
		panic("Error loading configuration file: %w, err\n")
	}

	urlDatabase := viper.GetString("ConnectionStrings.CockroachDB")

	uriPostgres := strings.Replace(urlDatabase, "postgresql://", "cockroachdb://", 1)

	m, err := migrate.New(
		"file://src/app/StockMarket.Infrastructure/Migrations",
		uriPostgres,
	)
	if err != nil {
		log.Fatalf("Error deleting migration: %v", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error executing migrations: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error executing migrations: %v", err)
	}
	log.Println("Migrations executed successfully!")

	var nextPage = ""
	var stockList []*models.StockMarketModel
	n := 0

	for n < 1 {
		var body, err = s.truoraStocksHttpRepository.FetchStocks(nextPage)

		if err == nil {
			var stockMarketResponse response.StockListResponse

			var errorResponse struct {
				Message string `json:"message"`
			}
			errUnmarshal := json.Unmarshal(body, &errorResponse)
			if errUnmarshal == nil && errorResponse.Message == "Service Unavailable" {
				log.Println("API returned Service Unavailable'. Repeting request...")
				time.Sleep(1 * time.Second)
				continue
			}

			err := json.Unmarshal(body, &stockMarketResponse)
			if err != nil {
				return err
			}
			fmt.Println(stockMarketResponse)
			stockList = append(stockList, response.MapStockItemsToModelsPtr(stockMarketResponse.Items)...)
			if stockMarketResponse.NextPage == "" {
				n = 1
			}
			nextPage = stockMarketResponse.NextPage
		}
	}

	return s.dbRepository.CreateStocks(stockList)
}
