package main

import (
	controllers "api-stock-market/src/app/StockMarket.Api/Controllers"
	errorMiddleware "api-stock-market/src/app/StockMarket.Api/Middlewares"
	serviceImpl "api-stock-market/src/app/StockMarket.Application/Services"
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Database"
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Http"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"

	_ "api-stock-market/src/app/StockMarket.Api/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func loadConfig() (string, error) {

	viper.SetConfigName("appsettings")
	viper.SetConfigType("json")
	viper.AddConfigPath("src/app/StockMarket.Api/")

	if err := viper.ReadInConfig(); err != nil {
		return "", fmt.Errorf("error loading configuration file: %w", err)
	}

	connString := viper.GetString("ConnectionStrings.CockroachDB")
	if connString == "" {
		return "", fmt.Errorf("connection string not defined in the configuration file")
	}
	return connString, nil
}

// @title Stock Market API
// @version 1.0
// @description Stock market API.
// @host localhost:5000
// @BasePath /
func main() {

	connString, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configurations: %v", err)
	}

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Error parsing connection string: %v", err)
	}
	config.MaxConns = 10

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	defer pool.Close()

	fmt.Println("Successfully connected to CockroachDB!")

	runMigrations(connString)

	initSwagger()
	initHttp(context.Background(), pool)
}

func initHttp(ctx context.Context, pool *pgxpool.Pool) {
	fmpApiKey := os.Getenv("FMP_API_KEY")
	openAiApiKey := os.Getenv("OPENAI_API_KEY")
	truoraApiKey := os.Getenv("TRUORA_API_KEY")

	openApiRepo := Http.NewOpenAiHttpRepository(openAiApiKey, "gpt-4o-mini")
	externalStockRepo := Http.NewFmpStocksRepository(fmpApiKey)
	httpRepo := Http.NewTruoraStocksHttpRepository(truoraApiKey)
	stockRepo := Database.NewCockroachDbRepository(ctx, pool)

	var service = serviceImpl.NewStockMarketApplicationService(stockRepo, httpRepo)
	var analysisService = serviceImpl.NewStockAnalysisApplicationService(externalStockRepo, openApiRepo, stockRepo)

	//job := &Jobs.DailyJob{
	//	Service: service,
	//}

	//go job.ScheduleDailyJob()

	controller := &controllers.StockMarketController{
		StockMarketService:         service,
		StockMarketAnalysisService: analysisService,
	}

	http.HandleFunc("/stocks", controller.GetStockMarkets)
	http.HandleFunc("/fetchstocks", controller.FetchAndSaveStocks)
	http.HandleFunc("/analysis/", controller.StockAnalysis)
	http.HandleFunc("/best-investments", controller.GetBestInvestments)

	log.Println("Listening on port 5000...")
	if err := http.ListenAndServe(":5000", errorMiddleware.RecoveryMiddleware(http.DefaultServeMux)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	//select {}
}

func initSwagger() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
}
func runMigrations(databaseURL string) {

	uriPostgres := strings.Replace(databaseURL, "postgresql://", "cockroachdb://", 1)

	m, err := migrate.New(
		"file://src/app/StockMarket.Infrastructure/Migrations",
		uriPostgres,
	)
	if err != nil {
		log.Fatalf("Error creating migrations: %v", err)
	}
	// Executes all migrations (up)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error executing migrations: %v", err)
	}
	log.Println("Migrations executed successfully!")
}
