package Http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TruoraStocksHttpRepository struct {
	TruoraBaseUrl string
	ApiKey        string
}

func NewTruoraStocksHttpRepository(apiKey, baseUrl string) *TruoraStocksHttpRepository {
	return &TruoraStocksHttpRepository{
		TruoraBaseUrl: baseUrl,
		ApiKey:        apiKey,
	}
}

func (h *TruoraStocksHttpRepository) FetchStocks(nextPage string) ([]byte, error) {

	params := url.Values{}
	params.Add("next_page", nextPage)
	h.TruoraBaseUrl = h.TruoraBaseUrl + "?" + params.Encode()
	req, err := http.NewRequest("GET", h.TruoraBaseUrl, nil)
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
