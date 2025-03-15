package Http

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
)

type TruoraStocksHttpRepository struct {
	ApiKey string
}

func NewTruoraStocksHttpRepository(apiKey string) *TruoraStocksHttpRepository {
	return &TruoraStocksHttpRepository{
		ApiKey: apiKey,
	}
}

func (h *TruoraStocksHttpRepository) FetchStocks(nextPage string) ([]byte, error) {
	viper.SetConfigName("appsettings")
	viper.SetConfigType("json")
	viper.AddConfigPath("src/app/StockMarket.Api/")

	urlString := viper.GetString("TruoraStocksHttpRepository.BaseAddress")
	if urlString == "" {
		panic("Truora http string not found")
	}
	params := url.Values{}
	params.Add("next_page", nextPage)
	urlString = urlString + "?" + params.Encode()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.ApiKey))
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		return body, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
	return nil, nil
}
