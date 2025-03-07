package Jobs

import (
	ApplicationServiceInterface "api-stock-market/src/app/StockMarket.Application/Interfaces"
	"log"
	"time"
)

type DailyJob struct {
	Service ApplicationServiceInterface.IStockMarketApplicationService
}

func (dj *DailyJob) ExecuteDailyJob() error {
	log.Println("Job executing!")
	err := dj.Service.FetchAndSaveStocks()
	if err != nil {
		return err
	}
	return nil
}

func (dj *DailyJob) ScheduleDailyJob() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	if err := dj.ExecuteDailyJob(); err != nil {
		log.Printf("Error executing job: %v", err)
	}

	for {
		select {
		case <-ticker.C:
			if err := dj.ExecuteDailyJob(); err != nil {
				log.Printf("Erro ao executar o job diário: %v", err)
			}
		}
	}

}
